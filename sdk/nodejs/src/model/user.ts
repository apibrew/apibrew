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

export const UserEntityInfo = {
    namespace: "system",
    resource: "User",
    restPath: "system-user",
}

export interface AuditData {
    createdBy: string
    updatedBy: string
    createdOn: string | Date
    updatedOn: string | Date
}


