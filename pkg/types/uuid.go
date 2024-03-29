package types

import (
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/structpb"
)

var emptyUuid = uuid.UUID{}

var UuidType = uuidType{}

// string
type uuidType struct {
}

func (u uuidType) Equals(a, b interface{}) bool {
	return u.String(a) == u.String(b)
}

func (u uuidType) Pack(value interface{}) (*structpb.Value, error) {
	return structpb.NewValue(value.(uuid.UUID).String())
}

func (u uuidType) UnPack(val *structpb.Value) (interface{}, error) {
	return uuid.Parse(val.GetStringValue())
}

func (u uuidType) Default() any {
	return emptyUuid
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
