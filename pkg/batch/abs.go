package batch

import (
	"context"
	"github.com/tislib/data-handler/pkg/model"
	"os"
)

type Writer interface {
	StartBatch(batch *model.BatchHeader) error
	EndBatch() error
	WriteResource(resource ...*model.Resource) error
	WriteRecord(namespace string, resourceName string, record ...*model.Record) error
}

type Reader interface {
}

type Executor interface {
	Restore(ctx context.Context, in *os.File) error
}
