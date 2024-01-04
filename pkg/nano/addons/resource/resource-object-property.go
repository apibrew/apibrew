package resource

import (
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/dop251/goja"
	log "github.com/sirupsen/logrus"
)

func (o *resourceObject) initPropertyMethods(object *goja.Object) {
	for _, prop := range o.resource.Properties {
		_ = object.Set(prop.Name, o.propertyObject(prop))
	}
}

func (o *resourceObject) aggFnForProp(property string, algorithm model.AggregationItem_Algorithm) func(filters map[string]string) goja.Value {
	return func(filters map[string]string) goja.Value {
		return o.simpleAggregateFn(property, filters, algorithm)
	}
}

func (o *resourceObject) computeFn(prop *model.ResourceProperty) func(fn func(call goja.FunctionCall) goja.Value, dependencies []interface{}) {
	return func(fn func(call goja.FunctionCall) goja.Value, dependencies []interface{}) {
		log.Print("compute called")

		for _, dependency := range dependencies {
			// locating dependency resource

			if itemMap, ok := dependency.(map[string]interface{}); ok {
				var depResourceObject = itemMap["resourceObject"].(*resourceObject)
				var depProp = itemMap["property"].(*model.ResourceProperty)
				var isSelf = depResourceObject.resource.Id == o.resource.Id

				if !isSelf && (depProp.Type != model.ResourceProperty_REFERENCE || depProp.Reference.Resource != o.resource.Name) {
					panic("Dependency must be a property which is pointing to computing resource")
				}

				o.registerDepHandlers(fn, depResourceObject, depProp, prop)
			} else {
				panic("Unknown compute dependency type, dependencies must be either string or resourceObject")
			}
		}

		if len(dependencies) == 0 {
			panic("Compute with no dependency")
		}
	}
}

func (o *resourceObject) registerDepHandlers(fn func(call goja.FunctionCall) goja.Value, dep *resourceObject, depProp *model.ResourceProperty, prop *model.ResourceProperty) {
	var isSelf = dep.resource.Id == o.resource.Id

	if isSelf {
		dep.registerHandler(99, true, true, model.Event_CREATE)(o.recordComputeHandlerFnForDep(fn, depProp, prop, isSelf))
		dep.registerHandler(99, true, true, model.Event_UPDATE)(o.recordComputeHandlerFnForDep(fn, depProp, prop, isSelf))
		dep.registerHandler(99, true, true, model.Event_DELETE)(o.recordComputeHandlerFnForDep(fn, depProp, prop, isSelf))
	} else {
		dep.registerHandler(101, true, true, model.Event_CREATE)(o.recordComputeHandlerFnForDep(fn, depProp, prop, isSelf))
		dep.registerHandler(101, true, true, model.Event_UPDATE)(o.recordComputeHandlerFnForDep(fn, depProp, prop, isSelf))
		dep.registerHandler(101, true, true, model.Event_DELETE)(o.recordComputeHandlerFnForDep(fn, depProp, prop, isSelf))
	}
}

func (o *resourceObject) recordComputeHandlerFnForDep(fn func(call goja.FunctionCall) goja.Value, depProp *model.ResourceProperty, prop *model.ResourceProperty, isSelf bool) func(call goja.FunctionCall) goja.Value {
	return func(call goja.FunctionCall) goja.Value {
		depEntity := call.Arguments[0].Export().(map[string]interface{})

		depPropValue := depEntity[depProp.Name]

		if depPropValue != nil {
			// locating referenced item
			if isSelf {
				depEntity[prop.Name] = fn(call).Export()

				return o.vm.ToValue(depEntity)
			} else {
				entityValue := o.loadFn(o.vm.ToValue(depPropValue))
				entity := make(map[string]interface{})

				entity["id"] = entityValue.Export().(map[string]interface{})["id"]
				entity[prop.Name] = fn(goja.FunctionCall{Arguments: []goja.Value{entityValue}}).Export()

				o.updateFn(o.vm.ToValue(entity))
			}
		}

		return call.Arguments[0]
	}
}

func (o *resourceObject) propertyObject(prop *model.ResourceProperty) goja.Value {
	object := o.vm.NewObject()

	_ = object.Set("resourceObject", o)
	_ = object.Set("property", prop)
	_ = object.Set("sum", o.aggFnForProp(prop.Name, model.AggregationItem_SUM))
	_ = object.Set("max", o.aggFnForProp(prop.Name, model.AggregationItem_MAX))
	_ = object.Set("min", o.aggFnForProp(prop.Name, model.AggregationItem_MIN))
	_ = object.Set("avg", o.aggFnForProp(prop.Name, model.AggregationItem_AVG))
	_ = object.Set("compute", o.computeFn(prop))

	return object
}
