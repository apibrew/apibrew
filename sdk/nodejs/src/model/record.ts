
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


