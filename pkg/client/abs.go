package client

import (
	"context"
	"github.com/tislib/data-handler/pkg/model"
	"github.com/tislib/data-handler/pkg/stub"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/structpb"
)

type ApplyInterface interface {
	Apply(ctx context.Context, msg proto.Message) error
	ApplyRecord(ctx context.Context, resource *model.Resource, record *model.Record) error
	ApplyNamespace(ctx context.Context, namespace *model.Namespace) error
	ApplyExtension(ctx context.Context, extension *model.Extension) error
	ApplyUser(ctx context.Context, user *model.User) error
	ApplyDataSource(ctx context.Context, dataSource *model.DataSource) error
	ApplyResource(ctx context.Context, resource *model.Resource, doMigration, forceMigration bool) error
}

type DhClient interface {
	ApplyInterface
	GetAuthenticationClient() stub.AuthenticationClient
	GetDataSourceClient() stub.DataSourceClient
	GetResourceClient() stub.ResourceClient
	GetRecordClient() stub.RecordClient
	GetGenericClient() stub.GenericClient
	GetNamespaceClient() stub.NamespaceClient
	GetExtensionClient() stub.ExtensionClient
	GetUserClient() stub.UserClient
	GetToken() string
	AuthenticateWithToken(token string)
	AuthenticateWithUsernameAndPassword(username string, password string) error
	NewExtension(host string) Extension
}

type Entity[T any] interface {
	ToRecord() *model.Record
	FromRecord(record *model.Record)
	FromProperties(properties map[string]*structpb.Value)
	ToProperties() map[string]*structpb.Value
	GetResourceName() string
	GetNamespace() string
	Equals(other T) bool
	Same(other T) bool
}

type Repository[T Entity[T]] interface {
	Create(ctx context.Context, entity T) (T, error)
	Update(ctx context.Context, entity T) (T, error)
	Save(ctx context.Context, entity T) (T, error)
	Get(ctx context.Context, id string) (T, error)
	Find(ctx context.Context, params FindParams) ([]T, error)
	Extend(extension Extension) RepositoryExtension[T]
}

type FindParams struct {
	Limit             uint32
	Offset            uint64
	UseHistory        bool
	Annotations       map[string]string
	ResolveReferences []string // default []string{"*"}
	Query             *model.BooleanExpression
}
