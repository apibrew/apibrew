package types

import "strconv"

type numericType struct {
}

func (n numericType) Pointer(required bool) any {
	if required {
		return new(int)
	} else {
		return new(*int)
	}
}

func (n numericType) String(val any) string {
	return strconv.Itoa(val.(int))
}

func (n numericType) IsEmpty(value any) bool {
	return value != nil
}

func (n numericType) ValidateValue(value any) error {
	return nil
}

func (n numericType) Default() any {
	return 0
}
