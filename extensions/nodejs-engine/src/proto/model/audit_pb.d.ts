// package: model
// file: model/audit.proto

/* tslint:disable */
/* eslint-disable */

import * as jspb from "google-protobuf";
import * as google_protobuf_timestamp_pb from "google-protobuf/google/protobuf/timestamp_pb";

export class AuditData extends jspb.Message { 

    hasCreatedOn(): boolean;
    clearCreatedOn(): void;
    getCreatedOn(): google_protobuf_timestamp_pb.Timestamp | undefined;
    setCreatedOn(value?: google_protobuf_timestamp_pb.Timestamp): AuditData;

    hasUpdatedOn(): boolean;
    clearUpdatedOn(): void;
    getUpdatedOn(): google_protobuf_timestamp_pb.Timestamp | undefined;
    setUpdatedOn(value?: google_protobuf_timestamp_pb.Timestamp): AuditData;
    getCreatedBy(): string;
    setCreatedBy(value: string): AuditData;
    getUpdatedBy(): string;
    setUpdatedBy(value: string): AuditData;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): AuditData.AsObject;
    static toObject(includeInstance: boolean, msg: AuditData): AuditData.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: AuditData, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): AuditData;
    static deserializeBinaryFromReader(message: AuditData, reader: jspb.BinaryReader): AuditData;
}

export namespace AuditData {
    export type AsObject = {
        createdOn?: google_protobuf_timestamp_pb.Timestamp.AsObject,
        updatedOn?: google_protobuf_timestamp_pb.Timestamp.AsObject,
        createdBy: string,
        updatedBy: string,
    }
}
