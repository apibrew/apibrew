// package: model
// file: model/data-source.proto

/* tslint:disable */
/* eslint-disable */

import * as jspb from "google-protobuf";
import * as model_audit_pb from "../model/audit_pb";
import * as model_common_pb from "../model/common_pb";
import * as model_annotations_pb from "../model/annotations_pb";

export class PostgresqlParams extends jspb.Message { 
    getUsername(): string;
    setUsername(value: string): PostgresqlParams;
    getPassword(): string;
    setPassword(value: string): PostgresqlParams;
    getHost(): string;
    setHost(value: string): PostgresqlParams;
    getPort(): number;
    setPort(value: number): PostgresqlParams;
    getDbname(): string;
    setDbname(value: string): PostgresqlParams;
    getDefaultschema(): string;
    setDefaultschema(value: string): PostgresqlParams;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): PostgresqlParams.AsObject;
    static toObject(includeInstance: boolean, msg: PostgresqlParams): PostgresqlParams.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: PostgresqlParams, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): PostgresqlParams;
    static deserializeBinaryFromReader(message: PostgresqlParams, reader: jspb.BinaryReader): PostgresqlParams;
}

export namespace PostgresqlParams {
    export type AsObject = {
        username: string,
        password: string,
        host: string,
        port: number,
        dbname: string,
        defaultschema: string,
    }
}

export class MysqlParams extends jspb.Message { 
    getUsername(): string;
    setUsername(value: string): MysqlParams;
    getPassword(): string;
    setPassword(value: string): MysqlParams;
    getHost(): string;
    setHost(value: string): MysqlParams;
    getPort(): number;
    setPort(value: number): MysqlParams;
    getDbname(): string;
    setDbname(value: string): MysqlParams;
    getDefaultschema(): string;
    setDefaultschema(value: string): MysqlParams;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): MysqlParams.AsObject;
    static toObject(includeInstance: boolean, msg: MysqlParams): MysqlParams.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: MysqlParams, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): MysqlParams;
    static deserializeBinaryFromReader(message: MysqlParams, reader: jspb.BinaryReader): MysqlParams;
}

export namespace MysqlParams {
    export type AsObject = {
        username: string,
        password: string,
        host: string,
        port: number,
        dbname: string,
        defaultschema: string,
    }
}

export class RedisParams extends jspb.Message { 
    getAddr(): string;
    setAddr(value: string): RedisParams;
    getPassword(): string;
    setPassword(value: string): RedisParams;
    getDb(): number;
    setDb(value: number): RedisParams;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): RedisParams.AsObject;
    static toObject(includeInstance: boolean, msg: RedisParams): RedisParams.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: RedisParams, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): RedisParams;
    static deserializeBinaryFromReader(message: RedisParams, reader: jspb.BinaryReader): RedisParams;
}

export namespace RedisParams {
    export type AsObject = {
        addr: string,
        password: string,
        db: number,
    }
}

export class MongoParams extends jspb.Message { 
    getUri(): string;
    setUri(value: string): MongoParams;
    getDbname(): string;
    setDbname(value: string): MongoParams;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): MongoParams.AsObject;
    static toObject(includeInstance: boolean, msg: MongoParams): MongoParams.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: MongoParams, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): MongoParams;
    static deserializeBinaryFromReader(message: MongoParams, reader: jspb.BinaryReader): MongoParams;
}

export namespace MongoParams {
    export type AsObject = {
        uri: string,
        dbname: string,
    }
}

export class VirtualParams extends jspb.Message { 
    getMode(): VirtualParams.Mode;
    setMode(value: VirtualParams.Mode): VirtualParams;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): VirtualParams.AsObject;
    static toObject(includeInstance: boolean, msg: VirtualParams): VirtualParams.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: VirtualParams, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): VirtualParams;
    static deserializeBinaryFromReader(message: VirtualParams, reader: jspb.BinaryReader): VirtualParams;
}

export namespace VirtualParams {
    export type AsObject = {
        mode: VirtualParams.Mode,
    }

    export enum Mode {
    DISCARD = 0,
    ERROR = 1,
    }

}

export class DataSource extends jspb.Message { 
    getId(): string;
    setId(value: string): DataSource;
    getBackend(): DataSourceBackendType;
    setBackend(value: DataSourceBackendType): DataSource;
    getName(): string;
    setName(value: string): DataSource;
    getDescription(): string;
    setDescription(value: string): DataSource;

    hasPostgresqlparams(): boolean;
    clearPostgresqlparams(): void;
    getPostgresqlparams(): PostgresqlParams | undefined;
    setPostgresqlparams(value?: PostgresqlParams): DataSource;

    hasMysqlparams(): boolean;
    clearMysqlparams(): void;
    getMysqlparams(): MysqlParams | undefined;
    setMysqlparams(value?: MysqlParams): DataSource;

    hasVirtualparams(): boolean;
    clearVirtualparams(): void;
    getVirtualparams(): VirtualParams | undefined;
    setVirtualparams(value?: VirtualParams): DataSource;

    hasRedisparams(): boolean;
    clearRedisparams(): void;
    getRedisparams(): RedisParams | undefined;
    setRedisparams(value?: RedisParams): DataSource;

    hasMongoparams(): boolean;
    clearMongoparams(): void;
    getMongoparams(): MongoParams | undefined;
    setMongoparams(value?: MongoParams): DataSource;

    hasAuditdata(): boolean;
    clearAuditdata(): void;
    getAuditdata(): model_audit_pb.AuditData | undefined;
    setAuditdata(value?: model_audit_pb.AuditData): DataSource;
    getVersion(): number;
    setVersion(value: number): DataSource;

    getParamsCase(): DataSource.ParamsCase;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): DataSource.AsObject;
    static toObject(includeInstance: boolean, msg: DataSource): DataSource.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: DataSource, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): DataSource;
    static deserializeBinaryFromReader(message: DataSource, reader: jspb.BinaryReader): DataSource;
}

export namespace DataSource {
    export type AsObject = {
        id: string,
        backend: DataSourceBackendType,
        name: string,
        description: string,
        postgresqlparams?: PostgresqlParams.AsObject,
        mysqlparams?: MysqlParams.AsObject,
        virtualparams?: VirtualParams.AsObject,
        redisparams?: RedisParams.AsObject,
        mongoparams?: MongoParams.AsObject,
        auditdata?: model_audit_pb.AuditData.AsObject,
        version: number,
    }

    export enum ParamsCase {
        PARAMS_NOT_SET = 0,
        POSTGRESQLPARAMS = 7,
        MYSQLPARAMS = 8,
        VIRTUALPARAMS = 9,
        REDISPARAMS = 10,
        MONGOPARAMS = 11,
    }

}

export class DataSourceEntity extends jspb.Message { 
    getName(): string;
    setName(value: string): DataSourceEntity;
    getReadonly(): boolean;
    setReadonly(value: boolean): DataSourceEntity;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): DataSourceEntity.AsObject;
    static toObject(includeInstance: boolean, msg: DataSourceEntity): DataSourceEntity.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: DataSourceEntity, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): DataSourceEntity;
    static deserializeBinaryFromReader(message: DataSourceEntity, reader: jspb.BinaryReader): DataSourceEntity;
}

export namespace DataSourceEntity {
    export type AsObject = {
        name: string,
        readonly: boolean,
    }
}

export class DataSourceCatalog extends jspb.Message { 
    getName(): string;
    setName(value: string): DataSourceCatalog;
    clearEntitiesList(): void;
    getEntitiesList(): Array<DataSourceEntity>;
    setEntitiesList(value: Array<DataSourceEntity>): DataSourceCatalog;
    addEntities(value?: DataSourceEntity, index?: number): DataSourceEntity;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): DataSourceCatalog.AsObject;
    static toObject(includeInstance: boolean, msg: DataSourceCatalog): DataSourceCatalog.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: DataSourceCatalog, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): DataSourceCatalog;
    static deserializeBinaryFromReader(message: DataSourceCatalog, reader: jspb.BinaryReader): DataSourceCatalog;
}

export namespace DataSourceCatalog {
    export type AsObject = {
        name: string,
        entitiesList: Array<DataSourceEntity.AsObject>,
    }
}

export enum DataSourceBackendType {
    POSTGRESQL = 0,
    VIRTUAL = 1,
    MYSQL = 2,
    ORACLE = 3,
    MONGODB = 4,
    REDIS = 5,
}
