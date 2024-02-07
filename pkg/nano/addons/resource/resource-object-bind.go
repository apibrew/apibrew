package resource

import (
	"context"
	"fmt"
	"github.com/apibrew/apibrew/pkg/errors"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/nano/abs"
	"github.com/apibrew/apibrew/pkg/service"
	backend_event_handler "github.com/apibrew/apibrew/pkg/service/backend-event-handler"
	"github.com/apibrew/apibrew/pkg/util"
	"github.com/dop251/goja"
)

type BindFunc func(resourceValue *resourceObject, mapFrom func(call goja.FunctionCall) goja.Value, mapTo func(call goja.FunctionCall) goja.Value)

func (o *resourceObject) registerBindHandler(action model.Event_Action) interface{} {
	return func(boundResource map[string]interface{}, mapFrom interface{}, mapTo interface{}) {
		handlerId := "nano-" + util.RandomHex(8)

		o.cec.AddHandlerId(handlerId)

		o.backendEventHandler.RegisterHandler(backend_event_handler.Handler{
			Id:        handlerId,
			Name:      handlerId,
			Fn:        o.bindRecordFn(boundResource["self"].(*resourceObject).resource, mapFrom.(func(call goja.FunctionCall) goja.Value), mapTo.(func(call goja.FunctionCall) goja.Value)),
			Selector:  o.handlerSelector(action),
			Order:     90,
			Sync:      true,
			Responds:  true,
			Finalizes: true,
		})
	}
}

func (o *resourceObject) bindRecordFn(boundResource *model.Resource, from func(call goja.FunctionCall) goja.Value, to func(call goja.FunctionCall) goja.Value) backend_event_handler.HandlerFunc {
	srv := o.container.GetRecordService()

	mapper := func(to func(call goja.FunctionCall) goja.Value, records []*model.Record) ([]*model.Record, errors.ServiceError) {
		return util.ArrayMapWithError(records, func(record *model.Record) (*model.Record, errors.ServiceError) {
			value := o.recordToValue(record)

			mapped := to(goja.FunctionCall{Arguments: []goja.Value{value}})

			return abs.ValueToRecord(o.resource, mapped.Export())
		})
	}

	return func(ctx context.Context, event *model.Event) (processedEvent *model.Event, err errors.ServiceError) {
		defer func() {
			if r := recover(); r != nil {
				err = errors.RecordValidationError.WithDetails(fmt.Sprintf("%v", r))
			}
		}()

		toMappedRecords, err := mapper(to, event.Records)

		if err != nil {
			return nil, err
		}

		var resultRecords []*model.Record

		switch event.Action {
		case model.Event_CREATE:
			result, err := srv.Create(util.SystemContext, service.RecordCreateParams{
				Namespace: boundResource.Namespace,
				Resource:  boundResource.Name,
				Records:   toMappedRecords,
			})

			if err != nil {
				return nil, err
			}

			resultRecords = result
		case model.Event_UPDATE:
			result, err := srv.Update(util.SystemContext, service.RecordUpdateParams{
				Namespace: boundResource.Namespace,
				Resource:  boundResource.Name,
				Records:   toMappedRecords,
			})

			if err != nil {
				return nil, err
			}

			resultRecords = result
		case model.Event_DELETE:
			err := srv.Delete(util.SystemContext, service.RecordDeleteParams{
				Namespace: boundResource.Namespace,
				Resource:  boundResource.Name,
				Ids:       util.ArrayMap(event.Records, util.GetRecordId),
			})

			if err != nil {
				return nil, err
			}
		case model.Event_GET:
			record, err := srv.Get(util.SystemContext, service.RecordGetParams{
				Namespace: boundResource.Namespace,
				Resource:  boundResource.Name,
				Id:        util.GetRecordId(event.Records[0]),
			})

			if err != nil {
				return nil, err
			}

			resultRecords = []*model.Record{record}
		case model.Event_LIST:
			list, total, err := srv.List(util.SystemContext, service.RecordListParams{
				Namespace:         boundResource.Namespace,
				Resource:          boundResource.Name,
				ResolveReferences: event.RecordSearchParams.ResolveReferences,
				Query:             event.RecordSearchParams.Query,
			})

			if err != nil {
				return nil, err
			}

			resultRecords = list
			event.Total = uint64(total)
		}

		fromMappedRecords, err := mapper(from, resultRecords)

		if err != nil {
			return nil, err
		}

		event.Records = fromMappedRecords

		return event, nil
	}
}

func (o *resourceObject) initBindMethods(object *goja.Object) {
	_ = object.Set("bindCreate", o.registerBindHandler(model.Event_CREATE))
	_ = object.Set("bindUpdate", o.registerBindHandler(model.Event_UPDATE))
	_ = object.Set("bindDelete", o.registerBindHandler(model.Event_DELETE))
	_ = object.Set("bindGet", o.registerBindHandler(model.Event_GET))
	_ = object.Set("bindList", o.registerBindHandler(model.Event_LIST))
}
