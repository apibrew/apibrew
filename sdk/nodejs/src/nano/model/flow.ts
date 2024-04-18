
export interface Flow {
    id: string
    name: string
    trigger: string
    version: number
    statements: Statement[]
}

export const FlowEntityInfo = {
    namespace: "nano",
    resource: "Flow",
    restPath: "nano-flow",
}

export interface EventParams {
    action: Action
    responds: boolean
    finalizes: boolean
    annotations: { [key: string]: string }
    sync: boolean
    type: string
    order: Order
}

export interface ActionParams {
    type: string
}

export interface ApiSaveParams {
    type: string
    payload: object
}

export interface ApiLoadParams {
    type: string
    match: object
    params: object
}

export interface AssignParams {
    left: string
    expression: string
}

export interface FunctionCallParams {
    name: string
    params: object
}

export interface CodeParams {
    content: string
}

export interface ConditionParams {
    fail: Statement[]
    pass: Statement[]
    condition: string
}

export interface FailParams {
    message: string
}

export interface Statement {
    variable: string
    checkResult: boolean
    kind: Kind
    params: object
}

export enum Action {
    CREATE = "CREATE",
    GET = "GET",
    LOAD = "LOAD",
    UPDATE = "UPDATE",
    DELETE = "DELETE",
    LIST = "LIST",
}

export enum Order {
    BEFORE = "BEFORE",
    AFTER = "AFTER",
}

export enum Kind {
    ACTION = "ACTION",
    EVENT = "EVENT",
    API_CREATE = "API_CREATE",
    API_GET = "API_GET",
    API_LOAD = "API_LOAD",
    API_UPDATE = "API_UPDATE",
    API_DELETE = "API_DELETE",
    API_LIST = "API_LIST",
    CODE = "CODE",
    FAIL = "FAIL",
    ASSIGN = "ASSIGN",
    CONDITION = "CONDITION",
    HTTP_CALL = "HTTP_CALL",
    FUNCTION_CALL = "FUNCTION_CALL",
    TEMPLATE_CALL = "TEMPLATE_CALL",
    END = "END",
}

export const FlowResource = {
  "auditData": {
    "createdBy": "admin",
    "createdOn": "2024-04-17T10:31:07Z"
  },
  "name": "Flow",
  "namespace": {
    "name": "nano"
  },
  "properties": {
    "id": {
      "type": "UUID",
      "primary": true,
      "required": true,
      "immutable": true,
      "exampleValue": "a39621a4-6d48-11ee-b962-0242ac120002",
      "description": "The unique identifier of the resource. It is randomly generated and immutable.",
      "annotations": {
        "SpecialProperty": "true"
      }
    },
    "name": {
      "type": "STRING",
      "required": true,
      "unique": true,
      "length": 255,
      "title": "Name",
      "description": "Name"
    },
    "statements": {
      "type": "LIST",
      "required": true,
      "item": {
        "type": "STRUCT",
        "typeRef": "Statement"
      },
      "defaultValue": null
    },
    "trigger": {
      "type": "STRING",
      "required": true,
      "length": 255,
      "title": "Trigger",
      "description": "Trigger"
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
  "types": [
    {
      "name": "EventParams",
      "title": "",
      "description": "",
      "properties": {
        "action": {
          "type": "ENUM",
          "enumValues": [
            "CREATE",
            "GET",
            "LOAD",
            "UPDATE",
            "DELETE",
            "LIST"
          ]
        },
        "annotations": {
          "type": "MAP",
          "item": {
            "type": "STRING"
          }
        },
        "finalizes": {
          "type": "BOOL"
        },
        "order": {
          "type": "ENUM",
          "enumValues": [
            "BEFORE",
            "AFTER"
          ]
        },
        "responds": {
          "type": "BOOL"
        },
        "sync": {
          "type": "BOOL"
        },
        "type": {
          "type": "STRING"
        }
      }
    },
    {
      "name": "ActionParams",
      "title": "",
      "description": "",
      "properties": {
        "type": {
          "type": "STRING"
        }
      }
    },
    {
      "name": "ApiSaveParams",
      "title": "",
      "description": "",
      "properties": {
        "payload": {
          "type": "OBJECT"
        },
        "type": {
          "type": "STRING"
        }
      }
    },
    {
      "name": "ApiLoadParams",
      "title": "",
      "description": "",
      "properties": {
        "match": {
          "type": "OBJECT"
        },
        "params": {
          "type": "OBJECT"
        },
        "type": {
          "type": "STRING"
        }
      }
    },
    {
      "name": "AssignParams",
      "title": "",
      "description": "",
      "properties": {
        "expression": {
          "type": "STRING"
        },
        "left": {
          "type": "STRING"
        }
      }
    },
    {
      "name": "FunctionCallParams",
      "title": "",
      "description": "",
      "properties": {
        "name": {
          "type": "STRING"
        },
        "params": {
          "type": "OBJECT"
        }
      }
    },
    {
      "name": "CodeParams",
      "title": "",
      "description": "",
      "properties": {
        "content": {
          "type": "STRING"
        }
      }
    },
    {
      "name": "ConditionParams",
      "title": "",
      "description": "",
      "properties": {
        "condition": {
          "type": "STRING"
        },
        "fail": {
          "type": "LIST",
          "required": true,
          "item": {
            "type": "STRUCT",
            "typeRef": "Statement"
          },
          "defaultValue": null
        },
        "pass": {
          "type": "LIST",
          "required": true,
          "item": {
            "type": "STRUCT",
            "typeRef": "Statement"
          },
          "defaultValue": null
        }
      }
    },
    {
      "name": "FailParams",
      "title": "",
      "description": "",
      "properties": {
        "message": {
          "type": "STRING"
        }
      }
    },
    {
      "name": "Statement",
      "title": "",
      "description": "",
      "properties": {
        "checkResult": {
          "type": "BOOL"
        },
        "kind": {
          "type": "ENUM",
          "enumValues": [
            "ACTION",
            "EVENT",
            "API_CREATE",
            "API_GET",
            "API_LOAD",
            "API_UPDATE",
            "API_DELETE",
            "API_LIST",
            "CODE",
            "FAIL",
            "ASSIGN",
            "CONDITION",
            "HTTP_CALL",
            "FUNCTION_CALL",
            "TEMPLATE_CALL",
            "END"
          ]
        },
        "params": {
          "type": "OBJECT"
        },
        "variable": {
          "type": "STRING"
        }
      }
    }
  ]
} as unknown

