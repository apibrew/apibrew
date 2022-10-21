package types

import (
	"time"
)

type timestampType struct {
}

func (t timestampType) Pack(value interface{}) (interface{}, error) {
	return t.String(value), nil
}

func (t timestampType) UnPack(value interface{}) (interface{}, error) {
	return time.Parse(time.RFC3339, value.(string))
}

func (t timestampType) Pointer(required bool) any {
	if required {
		return new(time.Time)
	} else {
		return new(*time.Time)
	}
}

func (t timestampType) String(val any) string {
	return val.(string)
}

func (t timestampType) IsEmpty(value any) bool {
	return value == nil
}

func (t timestampType) ValidateValue(value any) error {
	return ValidateDateTime(value)
}

func (t timestampType) Default() any {
	return time.Now().Format(time.RFC3339)
}
