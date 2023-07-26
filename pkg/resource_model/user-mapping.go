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

	var_74d9ea1a5ff5 := user.Id

	if var_74d9ea1a5ff5 != nil {
		var var_74d9ea1a5ff5_mapped *structpb.Value

		var var_74d9ea1a5ff5_err error
		var_74d9ea1a5ff5_mapped, var_74d9ea1a5ff5_err = types.ByResourcePropertyType(model.ResourceProperty_UUID).Pack(*var_74d9ea1a5ff5)
		if var_74d9ea1a5ff5_err != nil {
			panic(var_74d9ea1a5ff5_err)
		}
		properties["id"] = var_74d9ea1a5ff5_mapped
	}

	var_f728826f12ae := user.Version

	var var_f728826f12ae_mapped *structpb.Value

	var var_f728826f12ae_err error
	var_f728826f12ae_mapped, var_f728826f12ae_err = types.ByResourcePropertyType(model.ResourceProperty_INT32).Pack(var_f728826f12ae)
	if var_f728826f12ae_err != nil {
		panic(var_f728826f12ae_err)
	}
	properties["version"] = var_f728826f12ae_mapped

	var_b27ea3ec3e8a := user.CreatedBy

	if var_b27ea3ec3e8a != nil {
		var var_b27ea3ec3e8a_mapped *structpb.Value

		var var_b27ea3ec3e8a_err error
		var_b27ea3ec3e8a_mapped, var_b27ea3ec3e8a_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_b27ea3ec3e8a)
		if var_b27ea3ec3e8a_err != nil {
			panic(var_b27ea3ec3e8a_err)
		}
		properties["createdBy"] = var_b27ea3ec3e8a_mapped
	}

	var_24cd2c6389ee := user.UpdatedBy

	if var_24cd2c6389ee != nil {
		var var_24cd2c6389ee_mapped *structpb.Value

		var var_24cd2c6389ee_err error
		var_24cd2c6389ee_mapped, var_24cd2c6389ee_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_24cd2c6389ee)
		if var_24cd2c6389ee_err != nil {
			panic(var_24cd2c6389ee_err)
		}
		properties["updatedBy"] = var_24cd2c6389ee_mapped
	}

	var_34da7e57e1e2 := user.CreatedOn

	if var_34da7e57e1e2 != nil {
		var var_34da7e57e1e2_mapped *structpb.Value

		var var_34da7e57e1e2_err error
		var_34da7e57e1e2_mapped, var_34da7e57e1e2_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_34da7e57e1e2)
		if var_34da7e57e1e2_err != nil {
			panic(var_34da7e57e1e2_err)
		}
		properties["createdOn"] = var_34da7e57e1e2_mapped
	}

	var_d8c87fb1a6f9 := user.UpdatedOn

	if var_d8c87fb1a6f9 != nil {
		var var_d8c87fb1a6f9_mapped *structpb.Value

		var var_d8c87fb1a6f9_err error
		var_d8c87fb1a6f9_mapped, var_d8c87fb1a6f9_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_d8c87fb1a6f9)
		if var_d8c87fb1a6f9_err != nil {
			panic(var_d8c87fb1a6f9_err)
		}
		properties["updatedOn"] = var_d8c87fb1a6f9_mapped
	}

	var_45f199d87213 := user.Username

	var var_45f199d87213_mapped *structpb.Value

	var var_45f199d87213_err error
	var_45f199d87213_mapped, var_45f199d87213_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_45f199d87213)
	if var_45f199d87213_err != nil {
		panic(var_45f199d87213_err)
	}
	properties["username"] = var_45f199d87213_mapped

	var_e7356a2f5831 := user.Password

	if var_e7356a2f5831 != nil {
		var var_e7356a2f5831_mapped *structpb.Value

		var var_e7356a2f5831_err error
		var_e7356a2f5831_mapped, var_e7356a2f5831_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_e7356a2f5831)
		if var_e7356a2f5831_err != nil {
			panic(var_e7356a2f5831_err)
		}
		properties["password"] = var_e7356a2f5831_mapped
	}

	var_ba1acf21bbff := user.Roles

	if var_ba1acf21bbff != nil {
		var var_ba1acf21bbff_mapped *structpb.Value

		var var_ba1acf21bbff_l []*structpb.Value
		for _, value := range var_ba1acf21bbff {

			var_77dc2a4cf60d := value
			var var_77dc2a4cf60d_mapped *structpb.Value

			var_77dc2a4cf60d_mapped = structpb.NewStructValue(&structpb.Struct{Fields: RoleMapperInstance.ToProperties(var_77dc2a4cf60d)})

			var_ba1acf21bbff_l = append(var_ba1acf21bbff_l, var_77dc2a4cf60d_mapped)
		}
		var_ba1acf21bbff_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_ba1acf21bbff_l})
		properties["roles"] = var_ba1acf21bbff_mapped
	}

	var_d8da9ef09bd9 := user.SecurityConstraints

	if var_d8da9ef09bd9 != nil {
		var var_d8da9ef09bd9_mapped *structpb.Value

		var var_d8da9ef09bd9_l []*structpb.Value
		for _, value := range var_d8da9ef09bd9 {

			var_8669d46b8d3a := value
			var var_8669d46b8d3a_mapped *structpb.Value

			var_8669d46b8d3a_mapped = structpb.NewStructValue(&structpb.Struct{Fields: SecurityConstraintMapperInstance.ToProperties(var_8669d46b8d3a)})

			var_d8da9ef09bd9_l = append(var_d8da9ef09bd9_l, var_8669d46b8d3a_mapped)
		}
		var_d8da9ef09bd9_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_d8da9ef09bd9_l})
		properties["securityConstraints"] = var_d8da9ef09bd9_mapped
	}

	var_17a00d869537 := user.Details

	if var_17a00d869537 != nil {
		var var_17a00d869537_mapped *structpb.Value

		var var_17a00d869537_err error
		var_17a00d869537_mapped, var_17a00d869537_err = types.ByResourcePropertyType(model.ResourceProperty_OBJECT).Pack(*var_17a00d869537)
		if var_17a00d869537_err != nil {
			panic(var_17a00d869537_err)
		}
		properties["details"] = var_17a00d869537_mapped
	}
	return properties
}

func (m *UserMapper) FromProperties(properties map[string]*structpb.Value) *User {
	var s = m.New()
	if properties["id"] != nil {

		var_2fac28a5c940 := properties["id"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_UUID).UnPack(var_2fac28a5c940)

		if err != nil {
			panic(err)
		}

		var_2fac28a5c940_mapped := new(uuid.UUID)
		*var_2fac28a5c940_mapped = val.(uuid.UUID)

		s.Id = var_2fac28a5c940_mapped
	}
	if properties["version"] != nil {

		var_c3396f483320 := properties["version"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(var_c3396f483320)

		if err != nil {
			panic(err)
		}

		var_c3396f483320_mapped := val.(int32)

		s.Version = var_c3396f483320_mapped
	}
	if properties["createdBy"] != nil {

		var_e7792ccb66f4 := properties["createdBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_e7792ccb66f4)

		if err != nil {
			panic(err)
		}

		var_e7792ccb66f4_mapped := new(string)
		*var_e7792ccb66f4_mapped = val.(string)

		s.CreatedBy = var_e7792ccb66f4_mapped
	}
	if properties["updatedBy"] != nil {

		var_f1e68edbcc36 := properties["updatedBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_f1e68edbcc36)

		if err != nil {
			panic(err)
		}

		var_f1e68edbcc36_mapped := new(string)
		*var_f1e68edbcc36_mapped = val.(string)

		s.UpdatedBy = var_f1e68edbcc36_mapped
	}
	if properties["createdOn"] != nil {

		var_b261cd831ede := properties["createdOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_b261cd831ede)

		if err != nil {
			panic(err)
		}

		var_b261cd831ede_mapped := new(time.Time)
		*var_b261cd831ede_mapped = val.(time.Time)

		s.CreatedOn = var_b261cd831ede_mapped
	}
	if properties["updatedOn"] != nil {

		var_c52db6de2b22 := properties["updatedOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_c52db6de2b22)

		if err != nil {
			panic(err)
		}

		var_c52db6de2b22_mapped := new(time.Time)
		*var_c52db6de2b22_mapped = val.(time.Time)

		s.UpdatedOn = var_c52db6de2b22_mapped
	}
	if properties["username"] != nil {

		var_d622ff071905 := properties["username"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_d622ff071905)

		if err != nil {
			panic(err)
		}

		var_d622ff071905_mapped := val.(string)

		s.Username = var_d622ff071905_mapped
	}
	if properties["password"] != nil {

		var_d3a3d4e88349 := properties["password"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_d3a3d4e88349)

		if err != nil {
			panic(err)
		}

		var_d3a3d4e88349_mapped := new(string)
		*var_d3a3d4e88349_mapped = val.(string)

		s.Password = var_d3a3d4e88349_mapped
	}
	if properties["roles"] != nil {

		var_f30c053fdb72 := properties["roles"]
		var_f30c053fdb72_mapped := []*Role{}
		for _, v := range var_f30c053fdb72.GetListValue().Values {

			var_bbcbcfe866c7 := v
			var_bbcbcfe866c7_mapped := RoleMapperInstance.FromProperties(var_bbcbcfe866c7.GetStructValue().Fields)

			var_f30c053fdb72_mapped = append(var_f30c053fdb72_mapped, var_bbcbcfe866c7_mapped)
		}

		s.Roles = var_f30c053fdb72_mapped
	}
	if properties["securityConstraints"] != nil {

		var_ee7a7b9d3c11 := properties["securityConstraints"]
		var_ee7a7b9d3c11_mapped := []*SecurityConstraint{}
		for _, v := range var_ee7a7b9d3c11.GetListValue().Values {

			var_a58c922df53b := v
			var_a58c922df53b_mapped := SecurityConstraintMapperInstance.FromProperties(var_a58c922df53b.GetStructValue().Fields)

			var_ee7a7b9d3c11_mapped = append(var_ee7a7b9d3c11_mapped, var_a58c922df53b_mapped)
		}

		s.SecurityConstraints = var_ee7a7b9d3c11_mapped
	}
	if properties["details"] != nil {

		var_51c244b4383a := properties["details"]
		var_51c244b4383a_mapped := new(unstructured.Unstructured)
		*var_51c244b4383a_mapped = unstructured.FromStructValue(var_51c244b4383a.GetStructValue())

		s.Details = var_51c244b4383a_mapped
	}
	return s
}
