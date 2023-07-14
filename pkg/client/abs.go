package client

import (
	"context"
	"github.com/apibrew/apibrew/pkg/abs"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/stub"
)

type ApplyInterface interface {
	ApplyRecord(ctx context.Context, resource *model.Resource, record *model.Record) error
	ApplyResource(ctx context.Context, resource *model.Resource, doMigration, forceMigration bool) error
}

type DhClient interface {
	ApplyInterface
	GetAuthenticationClient() stub.AuthenticationClient
	GetDataSourceClient() stub.DataSourceClient
	GetResourceClient() stub.ResourceClient
	GetRecordClient() stub.RecordClient
	GetGenericClient() stub.GenericClient
	GetToken() string
	AuthenticateWithToken(token string)
	AuthenticateWithUsernameAndPassword(username string, password string) error
	NewExtension(host string) Extension
}

type Repository[T abs.Entity[T]] interface {
	Create(ctx context.Context, entity T) (T, error)
	Update(ctx context.Context, entity T) (T, error)
	Save(ctx context.Context, entity T) (T, error)
	Get(ctx context.Context, id string) (T, error)
	Find(ctx context.Context, params FindParams) ([]T, error)
}

type FindParams struct {
	Limit             uint32
	Offset            uint64
	UseHistory        bool
	Annotations       map[string]string
	ResolveReferences []string // default []string{"*"}
	Query             *model.BooleanExpression
}
