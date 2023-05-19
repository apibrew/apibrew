package output

import (
	"github.com/apibrew/apibrew/pkg/formats/batch"
	"github.com/apibrew/apibrew/pkg/formats/hclformat"
	yamlformat "github.com/apibrew/apibrew/pkg/formats/yaml"
	"github.com/apibrew/apibrew/pkg/model"
	"io"
	"log"
)

type Writer interface {
	WriteResource(resource ...*model.Resource) error
	WriteRecordsChan(resource *model.Resource, total uint32, recordsChan chan *model.Record) error
	IsBinary() bool
}

func NewOutputWriter(format string, w io.Writer, annotations map[string]string) Writer {
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
	case "yaml", "yml":
		return yamlformat.NewWriter(w, annotations)
	case "hcl":
		return hclformat.NewWriter(w, annotations)
	case "pb":
		return &protobufWriter{
			batchWriter: batch.NewWriter(w),
		}
	}

	log.Fatal("Writer not found: " + format)

	return nil
}
