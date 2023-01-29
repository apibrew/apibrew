package types

import (
	"google.golang.org/protobuf/types/known/structpb"
	"strconv"
)

// bool
type boolType struct {
}

func (u boolType) Equals(a, b interface{}) bool {
	return a == b
}

func (u boolType) Pack(value interface{}) (*structpb.Value, error) {
	return structpb.NewValue(value)
}

func (u boolType) UnPack(value *structpb.Value) (interface{}, error) {
	return value.GetBoolValue(), nil
}

func (u boolType) Default() any {
	return false
}

func (u boolType) Pointer(required bool) any {
	if required {
		return new(bool)
	} else {
		return new(bool)
	}
}

func (u boolType) String(val any) string {
	return strconv.FormatBool(u.typed(val))
}

func (u boolType) typed(val any) bool {
	return val.(bool)
}

func (u boolType) IsEmpty(value any) bool {
	return value == nil
}

func (u boolType) ValidatePackedValue(value *structpb.Value) error {
	return canCast[bool]("bool", value.AsInterface())
}
