package client

import (
	"context"
	"github.com/apibrew/apibrew/pkg/abs"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/service/annotations"
	"github.com/apibrew/apibrew/pkg/service/security"
	"github.com/apibrew/apibrew/pkg/stub"
	"github.com/apibrew/apibrew/pkg/util"
)

type repository[T abs.Entity[T]] struct {
	client DhClient
	params RepositoryParams[T]
}

func (r repository[T]) Create(ctx context.Context, entity T) (T, error) {
	resp, err := r.client.GetRecordClient().Create(ctx, &stub.CreateRecordRequest{
		Token:     r.client.GetToken(),
		Namespace: entity.GetNamespace(),
		Resource:  entity.GetResourceName(),
		Record:    entity.ToRecord(),
	})

	if err != nil {
		return entity, err
	}

	entity.FromRecord(resp.Record)

	return entity, nil
}

func (r repository[T]) Update(ctx context.Context, entity T) (T, error) {
	resp, err := r.client.GetRecordClient().Update(annotations.SetWithContext(security.SystemContext, annotations.CheckVersion, annotations.Enabled), &stub.UpdateRecordRequest{
		Token:     r.client.GetToken(),
		Namespace: entity.GetNamespace(),
		Resource:  entity.GetResourceName(),
		Record:    entity.ToRecord(),
	})

	if err != nil {
		return entity, err
	}

	entity.FromRecord(resp.Record)

	return entity, nil
}

func (r repository[T]) Save(ctx context.Context, entity T) (T, error) {
	resource, err := r.loadResource(ctx)

	if err != nil {
		return entity, err
	}

	return entity, r.client.ApplyRecord(ctx, resource, entity.ToRecord())
}

func (r repository[T]) Get(ctx context.Context, id string) (T, error) {
	instance := r.params.InstanceProvider()

	resp, err := r.client.GetRecordClient().Get(ctx, &stub.GetRecordRequest{
		Token:     r.client.GetToken(),
		Namespace: instance.GetNamespace(),
		Resource:  instance.GetResourceName(),
		Id:        id,
	})

	if err != nil {
		return instance, err
	}

	instance.FromRecord(resp.Record)

	return instance, nil
}

func (r repository[T]) Find(ctx context.Context, params FindParams) ([]T, error) {
	instance := r.params.InstanceProvider()

	if params.ResolveReferences == nil {
		params.ResolveReferences = []string{"*"}
	}

	resp, err := r.client.GetRecordClient().Search(ctx, &stub.SearchRecordRequest{
		Token:             r.client.GetToken(),
		Namespace:         instance.GetNamespace(),
		Resource:          instance.GetResourceName(),
		Query:             params.Query,
		Limit:             params.Limit,
		Offset:            params.Offset,
		UseHistory:        params.UseHistory,
		ResolveReferences: params.ResolveReferences,
		Annotations:       params.Annotations,
	})

	if err != nil {
		return []T{}, err
	}

	return util.ArrayMap(resp.Content, func(record *model.Record) T {
		var newInstance = r.params.InstanceProvider()

		newInstance.FromRecord(record)

		return newInstance
	}), nil
}

func (r repository[T]) Extend(extension Extension) RepositoryExtension[T] {
	instance := r.params.InstanceProvider()
	return &repositoryExtension[T]{client: r.client, instanceProvider: r.params.InstanceProvider, repository: r, extension: extension, resourceName: instance.GetResourceName(), namespace: instance.GetNamespace()}
}

func (r repository[T]) loadResource(ctx context.Context) (*model.Resource, error) {
	instance := r.params.InstanceProvider()

	resp, err := r.client.GetResourceClient().GetByName(ctx, &stub.GetResourceByNameRequest{
		Token:     r.client.GetToken(),
		Namespace: instance.GetNamespace(),
		Name:      instance.GetResourceName(),
	})

	if err != nil {
		return nil, err
	}

	return resp.Resource, nil
}

type RepositoryParams[T abs.Entity[T]] struct {
	UpdateCheckVersion bool
	InstanceProvider   func() T
}

func NewRepository[T abs.Entity[T]](client DhClient, params RepositoryParams[T]) Repository[T] {
	return repository[T]{client: client, params: params}
}
