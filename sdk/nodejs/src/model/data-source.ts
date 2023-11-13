
export interface DataSource {
    id: string
    version: number
    auditData?: AuditData
    name: string
    description?: string
    backend: Backend
    options: { [key: string]: string }
}

export const DataSourceEntityInfo = {
    namespace: "system",
    resource: "DataSource",
    restPath: "system-datasource",
}

export interface AuditData {
    createdBy: string
    updatedBy: string
    createdOn: string | Date
    updatedOn: string | Date
}

export enum Backend {
    POSTGRESQL = "POSTGRESQL",
    MYSQL = "MYSQL",
    MONGODB = "MONGODB",
    REDIS = "REDIS",
}

export const DataSourceResource = {
  "name": "DataSource",
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
        "createdOn": "2023-11-13T12:31:41+04:00",
        "updatedBy": "admin",
        "updatedOn": "2023-11-13T12:31:41+04:00"
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
      "length": 64,
      "annotations": {
        "IsHclLabel": "true"
      }
    },
    {
      "name": "description",
      "type": "STRING",
      "length": 64,
      "annotations": {
        "AllowEmptyPrimitive": "true"
      }
    },
    {
      "name": "backend",
      "type": "ENUM",
      "required": true,
      "enumValues": [
        "POSTGRESQL",
        "MYSQL",
        "MONGODB",
        "REDIS"
      ]
    },
    {
      "name": "options",
      "type": "MAP",
      "required": true,
      "item": {
        "name": "",
        "type": "STRING"
      }
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
          "exampleValue": "2023-11-13T12:31:41+04:00",
          "title": "Created On",
          "description": "The timestamp when the resource/record was created.",
          "annotations": {
            "SpecialProperty": "true"
          }
        },
        {
          "name": "updatedOn",
          "type": "TIMESTAMP",
          "exampleValue": "2023-11-13T12:31:41+04:00",
          "title": "Updated On",
          "description": "The timestamp when the resource/record was last updated.",
          "annotations": {
            "SpecialProperty": "true"
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

