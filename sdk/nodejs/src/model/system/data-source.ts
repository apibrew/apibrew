


export const DataSourceResource = {
    resource: "DataSource",
    namespace: "system",
};

// Sub Types

// Resource Type
export interface DataSource {
    id: string;
version: number;
createdBy?: string;
updatedBy?: string;
createdOn?: string;
updatedOn?: string;
name: string;
description?: string;
backend: 'POSTGRESQL' | 'MYSQL' | 'MONGODB' | 'REDIS';
options: object;

}
// Resource and Property Names
export const DataSourceName = "DataSource";

export const DataSourceIdName = "Id";

export const DataSourceVersionName = "Version";

export const DataSourceCreatedByName = "CreatedBy";

export const DataSourceUpdatedByName = "UpdatedBy";

export const DataSourceCreatedOnName = "CreatedOn";

export const DataSourceUpdatedOnName = "UpdatedOn";

export const DataSourceNameName = "Name";

export const DataSourceDescriptionName = "Description";

export const DataSourceBackendName = "Backend";

export const DataSourceOptionsName = "Options";


