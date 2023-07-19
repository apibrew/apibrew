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

	var_c1cf1270d846 := record.Id

	if var_c1cf1270d846 != nil {
		var var_c1cf1270d846_mapped *structpb.Value

		var var_c1cf1270d846_err error
		var_c1cf1270d846_mapped, var_c1cf1270d846_err = types.ByResourcePropertyType(model.ResourceProperty_UUID).Pack(*var_c1cf1270d846)
		if var_c1cf1270d846_err != nil {
			panic(var_c1cf1270d846_err)
		}
		properties["id"] = var_c1cf1270d846_mapped
	}

	var_dcc23bfef4ea := record.Properties

	var var_dcc23bfef4ea_mapped *structpb.Value

	var var_dcc23bfef4ea_err error
	var_dcc23bfef4ea_mapped, var_dcc23bfef4ea_err = types.ByResourcePropertyType(model.ResourceProperty_OBJECT).Pack(var_dcc23bfef4ea)
	if var_dcc23bfef4ea_err != nil {
		panic(var_dcc23bfef4ea_err)
	}
	properties["properties"] = var_dcc23bfef4ea_mapped

	var_3ef7bfe094b1 := record.PackedProperties

	if var_3ef7bfe094b1 != nil {
		var var_3ef7bfe094b1_mapped *structpb.Value

		var var_3ef7bfe094b1_l []*structpb.Value
		for _, value := range var_3ef7bfe094b1 {

			var_dca0756dd783 := value
			var var_dca0756dd783_mapped *structpb.Value

			var var_dca0756dd783_err error
			var_dca0756dd783_mapped, var_dca0756dd783_err = types.ByResourcePropertyType(model.ResourceProperty_OBJECT).Pack(var_dca0756dd783)
			if var_dca0756dd783_err != nil {
				panic(var_dca0756dd783_err)
			}

			var_3ef7bfe094b1_l = append(var_3ef7bfe094b1_l, var_dca0756dd783_mapped)
		}
		var_3ef7bfe094b1_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_3ef7bfe094b1_l})
		properties["packedProperties"] = var_3ef7bfe094b1_mapped
	}
	return properties
}

func (m *RecordMapper) FromProperties(properties map[string]*structpb.Value) *Record {
	var s = m.New()
	if properties["id"] != nil {

		var_9551b13a3096 := properties["id"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_UUID).UnPack(var_9551b13a3096)

		if err != nil {
			panic(err)
		}

		var_9551b13a3096_mapped := new(uuid.UUID)
		*var_9551b13a3096_mapped = val.(uuid.UUID)

		s.Id = var_9551b13a3096_mapped
	}
	if properties["properties"] != nil {

		var_affc59bcfafb := properties["properties"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_OBJECT).UnPack(var_affc59bcfafb)

		if err != nil {
			panic(err)
		}

		var_affc59bcfafb_mapped := val.(unstructured.Unstructured)

		s.Properties = var_affc59bcfafb_mapped
	}
	if properties["packedProperties"] != nil {

		var_b252327682d8 := properties["packedProperties"]
		var_b252327682d8_mapped := []unstructured.Unstructured{}
		for _, v := range var_b252327682d8.GetListValue().Values {

			var_38df77065174 := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_OBJECT).UnPack(var_38df77065174)

			if err != nil {
				panic(err)
			}

			var_38df77065174_mapped := val.(unstructured.Unstructured)

			var_b252327682d8_mapped = append(var_b252327682d8_mapped, var_38df77065174_mapped)
		}

		s.PackedProperties = var_b252327682d8_mapped
	}
	return s
}
