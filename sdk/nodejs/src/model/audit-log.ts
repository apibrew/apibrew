
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

export const AuditLogEntityInfo = {
    namespace: "system",
    resource: "AuditLog",
    restPath: "system-auditlog",
}

export enum Operation {
    CREATE = "CREATE",
    UPDATE = "UPDATE",
    DELETE = "DELETE",
}


