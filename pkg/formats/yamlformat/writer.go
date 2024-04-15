package yamlformat

import (
	"github.com/apibrew/apibrew/pkg/formats"
	"github.com/apibrew/apibrew/pkg/formats/unstructured"
	writer2 "github.com/apibrew/apibrew/pkg/formats/writer"
	"github.com/apibrew/apibrew/pkg/resource_model"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
	"io"
)

type writer struct {
	output             io.Writer
	hasMessageWritten  bool
	annotations        map[string]string
	unstructuredWriter writer2.Writer
}

func (w *writer) WriteRecord(typ string, records ...unstructured.Unstructured) error {
	for _, record := range records {
		data, err := w.unstructuredWriter.WriteRecord2(typ, record)

		if err != nil {
			return err
		}

		body, err := yaml.Marshal(data)

		if err != nil {
			return err
		}

		w.writePrefix()
		_, err = w.output.Write(body)

		if err != nil {
			return err
		}
	}

	return nil
}

func (w *writer) WriteRecords(resource *resource_model.Resource, total uint32, records []unstructured.Unstructured) error {
	for _, record := range records {
		var typ = resource.Namespace.Name + "/" + resource.Name

		if resource.Namespace.Name == "" || resource.Namespace.Name == "default" {
			typ = resource.Name
		}

		err := w.WriteRecord(typ, record)

		if err != nil {
			return err
		}
	}

	return nil
}

func (w *writer) IsBinary() bool {
	return false
}

func (w *writer) WriteResource(resources ...*resource_model.Resource) error {
	for _, resource := range resources {
		data, err := w.unstructuredWriter.WriteResource(resource)

		if err != nil {
			return err
		}

		out, err := yaml.Marshal(data)

		if err != nil {
			return err
		}

		w.writePrefix()

		_, err = w.output.Write(out)

		if err != nil {
			return err
		}
	}

	return nil
}

func (w *writer) writePrefix() {
	if w.hasMessageWritten {
		if _, err := w.output.Write([]byte("---\n")); err != nil {
			log.Fatal(err)
		}
	}

	w.hasMessageWritten = true
}

func NewWriter(output io.Writer, annotations map[string]string) formats.Writer {
	return &writer{output: output, annotations: annotations, unstructuredWriter: writer2.Writer{Annotations: annotations}}
}
