package nano

import (
	"context"
	"fmt"
	"github.com/apibrew/apibrew/pkg/errors"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/service"
	backend_event_handler "github.com/apibrew/apibrew/pkg/service/backend-event-handler"
	"github.com/apibrew/apibrew/pkg/util"
	"github.com/dop251/goja"
	"google.golang.org/protobuf/types/known/structpb"
)

type BindFunc func(resourceValue *resourceObject, mapFrom func(call goja.FunctionCall) goja.Value, mapTo func(call goja.FunctionCall) goja.Value)

type resourceObject struct {
	container service.Container
	resource  *model.Resource

	BeforeCreate func(fn func(call goja.FunctionCall) goja.Value) `json:"beforeCreate"`
	BeforeUpdate func(fn func(call goja.FunctionCall) goja.Value) `json:"beforeUpdate"`
	BeforeDelete func(fn func(call goja.FunctionCall) goja.Value) `json:"beforeDelete"`
	BeforeGet    func(fn func(call goja.FunctionCall) goja.Value) `json:"beforeGet"`
	BeforeList   func(fn func(call goja.FunctionCall) goja.Value) `json:"beforeList"`

	AfterCreate func(fn func(call goja.FunctionCall) goja.Value) `json:"afterCreate"`
	AfterUpdate func(fn func(call goja.FunctionCall) goja.Value) `json:"afterUpdate"`
	AfterDelete func(fn func(call goja.FunctionCall) goja.Value) `json:"afterDelete"`
	AfterGet    func(fn func(call goja.FunctionCall) goja.Value) `json:"afterGet"`
	AfterList   func(fn func(call goja.FunctionCall) goja.Value) `json:"afterList"`

	OnCreate func(fn func(call goja.FunctionCall) goja.Value) `json:"onCreate"`
	OnUpdate func(fn func(call goja.FunctionCall) goja.Value) `json:"onUpdate"`
	OnDelete func(fn func(call goja.FunctionCall) goja.Value) `json:"onDelete"`
	OnGet    func(fn func(call goja.FunctionCall) goja.Value) `json:"onGet"`
	OnList   func(fn func(call goja.FunctionCall) goja.Value) `json:"onList"`

	BindCreate BindFunc `jsbind:"bindCreate"`
	BindUpdate BindFunc `jsbind:"bindUpdate"`
	BindDelete BindFunc `jsbind:"bindDelete"`
	BindGet    BindFunc `jsbind:"bindGet"`
	BindList   BindFunc `jsbind:"bindList"`

	vm                  *goja.Runtime
	cec                 *codeExecutionContext
	backendEventHandler backend_event_handler.BackendEventHandler
}

func (o *resourceObject) registerBindHandler(action model.Event_Action) BindFunc {
	return func(boundResource *resourceObject, mapFrom func(call goja.FunctionCall) goja.Value, mapTo func(call goja.FunctionCall) goja.Value) {
		handlerId := "nano-" + util.RandomHex(8)

		o.cec.handlerIds = append(o.cec.handlerIds, handlerId)

		o.backendEventHandler.RegisterHandler(backend_event_handler.Handler{
			Id:        handlerId,
			Name:      handlerId,
			Fn:        o.bindRecordFn(boundResource.resource, mapFrom, mapTo),
			Selector:  o.handlerSelector(action),
			Order:     90,
			Sync:      true,
			Responds:  true,
			Finalizes: true,
		})
	}
}

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

func (o *resourceObject) handlerSelector(action model.Event_Action) *model.EventSelector {
	return &model.EventSelector{
		Actions:    []model.Event_Action{action},
		Namespaces: []string{o.resource.Namespace},
		Resources:  []string{o.resource.Name},
	}
}

func (o *resourceObject) recordHandlerFn(fn func(call goja.FunctionCall) goja.Value) backend_event_handler.HandlerFunc {
	return func(ctx context.Context, event *model.Event) (*model.Event, errors.ServiceError) {
		for idx := range event.Records {
			record := event.Records[idx]

			entityValue := o.recordToValue(record)

			result := fn(goja.FunctionCall{
				Arguments: []goja.Value{
					entityValue,
					o.vm.ToValue(event),
				},
			})

			if result != nil {
				updatedRecord, err := o.valueToRecord(result)
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

func (o *resourceObject) valueToRecord(result goja.Value) (*model.Record, errors.ServiceError) {
	recordObj, ok := result.Export().(map[string]interface{})

	if !ok {
		return nil, errors.LogicalError.WithDetails(fmt.Sprintf("Cannot accept nano function result: %v", result))
	}

	var record = new(model.Record)
	record.Properties = make(map[string]*structpb.Value)

	for key, value := range recordObj {
		sv, verr := structpb.NewValue(value)

		if verr != nil {
			return nil, errors.LogicalError.WithDetails(verr.Error())
		}

		record.Properties[key] = sv
	}
	return record, nil
}

func (o *resourceObject) recordToValue(record *model.Record) goja.Value {
	var recordObj = make(map[string]interface{})
	for key, value := range record.Properties {
		recordObj[key] = value.AsInterface()
	}
	return o.vm.ToValue(recordObj)
}

func (o *resourceObject) bindRecordFn(boundResource *model.Resource, from func(call goja.FunctionCall) goja.Value, to func(call goja.FunctionCall) goja.Value) backend_event_handler.HandlerFunc {
	srv := o.container.GetRecordService()

	mapper := func(to func(call goja.FunctionCall) goja.Value, records []*model.Record) ([]*model.Record, errors.ServiceError) {
		return util.ArrayMapWithError(records, func(record *model.Record) (*model.Record, errors.ServiceError) {
			value := o.recordToValue(record)

			mapped := to(goja.FunctionCall{Arguments: []goja.Value{value}})

			return o.valueToRecord(mapped)
		})
	}

	return func(ctx context.Context, event *model.Event) (*model.Event, errors.ServiceError) {
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

func (o *resourceObject) initHandlers() {
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

	o.BindCreate = o.registerBindHandler(model.Event_CREATE)
	o.BindUpdate = o.registerBindHandler(model.Event_UPDATE)
	o.BindDelete = o.registerBindHandler(model.Event_DELETE)
	o.BindGet = o.registerBindHandler(model.Event_GET)
	o.BindList = o.registerBindHandler(model.Event_LIST)
}

func resourceFn(container service.Container, vm *goja.Runtime, cec *codeExecutionContext, backendEventHandler backend_event_handler.BackendEventHandler) func(args ...string) *resourceObject {
	resourceService := container.GetResourceService()
	return func(args ...string) *resourceObject {
		var resourceName string
		var namespace string

		if len(args) == 0 || len(args) > 2 {
			panic("resource function needs 1 or 2 parameters")
		}

		if len(args) == 1 {
			namespace = "default"
			resourceName = args[0]
		} else {
			namespace = args[0]
			resourceName = args[1]
		}

		resource, err := resourceService.GetResourceByName(util.SystemContext, namespace, resourceName)

		if err != nil {
			panic(err)
		}

		return newResourceObject(resource, container, vm, cec, backendEventHandler)
	}
}

func newResourceObject(resource *model.Resource, container service.Container, vm *goja.Runtime, cec *codeExecutionContext, backendEventHandler backend_event_handler.BackendEventHandler) *resourceObject {
	ro := &resourceObject{resource: resource, container: container, vm: vm, cec: cec, backendEventHandler: backendEventHandler}

	ro.initHandlers()

	return ro
}
