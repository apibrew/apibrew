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

		var_ff2d5684e9a3 := properties["id"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_UUID).UnPack(var_ff2d5684e9a3)

		if err != nil {
			panic(err)
		}

		var_ff2d5684e9a3_mapped := new(uuid.UUID)
		*var_ff2d5684e9a3_mapped = val.(uuid.UUID)

		s.Id = var_ff2d5684e9a3_mapped
	}
	if properties["properties"] != nil {

		var_615b7c705441 := properties["properties"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_OBJECT).UnPack(var_615b7c705441)

		if err != nil {
			panic(err)
		}

		var_615b7c705441_mapped := val.(unstructured.Unstructured)

		s.Properties = var_615b7c705441_mapped
	}
	if properties["packedProperties"] != nil {

		var_a23147582751 := properties["packedProperties"]
		var_a23147582751_mapped := []unstructured.Unstructured{}
		for _, v := range var_a23147582751.GetListValue().Values {

			var_34117d4743b5 := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_OBJECT).UnPack(var_34117d4743b5)

			if err != nil {
				panic(err)
			}

			var_34117d4743b5_mapped := val.(unstructured.Unstructured)

			var_a23147582751_mapped = append(var_a23147582751_mapped, var_34117d4743b5_mapped)
		}

		s.PackedProperties = var_a23147582751_mapped
	}
	return s
}
