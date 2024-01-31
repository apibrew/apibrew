package resource

import (
	"github.com/apibrew/apibrew/pkg/api"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/nano/abs"
	"github.com/apibrew/apibrew/pkg/service"
	backend_event_handler "github.com/apibrew/apibrew/pkg/service/backend-event-handler"
	"github.com/apibrew/apibrew/pkg/util"
	"github.com/dop251/goja"
	"strings"
)

type resourceObject struct {
	goja.DynamicObject
	container service.Container
	resource  *model.Resource

	vm                  *goja.Runtime
	cec                 abs.CodeExecutionContext
	backendEventHandler backend_event_handler.BackendEventHandler
	global              abs.GlobalObject
	api                 api.Interface
}

func (o *resourceObject) handlerSelector(action model.Event_Action) *model.EventSelector {
	return &model.EventSelector{
		Actions:    []model.Event_Action{action},
		Namespaces: []string{o.resource.Namespace},
		Resources:  []string{o.resource.Name},
	}
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

	_ = object.Set("define", o.defineFn)
	_ = object.Set("call", o.callFn)
}

func (o *resourceObject) defineFn(name string, value interface{}) {
	o.global.Define(o.locateGlobalName(name), value)
}

func (o *resourceObject) callFn(name string, args ...goja.Value) goja.Value {
	fnObj := o.global.Get(o.locateGlobalName(name))

	if fnObj == nil {
		panic("definition not found: " + name)
	}

	if fn, ok := fnObj.(func(call goja.FunctionCall) goja.Value); ok {
		return fn(goja.FunctionCall{
			Arguments: args,
		})
	} else {
		panic("definition is not a function: " + name)
	}
}

func (o *resourceObject) locateGlobalName(name string) string {
	return o.resource.Namespace + "_" + o.resource.Name + "_" + name
}

func resourceFn(container service.Container, vm *goja.Runtime, cec abs.CodeExecutionContext, backendEventHandler backend_event_handler.BackendEventHandler, global abs.GlobalObject) func(args ...string) goja.Value {
	resourceService := container.GetResourceService()
	return func(args ...string) goja.Value {
		resource := ResourceByName(args, resourceService)

		ro := &resourceObject{
			resource:            resource,
			container:           container,
			api:                 api.NewInterface(container),
			vm:                  vm,
			cec:                 cec,
			backendEventHandler: backendEventHandler,
			global:              global,
		}

		value := vm.NewObject()

		ro.initValue(value)

		return value
	}
}

func ResourceByName(args []string, resourceService service.ResourceService) *model.Resource {
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
