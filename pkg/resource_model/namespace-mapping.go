package resource_model

import (
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/types"
	"google.golang.org/protobuf/types/known/structpb"
)

import "github.com/google/uuid"
import "time"
import "github.com/apibrew/apibrew/pkg/formats/unstructured"

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

		var_f673f35e7fac := properties["id"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_UUID).UnPack(var_f673f35e7fac)

		if err != nil {
			panic(err)
		}

		var_f673f35e7fac_mapped := new(uuid.UUID)
		*var_f673f35e7fac_mapped = val.(uuid.UUID)

		s.Id = var_f673f35e7fac_mapped
	}
	if properties["version"] != nil {

		var_f8919bf02153 := properties["version"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(var_f8919bf02153)

		if err != nil {
			panic(err)
		}

		var_f8919bf02153_mapped := new(int32)
		*var_f8919bf02153_mapped = val.(int32)

		s.Version = var_f8919bf02153_mapped
	}
	if properties["createdBy"] != nil {

		var_77aade2a50f1 := properties["createdBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_77aade2a50f1)

		if err != nil {
			panic(err)
		}

		var_77aade2a50f1_mapped := new(string)
		*var_77aade2a50f1_mapped = val.(string)

		s.CreatedBy = var_77aade2a50f1_mapped
	}
	if properties["updatedBy"] != nil {

		var_f98c504c628e := properties["updatedBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_f98c504c628e)

		if err != nil {
			panic(err)
		}

		var_f98c504c628e_mapped := new(string)
		*var_f98c504c628e_mapped = val.(string)

		s.UpdatedBy = var_f98c504c628e_mapped
	}
	if properties["createdOn"] != nil {

		var_dfccf0f12a84 := properties["createdOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_dfccf0f12a84)

		if err != nil {
			panic(err)
		}

		var_dfccf0f12a84_mapped := new(time.Time)
		*var_dfccf0f12a84_mapped = val.(time.Time)

		s.CreatedOn = var_dfccf0f12a84_mapped
	}
	if properties["updatedOn"] != nil {

		var_376a1b34b4c0 := properties["updatedOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_376a1b34b4c0)

		if err != nil {
			panic(err)
		}

		var_376a1b34b4c0_mapped := new(time.Time)
		*var_376a1b34b4c0_mapped = val.(time.Time)

		s.UpdatedOn = var_376a1b34b4c0_mapped
	}
	if properties["name"] != nil {

		var_1c13ea0b3bf9 := properties["name"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_1c13ea0b3bf9)

		if err != nil {
			panic(err)
		}

		var_1c13ea0b3bf9_mapped := val.(string)

		s.Name = var_1c13ea0b3bf9_mapped
	}
	if properties["description"] != nil {

		var_55c9dea44795 := properties["description"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_55c9dea44795)

		if err != nil {
			panic(err)
		}

		var_55c9dea44795_mapped := new(string)
		*var_55c9dea44795_mapped = val.(string)

		s.Description = var_55c9dea44795_mapped
	}
	if properties["details"] != nil {

		var_752c4072ef8e := properties["details"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_OBJECT).UnPack(var_752c4072ef8e)

		if err != nil {
			panic(err)
		}

		var_752c4072ef8e_mapped := new(unstructured.Unstructured)
		*var_752c4072ef8e_mapped = val.(unstructured.Unstructured)

		s.Details = var_752c4072ef8e_mapped
	}
	if properties["securityConstraints"] != nil {

		var_cce274e4747f := properties["securityConstraints"]
		var_cce274e4747f_mapped := []*SecurityConstraint{}
		for _, v := range var_cce274e4747f.GetListValue().Values {

			var_5acc4f17c5bf := v
			var_5acc4f17c5bf_mapped := SecurityConstraintMapperInstance.FromProperties(var_5acc4f17c5bf.GetStructValue().Fields)

			var_cce274e4747f_mapped = append(var_cce274e4747f_mapped, var_5acc4f17c5bf_mapped)
		}

		s.SecurityConstraints = var_cce274e4747f_mapped
	}
	return s
}
