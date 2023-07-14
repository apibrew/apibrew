package resource_model

import (
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/types"
	"google.golang.org/protobuf/types/known/structpb"
)

import "github.com/google/uuid"
import "time"

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

		var_0b7866c87508 := properties["id"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_UUID).UnPack(var_0b7866c87508)

		if err != nil {
			panic(err)
		}

		var_0b7866c87508_mapped := new(uuid.UUID)
		*var_0b7866c87508_mapped = val.(uuid.UUID)

		s.Id = var_0b7866c87508_mapped
	}
	if properties["version"] != nil {

		var_3e646d9b2572 := properties["version"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(var_3e646d9b2572)

		if err != nil {
			panic(err)
		}

		var_3e646d9b2572_mapped := new(int32)
		*var_3e646d9b2572_mapped = val.(int32)

		s.Version = var_3e646d9b2572_mapped
	}
	if properties["createdBy"] != nil {

		var_c4e244025f2d := properties["createdBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_c4e244025f2d)

		if err != nil {
			panic(err)
		}

		var_c4e244025f2d_mapped := new(string)
		*var_c4e244025f2d_mapped = val.(string)

		s.CreatedBy = var_c4e244025f2d_mapped
	}
	if properties["updatedBy"] != nil {

		var_5d8c7d3c1a11 := properties["updatedBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_5d8c7d3c1a11)

		if err != nil {
			panic(err)
		}

		var_5d8c7d3c1a11_mapped := new(string)
		*var_5d8c7d3c1a11_mapped = val.(string)

		s.UpdatedBy = var_5d8c7d3c1a11_mapped
	}
	if properties["createdOn"] != nil {

		var_b64a165626b0 := properties["createdOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_b64a165626b0)

		if err != nil {
			panic(err)
		}

		var_b64a165626b0_mapped := new(time.Time)
		*var_b64a165626b0_mapped = val.(time.Time)

		s.CreatedOn = var_b64a165626b0_mapped
	}
	if properties["updatedOn"] != nil {

		var_3bd9010ed8f4 := properties["updatedOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_3bd9010ed8f4)

		if err != nil {
			panic(err)
		}

		var_3bd9010ed8f4_mapped := new(time.Time)
		*var_3bd9010ed8f4_mapped = val.(time.Time)

		s.UpdatedOn = var_3bd9010ed8f4_mapped
	}
	if properties["username"] != nil {

		var_f7fb52ed84b8 := properties["username"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_f7fb52ed84b8)

		if err != nil {
			panic(err)
		}

		var_f7fb52ed84b8_mapped := val.(string)

		s.Username = var_f7fb52ed84b8_mapped
	}
	if properties["password"] != nil {

		var_e5c165aa2c79 := properties["password"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_e5c165aa2c79)

		if err != nil {
			panic(err)
		}

		var_e5c165aa2c79_mapped := new(string)
		*var_e5c165aa2c79_mapped = val.(string)

		s.Password = var_e5c165aa2c79_mapped
	}
	if properties["roles"] != nil {

		var_87d71eb82643 := properties["roles"]
		var_87d71eb82643_mapped := []string{}
		for _, v := range var_87d71eb82643.GetListValue().Values {

			var_5759dd0dc299 := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_5759dd0dc299)

			if err != nil {
				panic(err)
			}

			var_5759dd0dc299_mapped := val.(string)

			var_87d71eb82643_mapped = append(var_87d71eb82643_mapped, var_5759dd0dc299_mapped)
		}

		s.Roles = var_87d71eb82643_mapped
	}
	if properties["securityConstraints"] != nil {

		var_58914eb19ab8 := properties["securityConstraints"]
		var_58914eb19ab8_mapped := []*SecurityConstraint{}
		for _, v := range var_58914eb19ab8.GetListValue().Values {

			var_3bd782203a7e := v
			var_3bd782203a7e_mapped := SecurityConstraintMapperInstance.FromProperties(var_3bd782203a7e.GetStructValue().Fields)

			var_58914eb19ab8_mapped = append(var_58914eb19ab8_mapped, var_3bd782203a7e_mapped)
		}

		s.SecurityConstraints = var_58914eb19ab8_mapped
	}
	if properties["details"] != nil {

		var_9b7e013f00e3 := properties["details"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_OBJECT).UnPack(var_9b7e013f00e3)

		if err != nil {
			panic(err)
		}

		var_9b7e013f00e3_mapped := new(interface{})
		*var_9b7e013f00e3_mapped = val.(interface{})

		s.Details = var_9b7e013f00e3_mapped
	}
	return s
}
