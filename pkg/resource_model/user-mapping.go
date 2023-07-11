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

	if user.SecurityConstraints != nil {
	}

	if user.Details != nil {
	}

	if user.Roles != nil {
	}

	return properties
}

func (m *UserMapper) FromProperties(properties map[string]*structpb.Value) *User {
	var s = m.New()
	if properties["id"] != nil {
		val, err := types.ByResourcePropertyType(model.ResourceProperty_UUID).UnPack(properties["id"])

		if err != nil {
			panic(err)
		}

		s.Id = new(uuid.UUID)
		*s.Id = val.(uuid.UUID)
	}
	if properties["version"] != nil {
		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(properties["version"])

		if err != nil {
			panic(err)
		}

		s.Version = new(int32)
		*s.Version = val.(int32)
	}
	if properties["createdBy"] != nil {
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(properties["createdBy"])

		if err != nil {
			panic(err)
		}

		s.CreatedBy = new(string)
		*s.CreatedBy = val.(string)
	}
	if properties["updatedBy"] != nil {
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(properties["updatedBy"])

		if err != nil {
			panic(err)
		}

		s.UpdatedBy = new(string)
		*s.UpdatedBy = val.(string)
	}
	if properties["createdOn"] != nil {
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(properties["createdOn"])

		if err != nil {
			panic(err)
		}

		s.CreatedOn = new(time.Time)
		*s.CreatedOn = val.(time.Time)
	}
	if properties["updatedOn"] != nil {
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(properties["updatedOn"])

		if err != nil {
			panic(err)
		}

		s.UpdatedOn = new(time.Time)
		*s.UpdatedOn = val.(time.Time)
	}
	if properties["username"] != nil {
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(properties["username"])

		if err != nil {
			panic(err)
		}

		s.Username = val.(string)
	}
	if properties["password"] != nil {
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(properties["password"])

		if err != nil {
			panic(err)
		}

		s.Password = new(string)
		*s.Password = val.(string)
	}
	if properties["securityConstraints"] != nil {
		s.SecurityConstraints = []interface{}{}
		for _, v := range properties["securityConstraints"].AsInterface().([]interface{}) {
			s.SecurityConstraints = append(s.SecurityConstraints, v.(interface{}))
		}
	}
	if properties["details"] != nil {
		val, err := types.ByResourcePropertyType(model.ResourceProperty_OBJECT).UnPack(properties["details"])

		if err != nil {
			panic(err)
		}

		s.Details = new(interface{})
		*s.Details = val.(interface{})
	}
	if properties["roles"] != nil {
		s.Roles = []string{}
		for _, v := range properties["roles"].AsInterface().([]interface{}) {
			s.Roles = append(s.Roles, v.(string))
		}
	}
	return s
}
