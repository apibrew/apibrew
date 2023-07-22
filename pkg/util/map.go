package util

import (
	"github.com/apibrew/apibrew/pkg/errors"
)

func FlatMap[T interface{}](arr ...[]T) []T {
	var list []T

	for _, item := range arr {
		list = append(list, item...)
	}

	return list
}

func ArrayMap[T interface{}, R interface{}](arr []T, mapper func(T) R) []R {
	var list = make([]R, 0)

	for _, item := range arr {
		list = append(list, mapper(item))
	}

	return list
}

type HasId interface {
	GetId() string
}

func ArrayMapToId[T HasId](arr []T) []string {
	return ArrayMap(arr, func(t T) string {
		return t.GetId()
	})
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

func ArrayMapToString[T interface{}](arr []T, fn func(t T) string) []string {
	return ArrayMap(arr, fn)
}
