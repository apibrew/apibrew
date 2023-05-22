// package: rest
// file: stub/rest/record.proto

/* tslint:disable */
/* eslint-disable */

import * as grpc from "@grpc/grpc-js";
import * as stub_rest_record_pb from "../../stub/rest/record_pb";
import * as model_record_pb from "../../model/record_pb";
import * as model_query_pb from "../../model/query_pb";
import * as model_error_pb from "../../model/error_pb";
import * as google_protobuf_struct_pb from "google-protobuf/google/protobuf/struct_pb";

interface IRecordService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
    create: IRecordService_ICreate;
    update: IRecordService_IUpdate;
    apply: IRecordService_IApply;
    delete: IRecordService_IDelete;
}

interface IRecordService_ICreate extends grpc.MethodDefinition<stub_rest_record_pb.CreateRecordRequest, stub_rest_record_pb.CreateRecordResponse> {
    path: "/rest.Record/Create";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<stub_rest_record_pb.CreateRecordRequest>;
    requestDeserialize: grpc.deserialize<stub_rest_record_pb.CreateRecordRequest>;
    responseSerialize: grpc.serialize<stub_rest_record_pb.CreateRecordResponse>;
    responseDeserialize: grpc.deserialize<stub_rest_record_pb.CreateRecordResponse>;
}
interface IRecordService_IUpdate extends grpc.MethodDefinition<stub_rest_record_pb.UpdateRecordRequest, stub_rest_record_pb.UpdateRecordResponse> {
    path: "/rest.Record/Update";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<stub_rest_record_pb.UpdateRecordRequest>;
    requestDeserialize: grpc.deserialize<stub_rest_record_pb.UpdateRecordRequest>;
    responseSerialize: grpc.serialize<stub_rest_record_pb.UpdateRecordResponse>;
    responseDeserialize: grpc.deserialize<stub_rest_record_pb.UpdateRecordResponse>;
}
interface IRecordService_IApply extends grpc.MethodDefinition<stub_rest_record_pb.ApplyRecordRequest, stub_rest_record_pb.ApplyRecordResponse> {
    path: "/rest.Record/Apply";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<stub_rest_record_pb.ApplyRecordRequest>;
    requestDeserialize: grpc.deserialize<stub_rest_record_pb.ApplyRecordRequest>;
    responseSerialize: grpc.serialize<stub_rest_record_pb.ApplyRecordResponse>;
    responseDeserialize: grpc.deserialize<stub_rest_record_pb.ApplyRecordResponse>;
}
interface IRecordService_IDelete extends grpc.MethodDefinition<stub_rest_record_pb.DeleteRecordRequest, stub_rest_record_pb.DeleteRecordResponse> {
    path: "/rest.Record/Delete";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<stub_rest_record_pb.DeleteRecordRequest>;
    requestDeserialize: grpc.deserialize<stub_rest_record_pb.DeleteRecordRequest>;
    responseSerialize: grpc.serialize<stub_rest_record_pb.DeleteRecordResponse>;
    responseDeserialize: grpc.deserialize<stub_rest_record_pb.DeleteRecordResponse>;
}

export const RecordService: IRecordService;

export interface IRecordServer {
    create: grpc.handleUnaryCall<stub_rest_record_pb.CreateRecordRequest, stub_rest_record_pb.CreateRecordResponse>;
    update: grpc.handleUnaryCall<stub_rest_record_pb.UpdateRecordRequest, stub_rest_record_pb.UpdateRecordResponse>;
    apply: grpc.handleUnaryCall<stub_rest_record_pb.ApplyRecordRequest, stub_rest_record_pb.ApplyRecordResponse>;
    delete: grpc.handleUnaryCall<stub_rest_record_pb.DeleteRecordRequest, stub_rest_record_pb.DeleteRecordResponse>;
}

export interface IRecordClient {
    create(request: stub_rest_record_pb.CreateRecordRequest, callback: (error: grpc.ServiceError | null, response: stub_rest_record_pb.CreateRecordResponse) => void): grpc.ClientUnaryCall;
    create(request: stub_rest_record_pb.CreateRecordRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: stub_rest_record_pb.CreateRecordResponse) => void): grpc.ClientUnaryCall;
    create(request: stub_rest_record_pb.CreateRecordRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: stub_rest_record_pb.CreateRecordResponse) => void): grpc.ClientUnaryCall;
    update(request: stub_rest_record_pb.UpdateRecordRequest, callback: (error: grpc.ServiceError | null, response: stub_rest_record_pb.UpdateRecordResponse) => void): grpc.ClientUnaryCall;
    update(request: stub_rest_record_pb.UpdateRecordRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: stub_rest_record_pb.UpdateRecordResponse) => void): grpc.ClientUnaryCall;
    update(request: stub_rest_record_pb.UpdateRecordRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: stub_rest_record_pb.UpdateRecordResponse) => void): grpc.ClientUnaryCall;
    apply(request: stub_rest_record_pb.ApplyRecordRequest, callback: (error: grpc.ServiceError | null, response: stub_rest_record_pb.ApplyRecordResponse) => void): grpc.ClientUnaryCall;
    apply(request: stub_rest_record_pb.ApplyRecordRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: stub_rest_record_pb.ApplyRecordResponse) => void): grpc.ClientUnaryCall;
    apply(request: stub_rest_record_pb.ApplyRecordRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: stub_rest_record_pb.ApplyRecordResponse) => void): grpc.ClientUnaryCall;
    delete(request: stub_rest_record_pb.DeleteRecordRequest, callback: (error: grpc.ServiceError | null, response: stub_rest_record_pb.DeleteRecordResponse) => void): grpc.ClientUnaryCall;
    delete(request: stub_rest_record_pb.DeleteRecordRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: stub_rest_record_pb.DeleteRecordResponse) => void): grpc.ClientUnaryCall;
    delete(request: stub_rest_record_pb.DeleteRecordRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: stub_rest_record_pb.DeleteRecordResponse) => void): grpc.ClientUnaryCall;
}

export class RecordClient extends grpc.Client implements IRecordClient {
    constructor(address: string, credentials: grpc.ChannelCredentials, options?: object);
    public create(request: stub_rest_record_pb.CreateRecordRequest, callback: (error: grpc.ServiceError | null, response: stub_rest_record_pb.CreateRecordResponse) => void): grpc.ClientUnaryCall;
    public create(request: stub_rest_record_pb.CreateRecordRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: stub_rest_record_pb.CreateRecordResponse) => void): grpc.ClientUnaryCall;
    public create(request: stub_rest_record_pb.CreateRecordRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: stub_rest_record_pb.CreateRecordResponse) => void): grpc.ClientUnaryCall;
    public update(request: stub_rest_record_pb.UpdateRecordRequest, callback: (error: grpc.ServiceError | null, response: stub_rest_record_pb.UpdateRecordResponse) => void): grpc.ClientUnaryCall;
    public update(request: stub_rest_record_pb.UpdateRecordRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: stub_rest_record_pb.UpdateRecordResponse) => void): grpc.ClientUnaryCall;
    public update(request: stub_rest_record_pb.UpdateRecordRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: stub_rest_record_pb.UpdateRecordResponse) => void): grpc.ClientUnaryCall;
    public apply(request: stub_rest_record_pb.ApplyRecordRequest, callback: (error: grpc.ServiceError | null, response: stub_rest_record_pb.ApplyRecordResponse) => void): grpc.ClientUnaryCall;
    public apply(request: stub_rest_record_pb.ApplyRecordRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: stub_rest_record_pb.ApplyRecordResponse) => void): grpc.ClientUnaryCall;
    public apply(request: stub_rest_record_pb.ApplyRecordRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: stub_rest_record_pb.ApplyRecordResponse) => void): grpc.ClientUnaryCall;
    public delete(request: stub_rest_record_pb.DeleteRecordRequest, callback: (error: grpc.ServiceError | null, response: stub_rest_record_pb.DeleteRecordResponse) => void): grpc.ClientUnaryCall;
    public delete(request: stub_rest_record_pb.DeleteRecordRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: stub_rest_record_pb.DeleteRecordResponse) => void): grpc.ClientUnaryCall;
    public delete(request: stub_rest_record_pb.DeleteRecordRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: stub_rest_record_pb.DeleteRecordResponse) => void): grpc.ClientUnaryCall;
}
