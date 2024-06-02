package types

import (
	"google.golang.org/protobuf/types/known/structpb"
	"strconv"
)

var Int64Type = int64Type{}

// float64
type int64Type struct {
}

func (i int64Type) Equals(a, b interface{}) bool {
	return a == b
}

func (i int64Type) Pack(value interface{}) (interface{}, error) {
	return structpb.NewValue(value)
}

func (i int64Type) UnPack(value interface{}) (interface{}, error) {
	return value.(int64), nil
}

func (i int64Type) Default() any {
	return int64(0)
}

func (i int64Type) Pointer(required bool) any {
	if required {
		return new(int64)
	} else {
		return new(*int64)
	}
}

func (i int64Type) String(val any) string {
	return strconv.Itoa(int(val.(int64)))
}

func (i int64Type) IsEmpty(value any) bool {
	return value == nil
}
