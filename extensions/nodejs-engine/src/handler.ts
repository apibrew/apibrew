import {Event} from "./proto/model/event_pb";
import {Function, functionMap, handleFunctionCall} from "./function-registry";
import {Value} from "google-protobuf/google/protobuf/struct_pb";

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

            console.log('running', fn)

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

export async function handle(event: Event): Promise<Event> {
    if (!event) {
        return event
    }
    switch (event.getResource().getName()) {
        case "Function":
            return handleFunctionCall(event)
        case "FunctionExecution":
            return handleFunctionExecutionCall(event)
        default:
            throw new Error("Unknown resource: " + event.getResource().getName())
    }
}