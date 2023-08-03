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

	var_47212de5b394 := record.Id

	if var_47212de5b394 != nil {
		var var_47212de5b394_mapped *structpb.Value

		var var_47212de5b394_err error
		var_47212de5b394_mapped, var_47212de5b394_err = types.ByResourcePropertyType(model.ResourceProperty_UUID).Pack(*var_47212de5b394)
		if var_47212de5b394_err != nil {
			panic(var_47212de5b394_err)
		}
		properties["id"] = var_47212de5b394_mapped
	}

	var_af0002b37ffe := record.Properties

	var var_af0002b37ffe_mapped *structpb.Value

	var var_af0002b37ffe_err error
	var_af0002b37ffe_mapped, var_af0002b37ffe_err = types.ByResourcePropertyType(model.ResourceProperty_OBJECT).Pack(var_af0002b37ffe)
	if var_af0002b37ffe_err != nil {
		panic(var_af0002b37ffe_err)
	}
	properties["properties"] = var_af0002b37ffe_mapped

	var_29fcfec3ae81 := record.PackedProperties

	if var_29fcfec3ae81 != nil {
		var var_29fcfec3ae81_mapped *structpb.Value

		var var_29fcfec3ae81_l []*structpb.Value
		for _, value := range var_29fcfec3ae81 {

			var_bba77db8668b := value
			var var_bba77db8668b_mapped *structpb.Value

			var var_bba77db8668b_err error
			var_bba77db8668b_mapped, var_bba77db8668b_err = types.ByResourcePropertyType(model.ResourceProperty_OBJECT).Pack(var_bba77db8668b)
			if var_bba77db8668b_err != nil {
				panic(var_bba77db8668b_err)
			}

			var_29fcfec3ae81_l = append(var_29fcfec3ae81_l, var_bba77db8668b_mapped)
		}
		var_29fcfec3ae81_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_29fcfec3ae81_l})
		properties["packedProperties"] = var_29fcfec3ae81_mapped
	}
	return properties
}

func (m *RecordMapper) FromProperties(properties map[string]*structpb.Value) *Record {
	var s = m.New()
	if properties["id"] != nil {

		var_91bca3c05aaa := properties["id"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_UUID).UnPack(var_91bca3c05aaa)

		if err != nil {
			panic(err)
		}

		var_91bca3c05aaa_mapped := new(uuid.UUID)
		*var_91bca3c05aaa_mapped = val.(uuid.UUID)

		s.Id = var_91bca3c05aaa_mapped
	}
	if properties["properties"] != nil {

		var_a5e16e81ff90 := properties["properties"]
		var_a5e16e81ff90_mapped := unstructured.FromStructValue(var_a5e16e81ff90.GetStructValue())

		s.Properties = var_a5e16e81ff90_mapped
	}
	if properties["packedProperties"] != nil {

		var_08ed659020ac := properties["packedProperties"]
		var_08ed659020ac_mapped := []unstructured.Unstructured{}
		for _, v := range var_08ed659020ac.GetListValue().Values {

			var_ab6e72c96211 := v
			var_ab6e72c96211_mapped := unstructured.FromStructValue(var_ab6e72c96211.GetStructValue())

			var_08ed659020ac_mapped = append(var_08ed659020ac_mapped, var_ab6e72c96211_mapped)
		}

		s.PackedProperties = var_08ed659020ac_mapped
	}
	return s
}
