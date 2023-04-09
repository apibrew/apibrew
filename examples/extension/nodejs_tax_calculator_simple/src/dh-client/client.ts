import { credentials } from "@grpc/grpc-js";
import { stub } from "./stub/authentication";

export interface DhClientParams {
    addr: string;
    insecure: boolean;
    token: string;
}

export class DhClient {
    public readonly authenticationClient: stub.AuthenticationClient;

    public constructor(params: DhClientParams) {
        this.authenticationClient = new stub.AuthenticationClient(params.addr, credentials.createInsecure())
    }
}