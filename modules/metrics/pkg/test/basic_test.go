package test

import (
	"github.com/apibrew/apibrew/modules/metrics/pkg/model"
	"github.com/apibrew/apibrew/pkg/api"
	"github.com/apibrew/apibrew/pkg/util"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestBasicExecution(t *testing.T) {
	repo := api.NewRepository[*model.TestResource](apiInterface, model.TestResourceMapperInstance)

	var entity = &model.TestResource{
		Name: util.Pointer("test"),
	}
	result, err := repo.Create(util.SystemContext, entity)

	if err != nil {
		t.Error(err)
		return
	}

	assert.Equal(t, result.Name, entity.Name)
	assert.Equal(t, result.Description, entity.Name)
}
