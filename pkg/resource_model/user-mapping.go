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

	var_6436c40c91d1 := user.Id

	if var_6436c40c91d1 != nil {
		var var_6436c40c91d1_mapped *structpb.Value

		var var_6436c40c91d1_err error
		var_6436c40c91d1_mapped, var_6436c40c91d1_err = types.ByResourcePropertyType(model.ResourceProperty_UUID).Pack(*var_6436c40c91d1)
		if var_6436c40c91d1_err != nil {
			panic(var_6436c40c91d1_err)
		}
		properties["id"] = var_6436c40c91d1_mapped
	}

	var_87ab06213d26 := user.Version

	var var_87ab06213d26_mapped *structpb.Value

	var var_87ab06213d26_err error
	var_87ab06213d26_mapped, var_87ab06213d26_err = types.ByResourcePropertyType(model.ResourceProperty_INT32).Pack(var_87ab06213d26)
	if var_87ab06213d26_err != nil {
		panic(var_87ab06213d26_err)
	}
	properties["version"] = var_87ab06213d26_mapped

	var_aa1eaebad62e := user.CreatedBy

	if var_aa1eaebad62e != nil {
		var var_aa1eaebad62e_mapped *structpb.Value

		var var_aa1eaebad62e_err error
		var_aa1eaebad62e_mapped, var_aa1eaebad62e_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_aa1eaebad62e)
		if var_aa1eaebad62e_err != nil {
			panic(var_aa1eaebad62e_err)
		}
		properties["createdBy"] = var_aa1eaebad62e_mapped
	}

	var_6861b2aa251b := user.UpdatedBy

	if var_6861b2aa251b != nil {
		var var_6861b2aa251b_mapped *structpb.Value

		var var_6861b2aa251b_err error
		var_6861b2aa251b_mapped, var_6861b2aa251b_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_6861b2aa251b)
		if var_6861b2aa251b_err != nil {
			panic(var_6861b2aa251b_err)
		}
		properties["updatedBy"] = var_6861b2aa251b_mapped
	}

	var_8da02e001f30 := user.CreatedOn

	if var_8da02e001f30 != nil {
		var var_8da02e001f30_mapped *structpb.Value

		var var_8da02e001f30_err error
		var_8da02e001f30_mapped, var_8da02e001f30_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_8da02e001f30)
		if var_8da02e001f30_err != nil {
			panic(var_8da02e001f30_err)
		}
		properties["createdOn"] = var_8da02e001f30_mapped
	}

	var_bd32eed358a2 := user.UpdatedOn

	if var_bd32eed358a2 != nil {
		var var_bd32eed358a2_mapped *structpb.Value

		var var_bd32eed358a2_err error
		var_bd32eed358a2_mapped, var_bd32eed358a2_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_bd32eed358a2)
		if var_bd32eed358a2_err != nil {
			panic(var_bd32eed358a2_err)
		}
		properties["updatedOn"] = var_bd32eed358a2_mapped
	}

	var_b51b99fdf980 := user.Username

	var var_b51b99fdf980_mapped *structpb.Value

	var var_b51b99fdf980_err error
	var_b51b99fdf980_mapped, var_b51b99fdf980_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_b51b99fdf980)
	if var_b51b99fdf980_err != nil {
		panic(var_b51b99fdf980_err)
	}
	properties["username"] = var_b51b99fdf980_mapped

	var_511bf5cda485 := user.Password

	if var_511bf5cda485 != nil {
		var var_511bf5cda485_mapped *structpb.Value

		var var_511bf5cda485_err error
		var_511bf5cda485_mapped, var_511bf5cda485_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_511bf5cda485)
		if var_511bf5cda485_err != nil {
			panic(var_511bf5cda485_err)
		}
		properties["password"] = var_511bf5cda485_mapped
	}

	var_8827048ab677 := user.Roles

	if var_8827048ab677 != nil {
		var var_8827048ab677_mapped *structpb.Value

		var var_8827048ab677_l []*structpb.Value
		for _, value := range var_8827048ab677 {

			var_0541015068db := value
			var var_0541015068db_mapped *structpb.Value

			var_0541015068db_mapped = structpb.NewStructValue(&structpb.Struct{Fields: RoleMapperInstance.ToProperties(var_0541015068db)})

			var_8827048ab677_l = append(var_8827048ab677_l, var_0541015068db_mapped)
		}
		var_8827048ab677_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_8827048ab677_l})
		properties["roles"] = var_8827048ab677_mapped
	}

	var_65d47ec432ee := user.SecurityConstraints

	if var_65d47ec432ee != nil {
		var var_65d47ec432ee_mapped *structpb.Value

		var var_65d47ec432ee_l []*structpb.Value
		for _, value := range var_65d47ec432ee {

			var_687e0017149a := value
			var var_687e0017149a_mapped *structpb.Value

			var_687e0017149a_mapped = structpb.NewStructValue(&structpb.Struct{Fields: SecurityConstraintMapperInstance.ToProperties(var_687e0017149a)})

			var_65d47ec432ee_l = append(var_65d47ec432ee_l, var_687e0017149a_mapped)
		}
		var_65d47ec432ee_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_65d47ec432ee_l})
		properties["securityConstraints"] = var_65d47ec432ee_mapped
	}

	var_5af94e0247d4 := user.Details

	if var_5af94e0247d4 != nil {
		var var_5af94e0247d4_mapped *structpb.Value

		var var_5af94e0247d4_err error
		var_5af94e0247d4_mapped, var_5af94e0247d4_err = types.ByResourcePropertyType(model.ResourceProperty_OBJECT).Pack(*var_5af94e0247d4)
		if var_5af94e0247d4_err != nil {
			panic(var_5af94e0247d4_err)
		}
		properties["details"] = var_5af94e0247d4_mapped
	}
	return properties
}

func (m *UserMapper) FromProperties(properties map[string]*structpb.Value) *User {
	var s = m.New()
	if properties["id"] != nil {

		var_a67e98add17f := properties["id"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_UUID).UnPack(var_a67e98add17f)

		if err != nil {
			panic(err)
		}

		var_a67e98add17f_mapped := new(uuid.UUID)
		*var_a67e98add17f_mapped = val.(uuid.UUID)

		s.Id = var_a67e98add17f_mapped
	}
	if properties["version"] != nil {

		var_f04012192a34 := properties["version"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(var_f04012192a34)

		if err != nil {
			panic(err)
		}

		var_f04012192a34_mapped := val.(int32)

		s.Version = var_f04012192a34_mapped
	}
	if properties["createdBy"] != nil {

		var_dffe72792590 := properties["createdBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_dffe72792590)

		if err != nil {
			panic(err)
		}

		var_dffe72792590_mapped := new(string)
		*var_dffe72792590_mapped = val.(string)

		s.CreatedBy = var_dffe72792590_mapped
	}
	if properties["updatedBy"] != nil {

		var_2254741cb6f5 := properties["updatedBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_2254741cb6f5)

		if err != nil {
			panic(err)
		}

		var_2254741cb6f5_mapped := new(string)
		*var_2254741cb6f5_mapped = val.(string)

		s.UpdatedBy = var_2254741cb6f5_mapped
	}
	if properties["createdOn"] != nil {

		var_0af7fc8cc459 := properties["createdOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_0af7fc8cc459)

		if err != nil {
			panic(err)
		}

		var_0af7fc8cc459_mapped := new(time.Time)
		*var_0af7fc8cc459_mapped = val.(time.Time)

		s.CreatedOn = var_0af7fc8cc459_mapped
	}
	if properties["updatedOn"] != nil {

		var_a1c8dcb42e3b := properties["updatedOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_a1c8dcb42e3b)

		if err != nil {
			panic(err)
		}

		var_a1c8dcb42e3b_mapped := new(time.Time)
		*var_a1c8dcb42e3b_mapped = val.(time.Time)

		s.UpdatedOn = var_a1c8dcb42e3b_mapped
	}
	if properties["username"] != nil {

		var_a8b367fb7982 := properties["username"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_a8b367fb7982)

		if err != nil {
			panic(err)
		}

		var_a8b367fb7982_mapped := val.(string)

		s.Username = var_a8b367fb7982_mapped
	}
	if properties["password"] != nil {

		var_0acf8bd929a3 := properties["password"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_0acf8bd929a3)

		if err != nil {
			panic(err)
		}

		var_0acf8bd929a3_mapped := new(string)
		*var_0acf8bd929a3_mapped = val.(string)

		s.Password = var_0acf8bd929a3_mapped
	}
	if properties["roles"] != nil {

		var_fa97a5a5113d := properties["roles"]
		var_fa97a5a5113d_mapped := []*Role{}
		for _, v := range var_fa97a5a5113d.GetListValue().Values {

			var_b8b3fee45ea5 := v
			var_b8b3fee45ea5_mapped := RoleMapperInstance.FromProperties(var_b8b3fee45ea5.GetStructValue().Fields)

			var_fa97a5a5113d_mapped = append(var_fa97a5a5113d_mapped, var_b8b3fee45ea5_mapped)
		}

		s.Roles = var_fa97a5a5113d_mapped
	}
	if properties["securityConstraints"] != nil {

		var_2d093e73343e := properties["securityConstraints"]
		var_2d093e73343e_mapped := []*SecurityConstraint{}
		for _, v := range var_2d093e73343e.GetListValue().Values {

			var_79628a872476 := v
			var_79628a872476_mapped := SecurityConstraintMapperInstance.FromProperties(var_79628a872476.GetStructValue().Fields)

			var_2d093e73343e_mapped = append(var_2d093e73343e_mapped, var_79628a872476_mapped)
		}

		s.SecurityConstraints = var_2d093e73343e_mapped
	}
	if properties["details"] != nil {

		var_c3f129e9db4a := properties["details"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_OBJECT).UnPack(var_c3f129e9db4a)

		if err != nil {
			panic(err)
		}

		var_c3f129e9db4a_mapped := new(unstructured.Unstructured)
		*var_c3f129e9db4a_mapped = val.(unstructured.Unstructured)

		s.Details = var_c3f129e9db4a_mapped
	}
	return s
}
