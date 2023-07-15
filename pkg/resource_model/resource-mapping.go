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

		var_8c479c3540fc := properties["id"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_UUID).UnPack(var_8c479c3540fc)

		if err != nil {
			panic(err)
		}

		var_8c479c3540fc_mapped := new(uuid.UUID)
		*var_8c479c3540fc_mapped = val.(uuid.UUID)

		s.Id = var_8c479c3540fc_mapped
	}
	if properties["version"] != nil {

		var_da4c59fbc217 := properties["version"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(var_da4c59fbc217)

		if err != nil {
			panic(err)
		}

		var_da4c59fbc217_mapped := new(int32)
		*var_da4c59fbc217_mapped = val.(int32)

		s.Version = var_da4c59fbc217_mapped
	}
	if properties["createdBy"] != nil {

		var_08b8940a9bee := properties["createdBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_08b8940a9bee)

		if err != nil {
			panic(err)
		}

		var_08b8940a9bee_mapped := new(string)
		*var_08b8940a9bee_mapped = val.(string)

		s.CreatedBy = var_08b8940a9bee_mapped
	}
	if properties["updatedBy"] != nil {

		var_e96640e7439d := properties["updatedBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_e96640e7439d)

		if err != nil {
			panic(err)
		}

		var_e96640e7439d_mapped := new(string)
		*var_e96640e7439d_mapped = val.(string)

		s.UpdatedBy = var_e96640e7439d_mapped
	}
	if properties["createdOn"] != nil {

		var_bd006113dfaa := properties["createdOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_bd006113dfaa)

		if err != nil {
			panic(err)
		}

		var_bd006113dfaa_mapped := new(time.Time)
		*var_bd006113dfaa_mapped = val.(time.Time)

		s.CreatedOn = var_bd006113dfaa_mapped
	}
	if properties["updatedOn"] != nil {

		var_43a0947174dc := properties["updatedOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_43a0947174dc)

		if err != nil {
			panic(err)
		}

		var_43a0947174dc_mapped := new(time.Time)
		*var_43a0947174dc_mapped = val.(time.Time)

		s.UpdatedOn = var_43a0947174dc_mapped
	}
	if properties["name"] != nil {

		var_d325327aae0a := properties["name"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_d325327aae0a)

		if err != nil {
			panic(err)
		}

		var_d325327aae0a_mapped := val.(string)

		s.Name = var_d325327aae0a_mapped
	}
	if properties["namespace"] != nil {

		var_9497d1a48faa := properties["namespace"]
		var_9497d1a48faa_mapped := NamespaceMapperInstance.FromProperties(var_9497d1a48faa.GetStructValue().Fields)

		s.Namespace = var_9497d1a48faa_mapped
	}
	if properties["virtual"] != nil {

		var_e7aafd0723b8 := properties["virtual"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_BOOL).UnPack(var_e7aafd0723b8)

		if err != nil {
			panic(err)
		}

		var_e7aafd0723b8_mapped := val.(bool)

		s.Virtual = var_e7aafd0723b8_mapped
	}
	if properties["types"] != nil {

		var_fa61e8109e3f := properties["types"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_OBJECT).UnPack(var_fa61e8109e3f)

		if err != nil {
			panic(err)
		}

		var_fa61e8109e3f_mapped := new(unstructured.Unstructured)
		*var_fa61e8109e3f_mapped = val.(unstructured.Unstructured)

		s.Types = var_fa61e8109e3f_mapped
	}
	if properties["immutable"] != nil {

		var_8c3ea3ff6e60 := properties["immutable"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_BOOL).UnPack(var_8c3ea3ff6e60)

		if err != nil {
			panic(err)
		}

		var_8c3ea3ff6e60_mapped := val.(bool)

		s.Immutable = var_8c3ea3ff6e60_mapped
	}
	if properties["abstract"] != nil {

		var_4bd676ac0124 := properties["abstract"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_BOOL).UnPack(var_4bd676ac0124)

		if err != nil {
			panic(err)
		}

		var_4bd676ac0124_mapped := val.(bool)

		s.Abstract = var_4bd676ac0124_mapped
	}
	if properties["dataSource"] != nil {

		var_2ecee6ed9259 := properties["dataSource"]
		var_2ecee6ed9259_mapped := DataSourceMapperInstance.FromProperties(var_2ecee6ed9259.GetStructValue().Fields)

		s.DataSource = var_2ecee6ed9259_mapped
	}
	if properties["entity"] != nil {

		var_ee09ffdba784 := properties["entity"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_ee09ffdba784)

		if err != nil {
			panic(err)
		}

		var_ee09ffdba784_mapped := new(string)
		*var_ee09ffdba784_mapped = val.(string)

		s.Entity = var_ee09ffdba784_mapped
	}
	if properties["catalog"] != nil {

		var_2a3621acd1d2 := properties["catalog"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_2a3621acd1d2)

		if err != nil {
			panic(err)
		}

		var_2a3621acd1d2_mapped := new(string)
		*var_2a3621acd1d2_mapped = val.(string)

		s.Catalog = var_2a3621acd1d2_mapped
	}
	if properties["annotations"] != nil {

		var_73a191de9642 := properties["annotations"]
		var_73a191de9642_mapped := make(map[string]string)
		for k, v := range var_73a191de9642.GetStructValue().Fields {

			var_384278501475 := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_384278501475)

			if err != nil {
				panic(err)
			}

			var_384278501475_mapped := val.(string)

			var_73a191de9642_mapped[k] = var_384278501475_mapped
		}

		s.Annotations = var_73a191de9642_mapped
	}
	if properties["indexes"] != nil {

		var_d5cc08b1c750 := properties["indexes"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_OBJECT).UnPack(var_d5cc08b1c750)

		if err != nil {
			panic(err)
		}

		var_d5cc08b1c750_mapped := new(unstructured.Unstructured)
		*var_d5cc08b1c750_mapped = val.(unstructured.Unstructured)

		s.Indexes = var_d5cc08b1c750_mapped
	}
	if properties["securityConstraints"] != nil {

		var_3a377c6d37bf := properties["securityConstraints"]
		var_3a377c6d37bf_mapped := []*SecurityConstraint{}
		for _, v := range var_3a377c6d37bf.GetListValue().Values {

			var_326785b235bb := v
			var_326785b235bb_mapped := SecurityConstraintMapperInstance.FromProperties(var_326785b235bb.GetStructValue().Fields)

			var_3a377c6d37bf_mapped = append(var_3a377c6d37bf_mapped, var_326785b235bb_mapped)
		}

		s.SecurityConstraints = var_3a377c6d37bf_mapped
	}
	if properties["title"] != nil {

		var_1fdd7089d64e := properties["title"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_1fdd7089d64e)

		if err != nil {
			panic(err)
		}

		var_1fdd7089d64e_mapped := new(string)
		*var_1fdd7089d64e_mapped = val.(string)

		s.Title = var_1fdd7089d64e_mapped
	}
	if properties["description"] != nil {

		var_7c6879cad14c := properties["description"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_7c6879cad14c)

		if err != nil {
			panic(err)
		}

		var_7c6879cad14c_mapped := new(string)
		*var_7c6879cad14c_mapped = val.(string)

		s.Description = var_7c6879cad14c_mapped
	}
	return s
}
