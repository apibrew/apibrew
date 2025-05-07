package test

import (
	"github.com/apibrew/apibrew/modules/nano/pkg/model"
	"github.com/apibrew/apibrew/pkg/api"
	"github.com/apibrew/apibrew/pkg/util"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBasicAction(t *testing.T) {
	var api = api.NewInterface(container)

	var testAction = new(model.Action)
	testAction.Name = `TestAction`
	testAction.RestPath = `test`
	testAction.Source = `function({a, b}) {
	return a + b 
}
`
	testAction.Language = model.ActionLanguage_JAVASCRIPT
	testAction.ContentFormat = model.ActionContentFormat_TEXT

	_, err := api.Apply(util.SystemContext, model.ActionMapperInstance.ToUnstructured(testAction))

	if err != nil {
		t.Error(err)
		return
	}

	res, err := api.Create(util.SystemContext, map[string]interface{}{
		"type": "actions/TestAction",
		"input": map[string]interface{}{
			"a": 5,
			"b": 6,
		},
	})

	if err != nil {
		t.Error(err)
		return
	}

	assert.Equal(t, float64(11), res["output"])
}

func TestActionRename(t *testing.T) {
	var api = api.NewInterface(container)

	var testAction = new(model.Action)
	testAction.Name = `TestAction`
	testAction.RestPath = `test`
	testAction.Source = `function({a, b}) {
	return a + b 
}
`
	testAction.Language = model.ActionLanguage_JAVASCRIPT
	testAction.ContentFormat = model.ActionContentFormat_TEXT

	res, err := api.Apply(util.SystemContext, model.ActionMapperInstance.ToUnstructured(testAction))

	if err != nil {
		t.Error(err)
		return
	}

	testAction = new(model.Action)
	testAction.Id = new(uuid.UUID)
	*testAction.Id = uuid.MustParse(res["id"].(string))
	testAction.Name = `TestAction2`
	testAction.Source = `function({a, b}) {
	return a + b 
}
`
	testAction.Language = model.ActionLanguage_JAVASCRIPT
	testAction.ContentFormat = model.ActionContentFormat_TEXT

	_, err = api.Update(util.SystemContext, model.ActionMapperInstance.ToUnstructured(testAction))

	if err != nil {
		t.Error(err)
		return
	}

	res, err = api.Create(util.SystemContext, map[string]interface{}{
		"type": "actions/TestAction2",
		"input": map[string]interface{}{
			"a": 5,
			"b": 6,
		},
	})

	if err != nil {
		t.Error(err)
		return
	}

	assert.Equal(t, float64(11), res["output"])
}

func TestActionUpdateSource(t *testing.T) {
	var api = api.NewInterface(container)

	var testAction = new(model.Action)
	testAction.Name = `TestAction`
	testAction.RestPath = `test`
	testAction.Source = `function({a, b}) {
	return a + b 
}
`
	testAction.Language = model.ActionLanguage_JAVASCRIPT
	testAction.ContentFormat = model.ActionContentFormat_TEXT

	_, err := api.Apply(util.SystemContext, model.ActionMapperInstance.ToUnstructured(testAction))

	if err != nil {
		t.Error(err)
		return
	}

	testAction.Source = `function({a, b}) {
	return a - b 
}
`
	_, err = api.Apply(util.SystemContext, model.ActionMapperInstance.ToUnstructured(testAction))

	if err != nil {
		t.Error(err)
		return
	}

	res, err := api.Create(util.SystemContext, map[string]interface{}{
		"type": "actions/TestAction",
		"input": map[string]interface{}{
			"a": 5,
			"b": 6,
		},
	})

	if err != nil {
		t.Error(err)
		return
	}

	assert.Equal(t, float64(-1), res["output"])
}
