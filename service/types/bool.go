package types

import (
	"errors"
	"fmt"
	"strconv"
)

type boolType struct {
}

func (u boolType) Default() any {
	return false
}

func (u boolType) Pointer(required bool) any {
	if required {
		return new(bool)
	} else {
		return new(bool)
	}
}

func (u boolType) String(val any) string {
	return strconv.FormatBool(u.typed(val))
}

func (u boolType) typed(val any) bool {
	return val.(bool)
}

func (u boolType) IsEmpty(value any) bool {
	return value == nil
}

func (u boolType) ValidateValue(value any) error {
	if _, ok := value.(bool); ok {
		return nil
	} else {
		return errors.New(fmt.Sprintf("value is not bool: %s", value))
	}
}
