package types

import (
	"strconv"
)

type int64Type struct {
}

func (i int64Type) Default() any {
	return int64(0)
}

func (i int64Type) Pointer(required bool) any {
	if required {
		return new(int64)
	} else {
		return new(*int64)
	}
}

func (i int64Type) String(val any) string {
	return strconv.Itoa(int(val.(int64)))
}

func (i int64Type) IsEmpty(value any) bool {
	return value == nil
}

func (i int64Type) ValidateValue(value any) error {
	return canCastNumber[int64]("int64", value)
}
