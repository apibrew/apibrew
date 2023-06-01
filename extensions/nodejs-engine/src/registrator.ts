import {ExtensionClient} from "./proto/stub/extension_grpc_pb";
import {credentials} from "@grpc/grpc-js";
import {
    CreateExtensionRequest,
    ListExtensionRequest,
    ListExtensionResponse,
    UpdateExtensionRequest
} from "./proto/stub/extension_pb";
import {toPromise} from "./util";
import {Extension} from "./proto/model/extension_pb";
import {ExternalCall, FunctionCall} from "./proto/model/external_pb";
import {APBR_ADDR, ENGINE_REMOTE_ADDR, EXTENSION_NAME, TOKEN} from "./config";
import {Event, EventSelector} from "./proto/model/event_pb";
import {RecordClient} from "./proto/stub/record_grpc_pb";
import {ApplyRecordRequest, ApplyRecordResponse} from "./proto/stub/record_pb";
import {Record} from "./proto/model/record_pb";
import * as google_protobuf_struct_pb from "google-protobuf/google/protobuf/struct_pb";
import {initFunctionRegistry} from "./function-registry";
import {ExecuteFunction, WatchLogicResources} from "./const";

export let engineId: string

async function updateExtension(client: ExtensionClient, extension: Extension) {
    const request = new UpdateExtensionRequest()
    request.setExtensionsList([extension])
    await toPromise(client.update.bind(client))(request)
}

async function createExtension(client: ExtensionClient, extension: Extension) {
    const request = new CreateExtensionRequest()
    request.setExtensionsList([extension])
    await toPromise(client.create.bind(client))(request)
}

async function registerFunctionEngine() {
    const recordClient = new RecordClient(APBR_ADDR, credentials.createInsecure(), null)
    const applyFunctionEngineRequest = new ApplyRecordRequest()
    applyFunctionEngineRequest.setNamespace('logic')
    applyFunctionEngineRequest.setResource('FunctionExecutionEngine')
    applyFunctionEngineRequest.setToken(TOKEN)
    const record = new Record()
    record.getPropertiesMap().set('name', new google_protobuf_struct_pb.Value().setStringValue(EXTENSION_NAME))
    applyFunctionEngineRequest.setRecord(record)
    const result = await toPromise<ApplyRecordRequest, ApplyRecordResponse>(recordClient.apply.bind(recordClient))(applyFunctionEngineRequest)
    engineId = result.getRecord().getId()

    console.log('Engine registration Id:', engineId)
    await initFunctionRegistry(engineId)
}

export async function registerExtensions(extensions: Extension[]) {
    const extensionClient = new ExtensionClient(APBR_ADDR, credentials.createInsecure(), null)
    const listRequest = new ListExtensionRequest()

    listRequest.setToken(TOKEN)
    const existingExtensionResponse$ = await toPromise<ListExtensionRequest, ListExtensionResponse>(extensionClient.list.bind(extensionClient))(listRequest)

    const existingExtensions = (await existingExtensionResponse$).getContentList()

    for (const extension of extensions) {
        let found = false

        for (const existing of existingExtensions) {
            if (existing.getName() === extension.getName()) {
                console.log('updating extension')
                await updateExtension(extensionClient, extension)
                found = true
            }
        }

        if (!found) {
            console.log('creating extension')
            await createExtension(extensionClient, extension)
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
    extension.setName(EXTENSION_NAME)
    extension.setSync(true)
    const call = new ExternalCall()
    const fCall = new FunctionCall()
    fCall.setFunctionname(WatchLogicResources)
    fCall.setHost(ENGINE_REMOTE_ADDR)
    call.setFunctioncall(fCall)
    extension.setCall(call)
    extension.setOrder(10000)
    extension.setResponds(false)
    extension.setSync(false)
    const eventSelector = new EventSelector()
    eventSelector.setNamespacesList(['logic'])
    eventSelector.setActionsList([Event.Action.CREATE, Event.Action.UPDATE, Event.Action.DELETE])
    extension.setSelector(eventSelector)

    return extension
}

function prepareFunctionExecutionExtension(): Extension {
    const extension = new Extension()
    extension.setName(`${EXTENSION_NAME}-execution`)
    extension.setSync(true)
    const call = new ExternalCall()
    const fCall = new FunctionCall()
    fCall.setFunctionname(ExecuteFunction)
    fCall.setHost(ENGINE_REMOTE_ADDR)
    call.setFunctioncall(fCall)
    extension.setCall(call)
    extension.setOrder(10)
    extension.setResponds(true)
    extension.setFinalizes(true)
    const eventSelector = new EventSelector()
    eventSelector.setResourcesList(['FunctionExecution'])
    eventSelector.setNamespacesList(['logic'])
    extension.setSelector(eventSelector)

    return extension
}
