package util

import "github.com/apibrew/apibrew/pkg/errors"

func ArrayDiffer[T interface{}](existing []T, updated []T, hasSameId func(a, b T) bool, isEqual func(a, b T) bool, onNew func(rec T) errors.ServiceError, onUpdate func(e, u T) errors.ServiceError, onDelete func(rec T) errors.ServiceError) errors.ServiceError {
	// fixme do not match already matched items
	var passedToUpdated []T

	var isUpdated = func(u T) bool {
		for _, e := range passedToUpdated {
			if hasSameId(e, u) {
				return true
			}
		}

		return false
	}

	for _, e := range existing {
		found := false
		for _, u := range updated {
			if hasSameId(e, u) {
				if !isEqual(e, u) && !isUpdated(u) {
					passedToUpdated = append(passedToUpdated, u)
					err := onUpdate(e, u)

					if err != nil {
						return err
					}
				}

				found = true
				break
			}
		}

		if !found {
			err := onDelete(e)

			if err != nil {
				return err
			}
		}
	}

	for _, u := range updated {
		found := false
		for _, e := range existing {

			if hasSameId(e, u) {
				if !isEqual(e, u) && !isUpdated(e) {
					err := onUpdate(e, u)

					if err != nil {
						return err
					}
				}

				found = true
				break
			}
		}

		if !found {
			err := onNew(u)

			if err != nil {
				return err
			}
		}
	}

	return nil
}

func MapDiffer[T interface{}](existing map[string]T, updated map[string]T, hasSameId func(a, b T) bool, isEqual func(a, b T) bool, onNew func(rec MapEntry[T]) errors.ServiceError, onUpdate func(e, u MapEntry[T]) errors.ServiceError, onDelete func(rec MapEntry[T]) errors.ServiceError) errors.ServiceError {
	return ArrayDiffer(MapToList(existing), MapToList(updated), func(a, b MapEntry[T]) bool {
		return a.Key == b.Key || hasSameId(a.Value, b.Value)
	}, func(a, b MapEntry[T]) bool {
		return a.Key == b.Key && isEqual(a.Value, b.Value)
	}, func(rec MapEntry[T]) errors.ServiceError {
		return onNew(rec)
	}, func(e, u MapEntry[T]) errors.ServiceError {
		return onUpdate(e, u)
	}, func(rec MapEntry[T]) errors.ServiceError {
		return onDelete(rec)
	})
}

type MapEntry[T interface{}] struct {
	Key   string
	Value T
}

func MapToList[T interface{}](input map[string]T) []MapEntry[T] {
	var result []MapEntry[T]

	for key, value := range input {
		result = append(result, MapEntry[T]{Key: key, Value: value})
	}

	return result
}
