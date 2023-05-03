package util

import "github.com/apibrew/apibrew/pkg/errors"

func Coalesce[T interface{}](val ...*T) *T {
	for _, item := range val {
		if item != nil {
			return item
		}
	}

	return nil
}

func CoalesceThen[T interface{}](fn func(val *T) errors.ServiceError, val ...*T) errors.ServiceError {
	for _, item := range val {
		if item != nil {
			return fn(item)
		}
	}

	return nil
}
