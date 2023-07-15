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

type RoleMapper struct {
}

func NewRoleMapper() *RoleMapper {
	return &RoleMapper{}
}

var RoleMapperInstance = NewRoleMapper()

func (m *RoleMapper) New() *Role {
	return &Role{}
}

func (m *RoleMapper) ToRecord(role *Role) *model.Record {
	var rec = &model.Record{}
	rec.Properties = m.ToProperties(role)
	return rec
}

func (m *RoleMapper) FromRecord(record *model.Record) *Role {
	return m.FromProperties(record.Properties)
}

func (m *RoleMapper) ToProperties(role *Role) map[string]*structpb.Value {
	var properties = make(map[string]*structpb.Value)

	if role.Id != nil {
		id, err := types.ByResourcePropertyType(model.ResourceProperty_UUID).Pack(*role.Id)
		if err != nil {
			panic(err)
		}
		properties["id"] = id
	}

	if role.Version != nil {
		version, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).Pack(*role.Version)
		if err != nil {
			panic(err)
		}
		properties["version"] = version
	}

	if role.CreatedBy != nil {
		createdBy, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*role.CreatedBy)
		if err != nil {
			panic(err)
		}
		properties["createdBy"] = createdBy
	}

	if role.UpdatedBy != nil {
		updatedBy, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*role.UpdatedBy)
		if err != nil {
			panic(err)
		}
		properties["updatedBy"] = updatedBy
	}

	if role.CreatedOn != nil {
		createdOn, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*role.CreatedOn)
		if err != nil {
			panic(err)
		}
		properties["createdOn"] = createdOn
	}

	if role.UpdatedOn != nil {
		updatedOn, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*role.UpdatedOn)
		if err != nil {
			panic(err)
		}
		properties["updatedOn"] = updatedOn
	}

	name, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(role.Name)
	if err != nil {
		panic(err)
	}
	properties["name"] = name

	if role.SecurityConstraints != nil {
	}

	if role.Details != nil {
	}

	return properties
}

func (m *RoleMapper) FromProperties(properties map[string]*structpb.Value) *Role {
	var s = m.New()
	if properties["id"] != nil {

		var_62fbe487d808 := properties["id"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_UUID).UnPack(var_62fbe487d808)

		if err != nil {
			panic(err)
		}

		var_62fbe487d808_mapped := new(uuid.UUID)
		*var_62fbe487d808_mapped = val.(uuid.UUID)

		s.Id = var_62fbe487d808_mapped
	}
	if properties["version"] != nil {

		var_6557cca90548 := properties["version"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(var_6557cca90548)

		if err != nil {
			panic(err)
		}

		var_6557cca90548_mapped := new(int32)
		*var_6557cca90548_mapped = val.(int32)

		s.Version = var_6557cca90548_mapped
	}
	if properties["createdBy"] != nil {

		var_11accbdc1c54 := properties["createdBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_11accbdc1c54)

		if err != nil {
			panic(err)
		}

		var_11accbdc1c54_mapped := new(string)
		*var_11accbdc1c54_mapped = val.(string)

		s.CreatedBy = var_11accbdc1c54_mapped
	}
	if properties["updatedBy"] != nil {

		var_0e3e0720b1e8 := properties["updatedBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_0e3e0720b1e8)

		if err != nil {
			panic(err)
		}

		var_0e3e0720b1e8_mapped := new(string)
		*var_0e3e0720b1e8_mapped = val.(string)

		s.UpdatedBy = var_0e3e0720b1e8_mapped
	}
	if properties["createdOn"] != nil {

		var_558a71b4bbc4 := properties["createdOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_558a71b4bbc4)

		if err != nil {
			panic(err)
		}

		var_558a71b4bbc4_mapped := new(time.Time)
		*var_558a71b4bbc4_mapped = val.(time.Time)

		s.CreatedOn = var_558a71b4bbc4_mapped
	}
	if properties["updatedOn"] != nil {

		var_c130a63ad77a := properties["updatedOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_c130a63ad77a)

		if err != nil {
			panic(err)
		}

		var_c130a63ad77a_mapped := new(time.Time)
		*var_c130a63ad77a_mapped = val.(time.Time)

		s.UpdatedOn = var_c130a63ad77a_mapped
	}
	if properties["name"] != nil {

		var_bf20543506d0 := properties["name"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_bf20543506d0)

		if err != nil {
			panic(err)
		}

		var_bf20543506d0_mapped := val.(string)

		s.Name = var_bf20543506d0_mapped
	}
	if properties["securityConstraints"] != nil {

		var_eb00f7e13eed := properties["securityConstraints"]
		var_eb00f7e13eed_mapped := []*SecurityConstraint{}
		for _, v := range var_eb00f7e13eed.GetListValue().Values {

			var_c1a560ad48fd := v
			var_c1a560ad48fd_mapped := SecurityConstraintMapperInstance.FromProperties(var_c1a560ad48fd.GetStructValue().Fields)

			var_eb00f7e13eed_mapped = append(var_eb00f7e13eed_mapped, var_c1a560ad48fd_mapped)
		}

		s.SecurityConstraints = var_eb00f7e13eed_mapped
	}
	if properties["details"] != nil {

		var_fdb546e6fb2d := properties["details"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_OBJECT).UnPack(var_fdb546e6fb2d)

		if err != nil {
			panic(err)
		}

		var_fdb546e6fb2d_mapped := new(unstructured.Unstructured)
		*var_fdb546e6fb2d_mapped = val.(unstructured.Unstructured)

		s.Details = var_fdb546e6fb2d_mapped
	}
	return s
}
