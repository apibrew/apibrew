
export interface StorageObject {
    id: string
    name?: string
    annotations?: { [key: string]: string }
    contentType?: string
    size?: number
    allowDownloadPublicly: boolean
    allowUploadPublicly: boolean
    version: number
    auditData?: AuditData
}

export const StorageObjectEntityInfo = {
    namespace: "storage",
    resource: "StorageObject",
    restPath: "storage-storageobject",
}

export interface AuditData {
    createdBy: string
    updatedBy: string
    createdOn: string | Date
    updatedOn: string | Date
}

export const StorageObjectResource = {
  "auditData": {
    "createdBy": "2023-10-29T10:47:46Z",
    "updatedBy": "admin",
    "createdOn": "2023-10-28T20:42:57Z"
  },
  "name": "StorageObject",
  "namespace": {
    "name": "storage"
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
      "length": 255,
      "title": "File Name",
      "description": "File name"
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
      "name": "contentType",
      "type": "STRING",
      "length": 255,
      "title": "Content Type",
      "description": "Content type"
    },
    {
      "name": "size",
      "type": "INT64",
      "title": "Size",
      "description": "File size"
    },
    {
      "name": "allowDownloadPublicly",
      "type": "BOOL",
      "required": true,
      "defaultValue": false,
      "title": "Allow Download Publicly",
      "description": "Allow download publicly"
    },
    {
      "name": "allowUploadPublicly",
      "type": "BOOL",
      "required": true,
      "defaultValue": false,
      "title": "Allow Upload Publicly",
      "description": "Allow upload publicly"
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
        "createdOn": "2023-10-29T13:53:20+04:00",
        "updatedBy": "admin",
        "updatedOn": "2023-10-29T13:53:20+04:00"
      },
      "title": "Audit Data",
      "description": "The audit data of the resource/record. \nIt contains information about who created the resource/record, when it was created, who last updated the resource/record and when it was last updated.",
      "annotations": {
        "SpecialProperty": "true"
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
          "exampleValue": "2023-10-29T13:53:20+04:00",
          "title": "Created On",
          "description": "The timestamp when the resource/record was created.",
          "annotations": {
            "SpecialProperty": "true"
          }
        },
        {
          "name": "updatedOn",
          "type": "TIMESTAMP",
          "exampleValue": "2023-10-29T13:53:20+04:00",
          "title": "Updated On",
          "description": "The timestamp when the resource/record was last updated.",
          "annotations": {
            "SpecialProperty": "true"
          }
        }
      ]
    }
  ],
  "title": "StorageObject",
  "description": "Storage Object",
  "annotations": {
    "EnableAudit": "true",
    "NormalizedResource": "true"
  }
} as unknown

