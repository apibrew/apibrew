package test2

import "testing"

func check(err error, t *testing.T) {
	if err != nil {
		if t != nil {
			t.Error(err)
		} else {
			panic(err)
		}
	}
}
