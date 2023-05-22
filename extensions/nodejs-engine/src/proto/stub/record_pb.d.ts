// package: stub
// file: stub/record.proto

/* tslint:disable */
/* eslint-disable */

import * as jspb from "google-protobuf";
import * as model_record_pb from "../model/record_pb";
import * as model_query_pb from "../model/query_pb";
import * as model_error_pb from "../model/error_pb";
import * as google_protobuf_struct_pb from "google-protobuf/google/protobuf/struct_pb";

export class ListRecordRequest extends jspb.Message { 
    getToken(): string;
    setToken(value: string): ListRecordRequest;
    getNamespace(): string;
    setNamespace(value: string): ListRecordRequest;
    getResource(): string;
    setResource(value: string): ListRecordRequest;

    getFiltersMap(): jspb.Map<string, string>;
    clearFiltersMap(): void;
    getLimit(): number;
    setLimit(value: number): ListRecordRequest;
    getOffset(): number;
    setOffset(value: number): ListRecordRequest;
    getUsehistory(): boolean;
    setUsehistory(value: boolean): ListRecordRequest;
    clearResolvereferencesList(): void;
    getResolvereferencesList(): Array<string>;
    setResolvereferencesList(value: Array<string>): ListRecordRequest;
    addResolvereferences(value: string, index?: number): string;

    getAnnotationsMap(): jspb.Map<string, string>;
    clearAnnotationsMap(): void;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ListRecordRequest.AsObject;
    static toObject(includeInstance: boolean, msg: ListRecordRequest): ListRecordRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ListRecordRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ListRecordRequest;
    static deserializeBinaryFromReader(message: ListRecordRequest, reader: jspb.BinaryReader): ListRecordRequest;
}

export namespace ListRecordRequest {
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

export class ListRecordResponse extends jspb.Message { 
    getTotal(): number;
    setTotal(value: number): ListRecordResponse;
    clearContentList(): void;
    getContentList(): Array<model_record_pb.Record>;
    setContentList(value: Array<model_record_pb.Record>): ListRecordResponse;
    addContent(value?: model_record_pb.Record, index?: number): model_record_pb.Record;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ListRecordResponse.AsObject;
    static toObject(includeInstance: boolean, msg: ListRecordResponse): ListRecordResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ListRecordResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ListRecordResponse;
    static deserializeBinaryFromReader(message: ListRecordResponse, reader: jspb.BinaryReader): ListRecordResponse;
}

export namespace ListRecordResponse {
    export type AsObject = {
        total: number,
        contentList: Array<model_record_pb.Record.AsObject>,
    }
}

export class SearchRecordRequest extends jspb.Message { 
    getToken(): string;
    setToken(value: string): SearchRecordRequest;
    getNamespace(): string;
    setNamespace(value: string): SearchRecordRequest;
    getResource(): string;
    setResource(value: string): SearchRecordRequest;

    hasQuery(): boolean;
    clearQuery(): void;
    getQuery(): model_query_pb.BooleanExpression | undefined;
    setQuery(value?: model_query_pb.BooleanExpression): SearchRecordRequest;
    getLimit(): number;
    setLimit(value: number): SearchRecordRequest;
    getOffset(): number;
    setOffset(value: number): SearchRecordRequest;
    getUsehistory(): boolean;
    setUsehistory(value: boolean): SearchRecordRequest;
    clearResolvereferencesList(): void;
    getResolvereferencesList(): Array<string>;
    setResolvereferencesList(value: Array<string>): SearchRecordRequest;
    addResolvereferences(value: string, index?: number): string;

    getAnnotationsMap(): jspb.Map<string, string>;
    clearAnnotationsMap(): void;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): SearchRecordRequest.AsObject;
    static toObject(includeInstance: boolean, msg: SearchRecordRequest): SearchRecordRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: SearchRecordRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): SearchRecordRequest;
    static deserializeBinaryFromReader(message: SearchRecordRequest, reader: jspb.BinaryReader): SearchRecordRequest;
}

export namespace SearchRecordRequest {
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

export class ReadStreamRequest extends jspb.Message { 
    getToken(): string;
    setToken(value: string): ReadStreamRequest;
    getNamespace(): string;
    setNamespace(value: string): ReadStreamRequest;
    getResource(): string;
    setResource(value: string): ReadStreamRequest;

    hasQuery(): boolean;
    clearQuery(): void;
    getQuery(): model_query_pb.BooleanExpression | undefined;
    setQuery(value?: model_query_pb.BooleanExpression): ReadStreamRequest;
    getLimit(): number;
    setLimit(value: number): ReadStreamRequest;
    getOffset(): number;
    setOffset(value: number): ReadStreamRequest;
    getUsehistory(): boolean;
    setUsehistory(value: boolean): ReadStreamRequest;
    clearResolvereferencesList(): void;
    getResolvereferencesList(): Array<string>;
    setResolvereferencesList(value: Array<string>): ReadStreamRequest;
    addResolvereferences(value: string, index?: number): string;
    getUsetransaction(): boolean;
    setUsetransaction(value: boolean): ReadStreamRequest;
    getPackrecords(): boolean;
    setPackrecords(value: boolean): ReadStreamRequest;

    getAnnotationsMap(): jspb.Map<string, string>;
    clearAnnotationsMap(): void;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ReadStreamRequest.AsObject;
    static toObject(includeInstance: boolean, msg: ReadStreamRequest): ReadStreamRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ReadStreamRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ReadStreamRequest;
    static deserializeBinaryFromReader(message: ReadStreamRequest, reader: jspb.BinaryReader): ReadStreamRequest;
}

export namespace ReadStreamRequest {
    export type AsObject = {
        token: string,
        namespace: string,
        resource: string,
        query?: model_query_pb.BooleanExpression.AsObject,
        limit: number,
        offset: number,
        usehistory: boolean,
        resolvereferencesList: Array<string>,
        usetransaction: boolean,
        packrecords: boolean,

        annotationsMap: Array<[string, string]>,
    }
}

export class WriteStreamResponse extends jspb.Message { 
    clearSuccessList(): void;
    getSuccessList(): Array<boolean>;
    setSuccessList(value: Array<boolean>): WriteStreamResponse;
    addSuccess(value: boolean, index?: number): boolean;
    clearCreatedList(): void;
    getCreatedList(): Array<boolean>;
    setCreatedList(value: Array<boolean>): WriteStreamResponse;
    addCreated(value: boolean, index?: number): boolean;
    clearUpdatedList(): void;
    getUpdatedList(): Array<boolean>;
    setUpdatedList(value: Array<boolean>): WriteStreamResponse;
    addUpdated(value: boolean, index?: number): boolean;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): WriteStreamResponse.AsObject;
    static toObject(includeInstance: boolean, msg: WriteStreamResponse): WriteStreamResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: WriteStreamResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): WriteStreamResponse;
    static deserializeBinaryFromReader(message: WriteStreamResponse, reader: jspb.BinaryReader): WriteStreamResponse;
}

export namespace WriteStreamResponse {
    export type AsObject = {
        successList: Array<boolean>,
        createdList: Array<boolean>,
        updatedList: Array<boolean>,
    }
}

export class SearchRecordResponse extends jspb.Message { 
    getTotal(): number;
    setTotal(value: number): SearchRecordResponse;
    clearContentList(): void;
    getContentList(): Array<model_record_pb.Record>;
    setContentList(value: Array<model_record_pb.Record>): SearchRecordResponse;
    addContent(value?: model_record_pb.Record, index?: number): model_record_pb.Record;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): SearchRecordResponse.AsObject;
    static toObject(includeInstance: boolean, msg: SearchRecordResponse): SearchRecordResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: SearchRecordResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): SearchRecordResponse;
    static deserializeBinaryFromReader(message: SearchRecordResponse, reader: jspb.BinaryReader): SearchRecordResponse;
}

export namespace SearchRecordResponse {
    export type AsObject = {
        total: number,
        contentList: Array<model_record_pb.Record.AsObject>,
    }
}

export class CreateRecordRequest extends jspb.Message { 
    getToken(): string;
    setToken(value: string): CreateRecordRequest;
    getNamespace(): string;
    setNamespace(value: string): CreateRecordRequest;
    getResource(): string;
    setResource(value: string): CreateRecordRequest;

    hasRecord(): boolean;
    clearRecord(): void;
    getRecord(): model_record_pb.Record | undefined;
    setRecord(value?: model_record_pb.Record): CreateRecordRequest;
    clearRecordsList(): void;
    getRecordsList(): Array<model_record_pb.Record>;
    setRecordsList(value: Array<model_record_pb.Record>): CreateRecordRequest;
    addRecords(value?: model_record_pb.Record, index?: number): model_record_pb.Record;

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
        record?: model_record_pb.Record.AsObject,
        recordsList: Array<model_record_pb.Record.AsObject>,

        annotationsMap: Array<[string, string]>,
    }
}

export class CreateRecordResponse extends jspb.Message { 

    hasRecord(): boolean;
    clearRecord(): void;
    getRecord(): model_record_pb.Record | undefined;
    setRecord(value?: model_record_pb.Record): CreateRecordResponse;
    clearRecordsList(): void;
    getRecordsList(): Array<model_record_pb.Record>;
    setRecordsList(value: Array<model_record_pb.Record>): CreateRecordResponse;
    addRecords(value?: model_record_pb.Record, index?: number): model_record_pb.Record;
    clearInsertedList(): void;
    getInsertedList(): Array<boolean>;
    setInsertedList(value: Array<boolean>): CreateRecordResponse;
    addInserted(value: boolean, index?: number): boolean;

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
        record?: model_record_pb.Record.AsObject,
        recordsList: Array<model_record_pb.Record.AsObject>,
        insertedList: Array<boolean>,
    }
}

export class UpdateRecordRequest extends jspb.Message { 
    getToken(): string;
    setToken(value: string): UpdateRecordRequest;
    getNamespace(): string;
    setNamespace(value: string): UpdateRecordRequest;
    getResource(): string;
    setResource(value: string): UpdateRecordRequest;

    hasRecord(): boolean;
    clearRecord(): void;
    getRecord(): model_record_pb.Record | undefined;
    setRecord(value?: model_record_pb.Record): UpdateRecordRequest;
    clearRecordsList(): void;
    getRecordsList(): Array<model_record_pb.Record>;
    setRecordsList(value: Array<model_record_pb.Record>): UpdateRecordRequest;
    addRecords(value?: model_record_pb.Record, index?: number): model_record_pb.Record;

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
        record?: model_record_pb.Record.AsObject,
        recordsList: Array<model_record_pb.Record.AsObject>,

        annotationsMap: Array<[string, string]>,
    }
}

export class UpdateRecordResponse extends jspb.Message { 

    hasRecord(): boolean;
    clearRecord(): void;
    getRecord(): model_record_pb.Record | undefined;
    setRecord(value?: model_record_pb.Record): UpdateRecordResponse;
    clearRecordsList(): void;
    getRecordsList(): Array<model_record_pb.Record>;
    setRecordsList(value: Array<model_record_pb.Record>): UpdateRecordResponse;
    addRecords(value?: model_record_pb.Record, index?: number): model_record_pb.Record;

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
        record?: model_record_pb.Record.AsObject,
        recordsList: Array<model_record_pb.Record.AsObject>,
    }
}

export class ApplyRecordRequest extends jspb.Message { 
    getToken(): string;
    setToken(value: string): ApplyRecordRequest;
    getNamespace(): string;
    setNamespace(value: string): ApplyRecordRequest;
    getResource(): string;
    setResource(value: string): ApplyRecordRequest;

    hasRecord(): boolean;
    clearRecord(): void;
    getRecord(): model_record_pb.Record | undefined;
    setRecord(value?: model_record_pb.Record): ApplyRecordRequest;
    clearRecordsList(): void;
    getRecordsList(): Array<model_record_pb.Record>;
    setRecordsList(value: Array<model_record_pb.Record>): ApplyRecordRequest;
    addRecords(value?: model_record_pb.Record, index?: number): model_record_pb.Record;

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
        record?: model_record_pb.Record.AsObject,
        recordsList: Array<model_record_pb.Record.AsObject>,

        annotationsMap: Array<[string, string]>,
    }
}

export class ApplyRecordResponse extends jspb.Message { 

    hasRecord(): boolean;
    clearRecord(): void;
    getRecord(): model_record_pb.Record | undefined;
    setRecord(value?: model_record_pb.Record): ApplyRecordResponse;
    clearRecordsList(): void;
    getRecordsList(): Array<model_record_pb.Record>;
    setRecordsList(value: Array<model_record_pb.Record>): ApplyRecordResponse;
    addRecords(value?: model_record_pb.Record, index?: number): model_record_pb.Record;

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
        record?: model_record_pb.Record.AsObject,
        recordsList: Array<model_record_pb.Record.AsObject>,
    }
}

export class UpdateMultiRecordRequest extends jspb.Message { 
    getToken(): string;
    setToken(value: string): UpdateMultiRecordRequest;
    getNamespace(): string;
    setNamespace(value: string): UpdateMultiRecordRequest;
    getResource(): string;
    setResource(value: string): UpdateMultiRecordRequest;

    hasQuery(): boolean;
    clearQuery(): void;
    getQuery(): model_query_pb.BooleanExpression | undefined;
    setQuery(value?: model_query_pb.BooleanExpression): UpdateMultiRecordRequest;

    getPropertiesMap(): jspb.Map<string, google_protobuf_struct_pb.Value>;
    clearPropertiesMap(): void;

    getAnnotationsMap(): jspb.Map<string, string>;
    clearAnnotationsMap(): void;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): UpdateMultiRecordRequest.AsObject;
    static toObject(includeInstance: boolean, msg: UpdateMultiRecordRequest): UpdateMultiRecordRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: UpdateMultiRecordRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): UpdateMultiRecordRequest;
    static deserializeBinaryFromReader(message: UpdateMultiRecordRequest, reader: jspb.BinaryReader): UpdateMultiRecordRequest;
}

export namespace UpdateMultiRecordRequest {
    export type AsObject = {
        token: string,
        namespace: string,
        resource: string,
        query?: model_query_pb.BooleanExpression.AsObject,

        propertiesMap: Array<[string, google_protobuf_struct_pb.Value.AsObject]>,

        annotationsMap: Array<[string, string]>,
    }
}

export class UpdateMultiRecordResponse extends jspb.Message { 

    hasRecord(): boolean;
    clearRecord(): void;
    getRecord(): model_record_pb.Record | undefined;
    setRecord(value?: model_record_pb.Record): UpdateMultiRecordResponse;
    clearRecordsList(): void;
    getRecordsList(): Array<model_record_pb.Record>;
    setRecordsList(value: Array<model_record_pb.Record>): UpdateMultiRecordResponse;
    addRecords(value?: model_record_pb.Record, index?: number): model_record_pb.Record;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): UpdateMultiRecordResponse.AsObject;
    static toObject(includeInstance: boolean, msg: UpdateMultiRecordResponse): UpdateMultiRecordResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: UpdateMultiRecordResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): UpdateMultiRecordResponse;
    static deserializeBinaryFromReader(message: UpdateMultiRecordResponse, reader: jspb.BinaryReader): UpdateMultiRecordResponse;
}

export namespace UpdateMultiRecordResponse {
    export type AsObject = {
        record?: model_record_pb.Record.AsObject,
        recordsList: Array<model_record_pb.Record.AsObject>,
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
    clearIdsList(): void;
    getIdsList(): Array<string>;
    setIdsList(value: Array<string>): DeleteRecordRequest;
    addIds(value: string, index?: number): string;

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
        idsList: Array<string>,

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

export class GetRecordRequest extends jspb.Message { 
    getToken(): string;
    setToken(value: string): GetRecordRequest;
    getNamespace(): string;
    setNamespace(value: string): GetRecordRequest;
    getResource(): string;
    setResource(value: string): GetRecordRequest;
    getId(): string;
    setId(value: string): GetRecordRequest;

    getAnnotationsMap(): jspb.Map<string, string>;
    clearAnnotationsMap(): void;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): GetRecordRequest.AsObject;
    static toObject(includeInstance: boolean, msg: GetRecordRequest): GetRecordRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: GetRecordRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): GetRecordRequest;
    static deserializeBinaryFromReader(message: GetRecordRequest, reader: jspb.BinaryReader): GetRecordRequest;
}

export namespace GetRecordRequest {
    export type AsObject = {
        token: string,
        namespace: string,
        resource: string,
        id: string,

        annotationsMap: Array<[string, string]>,
    }
}

export class GetRecordResponse extends jspb.Message { 

    hasRecord(): boolean;
    clearRecord(): void;
    getRecord(): model_record_pb.Record | undefined;
    setRecord(value?: model_record_pb.Record): GetRecordResponse;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): GetRecordResponse.AsObject;
    static toObject(includeInstance: boolean, msg: GetRecordResponse): GetRecordResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: GetRecordResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): GetRecordResponse;
    static deserializeBinaryFromReader(message: GetRecordResponse, reader: jspb.BinaryReader): GetRecordResponse;
}

export namespace GetRecordResponse {
    export type AsObject = {
        record?: model_record_pb.Record.AsObject,
    }
}
