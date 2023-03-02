package util

func ArrToChan[T interface{}](arr []T) chan T {
	ch := make(chan T, len(arr))

	for _, item := range arr {
		ch <- item
	}
	close(ch)

	return ch
}
