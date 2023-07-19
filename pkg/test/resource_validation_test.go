package test

import (
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/stub"
	"github.com/apibrew/apibrew/pkg/test/setup"
	"github.com/apibrew/apibrew/pkg/util"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestCreateResourceValidationForResourceFields(t *testing.T) {

	testResource := &model.Resource{}

	_, err := resourceClient.Create(setup.Ctx, &stub.CreateResourceRequest{
		Resources:      []*model.Resource{testResource},
		DoMigration:    true,
		ForceMigration: false,
	})

	if util.GetErrorCode(err) != model.ErrorCode_RESOURCE_VALIDATION_ERROR {
		t.Error("Error should be RESOURCE_VALIDATION_ERROR")
		return
	}

	errorFields := util.GetErrorFields(err)

	if len(errorFields) != 2 {
		t.Error("There should be 3 errors")
		return
	}

	if errorFields[0].Property != "Name" {
		t.Error("errorFields[0].Property should be Name")
		return
	}

	if errorFields[1].Property != "SourceConfig.Entity" {
		t.Error("errorFields[1].Property should be SourceConfig: " + errorFields[1].Property)
		return
	}
}

func TestCreateResourceValidationForProperties(t *testing.T) {

	_, err := resourceClient.Create(setup.Ctx, &stub.CreateResourceRequest{
		Resources: []*model.Resource{{
			Properties: []*model.ResourceProperty{
				{
					Name: "Type123",
					Type: model.ResourceProperty_STRING,
				},
				{},
				{
					Name: "Type321",
					Type: model.ResourceProperty_REFERENCE,
				},
			},
		}},
		DoMigration:    true,
		ForceMigration: false,
	})

	if util.GetErrorCode(err) != model.ErrorCode_RESOURCE_VALIDATION_ERROR {
		t.Error("Error should be RESOURCE_VALIDATION_ERROR")
		return
	}

	errorFields := util.GetErrorFields(err)

	if len(errorFields) != 4 {
		t.Error("There should be 4 errors; but " + strconv.Itoa(len(errorFields)))
		return
	}

	assert.Equal(t, errorFields[0].Property, "Name")
	assert.Equal(t, errorFields[1].Property, "SourceConfig.Entity")
	assert.Equal(t, errorFields[2].Property, ".Name{index:1}")
	assert.Equal(t, errorFields[3].Property, "Type321.Reference")

}

func TestCreateResourceWithSameName(t *testing.T) {
	_, err := resourceClient.Create(setup.Ctx, &stub.CreateResourceRequest{
		Resources:      []*model.Resource{setup.RichResource1},
		DoMigration:    true,
		ForceMigration: false,
	})

	if err == nil {
		t.Error("Error should be provided for resource is already exits")
		return
	}

	if util.GetErrorCode(err) != model.ErrorCode_ALREADY_EXISTS {
		t.Error("Error code should be provided for ErrorCode_ALREADY_EXISTS but: "+util.GetErrorCode(err).String(), err.Error())
	}
}

func TestCreateResourceWithNonExistingRecordShouldFail(t *testing.T) {
	randUUid, _ := uuid.NewRandom()

	resp, err := resourceClient.Create(setup.Ctx, &stub.CreateResourceRequest{
		Resources: []*model.Resource{
			{
				Name: "non_existent_source",
				SourceConfig: &model.ResourceSourceConfig{
					DataSource: randUUid.String(),
					Catalog:    "catalog_1",
					Entity:     "entity_1",
				},
			},
		},
		DoMigration:    true,
		ForceMigration: true,
	})

	defer func() {
		if resp != nil && len(resp.Resources) > 0 {
			_, _ = resourceClient.Delete(setup.Ctx, &stub.DeleteResourceRequest{
				Ids:            []string{resp.Resources[0].Id},
				DoMigration:    true,
				ForceMigration: true,
			})
		}
	}()

	if err == nil {
		t.Error("Error should be provided for resource is already exits")
		return
	}

	if util.GetErrorCode(err) != model.ErrorCode_ALREADY_EXISTS {
		print(util.GetErrorCode(err))
		t.Error("Error code should be provided for ErrorCode_RESOURCE_VALIDATION_ERROR: " + util.GetErrorCode(err).String())
	}
}
