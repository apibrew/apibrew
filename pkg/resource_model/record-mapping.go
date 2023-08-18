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

	var_04522749d081 := record.Id

	if var_04522749d081 != nil {
		var var_04522749d081_mapped *structpb.Value

		var var_04522749d081_err error
		var_04522749d081_mapped, var_04522749d081_err = types.ByResourcePropertyType(model.ResourceProperty_UUID).Pack(*var_04522749d081)
		if var_04522749d081_err != nil {
			panic(var_04522749d081_err)
		}
		properties["id"] = var_04522749d081_mapped
	}

	var_d980f764355a := record.Properties

	var var_d980f764355a_mapped *structpb.Value

	var var_d980f764355a_err error
	var_d980f764355a_mapped, var_d980f764355a_err = types.ByResourcePropertyType(model.ResourceProperty_OBJECT).Pack(var_d980f764355a)
	if var_d980f764355a_err != nil {
		panic(var_d980f764355a_err)
	}
	properties["properties"] = var_d980f764355a_mapped

	var_59657bdd67ca := record.PackedProperties

	if var_59657bdd67ca != nil {
		var var_59657bdd67ca_mapped *structpb.Value

		var var_59657bdd67ca_l []*structpb.Value
		for _, value := range var_59657bdd67ca {

			var_6b11f95857cf := value
			var var_6b11f95857cf_mapped *structpb.Value

			var var_6b11f95857cf_err error
			var_6b11f95857cf_mapped, var_6b11f95857cf_err = types.ByResourcePropertyType(model.ResourceProperty_OBJECT).Pack(var_6b11f95857cf)
			if var_6b11f95857cf_err != nil {
				panic(var_6b11f95857cf_err)
			}

			var_59657bdd67ca_l = append(var_59657bdd67ca_l, var_6b11f95857cf_mapped)
		}
		var_59657bdd67ca_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_59657bdd67ca_l})
		properties["packedProperties"] = var_59657bdd67ca_mapped
	}
	return properties
}

func (m *RecordMapper) FromProperties(properties map[string]*structpb.Value) *Record {
	var s = m.New()
	if properties["id"] != nil {

		var_62b9751f299b := properties["id"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_UUID).UnPack(var_62b9751f299b)

		if err != nil {
			panic(err)
		}

		var_62b9751f299b_mapped := new(uuid.UUID)
		*var_62b9751f299b_mapped = val.(uuid.UUID)

		s.Id = var_62b9751f299b_mapped
	}
	if properties["properties"] != nil {

		var_89df1c0638a2 := properties["properties"]
		var_89df1c0638a2_mapped := unstructured.FromStructValue(var_89df1c0638a2.GetStructValue())

		s.Properties = var_89df1c0638a2_mapped
	}
	if properties["packedProperties"] != nil {

		var_ac32c5adca93 := properties["packedProperties"]
		var_ac32c5adca93_mapped := []unstructured.Unstructured{}
		for _, v := range var_ac32c5adca93.GetListValue().Values {

			var_494c89ffd0d4 := v
			var_494c89ffd0d4_mapped := unstructured.FromStructValue(var_494c89ffd0d4.GetStructValue())

			var_ac32c5adca93_mapped = append(var_ac32c5adca93_mapped, var_494c89ffd0d4_mapped)
		}

		s.PackedProperties = var_ac32c5adca93_mapped
	}
	return s
}
