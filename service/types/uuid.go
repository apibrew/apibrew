package types

import "github.com/google/uuid"

type uuidType struct {
}

func (u uuidType) Default() any {
	return uuid.New()
}

func (u uuidType) Pointer(required bool) any {
	if required {
		return new(uuid.UUID)
	} else {
		return new(*uuid.UUID)
	}
}

func (u uuidType) String(val any) string {
	return val.(uuid.UUID).String()
}

func (u uuidType) IsEmpty(value any) bool {
	return value == nil
}

func (u uuidType) ValidateValue(value any) error {
	return nil
}
