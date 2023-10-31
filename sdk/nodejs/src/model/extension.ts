import {Resource} from './resource';
import {Record} from './record';

export interface Extension {
    id: string
    version: number
    auditData?: AuditData
    name: string
    description?: string
    selector?: EventSelector
    order: number
    finalizes: boolean
    sync: boolean
    responds: boolean
    call: ExternalCall
    annotations?: { [key: string]: string }
}

export const NAMESPACE = "system";
export const RESOURCE = "Extension";
export const REST_PATH = "system-extension"

export const ExtensionEntityInfo = {
    namespace: NAMESPACE,
    resource: RESOURCE,
    restPath: REST_PATH,
}

export interface BooleanExpression {
    and: BooleanExpression[]
    or: BooleanExpression[]
    not: BooleanExpression
    equal: PairExpression
    lessThan: PairExpression
    greaterThan: PairExpression
    lessThanOrEqual: PairExpression
    greaterThanOrEqual: PairExpression
    $in: PairExpression
    isNull: Expression
    regexMatch: RegexMatchExpression
}

export interface PairExpression {
    left: Expression
    right: Expression
}

export interface RegexMatchExpression {
    pattern: string
    expression: Expression
}

export interface Expression {
    property: string
    value: object
}

export interface AuditData {
    createdBy: string
    updatedBy: string
    createdOn: string | Date
    updatedOn: string | Date
}

export interface FunctionCall {
    host: string
    functionName: string
}

export interface HttpCall {
    uri: string
    method: string
}

export interface ChannelCall {
    channelKey: string
}

export interface ExternalCall {
    functionCall: FunctionCall
    httpCall: HttpCall
    channelCall: ChannelCall
}

export interface EventSelector {
    actions: Action[]
    recordSelector: BooleanExpression
    namespaces: string[]
    resources: string[]
    ids: string[]
    annotations: { [key: string]: string }
}

export interface RecordSearchParams {
    query: BooleanExpression
    limit: number
    offset: number
    resolveReferences: string[]
}

export interface Event {
    id: string
    action: Action
    recordSearchParams: RecordSearchParams
    actionSummary: string
    actionDescription: string
    resource: Resource
    records: Record[]
    finalizes: boolean
    sync: boolean
    time: string | Date
    total: number
    actionName: string
    input: object
    output: object
    annotations: { [key: string]: string }
    error: Error
}

export interface ErrorField {
    recordId: string
    property: string
    message: string
    value: object
}

export interface Error {
    code: Code
    message: string
    fields: ErrorField[]
}

export enum Action {
    CREATE = "CREATE",
    UPDATE = "UPDATE",
    DELETE = "DELETE",
    GET = "GET",
    LIST = "LIST",
    OPERATE = "OPERATE",
}

export enum Code {
    UNKNOWN_ERROR = "UNKNOWN_ERROR",
    RECORD_NOT_FOUND = "RECORD_NOT_FOUND",
    UNABLE_TO_LOCATE_PRIMARY_KEY = "UNABLE_TO_LOCATE_PRIMARY_KEY",
    INTERNAL_ERROR = "INTERNAL_ERROR",
    PROPERTY_NOT_FOUND = "PROPERTY_NOT_FOUND",
    RECORD_VALIDATION_ERROR = "RECORD_VALIDATION_ERROR",
    RESOURCE_VALIDATION_ERROR = "RESOURCE_VALIDATION_ERROR",
    AUTHENTICATION_FAILED = "AUTHENTICATION_FAILED",
    ALREADY_EXISTS = "ALREADY_EXISTS",
    ACCESS_DENIED = "ACCESS_DENIED",
    BACKEND_ERROR = "BACKEND_ERROR",
    UNIQUE_VIOLATION = "UNIQUE_VIOLATION",
    REFERENCE_VIOLATION = "REFERENCE_VIOLATION",
    RESOURCE_NOT_FOUND = "RESOURCE_NOT_FOUND",
    UNSUPPORTED_OPERATION = "UNSUPPORTED_OPERATION",
    EXTERNAL_BACKEND_COMMUNICATION_ERROR = "EXTERNAL_BACKEND_COMMUNICATION_ERROR",
    EXTERNAL_BACKEND_ERROR = "EXTERNAL_BACKEND_ERROR",
    RATE_LIMIT_ERROR = "RATE_LIMIT_ERROR",
}


