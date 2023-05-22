// package: model
// file: model/common.proto

/* tslint:disable */
/* eslint-disable */

import * as jspb from "google-protobuf";
import * as google_protobuf_any_pb from "google-protobuf/google/protobuf/any_pb";

export class MapAnyWrap extends jspb.Message { 

    getContentMap(): jspb.Map<string, google_protobuf_any_pb.Any>;
    clearContentMap(): void;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): MapAnyWrap.AsObject;
    static toObject(includeInstance: boolean, msg: MapAnyWrap): MapAnyWrap.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: MapAnyWrap, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): MapAnyWrap;
    static deserializeBinaryFromReader(message: MapAnyWrap, reader: jspb.BinaryReader): MapAnyWrap;
}

export namespace MapAnyWrap {
    export type AsObject = {

        contentMap: Array<[string, google_protobuf_any_pb.Any.AsObject]>,
    }
}
