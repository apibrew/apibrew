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

	version, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).Pack(namespace.Version)
	if err != nil {
		panic(err)
	}
	properties["version"] = version

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

		var_b2f08f26815d := properties["id"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_UUID).UnPack(var_b2f08f26815d)

		if err != nil {
			panic(err)
		}

		var_b2f08f26815d_mapped := new(uuid.UUID)
		*var_b2f08f26815d_mapped = val.(uuid.UUID)

		s.Id = var_b2f08f26815d_mapped
	}
	if properties["version"] != nil {

		var_e6930c4eaf01 := properties["version"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(var_e6930c4eaf01)

		if err != nil {
			panic(err)
		}

		var_e6930c4eaf01_mapped := val.(int32)

		s.Version = var_e6930c4eaf01_mapped
	}
	if properties["createdBy"] != nil {

		var_1260ce6912e7 := properties["createdBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_1260ce6912e7)

		if err != nil {
			panic(err)
		}

		var_1260ce6912e7_mapped := new(string)
		*var_1260ce6912e7_mapped = val.(string)

		s.CreatedBy = var_1260ce6912e7_mapped
	}
	if properties["updatedBy"] != nil {

		var_6b00832615d4 := properties["updatedBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_6b00832615d4)

		if err != nil {
			panic(err)
		}

		var_6b00832615d4_mapped := new(string)
		*var_6b00832615d4_mapped = val.(string)

		s.UpdatedBy = var_6b00832615d4_mapped
	}
	if properties["createdOn"] != nil {

		var_b9da714bb049 := properties["createdOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_b9da714bb049)

		if err != nil {
			panic(err)
		}

		var_b9da714bb049_mapped := new(time.Time)
		*var_b9da714bb049_mapped = val.(time.Time)

		s.CreatedOn = var_b9da714bb049_mapped
	}
	if properties["updatedOn"] != nil {

		var_6afb574ebefa := properties["updatedOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_6afb574ebefa)

		if err != nil {
			panic(err)
		}

		var_6afb574ebefa_mapped := new(time.Time)
		*var_6afb574ebefa_mapped = val.(time.Time)

		s.UpdatedOn = var_6afb574ebefa_mapped
	}
	if properties["name"] != nil {

		var_9edad22440bf := properties["name"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_9edad22440bf)

		if err != nil {
			panic(err)
		}

		var_9edad22440bf_mapped := val.(string)

		s.Name = var_9edad22440bf_mapped
	}
	if properties["description"] != nil {

		var_d2aa54285336 := properties["description"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_d2aa54285336)

		if err != nil {
			panic(err)
		}

		var_d2aa54285336_mapped := new(string)
		*var_d2aa54285336_mapped = val.(string)

		s.Description = var_d2aa54285336_mapped
	}
	if properties["details"] != nil {

		var_25afe3c86eb9 := properties["details"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_OBJECT).UnPack(var_25afe3c86eb9)

		if err != nil {
			panic(err)
		}

		var_25afe3c86eb9_mapped := new(unstructured.Unstructured)
		*var_25afe3c86eb9_mapped = val.(unstructured.Unstructured)

		s.Details = var_25afe3c86eb9_mapped
	}
	if properties["securityConstraints"] != nil {

		var_18560cdf1695 := properties["securityConstraints"]
		var_18560cdf1695_mapped := []*SecurityConstraint{}
		for _, v := range var_18560cdf1695.GetListValue().Values {

			var_4fc9100bc17f := v
			var_4fc9100bc17f_mapped := SecurityConstraintMapperInstance.FromProperties(var_4fc9100bc17f.GetStructValue().Fields)

			var_18560cdf1695_mapped = append(var_18560cdf1695_mapped, var_4fc9100bc17f_mapped)
		}

		s.SecurityConstraints = var_18560cdf1695_mapped
	}
	return s
}
