import {Resource} from './resource';

export interface ResourceAction {
    id: string
    version: number
    auditData?: AuditData
    resource: Resource
    name: string
    title?: string
    description?: string
    internal: boolean
    types?: SubType[]
    input?: Property[]
    output?: Property
    annotations?: { [key: string]: string }
}

export const ResourceActionEntityInfo = {
    namespace: "system",
    resource: "ResourceAction",
    restPath: "system-resourceaction",
}

export interface SubType {
    name: string
    title: string
    description: string
    properties: Property[]
}

export interface Reference {
    resource: Resource
    cascade: boolean
    backReference: string
}

export interface Property {
    name: string
    type: Type
    typeRef: string
    primary: boolean
    required: boolean
    unique: boolean
    immutable: boolean
    length: number
    item: Property
    reference: Reference
    defaultValue: object
    enumValues: string[]
    exampleValue: object
    title: string
    description: string
    annotations: { [key: string]: string }
}

export interface AuditData {
    createdBy: string
    updatedBy: string
    createdOn: string | Date
    updatedOn: string | Date
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

export const ResourceActionResource = {
  "name": "ResourceAction",
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
      "name": "resource",
      "type": "REFERENCE",
      "required": true,
      "reference": {
        "resource": {
          "name": "Resource",
          "namespace": {
            "name": "system"
          }
        },
        "cascade": false
      }
    },
    {
      "name": "name",
      "type": "STRING",
      "required": true,
      "length": 256,
      "annotations": {
        "IsHclLabel": "true"
      }
    },
    {
      "name": "title",
      "type": "STRING",
      "length": 256,
      "annotations": {
        "IsHclLabel": "true"
      }
    },
    {
      "name": "description",
      "type": "STRING",
      "length": 256,
      "annotations": {
        "IsHclLabel": "true"
      }
    },
    {
      "name": "internal",
      "type": "BOOL",
      "required": true,
      "defaultValue": false,
      "annotations": {
        "IsHclLabel": "true"
      }
    },
    {
      "name": "types",
      "type": "LIST",
      "item": {
        "name": "",
        "type": "STRUCT",
        "typeRef": "SubType"
      }
    },
    {
      "name": "input",
      "type": "LIST",
      "item": {
        "name": "",
        "type": "STRUCT",
        "typeRef": "Property"
      }
    },
    {
      "name": "output",
      "type": "STRUCT",
      "typeRef": "Property"
    },
    {
      "name": "annotations",
      "type": "MAP",
      "item": {
        "name": "",
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
    }
  ],
  "indexes": [
    {
      "properties": [
        {
          "name": "resource",
          "order": "UNKNOWN"
        },
        {
          "name": "name",
          "order": "UNKNOWN"
        }
      ],
      "unique": true
    }
  ],
  "types": [
    {
      "name": "SubType",
      "title": "Sub Type",
      "description": "Sub Type is a type that represents a sub type of a resource. It is mostly used by STRUCT type to define the properties of the struct. ",
      "properties": [
        {
          "name": "name",
          "type": "STRING",
          "required": true,
          "exampleValue": "Book",
          "title": "Name",
          "description": "The name of the sub type. "
        },
        {
          "name": "title",
          "type": "STRING",
          "length": 256,
          "exampleValue": "Book",
          "title": "Title",
          "description": "The title of the sub type. It is used to have meaningful names for the sub types."
        },
        {
          "name": "description",
          "type": "STRING",
          "length": 256,
          "exampleValue": "Book is a sub type of Resource. It represents a book in the system. ",
          "title": "Description",
          "description": "The description of the sub type. It is used to have meaningful description for the sub types. "
        },
        {
          "name": "properties",
          "type": "LIST",
          "required": true,
          "item": {
            "name": "",
            "type": "STRUCT",
            "typeRef": "Property"
          },
          "exampleValue": [
            {
              "name": "title",
              "type": "STRING"
            }
          ],
          "title": "Properties",
          "description": "The properties of the sub type. It is used to define the properties of the sub type. "
        }
      ]
    },
    {
      "name": "Reference",
      "title": "Reference",
      "description": "Reference is a type that represents a reference to another resource. It is used to define the reference to another resource. ",
      "properties": [
        {
          "name": "resource",
          "type": "REFERENCE",
          "reference": {
            "resource": {
              "name": "Resource",
              "namespace": {
                "name": "system"
              }
            },
            "cascade": false
          },
          "exampleValue": {
            "name": "Book",
            "namespace": {
              "name": "test-namespace"
            }
          },
          "title": "Resource",
          "description": "This property indicates the resource of the reference.\nWhen providing resource, you need to provide namespace and resource name of the referenced resource.\nIf you don't provide namespace, it will be assumed as the same namespace with the resource.\n"
        },
        {
          "name": "cascade",
          "type": "BOOL",
          "title": "Cascade",
          "description": "This property indicates that whether or not given reference is cascade.\nIf it is true, when referenced resource record is deleted, all the records that are referencing to that resource will be deleted.\n"
        },
        {
          "name": "backReference",
          "type": "STRING",
          "exampleValue": "author",
          "title": "Back Reference",
          "description": "This property indicates that whether or not given reference is back reference.\nBack reference is reverse of reference, If resource A has reference to resource B, in that case resource B can have back reference to resource A.\nBack reference is used only as List.\nBackreference should be the name of property in the referenced resource. (like author inside book)\nFor example:\n\tBook -\u003e Author.\n\tBook will have reference to Author. And Author can have back reference to the list of books\n\n"
        }
      ]
    },
    {
      "name": "Property",
      "title": "Property",
      "description": "Property is a type that represents a property of a resource. It is like an API properties or properties of class in a programming language",
      "properties": [
        {
          "name": "name",
          "type": "STRING",
          "length": 256,
          "exampleValue": "title",
          "title": "Name",
          "description": "The name of the property. \nName is the main parameter of property, it is used to identify the property. It is also used to name record properties. \nFor example {\"title\": \"Lord of the Rings\"} And there \"title\" is a property, and it is defined by name \"title\", in its Resource \n\t\t"
        },
        {
          "name": "type",
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
          "exampleValue": "STRING",
          "title": "Type",
          "description": "The type of the property. Property Data Types can be one of it. Types can be written with all capital letters."
        },
        {
          "name": "typeRef",
          "type": "STRING",
          "length": 256,
          "exampleValue": "BookPublishingDetails",
          "title": "Type Reference",
          "description": "The type reference of the property. It is only used for STRUCT type. \nWhen you used STRUCT type, you need to define your type inside types of resource and then you can use its name as typeRef."
        },
        {
          "name": "primary",
          "type": "BOOL",
          "required": true,
          "defaultValue": false,
          "title": "Primary",
          "description": "The primary property of the resource. It is used to identify the resource. When it is not supplied, an id property is automatically created.\nNormally primary property should not be provided. It is only used for special cases. If provided, it can break some functionalities of system. \nIf Primary is provided, it should be a single property. It can not be a list or map.\nIf Primary is provided, internal id property will not be created.\n"
        },
        {
          "name": "required",
          "type": "BOOL",
          "required": true,
          "defaultValue": false,
          "title": "Required",
          "description": "This property indicates that whether or not given property is required.\nWhen creating/updating records, if required property is not and defaultValue is given in property, the system will allow request but will use default value instead.\n(In all cases if default value is provided it will be used in case of property absence)\n"
        },
        {
          "name": "unique",
          "type": "BOOL",
          "required": true,
          "defaultValue": false,
          "title": "Unique",
          "description": "This property indicates that whether or not given property is unique.\nUnique property is only working for single property, for combination of properties to become unique, you can use indexes with unique flag \n"
        },
        {
          "name": "immutable",
          "type": "BOOL",
          "required": true,
          "defaultValue": false,
          "title": "Immutable",
          "description": "This property indicates that whether or not given property is immutable. Immutable properties can not be updated."
        },
        {
          "name": "length",
          "type": "INT32",
          "required": true,
          "defaultValue": 256,
          "exampleValue": 256,
          "title": "Length",
          "description": "This property indicates the length of the property. It is only used for STRING type."
        },
        {
          "name": "item",
          "type": "STRUCT",
          "typeRef": "Property",
          "exampleValue": {
            "type": "STRING"
          },
          "title": "Item",
          "description": "This property indicates the item type of the property. It is only used for LIST and MAP types."
        },
        {
          "name": "reference",
          "type": "STRUCT",
          "typeRef": "Reference",
          "exampleValue": {
            "resource": {
              "namespace": "default",
              "resource": "Book"
            }
          },
          "title": "Reference",
          "description": "This property indicates the reference type of the property. It is only used for REFERENCE type.\nWhen you use REFERENCE type, you need to provide reference details.\nReference details is used to locate referenced resource\nWhen providing reference details, you need to provide namespace and resource name of the referenced resource.\nIf you don't provide namespace, it will be assumed as the same namespace with the resource.\n"
        },
        {
          "name": "defaultValue",
          "type": "OBJECT",
          "exampleValue": "Lord of the Rings",
          "title": "Default Value",
          "description": "This property indicates the default value of the property. \nIt is used when creating/updating records and property is not provided.\n"
        },
        {
          "name": "enumValues",
          "type": "LIST",
          "item": {
            "name": "",
            "type": "STRING"
          },
          "exampleValue": [
            "UNKNOWN",
            "ASC",
            "DESC"
          ],
          "title": "Enum Values",
          "description": "This property is only used with ENUM type. This property indicates the enum values of the property."
        },
        {
          "name": "exampleValue",
          "type": "OBJECT",
          "exampleValue": "no-book-name",
          "title": "Example Value",
          "description": "This property indicates the example value of the property."
        },
        {
          "name": "title",
          "type": "STRING",
          "length": 256,
          "exampleValue": "Book Title",
          "title": "Title",
          "description": "This property indicates the title of the property. It is used to have meaningful names for the properties."
        },
        {
          "name": "description",
          "type": "STRING",
          "length": 256,
          "exampleValue": "Book Title is a property of Book Resource. It represents the title of the book.",
          "title": "Description",
          "description": "This property indicates the description of the property. It is used to have meaningful description for the properties."
        },
        {
          "name": "annotations",
          "type": "MAP",
          "item": {
            "name": "",
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
        }
      ]
    },
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
    "OpenApiGroup": "meta",
    "RestApiDisabled": "true"
  }
} as unknown

