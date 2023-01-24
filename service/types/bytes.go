package types

import (
	"encoding/base64"
	"errors"
	"reflect"
)

// string: base64
type bytesType struct {
}

func (b bytesType) Equals(aBytes, bBytes interface{}) bool {
	aArr := aBytes.([]byte)
	bArr := bBytes.([]byte)

	if len(aArr) != len(bArr) {
		return false
	}

	isEqual := true

	for i := 0; i < len(aArr); i++ {
		isEqual = isEqual && (aArr[i] == bArr[i])
	}

	return isEqual
}

func (b bytesType) Pack(value interface{}) (interface{}, error) {
	return base64.StdEncoding.EncodeToString(value.([]byte)), nil
}

func (b bytesType) UnPack(value interface{}) (interface{}, error) {
	if str, ok := value.(string); ok {
		return base64.StdEncoding.DecodeString(str)
	} else {
		return nil, errors.New("wrong type: " + reflect.TypeOf(value).String())
	}
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
	return val == nil || len(val.([]byte)) == 0
}

func (b bytesType) ValidatePackedValue(value any) error {
	err := canCast[string]("string", value)

	if err != nil {
		return err
	}

	_, err = b.UnPack(value)

	return err
}

func (b bytesType) Default() any {
	return []byte("")
}
