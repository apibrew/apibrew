package types

type bytesType struct {
}

func (b bytesType) Pointer(required bool) any {
	if required {
		return new([]byte)
	} else {
		return new(*[]byte)
	}
}

func (b bytesType) String(val any) string {
	return string(val.([]byte))
}

func (b bytesType) IsEmpty(val any) bool {
	return val == nil || len(val.(string)) == 0
}

func (b bytesType) ValidateValue(value any) error {
	return canCast[string]("base64", value)
}

func (b bytesType) Default() any {
	return []byte("")
}
