package output

import (
	"github.com/apibrew/apibrew/pkg/formats/yamlformat"
	"github.com/apibrew/apibrew/pkg/model"
	"io"
	"log"
)

type Writer interface {
	WriteResource(resource ...*model.Resource) error
	WriteRecords(resource *model.Resource, total uint32, records []*model.Record) error
	IsBinary() bool
}

func NewOutputWriter(format string, w io.Writer, annotations map[string]string) Writer {
	switch format {
	case "console":
		return &consoleWriter{
			writer:   w,
			describe: false,
		}
	case "yaml", "yml":
		return yamlformat.NewWriter(w, annotations)
	}

	log.Fatal("Writer not found: " + format)

	return nil
}
