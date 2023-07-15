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

type NamespaceMapper struct {
}

func NewNamespaceMapper() *NamespaceMapper {
	return &NamespaceMapper{}
}

var NamespaceMapperInstance = NewNamespaceMapper()

func (m *NamespaceMapper) New() *Namespace {
	return &Namespace{}
}

func (m *NamespaceMapper) ToRecord(namespace *Namespace) *model.Record {
	var rec = &model.Record{}
	rec.Properties = m.ToProperties(namespace)
	return rec
}

func (m *NamespaceMapper) FromRecord(record *model.Record) *Namespace {
	return m.FromProperties(record.Properties)
}

func (m *NamespaceMapper) ToProperties(namespace *Namespace) map[string]*structpb.Value {
	var properties = make(map[string]*structpb.Value)

	if namespace.Id != nil {
		id, err := types.ByResourcePropertyType(model.ResourceProperty_UUID).Pack(*namespace.Id)
		if err != nil {
			panic(err)
		}
		properties["id"] = id
	}

	if namespace.Version != nil {
		version, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).Pack(*namespace.Version)
		if err != nil {
			panic(err)
		}
		properties["version"] = version
	}

	if namespace.CreatedBy != nil {
		createdBy, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*namespace.CreatedBy)
		if err != nil {
			panic(err)
		}
		properties["createdBy"] = createdBy
	}

	if namespace.UpdatedBy != nil {
		updatedBy, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*namespace.UpdatedBy)
		if err != nil {
			panic(err)
		}
		properties["updatedBy"] = updatedBy
	}

	if namespace.CreatedOn != nil {
		createdOn, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*namespace.CreatedOn)
		if err != nil {
			panic(err)
		}
		properties["createdOn"] = createdOn
	}

	if namespace.UpdatedOn != nil {
		updatedOn, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*namespace.UpdatedOn)
		if err != nil {
			panic(err)
		}
		properties["updatedOn"] = updatedOn
	}

	name, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(namespace.Name)
	if err != nil {
		panic(err)
	}
	properties["name"] = name

	if namespace.Description != nil {
		description, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*namespace.Description)
		if err != nil {
			panic(err)
		}
		properties["description"] = description
	}

	if namespace.Details != nil {
	}

	if namespace.SecurityConstraints != nil {
	}

	return properties
}

func (m *NamespaceMapper) FromProperties(properties map[string]*structpb.Value) *Namespace {
	var s = m.New()
	if properties["id"] != nil {

		var_8d24431cc143 := properties["id"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_UUID).UnPack(var_8d24431cc143)

		if err != nil {
			panic(err)
		}

		var_8d24431cc143_mapped := new(uuid.UUID)
		*var_8d24431cc143_mapped = val.(uuid.UUID)

		s.Id = var_8d24431cc143_mapped
	}
	if properties["version"] != nil {

		var_0809b5ef005a := properties["version"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(var_0809b5ef005a)

		if err != nil {
			panic(err)
		}

		var_0809b5ef005a_mapped := new(int32)
		*var_0809b5ef005a_mapped = val.(int32)

		s.Version = var_0809b5ef005a_mapped
	}
	if properties["createdBy"] != nil {

		var_247253a71098 := properties["createdBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_247253a71098)

		if err != nil {
			panic(err)
		}

		var_247253a71098_mapped := new(string)
		*var_247253a71098_mapped = val.(string)

		s.CreatedBy = var_247253a71098_mapped
	}
	if properties["updatedBy"] != nil {

		var_6e879bd02668 := properties["updatedBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_6e879bd02668)

		if err != nil {
			panic(err)
		}

		var_6e879bd02668_mapped := new(string)
		*var_6e879bd02668_mapped = val.(string)

		s.UpdatedBy = var_6e879bd02668_mapped
	}
	if properties["createdOn"] != nil {

		var_faa6820e6f99 := properties["createdOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_faa6820e6f99)

		if err != nil {
			panic(err)
		}

		var_faa6820e6f99_mapped := new(time.Time)
		*var_faa6820e6f99_mapped = val.(time.Time)

		s.CreatedOn = var_faa6820e6f99_mapped
	}
	if properties["updatedOn"] != nil {

		var_1b0973d5ef8d := properties["updatedOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_1b0973d5ef8d)

		if err != nil {
			panic(err)
		}

		var_1b0973d5ef8d_mapped := new(time.Time)
		*var_1b0973d5ef8d_mapped = val.(time.Time)

		s.UpdatedOn = var_1b0973d5ef8d_mapped
	}
	if properties["name"] != nil {

		var_c54eb48f4c5c := properties["name"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_c54eb48f4c5c)

		if err != nil {
			panic(err)
		}

		var_c54eb48f4c5c_mapped := val.(string)

		s.Name = var_c54eb48f4c5c_mapped
	}
	if properties["description"] != nil {

		var_a0afad0ec330 := properties["description"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_a0afad0ec330)

		if err != nil {
			panic(err)
		}

		var_a0afad0ec330_mapped := new(string)
		*var_a0afad0ec330_mapped = val.(string)

		s.Description = var_a0afad0ec330_mapped
	}
	if properties["details"] != nil {

		var_c7fc69db7dc6 := properties["details"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_OBJECT).UnPack(var_c7fc69db7dc6)

		if err != nil {
			panic(err)
		}

		var_c7fc69db7dc6_mapped := new(unstructured.Unstructured)
		*var_c7fc69db7dc6_mapped = val.(unstructured.Unstructured)

		s.Details = var_c7fc69db7dc6_mapped
	}
	if properties["securityConstraints"] != nil {

		var_e23f0987119c := properties["securityConstraints"]
		var_e23f0987119c_mapped := []*SecurityConstraint{}
		for _, v := range var_e23f0987119c.GetListValue().Values {

			var_c3ce834f1269 := v
			var_c3ce834f1269_mapped := SecurityConstraintMapperInstance.FromProperties(var_c3ce834f1269.GetStructValue().Fields)

			var_e23f0987119c_mapped = append(var_e23f0987119c_mapped, var_c3ce834f1269_mapped)
		}

		s.SecurityConstraints = var_e23f0987119c_mapped
	}
	return s
}
