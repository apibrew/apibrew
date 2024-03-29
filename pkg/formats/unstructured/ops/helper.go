package ops

import (
	"fmt"
	"github.com/apibrew/apibrew/pkg/formats/unstructured"
	"reflect"
)

func WalkUnstructured(in interface{}, visitor func(value interface{}) (interface{}, error)) (interface{}, error) {
	body, err := visitor(in)

	if err != nil {
		return nil, err
	}

	switch x := body.(type) {
	case unstructured.Unstructured:
		for key, value := range x {
			newVal, err := WalkUnstructured(value, visitor)
			if err != nil {
				return nil, err
			}

			visitedVal, err := visitor(newVal)
			if err != nil {
				return nil, err
			}

			x[key] = visitedVal
		}
	case []unstructured.Unstructured:
		for i, value := range x {
			newVal, err := WalkUnstructured(value, visitor)
			if err != nil {
				return nil, err
			}

			visitedVal, err := visitor(newVal)
			if err != nil {
				return nil, err
			}

			x[i] = visitedVal.(unstructured.Unstructured)
		}
	case []interface{}:
		for i, value := range x {
			newVal, err := WalkUnstructured(value, visitor)
			if err != nil {
				return nil, err
			}

			visitedVal, err := visitor(newVal)
			if err != nil {
				return nil, err
			}

			x[i] = visitedVal
		}
	case interface{}:
		return visitor(x)
	case nil:
		return nil, nil
	default:
		fmt.Print(reflect.TypeOf(x))
		panic("unknown type: ")
	}

	return body, nil
}
