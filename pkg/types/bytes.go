package types

import (
	"encoding/base64"
	"google.golang.org/protobuf/types/known/structpb"
)

var BytesType = bytesType{}

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

func (b bytesType) Pack(value interface{}) (*structpb.Value, error) {
	return structpb.NewValue(value)
}

func (b bytesType) UnPack(value *structpb.Value) (interface{}, error) {
	return base64.StdEncoding.DecodeString(value.GetStringValue())
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

func (b bytesType) Default() any {
	return []byte("")
}
