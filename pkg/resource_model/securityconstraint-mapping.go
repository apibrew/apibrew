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

	version, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).Pack(securityConstraint.Version)
	if err != nil {
		panic(err)
	}
	properties["version"] = version

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

		var_f962884e83bf := properties["id"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_UUID).UnPack(var_f962884e83bf)

		if err != nil {
			panic(err)
		}

		var_f962884e83bf_mapped := new(uuid.UUID)
		*var_f962884e83bf_mapped = val.(uuid.UUID)

		s.Id = var_f962884e83bf_mapped
	}
	if properties["version"] != nil {

		var_a53bba0b80d8 := properties["version"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(var_a53bba0b80d8)

		if err != nil {
			panic(err)
		}

		var_a53bba0b80d8_mapped := val.(int32)

		s.Version = var_a53bba0b80d8_mapped
	}
	if properties["createdBy"] != nil {

		var_9d72468db589 := properties["createdBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_9d72468db589)

		if err != nil {
			panic(err)
		}

		var_9d72468db589_mapped := new(string)
		*var_9d72468db589_mapped = val.(string)

		s.CreatedBy = var_9d72468db589_mapped
	}
	if properties["updatedBy"] != nil {

		var_91ad92332043 := properties["updatedBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_91ad92332043)

		if err != nil {
			panic(err)
		}

		var_91ad92332043_mapped := new(string)
		*var_91ad92332043_mapped = val.(string)

		s.UpdatedBy = var_91ad92332043_mapped
	}
	if properties["createdOn"] != nil {

		var_e271a0d1c1c5 := properties["createdOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_e271a0d1c1c5)

		if err != nil {
			panic(err)
		}

		var_e271a0d1c1c5_mapped := new(time.Time)
		*var_e271a0d1c1c5_mapped = val.(time.Time)

		s.CreatedOn = var_e271a0d1c1c5_mapped
	}
	if properties["updatedOn"] != nil {

		var_60200a1867c6 := properties["updatedOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_60200a1867c6)

		if err != nil {
			panic(err)
		}

		var_60200a1867c6_mapped := new(time.Time)
		*var_60200a1867c6_mapped = val.(time.Time)

		s.UpdatedOn = var_60200a1867c6_mapped
	}
	if properties["namespace"] != nil {

		var_3b654b015993 := properties["namespace"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_3b654b015993)

		if err != nil {
			panic(err)
		}

		var_3b654b015993_mapped := val.(string)

		s.Namespace = var_3b654b015993_mapped
	}
	if properties["resource"] != nil {

		var_a2a5e1e65d37 := properties["resource"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_a2a5e1e65d37)

		if err != nil {
			panic(err)
		}

		var_a2a5e1e65d37_mapped := val.(string)

		s.Resource = var_a2a5e1e65d37_mapped
	}
	if properties["property"] != nil {

		var_b104efa8138f := properties["property"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_b104efa8138f)

		if err != nil {
			panic(err)
		}

		var_b104efa8138f_mapped := val.(string)

		s.Property = var_b104efa8138f_mapped
	}
	if properties["propertyValue"] != nil {

		var_4aaf4ee01edc := properties["propertyValue"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_4aaf4ee01edc)

		if err != nil {
			panic(err)
		}

		var_4aaf4ee01edc_mapped := new(string)
		*var_4aaf4ee01edc_mapped = val.(string)

		s.PropertyValue = var_4aaf4ee01edc_mapped
	}
	if properties["propertyMode"] != nil {

		var_1926d9d47fb4 := properties["propertyMode"]
		var_1926d9d47fb4_mapped := new(SecurityConstraintPropertyMode)
		*var_1926d9d47fb4_mapped = (SecurityConstraintPropertyMode)(var_1926d9d47fb4.GetStringValue())

		s.PropertyMode = var_1926d9d47fb4_mapped
	}
	if properties["operation"] != nil {

		var_8a17b726e30b := properties["operation"]
		var_8a17b726e30b_mapped := (SecurityConstraintOperation)(var_8a17b726e30b.GetStringValue())

		s.Operation = var_8a17b726e30b_mapped
	}
	if properties["recordIds"] != nil {

		var_787f78a68065 := properties["recordIds"]
		var_787f78a68065_mapped := []string{}
		for _, v := range var_787f78a68065.GetListValue().Values {

			var_39b6464ce1a3 := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_39b6464ce1a3)

			if err != nil {
				panic(err)
			}

			var_39b6464ce1a3_mapped := val.(string)

			var_787f78a68065_mapped = append(var_787f78a68065_mapped, var_39b6464ce1a3_mapped)
		}

		s.RecordIds = var_787f78a68065_mapped
	}
	if properties["before"] != nil {

		var_7457c4c0ab11 := properties["before"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_7457c4c0ab11)

		if err != nil {
			panic(err)
		}

		var_7457c4c0ab11_mapped := new(time.Time)
		*var_7457c4c0ab11_mapped = val.(time.Time)

		s.Before = var_7457c4c0ab11_mapped
	}
	if properties["after"] != nil {

		var_39b10d52a342 := properties["after"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_39b10d52a342)

		if err != nil {
			panic(err)
		}

		var_39b10d52a342_mapped := new(time.Time)
		*var_39b10d52a342_mapped = val.(time.Time)

		s.After = var_39b10d52a342_mapped
	}
	if properties["username"] != nil {

		var_27d2f499d042 := properties["username"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_27d2f499d042)

		if err != nil {
			panic(err)
		}

		var_27d2f499d042_mapped := new(string)
		*var_27d2f499d042_mapped = val.(string)

		s.Username = var_27d2f499d042_mapped
	}
	if properties["role"] != nil {

		var_04bf247ce637 := properties["role"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_04bf247ce637)

		if err != nil {
			panic(err)
		}

		var_04bf247ce637_mapped := new(string)
		*var_04bf247ce637_mapped = val.(string)

		s.Role = var_04bf247ce637_mapped
	}
	if properties["permit"] != nil {

		var_7c85c31f4dcb := properties["permit"]
		var_7c85c31f4dcb_mapped := (SecurityConstraintPermit)(var_7c85c31f4dcb.GetStringValue())

		s.Permit = var_7c85c31f4dcb_mapped
	}
	if properties["localFlags"] != nil {

		var_f90af17725ec := properties["localFlags"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_OBJECT).UnPack(var_f90af17725ec)

		if err != nil {
			panic(err)
		}

		var_f90af17725ec_mapped := new(unstructured.Unstructured)
		*var_f90af17725ec_mapped = val.(unstructured.Unstructured)

		s.LocalFlags = var_f90af17725ec_mapped
	}
	return s
}
