package types

type objectType struct {
}

func (o objectType) Pointer(required bool) any {
	return nil
}

func (o objectType) String(val any) string {
	return ""
}

func (o objectType) IsEmpty(value any) bool {
	return false
}

func (o objectType) ValidateValue(value any) error {
	return nil
}

func (o objectType) Default() any {
	return nil
}
