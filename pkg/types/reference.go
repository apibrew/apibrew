package types

import (
	"fmt"
)

var ReferenceType = referenceType{}

// string
type referenceType struct {
}

func (u referenceType) Equals(a, b interface{}) bool {
	return u.String(a) == u.String(b) //fixme
}

func (u referenceType) Pack(value interface{}) (interface{}, error) {
	return value.(map[string]interface{}), nil
}

func (u referenceType) UnPack(val interface{}) (interface{}, error) {
	if val == nil {
		return nil, nil
	}

	if valueStr, ok := val.(string); ok {
		return map[string]interface{}{
			"id": valueStr,
		}, nil
	}

	return val, nil
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
