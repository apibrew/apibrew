package types

import (
	"google.golang.org/protobuf/types/known/structpb"
	"time"
)

// string
type timeType struct {
}

func (t timeType) Equals(a, b interface{}) bool {
	return a == b
}

func (t timeType) Pack(value interface{}) (*structpb.Value, error) {
	return structpb.NewValue(value.(time.Time).Format("15:04:05"))
}

func (t timeType) UnPack(value *structpb.Value) (interface{}, error) {
	return time.Parse("15:04:05", value.GetStringValue())
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

func (t timeType) ValidatePackedValue(value *structpb.Value) error {
	err := canCast[string]("string", value.AsInterface())

	if err != nil {
		return err
	}

	_, err = time.Parse("15:04:05", value.GetStringValue())

	return err
}

func (t timeType) Default() any {
	return time.Now()
}
