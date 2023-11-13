
export interface Code {
    id: string
    name: string
    language: Language
    content: string
    contentFormat: ContentFormat
    annotations?: { [key: string]: string }
    version: number
    auditData?: AuditData
}

export const CodeEntityInfo = {
    namespace: "nano",
    resource: "Code",
    restPath: "nano-code",
}

export interface AuditData {
    createdBy: string
    updatedBy: string
    createdOn: string | Date
    updatedOn: string | Date
}

export enum Language {
    PYTHON = "PYTHON",
    JAVASCRIPT = "JAVASCRIPT",
}

export enum ContentFormat {
    TEXT = "TEXT",
    TAR = "TAR",
    TAR_GZ = "TAR_GZ",
}

export const CodeResource = {
  "auditData": {
    "createdBy": "2023-10-20T11:10:03Z",
    "updatedBy": "admin",
    "createdOn": "2023-10-20T10:53:30Z"
  },
  "name": "Code",
  "namespace": {
    "name": "nano"
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
      "name": "name",
      "type": "STRING",
      "required": true,
      "unique": true,
      "length": 255,
      "title": "Name",
      "description": "Full Qualified Name of the code, it must be unique in the system"
    },
    {
      "name": "language",
      "type": "ENUM",
      "required": true,
      "enumValues": [
        "PYTHON",
        "JAVASCRIPT"
      ],
      "title": "Language",
      "description": "Code language"
    },
    {
      "name": "content",
      "type": "STRING",
      "required": true,
      "length": 64000,
      "title": "Content",
      "description": "Code content",
      "annotations": {
        "SQLType": "TEXT"
      }
    },
    {
      "name": "contentFormat",
      "type": "ENUM",
      "required": true,
      "defaultValue": "TEXT",
      "enumValues": [
        "TEXT",
        "TAR",
        "TAR_GZ"
      ],
      "title": "Content Format",
      "description": "Code content format"
    },
    {
      "name": "annotations",
      "type": "MAP",
      "item": {
        "name": "",
        "type": "STRING"
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
        "createdOn": "2023-10-20T15:07:20+04:00",
        "updatedBy": "admin",
        "updatedOn": "2023-10-20T15:07:20+04:00"
      },
      "title": "Audit Data",
      "description": "The audit data of the resource/record. \nIt contains information about who created the resource/record, when it was created, who last updated the resource/record and when it was last updated."
    }
  ],
  "types": [
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
          "exampleValue": "2023-10-20T15:07:20+04:00",
          "title": "Created On",
          "description": "The timestamp when the resource/record was created.",
          "annotations": {
            "SpecialProperty": "true"
          }
        },
        {
          "name": "updatedOn",
          "type": "TIMESTAMP",
          "exampleValue": "2023-10-20T15:07:20+04:00",
          "title": "Updated On",
          "description": "The timestamp when the resource/record was last updated.",
          "annotations": {
            "SpecialProperty": "true"
          }
        }
      ]
    }
  ],
  "title": "Code",
  "description": "Nano code",
  "annotations": {
    "EnableAudit": "true",
    "NormalizedResource": "true"
  }
} as unknown

