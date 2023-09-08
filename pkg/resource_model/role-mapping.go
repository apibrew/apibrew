// AUTOGENERATED FILE

//go:build !codeanalysis

package resource_model

import (
	"github.com/apibrew/apibrew/pkg/abs"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/types"
	"google.golang.org/protobuf/types/known/structpb"
)

import "github.com/google/uuid"
import "time"
import "github.com/apibrew/apibrew/pkg/formats/unstructured"

type RoleMapper struct {
}

func NewRoleMapper() *RoleMapper {
	return &RoleMapper{}
}

var RoleMapperInstance = NewRoleMapper()

func (m *RoleMapper) New() *Role {
	return &Role{}
}

func (m *RoleMapper) ResourceIdentity() abs.ResourceIdentity {
	return abs.ResourceIdentity{
		Namespace: "system",
		Name:      "Role",
	}
}

func (m *RoleMapper) ToRecord(role *Role) *model.Record {
	var rec = &model.Record{}
	rec.Properties = m.ToProperties(role)

	if role.Id != nil {
		rec.Id = role.Id.String()
	}

	return rec
}

func (m *RoleMapper) FromRecord(record *model.Record) *Role {
	return m.FromProperties(record.Properties)
}

func (m *RoleMapper) ToProperties(role *Role) map[string]*structpb.Value {
	var properties = make(map[string]*structpb.Value)

	var_Id := role.Id

	if var_Id != nil {
		var var_Id_mapped *structpb.Value

		var var_Id_err error
		var_Id_mapped, var_Id_err = types.ByResourcePropertyType(model.ResourceProperty_UUID).Pack(*var_Id)
		if var_Id_err != nil {
			panic(var_Id_err)
		}
		properties["id"] = var_Id_mapped
	}

	var_Version := role.Version

	var var_Version_mapped *structpb.Value

	var var_Version_err error
	var_Version_mapped, var_Version_err = types.ByResourcePropertyType(model.ResourceProperty_INT32).Pack(var_Version)
	if var_Version_err != nil {
		panic(var_Version_err)
	}
	properties["version"] = var_Version_mapped

	var_CreatedBy := role.CreatedBy

	if var_CreatedBy != nil {
		var var_CreatedBy_mapped *structpb.Value

		var var_CreatedBy_err error
		var_CreatedBy_mapped, var_CreatedBy_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_CreatedBy)
		if var_CreatedBy_err != nil {
			panic(var_CreatedBy_err)
		}
		properties["createdBy"] = var_CreatedBy_mapped
	}

	var_UpdatedBy := role.UpdatedBy

	if var_UpdatedBy != nil {
		var var_UpdatedBy_mapped *structpb.Value

		var var_UpdatedBy_err error
		var_UpdatedBy_mapped, var_UpdatedBy_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_UpdatedBy)
		if var_UpdatedBy_err != nil {
			panic(var_UpdatedBy_err)
		}
		properties["updatedBy"] = var_UpdatedBy_mapped
	}

	var_CreatedOn := role.CreatedOn

	if var_CreatedOn != nil {
		var var_CreatedOn_mapped *structpb.Value

		var var_CreatedOn_err error
		var_CreatedOn_mapped, var_CreatedOn_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_CreatedOn)
		if var_CreatedOn_err != nil {
			panic(var_CreatedOn_err)
		}
		properties["createdOn"] = var_CreatedOn_mapped
	}

	var_UpdatedOn := role.UpdatedOn

	if var_UpdatedOn != nil {
		var var_UpdatedOn_mapped *structpb.Value

		var var_UpdatedOn_err error
		var_UpdatedOn_mapped, var_UpdatedOn_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_UpdatedOn)
		if var_UpdatedOn_err != nil {
			panic(var_UpdatedOn_err)
		}
		properties["updatedOn"] = var_UpdatedOn_mapped
	}

	var_Name := role.Name

	var var_Name_mapped *structpb.Value

	var var_Name_err error
	var_Name_mapped, var_Name_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_Name)
	if var_Name_err != nil {
		panic(var_Name_err)
	}
	properties["name"] = var_Name_mapped

	var_Permissions := role.Permissions

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

	var_Details := role.Details

	if var_Details != nil {
		var var_Details_mapped *structpb.Value

		var var_Details_err error
		var_Details_mapped, var_Details_err = types.ByResourcePropertyType(model.ResourceProperty_OBJECT).Pack(*var_Details)
		if var_Details_err != nil {
			panic(var_Details_err)
		}
		properties["details"] = var_Details_mapped
	}
	return properties
}

func (m *RoleMapper) FromProperties(properties map[string]*structpb.Value) *Role {
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
	if properties["name"] != nil && properties["name"].AsInterface() != nil {

		var_Name := properties["name"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_Name)

		if err != nil {
			panic(err)
		}

		var_Name_mapped := val.(string)

		s.Name = var_Name_mapped
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
		var_Details_mapped := new(unstructured.Unstructured)
		*var_Details_mapped = unstructured.FromStructValue(var_Details.GetStructValue())

		s.Details = var_Details_mapped
	}
	return s
}
