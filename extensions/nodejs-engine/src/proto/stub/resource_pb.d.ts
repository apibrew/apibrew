// package: stub
// file: stub/resource.proto

/* tslint:disable */
/* eslint-disable */

import * as jspb from "google-protobuf";
import * as model_error_pb from "../model/error_pb";
import * as model_resource_pb from "../model/resource_pb";
import * as model_resource_migration_pb from "../model/resource-migration_pb";


export class PrepareResourceMigrationPlanRequest extends jspb.Message { 
    getToken(): string;
    setToken(value: string): PrepareResourceMigrationPlanRequest;
    getPreparefromdatasource(): boolean;
    setPreparefromdatasource(value: boolean): PrepareResourceMigrationPlanRequest;
    clearResourcesList(): void;
    getResourcesList(): Array<model_resource_pb.Resource>;
    setResourcesList(value: Array<model_resource_pb.Resource>): PrepareResourceMigrationPlanRequest;
    addResources(value?: model_resource_pb.Resource, index?: number): model_resource_pb.Resource;

    getAnnotationsMap(): jspb.Map<string, string>;
    clearAnnotationsMap(): void;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): PrepareResourceMigrationPlanRequest.AsObject;
    static toObject(includeInstance: boolean, msg: PrepareResourceMigrationPlanRequest): PrepareResourceMigrationPlanRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: PrepareResourceMigrationPlanRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): PrepareResourceMigrationPlanRequest;
    static deserializeBinaryFromReader(message: PrepareResourceMigrationPlanRequest, reader: jspb.BinaryReader): PrepareResourceMigrationPlanRequest;
}

export namespace PrepareResourceMigrationPlanRequest {
    export type AsObject = {
        token: string,
        preparefromdatasource: boolean,
        resourcesList: Array<model_resource_pb.Resource.AsObject>,

        annotationsMap: Array<[string, string]>,
    }
}

export class PrepareResourceMigrationPlanResponse extends jspb.Message { 
    clearPlansList(): void;
    getPlansList(): Array<model_resource_migration_pb.ResourceMigrationPlan>;
    setPlansList(value: Array<model_resource_migration_pb.ResourceMigrationPlan>): PrepareResourceMigrationPlanResponse;
    addPlans(value?: model_resource_migration_pb.ResourceMigrationPlan, index?: number): model_resource_migration_pb.ResourceMigrationPlan;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): PrepareResourceMigrationPlanResponse.AsObject;
    static toObject(includeInstance: boolean, msg: PrepareResourceMigrationPlanResponse): PrepareResourceMigrationPlanResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: PrepareResourceMigrationPlanResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): PrepareResourceMigrationPlanResponse;
    static deserializeBinaryFromReader(message: PrepareResourceMigrationPlanResponse, reader: jspb.BinaryReader): PrepareResourceMigrationPlanResponse;
}

export namespace PrepareResourceMigrationPlanResponse {
    export type AsObject = {
        plansList: Array<model_resource_migration_pb.ResourceMigrationPlan.AsObject>,
    }
}

export class CreateResourceRequest extends jspb.Message { 
    getToken(): string;
    setToken(value: string): CreateResourceRequest;
    clearResourcesList(): void;
    getResourcesList(): Array<model_resource_pb.Resource>;
    setResourcesList(value: Array<model_resource_pb.Resource>): CreateResourceRequest;
    addResources(value?: model_resource_pb.Resource, index?: number): model_resource_pb.Resource;
    getDomigration(): boolean;
    setDomigration(value: boolean): CreateResourceRequest;
    getForcemigration(): boolean;
    setForcemigration(value: boolean): CreateResourceRequest;

    getAnnotationsMap(): jspb.Map<string, string>;
    clearAnnotationsMap(): void;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): CreateResourceRequest.AsObject;
    static toObject(includeInstance: boolean, msg: CreateResourceRequest): CreateResourceRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: CreateResourceRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): CreateResourceRequest;
    static deserializeBinaryFromReader(message: CreateResourceRequest, reader: jspb.BinaryReader): CreateResourceRequest;
}

export namespace CreateResourceRequest {
    export type AsObject = {
        token: string,
        resourcesList: Array<model_resource_pb.Resource.AsObject>,
        domigration: boolean,
        forcemigration: boolean,

        annotationsMap: Array<[string, string]>,
    }
}

export class CreateResourceResponse extends jspb.Message { 
    clearResourcesList(): void;
    getResourcesList(): Array<model_resource_pb.Resource>;
    setResourcesList(value: Array<model_resource_pb.Resource>): CreateResourceResponse;
    addResources(value?: model_resource_pb.Resource, index?: number): model_resource_pb.Resource;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): CreateResourceResponse.AsObject;
    static toObject(includeInstance: boolean, msg: CreateResourceResponse): CreateResourceResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: CreateResourceResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): CreateResourceResponse;
    static deserializeBinaryFromReader(message: CreateResourceResponse, reader: jspb.BinaryReader): CreateResourceResponse;
}

export namespace CreateResourceResponse {
    export type AsObject = {
        resourcesList: Array<model_resource_pb.Resource.AsObject>,
    }
}

export class UpdateResourceRequest extends jspb.Message { 
    getToken(): string;
    setToken(value: string): UpdateResourceRequest;
    clearResourcesList(): void;
    getResourcesList(): Array<model_resource_pb.Resource>;
    setResourcesList(value: Array<model_resource_pb.Resource>): UpdateResourceRequest;
    addResources(value?: model_resource_pb.Resource, index?: number): model_resource_pb.Resource;
    getDomigration(): boolean;
    setDomigration(value: boolean): UpdateResourceRequest;
    getForcemigration(): boolean;
    setForcemigration(value: boolean): UpdateResourceRequest;

    getAnnotationsMap(): jspb.Map<string, string>;
    clearAnnotationsMap(): void;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): UpdateResourceRequest.AsObject;
    static toObject(includeInstance: boolean, msg: UpdateResourceRequest): UpdateResourceRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: UpdateResourceRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): UpdateResourceRequest;
    static deserializeBinaryFromReader(message: UpdateResourceRequest, reader: jspb.BinaryReader): UpdateResourceRequest;
}

export namespace UpdateResourceRequest {
    export type AsObject = {
        token: string,
        resourcesList: Array<model_resource_pb.Resource.AsObject>,
        domigration: boolean,
        forcemigration: boolean,

        annotationsMap: Array<[string, string]>,
    }
}

export class UpdateResourceResponse extends jspb.Message { 
    clearResourcesList(): void;
    getResourcesList(): Array<model_resource_pb.Resource>;
    setResourcesList(value: Array<model_resource_pb.Resource>): UpdateResourceResponse;
    addResources(value?: model_resource_pb.Resource, index?: number): model_resource_pb.Resource;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): UpdateResourceResponse.AsObject;
    static toObject(includeInstance: boolean, msg: UpdateResourceResponse): UpdateResourceResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: UpdateResourceResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): UpdateResourceResponse;
    static deserializeBinaryFromReader(message: UpdateResourceResponse, reader: jspb.BinaryReader): UpdateResourceResponse;
}

export namespace UpdateResourceResponse {
    export type AsObject = {
        resourcesList: Array<model_resource_pb.Resource.AsObject>,
    }
}

export class DeleteResourceRequest extends jspb.Message { 
    getToken(): string;
    setToken(value: string): DeleteResourceRequest;
    clearIdsList(): void;
    getIdsList(): Array<string>;
    setIdsList(value: Array<string>): DeleteResourceRequest;
    addIds(value: string, index?: number): string;
    getDomigration(): boolean;
    setDomigration(value: boolean): DeleteResourceRequest;
    getForcemigration(): boolean;
    setForcemigration(value: boolean): DeleteResourceRequest;

    getAnnotationsMap(): jspb.Map<string, string>;
    clearAnnotationsMap(): void;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): DeleteResourceRequest.AsObject;
    static toObject(includeInstance: boolean, msg: DeleteResourceRequest): DeleteResourceRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: DeleteResourceRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): DeleteResourceRequest;
    static deserializeBinaryFromReader(message: DeleteResourceRequest, reader: jspb.BinaryReader): DeleteResourceRequest;
}

export namespace DeleteResourceRequest {
    export type AsObject = {
        token: string,
        idsList: Array<string>,
        domigration: boolean,
        forcemigration: boolean,

        annotationsMap: Array<[string, string]>,
    }
}

export class DeleteResourceResponse extends jspb.Message { 

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): DeleteResourceResponse.AsObject;
    static toObject(includeInstance: boolean, msg: DeleteResourceResponse): DeleteResourceResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: DeleteResourceResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): DeleteResourceResponse;
    static deserializeBinaryFromReader(message: DeleteResourceResponse, reader: jspb.BinaryReader): DeleteResourceResponse;
}

export namespace DeleteResourceResponse {
    export type AsObject = {
    }
}

export class ListResourceRequest extends jspb.Message { 
    getToken(): string;
    setToken(value: string): ListResourceRequest;

    getAnnotationsMap(): jspb.Map<string, string>;
    clearAnnotationsMap(): void;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ListResourceRequest.AsObject;
    static toObject(includeInstance: boolean, msg: ListResourceRequest): ListResourceRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ListResourceRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ListResourceRequest;
    static deserializeBinaryFromReader(message: ListResourceRequest, reader: jspb.BinaryReader): ListResourceRequest;
}

export namespace ListResourceRequest {
    export type AsObject = {
        token: string,

        annotationsMap: Array<[string, string]>,
    }
}

export class ListResourceResponse extends jspb.Message { 
    clearResourcesList(): void;
    getResourcesList(): Array<model_resource_pb.Resource>;
    setResourcesList(value: Array<model_resource_pb.Resource>): ListResourceResponse;
    addResources(value?: model_resource_pb.Resource, index?: number): model_resource_pb.Resource;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ListResourceResponse.AsObject;
    static toObject(includeInstance: boolean, msg: ListResourceResponse): ListResourceResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ListResourceResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ListResourceResponse;
    static deserializeBinaryFromReader(message: ListResourceResponse, reader: jspb.BinaryReader): ListResourceResponse;
}

export namespace ListResourceResponse {
    export type AsObject = {
        resourcesList: Array<model_resource_pb.Resource.AsObject>,
    }
}

export class GetResourceRequest extends jspb.Message { 
    getToken(): string;
    setToken(value: string): GetResourceRequest;
    getId(): string;
    setId(value: string): GetResourceRequest;

    getAnnotationsMap(): jspb.Map<string, string>;
    clearAnnotationsMap(): void;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): GetResourceRequest.AsObject;
    static toObject(includeInstance: boolean, msg: GetResourceRequest): GetResourceRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: GetResourceRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): GetResourceRequest;
    static deserializeBinaryFromReader(message: GetResourceRequest, reader: jspb.BinaryReader): GetResourceRequest;
}

export namespace GetResourceRequest {
    export type AsObject = {
        token: string,
        id: string,

        annotationsMap: Array<[string, string]>,
    }
}

export class GetResourceResponse extends jspb.Message { 

    hasResource(): boolean;
    clearResource(): void;
    getResource(): model_resource_pb.Resource | undefined;
    setResource(value?: model_resource_pb.Resource): GetResourceResponse;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): GetResourceResponse.AsObject;
    static toObject(includeInstance: boolean, msg: GetResourceResponse): GetResourceResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: GetResourceResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): GetResourceResponse;
    static deserializeBinaryFromReader(message: GetResourceResponse, reader: jspb.BinaryReader): GetResourceResponse;
}

export namespace GetResourceResponse {
    export type AsObject = {
        resource?: model_resource_pb.Resource.AsObject,
    }
}

export class GetResourceByNameRequest extends jspb.Message { 
    getToken(): string;
    setToken(value: string): GetResourceByNameRequest;
    getNamespace(): string;
    setNamespace(value: string): GetResourceByNameRequest;
    getName(): string;
    setName(value: string): GetResourceByNameRequest;

    getAnnotationsMap(): jspb.Map<string, string>;
    clearAnnotationsMap(): void;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): GetResourceByNameRequest.AsObject;
    static toObject(includeInstance: boolean, msg: GetResourceByNameRequest): GetResourceByNameRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: GetResourceByNameRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): GetResourceByNameRequest;
    static deserializeBinaryFromReader(message: GetResourceByNameRequest, reader: jspb.BinaryReader): GetResourceByNameRequest;
}

export namespace GetResourceByNameRequest {
    export type AsObject = {
        token: string,
        namespace: string,
        name: string,

        annotationsMap: Array<[string, string]>,
    }
}

export class GetResourceByNameResponse extends jspb.Message { 

    hasResource(): boolean;
    clearResource(): void;
    getResource(): model_resource_pb.Resource | undefined;
    setResource(value?: model_resource_pb.Resource): GetResourceByNameResponse;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): GetResourceByNameResponse.AsObject;
    static toObject(includeInstance: boolean, msg: GetResourceByNameResponse): GetResourceByNameResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: GetResourceByNameResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): GetResourceByNameResponse;
    static deserializeBinaryFromReader(message: GetResourceByNameResponse, reader: jspb.BinaryReader): GetResourceByNameResponse;
}

export namespace GetResourceByNameResponse {
    export type AsObject = {
        resource?: model_resource_pb.Resource.AsObject,
    }
}

export class GetSystemResourceRequest extends jspb.Message { 
    getToken(): string;
    setToken(value: string): GetSystemResourceRequest;
    getName(): string;
    setName(value: string): GetSystemResourceRequest;

    getAnnotationsMap(): jspb.Map<string, string>;
    clearAnnotationsMap(): void;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): GetSystemResourceRequest.AsObject;
    static toObject(includeInstance: boolean, msg: GetSystemResourceRequest): GetSystemResourceRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: GetSystemResourceRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): GetSystemResourceRequest;
    static deserializeBinaryFromReader(message: GetSystemResourceRequest, reader: jspb.BinaryReader): GetSystemResourceRequest;
}

export namespace GetSystemResourceRequest {
    export type AsObject = {
        token: string,
        name: string,

        annotationsMap: Array<[string, string]>,
    }
}

export class GetSystemResourceResponse extends jspb.Message { 

    hasResource(): boolean;
    clearResource(): void;
    getResource(): model_resource_pb.Resource | undefined;
    setResource(value?: model_resource_pb.Resource): GetSystemResourceResponse;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): GetSystemResourceResponse.AsObject;
    static toObject(includeInstance: boolean, msg: GetSystemResourceResponse): GetSystemResourceResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: GetSystemResourceResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): GetSystemResourceResponse;
    static deserializeBinaryFromReader(message: GetSystemResourceResponse, reader: jspb.BinaryReader): GetSystemResourceResponse;
}

export namespace GetSystemResourceResponse {
    export type AsObject = {
        resource?: model_resource_pb.Resource.AsObject,
    }
}
