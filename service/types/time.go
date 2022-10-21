package types

import "time"

type timeType struct {
}

func (t timeType) Pointer(required bool) any {
	if required {
		return new(time.Time)
	} else {
		return new(*time.Time)
	}
}

func (t timeType) String(val any) string {
	return ""
}

func (t timeType) IsEmpty(value any) bool {
	return value == nil
}

func (t timeType) ValidateValue(value any) error {
	err := canCast[string]("string", value)

	if err != nil {
		return nil
	}

	_, err = time.Parse("15:04:05", value.(string))

	return err
}

func (t timeType) Default() any {
	return time.Now()
}
