package util

func ArrayPrepend[T interface{}](arr []*T, elem *T) []*T {
	if elem == nil {
		return arr
	}
	return append([]*T{elem}, arr...)
}

func ArrayFirst[T interface{}](arr []*T) *T {
	if len(arr) == 0 {
		return nil
	}

	return arr[0]
}

func ArrayFilter[T interface{}](arr []T, predicate func(elem T) bool) []T {
	var result []T

	for _, el := range arr {
		if predicate(el) {
			result = append(result, el)
		}
	}

	return result
}

func ArrayContains[T comparable](arr []T, elem T) bool {
	for _, el := range arr {
		if el == elem {
			return true
		}
	}

	return false
}

func LocateArrayElement[T interface{}](arr []*T, test func(elem *T) bool) *T {
	for _, el := range arr {
		if test(el) {
			return el
		}
	}

	return nil
}
