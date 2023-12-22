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

	Preprocess  func(fn func(call goja.FunctionCall) goja.Value) `json:"preprocess"`
	Postprocess func(fn func(call goja.FunctionCall) goja.Value) `json:"postprocess"`
	Check       func(fn func(call goja.FunctionCall) goja.Value) `json:"check"`

	BindCreate BindFunc `jsbind:"bindCreate"`
	BindUpdate BindFunc `jsbind:"bindUpdate"`
	BindDelete BindFunc `jsbind:"bindDelete"`
	BindGet    BindFunc `jsbind:"bindGet"`
	BindList   BindFunc `jsbind:"bindList"`

	Create func(record goja.Value) goja.Value                  `js:"create"`
	Update func(record goja.Value) goja.Value                  `js:"update"`
	Apply  func(record goja.Value) goja.Value                  `js:"apply"`
	Delete func(record goja.Value) goja.Value                  `js:"delete"`
	Get    func(recordId string, params goja.Value) goja.Value `js:"get"`
	List   func(params goja.Value) goja.Value                  `js:"list"`

	Count    func(params goja.Value) goja.Value                  `js:"count"`
	Load     func(params goja.Value) goja.Value                  `js:"load"`
	FindById func(recordId string, params goja.Value) goja.Value `js:"findById"`

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

func (o *resourceObject) initHandlers() {
	o.initHandlerMethods()

	o.initBindMethods()
	o.initRepositoryMethods()
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
