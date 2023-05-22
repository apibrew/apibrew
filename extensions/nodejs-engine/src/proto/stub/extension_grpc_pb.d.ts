// package: stub
// file: stub/extension.proto

/* tslint:disable */
/* eslint-disable */

import * as grpc from "@grpc/grpc-js";
import * as stub_extension_pb from "../stub/extension_pb";
import * as model_query_pb from "../model/query_pb";
import * as model_extension_pb from "../model/extension_pb";

interface IExtensionService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
    list: IExtensionService_IList;
    get: IExtensionService_IGet;
    create: IExtensionService_ICreate;
    update: IExtensionService_IUpdate;
    delete: IExtensionService_IDelete;
}

interface IExtensionService_IList extends grpc.MethodDefinition<stub_extension_pb.ListExtensionRequest, stub_extension_pb.ListExtensionResponse> {
    path: "/stub.Extension/List";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<stub_extension_pb.ListExtensionRequest>;
    requestDeserialize: grpc.deserialize<stub_extension_pb.ListExtensionRequest>;
    responseSerialize: grpc.serialize<stub_extension_pb.ListExtensionResponse>;
    responseDeserialize: grpc.deserialize<stub_extension_pb.ListExtensionResponse>;
}
interface IExtensionService_IGet extends grpc.MethodDefinition<stub_extension_pb.GetExtensionRequest, stub_extension_pb.GetExtensionResponse> {
    path: "/stub.Extension/Get";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<stub_extension_pb.GetExtensionRequest>;
    requestDeserialize: grpc.deserialize<stub_extension_pb.GetExtensionRequest>;
    responseSerialize: grpc.serialize<stub_extension_pb.GetExtensionResponse>;
    responseDeserialize: grpc.deserialize<stub_extension_pb.GetExtensionResponse>;
}
interface IExtensionService_ICreate extends grpc.MethodDefinition<stub_extension_pb.CreateExtensionRequest, stub_extension_pb.CreateExtensionResponse> {
    path: "/stub.Extension/Create";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<stub_extension_pb.CreateExtensionRequest>;
    requestDeserialize: grpc.deserialize<stub_extension_pb.CreateExtensionRequest>;
    responseSerialize: grpc.serialize<stub_extension_pb.CreateExtensionResponse>;
    responseDeserialize: grpc.deserialize<stub_extension_pb.CreateExtensionResponse>;
}
interface IExtensionService_IUpdate extends grpc.MethodDefinition<stub_extension_pb.UpdateExtensionRequest, stub_extension_pb.UpdateExtensionResponse> {
    path: "/stub.Extension/Update";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<stub_extension_pb.UpdateExtensionRequest>;
    requestDeserialize: grpc.deserialize<stub_extension_pb.UpdateExtensionRequest>;
    responseSerialize: grpc.serialize<stub_extension_pb.UpdateExtensionResponse>;
    responseDeserialize: grpc.deserialize<stub_extension_pb.UpdateExtensionResponse>;
}
interface IExtensionService_IDelete extends grpc.MethodDefinition<stub_extension_pb.DeleteExtensionRequest, stub_extension_pb.DeleteExtensionResponse> {
    path: "/stub.Extension/Delete";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<stub_extension_pb.DeleteExtensionRequest>;
    requestDeserialize: grpc.deserialize<stub_extension_pb.DeleteExtensionRequest>;
    responseSerialize: grpc.serialize<stub_extension_pb.DeleteExtensionResponse>;
    responseDeserialize: grpc.deserialize<stub_extension_pb.DeleteExtensionResponse>;
}

export const ExtensionService: IExtensionService;

export interface IExtensionServer {
    list: grpc.handleUnaryCall<stub_extension_pb.ListExtensionRequest, stub_extension_pb.ListExtensionResponse>;
    get: grpc.handleUnaryCall<stub_extension_pb.GetExtensionRequest, stub_extension_pb.GetExtensionResponse>;
    create: grpc.handleUnaryCall<stub_extension_pb.CreateExtensionRequest, stub_extension_pb.CreateExtensionResponse>;
    update: grpc.handleUnaryCall<stub_extension_pb.UpdateExtensionRequest, stub_extension_pb.UpdateExtensionResponse>;
    delete: grpc.handleUnaryCall<stub_extension_pb.DeleteExtensionRequest, stub_extension_pb.DeleteExtensionResponse>;
}

export interface IExtensionClient {
    list(request: stub_extension_pb.ListExtensionRequest, callback: (error: grpc.ServiceError | null, response: stub_extension_pb.ListExtensionResponse) => void): grpc.ClientUnaryCall;
    list(request: stub_extension_pb.ListExtensionRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: stub_extension_pb.ListExtensionResponse) => void): grpc.ClientUnaryCall;
    list(request: stub_extension_pb.ListExtensionRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: stub_extension_pb.ListExtensionResponse) => void): grpc.ClientUnaryCall;
    get(request: stub_extension_pb.GetExtensionRequest, callback: (error: grpc.ServiceError | null, response: stub_extension_pb.GetExtensionResponse) => void): grpc.ClientUnaryCall;
    get(request: stub_extension_pb.GetExtensionRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: stub_extension_pb.GetExtensionResponse) => void): grpc.ClientUnaryCall;
    get(request: stub_extension_pb.GetExtensionRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: stub_extension_pb.GetExtensionResponse) => void): grpc.ClientUnaryCall;
    create(request: stub_extension_pb.CreateExtensionRequest, callback: (error: grpc.ServiceError | null, response: stub_extension_pb.CreateExtensionResponse) => void): grpc.ClientUnaryCall;
    create(request: stub_extension_pb.CreateExtensionRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: stub_extension_pb.CreateExtensionResponse) => void): grpc.ClientUnaryCall;
    create(request: stub_extension_pb.CreateExtensionRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: stub_extension_pb.CreateExtensionResponse) => void): grpc.ClientUnaryCall;
    update(request: stub_extension_pb.UpdateExtensionRequest, callback: (error: grpc.ServiceError | null, response: stub_extension_pb.UpdateExtensionResponse) => void): grpc.ClientUnaryCall;
    update(request: stub_extension_pb.UpdateExtensionRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: stub_extension_pb.UpdateExtensionResponse) => void): grpc.ClientUnaryCall;
    update(request: stub_extension_pb.UpdateExtensionRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: stub_extension_pb.UpdateExtensionResponse) => void): grpc.ClientUnaryCall;
    delete(request: stub_extension_pb.DeleteExtensionRequest, callback: (error: grpc.ServiceError | null, response: stub_extension_pb.DeleteExtensionResponse) => void): grpc.ClientUnaryCall;
    delete(request: stub_extension_pb.DeleteExtensionRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: stub_extension_pb.DeleteExtensionResponse) => void): grpc.ClientUnaryCall;
    delete(request: stub_extension_pb.DeleteExtensionRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: stub_extension_pb.DeleteExtensionResponse) => void): grpc.ClientUnaryCall;
}

export class ExtensionClient extends grpc.Client implements IExtensionClient {
    constructor(address: string, credentials: grpc.ChannelCredentials, options?: object);
    public list(request: stub_extension_pb.ListExtensionRequest, callback: (error: grpc.ServiceError | null, response: stub_extension_pb.ListExtensionResponse) => void): grpc.ClientUnaryCall;
    public list(request: stub_extension_pb.ListExtensionRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: stub_extension_pb.ListExtensionResponse) => void): grpc.ClientUnaryCall;
    public list(request: stub_extension_pb.ListExtensionRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: stub_extension_pb.ListExtensionResponse) => void): grpc.ClientUnaryCall;
    public get(request: stub_extension_pb.GetExtensionRequest, callback: (error: grpc.ServiceError | null, response: stub_extension_pb.GetExtensionResponse) => void): grpc.ClientUnaryCall;
    public get(request: stub_extension_pb.GetExtensionRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: stub_extension_pb.GetExtensionResponse) => void): grpc.ClientUnaryCall;
    public get(request: stub_extension_pb.GetExtensionRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: stub_extension_pb.GetExtensionResponse) => void): grpc.ClientUnaryCall;
    public create(request: stub_extension_pb.CreateExtensionRequest, callback: (error: grpc.ServiceError | null, response: stub_extension_pb.CreateExtensionResponse) => void): grpc.ClientUnaryCall;
    public create(request: stub_extension_pb.CreateExtensionRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: stub_extension_pb.CreateExtensionResponse) => void): grpc.ClientUnaryCall;
    public create(request: stub_extension_pb.CreateExtensionRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: stub_extension_pb.CreateExtensionResponse) => void): grpc.ClientUnaryCall;
    public update(request: stub_extension_pb.UpdateExtensionRequest, callback: (error: grpc.ServiceError | null, response: stub_extension_pb.UpdateExtensionResponse) => void): grpc.ClientUnaryCall;
    public update(request: stub_extension_pb.UpdateExtensionRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: stub_extension_pb.UpdateExtensionResponse) => void): grpc.ClientUnaryCall;
    public update(request: stub_extension_pb.UpdateExtensionRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: stub_extension_pb.UpdateExtensionResponse) => void): grpc.ClientUnaryCall;
    public delete(request: stub_extension_pb.DeleteExtensionRequest, callback: (error: grpc.ServiceError | null, response: stub_extension_pb.DeleteExtensionResponse) => void): grpc.ClientUnaryCall;
    public delete(request: stub_extension_pb.DeleteExtensionRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: stub_extension_pb.DeleteExtensionResponse) => void): grpc.ClientUnaryCall;
    public delete(request: stub_extension_pb.DeleteExtensionRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: stub_extension_pb.DeleteExtensionResponse) => void): grpc.ClientUnaryCall;
}
