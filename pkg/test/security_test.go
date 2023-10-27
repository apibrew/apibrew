package test

import (
	"context"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/resource_model"
	"github.com/apibrew/apibrew/pkg/resources"
	"github.com/apibrew/apibrew/pkg/stub"
	"github.com/apibrew/apibrew/pkg/test/setup"
	"github.com/apibrew/apibrew/pkg/util"
	log "github.com/sirupsen/logrus"
	"testing"
)

func TestDhTestUserCannotCreateUser(t *testing.T) {
	userDhTestCtx := setup.WithUserAuthenticationContext(context.TODO(), "dh_test", "dh_test")

	record, err := recordClient.Apply(userDhTestCtx, &stub.ApplyRecordRequest{
		Namespace: resources.UserResource.Namespace,
		Resource:  resources.UserResource.Name,
		Records: []*model.Record{
			resource_model.UserMapperInstance.ToRecord(&resource_model.User{
				Username: "test123",
			}),
		},
	})

	if record != nil && record.Records != nil && len(record.Records) > 0 {
		_, err := recordClient.Delete(setup.Ctx, &stub.DeleteRecordRequest{
			Namespace: resources.UserResource.Namespace,
			Resource:  resources.UserResource.Name,
			Ids: []string{
				util.GetRecordId(nil, record.Records[0]),
			},
		})

		if err != nil {
			log.Print(err)
			t.Error("error while deleting test user")
		}
	}

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
	log.Info("Before TestDhTestUserCanReadUser")
	userDhTestCtx := setup.WithUserAuthenticationContext(setup.Ctx, "dh_test", "dh_test")
	log.Info("Prepare Context")

	_, err := recordClient.List(userDhTestCtx, &stub.ListRecordRequest{
		Namespace: resources.UserResource.Namespace,
		Resource:  resources.UserResource.Name,
	})

	log.Info("After recordClient.ListX", err)

	if err != nil {
		log.Print(err)
		t.Error("read operation should be successful")
		return
	}
}
