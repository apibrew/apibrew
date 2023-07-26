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

	var_58865f6d4aa9 := role.Id

	if var_58865f6d4aa9 != nil {
		var var_58865f6d4aa9_mapped *structpb.Value

		var var_58865f6d4aa9_err error
		var_58865f6d4aa9_mapped, var_58865f6d4aa9_err = types.ByResourcePropertyType(model.ResourceProperty_UUID).Pack(*var_58865f6d4aa9)
		if var_58865f6d4aa9_err != nil {
			panic(var_58865f6d4aa9_err)
		}
		properties["id"] = var_58865f6d4aa9_mapped
	}

	var_7b53420e1ba8 := role.Version

	var var_7b53420e1ba8_mapped *structpb.Value

	var var_7b53420e1ba8_err error
	var_7b53420e1ba8_mapped, var_7b53420e1ba8_err = types.ByResourcePropertyType(model.ResourceProperty_INT32).Pack(var_7b53420e1ba8)
	if var_7b53420e1ba8_err != nil {
		panic(var_7b53420e1ba8_err)
	}
	properties["version"] = var_7b53420e1ba8_mapped

	var_9600c5e14961 := role.CreatedBy

	if var_9600c5e14961 != nil {
		var var_9600c5e14961_mapped *structpb.Value

		var var_9600c5e14961_err error
		var_9600c5e14961_mapped, var_9600c5e14961_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_9600c5e14961)
		if var_9600c5e14961_err != nil {
			panic(var_9600c5e14961_err)
		}
		properties["createdBy"] = var_9600c5e14961_mapped
	}

	var_494aba29e9ce := role.UpdatedBy

	if var_494aba29e9ce != nil {
		var var_494aba29e9ce_mapped *structpb.Value

		var var_494aba29e9ce_err error
		var_494aba29e9ce_mapped, var_494aba29e9ce_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_494aba29e9ce)
		if var_494aba29e9ce_err != nil {
			panic(var_494aba29e9ce_err)
		}
		properties["updatedBy"] = var_494aba29e9ce_mapped
	}

	var_46d024736d9c := role.CreatedOn

	if var_46d024736d9c != nil {
		var var_46d024736d9c_mapped *structpb.Value

		var var_46d024736d9c_err error
		var_46d024736d9c_mapped, var_46d024736d9c_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_46d024736d9c)
		if var_46d024736d9c_err != nil {
			panic(var_46d024736d9c_err)
		}
		properties["createdOn"] = var_46d024736d9c_mapped
	}

	var_3f76699b2be0 := role.UpdatedOn

	if var_3f76699b2be0 != nil {
		var var_3f76699b2be0_mapped *structpb.Value

		var var_3f76699b2be0_err error
		var_3f76699b2be0_mapped, var_3f76699b2be0_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_3f76699b2be0)
		if var_3f76699b2be0_err != nil {
			panic(var_3f76699b2be0_err)
		}
		properties["updatedOn"] = var_3f76699b2be0_mapped
	}

	var_5efba981326d := role.Name

	var var_5efba981326d_mapped *structpb.Value

	var var_5efba981326d_err error
	var_5efba981326d_mapped, var_5efba981326d_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_5efba981326d)
	if var_5efba981326d_err != nil {
		panic(var_5efba981326d_err)
	}
	properties["name"] = var_5efba981326d_mapped

	var_495da6bf9e35 := role.SecurityConstraints

	if var_495da6bf9e35 != nil {
		var var_495da6bf9e35_mapped *structpb.Value

		var var_495da6bf9e35_l []*structpb.Value
		for _, value := range var_495da6bf9e35 {

			var_ceef1060b722 := value
			var var_ceef1060b722_mapped *structpb.Value

			var_ceef1060b722_mapped = structpb.NewStructValue(&structpb.Struct{Fields: SecurityConstraintMapperInstance.ToProperties(var_ceef1060b722)})

			var_495da6bf9e35_l = append(var_495da6bf9e35_l, var_ceef1060b722_mapped)
		}
		var_495da6bf9e35_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_495da6bf9e35_l})
		properties["securityConstraints"] = var_495da6bf9e35_mapped
	}

	var_0ba3e9469547 := role.Details

	if var_0ba3e9469547 != nil {
		var var_0ba3e9469547_mapped *structpb.Value

		var var_0ba3e9469547_err error
		var_0ba3e9469547_mapped, var_0ba3e9469547_err = types.ByResourcePropertyType(model.ResourceProperty_OBJECT).Pack(*var_0ba3e9469547)
		if var_0ba3e9469547_err != nil {
			panic(var_0ba3e9469547_err)
		}
		properties["details"] = var_0ba3e9469547_mapped
	}
	return properties
}

func (m *RoleMapper) FromProperties(properties map[string]*structpb.Value) *Role {
	var s = m.New()
	if properties["id"] != nil {

		var_6f0d661b04d9 := properties["id"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_UUID).UnPack(var_6f0d661b04d9)

		if err != nil {
			panic(err)
		}

		var_6f0d661b04d9_mapped := new(uuid.UUID)
		*var_6f0d661b04d9_mapped = val.(uuid.UUID)

		s.Id = var_6f0d661b04d9_mapped
	}
	if properties["version"] != nil {

		var_7060e7db1ef7 := properties["version"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(var_7060e7db1ef7)

		if err != nil {
			panic(err)
		}

		var_7060e7db1ef7_mapped := val.(int32)

		s.Version = var_7060e7db1ef7_mapped
	}
	if properties["createdBy"] != nil {

		var_7b67bbe8a7eb := properties["createdBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_7b67bbe8a7eb)

		if err != nil {
			panic(err)
		}

		var_7b67bbe8a7eb_mapped := new(string)
		*var_7b67bbe8a7eb_mapped = val.(string)

		s.CreatedBy = var_7b67bbe8a7eb_mapped
	}
	if properties["updatedBy"] != nil {

		var_e1056dfe4a6e := properties["updatedBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_e1056dfe4a6e)

		if err != nil {
			panic(err)
		}

		var_e1056dfe4a6e_mapped := new(string)
		*var_e1056dfe4a6e_mapped = val.(string)

		s.UpdatedBy = var_e1056dfe4a6e_mapped
	}
	if properties["createdOn"] != nil {

		var_97792d51fbc4 := properties["createdOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_97792d51fbc4)

		if err != nil {
			panic(err)
		}

		var_97792d51fbc4_mapped := new(time.Time)
		*var_97792d51fbc4_mapped = val.(time.Time)

		s.CreatedOn = var_97792d51fbc4_mapped
	}
	if properties["updatedOn"] != nil {

		var_9b20c1b27642 := properties["updatedOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_9b20c1b27642)

		if err != nil {
			panic(err)
		}

		var_9b20c1b27642_mapped := new(time.Time)
		*var_9b20c1b27642_mapped = val.(time.Time)

		s.UpdatedOn = var_9b20c1b27642_mapped
	}
	if properties["name"] != nil {

		var_bf72bca97593 := properties["name"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_bf72bca97593)

		if err != nil {
			panic(err)
		}

		var_bf72bca97593_mapped := val.(string)

		s.Name = var_bf72bca97593_mapped
	}
	if properties["securityConstraints"] != nil {

		var_925ee6fbcba8 := properties["securityConstraints"]
		var_925ee6fbcba8_mapped := []*SecurityConstraint{}
		for _, v := range var_925ee6fbcba8.GetListValue().Values {

			var_82b4480d62f6 := v
			var_82b4480d62f6_mapped := SecurityConstraintMapperInstance.FromProperties(var_82b4480d62f6.GetStructValue().Fields)

			var_925ee6fbcba8_mapped = append(var_925ee6fbcba8_mapped, var_82b4480d62f6_mapped)
		}

		s.SecurityConstraints = var_925ee6fbcba8_mapped
	}
	if properties["details"] != nil {

		var_e73d0a988dfa := properties["details"]
		var_e73d0a988dfa_mapped := new(unstructured.Unstructured)
		*var_e73d0a988dfa_mapped = unstructured.FromStructValue(var_e73d0a988dfa.GetStructValue())

		s.Details = var_e73d0a988dfa_mapped
	}
	return s
}
