package backend

import (
	"data-handler/stub/model"
)

type DataSourceBackend interface {
	GetDataSourceId() string
}

type AddResourceParams struct {
	Resource             *model.Resource
	AllowSystemAndStatic bool
	IgnoreIfExists       bool
	Migrate              bool
	ForceMigrate         bool
}

type AddRecordsParams struct {
	Resource *model.Resource
	Records  []*model.Record
}

type DataSourceLocator interface {
	GetDataSourceBackendById(dataSourceId string) (DataSourceBackend, error)
	GetSystemDataSourceBackend() DataSourceBackend
}

type ResourceServiceBackend interface {
	Init()
	AddResource(params AddResourceParams) (*model.Resource, error)
	AddRecords(params AddRecordsParams) ([]*model.Record, error)
	GetResourceByName(resourceName string) (*model.Resource, error)
	GetRecord(resource *model.Resource, id string) (*model.Record, error)
	DeleteRecords(resource *model.Resource, list []string) error
	DestroyDataSource(dataSourceId string)
	InjectDataSourceService(service DataSourceLocator)
}
