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

	if record.Id != nil {
		rec.Id = record.Id.String()
	}

	return rec
}

func (m *RecordMapper) FromRecord(record *model.Record) *Record {
	return m.FromProperties(record.Properties)
}

func (m *RecordMapper) ToProperties(record *Record) map[string]*structpb.Value {
	var properties = make(map[string]*structpb.Value)

	var_886c771b8d4e := record.Id

	if var_886c771b8d4e != nil {
		var var_886c771b8d4e_mapped *structpb.Value

		var var_886c771b8d4e_err error
		var_886c771b8d4e_mapped, var_886c771b8d4e_err = types.ByResourcePropertyType(model.ResourceProperty_UUID).Pack(*var_886c771b8d4e)
		if var_886c771b8d4e_err != nil {
			panic(var_886c771b8d4e_err)
		}
		properties["id"] = var_886c771b8d4e_mapped
	}

	var_713848463b8c := record.Properties

	var var_713848463b8c_mapped *structpb.Value

	var var_713848463b8c_err error
	var_713848463b8c_mapped, var_713848463b8c_err = types.ByResourcePropertyType(model.ResourceProperty_OBJECT).Pack(var_713848463b8c)
	if var_713848463b8c_err != nil {
		panic(var_713848463b8c_err)
	}
	properties["properties"] = var_713848463b8c_mapped

	var_18448e7f881f := record.PackedProperties

	if var_18448e7f881f != nil {
		var var_18448e7f881f_mapped *structpb.Value

		var var_18448e7f881f_l []*structpb.Value
		for _, value := range var_18448e7f881f {

			var_dbef9515bdc5 := value
			var var_dbef9515bdc5_mapped *structpb.Value

			var var_dbef9515bdc5_err error
			var_dbef9515bdc5_mapped, var_dbef9515bdc5_err = types.ByResourcePropertyType(model.ResourceProperty_OBJECT).Pack(var_dbef9515bdc5)
			if var_dbef9515bdc5_err != nil {
				panic(var_dbef9515bdc5_err)
			}

			var_18448e7f881f_l = append(var_18448e7f881f_l, var_dbef9515bdc5_mapped)
		}
		var_18448e7f881f_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_18448e7f881f_l})
		properties["packedProperties"] = var_18448e7f881f_mapped
	}
	return properties
}

func (m *RecordMapper) FromProperties(properties map[string]*structpb.Value) *Record {
	var s = m.New()
	if properties["id"] != nil {

		var_8ede5de22309 := properties["id"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_UUID).UnPack(var_8ede5de22309)

		if err != nil {
			panic(err)
		}

		var_8ede5de22309_mapped := new(uuid.UUID)
		*var_8ede5de22309_mapped = val.(uuid.UUID)

		s.Id = var_8ede5de22309_mapped
	}
	if properties["properties"] != nil {

		var_b3c599238915 := properties["properties"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_OBJECT).UnPack(var_b3c599238915)

		if err != nil {
			panic(err)
		}

		var_b3c599238915_mapped := val.(unstructured.Unstructured)

		s.Properties = var_b3c599238915_mapped
	}
	if properties["packedProperties"] != nil {

		var_e347aca59366 := properties["packedProperties"]
		var_e347aca59366_mapped := []unstructured.Unstructured{}
		for _, v := range var_e347aca59366.GetListValue().Values {

			var_b42af7aca404 := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_OBJECT).UnPack(var_b42af7aca404)

			if err != nil {
				panic(err)
			}

			var_b42af7aca404_mapped := val.(unstructured.Unstructured)

			var_e347aca59366_mapped = append(var_e347aca59366_mapped, var_b42af7aca404_mapped)
		}

		s.PackedProperties = var_e347aca59366_mapped
	}
	return s
}
