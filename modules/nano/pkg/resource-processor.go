package nano

import (
	"context"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/service"
	backend_event_handler "github.com/apibrew/apibrew/pkg/service/backend-event-handler"
	"github.com/apibrew/apibrew/pkg/util"
	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/types/known/structpb"
)

type ResourceProcessor[T any] interface {
	MapperTo(record *model.Record) T

	Register(ctx context.Context, entity T) error
	Update(ctx context.Context, entity T) error
	UnRegister(ctx context.Context, entity T) error
}

func RegisterResourceProcessor[T any](handlerName string,
	processor ResourceProcessor[T],
	backendEventHandler backend_event_handler.BackendEventHandler,
	container service.Container,
	resource *model.Resource) error {
	handler := func(ctx context.Context, event *model.Event) (*model.Event, error) {
		for _, record := range event.Records {

			switch event.Action {
			case model.Event_CREATE:
				entity := processor.MapperTo(record)
				err := processor.Register(ctx, entity)

				if err != nil {
					return nil, err
				}
			case model.Event_UPDATE:
				existing, err := container.GetRecordService().Load(ctx, event.Resource.Namespace, event.Resource.Name, record.Properties, service.RecordLoadParams{})

				if err != nil {
					return nil, err
				}

				record = mergeRecords(existing, record)

				entity := processor.MapperTo(record)

				err = processor.Update(ctx, entity)

				if err != nil {
					return nil, err
				}
			case model.Event_DELETE:
				existing, err := container.GetRecordService().Load(ctx, event.Resource.Namespace, event.Resource.Name, record.Properties, service.RecordLoadParams{})

				if err != nil {
					return nil, err
				}

				record = mergeRecords(existing, record)

				entity := processor.MapperTo(record)

				err = processor.UnRegister(ctx, entity)

				if err != nil {
					return nil, err
				}
			}
		}

		return event, nil
	}

	backendEventHandler.RegisterHandler(backend_event_handler.Handler{
		Id:   handlerName,
		Name: handlerName,
		Fn:   handler,
		Selector: &model.EventSelector{
			Actions: []model.Event_Action{
				model.Event_CREATE, model.Event_UPDATE, model.Event_DELETE,
			},
			Namespaces: []string{resource.Namespace},
			Resources:  []string{resource.Name},
		},
		Order:    90,
		Sync:     true,
		Internal: true,
	})

	var codeRecords, _, err = container.GetRecordService().List(util.SystemContext, service.RecordListParams{
		Namespace: resource.Namespace,
		Resource:  resource.Name,
		Limit:     1000000,
	})

	if err != nil {
		return err
	}

	for _, record := range codeRecords {
		entity := processor.MapperTo(record)

		err := processor.Register(util.SystemContext, entity)

		if err != nil {
			logrus.Error(err)
		}
	}

	return nil
}

func mergeRecords(existing *model.Record, record *model.Record) *model.Record {
	var result = &model.Record{
		Properties: make(map[string]*structpb.Value),
	}

	for key, value := range existing.Properties {
		result.Properties[key] = value
	}

	for key, value := range record.Properties {
		result.Properties[key] = value
	}

	return result
}
