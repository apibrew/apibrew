// package: stub
// file: stub/generic.proto

/* tslint:disable */
/* eslint-disable */

import * as grpc from "@grpc/grpc-js";
import * as stub_generic_pb from "../stub/generic_pb";
import * as model_query_pb from "../model/query_pb";
import * as model_error_pb from "../model/error_pb";
import * as google_protobuf_struct_pb from "google-protobuf/google/protobuf/struct_pb";
import * as google_protobuf_any_pb from "google-protobuf/google/protobuf/any_pb";

interface IGenericService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
    create: IGenericService_ICreate;
    update: IGenericService_IUpdate;
    updateMulti: IGenericService_IUpdateMulti;
    delete: IGenericService_IDelete;
    list: IGenericService_IList;
    search: IGenericService_ISearch;
    get: IGenericService_IGet;
}

interface IGenericService_ICreate extends grpc.MethodDefinition<stub_generic_pb.CreateRequest, stub_generic_pb.CreateResponse> {
    path: "/stub.Generic/Create";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<stub_generic_pb.CreateRequest>;
    requestDeserialize: grpc.deserialize<stub_generic_pb.CreateRequest>;
    responseSerialize: grpc.serialize<stub_generic_pb.CreateResponse>;
    responseDeserialize: grpc.deserialize<stub_generic_pb.CreateResponse>;
}
interface IGenericService_IUpdate extends grpc.MethodDefinition<stub_generic_pb.UpdateRequest, stub_generic_pb.UpdateResponse> {
    path: "/stub.Generic/Update";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<stub_generic_pb.UpdateRequest>;
    requestDeserialize: grpc.deserialize<stub_generic_pb.UpdateRequest>;
    responseSerialize: grpc.serialize<stub_generic_pb.UpdateResponse>;
    responseDeserialize: grpc.deserialize<stub_generic_pb.UpdateResponse>;
}
interface IGenericService_IUpdateMulti extends grpc.MethodDefinition<stub_generic_pb.UpdateMultiRequest, stub_generic_pb.UpdateMultiResponse> {
    path: "/stub.Generic/UpdateMulti";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<stub_generic_pb.UpdateMultiRequest>;
    requestDeserialize: grpc.deserialize<stub_generic_pb.UpdateMultiRequest>;
    responseSerialize: grpc.serialize<stub_generic_pb.UpdateMultiResponse>;
    responseDeserialize: grpc.deserialize<stub_generic_pb.UpdateMultiResponse>;
}
interface IGenericService_IDelete extends grpc.MethodDefinition<stub_generic_pb.DeleteRequest, stub_generic_pb.DeleteResponse> {
    path: "/stub.Generic/Delete";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<stub_generic_pb.DeleteRequest>;
    requestDeserialize: grpc.deserialize<stub_generic_pb.DeleteRequest>;
    responseSerialize: grpc.serialize<stub_generic_pb.DeleteResponse>;
    responseDeserialize: grpc.deserialize<stub_generic_pb.DeleteResponse>;
}
interface IGenericService_IList extends grpc.MethodDefinition<stub_generic_pb.ListRequest, stub_generic_pb.ListResponse> {
    path: "/stub.Generic/List";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<stub_generic_pb.ListRequest>;
    requestDeserialize: grpc.deserialize<stub_generic_pb.ListRequest>;
    responseSerialize: grpc.serialize<stub_generic_pb.ListResponse>;
    responseDeserialize: grpc.deserialize<stub_generic_pb.ListResponse>;
}
interface IGenericService_ISearch extends grpc.MethodDefinition<stub_generic_pb.SearchRequest, stub_generic_pb.SearchResponse> {
    path: "/stub.Generic/Search";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<stub_generic_pb.SearchRequest>;
    requestDeserialize: grpc.deserialize<stub_generic_pb.SearchRequest>;
    responseSerialize: grpc.serialize<stub_generic_pb.SearchResponse>;
    responseDeserialize: grpc.deserialize<stub_generic_pb.SearchResponse>;
}
interface IGenericService_IGet extends grpc.MethodDefinition<stub_generic_pb.GetRequest, stub_generic_pb.GetResponse> {
    path: "/stub.Generic/Get";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<stub_generic_pb.GetRequest>;
    requestDeserialize: grpc.deserialize<stub_generic_pb.GetRequest>;
    responseSerialize: grpc.serialize<stub_generic_pb.GetResponse>;
    responseDeserialize: grpc.deserialize<stub_generic_pb.GetResponse>;
}

export const GenericService: IGenericService;

export interface IGenericServer {
    create: grpc.handleUnaryCall<stub_generic_pb.CreateRequest, stub_generic_pb.CreateResponse>;
    update: grpc.handleUnaryCall<stub_generic_pb.UpdateRequest, stub_generic_pb.UpdateResponse>;
    updateMulti: grpc.handleUnaryCall<stub_generic_pb.UpdateMultiRequest, stub_generic_pb.UpdateMultiResponse>;
    delete: grpc.handleUnaryCall<stub_generic_pb.DeleteRequest, stub_generic_pb.DeleteResponse>;
    list: grpc.handleUnaryCall<stub_generic_pb.ListRequest, stub_generic_pb.ListResponse>;
    search: grpc.handleUnaryCall<stub_generic_pb.SearchRequest, stub_generic_pb.SearchResponse>;
    get: grpc.handleUnaryCall<stub_generic_pb.GetRequest, stub_generic_pb.GetResponse>;
}

export interface IGenericClient {
    create(request: stub_generic_pb.CreateRequest, callback: (error: grpc.ServiceError | null, response: stub_generic_pb.CreateResponse) => void): grpc.ClientUnaryCall;
    create(request: stub_generic_pb.CreateRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: stub_generic_pb.CreateResponse) => void): grpc.ClientUnaryCall;
    create(request: stub_generic_pb.CreateRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: stub_generic_pb.CreateResponse) => void): grpc.ClientUnaryCall;
    update(request: stub_generic_pb.UpdateRequest, callback: (error: grpc.ServiceError | null, response: stub_generic_pb.UpdateResponse) => void): grpc.ClientUnaryCall;
    update(request: stub_generic_pb.UpdateRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: stub_generic_pb.UpdateResponse) => void): grpc.ClientUnaryCall;
    update(request: stub_generic_pb.UpdateRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: stub_generic_pb.UpdateResponse) => void): grpc.ClientUnaryCall;
    updateMulti(request: stub_generic_pb.UpdateMultiRequest, callback: (error: grpc.ServiceError | null, response: stub_generic_pb.UpdateMultiResponse) => void): grpc.ClientUnaryCall;
    updateMulti(request: stub_generic_pb.UpdateMultiRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: stub_generic_pb.UpdateMultiResponse) => void): grpc.ClientUnaryCall;
    updateMulti(request: stub_generic_pb.UpdateMultiRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: stub_generic_pb.UpdateMultiResponse) => void): grpc.ClientUnaryCall;
    delete(request: stub_generic_pb.DeleteRequest, callback: (error: grpc.ServiceError | null, response: stub_generic_pb.DeleteResponse) => void): grpc.ClientUnaryCall;
    delete(request: stub_generic_pb.DeleteRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: stub_generic_pb.DeleteResponse) => void): grpc.ClientUnaryCall;
    delete(request: stub_generic_pb.DeleteRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: stub_generic_pb.DeleteResponse) => void): grpc.ClientUnaryCall;
    list(request: stub_generic_pb.ListRequest, callback: (error: grpc.ServiceError | null, response: stub_generic_pb.ListResponse) => void): grpc.ClientUnaryCall;
    list(request: stub_generic_pb.ListRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: stub_generic_pb.ListResponse) => void): grpc.ClientUnaryCall;
    list(request: stub_generic_pb.ListRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: stub_generic_pb.ListResponse) => void): grpc.ClientUnaryCall;
    search(request: stub_generic_pb.SearchRequest, callback: (error: grpc.ServiceError | null, response: stub_generic_pb.SearchResponse) => void): grpc.ClientUnaryCall;
    search(request: stub_generic_pb.SearchRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: stub_generic_pb.SearchResponse) => void): grpc.ClientUnaryCall;
    search(request: stub_generic_pb.SearchRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: stub_generic_pb.SearchResponse) => void): grpc.ClientUnaryCall;
    get(request: stub_generic_pb.GetRequest, callback: (error: grpc.ServiceError | null, response: stub_generic_pb.GetResponse) => void): grpc.ClientUnaryCall;
    get(request: stub_generic_pb.GetRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: stub_generic_pb.GetResponse) => void): grpc.ClientUnaryCall;
    get(request: stub_generic_pb.GetRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: stub_generic_pb.GetResponse) => void): grpc.ClientUnaryCall;
}

export class GenericClient extends grpc.Client implements IGenericClient {
    constructor(address: string, credentials: grpc.ChannelCredentials, options?: object);
    public create(request: stub_generic_pb.CreateRequest, callback: (error: grpc.ServiceError | null, response: stub_generic_pb.CreateResponse) => void): grpc.ClientUnaryCall;
    public create(request: stub_generic_pb.CreateRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: stub_generic_pb.CreateResponse) => void): grpc.ClientUnaryCall;
    public create(request: stub_generic_pb.CreateRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: stub_generic_pb.CreateResponse) => void): grpc.ClientUnaryCall;
    public update(request: stub_generic_pb.UpdateRequest, callback: (error: grpc.ServiceError | null, response: stub_generic_pb.UpdateResponse) => void): grpc.ClientUnaryCall;
    public update(request: stub_generic_pb.UpdateRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: stub_generic_pb.UpdateResponse) => void): grpc.ClientUnaryCall;
    public update(request: stub_generic_pb.UpdateRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: stub_generic_pb.UpdateResponse) => void): grpc.ClientUnaryCall;
    public updateMulti(request: stub_generic_pb.UpdateMultiRequest, callback: (error: grpc.ServiceError | null, response: stub_generic_pb.UpdateMultiResponse) => void): grpc.ClientUnaryCall;
    public updateMulti(request: stub_generic_pb.UpdateMultiRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: stub_generic_pb.UpdateMultiResponse) => void): grpc.ClientUnaryCall;
    public updateMulti(request: stub_generic_pb.UpdateMultiRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: stub_generic_pb.UpdateMultiResponse) => void): grpc.ClientUnaryCall;
    public delete(request: stub_generic_pb.DeleteRequest, callback: (error: grpc.ServiceError | null, response: stub_generic_pb.DeleteResponse) => void): grpc.ClientUnaryCall;
    public delete(request: stub_generic_pb.DeleteRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: stub_generic_pb.DeleteResponse) => void): grpc.ClientUnaryCall;
    public delete(request: stub_generic_pb.DeleteRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: stub_generic_pb.DeleteResponse) => void): grpc.ClientUnaryCall;
    public list(request: stub_generic_pb.ListRequest, callback: (error: grpc.ServiceError | null, response: stub_generic_pb.ListResponse) => void): grpc.ClientUnaryCall;
    public list(request: stub_generic_pb.ListRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: stub_generic_pb.ListResponse) => void): grpc.ClientUnaryCall;
    public list(request: stub_generic_pb.ListRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: stub_generic_pb.ListResponse) => void): grpc.ClientUnaryCall;
    public search(request: stub_generic_pb.SearchRequest, callback: (error: grpc.ServiceError | null, response: stub_generic_pb.SearchResponse) => void): grpc.ClientUnaryCall;
    public search(request: stub_generic_pb.SearchRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: stub_generic_pb.SearchResponse) => void): grpc.ClientUnaryCall;
    public search(request: stub_generic_pb.SearchRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: stub_generic_pb.SearchResponse) => void): grpc.ClientUnaryCall;
    public get(request: stub_generic_pb.GetRequest, callback: (error: grpc.ServiceError | null, response: stub_generic_pb.GetResponse) => void): grpc.ClientUnaryCall;
    public get(request: stub_generic_pb.GetRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: stub_generic_pb.GetResponse) => void): grpc.ClientUnaryCall;
    public get(request: stub_generic_pb.GetRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: stub_generic_pb.GetResponse) => void): grpc.ClientUnaryCall;
}
