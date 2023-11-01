
export interface DataSource {
    id: string
    version: number
    auditData?: AuditData
    name: string
    description?: string
    backend: Backend
    options: { [key: string]: string }
}

export const DataSourceEntityInfo = {
    namespace: "system",
    resource: "DataSource",
    restPath: "system-datasource",
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


