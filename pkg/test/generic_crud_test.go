package test

import (
	"github.com/tislib/data-handler/pkg/model"
	"github.com/tislib/data-handler/pkg/resources"
	"github.com/tislib/data-handler/pkg/resources/mapping"
	"github.com/tislib/data-handler/pkg/server/util"
	"github.com/tislib/data-handler/pkg/stub"
	"google.golang.org/protobuf/types/known/anypb"
	"testing"
)

func TestCreateUser1(t *testing.T) {
	user1 := new(model.User)
	user1.Username = "taleh999"
	user1.Password = "taleh999"

	any1, err := mapping.MessageToAny(user1)

	if err != nil {
		t.Error(err)
		return
	}

	_, err = genericServiceClient.Create(ctx, &stub.CreateRequest{
		Namespace: "system",
		Resource:  resources.UserResource.Name,
		Items:     []*anypb.Any{any1},
	})

	if err == nil {
		t.Error("Save should fail")
	}

	if util.GetErrorCode(err) != model.ErrorCode_RECORD_VALIDATION_ERROR {
		t.Error("Error code should be: " + model.ErrorCode_RECORD_VALIDATION_ERROR.String())
	}

	errorFields := util.GetErrorFields(err)

	if len(errorFields) != 13 {
		t.Error("There must be 14 error field")
	}
}
