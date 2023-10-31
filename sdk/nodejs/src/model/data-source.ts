
export interface DataSource {
    id: string
    version: number
    auditData?: AuditData
    name: string
    description?: string
    backend: Backend
    options: { [key: string]: string }
}

export const NAMESPACE = "system";
export const RESOURCE = "DataSource";
export const REST_PATH = "system-datasource"

export const DataSourceEntityInfo = {
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

export enum Backend {
    POSTGRESQL = "POSTGRESQL",
    MYSQL = "MYSQL",
    MONGODB = "MONGODB",
    REDIS = "REDIS",
}


