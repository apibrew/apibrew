import {User} from './user';
import {Role} from './role';

export interface Permission {
    id: string
    version: number
    auditData?: AuditData
    namespace?: string
    resource?: string
    property?: string
    propertyValue?: string
    propertyMode?: PropertyMode
    operation: Operation
    recordIds?: string[]
    before?: string | Date
    after?: string | Date
    user?: User
    role?: Role
    permit: Permit
    localFlags?: object
}

export const NAMESPACE = "system";
export const RESOURCE = "Permission";
export const REST_PATH = "system-permission"

export const PermissionEntityInfo = {
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

export enum PropertyMode {
    PROPERTY_MATCH_ONLY = "PROPERTY_MATCH_ONLY",
    PROPERTY_MATCH_ANY = "PROPERTY_MATCH_ANY",
}

export enum Operation {
    READ = "READ",
    CREATE = "CREATE",
    UPDATE = "UPDATE",
    DELETE = "DELETE",
    FULL = "FULL",
}

export enum Permit {
    ALLOW = "ALLOW",
    REJECT = "REJECT",
}


