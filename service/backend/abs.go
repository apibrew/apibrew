package backend

import (
	"data-handler/stub/model"
)

type DataSourceBackend interface {
	GetDataSourceId() string
}

type AddResourceParams struct {
	Backend              DataSourceBackend
	Resource             *model.Resource
	AllowSystemAndStatic bool
	IgnoreIfExists       bool
	Migrate              bool
	ForceMigrate         bool
}

type AddRecordsParams struct {
	Backend  DataSourceBackend
	Resource *model.Resource
	Records  []*model.Record
}

type ResourceServiceBackend interface {
	Init(backend DataSourceBackend)
	AddResource(params AddResourceParams) (*model.Resource, error)
	AddRecords(params AddRecordsParams) ([]*model.Record, error)
	GetResourceByName(resourceName string) (*model.Resource, error)
}
