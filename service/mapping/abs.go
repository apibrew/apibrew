package mapping

import (
	"github.com/tislib/data-handler/model"
	"google.golang.org/protobuf/proto"
)

func MapToRecord[T proto.Message](list []T, mapper func(T) *model.Record) []*model.Record {
	var result []*model.Record

	for _, item := range list {
		result = append(result, mapper(item))
	}

	return result
}

func MapFromRecord[T proto.Message](list []*model.Record, mapper func(*model.Record) T) []T {
	var result []T

	for _, item := range list {
		result = append(result, mapper(item))
	}

	return result
}
