package nano

import (
	"context"
	"fmt"
	"github.com/apibrew/apibrew/pkg/errors"
	"github.com/apibrew/apibrew/pkg/model"
	backend_event_handler "github.com/apibrew/apibrew/pkg/service/backend-event-handler"
	"github.com/apibrew/apibrew/pkg/util"
	"github.com/dop251/goja"
)

func (o *resourceObject) registerHandler(order int, sync bool, responds bool, action model.Event_Action) func(fn func(call goja.FunctionCall) goja.Value) {
	return func(fn func(call goja.FunctionCall) goja.Value) {
		handlerId := "nano-" + util.RandomHex(8)

		o.cec.handlerIds = append(o.cec.handlerIds, handlerId)

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
				err = errors.LogicalError.WithDetails(fmt.Sprintf("%v", r))
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
				updatedRecord, err := o.valueToRecord(resultExported)
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

func (o *resourceObject) initHandlerMethods() {
	o.BeforeCreate = o.registerHandler(10, true, true, model.Event_CREATE)
	o.BeforeUpdate = o.registerHandler(10, true, true, model.Event_UPDATE)
	o.BeforeDelete = o.registerHandler(10, true, true, model.Event_DELETE)
	o.BeforeGet = o.registerHandler(10, true, true, model.Event_GET)
	o.BeforeList = o.registerHandler(10, true, true, model.Event_LIST)

	o.AfterCreate = o.registerHandler(110, true, true, model.Event_CREATE)
	o.AfterUpdate = o.registerHandler(110, true, true, model.Event_UPDATE)
	o.AfterDelete = o.registerHandler(110, true, true, model.Event_DELETE)
	o.AfterGet = o.registerHandler(110, true, true, model.Event_GET)
	o.AfterList = o.registerHandler(110, true, true, model.Event_LIST)

	o.OnCreate = o.registerHandler(110, true, true, model.Event_CREATE)
	o.OnUpdate = o.registerHandler(110, true, true, model.Event_UPDATE)
	o.OnDelete = o.registerHandler(110, true, true, model.Event_DELETE)
	o.OnGet = o.registerHandler(110, true, true, model.Event_GET)
	o.OnList = o.registerHandler(110, true, true, model.Event_LIST)
}
