package yamlformat

import (
	"github.com/apibrew/apibrew/pkg/formats"
	"github.com/apibrew/apibrew/pkg/formats/unstructured"
	"github.com/apibrew/apibrew/pkg/model"
	log "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/encoding/protojson"
	"gopkg.in/yaml.v3"
	"io"
)

type writer struct {
	output             io.Writer
	hasMessageWritten  bool
	annotations        map[string]string
	unstructuredWriter unstructured.Writer
}

func (w *writer) WriteRecord(namespace string, resourceName string, records ...*model.Record) error {
	for _, record := range records {
		data, err := w.unstructuredWriter.WriteRecord(namespace, resourceName, record)

		if err != nil {
			return err
		}

		body, err := yaml.Marshal(data)

		if err != nil {
			return err
		}

		w.writePrefix()
		_, err = w.output.Write(body)
	}

	return nil
}

func (w *writer) WriteRecordsChan(resource *model.Resource, total uint32, recordsChan chan *model.Record) error {
	for record := range recordsChan {
		err := w.WriteRecord(resource.Namespace, resource.Name, record)

		if err != nil {
			return err
		}
	}

	return nil
}

func (w *writer) IsBinary() bool {
	return false
}

var jsonMo = protojson.MarshalOptions{
	Multiline:       true,
	EmitUnpopulated: false,
}

func (w *writer) WriteResource(resources ...*model.Resource) error {
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
	return &writer{output: output, annotations: annotations, unstructuredWriter: unstructured.Writer{Annotations: annotations}}
}
