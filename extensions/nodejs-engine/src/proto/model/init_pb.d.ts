// package: model
// file: model/init.proto

/* tslint:disable */
/* eslint-disable */

import * as jspb from "google-protobuf";
import * as model_data_source_pb from "../model/data-source_pb";
import * as model_user_pb from "../model/user_pb";
import * as model_resource_pb from "../model/resource_pb";
import * as model_record_pb from "../model/record_pb";
import * as model_namespace_pb from "../model/namespace_pb";

export class AppConfig extends jspb.Message { 
    getHost(): string;
    setHost(value: string): AppConfig;
    getPort(): number;
    setPort(value: number): AppConfig;
    getJwtprivatekey(): string;
    setJwtprivatekey(value: string): AppConfig;
    getJwtpublickey(): string;
    setJwtpublickey(value: string): AppConfig;
    getDisableauthentication(): boolean;
    setDisableauthentication(value: boolean): AppConfig;
    getDisablecache(): boolean;
    setDisablecache(value: boolean): AppConfig;
    getPluginspath(): string;
    setPluginspath(value: string): AppConfig;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): AppConfig.AsObject;
    static toObject(includeInstance: boolean, msg: AppConfig): AppConfig.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: AppConfig, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): AppConfig;
    static deserializeBinaryFromReader(message: AppConfig, reader: jspb.BinaryReader): AppConfig;
}

export namespace AppConfig {
    export type AsObject = {
        host: string,
        port: number,
        jwtprivatekey: string,
        jwtpublickey: string,
        disableauthentication: boolean,
        disablecache: boolean,
        pluginspath: string,
    }
}

export class InitData extends jspb.Message { 

    hasConfig(): boolean;
    clearConfig(): void;
    getConfig(): AppConfig | undefined;
    setConfig(value?: AppConfig): InitData;

    hasSystemdatasource(): boolean;
    clearSystemdatasource(): void;
    getSystemdatasource(): model_data_source_pb.DataSource | undefined;
    setSystemdatasource(value?: model_data_source_pb.DataSource): InitData;

    hasSystemnamespace(): boolean;
    clearSystemnamespace(): void;
    getSystemnamespace(): model_namespace_pb.Namespace | undefined;
    setSystemnamespace(value?: model_namespace_pb.Namespace): InitData;
    clearInitdatasourcesList(): void;
    getInitdatasourcesList(): Array<model_data_source_pb.DataSource>;
    setInitdatasourcesList(value: Array<model_data_source_pb.DataSource>): InitData;
    addInitdatasources(value?: model_data_source_pb.DataSource, index?: number): model_data_source_pb.DataSource;
    clearInitnamespacesList(): void;
    getInitnamespacesList(): Array<model_namespace_pb.Namespace>;
    setInitnamespacesList(value: Array<model_namespace_pb.Namespace>): InitData;
    addInitnamespaces(value?: model_namespace_pb.Namespace, index?: number): model_namespace_pb.Namespace;
    clearInitusersList(): void;
    getInitusersList(): Array<model_user_pb.User>;
    setInitusersList(value: Array<model_user_pb.User>): InitData;
    addInitusers(value?: model_user_pb.User, index?: number): model_user_pb.User;
    clearInitresourcesList(): void;
    getInitresourcesList(): Array<model_resource_pb.Resource>;
    setInitresourcesList(value: Array<model_resource_pb.Resource>): InitData;
    addInitresources(value?: model_resource_pb.Resource, index?: number): model_resource_pb.Resource;
    clearInitrecordsList(): void;
    getInitrecordsList(): Array<model_record_pb.Record>;
    setInitrecordsList(value: Array<model_record_pb.Record>): InitData;
    addInitrecords(value?: model_record_pb.Record, index?: number): model_record_pb.Record;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): InitData.AsObject;
    static toObject(includeInstance: boolean, msg: InitData): InitData.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: InitData, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): InitData;
    static deserializeBinaryFromReader(message: InitData, reader: jspb.BinaryReader): InitData;
}

export namespace InitData {
    export type AsObject = {
        config?: AppConfig.AsObject,
        systemdatasource?: model_data_source_pb.DataSource.AsObject,
        systemnamespace?: model_namespace_pb.Namespace.AsObject,
        initdatasourcesList: Array<model_data_source_pb.DataSource.AsObject>,
        initnamespacesList: Array<model_namespace_pb.Namespace.AsObject>,
        initusersList: Array<model_user_pb.User.AsObject>,
        initresourcesList: Array<model_resource_pb.Resource.AsObject>,
        initrecordsList: Array<model_record_pb.Record.AsObject>,
    }
}
