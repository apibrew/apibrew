// Code generated by apbr generate. DO NOT EDIT.
// versions:
// 	apbr generate v1.2

//go:build !codeanalysis

package resource_model

import (
	"github.com/apibrew/apibrew/pkg/abs"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/types"
	"google.golang.org/protobuf/types/known/structpb"
)

import "github.com/google/uuid"
import "github.com/apibrew/apibrew/pkg/formats/unstructured"
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

func (m *UserMapper) ResourceIdentity() abs.ResourceIdentity {
	return abs.ResourceIdentity{
		Namespace: "system",
		Name:      "User",
	}
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

	var_Id := user.Id

	if var_Id != nil {
		var var_Id_mapped *structpb.Value

		var var_Id_err error
		var_Id_mapped, var_Id_err = types.ByResourcePropertyType(model.ResourceProperty_UUID).Pack(*var_Id)
		if var_Id_err != nil {
			panic(var_Id_err)
		}
		properties["id"] = var_Id_mapped
	}

	var_Version := user.Version

	var var_Version_mapped *structpb.Value

	var var_Version_err error
	var_Version_mapped, var_Version_err = types.ByResourcePropertyType(model.ResourceProperty_INT32).Pack(var_Version)
	if var_Version_err != nil {
		panic(var_Version_err)
	}
	properties["version"] = var_Version_mapped

	var_AuditData := user.AuditData

	if var_AuditData != nil {
		var var_AuditData_mapped *structpb.Value

		var_AuditData_mapped = structpb.NewStructValue(&structpb.Struct{Fields: UserAuditDataMapperInstance.ToProperties(var_AuditData)})
		properties["auditData"] = var_AuditData_mapped
	}

	var_Username := user.Username

	var var_Username_mapped *structpb.Value

	var var_Username_err error
	var_Username_mapped, var_Username_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_Username)
	if var_Username_err != nil {
		panic(var_Username_err)
	}
	properties["username"] = var_Username_mapped

	var_Password := user.Password

	if var_Password != nil {
		var var_Password_mapped *structpb.Value

		var var_Password_err error
		var_Password_mapped, var_Password_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_Password)
		if var_Password_err != nil {
			panic(var_Password_err)
		}
		properties["password"] = var_Password_mapped
	}

	var_Roles := user.Roles

	if var_Roles != nil {
		var var_Roles_mapped *structpb.Value

		var var_Roles_l []*structpb.Value
		for _, value := range var_Roles {

			var_5x := value
			var var_5x_mapped *structpb.Value

			var_5x_mapped = structpb.NewStructValue(&structpb.Struct{Fields: RoleMapperInstance.ToProperties(var_5x)})

			var_Roles_l = append(var_Roles_l, var_5x_mapped)
		}
		var_Roles_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_Roles_l})
		properties["roles"] = var_Roles_mapped
	}

	var_Permissions := user.Permissions

	if var_Permissions != nil {
		var var_Permissions_mapped *structpb.Value

		var var_Permissions_l []*structpb.Value
		for _, value := range var_Permissions {

			var_5x := value
			var var_5x_mapped *structpb.Value

			var_5x_mapped = structpb.NewStructValue(&structpb.Struct{Fields: PermissionMapperInstance.ToProperties(var_5x)})

			var_Permissions_l = append(var_Permissions_l, var_5x_mapped)
		}
		var_Permissions_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_Permissions_l})
		properties["permissions"] = var_Permissions_mapped
	}

	var_Details := user.Details

	if var_Details != nil {
		var var_Details_mapped *structpb.Value

		var var_Details_err error
		var_Details_mapped, var_Details_err = types.ByResourcePropertyType(model.ResourceProperty_OBJECT).Pack(var_Details)
		if var_Details_err != nil {
			panic(var_Details_err)
		}
		properties["details"] = var_Details_mapped
	}
	return properties
}

func (m *UserMapper) FromProperties(properties map[string]*structpb.Value) *User {
	var s = m.New()
	if properties["id"] != nil && properties["id"].AsInterface() != nil {

		var_Id := properties["id"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_UUID).UnPack(var_Id)

		if err != nil {
			panic(err)
		}

		var_Id_mapped := new(uuid.UUID)
		*var_Id_mapped = val.(uuid.UUID)

		s.Id = var_Id_mapped
	}
	if properties["version"] != nil && properties["version"].AsInterface() != nil {

		var_Version := properties["version"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(var_Version)

		if err != nil {
			panic(err)
		}

		var_Version_mapped := val.(int32)

		s.Version = var_Version_mapped
	}
	if properties["auditData"] != nil && properties["auditData"].AsInterface() != nil {

		var_AuditData := properties["auditData"]
		var mappedValue = UserAuditDataMapperInstance.FromProperties(var_AuditData.GetStructValue().Fields)

		var_AuditData_mapped := mappedValue

		s.AuditData = var_AuditData_mapped
	}
	if properties["username"] != nil && properties["username"].AsInterface() != nil {

		var_Username := properties["username"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_Username)

		if err != nil {
			panic(err)
		}

		var_Username_mapped := val.(string)

		s.Username = var_Username_mapped
	}
	if properties["password"] != nil && properties["password"].AsInterface() != nil {

		var_Password := properties["password"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_Password)

		if err != nil {
			panic(err)
		}

		var_Password_mapped := new(string)
		*var_Password_mapped = val.(string)

		s.Password = var_Password_mapped
	}
	if properties["roles"] != nil && properties["roles"].AsInterface() != nil {

		var_Roles := properties["roles"]
		var_Roles_mapped := []*Role{}
		for _, v := range var_Roles.GetListValue().Values {

			var_4x := v
			var_4x_mapped := RoleMapperInstance.FromProperties(var_4x.GetStructValue().Fields)

			var_Roles_mapped = append(var_Roles_mapped, var_4x_mapped)
		}

		s.Roles = var_Roles_mapped
	}
	if properties["permissions"] != nil && properties["permissions"].AsInterface() != nil {

		var_Permissions := properties["permissions"]
		var_Permissions_mapped := []*Permission{}
		for _, v := range var_Permissions.GetListValue().Values {

			var_4x := v
			var_4x_mapped := PermissionMapperInstance.FromProperties(var_4x.GetStructValue().Fields)

			var_Permissions_mapped = append(var_Permissions_mapped, var_4x_mapped)
		}

		s.Permissions = var_Permissions_mapped
	}
	if properties["details"] != nil && properties["details"].AsInterface() != nil {

		var_Details := properties["details"]
		var_Details_mapped := new(interface{})
		*var_Details_mapped = unstructured.FromValue(var_Details)

		s.Details = var_Details_mapped
	}
	return s
}

type UserAuditDataMapper struct {
}

func NewUserAuditDataMapper() *UserAuditDataMapper {
	return &UserAuditDataMapper{}
}

var UserAuditDataMapperInstance = NewUserAuditDataMapper()

func (m *UserAuditDataMapper) New() *UserAuditData {
	return &UserAuditData{}
}

func (m *UserAuditDataMapper) ResourceIdentity() abs.ResourceIdentity {
	return abs.ResourceIdentity{
		Namespace: "system",
		Name:      "User",
	}
}

func (m *UserAuditDataMapper) ToProperties(userAuditData *UserAuditData) map[string]*structpb.Value {
	var properties = make(map[string]*structpb.Value)

	var_CreatedBy := userAuditData.CreatedBy

	if var_CreatedBy != nil {
		var var_CreatedBy_mapped *structpb.Value

		var var_CreatedBy_err error
		var_CreatedBy_mapped, var_CreatedBy_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_CreatedBy)
		if var_CreatedBy_err != nil {
			panic(var_CreatedBy_err)
		}
		properties["createdBy"] = var_CreatedBy_mapped
	}

	var_UpdatedBy := userAuditData.UpdatedBy

	if var_UpdatedBy != nil {
		var var_UpdatedBy_mapped *structpb.Value

		var var_UpdatedBy_err error
		var_UpdatedBy_mapped, var_UpdatedBy_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_UpdatedBy)
		if var_UpdatedBy_err != nil {
			panic(var_UpdatedBy_err)
		}
		properties["updatedBy"] = var_UpdatedBy_mapped
	}

	var_CreatedOn := userAuditData.CreatedOn

	if var_CreatedOn != nil {
		var var_CreatedOn_mapped *structpb.Value

		var var_CreatedOn_err error
		var_CreatedOn_mapped, var_CreatedOn_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_CreatedOn)
		if var_CreatedOn_err != nil {
			panic(var_CreatedOn_err)
		}
		properties["createdOn"] = var_CreatedOn_mapped
	}

	var_UpdatedOn := userAuditData.UpdatedOn

	if var_UpdatedOn != nil {
		var var_UpdatedOn_mapped *structpb.Value

		var var_UpdatedOn_err error
		var_UpdatedOn_mapped, var_UpdatedOn_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_UpdatedOn)
		if var_UpdatedOn_err != nil {
			panic(var_UpdatedOn_err)
		}
		properties["updatedOn"] = var_UpdatedOn_mapped
	}
	return properties
}

func (m *UserAuditDataMapper) FromProperties(properties map[string]*structpb.Value) *UserAuditData {
	var s = m.New()
	if properties["createdBy"] != nil && properties["createdBy"].AsInterface() != nil {

		var_CreatedBy := properties["createdBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_CreatedBy)

		if err != nil {
			panic(err)
		}

		var_CreatedBy_mapped := new(string)
		*var_CreatedBy_mapped = val.(string)

		s.CreatedBy = var_CreatedBy_mapped
	}
	if properties["updatedBy"] != nil && properties["updatedBy"].AsInterface() != nil {

		var_UpdatedBy := properties["updatedBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_UpdatedBy)

		if err != nil {
			panic(err)
		}

		var_UpdatedBy_mapped := new(string)
		*var_UpdatedBy_mapped = val.(string)

		s.UpdatedBy = var_UpdatedBy_mapped
	}
	if properties["createdOn"] != nil && properties["createdOn"].AsInterface() != nil {

		var_CreatedOn := properties["createdOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_CreatedOn)

		if err != nil {
			panic(err)
		}

		var_CreatedOn_mapped := new(time.Time)
		*var_CreatedOn_mapped = val.(time.Time)

		s.CreatedOn = var_CreatedOn_mapped
	}
	if properties["updatedOn"] != nil && properties["updatedOn"].AsInterface() != nil {

		var_UpdatedOn := properties["updatedOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_UpdatedOn)

		if err != nil {
			panic(err)
		}

		var_UpdatedOn_mapped := new(time.Time)
		*var_UpdatedOn_mapped = val.(time.Time)

		s.UpdatedOn = var_UpdatedOn_mapped
	}
	return s
}
