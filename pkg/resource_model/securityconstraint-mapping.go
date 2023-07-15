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

		var_ef365c2e8362 := properties["id"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_UUID).UnPack(var_ef365c2e8362)

		if err != nil {
			panic(err)
		}

		var_ef365c2e8362_mapped := new(uuid.UUID)
		*var_ef365c2e8362_mapped = val.(uuid.UUID)

		s.Id = var_ef365c2e8362_mapped
	}
	if properties["version"] != nil {

		var_f35d3114ab48 := properties["version"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(var_f35d3114ab48)

		if err != nil {
			panic(err)
		}

		var_f35d3114ab48_mapped := new(int32)
		*var_f35d3114ab48_mapped = val.(int32)

		s.Version = var_f35d3114ab48_mapped
	}
	if properties["createdBy"] != nil {

		var_37fbd3380380 := properties["createdBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_37fbd3380380)

		if err != nil {
			panic(err)
		}

		var_37fbd3380380_mapped := new(string)
		*var_37fbd3380380_mapped = val.(string)

		s.CreatedBy = var_37fbd3380380_mapped
	}
	if properties["updatedBy"] != nil {

		var_b4585a9a081d := properties["updatedBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_b4585a9a081d)

		if err != nil {
			panic(err)
		}

		var_b4585a9a081d_mapped := new(string)
		*var_b4585a9a081d_mapped = val.(string)

		s.UpdatedBy = var_b4585a9a081d_mapped
	}
	if properties["createdOn"] != nil {

		var_8caebb055c78 := properties["createdOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_8caebb055c78)

		if err != nil {
			panic(err)
		}

		var_8caebb055c78_mapped := new(time.Time)
		*var_8caebb055c78_mapped = val.(time.Time)

		s.CreatedOn = var_8caebb055c78_mapped
	}
	if properties["updatedOn"] != nil {

		var_36e3727d4bee := properties["updatedOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_36e3727d4bee)

		if err != nil {
			panic(err)
		}

		var_36e3727d4bee_mapped := new(time.Time)
		*var_36e3727d4bee_mapped = val.(time.Time)

		s.UpdatedOn = var_36e3727d4bee_mapped
	}
	if properties["namespace"] != nil {

		var_1808df546e3e := properties["namespace"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_1808df546e3e)

		if err != nil {
			panic(err)
		}

		var_1808df546e3e_mapped := val.(string)

		s.Namespace = var_1808df546e3e_mapped
	}
	if properties["resource"] != nil {

		var_a88bcce217e6 := properties["resource"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_a88bcce217e6)

		if err != nil {
			panic(err)
		}

		var_a88bcce217e6_mapped := val.(string)

		s.Resource = var_a88bcce217e6_mapped
	}
	if properties["property"] != nil {

		var_6872efd4b872 := properties["property"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_6872efd4b872)

		if err != nil {
			panic(err)
		}

		var_6872efd4b872_mapped := val.(string)

		s.Property = var_6872efd4b872_mapped
	}
	if properties["propertyValue"] != nil {

		var_4a7053347b75 := properties["propertyValue"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_4a7053347b75)

		if err != nil {
			panic(err)
		}

		var_4a7053347b75_mapped := new(string)
		*var_4a7053347b75_mapped = val.(string)

		s.PropertyValue = var_4a7053347b75_mapped
	}
	if properties["propertyMode"] != nil {

		var_50be0cb9e046 := properties["propertyMode"]
		var_50be0cb9e046_mapped := new(SecurityConstraintPropertyMode)
		*var_50be0cb9e046_mapped = (SecurityConstraintPropertyMode)(var_50be0cb9e046.GetStringValue())

		s.PropertyMode = var_50be0cb9e046_mapped
	}
	if properties["operation"] != nil {

		var_363967d16b15 := properties["operation"]
		var_363967d16b15_mapped := (SecurityConstraintOperation)(var_363967d16b15.GetStringValue())

		s.Operation = var_363967d16b15_mapped
	}
	if properties["recordIds"] != nil {

		var_ad5f68cafa8e := properties["recordIds"]
		var_ad5f68cafa8e_mapped := []string{}
		for _, v := range var_ad5f68cafa8e.GetListValue().Values {

			var_77f8a1d19ba5 := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_77f8a1d19ba5)

			if err != nil {
				panic(err)
			}

			var_77f8a1d19ba5_mapped := val.(string)

			var_ad5f68cafa8e_mapped = append(var_ad5f68cafa8e_mapped, var_77f8a1d19ba5_mapped)
		}

		s.RecordIds = var_ad5f68cafa8e_mapped
	}
	if properties["before"] != nil {

		var_b7bc4c65b185 := properties["before"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_b7bc4c65b185)

		if err != nil {
			panic(err)
		}

		var_b7bc4c65b185_mapped := new(time.Time)
		*var_b7bc4c65b185_mapped = val.(time.Time)

		s.Before = var_b7bc4c65b185_mapped
	}
	if properties["after"] != nil {

		var_76e367813924 := properties["after"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_76e367813924)

		if err != nil {
			panic(err)
		}

		var_76e367813924_mapped := new(time.Time)
		*var_76e367813924_mapped = val.(time.Time)

		s.After = var_76e367813924_mapped
	}
	if properties["username"] != nil {

		var_f9368aa92683 := properties["username"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_f9368aa92683)

		if err != nil {
			panic(err)
		}

		var_f9368aa92683_mapped := new(string)
		*var_f9368aa92683_mapped = val.(string)

		s.Username = var_f9368aa92683_mapped
	}
	if properties["role"] != nil {

		var_856a558e27fe := properties["role"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_856a558e27fe)

		if err != nil {
			panic(err)
		}

		var_856a558e27fe_mapped := new(string)
		*var_856a558e27fe_mapped = val.(string)

		s.Role = var_856a558e27fe_mapped
	}
	if properties["permit"] != nil {

		var_c8acd74f63d2 := properties["permit"]
		var_c8acd74f63d2_mapped := (SecurityConstraintPermit)(var_c8acd74f63d2.GetStringValue())

		s.Permit = var_c8acd74f63d2_mapped
	}
	if properties["localFlags"] != nil {

		var_e8a18dbce759 := properties["localFlags"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_OBJECT).UnPack(var_e8a18dbce759)

		if err != nil {
			panic(err)
		}

		var_e8a18dbce759_mapped := new(unstructured.Unstructured)
		*var_e8a18dbce759_mapped = val.(unstructured.Unstructured)

		s.LocalFlags = var_e8a18dbce759_mapped
	}
	return s
}
