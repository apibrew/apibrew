package types

import (
	"fmt"
	"google.golang.org/protobuf/types/known/structpb"
)

// float64
type float64Type struct {
}

func (d float64Type) Equals(a, b interface{}) bool {
	return a == b
}

func (d float64Type) Pack(value interface{}) (*structpb.Value, error) {
	return structpb.NewValue(value)
}

func (d float64Type) UnPack(value *structpb.Value) (interface{}, error) {
	return value.GetNumberValue(), nil
}

func (d float64Type) Pointer(required bool) any {
	if required {
		return new(float64)
	} else {
		return new(*float64)
	}
}

func (d float64Type) String(val any) string {
	return fmt.Sprintf("%f", val)
}

func (d float64Type) IsEmpty(value any) bool {
	return value == nil
}

func (d float64Type) ValidatePackedValue(value *structpb.Value) error {
	return canCastNumber[float64]("float64", value.AsInterface())
}

func (d float64Type) Default() any {
	return float64(0)
}
