// package: model
// file: model/resource.proto

/* tslint:disable */
/* eslint-disable */

import * as jspb from "google-protobuf";
import * as model_audit_pb from "../model/audit_pb";
import * as model_query_pb from "../model/query_pb";
import * as model_common_pb from "../model/common_pb";
import * as model_security_pb from "../model/security_pb";
import * as model_annotations_pb from "../model/annotations_pb";
import * as google_protobuf_struct_pb from "google-protobuf/google/protobuf/struct_pb";
import * as google_protobuf_descriptor_pb from "google-protobuf/google/protobuf/descriptor_pb";

export class ResourceProperty extends jspb.Message { 

    hasId(): boolean;
    clearId(): void;
    getId(): string | undefined;
    setId(value: string): ResourceProperty;
    getName(): string;
    setName(value: string): ResourceProperty;
    getType(): ResourceProperty.Type;
    setType(value: ResourceProperty.Type): ResourceProperty;

    hasTyperef(): boolean;
    clearTyperef(): void;
    getTyperef(): string | undefined;
    setTyperef(value: string): ResourceProperty;
    getMapping(): string;
    setMapping(value: string): ResourceProperty;
    getRequired(): boolean;
    setRequired(value: boolean): ResourceProperty;
    getPrimary(): boolean;
    setPrimary(value: boolean): ResourceProperty;
    getLength(): number;
    setLength(value: number): ResourceProperty;
    getUnique(): boolean;
    setUnique(value: boolean): ResourceProperty;
    getImmutable(): boolean;
    setImmutable(value: boolean): ResourceProperty;

    hasSecuritycontext(): boolean;
    clearSecuritycontext(): void;
    getSecuritycontext(): model_security_pb.SecurityContext | undefined;
    setSecuritycontext(value?: model_security_pb.SecurityContext): ResourceProperty;

    hasDefaultvalue(): boolean;
    clearDefaultvalue(): void;
    getDefaultvalue(): google_protobuf_struct_pb.Value | undefined;
    setDefaultvalue(value?: google_protobuf_struct_pb.Value): ResourceProperty;

    hasExamplevalue(): boolean;
    clearExamplevalue(): void;
    getExamplevalue(): google_protobuf_struct_pb.Value | undefined;
    setExamplevalue(value?: google_protobuf_struct_pb.Value): ResourceProperty;
    clearEnumvaluesList(): void;
    getEnumvaluesList(): Array<google_protobuf_struct_pb.Value>;
    setEnumvaluesList(value: Array<google_protobuf_struct_pb.Value>): ResourceProperty;
    addEnumvalues(value?: google_protobuf_struct_pb.Value, index?: number): google_protobuf_struct_pb.Value;

    hasReference(): boolean;
    clearReference(): void;
    getReference(): Reference | undefined;
    setReference(value?: Reference): ResourceProperty;
    clearPropertiesList(): void;
    getPropertiesList(): Array<ResourceProperty>;
    setPropertiesList(value: Array<ResourceProperty>): ResourceProperty;
    addProperties(value?: ResourceProperty, index?: number): ResourceProperty;

    hasItem(): boolean;
    clearItem(): void;
    getItem(): ResourceProperty | undefined;
    setItem(value?: ResourceProperty): ResourceProperty;

    hasTitle(): boolean;
    clearTitle(): void;
    getTitle(): string | undefined;
    setTitle(value: string): ResourceProperty;

    hasDescription(): boolean;
    clearDescription(): void;
    getDescription(): string | undefined;
    setDescription(value: string): ResourceProperty;

    getAnnotationsMap(): jspb.Map<string, string>;
    clearAnnotationsMap(): void;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ResourceProperty.AsObject;
    static toObject(includeInstance: boolean, msg: ResourceProperty): ResourceProperty.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ResourceProperty, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ResourceProperty;
    static deserializeBinaryFromReader(message: ResourceProperty, reader: jspb.BinaryReader): ResourceProperty;
}

export namespace ResourceProperty {
    export type AsObject = {
        id?: string,
        name: string,
        type: ResourceProperty.Type,
        typeref?: string,
        mapping: string,
        required: boolean,
        primary: boolean,
        length: number,
        unique: boolean,
        immutable: boolean,
        securitycontext?: model_security_pb.SecurityContext.AsObject,
        defaultvalue?: google_protobuf_struct_pb.Value.AsObject,
        examplevalue?: google_protobuf_struct_pb.Value.AsObject,
        enumvaluesList: Array<google_protobuf_struct_pb.Value.AsObject>,
        reference?: Reference.AsObject,
        propertiesList: Array<ResourceProperty.AsObject>,
        item?: ResourceProperty.AsObject,
        title?: string,
        description?: string,

        annotationsMap: Array<[string, string]>,
    }

    export enum Type {
    BOOL = 0,
    STRING = 1,
    FLOAT32 = 2,
    FLOAT64 = 3,
    INT32 = 4,
    INT64 = 5,
    BYTES = 6,
    UUID = 8,
    DATE = 9,
    TIME = 10,
    TIMESTAMP = 11,
    OBJECT = 12,
    MAP = 13,
    LIST = 14,
    REFERENCE = 15,
    ENUM = 16,
    STRUCT = 17,
    }

}

export class Reference extends jspb.Message { 
    getReferencedresource(): string;
    setReferencedresource(value: string): Reference;
    getCascade(): boolean;
    setCascade(value: boolean): Reference;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): Reference.AsObject;
    static toObject(includeInstance: boolean, msg: Reference): Reference.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: Reference, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): Reference;
    static deserializeBinaryFromReader(message: Reference, reader: jspb.BinaryReader): Reference;
}

export namespace Reference {
    export type AsObject = {
        referencedresource: string,
        cascade: boolean,
    }
}

export class ResourceSourceConfig extends jspb.Message { 
    getDatasource(): string;
    setDatasource(value: string): ResourceSourceConfig;
    getCatalog(): string;
    setCatalog(value: string): ResourceSourceConfig;
    getEntity(): string;
    setEntity(value: string): ResourceSourceConfig;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ResourceSourceConfig.AsObject;
    static toObject(includeInstance: boolean, msg: ResourceSourceConfig): ResourceSourceConfig.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ResourceSourceConfig, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ResourceSourceConfig;
    static deserializeBinaryFromReader(message: ResourceSourceConfig, reader: jspb.BinaryReader): ResourceSourceConfig;
}

export namespace ResourceSourceConfig {
    export type AsObject = {
        datasource: string,
        catalog: string,
        entity: string,
    }
}

export class ResourceIndexProperty extends jspb.Message { 
    getName(): string;
    setName(value: string): ResourceIndexProperty;
    getOrder(): Order;
    setOrder(value: Order): ResourceIndexProperty;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ResourceIndexProperty.AsObject;
    static toObject(includeInstance: boolean, msg: ResourceIndexProperty): ResourceIndexProperty.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ResourceIndexProperty, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ResourceIndexProperty;
    static deserializeBinaryFromReader(message: ResourceIndexProperty, reader: jspb.BinaryReader): ResourceIndexProperty;
}

export namespace ResourceIndexProperty {
    export type AsObject = {
        name: string,
        order: Order,
    }
}

export class ResourceIndex extends jspb.Message { 
    clearPropertiesList(): void;
    getPropertiesList(): Array<ResourceIndexProperty>;
    setPropertiesList(value: Array<ResourceIndexProperty>): ResourceIndex;
    addProperties(value?: ResourceIndexProperty, index?: number): ResourceIndexProperty;
    getIndextype(): ResourceIndexType;
    setIndextype(value: ResourceIndexType): ResourceIndex;
    getUnique(): boolean;
    setUnique(value: boolean): ResourceIndex;

    getAnnotationsMap(): jspb.Map<string, string>;
    clearAnnotationsMap(): void;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ResourceIndex.AsObject;
    static toObject(includeInstance: boolean, msg: ResourceIndex): ResourceIndex.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ResourceIndex, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ResourceIndex;
    static deserializeBinaryFromReader(message: ResourceIndex, reader: jspb.BinaryReader): ResourceIndex;
}

export namespace ResourceIndex {
    export type AsObject = {
        propertiesList: Array<ResourceIndexProperty.AsObject>,
        indextype: ResourceIndexType,
        unique: boolean,

        annotationsMap: Array<[string, string]>,
    }
}

export class ResourceSubType extends jspb.Message { 
    getName(): string;
    setName(value: string): ResourceSubType;
    clearPropertiesList(): void;
    getPropertiesList(): Array<ResourceProperty>;
    setPropertiesList(value: Array<ResourceProperty>): ResourceSubType;
    addProperties(value?: ResourceProperty, index?: number): ResourceProperty;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ResourceSubType.AsObject;
    static toObject(includeInstance: boolean, msg: ResourceSubType): ResourceSubType.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ResourceSubType, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ResourceSubType;
    static deserializeBinaryFromReader(message: ResourceSubType, reader: jspb.BinaryReader): ResourceSubType;
}

export namespace ResourceSubType {
    export type AsObject = {
        name: string,
        propertiesList: Array<ResourceProperty.AsObject>,
    }
}

export class Resource extends jspb.Message { 
    getId(): string;
    setId(value: string): Resource;
    getName(): string;
    setName(value: string): Resource;
    getNamespace(): string;
    setNamespace(value: string): Resource;

    hasSourceconfig(): boolean;
    clearSourceconfig(): void;
    getSourceconfig(): ResourceSourceConfig | undefined;
    setSourceconfig(value?: ResourceSourceConfig): Resource;
    clearPropertiesList(): void;
    getPropertiesList(): Array<ResourceProperty>;
    setPropertiesList(value: Array<ResourceProperty>): Resource;
    addProperties(value?: ResourceProperty, index?: number): ResourceProperty;
    clearTypesList(): void;
    getTypesList(): Array<ResourceSubType>;
    setTypesList(value: Array<ResourceSubType>): Resource;
    addTypes(value?: ResourceSubType, index?: number): ResourceSubType;
    clearIndexesList(): void;
    getIndexesList(): Array<ResourceIndex>;
    setIndexesList(value: Array<ResourceIndex>): Resource;
    addIndexes(value?: ResourceIndex, index?: number): ResourceIndex;

    hasSecuritycontext(): boolean;
    clearSecuritycontext(): void;
    getSecuritycontext(): model_security_pb.SecurityContext | undefined;
    setSecuritycontext(value?: model_security_pb.SecurityContext): Resource;
    getVirtual(): boolean;
    setVirtual(value: boolean): Resource;
    getImmutable(): boolean;
    setImmutable(value: boolean): Resource;
    getAbstract(): boolean;
    setAbstract(value: boolean): Resource;

    hasTitle(): boolean;
    clearTitle(): void;
    getTitle(): string | undefined;
    setTitle(value: string): Resource;

    hasDescription(): boolean;
    clearDescription(): void;
    getDescription(): string | undefined;
    setDescription(value: string): Resource;

    hasAuditdata(): boolean;
    clearAuditdata(): void;
    getAuditdata(): model_audit_pb.AuditData | undefined;
    setAuditdata(value?: model_audit_pb.AuditData): Resource;
    getVersion(): number;
    setVersion(value: number): Resource;

    getAnnotationsMap(): jspb.Map<string, string>;
    clearAnnotationsMap(): void;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): Resource.AsObject;
    static toObject(includeInstance: boolean, msg: Resource): Resource.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: Resource, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): Resource;
    static deserializeBinaryFromReader(message: Resource, reader: jspb.BinaryReader): Resource;
}

export namespace Resource {
    export type AsObject = {
        id: string,
        name: string,
        namespace: string,
        sourceconfig?: ResourceSourceConfig.AsObject,
        propertiesList: Array<ResourceProperty.AsObject>,
        typesList: Array<ResourceSubType.AsObject>,
        indexesList: Array<ResourceIndex.AsObject>,
        securitycontext?: model_security_pb.SecurityContext.AsObject,
        virtual: boolean,
        immutable: boolean,
        pb_abstract: boolean,
        title?: string,
        description?: string,
        auditdata?: model_audit_pb.AuditData.AsObject,
        version: number,

        annotationsMap: Array<[string, string]>,
    }
}

export enum Order {
    ORDER_UNKNOWN = 0,
    ORDER_ASC = 1,
    ORDER_DESC = 2,
}

export enum ResourceIndexType {
    BTREE = 0,
    HASH = 1,
}
