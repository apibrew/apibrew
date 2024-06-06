package types

import (
	"encoding/json"
	"google.golang.org/protobuf/types/known/structpb"
)

var ListType = listType{}

// string
type listType struct {
}

func (o listType) Equals(a, b interface{}) bool {
	v1, _ := o.Serialize(a)
	v2, _ := o.Serialize(b)

	return v1 == v2
}

func (o listType) Pack(value interface{}) (interface{}, error) {
	return value, nil
}

func (o listType) Serialize(value interface{}) (interface{}, error) {
	data, err := json.Marshal(value)

	if err != nil {
		return nil, err
	}

	return string(data), nil
}

func (o listType) UnPack(value interface{}) (interface{}, error) {
	return value, nil
}

func (o listType) PackStruct(value interface{}) (*structpb.Value, error) {
	return structpb.NewValue(value)
}

func (o listType) UnPackStruct(value *structpb.Value) (interface{}, error) {
	return value.AsInterface(), nil
}

func (o listType) Pointer(required bool) any {
	if required {
		return new(string)
	} else {
		return new(*string)
	}
}

func (o listType) String(val any) string {
	return ""
}

func (o listType) IsEmpty(value any) bool {
	return value == nil
}

func (o listType) Default() any {
	return []interface{}{}
}
