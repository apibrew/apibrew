// AUTOGENERATED FILE

//go:build !codeanalysis

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

	if securityConstraint.Id != nil {
		rec.Id = securityConstraint.Id.String()
	}

	return rec
}

func (m *SecurityConstraintMapper) FromRecord(record *model.Record) *SecurityConstraint {
	return m.FromProperties(record.Properties)
}

func (m *SecurityConstraintMapper) ToProperties(securityConstraint *SecurityConstraint) map[string]*structpb.Value {
	var properties = make(map[string]*structpb.Value)

	var_b4090d397e51 := securityConstraint.Id

	if var_b4090d397e51 != nil {
		var var_b4090d397e51_mapped *structpb.Value

		var var_b4090d397e51_err error
		var_b4090d397e51_mapped, var_b4090d397e51_err = types.ByResourcePropertyType(model.ResourceProperty_UUID).Pack(*var_b4090d397e51)
		if var_b4090d397e51_err != nil {
			panic(var_b4090d397e51_err)
		}
		properties["id"] = var_b4090d397e51_mapped
	}

	var_4a32ff92588e := securityConstraint.Version

	var var_4a32ff92588e_mapped *structpb.Value

	var var_4a32ff92588e_err error
	var_4a32ff92588e_mapped, var_4a32ff92588e_err = types.ByResourcePropertyType(model.ResourceProperty_INT32).Pack(var_4a32ff92588e)
	if var_4a32ff92588e_err != nil {
		panic(var_4a32ff92588e_err)
	}
	properties["version"] = var_4a32ff92588e_mapped

	var_a5faee161e3b := securityConstraint.CreatedBy

	if var_a5faee161e3b != nil {
		var var_a5faee161e3b_mapped *structpb.Value

		var var_a5faee161e3b_err error
		var_a5faee161e3b_mapped, var_a5faee161e3b_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_a5faee161e3b)
		if var_a5faee161e3b_err != nil {
			panic(var_a5faee161e3b_err)
		}
		properties["createdBy"] = var_a5faee161e3b_mapped
	}

	var_1f5aa0724b3e := securityConstraint.UpdatedBy

	if var_1f5aa0724b3e != nil {
		var var_1f5aa0724b3e_mapped *structpb.Value

		var var_1f5aa0724b3e_err error
		var_1f5aa0724b3e_mapped, var_1f5aa0724b3e_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_1f5aa0724b3e)
		if var_1f5aa0724b3e_err != nil {
			panic(var_1f5aa0724b3e_err)
		}
		properties["updatedBy"] = var_1f5aa0724b3e_mapped
	}

	var_a6c84b03e459 := securityConstraint.CreatedOn

	if var_a6c84b03e459 != nil {
		var var_a6c84b03e459_mapped *structpb.Value

		var var_a6c84b03e459_err error
		var_a6c84b03e459_mapped, var_a6c84b03e459_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_a6c84b03e459)
		if var_a6c84b03e459_err != nil {
			panic(var_a6c84b03e459_err)
		}
		properties["createdOn"] = var_a6c84b03e459_mapped
	}

	var_e067d58da6ff := securityConstraint.UpdatedOn

	if var_e067d58da6ff != nil {
		var var_e067d58da6ff_mapped *structpb.Value

		var var_e067d58da6ff_err error
		var_e067d58da6ff_mapped, var_e067d58da6ff_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_e067d58da6ff)
		if var_e067d58da6ff_err != nil {
			panic(var_e067d58da6ff_err)
		}
		properties["updatedOn"] = var_e067d58da6ff_mapped
	}

	var_885eaa66875e := securityConstraint.Namespace

	if var_885eaa66875e != nil {
		var var_885eaa66875e_mapped *structpb.Value

		var var_885eaa66875e_err error
		var_885eaa66875e_mapped, var_885eaa66875e_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_885eaa66875e)
		if var_885eaa66875e_err != nil {
			panic(var_885eaa66875e_err)
		}
		properties["namespace"] = var_885eaa66875e_mapped
	}

	var_8c7906024781 := securityConstraint.Resource

	if var_8c7906024781 != nil {
		var var_8c7906024781_mapped *structpb.Value

		var var_8c7906024781_err error
		var_8c7906024781_mapped, var_8c7906024781_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_8c7906024781)
		if var_8c7906024781_err != nil {
			panic(var_8c7906024781_err)
		}
		properties["resource"] = var_8c7906024781_mapped
	}

	var_1545e89ecce3 := securityConstraint.Property

	if var_1545e89ecce3 != nil {
		var var_1545e89ecce3_mapped *structpb.Value

		var var_1545e89ecce3_err error
		var_1545e89ecce3_mapped, var_1545e89ecce3_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_1545e89ecce3)
		if var_1545e89ecce3_err != nil {
			panic(var_1545e89ecce3_err)
		}
		properties["property"] = var_1545e89ecce3_mapped
	}

	var_854aafc8b003 := securityConstraint.PropertyValue

	if var_854aafc8b003 != nil {
		var var_854aafc8b003_mapped *structpb.Value

		var var_854aafc8b003_err error
		var_854aafc8b003_mapped, var_854aafc8b003_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_854aafc8b003)
		if var_854aafc8b003_err != nil {
			panic(var_854aafc8b003_err)
		}
		properties["propertyValue"] = var_854aafc8b003_mapped
	}

	var_893c59fec4c8 := securityConstraint.PropertyMode

	if var_893c59fec4c8 != nil {
		var var_893c59fec4c8_mapped *structpb.Value

		var var_893c59fec4c8_err error
		var_893c59fec4c8_mapped, var_893c59fec4c8_err = types.ByResourcePropertyType(model.ResourceProperty_ENUM).Pack(string(*var_893c59fec4c8))
		if var_893c59fec4c8_err != nil {
			panic(var_893c59fec4c8_err)
		}
		properties["propertyMode"] = var_893c59fec4c8_mapped
	}

	var_5e2ebd6e0659 := securityConstraint.Operation

	var var_5e2ebd6e0659_mapped *structpb.Value

	var var_5e2ebd6e0659_err error
	var_5e2ebd6e0659_mapped, var_5e2ebd6e0659_err = types.ByResourcePropertyType(model.ResourceProperty_ENUM).Pack(string(var_5e2ebd6e0659))
	if var_5e2ebd6e0659_err != nil {
		panic(var_5e2ebd6e0659_err)
	}
	properties["operation"] = var_5e2ebd6e0659_mapped

	var_995c95a2f9ec := securityConstraint.RecordIds

	if var_995c95a2f9ec != nil {
		var var_995c95a2f9ec_mapped *structpb.Value

		var var_995c95a2f9ec_l []*structpb.Value
		for _, value := range var_995c95a2f9ec {

			var_e1d02bbee1d5 := value
			var var_e1d02bbee1d5_mapped *structpb.Value

			var var_e1d02bbee1d5_err error
			var_e1d02bbee1d5_mapped, var_e1d02bbee1d5_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_e1d02bbee1d5)
			if var_e1d02bbee1d5_err != nil {
				panic(var_e1d02bbee1d5_err)
			}

			var_995c95a2f9ec_l = append(var_995c95a2f9ec_l, var_e1d02bbee1d5_mapped)
		}
		var_995c95a2f9ec_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_995c95a2f9ec_l})
		properties["recordIds"] = var_995c95a2f9ec_mapped
	}

	var_587b2e9ced1d := securityConstraint.Before

	if var_587b2e9ced1d != nil {
		var var_587b2e9ced1d_mapped *structpb.Value

		var var_587b2e9ced1d_err error
		var_587b2e9ced1d_mapped, var_587b2e9ced1d_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_587b2e9ced1d)
		if var_587b2e9ced1d_err != nil {
			panic(var_587b2e9ced1d_err)
		}
		properties["before"] = var_587b2e9ced1d_mapped
	}

	var_82beee7dcfa9 := securityConstraint.After

	if var_82beee7dcfa9 != nil {
		var var_82beee7dcfa9_mapped *structpb.Value

		var var_82beee7dcfa9_err error
		var_82beee7dcfa9_mapped, var_82beee7dcfa9_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_82beee7dcfa9)
		if var_82beee7dcfa9_err != nil {
			panic(var_82beee7dcfa9_err)
		}
		properties["after"] = var_82beee7dcfa9_mapped
	}

	var_fe937bf7b8c8 := securityConstraint.User

	if var_fe937bf7b8c8 != nil {
		var var_fe937bf7b8c8_mapped *structpb.Value

		var_fe937bf7b8c8_mapped = structpb.NewStructValue(&structpb.Struct{Fields: UserMapperInstance.ToProperties(var_fe937bf7b8c8)})
		properties["user"] = var_fe937bf7b8c8_mapped
	}

	var_cd6991a6e398 := securityConstraint.Role

	if var_cd6991a6e398 != nil {
		var var_cd6991a6e398_mapped *structpb.Value

		var_cd6991a6e398_mapped = structpb.NewStructValue(&structpb.Struct{Fields: RoleMapperInstance.ToProperties(var_cd6991a6e398)})
		properties["role"] = var_cd6991a6e398_mapped
	}

	var_4f110c425c7a := securityConstraint.Permit

	var var_4f110c425c7a_mapped *structpb.Value

	var var_4f110c425c7a_err error
	var_4f110c425c7a_mapped, var_4f110c425c7a_err = types.ByResourcePropertyType(model.ResourceProperty_ENUM).Pack(string(var_4f110c425c7a))
	if var_4f110c425c7a_err != nil {
		panic(var_4f110c425c7a_err)
	}
	properties["permit"] = var_4f110c425c7a_mapped

	var_a57c2b22a519 := securityConstraint.LocalFlags

	if var_a57c2b22a519 != nil {
		var var_a57c2b22a519_mapped *structpb.Value

		var var_a57c2b22a519_err error
		var_a57c2b22a519_mapped, var_a57c2b22a519_err = types.ByResourcePropertyType(model.ResourceProperty_OBJECT).Pack(*var_a57c2b22a519)
		if var_a57c2b22a519_err != nil {
			panic(var_a57c2b22a519_err)
		}
		properties["localFlags"] = var_a57c2b22a519_mapped
	}
	return properties
}

func (m *SecurityConstraintMapper) FromProperties(properties map[string]*structpb.Value) *SecurityConstraint {
	var s = m.New()
	if properties["id"] != nil {

		var_6a411e7b7502 := properties["id"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_UUID).UnPack(var_6a411e7b7502)

		if err != nil {
			panic(err)
		}

		var_6a411e7b7502_mapped := new(uuid.UUID)
		*var_6a411e7b7502_mapped = val.(uuid.UUID)

		s.Id = var_6a411e7b7502_mapped
	}
	if properties["version"] != nil {

		var_3f25fdd225d2 := properties["version"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(var_3f25fdd225d2)

		if err != nil {
			panic(err)
		}

		var_3f25fdd225d2_mapped := val.(int32)

		s.Version = var_3f25fdd225d2_mapped
	}
	if properties["createdBy"] != nil {

		var_0dbe3cf40bf0 := properties["createdBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_0dbe3cf40bf0)

		if err != nil {
			panic(err)
		}

		var_0dbe3cf40bf0_mapped := new(string)
		*var_0dbe3cf40bf0_mapped = val.(string)

		s.CreatedBy = var_0dbe3cf40bf0_mapped
	}
	if properties["updatedBy"] != nil {

		var_22c31bafaca0 := properties["updatedBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_22c31bafaca0)

		if err != nil {
			panic(err)
		}

		var_22c31bafaca0_mapped := new(string)
		*var_22c31bafaca0_mapped = val.(string)

		s.UpdatedBy = var_22c31bafaca0_mapped
	}
	if properties["createdOn"] != nil {

		var_904db8e2c66c := properties["createdOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_904db8e2c66c)

		if err != nil {
			panic(err)
		}

		var_904db8e2c66c_mapped := new(time.Time)
		*var_904db8e2c66c_mapped = val.(time.Time)

		s.CreatedOn = var_904db8e2c66c_mapped
	}
	if properties["updatedOn"] != nil {

		var_4b58f912004d := properties["updatedOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_4b58f912004d)

		if err != nil {
			panic(err)
		}

		var_4b58f912004d_mapped := new(time.Time)
		*var_4b58f912004d_mapped = val.(time.Time)

		s.UpdatedOn = var_4b58f912004d_mapped
	}
	if properties["namespace"] != nil {

		var_46fcbee56c18 := properties["namespace"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_46fcbee56c18)

		if err != nil {
			panic(err)
		}

		var_46fcbee56c18_mapped := new(string)
		*var_46fcbee56c18_mapped = val.(string)

		s.Namespace = var_46fcbee56c18_mapped
	}
	if properties["resource"] != nil {

		var_23d236cb0451 := properties["resource"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_23d236cb0451)

		if err != nil {
			panic(err)
		}

		var_23d236cb0451_mapped := new(string)
		*var_23d236cb0451_mapped = val.(string)

		s.Resource = var_23d236cb0451_mapped
	}
	if properties["property"] != nil {

		var_a229482fc8a8 := properties["property"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_a229482fc8a8)

		if err != nil {
			panic(err)
		}

		var_a229482fc8a8_mapped := new(string)
		*var_a229482fc8a8_mapped = val.(string)

		s.Property = var_a229482fc8a8_mapped
	}
	if properties["propertyValue"] != nil {

		var_75aadbe9947b := properties["propertyValue"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_75aadbe9947b)

		if err != nil {
			panic(err)
		}

		var_75aadbe9947b_mapped := new(string)
		*var_75aadbe9947b_mapped = val.(string)

		s.PropertyValue = var_75aadbe9947b_mapped
	}
	if properties["propertyMode"] != nil {

		var_e0264eeacbc2 := properties["propertyMode"]
		var_e0264eeacbc2_mapped := new(SecurityConstraintPropertyMode)
		*var_e0264eeacbc2_mapped = (SecurityConstraintPropertyMode)(var_e0264eeacbc2.GetStringValue())

		s.PropertyMode = var_e0264eeacbc2_mapped
	}
	if properties["operation"] != nil {

		var_bf56dd769101 := properties["operation"]
		var_bf56dd769101_mapped := (SecurityConstraintOperation)(var_bf56dd769101.GetStringValue())

		s.Operation = var_bf56dd769101_mapped
	}
	if properties["recordIds"] != nil {

		var_9cfb2a1b68e9 := properties["recordIds"]
		var_9cfb2a1b68e9_mapped := []string{}
		for _, v := range var_9cfb2a1b68e9.GetListValue().Values {

			var_05b2fde26c26 := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_05b2fde26c26)

			if err != nil {
				panic(err)
			}

			var_05b2fde26c26_mapped := val.(string)

			var_9cfb2a1b68e9_mapped = append(var_9cfb2a1b68e9_mapped, var_05b2fde26c26_mapped)
		}

		s.RecordIds = var_9cfb2a1b68e9_mapped
	}
	if properties["before"] != nil {

		var_bace661825a3 := properties["before"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_bace661825a3)

		if err != nil {
			panic(err)
		}

		var_bace661825a3_mapped := new(time.Time)
		*var_bace661825a3_mapped = val.(time.Time)

		s.Before = var_bace661825a3_mapped
	}
	if properties["after"] != nil {

		var_6b7d0bd861e1 := properties["after"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_6b7d0bd861e1)

		if err != nil {
			panic(err)
		}

		var_6b7d0bd861e1_mapped := new(time.Time)
		*var_6b7d0bd861e1_mapped = val.(time.Time)

		s.After = var_6b7d0bd861e1_mapped
	}
	if properties["user"] != nil {

		var_62613d579711 := properties["user"]
		var_62613d579711_mapped := UserMapperInstance.FromProperties(var_62613d579711.GetStructValue().Fields)

		s.User = var_62613d579711_mapped
	}
	if properties["role"] != nil {

		var_a9c951243f7a := properties["role"]
		var_a9c951243f7a_mapped := RoleMapperInstance.FromProperties(var_a9c951243f7a.GetStructValue().Fields)

		s.Role = var_a9c951243f7a_mapped
	}
	if properties["permit"] != nil {

		var_7fbbb00354e0 := properties["permit"]
		var_7fbbb00354e0_mapped := (SecurityConstraintPermit)(var_7fbbb00354e0.GetStringValue())

		s.Permit = var_7fbbb00354e0_mapped
	}
	if properties["localFlags"] != nil {

		var_38c81efdcc27 := properties["localFlags"]
		var_38c81efdcc27_mapped := new(unstructured.Unstructured)
		*var_38c81efdcc27_mapped = unstructured.FromStructValue(var_38c81efdcc27.GetStructValue())

		s.LocalFlags = var_38c81efdcc27_mapped
	}
	return s
}
