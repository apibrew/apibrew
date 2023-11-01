import {Permission} from './permission';

export interface Role {
    id: string
    version: number
    auditData?: AuditData
    name: string
    permissions?: Permission[]
    details?: object
}

export const RoleEntityInfo = {
    namespace: "system",
    resource: "Role",
    restPath: "system-role",
}

export interface AuditData {
    createdBy: string
    updatedBy: string
    createdOn: string | Date
    updatedOn: string | Date
}


