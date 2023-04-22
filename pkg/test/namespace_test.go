package test

import (
	"github.com/tislib/apibrew/pkg/model"
	"github.com/tislib/apibrew/pkg/stub"
	"github.com/tislib/apibrew/pkg/test/setup"
	"testing"
)

func TestNamespaceNameShouldNotBeUpdated(t *testing.T) {
	namespace1 := &model.Namespace{
		Name: "test-namespace",
	}

	res, err := namespaceClient.Create(setup.Ctx, &stub.CreateNamespaceRequest{
		Namespaces: []*model.Namespace{
			namespace1,
		},
	})

	if err != nil {
		t.Error(err)
		return
	}

	if res.Namespaces != nil {
		namespace1.Id = res.Namespaces[0].Id
	} else {
		t.Error("Namespace was not created")
		return
	}

	defer func() {
		_, _ = namespaceClient.Delete(setup.Ctx, &stub.DeleteNamespaceRequest{
			Ids: []string{
				namespace1.Id,
			},
		})
	}()

	// try to update

	namespace1.Name = "test-123321123"

	_, err = namespaceClient.Update(setup.Ctx, &stub.UpdateNamespaceRequest{
		Namespaces: []*model.Namespace{
			namespace1,
		},
	})

	if err != nil {
		t.Error(err)
		return
	}

	res2, err := namespaceClient.Get(setup.Ctx, &stub.GetNamespaceRequest{Id: namespace1.Id})

	if err != nil {
		t.Error(err)
		return
	}

	if res2.Namespace.Name != "test-namespace" {
		t.Error("Namespace name is immutable and it must not be updated")
	}
}
