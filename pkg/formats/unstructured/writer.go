package unstructured

import (
	"encoding/json"
	"github.com/apibrew/apibrew/pkg/formats"
	"github.com/apibrew/apibrew/pkg/model"
	"google.golang.org/protobuf/encoding/protojson"
	"reflect"
	"sort"
)

type Writer struct {
	Annotations map[string]string
}

var jsonMo = protojson.MarshalOptions{
	Multiline:       true,
	EmitUnpopulated: false,
}

func (w *Writer) isForApply() bool {
	return w.Annotations != nil && w.Annotations["for-apply"] == "true"
}

func (w *Writer) WriteRecord(namespace string, resourceName string, record *model.Record) (Unstructured, error) {
	body, err := jsonMo.Marshal(record)

	if err != nil {
		return nil, err
	}

	var data Unstructured

	err = json.Unmarshal(body, &data)

	if err != nil {
		return nil, err
	}

	data = fixBeforeWrite(data).(Unstructured)

	data = fixMaps(data).(Unstructured)

	data["type"] = "record"
	data["namespace"] = namespace
	data["resource"] = resourceName

	return data, nil
}

func (w *Writer) WriteResource(resource *model.Resource) (Unstructured, error) {
	if w.isForApply() {
		resource = formats.FixResourceForApply(resource)
	}

	body, err := jsonMo.Marshal(resource)

	if err != nil {
		return nil, err
	}

	var data Unstructured

	err = json.Unmarshal(body, &data)

	if err != nil {
		return nil, err
	}

	data = fixBeforeWrite(data).(Unstructured)

	data["type"] = "resource"

	return data, nil
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
