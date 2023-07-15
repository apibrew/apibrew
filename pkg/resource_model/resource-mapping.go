package resource_model

import (
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/types"
	"google.golang.org/protobuf/types/known/structpb"
)

import "github.com/google/uuid"
import "time"
import "github.com/apibrew/apibrew/pkg/formats/unstructured"

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

	version, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).Pack(resource.Version)
	if err != nil {
		panic(err)
	}
	properties["version"] = version

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

		var_fbe8aa1bc936 := properties["id"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_UUID).UnPack(var_fbe8aa1bc936)

		if err != nil {
			panic(err)
		}

		var_fbe8aa1bc936_mapped := new(uuid.UUID)
		*var_fbe8aa1bc936_mapped = val.(uuid.UUID)

		s.Id = var_fbe8aa1bc936_mapped
	}
	if properties["version"] != nil {

		var_3e249a078d37 := properties["version"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(var_3e249a078d37)

		if err != nil {
			panic(err)
		}

		var_3e249a078d37_mapped := val.(int32)

		s.Version = var_3e249a078d37_mapped
	}
	if properties["createdBy"] != nil {

		var_89c6b1bc2181 := properties["createdBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_89c6b1bc2181)

		if err != nil {
			panic(err)
		}

		var_89c6b1bc2181_mapped := new(string)
		*var_89c6b1bc2181_mapped = val.(string)

		s.CreatedBy = var_89c6b1bc2181_mapped
	}
	if properties["updatedBy"] != nil {

		var_425d06394ddd := properties["updatedBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_425d06394ddd)

		if err != nil {
			panic(err)
		}

		var_425d06394ddd_mapped := new(string)
		*var_425d06394ddd_mapped = val.(string)

		s.UpdatedBy = var_425d06394ddd_mapped
	}
	if properties["createdOn"] != nil {

		var_7dabba9d9998 := properties["createdOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_7dabba9d9998)

		if err != nil {
			panic(err)
		}

		var_7dabba9d9998_mapped := new(time.Time)
		*var_7dabba9d9998_mapped = val.(time.Time)

		s.CreatedOn = var_7dabba9d9998_mapped
	}
	if properties["updatedOn"] != nil {

		var_fc92238bdb3e := properties["updatedOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_fc92238bdb3e)

		if err != nil {
			panic(err)
		}

		var_fc92238bdb3e_mapped := new(time.Time)
		*var_fc92238bdb3e_mapped = val.(time.Time)

		s.UpdatedOn = var_fc92238bdb3e_mapped
	}
	if properties["name"] != nil {

		var_46e227e09e2e := properties["name"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_46e227e09e2e)

		if err != nil {
			panic(err)
		}

		var_46e227e09e2e_mapped := val.(string)

		s.Name = var_46e227e09e2e_mapped
	}
	if properties["namespace"] != nil {

		var_83154fae4070 := properties["namespace"]
		var_83154fae4070_mapped := NamespaceMapperInstance.FromProperties(var_83154fae4070.GetStructValue().Fields)

		s.Namespace = var_83154fae4070_mapped
	}
	if properties["virtual"] != nil {

		var_8cb8b9a7fbf6 := properties["virtual"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_BOOL).UnPack(var_8cb8b9a7fbf6)

		if err != nil {
			panic(err)
		}

		var_8cb8b9a7fbf6_mapped := val.(bool)

		s.Virtual = var_8cb8b9a7fbf6_mapped
	}
	if properties["types"] != nil {

		var_77fc3f67a570 := properties["types"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_OBJECT).UnPack(var_77fc3f67a570)

		if err != nil {
			panic(err)
		}

		var_77fc3f67a570_mapped := new(unstructured.Unstructured)
		*var_77fc3f67a570_mapped = val.(unstructured.Unstructured)

		s.Types = var_77fc3f67a570_mapped
	}
	if properties["immutable"] != nil {

		var_4490c1d9c7a0 := properties["immutable"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_BOOL).UnPack(var_4490c1d9c7a0)

		if err != nil {
			panic(err)
		}

		var_4490c1d9c7a0_mapped := val.(bool)

		s.Immutable = var_4490c1d9c7a0_mapped
	}
	if properties["abstract"] != nil {

		var_91f5978b457e := properties["abstract"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_BOOL).UnPack(var_91f5978b457e)

		if err != nil {
			panic(err)
		}

		var_91f5978b457e_mapped := val.(bool)

		s.Abstract = var_91f5978b457e_mapped
	}
	if properties["dataSource"] != nil {

		var_6929d464e5f1 := properties["dataSource"]
		var_6929d464e5f1_mapped := DataSourceMapperInstance.FromProperties(var_6929d464e5f1.GetStructValue().Fields)

		s.DataSource = var_6929d464e5f1_mapped
	}
	if properties["entity"] != nil {

		var_29ecd3197df0 := properties["entity"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_29ecd3197df0)

		if err != nil {
			panic(err)
		}

		var_29ecd3197df0_mapped := new(string)
		*var_29ecd3197df0_mapped = val.(string)

		s.Entity = var_29ecd3197df0_mapped
	}
	if properties["catalog"] != nil {

		var_222dc4f1411d := properties["catalog"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_222dc4f1411d)

		if err != nil {
			panic(err)
		}

		var_222dc4f1411d_mapped := new(string)
		*var_222dc4f1411d_mapped = val.(string)

		s.Catalog = var_222dc4f1411d_mapped
	}
	if properties["annotations"] != nil {

		var_c0ee975d9072 := properties["annotations"]
		var_c0ee975d9072_mapped := make(map[string]string)
		for k, v := range var_c0ee975d9072.GetStructValue().Fields {

			var_f0f0aa4f0c90 := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_f0f0aa4f0c90)

			if err != nil {
				panic(err)
			}

			var_f0f0aa4f0c90_mapped := val.(string)

			var_c0ee975d9072_mapped[k] = var_f0f0aa4f0c90_mapped
		}

		s.Annotations = var_c0ee975d9072_mapped
	}
	if properties["indexes"] != nil {

		var_bedec0467982 := properties["indexes"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_OBJECT).UnPack(var_bedec0467982)

		if err != nil {
			panic(err)
		}

		var_bedec0467982_mapped := new(unstructured.Unstructured)
		*var_bedec0467982_mapped = val.(unstructured.Unstructured)

		s.Indexes = var_bedec0467982_mapped
	}
	if properties["securityConstraints"] != nil {

		var_ff537794adea := properties["securityConstraints"]
		var_ff537794adea_mapped := []*SecurityConstraint{}
		for _, v := range var_ff537794adea.GetListValue().Values {

			var_8dce054b7db1 := v
			var_8dce054b7db1_mapped := SecurityConstraintMapperInstance.FromProperties(var_8dce054b7db1.GetStructValue().Fields)

			var_ff537794adea_mapped = append(var_ff537794adea_mapped, var_8dce054b7db1_mapped)
		}

		s.SecurityConstraints = var_ff537794adea_mapped
	}
	if properties["title"] != nil {

		var_7c3f3e3160dd := properties["title"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_7c3f3e3160dd)

		if err != nil {
			panic(err)
		}

		var_7c3f3e3160dd_mapped := new(string)
		*var_7c3f3e3160dd_mapped = val.(string)

		s.Title = var_7c3f3e3160dd_mapped
	}
	if properties["description"] != nil {

		var_126778a97ee1 := properties["description"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_126778a97ee1)

		if err != nil {
			panic(err)
		}

		var_126778a97ee1_mapped := new(string)
		*var_126778a97ee1_mapped = val.(string)

		s.Description = var_126778a97ee1_mapped
	}
	return s
}
