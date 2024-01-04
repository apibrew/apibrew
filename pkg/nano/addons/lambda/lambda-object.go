package lambda

import (
	"context"
	"fmt"
	"github.com/apibrew/apibrew/pkg/errors"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/nano/abs"
	resource2 "github.com/apibrew/apibrew/pkg/nano/addons/resource"
	"github.com/apibrew/apibrew/pkg/service"
	backend_event_handler "github.com/apibrew/apibrew/pkg/service/backend-event-handler"
	"github.com/apibrew/apibrew/pkg/service/validate"
	"github.com/apibrew/apibrew/pkg/util"
	"github.com/dop251/goja"
	log "github.com/sirupsen/logrus"
)

type lambdaObject struct {
	goja.DynamicObject
	container service.Container
	resource  *model.Resource

	Fire   func(event map[string]interface{})       `json:"fire"`
	Listen func(func(event map[string]interface{})) `json:"listen"`

	vm                  *goja.Runtime
	cec                 abs.CodeExecutionContext
	backendEventHandler backend_event_handler.BackendEventHandler
}

func (o *lambdaObject) handlerSelector(action model.Event_Action) *model.EventSelector {
	return &model.EventSelector{
		Actions:    []model.Event_Action{action},
		Namespaces: []string{o.resource.Namespace},
		Resources:  []string{o.resource.Name},
	}
}

func (o *lambdaObject) recordToObject(record *model.Record) map[string]interface{} {
	var recordObj = make(map[string]interface{})
	for key, value := range record.Properties {
		recordObj[key] = value.AsInterface()
	}
	return recordObj
}

func (o *lambdaObject) init() {
	o.Fire = o.fireFn
	o.Listen = o.listenFn
}

func (o *lambdaObject) fireFn(event map[string]interface{}) {
	record, err := abs.ValueToRecord(o.resource, event)

	if err != nil {
		panic(err)
	}

	err = validate.Records(o.resource, []*model.Record{
		record,
	}, false)

	if err != nil {
		panic(err)
	}

	go func() {

		_, err = o.container.GetRecordService().Create(util.SystemContext, service.RecordCreateParams{
			Namespace: o.resource.Namespace,
			Resource:  o.resource.Name,
			Records:   []*model.Record{record},
		})

		if err != nil {
			log.Error(err)
		}
	}()
}

func (o *lambdaObject) listenFn(f func(event map[string]interface{})) {
	handlerId := "nano-lambda-" + util.RandomHex(8)

	o.cec.AddHandlerId(handlerId)

	o.backendEventHandler.RegisterHandler(backend_event_handler.Handler{
		Id:       handlerId,
		Name:     handlerId,
		Fn:       o.recordHandlerFn(f),
		Selector: o.handlerSelector(model.Event_CREATE),
		Order:    99,
		Sync:     false,
		Responds: false,
	})
}

func (o *lambdaObject) recordHandlerFn(fn func(event map[string]interface{})) backend_event_handler.HandlerFunc {
	return func(ctx context.Context, event *model.Event) (processedEvent *model.Event, err errors.ServiceError) {
		defer func() {
			if r := recover(); r != nil {
				err = errors.RecordValidationError.WithDetails(fmt.Sprintf("%v", r))
			}
		}()

		for idx := range event.Records {
			record := event.Records[idx]

			entityValue := o.recordToObject(record)

			fn(entityValue)
		}
		return event, nil
	}
}

func lambdaFn(container service.Container, vm *goja.Runtime, cec abs.CodeExecutionContext, backendEventHandler backend_event_handler.BackendEventHandler) func(args ...string) goja.Value {
	resourceService := container.GetResourceService()
	return func(args ...string) goja.Value {
		resource := resource2.ResourceByName(args, resourceService)

		lo := &lambdaObject{resource: resource, container: container, vm: vm, cec: cec, backendEventHandler: backendEventHandler}

		lo.init()

		return vm.ToValue(lo)
	}
}
