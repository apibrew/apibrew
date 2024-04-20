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

type MapEntry[T interface{}] struct {
	Key string
	Val T
}

func MapToArray[T interface{}](arr map[string]T) []MapEntry[T] {
	var list = make([]MapEntry[T], 0)

	for key, item := range arr {
		list = append(list, MapEntry[T]{
			Key: key,
			Val: item,
		})
	}

	return list
}

func ArrayMapX[T interface{}, R interface{}](arr []*T, mapper func(*T) *R) []*R {
	var list = make([]*R, 0)

	for _, item := range arr {
		if arr != nil {
			list = append(list, mapper(item))
		} else {
			list = append(list, nil)
		}
	}

	return list
}

func ArrayToMap[T interface{}, R interface{}, K comparable](arr []T, keyFunc func(T) K, valueFunc func(T) R) map[K]R {
	var result = make(map[K]R, 0)

	for _, item := range arr {
		result[keyFunc(item)] = valueFunc(item)

	}

	return result
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

func Keys[T any](u map[string]T) []string {
	var keys []string

	for key := range u {
		keys = append(keys, key)
	}

	return keys
}
