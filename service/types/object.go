package types

import "encoding/json"

// string
type objectType struct {
}

func (o objectType) Equals(a, b interface{}) bool {
	v1, _ := o.Pack(a)
	v2, _ := o.Pack(b)

	return v1 == v2
}

func (o objectType) Pack(value interface{}) (interface{}, error) {
	data, err := json.Marshal(value)

	if err != nil {
		return nil, err
	}

	return string(data), nil
}

func (o objectType) UnPack(value interface{}) (interface{}, error) {
	if str, ok := value.(string); ok {
		var data = new(interface{})
		err := json.Unmarshal([]byte(str), data)

		return *data, err
	} else {
		return value, nil
	}
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
