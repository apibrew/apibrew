package types

// string
type passwordType struct {
}

func (s passwordType) Equals(a, b interface{}) bool {
	return a == b
}

func (s passwordType) Pack(value interface{}) (interface{}, error) {
	return value, nil
}

func (s passwordType) UnPack(value interface{}) (interface{}, error) {
	return value, nil
}

func (s passwordType) Default() any {
	return ""
}

func (s passwordType) Pointer(required bool) any {
	if required {
		return new(string)
	} else {
		return new(*string)
	}
}

func (s passwordType) String(val any) string {
	return val.(string)
}

func (s passwordType) IsEmpty(value any) bool {
	return value == nil || value == ""
}

func (s passwordType) ValidatePackedValue(value any) error {
	return canCast[string]("string", value)
}
