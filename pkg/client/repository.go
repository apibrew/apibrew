package client

import (
	"context"
	"github.com/tislib/data-handler/pkg/model"
	"github.com/tislib/data-handler/pkg/stub"
	"github.com/tislib/data-handler/pkg/util"
)

type repository[T Entity[T]] struct {
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
	resp, err := r.client.GetRecordClient().Update(ctx, &stub.UpdateRecordRequest{
		Token:        r.client.GetToken(),
		Namespace:    entity.GetNamespace(),
		Resource:     entity.GetResourceName(),
		Record:       entity.ToRecord(),
		CheckVersion: r.params.UpdateCheckVersion,
	})

	if err != nil {
		return entity, err
	}

	entity.FromRecord(resp.Record)

	return entity, nil
}

func (r repository[T]) Save(ctx context.Context, entity T) (T, error) {
	if entity.GetId() != "" {
		return r.Update(ctx, entity)
	} else {
		return r.Create(ctx, entity)
	}
}

func (r repository[T]) Get(ctx context.Context, id string) (T, error) {
	resp, err := r.client.GetRecordClient().Get(ctx, &stub.GetRecordRequest{
		Token:     r.client.GetToken(),
		Namespace: r.params.Instance.GetNamespace(),
		Resource:  r.params.Instance.GetResourceName(),
		Id:        id,
	})

	if err != nil {
		return r.params.Instance, err
	}

	var newInstance = r.params.Instance.Clone()

	newInstance.FromRecord(resp.Record)

	return newInstance, nil
}

func (r repository[T]) List(ctx context.Context) ([]T, error) {
	resp, err := r.client.GetRecordClient().List(ctx, &stub.ListRecordRequest{
		Token:             r.client.GetToken(),
		Namespace:         r.params.Instance.GetNamespace(),
		Resource:          r.params.Instance.GetResourceName(),
		ResolveReferences: []string{"*"},
	})

	if err != nil {
		return []T{}, err
	}

	return util.ArrayMap(resp.Content, func(record *model.Record) T {
		var newInstance = r.params.Instance.Clone()

		newInstance.FromRecord(record)

		return newInstance
	}), nil
}

type RepositoryParams[T Entity[T]] struct {
	UpdateCheckVersion bool
	Instance           T
}

func NewRepository[T Entity[T]](client DhClient, params RepositoryParams[T]) Repository[T] {
	return repository[T]{client: client, params: params}
}
