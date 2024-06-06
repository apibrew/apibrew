package util

import (
	"google.golang.org/protobuf/types/known/structpb"
)

func StructKv(key string, value interface{}) *structpb.Value {
	return MapStructValue(map[string]interface{}{
		key: value,
	})
}

func Kv(key string, value interface{}) map[string]interface{} {
	return map[string]interface{}{
		key: value,
	}
}

func StructKv2(key1 string, value1 interface{}, key2 string, value2 interface{}) *structpb.Value {
	return MapStructValue(map[string]interface{}{
		key1: value1,
		key2: value2,
	})
}

func MapStructValue(v map[string]interface{}) *structpb.Value {
	st, err := structpb.NewStruct(v)

	if err != nil {
		panic(err)
	}

	return structpb.NewStructValue(st)
}
