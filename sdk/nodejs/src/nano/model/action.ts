import {Resource} from "../../model";

export interface Action {
    id: string
    language: Language
    restPath?: string
    contentFormat: ContentFormat
    source: string
    annotations?: { [key: string]: string }
    inputSchema?: { [key: string]: Property }
    outputSchema?: { [key: string]: Property }
    name: string
    version: number
    auditData?: AuditData
    resource?: Resource
}

export const ActionEntityInfo = {
    namespace: "nano",
    resource: "Action",
    restPath: "nano-action",
}

export interface Property {
    typeRef: string
    unique: boolean
    length: number
    description: string
    reference: string
    virtual: boolean
    required: boolean
    immutable: boolean
    backReference: string
    type: Type
    title: string
    exampleValue: object
    item: Property
    enumValues: string[]
    annotations: { [key: string]: string }
    defaultValue: object
    primary: boolean
}

export interface AuditData {
    createdBy: string
    createdOn: string | Date
    updatedBy: string
    updatedOn: string | Date
}

export enum Language {
    JAVASCRIPT = "JAVASCRIPT",
    TYPESCRIPT = "TYPESCRIPT",
}

export enum ContentFormat {
    TEXT = "TEXT",
    TAR = "TAR",
    TAR_GZ = "TAR_GZ",
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

export const ActionResource = {
  "auditData": {
    "createdBy": "system",
    "updatedBy": "system",
    "createdOn": "2024-05-13T13:19:01Z",
    "updatedOn": "2024-05-14T09:23:41Z"
  },
  "name": "Action",
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
        "SourceMatchKey": "3c90d26fc8f4"
      }
    },
    "auditData": {
      "type": "STRUCT",
      "typeRef": "AuditData",
      "exampleValue": {
        "createdBy": "admin",
        "createdOn": "2024-05-13T14:49:06+04:00",
        "updatedBy": "admin",
        "updatedOn": "2024-05-13T14:49:06+04:00"
      },
      "annotations": {
        "SourceMatchKey": "8d062573b9ba",
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
        "SourceMatchKey": "fcca1ae8d8f8"
      }
    },
    "id": {
      "type": "UUID",
      "primary": true,
      "required": true,
      "immutable": true,
      "exampleValue": "a39621a4-6d48-11ee-b962-0242ac120002",
      "annotations": {
        "SourceMatchKey": "93ae987a4d33",
        "SpecialProperty": "true"
      }
    },
    "inputSchema": {
      "type": "MAP",
      "item": {
        "type": "STRUCT",
        "typeRef": "Property"
      },
      "annotations": {
        "SourceMatchKey": "17aa46b0acc9"
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
        "SourceMatchKey": "588d185dc18a"
      }
    },
    "name": {
      "type": "STRING",
      "required": true,
      "unique": true,
      "immutable": true,
      "length": 255,
      "annotations": {
        "SourceMatchKey": "90e649ffd1a9"
      }
    },
    "outputSchema": {
      "type": "MAP",
      "item": {
        "type": "STRUCT",
        "typeRef": "Property"
      },
      "annotations": {
        "SourceMatchKey": "b40ef2307e39"
      }
    },
    "resource": {
      "type": "REFERENCE",
      "reference": "system/Resource",
      "annotations": {
        "SourceMatchKey": "44bc23ddc72d"
      }
    },
    "restPath": {
      "type": "STRING",
      "length": 255,
      "annotations": {
        "SourceMatchKey": "318f4023e862"
      }
    },
    "source": {
      "type": "STRING",
      "required": true,
      "length": 64000,
      "annotations": {
        "SQLType": "TEXT",
        "SourceMatchKey": "38b3fea4b479"
      }
    },
    "version": {
      "type": "INT32",
      "required": true,
      "defaultValue": 1,
      "exampleValue": 1,
      "annotations": {
        "AllowEmptyPrimitive": "true",
        "SourceMatchKey": "7e645d401d21",
        "SpecialProperty": "true"
      }
    }
  },
  "types": [
    {
      "name": "Property",
      "title": "Property",
      "description": "Property is a type that represents a property of a resource. It is like an API properties or properties of class in a programming language",
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
          "annotations": {
            "SpecialProperty": "true"
          }
        },
        "backReference": {
          "type": "STRING",
          "exampleValue": "Book"
        },
        "defaultValue": {
          "type": "OBJECT",
          "exampleValue": "Lord of the Rings"
        },
        "description": {
          "type": "STRING",
          "length": 256,
          "exampleValue": "Book Title is a property of Book Resource. It represents the title of the book."
        },
        "enumValues": {
          "type": "LIST",
          "item": {
            "type": "STRING"
          },
          "exampleValue": [
            "UNKNOWN",
            "ASC",
            "DESC"
          ]
        },
        "exampleValue": {
          "type": "OBJECT",
          "exampleValue": "no-book-name"
        },
        "immutable": {
          "type": "BOOL",
          "required": true,
          "defaultValue": false
        },
        "item": {
          "type": "STRUCT",
          "typeRef": "Property",
          "exampleValue": {
            "type": "STRING"
          }
        },
        "length": {
          "type": "INT32",
          "required": true,
          "defaultValue": 256,
          "exampleValue": 256
        },
        "primary": {
          "type": "BOOL",
          "required": true,
          "defaultValue": false
        },
        "reference": {
          "type": "STRING",
          "exampleValue": "Book"
        },
        "required": {
          "type": "BOOL",
          "required": true,
          "defaultValue": false
        },
        "title": {
          "type": "STRING",
          "length": 256,
          "exampleValue": "Book Title"
        },
        "type": {
          "type": "ENUM",
          "required": true,
          "enumValues": [
            "BOOL",
            "STRING",
            "FLOAT32",
            "FLOAT64",
            "INT32",
            "INT64",
            "BYTES",
            "UUID",
            "DATE",
            "TIME",
            "TIMESTAMP",
            "OBJECT",
            "MAP",
            "LIST",
            "REFERENCE",
            "ENUM",
            "STRUCT"
          ],
          "exampleValue": "STRING"
        },
        "typeRef": {
          "type": "STRING",
          "length": 256,
          "exampleValue": "BookPublishingDetails"
        },
        "unique": {
          "type": "BOOL",
          "required": true,
          "defaultValue": false
        },
        "virtual": {
          "type": "BOOL",
          "required": true,
          "defaultValue": false
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
          "annotations": {
            "SpecialProperty": "true"
          }
        },
        "createdOn": {
          "type": "TIMESTAMP",
          "immutable": true,
          "exampleValue": "2024-05-13T14:49:06+04:00",
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
          "exampleValue": "2024-05-13T14:49:06+04:00",
          "annotations": {
            "SpecialProperty": "true"
          }
        }
      }
    }
  ],
  "title": "Action",
  "description": "Action",
  "annotations": {
    "EnableAudit": "true",
    "OpenApiGroup": "meta"
  }
} as unknown

