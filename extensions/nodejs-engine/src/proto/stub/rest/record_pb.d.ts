// package: rest
// file: stub/rest/record.proto

/* tslint:disable */
/* eslint-disable */

import * as jspb from "google-protobuf";
import * as model_record_pb from "../../model/record_pb";
import * as model_query_pb from "../../model/query_pb";
import * as model_error_pb from "../../model/error_pb";
import * as google_protobuf_struct_pb from "google-protobuf/google/protobuf/struct_pb";

export class CreateRecordRequest extends jspb.Message { 
    getToken(): string;
    setToken(value: string): CreateRecordRequest;
    getNamespace(): string;
    setNamespace(value: string): CreateRecordRequest;
    getResource(): string;
    setResource(value: string): CreateRecordRequest;

    getPropertiesMap(): jspb.Map<string, google_protobuf_struct_pb.Value>;
    clearPropertiesMap(): void;

    getAnnotationsMap(): jspb.Map<string, string>;
    clearAnnotationsMap(): void;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): CreateRecordRequest.AsObject;
    static toObject(includeInstance: boolean, msg: CreateRecordRequest): CreateRecordRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: CreateRecordRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): CreateRecordRequest;
    static deserializeBinaryFromReader(message: CreateRecordRequest, reader: jspb.BinaryReader): CreateRecordRequest;
}

export namespace CreateRecordRequest {
    export type AsObject = {
        token: string,
        namespace: string,
        resource: string,

        propertiesMap: Array<[string, google_protobuf_struct_pb.Value.AsObject]>,

        annotationsMap: Array<[string, string]>,
    }
}

export class CreateRecordResponse extends jspb.Message { 
    getId(): string;
    setId(value: string): CreateRecordResponse;

    getPropertiesMap(): jspb.Map<string, google_protobuf_struct_pb.Value>;
    clearPropertiesMap(): void;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): CreateRecordResponse.AsObject;
    static toObject(includeInstance: boolean, msg: CreateRecordResponse): CreateRecordResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: CreateRecordResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): CreateRecordResponse;
    static deserializeBinaryFromReader(message: CreateRecordResponse, reader: jspb.BinaryReader): CreateRecordResponse;
}

export namespace CreateRecordResponse {
    export type AsObject = {
        id: string,

        propertiesMap: Array<[string, google_protobuf_struct_pb.Value.AsObject]>,
    }
}

export class UpdateRecordRequest extends jspb.Message { 
    getToken(): string;
    setToken(value: string): UpdateRecordRequest;
    getNamespace(): string;
    setNamespace(value: string): UpdateRecordRequest;
    getResource(): string;
    setResource(value: string): UpdateRecordRequest;
    getId(): string;
    setId(value: string): UpdateRecordRequest;

    getPropertiesMap(): jspb.Map<string, google_protobuf_struct_pb.Value>;
    clearPropertiesMap(): void;
    getCheckversion(): boolean;
    setCheckversion(value: boolean): UpdateRecordRequest;

    getAnnotationsMap(): jspb.Map<string, string>;
    clearAnnotationsMap(): void;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): UpdateRecordRequest.AsObject;
    static toObject(includeInstance: boolean, msg: UpdateRecordRequest): UpdateRecordRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: UpdateRecordRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): UpdateRecordRequest;
    static deserializeBinaryFromReader(message: UpdateRecordRequest, reader: jspb.BinaryReader): UpdateRecordRequest;
}

export namespace UpdateRecordRequest {
    export type AsObject = {
        token: string,
        namespace: string,
        resource: string,
        id: string,

        propertiesMap: Array<[string, google_protobuf_struct_pb.Value.AsObject]>,
        checkversion: boolean,

        annotationsMap: Array<[string, string]>,
    }
}

export class UpdateRecordResponse extends jspb.Message { 

    getPropertiesMap(): jspb.Map<string, google_protobuf_struct_pb.Value>;
    clearPropertiesMap(): void;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): UpdateRecordResponse.AsObject;
    static toObject(includeInstance: boolean, msg: UpdateRecordResponse): UpdateRecordResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: UpdateRecordResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): UpdateRecordResponse;
    static deserializeBinaryFromReader(message: UpdateRecordResponse, reader: jspb.BinaryReader): UpdateRecordResponse;
}

export namespace UpdateRecordResponse {
    export type AsObject = {

        propertiesMap: Array<[string, google_protobuf_struct_pb.Value.AsObject]>,
    }
}

export class ApplyRecordRequest extends jspb.Message { 
    getToken(): string;
    setToken(value: string): ApplyRecordRequest;
    getNamespace(): string;
    setNamespace(value: string): ApplyRecordRequest;
    getResource(): string;
    setResource(value: string): ApplyRecordRequest;
    getId(): string;
    setId(value: string): ApplyRecordRequest;

    getPropertiesMap(): jspb.Map<string, google_protobuf_struct_pb.Value>;
    clearPropertiesMap(): void;
    getCheckversion(): boolean;
    setCheckversion(value: boolean): ApplyRecordRequest;

    getAnnotationsMap(): jspb.Map<string, string>;
    clearAnnotationsMap(): void;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ApplyRecordRequest.AsObject;
    static toObject(includeInstance: boolean, msg: ApplyRecordRequest): ApplyRecordRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ApplyRecordRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ApplyRecordRequest;
    static deserializeBinaryFromReader(message: ApplyRecordRequest, reader: jspb.BinaryReader): ApplyRecordRequest;
}

export namespace ApplyRecordRequest {
    export type AsObject = {
        token: string,
        namespace: string,
        resource: string,
        id: string,

        propertiesMap: Array<[string, google_protobuf_struct_pb.Value.AsObject]>,
        checkversion: boolean,

        annotationsMap: Array<[string, string]>,
    }
}

export class ApplyRecordResponse extends jspb.Message { 

    getPropertiesMap(): jspb.Map<string, google_protobuf_struct_pb.Value>;
    clearPropertiesMap(): void;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ApplyRecordResponse.AsObject;
    static toObject(includeInstance: boolean, msg: ApplyRecordResponse): ApplyRecordResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ApplyRecordResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ApplyRecordResponse;
    static deserializeBinaryFromReader(message: ApplyRecordResponse, reader: jspb.BinaryReader): ApplyRecordResponse;
}

export namespace ApplyRecordResponse {
    export type AsObject = {

        propertiesMap: Array<[string, google_protobuf_struct_pb.Value.AsObject]>,
    }
}

export class DeleteRecordRequest extends jspb.Message { 
    getToken(): string;
    setToken(value: string): DeleteRecordRequest;
    getNamespace(): string;
    setNamespace(value: string): DeleteRecordRequest;
    getResource(): string;
    setResource(value: string): DeleteRecordRequest;
    getId(): string;
    setId(value: string): DeleteRecordRequest;

    getAnnotationsMap(): jspb.Map<string, string>;
    clearAnnotationsMap(): void;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): DeleteRecordRequest.AsObject;
    static toObject(includeInstance: boolean, msg: DeleteRecordRequest): DeleteRecordRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: DeleteRecordRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): DeleteRecordRequest;
    static deserializeBinaryFromReader(message: DeleteRecordRequest, reader: jspb.BinaryReader): DeleteRecordRequest;
}

export namespace DeleteRecordRequest {
    export type AsObject = {
        token: string,
        namespace: string,
        resource: string,
        id: string,

        annotationsMap: Array<[string, string]>,
    }
}

export class DeleteRecordResponse extends jspb.Message { 

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): DeleteRecordResponse.AsObject;
    static toObject(includeInstance: boolean, msg: DeleteRecordResponse): DeleteRecordResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: DeleteRecordResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): DeleteRecordResponse;
    static deserializeBinaryFromReader(message: DeleteRecordResponse, reader: jspb.BinaryReader): DeleteRecordResponse;
}

export namespace DeleteRecordResponse {
    export type AsObject = {
    }
}
