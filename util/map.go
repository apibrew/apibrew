package util

import (
	"github.com/tislib/data-handler/service/errors"
)

func ArrayMap[T interface{}, R interface{}](arr []T, mapper func(T) R) []R {
	var list []R

	for _, item := range arr {
		list = append(list, mapper(item))
	}

	return list
}

func ArrayMapWithError[T interface{}, R interface{}](arr []T, mapper func(T) (R, errors.ServiceError)) ([]R, errors.ServiceError) {
	var list []R

	for _, item := range arr {
		mappedItem, err := mapper(item)
		if err != nil {
			return nil, err
		}
		list = append(list, mappedItem)
	}

	return list, nil
}

func ArrayMapString(arr []string, mapper func(string) string) []string {
	return ArrayMap[string, string](arr, mapper)
}

func ArrayMapToInterface[T interface{}](arr []T) []interface{} {
	return ArrayMap(arr, func(t T) interface{} {
		return t
	})
}
