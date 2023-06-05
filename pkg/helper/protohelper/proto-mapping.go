package protohelper

import (
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/resources/mapping"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/structpb"
)

type MappingHelper[T proto.Message] struct {
	Resource *model.Resource
	Instance func() T
}

func (h MappingHelper[T]) MapTo(t T) *model.Record {
	data, err := protojson.Marshal(t)

	if err != nil {
		panic(err)
	}

	var strucpb = new(structpb.Struct)
	err = protojson.Unmarshal(data, strucpb)

	if err != nil {
		panic(err)
	}

	return &model.Record{
		Properties: strucpb.GetFields(),
	}
}

func (h MappingHelper[T]) MapFrom(record *model.Record) T {
	var instance = h.Instance()
	mapping.MapSpecialColumnsFromRecord(instance, &record.Properties)

	delete(record.Properties, "createdBy")
	delete(record.Properties, "createdOn")
	delete(record.Properties, "updatedBy")
	delete(record.Properties, "updatedOn")

	data, err := protojson.Marshal(&structpb.Struct{Fields: record.Properties})

	if err != nil {
		panic(err)
	}

	err = protojson.Unmarshal(data, instance)

	if err != nil {
		panic(err)
	}

	return instance
}
