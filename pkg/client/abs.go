package client

import (
	"context"
	"github.com/apibrew/apibrew/pkg/abs"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/service"
	"github.com/apibrew/apibrew/pkg/stub"
	"google.golang.org/protobuf/types/known/structpb"
)

type ApplyInterface interface {
	ApplyResource(ctx context.Context, resource *model.Resource, doMigration, forceMigration bool) error
}

type Client interface {
	ApplyInterface
	AuthenticateWithToken(token string)
	AuthenticateWithUsernameAndPassword(username string, password string) error
	NewRemoteExtension(host string, remoteHost string) Extension
	NewPollExtension() Extension
	UpdateTokenFromContext(ctx context.Context)

	// record
	CreateRecord(ctx context.Context, namespace string, resource string, record *model.Record) (*model.Record, error)
	UpdateRecord(ctx context.Context, namespace string, resource string, record *model.Record) (*model.Record, error)
	DeleteRecord(ctx context.Context, namespace string, name string, record *model.Record) error
	ApplyRecord(ctx context.Context, namespace string, resource string, record *model.Record) (*model.Record, error)
	GetRecord(ctx context.Context, namespace string, resource string, id string) (*model.Record, error)
	LoadRecord(ctx context.Context, namespace string, resource string, properties map[string]*structpb.Value, params service.RecordLoadParams) (*model.Record, error)
	ListRecords(ctx context.Context, params service.RecordListParams) ([]*model.Record, uint32, error)
	ListenRecords(ctx context.Context, namespace string, resource string, consumer func(records []*model.Record)) error

	// resource
	GetResourceByName(ctx context.Context, namespace string, resourceName string) (*model.Resource, error)
	ListResources(ctx context.Context) ([]*model.Resource, error)
	ReadRecordStream(ctx context.Context, params service.RecordListParams, recordsChan chan *model.Record) error
	DeleteResource(ctx context.Context, id string, doMigration bool, forceMigration bool) error
	CreateResource(ctx context.Context, resource *model.Resource, migration bool, force bool) error
	UpdateResource(ctx context.Context, resource *model.Resource, migration bool, force bool) error

	Watch(ctx context.Context, request *stub.WatchRequest) (stub.Watch_WatchClient, error)
	PollEvents(ctx context.Context, key string) (<-chan *model.Event, error)
	WriteEvent(ctx context.Context, key string, event *model.Event) error
}

type Repository[Entity interface{}] interface {
	Mapper() abs.EntityMapper[Entity]
	Create(ctx context.Context, entity Entity) (Entity, error)
	Update(ctx context.Context, entity Entity) (Entity, error)
	Apply(ctx context.Context, entity Entity) (Entity, error)
	Get(ctx context.Context, id string) (Entity, error)
	Find(ctx context.Context, params FindParams) ([]Entity, uint32, error)
	Listen(ctx context.Context, consumer func(records []Entity)) error
}

type FindParams struct {
	Limit             uint32
	Offset            uint64
	UseHistory        bool
	Annotations       map[string]string
	ResolveReferences []string
	Filters           map[string]string
}
