package backend

import (
	"context"
	"data-handler/model"
	"data-handler/service/errors"
)

type DataSourceBackend interface {
	GetDataSourceId() string
	GetBackend() model.DataSourceBackend
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
	Resource          *model.Resource
	Query             *model.BooleanExpression
	Limit             uint32
	Offset            uint64
	UseHistory        bool
	ResolveReferences bool
}

type DataSourceLocator interface {
	GetDataSourceBackendById(ctx context.Context, dataSourceId string) (DataSourceBackend, errors.ServiceError)
	GetSystemDataSourceBackend(ctx context.Context) DataSourceBackend
}

type ResourceServiceBackend interface {
	Init()
	AddResource(ctx context.Context, params AddResourceParams) (*model.Resource, errors.ServiceError)
	AddRecords(ctx context.Context, params BulkRecordsParams) ([]*model.Record, bool, errors.ServiceError)
	UpdateRecords(ctx context.Context, params BulkRecordsParams) ([]*model.Record, errors.ServiceError)
	GetResourceByName(ctx context.Context, resourceName string, name string) (*model.Resource, errors.ServiceError)
	GetResource(ctx context.Context, workspace string, id string) (*model.Resource, errors.ServiceError)
	GetRecord(ctx context.Context, resource *model.Resource, id string) (*model.Record, errors.ServiceError)
	DeleteRecords(ctx context.Context, resource *model.Resource, list []string) errors.ServiceError
	DestroyDataSource(ctx context.Context, dataSourceId string)
	InjectDataSourceService(service DataSourceLocator)
	GetStatus(ctx context.Context, dataSourceId string) (connectionAlreadyInitiated bool, testConnection bool, err errors.ServiceError)
	ListRecords(ctx context.Context, params ListRecordParams) ([]*model.Record, uint32, errors.ServiceError)
	PrepareResourceFromEntity(ctx context.Context, dataSourceId string, entity string) (*model.Resource, errors.ServiceError)
	DeleteResources(ctx context.Context, workspace string, ids []string, migration bool, forceMigration bool) errors.ServiceError
	ListEntities(ctx context.Context, dataSourceId string) ([]string, errors.ServiceError)
	UpdateResource(ctx context.Context, resource *model.Resource, doMigration bool, forceMigration bool) errors.ServiceError
	ListResources(ctx context.Context) ([]*model.Resource, errors.ServiceError)
}
