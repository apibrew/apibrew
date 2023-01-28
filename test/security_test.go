package test

import (
	"data-handler/model"
	"data-handler/server/stub"
	"data-handler/server/util"
	log "github.com/sirupsen/logrus"
	"testing"
)

func TestDhTestUserCannotCreateUser(t *testing.T) {
	userDhTestCtx := withUserAuthenticationContext(ctx, "dh_test", "dh_test")

	_, err := userServiceClient.Create(userDhTestCtx, &stub.CreateUserRequest{
		Users: []*model.User{
			{
				Username: "test123",
			},
		},
	})

	if err == nil {
		t.Error("error expected ErrorCode_ACCESS_DENIED but it succeeded")
		return
	}

	if util.GetErrorCode(err) != model.ErrorCode_ACCESS_DENIED {
		log.Print(err)
		t.Error("error expected ErrorCode_ACCESS_DENIED")
	}
}

func TestDhTestUserCanReadUser(t *testing.T) {
	userDhTestCtx := withUserAuthenticationContext(ctx, "dh_test", "dh_test")

	_, err := userServiceClient.List(userDhTestCtx, &stub.ListUserRequest{})

	if err != nil {
		log.Print(err)
		t.Error("read operation should be successful")
		return
	}
}
