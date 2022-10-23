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
	Resource       *model.Resource
	IgnoreIfExists bool
	Migrate        bool
	ForceMigrate   bool
}

type BulkRecordsParams struct {
	Resource       *model.Resource
	Records        []*model.Record
	CheckVersion   bool
	IgnoreIfExists bool
}

type ListRecordParams struct {
	Resource   *model.Resource
	Query      *model.BooleanExpression
	Limit      uint32
	Offset     uint64
	UseHistory bool
}

type DataSourceLocator interface {
	GetDataSourceBackendById(dataSourceId string) (DataSourceBackend, error)
	GetSystemDataSourceBackend() DataSourceBackend
}

type ResourceServiceBackend interface {
	Init()
	AddResource(params AddResourceParams) (*model.Resource, error)
	AddRecords(params BulkRecordsParams) ([]*model.Record, bool, error)
	UpdateRecords(params BulkRecordsParams) ([]*model.Record, error)
	GetResourceByName(resourceName string) (*model.Resource, error)
	GetRecord(resource *model.Resource, id string) (*model.Record, error)
	DeleteRecords(resource *model.Resource, list []string) error
	DestroyDataSource(dataSourceId string)
	InjectDataSourceService(service DataSourceLocator)
	GetStatus(dataSourceId string) (*stub.StatusResponse, error)
	ListRecords(params ListRecordParams) ([]*model.Record, uint32, error)
	PrepareResourceFromEntity(ctx context.Context, dataSourceId string, entity string) (*model.Resource, error)
	DeleteResources(ctx context.Context, ids []string, migration bool, forceMigration bool) error
	ListEntities(ctx context.Context, dataSourceId string) ([]string, error)
	UpdateResource(ctx context.Context, resource *model.Resource, doMigration bool, forceMigration bool) error
	ListResources(ctx context.Context) ([]*model.Resource, error)
}
