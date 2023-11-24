
export interface AuditLog {
    id: string
    version: number
    namespace: string
    resource: string
    recordId: string
    time: string | Date
    username: string
    operation: Operation
    properties?: object
    annotations?: { [key: string]: string }
}

export const AuditLogEntityInfo = {
    namespace: "system",
    resource: "AuditLog",
    restPath: "system-auditlog",
}

export enum Operation {
    CREATE = "CREATE",
    UPDATE = "UPDATE",
    DELETE = "DELETE",
}

export const AuditLogResource = {
  "name": "AuditLog",
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
    "namespace": {
      "type": "STRING",
      "required": true,
      "length": 256
    },
    "operation": {
      "type": "ENUM",
      "required": true,
      "enumValues": [
        "CREATE",
        "UPDATE",
        "DELETE"
      ]
    },
    "properties": {
      "type": "OBJECT"
    },
    "recordId": {
      "type": "STRING",
      "required": true,
      "length": 256,
      "annotations": {
        "SourceDef": "record_id"
      }
    },
    "resource": {
      "type": "STRING",
      "required": true,
      "length": 256
    },
    "time": {
      "type": "TIMESTAMP",
      "required": true
    },
    "username": {
      "type": "STRING",
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
  "immutable": true,
  "annotations": {
    "BypassExtensions": "true",
    "OpenApiGroup": "internal"
  }
} as unknown

