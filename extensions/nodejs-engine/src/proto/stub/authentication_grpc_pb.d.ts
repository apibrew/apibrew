// package: stub
// file: stub/authentication.proto

/* tslint:disable */
/* eslint-disable */

import * as grpc from "@grpc/grpc-js";
import * as stub_authentication_pb from "../stub/authentication_pb";
import * as model_error_pb from "../model/error_pb";
import * as model_token_pb from "../model/token_pb";


interface IAuthenticationService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
    authenticate: IAuthenticationService_IAuthenticate;
    renewToken: IAuthenticationService_IRenewToken;
}

interface IAuthenticationService_IAuthenticate extends grpc.MethodDefinition<stub_authentication_pb.AuthenticationRequest, stub_authentication_pb.AuthenticationResponse> {
    path: "/stub.Authentication/Authenticate";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<stub_authentication_pb.AuthenticationRequest>;
    requestDeserialize: grpc.deserialize<stub_authentication_pb.AuthenticationRequest>;
    responseSerialize: grpc.serialize<stub_authentication_pb.AuthenticationResponse>;
    responseDeserialize: grpc.deserialize<stub_authentication_pb.AuthenticationResponse>;
}
interface IAuthenticationService_IRenewToken extends grpc.MethodDefinition<stub_authentication_pb.RenewTokenRequest, stub_authentication_pb.RenewTokenResponse> {
    path: "/stub.Authentication/RenewToken";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<stub_authentication_pb.RenewTokenRequest>;
    requestDeserialize: grpc.deserialize<stub_authentication_pb.RenewTokenRequest>;
    responseSerialize: grpc.serialize<stub_authentication_pb.RenewTokenResponse>;
    responseDeserialize: grpc.deserialize<stub_authentication_pb.RenewTokenResponse>;
}

export const AuthenticationService: IAuthenticationService;

export interface IAuthenticationServer {
    authenticate: grpc.handleUnaryCall<stub_authentication_pb.AuthenticationRequest, stub_authentication_pb.AuthenticationResponse>;
    renewToken: grpc.handleUnaryCall<stub_authentication_pb.RenewTokenRequest, stub_authentication_pb.RenewTokenResponse>;
}

export interface IAuthenticationClient {
    authenticate(request: stub_authentication_pb.AuthenticationRequest, callback: (error: grpc.ServiceError | null, response: stub_authentication_pb.AuthenticationResponse) => void): grpc.ClientUnaryCall;
    authenticate(request: stub_authentication_pb.AuthenticationRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: stub_authentication_pb.AuthenticationResponse) => void): grpc.ClientUnaryCall;
    authenticate(request: stub_authentication_pb.AuthenticationRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: stub_authentication_pb.AuthenticationResponse) => void): grpc.ClientUnaryCall;
    renewToken(request: stub_authentication_pb.RenewTokenRequest, callback: (error: grpc.ServiceError | null, response: stub_authentication_pb.RenewTokenResponse) => void): grpc.ClientUnaryCall;
    renewToken(request: stub_authentication_pb.RenewTokenRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: stub_authentication_pb.RenewTokenResponse) => void): grpc.ClientUnaryCall;
    renewToken(request: stub_authentication_pb.RenewTokenRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: stub_authentication_pb.RenewTokenResponse) => void): grpc.ClientUnaryCall;
}

export class AuthenticationClient extends grpc.Client implements IAuthenticationClient {
    constructor(address: string, credentials: grpc.ChannelCredentials, options?: object);
    public authenticate(request: stub_authentication_pb.AuthenticationRequest, callback: (error: grpc.ServiceError | null, response: stub_authentication_pb.AuthenticationResponse) => void): grpc.ClientUnaryCall;
    public authenticate(request: stub_authentication_pb.AuthenticationRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: stub_authentication_pb.AuthenticationResponse) => void): grpc.ClientUnaryCall;
    public authenticate(request: stub_authentication_pb.AuthenticationRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: stub_authentication_pb.AuthenticationResponse) => void): grpc.ClientUnaryCall;
    public renewToken(request: stub_authentication_pb.RenewTokenRequest, callback: (error: grpc.ServiceError | null, response: stub_authentication_pb.RenewTokenResponse) => void): grpc.ClientUnaryCall;
    public renewToken(request: stub_authentication_pb.RenewTokenRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: stub_authentication_pb.RenewTokenResponse) => void): grpc.ClientUnaryCall;
    public renewToken(request: stub_authentication_pb.RenewTokenRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: stub_authentication_pb.RenewTokenResponse) => void): grpc.ClientUnaryCall;
}
