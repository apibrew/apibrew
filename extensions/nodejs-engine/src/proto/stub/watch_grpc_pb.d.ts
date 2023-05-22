// package: stub
// file: stub/watch.proto

/* tslint:disable */
/* eslint-disable */

import * as grpc from "@grpc/grpc-js";
import * as stub_watch_pb from "../stub/watch_pb";
import * as model_event_pb from "../model/event_pb";
import * as model_query_pb from "../model/query_pb";


interface IWatchService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
    watch: IWatchService_IWatch;
}

interface IWatchService_IWatch extends grpc.MethodDefinition<stub_watch_pb.WatchRequest, model_event_pb.Event> {
    path: "/stub.Watch/Watch";
    requestStream: false;
    responseStream: true;
    requestSerialize: grpc.serialize<stub_watch_pb.WatchRequest>;
    requestDeserialize: grpc.deserialize<stub_watch_pb.WatchRequest>;
    responseSerialize: grpc.serialize<model_event_pb.Event>;
    responseDeserialize: grpc.deserialize<model_event_pb.Event>;
}

export const WatchService: IWatchService;

export interface IWatchServer {
    watch: grpc.handleServerStreamingCall<stub_watch_pb.WatchRequest, model_event_pb.Event>;
}

export interface IWatchClient {
    watch(request: stub_watch_pb.WatchRequest, options?: Partial<grpc.CallOptions>): grpc.ClientReadableStream<model_event_pb.Event>;
    watch(request: stub_watch_pb.WatchRequest, metadata?: grpc.Metadata, options?: Partial<grpc.CallOptions>): grpc.ClientReadableStream<model_event_pb.Event>;
}

export class WatchClient extends grpc.Client implements IWatchClient {
    constructor(address: string, credentials: grpc.ChannelCredentials, options?: object);
    public watch(request: stub_watch_pb.WatchRequest, options?: Partial<grpc.CallOptions>): grpc.ClientReadableStream<model_event_pb.Event>;
    public watch(request: stub_watch_pb.WatchRequest, metadata?: grpc.Metadata, options?: Partial<grpc.CallOptions>): grpc.ClientReadableStream<model_event_pb.Event>;
}
