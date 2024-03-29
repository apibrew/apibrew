package mapping

import (
	"github.com/apibrew/apibrew/pkg/model"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
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

func MessageToRecord(message proto.Message) *model.Record {
	return nil
}

func MessageFromRecord(resource, namespace string, record *model.Record) proto.Message {
	return nil
}

//goland:noinspection GoUnusedExportedFunction
func MessageToAny[T proto.Message](message T) (*anypb.Any, error) {
	any1 := new(anypb.Any)

	err := anypb.MarshalFrom(any1, message, proto.MarshalOptions{})

	if err != nil {
		return nil, err
	}

	return any1, nil
}

//goland:noinspection GoUnusedExportedFunction
func MessageFromAny[T proto.Message](any1 *anypb.Any, instance T) error {
	return anypb.UnmarshalTo(any1, instance, proto.UnmarshalOptions{})
}
