package types

import (
	"fmt"
	"google.golang.org/protobuf/types/known/structpb"
)

// float32
type floatType struct {
}

func (f floatType) Equals(a, b interface{}) bool {
	return a == b
}

func (f floatType) Pack(value interface{}) (*structpb.Value, error) {
	return structpb.NewValue(value)
}

func (f floatType) UnPack(value *structpb.Value) (interface{}, error) {
	return float32(value.GetNumberValue()), nil
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

func (f floatType) ValidatePackedValue(value *structpb.Value) error {
	return canCastNumber[float32]("float32", value.AsInterface())
}

func (f floatType) Default() any {
	return float32(0)
}
