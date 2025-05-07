package test

import (
	"github.com/apibrew/apibrew/modules/nano/pkg/model"
	"github.com/apibrew/apibrew/pkg/api"
	"github.com/apibrew/apibrew/pkg/formats/unstructured"
	"github.com/apibrew/apibrew/pkg/util"
	"testing"
)

func runScript(t testing.TB, source string) unstructured.Unstructured {
	return runScriptWithLanguage(t, source, model.ScriptLanguage_JAVASCRIPT)
}

func runScriptWithLanguage(t testing.TB, source string, lang model.ScriptLanguage) unstructured.Unstructured {
	var api = api.NewInterface(container)

	var testFn = new(model.Script)
	testFn.Source = source
	testFn.Language = lang
	testFn.ContentFormat = model.ScriptContentFormat_TEXT

	result, err := api.Apply(util.SystemContext, model.ScriptMapperInstance.ToUnstructured(testFn))

	if err != nil {
		t.Error(err)
		return nil
	}
	return result
}
