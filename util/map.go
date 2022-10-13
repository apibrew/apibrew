package util

func ArrayMap[T interface{}, R interface{}](arr []T, mapper func(T) R) []R {
	var list []R

	for _, item := range arr {
		list = append(list, mapper(item))
	}

	return list
}

func ArrayMapToInterface[T interface{}](arr []T) []interface{} {
	return ArrayMap(arr, func(t T) interface{} {
		return t
	})
}
