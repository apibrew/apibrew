package resource

import (
	"context"
	"fmt"
	"github.com/apibrew/apibrew/pkg/errors"
	"github.com/apibrew/apibrew/pkg/formats/unstructured"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/nano/abs"
	"github.com/apibrew/apibrew/pkg/resource_model"
	"github.com/apibrew/apibrew/pkg/resource_model/extramappings"
	backend_event_handler "github.com/apibrew/apibrew/pkg/service/backend-event-handler"
	"github.com/apibrew/apibrew/pkg/util"
	"runtime/debug"
)

type HandlerFunc func(entity map[string]interface{}, event *resource_model.Event) interface{}

type Handler struct {
	Name      string
	Fn        HandlerFunc
	Selector  resource_model.EventSelector
	Order     int
	Finalizes bool
	Sync      bool
	Responds  bool
}

func handle(cec abs.CodeExecutionContext, backendEventHandler backend_event_handler.BackendEventHandler) func(handler Handler) {
	return func(handler Handler) {
		handlerId := "nano-" + cec.GetCodeIdentifier() + "-" + util.RandomHex(8)

		cec.AddHandlerId(handlerId)

		backendEventHandler.RegisterHandler(backend_event_handler.Handler{
			Id:        handlerId,
			Name:      cec.GetCodeIdentifier() + "-" + handler.Name,
			Fn:        recordHandlerFn(handler),
			Selector:  extramappings.EventSelectorToProto(handler.Selector),
			Order:     handler.Order,
			Sync:      handler.Sync,
			Responds:  handler.Responds,
			Finalizes: handler.Finalizes,
		})
	}
}

func recordHandlerFn(handler Handler) backend_event_handler.HandlerFunc {
	fn := handler.Fn
	return func(ctx context.Context, event *model.Event) (processedEvent *model.Event, err errors.ServiceError) {
		defer func() {
			if r := recover(); r != nil {
				debug.Stack()
				err = errors.RecordValidationError.WithDetails(fmt.Sprintf("%v", r))
			}
		}()
		e := extramappings.EventFromProto(event)

		if len(event.Records) == 0 {
			result := fn(nil, e)

			if result != nil {
				resultObj := result.(map[string]interface{})

				if resultObj["total"] != nil {
					var total = resultObj["total"].(int64)
					e.Total = &total
				}

				if resultObj["content"] != nil {
					var records []*resource_model.Record
					for _, record := range resultObj["content"].([]interface{}) {
						records = append(records, &resource_model.Record{
							Properties: record.(unstructured.Unstructured)["properties"].(map[string]interface{}),
						})
					}
					e.Records = records
				}

				return extramappings.EventToProto(e), nil
			}

			return event, nil
		}

		var processedRecords []*model.Record
		for _, record := range event.Records {
			entity := recordToObject(record)

			result := fn(entity, e)

			if result == false {
				continue
			}

			if result != nil {
				updatedRecord, err := abs.ValueToRecord(event.Resource, result)
				if err != nil {
					return nil, err
				}

				processedRecords = append(processedRecords, updatedRecord)
			} else {
				processedRecords = append(processedRecords, record)
			}
		}

		if len(processedRecords) == 0 {
			return nil, nil
		}

		event.Records = processedRecords

		return event, nil
	}
}

func recordToObject(record *model.Record) map[string]interface{} {
	var recordObj = make(map[string]interface{})
	for key, value := range record.Properties {
		recordObj[key] = value.AsInterface()
	}
	return recordObj
}
