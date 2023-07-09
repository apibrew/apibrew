package grpc

import (
	"context"
	"github.com/apibrew/apibrew/pkg/abs"
	"github.com/apibrew/apibrew/pkg/stub"
	"github.com/apibrew/apibrew/pkg/util"
)

type dataSourceServer struct {
	stub.DataSourceServer
	service abs.DataSourceService
}

func (d *dataSourceServer) ListEntities(ctx context.Context, request *stub.ListEntitiesRequest) (*stub.ListEntitiesResponse, error) {
	res, err := d.service.ListEntities(ctx, request.Id)

	return &stub.ListEntitiesResponse{
		Catalogs: res,
	}, util.ToStatusError(err)
}

func (d *dataSourceServer) Status(ctx context.Context, request *stub.StatusRequest) (*stub.StatusResponse, error) {
	connectionAlreadyInitiated, testConnection, err := d.service.GetStatus(ctx, request.Id)

	return &stub.StatusResponse{
		ConnectionAlreadyInitiated: connectionAlreadyInitiated,
		TestConnection:             testConnection,
	}, util.ToStatusError(err)
}

func (d *dataSourceServer) PrepareResourceFromEntity(ctx context.Context, request *stub.PrepareResourceFromEntityRequest) (*stub.PrepareResourceFromEntityResponse, error) {
	resources, err := d.service.PrepareResourceFromEntity(ctx, request.Id, request.Catalog, request.Entity)

	return &stub.PrepareResourceFromEntityResponse{
		Resource: resources,
	}, util.ToStatusError(err)
}

func NewDataSourceServer(dataSourceService abs.DataSourceService) stub.DataSourceServer {
	return &dataSourceServer{
		service: dataSourceService,
	}
}
