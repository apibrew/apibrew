package service

import (
	"data-handler/service/backend"
	"data-handler/stub"
	model "data-handler/stub/model"
)

type DataSourceService interface {
	stub.DataSourceServiceServer
	InjectResourceService(service ResourceService)
	Init(data *model.InitData)
	GetDataSourceBackend(id string) (backend.DataSourceBackend, error)
	LocateDataSource(id string) (*model.DataSource, error)
}

type dataSourceService struct {
	stub.DataSourceServiceServer
	resourceService    ResourceService
	systemDataSource   *model.DataSource
	dataSourceResource *model.Resource
}

func (d *dataSourceService) LocateDataSource(dataSourceId string) (*model.DataSource, error) {
	if dataSourceId == "system" {
		return d.systemDataSource, nil
	} else {
		panic("not implemented")
	}
}

func (d *dataSourceService) GetDataSourceBackend(id string) (backend.DataSourceBackend, error) {
	if id != "system" {
		panic("not implemented")
	}

	switch d.systemDataSource.Backend {
	case model.DataSourceBackend_POSTGRESQL:
		return backend.NewPostgresDataSourceBackend(id, d.systemDataSource.Options.(*model.DataSource_PostgresqlParams).PostgresqlParams), nil
	case model.DataSourceBackend_MONGODB:
		panic("mongodb data-source not init")
	default:
		panic("unknown data-source type")
	}
}

func (d *dataSourceService) InjectResourceService(service ResourceService) {
	d.resourceService = service
}

func (d *dataSourceService) Init(data *model.InitData) {
	d.systemDataSource = data.SystemDataSource

	d.resourceService.InitResource(d.dataSourceResource)
}

func NewDataSourceService() DataSourceService {
	return &dataSourceService{
		dataSourceResource: prepareDataSourceResource(),
	}
}

func prepareDataSourceResource() *model.Resource {
	return &model.Resource{
		Name:      "data-source",
		Workspace: "system",
		Type:      model.DataType_SYSTEM,
		SourceConfig: &model.ResourceSourceConfig{
			DataSource: "system",
			Mapping:    "data_source",
		},
		Properties: []*model.ResourceProperty{
			{
				Name: "backend",
				SourceConfig: &model.ResourceProperty_Mapping{
					Mapping: "backend",
				},
				Type:     model.ResourcePropertyType_INT32,
				Required: true,
			},
			{
				Name: "options_postgres_username",
				SourceConfig: &model.ResourceProperty_Mapping{
					Mapping: "options_postgres_username",
				},
				Type:     model.ResourcePropertyType_STRING,
				Length:   64,
				Required: false,
			},
			{
				Name: "options_postgres_password",
				SourceConfig: &model.ResourceProperty_Mapping{
					Mapping: "options_postgres_password",
				},
				Type:     model.ResourcePropertyType_STRING,
				Length:   64,
				Required: false,
			},
			{
				Name: "options_postgres_host",
				SourceConfig: &model.ResourceProperty_Mapping{
					Mapping: "options_postgres_host",
				},
				Type:     model.ResourcePropertyType_STRING,
				Length:   64,
				Required: false,
			},
			{
				Name: "options_postgres_port",
				SourceConfig: &model.ResourceProperty_Mapping{
					Mapping: "options_postgres_port",
				},
				Type:     model.ResourcePropertyType_INT32,
				Required: false,
			},
			{
				Name: "options_postgres_db_name",
				SourceConfig: &model.ResourceProperty_Mapping{
					Mapping: "options_postgres_db_name",
				},
				Type:     model.ResourcePropertyType_STRING,
				Length:   64,
				Required: false,
			},
			{
				Name: "options_postgres_default_schema",
				SourceConfig: &model.ResourceProperty_Mapping{
					Mapping: "options_postgres_default_schema",
				},
				Type:     model.ResourcePropertyType_STRING,
				Length:   64,
				Required: false,
			},
		},
	}
}
