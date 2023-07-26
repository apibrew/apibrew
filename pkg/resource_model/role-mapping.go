package resource_model

import (
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/types"
	"google.golang.org/protobuf/types/known/structpb"
)

import "github.com/google/uuid"
import "time"
import "github.com/apibrew/apibrew/pkg/formats/unstructured"

type RoleMapper struct {
}

func NewRoleMapper() *RoleMapper {
	return &RoleMapper{}
}

var RoleMapperInstance = NewRoleMapper()

func (m *RoleMapper) New() *Role {
	return &Role{}
}

func (m *RoleMapper) ToRecord(role *Role) *model.Record {
	var rec = &model.Record{}
	rec.Properties = m.ToProperties(role)

	if role.Id != nil {
		rec.Id = role.Id.String()
	}

	return rec
}

func (m *RoleMapper) FromRecord(record *model.Record) *Role {
	return m.FromProperties(record.Properties)
}

func (m *RoleMapper) ToProperties(role *Role) map[string]*structpb.Value {
	var properties = make(map[string]*structpb.Value)

	var_db15246d143d := role.Id

	if var_db15246d143d != nil {
		var var_db15246d143d_mapped *structpb.Value

		var var_db15246d143d_err error
		var_db15246d143d_mapped, var_db15246d143d_err = types.ByResourcePropertyType(model.ResourceProperty_UUID).Pack(*var_db15246d143d)
		if var_db15246d143d_err != nil {
			panic(var_db15246d143d_err)
		}
		properties["id"] = var_db15246d143d_mapped
	}

	var_7feb8d084ced := role.Version

	var var_7feb8d084ced_mapped *structpb.Value

	var var_7feb8d084ced_err error
	var_7feb8d084ced_mapped, var_7feb8d084ced_err = types.ByResourcePropertyType(model.ResourceProperty_INT32).Pack(var_7feb8d084ced)
	if var_7feb8d084ced_err != nil {
		panic(var_7feb8d084ced_err)
	}
	properties["version"] = var_7feb8d084ced_mapped

	var_32761f14b33a := role.CreatedBy

	if var_32761f14b33a != nil {
		var var_32761f14b33a_mapped *structpb.Value

		var var_32761f14b33a_err error
		var_32761f14b33a_mapped, var_32761f14b33a_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_32761f14b33a)
		if var_32761f14b33a_err != nil {
			panic(var_32761f14b33a_err)
		}
		properties["createdBy"] = var_32761f14b33a_mapped
	}

	var_b1eff4ba6124 := role.UpdatedBy

	if var_b1eff4ba6124 != nil {
		var var_b1eff4ba6124_mapped *structpb.Value

		var var_b1eff4ba6124_err error
		var_b1eff4ba6124_mapped, var_b1eff4ba6124_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_b1eff4ba6124)
		if var_b1eff4ba6124_err != nil {
			panic(var_b1eff4ba6124_err)
		}
		properties["updatedBy"] = var_b1eff4ba6124_mapped
	}

	var_7dd467780030 := role.CreatedOn

	if var_7dd467780030 != nil {
		var var_7dd467780030_mapped *structpb.Value

		var var_7dd467780030_err error
		var_7dd467780030_mapped, var_7dd467780030_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_7dd467780030)
		if var_7dd467780030_err != nil {
			panic(var_7dd467780030_err)
		}
		properties["createdOn"] = var_7dd467780030_mapped
	}

	var_96e1a805cbce := role.UpdatedOn

	if var_96e1a805cbce != nil {
		var var_96e1a805cbce_mapped *structpb.Value

		var var_96e1a805cbce_err error
		var_96e1a805cbce_mapped, var_96e1a805cbce_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_96e1a805cbce)
		if var_96e1a805cbce_err != nil {
			panic(var_96e1a805cbce_err)
		}
		properties["updatedOn"] = var_96e1a805cbce_mapped
	}

	var_2dd5050aba6d := role.Name

	var var_2dd5050aba6d_mapped *structpb.Value

	var var_2dd5050aba6d_err error
	var_2dd5050aba6d_mapped, var_2dd5050aba6d_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_2dd5050aba6d)
	if var_2dd5050aba6d_err != nil {
		panic(var_2dd5050aba6d_err)
	}
	properties["name"] = var_2dd5050aba6d_mapped

	var_a62aff65d200 := role.SecurityConstraints

	if var_a62aff65d200 != nil {
		var var_a62aff65d200_mapped *structpb.Value

		var var_a62aff65d200_l []*structpb.Value
		for _, value := range var_a62aff65d200 {

			var_fc78303a7b57 := value
			var var_fc78303a7b57_mapped *structpb.Value

			var_fc78303a7b57_mapped = structpb.NewStructValue(&structpb.Struct{Fields: SecurityConstraintMapperInstance.ToProperties(var_fc78303a7b57)})

			var_a62aff65d200_l = append(var_a62aff65d200_l, var_fc78303a7b57_mapped)
		}
		var_a62aff65d200_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_a62aff65d200_l})
		properties["securityConstraints"] = var_a62aff65d200_mapped
	}

	var_79b5558433ac := role.Details

	if var_79b5558433ac != nil {
		var var_79b5558433ac_mapped *structpb.Value

		var var_79b5558433ac_err error
		var_79b5558433ac_mapped, var_79b5558433ac_err = types.ByResourcePropertyType(model.ResourceProperty_OBJECT).Pack(*var_79b5558433ac)
		if var_79b5558433ac_err != nil {
			panic(var_79b5558433ac_err)
		}
		properties["details"] = var_79b5558433ac_mapped
	}
	return properties
}

func (m *RoleMapper) FromProperties(properties map[string]*structpb.Value) *Role {
	var s = m.New()
	if properties["id"] != nil {

		var_b7106770e537 := properties["id"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_UUID).UnPack(var_b7106770e537)

		if err != nil {
			panic(err)
		}

		var_b7106770e537_mapped := new(uuid.UUID)
		*var_b7106770e537_mapped = val.(uuid.UUID)

		s.Id = var_b7106770e537_mapped
	}
	if properties["version"] != nil {

		var_58a1b34d0201 := properties["version"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(var_58a1b34d0201)

		if err != nil {
			panic(err)
		}

		var_58a1b34d0201_mapped := val.(int32)

		s.Version = var_58a1b34d0201_mapped
	}
	if properties["createdBy"] != nil {

		var_426a1495d306 := properties["createdBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_426a1495d306)

		if err != nil {
			panic(err)
		}

		var_426a1495d306_mapped := new(string)
		*var_426a1495d306_mapped = val.(string)

		s.CreatedBy = var_426a1495d306_mapped
	}
	if properties["updatedBy"] != nil {

		var_4c45d4d0a50f := properties["updatedBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_4c45d4d0a50f)

		if err != nil {
			panic(err)
		}

		var_4c45d4d0a50f_mapped := new(string)
		*var_4c45d4d0a50f_mapped = val.(string)

		s.UpdatedBy = var_4c45d4d0a50f_mapped
	}
	if properties["createdOn"] != nil {

		var_573d6d3b3254 := properties["createdOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_573d6d3b3254)

		if err != nil {
			panic(err)
		}

		var_573d6d3b3254_mapped := new(time.Time)
		*var_573d6d3b3254_mapped = val.(time.Time)

		s.CreatedOn = var_573d6d3b3254_mapped
	}
	if properties["updatedOn"] != nil {

		var_3c18532a73ac := properties["updatedOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_3c18532a73ac)

		if err != nil {
			panic(err)
		}

		var_3c18532a73ac_mapped := new(time.Time)
		*var_3c18532a73ac_mapped = val.(time.Time)

		s.UpdatedOn = var_3c18532a73ac_mapped
	}
	if properties["name"] != nil {

		var_52f166d617fa := properties["name"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_52f166d617fa)

		if err != nil {
			panic(err)
		}

		var_52f166d617fa_mapped := val.(string)

		s.Name = var_52f166d617fa_mapped
	}
	if properties["securityConstraints"] != nil {

		var_7baae6e20233 := properties["securityConstraints"]
		var_7baae6e20233_mapped := []*SecurityConstraint{}
		for _, v := range var_7baae6e20233.GetListValue().Values {

			var_8ba4c959b717 := v
			var_8ba4c959b717_mapped := SecurityConstraintMapperInstance.FromProperties(var_8ba4c959b717.GetStructValue().Fields)

			var_7baae6e20233_mapped = append(var_7baae6e20233_mapped, var_8ba4c959b717_mapped)
		}

		s.SecurityConstraints = var_7baae6e20233_mapped
	}
	if properties["details"] != nil {

		var_9a7b7bbfa76e := properties["details"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_OBJECT).UnPack(var_9a7b7bbfa76e)

		if err != nil {
			panic(err)
		}

		var_9a7b7bbfa76e_mapped := new(unstructured.Unstructured)
		*var_9a7b7bbfa76e_mapped = val.(unstructured.Unstructured)

		s.Details = var_9a7b7bbfa76e_mapped
	}
	return s
}
