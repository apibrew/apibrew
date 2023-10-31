
export interface Record {
    id: string
    properties: object
    packedProperties?: object[]
}

export const NAMESPACE = "system";
export const RESOURCE = "Record";
export const REST_PATH = "system-record"

export const RecordEntityInfo = {
    namespace: NAMESPACE,
    resource: RESOURCE,
    restPath: REST_PATH,
}


