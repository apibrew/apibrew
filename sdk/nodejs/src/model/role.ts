import {Permission} from './permission';

export interface Role {
    id: string
    version: number
    auditData?: AuditData
    name: string
    permissions?: Permission[]
    details?: object
}

export const NAMESPACE = "system";
export const RESOURCE = "Role";
export const REST_PATH = "system-role"

export const RoleEntityInfo = {
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


