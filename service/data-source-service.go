package service

import (
	"context"
	"data-handler/service/backend"
	"data-handler/service/backend/postgres"
	mapping "data-handler/service/mapping"
	"data-handler/service/security"
	"data-handler/service/system"
	"data-handler/stub"
	model "data-handler/stub/model"
	log "github.com/sirupsen/logrus"
)

type DataSourceService interface {
	stub.DataSourceServiceServer
	InjectResourceService(service ResourceService)
	InjectInitData(data *model.InitData)
	Init(*model.InitData)
	GetDataSourceBackend(dataSource *model.DataSource) backend.DataSourceBackend
	InjectRecordService(service RecordService)
	GetSystemDataSourceBackend() backend.DataSourceBackend
	InjectPostgresResourceServiceBackend(serviceBackend backend.ResourceServiceBackend)
	GetDataSourceBackendById(dataSourceId string) (backend.DataSourceBackend, error)
	InjectAuthenticationService(service AuthenticationService)
}

type dataSourceService struct {
	stub.DataSourceServiceServer
	resourceService                ResourceService
	recordService                  RecordService
	systemDataSource               *model.DataSource
	postgresResourceServiceBackend backend.ResourceServiceBackend
	authenticationService          AuthenticationService
	ServiceName                    string
}

func (d *dataSourceService) InjectAuthenticationService(service AuthenticationService) {
	d.authenticationService = service
}

func (d *dataSourceService) ListEntities(ctx context.Context, request *stub.ListEntitiesRequest) (*stub.ListEntitiesResponse, error) {
	res, err := d.postgresResourceServiceBackend.ListEntities(ctx, request.Id)

	if err != nil {
		return nil, err
	}

	return &stub.ListEntitiesResponse{
		Entities: res,
		Error:    nil,
	}, nil
}

func (d *dataSourceService) List(ctx context.Context, request *stub.ListDataSourceRequest) (*stub.ListDataSourceResponse, error) {
	systemCtx := security.WithSystemContext(ctx)
	result, err := d.recordService.List(systemCtx, &stub.ListRecordRequest{
		Resource: system.DataSourceResource.Name,
		Token:    request.Token,
	})

	if err != nil {
		return nil, err
	}

	return &stub.ListDataSourceResponse{
		Content: mapping.MapFromRecord(result.Content, mapping.DataSourceFromRecord),
		Error:   result.Error,
	}, err
}

func (d *dataSourceService) Status(ctx context.Context, request *stub.StatusRequest) (*stub.StatusResponse, error) {
	return d.postgresResourceServiceBackend.GetStatus(request.Id)
}

func (d *dataSourceService) Create(ctx context.Context, request *stub.CreateDataSourceRequest) (*stub.CreateDataSourceResponse, error) {
	// insert records via resource service
	records := mapping.MapToRecord(request.DataSources, mapping.DataSourceToRecord)
	systemCtx := security.WithSystemContext(ctx)
	result, err := d.recordService.Create(systemCtx, &stub.CreateRecordRequest{
		Token:   request.Token,
		Records: records,
	})

	if err != nil {
		return nil, err
	}

	return &stub.CreateDataSourceResponse{
		DataSources: mapping.MapFromRecord(result.Records, mapping.DataSourceFromRecord),
		Error:       result.Error,
	}, err
}

func (d *dataSourceService) Update(ctx context.Context, request *stub.UpdateDataSourceRequest) (*stub.UpdateDataSourceResponse, error) {
	// insert records via resource service
	records := mapping.MapToRecord(request.DataSources, mapping.DataSourceToRecord)
	systemCtx := security.WithSystemContext(ctx)
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
		DataSources: mapping.MapFromRecord(result.Records, mapping.DataSourceFromRecord),
		Error:       result.Error,
	}, err
}

func (d *dataSourceService) PrepareResourceFromEntity(ctx context.Context, request *stub.PrepareResourceFromEntityRequest) (*stub.PrepareResourceFromEntityResponse, error) {
	resource, err := d.postgresResourceServiceBackend.PrepareResourceFromEntity(ctx, request.Id, request.Entity)

	if err != nil {
		return nil, err
	}

	return &stub.PrepareResourceFromEntityResponse{
		Resource: resource,
		Error:    nil,
	}, nil
}

func (d *dataSourceService) Get(ctx context.Context, request *stub.GetDataSourceRequest) (*stub.GetDataSourceResponse, error) {
	systemCtx := security.WithSystemContext(ctx)
	record, err := d.recordService.Get(systemCtx, &stub.GetRecordRequest{
		Token:    request.Token,
		Resource: system.DataSourceResource.Name,
		Id:       request.Id,
	})

	if err != nil {
		return nil, err
	}

	return &stub.GetDataSourceResponse{
		DataSource: mapping.DataSourceFromRecord(record.Record),
		Error:      record.Error,
	}, nil
}

func (d *dataSourceService) Delete(ctx context.Context, request *stub.DeleteDataSourceRequest) (*stub.DeleteDataSourceResponse, error) {
	systemCtx := security.WithSystemContext(ctx)

	record, err := d.recordService.Delete(systemCtx, &stub.DeleteRecordRequest{
		Token:    request.Token,
		Resource: system.DataSourceResource.Name,
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

	systemCtx := security.WithSystemContext(context.TODO())
	record, err := d.recordService.Get(systemCtx, &stub.GetRecordRequest{
		Resource: system.DataSourceResource.Name,
		Id:       dataSourceId,
	})

	if err != nil {
		return nil, err
	}

	dataSource := mapping.DataSourceFromRecord(record.Record)

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

func (d *dataSourceService) Init(data *model.InitData) {
	d.resourceService.InitResource(system.DataSourceResource)

	if len(data.InitDataSources) > 0 {
		res, err := d.recordService.Create(security.SystemContext, &stub.CreateRecordRequest{
			Records:        mapping.MapToRecord(data.InitDataSources, mapping.DataSourceToRecord),
			IgnoreIfExists: true,
		})

		if err != nil || res.Error != nil {
			log.Error(err, res.Error)
		}
	}
}

func NewDataSourceService() DataSourceService {
	return &dataSourceService{
		ServiceName: "DataSourceService",
	}
}
