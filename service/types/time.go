package types

import "time"

type timeType struct {
}

func (t timeType) Pointer(required bool) any {
	return nil
}

func (t timeType) String(val any) string {
	return ""
}

func (t timeType) IsEmpty(value any) bool {
	return false
}

func (t timeType) ValidateValue(value any) error {
	return nil
}

func (t timeType) Default() any {
	return time.Now()
}
