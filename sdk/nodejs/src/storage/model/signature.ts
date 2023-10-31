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

export const NAMESPACE = "storage";
export const RESOURCE = "Signature";
export const REST_PATH = "storage-signature"

export const SignatureEntityInfo = {
    namespace: NAMESPACE,
    resource: RESOURCE,
    restPath: REST_PATH,
}

export enum Permission {
    DOWNLOAD = "DOWNLOAD",
    UPLOAD = "UPLOAD",
}


