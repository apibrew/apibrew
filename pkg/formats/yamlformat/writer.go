package yamlformat

import (
	"encoding/json"
	"github.com/apibrew/apibrew/pkg/formats"
	"github.com/apibrew/apibrew/pkg/model"
	log "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/encoding/protojson"
	"gopkg.in/yaml.v3"
	"io"
	"reflect"
	"sort"
)

type writer struct {
	output            io.Writer
	hasMessageWritten bool
	annotations       map[string]string
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

func (w *writer) isForApply() bool {
	return w.annotations != nil && w.annotations["for-apply"] == "true"
}

func (w *writer) WriteResource(resources ...*model.Resource) error {
	for _, resource := range resources {
		if w.isForApply() {
			resource = formats.FixResourceForApply(resource)
		}

		w.writePrefix()
		body, err := jsonMo.Marshal(resource)

		if err != nil {
			return err
		}

		var data map[string]interface{}

		err = json.Unmarshal(body, &data)

		if err != nil {
			return err
		}

		data = fixBeforeWrite(data).(map[string]interface{})

		data["type"] = "resource"

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

		if err != nil {
			return err
		}

		data = fixBeforeWrite(data).(map[string]interface{})

		data["type"] = "record"
		data["namespace"] = namespace
		data["resource"] = resourceName

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

func fixBeforeWrite(i interface{}) interface{} {
	switch x := i.(type) {
	case map[string]interface{}:
		m2 := map[string]interface{}{}
		var keys = make([]string, 0, len(x))
		for k := range x {
			keys = append(keys, k)
		}

		sort.Sort(sort.StringSlice(keys))

		for _, k := range keys {
			v := x[k]
			if IsZeroOfUnderlyingType(v) {
				continue
			}

			m2[k] = fixBeforeWrite(v)
		}
		return m2
	case []interface{}:
		for i, v := range x {
			x[i] = fixBeforeWrite(v)
		}
	}

	return i
}

func IsZeroOfUnderlyingType(x interface{}) bool {
	if x == nil {
		return true
	}
	return reflect.DeepEqual(x, reflect.Zero(reflect.TypeOf(x)).Interface())
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
	return &writer{output: output, annotations: annotations}
}
