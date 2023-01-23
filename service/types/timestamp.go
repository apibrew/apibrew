package types

import (
	"errors"
	"reflect"
	"time"
)

// string
type timestampType struct {
}

func (t timestampType) Equals(a, b interface{}) bool {
	return a == b
}

func (t timestampType) Pack(value interface{}) (interface{}, error) {
	return value.(time.Time).Format(time.RFC3339), nil
}

func (t timestampType) UnPack(value interface{}) (interface{}, error) {
	if str, ok := value.(string); ok {
		return time.Parse(time.RFC3339, str)
	} else {
		return nil, errors.New("wrong type: " + reflect.TypeOf(value).String())
	}
}

func (t timestampType) Pointer(required bool) any {
	if required {
		return new(time.Time)
	} else {
		return new(*time.Time)
	}
}

func (t timestampType) String(val any) string {
	return val.(time.Time).Format(time.RFC3339)
}

func (t timestampType) IsEmpty(value any) bool {
	return value == nil
}

func (t timestampType) ValidatePackedValue(value any) error {
	err := ValidateDateTime(value)

	if err != nil {
		return err
	}

	_, err = t.UnPack(value)

	return err
}

func (t timestampType) Default() any {
	return time.Now()
}
