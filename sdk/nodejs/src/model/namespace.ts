
export interface Namespace {
    id: string
    version: number
    auditData?: AuditData
    name: string
    description?: string
    details?: object
}

export const NAMESPACE = "system";
export const RESOURCE = "Namespace";
export const REST_PATH = "system-namespace"

export const NamespaceEntityInfo = {
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


