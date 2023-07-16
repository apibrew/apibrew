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

	var_e0f488fa2af9 := role.Id

	if var_e0f488fa2af9 != nil {
		var var_e0f488fa2af9_mapped *structpb.Value

		var var_e0f488fa2af9_err error
		var_e0f488fa2af9_mapped, var_e0f488fa2af9_err = types.ByResourcePropertyType(model.ResourceProperty_UUID).Pack(*var_e0f488fa2af9)
		if var_e0f488fa2af9_err != nil {
			panic(var_e0f488fa2af9_err)
		}
		properties["id"] = var_e0f488fa2af9_mapped
	}

	var_f8c7729b338e := role.Version

	var var_f8c7729b338e_mapped *structpb.Value

	var var_f8c7729b338e_err error
	var_f8c7729b338e_mapped, var_f8c7729b338e_err = types.ByResourcePropertyType(model.ResourceProperty_INT32).Pack(var_f8c7729b338e)
	if var_f8c7729b338e_err != nil {
		panic(var_f8c7729b338e_err)
	}
	properties["version"] = var_f8c7729b338e_mapped

	var_c2d65ee55447 := role.CreatedBy

	if var_c2d65ee55447 != nil {
		var var_c2d65ee55447_mapped *structpb.Value

		var var_c2d65ee55447_err error
		var_c2d65ee55447_mapped, var_c2d65ee55447_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_c2d65ee55447)
		if var_c2d65ee55447_err != nil {
			panic(var_c2d65ee55447_err)
		}
		properties["createdBy"] = var_c2d65ee55447_mapped
	}

	var_479791f6e359 := role.UpdatedBy

	if var_479791f6e359 != nil {
		var var_479791f6e359_mapped *structpb.Value

		var var_479791f6e359_err error
		var_479791f6e359_mapped, var_479791f6e359_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_479791f6e359)
		if var_479791f6e359_err != nil {
			panic(var_479791f6e359_err)
		}
		properties["updatedBy"] = var_479791f6e359_mapped
	}

	var_e72dce3099d4 := role.CreatedOn

	if var_e72dce3099d4 != nil {
		var var_e72dce3099d4_mapped *structpb.Value

		var var_e72dce3099d4_err error
		var_e72dce3099d4_mapped, var_e72dce3099d4_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_e72dce3099d4)
		if var_e72dce3099d4_err != nil {
			panic(var_e72dce3099d4_err)
		}
		properties["createdOn"] = var_e72dce3099d4_mapped
	}

	var_35eeb63e8535 := role.UpdatedOn

	if var_35eeb63e8535 != nil {
		var var_35eeb63e8535_mapped *structpb.Value

		var var_35eeb63e8535_err error
		var_35eeb63e8535_mapped, var_35eeb63e8535_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_35eeb63e8535)
		if var_35eeb63e8535_err != nil {
			panic(var_35eeb63e8535_err)
		}
		properties["updatedOn"] = var_35eeb63e8535_mapped
	}

	var_625d10ec7ed8 := role.Name

	var var_625d10ec7ed8_mapped *structpb.Value

	var var_625d10ec7ed8_err error
	var_625d10ec7ed8_mapped, var_625d10ec7ed8_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_625d10ec7ed8)
	if var_625d10ec7ed8_err != nil {
		panic(var_625d10ec7ed8_err)
	}
	properties["name"] = var_625d10ec7ed8_mapped

	var_ab54aef5a096 := role.SecurityConstraints

	if var_ab54aef5a096 != nil {
		var var_ab54aef5a096_mapped *structpb.Value

		var var_ab54aef5a096_l []*structpb.Value
		for _, value := range var_ab54aef5a096 {

			var_78077f8db528 := value
			var var_78077f8db528_mapped *structpb.Value

			var_78077f8db528_mapped = structpb.NewStructValue(&structpb.Struct{Fields: SecurityConstraintMapperInstance.ToProperties(var_78077f8db528)})

			var_ab54aef5a096_l = append(var_ab54aef5a096_l, var_78077f8db528_mapped)
		}
		var_ab54aef5a096_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_ab54aef5a096_l})
		properties["securityConstraints"] = var_ab54aef5a096_mapped
	}

	var_5b6ce69be9cc := role.Details

	if var_5b6ce69be9cc != nil {
		var var_5b6ce69be9cc_mapped *structpb.Value

		var var_5b6ce69be9cc_err error
		var_5b6ce69be9cc_mapped, var_5b6ce69be9cc_err = types.ByResourcePropertyType(model.ResourceProperty_OBJECT).Pack(*var_5b6ce69be9cc)
		if var_5b6ce69be9cc_err != nil {
			panic(var_5b6ce69be9cc_err)
		}
		properties["details"] = var_5b6ce69be9cc_mapped
	}
	return properties
}

func (m *RoleMapper) FromProperties(properties map[string]*structpb.Value) *Role {
	var s = m.New()
	if properties["id"] != nil {

		var_1906fa728790 := properties["id"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_UUID).UnPack(var_1906fa728790)

		if err != nil {
			panic(err)
		}

		var_1906fa728790_mapped := new(uuid.UUID)
		*var_1906fa728790_mapped = val.(uuid.UUID)

		s.Id = var_1906fa728790_mapped
	}
	if properties["version"] != nil {

		var_a5157d00ed7d := properties["version"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(var_a5157d00ed7d)

		if err != nil {
			panic(err)
		}

		var_a5157d00ed7d_mapped := val.(int32)

		s.Version = var_a5157d00ed7d_mapped
	}
	if properties["createdBy"] != nil {

		var_6e26c025450b := properties["createdBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_6e26c025450b)

		if err != nil {
			panic(err)
		}

		var_6e26c025450b_mapped := new(string)
		*var_6e26c025450b_mapped = val.(string)

		s.CreatedBy = var_6e26c025450b_mapped
	}
	if properties["updatedBy"] != nil {

		var_6e775c2b5a49 := properties["updatedBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_6e775c2b5a49)

		if err != nil {
			panic(err)
		}

		var_6e775c2b5a49_mapped := new(string)
		*var_6e775c2b5a49_mapped = val.(string)

		s.UpdatedBy = var_6e775c2b5a49_mapped
	}
	if properties["createdOn"] != nil {

		var_41440131a557 := properties["createdOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_41440131a557)

		if err != nil {
			panic(err)
		}

		var_41440131a557_mapped := new(time.Time)
		*var_41440131a557_mapped = val.(time.Time)

		s.CreatedOn = var_41440131a557_mapped
	}
	if properties["updatedOn"] != nil {

		var_97942a3e62c3 := properties["updatedOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_97942a3e62c3)

		if err != nil {
			panic(err)
		}

		var_97942a3e62c3_mapped := new(time.Time)
		*var_97942a3e62c3_mapped = val.(time.Time)

		s.UpdatedOn = var_97942a3e62c3_mapped
	}
	if properties["name"] != nil {

		var_affc495af12c := properties["name"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_affc495af12c)

		if err != nil {
			panic(err)
		}

		var_affc495af12c_mapped := val.(string)

		s.Name = var_affc495af12c_mapped
	}
	if properties["securityConstraints"] != nil {

		var_faa5b018ee63 := properties["securityConstraints"]
		var_faa5b018ee63_mapped := []*SecurityConstraint{}
		for _, v := range var_faa5b018ee63.GetListValue().Values {

			var_d7e8d742306e := v
			var_d7e8d742306e_mapped := SecurityConstraintMapperInstance.FromProperties(var_d7e8d742306e.GetStructValue().Fields)

			var_faa5b018ee63_mapped = append(var_faa5b018ee63_mapped, var_d7e8d742306e_mapped)
		}

		s.SecurityConstraints = var_faa5b018ee63_mapped
	}
	if properties["details"] != nil {

		var_abc6e9d01e0c := properties["details"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_OBJECT).UnPack(var_abc6e9d01e0c)

		if err != nil {
			panic(err)
		}

		var_abc6e9d01e0c_mapped := new(unstructured.Unstructured)
		*var_abc6e9d01e0c_mapped = val.(unstructured.Unstructured)

		s.Details = var_abc6e9d01e0c_mapped
	}
	return s
}
