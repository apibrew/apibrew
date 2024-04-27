package test

import (
	"github.com/apibrew/apibrew/pkg/api"
	"github.com/apibrew/apibrew/pkg/formats/unstructured"
	"github.com/apibrew/apibrew/pkg/test/setup"
	"github.com/apibrew/apibrew/pkg/util"
	"math/rand"
	"strconv"
	"testing"
)

func TestSimpleSorting(t *testing.T) {
	_, err := apiInterface.Apply(setup.Ctx, unstructured.Unstructured{
		"type": "resource",
		"name": "TestSimpleSortingList",
		"properties": map[string]interface{}{
			"name": map[string]interface{}{
				"type": "string",
			},
			"weight": map[string]interface{}{
				"type": "int32",
			},
		},
	})

	if err != nil {
		t.Error(err)
		return
	}

	result, err := apiInterface.List(setup.Ctx, api.ListParams{
		Type: "TestSimpleSortingList",
		Sorting: []api.SortingItem{
			{
				Property:  "weight",
				Direction: "DESC",
			},
		},
	})

	if err != nil {
		t.Error(err)
		return
	}

	for _, item := range result.Content {
		if err = apiInterface.Delete(setup.Ctx, item); err != nil {
			t.Error(err)
			return
		}
	}

	var weights = rand.Perm(25)

	for i := 0; i < 25; i++ {
		_, err = apiInterface.Create(setup.Ctx, unstructured.Unstructured{
			"type":   "TestSimpleSortingList",
			"name":   "Item" + strconv.Itoa(i),
			"weight": weights[i],
		})

		if err != nil {
			t.Error(err)
			return
		}
	}

	result, err = apiDirectInterface.List(util.SystemContext, api.ListParams{
		Type: "TestSimpleSortingList",
		Sorting: []api.SortingItem{
			{
				Property:  "weight",
				Direction: "DESC",
			},
		},
	})

	if err != nil {
		t.Error(err)
		return
	}

	var items = result.Content

	for i := 0; i < 24; i++ {
		if items[i]["weight"].(float64) < items[i+1]["weight"].(float64) {
			t.Error("Sorting failed")
			return
		}
	}

	result, err = apiDirectInterface.List(util.SystemContext, api.ListParams{
		Type: "TestSimpleSortingList",
		Sorting: []api.SortingItem{
			{
				Property:  "weight",
				Direction: "ASC",
			},
		},
	})

	if err != nil {
		t.Error(err)
		return
	}

	items = result.Content

	for i := 0; i < 24; i++ {
		if items[i]["weight"].(float64) > items[i+1]["weight"].(float64) {
			t.Error("Sorting failed")
			return
		}
	}

}

func TestNestedSorting(t *testing.T) {
	_, err := apiInterface.Apply(setup.Ctx, unstructured.Unstructured{
		"type": "resource",
		"name": "TestNestedSortingList",
		"properties": map[string]interface{}{
			"name": map[string]interface{}{
				"type": "string",
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

	result, err := apiInterface.List(setup.Ctx, api.ListParams{
		Type: "TestNestedSortingList",
	})

	if err != nil {
		t.Error(err)
		return
	}

	for _, item := range result.Content {
		if err = apiInterface.Delete(setup.Ctx, item); err != nil {
			t.Error(err)
			return
		}
	}

	var weights = rand.Perm(25)

	for i := 0; i < 25; i++ {
		_, err = apiInterface.Create(setup.Ctx, unstructured.Unstructured{
			"type": "TestNestedSortingList",
			"name": "Item" + strconv.Itoa(i),
			"details": map[string]interface{}{
				"weight": weights[i],
			},
		})

		if err != nil {
			t.Error(err)
			return
		}
	}

	result, err = apiDirectInterface.List(util.SystemContext, api.ListParams{
		Type: "TestNestedSortingList",
		Sorting: []api.SortingItem{
			{
				Property:  "details.weight",
				Direction: "DESC",
			},
		},
	})

	if err != nil {
		t.Error(err)
		return
	}

	var items = result.Content

	for i := 0; i < 24; i++ {
		if items[i]["details"].(map[string]interface{})["weight"].(float64) < items[i+1]["details"].(map[string]interface{})["weight"].(float64) {
			t.Error("Sorting failed")
			return
		}
	}

	result, err = apiDirectInterface.List(util.SystemContext, api.ListParams{
		Type: "TestNestedSortingList",
		Sorting: []api.SortingItem{
			{
				Property:  "details.weight",
				Direction: "ASC",
			},
		},
	})

	if err != nil {
		t.Error(err)
		return
	}

	items = result.Content

	for i := 0; i < 24; i++ {
		if items[i]["details"].(map[string]interface{})["weight"].(float64) > items[i+1]["details"].(map[string]interface{})["weight"].(float64) {
			t.Error("Sorting failed")
			return
		}
	}

}
