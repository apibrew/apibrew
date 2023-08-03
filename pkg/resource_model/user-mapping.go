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

	var_c6b17bcd4bca := user.Id

	if var_c6b17bcd4bca != nil {
		var var_c6b17bcd4bca_mapped *structpb.Value

		var var_c6b17bcd4bca_err error
		var_c6b17bcd4bca_mapped, var_c6b17bcd4bca_err = types.ByResourcePropertyType(model.ResourceProperty_UUID).Pack(*var_c6b17bcd4bca)
		if var_c6b17bcd4bca_err != nil {
			panic(var_c6b17bcd4bca_err)
		}
		properties["id"] = var_c6b17bcd4bca_mapped
	}

	var_a59386067ad6 := user.Version

	var var_a59386067ad6_mapped *structpb.Value

	var var_a59386067ad6_err error
	var_a59386067ad6_mapped, var_a59386067ad6_err = types.ByResourcePropertyType(model.ResourceProperty_INT32).Pack(var_a59386067ad6)
	if var_a59386067ad6_err != nil {
		panic(var_a59386067ad6_err)
	}
	properties["version"] = var_a59386067ad6_mapped

	var_64cf6a0a7d62 := user.CreatedBy

	if var_64cf6a0a7d62 != nil {
		var var_64cf6a0a7d62_mapped *structpb.Value

		var var_64cf6a0a7d62_err error
		var_64cf6a0a7d62_mapped, var_64cf6a0a7d62_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_64cf6a0a7d62)
		if var_64cf6a0a7d62_err != nil {
			panic(var_64cf6a0a7d62_err)
		}
		properties["createdBy"] = var_64cf6a0a7d62_mapped
	}

	var_2098e4c18ebb := user.UpdatedBy

	if var_2098e4c18ebb != nil {
		var var_2098e4c18ebb_mapped *structpb.Value

		var var_2098e4c18ebb_err error
		var_2098e4c18ebb_mapped, var_2098e4c18ebb_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_2098e4c18ebb)
		if var_2098e4c18ebb_err != nil {
			panic(var_2098e4c18ebb_err)
		}
		properties["updatedBy"] = var_2098e4c18ebb_mapped
	}

	var_8e4fba6ce01d := user.CreatedOn

	if var_8e4fba6ce01d != nil {
		var var_8e4fba6ce01d_mapped *structpb.Value

		var var_8e4fba6ce01d_err error
		var_8e4fba6ce01d_mapped, var_8e4fba6ce01d_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_8e4fba6ce01d)
		if var_8e4fba6ce01d_err != nil {
			panic(var_8e4fba6ce01d_err)
		}
		properties["createdOn"] = var_8e4fba6ce01d_mapped
	}

	var_1ee5de76bf55 := user.UpdatedOn

	if var_1ee5de76bf55 != nil {
		var var_1ee5de76bf55_mapped *structpb.Value

		var var_1ee5de76bf55_err error
		var_1ee5de76bf55_mapped, var_1ee5de76bf55_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_1ee5de76bf55)
		if var_1ee5de76bf55_err != nil {
			panic(var_1ee5de76bf55_err)
		}
		properties["updatedOn"] = var_1ee5de76bf55_mapped
	}

	var_d6d1eb9972bf := user.Username

	var var_d6d1eb9972bf_mapped *structpb.Value

	var var_d6d1eb9972bf_err error
	var_d6d1eb9972bf_mapped, var_d6d1eb9972bf_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_d6d1eb9972bf)
	if var_d6d1eb9972bf_err != nil {
		panic(var_d6d1eb9972bf_err)
	}
	properties["username"] = var_d6d1eb9972bf_mapped

	var_54de85345ded := user.Password

	if var_54de85345ded != nil {
		var var_54de85345ded_mapped *structpb.Value

		var var_54de85345ded_err error
		var_54de85345ded_mapped, var_54de85345ded_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_54de85345ded)
		if var_54de85345ded_err != nil {
			panic(var_54de85345ded_err)
		}
		properties["password"] = var_54de85345ded_mapped
	}

	var_6be681a2eb45 := user.Roles

	if var_6be681a2eb45 != nil {
		var var_6be681a2eb45_mapped *structpb.Value

		var var_6be681a2eb45_l []*structpb.Value
		for _, value := range var_6be681a2eb45 {

			var_4b26ad20f651 := value
			var var_4b26ad20f651_mapped *structpb.Value

			var_4b26ad20f651_mapped = structpb.NewStructValue(&structpb.Struct{Fields: RoleMapperInstance.ToProperties(var_4b26ad20f651)})

			var_6be681a2eb45_l = append(var_6be681a2eb45_l, var_4b26ad20f651_mapped)
		}
		var_6be681a2eb45_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_6be681a2eb45_l})
		properties["roles"] = var_6be681a2eb45_mapped
	}

	var_ceef7aa24ed3 := user.SecurityConstraints

	if var_ceef7aa24ed3 != nil {
		var var_ceef7aa24ed3_mapped *structpb.Value

		var var_ceef7aa24ed3_l []*structpb.Value
		for _, value := range var_ceef7aa24ed3 {

			var_aa8d59c75818 := value
			var var_aa8d59c75818_mapped *structpb.Value

			var_aa8d59c75818_mapped = structpb.NewStructValue(&structpb.Struct{Fields: SecurityConstraintMapperInstance.ToProperties(var_aa8d59c75818)})

			var_ceef7aa24ed3_l = append(var_ceef7aa24ed3_l, var_aa8d59c75818_mapped)
		}
		var_ceef7aa24ed3_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_ceef7aa24ed3_l})
		properties["securityConstraints"] = var_ceef7aa24ed3_mapped
	}

	var_46f4be046f8f := user.Details

	if var_46f4be046f8f != nil {
		var var_46f4be046f8f_mapped *structpb.Value

		var var_46f4be046f8f_err error
		var_46f4be046f8f_mapped, var_46f4be046f8f_err = types.ByResourcePropertyType(model.ResourceProperty_OBJECT).Pack(*var_46f4be046f8f)
		if var_46f4be046f8f_err != nil {
			panic(var_46f4be046f8f_err)
		}
		properties["details"] = var_46f4be046f8f_mapped
	}
	return properties
}

func (m *UserMapper) FromProperties(properties map[string]*structpb.Value) *User {
	var s = m.New()
	if properties["id"] != nil {

		var_fc1b3f4e6130 := properties["id"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_UUID).UnPack(var_fc1b3f4e6130)

		if err != nil {
			panic(err)
		}

		var_fc1b3f4e6130_mapped := new(uuid.UUID)
		*var_fc1b3f4e6130_mapped = val.(uuid.UUID)

		s.Id = var_fc1b3f4e6130_mapped
	}
	if properties["version"] != nil {

		var_d41a0136cd73 := properties["version"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(var_d41a0136cd73)

		if err != nil {
			panic(err)
		}

		var_d41a0136cd73_mapped := val.(int32)

		s.Version = var_d41a0136cd73_mapped
	}
	if properties["createdBy"] != nil {

		var_a58be79e4115 := properties["createdBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_a58be79e4115)

		if err != nil {
			panic(err)
		}

		var_a58be79e4115_mapped := new(string)
		*var_a58be79e4115_mapped = val.(string)

		s.CreatedBy = var_a58be79e4115_mapped
	}
	if properties["updatedBy"] != nil {

		var_04bc9836039c := properties["updatedBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_04bc9836039c)

		if err != nil {
			panic(err)
		}

		var_04bc9836039c_mapped := new(string)
		*var_04bc9836039c_mapped = val.(string)

		s.UpdatedBy = var_04bc9836039c_mapped
	}
	if properties["createdOn"] != nil {

		var_b1f016872349 := properties["createdOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_b1f016872349)

		if err != nil {
			panic(err)
		}

		var_b1f016872349_mapped := new(time.Time)
		*var_b1f016872349_mapped = val.(time.Time)

		s.CreatedOn = var_b1f016872349_mapped
	}
	if properties["updatedOn"] != nil {

		var_d0da9ea85c39 := properties["updatedOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_d0da9ea85c39)

		if err != nil {
			panic(err)
		}

		var_d0da9ea85c39_mapped := new(time.Time)
		*var_d0da9ea85c39_mapped = val.(time.Time)

		s.UpdatedOn = var_d0da9ea85c39_mapped
	}
	if properties["username"] != nil {

		var_960b0a0bc6df := properties["username"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_960b0a0bc6df)

		if err != nil {
			panic(err)
		}

		var_960b0a0bc6df_mapped := val.(string)

		s.Username = var_960b0a0bc6df_mapped
	}
	if properties["password"] != nil {

		var_1671130d09b9 := properties["password"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_1671130d09b9)

		if err != nil {
			panic(err)
		}

		var_1671130d09b9_mapped := new(string)
		*var_1671130d09b9_mapped = val.(string)

		s.Password = var_1671130d09b9_mapped
	}
	if properties["roles"] != nil {

		var_f195f021d91b := properties["roles"]
		var_f195f021d91b_mapped := []*Role{}
		for _, v := range var_f195f021d91b.GetListValue().Values {

			var_555e011db95f := v
			var_555e011db95f_mapped := RoleMapperInstance.FromProperties(var_555e011db95f.GetStructValue().Fields)

			var_f195f021d91b_mapped = append(var_f195f021d91b_mapped, var_555e011db95f_mapped)
		}

		s.Roles = var_f195f021d91b_mapped
	}
	if properties["securityConstraints"] != nil {

		var_66438042cad9 := properties["securityConstraints"]
		var_66438042cad9_mapped := []*SecurityConstraint{}
		for _, v := range var_66438042cad9.GetListValue().Values {

			var_bb176fd61092 := v
			var_bb176fd61092_mapped := SecurityConstraintMapperInstance.FromProperties(var_bb176fd61092.GetStructValue().Fields)

			var_66438042cad9_mapped = append(var_66438042cad9_mapped, var_bb176fd61092_mapped)
		}

		s.SecurityConstraints = var_66438042cad9_mapped
	}
	if properties["details"] != nil {

		var_f63234880147 := properties["details"]
		var_f63234880147_mapped := new(unstructured.Unstructured)
		*var_f63234880147_mapped = unstructured.FromStructValue(var_f63234880147.GetStructValue())

		s.Details = var_f63234880147_mapped
	}
	return s
}
