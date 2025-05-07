package test

import (
	"github.com/apibrew/apibrew/modules/nano/pkg/model"
	"github.com/apibrew/apibrew/pkg/api"
	"github.com/apibrew/apibrew/pkg/util"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTryCatch(t *testing.T) {
	var api = api.NewInterface(container)

	var testFn = new(model.Script)
	testFn.Source = `
		var actualError
		try {
			throwError("error123");
		} catch(e) {
			actualError = e.message
		}

		actualError
	`
	testFn.Language = model.ScriptLanguage_JAVASCRIPT
	testFn.ContentFormat = model.ScriptContentFormat_TEXT

	result, err := api.Apply(util.SystemContext, model.ScriptMapperInstance.ToUnstructured(testFn))

	if err != nil {
		t.Error(err)
		return
	}

	assert.NotNil(t, result["output"])

	if t.Failed() {
		return
	}

	output := result["output"]

	assert.Equal(t, "error123", output)
}
