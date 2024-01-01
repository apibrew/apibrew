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
  "properties": {
    "annotations": {
      "type": "MAP",
      "item": {
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
    "auditData": {
      "type": "STRUCT",
      "typeRef": "AuditData",
      "exampleValue": {
        "createdBy": "admin",
        "createdOn": "2024-01-02T02:32:09+04:00",
        "updatedBy": "admin",
        "updatedOn": "2024-01-02T02:32:09+04:00"
      },
      "title": "Audit Data",
      "description": "The audit data of the resource/record. \nIt contains information about who created the resource/record, when it was created, who last updated the resource/record and when it was last updated.",
      "annotations": {
        "SpecialProperty": "true"
      }
    },
    "call": {
      "type": "STRUCT",
      "typeRef": "ExternalCall",
      "required": true
    },
    "description": {
      "type": "STRING",
      "length": 1024
    },
    "finalizes": {
      "type": "BOOL",
      "required": true
    },
    "id": {
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
    "name": {
      "type": "STRING",
      "required": true,
      "unique": true,
      "length": 256,
      "annotations": {
        "IsHclLabel": "true"
      }
    },
    "order": {
      "type": "INT32",
      "required": true
    },
    "responds": {
      "type": "BOOL",
      "required": true
    },
    "selector": {
      "type": "STRUCT",
      "typeRef": "EventSelector"
    },
    "sync": {
      "type": "BOOL",
      "required": true
    },
    "version": {
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
    }
  },
  "types": [
    {
      "name": "BooleanExpression",
      "title": "",
      "description": "",
      "properties": {
        "and": {
          "type": "LIST",
          "item": {
            "type": "STRUCT",
            "typeRef": "BooleanExpression"
          }
        },
        "equal": {
          "type": "STRUCT",
          "typeRef": "PairExpression"
        },
        "greaterThan": {
          "type": "STRUCT",
          "typeRef": "PairExpression"
        },
        "greaterThanOrEqual": {
          "type": "STRUCT",
          "typeRef": "PairExpression"
        },
        "in": {
          "type": "STRUCT",
          "typeRef": "PairExpression"
        },
        "isNull": {
          "type": "STRUCT",
          "typeRef": "Expression"
        },
        "lessThan": {
          "type": "STRUCT",
          "typeRef": "PairExpression"
        },
        "lessThanOrEqual": {
          "type": "STRUCT",
          "typeRef": "PairExpression"
        },
        "not": {
          "type": "STRUCT",
          "typeRef": "BooleanExpression"
        },
        "or": {
          "type": "LIST",
          "item": {
            "type": "STRUCT",
            "typeRef": "BooleanExpression"
          }
        },
        "regexMatch": {
          "type": "STRUCT",
          "typeRef": "RegexMatchExpression"
        }
      }
    },
    {
      "name": "PairExpression",
      "title": "",
      "description": "",
      "properties": {
        "left": {
          "type": "STRUCT",
          "typeRef": "Expression"
        },
        "right": {
          "type": "STRUCT",
          "typeRef": "Expression"
        }
      }
    },
    {
      "name": "RegexMatchExpression",
      "title": "",
      "description": "",
      "properties": {
        "expression": {
          "type": "STRUCT",
          "typeRef": "Expression"
        },
        "pattern": {
          "type": "STRING"
        }
      }
    },
    {
      "name": "Expression",
      "title": "",
      "description": "",
      "properties": {
        "property": {
          "type": "STRING"
        },
        "value": {
          "type": "OBJECT"
        }
      }
    },
    {
      "name": "AuditData",
      "title": "Audit Data",
      "description": "Audit Data is a type that represents the audit data of a resource/record. ",
      "properties": {
        "createdBy": {
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
        "createdOn": {
          "type": "TIMESTAMP",
          "immutable": true,
          "exampleValue": "2024-01-02T02:32:09+04:00",
          "title": "Created On",
          "description": "The timestamp when the resource/record was created.",
          "annotations": {
            "SpecialProperty": "true"
          }
        },
        "updatedBy": {
          "type": "STRING",
          "length": 256,
          "exampleValue": "admin",
          "title": "Updated By",
          "description": "The user who last updated the resource/record.",
          "annotations": {
            "SpecialProperty": "true"
          }
        },
        "updatedOn": {
          "type": "TIMESTAMP",
          "exampleValue": "2024-01-02T02:32:09+04:00",
          "title": "Updated On",
          "description": "The timestamp when the resource/record was last updated.",
          "annotations": {
            "SpecialProperty": "true"
          }
        }
      }
    },
    {
      "name": "FunctionCall",
      "title": "",
      "description": "",
      "properties": {
        "functionName": {
          "type": "STRING",
          "required": true
        },
        "host": {
          "type": "STRING",
          "required": true
        }
      }
    },
    {
      "name": "HttpCall",
      "title": "",
      "description": "",
      "properties": {
        "method": {
          "type": "STRING",
          "required": true
        },
        "uri": {
          "type": "STRING",
          "required": true
        }
      }
    },
    {
      "name": "ChannelCall",
      "title": "",
      "description": "",
      "properties": {
        "channelKey": {
          "type": "STRING",
          "required": true
        }
      }
    },
    {
      "name": "ExternalCall",
      "title": "",
      "description": "",
      "properties": {
        "channelCall": {
          "type": "STRUCT",
          "typeRef": "ChannelCall"
        },
        "functionCall": {
          "type": "STRUCT",
          "typeRef": "FunctionCall"
        },
        "httpCall": {
          "type": "STRUCT",
          "typeRef": "HttpCall"
        }
      }
    },
    {
      "name": "EventSelector",
      "title": "",
      "description": "",
      "properties": {
        "actions": {
          "type": "LIST",
          "item": {
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
        "annotations": {
          "type": "MAP",
          "item": {
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
        "ids": {
          "type": "LIST",
          "item": {
            "type": "STRING"
          }
        },
        "namespaces": {
          "type": "LIST",
          "item": {
            "type": "STRING"
          }
        },
        "recordSelector": {
          "type": "STRUCT",
          "typeRef": "BooleanExpression"
        },
        "resources": {
          "type": "LIST",
          "item": {
            "type": "STRING"
          }
        }
      }
    },
    {
      "name": "RecordSearchParams",
      "title": "",
      "description": "",
      "properties": {
        "limit": {
          "type": "INT32"
        },
        "offset": {
          "type": "INT32"
        },
        "query": {
          "type": "STRUCT",
          "typeRef": "BooleanExpression"
        },
        "resolveReferences": {
          "type": "LIST",
          "item": {
            "type": "STRING"
          }
        }
      }
    },
    {
      "name": "Event",
      "title": "",
      "description": "",
      "properties": {
        "action": {
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
        "actionDescription": {
          "type": "STRING"
        },
        "actionName": {
          "type": "STRING"
        },
        "actionSummary": {
          "type": "STRING"
        },
        "annotations": {
          "type": "MAP",
          "item": {
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
        "error": {
          "type": "STRUCT",
          "typeRef": "Error"
        },
        "finalizes": {
          "type": "BOOL"
        },
        "id": {
          "type": "STRING",
          "required": true,
          "immutable": true
        },
        "input": {
          "type": "OBJECT"
        },
        "output": {
          "type": "OBJECT"
        },
        "recordSearchParams": {
          "type": "STRUCT",
          "typeRef": "RecordSearchParams"
        },
        "records": {
          "type": "LIST",
          "item": {
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
        "resource": {
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
        "sync": {
          "type": "BOOL"
        },
        "time": {
          "type": "TIMESTAMP"
        },
        "total": {
          "type": "INT64"
        }
      }
    },
    {
      "name": "ErrorField",
      "title": "",
      "description": "",
      "properties": {
        "message": {
          "type": "STRING"
        },
        "property": {
          "type": "STRING"
        },
        "recordId": {
          "type": "STRING"
        },
        "value": {
          "type": "OBJECT"
        }
      }
    },
    {
      "name": "Error",
      "title": "",
      "description": "",
      "properties": {
        "code": {
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
        "fields": {
          "type": "LIST",
          "item": {
            "type": "STRUCT",
            "typeRef": "ErrorField"
          }
        },
        "message": {
          "type": "STRING"
        }
      }
    }
  ],
  "annotations": {
    "EnableAudit": "true",
    "OpenApiGroup": "internal"
  }
} as unknown

