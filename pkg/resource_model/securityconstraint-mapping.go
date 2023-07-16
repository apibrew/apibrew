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

	var_89703ba0441d := securityConstraint.Id

	if var_89703ba0441d != nil {
		var var_89703ba0441d_mapped *structpb.Value

		var var_89703ba0441d_err error
		var_89703ba0441d_mapped, var_89703ba0441d_err = types.ByResourcePropertyType(model.ResourceProperty_UUID).Pack(*var_89703ba0441d)
		if var_89703ba0441d_err != nil {
			panic(var_89703ba0441d_err)
		}
		properties["id"] = var_89703ba0441d_mapped
	}

	var_5986eafa652b := securityConstraint.Version

	var var_5986eafa652b_mapped *structpb.Value

	var var_5986eafa652b_err error
	var_5986eafa652b_mapped, var_5986eafa652b_err = types.ByResourcePropertyType(model.ResourceProperty_INT32).Pack(var_5986eafa652b)
	if var_5986eafa652b_err != nil {
		panic(var_5986eafa652b_err)
	}
	properties["version"] = var_5986eafa652b_mapped

	var_0e5b815da53c := securityConstraint.CreatedBy

	if var_0e5b815da53c != nil {
		var var_0e5b815da53c_mapped *structpb.Value

		var var_0e5b815da53c_err error
		var_0e5b815da53c_mapped, var_0e5b815da53c_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_0e5b815da53c)
		if var_0e5b815da53c_err != nil {
			panic(var_0e5b815da53c_err)
		}
		properties["createdBy"] = var_0e5b815da53c_mapped
	}

	var_2fbfadd7a3b5 := securityConstraint.UpdatedBy

	if var_2fbfadd7a3b5 != nil {
		var var_2fbfadd7a3b5_mapped *structpb.Value

		var var_2fbfadd7a3b5_err error
		var_2fbfadd7a3b5_mapped, var_2fbfadd7a3b5_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_2fbfadd7a3b5)
		if var_2fbfadd7a3b5_err != nil {
			panic(var_2fbfadd7a3b5_err)
		}
		properties["updatedBy"] = var_2fbfadd7a3b5_mapped
	}

	var_b0005913b5d7 := securityConstraint.CreatedOn

	if var_b0005913b5d7 != nil {
		var var_b0005913b5d7_mapped *structpb.Value

		var var_b0005913b5d7_err error
		var_b0005913b5d7_mapped, var_b0005913b5d7_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_b0005913b5d7)
		if var_b0005913b5d7_err != nil {
			panic(var_b0005913b5d7_err)
		}
		properties["createdOn"] = var_b0005913b5d7_mapped
	}

	var_3c135ef5b0b6 := securityConstraint.UpdatedOn

	if var_3c135ef5b0b6 != nil {
		var var_3c135ef5b0b6_mapped *structpb.Value

		var var_3c135ef5b0b6_err error
		var_3c135ef5b0b6_mapped, var_3c135ef5b0b6_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_3c135ef5b0b6)
		if var_3c135ef5b0b6_err != nil {
			panic(var_3c135ef5b0b6_err)
		}
		properties["updatedOn"] = var_3c135ef5b0b6_mapped
	}

	var_01f9e8f4b92d := securityConstraint.Namespace

	var var_01f9e8f4b92d_mapped *structpb.Value

	var var_01f9e8f4b92d_err error
	var_01f9e8f4b92d_mapped, var_01f9e8f4b92d_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_01f9e8f4b92d)
	if var_01f9e8f4b92d_err != nil {
		panic(var_01f9e8f4b92d_err)
	}
	properties["namespace"] = var_01f9e8f4b92d_mapped

	var_c5724cab74f6 := securityConstraint.Resource

	var var_c5724cab74f6_mapped *structpb.Value

	var var_c5724cab74f6_err error
	var_c5724cab74f6_mapped, var_c5724cab74f6_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_c5724cab74f6)
	if var_c5724cab74f6_err != nil {
		panic(var_c5724cab74f6_err)
	}
	properties["resource"] = var_c5724cab74f6_mapped

	var_4fc1d5035d3f := securityConstraint.Property

	var var_4fc1d5035d3f_mapped *structpb.Value

	var var_4fc1d5035d3f_err error
	var_4fc1d5035d3f_mapped, var_4fc1d5035d3f_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_4fc1d5035d3f)
	if var_4fc1d5035d3f_err != nil {
		panic(var_4fc1d5035d3f_err)
	}
	properties["property"] = var_4fc1d5035d3f_mapped

	var_87ae2e382f0a := securityConstraint.PropertyValue

	if var_87ae2e382f0a != nil {
		var var_87ae2e382f0a_mapped *structpb.Value

		var var_87ae2e382f0a_err error
		var_87ae2e382f0a_mapped, var_87ae2e382f0a_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_87ae2e382f0a)
		if var_87ae2e382f0a_err != nil {
			panic(var_87ae2e382f0a_err)
		}
		properties["propertyValue"] = var_87ae2e382f0a_mapped
	}

	var_c5d832536c20 := securityConstraint.PropertyMode

	if var_c5d832536c20 != nil {
		var var_c5d832536c20_mapped *structpb.Value

		var var_c5d832536c20_err error
		var_c5d832536c20_mapped, var_c5d832536c20_err = types.ByResourcePropertyType(model.ResourceProperty_ENUM).Pack(string(*var_c5d832536c20))
		if var_c5d832536c20_err != nil {
			panic(var_c5d832536c20_err)
		}
		properties["propertyMode"] = var_c5d832536c20_mapped
	}

	var_3c16a847c826 := securityConstraint.Operation

	var var_3c16a847c826_mapped *structpb.Value

	var var_3c16a847c826_err error
	var_3c16a847c826_mapped, var_3c16a847c826_err = types.ByResourcePropertyType(model.ResourceProperty_ENUM).Pack(string(var_3c16a847c826))
	if var_3c16a847c826_err != nil {
		panic(var_3c16a847c826_err)
	}
	properties["operation"] = var_3c16a847c826_mapped

	var_6c0c39f00bbe := securityConstraint.RecordIds

	if var_6c0c39f00bbe != nil {
		var var_6c0c39f00bbe_mapped *structpb.Value

		var var_6c0c39f00bbe_l []*structpb.Value
		for _, value := range var_6c0c39f00bbe {

			var_1800d8fa2711 := value
			var var_1800d8fa2711_mapped *structpb.Value

			var var_1800d8fa2711_err error
			var_1800d8fa2711_mapped, var_1800d8fa2711_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_1800d8fa2711)
			if var_1800d8fa2711_err != nil {
				panic(var_1800d8fa2711_err)
			}

			var_6c0c39f00bbe_l = append(var_6c0c39f00bbe_l, var_1800d8fa2711_mapped)
		}
		var_6c0c39f00bbe_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_6c0c39f00bbe_l})
		properties["recordIds"] = var_6c0c39f00bbe_mapped
	}

	var_e40ebdf1db0c := securityConstraint.Before

	if var_e40ebdf1db0c != nil {
		var var_e40ebdf1db0c_mapped *structpb.Value

		var var_e40ebdf1db0c_err error
		var_e40ebdf1db0c_mapped, var_e40ebdf1db0c_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_e40ebdf1db0c)
		if var_e40ebdf1db0c_err != nil {
			panic(var_e40ebdf1db0c_err)
		}
		properties["before"] = var_e40ebdf1db0c_mapped
	}

	var_c10dd5d6bdec := securityConstraint.After

	if var_c10dd5d6bdec != nil {
		var var_c10dd5d6bdec_mapped *structpb.Value

		var var_c10dd5d6bdec_err error
		var_c10dd5d6bdec_mapped, var_c10dd5d6bdec_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_c10dd5d6bdec)
		if var_c10dd5d6bdec_err != nil {
			panic(var_c10dd5d6bdec_err)
		}
		properties["after"] = var_c10dd5d6bdec_mapped
	}

	var_86b0c7204af3 := securityConstraint.Username

	if var_86b0c7204af3 != nil {
		var var_86b0c7204af3_mapped *structpb.Value

		var var_86b0c7204af3_err error
		var_86b0c7204af3_mapped, var_86b0c7204af3_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_86b0c7204af3)
		if var_86b0c7204af3_err != nil {
			panic(var_86b0c7204af3_err)
		}
		properties["username"] = var_86b0c7204af3_mapped
	}

	var_7f5ca49e5126 := securityConstraint.Role

	if var_7f5ca49e5126 != nil {
		var var_7f5ca49e5126_mapped *structpb.Value

		var var_7f5ca49e5126_err error
		var_7f5ca49e5126_mapped, var_7f5ca49e5126_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_7f5ca49e5126)
		if var_7f5ca49e5126_err != nil {
			panic(var_7f5ca49e5126_err)
		}
		properties["role"] = var_7f5ca49e5126_mapped
	}

	var_5593ad8ed99c := securityConstraint.Permit

	var var_5593ad8ed99c_mapped *structpb.Value

	var var_5593ad8ed99c_err error
	var_5593ad8ed99c_mapped, var_5593ad8ed99c_err = types.ByResourcePropertyType(model.ResourceProperty_ENUM).Pack(string(var_5593ad8ed99c))
	if var_5593ad8ed99c_err != nil {
		panic(var_5593ad8ed99c_err)
	}
	properties["permit"] = var_5593ad8ed99c_mapped

	var_e4be15b7eff6 := securityConstraint.LocalFlags

	if var_e4be15b7eff6 != nil {
		var var_e4be15b7eff6_mapped *structpb.Value

		var var_e4be15b7eff6_err error
		var_e4be15b7eff6_mapped, var_e4be15b7eff6_err = types.ByResourcePropertyType(model.ResourceProperty_OBJECT).Pack(*var_e4be15b7eff6)
		if var_e4be15b7eff6_err != nil {
			panic(var_e4be15b7eff6_err)
		}
		properties["localFlags"] = var_e4be15b7eff6_mapped
	}
	return properties
}

func (m *SecurityConstraintMapper) FromProperties(properties map[string]*structpb.Value) *SecurityConstraint {
	var s = m.New()
	if properties["id"] != nil {

		var_961bef3afbf6 := properties["id"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_UUID).UnPack(var_961bef3afbf6)

		if err != nil {
			panic(err)
		}

		var_961bef3afbf6_mapped := new(uuid.UUID)
		*var_961bef3afbf6_mapped = val.(uuid.UUID)

		s.Id = var_961bef3afbf6_mapped
	}
	if properties["version"] != nil {

		var_0ce93691e029 := properties["version"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(var_0ce93691e029)

		if err != nil {
			panic(err)
		}

		var_0ce93691e029_mapped := val.(int32)

		s.Version = var_0ce93691e029_mapped
	}
	if properties["createdBy"] != nil {

		var_e943f28e6e46 := properties["createdBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_e943f28e6e46)

		if err != nil {
			panic(err)
		}

		var_e943f28e6e46_mapped := new(string)
		*var_e943f28e6e46_mapped = val.(string)

		s.CreatedBy = var_e943f28e6e46_mapped
	}
	if properties["updatedBy"] != nil {

		var_4259006837c0 := properties["updatedBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_4259006837c0)

		if err != nil {
			panic(err)
		}

		var_4259006837c0_mapped := new(string)
		*var_4259006837c0_mapped = val.(string)

		s.UpdatedBy = var_4259006837c0_mapped
	}
	if properties["createdOn"] != nil {

		var_0e9e4e11f84a := properties["createdOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_0e9e4e11f84a)

		if err != nil {
			panic(err)
		}

		var_0e9e4e11f84a_mapped := new(time.Time)
		*var_0e9e4e11f84a_mapped = val.(time.Time)

		s.CreatedOn = var_0e9e4e11f84a_mapped
	}
	if properties["updatedOn"] != nil {

		var_6a3a215d66d2 := properties["updatedOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_6a3a215d66d2)

		if err != nil {
			panic(err)
		}

		var_6a3a215d66d2_mapped := new(time.Time)
		*var_6a3a215d66d2_mapped = val.(time.Time)

		s.UpdatedOn = var_6a3a215d66d2_mapped
	}
	if properties["namespace"] != nil {

		var_3bf424855ff3 := properties["namespace"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_3bf424855ff3)

		if err != nil {
			panic(err)
		}

		var_3bf424855ff3_mapped := val.(string)

		s.Namespace = var_3bf424855ff3_mapped
	}
	if properties["resource"] != nil {

		var_628c44c2858b := properties["resource"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_628c44c2858b)

		if err != nil {
			panic(err)
		}

		var_628c44c2858b_mapped := val.(string)

		s.Resource = var_628c44c2858b_mapped
	}
	if properties["property"] != nil {

		var_8be691bfaedf := properties["property"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_8be691bfaedf)

		if err != nil {
			panic(err)
		}

		var_8be691bfaedf_mapped := val.(string)

		s.Property = var_8be691bfaedf_mapped
	}
	if properties["propertyValue"] != nil {

		var_8813578bdc97 := properties["propertyValue"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_8813578bdc97)

		if err != nil {
			panic(err)
		}

		var_8813578bdc97_mapped := new(string)
		*var_8813578bdc97_mapped = val.(string)

		s.PropertyValue = var_8813578bdc97_mapped
	}
	if properties["propertyMode"] != nil {

		var_652fb09ea54f := properties["propertyMode"]
		var_652fb09ea54f_mapped := new(SecurityConstraintPropertyMode)
		*var_652fb09ea54f_mapped = (SecurityConstraintPropertyMode)(var_652fb09ea54f.GetStringValue())

		s.PropertyMode = var_652fb09ea54f_mapped
	}
	if properties["operation"] != nil {

		var_9268bc5f8a07 := properties["operation"]
		var_9268bc5f8a07_mapped := (SecurityConstraintOperation)(var_9268bc5f8a07.GetStringValue())

		s.Operation = var_9268bc5f8a07_mapped
	}
	if properties["recordIds"] != nil {

		var_aead1a6dd0de := properties["recordIds"]
		var_aead1a6dd0de_mapped := []string{}
		for _, v := range var_aead1a6dd0de.GetListValue().Values {

			var_bfe782109ff0 := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_bfe782109ff0)

			if err != nil {
				panic(err)
			}

			var_bfe782109ff0_mapped := val.(string)

			var_aead1a6dd0de_mapped = append(var_aead1a6dd0de_mapped, var_bfe782109ff0_mapped)
		}

		s.RecordIds = var_aead1a6dd0de_mapped
	}
	if properties["before"] != nil {

		var_8f2223894535 := properties["before"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_8f2223894535)

		if err != nil {
			panic(err)
		}

		var_8f2223894535_mapped := new(time.Time)
		*var_8f2223894535_mapped = val.(time.Time)

		s.Before = var_8f2223894535_mapped
	}
	if properties["after"] != nil {

		var_899e30e3c652 := properties["after"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_899e30e3c652)

		if err != nil {
			panic(err)
		}

		var_899e30e3c652_mapped := new(time.Time)
		*var_899e30e3c652_mapped = val.(time.Time)

		s.After = var_899e30e3c652_mapped
	}
	if properties["username"] != nil {

		var_c13bdcc027dd := properties["username"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_c13bdcc027dd)

		if err != nil {
			panic(err)
		}

		var_c13bdcc027dd_mapped := new(string)
		*var_c13bdcc027dd_mapped = val.(string)

		s.Username = var_c13bdcc027dd_mapped
	}
	if properties["role"] != nil {

		var_a7b3fd324a0b := properties["role"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_a7b3fd324a0b)

		if err != nil {
			panic(err)
		}

		var_a7b3fd324a0b_mapped := new(string)
		*var_a7b3fd324a0b_mapped = val.(string)

		s.Role = var_a7b3fd324a0b_mapped
	}
	if properties["permit"] != nil {

		var_c8d417fc1857 := properties["permit"]
		var_c8d417fc1857_mapped := (SecurityConstraintPermit)(var_c8d417fc1857.GetStringValue())

		s.Permit = var_c8d417fc1857_mapped
	}
	if properties["localFlags"] != nil {

		var_e55d1e38bcdd := properties["localFlags"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_OBJECT).UnPack(var_e55d1e38bcdd)

		if err != nil {
			panic(err)
		}

		var_e55d1e38bcdd_mapped := new(unstructured.Unstructured)
		*var_e55d1e38bcdd_mapped = val.(unstructured.Unstructured)

		s.LocalFlags = var_e55d1e38bcdd_mapped
	}
	return s
}
