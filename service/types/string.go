package types

type stringType struct {
}

func (s stringType) Default() any {
	return ""
}

func (s stringType) Pointer(required bool) any {
	if required {
		return new(string)
	} else {
		return new(*string)
	}
}

func (s stringType) String(val any) string {
	return val.(string)
}

func (s stringType) IsEmpty(value any) bool {
	return value == nil || value == ""
}

func (s stringType) ValidateValue(value any) error {
	return canCast[string]("string", value)
}
