package service

import (
	"context"
	"github.com/apibrew/apibrew/pkg/abs"
	"github.com/apibrew/apibrew/pkg/formats/unstructured"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/resource_model"
	"github.com/apibrew/apibrew/pkg/stub"
	"github.com/apibrew/apibrew/pkg/util/jwt-model"
	"google.golang.org/protobuf/types/known/structpb"
)

type AuthenticationService interface {
	Init(config *model.AppConfig)
	Authenticate(ctx context.Context, username string, password string, term model.TokenTerm, minimizeToken bool) (*model.Token, error)
	AuthenticateWithoutPassword(ctx context.Context, username string, term model.TokenTerm) (*model.Token, error)
	RenewToken(ctx context.Context, token string, term model.TokenTerm) (*model.Token, error)
	GetToken(ctx context.Context) (*jwt_model.UserDetails, error)
	ParseAndVerifyToken(token string) (*jwt_model.UserDetails, error)
	AuthenticationDisabled() bool
}

type AuthorizationService interface {
	CheckRecordAccess(ctx context.Context, params CheckRecordAccessParams) error
	CheckRecordAccessWithRecordSelector(ctx context.Context, params CheckRecordAccessParams) (*resource_model.BooleanExpression, error)
	CheckIsExtensionController(ctx context.Context) error
}

type BackendProviderService interface {
	abs.BackendRecordsInterface
	abs.BackendActionExecutor
	DestroyDataSource(ctx context.Context, dataSourceName string) error
	ListEntities(ctx context.Context, dataSourceId string) ([]*model.DataSourceCatalog, error)
	PrepareResourceFromEntity(ctx context.Context, dataSourceName string, catalog, entity string) (*model.Resource, error)
	UpgradeResource(ctx context.Context, dataSourceName string, params abs.UpgradeResourceParams) error
	GetStatus(ctx context.Context, dataSourceId string) (connectionAlreadyInitiated bool, testConnection bool, err error)
	Init(config *model.AppConfig)
	SetSchema(schema *abs.Schema)
}

type DataSourceService interface {
	Init(config *model.AppConfig)
	ListEntities(ctx context.Context, id string) ([]*model.DataSourceCatalog, error)
	GetStatus(ctx context.Context, id string) (connectionAlreadyInitiated bool, testConnection bool, err error)
	PrepareResourceFromEntity(ctx context.Context, dataSourceId string, catalog, entity string) (*model.Resource, error)
	Delete(ctx context.Context, ids []string) error
}

type RecordService interface {
	Init(config *model.AppConfig)
	PrepareQuery(resource *model.Resource, queryMap map[string]interface{}) (*model.BooleanExpression, error)
	GetRecord(ctx context.Context, namespace, resourceName, id string, references []string) (*model.Record, error)
	FindBy(ctx context.Context, namespace, resourceName, propertyName string, value string) (*model.Record, error)
	ResolveReferences(ctx context.Context, resource *model.Resource, records []*model.Record, referencesToResolve []string) error
	List(ctx context.Context, params RecordListParams) ([]*model.Record, uint32, error)
	Create(ctx context.Context, params RecordCreateParams) ([]*model.Record, error)
	Update(ctx context.Context, params RecordUpdateParams) ([]*model.Record, error)
	Apply(ctx context.Context, params RecordUpdateParams) ([]*model.Record, error)
	Get(ctx context.Context, params RecordGetParams) (*model.Record, error)
	Delete(ctx context.Context, params RecordDeleteParams) error
	Load(ctx context.Context, namespace string, name string, properties map[string]*structpb.Value, listParams RecordLoadParams) (*model.Record, error)
}

type ResourceService interface {
	Init(config *model.AppConfig)
	GetResourceByName(ctx context.Context, namespace, resource string) (*model.Resource, error)
	GetSystemResourceByName(ctx context.Context, resourceName string) (*model.Resource, error)
	Create(ctx context.Context, resource *model.Resource, doMigration bool, forceMigration bool) (*model.Resource, error)
	Update(ctx context.Context, resource *model.Resource, doMigration bool, forceMigration bool) error
	Delete(ctx context.Context, ids []string, doMigration bool, forceMigration bool) error
	List(ctx context.Context) ([]*model.Resource, error)
	Get(ctx context.Context, id string) (*model.Resource, error)
	GetSchema() *abs.Schema
	PrepareResourceMigrationPlan(ctx context.Context, resources []*model.Resource, prepareFromDataSource bool) ([]*model.ResourceMigrationPlan, error)
	LocateReferences(resource *model.Resource, resolve []string) []string
	LocateLocalReferences(resource *model.Resource) []string
	LocateResourceByReference(resource *model.Resource, reference *model.Reference) *model.Resource
}

type ResourceMigrationService interface {
	PreparePlan(ctx context.Context, existingResource *model.Resource, resource *model.Resource) (*model.ResourceMigrationPlan, error)
}

type AuditService interface {
	Init(config *model.AppConfig)
}

type StatsService interface {
	Init(config *model.AppConfig)
}

type WatchService interface {
	Watch(ctx context.Context, params WatchParams) (<-chan *model.Event, error)
	WatchResource(ctx context.Context, params WatchParams) (<-chan *model.Event, error)
}

type EventChannelService interface {
	Exec(ctx context.Context, channelKey string, event *model.Event) (*model.Event, error)
	PollEvents(ctx context.Context, channelKey string) (chan *model.Event, error)
	WriteEvent(ctx context.Context, proto *model.Event) error
	Init(config *model.AppConfig)
}

type ExternalService interface {
	Call(ctx context.Context, all resource_model.ExternalCall, event *model.Event) (*model.Event, error)
}

type ExtensionService interface {
	RegisterExtension(*resource_model.Extension)
	UnRegisterExtension(*resource_model.Extension)
	Init(config *model.AppConfig)
	Reload()
}

type CheckRecordAccessParams struct {
	Resource  *model.Resource
	Records   *[]*model.Record
	Operation resource_model.PermissionOperation
}

type WatchParams struct {
	Selector   *model.EventSelector
	BufferSize int
}

type WatchResourceParams struct {
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
	Filters           map[string]interface{}
	Aggregation       *model.Aggregation
	Sorting           *model.Sorting
}

type RecordLoadParams struct {
	UseHistory        bool
	ResolveReferences []string
}

func (p RecordListParams) ToRequest() *stub.ListRecordRequest {
	var filters = make(map[string]*structpb.Value)

	for k, v := range p.Filters {
		val, err := unstructured.ToValue(v)

		if err != nil {
			panic(err)
		}

		filters[k] = val
	}

	return &stub.ListRecordRequest{
		Namespace:         p.Namespace,
		Resource:          p.Resource,
		Filters:           filters,
		Limit:             p.Limit,
		Offset:            p.Offset,
		UseHistory:        p.UseHistory,
		ResolveReferences: p.ResolveReferences,
		Annotations:       nil,
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
	Namespace         string
	Resource          string
	Id                string
	ResolveReferences []string
}

type RecordDeleteParams struct {
	Namespace string
	Resource  string
	Ids       []string
}

type ExecuteActionParams struct {
	Namespace  string
	Resource   string
	Id         string
	ActionName string
	Input      unstructured.Unstructured
}

func (p RecordDeleteParams) ToRequest() *stub.DeleteRecordRequest {
	return &stub.DeleteRecordRequest{
		Namespace: p.Namespace,
		Resource:  p.Resource,
		Ids:       p.Ids,
	}
}

type Module interface {
	Init()
}

type ModuleConstructor func(container Container) Module
