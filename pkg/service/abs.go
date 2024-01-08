package service

import (
	"context"
	"github.com/apibrew/apibrew/pkg/abs"
	"github.com/apibrew/apibrew/pkg/errors"
	"github.com/apibrew/apibrew/pkg/formats/unstructured"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/resource_model"
	"github.com/apibrew/apibrew/pkg/stub"
	"github.com/apibrew/apibrew/pkg/util/jwt-model"
	"google.golang.org/protobuf/types/known/structpb"
	"time"
)

type AuthenticationService interface {
	Init(config *model.AppConfig)
	Authenticate(ctx context.Context, username string, password string, term model.TokenTerm) (*model.Token, errors.ServiceError)
	AuthenticateWithoutPassword(ctx context.Context, username string, term model.TokenTerm) (*model.Token, errors.ServiceError)
	RenewToken(ctx context.Context, token string, term model.TokenTerm) (*model.Token, errors.ServiceError)
	ParseAndVerifyToken(token string) (*jwt_model.UserDetails, errors.ServiceError)
	AuthenticationDisabled() bool
}

type AuthorizationService interface {
	CheckRecordAccess(ctx context.Context, params CheckRecordAccessParams) errors.ServiceError
	CheckRecordAccessWithRecordSelector(ctx context.Context, params CheckRecordAccessParams) (*resource_model.BooleanExpression, errors.ServiceError)
	CheckIsExtensionController(ctx context.Context) errors.ServiceError
}

type BackendProviderService interface {
	abs.BackendRecordsInterface
	abs.BackendActionExecutor
	DestroyDataSource(ctx context.Context, dataSourceName string) errors.ServiceError
	ListEntities(ctx context.Context, dataSourceId string) ([]*model.DataSourceCatalog, errors.ServiceError)
	PrepareResourceFromEntity(ctx context.Context, dataSourceName string, catalog, entity string) (*model.Resource, errors.ServiceError)
	UpgradeResource(ctx context.Context, dataSourceName string, params abs.UpgradeResourceParams) errors.ServiceError
	GetStatus(ctx context.Context, dataSourceId string) (connectionAlreadyInitiated bool, testConnection bool, err errors.ServiceError)
	Init(config *model.AppConfig)
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
	PrepareQuery(resource *model.Resource, queryMap map[string]string) (*model.BooleanExpression, errors.ServiceError)
	GetRecord(ctx context.Context, namespace, resourceName, id string, references []string) (*model.Record, errors.ServiceError)
	FindBy(ctx context.Context, namespace, resourceName, propertyName string, value string) (*model.Record, errors.ServiceError)
	ResolveReferences(ctx context.Context, resource *model.Resource, records []*model.Record, referencesToResolve []string) errors.ServiceError
	List(ctx context.Context, params RecordListParams) ([]*model.Record, uint32, errors.ServiceError)
	Create(ctx context.Context, params RecordCreateParams) ([]*model.Record, errors.ServiceError)
	Update(ctx context.Context, params RecordUpdateParams) ([]*model.Record, errors.ServiceError)
	Apply(ctx context.Context, params RecordUpdateParams) ([]*model.Record, errors.ServiceError)
	Get(ctx context.Context, params RecordGetParams) (*model.Record, errors.ServiceError)
	Delete(ctx context.Context, params RecordDeleteParams) errors.ServiceError
	ExecuteAction(ctx context.Context, params ExecuteActionParams) (unstructured.Unstructured, errors.ServiceError)
	Load(ctx context.Context, namespace string, name string, properties map[string]*structpb.Value) (*model.Record, errors.ServiceError)
}

type ResourceService interface {
	Init(config *model.AppConfig)
	GetResourceByName(ctx context.Context, namespace, resource string) (*model.Resource, errors.ServiceError)
	GetSystemResourceByName(ctx context.Context, resourceName string) (*model.Resource, errors.ServiceError)
	Create(ctx context.Context, resource *model.Resource, doMigration bool, forceMigration bool) (*model.Resource, errors.ServiceError)
	Update(ctx context.Context, resource *model.Resource, doMigration bool, forceMigration bool) errors.ServiceError
	Delete(ctx context.Context, ids []string, doMigration bool, forceMigration bool) errors.ServiceError
	List(ctx context.Context) ([]*model.Resource, errors.ServiceError)
	Get(ctx context.Context, id string) (*model.Resource, errors.ServiceError)
	GetSchema() *abs.Schema
	PrepareResourceMigrationPlan(ctx context.Context, resources []*model.Resource, prepareFromDataSource bool) ([]*model.ResourceMigrationPlan, errors.ServiceError)
	LocateReferences(resource *model.Resource, resolve []string) []string
	LocateLocalReferences(resource *model.Resource) []string
	LocateResourceByReference(resource *model.Resource, reference *model.Reference) *model.Resource
}

type ResourceMigrationService interface {
	PreparePlan(ctx context.Context, existingResource *model.Resource, resource *model.Resource) (*model.ResourceMigrationPlan, errors.ServiceError)
}

type MetricsService interface {
	Init(config *model.AppConfig)
	GetMetrics(req MetricsRequest) ([]MetricsResponseItem, errors.ServiceError)
}

type AuditService interface {
	Init(config *model.AppConfig)
}

type StatsService interface {
	Init(config *model.AppConfig)
}

type WatchService interface {
	Watch(ctx context.Context, params WatchParams) (<-chan *model.Event, errors.ServiceError)
	WatchResource(ctx context.Context, params WatchParams) (<-chan *model.Event, errors.ServiceError)
}

type EventChannelService interface {
	Exec(ctx context.Context, channelKey string, event *model.Event) (*model.Event, errors.ServiceError)
	PollEvents(ctx context.Context, channelKey string) (chan *model.Event, errors.ServiceError)
	WriteEvent(ctx context.Context, proto *model.Event) errors.ServiceError
	Init(config *model.AppConfig)
}

type ExternalService interface {
	Call(ctx context.Context, all resource_model.ExternalCall, event *model.Event) (*model.Event, errors.ServiceError)
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
	Filters           map[string]string
	Aggregation       *model.Aggregation
}

func (p RecordListParams) ToRequest() *stub.ListRecordRequest {
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

type MetricsOperation string

const (
	MetricsOperationRead   MetricsOperation = "read"
	MetricsOperationWrite  MetricsOperation = "write"
	MetricsOperationDelete MetricsOperation = "delete"
)

type MetricsInterval string

const (
	MetricsIntervalMinute MetricsInterval = "minute"
	MetricsIntervalHour   MetricsInterval = "hour"
	MetricsIntervalDay    MetricsInterval = "day"
	MetricsIntervalWeek   MetricsInterval = "week"
	MetricsIntervalMonth  MetricsInterval = "month"
)

type MetricsRequest struct {
	Namespace *string           `json:"namespace"`
	Resource  *string           `json:"resource"`
	Operation *MetricsOperation `json:"operation"`
	Interval  *MetricsInterval  `json:"interval"`
	From      *time.Time
	To        *time.Time
}

type MetricsResponseItem struct {
	Namespace string           `json:"namespace"`
	Resource  string           `json:"resource"`
	Interval  MetricsInterval  `json:"interval"`
	Operation MetricsOperation `json:"operation"`
	Time      time.Time        `json:"time"`
	Count     uint64           `json:"count"`
}

type Module interface {
	Init()
}

type ModuleConstructor func(container Container) Module
