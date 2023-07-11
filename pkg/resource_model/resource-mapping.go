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
		val, err := types.ByResourcePropertyType(model.ResourceProperty_UUID).UnPack(properties["id"])

		if err != nil {
			panic(err)
		}

		s.Id = new(uuid.UUID)
		*s.Id = val.(uuid.UUID)
	}
	if properties["version"] != nil {
		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(properties["version"])

		if err != nil {
			panic(err)
		}

		s.Version = new(int32)
		*s.Version = val.(int32)
	}
	if properties["createdBy"] != nil {
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(properties["createdBy"])

		if err != nil {
			panic(err)
		}

		s.CreatedBy = new(string)
		*s.CreatedBy = val.(string)
	}
	if properties["updatedBy"] != nil {
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(properties["updatedBy"])

		if err != nil {
			panic(err)
		}

		s.UpdatedBy = new(string)
		*s.UpdatedBy = val.(string)
	}
	if properties["createdOn"] != nil {
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(properties["createdOn"])

		if err != nil {
			panic(err)
		}

		s.CreatedOn = new(time.Time)
		*s.CreatedOn = val.(time.Time)
	}
	if properties["updatedOn"] != nil {
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(properties["updatedOn"])

		if err != nil {
			panic(err)
		}

		s.UpdatedOn = new(time.Time)
		*s.UpdatedOn = val.(time.Time)
	}
	if properties["name"] != nil {
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(properties["name"])

		if err != nil {
			panic(err)
		}

		s.Name = val.(string)
	}
	if properties["namespace"] != nil {
		s.Namespace = NamespaceMapperInstance.FromProperties(properties["namespace"].GetStructValue().Fields)
	}
	if properties["virtual"] != nil {
		val, err := types.ByResourcePropertyType(model.ResourceProperty_BOOL).UnPack(properties["virtual"])

		if err != nil {
			panic(err)
		}

		s.Virtual = val.(bool)
	}
	if properties["types"] != nil {
		val, err := types.ByResourcePropertyType(model.ResourceProperty_OBJECT).UnPack(properties["types"])

		if err != nil {
			panic(err)
		}

		s.Types = new(interface{})
		*s.Types = val.(interface{})
	}
	if properties["immutable"] != nil {
		val, err := types.ByResourcePropertyType(model.ResourceProperty_BOOL).UnPack(properties["immutable"])

		if err != nil {
			panic(err)
		}

		s.Immutable = val.(bool)
	}
	if properties["abstract"] != nil {
		val, err := types.ByResourcePropertyType(model.ResourceProperty_BOOL).UnPack(properties["abstract"])

		if err != nil {
			panic(err)
		}

		s.Abstract = val.(bool)
	}
	if properties["dataSource"] != nil {
		s.DataSource = DataSourceMapperInstance.FromProperties(properties["dataSource"].GetStructValue().Fields)
	}
	if properties["entity"] != nil {
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(properties["entity"])

		if err != nil {
			panic(err)
		}

		s.Entity = new(string)
		*s.Entity = val.(string)
	}
	if properties["catalog"] != nil {
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(properties["catalog"])

		if err != nil {
			panic(err)
		}

		s.Catalog = new(string)
		*s.Catalog = val.(string)
	}
	if properties["annotations"] != nil {
		s.Annotations = make(map[string]string)
		for k, v := range properties["annotations"].GetStructValue().Fields {
			s.Annotations[k] = v.AsInterface().(string)
		}
	}
	if properties["indexes"] != nil {
		val, err := types.ByResourcePropertyType(model.ResourceProperty_OBJECT).UnPack(properties["indexes"])

		if err != nil {
			panic(err)
		}

		s.Indexes = new(interface{})
		*s.Indexes = val.(interface{})
	}
	if properties["securityConstraints"] != nil {
		s.SecurityConstraints = []ResourceSecurityConstraint{}
		for _, v := range properties["securityConstraints"].AsInterface().([]interface{}) {
			s.SecurityConstraints = append(s.SecurityConstraints, v.(ResourceSecurityConstraint))
		}
	}
	if properties["title"] != nil {
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(properties["title"])

		if err != nil {
			panic(err)
		}

		s.Title = new(string)
		*s.Title = val.(string)
	}
	if properties["description"] != nil {
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(properties["description"])

		if err != nil {
			panic(err)
		}

		s.Description = new(string)
		*s.Description = val.(string)
	}
	return s
}

type ResourceSecurityConstraintMapper struct {
}

func NewResourceSecurityConstraintMapper() *ResourceSecurityConstraintMapper {
	return &ResourceSecurityConstraintMapper{}
}

var ResourceSecurityConstraintMapperInstance = NewResourceSecurityConstraintMapper()

func (m *ResourceSecurityConstraintMapper) New() *ResourceSecurityConstraint {
	return &ResourceSecurityConstraint{}
}

func (m *ResourceSecurityConstraintMapper) ToRecord(resourceSecurityConstraint *ResourceSecurityConstraint) *model.Record {
	var rec = &model.Record{}
	rec.Properties = m.ToProperties(resourceSecurityConstraint)
	return rec
}

func (m *ResourceSecurityConstraintMapper) FromRecord(record *model.Record) *ResourceSecurityConstraint {
	return m.FromProperties(record.Properties)
}

func (m *ResourceSecurityConstraintMapper) ToProperties(resourceSecurityConstraint *ResourceSecurityConstraint) map[string]*structpb.Value {
	var properties = make(map[string]*structpb.Value)

	namespace, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(resourceSecurityConstraint.Namespace)
	if err != nil {
		panic(err)
	}
	properties["namespace"] = namespace

	resource, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(resourceSecurityConstraint.Resource)
	if err != nil {
		panic(err)
	}
	properties["resource"] = resource

	property, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(resourceSecurityConstraint.Property)
	if err != nil {
		panic(err)
	}
	properties["property"] = property

	if resourceSecurityConstraint.PropertyValue != nil {
		propertyValue, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*resourceSecurityConstraint.PropertyValue)
		if err != nil {
			panic(err)
		}
		properties["propertyValue"] = propertyValue
	}

	if resourceSecurityConstraint.PropertyMode != nil {
		propertyMode, err := types.ByResourcePropertyType(model.ResourceProperty_ENUM).Pack(*resourceSecurityConstraint.PropertyMode)
		if err != nil {
			panic(err)
		}
		properties["propertyMode"] = propertyMode
	}

	if resourceSecurityConstraint.Operation != nil {
		operation, err := types.ByResourcePropertyType(model.ResourceProperty_ENUM).Pack(*resourceSecurityConstraint.Operation)
		if err != nil {
			panic(err)
		}
		properties["operation"] = operation
	}

	if resourceSecurityConstraint.RecordIds != nil {
	}

	if resourceSecurityConstraint.Before != nil {
		before, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*resourceSecurityConstraint.Before)
		if err != nil {
			panic(err)
		}
		properties["before"] = before
	}

	if resourceSecurityConstraint.After != nil {
		after, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*resourceSecurityConstraint.After)
		if err != nil {
			panic(err)
		}
		properties["after"] = after
	}

	if resourceSecurityConstraint.Username != nil {
		username, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*resourceSecurityConstraint.Username)
		if err != nil {
			panic(err)
		}
		properties["username"] = username
	}

	if resourceSecurityConstraint.Role != nil {
		role, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*resourceSecurityConstraint.Role)
		if err != nil {
			panic(err)
		}
		properties["role"] = role
	}

	permit, err := types.ByResourcePropertyType(model.ResourceProperty_ENUM).Pack(resourceSecurityConstraint.Permit)
	if err != nil {
		panic(err)
	}
	properties["permit"] = permit

	if resourceSecurityConstraint.LocalFlags != nil {
	}

	return properties
}

func (m *ResourceSecurityConstraintMapper) FromProperties(properties map[string]*structpb.Value) *ResourceSecurityConstraint {
	var s = m.New()
	if properties["namespace"] != nil {
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(properties["namespace"])

		if err != nil {
			panic(err)
		}

		s.Namespace = val.(string)
	}
	if properties["resource"] != nil {
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(properties["resource"])

		if err != nil {
			panic(err)
		}

		s.Resource = val.(string)
	}
	if properties["property"] != nil {
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(properties["property"])

		if err != nil {
			panic(err)
		}

		s.Property = val.(string)
	}
	if properties["propertyValue"] != nil {
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(properties["propertyValue"])

		if err != nil {
			panic(err)
		}

		s.PropertyValue = new(string)
		*s.PropertyValue = val.(string)
	}
	if properties["propertyMode"] != nil {
		s.PropertyMode = new(ResourcePropertyMode)
		*s.PropertyMode = (ResourcePropertyMode)(properties["propertyMode"].GetStringValue())
	}
	if properties["operation"] != nil {
		s.Operation = new(ResourceOperation)
		*s.Operation = (ResourceOperation)(properties["operation"].GetStringValue())
	}
	if properties["recordIds"] != nil {
		s.RecordIds = []string{}
		for _, v := range properties["recordIds"].AsInterface().([]interface{}) {
			s.RecordIds = append(s.RecordIds, v.(string))
		}
	}
	if properties["before"] != nil {
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(properties["before"])

		if err != nil {
			panic(err)
		}

		s.Before = new(time.Time)
		*s.Before = val.(time.Time)
	}
	if properties["after"] != nil {
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(properties["after"])

		if err != nil {
			panic(err)
		}

		s.After = new(time.Time)
		*s.After = val.(time.Time)
	}
	if properties["username"] != nil {
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(properties["username"])

		if err != nil {
			panic(err)
		}

		s.Username = new(string)
		*s.Username = val.(string)
	}
	if properties["role"] != nil {
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(properties["role"])

		if err != nil {
			panic(err)
		}

		s.Role = new(string)
		*s.Role = val.(string)
	}
	if properties["permit"] != nil {
		s.Permit = (ResourcePermit)(properties["permit"].GetStringValue())
	}
	if properties["localFlags"] != nil {
		val, err := types.ByResourcePropertyType(model.ResourceProperty_OBJECT).UnPack(properties["localFlags"])

		if err != nil {
			panic(err)
		}

		s.LocalFlags = new(interface{})
		*s.LocalFlags = val.(interface{})
	}
	return s
}
