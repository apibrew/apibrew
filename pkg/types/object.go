package types

import (
	"encoding/json"
	"google.golang.org/protobuf/types/known/structpb"
)

// string
type objectType struct {
}

func (o objectType) Equals(a, b interface{}) bool {
	v1, _ := o.Serialize(a)
	v2, _ := o.Serialize(b)

	return v1 == v2
}

func (o objectType) Pack(value interface{}) (*structpb.Value, error) {
	return structpb.NewValue(value)
}

func (o objectType) Serialize(value interface{}) (interface{}, error) {
	data, err := json.Marshal(value)

	if err != nil {
		return nil, err
	}

	return string(data), nil
}

func (o objectType) UnPack(value *structpb.Value) (interface{}, error) {
	return value.AsInterface(), nil
}

func (o objectType) Pointer(required bool) any {
	if required {
		return new(string)
	} else {
		return new(*string)
	}
}

func (o objectType) String(val any) string {
	return ""
}

func (o objectType) IsEmpty(value any) bool {
	return value == nil
}

func (o objectType) ValidatePackedValue(value *structpb.Value) error {
	return nil
}

func (o objectType) Default() any {
	return nil
}
