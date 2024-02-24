package api

import (
	"github.com/apibrew/apibrew/pkg/api"
	"github.com/apibrew/apibrew/pkg/formats/unstructured"
	"github.com/apibrew/apibrew/pkg/nano/abs"
	"github.com/apibrew/apibrew/pkg/resource_model"
	"github.com/apibrew/apibrew/pkg/util"
)

func create(cec abs.CodeExecutionContext, apiInterface api.Interface) func(record unstructured.Unstructured) unstructured.Unstructured {
	return func(record unstructured.Unstructured) unstructured.Unstructured {
		result, err := apiInterface.Create(cec.Context(), record)

		if err != nil {
			panic(err)
		}

		return result
	}
}

func update(cec abs.CodeExecutionContext, apiInterface api.Interface) func(record unstructured.Unstructured) unstructured.Unstructured {
	return func(record unstructured.Unstructured) unstructured.Unstructured {
		result, err := apiInterface.Update(cec.Context(), record)

		if err != nil {
			panic(err)
		}

		return result
	}
}

func apply(cec abs.CodeExecutionContext, apiInterface api.Interface) func(record unstructured.Unstructured) unstructured.Unstructured {
	return func(record unstructured.Unstructured) unstructured.Unstructured {
		result, err := apiInterface.Apply(cec.Context(), record)

		if err != nil {
			panic(err)
		}

		return result
	}
}

func delete(cec abs.CodeExecutionContext, apiInterface api.Interface) func(record unstructured.Unstructured) {
	return func(record unstructured.Unstructured) {
		err := apiInterface.Delete(cec.Context(), record)

		if err != nil {
			panic(err)
		}
	}
}

func load(cec abs.CodeExecutionContext, apiInterface api.Interface) func(record unstructured.Unstructured, params api.LoadParams) unstructured.Unstructured {
	return func(record unstructured.Unstructured, params api.LoadParams) unstructured.Unstructured {
		result, err := apiInterface.Load(cec.Context(), record, params)

		if err != nil {
			panic(err)
		}

		return result
	}
}

func list(cec abs.CodeExecutionContext, apiInterface api.Interface) func(params api.ListParams) unstructured.Unstructured {
	return func(params api.ListParams) unstructured.Unstructured {
		result, err := apiInterface.List(cec.Context(), params)

		if err != nil {
			panic(err)
		}

		return unstructured.Unstructured{
			"content": util.ArrayMap(result.Content, func(item unstructured.Unstructured) interface{} {
				return item
			}),
			"total": result.Total,
		}
	}
}

func resourceByName(cec abs.CodeExecutionContext, apiInterface api.Interface) func(typeName string) *resource_model.Resource {
	return func(typeName string) *resource_model.Resource {
		result, err := apiInterface.GetResourceByType(cec.Context(), typeName)

		if err != nil {
			panic(err)
		}

		return result
	}
}
