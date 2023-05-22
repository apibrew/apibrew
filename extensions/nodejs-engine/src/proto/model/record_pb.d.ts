// package: model
// file: model/record.proto

/* tslint:disable */
/* eslint-disable */

import * as jspb from "google-protobuf";
import * as google_protobuf_struct_pb from "google-protobuf/google/protobuf/struct_pb";
import * as model_audit_pb from "../model/audit_pb";
import * as model_common_pb from "../model/common_pb";

export class Record extends jspb.Message { 
    getId(): string;
    setId(value: string): Record;

    getPropertiesMap(): jspb.Map<string, google_protobuf_struct_pb.Value>;
    clearPropertiesMap(): void;
    clearPropertiespackedList(): void;
    getPropertiespackedList(): Array<google_protobuf_struct_pb.Value>;
    setPropertiespackedList(value: Array<google_protobuf_struct_pb.Value>): Record;
    addPropertiespacked(value?: google_protobuf_struct_pb.Value, index?: number): google_protobuf_struct_pb.Value;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): Record.AsObject;
    static toObject(includeInstance: boolean, msg: Record): Record.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: Record, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): Record;
    static deserializeBinaryFromReader(message: Record, reader: jspb.BinaryReader): Record;
}

export namespace Record {
    export type AsObject = {
        id: string,

        propertiesMap: Array<[string, google_protobuf_struct_pb.Value.AsObject]>,
        propertiespackedList: Array<google_protobuf_struct_pb.Value.AsObject>,
    }
}
