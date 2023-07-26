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

	var_dbc8fd43c8f4 := record.Id

	if var_dbc8fd43c8f4 != nil {
		var var_dbc8fd43c8f4_mapped *structpb.Value

		var var_dbc8fd43c8f4_err error
		var_dbc8fd43c8f4_mapped, var_dbc8fd43c8f4_err = types.ByResourcePropertyType(model.ResourceProperty_UUID).Pack(*var_dbc8fd43c8f4)
		if var_dbc8fd43c8f4_err != nil {
			panic(var_dbc8fd43c8f4_err)
		}
		properties["id"] = var_dbc8fd43c8f4_mapped
	}

	var_04d462f0516e := record.Properties

	var var_04d462f0516e_mapped *structpb.Value

	var var_04d462f0516e_err error
	var_04d462f0516e_mapped, var_04d462f0516e_err = types.ByResourcePropertyType(model.ResourceProperty_OBJECT).Pack(var_04d462f0516e)
	if var_04d462f0516e_err != nil {
		panic(var_04d462f0516e_err)
	}
	properties["properties"] = var_04d462f0516e_mapped

	var_4cc547e2483a := record.PackedProperties

	if var_4cc547e2483a != nil {
		var var_4cc547e2483a_mapped *structpb.Value

		var var_4cc547e2483a_l []*structpb.Value
		for _, value := range var_4cc547e2483a {

			var_e10a35f336b8 := value
			var var_e10a35f336b8_mapped *structpb.Value

			var var_e10a35f336b8_err error
			var_e10a35f336b8_mapped, var_e10a35f336b8_err = types.ByResourcePropertyType(model.ResourceProperty_OBJECT).Pack(var_e10a35f336b8)
			if var_e10a35f336b8_err != nil {
				panic(var_e10a35f336b8_err)
			}

			var_4cc547e2483a_l = append(var_4cc547e2483a_l, var_e10a35f336b8_mapped)
		}
		var_4cc547e2483a_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_4cc547e2483a_l})
		properties["packedProperties"] = var_4cc547e2483a_mapped
	}
	return properties
}

func (m *RecordMapper) FromProperties(properties map[string]*structpb.Value) *Record {
	var s = m.New()
	if properties["id"] != nil {

		var_085a2564fc3e := properties["id"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_UUID).UnPack(var_085a2564fc3e)

		if err != nil {
			panic(err)
		}

		var_085a2564fc3e_mapped := new(uuid.UUID)
		*var_085a2564fc3e_mapped = val.(uuid.UUID)

		s.Id = var_085a2564fc3e_mapped
	}
	if properties["properties"] != nil {

		var_7438caba67e5 := properties["properties"]
		var_7438caba67e5_mapped := unstructured.FromStructValue(var_7438caba67e5.GetStructValue())

		s.Properties = var_7438caba67e5_mapped
	}
	if properties["packedProperties"] != nil {

		var_06a4f94c3004 := properties["packedProperties"]
		var_06a4f94c3004_mapped := []unstructured.Unstructured{}
		for _, v := range var_06a4f94c3004.GetListValue().Values {

			var_aa760c4e39c5 := v
			var_aa760c4e39c5_mapped := unstructured.FromStructValue(var_aa760c4e39c5.GetStructValue())

			var_06a4f94c3004_mapped = append(var_06a4f94c3004_mapped, var_aa760c4e39c5_mapped)
		}

		s.PackedProperties = var_06a4f94c3004_mapped
	}
	return s
}
