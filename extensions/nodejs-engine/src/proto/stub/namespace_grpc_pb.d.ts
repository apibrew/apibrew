// package: stub
// file: stub/namespace.proto

/* tslint:disable */
/* eslint-disable */

import * as grpc from "@grpc/grpc-js";
import * as stub_namespace_pb from "../stub/namespace_pb";
import * as model_namespace_pb from "../model/namespace_pb";
import * as model_error_pb from "../model/error_pb";


interface INamespaceService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
    create: INamespaceService_ICreate;
    list: INamespaceService_IList;
    update: INamespaceService_IUpdate;
    delete: INamespaceService_IDelete;
    get: INamespaceService_IGet;
}

interface INamespaceService_ICreate extends grpc.MethodDefinition<stub_namespace_pb.CreateNamespaceRequest, stub_namespace_pb.CreateNamespaceResponse> {
    path: "/stub.Namespace/Create";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<stub_namespace_pb.CreateNamespaceRequest>;
    requestDeserialize: grpc.deserialize<stub_namespace_pb.CreateNamespaceRequest>;
    responseSerialize: grpc.serialize<stub_namespace_pb.CreateNamespaceResponse>;
    responseDeserialize: grpc.deserialize<stub_namespace_pb.CreateNamespaceResponse>;
}
interface INamespaceService_IList extends grpc.MethodDefinition<stub_namespace_pb.ListNamespaceRequest, stub_namespace_pb.ListNamespaceResponse> {
    path: "/stub.Namespace/List";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<stub_namespace_pb.ListNamespaceRequest>;
    requestDeserialize: grpc.deserialize<stub_namespace_pb.ListNamespaceRequest>;
    responseSerialize: grpc.serialize<stub_namespace_pb.ListNamespaceResponse>;
    responseDeserialize: grpc.deserialize<stub_namespace_pb.ListNamespaceResponse>;
}
interface INamespaceService_IUpdate extends grpc.MethodDefinition<stub_namespace_pb.UpdateNamespaceRequest, stub_namespace_pb.UpdateNamespaceResponse> {
    path: "/stub.Namespace/Update";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<stub_namespace_pb.UpdateNamespaceRequest>;
    requestDeserialize: grpc.deserialize<stub_namespace_pb.UpdateNamespaceRequest>;
    responseSerialize: grpc.serialize<stub_namespace_pb.UpdateNamespaceResponse>;
    responseDeserialize: grpc.deserialize<stub_namespace_pb.UpdateNamespaceResponse>;
}
interface INamespaceService_IDelete extends grpc.MethodDefinition<stub_namespace_pb.DeleteNamespaceRequest, stub_namespace_pb.DeleteNamespaceResponse> {
    path: "/stub.Namespace/Delete";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<stub_namespace_pb.DeleteNamespaceRequest>;
    requestDeserialize: grpc.deserialize<stub_namespace_pb.DeleteNamespaceRequest>;
    responseSerialize: grpc.serialize<stub_namespace_pb.DeleteNamespaceResponse>;
    responseDeserialize: grpc.deserialize<stub_namespace_pb.DeleteNamespaceResponse>;
}
interface INamespaceService_IGet extends grpc.MethodDefinition<stub_namespace_pb.GetNamespaceRequest, stub_namespace_pb.GetNamespaceResponse> {
    path: "/stub.Namespace/Get";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<stub_namespace_pb.GetNamespaceRequest>;
    requestDeserialize: grpc.deserialize<stub_namespace_pb.GetNamespaceRequest>;
    responseSerialize: grpc.serialize<stub_namespace_pb.GetNamespaceResponse>;
    responseDeserialize: grpc.deserialize<stub_namespace_pb.GetNamespaceResponse>;
}

export const NamespaceService: INamespaceService;

export interface INamespaceServer {
    create: grpc.handleUnaryCall<stub_namespace_pb.CreateNamespaceRequest, stub_namespace_pb.CreateNamespaceResponse>;
    list: grpc.handleUnaryCall<stub_namespace_pb.ListNamespaceRequest, stub_namespace_pb.ListNamespaceResponse>;
    update: grpc.handleUnaryCall<stub_namespace_pb.UpdateNamespaceRequest, stub_namespace_pb.UpdateNamespaceResponse>;
    delete: grpc.handleUnaryCall<stub_namespace_pb.DeleteNamespaceRequest, stub_namespace_pb.DeleteNamespaceResponse>;
    get: grpc.handleUnaryCall<stub_namespace_pb.GetNamespaceRequest, stub_namespace_pb.GetNamespaceResponse>;
}

export interface INamespaceClient {
    create(request: stub_namespace_pb.CreateNamespaceRequest, callback: (error: grpc.ServiceError | null, response: stub_namespace_pb.CreateNamespaceResponse) => void): grpc.ClientUnaryCall;
    create(request: stub_namespace_pb.CreateNamespaceRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: stub_namespace_pb.CreateNamespaceResponse) => void): grpc.ClientUnaryCall;
    create(request: stub_namespace_pb.CreateNamespaceRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: stub_namespace_pb.CreateNamespaceResponse) => void): grpc.ClientUnaryCall;
    list(request: stub_namespace_pb.ListNamespaceRequest, callback: (error: grpc.ServiceError | null, response: stub_namespace_pb.ListNamespaceResponse) => void): grpc.ClientUnaryCall;
    list(request: stub_namespace_pb.ListNamespaceRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: stub_namespace_pb.ListNamespaceResponse) => void): grpc.ClientUnaryCall;
    list(request: stub_namespace_pb.ListNamespaceRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: stub_namespace_pb.ListNamespaceResponse) => void): grpc.ClientUnaryCall;
    update(request: stub_namespace_pb.UpdateNamespaceRequest, callback: (error: grpc.ServiceError | null, response: stub_namespace_pb.UpdateNamespaceResponse) => void): grpc.ClientUnaryCall;
    update(request: stub_namespace_pb.UpdateNamespaceRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: stub_namespace_pb.UpdateNamespaceResponse) => void): grpc.ClientUnaryCall;
    update(request: stub_namespace_pb.UpdateNamespaceRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: stub_namespace_pb.UpdateNamespaceResponse) => void): grpc.ClientUnaryCall;
    delete(request: stub_namespace_pb.DeleteNamespaceRequest, callback: (error: grpc.ServiceError | null, response: stub_namespace_pb.DeleteNamespaceResponse) => void): grpc.ClientUnaryCall;
    delete(request: stub_namespace_pb.DeleteNamespaceRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: stub_namespace_pb.DeleteNamespaceResponse) => void): grpc.ClientUnaryCall;
    delete(request: stub_namespace_pb.DeleteNamespaceRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: stub_namespace_pb.DeleteNamespaceResponse) => void): grpc.ClientUnaryCall;
    get(request: stub_namespace_pb.GetNamespaceRequest, callback: (error: grpc.ServiceError | null, response: stub_namespace_pb.GetNamespaceResponse) => void): grpc.ClientUnaryCall;
    get(request: stub_namespace_pb.GetNamespaceRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: stub_namespace_pb.GetNamespaceResponse) => void): grpc.ClientUnaryCall;
    get(request: stub_namespace_pb.GetNamespaceRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: stub_namespace_pb.GetNamespaceResponse) => void): grpc.ClientUnaryCall;
}

export class NamespaceClient extends grpc.Client implements INamespaceClient {
    constructor(address: string, credentials: grpc.ChannelCredentials, options?: object);
    public create(request: stub_namespace_pb.CreateNamespaceRequest, callback: (error: grpc.ServiceError | null, response: stub_namespace_pb.CreateNamespaceResponse) => void): grpc.ClientUnaryCall;
    public create(request: stub_namespace_pb.CreateNamespaceRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: stub_namespace_pb.CreateNamespaceResponse) => void): grpc.ClientUnaryCall;
    public create(request: stub_namespace_pb.CreateNamespaceRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: stub_namespace_pb.CreateNamespaceResponse) => void): grpc.ClientUnaryCall;
    public list(request: stub_namespace_pb.ListNamespaceRequest, callback: (error: grpc.ServiceError | null, response: stub_namespace_pb.ListNamespaceResponse) => void): grpc.ClientUnaryCall;
    public list(request: stub_namespace_pb.ListNamespaceRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: stub_namespace_pb.ListNamespaceResponse) => void): grpc.ClientUnaryCall;
    public list(request: stub_namespace_pb.ListNamespaceRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: stub_namespace_pb.ListNamespaceResponse) => void): grpc.ClientUnaryCall;
    public update(request: stub_namespace_pb.UpdateNamespaceRequest, callback: (error: grpc.ServiceError | null, response: stub_namespace_pb.UpdateNamespaceResponse) => void): grpc.ClientUnaryCall;
    public update(request: stub_namespace_pb.UpdateNamespaceRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: stub_namespace_pb.UpdateNamespaceResponse) => void): grpc.ClientUnaryCall;
    public update(request: stub_namespace_pb.UpdateNamespaceRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: stub_namespace_pb.UpdateNamespaceResponse) => void): grpc.ClientUnaryCall;
    public delete(request: stub_namespace_pb.DeleteNamespaceRequest, callback: (error: grpc.ServiceError | null, response: stub_namespace_pb.DeleteNamespaceResponse) => void): grpc.ClientUnaryCall;
    public delete(request: stub_namespace_pb.DeleteNamespaceRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: stub_namespace_pb.DeleteNamespaceResponse) => void): grpc.ClientUnaryCall;
    public delete(request: stub_namespace_pb.DeleteNamespaceRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: stub_namespace_pb.DeleteNamespaceResponse) => void): grpc.ClientUnaryCall;
    public get(request: stub_namespace_pb.GetNamespaceRequest, callback: (error: grpc.ServiceError | null, response: stub_namespace_pb.GetNamespaceResponse) => void): grpc.ClientUnaryCall;
    public get(request: stub_namespace_pb.GetNamespaceRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: stub_namespace_pb.GetNamespaceResponse) => void): grpc.ClientUnaryCall;
    public get(request: stub_namespace_pb.GetNamespaceRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: stub_namespace_pb.GetNamespaceResponse) => void): grpc.ClientUnaryCall;
}
