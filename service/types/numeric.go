package types

import (
	"fmt"
)

type numericType struct {
}

func (n numericType) Pack(value interface{}) (interface{}, error) {
	return value, nil
}

func (n numericType) UnPack(value interface{}) (interface{}, error) {
	return value, nil
}

func (n numericType) Pointer(required bool) any {
	if required {
		return new(int)
	} else {
		return new(*int)
	}
}

func (n numericType) String(val any) string {
	return fmt.Sprintf("%f", val)
}

func (n numericType) IsEmpty(value any) bool {
	return value == nil
}

func (n numericType) ValidateValue(value any) error {
	return canCastNumber[float64]("float64", value)
}

func (n numericType) Default() any {
	return float64(0)
}
