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
	if properties["description"] != nil {
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(properties["description"])

		if err != nil {
			panic(err)
		}

		s.Description = new(string)
		*s.Description = val.(string)
	}
	if properties["details"] != nil {
		val, err := types.ByResourcePropertyType(model.ResourceProperty_OBJECT).UnPack(properties["details"])

		if err != nil {
			panic(err)
		}

		s.Details = new(interface{})
		*s.Details = val.(interface{})
	}
	if properties["securityConstraints"] != nil {
		s.SecurityConstraints = []NamespaceSecurityConstraint{}
		for _, v := range properties["securityConstraints"].AsInterface().([]interface{}) {
			s.SecurityConstraints = append(s.SecurityConstraints, v.(NamespaceSecurityConstraint))
		}
	}
	return s
}

type NamespaceSecurityConstraintMapper struct {
}

func NewNamespaceSecurityConstraintMapper() *NamespaceSecurityConstraintMapper {
	return &NamespaceSecurityConstraintMapper{}
}

var NamespaceSecurityConstraintMapperInstance = NewNamespaceSecurityConstraintMapper()

func (m *NamespaceSecurityConstraintMapper) New() *NamespaceSecurityConstraint {
	return &NamespaceSecurityConstraint{}
}

func (m *NamespaceSecurityConstraintMapper) ToRecord(namespaceSecurityConstraint *NamespaceSecurityConstraint) *model.Record {
	var rec = &model.Record{}
	rec.Properties = m.ToProperties(namespaceSecurityConstraint)
	return rec
}

func (m *NamespaceSecurityConstraintMapper) FromRecord(record *model.Record) *NamespaceSecurityConstraint {
	return m.FromProperties(record.Properties)
}

func (m *NamespaceSecurityConstraintMapper) ToProperties(namespaceSecurityConstraint *NamespaceSecurityConstraint) map[string]*structpb.Value {
	var properties = make(map[string]*structpb.Value)

	namespace, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(namespaceSecurityConstraint.Namespace)
	if err != nil {
		panic(err)
	}
	properties["namespace"] = namespace

	resource, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(namespaceSecurityConstraint.Resource)
	if err != nil {
		panic(err)
	}
	properties["resource"] = resource

	property, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(namespaceSecurityConstraint.Property)
	if err != nil {
		panic(err)
	}
	properties["property"] = property

	if namespaceSecurityConstraint.PropertyValue != nil {
		propertyValue, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*namespaceSecurityConstraint.PropertyValue)
		if err != nil {
			panic(err)
		}
		properties["propertyValue"] = propertyValue
	}

	if namespaceSecurityConstraint.PropertyMode != nil {
		propertyMode, err := types.ByResourcePropertyType(model.ResourceProperty_ENUM).Pack(*namespaceSecurityConstraint.PropertyMode)
		if err != nil {
			panic(err)
		}
		properties["propertyMode"] = propertyMode
	}

	if namespaceSecurityConstraint.Operation != nil {
		operation, err := types.ByResourcePropertyType(model.ResourceProperty_ENUM).Pack(*namespaceSecurityConstraint.Operation)
		if err != nil {
			panic(err)
		}
		properties["operation"] = operation
	}

	if namespaceSecurityConstraint.RecordIds != nil {
	}

	if namespaceSecurityConstraint.Before != nil {
		before, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*namespaceSecurityConstraint.Before)
		if err != nil {
			panic(err)
		}
		properties["before"] = before
	}

	if namespaceSecurityConstraint.After != nil {
		after, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*namespaceSecurityConstraint.After)
		if err != nil {
			panic(err)
		}
		properties["after"] = after
	}

	if namespaceSecurityConstraint.Username != nil {
		username, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*namespaceSecurityConstraint.Username)
		if err != nil {
			panic(err)
		}
		properties["username"] = username
	}

	if namespaceSecurityConstraint.Role != nil {
		role, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*namespaceSecurityConstraint.Role)
		if err != nil {
			panic(err)
		}
		properties["role"] = role
	}

	permit, err := types.ByResourcePropertyType(model.ResourceProperty_ENUM).Pack(namespaceSecurityConstraint.Permit)
	if err != nil {
		panic(err)
	}
	properties["permit"] = permit

	if namespaceSecurityConstraint.LocalFlags != nil {
	}

	return properties
}

func (m *NamespaceSecurityConstraintMapper) FromProperties(properties map[string]*structpb.Value) *NamespaceSecurityConstraint {
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
		val, err := types.ByResourcePropertyType(model.ResourceProperty_ENUM).UnPack(properties["propertyMode"])

		if err != nil {
			panic(err)
		}

		s.PropertyMode = new(NamespacePropertyMode)
		*s.PropertyMode = val.(NamespacePropertyMode)
	}
	if properties["operation"] != nil {
		val, err := types.ByResourcePropertyType(model.ResourceProperty_ENUM).UnPack(properties["operation"])

		if err != nil {
			panic(err)
		}

		s.Operation = new(NamespaceOperation)
		*s.Operation = val.(NamespaceOperation)
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
		val, err := types.ByResourcePropertyType(model.ResourceProperty_ENUM).UnPack(properties["permit"])

		if err != nil {
			panic(err)
		}

		s.Permit = val.(NamespacePermit)
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
