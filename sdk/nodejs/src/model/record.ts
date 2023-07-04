export interface RecordResourceInfo {
    namespace: string
    resource: string
}

export interface RecordCoreParams {
    id?: string


    version?: number
    createdAt?: string
    updatedAt?: string
    createdBy?: string
    updatedBy?: string
}

export type Record<T> = RecordCoreParams & T

export interface RecordListContainer<T> {
    content: T[]
    total: number
}
