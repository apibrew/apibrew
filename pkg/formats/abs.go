package formats

import (
	"context"
	"github.com/apibrew/apibrew/pkg/formats/unstructured"
	"github.com/apibrew/apibrew/pkg/resource_model"
)

type Writer interface {
	WriteResource(resource ...*resource_model.Resource) error
	WriteRecord(typ string, record ...unstructured.Unstructured) error
	WriteRecords(resource *resource_model.Resource, total uint32, records []unstructured.Unstructured) error
	IsBinary() bool
}

type Reader interface {
}

type Executor interface {
	Restore(ctx context.Context) error
}
