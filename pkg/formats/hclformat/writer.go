package hclformat

import (
	"github.com/apibrew/apibrew/pkg/formats"
	"github.com/apibrew/apibrew/pkg/model"
	"io"
)

type writer struct {
	output      io.Writer
	batch       *model.Batch
	annotations map[string]string
}

func (w writer) WriteResource(resource ...*model.Resource) error {
	//TODO implement me
	panic("implement me")
}

func (w writer) WriteRecord(namespace string, resourceName string, record ...*model.Record) error {
	//TODO implement me
	panic("implement me")
}

func (w writer) WriteRecordsChan(resource *model.Resource, total uint32, recordsChan chan *model.Record) error {
	//TODO implement me
	panic("implement me")
}

func (w writer) IsBinary() bool {
	//TODO implement me
	panic("implement me")
}

func NewWriter(output io.Writer, annotations map[string]string) formats.Writer {
	return &writer{output: output, annotations: annotations}
}
