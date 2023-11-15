package abs

import (
	"context"
	"github.com/apibrew/apibrew/pkg/errors"
	"github.com/apibrew/apibrew/pkg/formats/unstructured"
	"github.com/apibrew/apibrew/pkg/model"
)

type Backend interface {
	BackendGenericInterface
	BackendRecordsInterface
	BackendSchemaInterface

	SetSchema(schema *Schema)
}

type BackendActionExecutor interface {
	ExecuteAction(ctx context.Context, resource *model.Resource, rec *model.Record, actionName string, input unstructured.Any) (unstructured.Unstructured, errors.ServiceError)
}

type BackendGenericInterface interface {
	GetStatus(ctx context.Context) (connectionAlreadyInitiated bool, testConnection bool, err errors.ServiceError)
	DestroyDataSource(ctx context.Context)
}

type BackendRecordsInterface interface {
	AddRecords(ctx context.Context, resource *model.Resource, records []*model.Record) ([]*model.Record, errors.ServiceError)
	UpdateRecords(ctx context.Context, resource *model.Resource, records []*model.Record) ([]*model.Record, errors.ServiceError)
	GetRecord(ctx context.Context, resource *model.Resource, id string, resolveReferences []string) (*model.Record, errors.ServiceError)
	DeleteRecords(ctx context.Context, resource *model.Resource, ids []*model.Record) errors.ServiceError
	ListRecords(ctx context.Context, resource *model.Resource, params ListRecordParams, resultChan chan<- *model.Record) ([]*model.Record, uint32, errors.ServiceError)
}

type BackendSchemaInterface interface {
	ListEntities(ctx context.Context) ([]*model.DataSourceCatalog, errors.ServiceError)
	PrepareResourceFromEntity(ctx context.Context, catalog, entity string) (*model.Resource, errors.ServiceError)
	UpgradeResource(ctx context.Context, params UpgradeResourceParams) errors.ServiceError
}

type ListRecordParams struct {
	Query             *model.BooleanExpression
	Limit             uint32
	Offset            uint64
	ResolveReferences []string
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
