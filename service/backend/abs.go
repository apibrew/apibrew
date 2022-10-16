package backend

import (
	"context"
	"data-handler/stub"
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

type BulkRecordsParams struct {
	Resource *model.Resource
	Records  []*model.Record
}

type ListRecordParams struct {
	Resource *model.Resource
	Query    *model.BooleanExpression
	Limit    uint32
	Offset   uint64
}

type DataSourceLocator interface {
	GetDataSourceBackendById(dataSourceId string) (DataSourceBackend, error)
	GetSystemDataSourceBackend() DataSourceBackend
}

type ResourceServiceBackend interface {
	Init()
	AddResource(params AddResourceParams) (*model.Resource, error)
	AddRecords(params BulkRecordsParams) ([]*model.Record, error)
	UpdateRecords(params BulkRecordsParams) ([]*model.Record, error)
	GetResourceByName(resourceName string) (*model.Resource, error)
	GetRecord(resource *model.Resource, id string) (*model.Record, error)
	DeleteRecords(resource *model.Resource, list []string) error
	DestroyDataSource(dataSourceId string)
	InjectDataSourceService(service DataSourceLocator)
	GetStatus(dataSourceId string) (*stub.StatusResponse, error)
	ListRecords(params ListRecordParams) ([]*model.Record, uint32, error)
	PrepareResourceFromEntity(ctx context.Context, dataSourceId string, entity string) (*model.Resource, error)
	DeleteResources(ctx context.Context, ids []string) error
}
