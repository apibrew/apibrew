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

	var_87b12f05d8fd := user.Id

	if var_87b12f05d8fd != nil {
		var var_87b12f05d8fd_mapped *structpb.Value

		var var_87b12f05d8fd_err error
		var_87b12f05d8fd_mapped, var_87b12f05d8fd_err = types.ByResourcePropertyType(model.ResourceProperty_UUID).Pack(*var_87b12f05d8fd)
		if var_87b12f05d8fd_err != nil {
			panic(var_87b12f05d8fd_err)
		}
		properties["id"] = var_87b12f05d8fd_mapped
	}

	var_c769c466f207 := user.Version

	var var_c769c466f207_mapped *structpb.Value

	var var_c769c466f207_err error
	var_c769c466f207_mapped, var_c769c466f207_err = types.ByResourcePropertyType(model.ResourceProperty_INT32).Pack(var_c769c466f207)
	if var_c769c466f207_err != nil {
		panic(var_c769c466f207_err)
	}
	properties["version"] = var_c769c466f207_mapped

	var_61c979840230 := user.CreatedBy

	if var_61c979840230 != nil {
		var var_61c979840230_mapped *structpb.Value

		var var_61c979840230_err error
		var_61c979840230_mapped, var_61c979840230_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_61c979840230)
		if var_61c979840230_err != nil {
			panic(var_61c979840230_err)
		}
		properties["createdBy"] = var_61c979840230_mapped
	}

	var_d077f25bd5b6 := user.UpdatedBy

	if var_d077f25bd5b6 != nil {
		var var_d077f25bd5b6_mapped *structpb.Value

		var var_d077f25bd5b6_err error
		var_d077f25bd5b6_mapped, var_d077f25bd5b6_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_d077f25bd5b6)
		if var_d077f25bd5b6_err != nil {
			panic(var_d077f25bd5b6_err)
		}
		properties["updatedBy"] = var_d077f25bd5b6_mapped
	}

	var_be74a8220cf4 := user.CreatedOn

	if var_be74a8220cf4 != nil {
		var var_be74a8220cf4_mapped *structpb.Value

		var var_be74a8220cf4_err error
		var_be74a8220cf4_mapped, var_be74a8220cf4_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_be74a8220cf4)
		if var_be74a8220cf4_err != nil {
			panic(var_be74a8220cf4_err)
		}
		properties["createdOn"] = var_be74a8220cf4_mapped
	}

	var_04f82b2c7267 := user.UpdatedOn

	if var_04f82b2c7267 != nil {
		var var_04f82b2c7267_mapped *structpb.Value

		var var_04f82b2c7267_err error
		var_04f82b2c7267_mapped, var_04f82b2c7267_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_04f82b2c7267)
		if var_04f82b2c7267_err != nil {
			panic(var_04f82b2c7267_err)
		}
		properties["updatedOn"] = var_04f82b2c7267_mapped
	}

	var_bff1b97c5419 := user.Username

	var var_bff1b97c5419_mapped *structpb.Value

	var var_bff1b97c5419_err error
	var_bff1b97c5419_mapped, var_bff1b97c5419_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_bff1b97c5419)
	if var_bff1b97c5419_err != nil {
		panic(var_bff1b97c5419_err)
	}
	properties["username"] = var_bff1b97c5419_mapped

	var_f7c7e9895f5e := user.Password

	if var_f7c7e9895f5e != nil {
		var var_f7c7e9895f5e_mapped *structpb.Value

		var var_f7c7e9895f5e_err error
		var_f7c7e9895f5e_mapped, var_f7c7e9895f5e_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_f7c7e9895f5e)
		if var_f7c7e9895f5e_err != nil {
			panic(var_f7c7e9895f5e_err)
		}
		properties["password"] = var_f7c7e9895f5e_mapped
	}

	var_6479b738b402 := user.Roles

	if var_6479b738b402 != nil {
		var var_6479b738b402_mapped *structpb.Value

		var var_6479b738b402_l []*structpb.Value
		for _, value := range var_6479b738b402 {

			var_6c248eda937e := value
			var var_6c248eda937e_mapped *structpb.Value

			var_6c248eda937e_mapped = structpb.NewStructValue(&structpb.Struct{Fields: RoleMapperInstance.ToProperties(var_6c248eda937e)})

			var_6479b738b402_l = append(var_6479b738b402_l, var_6c248eda937e_mapped)
		}
		var_6479b738b402_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_6479b738b402_l})
		properties["roles"] = var_6479b738b402_mapped
	}

	var_8cd2617317c5 := user.SecurityConstraints

	if var_8cd2617317c5 != nil {
		var var_8cd2617317c5_mapped *structpb.Value

		var var_8cd2617317c5_l []*structpb.Value
		for _, value := range var_8cd2617317c5 {

			var_87c404e9c6d6 := value
			var var_87c404e9c6d6_mapped *structpb.Value

			var_87c404e9c6d6_mapped = structpb.NewStructValue(&structpb.Struct{Fields: SecurityConstraintMapperInstance.ToProperties(var_87c404e9c6d6)})

			var_8cd2617317c5_l = append(var_8cd2617317c5_l, var_87c404e9c6d6_mapped)
		}
		var_8cd2617317c5_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_8cd2617317c5_l})
		properties["securityConstraints"] = var_8cd2617317c5_mapped
	}

	var_56b69de992da := user.Details

	if var_56b69de992da != nil {
		var var_56b69de992da_mapped *structpb.Value

		var var_56b69de992da_err error
		var_56b69de992da_mapped, var_56b69de992da_err = types.ByResourcePropertyType(model.ResourceProperty_OBJECT).Pack(*var_56b69de992da)
		if var_56b69de992da_err != nil {
			panic(var_56b69de992da_err)
		}
		properties["details"] = var_56b69de992da_mapped
	}
	return properties
}

func (m *UserMapper) FromProperties(properties map[string]*structpb.Value) *User {
	var s = m.New()
	if properties["id"] != nil {

		var_74a86d1789e7 := properties["id"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_UUID).UnPack(var_74a86d1789e7)

		if err != nil {
			panic(err)
		}

		var_74a86d1789e7_mapped := new(uuid.UUID)
		*var_74a86d1789e7_mapped = val.(uuid.UUID)

		s.Id = var_74a86d1789e7_mapped
	}
	if properties["version"] != nil {

		var_42c4bbb7bae0 := properties["version"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(var_42c4bbb7bae0)

		if err != nil {
			panic(err)
		}

		var_42c4bbb7bae0_mapped := val.(int32)

		s.Version = var_42c4bbb7bae0_mapped
	}
	if properties["createdBy"] != nil {

		var_c72555fcb9e4 := properties["createdBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_c72555fcb9e4)

		if err != nil {
			panic(err)
		}

		var_c72555fcb9e4_mapped := new(string)
		*var_c72555fcb9e4_mapped = val.(string)

		s.CreatedBy = var_c72555fcb9e4_mapped
	}
	if properties["updatedBy"] != nil {

		var_299b8a91de05 := properties["updatedBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_299b8a91de05)

		if err != nil {
			panic(err)
		}

		var_299b8a91de05_mapped := new(string)
		*var_299b8a91de05_mapped = val.(string)

		s.UpdatedBy = var_299b8a91de05_mapped
	}
	if properties["createdOn"] != nil {

		var_53a38a6a49ad := properties["createdOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_53a38a6a49ad)

		if err != nil {
			panic(err)
		}

		var_53a38a6a49ad_mapped := new(time.Time)
		*var_53a38a6a49ad_mapped = val.(time.Time)

		s.CreatedOn = var_53a38a6a49ad_mapped
	}
	if properties["updatedOn"] != nil {

		var_333d7b22e323 := properties["updatedOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_333d7b22e323)

		if err != nil {
			panic(err)
		}

		var_333d7b22e323_mapped := new(time.Time)
		*var_333d7b22e323_mapped = val.(time.Time)

		s.UpdatedOn = var_333d7b22e323_mapped
	}
	if properties["username"] != nil {

		var_5697aa36f836 := properties["username"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_5697aa36f836)

		if err != nil {
			panic(err)
		}

		var_5697aa36f836_mapped := val.(string)

		s.Username = var_5697aa36f836_mapped
	}
	if properties["password"] != nil {

		var_890104f182bb := properties["password"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_890104f182bb)

		if err != nil {
			panic(err)
		}

		var_890104f182bb_mapped := new(string)
		*var_890104f182bb_mapped = val.(string)

		s.Password = var_890104f182bb_mapped
	}
	if properties["roles"] != nil {

		var_6761d9020db6 := properties["roles"]
		var_6761d9020db6_mapped := []*Role{}
		for _, v := range var_6761d9020db6.GetListValue().Values {

			var_0ed719dbc013 := v
			var_0ed719dbc013_mapped := RoleMapperInstance.FromProperties(var_0ed719dbc013.GetStructValue().Fields)

			var_6761d9020db6_mapped = append(var_6761d9020db6_mapped, var_0ed719dbc013_mapped)
		}

		s.Roles = var_6761d9020db6_mapped
	}
	if properties["securityConstraints"] != nil {

		var_437929b05a26 := properties["securityConstraints"]
		var_437929b05a26_mapped := []*SecurityConstraint{}
		for _, v := range var_437929b05a26.GetListValue().Values {

			var_514bd5a788bb := v
			var_514bd5a788bb_mapped := SecurityConstraintMapperInstance.FromProperties(var_514bd5a788bb.GetStructValue().Fields)

			var_437929b05a26_mapped = append(var_437929b05a26_mapped, var_514bd5a788bb_mapped)
		}

		s.SecurityConstraints = var_437929b05a26_mapped
	}
	if properties["details"] != nil {

		var_9bc4fab9829f := properties["details"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_OBJECT).UnPack(var_9bc4fab9829f)

		if err != nil {
			panic(err)
		}

		var_9bc4fab9829f_mapped := new(unstructured.Unstructured)
		*var_9bc4fab9829f_mapped = val.(unstructured.Unstructured)

		s.Details = var_9bc4fab9829f_mapped
	}
	return s
}
