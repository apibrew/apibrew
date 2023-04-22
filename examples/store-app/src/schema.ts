import {Entity} from "apibrew-client";


export interface Pet extends Entity {

    id: string;
    name: string;
    description: string;
    tags: string[];
    status: string;
    version: number;
    
}

export interface Namespace extends Entity {

    id: string;
    version: number;
    createdBy: string;
    updatedBy: string;
    createdOn: string;
    updatedOn: string;
    name: string;
    description: string;
    details: object;
    securityContext: object;
    
}

export interface User extends Entity {

    id: string;
    version: number;
    createdBy: string;
    updatedBy: string;
    createdOn: string;
    updatedOn: string;
    username: string;
    password: string;
    securityContext: object;
    details: object;
    
}

export interface DataSource extends Entity {

    id: string;
    version: number;
    createdBy: string;
    updatedBy: string;
    createdOn: string;
    updatedOn: string;
    name: string;
    description: string;
    backend: number;
    optionsPostgresUsername: string;
    optionsPostgresPassword: string;
    optionsPostgresHost: string;
    optionsPostgresPort: number;
    optionsPostgresDbName: string;
    optionsPostgresDefaultSchema: string;
    optionsMysqlUsername: string;
    optionsMysqlPassword: string;
    optionsMysqlHost: string;
    optionsMysqlPort: number;
    optionsMysqlDbName: string;
    optionsMysqlDefaultSchema: string;
    optionsRedisAddr: string;
    optionsRedisPassword: string;
    optionsRedisDb: number;
    optionsMongoUri: string;
    optionsMongoDbName: string;
    
}

export interface Extension extends Entity {

    id: string;
    version: number;
    createdBy: string;
    updatedBy: string;
    createdOn: string;
    updatedOn: string;
    name: string;
    description: string;
    namespace: string;
    resource: string;
    before: object;
    after: object;
    instead: object;
    
}

export interface Resource extends Entity {

    id: string;
    version: number;
    createdBy: string;
    updatedBy: string;
    createdOn: string;
    updatedOn: string;
    name: string;
    namespace: Namespace;
    virtual: boolean;
    immutable: boolean;
    abstract: boolean;
    dataSource: DataSource;
    entity: string;
    catalog: string;
    annotations: object;
    indexes: object;
    securityContext: object;
    title: string;
    description: string;
    
}

