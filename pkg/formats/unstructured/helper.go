package unstructured

import (
	"fmt"
	"reflect"
)

func WalkUnstructured(body interface{}, visitor func(value interface{}) (interface{}, error)) (interface{}, error) {
	switch x := body.(type) {
	case Unstructured:
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

func fixMaps(i interface{}) interface{} {
	switch x := i.(type) {
	case map[interface{}]interface{}:
		m2 := Unstructured{}
		for k, v := range x {
			// TODO: check if key is string
			m2[k.(string)] = fixMaps(v)
		}
		return m2
	case map[string]interface{}:
		m2 := Unstructured{}
		for k, v := range x {
			// TODO: check if key is string
			m2[k] = fixMaps(v)
		}
		return m2
	case Unstructured:
		m2 := Unstructured{}
		for k, v := range x {
			m2[k] = fixMaps(v)
		}
		return m2
	case []interface{}:
		for i, v := range x {
			x[i] = fixMaps(v)
		}
	}
	return i
}
