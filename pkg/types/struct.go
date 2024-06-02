package types

import (
	"encoding/json"
)

var StructType = structType{}

// string
type structType struct {
}

func (o structType) Equals(a, b interface{}) bool {
	v1, _ := o.Serialize(a)
	v2, _ := o.Serialize(b)

	return v1 == v2
}

func (o structType) Pack(value interface{}) (interface{}, error) {
	return value, nil
}

func (o structType) Serialize(value interface{}) (interface{}, error) {
	data, err := json.Marshal(value)

	if err != nil {
		return nil, err
	}

	return string(data), nil
}

func (o structType) UnPack(value interface{}) (interface{}, error) {
	return value, nil
}

func (o structType) Pointer(required bool) any {
	if required {
		return new(string)
	} else {
		return new(*string)
	}
}

func (o structType) String(val any) string {
	return ""
}

func (o structType) IsEmpty(value any) bool {
	return value == nil
}

func (o structType) Default() any {
	return map[string]interface{}{}
}
