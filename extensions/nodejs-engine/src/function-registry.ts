import {Event} from "./proto/model/event_pb";
import {credentials} from "@grpc/grpc-js";
import {RecordClient} from "./proto/stub/record_grpc_pb";
import {ListRecordRequest, ListRecordResponse} from "./proto/stub/record_pb";
import {toPromise} from "./util";
import {APBR_ADDR, TOKEN} from "./config";

let engineId: string

const recordClient = new RecordClient(APBR_ADDR, credentials.createInsecure(), null)

export interface Function {
    name: string
    packageName: string
    script: string
}

export let functionMap: { [key: string]: Function } = {}

export async function reloadFunctions() {
    console.log('Reloading Functions')
    console.log('`${APBR_HOST}:${APBR_PORT}`', APBR_ADDR)
    const request = new ListRecordRequest()
    request.setNamespace('extensions')
    request.setResource('Function')
    request.setToken(TOKEN)
    const functions = await toPromise<ListRecordRequest, ListRecordResponse>(recordClient.list.bind(recordClient))(request)

    functionMap = {}

    functions.getContentList().forEach((record) => {
        const name = record.getPropertiesMap().get('name').getStringValue()
        const packageName = record.getPropertiesMap().get('package').getStringValue()
        const script = record.getPropertiesMap().get('script').getStringValue()
        const fnEngineId = record.getPropertiesMap().get('engine').getStructValue().getFieldsMap().get('id').getStringValue()

        if (fnEngineId === engineId) {
            functionMap[packageName + '/' + name] = {
                name,
                packageName,
                script
            }
        }
    })

    console.log('Functions reloaded')
    console.log(functionMap)
}

export async function initFunctionRegistry(_engineId: string) {
    engineId = _engineId

    await reloadFunctions()
}

export async function handleFunctionCall(event: Event): Promise<Event> {
    if (event.getAction() == Event.Action.CREATE || event.getAction() == Event.Action.UPDATE || event.getAction() == Event.Action.DELETE) {
        setTimeout(async () => {
            await reloadFunctions()
        }, 100)
    }
    return event;
}