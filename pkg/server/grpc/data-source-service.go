package grpc

import (
	"context"
	"github.com/tislib/data-handler/pkg/abs"
	"github.com/tislib/data-handler/pkg/server/util"
	"github.com/tislib/data-handler/pkg/stub"
)

type dataSourceServiceServer struct {
	stub.DataSourceServiceServer
	service abs.DataSourceService
}

func (d *dataSourceServiceServer) ListEntities(ctx context.Context, request *stub.ListEntitiesRequest) (*stub.ListEntitiesResponse, error) {
	res, err := d.service.ListEntities(ctx, request.Id)

	return &stub.ListEntitiesResponse{
		Entities: res,
	}, util.ToStatusError(err)
}

func (d *dataSourceServiceServer) List(ctx context.Context, request *stub.ListDataSourceRequest) (*stub.ListDataSourceResponse, error) {
	result, err := d.service.List(ctx)

	return &stub.ListDataSourceResponse{
		Content: result,
	}, util.ToStatusError(err)
}

func (d *dataSourceServiceServer) Status(ctx context.Context, request *stub.StatusRequest) (*stub.StatusResponse, error) {
	connectionAlreadyInitiated, testConnection, err := d.service.GetStatus(ctx, request.Id)

	return &stub.StatusResponse{
		ConnectionAlreadyInitiated: connectionAlreadyInitiated,
		TestConnection:             testConnection,
	}, util.ToStatusError(err)
}

func (d *dataSourceServiceServer) Create(ctx context.Context, request *stub.CreateDataSourceRequest) (*stub.CreateDataSourceResponse, error) {
	res, err := d.service.Create(ctx, request.DataSources)

	return &stub.CreateDataSourceResponse{
		DataSources: res,
	}, util.ToStatusError(err)
}

func (d *dataSourceServiceServer) Update(ctx context.Context, request *stub.UpdateDataSourceRequest) (*stub.UpdateDataSourceResponse, error) {
	res, err := d.service.Update(ctx, request.DataSources)

	return &stub.UpdateDataSourceResponse{
		DataSources: res,
	}, util.ToStatusError(err)
}

func (d *dataSourceServiceServer) PrepareResourceFromEntity(ctx context.Context, request *stub.PrepareResourceFromEntityRequest) (*stub.PrepareResourceFromEntityResponse, error) {
	resources, err := d.service.PrepareResourceFromEntity(ctx, request.Id, request.Catalog, request.Entity)

	return &stub.PrepareResourceFromEntityResponse{
		Resource: resources,
	}, util.ToStatusError(err)
}

func (d *dataSourceServiceServer) Get(ctx context.Context, request *stub.GetDataSourceRequest) (*stub.GetDataSourceResponse, error) {
	dataSource, err := d.service.Get(ctx, request.Id)

	return &stub.GetDataSourceResponse{
		DataSource: dataSource,
	}, util.ToStatusError(err)
}

func (d *dataSourceServiceServer) Delete(ctx context.Context, request *stub.DeleteDataSourceRequest) (*stub.DeleteDataSourceResponse, error) {
	err := d.service.Delete(ctx, request.Ids)

	return &stub.DeleteDataSourceResponse{}, util.ToStatusError(err)
}

func NewDataSourceServiceServer(dataSourceService abs.DataSourceService) stub.DataSourceServiceServer {
	return &dataSourceServiceServer{
		service: dataSourceService,
	}
}
