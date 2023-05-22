import {Resource} from "./proto/model/resource_pb";
import {ResourceClient} from "./proto/stub/resource_grpc_pb";

const resource = new Resource()

resource.setName('Taleh123')

console.log(resource.getName())

const client = new ResourceClient('localhost:9009', null, null)

// client.