package test

import (
	pkg "github.com/apibrew/apibrew/modules/tenant"
	"github.com/apibrew/apibrew/pkg/api"
	"github.com/apibrew/apibrew/pkg/formats/unstructured"
	"github.com/apibrew/apibrew/pkg/test/setup"
	"github.com/apibrew/apibrew/pkg/util"
	"testing"
)

func TestBasicExecution(t *testing.T) {
	apiInt := api.NewInterface(container)

	// create new tenants
	_, err := apiInt.Apply(util.WithSystemContext(setup.Ctx), map[string]unstructured.Any{
		"type": "Tenant",
		"name": "test1",
	})

	if err != nil {
		t.Error(err)
		return
	}
	_, err = apiInt.Apply(util.WithSystemContext(setup.Ctx), map[string]unstructured.Any{
		"type": "Tenant",
		"name": "test2",
	})

	if err != nil {
		t.Error(err)
		return
	}

	_, err = apiInt.Apply(pkg.WithTenant(util.WithSystemContext(setup.Ctx), "test1"), map[string]unstructured.Any{
		"type": "system/Resource",
		"name": "TestResourceForTenant1",
		"properties": map[string]interface{}{
			"test1": map[string]interface{}{
				"type": "string",
			},
		},
	})

	if err != nil {
		t.Error(err)
		return
	}

	_, err = apiInt.Apply(pkg.WithTenant(util.WithSystemContext(setup.Ctx), "test2"), map[string]unstructured.Any{
		"type": "system/Resource",
		"name": "TestResourceForTenant2",
		"properties": map[string]interface{}{
			"test1": map[string]interface{}{
				"type": "string",
			},
		},
	})

	if err != nil {
		t.Error(err)
		return
	}

	_, err = apiInt.Apply(pkg.WithTenant(util.WithSystemContext(setup.Ctx), "test2"), map[string]unstructured.Any{
		"type":  "TestResourceForTenant2",
		"test1": "value1",
	})

	if err != nil {
		t.Error(err)
		return
	}
}
