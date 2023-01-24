package types

import (
	"fmt"
	"google.golang.org/protobuf/types/known/structpb"
)

// float64
type doubleType struct {
}

func (d doubleType) Equals(a, b interface{}) bool {
	return a == b
}

func (d doubleType) Pack(value interface{}) (*structpb.Value, error) {
	return structpb.NewValue(value)
}

func (d doubleType) UnPack(value *structpb.Value) (interface{}, error) {
	return value.GetNumberValue(), nil
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

func (d doubleType) ValidatePackedValue(value *structpb.Value) error {
	return canCastNumber[float64]("float64", value.AsInterface())
}

func (d doubleType) Default() any {
	return float64(0)
}
