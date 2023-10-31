
export interface AuditLog {
    id: string
    version: number
    namespace: string
    resource: string
    recordId: string
    time: string | Date
    username: string
    operation: Operation
    properties?: object
    annotations?: { [key: string]: string }
}

export const NAMESPACE = "system";
export const RESOURCE = "AuditLog";
export const REST_PATH = "system-auditlog"

export const AuditLogEntityInfo = {
    namespace: NAMESPACE,
    resource: RESOURCE,
    restPath: REST_PATH,
}

export enum Operation {
    CREATE = "CREATE",
    UPDATE = "UPDATE",
    DELETE = "DELETE",
}


