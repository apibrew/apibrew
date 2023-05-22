// package: stub
// file: stub/resource.proto

/* tslint:disable */
/* eslint-disable */

import * as grpc from "@grpc/grpc-js";
import * as stub_resource_pb from "../stub/resource_pb";
import * as model_error_pb from "../model/error_pb";
import * as model_resource_pb from "../model/resource_pb";
import * as model_resource_migration_pb from "../model/resource-migration_pb";


interface IResourceService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
    create: IResourceService_ICreate;
    update: IResourceService_IUpdate;
    delete: IResourceService_IDelete;
    list: IResourceService_IList;
    prepareResourceMigrationPlan: IResourceService_IPrepareResourceMigrationPlan;
    get: IResourceService_IGet;
    getByName: IResourceService_IGetByName;
    getSystemResource: IResourceService_IGetSystemResource;
}

interface IResourceService_ICreate extends grpc.MethodDefinition<stub_resource_pb.CreateResourceRequest, stub_resource_pb.CreateResourceResponse> {
    path: "/stub.Resource/Create";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<stub_resource_pb.CreateResourceRequest>;
    requestDeserialize: grpc.deserialize<stub_resource_pb.CreateResourceRequest>;
    responseSerialize: grpc.serialize<stub_resource_pb.CreateResourceResponse>;
    responseDeserialize: grpc.deserialize<stub_resource_pb.CreateResourceResponse>;
}
interface IResourceService_IUpdate extends grpc.MethodDefinition<stub_resource_pb.UpdateResourceRequest, stub_resource_pb.UpdateResourceResponse> {
    path: "/stub.Resource/Update";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<stub_resource_pb.UpdateResourceRequest>;
    requestDeserialize: grpc.deserialize<stub_resource_pb.UpdateResourceRequest>;
    responseSerialize: grpc.serialize<stub_resource_pb.UpdateResourceResponse>;
    responseDeserialize: grpc.deserialize<stub_resource_pb.UpdateResourceResponse>;
}
interface IResourceService_IDelete extends grpc.MethodDefinition<stub_resource_pb.DeleteResourceRequest, stub_resource_pb.DeleteResourceResponse> {
    path: "/stub.Resource/Delete";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<stub_resource_pb.DeleteResourceRequest>;
    requestDeserialize: grpc.deserialize<stub_resource_pb.DeleteResourceRequest>;
    responseSerialize: grpc.serialize<stub_resource_pb.DeleteResourceResponse>;
    responseDeserialize: grpc.deserialize<stub_resource_pb.DeleteResourceResponse>;
}
interface IResourceService_IList extends grpc.MethodDefinition<stub_resource_pb.ListResourceRequest, stub_resource_pb.ListResourceResponse> {
    path: "/stub.Resource/List";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<stub_resource_pb.ListResourceRequest>;
    requestDeserialize: grpc.deserialize<stub_resource_pb.ListResourceRequest>;
    responseSerialize: grpc.serialize<stub_resource_pb.ListResourceResponse>;
    responseDeserialize: grpc.deserialize<stub_resource_pb.ListResourceResponse>;
}
interface IResourceService_IPrepareResourceMigrationPlan extends grpc.MethodDefinition<stub_resource_pb.PrepareResourceMigrationPlanRequest, stub_resource_pb.PrepareResourceMigrationPlanResponse> {
    path: "/stub.Resource/PrepareResourceMigrationPlan";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<stub_resource_pb.PrepareResourceMigrationPlanRequest>;
    requestDeserialize: grpc.deserialize<stub_resource_pb.PrepareResourceMigrationPlanRequest>;
    responseSerialize: grpc.serialize<stub_resource_pb.PrepareResourceMigrationPlanResponse>;
    responseDeserialize: grpc.deserialize<stub_resource_pb.PrepareResourceMigrationPlanResponse>;
}
interface IResourceService_IGet extends grpc.MethodDefinition<stub_resource_pb.GetResourceRequest, stub_resource_pb.GetResourceResponse> {
    path: "/stub.Resource/Get";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<stub_resource_pb.GetResourceRequest>;
    requestDeserialize: grpc.deserialize<stub_resource_pb.GetResourceRequest>;
    responseSerialize: grpc.serialize<stub_resource_pb.GetResourceResponse>;
    responseDeserialize: grpc.deserialize<stub_resource_pb.GetResourceResponse>;
}
interface IResourceService_IGetByName extends grpc.MethodDefinition<stub_resource_pb.GetResourceByNameRequest, stub_resource_pb.GetResourceByNameResponse> {
    path: "/stub.Resource/GetByName";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<stub_resource_pb.GetResourceByNameRequest>;
    requestDeserialize: grpc.deserialize<stub_resource_pb.GetResourceByNameRequest>;
    responseSerialize: grpc.serialize<stub_resource_pb.GetResourceByNameResponse>;
    responseDeserialize: grpc.deserialize<stub_resource_pb.GetResourceByNameResponse>;
}
interface IResourceService_IGetSystemResource extends grpc.MethodDefinition<stub_resource_pb.GetSystemResourceRequest, stub_resource_pb.GetSystemResourceResponse> {
    path: "/stub.Resource/GetSystemResource";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<stub_resource_pb.GetSystemResourceRequest>;
    requestDeserialize: grpc.deserialize<stub_resource_pb.GetSystemResourceRequest>;
    responseSerialize: grpc.serialize<stub_resource_pb.GetSystemResourceResponse>;
    responseDeserialize: grpc.deserialize<stub_resource_pb.GetSystemResourceResponse>;
}

export const ResourceService: IResourceService;

export interface IResourceServer {
    create: grpc.handleUnaryCall<stub_resource_pb.CreateResourceRequest, stub_resource_pb.CreateResourceResponse>;
    update: grpc.handleUnaryCall<stub_resource_pb.UpdateResourceRequest, stub_resource_pb.UpdateResourceResponse>;
    delete: grpc.handleUnaryCall<stub_resource_pb.DeleteResourceRequest, stub_resource_pb.DeleteResourceResponse>;
    list: grpc.handleUnaryCall<stub_resource_pb.ListResourceRequest, stub_resource_pb.ListResourceResponse>;
    prepareResourceMigrationPlan: grpc.handleUnaryCall<stub_resource_pb.PrepareResourceMigrationPlanRequest, stub_resource_pb.PrepareResourceMigrationPlanResponse>;
    get: grpc.handleUnaryCall<stub_resource_pb.GetResourceRequest, stub_resource_pb.GetResourceResponse>;
    getByName: grpc.handleUnaryCall<stub_resource_pb.GetResourceByNameRequest, stub_resource_pb.GetResourceByNameResponse>;
    getSystemResource: grpc.handleUnaryCall<stub_resource_pb.GetSystemResourceRequest, stub_resource_pb.GetSystemResourceResponse>;
}

export interface IResourceClient {
    create(request: stub_resource_pb.CreateResourceRequest, callback: (error: grpc.ServiceError | null, response: stub_resource_pb.CreateResourceResponse) => void): grpc.ClientUnaryCall;
    create(request: stub_resource_pb.CreateResourceRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: stub_resource_pb.CreateResourceResponse) => void): grpc.ClientUnaryCall;
    create(request: stub_resource_pb.CreateResourceRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: stub_resource_pb.CreateResourceResponse) => void): grpc.ClientUnaryCall;
    update(request: stub_resource_pb.UpdateResourceRequest, callback: (error: grpc.ServiceError | null, response: stub_resource_pb.UpdateResourceResponse) => void): grpc.ClientUnaryCall;
    update(request: stub_resource_pb.UpdateResourceRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: stub_resource_pb.UpdateResourceResponse) => void): grpc.ClientUnaryCall;
    update(request: stub_resource_pb.UpdateResourceRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: stub_resource_pb.UpdateResourceResponse) => void): grpc.ClientUnaryCall;
    delete(request: stub_resource_pb.DeleteResourceRequest, callback: (error: grpc.ServiceError | null, response: stub_resource_pb.DeleteResourceResponse) => void): grpc.ClientUnaryCall;
    delete(request: stub_resource_pb.DeleteResourceRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: stub_resource_pb.DeleteResourceResponse) => void): grpc.ClientUnaryCall;
    delete(request: stub_resource_pb.DeleteResourceRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: stub_resource_pb.DeleteResourceResponse) => void): grpc.ClientUnaryCall;
    list(request: stub_resource_pb.ListResourceRequest, callback: (error: grpc.ServiceError | null, response: stub_resource_pb.ListResourceResponse) => void): grpc.ClientUnaryCall;
    list(request: stub_resource_pb.ListResourceRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: stub_resource_pb.ListResourceResponse) => void): grpc.ClientUnaryCall;
    list(request: stub_resource_pb.ListResourceRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: stub_resource_pb.ListResourceResponse) => void): grpc.ClientUnaryCall;
    prepareResourceMigrationPlan(request: stub_resource_pb.PrepareResourceMigrationPlanRequest, callback: (error: grpc.ServiceError | null, response: stub_resource_pb.PrepareResourceMigrationPlanResponse) => void): grpc.ClientUnaryCall;
    prepareResourceMigrationPlan(request: stub_resource_pb.PrepareResourceMigrationPlanRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: stub_resource_pb.PrepareResourceMigrationPlanResponse) => void): grpc.ClientUnaryCall;
    prepareResourceMigrationPlan(request: stub_resource_pb.PrepareResourceMigrationPlanRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: stub_resource_pb.PrepareResourceMigrationPlanResponse) => void): grpc.ClientUnaryCall;
    get(request: stub_resource_pb.GetResourceRequest, callback: (error: grpc.ServiceError | null, response: stub_resource_pb.GetResourceResponse) => void): grpc.ClientUnaryCall;
    get(request: stub_resource_pb.GetResourceRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: stub_resource_pb.GetResourceResponse) => void): grpc.ClientUnaryCall;
    get(request: stub_resource_pb.GetResourceRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: stub_resource_pb.GetResourceResponse) => void): grpc.ClientUnaryCall;
    getByName(request: stub_resource_pb.GetResourceByNameRequest, callback: (error: grpc.ServiceError | null, response: stub_resource_pb.GetResourceByNameResponse) => void): grpc.ClientUnaryCall;
    getByName(request: stub_resource_pb.GetResourceByNameRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: stub_resource_pb.GetResourceByNameResponse) => void): grpc.ClientUnaryCall;
    getByName(request: stub_resource_pb.GetResourceByNameRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: stub_resource_pb.GetResourceByNameResponse) => void): grpc.ClientUnaryCall;
    getSystemResource(request: stub_resource_pb.GetSystemResourceRequest, callback: (error: grpc.ServiceError | null, response: stub_resource_pb.GetSystemResourceResponse) => void): grpc.ClientUnaryCall;
    getSystemResource(request: stub_resource_pb.GetSystemResourceRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: stub_resource_pb.GetSystemResourceResponse) => void): grpc.ClientUnaryCall;
    getSystemResource(request: stub_resource_pb.GetSystemResourceRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: stub_resource_pb.GetSystemResourceResponse) => void): grpc.ClientUnaryCall;
}

export class ResourceClient extends grpc.Client implements IResourceClient {
    constructor(address: string, credentials: grpc.ChannelCredentials, options?: object);
    public create(request: stub_resource_pb.CreateResourceRequest, callback: (error: grpc.ServiceError | null, response: stub_resource_pb.CreateResourceResponse) => void): grpc.ClientUnaryCall;
    public create(request: stub_resource_pb.CreateResourceRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: stub_resource_pb.CreateResourceResponse) => void): grpc.ClientUnaryCall;
    public create(request: stub_resource_pb.CreateResourceRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: stub_resource_pb.CreateResourceResponse) => void): grpc.ClientUnaryCall;
    public update(request: stub_resource_pb.UpdateResourceRequest, callback: (error: grpc.ServiceError | null, response: stub_resource_pb.UpdateResourceResponse) => void): grpc.ClientUnaryCall;
    public update(request: stub_resource_pb.UpdateResourceRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: stub_resource_pb.UpdateResourceResponse) => void): grpc.ClientUnaryCall;
    public update(request: stub_resource_pb.UpdateResourceRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: stub_resource_pb.UpdateResourceResponse) => void): grpc.ClientUnaryCall;
    public delete(request: stub_resource_pb.DeleteResourceRequest, callback: (error: grpc.ServiceError | null, response: stub_resource_pb.DeleteResourceResponse) => void): grpc.ClientUnaryCall;
    public delete(request: stub_resource_pb.DeleteResourceRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: stub_resource_pb.DeleteResourceResponse) => void): grpc.ClientUnaryCall;
    public delete(request: stub_resource_pb.DeleteResourceRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: stub_resource_pb.DeleteResourceResponse) => void): grpc.ClientUnaryCall;
    public list(request: stub_resource_pb.ListResourceRequest, callback: (error: grpc.ServiceError | null, response: stub_resource_pb.ListResourceResponse) => void): grpc.ClientUnaryCall;
    public list(request: stub_resource_pb.ListResourceRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: stub_resource_pb.ListResourceResponse) => void): grpc.ClientUnaryCall;
    public list(request: stub_resource_pb.ListResourceRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: stub_resource_pb.ListResourceResponse) => void): grpc.ClientUnaryCall;
    public prepareResourceMigrationPlan(request: stub_resource_pb.PrepareResourceMigrationPlanRequest, callback: (error: grpc.ServiceError | null, response: stub_resource_pb.PrepareResourceMigrationPlanResponse) => void): grpc.ClientUnaryCall;
    public prepareResourceMigrationPlan(request: stub_resource_pb.PrepareResourceMigrationPlanRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: stub_resource_pb.PrepareResourceMigrationPlanResponse) => void): grpc.ClientUnaryCall;
    public prepareResourceMigrationPlan(request: stub_resource_pb.PrepareResourceMigrationPlanRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: stub_resource_pb.PrepareResourceMigrationPlanResponse) => void): grpc.ClientUnaryCall;
    public get(request: stub_resource_pb.GetResourceRequest, callback: (error: grpc.ServiceError | null, response: stub_resource_pb.GetResourceResponse) => void): grpc.ClientUnaryCall;
    public get(request: stub_resource_pb.GetResourceRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: stub_resource_pb.GetResourceResponse) => void): grpc.ClientUnaryCall;
    public get(request: stub_resource_pb.GetResourceRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: stub_resource_pb.GetResourceResponse) => void): grpc.ClientUnaryCall;
    public getByName(request: stub_resource_pb.GetResourceByNameRequest, callback: (error: grpc.ServiceError | null, response: stub_resource_pb.GetResourceByNameResponse) => void): grpc.ClientUnaryCall;
    public getByName(request: stub_resource_pb.GetResourceByNameRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: stub_resource_pb.GetResourceByNameResponse) => void): grpc.ClientUnaryCall;
    public getByName(request: stub_resource_pb.GetResourceByNameRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: stub_resource_pb.GetResourceByNameResponse) => void): grpc.ClientUnaryCall;
    public getSystemResource(request: stub_resource_pb.GetSystemResourceRequest, callback: (error: grpc.ServiceError | null, response: stub_resource_pb.GetSystemResourceResponse) => void): grpc.ClientUnaryCall;
    public getSystemResource(request: stub_resource_pb.GetSystemResourceRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: stub_resource_pb.GetSystemResourceResponse) => void): grpc.ClientUnaryCall;
    public getSystemResource(request: stub_resource_pb.GetSystemResourceRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: stub_resource_pb.GetSystemResourceResponse) => void): grpc.ClientUnaryCall;
}
