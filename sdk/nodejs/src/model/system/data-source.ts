


export const DataSourceResource = {
    resource: "data-source",
    namespace: "system",
};

// Sub Types

// Resource Type
export interface DataSource {
    id: string;
version: number;
createdBy: string;
updatedBy?: string;
createdOn: string;
updatedOn?: string;
name: string;
description?: string;
backend: number;
optionsPostgresUsername?: string;
optionsPostgresPassword?: string;
optionsPostgresHost?: string;
optionsPostgresPort?: number;
optionsPostgresDbName?: string;
optionsPostgresDefaultSchema?: string;
optionsMysqlUsername?: string;
optionsMysqlPassword?: string;
optionsMysqlHost?: string;
optionsMysqlPort?: number;
optionsMysqlDbName?: string;
optionsMysqlDefaultSchema?: string;
optionsRedisAddr?: string;
optionsRedisPassword?: string;
optionsRedisDb?: number;
optionsMongoUri?: string;
optionsMongoDbName?: string;

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

export const DataSourceOptionsPostgresUsernameName = "OptionsPostgresUsername";

export const DataSourceOptionsPostgresPasswordName = "OptionsPostgresPassword";

export const DataSourceOptionsPostgresHostName = "OptionsPostgresHost";

export const DataSourceOptionsPostgresPortName = "OptionsPostgresPort";

export const DataSourceOptionsPostgresDbNameName = "OptionsPostgresDbName";

export const DataSourceOptionsPostgresDefaultSchemaName = "OptionsPostgresDefaultSchema";

export const DataSourceOptionsMysqlUsernameName = "OptionsMysqlUsername";

export const DataSourceOptionsMysqlPasswordName = "OptionsMysqlPassword";

export const DataSourceOptionsMysqlHostName = "OptionsMysqlHost";

export const DataSourceOptionsMysqlPortName = "OptionsMysqlPort";

export const DataSourceOptionsMysqlDbNameName = "OptionsMysqlDbName";

export const DataSourceOptionsMysqlDefaultSchemaName = "OptionsMysqlDefaultSchema";

export const DataSourceOptionsRedisAddrName = "OptionsRedisAddr";

export const DataSourceOptionsRedisPasswordName = "OptionsRedisPassword";

export const DataSourceOptionsRedisDbName = "OptionsRedisDb";

export const DataSourceOptionsMongoUriName = "OptionsMongoUri";

export const DataSourceOptionsMongoDbNameName = "OptionsMongoDbName";


