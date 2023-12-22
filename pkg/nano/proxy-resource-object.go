package nano

import (
	"context"
	"github.com/apibrew/apibrew/pkg/errors"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/service"
	backend_event_handler "github.com/apibrew/apibrew/pkg/service/backend-event-handler"
	"github.com/apibrew/apibrew/pkg/util"
	"github.com/dop251/goja"
	log "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/types/known/structpb"
)

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

	vm                  *goja.Runtime
	cec                 *codeExecutionContext
	backendEventHandler backend_event_handler.BackendEventHandler
}

func (o resourceObject) registerHandler(order int, sync bool, responds bool, action model.Event_Action) func(fn func(call goja.FunctionCall) goja.Value) {
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

func (o resourceObject) handlerSelector(action model.Event_Action) *model.EventSelector {
	return &model.EventSelector{
		Actions:    []model.Event_Action{action},
		Namespaces: []string{o.resource.Namespace},
		Resources:  []string{o.resource.Name},
	}
}

func (o resourceObject) recordHandlerFn(fn func(call goja.FunctionCall) goja.Value) backend_event_handler.HandlerFunc {
	return func(ctx context.Context, event *model.Event) (*model.Event, errors.ServiceError) {
		for idx := range event.Records {
			var recordObj = make(map[string]interface{})

			record := event.Records[idx]

			for key, value := range record.Properties {
				recordObj[key] = value.AsInterface()
			}

			result := fn(goja.FunctionCall{
				Arguments: []goja.Value{
					o.vm.ToValue(recordObj),
					o.vm.ToValue(event),
				},
			})

			if result != nil {
				resultObj, ok := result.Export().(map[string]interface{})

				if ok {
					recordObj = resultObj
				} else {
					log.Warn("Cannot accept nano function result: ", result)
				}
			}

			for key, value := range recordObj {
				sv, verr := structpb.NewValue(value)

				if verr != nil {
					return nil, errors.LogicalError.WithDetails(verr.Error())
				}

				record.Properties[key] = sv
			}

			event.Records[idx] = record
		}

		if len(event.Records) == 0 {

		}

		return event, nil
	}
}

func (o resourceObject) initHandlers() {
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

func resourceFn(container service.Container, vm *goja.Runtime, cec *codeExecutionContext, backendEventHandler backend_event_handler.BackendEventHandler) func(args ...string) *resourceObject {
	resourceService := container.GetResourceService()
	return func(args ...string) *resourceObject {
		var resourceName string
		var namespace string

		if len(args) == 0 || len(args) > 2 {
			panic("resource function needs 1 or 2 parameters")
		}

		resourceName = args[0]
		if len(args) == 1 {
			namespace = "default"
		} else {
			namespace = args[1]
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
