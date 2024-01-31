package resource

import (
	"github.com/apibrew/apibrew/pkg/api"
	"github.com/apibrew/apibrew/pkg/formats/unstructured"
	"github.com/apibrew/apibrew/pkg/util"
	"github.com/dop251/goja"
)

func (o *resourceObject) createFn(record unstructured.Unstructured) unstructured.Unstructured {
	record["type"] = o.resource.Namespace + "/" + o.resource.Name

	result, err := o.api.Create(util.SystemContext, record)

	if err != nil {
		panic(err)
	}

	return result
}

func (o *resourceObject) updateFn(record unstructured.Unstructured) unstructured.Unstructured {
	record["type"] = o.resource.Namespace + "/" + o.resource.Name

	result, err := o.api.Update(util.SystemContext, record)

	if err != nil {
		panic(err)
	}

	return result
}

func (o *resourceObject) applyFn(record unstructured.Unstructured) unstructured.Unstructured {
	record["type"] = o.resource.Namespace + "/" + o.resource.Name

	result, err := o.api.Apply(util.SystemContext, record)

	if err != nil {
		panic(err)
	}

	return result
}

func (o *resourceObject) deleteFn(record unstructured.Unstructured) {
	record["type"] = o.resource.Namespace + "/" + o.resource.Name

	err := o.api.Delete(util.SystemContext, record)

	if err != nil {
		panic(err)
	}
}

func (o *resourceObject) getFn(recordId string, params api.LoadParams) unstructured.Unstructured {
	record, err := o.api.Load(util.SystemContext, map[string]interface{}{
		"type": o.resource.Namespace + "/" + o.resource.Name,
		"id":   recordId,
	}, params)

	if err != nil {
		panic(err)
	}

	return record
}

func (o *resourceObject) listFn(params api.ListParams) api.RecordListResult {
	params.Type = o.resource.Namespace + "/" + o.resource.Name

	result, err := o.api.List(util.SystemContext, params)

	if err != nil {
		panic(err)
	}

	return result
}

func (o *resourceObject) loadFn(record unstructured.Unstructured, params api.LoadParams) unstructured.Unstructured {
	record["type"] = o.resource.Namespace + "/" + o.resource.Name
	record, err := o.api.Load(util.SystemContext, record, params)

	if err != nil {
		panic(err)
	}

	return record
}
func (o *resourceObject) countFn(params api.ListParams) uint32 {
	params.Type = o.resource.Namespace + "/" + o.resource.Name

	params.Limit = 0

	result, err := o.api.List(util.SystemContext, params)

	if err != nil {
		panic(err)
	}

	return result.Total
}

func (o *resourceObject) simpleAggregateFn(property string, filters map[string]string, algorithm api.AggregationAlgorithm) unstructured.Any {
	var params = api.ListParams{}
	params.Filters = filters

	params.Aggregation = &api.Aggregation{
		Items: []api.AggregationItem{
			{
				Name:      "result",
				Algorithm: algorithm,
				Property:  property,
			},
		},
	}

	params.Type = o.resource.Namespace + "/" + o.resource.Name

	result, err := o.api.List(util.SystemContext, params)

	if err != nil {
		panic(err)
	}

	if result.Total == 0 {
		return o.vm.ToValue(0)
	}

	return result.Content[0][property]
}

func (o *resourceObject) sumFn(property string, filters map[string]string) unstructured.Any {
	return o.simpleAggregateFn(property, filters, api.Sum)
}

func (o *resourceObject) maxFn(property string, filters map[string]string) unstructured.Any {
	return o.simpleAggregateFn(property, filters, api.Max)
}

func (o *resourceObject) minFn(property string, filters map[string]string) unstructured.Any {
	return o.simpleAggregateFn(property, filters, api.Min)
}

func (o *resourceObject) avgFn(property string, filters map[string]string) unstructured.Any {
	return o.simpleAggregateFn(property, filters, api.Avg)
}

func (o *resourceObject) initRepositoryMethods(object *goja.Object) {
	_ = object.Set("create", o.createFn)
	_ = object.Set("update", o.updateFn)
	_ = object.Set("apply", o.applyFn)
	_ = object.Set("delete", o.deleteFn)
	_ = object.Set("get", o.getFn)
	_ = object.Set("findById", o.getFn)
	_ = object.Set("list", o.listFn)
	_ = object.Set("load", o.loadFn)
	_ = object.Set("count", o.countFn)
	_ = object.Set("sum", o.sumFn)
	_ = object.Set("max", o.maxFn)
	_ = object.Set("min", o.minFn)
	_ = object.Set("avg", o.avgFn)
}
