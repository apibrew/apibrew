package api

import (
	"context"
	"github.com/apibrew/apibrew/pkg/abs"
	"github.com/apibrew/apibrew/pkg/errors"
	"github.com/apibrew/apibrew/pkg/formats/unstructured"
	"github.com/apibrew/apibrew/pkg/resource_model"
)

type Entity interface {
}

type Mapper[T Entity] interface {
	ToRecord(user T) abs.RecordLike
	FromRecord(record abs.RecordLike) T
}

type EntityListResult[T Entity] struct {
	Total   uint32 `json:"total"`
	Content []T    `json:"content"`
}

type Repository[T Entity] interface {
	Create(ctx context.Context, entity T) (T, error)
	Update(ctx context.Context, entity T) (T, error)
	Apply(ctx context.Context, entity T) (T, error)
	Load(ctx context.Context, entity T, params LoadParams) (T, error)
	Delete(ctx context.Context, entity T) error
	List(ctx context.Context, params ListParams) (EntityListResult[T], error)
	GetResourceByType(ctx context.Context, typeName string) (*resource_model.Resource, error)
}

func (r repository[T]) Create(ctx context.Context, entity T) (T, error) {
	result, err := r.api.Create(ctx, r.mapper.ToRecord(entity).MapCopy())

	if err != nil {
		return r.entityDefault, err
	}

	record, err2 := unstructured.ToRecord(result)

	if err2 != nil {
		return r.entityDefault, errors.RecordValidationError.WithMessage(err2.Error())
	}

	return r.mapper.FromRecord(record), nil
}

func (r repository[T]) Update(ctx context.Context, entity T) (T, error) {
	result, err := r.api.Update(ctx, r.mapper.ToRecord(entity).MapCopy())

	if err != nil {
		return r.entityDefault, err
	}

	record, err2 := unstructured.ToRecord(result)

	if err2 != nil {
		return r.entityDefault, errors.RecordValidationError.WithMessage(err2.Error())
	}

	return r.mapper.FromRecord(record), nil
}

func (r repository[T]) Apply(ctx context.Context, entity T) (T, error) {
	result, err := r.api.Apply(ctx, r.mapper.ToRecord(entity).MapCopy())

	if err != nil {
		return r.entityDefault, err
	}

	record, err2 := unstructured.ToRecord(result)

	if err2 != nil {
		return r.entityDefault, errors.RecordValidationError.WithMessage(err2.Error())
	}

	return r.mapper.FromRecord(record), nil
}

func (r repository[T]) Load(ctx context.Context, entity T, params LoadParams) (T, error) {
	result, err := r.api.Load(ctx, r.mapper.ToRecord(entity).MapCopy(), params)

	if err != nil {
		return r.entityDefault, err
	}

	record, err2 := unstructured.ToRecord(result)

	if err2 != nil {
		return r.entityDefault, errors.RecordValidationError.WithMessage(err2.Error())
	}

	return r.mapper.FromRecord(record), nil
}

func (r repository[T]) Delete(ctx context.Context, entity T) error {
	return r.api.Delete(ctx, r.mapper.ToRecord(entity).MapCopy())
}

func (r repository[T]) List(ctx context.Context, params ListParams) (EntityListResult[T], error) {
	result, err := r.api.List(ctx, params)

	if err != nil {
		return EntityListResult[T]{}, err
	}

	records := make([]T, 0)

	for _, obj := range result.Content {
		record, err2 := unstructured.ToRecord(obj)

		if err2 != nil {
			return EntityListResult[T]{}, errors.RecordValidationError.WithMessage(err2.Error())
		}

		entity := r.mapper.FromRecord(record)
		records = append(records, entity)
	}

	return EntityListResult[T]{Total: result.Total, Content: records}, nil
}

func (r repository[T]) GetResourceByType(ctx context.Context, typeName string) (*resource_model.Resource, error) {
	return r.api.GetResourceByType(ctx, typeName)
}

type repository[T Entity] struct {
	api           Interface
	mapper        Mapper[T]
	entityDefault T
}

func NewRepository[T Entity](api Interface, mapper Mapper[T]) Repository[T] {
	return &repository[T]{api: api, mapper: mapper}
}
