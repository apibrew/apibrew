package test

import (
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/resource_model"
	"github.com/apibrew/apibrew/pkg/stub"
	"github.com/apibrew/apibrew/pkg/test/setup"
	"github.com/google/uuid"
	"testing"
)

func TestNamespaceNameShouldNotBeUpdated(t *testing.T) {
	namespace1 := &resource_model.Namespace{
		Name: "test-namespace",
	}

	res, err := recordClient.Create(setup.Ctx, &stub.CreateRecordRequest{
		Records: []*model.Record{
			resource_model.NamespaceMapperInstance.ToRecord(namespace1),
		},
	})

	if err != nil {
		t.Error(err)
		return
	}

	if res.Records != nil {
		namespace1.Id = new(uuid.UUID)
		*namespace1.Id = uuid.MustParse(res.Records[0].Id)
	} else {
		t.Error("Namespace was not created")
		return
	}

	defer func() {
		_, _ = recordClient.Delete(setup.Ctx, &stub.DeleteRecordRequest{
			Ids: []string{
				namespace1.Id.String(),
			},
		})
	}()

	// try to update

	namespace1.Name = "test-123321123"

	_, err = recordClient.Update(setup.Ctx, &stub.UpdateRecordRequest{
		Records: []*model.Record{
			resource_model.NamespaceMapperInstance.ToRecord(namespace1),
		},
	})

	if err != nil {
		t.Error(err)
		return
	}

	res2, err := recordClient.Get(setup.Ctx, &stub.GetRecordRequest{Id: namespace1.Id.String()})

	if err != nil {
		t.Error(err)
		return
	}

	if res2.Record.Properties["Name"].GetStringValue() != "test-namespace" {
		t.Error("Namespace name is immutable and it must not be updated")
	}
}
