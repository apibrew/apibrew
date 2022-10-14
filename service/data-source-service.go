package service

import (
	"context"
	"data-handler/service/backend"
	"data-handler/service/backend/postgres"
	"data-handler/stub"
	model "data-handler/stub/model"
)

type DataSourceService interface {
	stub.DataSourceServiceServer
	InjectResourceService(service ResourceService)
	InjectInitData(data *model.InitData)
	Init()
	GetDataSourceBackend(dataSource *model.DataSource) backend.DataSourceBackend
	InjectRecordService(service RecordService)
	GetSystemDataSourceBackend() backend.DataSourceBackend
	InjectPostgresResourceServiceBackend(serviceBackend backend.ResourceServiceBackend)
	GetDataSourceBackendById(dataSourceId string) (backend.DataSourceBackend, error)
}

type dataSourceService struct {
	stub.DataSourceServiceServer
	resourceService                ResourceService
	recordService                  RecordService
	systemDataSource               *model.DataSource
	postgresResourceServiceBackend backend.ResourceServiceBackend
}

/*
  rpc TestConnection(TestConnectionRequest) returns (TestConnectionResponse) {}
  rpc ListEntities(TestConnectionRequest) returns (TestConnectionResponse) {}
  rpc PrepareResourceFromEntity(PrepareResourceFromEntityRequest) returns (PrepareResourceFromEntityResponse) {}
*/

func (d *dataSourceService) List(ctx context.Context, request *stub.ListDataSourceRequest) (*stub.ListDataSourceResponse, error) {
	systemCtx := withSystemContext(ctx)
	result, err := d.recordService.List(systemCtx, &stub.ListRecordRequest{
		Resource: dataSourceResource.Name,
		Token:    request.Token,
	})

	if err != nil {
		return nil, err
	}

	return &stub.ListDataSourceResponse{
		Content: mapFromRecord(result.Content, dataSourceFromRecord),
		Error:   result.Error,
	}, err
}

func (d *dataSourceService) Status(ctx context.Context, request *stub.StatusRequest) (*stub.StatusResponse, error) {
	return d.postgresResourceServiceBackend.GetStatus(request.Id)
}

func (d *dataSourceService) Create(ctx context.Context, request *stub.CreateDataSourceRequest) (*stub.CreateDataSourceResponse, error) {
	// insert records via resource service
	records := mapToRecord(request.DataSources, dataSourceToRecord)
	systemCtx := withSystemContext(ctx)
	result, err := d.recordService.Create(systemCtx, &stub.CreateRecordRequest{
		Token:   request.Token,
		Records: records,
	})

	if err != nil {
		return nil, err
	}

	return &stub.CreateDataSourceResponse{
		DataSources: mapFromRecord(result.Records, dataSourceFromRecord),
		Error:       result.Error,
	}, err
}

func (d *dataSourceService) Update(ctx context.Context, request *stub.UpdateDataSourceRequest) (*stub.UpdateDataSourceResponse, error) {
	// insert records via resource service
	records := mapToRecord(request.DataSources, dataSourceToRecord)
	systemCtx := withSystemContext(ctx)
	result, err := d.recordService.Update(systemCtx, &stub.UpdateRecordRequest{
		Token:   request.Token,
		Records: records,
	})

	if err != nil {
		return nil, err
	}

	for _, item := range request.DataSources {
		d.postgresResourceServiceBackend.DestroyDataSource(item.Id)
	}

	return &stub.UpdateDataSourceResponse{
		DataSources: mapFromRecord(result.Records, dataSourceFromRecord),
		Error:       result.Error,
	}, err
}

func (d *dataSourceService) Get(ctx context.Context, request *stub.GetDataSourceRequest) (*stub.GetDataSourceResponse, error) {
	systemCtx := withSystemContext(ctx)
	record, err := d.recordService.Get(systemCtx, &stub.GetRecordRequest{
		Token:    request.Token,
		Resource: dataSourceResource.Name,
		Id:       request.Id,
	})

	if err != nil {
		return nil, err
	}

	return &stub.GetDataSourceResponse{
		DataSource: dataSourceFromRecord(record.Record),
		Error:      record.Error,
	}, nil
}

func (d *dataSourceService) Delete(ctx context.Context, request *stub.DeleteDataSourceRequest) (*stub.DeleteDataSourceResponse, error) {
	systemCtx := withSystemContext(ctx)

	record, err := d.recordService.Delete(systemCtx, &stub.DeleteRecordRequest{
		Token:    request.Token,
		Resource: dataSourceResource.Name,
		Ids:      request.Ids,
	})

	if err != nil {
		return nil, err
	}

	for _, dataSourceId := range request.Ids {
		d.postgresResourceServiceBackend.DestroyDataSource(dataSourceId)
	}

	return &stub.DeleteDataSourceResponse{
		Error: record.Error,
	}, nil
}

func (d *dataSourceService) GetDataSourceBackendById(dataSourceId string) (backend.DataSourceBackend, error) {
	if dataSourceId == d.systemDataSource.Id {
		return d.GetSystemDataSourceBackend(), nil
	}

	systemCtx := withSystemContext(context.TODO())
	record, err := d.recordService.Get(systemCtx, &stub.GetRecordRequest{
		Resource: dataSourceResource.Name,
		Id:       dataSourceId,
	})

	if err != nil {
		return nil, err
	}

	dataSource := dataSourceFromRecord(record.Record)

	return d.GetDataSourceBackend(dataSource), nil
}

func (d *dataSourceService) GetDataSourceBackend(dataSource *model.DataSource) backend.DataSourceBackend {
	switch d.systemDataSource.Backend {
	case model.DataSourceBackend_POSTGRESQL:
		return postgres.NewPostgresDataSourceBackend(dataSource.Id, dataSource.Options.(*model.DataSource_PostgresqlParams).PostgresqlParams)
	case model.DataSourceBackend_MONGODB:
		panic("mongodb data-source not init")
	default:
		panic("unknown data-source type")
	}
}

func (d *dataSourceService) GetSystemDataSourceBackend() backend.DataSourceBackend {
	return d.GetDataSourceBackend(d.systemDataSource)
}

func (d *dataSourceService) InjectPostgresResourceServiceBackend(serviceBackend backend.ResourceServiceBackend) {
	d.postgresResourceServiceBackend = serviceBackend
}

func (d *dataSourceService) InjectResourceService(service ResourceService) {
	d.resourceService = service
}

func (d *dataSourceService) InjectRecordService(service RecordService) {
	d.recordService = service
}

func (d *dataSourceService) InjectInitData(data *model.InitData) {
	d.systemDataSource = data.SystemDataSource
}

func (d *dataSourceService) Init() {
	d.resourceService.InitResource(dataSourceResource)
}

func NewDataSourceService() DataSourceService {
	return &dataSourceService{}
}
