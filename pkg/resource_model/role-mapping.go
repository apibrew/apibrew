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

	var_5129d3c1e0e4 := role.Id

	if var_5129d3c1e0e4 != nil {
		var var_5129d3c1e0e4_mapped *structpb.Value

		var var_5129d3c1e0e4_err error
		var_5129d3c1e0e4_mapped, var_5129d3c1e0e4_err = types.ByResourcePropertyType(model.ResourceProperty_UUID).Pack(*var_5129d3c1e0e4)
		if var_5129d3c1e0e4_err != nil {
			panic(var_5129d3c1e0e4_err)
		}
		properties["id"] = var_5129d3c1e0e4_mapped
	}

	var_578704718274 := role.Version

	var var_578704718274_mapped *structpb.Value

	var var_578704718274_err error
	var_578704718274_mapped, var_578704718274_err = types.ByResourcePropertyType(model.ResourceProperty_INT32).Pack(var_578704718274)
	if var_578704718274_err != nil {
		panic(var_578704718274_err)
	}
	properties["version"] = var_578704718274_mapped

	var_ef4b1800599f := role.CreatedBy

	if var_ef4b1800599f != nil {
		var var_ef4b1800599f_mapped *structpb.Value

		var var_ef4b1800599f_err error
		var_ef4b1800599f_mapped, var_ef4b1800599f_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_ef4b1800599f)
		if var_ef4b1800599f_err != nil {
			panic(var_ef4b1800599f_err)
		}
		properties["createdBy"] = var_ef4b1800599f_mapped
	}

	var_fc17c618461c := role.UpdatedBy

	if var_fc17c618461c != nil {
		var var_fc17c618461c_mapped *structpb.Value

		var var_fc17c618461c_err error
		var_fc17c618461c_mapped, var_fc17c618461c_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_fc17c618461c)
		if var_fc17c618461c_err != nil {
			panic(var_fc17c618461c_err)
		}
		properties["updatedBy"] = var_fc17c618461c_mapped
	}

	var_62fcffccd25c := role.CreatedOn

	if var_62fcffccd25c != nil {
		var var_62fcffccd25c_mapped *structpb.Value

		var var_62fcffccd25c_err error
		var_62fcffccd25c_mapped, var_62fcffccd25c_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_62fcffccd25c)
		if var_62fcffccd25c_err != nil {
			panic(var_62fcffccd25c_err)
		}
		properties["createdOn"] = var_62fcffccd25c_mapped
	}

	var_bc34fa453a57 := role.UpdatedOn

	if var_bc34fa453a57 != nil {
		var var_bc34fa453a57_mapped *structpb.Value

		var var_bc34fa453a57_err error
		var_bc34fa453a57_mapped, var_bc34fa453a57_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_bc34fa453a57)
		if var_bc34fa453a57_err != nil {
			panic(var_bc34fa453a57_err)
		}
		properties["updatedOn"] = var_bc34fa453a57_mapped
	}

	var_42c450eb105e := role.Name

	var var_42c450eb105e_mapped *structpb.Value

	var var_42c450eb105e_err error
	var_42c450eb105e_mapped, var_42c450eb105e_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_42c450eb105e)
	if var_42c450eb105e_err != nil {
		panic(var_42c450eb105e_err)
	}
	properties["name"] = var_42c450eb105e_mapped

	var_1a80422a5953 := role.SecurityConstraints

	if var_1a80422a5953 != nil {
		var var_1a80422a5953_mapped *structpb.Value

		var var_1a80422a5953_l []*structpb.Value
		for _, value := range var_1a80422a5953 {

			var_101aa8d215b7 := value
			var var_101aa8d215b7_mapped *structpb.Value

			var_101aa8d215b7_mapped = structpb.NewStructValue(&structpb.Struct{Fields: SecurityConstraintMapperInstance.ToProperties(var_101aa8d215b7)})

			var_1a80422a5953_l = append(var_1a80422a5953_l, var_101aa8d215b7_mapped)
		}
		var_1a80422a5953_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_1a80422a5953_l})
		properties["securityConstraints"] = var_1a80422a5953_mapped
	}

	var_0e2cc9fa43e2 := role.Details

	if var_0e2cc9fa43e2 != nil {
		var var_0e2cc9fa43e2_mapped *structpb.Value

		var var_0e2cc9fa43e2_err error
		var_0e2cc9fa43e2_mapped, var_0e2cc9fa43e2_err = types.ByResourcePropertyType(model.ResourceProperty_OBJECT).Pack(*var_0e2cc9fa43e2)
		if var_0e2cc9fa43e2_err != nil {
			panic(var_0e2cc9fa43e2_err)
		}
		properties["details"] = var_0e2cc9fa43e2_mapped
	}
	return properties
}

func (m *RoleMapper) FromProperties(properties map[string]*structpb.Value) *Role {
	var s = m.New()
	if properties["id"] != nil {

		var_7fb6b3a7b34d := properties["id"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_UUID).UnPack(var_7fb6b3a7b34d)

		if err != nil {
			panic(err)
		}

		var_7fb6b3a7b34d_mapped := new(uuid.UUID)
		*var_7fb6b3a7b34d_mapped = val.(uuid.UUID)

		s.Id = var_7fb6b3a7b34d_mapped
	}
	if properties["version"] != nil {

		var_a887d8fe7dc1 := properties["version"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(var_a887d8fe7dc1)

		if err != nil {
			panic(err)
		}

		var_a887d8fe7dc1_mapped := val.(int32)

		s.Version = var_a887d8fe7dc1_mapped
	}
	if properties["createdBy"] != nil {

		var_5c1095f4413c := properties["createdBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_5c1095f4413c)

		if err != nil {
			panic(err)
		}

		var_5c1095f4413c_mapped := new(string)
		*var_5c1095f4413c_mapped = val.(string)

		s.CreatedBy = var_5c1095f4413c_mapped
	}
	if properties["updatedBy"] != nil {

		var_fee06e37de6f := properties["updatedBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_fee06e37de6f)

		if err != nil {
			panic(err)
		}

		var_fee06e37de6f_mapped := new(string)
		*var_fee06e37de6f_mapped = val.(string)

		s.UpdatedBy = var_fee06e37de6f_mapped
	}
	if properties["createdOn"] != nil {

		var_e4e7c59a83dc := properties["createdOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_e4e7c59a83dc)

		if err != nil {
			panic(err)
		}

		var_e4e7c59a83dc_mapped := new(time.Time)
		*var_e4e7c59a83dc_mapped = val.(time.Time)

		s.CreatedOn = var_e4e7c59a83dc_mapped
	}
	if properties["updatedOn"] != nil {

		var_a97b7b84b0ac := properties["updatedOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_a97b7b84b0ac)

		if err != nil {
			panic(err)
		}

		var_a97b7b84b0ac_mapped := new(time.Time)
		*var_a97b7b84b0ac_mapped = val.(time.Time)

		s.UpdatedOn = var_a97b7b84b0ac_mapped
	}
	if properties["name"] != nil {

		var_a0a19d49c9fd := properties["name"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_a0a19d49c9fd)

		if err != nil {
			panic(err)
		}

		var_a0a19d49c9fd_mapped := val.(string)

		s.Name = var_a0a19d49c9fd_mapped
	}
	if properties["securityConstraints"] != nil {

		var_e10e5890db3e := properties["securityConstraints"]
		var_e10e5890db3e_mapped := []*SecurityConstraint{}
		for _, v := range var_e10e5890db3e.GetListValue().Values {

			var_301bec8765f6 := v
			var_301bec8765f6_mapped := SecurityConstraintMapperInstance.FromProperties(var_301bec8765f6.GetStructValue().Fields)

			var_e10e5890db3e_mapped = append(var_e10e5890db3e_mapped, var_301bec8765f6_mapped)
		}

		s.SecurityConstraints = var_e10e5890db3e_mapped
	}
	if properties["details"] != nil {

		var_e21a995f1dde := properties["details"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_OBJECT).UnPack(var_e21a995f1dde)

		if err != nil {
			panic(err)
		}

		var_e21a995f1dde_mapped := new(unstructured.Unstructured)
		*var_e21a995f1dde_mapped = val.(unstructured.Unstructured)

		s.Details = var_e21a995f1dde_mapped
	}
	return s
}
