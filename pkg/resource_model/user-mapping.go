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

		var_0b5db3d7e6ff := properties["id"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_UUID).UnPack(var_0b5db3d7e6ff)

		if err != nil {
			panic(err)
		}

		var_0b5db3d7e6ff_mapped := new(uuid.UUID)
		*var_0b5db3d7e6ff_mapped = val.(uuid.UUID)

		s.Id = var_0b5db3d7e6ff_mapped
	}
	if properties["version"] != nil {

		var_5fababce22ba := properties["version"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(var_5fababce22ba)

		if err != nil {
			panic(err)
		}

		var_5fababce22ba_mapped := new(int32)
		*var_5fababce22ba_mapped = val.(int32)

		s.Version = var_5fababce22ba_mapped
	}
	if properties["createdBy"] != nil {

		var_be01109c258d := properties["createdBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_be01109c258d)

		if err != nil {
			panic(err)
		}

		var_be01109c258d_mapped := new(string)
		*var_be01109c258d_mapped = val.(string)

		s.CreatedBy = var_be01109c258d_mapped
	}
	if properties["updatedBy"] != nil {

		var_464fe94ca825 := properties["updatedBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_464fe94ca825)

		if err != nil {
			panic(err)
		}

		var_464fe94ca825_mapped := new(string)
		*var_464fe94ca825_mapped = val.(string)

		s.UpdatedBy = var_464fe94ca825_mapped
	}
	if properties["createdOn"] != nil {

		var_42109cf4d1b7 := properties["createdOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_42109cf4d1b7)

		if err != nil {
			panic(err)
		}

		var_42109cf4d1b7_mapped := new(time.Time)
		*var_42109cf4d1b7_mapped = val.(time.Time)

		s.CreatedOn = var_42109cf4d1b7_mapped
	}
	if properties["updatedOn"] != nil {

		var_c340f14acb09 := properties["updatedOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_c340f14acb09)

		if err != nil {
			panic(err)
		}

		var_c340f14acb09_mapped := new(time.Time)
		*var_c340f14acb09_mapped = val.(time.Time)

		s.UpdatedOn = var_c340f14acb09_mapped
	}
	if properties["username"] != nil {

		var_1691680209fa := properties["username"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_1691680209fa)

		if err != nil {
			panic(err)
		}

		var_1691680209fa_mapped := val.(string)

		s.Username = var_1691680209fa_mapped
	}
	if properties["password"] != nil {

		var_00bea2b7f3dd := properties["password"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_00bea2b7f3dd)

		if err != nil {
			panic(err)
		}

		var_00bea2b7f3dd_mapped := new(string)
		*var_00bea2b7f3dd_mapped = val.(string)

		s.Password = var_00bea2b7f3dd_mapped
	}
	if properties["roles"] != nil {

		var_efd4fdd5b275 := properties["roles"]
		var_efd4fdd5b275_mapped := []string{}
		for _, v := range var_efd4fdd5b275.GetListValue().Values {

			var_7cc2a53db92e := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_7cc2a53db92e)

			if err != nil {
				panic(err)
			}

			var_7cc2a53db92e_mapped := val.(string)

			var_efd4fdd5b275_mapped = append(var_efd4fdd5b275_mapped, var_7cc2a53db92e_mapped)
		}

		s.Roles = var_efd4fdd5b275_mapped
	}
	if properties["securityConstraints"] != nil {

		var_7c63e647ae19 := properties["securityConstraints"]
		var_7c63e647ae19_mapped := []*SecurityConstraint{}
		for _, v := range var_7c63e647ae19.GetListValue().Values {

			var_ee38dd7889f3 := v
			var_ee38dd7889f3_mapped := SecurityConstraintMapperInstance.FromProperties(var_ee38dd7889f3.GetStructValue().Fields)

			var_7c63e647ae19_mapped = append(var_7c63e647ae19_mapped, var_ee38dd7889f3_mapped)
		}

		s.SecurityConstraints = var_7c63e647ae19_mapped
	}
	if properties["details"] != nil {

		var_05c8708dfa8a := properties["details"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_OBJECT).UnPack(var_05c8708dfa8a)

		if err != nil {
			panic(err)
		}

		var_05c8708dfa8a_mapped := new(unstructured.Unstructured)
		*var_05c8708dfa8a_mapped = val.(unstructured.Unstructured)

		s.Details = var_05c8708dfa8a_mapped
	}
	return s
}
