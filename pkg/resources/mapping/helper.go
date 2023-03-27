package mapping

func convertMap[T interface{}, K interface{}](annotations map[string]T, mapper func(v T) K) map[string]K {
	var result = make(map[string]K)

	for k, v := range annotations {
		result[k] = mapper(v)
	}

	return result
}
