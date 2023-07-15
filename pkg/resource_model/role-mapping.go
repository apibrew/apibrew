package resource_model

import (
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

	version, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).Pack(role.Version)
	if err != nil {
		panic(err)
	}
	properties["version"] = version

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

		var_da26b6cf3036 := properties["id"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_UUID).UnPack(var_da26b6cf3036)

		if err != nil {
			panic(err)
		}

		var_da26b6cf3036_mapped := new(uuid.UUID)
		*var_da26b6cf3036_mapped = val.(uuid.UUID)

		s.Id = var_da26b6cf3036_mapped
	}
	if properties["version"] != nil {

		var_5003e1013b7e := properties["version"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(var_5003e1013b7e)

		if err != nil {
			panic(err)
		}

		var_5003e1013b7e_mapped := val.(int32)

		s.Version = var_5003e1013b7e_mapped
	}
	if properties["createdBy"] != nil {

		var_8bce88422e13 := properties["createdBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_8bce88422e13)

		if err != nil {
			panic(err)
		}

		var_8bce88422e13_mapped := new(string)
		*var_8bce88422e13_mapped = val.(string)

		s.CreatedBy = var_8bce88422e13_mapped
	}
	if properties["updatedBy"] != nil {

		var_308408836181 := properties["updatedBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_308408836181)

		if err != nil {
			panic(err)
		}

		var_308408836181_mapped := new(string)
		*var_308408836181_mapped = val.(string)

		s.UpdatedBy = var_308408836181_mapped
	}
	if properties["createdOn"] != nil {

		var_3e46d32cae44 := properties["createdOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_3e46d32cae44)

		if err != nil {
			panic(err)
		}

		var_3e46d32cae44_mapped := new(time.Time)
		*var_3e46d32cae44_mapped = val.(time.Time)

		s.CreatedOn = var_3e46d32cae44_mapped
	}
	if properties["updatedOn"] != nil {

		var_12f1a624ee2b := properties["updatedOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_12f1a624ee2b)

		if err != nil {
			panic(err)
		}

		var_12f1a624ee2b_mapped := new(time.Time)
		*var_12f1a624ee2b_mapped = val.(time.Time)

		s.UpdatedOn = var_12f1a624ee2b_mapped
	}
	if properties["name"] != nil {

		var_ebe940fecbab := properties["name"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_ebe940fecbab)

		if err != nil {
			panic(err)
		}

		var_ebe940fecbab_mapped := val.(string)

		s.Name = var_ebe940fecbab_mapped
	}
	if properties["securityConstraints"] != nil {

		var_d2ef874239c1 := properties["securityConstraints"]
		var_d2ef874239c1_mapped := []*SecurityConstraint{}
		for _, v := range var_d2ef874239c1.GetListValue().Values {

			var_f307f2e77b76 := v
			var_f307f2e77b76_mapped := SecurityConstraintMapperInstance.FromProperties(var_f307f2e77b76.GetStructValue().Fields)

			var_d2ef874239c1_mapped = append(var_d2ef874239c1_mapped, var_f307f2e77b76_mapped)
		}

		s.SecurityConstraints = var_d2ef874239c1_mapped
	}
	if properties["details"] != nil {

		var_a3c4e5b9ab67 := properties["details"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_OBJECT).UnPack(var_a3c4e5b9ab67)

		if err != nil {
			panic(err)
		}

		var_a3c4e5b9ab67_mapped := new(unstructured.Unstructured)
		*var_a3c4e5b9ab67_mapped = val.(unstructured.Unstructured)

		s.Details = var_a3c4e5b9ab67_mapped
	}
	return s
}
