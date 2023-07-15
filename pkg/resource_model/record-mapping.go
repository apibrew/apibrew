package resource_model

import (
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/types"
	"google.golang.org/protobuf/types/known/structpb"
)

import "github.com/google/uuid"
import "github.com/apibrew/apibrew/pkg/formats/unstructured"
import "encoding/json"

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

		var_65211d43dfa0 := properties["id"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_UUID).UnPack(var_65211d43dfa0)

		if err != nil {
			panic(err)
		}

		var_65211d43dfa0_mapped := new(uuid.UUID)
		*var_65211d43dfa0_mapped = val.(uuid.UUID)

		s.Id = var_65211d43dfa0_mapped
	}
	if properties["properties"] != nil {

		var_41a3eca5294a := properties["properties"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_OBJECT).UnPack(var_41a3eca5294a)

		if err != nil {
			panic(err)
		}

		var_41a3eca5294a_mapped := val.(unstructured.Unstructured)

		s.Properties = var_41a3eca5294a_mapped
	}
	if properties["packedProperties"] != nil {

		var_8caf7c79756a := properties["packedProperties"]
		var_8caf7c79756a_mapped := []unstructured.Unstructured{}
		for _, v := range var_8caf7c79756a.GetListValue().Values {

			var_5fac1524b4db := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_OBJECT).UnPack(var_5fac1524b4db)

			if err != nil {
				panic(err)
			}

			var_5fac1524b4db_mapped := val.(unstructured.Unstructured)

			var_8caf7c79756a_mapped = append(var_8caf7c79756a_mapped, var_5fac1524b4db_mapped)
		}

		s.PackedProperties = var_8caf7c79756a_mapped
	}
	return s
}
