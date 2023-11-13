import {User} from './user';
import {Role} from './role';

export interface Permission {
    id: string
    version: number
    auditData?: AuditData
    namespace?: string
    resource?: string
    recordSelector?: BooleanExpression
    operation: Operation
    before?: string | Date
    after?: string | Date
    user?: User
    role?: Role
    permit: Permit
    localFlags?: object
}

export const PermissionEntityInfo = {
    namespace: "system",
    resource: "Permission",
    restPath: "system-permission",
}

export interface AuditData {
    createdBy: string
    updatedBy: string
    createdOn: string | Date
    updatedOn: string | Date
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

export enum Operation {
    READ = "READ",
    CREATE = "CREATE",
    UPDATE = "UPDATE",
    DELETE = "DELETE",
    FULL = "FULL",
}

export enum Permit {
    ALLOW = "ALLOW",
    REJECT = "REJECT",
}

export const PermissionResource = {
  "name": "Permission",
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
      "name": "namespace",
      "type": "STRING",
      "length": 255,
      "exampleValue": "default",
      "title": "Namespace",
      "description": "The namespace(name) of the resource. If given it will be used to match the resource by namespace."
    },
    {
      "name": "resource",
      "type": "STRING",
      "length": 255,
      "exampleValue": "Book",
      "title": "Resource",
      "description": "The name of the resource. If given it will be used to match the resource by name."
    },
    {
      "name": "recordSelector",
      "type": "STRUCT",
      "typeRef": "BooleanExpression"
    },
    {
      "name": "operation",
      "type": "ENUM",
      "required": true,
      "length": 255,
      "defaultValue": "FULL",
      "enumValues": [
        "READ",
        "CREATE",
        "UPDATE",
        "DELETE",
        "FULL"
      ],
      "exampleValue": "READ",
      "title": "Operation",
      "description": "The operation of the permission. It is used to match the operation of the request. If given it will be used to match the operation of the request."
    },
    {
      "name": "before",
      "type": "TIMESTAMP",
      "title": "Before",
      "description": "The timestamp before which the permission is valid. If given it will be used to match the timestamp of the request."
    },
    {
      "name": "after",
      "type": "TIMESTAMP",
      "title": "After",
      "description": "The timestamp after which the permission is valid. If given it will be used to match the timestamp of the request."
    },
    {
      "name": "user",
      "type": "REFERENCE",
      "reference": {
        "resource": {
          "name": "User",
          "namespace": {
            "name": "system"
          }
        },
        "cascade": false
      },
      "title": "User",
      "description": "The user who has the permission. If given it will be used to match the user of the request. It is ignored by default, because if permissions is set through User this property is overrides and auto-populated by system"
    },
    {
      "name": "role",
      "type": "REFERENCE",
      "reference": {
        "resource": {
          "name": "Role",
          "namespace": {
            "name": "system"
          }
        },
        "cascade": false
      },
      "title": "Role",
      "description": "The role who has the permission. If given it will be used to match the role of the request. It is ignored by default, because if permissions is set through Role this property is overrides and auto-populated by system"
    },
    {
      "name": "permit",
      "type": "ENUM",
      "required": true,
      "length": 255,
      "enumValues": [
        "ALLOW",
        "REJECT"
      ],
      "title": "Permit",
      "description": "The permit of the permission. If permission is matched, this property is judging field to indicate that if operation is allowed or not"
    },
    {
      "name": "localFlags",
      "type": "OBJECT"
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
    },
    {
      "name": "BooleanExpression",
      "title": "",
      "description": "",
      "properties": [
        {
          "name": "and",
          "type": "LIST",
          "item": {
            "name": "",
            "type": "STRUCT",
            "typeRef": "BooleanExpression"
          }
        },
        {
          "name": "or",
          "type": "LIST",
          "item": {
            "name": "",
            "type": "STRUCT",
            "typeRef": "BooleanExpression"
          }
        },
        {
          "name": "not",
          "type": "STRUCT",
          "typeRef": "BooleanExpression"
        },
        {
          "name": "equal",
          "type": "STRUCT",
          "typeRef": "PairExpression"
        },
        {
          "name": "lessThan",
          "type": "STRUCT",
          "typeRef": "PairExpression"
        },
        {
          "name": "greaterThan",
          "type": "STRUCT",
          "typeRef": "PairExpression"
        },
        {
          "name": "lessThanOrEqual",
          "type": "STRUCT",
          "typeRef": "PairExpression"
        },
        {
          "name": "greaterThanOrEqual",
          "type": "STRUCT",
          "typeRef": "PairExpression"
        },
        {
          "name": "in",
          "type": "STRUCT",
          "typeRef": "PairExpression"
        },
        {
          "name": "isNull",
          "type": "STRUCT",
          "typeRef": "Expression"
        },
        {
          "name": "regexMatch",
          "type": "STRUCT",
          "typeRef": "RegexMatchExpression"
        }
      ]
    },
    {
      "name": "PairExpression",
      "title": "",
      "description": "",
      "properties": [
        {
          "name": "left",
          "type": "STRUCT",
          "typeRef": "Expression"
        },
        {
          "name": "right",
          "type": "STRUCT",
          "typeRef": "Expression"
        }
      ]
    },
    {
      "name": "RegexMatchExpression",
      "title": "",
      "description": "",
      "properties": [
        {
          "name": "pattern",
          "type": "STRING"
        },
        {
          "name": "expression",
          "type": "STRUCT",
          "typeRef": "Expression"
        }
      ]
    },
    {
      "name": "Expression",
      "title": "",
      "description": "",
      "properties": [
        {
          "name": "property",
          "type": "STRING"
        },
        {
          "name": "value",
          "type": "OBJECT"
        }
      ]
    }
  ],
  "title": "Permission",
  "description": "Permission is a resource that defines the access control rules for resources for users.",
  "annotations": {
    "EnableAudit": "true",
    "OpenApiGroup": "meta"
  }
} as unknown

