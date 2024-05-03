package abs

import (
	"context"
	"github.com/apibrew/apibrew/pkg/model"
)

type Backend interface {
	BackendGenericInterface
	BackendRecordsInterface
	BackendSchemaInterface

	SetSchema(schema *Schema)
}

type BackendGenericInterface interface {
	GetStatus(ctx context.Context) (connectionAlreadyInitiated bool, testConnection bool, err error)
	DestroyDataSource(ctx context.Context)
}

type BackendRecordsInterface interface {
	AddRecords(ctx context.Context, resource *model.Resource, records []*model.Record) ([]*model.Record, error)
	UpdateRecords(ctx context.Context, resource *model.Resource, records []*model.Record) ([]*model.Record, error)
	GetRecord(ctx context.Context, resource *model.Resource, id string, resolveReferences []string) (*model.Record, error)
	DeleteRecords(ctx context.Context, resource *model.Resource, ids []*model.Record) error
	ListRecords(ctx context.Context, resource *model.Resource, params ListRecordParams, resultChan chan<- *model.Record) ([]*model.Record, uint32, error)
}

type BackendSchemaInterface interface {
	ListEntities(ctx context.Context) ([]*model.DataSourceCatalog, error)
	PrepareResourceFromEntity(ctx context.Context, catalog, entity string) (*model.Resource, error)
	UpgradeResource(ctx context.Context, params UpgradeResourceParams) error
}

type ListRecordParams struct {
	Query             *model.BooleanExpression
	Limit             uint32
	Offset            uint64
	ResolveReferences []string
	Aggregation       *model.Aggregation
	Sorting           *model.Sorting
}

type UpgradeResourceParams struct {
	ForceMigration bool
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

type DataSource interface {
	GetName() string
}

type BackendConstructor func(dataSourceRecord DataSource) Backend
