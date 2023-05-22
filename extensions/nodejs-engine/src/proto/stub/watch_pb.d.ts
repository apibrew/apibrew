// package: stub
// file: stub/watch.proto

/* tslint:disable */
/* eslint-disable */

import * as jspb from "google-protobuf";
import * as model_event_pb from "../model/event_pb";
import * as model_query_pb from "../model/query_pb";


export class WatchRequest extends jspb.Message { 
    getToken(): string;
    setToken(value: string): WatchRequest;

    hasSelector(): boolean;
    clearSelector(): void;
    getSelector(): model_event_pb.EventSelector | undefined;
    setSelector(value?: model_event_pb.EventSelector): WatchRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): WatchRequest.AsObject;
    static toObject(includeInstance: boolean, msg: WatchRequest): WatchRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: WatchRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): WatchRequest;
    static deserializeBinaryFromReader(message: WatchRequest, reader: jspb.BinaryReader): WatchRequest;
}

export namespace WatchRequest {
    export type AsObject = {
        token: string,
        selector?: model_event_pb.EventSelector.AsObject,
    }
}
