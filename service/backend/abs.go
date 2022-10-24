package backend

import (
	"context"
	"data-handler/service/errors"
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
	GetDataSourceBackendById(dataSourceId string) (DataSourceBackend, errors.ServiceError)
	GetSystemDataSourceBackend() DataSourceBackend
}

type ResourceServiceBackend interface {
	Init()
	AddResource(params AddResourceParams) (*model.Resource, errors.ServiceError)
	AddRecords(params BulkRecordsParams) ([]*model.Record, bool, errors.ServiceError)
	UpdateRecords(params BulkRecordsParams) ([]*model.Record, errors.ServiceError)
	GetResourceByName(ctx context.Context, resourceName string, name string) (*model.Resource, errors.ServiceError)
	GetRecord(resource *model.Resource, id string) (*model.Record, errors.ServiceError)
	DeleteRecords(resource *model.Resource, list []string) errors.ServiceError
	DestroyDataSource(dataSourceId string)
	InjectDataSourceService(service DataSourceLocator)
	GetStatus(dataSourceId string) (*stub.StatusResponse, errors.ServiceError)
	ListRecords(params ListRecordParams) ([]*model.Record, uint32, errors.ServiceError)
	PrepareResourceFromEntity(ctx context.Context, dataSourceId string, entity string) (*model.Resource, errors.ServiceError)
	DeleteResources(ctx context.Context, workspace string, ids []string, migration bool, forceMigration bool) errors.ServiceError
	ListEntities(ctx context.Context, dataSourceId string) ([]string, errors.ServiceError)
	UpdateResource(ctx context.Context, resource *model.Resource, doMigration bool, forceMigration bool) errors.ServiceError
	ListResources(ctx context.Context) ([]*model.Resource, errors.ServiceError)
}
