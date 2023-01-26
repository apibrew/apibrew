package types

import (
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/structpb"
)

// string
type uuidType struct {
}

func (u uuidType) Equals(a, b interface{}) bool {
	return a == b
}

func (u uuidType) Pack(value interface{}) (*structpb.Value, error) {
	return structpb.NewValue(value.(uuid.UUID).String())
}

func (u uuidType) UnPack(val *structpb.Value) (interface{}, error) {
	return uuid.Parse(val.GetStringValue())
}

func (u uuidType) Default() any {
	return uuid.New()
}

func (u uuidType) Pointer(required bool) any {
	if required {
		return new(uuid.UUID)
	} else {
		return new(*uuid.UUID)
	}
}

func (u uuidType) String(val any) string {
	return val.(uuid.UUID).String()
}

func (u uuidType) IsEmpty(value any) bool {
	return value == nil
}

func (u uuidType) ValidatePackedValue(value *structpb.Value) error {
	err := canCast[string]("string", value.AsInterface())

	if err != nil {
		return err
	}

	_, err = uuid.Parse(value.GetStringValue())

	return err
}
