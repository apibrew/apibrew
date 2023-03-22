package util

import (
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"log"
)

func FromAny[T proto.Message](anyItem *anypb.Any, instance T) T {
	err := anyItem.UnmarshalTo(instance)

	if err != nil {
		log.Fatal(err)
	}

	return instance
}

func ToAny(instance proto.Message) *anypb.Any {
	anyItem := &anypb.Any{}

	err := anyItem.MarshalFrom(instance)

	if err != nil {
		log.Fatal(err)
	}

	return anyItem
}
