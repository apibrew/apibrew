// package: stub
// file: stub/extension.proto

/* tslint:disable */
/* eslint-disable */

import * as jspb from "google-protobuf";
import * as model_query_pb from "../model/query_pb";
import * as model_extension_pb from "../model/extension_pb";

export class ListExtensionRequest extends jspb.Message { 
    getToken(): string;
    setToken(value: string): ListExtensionRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ListExtensionRequest.AsObject;
    static toObject(includeInstance: boolean, msg: ListExtensionRequest): ListExtensionRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ListExtensionRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ListExtensionRequest;
    static deserializeBinaryFromReader(message: ListExtensionRequest, reader: jspb.BinaryReader): ListExtensionRequest;
}

export namespace ListExtensionRequest {
    export type AsObject = {
        token: string,
    }
}

export class ListExtensionResponse extends jspb.Message { 
    clearContentList(): void;
    getContentList(): Array<model_extension_pb.Extension>;
    setContentList(value: Array<model_extension_pb.Extension>): ListExtensionResponse;
    addContent(value?: model_extension_pb.Extension, index?: number): model_extension_pb.Extension;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ListExtensionResponse.AsObject;
    static toObject(includeInstance: boolean, msg: ListExtensionResponse): ListExtensionResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ListExtensionResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ListExtensionResponse;
    static deserializeBinaryFromReader(message: ListExtensionResponse, reader: jspb.BinaryReader): ListExtensionResponse;
}

export namespace ListExtensionResponse {
    export type AsObject = {
        contentList: Array<model_extension_pb.Extension.AsObject>,
    }
}

export class GetExtensionRequest extends jspb.Message { 
    getToken(): string;
    setToken(value: string): GetExtensionRequest;
    getId(): string;
    setId(value: string): GetExtensionRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): GetExtensionRequest.AsObject;
    static toObject(includeInstance: boolean, msg: GetExtensionRequest): GetExtensionRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: GetExtensionRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): GetExtensionRequest;
    static deserializeBinaryFromReader(message: GetExtensionRequest, reader: jspb.BinaryReader): GetExtensionRequest;
}

export namespace GetExtensionRequest {
    export type AsObject = {
        token: string,
        id: string,
    }
}

export class GetExtensionResponse extends jspb.Message { 

    hasExtension$(): boolean;
    clearExtension$(): void;
    getExtension$(): model_extension_pb.Extension | undefined;
    setExtension$(value?: model_extension_pb.Extension): GetExtensionResponse;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): GetExtensionResponse.AsObject;
    static toObject(includeInstance: boolean, msg: GetExtensionResponse): GetExtensionResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: GetExtensionResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): GetExtensionResponse;
    static deserializeBinaryFromReader(message: GetExtensionResponse, reader: jspb.BinaryReader): GetExtensionResponse;
}

export namespace GetExtensionResponse {
    export type AsObject = {
        extension?: model_extension_pb.Extension.AsObject,
    }
}

export class CreateExtensionRequest extends jspb.Message { 
    getToken(): string;
    setToken(value: string): CreateExtensionRequest;
    clearExtensionsList(): void;
    getExtensionsList(): Array<model_extension_pb.Extension>;
    setExtensionsList(value: Array<model_extension_pb.Extension>): CreateExtensionRequest;
    addExtensions(value?: model_extension_pb.Extension, index?: number): model_extension_pb.Extension;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): CreateExtensionRequest.AsObject;
    static toObject(includeInstance: boolean, msg: CreateExtensionRequest): CreateExtensionRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: CreateExtensionRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): CreateExtensionRequest;
    static deserializeBinaryFromReader(message: CreateExtensionRequest, reader: jspb.BinaryReader): CreateExtensionRequest;
}

export namespace CreateExtensionRequest {
    export type AsObject = {
        token: string,
        extensionsList: Array<model_extension_pb.Extension.AsObject>,
    }
}

export class CreateExtensionResponse extends jspb.Message { 
    clearExtensionsList(): void;
    getExtensionsList(): Array<model_extension_pb.Extension>;
    setExtensionsList(value: Array<model_extension_pb.Extension>): CreateExtensionResponse;
    addExtensions(value?: model_extension_pb.Extension, index?: number): model_extension_pb.Extension;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): CreateExtensionResponse.AsObject;
    static toObject(includeInstance: boolean, msg: CreateExtensionResponse): CreateExtensionResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: CreateExtensionResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): CreateExtensionResponse;
    static deserializeBinaryFromReader(message: CreateExtensionResponse, reader: jspb.BinaryReader): CreateExtensionResponse;
}

export namespace CreateExtensionResponse {
    export type AsObject = {
        extensionsList: Array<model_extension_pb.Extension.AsObject>,
    }
}

export class UpdateExtensionRequest extends jspb.Message { 
    getToken(): string;
    setToken(value: string): UpdateExtensionRequest;
    clearExtensionsList(): void;
    getExtensionsList(): Array<model_extension_pb.Extension>;
    setExtensionsList(value: Array<model_extension_pb.Extension>): UpdateExtensionRequest;
    addExtensions(value?: model_extension_pb.Extension, index?: number): model_extension_pb.Extension;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): UpdateExtensionRequest.AsObject;
    static toObject(includeInstance: boolean, msg: UpdateExtensionRequest): UpdateExtensionRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: UpdateExtensionRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): UpdateExtensionRequest;
    static deserializeBinaryFromReader(message: UpdateExtensionRequest, reader: jspb.BinaryReader): UpdateExtensionRequest;
}

export namespace UpdateExtensionRequest {
    export type AsObject = {
        token: string,
        extensionsList: Array<model_extension_pb.Extension.AsObject>,
    }
}

export class UpdateExtensionResponse extends jspb.Message { 
    clearExtensionsList(): void;
    getExtensionsList(): Array<model_extension_pb.Extension>;
    setExtensionsList(value: Array<model_extension_pb.Extension>): UpdateExtensionResponse;
    addExtensions(value?: model_extension_pb.Extension, index?: number): model_extension_pb.Extension;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): UpdateExtensionResponse.AsObject;
    static toObject(includeInstance: boolean, msg: UpdateExtensionResponse): UpdateExtensionResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: UpdateExtensionResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): UpdateExtensionResponse;
    static deserializeBinaryFromReader(message: UpdateExtensionResponse, reader: jspb.BinaryReader): UpdateExtensionResponse;
}

export namespace UpdateExtensionResponse {
    export type AsObject = {
        extensionsList: Array<model_extension_pb.Extension.AsObject>,
    }
}

export class DeleteExtensionRequest extends jspb.Message { 
    getToken(): string;
    setToken(value: string): DeleteExtensionRequest;
    clearIdsList(): void;
    getIdsList(): Array<string>;
    setIdsList(value: Array<string>): DeleteExtensionRequest;
    addIds(value: string, index?: number): string;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): DeleteExtensionRequest.AsObject;
    static toObject(includeInstance: boolean, msg: DeleteExtensionRequest): DeleteExtensionRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: DeleteExtensionRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): DeleteExtensionRequest;
    static deserializeBinaryFromReader(message: DeleteExtensionRequest, reader: jspb.BinaryReader): DeleteExtensionRequest;
}

export namespace DeleteExtensionRequest {
    export type AsObject = {
        token: string,
        idsList: Array<string>,
    }
}

export class DeleteExtensionResponse extends jspb.Message { 

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): DeleteExtensionResponse.AsObject;
    static toObject(includeInstance: boolean, msg: DeleteExtensionResponse): DeleteExtensionResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: DeleteExtensionResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): DeleteExtensionResponse;
    static deserializeBinaryFromReader(message: DeleteExtensionResponse, reader: jspb.BinaryReader): DeleteExtensionResponse;
}

export namespace DeleteExtensionResponse {
    export type AsObject = {
    }
}
