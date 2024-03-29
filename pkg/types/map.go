package types

import (
	"encoding/json"
	"google.golang.org/protobuf/types/known/structpb"
)

var MapType = mapType{}

// string
type mapType struct {
}

func (o mapType) Equals(a, b interface{}) bool {
	v1, _ := o.Serialize(a)
	v2, _ := o.Serialize(b)

	return v1 == v2
}

func (o mapType) Pack(value interface{}) (*structpb.Value, error) {
	return structpb.NewValue(value)
}

func (o mapType) Serialize(value interface{}) (interface{}, error) {
	data, err := json.Marshal(value)

	if err != nil {
		return nil, err
	}

	return string(data), nil
}

func (o mapType) UnPack(value *structpb.Value) (interface{}, error) {
	return value.AsInterface(), nil
}

func (o mapType) Pointer(required bool) any {
	if required {
		return new(string)
	} else {
		return new(*string)
	}
}

func (o mapType) String(val any) string {
	return ""
}

func (o mapType) IsEmpty(value any) bool {
	return value == nil
}

func (o mapType) Default() any {
	return map[string]interface{}{}
}
