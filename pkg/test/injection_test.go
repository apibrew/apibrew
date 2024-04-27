package test

import (
	"github.com/apibrew/apibrew/pkg/api"
	"github.com/apibrew/apibrew/pkg/formats/unstructured"
	"github.com/apibrew/apibrew/pkg/test/setup"
	"github.com/apibrew/apibrew/pkg/util"
	"testing"
)

func TestSortingInjection(t *testing.T) {
	_, err := apiInterface.Apply(setup.Ctx, unstructured.Unstructured{
		"type": "resource",
		"name": "TestSortingInjectionRecord",
		"properties": map[string]interface{}{
			"name": map[string]interface{}{
				"type": "string",
			},
			"weight": map[string]interface{}{
				"type": "int32",
			},
			"details": map[string]interface{}{
				"type":    "struct",
				"typeRef": "Details",
			},
		},
		"types": []interface{}{
			map[string]interface{}{
				"name": "Details",
				"properties": map[string]interface{}{
					"weight": map[string]interface{}{
						"type": "int32",
					},
				},
			},
		},
	})

	if err != nil {
		t.Error(err)
		return
	}

	_, err = apiInterface.Create(setup.Ctx, unstructured.Unstructured{
		"type":   "TestSortingInjectionRecord",
		"name":   "Item1",
		"weight": 321,
		"details": map[string]interface{}{
			"weight": 123,
		},
	})

	if err != nil {
		t.Error(err)
		return
	}

	_, err = apiDirectInterface.List(util.SystemContext, api.ListParams{
		Type: "TestSortingInjectionRecord",
		Sorting: []api.SortingItem{
			{
				Property: "weight1",
			},
		},
	})

	if err == nil {
		t.Error("Sorting injection test failed")
		return
	}

	if err.Error() != "Record Validation failed: Sorting property weight1 is not a valid property path" {
		t.Error("Error should be 'Record Validation failed: Sorting property weight1 is not a valid property path' but: " + err.Error())
		return
	}

	_, err = apiDirectInterface.List(util.SystemContext, api.ListParams{
		Type: "TestSortingInjectionRecord",
		Sorting: []api.SortingItem{
			{
				Property: "details.weight",
			},
			{
				Property: "details.weight2",
			},
		},
	})

	if err == nil {
		t.Error("Sorting injection test failed")
		return
	}

	if err.Error() != "Record Validation failed: Sorting property details.weight2 is not a valid property path" {
		t.Error("Error should be 'Record Validation failed: Sorting property details.weight2 is not a valid property path' but: " + err.Error())
		return
	}
}
