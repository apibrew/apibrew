// package: ext
// file: ext/function.proto

/* tslint:disable */
/* eslint-disable */

import * as jspb from "google-protobuf";
import * as model_record_pb from "../model/record_pb";
import * as model_resource_pb from "../model/resource_pb";
import * as model_query_pb from "../model/query_pb";
import * as model_error_pb from "../model/error_pb";
import * as google_protobuf_any_pb from "google-protobuf/google/protobuf/any_pb";
import * as google_protobuf_struct_pb from "google-protobuf/google/protobuf/struct_pb";
import * as model_event_pb from "../model/event_pb";

export class FunctionCallRequest extends jspb.Message { 
    getName(): string;
    setName(value: string): FunctionCallRequest;

    hasEvent(): boolean;
    clearEvent(): void;
    getEvent(): model_event_pb.Event | undefined;
    setEvent(value?: model_event_pb.Event): FunctionCallRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): FunctionCallRequest.AsObject;
    static toObject(includeInstance: boolean, msg: FunctionCallRequest): FunctionCallRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: FunctionCallRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): FunctionCallRequest;
    static deserializeBinaryFromReader(message: FunctionCallRequest, reader: jspb.BinaryReader): FunctionCallRequest;
}

export namespace FunctionCallRequest {
    export type AsObject = {
        name: string,
        event?: model_event_pb.Event.AsObject,
    }
}

export class FunctionCallResponse extends jspb.Message { 

    hasEvent(): boolean;
    clearEvent(): void;
    getEvent(): model_event_pb.Event | undefined;
    setEvent(value?: model_event_pb.Event): FunctionCallResponse;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): FunctionCallResponse.AsObject;
    static toObject(includeInstance: boolean, msg: FunctionCallResponse): FunctionCallResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: FunctionCallResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): FunctionCallResponse;
    static deserializeBinaryFromReader(message: FunctionCallResponse, reader: jspb.BinaryReader): FunctionCallResponse;
}

export namespace FunctionCallResponse {
    export type AsObject = {
        event?: model_event_pb.Event.AsObject,
    }
}
