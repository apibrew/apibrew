package output

import (
	"github.com/tislib/data-handler/pkg/model"
	"io"
	"log"
)

type Element struct {
	typ      string
	record   *model.Record
	resource *model.Resource
}

type Writer interface {
	WriteResources(resources []*model.Resource)
	WriteRecords(resource *model.Resource, recordsChan chan *model.Record)
	IsBinary() bool
}

func NewOutputReader(format string, w io.Writer) Writer {
	switch format {
	case "yaml":
		return &yamlReader{
			writer: w,
		}
	case "yml":
		return &yamlReader{
			writer: w,
		}
	case "pb":
		return &protobufReader{
			writer: w,
		}
	}

	log.Fatal("Writer not found: " + format)

	return nil
}
