package grpc_service

import (
	"context"
	"data-handler/grpc/stub"
	"data-handler/service"
)

type dataSourceServiceServer struct {
	stub.DataSourceServiceServer
	service service.DataSourceService
}

func (d *dataSourceServiceServer) ListEntities(ctx context.Context, request *stub.ListEntitiesRequest) (*stub.ListEntitiesResponse, error) {
	res, err := d.service.ListEntities(ctx, request.Id)

	if err != nil {
		return nil, err
	}

	return &stub.ListEntitiesResponse{
		Entities: res,
		Error:    nil,
	}, nil
}

func (d *dataSourceServiceServer) List(ctx context.Context, request *stub.ListDataSourceRequest) (*stub.ListDataSourceResponse, error) {
	result, err := d.service.List(ctx)

	return &stub.ListDataSourceResponse{
		Content: result,
		Error:   toProtoError(err),
	}, nil
}

func (d *dataSourceServiceServer) Status(ctx context.Context, request *stub.StatusRequest) (*stub.StatusResponse, error) {
	connectionAlreadyInitiated, testConnection, err := d.service.GetStatus(ctx, request.Id)

	return &stub.StatusResponse{
		ConnectionAlreadyInitiated: connectionAlreadyInitiated,
		TestConnection:             testConnection,
		Error:                      toProtoError(err),
	}, nil
}

func (d *dataSourceServiceServer) Create(ctx context.Context, request *stub.CreateDataSourceRequest) (*stub.CreateDataSourceResponse, error) {
	res, err := d.service.Create(ctx, request.DataSources)

	return &stub.CreateDataSourceResponse{
		DataSources: res,
		Error:       toProtoError(err),
	}, nil
}

func (d *dataSourceServiceServer) Update(ctx context.Context, request *stub.UpdateDataSourceRequest) (*stub.UpdateDataSourceResponse, error) {
	res, err := d.service.Update(ctx, request.DataSources)

	return &stub.UpdateDataSourceResponse{
		DataSources: res,
		Error:       toProtoError(err),
	}, nil
}

func (d *dataSourceServiceServer) PrepareResourceFromEntity(ctx context.Context, request *stub.PrepareResourceFromEntityRequest) (*stub.PrepareResourceFromEntityResponse, error) {
	resources, err := d.service.PrepareResourceFromEntity(ctx, request.Id, request.Catalog, request.Entity)

	return &stub.PrepareResourceFromEntityResponse{
		Resource: resources,
		Error:    toProtoError(err),
	}, nil
}

func (d *dataSourceServiceServer) Get(ctx context.Context, request *stub.GetDataSourceRequest) (*stub.GetDataSourceResponse, error) {
	dataSource, err := d.service.Get(ctx, request.Id)

	return &stub.GetDataSourceResponse{
		DataSource: dataSource,
		Error:      toProtoError(err),
	}, nil
}

func (d *dataSourceServiceServer) Delete(ctx context.Context, request *stub.DeleteDataSourceRequest) (*stub.DeleteDataSourceResponse, error) {
	err := d.service.Delete(ctx, request.Ids)

	return &stub.DeleteDataSourceResponse{
		Error: toProtoError(err),
	}, nil
}

func NewDataSourceServiceServer(dataSourceService service.DataSourceService) stub.DataSourceServiceServer {
	return &dataSourceServiceServer{
		service: dataSourceService,
	}
}
