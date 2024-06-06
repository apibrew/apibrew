package types

import (
	"fmt"
	"strconv"
)

var Int32Type = int32Type{}

// float64
type int32Type struct {
}

func (i int32Type) Equals(a, b interface{}) bool {
	return a == b
}

func (i int32Type) Pack(value interface{}) (interface{}, error) {
	return convertToInt32(value)
}

func (i int32Type) UnPack(value interface{}) (interface{}, error) {
	return convertToInt32(value)
}

func (i int32Type) Default() any {
	return int32(0)
}

func (i int32Type) Pointer(required bool) any {
	if required {
		return new(int32)
	} else {
		return new(*int32)
	}
}

func (i int32Type) String(val any) string {
	return strconv.Itoa(val.(int))
}

func (i int32Type) IsEmpty(value any) bool {
	return value == nil
}

func convertToInt32(val interface{}) (int32, error) {
	switch v := val.(type) {
	case int:
		return int32(v), nil
	case int8:
		return int32(v), nil
	case int16:
		return int32(v), nil
	case int32:
		return int32(v), nil
	case int64:
		return int32(v), nil
	case uint:
		return int32(v), nil
	case uint8:
		return int32(v), nil
	case uint16:
		return int32(v), nil
	case uint32:
		return int32(v), nil
	case uint64:
		return int32(v), nil
	case float32:
		return int32(v), nil
	case float64:
		return int32(v), nil
	default:
		return 0, fmt.Errorf("unsupported type: %T", v)
	}
}
