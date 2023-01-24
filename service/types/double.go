package types

import "fmt"

// float64
type doubleType struct {
}

func (d doubleType) Equals(a, b interface{}) bool {
	return a == b
}

func (d doubleType) Pack(value interface{}) (interface{}, error) {
	return value, nil
}

func (d doubleType) UnPack(value interface{}) (interface{}, error) {
	return value, nil
}

func (d doubleType) Pointer(required bool) any {
	if required {
		return new(float64)
	} else {
		return new(*float64)
	}
}

func (d doubleType) String(val any) string {
	return fmt.Sprintf("%f", val)
}

func (d doubleType) IsEmpty(value any) bool {
	return value == nil
}

func (d doubleType) ValidatePackedValue(value any) error {
	return canCastNumber[float64]("float64", value)
}

func (d doubleType) Default() any {
	return float64(0)
}
