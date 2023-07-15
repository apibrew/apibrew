package resource_model

import (
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/types"
	"google.golang.org/protobuf/types/known/structpb"
)

import "github.com/google/uuid"
import "github.com/apibrew/apibrew/pkg/formats/unstructured"

type RecordMapper struct {
}

func NewRecordMapper() *RecordMapper {
	return &RecordMapper{}
}

var RecordMapperInstance = NewRecordMapper()

func (m *RecordMapper) New() *Record {
	return &Record{}
}

func (m *RecordMapper) ToRecord(record *Record) *model.Record {
	var rec = &model.Record{}
	rec.Properties = m.ToProperties(record)
	return rec
}

func (m *RecordMapper) FromRecord(record *model.Record) *Record {
	return m.FromProperties(record.Properties)
}

func (m *RecordMapper) ToProperties(record *Record) map[string]*structpb.Value {
	var properties = make(map[string]*structpb.Value)

	if record.Id != nil {
		id, err := types.ByResourcePropertyType(model.ResourceProperty_UUID).Pack(*record.Id)
		if err != nil {
			panic(err)
		}
		properties["id"] = id
	}

	if record.PackedProperties != nil {
	}

	return properties
}

func (m *RecordMapper) FromProperties(properties map[string]*structpb.Value) *Record {
	var s = m.New()
	if properties["id"] != nil {

		var_2892878dee7c := properties["id"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_UUID).UnPack(var_2892878dee7c)

		if err != nil {
			panic(err)
		}

		var_2892878dee7c_mapped := new(uuid.UUID)
		*var_2892878dee7c_mapped = val.(uuid.UUID)

		s.Id = var_2892878dee7c_mapped
	}
	if properties["properties"] != nil {

		var_482ac1801cfa := properties["properties"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_OBJECT).UnPack(var_482ac1801cfa)

		if err != nil {
			panic(err)
		}

		var_482ac1801cfa_mapped := val.(unstructured.Unstructured)

		s.Properties = var_482ac1801cfa_mapped
	}
	if properties["packedProperties"] != nil {

		var_5a0a96d0e398 := properties["packedProperties"]
		var_5a0a96d0e398_mapped := []unstructured.Unstructured{}
		for _, v := range var_5a0a96d0e398.GetListValue().Values {

			var_fab3836afff0 := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_OBJECT).UnPack(var_fab3836afff0)

			if err != nil {
				panic(err)
			}

			var_fab3836afff0_mapped := val.(unstructured.Unstructured)

			var_5a0a96d0e398_mapped = append(var_5a0a96d0e398_mapped, var_fab3836afff0_mapped)
		}

		s.PackedProperties = var_5a0a96d0e398_mapped
	}
	return s
}
