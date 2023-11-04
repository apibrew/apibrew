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

export const ExtensionEntityInfo = {
    namespace: "system",
    resource: "Extension",
    restPath: "system-extension",
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
    in: PairExpression
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

export const ExtensionResource = {
  "name": "Extension",
  "namespace": {
    "name": "system"
  },
  "properties": [
    {
      "name": "id",
      "type": "UUID",
      "required": true,
      "immutable": true,
      "exampleValue": "a39621a4-6d48-11ee-b962-0242ac120002",
      "description": "The unique identifier of the resource. It is randomly generated and immutable.",
      "annotations": {
        "PrimaryProperty": "true",
        "SpecialProperty": "true"
      }
    },
    {
      "name": "version",
      "type": "INT32",
      "required": true,
      "defaultValue": 1,
      "exampleValue": 1,
      "title": "Version",
      "description": "The version of the resource/record. It is incremented on every update.",
      "annotations": {
        "AllowEmptyPrimitive": "true",
        "SpecialProperty": "true"
      }
    },
    {
      "name": "auditData",
      "type": "STRUCT",
      "typeRef": "AuditData",
      "exampleValue": {
        "createdBy": "admin",
        "createdOn": "2023-11-04T03:41:34+04:00",
        "updatedBy": "admin",
        "updatedOn": "2023-11-04T03:41:34+04:00"
      },
      "title": "Audit Data",
      "description": "The audit data of the resource/record. \nIt contains information about who created the resource/record, when it was created, who last updated the resource/record and when it was last updated.",
      "annotations": {
        "SpecialProperty": "true"
      }
    },
    {
      "name": "name",
      "type": "STRING",
      "required": true,
      "unique": true,
      "length": 256,
      "annotations": {
        "IsHclLabel": "true"
      }
    },
    {
      "name": "description",
      "type": "STRING",
      "length": 1024
    },
    {
      "name": "selector",
      "type": "STRUCT",
      "typeRef": "EventSelector"
    },
    {
      "name": "order",
      "type": "INT32",
      "required": true
    },
    {
      "name": "finalizes",
      "type": "BOOL",
      "required": true
    },
    {
      "name": "sync",
      "type": "BOOL",
      "required": true
    },
    {
      "name": "responds",
      "type": "BOOL",
      "required": true
    },
    {
      "name": "call",
      "type": "STRUCT",
      "typeRef": "ExternalCall",
      "required": true
    },
    {
      "name": "annotations",
      "type": "MAP",
      "item": {
        "name": "",
        "type": "STRING"
      },
      "exampleValue": {
        "CheckVersion": "true",
        "CommonType": "testType",
        "IgnoreIfExists": "true"
      },
      "title": "Annotations",
      "description": "The annotations of the resource/record. It contains information about the resource/record. For example, it can contain information about the UI representation of the resource/record.",
      "annotations": {
        "SpecialProperty": "true"
      }
    }
  ],
  "types": [
    {
      "name": "BooleanExpression",
      "title": "",
      "description": "",
      "properties": [
        {
          "name": "and",
          "type": "LIST",
          "item": {
            "name": "",
            "type": "STRUCT",
            "typeRef": "BooleanExpression"
          }
        },
        {
          "name": "or",
          "type": "LIST",
          "item": {
            "name": "",
            "type": "STRUCT",
            "typeRef": "BooleanExpression"
          }
        },
        {
          "name": "not",
          "type": "STRUCT",
          "typeRef": "BooleanExpression"
        },
        {
          "name": "equal",
          "type": "STRUCT",
          "typeRef": "PairExpression"
        },
        {
          "name": "lessThan",
          "type": "STRUCT",
          "typeRef": "PairExpression"
        },
        {
          "name": "greaterThan",
          "type": "STRUCT",
          "typeRef": "PairExpression"
        },
        {
          "name": "lessThanOrEqual",
          "type": "STRUCT",
          "typeRef": "PairExpression"
        },
        {
          "name": "greaterThanOrEqual",
          "type": "STRUCT",
          "typeRef": "PairExpression"
        },
        {
          "name": "in",
          "type": "STRUCT",
          "typeRef": "PairExpression"
        },
        {
          "name": "isNull",
          "type": "STRUCT",
          "typeRef": "Expression"
        },
        {
          "name": "regexMatch",
          "type": "STRUCT",
          "typeRef": "RegexMatchExpression"
        }
      ]
    },
    {
      "name": "PairExpression",
      "title": "",
      "description": "",
      "properties": [
        {
          "name": "left",
          "type": "STRUCT",
          "typeRef": "Expression"
        },
        {
          "name": "right",
          "type": "STRUCT",
          "typeRef": "Expression"
        }
      ]
    },
    {
      "name": "RegexMatchExpression",
      "title": "",
      "description": "",
      "properties": [
        {
          "name": "pattern",
          "type": "STRING"
        },
        {
          "name": "expression",
          "type": "STRUCT",
          "typeRef": "Expression"
        }
      ]
    },
    {
      "name": "Expression",
      "title": "",
      "description": "",
      "properties": [
        {
          "name": "property",
          "type": "STRING"
        },
        {
          "name": "value",
          "type": "OBJECT"
        }
      ]
    },
    {
      "name": "AuditData",
      "title": "Audit Data",
      "description": "Audit Data is a type that represents the audit data of a resource/record. ",
      "properties": [
        {
          "name": "createdBy",
          "type": "STRING",
          "immutable": true,
          "length": 256,
          "exampleValue": "admin",
          "title": "Created By",
          "description": "The user who created the resource/record.",
          "annotations": {
            "SpecialProperty": "true"
          }
        },
        {
          "name": "updatedBy",
          "type": "STRING",
          "length": 256,
          "exampleValue": "admin",
          "title": "Updated By",
          "description": "The user who last updated the resource/record.",
          "annotations": {
            "SpecialProperty": "true"
          }
        },
        {
          "name": "createdOn",
          "type": "TIMESTAMP",
          "immutable": true,
          "exampleValue": "2023-11-04T03:41:34+04:00",
          "title": "Created On",
          "description": "The timestamp when the resource/record was created.",
          "annotations": {
            "SpecialProperty": "true"
          }
        },
        {
          "name": "updatedOn",
          "type": "TIMESTAMP",
          "exampleValue": "2023-11-04T03:41:34+04:00",
          "title": "Updated On",
          "description": "The timestamp when the resource/record was last updated.",
          "annotations": {
            "SpecialProperty": "true"
          }
        }
      ]
    },
    {
      "name": "FunctionCall",
      "title": "",
      "description": "",
      "properties": [
        {
          "name": "host",
          "type": "STRING",
          "required": true
        },
        {
          "name": "functionName",
          "type": "STRING",
          "required": true
        }
      ]
    },
    {
      "name": "HttpCall",
      "title": "",
      "description": "",
      "properties": [
        {
          "name": "uri",
          "type": "STRING",
          "required": true
        },
        {
          "name": "method",
          "type": "STRING",
          "required": true
        }
      ]
    },
    {
      "name": "ChannelCall",
      "title": "",
      "description": "",
      "properties": [
        {
          "name": "channelKey",
          "type": "STRING",
          "required": true
        }
      ]
    },
    {
      "name": "ExternalCall",
      "title": "",
      "description": "",
      "properties": [
        {
          "name": "functionCall",
          "type": "STRUCT",
          "typeRef": "FunctionCall"
        },
        {
          "name": "httpCall",
          "type": "STRUCT",
          "typeRef": "HttpCall"
        },
        {
          "name": "channelCall",
          "type": "STRUCT",
          "typeRef": "ChannelCall"
        }
      ]
    },
    {
      "name": "EventSelector",
      "title": "",
      "description": "",
      "properties": [
        {
          "name": "actions",
          "type": "LIST",
          "item": {
            "name": "action",
            "type": "ENUM",
            "enumValues": [
              "CREATE",
              "UPDATE",
              "DELETE",
              "GET",
              "LIST",
              "OPERATE"
            ],
            "annotations": {
              "TypeName": "EventAction"
            }
          }
        },
        {
          "name": "recordSelector",
          "type": "STRUCT",
          "typeRef": "BooleanExpression"
        },
        {
          "name": "namespaces",
          "type": "LIST",
          "item": {
            "name": "",
            "type": "STRING"
          }
        },
        {
          "name": "resources",
          "type": "LIST",
          "item": {
            "name": "",
            "type": "STRING"
          }
        },
        {
          "name": "ids",
          "type": "LIST",
          "item": {
            "name": "",
            "type": "STRING"
          }
        },
        {
          "name": "annotations",
          "type": "MAP",
          "item": {
            "name": "",
            "type": "STRING"
          },
          "exampleValue": {
            "CheckVersion": "true",
            "CommonType": "testType",
            "IgnoreIfExists": "true"
          },
          "title": "Annotations",
          "description": "The annotations of the resource/record. It contains information about the resource/record. For example, it can contain information about the UI representation of the resource/record.",
          "annotations": {
            "SpecialProperty": "true"
          }
        }
      ]
    },
    {
      "name": "RecordSearchParams",
      "title": "",
      "description": "",
      "properties": [
        {
          "name": "query",
          "type": "STRUCT",
          "typeRef": "BooleanExpression"
        },
        {
          "name": "limit",
          "type": "INT32"
        },
        {
          "name": "offset",
          "type": "INT32"
        },
        {
          "name": "resolveReferences",
          "type": "LIST",
          "item": {
            "name": "",
            "type": "STRING"
          }
        }
      ]
    },
    {
      "name": "Event",
      "title": "",
      "description": "",
      "properties": [
        {
          "name": "id",
          "type": "STRING",
          "required": true,
          "immutable": true
        },
        {
          "name": "action",
          "type": "ENUM",
          "required": true,
          "enumValues": [
            "CREATE",
            "UPDATE",
            "DELETE",
            "GET",
            "LIST",
            "OPERATE"
          ]
        },
        {
          "name": "recordSearchParams",
          "type": "STRUCT",
          "typeRef": "RecordSearchParams"
        },
        {
          "name": "actionSummary",
          "type": "STRING"
        },
        {
          "name": "actionDescription",
          "type": "STRING"
        },
        {
          "name": "resource",
          "type": "REFERENCE",
          "reference": {
            "resource": {
              "name": "Resource",
              "namespace": {
                "name": "system"
              }
            },
            "cascade": false
          }
        },
        {
          "name": "records",
          "type": "LIST",
          "item": {
            "name": "",
            "type": "REFERENCE",
            "reference": {
              "resource": {
                "name": "Record",
                "namespace": {
                  "name": "system"
                }
              },
              "cascade": false
            }
          }
        },
        {
          "name": "finalizes",
          "type": "BOOL"
        },
        {
          "name": "sync",
          "type": "BOOL"
        },
        {
          "name": "time",
          "type": "TIMESTAMP"
        },
        {
          "name": "total",
          "type": "INT64"
        },
        {
          "name": "actionName",
          "type": "STRING"
        },
        {
          "name": "input",
          "type": "OBJECT"
        },
        {
          "name": "output",
          "type": "OBJECT"
        },
        {
          "name": "annotations",
          "type": "MAP",
          "item": {
            "name": "",
            "type": "STRING"
          },
          "exampleValue": {
            "CheckVersion": "true",
            "CommonType": "testType",
            "IgnoreIfExists": "true"
          },
          "title": "Annotations",
          "description": "The annotations of the resource/record. It contains information about the resource/record. For example, it can contain information about the UI representation of the resource/record.",
          "annotations": {
            "SpecialProperty": "true"
          }
        },
        {
          "name": "error",
          "type": "STRUCT",
          "typeRef": "Error"
        }
      ]
    },
    {
      "name": "ErrorField",
      "title": "",
      "description": "",
      "properties": [
        {
          "name": "recordId",
          "type": "STRING"
        },
        {
          "name": "property",
          "type": "STRING"
        },
        {
          "name": "message",
          "type": "STRING"
        },
        {
          "name": "value",
          "type": "OBJECT"
        }
      ]
    },
    {
      "name": "Error",
      "title": "",
      "description": "",
      "properties": [
        {
          "name": "code",
          "type": "ENUM",
          "enumValues": [
            "UNKNOWN_ERROR",
            "RECORD_NOT_FOUND",
            "UNABLE_TO_LOCATE_PRIMARY_KEY",
            "INTERNAL_ERROR",
            "PROPERTY_NOT_FOUND",
            "RECORD_VALIDATION_ERROR",
            "RESOURCE_VALIDATION_ERROR",
            "AUTHENTICATION_FAILED",
            "ALREADY_EXISTS",
            "ACCESS_DENIED",
            "BACKEND_ERROR",
            "UNIQUE_VIOLATION",
            "REFERENCE_VIOLATION",
            "RESOURCE_NOT_FOUND",
            "UNSUPPORTED_OPERATION",
            "EXTERNAL_BACKEND_COMMUNICATION_ERROR",
            "EXTERNAL_BACKEND_ERROR",
            "RATE_LIMIT_ERROR"
          ]
        },
        {
          "name": "message",
          "type": "STRING"
        },
        {
          "name": "fields",
          "type": "LIST",
          "item": {
            "name": "",
            "type": "STRUCT",
            "typeRef": "ErrorField"
          }
        }
      ]
    }
  ],
  "annotations": {
    "EnableAudit": "true",
    "OpenApiGroup": "internal"
  }
} as unknown

