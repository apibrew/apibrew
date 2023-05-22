// package: stub
// file: stub/data-source.proto

/* tslint:disable */
/* eslint-disable */

import * as grpc from "@grpc/grpc-js";
import * as stub_data_source_pb from "../stub/data-source_pb";
import * as model_data_source_pb from "../model/data-source_pb";
import * as model_error_pb from "../model/error_pb";
import * as model_resource_pb from "../model/resource_pb";


interface IDataSourceService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
    create: IDataSourceService_ICreate;
    list: IDataSourceService_IList;
    update: IDataSourceService_IUpdate;
    delete: IDataSourceService_IDelete;
    get: IDataSourceService_IGet;
    status: IDataSourceService_IStatus;
    listEntities: IDataSourceService_IListEntities;
    prepareResourceFromEntity: IDataSourceService_IPrepareResourceFromEntity;
}

interface IDataSourceService_ICreate extends grpc.MethodDefinition<stub_data_source_pb.CreateDataSourceRequest, stub_data_source_pb.CreateDataSourceResponse> {
    path: "/stub.DataSource/Create";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<stub_data_source_pb.CreateDataSourceRequest>;
    requestDeserialize: grpc.deserialize<stub_data_source_pb.CreateDataSourceRequest>;
    responseSerialize: grpc.serialize<stub_data_source_pb.CreateDataSourceResponse>;
    responseDeserialize: grpc.deserialize<stub_data_source_pb.CreateDataSourceResponse>;
}
interface IDataSourceService_IList extends grpc.MethodDefinition<stub_data_source_pb.ListDataSourceRequest, stub_data_source_pb.ListDataSourceResponse> {
    path: "/stub.DataSource/List";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<stub_data_source_pb.ListDataSourceRequest>;
    requestDeserialize: grpc.deserialize<stub_data_source_pb.ListDataSourceRequest>;
    responseSerialize: grpc.serialize<stub_data_source_pb.ListDataSourceResponse>;
    responseDeserialize: grpc.deserialize<stub_data_source_pb.ListDataSourceResponse>;
}
interface IDataSourceService_IUpdate extends grpc.MethodDefinition<stub_data_source_pb.UpdateDataSourceRequest, stub_data_source_pb.UpdateDataSourceResponse> {
    path: "/stub.DataSource/Update";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<stub_data_source_pb.UpdateDataSourceRequest>;
    requestDeserialize: grpc.deserialize<stub_data_source_pb.UpdateDataSourceRequest>;
    responseSerialize: grpc.serialize<stub_data_source_pb.UpdateDataSourceResponse>;
    responseDeserialize: grpc.deserialize<stub_data_source_pb.UpdateDataSourceResponse>;
}
interface IDataSourceService_IDelete extends grpc.MethodDefinition<stub_data_source_pb.DeleteDataSourceRequest, stub_data_source_pb.DeleteDataSourceResponse> {
    path: "/stub.DataSource/Delete";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<stub_data_source_pb.DeleteDataSourceRequest>;
    requestDeserialize: grpc.deserialize<stub_data_source_pb.DeleteDataSourceRequest>;
    responseSerialize: grpc.serialize<stub_data_source_pb.DeleteDataSourceResponse>;
    responseDeserialize: grpc.deserialize<stub_data_source_pb.DeleteDataSourceResponse>;
}
interface IDataSourceService_IGet extends grpc.MethodDefinition<stub_data_source_pb.GetDataSourceRequest, stub_data_source_pb.GetDataSourceResponse> {
    path: "/stub.DataSource/Get";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<stub_data_source_pb.GetDataSourceRequest>;
    requestDeserialize: grpc.deserialize<stub_data_source_pb.GetDataSourceRequest>;
    responseSerialize: grpc.serialize<stub_data_source_pb.GetDataSourceResponse>;
    responseDeserialize: grpc.deserialize<stub_data_source_pb.GetDataSourceResponse>;
}
interface IDataSourceService_IStatus extends grpc.MethodDefinition<stub_data_source_pb.StatusRequest, stub_data_source_pb.StatusResponse> {
    path: "/stub.DataSource/Status";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<stub_data_source_pb.StatusRequest>;
    requestDeserialize: grpc.deserialize<stub_data_source_pb.StatusRequest>;
    responseSerialize: grpc.serialize<stub_data_source_pb.StatusResponse>;
    responseDeserialize: grpc.deserialize<stub_data_source_pb.StatusResponse>;
}
interface IDataSourceService_IListEntities extends grpc.MethodDefinition<stub_data_source_pb.ListEntitiesRequest, stub_data_source_pb.ListEntitiesResponse> {
    path: "/stub.DataSource/ListEntities";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<stub_data_source_pb.ListEntitiesRequest>;
    requestDeserialize: grpc.deserialize<stub_data_source_pb.ListEntitiesRequest>;
    responseSerialize: grpc.serialize<stub_data_source_pb.ListEntitiesResponse>;
    responseDeserialize: grpc.deserialize<stub_data_source_pb.ListEntitiesResponse>;
}
interface IDataSourceService_IPrepareResourceFromEntity extends grpc.MethodDefinition<stub_data_source_pb.PrepareResourceFromEntityRequest, stub_data_source_pb.PrepareResourceFromEntityResponse> {
    path: "/stub.DataSource/PrepareResourceFromEntity";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<stub_data_source_pb.PrepareResourceFromEntityRequest>;
    requestDeserialize: grpc.deserialize<stub_data_source_pb.PrepareResourceFromEntityRequest>;
    responseSerialize: grpc.serialize<stub_data_source_pb.PrepareResourceFromEntityResponse>;
    responseDeserialize: grpc.deserialize<stub_data_source_pb.PrepareResourceFromEntityResponse>;
}

export const DataSourceService: IDataSourceService;

export interface IDataSourceServer {
    create: grpc.handleUnaryCall<stub_data_source_pb.CreateDataSourceRequest, stub_data_source_pb.CreateDataSourceResponse>;
    list: grpc.handleUnaryCall<stub_data_source_pb.ListDataSourceRequest, stub_data_source_pb.ListDataSourceResponse>;
    update: grpc.handleUnaryCall<stub_data_source_pb.UpdateDataSourceRequest, stub_data_source_pb.UpdateDataSourceResponse>;
    delete: grpc.handleUnaryCall<stub_data_source_pb.DeleteDataSourceRequest, stub_data_source_pb.DeleteDataSourceResponse>;
    get: grpc.handleUnaryCall<stub_data_source_pb.GetDataSourceRequest, stub_data_source_pb.GetDataSourceResponse>;
    status: grpc.handleUnaryCall<stub_data_source_pb.StatusRequest, stub_data_source_pb.StatusResponse>;
    listEntities: grpc.handleUnaryCall<stub_data_source_pb.ListEntitiesRequest, stub_data_source_pb.ListEntitiesResponse>;
    prepareResourceFromEntity: grpc.handleUnaryCall<stub_data_source_pb.PrepareResourceFromEntityRequest, stub_data_source_pb.PrepareResourceFromEntityResponse>;
}

export interface IDataSourceClient {
    create(request: stub_data_source_pb.CreateDataSourceRequest, callback: (error: grpc.ServiceError | null, response: stub_data_source_pb.CreateDataSourceResponse) => void): grpc.ClientUnaryCall;
    create(request: stub_data_source_pb.CreateDataSourceRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: stub_data_source_pb.CreateDataSourceResponse) => void): grpc.ClientUnaryCall;
    create(request: stub_data_source_pb.CreateDataSourceRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: stub_data_source_pb.CreateDataSourceResponse) => void): grpc.ClientUnaryCall;
    list(request: stub_data_source_pb.ListDataSourceRequest, callback: (error: grpc.ServiceError | null, response: stub_data_source_pb.ListDataSourceResponse) => void): grpc.ClientUnaryCall;
    list(request: stub_data_source_pb.ListDataSourceRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: stub_data_source_pb.ListDataSourceResponse) => void): grpc.ClientUnaryCall;
    list(request: stub_data_source_pb.ListDataSourceRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: stub_data_source_pb.ListDataSourceResponse) => void): grpc.ClientUnaryCall;
    update(request: stub_data_source_pb.UpdateDataSourceRequest, callback: (error: grpc.ServiceError | null, response: stub_data_source_pb.UpdateDataSourceResponse) => void): grpc.ClientUnaryCall;
    update(request: stub_data_source_pb.UpdateDataSourceRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: stub_data_source_pb.UpdateDataSourceResponse) => void): grpc.ClientUnaryCall;
    update(request: stub_data_source_pb.UpdateDataSourceRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: stub_data_source_pb.UpdateDataSourceResponse) => void): grpc.ClientUnaryCall;
    delete(request: stub_data_source_pb.DeleteDataSourceRequest, callback: (error: grpc.ServiceError | null, response: stub_data_source_pb.DeleteDataSourceResponse) => void): grpc.ClientUnaryCall;
    delete(request: stub_data_source_pb.DeleteDataSourceRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: stub_data_source_pb.DeleteDataSourceResponse) => void): grpc.ClientUnaryCall;
    delete(request: stub_data_source_pb.DeleteDataSourceRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: stub_data_source_pb.DeleteDataSourceResponse) => void): grpc.ClientUnaryCall;
    get(request: stub_data_source_pb.GetDataSourceRequest, callback: (error: grpc.ServiceError | null, response: stub_data_source_pb.GetDataSourceResponse) => void): grpc.ClientUnaryCall;
    get(request: stub_data_source_pb.GetDataSourceRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: stub_data_source_pb.GetDataSourceResponse) => void): grpc.ClientUnaryCall;
    get(request: stub_data_source_pb.GetDataSourceRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: stub_data_source_pb.GetDataSourceResponse) => void): grpc.ClientUnaryCall;
    status(request: stub_data_source_pb.StatusRequest, callback: (error: grpc.ServiceError | null, response: stub_data_source_pb.StatusResponse) => void): grpc.ClientUnaryCall;
    status(request: stub_data_source_pb.StatusRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: stub_data_source_pb.StatusResponse) => void): grpc.ClientUnaryCall;
    status(request: stub_data_source_pb.StatusRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: stub_data_source_pb.StatusResponse) => void): grpc.ClientUnaryCall;
    listEntities(request: stub_data_source_pb.ListEntitiesRequest, callback: (error: grpc.ServiceError | null, response: stub_data_source_pb.ListEntitiesResponse) => void): grpc.ClientUnaryCall;
    listEntities(request: stub_data_source_pb.ListEntitiesRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: stub_data_source_pb.ListEntitiesResponse) => void): grpc.ClientUnaryCall;
    listEntities(request: stub_data_source_pb.ListEntitiesRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: stub_data_source_pb.ListEntitiesResponse) => void): grpc.ClientUnaryCall;
    prepareResourceFromEntity(request: stub_data_source_pb.PrepareResourceFromEntityRequest, callback: (error: grpc.ServiceError | null, response: stub_data_source_pb.PrepareResourceFromEntityResponse) => void): grpc.ClientUnaryCall;
    prepareResourceFromEntity(request: stub_data_source_pb.PrepareResourceFromEntityRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: stub_data_source_pb.PrepareResourceFromEntityResponse) => void): grpc.ClientUnaryCall;
    prepareResourceFromEntity(request: stub_data_source_pb.PrepareResourceFromEntityRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: stub_data_source_pb.PrepareResourceFromEntityResponse) => void): grpc.ClientUnaryCall;
}

export class DataSourceClient extends grpc.Client implements IDataSourceClient {
    constructor(address: string, credentials: grpc.ChannelCredentials, options?: object);
    public create(request: stub_data_source_pb.CreateDataSourceRequest, callback: (error: grpc.ServiceError | null, response: stub_data_source_pb.CreateDataSourceResponse) => void): grpc.ClientUnaryCall;
    public create(request: stub_data_source_pb.CreateDataSourceRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: stub_data_source_pb.CreateDataSourceResponse) => void): grpc.ClientUnaryCall;
    public create(request: stub_data_source_pb.CreateDataSourceRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: stub_data_source_pb.CreateDataSourceResponse) => void): grpc.ClientUnaryCall;
    public list(request: stub_data_source_pb.ListDataSourceRequest, callback: (error: grpc.ServiceError | null, response: stub_data_source_pb.ListDataSourceResponse) => void): grpc.ClientUnaryCall;
    public list(request: stub_data_source_pb.ListDataSourceRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: stub_data_source_pb.ListDataSourceResponse) => void): grpc.ClientUnaryCall;
    public list(request: stub_data_source_pb.ListDataSourceRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: stub_data_source_pb.ListDataSourceResponse) => void): grpc.ClientUnaryCall;
    public update(request: stub_data_source_pb.UpdateDataSourceRequest, callback: (error: grpc.ServiceError | null, response: stub_data_source_pb.UpdateDataSourceResponse) => void): grpc.ClientUnaryCall;
    public update(request: stub_data_source_pb.UpdateDataSourceRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: stub_data_source_pb.UpdateDataSourceResponse) => void): grpc.ClientUnaryCall;
    public update(request: stub_data_source_pb.UpdateDataSourceRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: stub_data_source_pb.UpdateDataSourceResponse) => void): grpc.ClientUnaryCall;
    public delete(request: stub_data_source_pb.DeleteDataSourceRequest, callback: (error: grpc.ServiceError | null, response: stub_data_source_pb.DeleteDataSourceResponse) => void): grpc.ClientUnaryCall;
    public delete(request: stub_data_source_pb.DeleteDataSourceRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: stub_data_source_pb.DeleteDataSourceResponse) => void): grpc.ClientUnaryCall;
    public delete(request: stub_data_source_pb.DeleteDataSourceRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: stub_data_source_pb.DeleteDataSourceResponse) => void): grpc.ClientUnaryCall;
    public get(request: stub_data_source_pb.GetDataSourceRequest, callback: (error: grpc.ServiceError | null, response: stub_data_source_pb.GetDataSourceResponse) => void): grpc.ClientUnaryCall;
    public get(request: stub_data_source_pb.GetDataSourceRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: stub_data_source_pb.GetDataSourceResponse) => void): grpc.ClientUnaryCall;
    public get(request: stub_data_source_pb.GetDataSourceRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: stub_data_source_pb.GetDataSourceResponse) => void): grpc.ClientUnaryCall;
    public status(request: stub_data_source_pb.StatusRequest, callback: (error: grpc.ServiceError | null, response: stub_data_source_pb.StatusResponse) => void): grpc.ClientUnaryCall;
    public status(request: stub_data_source_pb.StatusRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: stub_data_source_pb.StatusResponse) => void): grpc.ClientUnaryCall;
    public status(request: stub_data_source_pb.StatusRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: stub_data_source_pb.StatusResponse) => void): grpc.ClientUnaryCall;
    public listEntities(request: stub_data_source_pb.ListEntitiesRequest, callback: (error: grpc.ServiceError | null, response: stub_data_source_pb.ListEntitiesResponse) => void): grpc.ClientUnaryCall;
    public listEntities(request: stub_data_source_pb.ListEntitiesRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: stub_data_source_pb.ListEntitiesResponse) => void): grpc.ClientUnaryCall;
    public listEntities(request: stub_data_source_pb.ListEntitiesRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: stub_data_source_pb.ListEntitiesResponse) => void): grpc.ClientUnaryCall;
    public prepareResourceFromEntity(request: stub_data_source_pb.PrepareResourceFromEntityRequest, callback: (error: grpc.ServiceError | null, response: stub_data_source_pb.PrepareResourceFromEntityResponse) => void): grpc.ClientUnaryCall;
    public prepareResourceFromEntity(request: stub_data_source_pb.PrepareResourceFromEntityRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: stub_data_source_pb.PrepareResourceFromEntityResponse) => void): grpc.ClientUnaryCall;
    public prepareResourceFromEntity(request: stub_data_source_pb.PrepareResourceFromEntityRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: stub_data_source_pb.PrepareResourceFromEntityResponse) => void): grpc.ClientUnaryCall;
}
