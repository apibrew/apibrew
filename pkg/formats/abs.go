package formats

import (
	"context"
	"github.com/apibrew/apibrew/pkg/model"
)

type Writer interface {
	WriteResource(resource ...*model.Resource) error
	WriteRecord(namespace string, resourceName string, record ...*model.Record) error
	WriteRecordsChan(resource *model.Resource, total uint32, recordsChan chan *model.Record) error
	IsBinary() bool
}

type Reader interface {
}

type Executor interface {
	Restore(ctx context.Context) error
}
