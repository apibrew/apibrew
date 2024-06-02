package types

import (
	"google.golang.org/protobuf/types/known/structpb"
	"time"
)

var TimeType = timeType{}

// string
type timeType struct {
}

func (t timeType) Equals(a, b interface{}) bool {
	return a.(time.Time).Equal(b.(time.Time))
}

func (t timeType) Pack(value interface{}) (interface{}, error) {
	return structpb.NewValue(value.(time.Time).Format("15:04:05"))
}

func (t timeType) UnPack(value interface{}) (interface{}, error) {
	return time.Parse("15:04:05", value.(string))
}

func (t timeType) Pointer(required bool) any {
	if required {
		return new(time.Time)
	} else {
		return new(*time.Time)
	}
}

func (t timeType) String(val any) string {
	return val.(time.Time).Format("15:04:05")
}

func (t timeType) IsEmpty(value any) bool {
	return value == nil
}

func (t timeType) Default() any {
	return time.Now()
}
