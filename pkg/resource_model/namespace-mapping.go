package resource_model

import (
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/types"
	"google.golang.org/protobuf/types/known/structpb"
)

import "github.com/google/uuid"
import "time"

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

		var_8381236f16f5 := properties["id"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_UUID).UnPack(var_8381236f16f5)

		if err != nil {
			panic(err)
		}

		var_8381236f16f5_mapped := new(uuid.UUID)
		*var_8381236f16f5_mapped = val.(uuid.UUID)

		s.Id = var_8381236f16f5_mapped
	}
	if properties["version"] != nil {

		var_4c18b9e9b8fb := properties["version"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(var_4c18b9e9b8fb)

		if err != nil {
			panic(err)
		}

		var_4c18b9e9b8fb_mapped := new(int32)
		*var_4c18b9e9b8fb_mapped = val.(int32)

		s.Version = var_4c18b9e9b8fb_mapped
	}
	if properties["createdBy"] != nil {

		var_7641f22990ee := properties["createdBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_7641f22990ee)

		if err != nil {
			panic(err)
		}

		var_7641f22990ee_mapped := new(string)
		*var_7641f22990ee_mapped = val.(string)

		s.CreatedBy = var_7641f22990ee_mapped
	}
	if properties["updatedBy"] != nil {

		var_09835a5472e7 := properties["updatedBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_09835a5472e7)

		if err != nil {
			panic(err)
		}

		var_09835a5472e7_mapped := new(string)
		*var_09835a5472e7_mapped = val.(string)

		s.UpdatedBy = var_09835a5472e7_mapped
	}
	if properties["createdOn"] != nil {

		var_3d3101dd274e := properties["createdOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_3d3101dd274e)

		if err != nil {
			panic(err)
		}

		var_3d3101dd274e_mapped := new(time.Time)
		*var_3d3101dd274e_mapped = val.(time.Time)

		s.CreatedOn = var_3d3101dd274e_mapped
	}
	if properties["updatedOn"] != nil {

		var_4c9a404e06cc := properties["updatedOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_4c9a404e06cc)

		if err != nil {
			panic(err)
		}

		var_4c9a404e06cc_mapped := new(time.Time)
		*var_4c9a404e06cc_mapped = val.(time.Time)

		s.UpdatedOn = var_4c9a404e06cc_mapped
	}
	if properties["name"] != nil {

		var_ba1c351cc02a := properties["name"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_ba1c351cc02a)

		if err != nil {
			panic(err)
		}

		var_ba1c351cc02a_mapped := val.(string)

		s.Name = var_ba1c351cc02a_mapped
	}
	if properties["description"] != nil {

		var_75c9d12d6d1d := properties["description"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_75c9d12d6d1d)

		if err != nil {
			panic(err)
		}

		var_75c9d12d6d1d_mapped := new(string)
		*var_75c9d12d6d1d_mapped = val.(string)

		s.Description = var_75c9d12d6d1d_mapped
	}
	if properties["details"] != nil {

		var_1f5b34ac80cd := properties["details"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_OBJECT).UnPack(var_1f5b34ac80cd)

		if err != nil {
			panic(err)
		}

		var_1f5b34ac80cd_mapped := new(interface{})
		*var_1f5b34ac80cd_mapped = val.(interface{})

		s.Details = var_1f5b34ac80cd_mapped
	}
	if properties["securityConstraints"] != nil {

		var_c91df897506b := properties["securityConstraints"]
		var_c91df897506b_mapped := []*SecurityConstraint{}
		for _, v := range var_c91df897506b.GetListValue().Values {

			var_2b9bb0871122 := v
			var_2b9bb0871122_mapped := SecurityConstraintMapperInstance.FromProperties(var_2b9bb0871122.GetStructValue().Fields)

			var_c91df897506b_mapped = append(var_c91df897506b_mapped, var_2b9bb0871122_mapped)
		}

		s.SecurityConstraints = var_c91df897506b_mapped
	}
	return s
}
