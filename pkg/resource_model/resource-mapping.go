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

type ResourceMapper struct {
}

func NewResourceMapper() *ResourceMapper {
	return &ResourceMapper{}
}

var ResourceMapperInstance = NewResourceMapper()

func (m *ResourceMapper) New() *Resource {
	return &Resource{}
}

func (m *ResourceMapper) ToRecord(resource *Resource) *model.Record {
	var rec = &model.Record{}
	rec.Properties = m.ToProperties(resource)
	return rec
}

func (m *ResourceMapper) FromRecord(record *model.Record) *Resource {
	return m.FromProperties(record.Properties)
}

func (m *ResourceMapper) ToProperties(resource *Resource) map[string]*structpb.Value {
	var properties = make(map[string]*structpb.Value)

	if resource.Id != nil {
		id, err := types.ByResourcePropertyType(model.ResourceProperty_UUID).Pack(*resource.Id)
		if err != nil {
			panic(err)
		}
		properties["id"] = id
	}

	if resource.Version != nil {
		version, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).Pack(*resource.Version)
		if err != nil {
			panic(err)
		}
		properties["version"] = version
	}

	if resource.CreatedBy != nil {
		createdBy, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*resource.CreatedBy)
		if err != nil {
			panic(err)
		}
		properties["createdBy"] = createdBy
	}

	if resource.UpdatedBy != nil {
		updatedBy, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*resource.UpdatedBy)
		if err != nil {
			panic(err)
		}
		properties["updatedBy"] = updatedBy
	}

	if resource.CreatedOn != nil {
		createdOn, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*resource.CreatedOn)
		if err != nil {
			panic(err)
		}
		properties["createdOn"] = createdOn
	}

	if resource.UpdatedOn != nil {
		updatedOn, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*resource.UpdatedOn)
		if err != nil {
			panic(err)
		}
		properties["updatedOn"] = updatedOn
	}

	name, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(resource.Name)
	if err != nil {
		panic(err)
	}
	properties["name"] = name

	if resource.Namespace != nil {
	}

	virtual, err := types.ByResourcePropertyType(model.ResourceProperty_BOOL).Pack(resource.Virtual)
	if err != nil {
		panic(err)
	}
	properties["virtual"] = virtual

	if resource.Types != nil {
	}

	immutable, err := types.ByResourcePropertyType(model.ResourceProperty_BOOL).Pack(resource.Immutable)
	if err != nil {
		panic(err)
	}
	properties["immutable"] = immutable

	abstract, err := types.ByResourcePropertyType(model.ResourceProperty_BOOL).Pack(resource.Abstract)
	if err != nil {
		panic(err)
	}
	properties["abstract"] = abstract

	if resource.DataSource != nil {
	}

	if resource.Entity != nil {
		entity, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*resource.Entity)
		if err != nil {
			panic(err)
		}
		properties["entity"] = entity
	}

	if resource.Catalog != nil {
		catalog, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*resource.Catalog)
		if err != nil {
			panic(err)
		}
		properties["catalog"] = catalog
	}

	if resource.Annotations != nil {
	}

	if resource.Indexes != nil {
	}

	if resource.SecurityConstraints != nil {
	}

	if resource.Title != nil {
		title, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*resource.Title)
		if err != nil {
			panic(err)
		}
		properties["title"] = title
	}

	if resource.Description != nil {
		description, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*resource.Description)
		if err != nil {
			panic(err)
		}
		properties["description"] = description
	}

	return properties
}

func (m *ResourceMapper) FromProperties(properties map[string]*structpb.Value) *Resource {
	var s = m.New()
	if properties["id"] != nil {

		var_7af1a3a5144a := properties["id"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_UUID).UnPack(var_7af1a3a5144a)

		if err != nil {
			panic(err)
		}

		var_7af1a3a5144a_mapped := new(uuid.UUID)
		*var_7af1a3a5144a_mapped = val.(uuid.UUID)

		s.Id = var_7af1a3a5144a_mapped
	}
	if properties["version"] != nil {

		var_a161c660ab26 := properties["version"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(var_a161c660ab26)

		if err != nil {
			panic(err)
		}

		var_a161c660ab26_mapped := new(int32)
		*var_a161c660ab26_mapped = val.(int32)

		s.Version = var_a161c660ab26_mapped
	}
	if properties["createdBy"] != nil {

		var_7df42c1b4cfd := properties["createdBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_7df42c1b4cfd)

		if err != nil {
			panic(err)
		}

		var_7df42c1b4cfd_mapped := new(string)
		*var_7df42c1b4cfd_mapped = val.(string)

		s.CreatedBy = var_7df42c1b4cfd_mapped
	}
	if properties["updatedBy"] != nil {

		var_e6cc6b1ee9d9 := properties["updatedBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_e6cc6b1ee9d9)

		if err != nil {
			panic(err)
		}

		var_e6cc6b1ee9d9_mapped := new(string)
		*var_e6cc6b1ee9d9_mapped = val.(string)

		s.UpdatedBy = var_e6cc6b1ee9d9_mapped
	}
	if properties["createdOn"] != nil {

		var_8c131797391f := properties["createdOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_8c131797391f)

		if err != nil {
			panic(err)
		}

		var_8c131797391f_mapped := new(time.Time)
		*var_8c131797391f_mapped = val.(time.Time)

		s.CreatedOn = var_8c131797391f_mapped
	}
	if properties["updatedOn"] != nil {

		var_2b7fea306e81 := properties["updatedOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_2b7fea306e81)

		if err != nil {
			panic(err)
		}

		var_2b7fea306e81_mapped := new(time.Time)
		*var_2b7fea306e81_mapped = val.(time.Time)

		s.UpdatedOn = var_2b7fea306e81_mapped
	}
	if properties["name"] != nil {

		var_e0fbbd7b2d10 := properties["name"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_e0fbbd7b2d10)

		if err != nil {
			panic(err)
		}

		var_e0fbbd7b2d10_mapped := val.(string)

		s.Name = var_e0fbbd7b2d10_mapped
	}
	if properties["namespace"] != nil {

		var_0c57f6f2b379 := properties["namespace"]
		var_0c57f6f2b379_mapped := NamespaceMapperInstance.FromProperties(var_0c57f6f2b379.GetStructValue().Fields)

		s.Namespace = var_0c57f6f2b379_mapped
	}
	if properties["virtual"] != nil {

		var_16b1c4d9f435 := properties["virtual"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_BOOL).UnPack(var_16b1c4d9f435)

		if err != nil {
			panic(err)
		}

		var_16b1c4d9f435_mapped := val.(bool)

		s.Virtual = var_16b1c4d9f435_mapped
	}
	if properties["types"] != nil {

		var_dabf526bcd1a := properties["types"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_OBJECT).UnPack(var_dabf526bcd1a)

		if err != nil {
			panic(err)
		}

		var_dabf526bcd1a_mapped := new(unstructured.Unstructured)
		*var_dabf526bcd1a_mapped = val.(unstructured.Unstructured)

		s.Types = var_dabf526bcd1a_mapped
	}
	if properties["immutable"] != nil {

		var_3f4ec8799c7c := properties["immutable"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_BOOL).UnPack(var_3f4ec8799c7c)

		if err != nil {
			panic(err)
		}

		var_3f4ec8799c7c_mapped := val.(bool)

		s.Immutable = var_3f4ec8799c7c_mapped
	}
	if properties["abstract"] != nil {

		var_f05a39e7016e := properties["abstract"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_BOOL).UnPack(var_f05a39e7016e)

		if err != nil {
			panic(err)
		}

		var_f05a39e7016e_mapped := val.(bool)

		s.Abstract = var_f05a39e7016e_mapped
	}
	if properties["dataSource"] != nil {

		var_7e09abbde5ce := properties["dataSource"]
		var_7e09abbde5ce_mapped := DataSourceMapperInstance.FromProperties(var_7e09abbde5ce.GetStructValue().Fields)

		s.DataSource = var_7e09abbde5ce_mapped
	}
	if properties["entity"] != nil {

		var_0fa7a7d0a1d6 := properties["entity"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_0fa7a7d0a1d6)

		if err != nil {
			panic(err)
		}

		var_0fa7a7d0a1d6_mapped := new(string)
		*var_0fa7a7d0a1d6_mapped = val.(string)

		s.Entity = var_0fa7a7d0a1d6_mapped
	}
	if properties["catalog"] != nil {

		var_1bbba61fd843 := properties["catalog"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_1bbba61fd843)

		if err != nil {
			panic(err)
		}

		var_1bbba61fd843_mapped := new(string)
		*var_1bbba61fd843_mapped = val.(string)

		s.Catalog = var_1bbba61fd843_mapped
	}
	if properties["annotations"] != nil {

		var_c99bd5820a81 := properties["annotations"]
		var_c99bd5820a81_mapped := make(map[string]string)
		for k, v := range var_c99bd5820a81.GetStructValue().Fields {

			var_7fd121b46b6b := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_7fd121b46b6b)

			if err != nil {
				panic(err)
			}

			var_7fd121b46b6b_mapped := val.(string)

			var_c99bd5820a81_mapped[k] = var_7fd121b46b6b_mapped
		}

		s.Annotations = var_c99bd5820a81_mapped
	}
	if properties["indexes"] != nil {

		var_566ce64ba652 := properties["indexes"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_OBJECT).UnPack(var_566ce64ba652)

		if err != nil {
			panic(err)
		}

		var_566ce64ba652_mapped := new(unstructured.Unstructured)
		*var_566ce64ba652_mapped = val.(unstructured.Unstructured)

		s.Indexes = var_566ce64ba652_mapped
	}
	if properties["securityConstraints"] != nil {

		var_5d72b5fac888 := properties["securityConstraints"]
		var_5d72b5fac888_mapped := []*SecurityConstraint{}
		for _, v := range var_5d72b5fac888.GetListValue().Values {

			var_accf712e0941 := v
			var_accf712e0941_mapped := SecurityConstraintMapperInstance.FromProperties(var_accf712e0941.GetStructValue().Fields)

			var_5d72b5fac888_mapped = append(var_5d72b5fac888_mapped, var_accf712e0941_mapped)
		}

		s.SecurityConstraints = var_5d72b5fac888_mapped
	}
	if properties["title"] != nil {

		var_85a1345fb4eb := properties["title"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_85a1345fb4eb)

		if err != nil {
			panic(err)
		}

		var_85a1345fb4eb_mapped := new(string)
		*var_85a1345fb4eb_mapped = val.(string)

		s.Title = var_85a1345fb4eb_mapped
	}
	if properties["description"] != nil {

		var_66e95dce76ef := properties["description"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_66e95dce76ef)

		if err != nil {
			panic(err)
		}

		var_66e95dce76ef_mapped := new(string)
		*var_66e95dce76ef_mapped = val.(string)

		s.Description = var_66e95dce76ef_mapped
	}
	return s
}
