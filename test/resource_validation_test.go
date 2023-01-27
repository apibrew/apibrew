package test

import (
	"data-handler/model"
	"data-handler/server/stub"
	"data-handler/server/util"
	"strconv"
	"testing"
)

func TestCreateResourceValidationForResourceFields(t *testing.T) {
	ctx := prepareTextContext()

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
	ctx := prepareTextContext()

	_, err := resourceServiceClient.Create(ctx, &stub.CreateResourceRequest{
		Resources: []*model.Resource{&model.Resource{
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
	ctx := prepareTextContext()

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
