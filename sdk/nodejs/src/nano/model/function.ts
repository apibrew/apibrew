
export interface Function {
    id: string
    name: string
    source: string
    version: number
    language: Language
    auditData?: AuditData
    annotations?: { [key: string]: string }
    contentFormat: ContentFormat
}

export const FunctionEntityInfo = {
    namespace: "nano",
    resource: "Function",
    restPath: "nano-function",
}

export interface AuditData {
    updatedOn: string | Date
    createdBy: string
    createdOn: string | Date
    updatedBy: string
}

export enum Language {
    JAVASCRIPT = "JAVASCRIPT",
}

export enum ContentFormat {
    TEXT = "TEXT",
    TAR = "TAR",
    TAR_GZ = "TAR_GZ",
}

export const FunctionResource = {
  "auditData": {
    "createdBy": "system",
    "updatedBy": "system",
    "createdOn": "2024-04-11T10:48:24Z",
    "updatedOn": "2024-04-17T10:30:03Z"
  },
  "name": "Function",
  "namespace": {
    "name": "nano"
  },
  "properties": {
    "annotations": {
      "type": "MAP",
      "item": {
        "type": "STRING"
      }
    },
    "auditData": {
      "type": "STRUCT",
      "typeRef": "AuditData",
      "exampleValue": {
        "createdBy": "admin",
        "createdOn": "2024-04-14T10:47:09+04:00",
        "updatedBy": "admin",
        "updatedOn": "2024-04-14T10:47:09+04:00"
      },
      "annotations": {
        "SpecialProperty": "true"
      }
    },
    "contentFormat": {
      "type": "ENUM",
      "required": true,
      "defaultValue": "TEXT",
      "enumValues": [
        "TEXT",
        "TAR",
        "TAR_GZ"
      ]
    },
    "id": {
      "type": "UUID",
      "primary": true,
      "required": true,
      "immutable": true,
      "exampleValue": "a39621a4-6d48-11ee-b962-0242ac120002",
      "annotations": {
        "SpecialProperty": "true"
      }
    },
    "language": {
      "type": "ENUM",
      "required": true,
      "defaultValue": "JAVASCRIPT",
      "enumValues": [
        "JAVASCRIPT"
      ]
    },
    "name": {
      "type": "STRING",
      "required": true,
      "unique": true,
      "immutable": true,
      "length": 255
    },
    "source": {
      "type": "STRING",
      "required": true,
      "length": 64000,
      "annotations": {
        "SQLType": "TEXT"
      }
    },
    "version": {
      "type": "INT32",
      "required": true,
      "defaultValue": 1,
      "exampleValue": 1,
      "annotations": {
        "AllowEmptyPrimitive": "true",
        "SpecialProperty": "true"
      }
    }
  },
  "types": [
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
          "annotations": {
            "SpecialProperty": "true"
          }
        },
        "createdOn": {
          "type": "TIMESTAMP",
          "immutable": true,
          "exampleValue": "2024-04-14T10:47:09+04:00",
          "annotations": {
            "SpecialProperty": "true"
          }
        },
        "updatedBy": {
          "type": "STRING",
          "length": 256,
          "exampleValue": "admin",
          "annotations": {
            "SpecialProperty": "true"
          }
        },
        "updatedOn": {
          "type": "TIMESTAMP",
          "exampleValue": "2024-04-14T10:47:09+04:00",
          "annotations": {
            "SpecialProperty": "true"
          }
        }
      }
    }
  ],
  "title": "Function",
  "description": "Nano function",
  "annotations": {
    "EnableAudit": "true",
    "OpenApiGroup": "meta"
  }
} as unknown

