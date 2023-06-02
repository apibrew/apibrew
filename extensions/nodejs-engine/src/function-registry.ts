import {read, filter} from "./store";
import {Function, FunctionName} from "./model/function";
import {FunctionTrigger, FunctionTriggerName} from "./model/function-trigger";
import {ResourceRule, ResourceRuleName} from "./model/resource-rule";
import {Extension} from "./proto/model/extension_pb";
import {ENGINE_REMOTE_ADDR, EXTENSION_NAME} from "./config";
import {ExternalCall, FunctionCall} from "./proto/model/external_pb";
import {Event, EventSelector} from "./proto/model/event_pb";
import {registerExtensions} from "./registrator";
import {ResourceOperationRule, ResourceOperationTrigger} from "./const";

let engineId: string

export let functionMap: { [key: string]: Function } = {}
export let functionIdMap: { [key: string]: Function } = {}


export async function reloadInternal() {
    filter('logic', FunctionName, (record: Function) => record.engine.id === engineId)
    const functions = read<Function>('logic', FunctionName)

    filter('logic', FunctionTriggerName, (record: FunctionTrigger) => functions.some(fn => fn.id === record.function.id))
    const triggers = read<FunctionTrigger>('logic', FunctionTriggerName)

    filter('logic', ResourceRuleName, (record: ResourceRule) => functions.some(fn => fn.id === record.conditionFunction.id))
    const rules = read<ResourceRule>('logic', ResourceRuleName)

    functions.forEach((record) => {
        if (record.engine.id === engineId) {
            functionMap[record.package + '/' + record.name] = record
            functionIdMap[record.id] = record
        }
    })

    let extensions: Extension[] = []

    for (const trigger of triggers) {
        extensions.push(prepareExtensionFromTrigger(trigger))
    }

    for (const rule of rules) {
        extensions.push(prepareExtensionFromRule(rule))
    }

    extensions = extensions.filter((item, index) => extensions.findIndex(item2 => JSON.stringify(item) === JSON.stringify(item2)) === index)

    await registerExtensions(extensions)

    console.log('Configuring extensions: ', extensions.map(item => item.getName()))
}

function prepareExtensionFromTrigger(trigger: FunctionTrigger): Extension {
    const extension = new Extension()
    extension.setName(`${EXTENSION_NAME}_trigger_${trigger.namespace}_${trigger.resource}`)
    extension.setSync(!trigger.async)
    const call = new ExternalCall()
    const fCall = new FunctionCall()
    fCall.setFunctionname(ResourceOperationTrigger)
    fCall.setHost(ENGINE_REMOTE_ADDR)
    call.setFunctioncall(fCall)
    extension.setCall(call)
    extension.setResponds(true)

    if (trigger.order) {
        switch (trigger.order) {
            case 'before':
                extension.setOrder(10)

                break
            case 'after':
                extension.setOrder(200)
                break
            case 'instead':
                extension.setOrder(80)
                extension.setFinalizes(true)
                break
        }
    }

    const eventSelector = new EventSelector()
    eventSelector.setNamespacesList([trigger.namespace])
    eventSelector.setResourcesList([trigger.resource])
    let action: Event.Action
    switch (trigger.action) {
        case 'create':
            action = Event.Action.CREATE
            break
        case 'update':
            action = Event.Action.UPDATE
            break
        case 'delete':
            action = Event.Action.DELETE
            break
        default:
            throw new Error('Unknown action: ' + trigger.action)
    }
    eventSelector.setActionsList([action])
    extension.setSelector(eventSelector)

    return extension
}

function prepareExtensionFromRule(rule: ResourceRule): Extension {
    const extension = new Extension()
    extension.setName(`${EXTENSION_NAME}_rule_${rule.namespace}_${rule.resource}`)
    extension.setSync(true)
    const call = new ExternalCall()
    const fCall = new FunctionCall()
    fCall.setFunctionname(ResourceOperationRule)
    fCall.setHost(ENGINE_REMOTE_ADDR)
    call.setFunctioncall(fCall)
    extension.setCall(call)
    extension.setOrder(85)

    const eventSelector = new EventSelector()
    eventSelector.setNamespacesList([rule.namespace])
    eventSelector.setResourcesList([rule.resource])
    eventSelector.setActionsList([Event.Action.CREATE, Event.Action.UPDATE])
    extension.setSelector(eventSelector)

    return extension
}

export async function initFunctionRegistry(_engineId: string) {
    engineId = _engineId
}