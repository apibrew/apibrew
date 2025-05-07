package test

import (
	"github.com/apibrew/apibrew/modules/nano/pkg/model"
	"github.com/apibrew/apibrew/pkg/api"
	"github.com/apibrew/apibrew/pkg/util"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNanoScriptBasic(t *testing.T) {
	result := runScript(t, `5 + 6`)
	if t.Failed() {
		return
	}

	assert.NotNil(t, result["output"])

	if t.Failed() {
		return
	}

	output := result["output"]

	assert.Equal(t, float64(11), output)
}

func TestNanoScriptMultiLine(t *testing.T) {
	var api = api.NewInterface(container)

	var testFn = new(model.Script)
	testFn.Source = `const a = 5 + 6;
const b = 2 + 3;
a + b;
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

	assert.Equal(t, float64(16), output)
}

func TestNanoScriptWithFunctionWrapper(t *testing.T) {
	var api = api.NewInterface(container)

	var testFn = new(model.Script)
	testFn.Source = `
(function() {
	  return 5 + 6;
})()
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

	assert.Equal(t, float64(11), output)
}
