import {Resource} from './resource';

export interface ResourceAction {
    id: string
    version: number
    auditData?: AuditData
    resource: Resource
    name: string
    title?: string
    description?: string
    internal: boolean
    types?: SubType[]
    input?: Property[]
    output?: Property
    annotations?: { [key: string]: string }
}

export const ResourceActionEntityInfo = {
    namespace: "system",
    resource: "ResourceAction",
    restPath: "system-resourceaction",
}

export interface SubType {
    name: string
    title: string
    description: string
    properties: Property[]
}

export interface Reference {
    resource: Resource
    cascade: boolean
    backReference: string
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

export interface AuditData {
    createdBy: string
    updatedBy: string
    createdOn: string | Date
    updatedOn: string | Date
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


