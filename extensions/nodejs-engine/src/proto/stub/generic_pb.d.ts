// package: stub
// file: stub/generic.proto

/* tslint:disable */
/* eslint-disable */

import * as jspb from "google-protobuf";
import * as model_query_pb from "../model/query_pb";
import * as model_error_pb from "../model/error_pb";
import * as google_protobuf_struct_pb from "google-protobuf/google/protobuf/struct_pb";
import * as google_protobuf_any_pb from "google-protobuf/google/protobuf/any_pb";

export class ListRequest extends jspb.Message { 
    getToken(): string;
    setToken(value: string): ListRequest;
    getNamespace(): string;
    setNamespace(value: string): ListRequest;
    getResource(): string;
    setResource(value: string): ListRequest;

    getFiltersMap(): jspb.Map<string, string>;
    clearFiltersMap(): void;
    getLimit(): number;
    setLimit(value: number): ListRequest;
    getOffset(): number;
    setOffset(value: number): ListRequest;
    getUsehistory(): boolean;
    setUsehistory(value: boolean): ListRequest;
    clearResolvereferencesList(): void;
    getResolvereferencesList(): Array<string>;
    setResolvereferencesList(value: Array<string>): ListRequest;
    addResolvereferences(value: string, index?: number): string;

    getAnnotationsMap(): jspb.Map<string, string>;
    clearAnnotationsMap(): void;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ListRequest.AsObject;
    static toObject(includeInstance: boolean, msg: ListRequest): ListRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ListRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ListRequest;
    static deserializeBinaryFromReader(message: ListRequest, reader: jspb.BinaryReader): ListRequest;
}

export namespace ListRequest {
    export type AsObject = {
        token: string,
        namespace: string,
        resource: string,

        filtersMap: Array<[string, string]>,
        limit: number,
        offset: number,
        usehistory: boolean,
        resolvereferencesList: Array<string>,

        annotationsMap: Array<[string, string]>,
    }
}

export class ListResponse extends jspb.Message { 
    getTotal(): number;
    setTotal(value: number): ListResponse;
    clearContentList(): void;
    getContentList(): Array<google_protobuf_any_pb.Any>;
    setContentList(value: Array<google_protobuf_any_pb.Any>): ListResponse;
    addContent(value?: google_protobuf_any_pb.Any, index?: number): google_protobuf_any_pb.Any;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ListResponse.AsObject;
    static toObject(includeInstance: boolean, msg: ListResponse): ListResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ListResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ListResponse;
    static deserializeBinaryFromReader(message: ListResponse, reader: jspb.BinaryReader): ListResponse;
}

export namespace ListResponse {
    export type AsObject = {
        total: number,
        contentList: Array<google_protobuf_any_pb.Any.AsObject>,
    }
}

export class SearchRequest extends jspb.Message { 
    getToken(): string;
    setToken(value: string): SearchRequest;
    getNamespace(): string;
    setNamespace(value: string): SearchRequest;
    getResource(): string;
    setResource(value: string): SearchRequest;

    hasQuery(): boolean;
    clearQuery(): void;
    getQuery(): model_query_pb.BooleanExpression | undefined;
    setQuery(value?: model_query_pb.BooleanExpression): SearchRequest;
    getLimit(): number;
    setLimit(value: number): SearchRequest;
    getOffset(): number;
    setOffset(value: number): SearchRequest;
    getUsehistory(): boolean;
    setUsehistory(value: boolean): SearchRequest;
    clearResolvereferencesList(): void;
    getResolvereferencesList(): Array<string>;
    setResolvereferencesList(value: Array<string>): SearchRequest;
    addResolvereferences(value: string, index?: number): string;

    getAnnotationsMap(): jspb.Map<string, string>;
    clearAnnotationsMap(): void;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): SearchRequest.AsObject;
    static toObject(includeInstance: boolean, msg: SearchRequest): SearchRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: SearchRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): SearchRequest;
    static deserializeBinaryFromReader(message: SearchRequest, reader: jspb.BinaryReader): SearchRequest;
}

export namespace SearchRequest {
    export type AsObject = {
        token: string,
        namespace: string,
        resource: string,
        query?: model_query_pb.BooleanExpression.AsObject,
        limit: number,
        offset: number,
        usehistory: boolean,
        resolvereferencesList: Array<string>,

        annotationsMap: Array<[string, string]>,
    }
}

export class SearchResponse extends jspb.Message { 
    getTotal(): number;
    setTotal(value: number): SearchResponse;
    clearContentList(): void;
    getContentList(): Array<google_protobuf_any_pb.Any>;
    setContentList(value: Array<google_protobuf_any_pb.Any>): SearchResponse;
    addContent(value?: google_protobuf_any_pb.Any, index?: number): google_protobuf_any_pb.Any;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): SearchResponse.AsObject;
    static toObject(includeInstance: boolean, msg: SearchResponse): SearchResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: SearchResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): SearchResponse;
    static deserializeBinaryFromReader(message: SearchResponse, reader: jspb.BinaryReader): SearchResponse;
}

export namespace SearchResponse {
    export type AsObject = {
        total: number,
        contentList: Array<google_protobuf_any_pb.Any.AsObject>,
    }
}

export class CreateRequest extends jspb.Message { 
    getToken(): string;
    setToken(value: string): CreateRequest;
    getNamespace(): string;
    setNamespace(value: string): CreateRequest;
    getResource(): string;
    setResource(value: string): CreateRequest;
    clearItemsList(): void;
    getItemsList(): Array<google_protobuf_any_pb.Any>;
    setItemsList(value: Array<google_protobuf_any_pb.Any>): CreateRequest;
    addItems(value?: google_protobuf_any_pb.Any, index?: number): google_protobuf_any_pb.Any;
    getIgnoreifexists(): boolean;
    setIgnoreifexists(value: boolean): CreateRequest;

    getAnnotationsMap(): jspb.Map<string, string>;
    clearAnnotationsMap(): void;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): CreateRequest.AsObject;
    static toObject(includeInstance: boolean, msg: CreateRequest): CreateRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: CreateRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): CreateRequest;
    static deserializeBinaryFromReader(message: CreateRequest, reader: jspb.BinaryReader): CreateRequest;
}

export namespace CreateRequest {
    export type AsObject = {
        token: string,
        namespace: string,
        resource: string,
        itemsList: Array<google_protobuf_any_pb.Any.AsObject>,
        ignoreifexists: boolean,

        annotationsMap: Array<[string, string]>,
    }
}

export class CreateResponse extends jspb.Message { 
    clearItemsList(): void;
    getItemsList(): Array<google_protobuf_any_pb.Any>;
    setItemsList(value: Array<google_protobuf_any_pb.Any>): CreateResponse;
    addItems(value?: google_protobuf_any_pb.Any, index?: number): google_protobuf_any_pb.Any;
    clearInsertedList(): void;
    getInsertedList(): Array<boolean>;
    setInsertedList(value: Array<boolean>): CreateResponse;
    addInserted(value: boolean, index?: number): boolean;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): CreateResponse.AsObject;
    static toObject(includeInstance: boolean, msg: CreateResponse): CreateResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: CreateResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): CreateResponse;
    static deserializeBinaryFromReader(message: CreateResponse, reader: jspb.BinaryReader): CreateResponse;
}

export namespace CreateResponse {
    export type AsObject = {
        itemsList: Array<google_protobuf_any_pb.Any.AsObject>,
        insertedList: Array<boolean>,
    }
}

export class UpdateRequest extends jspb.Message { 
    getToken(): string;
    setToken(value: string): UpdateRequest;
    getNamespace(): string;
    setNamespace(value: string): UpdateRequest;
    getResource(): string;
    setResource(value: string): UpdateRequest;
    clearItemsList(): void;
    getItemsList(): Array<google_protobuf_any_pb.Any>;
    setItemsList(value: Array<google_protobuf_any_pb.Any>): UpdateRequest;
    addItems(value?: google_protobuf_any_pb.Any, index?: number): google_protobuf_any_pb.Any;
    getCheckversion(): boolean;
    setCheckversion(value: boolean): UpdateRequest;

    getAnnotationsMap(): jspb.Map<string, string>;
    clearAnnotationsMap(): void;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): UpdateRequest.AsObject;
    static toObject(includeInstance: boolean, msg: UpdateRequest): UpdateRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: UpdateRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): UpdateRequest;
    static deserializeBinaryFromReader(message: UpdateRequest, reader: jspb.BinaryReader): UpdateRequest;
}

export namespace UpdateRequest {
    export type AsObject = {
        token: string,
        namespace: string,
        resource: string,
        itemsList: Array<google_protobuf_any_pb.Any.AsObject>,
        checkversion: boolean,

        annotationsMap: Array<[string, string]>,
    }
}

export class UpdateResponse extends jspb.Message { 
    clearItemsList(): void;
    getItemsList(): Array<google_protobuf_any_pb.Any>;
    setItemsList(value: Array<google_protobuf_any_pb.Any>): UpdateResponse;
    addItems(value?: google_protobuf_any_pb.Any, index?: number): google_protobuf_any_pb.Any;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): UpdateResponse.AsObject;
    static toObject(includeInstance: boolean, msg: UpdateResponse): UpdateResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: UpdateResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): UpdateResponse;
    static deserializeBinaryFromReader(message: UpdateResponse, reader: jspb.BinaryReader): UpdateResponse;
}

export namespace UpdateResponse {
    export type AsObject = {
        itemsList: Array<google_protobuf_any_pb.Any.AsObject>,
    }
}

export class UpdateMultiRequest extends jspb.Message { 
    getToken(): string;
    setToken(value: string): UpdateMultiRequest;
    getNamespace(): string;
    setNamespace(value: string): UpdateMultiRequest;
    getResource(): string;
    setResource(value: string): UpdateMultiRequest;

    hasQuery(): boolean;
    clearQuery(): void;
    getQuery(): model_query_pb.BooleanExpression | undefined;
    setQuery(value?: model_query_pb.BooleanExpression): UpdateMultiRequest;

    getPropertiesMap(): jspb.Map<string, google_protobuf_struct_pb.Value>;
    clearPropertiesMap(): void;

    getAnnotationsMap(): jspb.Map<string, string>;
    clearAnnotationsMap(): void;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): UpdateMultiRequest.AsObject;
    static toObject(includeInstance: boolean, msg: UpdateMultiRequest): UpdateMultiRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: UpdateMultiRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): UpdateMultiRequest;
    static deserializeBinaryFromReader(message: UpdateMultiRequest, reader: jspb.BinaryReader): UpdateMultiRequest;
}

export namespace UpdateMultiRequest {
    export type AsObject = {
        token: string,
        namespace: string,
        resource: string,
        query?: model_query_pb.BooleanExpression.AsObject,

        propertiesMap: Array<[string, google_protobuf_struct_pb.Value.AsObject]>,

        annotationsMap: Array<[string, string]>,
    }
}

export class UpdateMultiResponse extends jspb.Message { 
    clearItemsList(): void;
    getItemsList(): Array<google_protobuf_any_pb.Any>;
    setItemsList(value: Array<google_protobuf_any_pb.Any>): UpdateMultiResponse;
    addItems(value?: google_protobuf_any_pb.Any, index?: number): google_protobuf_any_pb.Any;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): UpdateMultiResponse.AsObject;
    static toObject(includeInstance: boolean, msg: UpdateMultiResponse): UpdateMultiResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: UpdateMultiResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): UpdateMultiResponse;
    static deserializeBinaryFromReader(message: UpdateMultiResponse, reader: jspb.BinaryReader): UpdateMultiResponse;
}

export namespace UpdateMultiResponse {
    export type AsObject = {
        itemsList: Array<google_protobuf_any_pb.Any.AsObject>,
    }
}

export class DeleteRequest extends jspb.Message { 
    getToken(): string;
    setToken(value: string): DeleteRequest;
    getNamespace(): string;
    setNamespace(value: string): DeleteRequest;
    getResource(): string;
    setResource(value: string): DeleteRequest;
    clearIdList(): void;
    getIdList(): Array<string>;
    setIdList(value: Array<string>): DeleteRequest;
    addId(value: string, index?: number): string;
    clearIdsList(): void;
    getIdsList(): Array<string>;
    setIdsList(value: Array<string>): DeleteRequest;
    addIds(value: string, index?: number): string;

    getAnnotationsMap(): jspb.Map<string, string>;
    clearAnnotationsMap(): void;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): DeleteRequest.AsObject;
    static toObject(includeInstance: boolean, msg: DeleteRequest): DeleteRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: DeleteRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): DeleteRequest;
    static deserializeBinaryFromReader(message: DeleteRequest, reader: jspb.BinaryReader): DeleteRequest;
}

export namespace DeleteRequest {
    export type AsObject = {
        token: string,
        namespace: string,
        resource: string,
        idList: Array<string>,
        idsList: Array<string>,

        annotationsMap: Array<[string, string]>,
    }
}

export class DeleteResponse extends jspb.Message { 

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): DeleteResponse.AsObject;
    static toObject(includeInstance: boolean, msg: DeleteResponse): DeleteResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: DeleteResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): DeleteResponse;
    static deserializeBinaryFromReader(message: DeleteResponse, reader: jspb.BinaryReader): DeleteResponse;
}

export namespace DeleteResponse {
    export type AsObject = {
    }
}

export class GetRequest extends jspb.Message { 
    getToken(): string;
    setToken(value: string): GetRequest;
    getNamespace(): string;
    setNamespace(value: string): GetRequest;
    getResource(): string;
    setResource(value: string): GetRequest;
    getId(): string;
    setId(value: string): GetRequest;

    getAnnotationsMap(): jspb.Map<string, string>;
    clearAnnotationsMap(): void;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): GetRequest.AsObject;
    static toObject(includeInstance: boolean, msg: GetRequest): GetRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: GetRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): GetRequest;
    static deserializeBinaryFromReader(message: GetRequest, reader: jspb.BinaryReader): GetRequest;
}

export namespace GetRequest {
    export type AsObject = {
        token: string,
        namespace: string,
        resource: string,
        id: string,

        annotationsMap: Array<[string, string]>,
    }
}

export class GetResponse extends jspb.Message { 

    hasItem(): boolean;
    clearItem(): void;
    getItem(): google_protobuf_any_pb.Any | undefined;
    setItem(value?: google_protobuf_any_pb.Any): GetResponse;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): GetResponse.AsObject;
    static toObject(includeInstance: boolean, msg: GetResponse): GetResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: GetResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): GetResponse;
    static deserializeBinaryFromReader(message: GetResponse, reader: jspb.BinaryReader): GetResponse;
}

export namespace GetResponse {
    export type AsObject = {
        item?: google_protobuf_any_pb.Any.AsObject,
    }
}
