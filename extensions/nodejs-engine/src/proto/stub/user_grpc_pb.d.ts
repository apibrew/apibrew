// package: stub
// file: stub/user.proto

/* tslint:disable */
/* eslint-disable */

import * as grpc from "@grpc/grpc-js";
import * as stub_user_pb from "../stub/user_pb";
import * as model_error_pb from "../model/error_pb";
import * as model_user_pb from "../model/user_pb";
import * as model_query_pb from "../model/query_pb";


interface IUserService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
    create: IUserService_ICreate;
    update: IUserService_IUpdate;
    delete: IUserService_IDelete;
    list: IUserService_IList;
    get: IUserService_IGet;
}

interface IUserService_ICreate extends grpc.MethodDefinition<stub_user_pb.CreateUserRequest, stub_user_pb.CreateUserResponse> {
    path: "/stub.User/Create";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<stub_user_pb.CreateUserRequest>;
    requestDeserialize: grpc.deserialize<stub_user_pb.CreateUserRequest>;
    responseSerialize: grpc.serialize<stub_user_pb.CreateUserResponse>;
    responseDeserialize: grpc.deserialize<stub_user_pb.CreateUserResponse>;
}
interface IUserService_IUpdate extends grpc.MethodDefinition<stub_user_pb.UpdateUserRequest, stub_user_pb.UpdateUserResponse> {
    path: "/stub.User/Update";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<stub_user_pb.UpdateUserRequest>;
    requestDeserialize: grpc.deserialize<stub_user_pb.UpdateUserRequest>;
    responseSerialize: grpc.serialize<stub_user_pb.UpdateUserResponse>;
    responseDeserialize: grpc.deserialize<stub_user_pb.UpdateUserResponse>;
}
interface IUserService_IDelete extends grpc.MethodDefinition<stub_user_pb.DeleteUserRequest, stub_user_pb.DeleteUserResponse> {
    path: "/stub.User/Delete";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<stub_user_pb.DeleteUserRequest>;
    requestDeserialize: grpc.deserialize<stub_user_pb.DeleteUserRequest>;
    responseSerialize: grpc.serialize<stub_user_pb.DeleteUserResponse>;
    responseDeserialize: grpc.deserialize<stub_user_pb.DeleteUserResponse>;
}
interface IUserService_IList extends grpc.MethodDefinition<stub_user_pb.ListUserRequest, stub_user_pb.ListUserResponse> {
    path: "/stub.User/List";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<stub_user_pb.ListUserRequest>;
    requestDeserialize: grpc.deserialize<stub_user_pb.ListUserRequest>;
    responseSerialize: grpc.serialize<stub_user_pb.ListUserResponse>;
    responseDeserialize: grpc.deserialize<stub_user_pb.ListUserResponse>;
}
interface IUserService_IGet extends grpc.MethodDefinition<stub_user_pb.GetUserRequest, stub_user_pb.GetUserResponse> {
    path: "/stub.User/Get";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<stub_user_pb.GetUserRequest>;
    requestDeserialize: grpc.deserialize<stub_user_pb.GetUserRequest>;
    responseSerialize: grpc.serialize<stub_user_pb.GetUserResponse>;
    responseDeserialize: grpc.deserialize<stub_user_pb.GetUserResponse>;
}

export const UserService: IUserService;

export interface IUserServer {
    create: grpc.handleUnaryCall<stub_user_pb.CreateUserRequest, stub_user_pb.CreateUserResponse>;
    update: grpc.handleUnaryCall<stub_user_pb.UpdateUserRequest, stub_user_pb.UpdateUserResponse>;
    delete: grpc.handleUnaryCall<stub_user_pb.DeleteUserRequest, stub_user_pb.DeleteUserResponse>;
    list: grpc.handleUnaryCall<stub_user_pb.ListUserRequest, stub_user_pb.ListUserResponse>;
    get: grpc.handleUnaryCall<stub_user_pb.GetUserRequest, stub_user_pb.GetUserResponse>;
}

export interface IUserClient {
    create(request: stub_user_pb.CreateUserRequest, callback: (error: grpc.ServiceError | null, response: stub_user_pb.CreateUserResponse) => void): grpc.ClientUnaryCall;
    create(request: stub_user_pb.CreateUserRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: stub_user_pb.CreateUserResponse) => void): grpc.ClientUnaryCall;
    create(request: stub_user_pb.CreateUserRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: stub_user_pb.CreateUserResponse) => void): grpc.ClientUnaryCall;
    update(request: stub_user_pb.UpdateUserRequest, callback: (error: grpc.ServiceError | null, response: stub_user_pb.UpdateUserResponse) => void): grpc.ClientUnaryCall;
    update(request: stub_user_pb.UpdateUserRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: stub_user_pb.UpdateUserResponse) => void): grpc.ClientUnaryCall;
    update(request: stub_user_pb.UpdateUserRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: stub_user_pb.UpdateUserResponse) => void): grpc.ClientUnaryCall;
    delete(request: stub_user_pb.DeleteUserRequest, callback: (error: grpc.ServiceError | null, response: stub_user_pb.DeleteUserResponse) => void): grpc.ClientUnaryCall;
    delete(request: stub_user_pb.DeleteUserRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: stub_user_pb.DeleteUserResponse) => void): grpc.ClientUnaryCall;
    delete(request: stub_user_pb.DeleteUserRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: stub_user_pb.DeleteUserResponse) => void): grpc.ClientUnaryCall;
    list(request: stub_user_pb.ListUserRequest, callback: (error: grpc.ServiceError | null, response: stub_user_pb.ListUserResponse) => void): grpc.ClientUnaryCall;
    list(request: stub_user_pb.ListUserRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: stub_user_pb.ListUserResponse) => void): grpc.ClientUnaryCall;
    list(request: stub_user_pb.ListUserRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: stub_user_pb.ListUserResponse) => void): grpc.ClientUnaryCall;
    get(request: stub_user_pb.GetUserRequest, callback: (error: grpc.ServiceError | null, response: stub_user_pb.GetUserResponse) => void): grpc.ClientUnaryCall;
    get(request: stub_user_pb.GetUserRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: stub_user_pb.GetUserResponse) => void): grpc.ClientUnaryCall;
    get(request: stub_user_pb.GetUserRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: stub_user_pb.GetUserResponse) => void): grpc.ClientUnaryCall;
}

export class UserClient extends grpc.Client implements IUserClient {
    constructor(address: string, credentials: grpc.ChannelCredentials, options?: object);
    public create(request: stub_user_pb.CreateUserRequest, callback: (error: grpc.ServiceError | null, response: stub_user_pb.CreateUserResponse) => void): grpc.ClientUnaryCall;
    public create(request: stub_user_pb.CreateUserRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: stub_user_pb.CreateUserResponse) => void): grpc.ClientUnaryCall;
    public create(request: stub_user_pb.CreateUserRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: stub_user_pb.CreateUserResponse) => void): grpc.ClientUnaryCall;
    public update(request: stub_user_pb.UpdateUserRequest, callback: (error: grpc.ServiceError | null, response: stub_user_pb.UpdateUserResponse) => void): grpc.ClientUnaryCall;
    public update(request: stub_user_pb.UpdateUserRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: stub_user_pb.UpdateUserResponse) => void): grpc.ClientUnaryCall;
    public update(request: stub_user_pb.UpdateUserRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: stub_user_pb.UpdateUserResponse) => void): grpc.ClientUnaryCall;
    public delete(request: stub_user_pb.DeleteUserRequest, callback: (error: grpc.ServiceError | null, response: stub_user_pb.DeleteUserResponse) => void): grpc.ClientUnaryCall;
    public delete(request: stub_user_pb.DeleteUserRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: stub_user_pb.DeleteUserResponse) => void): grpc.ClientUnaryCall;
    public delete(request: stub_user_pb.DeleteUserRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: stub_user_pb.DeleteUserResponse) => void): grpc.ClientUnaryCall;
    public list(request: stub_user_pb.ListUserRequest, callback: (error: grpc.ServiceError | null, response: stub_user_pb.ListUserResponse) => void): grpc.ClientUnaryCall;
    public list(request: stub_user_pb.ListUserRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: stub_user_pb.ListUserResponse) => void): grpc.ClientUnaryCall;
    public list(request: stub_user_pb.ListUserRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: stub_user_pb.ListUserResponse) => void): grpc.ClientUnaryCall;
    public get(request: stub_user_pb.GetUserRequest, callback: (error: grpc.ServiceError | null, response: stub_user_pb.GetUserResponse) => void): grpc.ClientUnaryCall;
    public get(request: stub_user_pb.GetUserRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: stub_user_pb.GetUserResponse) => void): grpc.ClientUnaryCall;
    public get(request: stub_user_pb.GetUserRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: stub_user_pb.GetUserResponse) => void): grpc.ClientUnaryCall;
}
