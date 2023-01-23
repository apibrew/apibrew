package types

import "github.com/google/uuid"

// string
type uuidType struct {
}

func (u uuidType) Equals(a, b interface{}) bool {
	return a == b
}

func (u uuidType) Pack(value interface{}) (interface{}, error) {
	return value, nil
}

func (u uuidType) UnPack(val interface{}) (interface{}, error) {
	if _, ok := val.(string); ok {
		return uuid.Parse(val.(string))
	}

	return val, nil
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

func (u uuidType) ValidatePackedValue(value any) error {
	err := canCast[string]("string", value)

	if err != nil {
		return nil
	}

	_, err = uuid.Parse(value.(string))

	return err
}
