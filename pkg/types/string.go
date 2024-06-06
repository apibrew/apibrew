package types

import "google.golang.org/protobuf/types/known/structpb"

var StringType = stringType{}

// string
type stringType struct {
}

func (s stringType) Equals(a, b interface{}) bool {
	if _, ok := a.(string); !ok {
		return false
	}
	if _, ok := b.(string); !ok {
		return false
	}
	return a == b
}

func (s stringType) Pack(value interface{}) (interface{}, error) {
	return value, nil
}

func (s stringType) UnPack(value interface{}) (interface{}, error) {
	return value, nil
}

func (s stringType) PackStruct(value interface{}) (*structpb.Value, error) {
	return structpb.NewValue(value)
}

func (s stringType) UnPackStruct(value *structpb.Value) (interface{}, error) {
	return value.GetStringValue(), nil
}

func (s stringType) Default() any {
	return ""
}

func (s stringType) Pointer(required bool) any {
	if required {
		return new(string)
	} else {
		return new(*string)
	}
}

func (s stringType) String(val any) string {
	return val.(string)
}

func (s stringType) IsEmpty(value any) bool {
	return value == nil || value == ""
}
