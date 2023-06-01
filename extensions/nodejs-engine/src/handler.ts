import {Event} from "./proto/model/event_pb";
import {functionIdMap, functionMap} from "./function-registry";
import {Value} from "google-protobuf/google/protobuf/struct_pb";
import {load, read} from "./store";
import {ExecuteFunction, ResourceOperation, WatchLogicResources} from "./const";
import {Function} from './model/function'
import {ResourceRule} from "./model/resource-rule";
import Action = Event.Action;

const {VM} = require('vm2');

function locateFunction(packageName: string, name: string): Function {
    return functionMap[packageName + '/' + name]
}

async function handleFunctionExecutionCall(event: Event) {
    for (const record of event.getRecordsList()) {
        try {
            const fnFields = record.getPropertiesMap().get('function').getStructValue().getFieldsMap()
            const packageName = fnFields.get('package').getStringValue()
            const name = fnFields.get('name').getStringValue()
            const input = record.getPropertiesMap().get('input')?.getStructValue()?.toJavaScript()

            const fn = locateFunction(packageName, name)

            const output = new VM({
                sandbox: {
                    ...input
                },
            }).run(`(function () {
                ${fn.script}
            })()`, {
                timeout: 1000,
            }) ?? {ok: true};

            record.getPropertiesMap().set('output', Value.fromJavaScript(output))
            record.getPropertiesMap().set('status', Value.fromJavaScript('success'))
        } catch (e) {
            console.log(e)

            record.getPropertiesMap().set('error', Value.fromJavaScript(e.message))
            record.getPropertiesMap().set('status', Value.fromJavaScript('error'))
        }
    }

    return event;
}

function locateRulesListeningToEvent(event: Event): ResourceRule[] {
    if (event.getAction() != Action.CREATE && event.getAction() != Action.UPDATE) {
        return []
    }
    const rules = read<ResourceRule>('logic', 'ResourceRule')

    return rules.filter(rule => rule.namespace == event.getResource().getNamespace() && rule.resource == event.getResource().getName())
}

async function handleResourceOperation(event: Event): Promise<Event> {
    const rules = locateRulesListeningToEvent(event)

    for (const record of event.getRecordsList()) {
        for (const rule of rules) {
            const fn = functionIdMap[rule.conditionFunction.id]

            const entity = {}

            record.getPropertiesMap().forEach((entry, key) => {
                entity[key] = entry.toJavaScript()
            })

            console.log(fn)

            const output = new VM({
                sandbox: {
                    entity: entity
                },
            }).run(`(function () {
                ${fn.script}
            })()`, {
                timeout: 1000,
            })

            console.log('e', output)

            if (output === false) {
                throw new Error(`Rule condition failed: ${rule.name}`)
            }
        }
    }

    return event
}

export async function handle(fnName: string, event: Event): Promise<Event> {
    if (!event) {
        return event
    }

    console.log('Handling event: ' + fnName)

    switch (fnName) {
        case WatchLogicResources:
            if (event.getResource().getNamespace() != 'logic') {
                throw new Error('WatchLogicResources can only be called for logic namespace')
            }
            load('logic', event.getResource().getName())
            return
        case ExecuteFunction:
            return handleFunctionExecutionCall(event)
        case ResourceOperation:
            return handleResourceOperation(event)
        default:
            throw new Error('Unknown function name: ' + fnName)
    }
}