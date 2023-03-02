package client

import (
	"context"
	"github.com/tislib/data-handler/pkg/model"
	"github.com/tislib/data-handler/pkg/stub"
	"google.golang.org/protobuf/types/known/structpb"
)

type DhClient interface {
	GetAuthenticationServiceClient() stub.AuthenticationServiceClient
	GetDataSourceServiceClient() stub.DataSourceServiceClient
	GetResourceServiceClient() stub.ResourceServiceClient
	GetRecordServiceClient() stub.RecordServiceClient
	GetGenericServiceClient() stub.GenericServiceClient
	GetNamespaceServiceClient() stub.NamespaceServiceClient
	GetExtensionServiceClient() stub.ExtensionServiceClient
	GetUserServiceClient() stub.UserServiceClient
	GetToken() string
	AuthenticateWithToken(token string)
	AuthenticateWithUsernameAndPassword(username string, password string) error
}

type Entity interface {
	ToRecord() *model.Record
	FromRecord(record *model.Record)
	FromProperties(properties map[string]*structpb.Value)
	ToProperties(includeTopProperties bool) map[string]*structpb.Value
	GetResourceName() string
	GetNamespace() string
	GetId() string
}

type Repository[T Entity] interface {
	Create(ctx context.Context, entity T) (T, error)
	Update(ctx context.Context, entity T) (T, error)
	Save(ctx context.Context, entity T) (T, error)
	Get(ctx context.Context, id string) (T, error)
	List(ctx context.Context) ([]T, error)
}
