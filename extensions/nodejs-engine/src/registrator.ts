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
import {APBR_HOST, APBR_PORT, ENGINE_REMOTE_ADDR, EXTENSION_NAME, TOKEN} from "./config";
import {Event, EventSelector} from "./proto/model/event_pb";
import {RecordClient} from "./proto/stub/record_grpc_pb";
import {ApplyRecordRequest, ApplyRecordResponse} from "./proto/stub/record_pb";
import {Record} from "./proto/model/record_pb";
import * as google_protobuf_struct_pb from "google-protobuf/google/protobuf/struct_pb";
import {initFunctionRegistry} from "./function-registry";

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

async function registerFunctionEngine(recordClient: RecordClient) {
    const applyFunctionEngineRequest = new ApplyRecordRequest()
    applyFunctionEngineRequest.setNamespace('extensions')
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

export async function registerExtension() {
    const extensionClient = new ExtensionClient(`${APBR_HOST}:${APBR_PORT}`, credentials.createInsecure(), null)
    const recordClient = new RecordClient(`${APBR_HOST}:${APBR_PORT}`, credentials.createInsecure(), null)
    await registerFunctionEngine(recordClient);
    const listRequest = new ListExtensionRequest()
    listRequest.setToken(TOKEN)

    const extensionList = await toPromise<ListExtensionRequest, ListExtensionResponse>(extensionClient.list.bind(extensionClient))(listRequest)

    const functionExtension = prepareFunctionExtension()
    const functionExecutionExtension = prepareFunctionExecutionExtension()
    let functionExtensionExists = false
    let functionExecutionExtensionExists = false

    if (extensionList.getContentList()) {
        for (const existing of extensionList.getContentList()) {
            if (existing.getName() === functionExtension.getName()) {
                console.log('updating function extension')
                await updateExtension(extensionClient, functionExtension)
                functionExtensionExists = true
            }

            if (existing.getName() === functionExecutionExtension.getName()) {
                console.log('updating function extension')
                await updateExtension(extensionClient, functionExecutionExtension)
                functionExecutionExtensionExists = true
            }
        }
    }

    if (!functionExtensionExists) {
        console.log('creating extension')
        await createExtension(extensionClient, functionExtension)
    }

    if (!functionExecutionExtensionExists) {
        console.log('creating extension')
        await createExtension(extensionClient, functionExecutionExtension)
    }
}

function prepareFunctionExtension(): Extension {
    const extension = new Extension()
    extension.setName(EXTENSION_NAME)
    extension.setSync(true)
    const call = new ExternalCall()
    const fCall = new FunctionCall()
    fCall.setFunctionname('external-call')
    fCall.setHost(ENGINE_REMOTE_ADDR)
    call.setFunctioncall(fCall)
    extension.setCall(call)
    extension.setOrder(10)
    extension.setResponds(true)
    const eventSelector = new EventSelector()
    eventSelector.setResourcesList(['Function'])
    eventSelector.setNamespacesList(['extensions'])
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
    fCall.setFunctionname('external-call')
    fCall.setHost(ENGINE_REMOTE_ADDR)
    call.setFunctioncall(fCall)
    extension.setCall(call)
    extension.setOrder(10)
    extension.setResponds(true)
    extension.setFinalizes(true)
    const eventSelector = new EventSelector()
    eventSelector.setResourcesList(['FunctionExecution'])
    eventSelector.setNamespacesList(['extensions'])
    extension.setSelector(eventSelector)

    return extension
}
