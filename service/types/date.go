package types

import (
	"time"
)

type dateType struct {
}

func (u dateType) Pack(value interface{}) (interface{}, error) {
	return u.String(value), nil
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

func (u dateType) ValidateValue(value any) error {
	err := canCast[string]("string", value)

	if err != nil {
		return nil
	}

	_, err = u.UnPack(value)

	return err
}
