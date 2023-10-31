
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

export const NAMESPACE = "storage";
export const RESOURCE = "StorageObject";
export const REST_PATH = "storage-storageobject"

export const StorageObjectEntityInfo = {
    namespace: NAMESPACE,
    resource: RESOURCE,
    restPath: REST_PATH,
}

export interface AuditData {
    createdBy: string
    updatedBy: string
    createdOn: string | Date
    updatedOn: string | Date
}


