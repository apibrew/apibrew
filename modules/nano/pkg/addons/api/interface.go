package api

import (
	"github.com/apibrew/apibrew/modules/nano/pkg/abs"
	util2 "github.com/apibrew/apibrew/modules/nano/pkg/addons/util"
	"github.com/apibrew/apibrew/pkg/api"
	"github.com/apibrew/apibrew/pkg/formats/unstructured"
	"github.com/apibrew/apibrew/pkg/resource_model"
	"github.com/apibrew/apibrew/pkg/util"
	"github.com/dop251/goja"
)

func create(cec abs.CodeExecutionContext, vm *goja.Runtime, apiInterface api.Interface) func(record unstructured.Unstructured) unstructured.Unstructured {
	return func(record unstructured.Unstructured) unstructured.Unstructured {
		var typ = record["type"].(string)

		result, err := apiInterface.Create(cec.LocalContext(), record)

		if err != nil {
			util2.ThrowError(vm, err.Error())
		}

		if cec.TransactionalEnabled() {
			cec.RegisterRevert(func() error {
				result["type"] = typ
				return apiInterface.Delete(cec.LocalContext(), result)
			})
		}

		return result
	}
}

func update(cec abs.CodeExecutionContext, vm *goja.Runtime, apiInterface api.Interface) func(record unstructured.Unstructured) unstructured.Unstructured {
	return func(record unstructured.Unstructured) unstructured.Unstructured {
		var typ = record["type"].(string)
		var existingRecord unstructured.Unstructured

		if cec.TransactionalEnabled() {
			var err error
			existingRecord, err = apiInterface.Load(cec.LocalContext(), record, api.LoadParams{})

			if err != nil {
				util2.ThrowError(vm, err.Error())
			}
		}

		record["type"] = typ
		result, err := apiInterface.Update(cec.LocalContext(), record)

		if err != nil {
			util2.ThrowError(vm, err.Error())
		}

		if cec.TransactionalEnabled() {
			cec.RegisterRevert(func() error {
				existingRecord["type"] = typ
				_, err := apiInterface.Update(cec.LocalContext(), existingRecord)

				return err
			})
		}

		return result
	}
}

func apply(cec abs.CodeExecutionContext, vm *goja.Runtime, apiInterface api.Interface) func(record unstructured.Unstructured) unstructured.Unstructured {
	return func(record unstructured.Unstructured) unstructured.Unstructured {
		var typ = record["type"].(string)
		var existingRecord unstructured.Unstructured = nil

		if cec.TransactionalEnabled() {
			existingRecord, _ = apiInterface.Load(cec.LocalContext(), record, api.LoadParams{})
		}

		result, err := apiInterface.Apply(cec.LocalContext(), record)

		if err != nil {
			util2.ThrowError(vm, err.Error())
		}

		if cec.TransactionalEnabled() {
			cec.RegisterRevert(func() error {
				if existingRecord == nil {
					result["type"] = typ
					return apiInterface.Delete(cec.LocalContext(), result)
				} else {
					existingRecord["type"] = typ
					_, err := apiInterface.Update(cec.LocalContext(), existingRecord)
					return err
				}
			})
		}

		return result
	}
}

func delete_(cec abs.CodeExecutionContext, vm *goja.Runtime, apiInterface api.Interface) func(record unstructured.Unstructured) {
	return func(record unstructured.Unstructured) {
		var typ = record["type"].(string)
		var existingRecord unstructured.Unstructured

		if cec.TransactionalEnabled() {
			var err error
			existingRecord, err = apiInterface.Load(cec.LocalContext(), record, api.LoadParams{})

			if err != nil {
				util2.ThrowError(vm, err.Error())
			}
		}

		record["type"] = typ
		err := apiInterface.Delete(cec.LocalContext(), record)

		if err != nil {
			util2.ThrowError(vm, err.Error())
		}

		if cec.TransactionalEnabled() {
			cec.RegisterRevert(func() error {
				existingRecord["type"] = typ
				_, err := apiInterface.Create(cec.LocalContext(), existingRecord)

				return err
			})

		}
	}
}

func load(cec abs.CodeExecutionContext, vm *goja.Runtime, apiInterface api.Interface) func(record unstructured.Unstructured, params api.LoadParams) unstructured.Unstructured {
	return func(record unstructured.Unstructured, params api.LoadParams) unstructured.Unstructured {
		result, err := apiInterface.Load(cec.LocalContext(), record, params)

		if err != nil {
			util2.ThrowError(vm, err.Error())
		}

		return result
	}
}

func list(cec abs.CodeExecutionContext, vm *goja.Runtime, apiInterface api.Interface) func(params api.ListParams) unstructured.Unstructured {
	return func(params api.ListParams) unstructured.Unstructured {
		result, err := apiInterface.List(cec.LocalContext(), params)

		if err != nil {
			util2.ThrowError(vm, err.Error())
		}

		return unstructured.Unstructured{
			"content": util.ArrayMap(result.Content, func(item unstructured.Unstructured) interface{} {
				return item
			}),
			"total": result.Total,
		}
	}
}

func begin(cec abs.CodeExecutionContext, vm *goja.Runtime, apiInterface api.Interface) func() {
	return func() {
		if err := cec.BeginTransaction(); err != nil {
			util2.ThrowError(vm, err.Error())
		}
	}
}

func commit(cec abs.CodeExecutionContext, vm *goja.Runtime, apiInterface api.Interface) func() {
	return func() {
		if err := cec.CommitTransaction(); err != nil {
			util2.ThrowError(vm, err.Error())
		}
	}
}

func rollback(cec abs.CodeExecutionContext, vm *goja.Runtime, apiInterface api.Interface) func() {
	return func() {
		if err := cec.RollbackTransaction(); err != nil {
			util2.ThrowError(vm, err.Error())
		}
	}
}

func resourceByName(cec abs.CodeExecutionContext, vm *goja.Runtime, apiInterface api.Interface) func(typeName string) *resource_model.Resource {
	return func(typeName string) *resource_model.Resource {
		result, err := apiInterface.GetResourceByType(cec.LocalContext(), typeName)

		if err != nil {
			util2.ThrowError(vm, err.Error())
		}

		return result
	}
}
