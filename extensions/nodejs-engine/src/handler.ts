import { functionIdMap, functionMap, reloadInternal } from "./function-registry";
import { load, read } from "./store";
import {
    ExecuteFunction,
    ResourceOperationRule,
    ResourceOperationTrigger,
    WatchLogicResources
} from "./const";
import { ResourceRule, ResourceRuleName } from "./model/resource-rule";
import { FunctionTrigger, FunctionTriggerName } from "./model/function-trigger";
import { components } from "./model/base-schema";
import { executeFunction, locateFunction } from "./function-execute";

type Event = components['schemas']['Event']

const { VM } = require('vm2');

export async function handleFunctionExecutionCall(event: Event) {
    for (const record of event.records) {
        try {
            const packageName = record.properties.function.package
            const name = record.properties.function.name
            const input = record.properties.input
            const fn = locateFunction(packageName, name)

            const output = (await executeFunction(fn, input) ?? { ok: true });

            record.properties.output = output

            record.properties.status = 'success'
        } catch (e) {
            console.error(e)

            record.properties.error = e.message
            record.properties.status = 'error'
        }
    }

    return event;
}

function locateRulesListeningToEvent(event: Event): ResourceRule[] {
    if (event.action != 'CREATE' && event.action != 'UPDATE') {
        return []
    }
    const rules = read<ResourceRule>('logic', ResourceRuleName)

    return rules.filter(rule => rule.namespace == event.resource.namespace && rule.resource == event.resource.name)
}

function locateTriggersListeningToEvent(event: Event): FunctionTrigger[] {
    if (event.action != 'CREATE' && event.action != 'UPDATE') {
        return []
    }
    const triggers = read<FunctionTrigger>('logic', FunctionTriggerName)

    return triggers.filter(trigger => {
        if (trigger.async == event.sync) {
            console.log('trigger async mismatch', trigger.async, event.sync)
            return false
        }

        if (trigger.resource != '*' && trigger.resource != event.resource.name) {
            console.log('trigger resource mismatch', trigger.resource, event.resource.name)
            return false
        }

        if (trigger.namespace != '*' && trigger.namespace != event.resource.namespace) {
            console.log('trigger namespace mismatch', trigger.namespace, event.resource.namespace)
            return false
        }

        if (trigger.action.toLowerCase() != event.action.toLowerCase()) {
            console.log('trigger action mismatch', trigger.action, event.action.toString().toLowerCase())
            return false
        }

        return true
    })
}

function recordToEntity(record: any) {
    const entity = {}

    for (const key of Object.keys(record.properties)) {
        const entry = record.properties[key]

        entity[key] = entry.toJson()
    }

    return entity;
}

async function handleResourceOperationTrigger(event: Event): Promise<Event> {
    console.log('Handling resource operation trigger', event.action)

    const triggers = locateTriggersListeningToEvent(event)

    for (const record of event.records) {
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

            for (const property of event.resource.properties) {
                if (output[property.name] === record.properties[property.name]) {
                    continue
                }
                record.properties[property.name] = output[property.name]
            }
        }
    }

    return event
}

async function handleResourceOperationRule(event: Event): Promise<Event> {
    console.log('Handling resource operation rule', event.action)

    const rules = locateRulesListeningToEvent(event)

    for (const record of event.records) {
        for (const rule of rules) {
            console.log('Handling rule', rule.name)
            const fn = functionIdMap[rule.conditionFunction.id]

            const entity = recordToEntity(record);

            const output = await executeFunction(fn, {
                entity: entity
            })

            console.log('output', output)

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
            if (event.resource.namespace != 'logic') {
                throw new Error('WatchLogicResources can only be called for logic namespace')
            }
            load('logic', event.resource.name).then(() => {
                reloadInternal()
            })
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