package resource_model

import (
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/types"
	"google.golang.org/protobuf/types/known/structpb"
)

import "github.com/google/uuid"
import "time"

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

		var_b37a1059ef5b := properties["id"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_UUID).UnPack(var_b37a1059ef5b)

		if err != nil {
			panic(err)
		}

		var_b37a1059ef5b_mapped := new(uuid.UUID)
		*var_b37a1059ef5b_mapped = val.(uuid.UUID)

		s.Id = var_b37a1059ef5b_mapped
	}
	if properties["version"] != nil {

		var_a4cf13770cdb := properties["version"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(var_a4cf13770cdb)

		if err != nil {
			panic(err)
		}

		var_a4cf13770cdb_mapped := new(int32)
		*var_a4cf13770cdb_mapped = val.(int32)

		s.Version = var_a4cf13770cdb_mapped
	}
	if properties["createdBy"] != nil {

		var_ff358c5f6d4d := properties["createdBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_ff358c5f6d4d)

		if err != nil {
			panic(err)
		}

		var_ff358c5f6d4d_mapped := new(string)
		*var_ff358c5f6d4d_mapped = val.(string)

		s.CreatedBy = var_ff358c5f6d4d_mapped
	}
	if properties["updatedBy"] != nil {

		var_5ef32d91b2a0 := properties["updatedBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_5ef32d91b2a0)

		if err != nil {
			panic(err)
		}

		var_5ef32d91b2a0_mapped := new(string)
		*var_5ef32d91b2a0_mapped = val.(string)

		s.UpdatedBy = var_5ef32d91b2a0_mapped
	}
	if properties["createdOn"] != nil {

		var_2539cf142926 := properties["createdOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_2539cf142926)

		if err != nil {
			panic(err)
		}

		var_2539cf142926_mapped := new(time.Time)
		*var_2539cf142926_mapped = val.(time.Time)

		s.CreatedOn = var_2539cf142926_mapped
	}
	if properties["updatedOn"] != nil {

		var_8b5aff622ccd := properties["updatedOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_8b5aff622ccd)

		if err != nil {
			panic(err)
		}

		var_8b5aff622ccd_mapped := new(time.Time)
		*var_8b5aff622ccd_mapped = val.(time.Time)

		s.UpdatedOn = var_8b5aff622ccd_mapped
	}
	if properties["name"] != nil {

		var_08b8bc564857 := properties["name"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_08b8bc564857)

		if err != nil {
			panic(err)
		}

		var_08b8bc564857_mapped := val.(string)

		s.Name = var_08b8bc564857_mapped
	}
	if properties["securityConstraints"] != nil {

		var_70eabdb66110 := properties["securityConstraints"]
		var_70eabdb66110_mapped := []*SecurityConstraint{}
		for _, v := range var_70eabdb66110.GetListValue().Values {

			var_dea18a39e070 := v
			var_dea18a39e070_mapped := SecurityConstraintMapperInstance.FromProperties(var_dea18a39e070.GetStructValue().Fields)

			var_70eabdb66110_mapped = append(var_70eabdb66110_mapped, var_dea18a39e070_mapped)
		}

		s.SecurityConstraints = var_70eabdb66110_mapped
	}
	if properties["details"] != nil {

		var_6dce3b522db9 := properties["details"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_OBJECT).UnPack(var_6dce3b522db9)

		if err != nil {
			panic(err)
		}

		var_6dce3b522db9_mapped := new(interface{})
		*var_6dce3b522db9_mapped = val.(interface{})

		s.Details = var_6dce3b522db9_mapped
	}
	return s
}
