package test

import (
	"errors"
	errors2 "github.com/apibrew/apibrew/pkg/errors"
	"github.com/apibrew/apibrew/pkg/formats/unstructured"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/service/annotations"
	"github.com/apibrew/apibrew/pkg/test/setup"
	"testing"
)

func TestAddReferenceCascade(t *testing.T) {
	_, err := apiInterface.Apply(setup.Ctx, unstructured.Unstructured{
		"type": "resource",
		"name": "TestAddReferenceAuthor",
		"properties": map[string]interface{}{
			"name": map[string]interface{}{
				"type":   "string",
				"unique": true,
			},
		},
	})

	if err != nil {
		t.Error(err)
		return
	}

	_, err = apiInterface.Apply(setup.Ctx, unstructured.Unstructured{
		"type": "resource",
		"name": "TestAddReferenceBook",
		"properties": map[string]interface{}{
			"name": map[string]interface{}{
				"type": "string",
			},
			"author": map[string]interface{}{
				"type":      "reference",
				"reference": "TestAddReferenceAuthor",
			},
		},
	})

	if err != nil {
		t.Error(err)
		return
	}

	_, err = apiInterface.Apply(setup.Ctx, unstructured.Unstructured{
		"type": "TestAddReferenceAuthor",
		"name": "Author1",
	})

	if err != nil {
		t.Error(err)
		return
	}

	_, err = apiInterface.Create(setup.Ctx, unstructured.Unstructured{
		"type": "TestAddReferenceBook",
		"name": "Book1",
		"author": map[string]interface{}{
			"name": "Author1",
		},
	})

	if err != nil {
		t.Error(err)
		return
	}

	err = apiInterface.Delete(setup.Ctx, unstructured.Unstructured{
		"type": "TestAddReferenceAuthor",
		"name": "Author1",
	})

	if err == nil {
		t.Error(errors.New("error expected ErrorCode_REFERENCE_EXISTS but it succeeded"))
		return
	}

	if errors2.FromServiceError(err).Code() != model.ErrorCode_REFERENCE_VIOLATION {
		t.Error(errors.New("error expected ErrorCode_REFERENCE_VIOLATION but got " + errors2.FromServiceError(err).Code().String()))
		return
	}

	_, err = apiInterface.Apply(setup.Ctx, unstructured.Unstructured{
		"type": "resource",
		"name": "TestAddReferenceBook",
		"properties": map[string]interface{}{
			"name": map[string]interface{}{
				"type": "string",
			},
			"author": map[string]interface{}{
				"type":      "reference",
				"reference": "TestAddReferenceAuthor",
				"annotations": map[string]interface{}{
					annotations.CascadeReference: "true",
				},
			},
		},
	})

	if err != nil {
		t.Error(err)
		return
	}

	err = apiInterface.Delete(setup.Ctx, unstructured.Unstructured{
		"type": "TestAddReferenceAuthor",
		"name": "Author1",
	})

	if err != nil {
		t.Error(err)
		return
	}
}

func TestRemoveReferenceCascade(t *testing.T) {
	_, err := apiInterface.Apply(setup.Ctx, unstructured.Unstructured{
		"type": "resource",
		"name": "TestAddReferenceAuthor",
		"properties": map[string]interface{}{
			"name": map[string]interface{}{
				"type":   "string",
				"unique": true,
			},
		},
	})

	if err != nil {
		t.Error(err)
		return
	}

	_, err = apiInterface.Apply(setup.Ctx, unstructured.Unstructured{
		"type": "resource",
		"name": "TestAddReferenceBook",
		"properties": map[string]interface{}{
			"name": map[string]interface{}{
				"type": "string",
			},
			"author": map[string]interface{}{
				"type":      "reference",
				"reference": "TestAddReferenceAuthor",
				"annotations": map[string]interface{}{
					annotations.CascadeReference: "true",
				},
			},
		},
	})

	if err != nil {
		t.Error(err)
		return
	}

	_, err = apiInterface.Apply(setup.Ctx, unstructured.Unstructured{
		"type": "TestAddReferenceAuthor",
		"name": "Author1",
	})

	if err != nil {
		t.Error(err)
		return
	}

	_, err = apiInterface.Create(setup.Ctx, unstructured.Unstructured{
		"type": "TestAddReferenceBook",
		"name": "Book1",
		"author": map[string]interface{}{
			"name": "Author1",
		},
	})

	if err != nil {
		t.Error(err)
		return
	}

	err = apiInterface.Delete(setup.Ctx, unstructured.Unstructured{
		"type": "TestAddReferenceAuthor",
		"name": "Author1",
	})

	if err != nil {
		t.Error(err)
		return
	}

	_, err = apiInterface.Apply(setup.Ctx, unstructured.Unstructured{
		"type": "TestAddReferenceAuthor",
		"name": "Author1",
	})

	if err != nil {
		t.Error(err)
		return
	}

	_, err = apiInterface.Create(setup.Ctx, unstructured.Unstructured{
		"type": "TestAddReferenceBook",
		"name": "Book1",
		"author": map[string]interface{}{
			"name": "Author1",
		},
	})

	if err != nil {
		t.Error(err)
		return
	}

	_, err = apiInterface.Apply(setup.Ctx, unstructured.Unstructured{
		"type": "resource",
		"name": "TestAddReferenceBook",
		"properties": map[string]interface{}{
			"name": map[string]interface{}{
				"type": "string",
			},
			"author": map[string]interface{}{
				"type":      "reference",
				"reference": "TestAddReferenceAuthor",
			},
		},
	})

	if err != nil {
		t.Error(err)
		return
	}

	err = apiInterface.Delete(setup.Ctx, unstructured.Unstructured{
		"type": "TestAddReferenceAuthor",
		"name": "Author1",
	})

	if err == nil {
		t.Error(errors.New("error expected ErrorCode_REFERENCE_EXISTS but it succeeded"))
		return
	}

	if errors2.FromServiceError(err).Code() != model.ErrorCode_REFERENCE_VIOLATION {
		t.Error(errors.New("error expected ErrorCode_REFERENCE_VIOLATION but got " + errors2.FromServiceError(err).Code().String()))
		return
	}
}
