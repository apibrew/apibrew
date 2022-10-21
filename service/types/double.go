package types

import "fmt"

type doubleType struct {
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

func (d doubleType) ValidateValue(value any) error {
	return nil
}

func (d doubleType) Default() any {
	return float64(0)
}
