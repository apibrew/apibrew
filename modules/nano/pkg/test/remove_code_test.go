package test

import (
	"github.com/apibrew/apibrew/pkg/api"
	"github.com/apibrew/apibrew/pkg/formats/unstructured"
	"github.com/apibrew/apibrew/pkg/util"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRemoveCote(t *testing.T) {
	var api = api.NewInterface(container)

	codeApplyRes, err := api.Apply(util.SystemContext, unstructured.Unstructured{
		"type":             "nano/Code",
		"name":             "TestAction1.js",
		"concurrencyLevel": 1,
		"content": `
		const TestAction1 = resource({
			name: "TestAction1",
			virtual: true,
			properties: {
				output: {
					type: "int32",
				}
			}
		})

		TestAction1.beforeCreate(req => {
			req.output = 5 + 6

			return req
		})
`,
	})

	if err != nil {
		t.Error(err)
		return
	}

	recordCreateRes, err := api.Create(util.SystemContext, unstructured.Unstructured{
		"type": "TestAction1",
	})

	if err != nil {
		t.Error(err)
		return
	}

	t.Log("Record create result: ", recordCreateRes)

	assert.NotNil(t, recordCreateRes["output"])

	assert.Equal(t, float64(11), recordCreateRes["output"])

	if t.Failed() {
		return
	}

	codeApplyRes["type"] = "nano/Code"

	err = api.Delete(util.SystemContext, codeApplyRes)

	if err != nil {
		t.Error(err)
		return
	}

	recordCreate2Res, err := api.Create(util.SystemContext, unstructured.Unstructured{
		"type": "TestAction1",
	})

	if err != nil {
		t.Error(err)
		return
	}

	t.Log("Record create 2 result: ", recordCreate2Res)

	assert.Nil(t, recordCreate2Res["output"])
}
