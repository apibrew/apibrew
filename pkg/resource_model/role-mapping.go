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

	var_bc697359dc96 := role.Id

	if var_bc697359dc96 != nil {
		var var_bc697359dc96_mapped *structpb.Value

		var var_bc697359dc96_err error
		var_bc697359dc96_mapped, var_bc697359dc96_err = types.ByResourcePropertyType(model.ResourceProperty_UUID).Pack(*var_bc697359dc96)
		if var_bc697359dc96_err != nil {
			panic(var_bc697359dc96_err)
		}
		properties["id"] = var_bc697359dc96_mapped
	}

	var_7578347bdf41 := role.Version

	var var_7578347bdf41_mapped *structpb.Value

	var var_7578347bdf41_err error
	var_7578347bdf41_mapped, var_7578347bdf41_err = types.ByResourcePropertyType(model.ResourceProperty_INT32).Pack(var_7578347bdf41)
	if var_7578347bdf41_err != nil {
		panic(var_7578347bdf41_err)
	}
	properties["version"] = var_7578347bdf41_mapped

	var_6ae8163817c4 := role.CreatedBy

	if var_6ae8163817c4 != nil {
		var var_6ae8163817c4_mapped *structpb.Value

		var var_6ae8163817c4_err error
		var_6ae8163817c4_mapped, var_6ae8163817c4_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_6ae8163817c4)
		if var_6ae8163817c4_err != nil {
			panic(var_6ae8163817c4_err)
		}
		properties["createdBy"] = var_6ae8163817c4_mapped
	}

	var_184d3b024010 := role.UpdatedBy

	if var_184d3b024010 != nil {
		var var_184d3b024010_mapped *structpb.Value

		var var_184d3b024010_err error
		var_184d3b024010_mapped, var_184d3b024010_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_184d3b024010)
		if var_184d3b024010_err != nil {
			panic(var_184d3b024010_err)
		}
		properties["updatedBy"] = var_184d3b024010_mapped
	}

	var_0d076dfd409b := role.CreatedOn

	if var_0d076dfd409b != nil {
		var var_0d076dfd409b_mapped *structpb.Value

		var var_0d076dfd409b_err error
		var_0d076dfd409b_mapped, var_0d076dfd409b_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_0d076dfd409b)
		if var_0d076dfd409b_err != nil {
			panic(var_0d076dfd409b_err)
		}
		properties["createdOn"] = var_0d076dfd409b_mapped
	}

	var_a9d55727c841 := role.UpdatedOn

	if var_a9d55727c841 != nil {
		var var_a9d55727c841_mapped *structpb.Value

		var var_a9d55727c841_err error
		var_a9d55727c841_mapped, var_a9d55727c841_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_a9d55727c841)
		if var_a9d55727c841_err != nil {
			panic(var_a9d55727c841_err)
		}
		properties["updatedOn"] = var_a9d55727c841_mapped
	}

	var_92663a611971 := role.Name

	var var_92663a611971_mapped *structpb.Value

	var var_92663a611971_err error
	var_92663a611971_mapped, var_92663a611971_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_92663a611971)
	if var_92663a611971_err != nil {
		panic(var_92663a611971_err)
	}
	properties["name"] = var_92663a611971_mapped

	var_95aacf58dcd3 := role.SecurityConstraints

	if var_95aacf58dcd3 != nil {
		var var_95aacf58dcd3_mapped *structpb.Value

		var var_95aacf58dcd3_l []*structpb.Value
		for _, value := range var_95aacf58dcd3 {

			var_a4e4ea8376a8 := value
			var var_a4e4ea8376a8_mapped *structpb.Value

			var_a4e4ea8376a8_mapped = structpb.NewStructValue(&structpb.Struct{Fields: SecurityConstraintMapperInstance.ToProperties(var_a4e4ea8376a8)})

			var_95aacf58dcd3_l = append(var_95aacf58dcd3_l, var_a4e4ea8376a8_mapped)
		}
		var_95aacf58dcd3_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_95aacf58dcd3_l})
		properties["securityConstraints"] = var_95aacf58dcd3_mapped
	}

	var_f13893fb654c := role.Details

	if var_f13893fb654c != nil {
		var var_f13893fb654c_mapped *structpb.Value

		var var_f13893fb654c_err error
		var_f13893fb654c_mapped, var_f13893fb654c_err = types.ByResourcePropertyType(model.ResourceProperty_OBJECT).Pack(*var_f13893fb654c)
		if var_f13893fb654c_err != nil {
			panic(var_f13893fb654c_err)
		}
		properties["details"] = var_f13893fb654c_mapped
	}
	return properties
}

func (m *RoleMapper) FromProperties(properties map[string]*structpb.Value) *Role {
	var s = m.New()
	if properties["id"] != nil {

		var_424954a179b0 := properties["id"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_UUID).UnPack(var_424954a179b0)

		if err != nil {
			panic(err)
		}

		var_424954a179b0_mapped := new(uuid.UUID)
		*var_424954a179b0_mapped = val.(uuid.UUID)

		s.Id = var_424954a179b0_mapped
	}
	if properties["version"] != nil {

		var_7ecdde026e00 := properties["version"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(var_7ecdde026e00)

		if err != nil {
			panic(err)
		}

		var_7ecdde026e00_mapped := val.(int32)

		s.Version = var_7ecdde026e00_mapped
	}
	if properties["createdBy"] != nil {

		var_910099280086 := properties["createdBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_910099280086)

		if err != nil {
			panic(err)
		}

		var_910099280086_mapped := new(string)
		*var_910099280086_mapped = val.(string)

		s.CreatedBy = var_910099280086_mapped
	}
	if properties["updatedBy"] != nil {

		var_60b2e7cd4913 := properties["updatedBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_60b2e7cd4913)

		if err != nil {
			panic(err)
		}

		var_60b2e7cd4913_mapped := new(string)
		*var_60b2e7cd4913_mapped = val.(string)

		s.UpdatedBy = var_60b2e7cd4913_mapped
	}
	if properties["createdOn"] != nil {

		var_d27a6f4c472b := properties["createdOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_d27a6f4c472b)

		if err != nil {
			panic(err)
		}

		var_d27a6f4c472b_mapped := new(time.Time)
		*var_d27a6f4c472b_mapped = val.(time.Time)

		s.CreatedOn = var_d27a6f4c472b_mapped
	}
	if properties["updatedOn"] != nil {

		var_42f816164b2d := properties["updatedOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_42f816164b2d)

		if err != nil {
			panic(err)
		}

		var_42f816164b2d_mapped := new(time.Time)
		*var_42f816164b2d_mapped = val.(time.Time)

		s.UpdatedOn = var_42f816164b2d_mapped
	}
	if properties["name"] != nil {

		var_d2708b9f5f45 := properties["name"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_d2708b9f5f45)

		if err != nil {
			panic(err)
		}

		var_d2708b9f5f45_mapped := val.(string)

		s.Name = var_d2708b9f5f45_mapped
	}
	if properties["securityConstraints"] != nil {

		var_ac8dfaba5ea9 := properties["securityConstraints"]
		var_ac8dfaba5ea9_mapped := []*SecurityConstraint{}
		for _, v := range var_ac8dfaba5ea9.GetListValue().Values {

			var_e5e1df8d1b4a := v
			var_e5e1df8d1b4a_mapped := SecurityConstraintMapperInstance.FromProperties(var_e5e1df8d1b4a.GetStructValue().Fields)

			var_ac8dfaba5ea9_mapped = append(var_ac8dfaba5ea9_mapped, var_e5e1df8d1b4a_mapped)
		}

		s.SecurityConstraints = var_ac8dfaba5ea9_mapped
	}
	if properties["details"] != nil {

		var_6c6781a4b6a5 := properties["details"]
		var_6c6781a4b6a5_mapped := new(unstructured.Unstructured)
		*var_6c6781a4b6a5_mapped = unstructured.FromStructValue(var_6c6781a4b6a5.GetStructValue())

		s.Details = var_6c6781a4b6a5_mapped
	}
	return s
}
