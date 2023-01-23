package types

import (
	"strconv"
)

// float64
type int32Type struct {
}

func (i int32Type) Equals(a, b interface{}) bool {
	return a == b
}

func (i int32Type) Pack(value interface{}) (interface{}, error) {
	return value, nil
}

func (i int32Type) UnPack(value interface{}) (interface{}, error) {
	return value, nil
}

func (i int32Type) Default() any {
	return int32(0)
}

func (i int32Type) Pointer(required bool) any {
	if required {
		return new(int32)
	} else {
		return new(*int32)
	}
}

func (i int32Type) String(val any) string {
	return strconv.Itoa(int(val.(int32)))
}

func (i int32Type) IsEmpty(value any) bool {
	return value == nil
}

func (i int32Type) ValidatePackedValue(value any) error {
	return canCastNumber[int32]("int32", value)
}
