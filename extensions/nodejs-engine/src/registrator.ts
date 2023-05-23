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
import {EXTENSION_NAME, HOST, PORT} from "./config";
import {EventSelector} from "./proto/model/event_pb";
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
    const record = new Record()
    record.getPropertiesMap().set('name', new google_protobuf_struct_pb.Value().setStringValue(EXTENSION_NAME))
    applyFunctionEngineRequest.setRecord(record)
    const result = await toPromise<ApplyRecordRequest, ApplyRecordResponse>(recordClient.apply.bind(recordClient))(applyFunctionEngineRequest)
    engineId = result.getRecord().getId()

    console.log('Engine registration Id:', engineId)
    await initFunctionRegistry(engineId)
}

export async function registerExtension() {
    const extensionClient = new ExtensionClient('localhost:9009', credentials.createInsecure(), null)
    const recordClient = new RecordClient('localhost:9009', credentials.createInsecure(), null)
    await registerFunctionEngine(recordClient);

    const extensionList = await toPromise<ListExtensionRequest, ListExtensionResponse>(extensionClient.list.bind(extensionClient))(new ListExtensionRequest())

    const extension = prepareExtension()
    if (extensionList.getContentList()) {
        for (const existing of extensionList.getContentList()) {
            if (existing.getName() === extension.getName()) {
                console.log('updating extension')
                return await updateExtension(extensionClient, extension)
            }
        }
    }

    console.log('creating extension')
    return await createExtension(extensionClient, extension)

}

function prepareExtension(): Extension {
    const extension = new Extension()
    extension.setName(EXTENSION_NAME)
    extension.setSync(true)
    const call = new ExternalCall()
    const fCall = new FunctionCall()
    fCall.setFunctionname('external-call')
    fCall.setHost(`${HOST}:${PORT}`)
    call.setFunctioncall(fCall)
    extension.setCall(call)
    extension.setOrder(10)
    extension.setResponds(true)
    const eventSelector = new EventSelector()
    eventSelector.setResourcesList(['Function', 'FunctionExecution'])
    eventSelector.setNamespacesList(['extensions'])
    extension.setSelector(eventSelector)

    return extension
}
