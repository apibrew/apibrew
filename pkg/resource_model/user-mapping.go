package resource_model

import (
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/types"
	"google.golang.org/protobuf/types/known/structpb"
)

import "github.com/google/uuid"
import "time"
import "github.com/apibrew/apibrew/pkg/formats/unstructured"
import "encoding/json"

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

	if user.Version != nil {
		version, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).Pack(*user.Version)
		if err != nil {
			panic(err)
		}
		properties["version"] = version
	}

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

		var_77fa14ad8f6a := properties["id"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_UUID).UnPack(var_77fa14ad8f6a)

		if err != nil {
			panic(err)
		}

		var_77fa14ad8f6a_mapped := new(uuid.UUID)
		*var_77fa14ad8f6a_mapped = val.(uuid.UUID)

		s.Id = var_77fa14ad8f6a_mapped
	}
	if properties["version"] != nil {

		var_18c6d6e128a1 := properties["version"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(var_18c6d6e128a1)

		if err != nil {
			panic(err)
		}

		var_18c6d6e128a1_mapped := new(int32)
		*var_18c6d6e128a1_mapped = val.(int32)

		s.Version = var_18c6d6e128a1_mapped
	}
	if properties["createdBy"] != nil {

		var_e9c9342c7d25 := properties["createdBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_e9c9342c7d25)

		if err != nil {
			panic(err)
		}

		var_e9c9342c7d25_mapped := new(string)
		*var_e9c9342c7d25_mapped = val.(string)

		s.CreatedBy = var_e9c9342c7d25_mapped
	}
	if properties["updatedBy"] != nil {

		var_1c51216389a0 := properties["updatedBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_1c51216389a0)

		if err != nil {
			panic(err)
		}

		var_1c51216389a0_mapped := new(string)
		*var_1c51216389a0_mapped = val.(string)

		s.UpdatedBy = var_1c51216389a0_mapped
	}
	if properties["createdOn"] != nil {

		var_8b277c9994d7 := properties["createdOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_8b277c9994d7)

		if err != nil {
			panic(err)
		}

		var_8b277c9994d7_mapped := new(time.Time)
		*var_8b277c9994d7_mapped = val.(time.Time)

		s.CreatedOn = var_8b277c9994d7_mapped
	}
	if properties["updatedOn"] != nil {

		var_2d0a7f3d9e0d := properties["updatedOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_2d0a7f3d9e0d)

		if err != nil {
			panic(err)
		}

		var_2d0a7f3d9e0d_mapped := new(time.Time)
		*var_2d0a7f3d9e0d_mapped = val.(time.Time)

		s.UpdatedOn = var_2d0a7f3d9e0d_mapped
	}
	if properties["username"] != nil {

		var_cb2724dda80e := properties["username"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_cb2724dda80e)

		if err != nil {
			panic(err)
		}

		var_cb2724dda80e_mapped := val.(string)

		s.Username = var_cb2724dda80e_mapped
	}
	if properties["password"] != nil {

		var_746edd358288 := properties["password"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_746edd358288)

		if err != nil {
			panic(err)
		}

		var_746edd358288_mapped := new(string)
		*var_746edd358288_mapped = val.(string)

		s.Password = var_746edd358288_mapped
	}
	if properties["roles"] != nil {

		var_ec6db218cc27 := properties["roles"]
		var_ec6db218cc27_mapped := []string{}
		for _, v := range var_ec6db218cc27.GetListValue().Values {

			var_75e55681edae := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_75e55681edae)

			if err != nil {
				panic(err)
			}

			var_75e55681edae_mapped := val.(string)

			var_ec6db218cc27_mapped = append(var_ec6db218cc27_mapped, var_75e55681edae_mapped)
		}

		s.Roles = var_ec6db218cc27_mapped
	}
	if properties["securityConstraints"] != nil {

		var_d0ee2aeed200 := properties["securityConstraints"]
		var_d0ee2aeed200_mapped := []*SecurityConstraint{}
		for _, v := range var_d0ee2aeed200.GetListValue().Values {

			var_e23c46bde30f := v
			var_e23c46bde30f_mapped := SecurityConstraintMapperInstance.FromProperties(var_e23c46bde30f.GetStructValue().Fields)

			var_d0ee2aeed200_mapped = append(var_d0ee2aeed200_mapped, var_e23c46bde30f_mapped)
		}

		s.SecurityConstraints = var_d0ee2aeed200_mapped
	}
	if properties["details"] != nil {

		var_0165c9434c8c := properties["details"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_OBJECT).UnPack(var_0165c9434c8c)

		if err != nil {
			panic(err)
		}

		var_0165c9434c8c_mapped := new(unstructured.Unstructured)
		*var_0165c9434c8c_mapped = val.(unstructured.Unstructured)

		s.Details = var_0165c9434c8c_mapped
	}
	return s
}
