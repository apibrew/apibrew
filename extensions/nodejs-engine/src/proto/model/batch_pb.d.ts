// package: model
// file: model/batch.proto

/* tslint:disable */
/* eslint-disable */

import * as jspb from "google-protobuf";
import * as model_record_pb from "../model/record_pb";
import * as model_resource_pb from "../model/resource_pb";
import * as google_protobuf_struct_pb from "google-protobuf/google/protobuf/struct_pb";

export class BatchHeader extends jspb.Message { 
    getMode(): BatchHeader.BatchMode;
    setMode(value: BatchHeader.BatchMode): BatchHeader;

    getAnnotationsMap(): jspb.Map<string, string>;
    clearAnnotationsMap(): void;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): BatchHeader.AsObject;
    static toObject(includeInstance: boolean, msg: BatchHeader): BatchHeader.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: BatchHeader, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): BatchHeader;
    static deserializeBinaryFromReader(message: BatchHeader, reader: jspb.BinaryReader): BatchHeader;
}

export namespace BatchHeader {
    export type AsObject = {
        mode: BatchHeader.BatchMode,

        annotationsMap: Array<[string, string]>,
    }

    export enum BatchMode {
    CREATE = 0,
    UPDATE = 1,
    DELETE = 2,
    }

}

export class BatchRecordsPart extends jspb.Message { 
    getNamespace(): string;
    setNamespace(value: string): BatchRecordsPart;
    getResource(): string;
    setResource(value: string): BatchRecordsPart;
    clearValuesList(): void;
    getValuesList(): Array<google_protobuf_struct_pb.Value>;
    setValuesList(value: Array<google_protobuf_struct_pb.Value>): BatchRecordsPart;
    addValues(value?: google_protobuf_struct_pb.Value, index?: number): google_protobuf_struct_pb.Value;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): BatchRecordsPart.AsObject;
    static toObject(includeInstance: boolean, msg: BatchRecordsPart): BatchRecordsPart.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: BatchRecordsPart, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): BatchRecordsPart;
    static deserializeBinaryFromReader(message: BatchRecordsPart, reader: jspb.BinaryReader): BatchRecordsPart;
}

export namespace BatchRecordsPart {
    export type AsObject = {
        namespace: string,
        resource: string,
        valuesList: Array<google_protobuf_struct_pb.Value.AsObject>,
    }
}

export class Batch extends jspb.Message { 

    hasHeader(): boolean;
    clearHeader(): void;
    getHeader(): BatchHeader | undefined;
    setHeader(value?: BatchHeader): Batch;
    clearResourcesList(): void;
    getResourcesList(): Array<model_resource_pb.Resource>;
    setResourcesList(value: Array<model_resource_pb.Resource>): Batch;
    addResources(value?: model_resource_pb.Resource, index?: number): model_resource_pb.Resource;
    clearBatchrecordsList(): void;
    getBatchrecordsList(): Array<BatchRecordsPart>;
    setBatchrecordsList(value: Array<BatchRecordsPart>): Batch;
    addBatchrecords(value?: BatchRecordsPart, index?: number): BatchRecordsPart;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): Batch.AsObject;
    static toObject(includeInstance: boolean, msg: Batch): Batch.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: Batch, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): Batch;
    static deserializeBinaryFromReader(message: Batch, reader: jspb.BinaryReader): Batch;
}

export namespace Batch {
    export type AsObject = {
        header?: BatchHeader.AsObject,
        resourcesList: Array<model_resource_pb.Resource.AsObject>,
        batchrecordsList: Array<BatchRecordsPart.AsObject>,
    }
}
