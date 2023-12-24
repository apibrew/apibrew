package nano

import (
	"encoding/json"
	"fmt"
	"github.com/apibrew/apibrew/pkg/errors"
	"github.com/apibrew/apibrew/pkg/helper"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/service"
	"github.com/apibrew/apibrew/pkg/util"
	"github.com/dop251/goja"
)

func (o *resourceObject) createFn(recordValue goja.Value) goja.Value {
	record, err := o.valueToRecord(recordValue.Export())

	if err != nil {
		panic(err)
	}

	result, err := o.container.GetRecordService().Create(util.SystemContext, service.RecordCreateParams{
		Namespace: o.resource.Namespace,
		Resource:  o.resource.Name,
		Records:   []*model.Record{record},
	})

	if err != nil {
		panic(err)
	}

	return o.recordToValue(result[0])
}

func (o *resourceObject) updateFn(recordValue goja.Value) goja.Value {
	record, err := o.valueToRecord(recordValue.Export())

	if err != nil {
		panic(err)
	}

	result, err := o.container.GetRecordService().Update(util.SystemContext, service.RecordUpdateParams{
		Namespace: o.resource.Namespace,
		Resource:  o.resource.Name,
		Records:   []*model.Record{record},
	})

	if err != nil {
		panic(err)
	}

	return o.recordToValue(result[0])
}

func (o *resourceObject) applyFn(recordValue goja.Value) goja.Value {
	record, err := o.valueToRecord(recordValue.Export())

	if err != nil {
		panic(err)
	}

	result, err := o.container.GetRecordService().Apply(util.SystemContext, service.RecordUpdateParams{
		Namespace: o.resource.Namespace,
		Resource:  o.resource.Name,
		Records:   []*model.Record{record},
	})

	if err != nil {
		panic(err)
	}

	return o.recordToValue(result[0])
}

func (o *resourceObject) deleteFn(recordValue goja.Value) goja.Value {
	record, err := o.valueToRecord(recordValue.Export())

	if err != nil {
		panic(err)
	}

	var id = util.GetRecordId(record)
	if id == "" {
		return o.deleteFn(o.loadFn(recordValue))
	}

	err = o.container.GetRecordService().Delete(util.SystemContext, service.RecordDeleteParams{
		Namespace: o.resource.Namespace,
		Resource:  o.resource.Name,
		Ids:       []string{id},
	})

	if err != nil {
		panic(err)
	}

	return goja.Undefined()
}

func (o *resourceObject) getFn(recordId string, params goja.Value) goja.Value {
	var resolveReferences []string

	if params.Export() != nil {
		paramsMap, ok := params.Export().(map[string]interface{})

		if !ok {
			panic(fmt.Sprintf("Params must be an object: %v", params.Export()))
		}

		if paramsMap["resolveReferences"] != nil {
			resolveReferences = paramsMap["resolveReferences"].([]string)
		}
	}

	record, err := o.container.GetRecordService().Get(util.SystemContext, service.RecordGetParams{
		Namespace:         o.resource.Namespace,
		Resource:          o.resource.Name,
		Id:                recordId,
		ResolveReferences: resolveReferences,
	})

	if err != nil {
		panic(err)
	}

	return o.recordToValue(record)
}

func (o *resourceObject) listFn(params goja.Value) goja.Value {
	var listParams = service.RecordListParams{}
	if params != nil {
		paramsStr, err := json.Marshal(params.Export())

		if err != nil {
			panic(err)
		}

		err = json.Unmarshal(paramsStr, &listParams)

		if err != nil {
			panic(err)
		}
	}

	listParams.Namespace = o.resource.Namespace
	listParams.Resource = o.resource.Name

	content, total, err := o.container.GetRecordService().List(util.SystemContext, listParams)

	if err != nil {
		panic(err)
	}

	var resultMap = make(map[string]interface{})

	resultMap["content"] = util.ArrayMap(content, o.recordToObject)
	resultMap["total"] = total
	resultMap["total"] = total

	return o.vm.ToValue(resultMap)
}

func (o *resourceObject) loadFn(params goja.Value) goja.Value {
	recordToLoad, err := o.valueToRecord(params.Export())

	if err != nil {
		panic(err)
	}

	identifierProps, rerr := util.RecordIdentifierProperties(o.resource, recordToLoad.Properties)

	if rerr != nil {
		panic(rerr)
	}

	qb := helper.NewQueryBuilder()

	searchRes, total, serr := o.container.GetRecordService().List(util.SystemContext, service.RecordListParams{
		Namespace: o.resource.Namespace,
		Resource:  o.resource.Name,
		Limit:     1,
		Query:     qb.FromProperties(o.resource, identifierProps),
	})

	if err != nil {
		panic(serr)
	}

	if total == 0 {
		panic(errors.LogicalError.WithDetails(fmt.Sprintf("Record not found with params: %v", identifierProps)))
	}

	return o.recordToValue(searchRes[0])
}
func (o *resourceObject) countFn(filters map[string]string) goja.Value {
	var listParams = service.RecordListParams{}

	listParams.Namespace = o.resource.Namespace
	listParams.Resource = o.resource.Name
	listParams.Filters = filters
	listParams.Limit = 1

	_, total, err := o.container.GetRecordService().List(util.SystemContext, listParams)

	if err != nil {
		panic(err)
	}

	return o.vm.ToValue(total)
}

func (o *resourceObject) simpleAggregateFn(property string, filters map[string]string, algorithm model.AggregationItem_Algorithm) goja.Value {
	var listParams = service.RecordListParams{}

	listParams.Namespace = o.resource.Namespace
	listParams.Resource = o.resource.Name
	listParams.Limit = 1
	listParams.Filters = filters
	listParams.Aggregation = &model.Aggregation{
		Items: []*model.AggregationItem{
			{
				Name:      "result",
				Algorithm: algorithm,
				Property:  property,
			},
		},
	}

	result, _, err := o.container.GetRecordService().List(util.SystemContext, listParams)

	if err != nil {
		panic(err)
	}

	if len(result) == 0 {
		return o.vm.ToValue(0)
	}

	return o.vm.ToValue(result[0].Properties["result"].AsInterface())
}

func (o *resourceObject) sumFn(property string, filters map[string]string) goja.Value {
	return o.simpleAggregateFn(property, filters, model.AggregationItem_SUM)
}

func (o *resourceObject) maxFn(property string, filters map[string]string) goja.Value {
	return o.simpleAggregateFn(property, filters, model.AggregationItem_MAX)
}

func (o *resourceObject) minFn(property string, filters map[string]string) goja.Value {
	return o.simpleAggregateFn(property, filters, model.AggregationItem_MIN)
}

func (o *resourceObject) avgFn(property string, filters map[string]string) goja.Value {
	return o.simpleAggregateFn(property, filters, model.AggregationItem_AVG)
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
