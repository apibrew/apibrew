
export interface Job {
    contentFormat: ContentFormat
    lastExecutionTime?: string | Date
    nextExecutionTime: string | Date
    lastExecutionError?: string
    name: string
    version: number
    id: string
    source: string
    annotations?: { [key: string]: string }
    language: Language
    auditData?: AuditData
}

export const JobEntityInfo = {
    namespace: "nano",
    resource: "Job",
    restPath: "nano-job",
}

export interface AuditData {
    createdBy: string
    createdOn: string | Date
    updatedBy: string
    updatedOn: string | Date
}

export enum ContentFormat {
    TEXT = "TEXT",
    TAR = "TAR",
    TAR_GZ = "TAR_GZ",
}

export enum Language {
    JAVASCRIPT = "JAVASCRIPT",
    TYPESCRIPT = "TYPESCRIPT",
}

export const JobResource = {
  "auditData": {
    "createdBy": "system",
    "createdOn": "2024-05-15T09:23:49Z"
  },
  "name": "Job",
  "namespace": {
    "name": "nano"
  },
  "properties": {
    "annotations": {
      "type": "MAP",
      "item": {
        "type": "STRING"
      },
      "annotations": {
        "SourceMatchKey": "f9b03cb220e3"
      }
    },
    "auditData": {
      "type": "STRUCT",
      "typeRef": "AuditData",
      "exampleValue": {
        "createdBy": "admin",
        "createdOn": "2024-05-15T13:05:39+04:00",
        "updatedBy": "admin",
        "updatedOn": "2024-05-15T13:05:39+04:00"
      },
      "annotations": {
        "SourceMatchKey": "66658db834e5",
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
      ],
      "annotations": {
        "SourceMatchKey": "d2ff25026060"
      }
    },
    "id": {
      "type": "UUID",
      "primary": true,
      "required": true,
      "immutable": true,
      "exampleValue": "a39621a4-6d48-11ee-b962-0242ac120002",
      "annotations": {
        "SourceMatchKey": "493db3ecf1d9",
        "SpecialProperty": "true"
      }
    },
    "language": {
      "type": "ENUM",
      "required": true,
      "defaultValue": "JAVASCRIPT",
      "enumValues": [
        "JAVASCRIPT",
        "TYPESCRIPT"
      ],
      "annotations": {
        "SourceMatchKey": "c3420ae1636e"
      }
    },
    "lastExecutionError": {
      "type": "STRING",
      "annotations": {
        "SourceMatchKey": "bb3b131dceb6"
      }
    },
    "lastExecutionTime": {
      "type": "TIMESTAMP",
      "annotations": {
        "SourceMatchKey": "71eec289d6ff"
      }
    },
    "name": {
      "type": "STRING",
      "required": true,
      "unique": true,
      "immutable": true,
      "length": 255,
      "annotations": {
        "SourceMatchKey": "7ef6c1e20810"
      }
    },
    "nextExecutionTime": {
      "type": "TIMESTAMP",
      "required": true,
      "annotations": {
        "SourceMatchKey": "c06be26ad8c4"
      }
    },
    "source": {
      "type": "STRING",
      "required": true,
      "length": 64000,
      "annotations": {
        "SQLType": "TEXT",
        "SourceMatchKey": "56072dc2acd4"
      }
    },
    "version": {
      "type": "INT32",
      "required": true,
      "defaultValue": 1,
      "exampleValue": 1,
      "annotations": {
        "AllowEmptyPrimitive": "true",
        "SourceMatchKey": "5c11e478313a",
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
          "exampleValue": "2024-05-15T13:05:39+04:00",
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
          "exampleValue": "2024-05-15T13:05:39+04:00",
          "annotations": {
            "SpecialProperty": "true"
          }
        }
      }
    }
  ],
  "title": "Job",
  "description": "Job",
  "annotations": {
    "EnableAudit": "true",
    "OpenApiGroup": "meta"
  }
} as unknown

