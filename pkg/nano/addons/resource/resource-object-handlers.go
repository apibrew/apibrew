package resource

import (
	"context"
	"fmt"
	"github.com/apibrew/apibrew/pkg/errors"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/nano/abs"
	backend_event_handler "github.com/apibrew/apibrew/pkg/service/backend-event-handler"
	"github.com/apibrew/apibrew/pkg/util"
	"github.com/dop251/goja"
)

func (o *resourceObject) registerHandler(order int, sync bool, responds bool, action model.Event_Action) func(fn func(call goja.FunctionCall) goja.Value) {
	return func(fn func(call goja.FunctionCall) goja.Value) {
		handlerId := "nano-" + util.RandomHex(8)

		o.cec.AddHandlerId(handlerId)

		o.backendEventHandler.RegisterHandler(backend_event_handler.Handler{
			Id:       handlerId,
			Name:     handlerId,
			Fn:       o.recordHandlerFn(fn),
			Selector: o.handlerSelector(action),
			Order:    order,
			Sync:     sync,
			Responds: responds,
		})
	}
}

func (o *resourceObject) recordHandlerFn(fn func(call goja.FunctionCall) goja.Value) backend_event_handler.HandlerFunc {
	return func(ctx context.Context, event *model.Event) (processedEvent *model.Event, err errors.ServiceError) {
		defer func() {
			if r := recover(); r != nil {
				err = errors.RecordValidationError.WithDetails(fmt.Sprintf("%v", r))
			}
		}()

		for idx := range event.Records {
			record := event.Records[idx]

			entityValue := o.recordToValue(record)

			result := fn(goja.FunctionCall{
				Arguments: []goja.Value{
					entityValue,
					o.vm.ToValue(event),
				},
			})

			resultExported := result.Export()

			if resultExported != nil {
				updatedRecord, err := abs.ValueToRecord(o.resource, resultExported)
				if err != nil {
					return nil, err
				}

				event.Records[idx] = updatedRecord
			}
		}

		if len(event.Records) == 0 {
			fn(goja.FunctionCall{
				Arguments: []goja.Value{
					o.vm.ToValue(event),
				},
			})
		}

		return event, nil
	}
}

func (o *resourceObject) initHandlerMethods(object *goja.Object) {
	_ = object.Set("beforeCreate", o.registerHandler(10, true, true, model.Event_CREATE))
	_ = object.Set("beforeUpdate", o.registerHandler(10, true, true, model.Event_UPDATE))
	_ = object.Set("beforeDelete", o.registerHandler(10, true, true, model.Event_DELETE))
	_ = object.Set("beforeGet", o.registerHandler(10, true, true, model.Event_GET))
	_ = object.Set("beforeList", o.registerHandler(10, true, true, model.Event_LIST))

	_ = object.Set("afterCreate", o.registerHandler(110, true, true, model.Event_CREATE))
	_ = object.Set("afterUpdate", o.registerHandler(110, true, true, model.Event_UPDATE))
	_ = object.Set("afterDelete", o.registerHandler(110, true, true, model.Event_DELETE))
	_ = object.Set("afterGet", o.registerHandler(110, true, true, model.Event_GET))
	_ = object.Set("afterList", o.registerHandler(110, true, true, model.Event_LIST))

	_ = object.Set("onCreate", o.registerHandler(110, true, true, model.Event_CREATE))
	_ = object.Set("onUpdate", o.registerHandler(110, true, true, model.Event_UPDATE))
	_ = object.Set("onDelete", o.registerHandler(110, true, true, model.Event_DELETE))
	_ = object.Set("onGet", o.registerHandler(110, true, true, model.Event_GET))
	_ = object.Set("onList", o.registerHandler(110, true, true, model.Event_LIST))
}
