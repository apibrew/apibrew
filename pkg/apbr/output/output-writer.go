package output

import (
	"github.com/apibrew/apibrew/pkg/formats/unstructured"
	"github.com/apibrew/apibrew/pkg/formats/yamlformat"
	"github.com/apibrew/apibrew/pkg/resource_model"
	"io"
	"log"
)

type Writer interface {
	WriteResource(resource ...*resource_model.Resource) error
	WriteRecords(resource *resource_model.Resource, total uint32, records []unstructured.Unstructured) error
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
		var a = yamlformat.NewWriter(w, annotations)

		return a
	}

	log.Fatal("Writer not found: " + format)

	return nil
}
