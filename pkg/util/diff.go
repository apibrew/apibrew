package util

func ArrayDiffer[T interface{}](existing []T, updated []T, hasSameId func(a, b T) bool, isEqual func(a, b T) bool, onNew func(rec T) error, onUpdate func(e, u T) error, onDelete func(rec T) error) error {
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
