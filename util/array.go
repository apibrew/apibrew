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

func ArrayCut[T interface{}](arr []T, index int) []T {
	if len(arr) <= index+1 {
		return nil
	}

	return arr[index+1:]
}
