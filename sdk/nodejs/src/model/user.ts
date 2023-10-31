import {Role} from './role';
import {Permission} from './permission';

export interface User {
    id: string
    version: number
    auditData?: AuditData
    username: string
    password?: string
    roles?: Role[]
    permissions?: Permission[]
    details?: object
}

export const NAMESPACE = "system";
export const RESOURCE = "User";
export const REST_PATH = "system-user"

export const UserEntityInfo = {
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


