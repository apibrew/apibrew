package writer

import (
	"github.com/apibrew/apibrew/pkg/formats"
	"github.com/apibrew/apibrew/pkg/formats/unstructured"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/resource_model"
	"reflect"
	"sort"
)

type Writer struct {
	Annotations map[string]string
}

func (w *Writer) isForApply() bool {
	return w.Annotations != nil && w.Annotations["for-apply"] == "true"
}

func (w *Writer) WriteRecord(namespace string, resourceName string, record unstructured.Unstructured) (unstructured.Unstructured, error) {
	record = fixBeforeWrite(record).(unstructured.Unstructured)

	record["type"] = "record"
	record["namespace"] = namespace
	record["resource"] = resourceName

	return record, nil
}

func (w *Writer) WriteResource(resource *resource_model.Resource) (unstructured.Unstructured, error) {
	if w.isForApply() {
		resource = formats.FixResourceForApply(resource)
	}

	var data = resource_model.ResourceMapperInstance.ToUnstructured(resource)

	data = fixBeforeWrite(data).(unstructured.Unstructured)

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

		sort.Strings(keys)

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

	if _, ok := x.(*model.ResourceProperty_Type); ok {
		return false
	}

	return reflect.DeepEqual(x, reflect.Zero(reflect.TypeOf(x)).Interface())
}
