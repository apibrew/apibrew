package client

import (
	"context"
	"github.com/apibrew/apibrew/pkg/abs"
	"github.com/apibrew/apibrew/pkg/model"
)

type repositoryExtension[T abs.Entity[T]] struct {
	repository       Repository[T]
	extension        Extension
	resourceName     string
	namespace        string
	instanceProvider func() T
	client           DhClient
}

func (r repositoryExtension[T]) OnCreate(handler func(ctx context.Context, elem T) (T, error)) error {
	extensionName := r.getExtensionName("OnCreate")

	r.extension.RegisterFunction(extensionName, CreateRecordTypedFunction(r.instanceProvider, handler))

	ext := &model.Extension{
		Name: extensionName,
		Selector: &model.EventSelector{
			Namespaces: []string{r.namespace},
			Resources:  []string{r.resourceName},
		},
		Call: &model.ExternalCall{
			FunctionCall: &model.FunctionCall{
				Host:         r.extension.GetRemoteHost(),
				FunctionName: extensionName,
			},
		},
	}

	return r.client.ApplyExtension(context.TODO(), ext)
}

func (r repositoryExtension[T]) getExtensionName(action string) string {
	return r.namespace + "-" + r.resourceName + "-" + action
}

func (r repositoryExtension[T]) OnUpdate(handler func(ctx context.Context, elem T) (T, error)) {
	//TODO implement me
	panic("implement me")
}

func (r repositoryExtension[T]) OnDelete(handler func(ctx context.Context, elem T) (T, error)) {
	//TODO implement me
	panic("implement me")
}

func (r repositoryExtension[T]) OnGet(handler func(ctx context.Context, id string) (T, error)) {
	//TODO implement me
	panic("implement me")
}

func (r repositoryExtension[T]) OnList(handler func(ctx context.Context) (T, error)) {
	//TODO implement me
	panic("implement me")
}

type RepositoryExtension[T abs.Entity[T]] interface {
	OnCreate(handler func(ctx context.Context, elem T) (T, error)) error
	OnUpdate(handler func(ctx context.Context, elem T) (T, error))
	OnDelete(handler func(ctx context.Context, elem T) (T, error))
	OnGet(handler func(ctx context.Context, id string) (T, error))
	OnList(handler func(ctx context.Context) (T, error))
}
