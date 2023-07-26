package resource_model

import (
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/types"
	"google.golang.org/protobuf/types/known/structpb"
)

import "github.com/google/uuid"
import "time"
import "github.com/apibrew/apibrew/pkg/formats/unstructured"

type UserMapper struct {
}

func NewUserMapper() *UserMapper {
	return &UserMapper{}
}

var UserMapperInstance = NewUserMapper()

func (m *UserMapper) New() *User {
	return &User{}
}

func (m *UserMapper) ToRecord(user *User) *model.Record {
	var rec = &model.Record{}
	rec.Properties = m.ToProperties(user)

	if user.Id != nil {
		rec.Id = user.Id.String()
	}

	return rec
}

func (m *UserMapper) FromRecord(record *model.Record) *User {
	return m.FromProperties(record.Properties)
}

func (m *UserMapper) ToProperties(user *User) map[string]*structpb.Value {
	var properties = make(map[string]*structpb.Value)

	var_bdd559906168 := user.Id

	if var_bdd559906168 != nil {
		var var_bdd559906168_mapped *structpb.Value

		var var_bdd559906168_err error
		var_bdd559906168_mapped, var_bdd559906168_err = types.ByResourcePropertyType(model.ResourceProperty_UUID).Pack(*var_bdd559906168)
		if var_bdd559906168_err != nil {
			panic(var_bdd559906168_err)
		}
		properties["id"] = var_bdd559906168_mapped
	}

	var_55c3919b0bca := user.Version

	var var_55c3919b0bca_mapped *structpb.Value

	var var_55c3919b0bca_err error
	var_55c3919b0bca_mapped, var_55c3919b0bca_err = types.ByResourcePropertyType(model.ResourceProperty_INT32).Pack(var_55c3919b0bca)
	if var_55c3919b0bca_err != nil {
		panic(var_55c3919b0bca_err)
	}
	properties["version"] = var_55c3919b0bca_mapped

	var_88a37710570e := user.CreatedBy

	if var_88a37710570e != nil {
		var var_88a37710570e_mapped *structpb.Value

		var var_88a37710570e_err error
		var_88a37710570e_mapped, var_88a37710570e_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_88a37710570e)
		if var_88a37710570e_err != nil {
			panic(var_88a37710570e_err)
		}
		properties["createdBy"] = var_88a37710570e_mapped
	}

	var_dfd88d1b57c4 := user.UpdatedBy

	if var_dfd88d1b57c4 != nil {
		var var_dfd88d1b57c4_mapped *structpb.Value

		var var_dfd88d1b57c4_err error
		var_dfd88d1b57c4_mapped, var_dfd88d1b57c4_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_dfd88d1b57c4)
		if var_dfd88d1b57c4_err != nil {
			panic(var_dfd88d1b57c4_err)
		}
		properties["updatedBy"] = var_dfd88d1b57c4_mapped
	}

	var_d9a7f0b40e3d := user.CreatedOn

	if var_d9a7f0b40e3d != nil {
		var var_d9a7f0b40e3d_mapped *structpb.Value

		var var_d9a7f0b40e3d_err error
		var_d9a7f0b40e3d_mapped, var_d9a7f0b40e3d_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_d9a7f0b40e3d)
		if var_d9a7f0b40e3d_err != nil {
			panic(var_d9a7f0b40e3d_err)
		}
		properties["createdOn"] = var_d9a7f0b40e3d_mapped
	}

	var_7cb00cadfc0f := user.UpdatedOn

	if var_7cb00cadfc0f != nil {
		var var_7cb00cadfc0f_mapped *structpb.Value

		var var_7cb00cadfc0f_err error
		var_7cb00cadfc0f_mapped, var_7cb00cadfc0f_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_7cb00cadfc0f)
		if var_7cb00cadfc0f_err != nil {
			panic(var_7cb00cadfc0f_err)
		}
		properties["updatedOn"] = var_7cb00cadfc0f_mapped
	}

	var_ba7fdc8b44b2 := user.Username

	var var_ba7fdc8b44b2_mapped *structpb.Value

	var var_ba7fdc8b44b2_err error
	var_ba7fdc8b44b2_mapped, var_ba7fdc8b44b2_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_ba7fdc8b44b2)
	if var_ba7fdc8b44b2_err != nil {
		panic(var_ba7fdc8b44b2_err)
	}
	properties["username"] = var_ba7fdc8b44b2_mapped

	var_4e9d2a2e0539 := user.Password

	if var_4e9d2a2e0539 != nil {
		var var_4e9d2a2e0539_mapped *structpb.Value

		var var_4e9d2a2e0539_err error
		var_4e9d2a2e0539_mapped, var_4e9d2a2e0539_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_4e9d2a2e0539)
		if var_4e9d2a2e0539_err != nil {
			panic(var_4e9d2a2e0539_err)
		}
		properties["password"] = var_4e9d2a2e0539_mapped
	}

	var_6d48878aeb32 := user.Roles

	if var_6d48878aeb32 != nil {
		var var_6d48878aeb32_mapped *structpb.Value

		var var_6d48878aeb32_l []*structpb.Value
		for _, value := range var_6d48878aeb32 {

			var_1b540998c6dc := value
			var var_1b540998c6dc_mapped *structpb.Value

			var_1b540998c6dc_mapped = structpb.NewStructValue(&structpb.Struct{Fields: RoleMapperInstance.ToProperties(var_1b540998c6dc)})

			var_6d48878aeb32_l = append(var_6d48878aeb32_l, var_1b540998c6dc_mapped)
		}
		var_6d48878aeb32_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_6d48878aeb32_l})
		properties["roles"] = var_6d48878aeb32_mapped
	}

	var_c1df7f127c9a := user.SecurityConstraints

	if var_c1df7f127c9a != nil {
		var var_c1df7f127c9a_mapped *structpb.Value

		var var_c1df7f127c9a_l []*structpb.Value
		for _, value := range var_c1df7f127c9a {

			var_63d5c2cd635f := value
			var var_63d5c2cd635f_mapped *structpb.Value

			var_63d5c2cd635f_mapped = structpb.NewStructValue(&structpb.Struct{Fields: SecurityConstraintMapperInstance.ToProperties(var_63d5c2cd635f)})

			var_c1df7f127c9a_l = append(var_c1df7f127c9a_l, var_63d5c2cd635f_mapped)
		}
		var_c1df7f127c9a_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_c1df7f127c9a_l})
		properties["securityConstraints"] = var_c1df7f127c9a_mapped
	}

	var_b1035b2a5cfe := user.Details

	if var_b1035b2a5cfe != nil {
		var var_b1035b2a5cfe_mapped *structpb.Value

		var var_b1035b2a5cfe_err error
		var_b1035b2a5cfe_mapped, var_b1035b2a5cfe_err = types.ByResourcePropertyType(model.ResourceProperty_OBJECT).Pack(*var_b1035b2a5cfe)
		if var_b1035b2a5cfe_err != nil {
			panic(var_b1035b2a5cfe_err)
		}
		properties["details"] = var_b1035b2a5cfe_mapped
	}
	return properties
}

func (m *UserMapper) FromProperties(properties map[string]*structpb.Value) *User {
	var s = m.New()
	if properties["id"] != nil {

		var_9e8e6a3f00cb := properties["id"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_UUID).UnPack(var_9e8e6a3f00cb)

		if err != nil {
			panic(err)
		}

		var_9e8e6a3f00cb_mapped := new(uuid.UUID)
		*var_9e8e6a3f00cb_mapped = val.(uuid.UUID)

		s.Id = var_9e8e6a3f00cb_mapped
	}
	if properties["version"] != nil {

		var_235594e72601 := properties["version"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(var_235594e72601)

		if err != nil {
			panic(err)
		}

		var_235594e72601_mapped := val.(int32)

		s.Version = var_235594e72601_mapped
	}
	if properties["createdBy"] != nil {

		var_b736adf0e16c := properties["createdBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_b736adf0e16c)

		if err != nil {
			panic(err)
		}

		var_b736adf0e16c_mapped := new(string)
		*var_b736adf0e16c_mapped = val.(string)

		s.CreatedBy = var_b736adf0e16c_mapped
	}
	if properties["updatedBy"] != nil {

		var_f0cffbcef397 := properties["updatedBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_f0cffbcef397)

		if err != nil {
			panic(err)
		}

		var_f0cffbcef397_mapped := new(string)
		*var_f0cffbcef397_mapped = val.(string)

		s.UpdatedBy = var_f0cffbcef397_mapped
	}
	if properties["createdOn"] != nil {

		var_c51972d282a6 := properties["createdOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_c51972d282a6)

		if err != nil {
			panic(err)
		}

		var_c51972d282a6_mapped := new(time.Time)
		*var_c51972d282a6_mapped = val.(time.Time)

		s.CreatedOn = var_c51972d282a6_mapped
	}
	if properties["updatedOn"] != nil {

		var_23e239fb07c6 := properties["updatedOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_23e239fb07c6)

		if err != nil {
			panic(err)
		}

		var_23e239fb07c6_mapped := new(time.Time)
		*var_23e239fb07c6_mapped = val.(time.Time)

		s.UpdatedOn = var_23e239fb07c6_mapped
	}
	if properties["username"] != nil {

		var_159682ef8260 := properties["username"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_159682ef8260)

		if err != nil {
			panic(err)
		}

		var_159682ef8260_mapped := val.(string)

		s.Username = var_159682ef8260_mapped
	}
	if properties["password"] != nil {

		var_32a63a52f716 := properties["password"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_32a63a52f716)

		if err != nil {
			panic(err)
		}

		var_32a63a52f716_mapped := new(string)
		*var_32a63a52f716_mapped = val.(string)

		s.Password = var_32a63a52f716_mapped
	}
	if properties["roles"] != nil {

		var_4d2a04b07a59 := properties["roles"]
		var_4d2a04b07a59_mapped := []*Role{}
		for _, v := range var_4d2a04b07a59.GetListValue().Values {

			var_f0abfc16a4d1 := v
			var_f0abfc16a4d1_mapped := RoleMapperInstance.FromProperties(var_f0abfc16a4d1.GetStructValue().Fields)

			var_4d2a04b07a59_mapped = append(var_4d2a04b07a59_mapped, var_f0abfc16a4d1_mapped)
		}

		s.Roles = var_4d2a04b07a59_mapped
	}
	if properties["securityConstraints"] != nil {

		var_d244e80d8403 := properties["securityConstraints"]
		var_d244e80d8403_mapped := []*SecurityConstraint{}
		for _, v := range var_d244e80d8403.GetListValue().Values {

			var_51a8ff30b8f7 := v
			var_51a8ff30b8f7_mapped := SecurityConstraintMapperInstance.FromProperties(var_51a8ff30b8f7.GetStructValue().Fields)

			var_d244e80d8403_mapped = append(var_d244e80d8403_mapped, var_51a8ff30b8f7_mapped)
		}

		s.SecurityConstraints = var_d244e80d8403_mapped
	}
	if properties["details"] != nil {

		var_9a88d048df37 := properties["details"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_OBJECT).UnPack(var_9a88d048df37)

		if err != nil {
			panic(err)
		}

		var_9a88d048df37_mapped := new(unstructured.Unstructured)
		*var_9a88d048df37_mapped = val.(unstructured.Unstructured)

		s.Details = var_9a88d048df37_mapped
	}
	return s
}
