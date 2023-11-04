
export interface Record {
    id: string
    properties: object
    packedProperties?: object[]
}

export const RecordEntityInfo = {
    namespace: "system",
    resource: "Record",
    restPath: "system-record",
}

export const RecordResource = {
  "name": "Record",
  "namespace": {
    "name": "system"
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
      "name": "properties",
      "type": "OBJECT",
      "required": true,
      "title": "Properties",
      "description": "The properties of the record. The schema of properties are defined in the resource definition. \nHere you will put the payload corresponding to the resource definition.\n"
    },
    {
      "name": "packedProperties",
      "type": "LIST",
      "item": {
        "name": "",
        "type": "OBJECT"
      },
      "annotations": {
        "OpenApiHide": "true"
      }
    }
  ],
  "title": "Generic Record",
  "description": "A generic record resource. All Apis are extended from Generic Record resource",
  "annotations": {
    "RestApiDisabled": "true"
  }
} as unknown

