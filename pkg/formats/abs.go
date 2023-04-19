package formats

import (
	"context"
	"github.com/tislib/data-handler/pkg/model"
)

type Writer interface {
	WriteResource(resource ...*model.Resource) error
	WriteRecord(namespace string, resourceName string, record ...*model.Record) error
}

type Reader interface {
}

type Executor interface {
	Restore(ctx context.Context) error
}
