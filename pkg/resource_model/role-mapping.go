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

		var_8113bb29ea65 := properties["id"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_UUID).UnPack(var_8113bb29ea65)

		if err != nil {
			panic(err)
		}

		var_8113bb29ea65_mapped := new(uuid.UUID)
		*var_8113bb29ea65_mapped = val.(uuid.UUID)

		s.Id = var_8113bb29ea65_mapped
	}
	if properties["version"] != nil {

		var_4d1570faf1df := properties["version"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(var_4d1570faf1df)

		if err != nil {
			panic(err)
		}

		var_4d1570faf1df_mapped := new(int32)
		*var_4d1570faf1df_mapped = val.(int32)

		s.Version = var_4d1570faf1df_mapped
	}
	if properties["createdBy"] != nil {

		var_9be3a6e70b1f := properties["createdBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_9be3a6e70b1f)

		if err != nil {
			panic(err)
		}

		var_9be3a6e70b1f_mapped := new(string)
		*var_9be3a6e70b1f_mapped = val.(string)

		s.CreatedBy = var_9be3a6e70b1f_mapped
	}
	if properties["updatedBy"] != nil {

		var_eafc160a284d := properties["updatedBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_eafc160a284d)

		if err != nil {
			panic(err)
		}

		var_eafc160a284d_mapped := new(string)
		*var_eafc160a284d_mapped = val.(string)

		s.UpdatedBy = var_eafc160a284d_mapped
	}
	if properties["createdOn"] != nil {

		var_a225ecc35c5a := properties["createdOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_a225ecc35c5a)

		if err != nil {
			panic(err)
		}

		var_a225ecc35c5a_mapped := new(time.Time)
		*var_a225ecc35c5a_mapped = val.(time.Time)

		s.CreatedOn = var_a225ecc35c5a_mapped
	}
	if properties["updatedOn"] != nil {

		var_a0e0ecf21a9d := properties["updatedOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_a0e0ecf21a9d)

		if err != nil {
			panic(err)
		}

		var_a0e0ecf21a9d_mapped := new(time.Time)
		*var_a0e0ecf21a9d_mapped = val.(time.Time)

		s.UpdatedOn = var_a0e0ecf21a9d_mapped
	}
	if properties["name"] != nil {

		var_6222181ba12f := properties["name"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_6222181ba12f)

		if err != nil {
			panic(err)
		}

		var_6222181ba12f_mapped := val.(string)

		s.Name = var_6222181ba12f_mapped
	}
	if properties["securityConstraints"] != nil {

		var_66e23a2b9467 := properties["securityConstraints"]
		var_66e23a2b9467_mapped := []*SecurityConstraint{}
		for _, v := range var_66e23a2b9467.GetListValue().Values {

			var_0cfee9e7cd7d := v
			var_0cfee9e7cd7d_mapped := SecurityConstraintMapperInstance.FromProperties(var_0cfee9e7cd7d.GetStructValue().Fields)

			var_66e23a2b9467_mapped = append(var_66e23a2b9467_mapped, var_0cfee9e7cd7d_mapped)
		}

		s.SecurityConstraints = var_66e23a2b9467_mapped
	}
	if properties["details"] != nil {

		var_b7867bf7f52e := properties["details"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_OBJECT).UnPack(var_b7867bf7f52e)

		if err != nil {
			panic(err)
		}

		var_b7867bf7f52e_mapped := new(unstructured.Unstructured)
		*var_b7867bf7f52e_mapped = val.(unstructured.Unstructured)

		s.Details = var_b7867bf7f52e_mapped
	}
	return s
}
