import {StorageObject} from './storage-object';

export interface Signature {
    id: string
    object: StorageObject
    permissions: Permission[]
    expiration: string | Date
    signature: string
    annotations?: { [key: string]: string }
    version: number
}

export const SignatureEntityInfo = {
    namespace: "storage",
    resource: "Signature",
    restPath: "storage-signature",
}

export enum Permission {
    DOWNLOAD = "DOWNLOAD",
    UPLOAD = "UPLOAD",
}

export const SignatureResource = {
  "auditData": {
    "createdBy": "2023-10-29T10:47:46Z",
    "updatedBy": "admin",
    "createdOn": "2023-10-28T20:42:57Z"
  },
  "name": "Signature",
  "namespace": {
    "name": "storage"
  },
  "virtual": true,
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
      "name": "object",
      "type": "REFERENCE",
      "required": true,
      "reference": {
        "resource": {
          "name": "StorageObject",
          "namespace": {
            "name": "storage"
          }
        },
        "cascade": false
      },
      "title": "Object",
      "description": "Object"
    },
    {
      "name": "permissions",
      "type": "LIST",
      "required": true,
      "item": {
        "name": "Permission",
        "type": "ENUM",
        "enumValues": [
          "DOWNLOAD",
          "UPLOAD"
        ]
      },
      "title": "Permissions",
      "description": "Permissions"
    },
    {
      "name": "expiration",
      "type": "TIMESTAMP",
      "required": true,
      "title": "Expiration",
      "description": "Expiration"
    },
    {
      "name": "signature",
      "type": "STRING",
      "required": true,
      "title": "Signature",
      "description": "Signature"
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
    }
  ],
  "title": "Signature",
  "description": "Signature",
  "annotations": {
    "NormalizedResource": "true"
  }
} as unknown

