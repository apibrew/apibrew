// package: model
// file: model/resource-migration.proto

/* tslint:disable */
/* eslint-disable */

import * as jspb from "google-protobuf";
import * as model_annotations_pb from "../model/annotations_pb";
import * as model_resource_pb from "../model/resource_pb";
import * as google_protobuf_struct_pb from "google-protobuf/google/protobuf/struct_pb";
import * as google_protobuf_descriptor_pb from "google-protobuf/google/protobuf/descriptor_pb";

export class ResourceMigrationCreateResource extends jspb.Message { 

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ResourceMigrationCreateResource.AsObject;
    static toObject(includeInstance: boolean, msg: ResourceMigrationCreateResource): ResourceMigrationCreateResource.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ResourceMigrationCreateResource, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ResourceMigrationCreateResource;
    static deserializeBinaryFromReader(message: ResourceMigrationCreateResource, reader: jspb.BinaryReader): ResourceMigrationCreateResource;
}

export namespace ResourceMigrationCreateResource {
    export type AsObject = {
    }
}

export class ResourceMigrationDeleteResource extends jspb.Message { 

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ResourceMigrationDeleteResource.AsObject;
    static toObject(includeInstance: boolean, msg: ResourceMigrationDeleteResource): ResourceMigrationDeleteResource.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ResourceMigrationDeleteResource, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ResourceMigrationDeleteResource;
    static deserializeBinaryFromReader(message: ResourceMigrationDeleteResource, reader: jspb.BinaryReader): ResourceMigrationDeleteResource;
}

export namespace ResourceMigrationDeleteResource {
    export type AsObject = {
    }
}

export class ResourceMigrationUpdateResource extends jspb.Message { 
    clearChangedfieldsList(): void;
    getChangedfieldsList(): Array<string>;
    setChangedfieldsList(value: Array<string>): ResourceMigrationUpdateResource;
    addChangedfields(value: string, index?: number): string;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ResourceMigrationUpdateResource.AsObject;
    static toObject(includeInstance: boolean, msg: ResourceMigrationUpdateResource): ResourceMigrationUpdateResource.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ResourceMigrationUpdateResource, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ResourceMigrationUpdateResource;
    static deserializeBinaryFromReader(message: ResourceMigrationUpdateResource, reader: jspb.BinaryReader): ResourceMigrationUpdateResource;
}

export namespace ResourceMigrationUpdateResource {
    export type AsObject = {
        changedfieldsList: Array<string>,
    }
}

export class ResourceMigrationCreateProperty extends jspb.Message { 
    getProperty(): string;
    setProperty(value: string): ResourceMigrationCreateProperty;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ResourceMigrationCreateProperty.AsObject;
    static toObject(includeInstance: boolean, msg: ResourceMigrationCreateProperty): ResourceMigrationCreateProperty.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ResourceMigrationCreateProperty, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ResourceMigrationCreateProperty;
    static deserializeBinaryFromReader(message: ResourceMigrationCreateProperty, reader: jspb.BinaryReader): ResourceMigrationCreateProperty;
}

export namespace ResourceMigrationCreateProperty {
    export type AsObject = {
        property: string,
    }
}

export class ResourceMigrationDeleteProperty extends jspb.Message { 
    getExistingproperty(): string;
    setExistingproperty(value: string): ResourceMigrationDeleteProperty;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ResourceMigrationDeleteProperty.AsObject;
    static toObject(includeInstance: boolean, msg: ResourceMigrationDeleteProperty): ResourceMigrationDeleteProperty.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ResourceMigrationDeleteProperty, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ResourceMigrationDeleteProperty;
    static deserializeBinaryFromReader(message: ResourceMigrationDeleteProperty, reader: jspb.BinaryReader): ResourceMigrationDeleteProperty;
}

export namespace ResourceMigrationDeleteProperty {
    export type AsObject = {
        existingproperty: string,
    }
}

export class ResourceMigrationUpdateProperty extends jspb.Message { 
    getExistingproperty(): string;
    setExistingproperty(value: string): ResourceMigrationUpdateProperty;
    getProperty(): string;
    setProperty(value: string): ResourceMigrationUpdateProperty;
    clearChangedfieldsList(): void;
    getChangedfieldsList(): Array<string>;
    setChangedfieldsList(value: Array<string>): ResourceMigrationUpdateProperty;
    addChangedfields(value: string, index?: number): string;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ResourceMigrationUpdateProperty.AsObject;
    static toObject(includeInstance: boolean, msg: ResourceMigrationUpdateProperty): ResourceMigrationUpdateProperty.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ResourceMigrationUpdateProperty, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ResourceMigrationUpdateProperty;
    static deserializeBinaryFromReader(message: ResourceMigrationUpdateProperty, reader: jspb.BinaryReader): ResourceMigrationUpdateProperty;
}

export namespace ResourceMigrationUpdateProperty {
    export type AsObject = {
        existingproperty: string,
        property: string,
        changedfieldsList: Array<string>,
    }
}

export class ResourceMigrationCreateIndex extends jspb.Message { 
    getIndex(): number;
    setIndex(value: number): ResourceMigrationCreateIndex;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ResourceMigrationCreateIndex.AsObject;
    static toObject(includeInstance: boolean, msg: ResourceMigrationCreateIndex): ResourceMigrationCreateIndex.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ResourceMigrationCreateIndex, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ResourceMigrationCreateIndex;
    static deserializeBinaryFromReader(message: ResourceMigrationCreateIndex, reader: jspb.BinaryReader): ResourceMigrationCreateIndex;
}

export namespace ResourceMigrationCreateIndex {
    export type AsObject = {
        index: number,
    }
}

export class ResourceMigrationDeleteIndex extends jspb.Message { 
    getExistingindex(): number;
    setExistingindex(value: number): ResourceMigrationDeleteIndex;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ResourceMigrationDeleteIndex.AsObject;
    static toObject(includeInstance: boolean, msg: ResourceMigrationDeleteIndex): ResourceMigrationDeleteIndex.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ResourceMigrationDeleteIndex, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ResourceMigrationDeleteIndex;
    static deserializeBinaryFromReader(message: ResourceMigrationDeleteIndex, reader: jspb.BinaryReader): ResourceMigrationDeleteIndex;
}

export namespace ResourceMigrationDeleteIndex {
    export type AsObject = {
        existingindex: number,
    }
}

export class ResourceMigrationStep extends jspb.Message { 

    hasCreateresource(): boolean;
    clearCreateresource(): void;
    getCreateresource(): ResourceMigrationCreateResource | undefined;
    setCreateresource(value?: ResourceMigrationCreateResource): ResourceMigrationStep;

    hasDeleteresource(): boolean;
    clearDeleteresource(): void;
    getDeleteresource(): ResourceMigrationDeleteResource | undefined;
    setDeleteresource(value?: ResourceMigrationDeleteResource): ResourceMigrationStep;

    hasUpdateresource(): boolean;
    clearUpdateresource(): void;
    getUpdateresource(): ResourceMigrationUpdateResource | undefined;
    setUpdateresource(value?: ResourceMigrationUpdateResource): ResourceMigrationStep;

    hasCreateproperty(): boolean;
    clearCreateproperty(): void;
    getCreateproperty(): ResourceMigrationCreateProperty | undefined;
    setCreateproperty(value?: ResourceMigrationCreateProperty): ResourceMigrationStep;

    hasDeleteproperty(): boolean;
    clearDeleteproperty(): void;
    getDeleteproperty(): ResourceMigrationDeleteProperty | undefined;
    setDeleteproperty(value?: ResourceMigrationDeleteProperty): ResourceMigrationStep;

    hasUpdateproperty(): boolean;
    clearUpdateproperty(): void;
    getUpdateproperty(): ResourceMigrationUpdateProperty | undefined;
    setUpdateproperty(value?: ResourceMigrationUpdateProperty): ResourceMigrationStep;

    hasCreateindex(): boolean;
    clearCreateindex(): void;
    getCreateindex(): ResourceMigrationCreateIndex | undefined;
    setCreateindex(value?: ResourceMigrationCreateIndex): ResourceMigrationStep;

    hasDeleteindex(): boolean;
    clearDeleteindex(): void;
    getDeleteindex(): ResourceMigrationDeleteIndex | undefined;
    setDeleteindex(value?: ResourceMigrationDeleteIndex): ResourceMigrationStep;

    getKindCase(): ResourceMigrationStep.KindCase;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ResourceMigrationStep.AsObject;
    static toObject(includeInstance: boolean, msg: ResourceMigrationStep): ResourceMigrationStep.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ResourceMigrationStep, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ResourceMigrationStep;
    static deserializeBinaryFromReader(message: ResourceMigrationStep, reader: jspb.BinaryReader): ResourceMigrationStep;
}

export namespace ResourceMigrationStep {
    export type AsObject = {
        createresource?: ResourceMigrationCreateResource.AsObject,
        deleteresource?: ResourceMigrationDeleteResource.AsObject,
        updateresource?: ResourceMigrationUpdateResource.AsObject,
        createproperty?: ResourceMigrationCreateProperty.AsObject,
        deleteproperty?: ResourceMigrationDeleteProperty.AsObject,
        updateproperty?: ResourceMigrationUpdateProperty.AsObject,
        createindex?: ResourceMigrationCreateIndex.AsObject,
        deleteindex?: ResourceMigrationDeleteIndex.AsObject,
    }

    export enum KindCase {
        KIND_NOT_SET = 0,
        CREATERESOURCE = 1,
        DELETERESOURCE = 2,
        UPDATERESOURCE = 3,
        CREATEPROPERTY = 4,
        DELETEPROPERTY = 5,
        UPDATEPROPERTY = 6,
        CREATEINDEX = 7,
        DELETEINDEX = 8,
    }

}

export class ResourceMigrationPlan extends jspb.Message { 

    hasExistingresource(): boolean;
    clearExistingresource(): void;
    getExistingresource(): model_resource_pb.Resource | undefined;
    setExistingresource(value?: model_resource_pb.Resource): ResourceMigrationPlan;

    hasCurrentresource(): boolean;
    clearCurrentresource(): void;
    getCurrentresource(): model_resource_pb.Resource | undefined;
    setCurrentresource(value?: model_resource_pb.Resource): ResourceMigrationPlan;
    clearStepsList(): void;
    getStepsList(): Array<ResourceMigrationStep>;
    setStepsList(value: Array<ResourceMigrationStep>): ResourceMigrationPlan;
    addSteps(value?: ResourceMigrationStep, index?: number): ResourceMigrationStep;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ResourceMigrationPlan.AsObject;
    static toObject(includeInstance: boolean, msg: ResourceMigrationPlan): ResourceMigrationPlan.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ResourceMigrationPlan, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ResourceMigrationPlan;
    static deserializeBinaryFromReader(message: ResourceMigrationPlan, reader: jspb.BinaryReader): ResourceMigrationPlan;
}

export namespace ResourceMigrationPlan {
    export type AsObject = {
        existingresource?: model_resource_pb.Resource.AsObject,
        currentresource?: model_resource_pb.Resource.AsObject,
        stepsList: Array<ResourceMigrationStep.AsObject>,
    }
}
