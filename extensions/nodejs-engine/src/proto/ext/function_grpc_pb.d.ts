// package: ext
// file: ext/function.proto

/* tslint:disable */
/* eslint-disable */

import * as grpc from "@grpc/grpc-js";
import * as ext_function_pb from "../ext/function_pb";
import * as model_record_pb from "../model/record_pb";
import * as model_resource_pb from "../model/resource_pb";
import * as model_query_pb from "../model/query_pb";
import * as model_error_pb from "../model/error_pb";
import * as google_protobuf_any_pb from "google-protobuf/google/protobuf/any_pb";
import * as google_protobuf_struct_pb from "google-protobuf/google/protobuf/struct_pb";
import * as model_event_pb from "../model/event_pb";

interface IFunctionService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
    functionCall: IFunctionService_IFunctionCall;
}

interface IFunctionService_IFunctionCall extends grpc.MethodDefinition<ext_function_pb.FunctionCallRequest, ext_function_pb.FunctionCallResponse> {
    path: "/ext.Function/FunctionCall";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<ext_function_pb.FunctionCallRequest>;
    requestDeserialize: grpc.deserialize<ext_function_pb.FunctionCallRequest>;
    responseSerialize: grpc.serialize<ext_function_pb.FunctionCallResponse>;
    responseDeserialize: grpc.deserialize<ext_function_pb.FunctionCallResponse>;
}

export const FunctionService: IFunctionService;

export interface IFunctionServer {
    functionCall: grpc.handleUnaryCall<ext_function_pb.FunctionCallRequest, ext_function_pb.FunctionCallResponse>;
}

export interface IFunctionClient {
    functionCall(request: ext_function_pb.FunctionCallRequest, callback: (error: grpc.ServiceError | null, response: ext_function_pb.FunctionCallResponse) => void): grpc.ClientUnaryCall;
    functionCall(request: ext_function_pb.FunctionCallRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: ext_function_pb.FunctionCallResponse) => void): grpc.ClientUnaryCall;
    functionCall(request: ext_function_pb.FunctionCallRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: ext_function_pb.FunctionCallResponse) => void): grpc.ClientUnaryCall;
}

export class FunctionClient extends grpc.Client implements IFunctionClient {
    constructor(address: string, credentials: grpc.ChannelCredentials, options?: object);
    public functionCall(request: ext_function_pb.FunctionCallRequest, callback: (error: grpc.ServiceError | null, response: ext_function_pb.FunctionCallResponse) => void): grpc.ClientUnaryCall;
    public functionCall(request: ext_function_pb.FunctionCallRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: ext_function_pb.FunctionCallResponse) => void): grpc.ClientUnaryCall;
    public functionCall(request: ext_function_pb.FunctionCallRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: ext_function_pb.FunctionCallResponse) => void): grpc.ClientUnaryCall;
}
