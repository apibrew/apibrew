package resource_model

import (
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/types"
	"google.golang.org/protobuf/types/known/structpb"
)

import "github.com/google/uuid"
import "time"
import "github.com/apibrew/apibrew/pkg/formats/unstructured"

type SecurityConstraintMapper struct {
}

func NewSecurityConstraintMapper() *SecurityConstraintMapper {
	return &SecurityConstraintMapper{}
}

var SecurityConstraintMapperInstance = NewSecurityConstraintMapper()

func (m *SecurityConstraintMapper) New() *SecurityConstraint {
	return &SecurityConstraint{}
}

func (m *SecurityConstraintMapper) ToRecord(securityConstraint *SecurityConstraint) *model.Record {
	var rec = &model.Record{}
	rec.Properties = m.ToProperties(securityConstraint)
	return rec
}

func (m *SecurityConstraintMapper) FromRecord(record *model.Record) *SecurityConstraint {
	return m.FromProperties(record.Properties)
}

func (m *SecurityConstraintMapper) ToProperties(securityConstraint *SecurityConstraint) map[string]*structpb.Value {
	var properties = make(map[string]*structpb.Value)

	if securityConstraint.Id != nil {
		id, err := types.ByResourcePropertyType(model.ResourceProperty_UUID).Pack(*securityConstraint.Id)
		if err != nil {
			panic(err)
		}
		properties["id"] = id
	}

	if securityConstraint.Version != nil {
		version, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).Pack(*securityConstraint.Version)
		if err != nil {
			panic(err)
		}
		properties["version"] = version
	}

	if securityConstraint.CreatedBy != nil {
		createdBy, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*securityConstraint.CreatedBy)
		if err != nil {
			panic(err)
		}
		properties["createdBy"] = createdBy
	}

	if securityConstraint.UpdatedBy != nil {
		updatedBy, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*securityConstraint.UpdatedBy)
		if err != nil {
			panic(err)
		}
		properties["updatedBy"] = updatedBy
	}

	if securityConstraint.CreatedOn != nil {
		createdOn, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*securityConstraint.CreatedOn)
		if err != nil {
			panic(err)
		}
		properties["createdOn"] = createdOn
	}

	if securityConstraint.UpdatedOn != nil {
		updatedOn, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*securityConstraint.UpdatedOn)
		if err != nil {
			panic(err)
		}
		properties["updatedOn"] = updatedOn
	}

	namespace, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(securityConstraint.Namespace)
	if err != nil {
		panic(err)
	}
	properties["namespace"] = namespace

	resource, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(securityConstraint.Resource)
	if err != nil {
		panic(err)
	}
	properties["resource"] = resource

	property, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(securityConstraint.Property)
	if err != nil {
		panic(err)
	}
	properties["property"] = property

	if securityConstraint.PropertyValue != nil {
		propertyValue, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*securityConstraint.PropertyValue)
		if err != nil {
			panic(err)
		}
		properties["propertyValue"] = propertyValue
	}

	if securityConstraint.PropertyMode != nil {
		propertyMode, err := types.ByResourcePropertyType(model.ResourceProperty_ENUM).Pack(*securityConstraint.PropertyMode)
		if err != nil {
			panic(err)
		}
		properties["propertyMode"] = propertyMode
	}

	operation, err := types.ByResourcePropertyType(model.ResourceProperty_ENUM).Pack(securityConstraint.Operation)
	if err != nil {
		panic(err)
	}
	properties["operation"] = operation

	if securityConstraint.RecordIds != nil {
	}

	if securityConstraint.Before != nil {
		before, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*securityConstraint.Before)
		if err != nil {
			panic(err)
		}
		properties["before"] = before
	}

	if securityConstraint.After != nil {
		after, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*securityConstraint.After)
		if err != nil {
			panic(err)
		}
		properties["after"] = after
	}

	if securityConstraint.Username != nil {
		username, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*securityConstraint.Username)
		if err != nil {
			panic(err)
		}
		properties["username"] = username
	}

	if securityConstraint.Role != nil {
		role, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*securityConstraint.Role)
		if err != nil {
			panic(err)
		}
		properties["role"] = role
	}

	permit, err := types.ByResourcePropertyType(model.ResourceProperty_ENUM).Pack(securityConstraint.Permit)
	if err != nil {
		panic(err)
	}
	properties["permit"] = permit

	if securityConstraint.LocalFlags != nil {
	}

	return properties
}

func (m *SecurityConstraintMapper) FromProperties(properties map[string]*structpb.Value) *SecurityConstraint {
	var s = m.New()
	if properties["id"] != nil {

		var_bc2c2cffdc9b := properties["id"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_UUID).UnPack(var_bc2c2cffdc9b)

		if err != nil {
			panic(err)
		}

		var_bc2c2cffdc9b_mapped := new(uuid.UUID)
		*var_bc2c2cffdc9b_mapped = val.(uuid.UUID)

		s.Id = var_bc2c2cffdc9b_mapped
	}
	if properties["version"] != nil {

		var_14d43f3f169a := properties["version"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(var_14d43f3f169a)

		if err != nil {
			panic(err)
		}

		var_14d43f3f169a_mapped := new(int32)
		*var_14d43f3f169a_mapped = val.(int32)

		s.Version = var_14d43f3f169a_mapped
	}
	if properties["createdBy"] != nil {

		var_8a0861a79915 := properties["createdBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_8a0861a79915)

		if err != nil {
			panic(err)
		}

		var_8a0861a79915_mapped := new(string)
		*var_8a0861a79915_mapped = val.(string)

		s.CreatedBy = var_8a0861a79915_mapped
	}
	if properties["updatedBy"] != nil {

		var_49f187559f8a := properties["updatedBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_49f187559f8a)

		if err != nil {
			panic(err)
		}

		var_49f187559f8a_mapped := new(string)
		*var_49f187559f8a_mapped = val.(string)

		s.UpdatedBy = var_49f187559f8a_mapped
	}
	if properties["createdOn"] != nil {

		var_102d45ee2d1f := properties["createdOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_102d45ee2d1f)

		if err != nil {
			panic(err)
		}

		var_102d45ee2d1f_mapped := new(time.Time)
		*var_102d45ee2d1f_mapped = val.(time.Time)

		s.CreatedOn = var_102d45ee2d1f_mapped
	}
	if properties["updatedOn"] != nil {

		var_62964896d580 := properties["updatedOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_62964896d580)

		if err != nil {
			panic(err)
		}

		var_62964896d580_mapped := new(time.Time)
		*var_62964896d580_mapped = val.(time.Time)

		s.UpdatedOn = var_62964896d580_mapped
	}
	if properties["namespace"] != nil {

		var_6a8425a7d4ad := properties["namespace"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_6a8425a7d4ad)

		if err != nil {
			panic(err)
		}

		var_6a8425a7d4ad_mapped := val.(string)

		s.Namespace = var_6a8425a7d4ad_mapped
	}
	if properties["resource"] != nil {

		var_212ceb74d981 := properties["resource"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_212ceb74d981)

		if err != nil {
			panic(err)
		}

		var_212ceb74d981_mapped := val.(string)

		s.Resource = var_212ceb74d981_mapped
	}
	if properties["property"] != nil {

		var_2f228d53644e := properties["property"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_2f228d53644e)

		if err != nil {
			panic(err)
		}

		var_2f228d53644e_mapped := val.(string)

		s.Property = var_2f228d53644e_mapped
	}
	if properties["propertyValue"] != nil {

		var_d0f55e9b19f9 := properties["propertyValue"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_d0f55e9b19f9)

		if err != nil {
			panic(err)
		}

		var_d0f55e9b19f9_mapped := new(string)
		*var_d0f55e9b19f9_mapped = val.(string)

		s.PropertyValue = var_d0f55e9b19f9_mapped
	}
	if properties["propertyMode"] != nil {

		var_50f35ecaa171 := properties["propertyMode"]
		var_50f35ecaa171_mapped := new(SecurityConstraintPropertyMode)
		*var_50f35ecaa171_mapped = (SecurityConstraintPropertyMode)(var_50f35ecaa171.GetStringValue())

		s.PropertyMode = var_50f35ecaa171_mapped
	}
	if properties["operation"] != nil {

		var_c734f5155228 := properties["operation"]
		var_c734f5155228_mapped := (SecurityConstraintOperation)(var_c734f5155228.GetStringValue())

		s.Operation = var_c734f5155228_mapped
	}
	if properties["recordIds"] != nil {

		var_b2590b3bab1a := properties["recordIds"]
		var_b2590b3bab1a_mapped := []string{}
		for _, v := range var_b2590b3bab1a.GetListValue().Values {

			var_03f7b9ee5ee1 := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_03f7b9ee5ee1)

			if err != nil {
				panic(err)
			}

			var_03f7b9ee5ee1_mapped := val.(string)

			var_b2590b3bab1a_mapped = append(var_b2590b3bab1a_mapped, var_03f7b9ee5ee1_mapped)
		}

		s.RecordIds = var_b2590b3bab1a_mapped
	}
	if properties["before"] != nil {

		var_9dc24b48f427 := properties["before"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_9dc24b48f427)

		if err != nil {
			panic(err)
		}

		var_9dc24b48f427_mapped := new(time.Time)
		*var_9dc24b48f427_mapped = val.(time.Time)

		s.Before = var_9dc24b48f427_mapped
	}
	if properties["after"] != nil {

		var_8d7f21b029be := properties["after"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_8d7f21b029be)

		if err != nil {
			panic(err)
		}

		var_8d7f21b029be_mapped := new(time.Time)
		*var_8d7f21b029be_mapped = val.(time.Time)

		s.After = var_8d7f21b029be_mapped
	}
	if properties["username"] != nil {

		var_a033c4df36a7 := properties["username"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_a033c4df36a7)

		if err != nil {
			panic(err)
		}

		var_a033c4df36a7_mapped := new(string)
		*var_a033c4df36a7_mapped = val.(string)

		s.Username = var_a033c4df36a7_mapped
	}
	if properties["role"] != nil {

		var_6b17fce3d9e1 := properties["role"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_6b17fce3d9e1)

		if err != nil {
			panic(err)
		}

		var_6b17fce3d9e1_mapped := new(string)
		*var_6b17fce3d9e1_mapped = val.(string)

		s.Role = var_6b17fce3d9e1_mapped
	}
	if properties["permit"] != nil {

		var_7d68c76238ce := properties["permit"]
		var_7d68c76238ce_mapped := (SecurityConstraintPermit)(var_7d68c76238ce.GetStringValue())

		s.Permit = var_7d68c76238ce_mapped
	}
	if properties["localFlags"] != nil {

		var_4420e92dcd54 := properties["localFlags"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_OBJECT).UnPack(var_4420e92dcd54)

		if err != nil {
			panic(err)
		}

		var_4420e92dcd54_mapped := new(unstructured.Unstructured)
		*var_4420e92dcd54_mapped = val.(unstructured.Unstructured)

		s.LocalFlags = var_4420e92dcd54_mapped
	}
	return s
}
