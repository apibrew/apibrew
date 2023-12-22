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

func (o *resourceObject) countFn(params goja.Value) goja.Value {
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
	listParams.Limit = 1

	_, total, err := o.container.GetRecordService().List(util.SystemContext, listParams)

	if err != nil {
		panic(err)
	}

	return o.vm.ToValue(total)
}

func (o *resourceObject) initRepositoryMethods() {
	o.Create = o.createFn
	o.Update = o.updateFn
	o.Apply = o.applyFn
	o.Delete = o.deleteFn
	o.Get = o.getFn
	o.FindById = o.getFn
	o.List = o.listFn
	o.Load = o.loadFn
	o.Count = o.countFn
}
