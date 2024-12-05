package test

import (
	"context"
	"github.com/apibrew/apibrew/pkg/api"
	"github.com/apibrew/apibrew/pkg/resource_model"
	"github.com/apibrew/apibrew/pkg/util"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestInternalApiCalls(t *testing.T) {
	repo := api.NewRepository[*resource_model.User](apiInterface, resource_model.UserMapperInstance)

	var entity = &resource_model.User{
		Username: "test" + util.RandomHex(5),
		Password: util.Pointer("test"),
	}
	result, err := repo.Create(util.SystemContext, entity)

	if err != nil {
		t.Error(err)
		return
	}

	assert.Equal(t, result.Username, entity.Username)

	list, err := repo.List(context.TODO(), api.ListParams{
		Filters: map[string]interface{}{
			"username": entity.Username,
		},
	})

	if err != nil {
		t.Error(err)
		return
	}

	assert.Equal(t, len(list.Content), 1)
	assert.Equal(t, list.Content[0].Username, entity.Username)

}
