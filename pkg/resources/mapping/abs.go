package mapping

import (
	"github.com/apibrew/apibrew/pkg/abs"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

func MapToRecord[T proto.Message](list []T, mapper func(T) abs.RecordLike) []abs.RecordLike {
	var result []abs.RecordLike

	for _, item := range list {
		result = append(result, mapper(item))
	}

	return result
}

func MapFromRecord[T proto.Message](list []abs.RecordLike, mapper func(abs.RecordLike) T) []T {
	var result []T

	for _, item := range list {
		result = append(result, mapper(item))
	}

	return result
}

func MessageToRecord(message proto.Message) abs.RecordLike {
	return nil
}

func MessageFromRecord(resource, namespace string, record abs.RecordLike) proto.Message {
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
