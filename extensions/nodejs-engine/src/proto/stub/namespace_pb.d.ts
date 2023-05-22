// package: stub
// file: stub/namespace.proto

/* tslint:disable */
/* eslint-disable */

import * as jspb from "google-protobuf";
import * as model_namespace_pb from "../model/namespace_pb";
import * as model_error_pb from "../model/error_pb";


export class ListNamespaceRequest extends jspb.Message { 
    getToken(): string;
    setToken(value: string): ListNamespaceRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ListNamespaceRequest.AsObject;
    static toObject(includeInstance: boolean, msg: ListNamespaceRequest): ListNamespaceRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ListNamespaceRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ListNamespaceRequest;
    static deserializeBinaryFromReader(message: ListNamespaceRequest, reader: jspb.BinaryReader): ListNamespaceRequest;
}

export namespace ListNamespaceRequest {
    export type AsObject = {
        token: string,
    }
}

export class ListNamespaceResponse extends jspb.Message { 
    clearContentList(): void;
    getContentList(): Array<model_namespace_pb.Namespace>;
    setContentList(value: Array<model_namespace_pb.Namespace>): ListNamespaceResponse;
    addContent(value?: model_namespace_pb.Namespace, index?: number): model_namespace_pb.Namespace;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ListNamespaceResponse.AsObject;
    static toObject(includeInstance: boolean, msg: ListNamespaceResponse): ListNamespaceResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ListNamespaceResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ListNamespaceResponse;
    static deserializeBinaryFromReader(message: ListNamespaceResponse, reader: jspb.BinaryReader): ListNamespaceResponse;
}

export namespace ListNamespaceResponse {
    export type AsObject = {
        contentList: Array<model_namespace_pb.Namespace.AsObject>,
    }
}

export class CreateNamespaceRequest extends jspb.Message { 
    getToken(): string;
    setToken(value: string): CreateNamespaceRequest;
    clearNamespacesList(): void;
    getNamespacesList(): Array<model_namespace_pb.Namespace>;
    setNamespacesList(value: Array<model_namespace_pb.Namespace>): CreateNamespaceRequest;
    addNamespaces(value?: model_namespace_pb.Namespace, index?: number): model_namespace_pb.Namespace;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): CreateNamespaceRequest.AsObject;
    static toObject(includeInstance: boolean, msg: CreateNamespaceRequest): CreateNamespaceRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: CreateNamespaceRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): CreateNamespaceRequest;
    static deserializeBinaryFromReader(message: CreateNamespaceRequest, reader: jspb.BinaryReader): CreateNamespaceRequest;
}

export namespace CreateNamespaceRequest {
    export type AsObject = {
        token: string,
        namespacesList: Array<model_namespace_pb.Namespace.AsObject>,
    }
}

export class CreateNamespaceResponse extends jspb.Message { 
    clearNamespacesList(): void;
    getNamespacesList(): Array<model_namespace_pb.Namespace>;
    setNamespacesList(value: Array<model_namespace_pb.Namespace>): CreateNamespaceResponse;
    addNamespaces(value?: model_namespace_pb.Namespace, index?: number): model_namespace_pb.Namespace;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): CreateNamespaceResponse.AsObject;
    static toObject(includeInstance: boolean, msg: CreateNamespaceResponse): CreateNamespaceResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: CreateNamespaceResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): CreateNamespaceResponse;
    static deserializeBinaryFromReader(message: CreateNamespaceResponse, reader: jspb.BinaryReader): CreateNamespaceResponse;
}

export namespace CreateNamespaceResponse {
    export type AsObject = {
        namespacesList: Array<model_namespace_pb.Namespace.AsObject>,
    }
}

export class UpdateNamespaceRequest extends jspb.Message { 
    getToken(): string;
    setToken(value: string): UpdateNamespaceRequest;
    clearNamespacesList(): void;
    getNamespacesList(): Array<model_namespace_pb.Namespace>;
    setNamespacesList(value: Array<model_namespace_pb.Namespace>): UpdateNamespaceRequest;
    addNamespaces(value?: model_namespace_pb.Namespace, index?: number): model_namespace_pb.Namespace;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): UpdateNamespaceRequest.AsObject;
    static toObject(includeInstance: boolean, msg: UpdateNamespaceRequest): UpdateNamespaceRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: UpdateNamespaceRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): UpdateNamespaceRequest;
    static deserializeBinaryFromReader(message: UpdateNamespaceRequest, reader: jspb.BinaryReader): UpdateNamespaceRequest;
}

export namespace UpdateNamespaceRequest {
    export type AsObject = {
        token: string,
        namespacesList: Array<model_namespace_pb.Namespace.AsObject>,
    }
}

export class UpdateNamespaceResponse extends jspb.Message { 
    clearNamespacesList(): void;
    getNamespacesList(): Array<model_namespace_pb.Namespace>;
    setNamespacesList(value: Array<model_namespace_pb.Namespace>): UpdateNamespaceResponse;
    addNamespaces(value?: model_namespace_pb.Namespace, index?: number): model_namespace_pb.Namespace;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): UpdateNamespaceResponse.AsObject;
    static toObject(includeInstance: boolean, msg: UpdateNamespaceResponse): UpdateNamespaceResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: UpdateNamespaceResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): UpdateNamespaceResponse;
    static deserializeBinaryFromReader(message: UpdateNamespaceResponse, reader: jspb.BinaryReader): UpdateNamespaceResponse;
}

export namespace UpdateNamespaceResponse {
    export type AsObject = {
        namespacesList: Array<model_namespace_pb.Namespace.AsObject>,
    }
}

export class DeleteNamespaceRequest extends jspb.Message { 
    getToken(): string;
    setToken(value: string): DeleteNamespaceRequest;
    clearIdsList(): void;
    getIdsList(): Array<string>;
    setIdsList(value: Array<string>): DeleteNamespaceRequest;
    addIds(value: string, index?: number): string;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): DeleteNamespaceRequest.AsObject;
    static toObject(includeInstance: boolean, msg: DeleteNamespaceRequest): DeleteNamespaceRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: DeleteNamespaceRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): DeleteNamespaceRequest;
    static deserializeBinaryFromReader(message: DeleteNamespaceRequest, reader: jspb.BinaryReader): DeleteNamespaceRequest;
}

export namespace DeleteNamespaceRequest {
    export type AsObject = {
        token: string,
        idsList: Array<string>,
    }
}

export class DeleteNamespaceResponse extends jspb.Message { 
    clearNamespacesList(): void;
    getNamespacesList(): Array<model_namespace_pb.Namespace>;
    setNamespacesList(value: Array<model_namespace_pb.Namespace>): DeleteNamespaceResponse;
    addNamespaces(value?: model_namespace_pb.Namespace, index?: number): model_namespace_pb.Namespace;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): DeleteNamespaceResponse.AsObject;
    static toObject(includeInstance: boolean, msg: DeleteNamespaceResponse): DeleteNamespaceResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: DeleteNamespaceResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): DeleteNamespaceResponse;
    static deserializeBinaryFromReader(message: DeleteNamespaceResponse, reader: jspb.BinaryReader): DeleteNamespaceResponse;
}

export namespace DeleteNamespaceResponse {
    export type AsObject = {
        namespacesList: Array<model_namespace_pb.Namespace.AsObject>,
    }
}

export class GetNamespaceRequest extends jspb.Message { 
    getToken(): string;
    setToken(value: string): GetNamespaceRequest;
    getId(): string;
    setId(value: string): GetNamespaceRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): GetNamespaceRequest.AsObject;
    static toObject(includeInstance: boolean, msg: GetNamespaceRequest): GetNamespaceRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: GetNamespaceRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): GetNamespaceRequest;
    static deserializeBinaryFromReader(message: GetNamespaceRequest, reader: jspb.BinaryReader): GetNamespaceRequest;
}

export namespace GetNamespaceRequest {
    export type AsObject = {
        token: string,
        id: string,
    }
}

export class GetNamespaceResponse extends jspb.Message { 

    hasNamespace(): boolean;
    clearNamespace(): void;
    getNamespace(): model_namespace_pb.Namespace | undefined;
    setNamespace(value?: model_namespace_pb.Namespace): GetNamespaceResponse;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): GetNamespaceResponse.AsObject;
    static toObject(includeInstance: boolean, msg: GetNamespaceResponse): GetNamespaceResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: GetNamespaceResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): GetNamespaceResponse;
    static deserializeBinaryFromReader(message: GetNamespaceResponse, reader: jspb.BinaryReader): GetNamespaceResponse;
}

export namespace GetNamespaceResponse {
    export type AsObject = {
        namespace?: model_namespace_pb.Namespace.AsObject,
    }
}
