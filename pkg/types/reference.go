package types

import (
	"fmt"
	"google.golang.org/protobuf/types/known/structpb"
)

var ReferenceType = referenceType{}

// string
type referenceType struct {
}

func (u referenceType) Equals(a, b interface{}) bool {
	return u.String(a) == u.String(b) //fixme
}

func (u referenceType) Pack(value interface{}) (*structpb.Value, error) {
	st, err := structpb.NewStruct(value.(map[string]interface{}))

	if err != nil {
		return nil, err
	}

	return structpb.NewStructValue(st), nil
}

func (u referenceType) UnPack(val *structpb.Value) (interface{}, error) {
	if val == nil {
		return nil, nil
	}

	return val.GetStructValue().AsMap(), nil
}

func (u referenceType) Default() any {
	return make(map[string]interface{})
}

func (u referenceType) Pointer(required bool) any {
	if required {
		return new(map[string]interface{})
	} else {
		return new(*map[string]interface{})
	}
}

func (u referenceType) String(val any) string {
	return fmt.Sprintf("%v", val.(map[string]interface{}))
}

func (u referenceType) IsEmpty(value any) bool {
	return value == nil
}

func (u referenceType) ValidatePackedValue(value *structpb.Value) error {
	return canCast[map[string]interface{}]("ReferenceType", value.AsInterface())
}
