// Code generated by apbr generate. DO NOT EDIT.
// versions:
// 	apbr generate v1.2

//go:build !codeanalysis

package resource_model

import (
	"github.com/apibrew/apibrew/pkg/abs"
	"github.com/apibrew/apibrew/pkg/formats/unstructured"
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

func (m *RecordMapper) ResourceIdentity() abs.ResourceIdentity {
	return abs.ResourceIdentity{
		Namespace: "system",
		Name:      "Record",
	}
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

	var_Id := record.Id

	if var_Id != nil {
		var var_Id_mapped *structpb.Value

		var var_Id_err error
		var_Id_mapped, var_Id_err = types.ByResourcePropertyType(model.ResourceProperty_UUID).Pack(*var_Id)
		if var_Id_err != nil {
			panic(var_Id_err)
		}
		properties["id"] = var_Id_mapped
	}

	var_Properties := record.Properties

	var var_Properties_mapped *structpb.Value

	var var_Properties_err error
	var_Properties_mapped, var_Properties_err = types.ByResourcePropertyType(model.ResourceProperty_OBJECT).Pack(var_Properties)
	if var_Properties_err != nil {
		panic(var_Properties_err)
	}
	properties["properties"] = var_Properties_mapped

	var_PackedProperties := record.PackedProperties

	if var_PackedProperties != nil {
		var var_PackedProperties_mapped *structpb.Value

		var var_PackedProperties_l []*structpb.Value
		for _, value := range var_PackedProperties {

			var_5x := value
			var var_5x_mapped *structpb.Value

			var var_5x_err error
			var_5x_mapped, var_5x_err = types.ByResourcePropertyType(model.ResourceProperty_OBJECT).Pack(var_5x)
			if var_5x_err != nil {
				panic(var_5x_err)
			}

			var_PackedProperties_l = append(var_PackedProperties_l, var_5x_mapped)
		}
		var_PackedProperties_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_PackedProperties_l})
		properties["packedProperties"] = var_PackedProperties_mapped
	}
	return properties
}

func (m *RecordMapper) FromProperties(properties map[string]*structpb.Value) *Record {
	var s = m.New()
	if properties["id"] != nil && properties["id"].AsInterface() != nil {

		var_Id := properties["id"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_UUID).UnPack(var_Id)

		if err != nil {
			panic(err)
		}

		var_Id_mapped := new(uuid.UUID)
		*var_Id_mapped = val.(uuid.UUID)

		s.Id = var_Id_mapped
	}
	if properties["properties"] != nil && properties["properties"].AsInterface() != nil {

		var_Properties := properties["properties"]
		var_Properties_mapped := unstructured.FromValue(var_Properties)

		s.Properties = var_Properties_mapped
	}
	if properties["packedProperties"] != nil && properties["packedProperties"].AsInterface() != nil {

		var_PackedProperties := properties["packedProperties"]
		var_PackedProperties_mapped := []interface{}{}
		for _, v := range var_PackedProperties.GetListValue().Values {

			var_4x := v
			var_4x_mapped := unstructured.FromValue(var_4x)

			var_PackedProperties_mapped = append(var_PackedProperties_mapped, var_4x_mapped)
		}

		s.PackedProperties = var_PackedProperties_mapped
	}
	return s
}

func (m *RecordMapper) ToUnstructured(record *Record) unstructured.Unstructured {
	var properties unstructured.Unstructured = make(unstructured.Unstructured)
	properties["type"] = "system/Record"

	var_Id := record.Id

	if var_Id != nil {
		var var_Id_mapped interface{}

		var_Id_mapped = var_Id.String()
		properties["id"] = var_Id_mapped
	}

	var_Properties := record.Properties

	var var_Properties_mapped interface{}

	var_Properties_mapped = var_Properties
	properties["properties"] = var_Properties_mapped

	var_PackedProperties := record.PackedProperties

	if var_PackedProperties != nil {
		var var_PackedProperties_mapped interface{}

		var var_PackedProperties_l []interface{}
		for _, value := range var_PackedProperties {

			var_5x := value
			var var_5x_mapped interface{}

			var_5x_mapped = var_5x

			var_PackedProperties_l = append(var_PackedProperties_l, var_5x_mapped)
		}
		var_PackedProperties_mapped = var_PackedProperties_l
		properties["packedProperties"] = var_PackedProperties_mapped
	}

	return properties
}
