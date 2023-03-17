package abs

import (
	"context"
	"github.com/tislib/data-handler/pkg/errors"
	"github.com/tislib/data-handler/pkg/model"
)

type Backend interface {
	BackendGenericInterface
	BackendRecordsInterface
	BackendSchemaInterface
	BackendTransactionInterface
}

type BackendGenericInterface interface {
	GetStatus(ctx context.Context) (connectionAlreadyInitiated bool, testConnection bool, err errors.ServiceError)
	DestroyDataSource(ctx context.Context)
}

type BackendRecordsInterface interface {
	AddRecords(ctx context.Context, params BulkRecordsParams) ([]*model.Record, bool, errors.ServiceError)
	UpdateRecords(ctx context.Context, params BulkRecordsParams) ([]*model.Record, errors.ServiceError)
	GetRecord(ctx context.Context, resource *model.Resource, schema *Schema, id string) (*model.Record, errors.ServiceError)
	DeleteRecords(ctx context.Context, resource *model.Resource, list []string) errors.ServiceError
	ListRecords(ctx context.Context, params ListRecordParams) ([]*model.Record, uint32, errors.ServiceError)
}

type BackendSchemaInterface interface {
	ListEntities(ctx context.Context) ([]*model.DataSourceCatalog, errors.ServiceError)
	PrepareResourceFromEntity(ctx context.Context, catalog, entity string) (*model.Resource, errors.ServiceError)
	UpgradeResource(ctx context.Context, params UpgradeResourceParams) errors.ServiceError
}

type BackendTransactionInterface interface {
	BeginTransaction(ctx context.Context, readOnly bool) (transactionKey string, serviceError errors.ServiceError)
	CommitTransaction(ctx context.Context) (serviceError errors.ServiceError)
	RollbackTransaction(ctx context.Context) (serviceError errors.ServiceError)
	IsTransactionAlive(ctx context.Context) (isAlive bool, serviceError errors.ServiceError)
}

type BulkRecordsParams struct {
	Resource       *model.Resource
	Records        []*model.Record
	CheckVersion   bool
	IgnoreIfExists bool
	Schema         *Schema
}

type ListRecordParams struct {
	Resource          *model.Resource
	Query             *model.BooleanExpression
	Limit             uint32
	Offset            uint64
	UseHistory        bool
	ResolveReferences []string
	Schema            *Schema
	ResultChan        chan<- *model.Record
	PackRecords       bool
}

type UpgradeResourceParams struct {
	ForceMigration bool
	Schema         *Schema
	MigrationPlan  *model.ResourceMigrationPlan
}

type AddResourceParams struct {
	Resource       *model.Resource
	IgnoreIfExists bool
	Migrate        bool
	ForceMigrate   bool
}

type DataSourceConnectionDetails interface {
}

type BackendConstructor func(dataSource *model.DataSource) Backend
