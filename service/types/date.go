package types

import "time"

type dateType struct {
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
	return (val.(time.Time)).Format(time.RFC3339)
}

func (u dateType) IsEmpty(value any) bool {
	return value == nil
}

func (u dateType) ValidateValue(value any) error {
	return nil
}
