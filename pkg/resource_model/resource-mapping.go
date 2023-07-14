package resource_model

import (
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/types"
	"google.golang.org/protobuf/types/known/structpb"
)

import "github.com/google/uuid"
import "time"

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

		var_421029505ad0 := properties["id"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_UUID).UnPack(var_421029505ad0)

		if err != nil {
			panic(err)
		}

		var_421029505ad0_mapped := new(uuid.UUID)
		*var_421029505ad0_mapped = val.(uuid.UUID)

		s.Id = var_421029505ad0_mapped
	}
	if properties["version"] != nil {

		var_fdd37c41b9ef := properties["version"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(var_fdd37c41b9ef)

		if err != nil {
			panic(err)
		}

		var_fdd37c41b9ef_mapped := new(int32)
		*var_fdd37c41b9ef_mapped = val.(int32)

		s.Version = var_fdd37c41b9ef_mapped
	}
	if properties["createdBy"] != nil {

		var_f095505ac5d4 := properties["createdBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_f095505ac5d4)

		if err != nil {
			panic(err)
		}

		var_f095505ac5d4_mapped := new(string)
		*var_f095505ac5d4_mapped = val.(string)

		s.CreatedBy = var_f095505ac5d4_mapped
	}
	if properties["updatedBy"] != nil {

		var_deb0ee165afd := properties["updatedBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_deb0ee165afd)

		if err != nil {
			panic(err)
		}

		var_deb0ee165afd_mapped := new(string)
		*var_deb0ee165afd_mapped = val.(string)

		s.UpdatedBy = var_deb0ee165afd_mapped
	}
	if properties["createdOn"] != nil {

		var_5a2d77864b42 := properties["createdOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_5a2d77864b42)

		if err != nil {
			panic(err)
		}

		var_5a2d77864b42_mapped := new(time.Time)
		*var_5a2d77864b42_mapped = val.(time.Time)

		s.CreatedOn = var_5a2d77864b42_mapped
	}
	if properties["updatedOn"] != nil {

		var_0127ef342243 := properties["updatedOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_0127ef342243)

		if err != nil {
			panic(err)
		}

		var_0127ef342243_mapped := new(time.Time)
		*var_0127ef342243_mapped = val.(time.Time)

		s.UpdatedOn = var_0127ef342243_mapped
	}
	if properties["name"] != nil {

		var_fbea309f8bf0 := properties["name"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_fbea309f8bf0)

		if err != nil {
			panic(err)
		}

		var_fbea309f8bf0_mapped := val.(string)

		s.Name = var_fbea309f8bf0_mapped
	}
	if properties["namespace"] != nil {

		var_0f120541da64 := properties["namespace"]
		var_0f120541da64_mapped := NamespaceMapperInstance.FromProperties(var_0f120541da64.GetStructValue().Fields)

		s.Namespace = var_0f120541da64_mapped
	}
	if properties["virtual"] != nil {

		var_7d66629e971c := properties["virtual"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_BOOL).UnPack(var_7d66629e971c)

		if err != nil {
			panic(err)
		}

		var_7d66629e971c_mapped := val.(bool)

		s.Virtual = var_7d66629e971c_mapped
	}
	if properties["types"] != nil {

		var_c783b8d29691 := properties["types"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_OBJECT).UnPack(var_c783b8d29691)

		if err != nil {
			panic(err)
		}

		var_c783b8d29691_mapped := new(interface{})
		*var_c783b8d29691_mapped = val.(interface{})

		s.Types = var_c783b8d29691_mapped
	}
	if properties["immutable"] != nil {

		var_1743f4fa89ea := properties["immutable"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_BOOL).UnPack(var_1743f4fa89ea)

		if err != nil {
			panic(err)
		}

		var_1743f4fa89ea_mapped := val.(bool)

		s.Immutable = var_1743f4fa89ea_mapped
	}
	if properties["abstract"] != nil {

		var_102654a90736 := properties["abstract"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_BOOL).UnPack(var_102654a90736)

		if err != nil {
			panic(err)
		}

		var_102654a90736_mapped := val.(bool)

		s.Abstract = var_102654a90736_mapped
	}
	if properties["dataSource"] != nil {

		var_206d5682db6d := properties["dataSource"]
		var_206d5682db6d_mapped := DataSourceMapperInstance.FromProperties(var_206d5682db6d.GetStructValue().Fields)

		s.DataSource = var_206d5682db6d_mapped
	}
	if properties["entity"] != nil {

		var_8354f9ff1216 := properties["entity"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_8354f9ff1216)

		if err != nil {
			panic(err)
		}

		var_8354f9ff1216_mapped := new(string)
		*var_8354f9ff1216_mapped = val.(string)

		s.Entity = var_8354f9ff1216_mapped
	}
	if properties["catalog"] != nil {

		var_911785e3efba := properties["catalog"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_911785e3efba)

		if err != nil {
			panic(err)
		}

		var_911785e3efba_mapped := new(string)
		*var_911785e3efba_mapped = val.(string)

		s.Catalog = var_911785e3efba_mapped
	}
	if properties["annotations"] != nil {

		var_6a2f2efd3201 := properties["annotations"]
		var_6a2f2efd3201_mapped := make(map[string]string)
		for k, v := range var_6a2f2efd3201.GetStructValue().Fields {

			var_dcc3e895fca0 := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_dcc3e895fca0)

			if err != nil {
				panic(err)
			}

			var_dcc3e895fca0_mapped := val.(string)

			var_6a2f2efd3201_mapped[k] = var_dcc3e895fca0_mapped
		}

		s.Annotations = var_6a2f2efd3201_mapped
	}
	if properties["indexes"] != nil {

		var_999740b5b773 := properties["indexes"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_OBJECT).UnPack(var_999740b5b773)

		if err != nil {
			panic(err)
		}

		var_999740b5b773_mapped := new(interface{})
		*var_999740b5b773_mapped = val.(interface{})

		s.Indexes = var_999740b5b773_mapped
	}
	if properties["securityConstraints"] != nil {

		var_e3d2cb25617e := properties["securityConstraints"]
		var_e3d2cb25617e_mapped := []*SecurityConstraint{}
		for _, v := range var_e3d2cb25617e.GetListValue().Values {

			var_3b18faade138 := v
			var_3b18faade138_mapped := SecurityConstraintMapperInstance.FromProperties(var_3b18faade138.GetStructValue().Fields)

			var_e3d2cb25617e_mapped = append(var_e3d2cb25617e_mapped, var_3b18faade138_mapped)
		}

		s.SecurityConstraints = var_e3d2cb25617e_mapped
	}
	if properties["title"] != nil {

		var_40fe9e479494 := properties["title"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_40fe9e479494)

		if err != nil {
			panic(err)
		}

		var_40fe9e479494_mapped := new(string)
		*var_40fe9e479494_mapped = val.(string)

		s.Title = var_40fe9e479494_mapped
	}
	if properties["description"] != nil {

		var_2a138a16978a := properties["description"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_2a138a16978a)

		if err != nil {
			panic(err)
		}

		var_2a138a16978a_mapped := new(string)
		*var_2a138a16978a_mapped = val.(string)

		s.Description = var_2a138a16978a_mapped
	}
	return s
}
