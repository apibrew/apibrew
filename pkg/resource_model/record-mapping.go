package resource_model

import (
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/types"
	"google.golang.org/protobuf/types/known/structpb"
)

import "github.com/google/uuid"

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

	if record.Properties != nil {
	}

	if record.PackedProperties != nil {
	}

	return properties
}

func (m *RecordMapper) FromProperties(properties map[string]*structpb.Value) *Record {
	var s = m.New()
	if properties["id"] != nil {

		var_15ae3e6c1b13 := properties["id"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_UUID).UnPack(var_15ae3e6c1b13)

		if err != nil {
			panic(err)
		}

		var_15ae3e6c1b13_mapped := new(uuid.UUID)
		*var_15ae3e6c1b13_mapped = val.(uuid.UUID)

		s.Id = var_15ae3e6c1b13_mapped
	}
	if properties["properties"] != nil {

		var_c5ebbffd6a69 := properties["properties"]
		var_c5ebbffd6a69_mapped := make(map[string]interface{})
		for k, v := range var_c5ebbffd6a69.GetStructValue().Fields {

			var_e95b98f6b5e5 := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_OBJECT).UnPack(var_e95b98f6b5e5)

			if err != nil {
				panic(err)
			}

			var_e95b98f6b5e5_mapped := val.(interface{})

			var_c5ebbffd6a69_mapped[k] = var_e95b98f6b5e5_mapped
		}

		s.Properties = var_c5ebbffd6a69_mapped
	}
	if properties["packedProperties"] != nil {

		var_fb2b6af1e372 := properties["packedProperties"]
		var_fb2b6af1e372_mapped := []interface{}{}
		for _, v := range var_fb2b6af1e372.GetListValue().Values {

			var_67ebdf268f2d := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_OBJECT).UnPack(var_67ebdf268f2d)

			if err != nil {
				panic(err)
			}

			var_67ebdf268f2d_mapped := val.(interface{})

			var_fb2b6af1e372_mapped = append(var_fb2b6af1e372_mapped, var_67ebdf268f2d_mapped)
		}

		s.PackedProperties = var_fb2b6af1e372_mapped
	}
	return s
}
