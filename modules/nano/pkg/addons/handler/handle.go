package handler

import (
	"context"
	"fmt"
	"github.com/apibrew/apibrew/modules/nano/pkg/abs"
	util2 "github.com/apibrew/apibrew/modules/nano/pkg/addons/util"
	"github.com/apibrew/apibrew/pkg/errors"
	"github.com/apibrew/apibrew/pkg/formats/unstructured"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/resource_model"
	"github.com/apibrew/apibrew/pkg/resource_model/extramappings"
	backend_event_handler "github.com/apibrew/apibrew/pkg/service/backend-event-handler"
	"github.com/apibrew/apibrew/pkg/util"
	"github.com/dop251/goja"
	"github.com/hashicorp/go-metrics"
	log "github.com/sirupsen/logrus"
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

func handle(vm *goja.Runtime, cec abs.CodeExecutionContext, backendEventHandler backend_event_handler.BackendEventHandler) func(handler Handler) {
	return func(handler Handler) {
		metrics.IncrCounterWithLabels([]string{"NanoMetrics"}, float32(1), []metrics.Label{
			{Name: "type", Value: "registerHandler"},
			{Name: "name", Value: handler.Name},
			{Name: "cecId", Value: cec.GetCodeIdentifier()},
		})

		if cec.IsScriptMode() {
			util2.ThrowError(vm, "Handlers are not supported in script mode")
		}

		handlerData := cec.HandlerMap().Get(handler.Name)

		if handlerData == nil {
			handlerData = &abs.HandlerData{}
			handlerData.Ch = make(chan *abs.EventWithContext, 100)
			cec.HandlerMap().Set(handler.Name, handlerData)

			handlerId := "nano-" + cec.GetCodeIdentifier() + "-" + util.RandomHex(8)

			var handlerTemplate = backend_event_handler.Handler{
				Id:        handlerId,
				Name:      cec.GetCodeIdentifier() + "-" + handler.Name,
				Selector:  extramappings.EventSelectorToProto(handler.Selector),
				Order:     handler.Order,
				Sync:      handler.Sync,
				Responds:  handler.Responds,
				Finalizes: handler.Finalizes,
			}

			handlerTemplate.Id = handlerId
			handlerData.Id = handlerId
			handlerTemplate.Fn = processThrowHandlerData(cec, handler, handlerData)

			backendEventHandler.RegisterHandler(handlerTemplate)

			go func() {
				<-cec.CodeContext().Done()
				log.Println("Closing Handler: ", handlerData.Id)
				backendEventHandler.UnRegisterHandler(handlerTemplate)
				handlerData.Closed = true
				close(handlerData.Ch)
				log.Println("Close Handler: ", handlerData.Id)
			}()
		}

		go func() {
			for item := range handlerData.Ch {
				processedEvent, err := recordHandlerFn(cec, handler.Fn)(item.Ctx, item.Event)

				item.Signal <- abs.EventWithContextSignal{
					ProcessedEvent: processedEvent,
					Err:            err,
				}
			}
		}()
	}
}

func recordHandlerFn(cec abs.CodeExecutionContext, fn HandlerFunc) backend_event_handler.HandlerFunc {
	return func(ctx context.Context, event *model.Event) (processedEvent *model.Event, err error) {
		cleanUpContext := cec.WithContext(util.WithSystemContext(ctx))
		defer cleanUpContext()

		defer func() {
			if r := recover(); r != nil {
				debug.PrintStack()
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

func processThrowHandlerData(cec abs.CodeExecutionContext, handler Handler, data *abs.HandlerData) backend_event_handler.HandlerFunc {
	return func(ctx context.Context, event *model.Event) (*model.Event, error) {
		metrics.IncrCounterWithLabels([]string{"NanoMetrics"}, float32(1), []metrics.Label{
			{Name: "type", Value: "handleEvent"},
			{Name: "name", Value: handler.Name},
			{Name: "cecId", Value: cec.GetCodeIdentifier()},
		})

		log.Debug("Begin dispatching event: " + event.Id)
		ec := &abs.EventWithContext{
			Ctx:    ctx,
			Event:  event,
			Signal: make(chan abs.EventWithContextSignal),
		}

		if data.Closed {
			return nil, errors.RecordValidationError.WithDetails("Handler is closed: " + data.Id)
		}
		data.Ch <- ec

		log.Debug("Starting to wait for signal: " + event.Id)
		res := <-ec.Signal

		log.Debug("Received signal: " + event.Id)
		return res.ProcessedEvent, res.Err
	}
}

func recordToObject(record *model.Record) map[string]interface{} {
	var recordObj = make(map[string]interface{})
	for key, value := range record.Properties {
		recordObj[key] = value.AsInterface()
	}
	return recordObj
}
