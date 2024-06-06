package types

import (
	"time"
)

var TimestampType = timestampType{}

// string
type timestampType struct {
}

func (t timestampType) Equals(a, b interface{}) bool {
	return a.(time.Time).Equal(b.(time.Time))
}

func (t timestampType) Pack(value interface{}) (interface{}, error) {
	return value.(time.Time).Format(time.RFC3339), nil
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
	return val.(time.Time).Format(time.RFC3339)
}

func (t timestampType) IsEmpty(value any) bool {
	return value == nil
}

func (t timestampType) Default() any {
	return time.Now()
}
