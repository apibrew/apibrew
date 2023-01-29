package test

import (
	"github.com/google/uuid"
	"github.com/tislib/data-handler/pkg/model"
	"github.com/tislib/data-handler/pkg/server/util"
	"github.com/tislib/data-handler/pkg/stub"
	"strconv"
	"testing"
)

func TestCreateResourceValidationForResourceFields(t *testing.T) {

	testResource := &model.Resource{}

	_, err := resourceServiceClient.Create(ctx, &stub.CreateResourceRequest{
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

	if errorFields[1].Property != "SourceConfig" {
		t.Error("errorFields[1].Property should be SourceConfig")
		return
	}
}

func TestCreateResourceValidationForProperties(t *testing.T) {

	_, err := resourceServiceClient.Create(ctx, &stub.CreateResourceRequest{
		Resources: []*model.Resource{{
			Properties: []*model.ResourceProperty{
				{
					Name: "Type123",
					Type: model.ResourcePropertyType_TYPE_STRING,
				},
				{},
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

	if len(errorFields) != 5 {
		t.Error("There should be 5 errors; but " + strconv.Itoa(len(errorFields)))
		return
	}

	if errorFields[0].Property != "Name" {
		t.Error("errorFields[0].Property should be Name: " + errorFields[0].Property)
	}

	if errorFields[1].Property != "SourceConfig" {
		t.Error("errorFields[1].Property should be SourceConfig: " + errorFields[1].Property)
	}

	if errorFields[2].Property != "Properties[0].SourceConfig" {
		t.Error("errorFields[2].Property should be Properties[0].SourceConfig: " + errorFields[2].Property)
	}

	if errorFields[3].Property != "Properties[1].Name" {
		t.Error("errorFields[3].Property should be Properties[1].Name: " + errorFields[3].Property)
	}

	if errorFields[4].Property != "Properties[1].SourceConfig" {
		t.Error("errorFields[4].Property should be Properties[1].SourceConfig: " + errorFields[4].Property)
	}
}

func TestCreateResourceWithSameName(t *testing.T) {

	_, err := resourceServiceClient.Create(ctx, &stub.CreateResourceRequest{
		Resources:      []*model.Resource{richResource1},
		DoMigration:    true,
		ForceMigration: false,
	})

	if err == nil {
		t.Error("Error should be provided for Resource is already exits")
		return
	}

	if util.GetErrorCode(err) != model.ErrorCode_ALREADY_EXISTS {
		t.Error("Error code should be provided for ErrorCode_ALREADY_EXISTS")
	}
}

func TestCreateResourceWithNonExistingDatasourceShouldFail(t *testing.T) {

	randUUid, _ := uuid.NewRandom()

	resp, err := resourceServiceClient.Create(ctx, &stub.CreateResourceRequest{
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
			_, _ = resourceServiceClient.Delete(ctx, &stub.DeleteResourceRequest{
				Ids:            []string{resp.Resources[0].Id},
				DoMigration:    true,
				ForceMigration: true,
			})
		}
	}()

	if err == nil {
		t.Error("Error should be provided for Resource is already exits")
		return
	}

	if util.GetErrorCode(err) != model.ErrorCode_RESOURCE_VALIDATION_ERROR {
		t.Error("Error code should be provided for ErrorCode_RESOURCE_VALIDATION_ERROR: " + util.GetErrorCode(err).String())
	}
}
