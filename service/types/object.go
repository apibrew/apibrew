package types

import "encoding/json"

type objectType struct {
}

func (o objectType) Pack(value interface{}) (interface{}, error) {
	var data = new(interface{})
	err := json.Unmarshal([]byte(value.(string)), data)

	return *data, err
}

func (o objectType) UnPack(value interface{}) (interface{}, error) {
	data, err := json.Marshal(value)

	if err != nil {
		return nil, err
	}

	return string(data), nil
}

func (o objectType) Pointer(required bool) any {
	if required {
		return new(string)
	} else {
		return new(*string)
	}
}

func (o objectType) String(val any) string {
	return ""
}

func (o objectType) IsEmpty(value any) bool {
	return value == nil
}

func (o objectType) ValidateValue(value any) error {
	return nil
}

func (o objectType) Default() any {
	return nil
}
