package service

import (
	"context"
	"github.com/apibrew/apibrew/pkg/abs"
	"github.com/apibrew/apibrew/pkg/errors"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/resource_model"
	"github.com/apibrew/apibrew/pkg/stub"
	"google.golang.org/protobuf/proto"
)

type AuthenticationService interface {
	Init(config *model.AppConfig)
	Authenticate(ctx context.Context, username string, password string, term model.TokenTerm) (*model.Token, errors.ServiceError)
	RenewToken(ctx context.Context, token string, term model.TokenTerm) (*model.Token, errors.ServiceError)
	ParseAndVerifyToken(token string) (*abs.UserDetails, errors.ServiceError)
	AuthenticationDisabled() bool
}

type AuthorizationService interface {
	CheckRecordAccess(ctx context.Context, params CheckRecordAccessParams) errors.ServiceError
}

type BackendProviderService interface {
	Init(config *model.AppConfig)
	GetSystemBackend(ctx context.Context) abs.Backend
	GetBackendByDataSourceId(ctx context.Context, dataSourceId string) (abs.Backend, errors.ServiceError)
	GetBackendByDataSourceName(ctx context.Context, dataSourceId string) (abs.Backend, errors.ServiceError)
	DestroyBackend(ctx context.Context, id string) error
	SetSchema(schema *abs.Schema)
}

type DataSourceService interface {
	Init(config *model.AppConfig)
	ListEntities(ctx context.Context, id string) ([]*model.DataSourceCatalog, errors.ServiceError)
	GetStatus(ctx context.Context, id string) (connectionAlreadyInitiated bool, testConnection bool, err errors.ServiceError)
	PrepareResourceFromEntity(ctx context.Context, dataSourceId string, catalog, entity string) (*model.Resource, errors.ServiceError)
	Delete(ctx context.Context, ids []string) errors.ServiceError
}

type RecordService interface {
	Init(config *model.AppConfig)
	PrepareQuery(resource *model.Resource, queryMap map[string]interface{}) (*model.BooleanExpression, errors.ServiceError)
	GetRecord(ctx context.Context, namespace, resourceName, id string) (*model.Record, errors.ServiceError)
	FindBy(ctx context.Context, namespace, resourceName, propertyName string, value interface{}) (*model.Record, errors.ServiceError)

	List(ctx context.Context, params RecordListParams) ([]*model.Record, uint32, errors.ServiceError)
	Create(ctx context.Context, params RecordCreateParams) ([]*model.Record, errors.ServiceError)
	Update(ctx context.Context, params RecordUpdateParams) ([]*model.Record, errors.ServiceError)
	Apply(ctx context.Context, params RecordUpdateParams) ([]*model.Record, errors.ServiceError)
	Get(ctx context.Context, params RecordGetParams) (*model.Record, errors.ServiceError)
	Delete(ctx context.Context, params RecordDeleteParams) errors.ServiceError
}

type ResourceService interface {
	Init(config *model.AppConfig)
	CheckResourceExists(ctx context.Context, namespace, name string) bool
	GetResourceByName(ctx context.Context, namespace, resource string) *model.Resource
	GetSystemResourceByName(ctx context.Context, resourceName string) *model.Resource
	Create(ctx context.Context, resource *model.Resource, doMigration bool, forceMigration bool) (*model.Resource, errors.ServiceError)
	Update(ctx context.Context, resource *model.Resource, doMigration bool, forceMigration bool) errors.ServiceError
	Delete(ctx context.Context, ids []string, doMigration bool, forceMigration bool) errors.ServiceError
	List(ctx context.Context) []*model.Resource
	Get(ctx context.Context, id string) *model.Resource
	GetSchema() *abs.Schema
	PrepareResourceMigrationPlan(ctx context.Context, resources []*model.Resource, prepareFromDataSource bool) ([]*model.ResourceMigrationPlan, errors.ServiceError)
}

type ResourceMigrationService interface {
	PreparePlan(ctx context.Context, existingResource *model.Resource, resource *model.Resource) (*model.ResourceMigrationPlan, errors.ServiceError)
}

type GenericRecordService[T proto.Message] interface {
	Init(records []T)
	Create(ctx context.Context, users []T) ([]T, errors.ServiceError)
	Update(ctx context.Context, users []T) ([]T, errors.ServiceError)
	Apply(ctx context.Context, users []T) ([]T, errors.ServiceError)
	Delete(ctx context.Context, ids []string) errors.ServiceError
	Get(ctx context.Context, id string) (T, errors.ServiceError)
	List(ctx context.Context, query *model.BooleanExpression, limit uint32, offset uint64) ([]T, errors.ServiceError)
}

type WatchService interface {
	Watch(ctx context.Context, params WatchParams) <-chan *model.Event
}

type ExternalService interface {
	Call(ctx context.Context, all resource_model.ExtensionExternalCall, event *model.Event) (*model.Event, errors.ServiceError)
}

type ExtensionService interface {
}

type CheckRecordAccessParams struct {
	Resource  *model.Resource
	Records   *[]*model.Record
	Operation model.OperationType
}

type WatchParams struct {
	Selector   *model.EventSelector
	BufferSize int
}

type RecordListParams struct {
	Query             *model.BooleanExpression
	Namespace         string
	Resource          string
	Limit             uint32
	Offset            uint64
	UseHistory        bool
	ResolveReferences []string
	ResultChan        chan<- *model.Record
	PackRecords       bool
	Filters           map[string]string
}

func (p RecordListParams) ToRequest() proto.Message {
	return &stub.ListRecordRequest{
		Namespace:         p.Namespace,
		Resource:          p.Resource,
		Filters:           p.Filters,
		Limit:             p.Limit,
		Offset:            p.Offset,
		UseHistory:        p.UseHistory,
		ResolveReferences: p.ResolveReferences,
	}
}

type RecordCreateParams struct {
	Namespace string
	Resource  string
	Records   []*model.Record
}

func (p RecordCreateParams) ToRequest() *stub.CreateRecordRequest {
	return &stub.CreateRecordRequest{
		Namespace: p.Namespace,
		Resource:  p.Resource,
		Records:   p.Records,
	}
}

type RecordUpdateParams struct {
	Namespace string
	Resource  string
	Records   []*model.Record
}

func (p RecordUpdateParams) ToRequest() *stub.UpdateRecordRequest {
	return &stub.UpdateRecordRequest{
		Namespace: p.Namespace,
		Resource:  p.Resource,
		Records:   p.Records,
	}
}

type RecordGetParams struct {
	Namespace string
	Resource  string
	Id        string
}

type RecordDeleteParams struct {
	Namespace string
	Resource  string
	Ids       []string
}

func (p RecordDeleteParams) ToRequest() *stub.DeleteRecordRequest {
	return &stub.DeleteRecordRequest{
		Namespace: p.Namespace,
		Resource:  p.Resource,
		Ids:       p.Ids,
	}
}
