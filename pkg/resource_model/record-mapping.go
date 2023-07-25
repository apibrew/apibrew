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

	var_b4d7ae7830ea := record.Id

	if var_b4d7ae7830ea != nil {
		var var_b4d7ae7830ea_mapped *structpb.Value

		var var_b4d7ae7830ea_err error
		var_b4d7ae7830ea_mapped, var_b4d7ae7830ea_err = types.ByResourcePropertyType(model.ResourceProperty_UUID).Pack(*var_b4d7ae7830ea)
		if var_b4d7ae7830ea_err != nil {
			panic(var_b4d7ae7830ea_err)
		}
		properties["id"] = var_b4d7ae7830ea_mapped
	}

	var_9ada106fcfc4 := record.Properties

	var var_9ada106fcfc4_mapped *structpb.Value

	var var_9ada106fcfc4_err error
	var_9ada106fcfc4_mapped, var_9ada106fcfc4_err = types.ByResourcePropertyType(model.ResourceProperty_OBJECT).Pack(var_9ada106fcfc4)
	if var_9ada106fcfc4_err != nil {
		panic(var_9ada106fcfc4_err)
	}
	properties["properties"] = var_9ada106fcfc4_mapped

	var_0cf71d799dbe := record.PackedProperties

	if var_0cf71d799dbe != nil {
		var var_0cf71d799dbe_mapped *structpb.Value

		var var_0cf71d799dbe_l []*structpb.Value
		for _, value := range var_0cf71d799dbe {

			var_1c22f7f33b6b := value
			var var_1c22f7f33b6b_mapped *structpb.Value

			var var_1c22f7f33b6b_err error
			var_1c22f7f33b6b_mapped, var_1c22f7f33b6b_err = types.ByResourcePropertyType(model.ResourceProperty_OBJECT).Pack(var_1c22f7f33b6b)
			if var_1c22f7f33b6b_err != nil {
				panic(var_1c22f7f33b6b_err)
			}

			var_0cf71d799dbe_l = append(var_0cf71d799dbe_l, var_1c22f7f33b6b_mapped)
		}
		var_0cf71d799dbe_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_0cf71d799dbe_l})
		properties["packedProperties"] = var_0cf71d799dbe_mapped
	}
	return properties
}

func (m *RecordMapper) FromProperties(properties map[string]*structpb.Value) *Record {
	var s = m.New()
	if properties["id"] != nil {

		var_1dd7f20ea547 := properties["id"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_UUID).UnPack(var_1dd7f20ea547)

		if err != nil {
			panic(err)
		}

		var_1dd7f20ea547_mapped := new(uuid.UUID)
		*var_1dd7f20ea547_mapped = val.(uuid.UUID)

		s.Id = var_1dd7f20ea547_mapped
	}
	if properties["properties"] != nil {

		var_f202b701ea5e := properties["properties"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_OBJECT).UnPack(var_f202b701ea5e)

		if err != nil {
			panic(err)
		}

		var_f202b701ea5e_mapped := val.(unstructured.Unstructured)

		s.Properties = var_f202b701ea5e_mapped
	}
	if properties["packedProperties"] != nil {

		var_7af313e0ea04 := properties["packedProperties"]
		var_7af313e0ea04_mapped := []unstructured.Unstructured{}
		for _, v := range var_7af313e0ea04.GetListValue().Values {

			var_504137fbec4d := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_OBJECT).UnPack(var_504137fbec4d)

			if err != nil {
				panic(err)
			}

			var_504137fbec4d_mapped := val.(unstructured.Unstructured)

			var_7af313e0ea04_mapped = append(var_7af313e0ea04_mapped, var_504137fbec4d_mapped)
		}

		s.PackedProperties = var_7af313e0ea04_mapped
	}
	return s
}
