
export interface Code {
    language: Language
    auditData?: AuditData
    annotations?: { [key: string]: string }
    contentFormat: ContentFormat
    id: string
    name: string
    content: string
    version: number
}

export const CodeEntityInfo = {
    namespace: "nano",
    resource: "Code",
    restPath: "nano-code",
}

export interface AuditData {
    createdOn: string | Date
    updatedBy: string
    updatedOn: string | Date
    createdBy: string
}

export enum Language {
    JAVASCRIPT = "JAVASCRIPT",
}

export enum ContentFormat {
    TEXT = "TEXT",
    TAR = "TAR",
    TAR_GZ = "TAR_GZ",
}

export const CodeResource = {
  "auditData": {
    "createdBy": "system",
    "updatedBy": "system",
    "createdOn": "2023-12-20T17:23:46Z",
    "updatedOn": "2024-01-01T22:32:10Z"
  },
  "name": "Code",
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
        "createdOn": "2023-12-13T21:18:29Z",
        "updatedBy": "admin",
        "updatedOn": "2023-12-13T21:18:29Z"
      },
      "annotations": {
        "SpecialProperty": "true"
      }
    },
    "content": {
      "type": "STRING",
      "required": true,
      "length": 64000,
      "annotations": {
        "SQLType": "TEXT"
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
      "required": true,
      "immutable": true,
      "exampleValue": "a39621a4-6d48-11ee-b962-0242ac120002",
      "annotations": {
        "PrimaryProperty": "true",
        "SpecialProperty": "true"
      }
    },
    "language": {
      "type": "ENUM",
      "required": true,
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
          "exampleValue": "2023-12-13T21:18:29Z",
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
          "exampleValue": "2023-12-13T21:18:29Z",
          "annotations": {
            "SpecialProperty": "true"
          }
        }
      }
    }
  ],
  "title": "Code",
  "description": "Nano code",
  "annotations": {
    "EnableAudit": "true",
    "NormalizedResource": "true"
  }
} as unknown

