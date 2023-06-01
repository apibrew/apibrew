import {Resource} from "./proto/model/resource_pb";
import {Server, ServerCredentials} from "@grpc/grpc-js";
import {FunctionService, IFunctionServer} from "./proto/ext/function_grpc_pb";
import {FunctionCallResponse} from "./proto/ext/function_pb";
import {initExtensions} from "./registrator";
import {handle} from "./handler";
import {ENGINE_ADDR} from "./config";
import {load} from "./store";
import {reloadInternal} from "./function-registry";

const promises = [
    initExtensions(),
    load('logic', 'Function'),
    load('logic', 'FunctionTrigger'),
    load('logic', 'ResourceRule'),
]

Promise.all(promises).then(() => {
    console.log('All resources loaded')
    reloadInternal()
})

const server = new Server();

const functionCallHandler: IFunctionServer['functionCall'] = (call, callback) => {
    console.log('Function call:', call.request.getName())


    handle(call.request.getName(), call.request.getEvent()).then(processedEvent => {
        const response = new FunctionCallResponse()
        response.setEvent(processedEvent)
        callback(null, response)
    }, err => {
        callback(err, null)
    })
}

server.addService(FunctionService, {
    functionCall: functionCallHandler
})

server.bindAsync(ENGINE_ADDR, ServerCredentials.createInsecure(), () => {
    server.start();
});