package output

import (
	"github.com/tislib/data-handler/pkg/formats/batch"
	"github.com/tislib/data-handler/pkg/model"
	"io"
	"log"
)

type Writer interface {
	WriteResources(resources []*model.Resource)
	WriteRecords(resource *model.Resource, total uint32, recordsChan chan *model.Record)
	IsBinary() bool
}

func NewOutputWriter(format string, w io.Writer) Writer {
	switch format {
	case "console":
		return &consoleWriter{
			writer:   w,
			describe: false,
		}
	case "describe":
		return &consoleWriter{
			writer:   w,
			describe: true,
		}
	case "yaml":
		return &yamlWriter{
			writer: w,
		}
	case "yml":
		return &yamlWriter{
			writer: w,
		}
	case "pb":
		return &protobufWriter{
			batchWriter: batch.NewWriter(w),
		}
	}

	log.Fatal("Writer not found: " + format)

	return nil
}
