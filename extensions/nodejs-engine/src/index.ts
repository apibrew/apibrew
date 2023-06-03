import {initExtensions} from "./registrator";
import {handle} from "./handler";
import {load} from "./store";
import {reloadInternal} from "./function-registry";
import * as http2 from "http2";
import {connectNodeAdapter} from "@bufbuild/connect-node";
import {ConnectRouter} from "@bufbuild/connect";
import {Function} from "./gen/ext/function_connect";
import {FunctionCallRequest, FunctionCallResponse} from "./gen/ext/function_pb";

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

const routes = (router: ConnectRouter) =>
    router.service(Function, {
        // implements rpc Say
        async functionCall(req: FunctionCallRequest) {
            const response = new FunctionCallResponse()
            response.event = await handle(req.name, req.event)
            return response
        },
    });

http2.createServer(
    connectNodeAdapter({
        routes,
        grpc: true,
        grpcWeb: false,
        connect: false,
    }) // responds with 404 for other requests
).listen(23619);