package test

import (
	"data-handler/model"
	"data-handler/server/stub"
	"data-handler/server/util"
	log "github.com/sirupsen/logrus"
	"testing"
)

func TestCreateResourceValidation(t *testing.T) {
	ctx := prepareTextContext()

	testResource := &model.Resource{}

	resp, err := resourceServiceClient.Create(ctx, &stub.CreateResourceRequest{
		Resources:      []*model.Resource{testResource},
		DoMigration:    true,
		ForceMigration: false,
	})

	if util.GetErrorCode(err) != model.ErrorCode_RESOURCE_VALIDATION_ERROR {
		t.Error("Error should be RESOURCE_VALIDATION_ERROR")
		return
	}

	errorFields := util.GetErrorFields(err)

	log.Print(resp, err, errorFields)
}
