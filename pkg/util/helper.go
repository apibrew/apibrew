package util

import (
	"google.golang.org/protobuf/types/known/structpb"
)

func StructKv(key string, value interface{}) *structpb.Value {
	return MapStructValue(map[string]interface{}{
		key: value,
	})
}

func MapStructValue(v map[string]interface{}) *structpb.Value {
	st, err := structpb.NewStruct(v)

	if err != nil {
		panic(err)
	}

	return structpb.NewStructValue(st)
}
