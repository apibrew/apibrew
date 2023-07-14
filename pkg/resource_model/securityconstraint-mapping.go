package resource_model

import (
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/types"
	"google.golang.org/protobuf/types/known/structpb"
)

import "github.com/google/uuid"
import "time"

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

		var_fef450e1581a := properties["id"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_UUID).UnPack(var_fef450e1581a)

		if err != nil {
			panic(err)
		}

		var_fef450e1581a_mapped := new(uuid.UUID)
		*var_fef450e1581a_mapped = val.(uuid.UUID)

		s.Id = var_fef450e1581a_mapped
	}
	if properties["version"] != nil {

		var_1103e8642481 := properties["version"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(var_1103e8642481)

		if err != nil {
			panic(err)
		}

		var_1103e8642481_mapped := new(int32)
		*var_1103e8642481_mapped = val.(int32)

		s.Version = var_1103e8642481_mapped
	}
	if properties["createdBy"] != nil {

		var_1de21dc75747 := properties["createdBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_1de21dc75747)

		if err != nil {
			panic(err)
		}

		var_1de21dc75747_mapped := new(string)
		*var_1de21dc75747_mapped = val.(string)

		s.CreatedBy = var_1de21dc75747_mapped
	}
	if properties["updatedBy"] != nil {

		var_60fca6ee532e := properties["updatedBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_60fca6ee532e)

		if err != nil {
			panic(err)
		}

		var_60fca6ee532e_mapped := new(string)
		*var_60fca6ee532e_mapped = val.(string)

		s.UpdatedBy = var_60fca6ee532e_mapped
	}
	if properties["createdOn"] != nil {

		var_be6632ba3c83 := properties["createdOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_be6632ba3c83)

		if err != nil {
			panic(err)
		}

		var_be6632ba3c83_mapped := new(time.Time)
		*var_be6632ba3c83_mapped = val.(time.Time)

		s.CreatedOn = var_be6632ba3c83_mapped
	}
	if properties["updatedOn"] != nil {

		var_73a4e3a486ca := properties["updatedOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_73a4e3a486ca)

		if err != nil {
			panic(err)
		}

		var_73a4e3a486ca_mapped := new(time.Time)
		*var_73a4e3a486ca_mapped = val.(time.Time)

		s.UpdatedOn = var_73a4e3a486ca_mapped
	}
	if properties["namespace"] != nil {

		var_8c414e5c395f := properties["namespace"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_8c414e5c395f)

		if err != nil {
			panic(err)
		}

		var_8c414e5c395f_mapped := val.(string)

		s.Namespace = var_8c414e5c395f_mapped
	}
	if properties["resource"] != nil {

		var_98c48d5fd7d4 := properties["resource"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_98c48d5fd7d4)

		if err != nil {
			panic(err)
		}

		var_98c48d5fd7d4_mapped := val.(string)

		s.Resource = var_98c48d5fd7d4_mapped
	}
	if properties["property"] != nil {

		var_f1a24a6c10a1 := properties["property"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_f1a24a6c10a1)

		if err != nil {
			panic(err)
		}

		var_f1a24a6c10a1_mapped := val.(string)

		s.Property = var_f1a24a6c10a1_mapped
	}
	if properties["propertyValue"] != nil {

		var_b625adb6675d := properties["propertyValue"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_b625adb6675d)

		if err != nil {
			panic(err)
		}

		var_b625adb6675d_mapped := new(string)
		*var_b625adb6675d_mapped = val.(string)

		s.PropertyValue = var_b625adb6675d_mapped
	}
	if properties["propertyMode"] != nil {

		var_d62fcc6b8bbf := properties["propertyMode"]
		var_d62fcc6b8bbf_mapped := new(SecurityConstraintPropertyMode)
		*var_d62fcc6b8bbf_mapped = (SecurityConstraintPropertyMode)(var_d62fcc6b8bbf.GetStringValue())

		s.PropertyMode = var_d62fcc6b8bbf_mapped
	}
	if properties["operation"] != nil {

		var_a4df25afb0ce := properties["operation"]
		var_a4df25afb0ce_mapped := (SecurityConstraintOperation)(var_a4df25afb0ce.GetStringValue())

		s.Operation = var_a4df25afb0ce_mapped
	}
	if properties["recordIds"] != nil {

		var_79ee09d233f1 := properties["recordIds"]
		var_79ee09d233f1_mapped := []string{}
		for _, v := range var_79ee09d233f1.GetListValue().Values {

			var_3614f2c594f8 := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_3614f2c594f8)

			if err != nil {
				panic(err)
			}

			var_3614f2c594f8_mapped := val.(string)

			var_79ee09d233f1_mapped = append(var_79ee09d233f1_mapped, var_3614f2c594f8_mapped)
		}

		s.RecordIds = var_79ee09d233f1_mapped
	}
	if properties["before"] != nil {

		var_eaa7a97ad8b3 := properties["before"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_eaa7a97ad8b3)

		if err != nil {
			panic(err)
		}

		var_eaa7a97ad8b3_mapped := new(time.Time)
		*var_eaa7a97ad8b3_mapped = val.(time.Time)

		s.Before = var_eaa7a97ad8b3_mapped
	}
	if properties["after"] != nil {

		var_6abaea4046f9 := properties["after"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_6abaea4046f9)

		if err != nil {
			panic(err)
		}

		var_6abaea4046f9_mapped := new(time.Time)
		*var_6abaea4046f9_mapped = val.(time.Time)

		s.After = var_6abaea4046f9_mapped
	}
	if properties["username"] != nil {

		var_170a333b5f6f := properties["username"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_170a333b5f6f)

		if err != nil {
			panic(err)
		}

		var_170a333b5f6f_mapped := new(string)
		*var_170a333b5f6f_mapped = val.(string)

		s.Username = var_170a333b5f6f_mapped
	}
	if properties["role"] != nil {

		var_b19dcc88a507 := properties["role"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_b19dcc88a507)

		if err != nil {
			panic(err)
		}

		var_b19dcc88a507_mapped := new(string)
		*var_b19dcc88a507_mapped = val.(string)

		s.Role = var_b19dcc88a507_mapped
	}
	if properties["permit"] != nil {

		var_f77939766c2e := properties["permit"]
		var_f77939766c2e_mapped := (SecurityConstraintPermit)(var_f77939766c2e.GetStringValue())

		s.Permit = var_f77939766c2e_mapped
	}
	if properties["localFlags"] != nil {

		var_38ecb40b5057 := properties["localFlags"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_OBJECT).UnPack(var_38ecb40b5057)

		if err != nil {
			panic(err)
		}

		var_38ecb40b5057_mapped := new(interface{})
		*var_38ecb40b5057_mapped = val.(interface{})

		s.LocalFlags = var_38ecb40b5057_mapped
	}
	return s
}
