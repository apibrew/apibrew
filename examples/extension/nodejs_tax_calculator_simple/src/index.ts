import { credentials } from "@grpc/grpc-js";
import { DhClientImpl, DhClientParams } from "./dh-client/client";
import { ListResourceRequest } from "./dh-client/stub/resource";


const dhClient = new DhClientImpl({ Addr: '127.0.0.1:9009' } as DhClientParams)

async function run() {
    await dhClient.AuthenticateWithUsernameAndPassword("admin", "admin")

    dhClient.GetResourceClient().List(new ListResourceRequest({
        token: dhClient.params.Token
    }), (err, resp) => {
        console.log(resp?.resources)
    })
}


run()