
export interface Namespace {
    id: string
    version: number
    auditData?: AuditData
    name: string
    description?: string
    details?: object
}

export const NamespaceEntityInfo = {
    namespace: "system",
    resource: "Namespace",
    restPath: "system-namespace",
}

export interface AuditData {
    createdBy: string
    updatedBy: string
    createdOn: string | Date
    updatedOn: string | Date
}


