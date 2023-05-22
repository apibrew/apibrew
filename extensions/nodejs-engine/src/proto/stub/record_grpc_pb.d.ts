// package: stub
// file: stub/record.proto

/* tslint:disable */
/* eslint-disable */

import * as grpc from "@grpc/grpc-js";
import * as stub_record_pb from "../stub/record_pb";
import * as model_record_pb from "../model/record_pb";
import * as model_query_pb from "../model/query_pb";
import * as model_error_pb from "../model/error_pb";
import * as google_protobuf_struct_pb from "google-protobuf/google/protobuf/struct_pb";

interface IRecordService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
    create: IRecordService_ICreate;
    update: IRecordService_IUpdate;
    apply: IRecordService_IApply;
    updateMulti: IRecordService_IUpdateMulti;
    delete: IRecordService_IDelete;
    list: IRecordService_IList;
    search: IRecordService_ISearch;
    readStream: IRecordService_IReadStream;
    writeStream: IRecordService_IWriteStream;
    get: IRecordService_IGet;
}

interface IRecordService_ICreate extends grpc.MethodDefinition<stub_record_pb.CreateRecordRequest, stub_record_pb.CreateRecordResponse> {
    path: "/stub.Record/Create";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<stub_record_pb.CreateRecordRequest>;
    requestDeserialize: grpc.deserialize<stub_record_pb.CreateRecordRequest>;
    responseSerialize: grpc.serialize<stub_record_pb.CreateRecordResponse>;
    responseDeserialize: grpc.deserialize<stub_record_pb.CreateRecordResponse>;
}
interface IRecordService_IUpdate extends grpc.MethodDefinition<stub_record_pb.UpdateRecordRequest, stub_record_pb.UpdateRecordResponse> {
    path: "/stub.Record/Update";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<stub_record_pb.UpdateRecordRequest>;
    requestDeserialize: grpc.deserialize<stub_record_pb.UpdateRecordRequest>;
    responseSerialize: grpc.serialize<stub_record_pb.UpdateRecordResponse>;
    responseDeserialize: grpc.deserialize<stub_record_pb.UpdateRecordResponse>;
}
interface IRecordService_IApply extends grpc.MethodDefinition<stub_record_pb.ApplyRecordRequest, stub_record_pb.ApplyRecordResponse> {
    path: "/stub.Record/Apply";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<stub_record_pb.ApplyRecordRequest>;
    requestDeserialize: grpc.deserialize<stub_record_pb.ApplyRecordRequest>;
    responseSerialize: grpc.serialize<stub_record_pb.ApplyRecordResponse>;
    responseDeserialize: grpc.deserialize<stub_record_pb.ApplyRecordResponse>;
}
interface IRecordService_IUpdateMulti extends grpc.MethodDefinition<stub_record_pb.UpdateMultiRecordRequest, stub_record_pb.UpdateMultiRecordResponse> {
    path: "/stub.Record/UpdateMulti";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<stub_record_pb.UpdateMultiRecordRequest>;
    requestDeserialize: grpc.deserialize<stub_record_pb.UpdateMultiRecordRequest>;
    responseSerialize: grpc.serialize<stub_record_pb.UpdateMultiRecordResponse>;
    responseDeserialize: grpc.deserialize<stub_record_pb.UpdateMultiRecordResponse>;
}
interface IRecordService_IDelete extends grpc.MethodDefinition<stub_record_pb.DeleteRecordRequest, stub_record_pb.DeleteRecordResponse> {
    path: "/stub.Record/Delete";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<stub_record_pb.DeleteRecordRequest>;
    requestDeserialize: grpc.deserialize<stub_record_pb.DeleteRecordRequest>;
    responseSerialize: grpc.serialize<stub_record_pb.DeleteRecordResponse>;
    responseDeserialize: grpc.deserialize<stub_record_pb.DeleteRecordResponse>;
}
interface IRecordService_IList extends grpc.MethodDefinition<stub_record_pb.ListRecordRequest, stub_record_pb.ListRecordResponse> {
    path: "/stub.Record/List";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<stub_record_pb.ListRecordRequest>;
    requestDeserialize: grpc.deserialize<stub_record_pb.ListRecordRequest>;
    responseSerialize: grpc.serialize<stub_record_pb.ListRecordResponse>;
    responseDeserialize: grpc.deserialize<stub_record_pb.ListRecordResponse>;
}
interface IRecordService_ISearch extends grpc.MethodDefinition<stub_record_pb.SearchRecordRequest, stub_record_pb.SearchRecordResponse> {
    path: "/stub.Record/Search";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<stub_record_pb.SearchRecordRequest>;
    requestDeserialize: grpc.deserialize<stub_record_pb.SearchRecordRequest>;
    responseSerialize: grpc.serialize<stub_record_pb.SearchRecordResponse>;
    responseDeserialize: grpc.deserialize<stub_record_pb.SearchRecordResponse>;
}
interface IRecordService_IReadStream extends grpc.MethodDefinition<stub_record_pb.ReadStreamRequest, model_record_pb.Record> {
    path: "/stub.Record/ReadStream";
    requestStream: false;
    responseStream: true;
    requestSerialize: grpc.serialize<stub_record_pb.ReadStreamRequest>;
    requestDeserialize: grpc.deserialize<stub_record_pb.ReadStreamRequest>;
    responseSerialize: grpc.serialize<model_record_pb.Record>;
    responseDeserialize: grpc.deserialize<model_record_pb.Record>;
}
interface IRecordService_IWriteStream extends grpc.MethodDefinition<model_record_pb.Record, stub_record_pb.WriteStreamResponse> {
    path: "/stub.Record/WriteStream";
    requestStream: true;
    responseStream: false;
    requestSerialize: grpc.serialize<model_record_pb.Record>;
    requestDeserialize: grpc.deserialize<model_record_pb.Record>;
    responseSerialize: grpc.serialize<stub_record_pb.WriteStreamResponse>;
    responseDeserialize: grpc.deserialize<stub_record_pb.WriteStreamResponse>;
}
interface IRecordService_IGet extends grpc.MethodDefinition<stub_record_pb.GetRecordRequest, stub_record_pb.GetRecordResponse> {
    path: "/stub.Record/Get";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<stub_record_pb.GetRecordRequest>;
    requestDeserialize: grpc.deserialize<stub_record_pb.GetRecordRequest>;
    responseSerialize: grpc.serialize<stub_record_pb.GetRecordResponse>;
    responseDeserialize: grpc.deserialize<stub_record_pb.GetRecordResponse>;
}

export const RecordService: IRecordService;

export interface IRecordServer {
    create: grpc.handleUnaryCall<stub_record_pb.CreateRecordRequest, stub_record_pb.CreateRecordResponse>;
    update: grpc.handleUnaryCall<stub_record_pb.UpdateRecordRequest, stub_record_pb.UpdateRecordResponse>;
    apply: grpc.handleUnaryCall<stub_record_pb.ApplyRecordRequest, stub_record_pb.ApplyRecordResponse>;
    updateMulti: grpc.handleUnaryCall<stub_record_pb.UpdateMultiRecordRequest, stub_record_pb.UpdateMultiRecordResponse>;
    delete: grpc.handleUnaryCall<stub_record_pb.DeleteRecordRequest, stub_record_pb.DeleteRecordResponse>;
    list: grpc.handleUnaryCall<stub_record_pb.ListRecordRequest, stub_record_pb.ListRecordResponse>;
    search: grpc.handleUnaryCall<stub_record_pb.SearchRecordRequest, stub_record_pb.SearchRecordResponse>;
    readStream: grpc.handleServerStreamingCall<stub_record_pb.ReadStreamRequest, model_record_pb.Record>;
    writeStream: grpc.handleClientStreamingCall<model_record_pb.Record, stub_record_pb.WriteStreamResponse>;
    get: grpc.handleUnaryCall<stub_record_pb.GetRecordRequest, stub_record_pb.GetRecordResponse>;
}

export interface IRecordClient {
    create(request: stub_record_pb.CreateRecordRequest, callback: (error: grpc.ServiceError | null, response: stub_record_pb.CreateRecordResponse) => void): grpc.ClientUnaryCall;
    create(request: stub_record_pb.CreateRecordRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: stub_record_pb.CreateRecordResponse) => void): grpc.ClientUnaryCall;
    create(request: stub_record_pb.CreateRecordRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: stub_record_pb.CreateRecordResponse) => void): grpc.ClientUnaryCall;
    update(request: stub_record_pb.UpdateRecordRequest, callback: (error: grpc.ServiceError | null, response: stub_record_pb.UpdateRecordResponse) => void): grpc.ClientUnaryCall;
    update(request: stub_record_pb.UpdateRecordRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: stub_record_pb.UpdateRecordResponse) => void): grpc.ClientUnaryCall;
    update(request: stub_record_pb.UpdateRecordRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: stub_record_pb.UpdateRecordResponse) => void): grpc.ClientUnaryCall;
    apply(request: stub_record_pb.ApplyRecordRequest, callback: (error: grpc.ServiceError | null, response: stub_record_pb.ApplyRecordResponse) => void): grpc.ClientUnaryCall;
    apply(request: stub_record_pb.ApplyRecordRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: stub_record_pb.ApplyRecordResponse) => void): grpc.ClientUnaryCall;
    apply(request: stub_record_pb.ApplyRecordRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: stub_record_pb.ApplyRecordResponse) => void): grpc.ClientUnaryCall;
    updateMulti(request: stub_record_pb.UpdateMultiRecordRequest, callback: (error: grpc.ServiceError | null, response: stub_record_pb.UpdateMultiRecordResponse) => void): grpc.ClientUnaryCall;
    updateMulti(request: stub_record_pb.UpdateMultiRecordRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: stub_record_pb.UpdateMultiRecordResponse) => void): grpc.ClientUnaryCall;
    updateMulti(request: stub_record_pb.UpdateMultiRecordRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: stub_record_pb.UpdateMultiRecordResponse) => void): grpc.ClientUnaryCall;
    delete(request: stub_record_pb.DeleteRecordRequest, callback: (error: grpc.ServiceError | null, response: stub_record_pb.DeleteRecordResponse) => void): grpc.ClientUnaryCall;
    delete(request: stub_record_pb.DeleteRecordRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: stub_record_pb.DeleteRecordResponse) => void): grpc.ClientUnaryCall;
    delete(request: stub_record_pb.DeleteRecordRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: stub_record_pb.DeleteRecordResponse) => void): grpc.ClientUnaryCall;
    list(request: stub_record_pb.ListRecordRequest, callback: (error: grpc.ServiceError | null, response: stub_record_pb.ListRecordResponse) => void): grpc.ClientUnaryCall;
    list(request: stub_record_pb.ListRecordRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: stub_record_pb.ListRecordResponse) => void): grpc.ClientUnaryCall;
    list(request: stub_record_pb.ListRecordRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: stub_record_pb.ListRecordResponse) => void): grpc.ClientUnaryCall;
    search(request: stub_record_pb.SearchRecordRequest, callback: (error: grpc.ServiceError | null, response: stub_record_pb.SearchRecordResponse) => void): grpc.ClientUnaryCall;
    search(request: stub_record_pb.SearchRecordRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: stub_record_pb.SearchRecordResponse) => void): grpc.ClientUnaryCall;
    search(request: stub_record_pb.SearchRecordRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: stub_record_pb.SearchRecordResponse) => void): grpc.ClientUnaryCall;
    readStream(request: stub_record_pb.ReadStreamRequest, options?: Partial<grpc.CallOptions>): grpc.ClientReadableStream<model_record_pb.Record>;
    readStream(request: stub_record_pb.ReadStreamRequest, metadata?: grpc.Metadata, options?: Partial<grpc.CallOptions>): grpc.ClientReadableStream<model_record_pb.Record>;
    writeStream(callback: (error: grpc.ServiceError | null, response: stub_record_pb.WriteStreamResponse) => void): grpc.ClientWritableStream<model_record_pb.Record>;
    writeStream(metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: stub_record_pb.WriteStreamResponse) => void): grpc.ClientWritableStream<model_record_pb.Record>;
    writeStream(options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: stub_record_pb.WriteStreamResponse) => void): grpc.ClientWritableStream<model_record_pb.Record>;
    writeStream(metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: stub_record_pb.WriteStreamResponse) => void): grpc.ClientWritableStream<model_record_pb.Record>;
    get(request: stub_record_pb.GetRecordRequest, callback: (error: grpc.ServiceError | null, response: stub_record_pb.GetRecordResponse) => void): grpc.ClientUnaryCall;
    get(request: stub_record_pb.GetRecordRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: stub_record_pb.GetRecordResponse) => void): grpc.ClientUnaryCall;
    get(request: stub_record_pb.GetRecordRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: stub_record_pb.GetRecordResponse) => void): grpc.ClientUnaryCall;
}

export class RecordClient extends grpc.Client implements IRecordClient {
    constructor(address: string, credentials: grpc.ChannelCredentials, options?: object);
    public create(request: stub_record_pb.CreateRecordRequest, callback: (error: grpc.ServiceError | null, response: stub_record_pb.CreateRecordResponse) => void): grpc.ClientUnaryCall;
    public create(request: stub_record_pb.CreateRecordRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: stub_record_pb.CreateRecordResponse) => void): grpc.ClientUnaryCall;
    public create(request: stub_record_pb.CreateRecordRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: stub_record_pb.CreateRecordResponse) => void): grpc.ClientUnaryCall;
    public update(request: stub_record_pb.UpdateRecordRequest, callback: (error: grpc.ServiceError | null, response: stub_record_pb.UpdateRecordResponse) => void): grpc.ClientUnaryCall;
    public update(request: stub_record_pb.UpdateRecordRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: stub_record_pb.UpdateRecordResponse) => void): grpc.ClientUnaryCall;
    public update(request: stub_record_pb.UpdateRecordRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: stub_record_pb.UpdateRecordResponse) => void): grpc.ClientUnaryCall;
    public apply(request: stub_record_pb.ApplyRecordRequest, callback: (error: grpc.ServiceError | null, response: stub_record_pb.ApplyRecordResponse) => void): grpc.ClientUnaryCall;
    public apply(request: stub_record_pb.ApplyRecordRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: stub_record_pb.ApplyRecordResponse) => void): grpc.ClientUnaryCall;
    public apply(request: stub_record_pb.ApplyRecordRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: stub_record_pb.ApplyRecordResponse) => void): grpc.ClientUnaryCall;
    public updateMulti(request: stub_record_pb.UpdateMultiRecordRequest, callback: (error: grpc.ServiceError | null, response: stub_record_pb.UpdateMultiRecordResponse) => void): grpc.ClientUnaryCall;
    public updateMulti(request: stub_record_pb.UpdateMultiRecordRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: stub_record_pb.UpdateMultiRecordResponse) => void): grpc.ClientUnaryCall;
    public updateMulti(request: stub_record_pb.UpdateMultiRecordRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: stub_record_pb.UpdateMultiRecordResponse) => void): grpc.ClientUnaryCall;
    public delete(request: stub_record_pb.DeleteRecordRequest, callback: (error: grpc.ServiceError | null, response: stub_record_pb.DeleteRecordResponse) => void): grpc.ClientUnaryCall;
    public delete(request: stub_record_pb.DeleteRecordRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: stub_record_pb.DeleteRecordResponse) => void): grpc.ClientUnaryCall;
    public delete(request: stub_record_pb.DeleteRecordRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: stub_record_pb.DeleteRecordResponse) => void): grpc.ClientUnaryCall;
    public list(request: stub_record_pb.ListRecordRequest, callback: (error: grpc.ServiceError | null, response: stub_record_pb.ListRecordResponse) => void): grpc.ClientUnaryCall;
    public list(request: stub_record_pb.ListRecordRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: stub_record_pb.ListRecordResponse) => void): grpc.ClientUnaryCall;
    public list(request: stub_record_pb.ListRecordRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: stub_record_pb.ListRecordResponse) => void): grpc.ClientUnaryCall;
    public search(request: stub_record_pb.SearchRecordRequest, callback: (error: grpc.ServiceError | null, response: stub_record_pb.SearchRecordResponse) => void): grpc.ClientUnaryCall;
    public search(request: stub_record_pb.SearchRecordRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: stub_record_pb.SearchRecordResponse) => void): grpc.ClientUnaryCall;
    public search(request: stub_record_pb.SearchRecordRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: stub_record_pb.SearchRecordResponse) => void): grpc.ClientUnaryCall;
    public readStream(request: stub_record_pb.ReadStreamRequest, options?: Partial<grpc.CallOptions>): grpc.ClientReadableStream<model_record_pb.Record>;
    public readStream(request: stub_record_pb.ReadStreamRequest, metadata?: grpc.Metadata, options?: Partial<grpc.CallOptions>): grpc.ClientReadableStream<model_record_pb.Record>;
    public writeStream(callback: (error: grpc.ServiceError | null, response: stub_record_pb.WriteStreamResponse) => void): grpc.ClientWritableStream<model_record_pb.Record>;
    public writeStream(metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: stub_record_pb.WriteStreamResponse) => void): grpc.ClientWritableStream<model_record_pb.Record>;
    public writeStream(options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: stub_record_pb.WriteStreamResponse) => void): grpc.ClientWritableStream<model_record_pb.Record>;
    public writeStream(metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: stub_record_pb.WriteStreamResponse) => void): grpc.ClientWritableStream<model_record_pb.Record>;
    public get(request: stub_record_pb.GetRecordRequest, callback: (error: grpc.ServiceError | null, response: stub_record_pb.GetRecordResponse) => void): grpc.ClientUnaryCall;
    public get(request: stub_record_pb.GetRecordRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: stub_record_pb.GetRecordResponse) => void): grpc.ClientUnaryCall;
    public get(request: stub_record_pb.GetRecordRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: stub_record_pb.GetRecordResponse) => void): grpc.ClientUnaryCall;
}
