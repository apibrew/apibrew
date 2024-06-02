package types

import (
	"google.golang.org/protobuf/types/known/structpb"
	"time"
)

var DateType = dateType{}

// string
type dateType struct {
}

func (u dateType) Equals(a, b interface{}) bool {
	return a.(time.Time).Equal(b.(time.Time))
}

func (u dateType) Pack(value interface{}) (interface{}, error) {
	return structpb.NewValue(u.String(value))
}

func (u dateType) UnPack(value interface{}) (interface{}, error) {
	return time.Parse("2006-01-02", value.(string))
}

func (u dateType) Default() any {
	return time.Now()
}

func (u dateType) Pointer(required bool) any {
	if required {
		return new(time.Time)
	} else {
		return new(*time.Time)
	}
}

func (u dateType) String(val any) string {
	return val.(time.Time).Format("2006-01-02")
}

func (u dateType) IsEmpty(value any) bool {
	return value == nil
}
