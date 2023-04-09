import { credentials } from "@grpc/grpc-js";
import { stub } from "./dh-client/stub/authentication";

const client = new stub.AuthenticationClient('127.0.0.1:9009', credentials.createInsecure())

client.Authenticate(stub.AuthenticationRequest.fromObject({
    "username": "admin",
    "password": "admin"
}), (err, resp) => {
    console.log(err, resp?.token.content)
})