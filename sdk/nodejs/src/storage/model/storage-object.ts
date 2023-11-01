
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


