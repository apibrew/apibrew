// package: stub
// file: stub/data-source.proto

/* tslint:disable */
/* eslint-disable */

import * as jspb from "google-protobuf";
import * as model_data_source_pb from "../model/data-source_pb";
import * as model_error_pb from "../model/error_pb";
import * as model_resource_pb from "../model/resource_pb";


export class PrepareResourceFromEntityRequest extends jspb.Message { 
    getToken(): string;
    setToken(value: string): PrepareResourceFromEntityRequest;
    getId(): string;
    setId(value: string): PrepareResourceFromEntityRequest;
    getCatalog(): string;
    setCatalog(value: string): PrepareResourceFromEntityRequest;
    getEntity(): string;
    setEntity(value: string): PrepareResourceFromEntityRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): PrepareResourceFromEntityRequest.AsObject;
    static toObject(includeInstance: boolean, msg: PrepareResourceFromEntityRequest): PrepareResourceFromEntityRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: PrepareResourceFromEntityRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): PrepareResourceFromEntityRequest;
    static deserializeBinaryFromReader(message: PrepareResourceFromEntityRequest, reader: jspb.BinaryReader): PrepareResourceFromEntityRequest;
}

export namespace PrepareResourceFromEntityRequest {
    export type AsObject = {
        token: string,
        id: string,
        catalog: string,
        entity: string,
    }
}

export class PrepareResourceFromEntityResponse extends jspb.Message { 

    hasResource(): boolean;
    clearResource(): void;
    getResource(): model_resource_pb.Resource | undefined;
    setResource(value?: model_resource_pb.Resource): PrepareResourceFromEntityResponse;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): PrepareResourceFromEntityResponse.AsObject;
    static toObject(includeInstance: boolean, msg: PrepareResourceFromEntityResponse): PrepareResourceFromEntityResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: PrepareResourceFromEntityResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): PrepareResourceFromEntityResponse;
    static deserializeBinaryFromReader(message: PrepareResourceFromEntityResponse, reader: jspb.BinaryReader): PrepareResourceFromEntityResponse;
}

export namespace PrepareResourceFromEntityResponse {
    export type AsObject = {
        resource?: model_resource_pb.Resource.AsObject,
    }
}

export class StatusRequest extends jspb.Message { 
    getToken(): string;
    setToken(value: string): StatusRequest;
    getId(): string;
    setId(value: string): StatusRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): StatusRequest.AsObject;
    static toObject(includeInstance: boolean, msg: StatusRequest): StatusRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: StatusRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): StatusRequest;
    static deserializeBinaryFromReader(message: StatusRequest, reader: jspb.BinaryReader): StatusRequest;
}

export namespace StatusRequest {
    export type AsObject = {
        token: string,
        id: string,
    }
}

export class StatusResponse extends jspb.Message { 
    getConnectionalreadyinitiated(): boolean;
    setConnectionalreadyinitiated(value: boolean): StatusResponse;
    getTestconnection(): boolean;
    setTestconnection(value: boolean): StatusResponse;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): StatusResponse.AsObject;
    static toObject(includeInstance: boolean, msg: StatusResponse): StatusResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: StatusResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): StatusResponse;
    static deserializeBinaryFromReader(message: StatusResponse, reader: jspb.BinaryReader): StatusResponse;
}

export namespace StatusResponse {
    export type AsObject = {
        connectionalreadyinitiated: boolean,
        testconnection: boolean,
    }
}

export class ListEntitiesRequest extends jspb.Message { 
    getToken(): string;
    setToken(value: string): ListEntitiesRequest;
    getId(): string;
    setId(value: string): ListEntitiesRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ListEntitiesRequest.AsObject;
    static toObject(includeInstance: boolean, msg: ListEntitiesRequest): ListEntitiesRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ListEntitiesRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ListEntitiesRequest;
    static deserializeBinaryFromReader(message: ListEntitiesRequest, reader: jspb.BinaryReader): ListEntitiesRequest;
}

export namespace ListEntitiesRequest {
    export type AsObject = {
        token: string,
        id: string,
    }
}

export class ListEntitiesResponse extends jspb.Message { 
    clearCatalogsList(): void;
    getCatalogsList(): Array<model_data_source_pb.DataSourceCatalog>;
    setCatalogsList(value: Array<model_data_source_pb.DataSourceCatalog>): ListEntitiesResponse;
    addCatalogs(value?: model_data_source_pb.DataSourceCatalog, index?: number): model_data_source_pb.DataSourceCatalog;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ListEntitiesResponse.AsObject;
    static toObject(includeInstance: boolean, msg: ListEntitiesResponse): ListEntitiesResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ListEntitiesResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ListEntitiesResponse;
    static deserializeBinaryFromReader(message: ListEntitiesResponse, reader: jspb.BinaryReader): ListEntitiesResponse;
}

export namespace ListEntitiesResponse {
    export type AsObject = {
        catalogsList: Array<model_data_source_pb.DataSourceCatalog.AsObject>,
    }
}

export class ListDataSourceRequest extends jspb.Message { 
    getToken(): string;
    setToken(value: string): ListDataSourceRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ListDataSourceRequest.AsObject;
    static toObject(includeInstance: boolean, msg: ListDataSourceRequest): ListDataSourceRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ListDataSourceRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ListDataSourceRequest;
    static deserializeBinaryFromReader(message: ListDataSourceRequest, reader: jspb.BinaryReader): ListDataSourceRequest;
}

export namespace ListDataSourceRequest {
    export type AsObject = {
        token: string,
    }
}

export class ListDataSourceResponse extends jspb.Message { 
    clearContentList(): void;
    getContentList(): Array<model_data_source_pb.DataSource>;
    setContentList(value: Array<model_data_source_pb.DataSource>): ListDataSourceResponse;
    addContent(value?: model_data_source_pb.DataSource, index?: number): model_data_source_pb.DataSource;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ListDataSourceResponse.AsObject;
    static toObject(includeInstance: boolean, msg: ListDataSourceResponse): ListDataSourceResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ListDataSourceResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ListDataSourceResponse;
    static deserializeBinaryFromReader(message: ListDataSourceResponse, reader: jspb.BinaryReader): ListDataSourceResponse;
}

export namespace ListDataSourceResponse {
    export type AsObject = {
        contentList: Array<model_data_source_pb.DataSource.AsObject>,
    }
}

export class CreateDataSourceRequest extends jspb.Message { 
    getToken(): string;
    setToken(value: string): CreateDataSourceRequest;
    clearDatasourcesList(): void;
    getDatasourcesList(): Array<model_data_source_pb.DataSource>;
    setDatasourcesList(value: Array<model_data_source_pb.DataSource>): CreateDataSourceRequest;
    addDatasources(value?: model_data_source_pb.DataSource, index?: number): model_data_source_pb.DataSource;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): CreateDataSourceRequest.AsObject;
    static toObject(includeInstance: boolean, msg: CreateDataSourceRequest): CreateDataSourceRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: CreateDataSourceRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): CreateDataSourceRequest;
    static deserializeBinaryFromReader(message: CreateDataSourceRequest, reader: jspb.BinaryReader): CreateDataSourceRequest;
}

export namespace CreateDataSourceRequest {
    export type AsObject = {
        token: string,
        datasourcesList: Array<model_data_source_pb.DataSource.AsObject>,
    }
}

export class CreateDataSourceResponse extends jspb.Message { 
    clearDatasourcesList(): void;
    getDatasourcesList(): Array<model_data_source_pb.DataSource>;
    setDatasourcesList(value: Array<model_data_source_pb.DataSource>): CreateDataSourceResponse;
    addDatasources(value?: model_data_source_pb.DataSource, index?: number): model_data_source_pb.DataSource;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): CreateDataSourceResponse.AsObject;
    static toObject(includeInstance: boolean, msg: CreateDataSourceResponse): CreateDataSourceResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: CreateDataSourceResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): CreateDataSourceResponse;
    static deserializeBinaryFromReader(message: CreateDataSourceResponse, reader: jspb.BinaryReader): CreateDataSourceResponse;
}

export namespace CreateDataSourceResponse {
    export type AsObject = {
        datasourcesList: Array<model_data_source_pb.DataSource.AsObject>,
    }
}

export class UpdateDataSourceRequest extends jspb.Message { 
    getToken(): string;
    setToken(value: string): UpdateDataSourceRequest;
    clearDatasourcesList(): void;
    getDatasourcesList(): Array<model_data_source_pb.DataSource>;
    setDatasourcesList(value: Array<model_data_source_pb.DataSource>): UpdateDataSourceRequest;
    addDatasources(value?: model_data_source_pb.DataSource, index?: number): model_data_source_pb.DataSource;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): UpdateDataSourceRequest.AsObject;
    static toObject(includeInstance: boolean, msg: UpdateDataSourceRequest): UpdateDataSourceRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: UpdateDataSourceRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): UpdateDataSourceRequest;
    static deserializeBinaryFromReader(message: UpdateDataSourceRequest, reader: jspb.BinaryReader): UpdateDataSourceRequest;
}

export namespace UpdateDataSourceRequest {
    export type AsObject = {
        token: string,
        datasourcesList: Array<model_data_source_pb.DataSource.AsObject>,
    }
}

export class UpdateDataSourceResponse extends jspb.Message { 
    clearDatasourcesList(): void;
    getDatasourcesList(): Array<model_data_source_pb.DataSource>;
    setDatasourcesList(value: Array<model_data_source_pb.DataSource>): UpdateDataSourceResponse;
    addDatasources(value?: model_data_source_pb.DataSource, index?: number): model_data_source_pb.DataSource;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): UpdateDataSourceResponse.AsObject;
    static toObject(includeInstance: boolean, msg: UpdateDataSourceResponse): UpdateDataSourceResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: UpdateDataSourceResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): UpdateDataSourceResponse;
    static deserializeBinaryFromReader(message: UpdateDataSourceResponse, reader: jspb.BinaryReader): UpdateDataSourceResponse;
}

export namespace UpdateDataSourceResponse {
    export type AsObject = {
        datasourcesList: Array<model_data_source_pb.DataSource.AsObject>,
    }
}

export class DeleteDataSourceRequest extends jspb.Message { 
    getToken(): string;
    setToken(value: string): DeleteDataSourceRequest;
    clearIdsList(): void;
    getIdsList(): Array<string>;
    setIdsList(value: Array<string>): DeleteDataSourceRequest;
    addIds(value: string, index?: number): string;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): DeleteDataSourceRequest.AsObject;
    static toObject(includeInstance: boolean, msg: DeleteDataSourceRequest): DeleteDataSourceRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: DeleteDataSourceRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): DeleteDataSourceRequest;
    static deserializeBinaryFromReader(message: DeleteDataSourceRequest, reader: jspb.BinaryReader): DeleteDataSourceRequest;
}

export namespace DeleteDataSourceRequest {
    export type AsObject = {
        token: string,
        idsList: Array<string>,
    }
}

export class DeleteDataSourceResponse extends jspb.Message { 

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): DeleteDataSourceResponse.AsObject;
    static toObject(includeInstance: boolean, msg: DeleteDataSourceResponse): DeleteDataSourceResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: DeleteDataSourceResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): DeleteDataSourceResponse;
    static deserializeBinaryFromReader(message: DeleteDataSourceResponse, reader: jspb.BinaryReader): DeleteDataSourceResponse;
}

export namespace DeleteDataSourceResponse {
    export type AsObject = {
    }
}

export class GetDataSourceRequest extends jspb.Message { 
    getToken(): string;
    setToken(value: string): GetDataSourceRequest;
    getId(): string;
    setId(value: string): GetDataSourceRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): GetDataSourceRequest.AsObject;
    static toObject(includeInstance: boolean, msg: GetDataSourceRequest): GetDataSourceRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: GetDataSourceRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): GetDataSourceRequest;
    static deserializeBinaryFromReader(message: GetDataSourceRequest, reader: jspb.BinaryReader): GetDataSourceRequest;
}

export namespace GetDataSourceRequest {
    export type AsObject = {
        token: string,
        id: string,
    }
}

export class GetDataSourceResponse extends jspb.Message { 

    hasDatasource(): boolean;
    clearDatasource(): void;
    getDatasource(): model_data_source_pb.DataSource | undefined;
    setDatasource(value?: model_data_source_pb.DataSource): GetDataSourceResponse;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): GetDataSourceResponse.AsObject;
    static toObject(includeInstance: boolean, msg: GetDataSourceResponse): GetDataSourceResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: GetDataSourceResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): GetDataSourceResponse;
    static deserializeBinaryFromReader(message: GetDataSourceResponse, reader: jspb.BinaryReader): GetDataSourceResponse;
}

export namespace GetDataSourceResponse {
    export type AsObject = {
        datasource?: model_data_source_pb.DataSource.AsObject,
    }
}
