package service

import (
	"context"
	"github.com/apibrew/apibrew/pkg/abs"
	"github.com/apibrew/apibrew/pkg/errors"
	"github.com/apibrew/apibrew/pkg/helper/protohelper"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/resources/mapping"
	"github.com/apibrew/apibrew/pkg/service/annotations"
	"github.com/apibrew/apibrew/pkg/service/security"
	log "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/proto"
)

type genericRecordService[T proto.Message] struct {
	recordService   abs.RecordService
	resourceService abs.ResourceService
	namespace       string
	resource        string
	protoHelper     protohelper.MappingHelper[T]
}

func (g genericRecordService[T]) Init(records []T) {
	if len(records) > 0 {
		_, err := g.recordService.Create(annotations.SetWithContext(security.SystemContext, annotations.IgnoreIfExists, annotations.Enabled), abs.RecordCreateParams{
			Namespace: g.namespace,
			Resource:  g.resource,
			Records:   mapping.MapToRecord(records, g.mapTo),
		})

		if err != nil {
			log.Fatal(err)
		}
	}
}

func (g genericRecordService[T]) Create(ctx context.Context, items []T) ([]T, errors.ServiceError) {
	// insert records via resource service
	records := mapping.MapToRecord(items, g.mapTo)

	result, err := g.recordService.Create(ctx, abs.RecordCreateParams{
		Namespace: g.namespace,
		Resource:  g.resource,
		Records:   records,
	})

	if err != nil {
		return nil, err
	}

	response := mapping.MapFromRecord(result, g.mapFrom)

	return response, nil
}

func (g genericRecordService[T]) Update(ctx context.Context, items []T) ([]T, errors.ServiceError) {
	// update records via resource service
	records := mapping.MapToRecord(items, g.mapTo)

	result, err := g.recordService.Update(ctx, abs.RecordUpdateParams{
		Namespace: g.namespace,
		Resource:  g.resource,
		Records:   records,
	})

	if err != nil {
		return nil, err
	}

	response := mapping.MapFromRecord(result, g.mapFrom)

	return response, nil
}

func (g genericRecordService[T]) Delete(ctx context.Context, ids []string) errors.ServiceError {
	return g.recordService.Delete(ctx, abs.RecordDeleteParams{
		Namespace: g.namespace,
		Resource:  g.resource,
		Ids:       ids,
	})
}

func (g genericRecordService[T]) Get(ctx context.Context, id string) (T, errors.ServiceError) {
	record, err := g.recordService.Get(ctx, abs.RecordGetParams{
		Namespace: g.namespace,
		Resource:  g.resource,
		Id:        id,
	})

	if err != nil {
		var nilMessage T
		return nilMessage, err
	}

	response := g.mapFrom(record)

	return response, nil
}

func (g genericRecordService[T]) List(ctx context.Context, query *model.BooleanExpression, limit uint32, offset uint64) ([]T, errors.ServiceError) {
	result, _, err := g.recordService.List(ctx, abs.RecordListParams{
		Query:     query,
		Namespace: g.namespace,
		Resource:  g.resource,
		Limit:     limit,
		Offset:    offset,
	})

	if err != nil {
		return nil, err
	}

	response := mapping.MapFromRecord(result, g.mapFrom)

	return response, nil
}

func (g genericRecordService[T]) mapTo(t T) *model.Record {
	return g.protoHelper.MapTo(t)
}

func (g genericRecordService[T]) mapFrom(record *model.Record) T {
	return g.protoHelper.MapFrom(record)
}

func NewGenericRecordService[T proto.Message](recordService abs.RecordService, resource *model.Resource, instance func() T) abs.GenericRecordService[T] {
	return &genericRecordService[T]{
		recordService: recordService,
		namespace:     resource.Namespace,
		resource:      resource.Name,
		protoHelper: protohelper.MappingHelper[T]{
			Resource: resource,
			Instance: instance,
		},
	}
}
