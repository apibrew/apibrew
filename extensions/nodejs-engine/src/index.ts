import {Resource} from "./proto/model/resource_pb";
import {ResourceClient} from "./proto/stub/resource_grpc_pb";
import {credentials, Server, ServerCredentials} from "@grpc/grpc-js";
import {ListResourceRequest} from "./proto/stub/resource_pb";
import {FunctionService, IFunctionServer} from "./proto/ext/function_grpc_pb";
import {FunctionCallResponse} from "./proto/ext/function_pb";
import {ExtensionClient} from "./proto/stub/extension_grpc_pb";
import {registerExtension} from "./registrator";
import {HOST, PORT} from "./config";

const resource = new Resource()

resource.setName('Taleh123')

console.log(resource.getName())

registerExtension()

const server = new Server();

const functionCallHandler: IFunctionServer['functionCall'] = (call, callback) => {
    console.log(call.request.getName())
    console.log(call.request.getEvent())

    const response = new FunctionCallResponse()

    response.setEvent(call.request.getEvent())

    console.log(response.getEvent())

    callback(null, response)
}

server.addService(FunctionService, {
    functionCall: functionCallHandler
})

server.bindAsync(`${HOST}:${PORT}`, ServerCredentials.createInsecure(), () => {
    server.start();
});