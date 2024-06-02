package types

import (
	"fmt"
	"google.golang.org/protobuf/types/known/structpb"
)

var Float32Type = float32Type{}

// float32
type float32Type struct {
}

func (f float32Type) Equals(a, b interface{}) bool {
	return a == b
}

func (f float32Type) Pack(value interface{}) (interface{}, error) {
	return structpb.NewValue(value)
}

func (f float32Type) UnPack(value interface{}) (interface{}, error) {
	return value.(float32), nil
}

func (f float32Type) Pointer(required bool) any {
	if required {
		return new(float32)
	} else {
		return new(*float32)
	}
}

func (f float32Type) String(val any) string {
	return fmt.Sprintf("%f", val)
}

func (f float32Type) IsEmpty(value any) bool {
	return value == nil
}

func (f float32Type) Default() any {
	return float32(0)
}
