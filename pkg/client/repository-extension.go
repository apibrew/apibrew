package client

import (
	"context"
	"github.com/tislib/data-handler/pkg/model"
)

type repositoryExtension[T Entity[T]] struct {
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
		Name:      extensionName,
		Namespace: r.namespace,
		Resource:  r.resourceName,
		Instead: &model.Extension_Instead{
			Create: &model.ExternalCall{
				Kind: &model.ExternalCall_FunctionCall{
					FunctionCall: &model.FunctionCall{
						Host:         r.extension.GetRemoteHost(),
						FunctionName: extensionName,
					},
				},
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

type RepositoryExtension[T Entity[T]] interface {
	OnCreate(handler func(ctx context.Context, elem T) (T, error)) error
	OnUpdate(handler func(ctx context.Context, elem T) (T, error))
	OnDelete(handler func(ctx context.Context, elem T) (T, error))
	OnGet(handler func(ctx context.Context, id string) (T, error))
	OnList(handler func(ctx context.Context) (T, error))
}
