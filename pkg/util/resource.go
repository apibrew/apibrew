package util

type Named interface {
	GetName() string
}

func GetNamedMap[T Named](items []T) map[string]T {
	var result = make(map[string]T)

	for _, prop := range items {
		result[prop.GetName()] = prop
	}

	return result
}

func GetArrayIndex[T comparable](items []T, item T, comparator func(a, b T) bool) int {
	for i, elem := range items {
		if comparator(elem, item) {
			return i
		}
	}

	return -1
}
