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

	var_e24333f11a44 := record.Id

	if var_e24333f11a44 != nil {
		var var_e24333f11a44_mapped *structpb.Value

		var var_e24333f11a44_err error
		var_e24333f11a44_mapped, var_e24333f11a44_err = types.ByResourcePropertyType(model.ResourceProperty_UUID).Pack(*var_e24333f11a44)
		if var_e24333f11a44_err != nil {
			panic(var_e24333f11a44_err)
		}
		properties["id"] = var_e24333f11a44_mapped
	}

	var_534c572a5b68 := record.Properties

	var var_534c572a5b68_mapped *structpb.Value

	var var_534c572a5b68_err error
	var_534c572a5b68_mapped, var_534c572a5b68_err = types.ByResourcePropertyType(model.ResourceProperty_OBJECT).Pack(var_534c572a5b68)
	if var_534c572a5b68_err != nil {
		panic(var_534c572a5b68_err)
	}
	properties["properties"] = var_534c572a5b68_mapped

	var_f53d0e4c0bce := record.PackedProperties

	if var_f53d0e4c0bce != nil {
		var var_f53d0e4c0bce_mapped *structpb.Value

		var var_f53d0e4c0bce_l []*structpb.Value
		for _, value := range var_f53d0e4c0bce {

			var_d73401b16076 := value
			var var_d73401b16076_mapped *structpb.Value

			var var_d73401b16076_err error
			var_d73401b16076_mapped, var_d73401b16076_err = types.ByResourcePropertyType(model.ResourceProperty_OBJECT).Pack(var_d73401b16076)
			if var_d73401b16076_err != nil {
				panic(var_d73401b16076_err)
			}

			var_f53d0e4c0bce_l = append(var_f53d0e4c0bce_l, var_d73401b16076_mapped)
		}
		var_f53d0e4c0bce_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_f53d0e4c0bce_l})
		properties["packedProperties"] = var_f53d0e4c0bce_mapped
	}
	return properties
}

func (m *RecordMapper) FromProperties(properties map[string]*structpb.Value) *Record {
	var s = m.New()
	if properties["id"] != nil {

		var_452ae5cfcd5f := properties["id"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_UUID).UnPack(var_452ae5cfcd5f)

		if err != nil {
			panic(err)
		}

		var_452ae5cfcd5f_mapped := new(uuid.UUID)
		*var_452ae5cfcd5f_mapped = val.(uuid.UUID)

		s.Id = var_452ae5cfcd5f_mapped
	}
	if properties["properties"] != nil {

		var_abc50302a62c := properties["properties"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_OBJECT).UnPack(var_abc50302a62c)

		if err != nil {
			panic(err)
		}

		var_abc50302a62c_mapped := val.(unstructured.Unstructured)

		s.Properties = var_abc50302a62c_mapped
	}
	if properties["packedProperties"] != nil {

		var_14c5aa029e6e := properties["packedProperties"]
		var_14c5aa029e6e_mapped := []unstructured.Unstructured{}
		for _, v := range var_14c5aa029e6e.GetListValue().Values {

			var_619822761d84 := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_OBJECT).UnPack(var_619822761d84)

			if err != nil {
				panic(err)
			}

			var_619822761d84_mapped := val.(unstructured.Unstructured)

			var_14c5aa029e6e_mapped = append(var_14c5aa029e6e_mapped, var_619822761d84_mapped)
		}

		s.PackedProperties = var_14c5aa029e6e_mapped
	}
	return s
}
