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
	GetDataSourceBackend(id string) (backend.DataSourceBackend, error)
	LocateDataSource(id string) (*model.DataSource, error)
	InjectRecordService(service RecordService)
	GetSystemDataSourceBackend() backend.DataSourceBackend
}

type dataSourceService struct {
	stub.DataSourceServiceServer
	resourceService  ResourceService
	recordService    RecordService
	systemDataSource *model.DataSource
}

func (d *dataSourceService) GetSystemDataSourceBackend() backend.DataSourceBackend {
	bck, err := d.GetDataSourceBackend(d.systemDataSource.Id)

	if err != nil {
		panic(err)
	}

	return bck
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

func (d *dataSourceService) LocateDataSource(dataSourceId string) (*model.DataSource, error) {
	if dataSourceId == d.systemDataSource.Id {
		return d.systemDataSource, nil
	} else {
		panic("not implemented")
	}
}

func (d *dataSourceService) GetDataSourceBackend(id string) (backend.DataSourceBackend, error) {
	if id != d.systemDataSource.Id {
		panic("not implemented")
	}

	switch d.systemDataSource.Backend {
	case model.DataSourceBackend_POSTGRESQL:
		return postgres.NewPostgresDataSourceBackend(id, d.systemDataSource.Options.(*model.DataSource_PostgresqlParams).PostgresqlParams), nil
	case model.DataSourceBackend_MONGODB:
		panic("mongodb data-source not init")
	default:
		panic("unknown data-source type")
	}
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
