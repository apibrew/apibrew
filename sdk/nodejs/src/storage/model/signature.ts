import {StorageObject} from './storage-object';

export interface Signature {
    id: string
    object: StorageObject
    permissions: Permission[]
    expiration: string | Date
    signature: string
    annotations?: { [key: string]: string }
    version: number
}

export const SignatureEntityInfo = {
    namespace: "storage",
    resource: "Signature",
    restPath: "storage-signature",
}

export enum Permission {
    DOWNLOAD = "DOWNLOAD",
    UPLOAD = "UPLOAD",
}


