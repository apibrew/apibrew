package nano

import (
	"fmt"
	"github.com/apibrew/apibrew/pkg/errors"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/service"
	backend_event_handler "github.com/apibrew/apibrew/pkg/service/backend-event-handler"
	"github.com/apibrew/apibrew/pkg/util"
	"github.com/dop251/goja"
	"google.golang.org/protobuf/types/known/structpb"
	"strings"
)

type resourceObject struct {
	goja.DynamicObject
	container service.Container
	resource  *model.Resource

	vm                  *goja.Runtime
	cec                 *codeExecutionContext
	backendEventHandler backend_event_handler.BackendEventHandler
}

func (o *resourceObject) handlerSelector(action model.Event_Action) *model.EventSelector {
	return &model.EventSelector{
		Actions:    []model.Event_Action{action},
		Namespaces: []string{o.resource.Namespace},
		Resources:  []string{o.resource.Name},
	}
}

func (o *resourceObject) valueToRecord(resultExported interface{}) (*model.Record, errors.ServiceError) {
	recordObj, ok := resultExported.(map[string]interface{})

	if !ok {
		return nil, errors.LogicalError.WithDetails(fmt.Sprintf("Cannot accept nano function result: %v", resultExported))
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
	return o.vm.ToValue(o.recordToObject(record))
}

func (o *resourceObject) recordToObject(record *model.Record) map[string]interface{} {
	var recordObj = make(map[string]interface{})
	for key, value := range record.Properties {
		recordObj[key] = value.AsInterface()
	}
	return recordObj
}

func (o *resourceObject) initValue(object *goja.Object) {
	_ = object.Set("self", o)
	o.initHandlerMethods(object)

	o.initBindMethods(object)
	o.initRepositoryMethods(object)

	o.initPropertyMethods(object)
}

func resourceFn(container service.Container, vm *goja.Runtime, cec *codeExecutionContext, backendEventHandler backend_event_handler.BackendEventHandler) func(args ...string) goja.Value {
	resourceService := container.GetResourceService()
	return func(args ...string) goja.Value {
		resource := resourceByName(args, resourceService)

		ro := &resourceObject{resource: resource, container: container, vm: vm, cec: cec, backendEventHandler: backendEventHandler}

		value := vm.NewObject()

		ro.initValue(value)

		return value
	}
}

func resourceByName(args []string, resourceService service.ResourceService) *model.Resource {
	var resourceName string
	var namespace string

	if len(args) == 0 || len(args) > 2 {
		panic("resource function needs 1 or 2 parameters")
	}

	if strings.Contains(args[0], "/") {
		args = strings.Split(args[0], "/")
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
	return resource
}
