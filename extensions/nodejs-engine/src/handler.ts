import {Event} from "./proto/model/event_pb";
import {functionIdMap, functionMap} from "./function-registry";
import {Value} from "google-protobuf/google/protobuf/struct_pb";
import {Record} from './proto/model/record_pb'
import {load, read} from "./store";
import {
    ExecuteFunction,
    ResourceOperationRule,
    ResourceOperationTrigger,
    WatchLogicResources
} from "./const";
import {Function} from './model/function'
import {ResourceRule, ResourceRuleName} from "./model/resource-rule";
import Action = Event.Action;
import {FunctionTrigger, FunctionTriggerName} from "./model/function-trigger";

const {VM} = require('vm2');

function locateFunction(packageName: string, name: string): Function {
    return functionMap[packageName + '/' + name]
}

export async function executeFunction<R>(fn: Function, params: object): Promise<R> {
    return new VM({
        sandbox: {
            ...params
        },
    }).run(`(function () {
                ${fn.script}
            })()`, {
        timeout: 1000,
    })

}

async function handleFunctionExecutionCall(event: Event) {
    for (const record of event.getRecordsList()) {
        try {
            const fnFields = record.getPropertiesMap().get('function').getStructValue().getFieldsMap()
            const packageName = fnFields.get('package').getStringValue()
            const name = fnFields.get('name').getStringValue()
            const input = record.getPropertiesMap().get('input')?.getStructValue()?.toJavaScript()

            const fn = locateFunction(packageName, name)

            const output = (await executeFunction(fn, input) ?? {ok: true});

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
    const rules = read<ResourceRule>('logic', ResourceRuleName)

    return rules.filter(rule => rule.namespace == event.getResource().getNamespace() && rule.resource == event.getResource().getName())
}

function locateTriggersListeningToEvent(event: Event): FunctionTrigger[] {
    if (event.getAction() != Action.CREATE && event.getAction() != Action.UPDATE) {
        return []
    }
    const triggers = read<FunctionTrigger>('logic', FunctionTriggerName)

    return triggers.filter(trigger => {
        if (trigger.async == event.getSync()) {
            console.log('trigger async mismatch', trigger.async, event.getSync())
            return false
        }

        if (trigger.resource != '*' && trigger.resource != event.getResource().getName()) {
            console.log('trigger resource mismatch', trigger.resource, event.getResource().getName())
            return false
        }

        if (trigger.namespace != '*' && trigger.namespace != event.getResource().getNamespace()) {
            console.log('trigger namespace mismatch', trigger.namespace, event.getResource().getNamespace())
            return false
        }

        if (trigger.action != eventActionToString(event.getAction())) {
            console.log('trigger action mismatch', trigger.action, event.getAction().toString().toLowerCase())
            return false
        }

        return true
    })
}

export function eventActionToString(action: Event.Action) {
    switch (action) {
        case Event.Action.CREATE:
            return 'create'
        case Event.Action.UPDATE:
            return 'update'
        case Event.Action.DELETE:
            return 'delete'
        case Event.Action.LIST:
            return 'list'
        case Event.Action.GET:
            return 'get'
        default:
            return 'unknown'
    }
}

function recordToEntity(record: Record) {
    const entity = {}

    record.getPropertiesMap().forEach((entry, key) => {
        entity[key] = entry.toJavaScript()
    })
    return entity;
}

async function handleResourceOperationTrigger(event: Event): Promise<Event> {
    console.log('Handling resource operation trigger', eventActionToString(event.getAction()))

    const triggers = locateTriggersListeningToEvent(event)

    for (const record of event.getRecordsList()) {
        for (const trigger of triggers) {
            console.log('Handling trigger', trigger.name)
            const fn = functionIdMap[trigger.function.id]

            const entity = recordToEntity(record);

            let output = await executeFunction(fn, {
                entity: entity
            })

            console.log('output', output)

            if (!output) {
                output = entity

            }

            for (const property of event.getResource().getPropertiesList()) {
                if (output[property.getName()] === record.getPropertiesMap().get(property.getName())) {
                    continue
                }
                record.getPropertiesMap().set(property.getName(), Value.fromJavaScript(output[property.getName()]))
            }
        }
    }

    return event
}

async function handleResourceOperationRule(event: Event): Promise<Event> {
    console.log('Handling resource operation rule', eventActionToString(event.getAction()))

    const rules = locateRulesListeningToEvent(event)

    for (const record of event.getRecordsList()) {
        for (const rule of rules) {
            console.log('Handling rule', rule.name)
            const fn = functionIdMap[rule.conditionFunction.id]

            const entity = recordToEntity(record);

            const output = await executeFunction(fn, {
                entity: entity
            })

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

    console.log(`Handling event: ${fnName}`)

    switch (fnName) {
        case WatchLogicResources:
            if (event.getResource().getNamespace() != 'logic') {
                throw new Error('WatchLogicResources can only be called for logic namespace')
            }
            load('logic', event.getResource().getName())
            return
        case ExecuteFunction:
            return handleFunctionExecutionCall(event)
        case ResourceOperationTrigger:
            return handleResourceOperationTrigger(event)
        case ResourceOperationRule:
            return handleResourceOperationRule(event)
        default:
            throw new Error('Unknown function name: ' + fnName)
    }
}