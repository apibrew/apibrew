import {ENGINE_REMOTE_ADDR, EXTENSION_NAME, TOKEN} from "./config";
import {initFunctionRegistry} from "./function-registry";
import {ExecuteFunction, WatchLogicResources} from "./const";
import {Event_Action, EventSelector} from "./gen/model/event_pb";
import {Extension} from "./gen/model/extension_pb";
import {
    CreateExtensionRequest,
    ListExtensionRequest,
    ListExtensionResponse,
    UpdateExtensionRequest
} from "./gen/stub/extension_pb";
import {extensionClient, recordClient} from "./client";
import {ApplyRecordRequest} from "./gen/stub/record_pb";
import {Record} from "./gen/model/record_pb";
import {Value} from "@bufbuild/protobuf";
import {ApplyRecordResponse} from "./gen/stub/record_pb";
import {ExternalCall, FunctionCall} from "./gen/model/external_pb";

export let engineId: string

async function registerFunctionEngine() {
    const applyFunctionEngineRequest = new ApplyRecordRequest()
    applyFunctionEngineRequest.namespace = ('logic')
    applyFunctionEngineRequest.resource = ('FunctionExecutionEngine')
    applyFunctionEngineRequest.token = (TOKEN)
    const record = new Record()
    record.properties.name = Value.fromJson(EXTENSION_NAME)
    applyFunctionEngineRequest.record = record
    const Apply = recordClient.apply as any
    const result = (await Apply(applyFunctionEngineRequest)) as ApplyRecordResponse
    engineId = result.record.id

    console.log('Engine registration Id:', engineId)
    await initFunctionRegistry(engineId)
}

export async function registerExtensions(extensions: Extension[]) {

    const listRequest = new ListExtensionRequest()

    listRequest.token = TOKEN
    const existingExtensionResponse$ = await extensionClient.list(listRequest) as ListExtensionResponse

    const existingExtensions = existingExtensionResponse$.content

    for (const extension of extensions) {
        let found = false

        for (const existing of existingExtensions) {
            if (existing.name === extension.name) {
                console.log('updating extension')
                await extensionClient.update(new UpdateExtensionRequest({
                    extensions: [{
                        ...extension,
                    }]
                }))
                found = true
            }
        }

        if (!found) {
            console.log('creating extension')
            await extensionClient.create(new CreateExtensionRequest({
                extensions: [{
                    ...extension,
                }]
            }))
        }
    }
}

export async function initExtensions() {
    await registerFunctionEngine();

    const extensions = prepareExtensions()

    await registerExtensions(extensions);
}

function prepareExtensions(): Extension[] {
    return [
        prepareFunctionExtension(),
        prepareFunctionExecutionExtension()
    ]
}

function prepareFunctionExtension(): Extension {
    const extension = new Extension()
    extension.name = (EXTENSION_NAME)
    extension.sync = (true)
    const call = new ExternalCall()
    const fCall = new FunctionCall()
    fCall.functionName = (WatchLogicResources)
    fCall.host = (ENGINE_REMOTE_ADDR)
    call.functionCall = (fCall)
    extension.call = (call)
    extension.order = (10000)
    extension.responds = (false)
    extension.sync = (false)
    const eventSelector = new EventSelector()
    eventSelector.namespaces = (['logic'])
    eventSelector.actions = ([Event_Action.CREATE, Event_Action.UPDATE, Event_Action.DELETE])
    extension.selector = (eventSelector)

    return extension
}

function prepareFunctionExecutionExtension(): Extension {
    const extension = new Extension()
    extension.name = (`${EXTENSION_NAME}-execution`)
    extension.sync = (true)
    const call = new ExternalCall()
    const fCall = new FunctionCall()
    fCall.functionName = (ExecuteFunction)
    fCall.host = (ENGINE_REMOTE_ADDR)
    call.functionCall = (fCall)
    extension.call = (call)
    extension.order = (10)
    extension.responds = (true)
    extension.finalizes = (true)
    const eventSelector = new EventSelector()
    eventSelector.resources = (['FunctionExecution'])
    eventSelector.namespaces = (['logic'])
    extension.selector = (eventSelector)

    return extension
}
