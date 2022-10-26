package util

func Combine[T interface{}](left []T, right []T) []T {
	var result = left

	for _, item := range right {
		result = append(result, item)
	}

	return result
}
