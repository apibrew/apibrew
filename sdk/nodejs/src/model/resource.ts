import {Namespace} from './namespace';
import {DataSource} from './data-source';

export interface Resource {
    id: string
    version: number
    auditData?: AuditData
    name: string
    namespace: Namespace
    virtual: boolean
    properties: Property[]
    indexes?: Index[]
    types?: SubType[]
    immutable: boolean
    abstract: boolean
    checkReferences: boolean
    dataSource?: DataSource
    entity?: string
    catalog?: string
    title?: string
    description?: string
    annotations?: { [key: string]: string }
}

export const ResourceEntityInfo = {
    namespace: "system",
    resource: "Resource",
    restPath: "resources",
}

export interface Property {
    name: string
    type: Type
    typeRef: string
    primary: boolean
    required: boolean
    unique: boolean
    immutable: boolean
    length: number
    item: Property
    reference: Reference
    defaultValue: object
    enumValues: string[]
    exampleValue: object
    title: string
    description: string
    annotations: { [key: string]: string }
}

export interface SubType {
    name: string
    title: string
    description: string
    properties: Property[]
}

export interface AuditData {
    createdBy: string
    updatedBy: string
    createdOn: string | Date
    updatedOn: string | Date
}

export interface IndexProperty {
    name: string
    order: Order
}

export interface Index {
    properties: IndexProperty[]
    indexType: IndexType
    unique: boolean
    annotations: { [key: string]: string }
}

export interface Reference {
    resource: Resource
    cascade: boolean
    backReference: string
}

export enum Type {
    BOOL = "BOOL",
    STRING = "STRING",
    FLOAT32 = "FLOAT32",
    FLOAT64 = "FLOAT64",
    INT32 = "INT32",
    INT64 = "INT64",
    BYTES = "BYTES",
    UUID = "UUID",
    DATE = "DATE",
    TIME = "TIME",
    TIMESTAMP = "TIMESTAMP",
    OBJECT = "OBJECT",
    MAP = "MAP",
    LIST = "LIST",
    REFERENCE = "REFERENCE",
    ENUM = "ENUM",
    STRUCT = "STRUCT",
}

export enum Order {
    UNKNOWN = "UNKNOWN",
    ASC = "ASC",
    DESC = "DESC",
}

export enum IndexType {
    BTREE = "BTREE",
    HASH = "HASH",
}


