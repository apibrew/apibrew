package ops

import "github.com/apibrew/apibrew/pkg/formats/unstructured"

func ParseRef(un unstructured.Unstructured, path string) (unstructured.Unstructured, error) {
	data, err := JsonPathRead(un, path)

	if err != nil {
		return nil, err
	}

	return data.(unstructured.Unstructured), nil
}
