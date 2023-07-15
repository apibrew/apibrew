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
	return rec
}

func (m *UserMapper) FromRecord(record *model.Record) *User {
	return m.FromProperties(record.Properties)
}

func (m *UserMapper) ToProperties(user *User) map[string]*structpb.Value {
	var properties = make(map[string]*structpb.Value)

	if user.Id != nil {
		id, err := types.ByResourcePropertyType(model.ResourceProperty_UUID).Pack(*user.Id)
		if err != nil {
			panic(err)
		}
		properties["id"] = id
	}

	version, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).Pack(user.Version)
	if err != nil {
		panic(err)
	}
	properties["version"] = version

	if user.CreatedBy != nil {
		createdBy, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*user.CreatedBy)
		if err != nil {
			panic(err)
		}
		properties["createdBy"] = createdBy
	}

	if user.UpdatedBy != nil {
		updatedBy, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*user.UpdatedBy)
		if err != nil {
			panic(err)
		}
		properties["updatedBy"] = updatedBy
	}

	if user.CreatedOn != nil {
		createdOn, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*user.CreatedOn)
		if err != nil {
			panic(err)
		}
		properties["createdOn"] = createdOn
	}

	if user.UpdatedOn != nil {
		updatedOn, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*user.UpdatedOn)
		if err != nil {
			panic(err)
		}
		properties["updatedOn"] = updatedOn
	}

	username, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(user.Username)
	if err != nil {
		panic(err)
	}
	properties["username"] = username

	if user.Password != nil {
		password, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*user.Password)
		if err != nil {
			panic(err)
		}
		properties["password"] = password
	}

	if user.Roles != nil {
	}

	if user.SecurityConstraints != nil {
	}

	if user.Details != nil {
	}

	return properties
}

func (m *UserMapper) FromProperties(properties map[string]*structpb.Value) *User {
	var s = m.New()
	if properties["id"] != nil {

		var_ff01a276ce1a := properties["id"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_UUID).UnPack(var_ff01a276ce1a)

		if err != nil {
			panic(err)
		}

		var_ff01a276ce1a_mapped := new(uuid.UUID)
		*var_ff01a276ce1a_mapped = val.(uuid.UUID)

		s.Id = var_ff01a276ce1a_mapped
	}
	if properties["version"] != nil {

		var_50ea22a28e57 := properties["version"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(var_50ea22a28e57)

		if err != nil {
			panic(err)
		}

		var_50ea22a28e57_mapped := val.(int32)

		s.Version = var_50ea22a28e57_mapped
	}
	if properties["createdBy"] != nil {

		var_c6808ae28b1c := properties["createdBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_c6808ae28b1c)

		if err != nil {
			panic(err)
		}

		var_c6808ae28b1c_mapped := new(string)
		*var_c6808ae28b1c_mapped = val.(string)

		s.CreatedBy = var_c6808ae28b1c_mapped
	}
	if properties["updatedBy"] != nil {

		var_fe0361dd392a := properties["updatedBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_fe0361dd392a)

		if err != nil {
			panic(err)
		}

		var_fe0361dd392a_mapped := new(string)
		*var_fe0361dd392a_mapped = val.(string)

		s.UpdatedBy = var_fe0361dd392a_mapped
	}
	if properties["createdOn"] != nil {

		var_21be0490a095 := properties["createdOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_21be0490a095)

		if err != nil {
			panic(err)
		}

		var_21be0490a095_mapped := new(time.Time)
		*var_21be0490a095_mapped = val.(time.Time)

		s.CreatedOn = var_21be0490a095_mapped
	}
	if properties["updatedOn"] != nil {

		var_df4dd8c0ed2a := properties["updatedOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_df4dd8c0ed2a)

		if err != nil {
			panic(err)
		}

		var_df4dd8c0ed2a_mapped := new(time.Time)
		*var_df4dd8c0ed2a_mapped = val.(time.Time)

		s.UpdatedOn = var_df4dd8c0ed2a_mapped
	}
	if properties["username"] != nil {

		var_6f54caefb791 := properties["username"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_6f54caefb791)

		if err != nil {
			panic(err)
		}

		var_6f54caefb791_mapped := val.(string)

		s.Username = var_6f54caefb791_mapped
	}
	if properties["password"] != nil {

		var_d899b66418a1 := properties["password"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_d899b66418a1)

		if err != nil {
			panic(err)
		}

		var_d899b66418a1_mapped := new(string)
		*var_d899b66418a1_mapped = val.(string)

		s.Password = var_d899b66418a1_mapped
	}
	if properties["roles"] != nil {

		var_6930f57501cd := properties["roles"]
		var_6930f57501cd_mapped := []string{}
		for _, v := range var_6930f57501cd.GetListValue().Values {

			var_78d064809640 := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_78d064809640)

			if err != nil {
				panic(err)
			}

			var_78d064809640_mapped := val.(string)

			var_6930f57501cd_mapped = append(var_6930f57501cd_mapped, var_78d064809640_mapped)
		}

		s.Roles = var_6930f57501cd_mapped
	}
	if properties["securityConstraints"] != nil {

		var_87af61e75803 := properties["securityConstraints"]
		var_87af61e75803_mapped := []*SecurityConstraint{}
		for _, v := range var_87af61e75803.GetListValue().Values {

			var_41b092772f36 := v
			var_41b092772f36_mapped := SecurityConstraintMapperInstance.FromProperties(var_41b092772f36.GetStructValue().Fields)

			var_87af61e75803_mapped = append(var_87af61e75803_mapped, var_41b092772f36_mapped)
		}

		s.SecurityConstraints = var_87af61e75803_mapped
	}
	if properties["details"] != nil {

		var_2ea134d4e3e6 := properties["details"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_OBJECT).UnPack(var_2ea134d4e3e6)

		if err != nil {
			panic(err)
		}

		var_2ea134d4e3e6_mapped := new(unstructured.Unstructured)
		*var_2ea134d4e3e6_mapped = val.(unstructured.Unstructured)

		s.Details = var_2ea134d4e3e6_mapped
	}
	return s
}
