export declare const DataSourceResource: {
    resource: string;
    namespace: string;
};
export interface DataSource {
    id: string;
    version: number;
    createdBy: string;
    updatedBy?: string;
    createdOn: string;
    updatedOn?: string;
    name: string;
    description?: string;
    backend: 'POSTGRESQL' | 'MYSQL' | 'MONGODB' | 'REDIS';
    options: object;
}
export declare const DataSourceName = "DataSource";
export declare const DataSourceIdName = "Id";
export declare const DataSourceVersionName = "Version";
export declare const DataSourceCreatedByName = "CreatedBy";
export declare const DataSourceUpdatedByName = "UpdatedBy";
export declare const DataSourceCreatedOnName = "CreatedOn";
export declare const DataSourceUpdatedOnName = "UpdatedOn";
export declare const DataSourceNameName = "Name";
export declare const DataSourceDescriptionName = "Description";
export declare const DataSourceBackendName = "Backend";
export declare const DataSourceOptionsName = "Options";
