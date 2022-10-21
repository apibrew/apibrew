package types

import "time"

type timestampType struct {
}

func (t timestampType) Pointer(required bool) any {
	return nil
}

func (t timestampType) String(val any) string {
	return ""
}

func (t timestampType) IsEmpty(value any) bool {
	return false
}

func (t timestampType) ValidateValue(value any) error {
	return nil
}

func (t timestampType) Default() any {
	return time.Now()
}
