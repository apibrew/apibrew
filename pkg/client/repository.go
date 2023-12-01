package client

import (
	"context"
	"github.com/apibrew/apibrew/pkg/abs"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/service"
)

type repository[Entity interface{}] struct {
	client             Client
	UpdateCheckVersion bool
	mapper             abs.EntityMapper[Entity]
}

func (r repository[T]) Create(ctx context.Context, entity T) (T, error) {
	resp, err := r.client.CreateRecord(ctx, r.mapper.ResourceIdentity().Namespace, r.mapper.ResourceIdentity().Name, r.mapper.ToRecord(entity))

	if err != nil {
		return entity, err
	}

	updatedEntity := r.mapper.FromRecord(resp)

	return updatedEntity, nil
}

func (r repository[T]) Update(ctx context.Context, entity T) (T, error) {
	resp, err := r.client.UpdateRecord(ctx, r.mapper.ResourceIdentity().Namespace, r.mapper.ResourceIdentity().Name, r.mapper.ToRecord(entity))

	if err != nil {
		return entity, err
	}

	updatedEntity := r.mapper.FromRecord(resp)

	return updatedEntity, nil
}

func (r repository[T]) Apply(ctx context.Context, entity T) (T, error) {
	resp, err := r.client.ApplyRecord(ctx, r.mapper.ResourceIdentity().Namespace, r.mapper.ResourceIdentity().Name, r.mapper.ToRecord(entity))

	if err != nil {
		return entity, err
	}

	updatedEntity := r.mapper.FromRecord(resp)

	return updatedEntity, nil
}

func (r repository[T]) Get(ctx context.Context, id string) (T, error) {
	resp, err := r.client.GetRecord(ctx, r.mapper.ResourceIdentity().Namespace, r.mapper.ResourceIdentity().Name, id)

	if err != nil {
		return r.mapper.New(), err
	}

	updatedEntity := r.mapper.FromRecord(resp)

	return updatedEntity, nil
}

func (r repository[T]) Find(ctx context.Context, params FindParams) ([]T, uint32, error) {
	listParams := service.RecordListParams{
		Namespace:         r.mapper.ResourceIdentity().Namespace,
		Resource:          r.mapper.ResourceIdentity().Name,
		Filters:           params.Filters,
		Limit:             params.Limit,
		Offset:            params.Offset,
		UseHistory:        params.UseHistory,
		ResolveReferences: params.ResolveReferences,
	}

	resp, total, err := r.client.ListRecords(ctx, listParams)

	if err != nil {
		return nil, 0, err
	}

	var result = make([]T, len(resp))

	for i, record := range resp {
		result[i] = r.mapper.FromRecord(record)
	}

	return result, total, nil
}

func (r repository[T]) Listen(ctx context.Context, consumer func(records []T)) error {
	return r.client.ListenRecords(ctx, r.mapper.ResourceIdentity().Namespace, r.mapper.ResourceIdentity().Name, func(records []*model.Record) {
		var result = make([]T, len(records))

		for i, record := range records {
			result[i] = r.mapper.FromRecord(record)
		}

		consumer(result)
	})
}

func (r repository[T]) Mapper() abs.EntityMapper[T] {
	return r.mapper
}

func R[Entity interface{}](client Client, mapper abs.EntityMapper[Entity]) Repository[Entity] {
	return NewRepository(client, mapper)
}

func NewRepository[Entity interface{}](client Client, mapper abs.EntityMapper[Entity]) Repository[Entity] {
	return repository[Entity]{client: client, mapper: mapper}
}
