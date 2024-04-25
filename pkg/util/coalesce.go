package util

func Coalesce[T interface{}](val ...*T) *T {
	for _, item := range val {
		if item != nil {
			return item
		}
	}

	return nil
}

func CoalesceThen[T interface{}](fn func(val *T) error, val ...*T) error {
	for _, item := range val {
		if item != nil {
			return fn(item)
		}
	}

	return nil
}
