package types

import (
	"fmt"
)

// float64
type floatType struct {
}

func (f floatType) Equals(a, b interface{}) bool {
	return a == b
}

func (f floatType) Pack(value interface{}) (interface{}, error) {
	return value, nil
}

func (f floatType) UnPack(value interface{}) (interface{}, error) {
	return value, nil
}

func (f floatType) Pointer(required bool) any {
	if required {
		return new(float32)
	} else {
		return new(*float32)
	}
}

func (f floatType) String(val any) string {
	return fmt.Sprintf("%f", val)
}

func (f floatType) IsEmpty(value any) bool {
	return value == nil
}

func (f floatType) ValidateValue(value any) error {
	return canCastNumber[float32]("float32", value)
}

func (f floatType) Default() any {
	return float32(0)
}
