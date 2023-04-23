package yamlformat

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"github.com/tislib/apibrew/pkg/formats"
	"github.com/tislib/apibrew/pkg/model"
	"google.golang.org/protobuf/encoding/protojson"
	"gopkg.in/yaml.v3"
	"io"
)

type writer struct {
	output            io.Writer
	hasMessageWritten bool
}

var jsonMo = protojson.MarshalOptions{
	Multiline:       true,
	EmitUnpopulated: false,
}

func (w *writer) WriteResource(resources ...*model.Resource) error {
	for _, resource := range resources {
		w.writePrefix()
		body, err := jsonMo.Marshal(resource)

		if err != nil {
			return err
		}

		var data map[string]interface{}

		err = json.Unmarshal(body, &data)

		data["type"] = "resource"

		if err != nil {
			return err
		}

		out, err := yaml.Marshal(data)

		if err != nil {
			return err
		}

		_, err = w.output.Write(out)

		if err != nil {
			return err
		}
	}

	return nil
}

func (w *writer) WriteRecord(namespace string, resourceName string, records ...*model.Record) error {
	for _, record := range records {
		w.writePrefix()
		body, err := jsonMo.Marshal(record)

		if err != nil {
			return err
		}

		var data map[string]interface{}

		err = json.Unmarshal(body, &data)

		data["type"] = "record"
		data["namespace"] = namespace
		data["resource"] = resourceName

		if err != nil {
			return err
		}

		out, err := yaml.Marshal(data)

		if err != nil {
			return err
		}

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

func NewWriter(output io.Writer) formats.Writer {
	return &writer{output: output}
}
