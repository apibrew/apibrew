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

	var_79e5f0e033c0 := securityConstraint.Id

	if var_79e5f0e033c0 != nil {
		var var_79e5f0e033c0_mapped *structpb.Value

		var var_79e5f0e033c0_err error
		var_79e5f0e033c0_mapped, var_79e5f0e033c0_err = types.ByResourcePropertyType(model.ResourceProperty_UUID).Pack(*var_79e5f0e033c0)
		if var_79e5f0e033c0_err != nil {
			panic(var_79e5f0e033c0_err)
		}
		properties["id"] = var_79e5f0e033c0_mapped
	}

	var_01d7cc57effe := securityConstraint.Version

	var var_01d7cc57effe_mapped *structpb.Value

	var var_01d7cc57effe_err error
	var_01d7cc57effe_mapped, var_01d7cc57effe_err = types.ByResourcePropertyType(model.ResourceProperty_INT32).Pack(var_01d7cc57effe)
	if var_01d7cc57effe_err != nil {
		panic(var_01d7cc57effe_err)
	}
	properties["version"] = var_01d7cc57effe_mapped

	var_4709f2475e1b := securityConstraint.CreatedBy

	if var_4709f2475e1b != nil {
		var var_4709f2475e1b_mapped *structpb.Value

		var var_4709f2475e1b_err error
		var_4709f2475e1b_mapped, var_4709f2475e1b_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_4709f2475e1b)
		if var_4709f2475e1b_err != nil {
			panic(var_4709f2475e1b_err)
		}
		properties["createdBy"] = var_4709f2475e1b_mapped
	}

	var_fce6995f0a97 := securityConstraint.UpdatedBy

	if var_fce6995f0a97 != nil {
		var var_fce6995f0a97_mapped *structpb.Value

		var var_fce6995f0a97_err error
		var_fce6995f0a97_mapped, var_fce6995f0a97_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_fce6995f0a97)
		if var_fce6995f0a97_err != nil {
			panic(var_fce6995f0a97_err)
		}
		properties["updatedBy"] = var_fce6995f0a97_mapped
	}

	var_bfceb348a82b := securityConstraint.CreatedOn

	if var_bfceb348a82b != nil {
		var var_bfceb348a82b_mapped *structpb.Value

		var var_bfceb348a82b_err error
		var_bfceb348a82b_mapped, var_bfceb348a82b_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_bfceb348a82b)
		if var_bfceb348a82b_err != nil {
			panic(var_bfceb348a82b_err)
		}
		properties["createdOn"] = var_bfceb348a82b_mapped
	}

	var_2099ac4feed6 := securityConstraint.UpdatedOn

	if var_2099ac4feed6 != nil {
		var var_2099ac4feed6_mapped *structpb.Value

		var var_2099ac4feed6_err error
		var_2099ac4feed6_mapped, var_2099ac4feed6_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_2099ac4feed6)
		if var_2099ac4feed6_err != nil {
			panic(var_2099ac4feed6_err)
		}
		properties["updatedOn"] = var_2099ac4feed6_mapped
	}

	var_6a9669038a86 := securityConstraint.Namespace

	if var_6a9669038a86 != nil {
		var var_6a9669038a86_mapped *structpb.Value

		var_6a9669038a86_mapped = structpb.NewStructValue(&structpb.Struct{Fields: NamespaceMapperInstance.ToProperties(var_6a9669038a86)})
		properties["namespace"] = var_6a9669038a86_mapped
	}

	var_01aacdf2a40b := securityConstraint.Resource

	if var_01aacdf2a40b != nil {
		var var_01aacdf2a40b_mapped *structpb.Value

		var_01aacdf2a40b_mapped = structpb.NewStructValue(&structpb.Struct{Fields: ResourceMapperInstance.ToProperties(var_01aacdf2a40b)})
		properties["resource"] = var_01aacdf2a40b_mapped
	}

	var_8c0672e0925b := securityConstraint.Property

	if var_8c0672e0925b != nil {
		var var_8c0672e0925b_mapped *structpb.Value

		var var_8c0672e0925b_err error
		var_8c0672e0925b_mapped, var_8c0672e0925b_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_8c0672e0925b)
		if var_8c0672e0925b_err != nil {
			panic(var_8c0672e0925b_err)
		}
		properties["property"] = var_8c0672e0925b_mapped
	}

	var_4c2ffc0821b0 := securityConstraint.PropertyValue

	if var_4c2ffc0821b0 != nil {
		var var_4c2ffc0821b0_mapped *structpb.Value

		var var_4c2ffc0821b0_err error
		var_4c2ffc0821b0_mapped, var_4c2ffc0821b0_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_4c2ffc0821b0)
		if var_4c2ffc0821b0_err != nil {
			panic(var_4c2ffc0821b0_err)
		}
		properties["propertyValue"] = var_4c2ffc0821b0_mapped
	}

	var_ffd3b5037712 := securityConstraint.PropertyMode

	if var_ffd3b5037712 != nil {
		var var_ffd3b5037712_mapped *structpb.Value

		var var_ffd3b5037712_err error
		var_ffd3b5037712_mapped, var_ffd3b5037712_err = types.ByResourcePropertyType(model.ResourceProperty_ENUM).Pack(string(*var_ffd3b5037712))
		if var_ffd3b5037712_err != nil {
			panic(var_ffd3b5037712_err)
		}
		properties["propertyMode"] = var_ffd3b5037712_mapped
	}

	var_26c4d9139fc0 := securityConstraint.Operation

	var var_26c4d9139fc0_mapped *structpb.Value

	var var_26c4d9139fc0_err error
	var_26c4d9139fc0_mapped, var_26c4d9139fc0_err = types.ByResourcePropertyType(model.ResourceProperty_ENUM).Pack(string(var_26c4d9139fc0))
	if var_26c4d9139fc0_err != nil {
		panic(var_26c4d9139fc0_err)
	}
	properties["operation"] = var_26c4d9139fc0_mapped

	var_1a107751c97d := securityConstraint.RecordIds

	if var_1a107751c97d != nil {
		var var_1a107751c97d_mapped *structpb.Value

		var var_1a107751c97d_l []*structpb.Value
		for _, value := range var_1a107751c97d {

			var_438715b07f7e := value
			var var_438715b07f7e_mapped *structpb.Value

			var var_438715b07f7e_err error
			var_438715b07f7e_mapped, var_438715b07f7e_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_438715b07f7e)
			if var_438715b07f7e_err != nil {
				panic(var_438715b07f7e_err)
			}

			var_1a107751c97d_l = append(var_1a107751c97d_l, var_438715b07f7e_mapped)
		}
		var_1a107751c97d_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_1a107751c97d_l})
		properties["recordIds"] = var_1a107751c97d_mapped
	}

	var_0688b16f8233 := securityConstraint.Before

	if var_0688b16f8233 != nil {
		var var_0688b16f8233_mapped *structpb.Value

		var var_0688b16f8233_err error
		var_0688b16f8233_mapped, var_0688b16f8233_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_0688b16f8233)
		if var_0688b16f8233_err != nil {
			panic(var_0688b16f8233_err)
		}
		properties["before"] = var_0688b16f8233_mapped
	}

	var_09bc78ddd8a1 := securityConstraint.After

	if var_09bc78ddd8a1 != nil {
		var var_09bc78ddd8a1_mapped *structpb.Value

		var var_09bc78ddd8a1_err error
		var_09bc78ddd8a1_mapped, var_09bc78ddd8a1_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_09bc78ddd8a1)
		if var_09bc78ddd8a1_err != nil {
			panic(var_09bc78ddd8a1_err)
		}
		properties["after"] = var_09bc78ddd8a1_mapped
	}

	var_ff12a9514501 := securityConstraint.User

	if var_ff12a9514501 != nil {
		var var_ff12a9514501_mapped *structpb.Value

		var_ff12a9514501_mapped = structpb.NewStructValue(&structpb.Struct{Fields: UserMapperInstance.ToProperties(var_ff12a9514501)})
		properties["user"] = var_ff12a9514501_mapped
	}

	var_cd5b87362cbd := securityConstraint.Role

	if var_cd5b87362cbd != nil {
		var var_cd5b87362cbd_mapped *structpb.Value

		var_cd5b87362cbd_mapped = structpb.NewStructValue(&structpb.Struct{Fields: RoleMapperInstance.ToProperties(var_cd5b87362cbd)})
		properties["role"] = var_cd5b87362cbd_mapped
	}

	var_68f3dea58b54 := securityConstraint.Permit

	var var_68f3dea58b54_mapped *structpb.Value

	var var_68f3dea58b54_err error
	var_68f3dea58b54_mapped, var_68f3dea58b54_err = types.ByResourcePropertyType(model.ResourceProperty_ENUM).Pack(string(var_68f3dea58b54))
	if var_68f3dea58b54_err != nil {
		panic(var_68f3dea58b54_err)
	}
	properties["permit"] = var_68f3dea58b54_mapped

	var_cb4922581139 := securityConstraint.LocalFlags

	if var_cb4922581139 != nil {
		var var_cb4922581139_mapped *structpb.Value

		var var_cb4922581139_err error
		var_cb4922581139_mapped, var_cb4922581139_err = types.ByResourcePropertyType(model.ResourceProperty_OBJECT).Pack(*var_cb4922581139)
		if var_cb4922581139_err != nil {
			panic(var_cb4922581139_err)
		}
		properties["localFlags"] = var_cb4922581139_mapped
	}
	return properties
}

func (m *SecurityConstraintMapper) FromProperties(properties map[string]*structpb.Value) *SecurityConstraint {
	var s = m.New()
	if properties["id"] != nil {

		var_495d7fd6fdb5 := properties["id"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_UUID).UnPack(var_495d7fd6fdb5)

		if err != nil {
			panic(err)
		}

		var_495d7fd6fdb5_mapped := new(uuid.UUID)
		*var_495d7fd6fdb5_mapped = val.(uuid.UUID)

		s.Id = var_495d7fd6fdb5_mapped
	}
	if properties["version"] != nil {

		var_a8cd03727ad8 := properties["version"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(var_a8cd03727ad8)

		if err != nil {
			panic(err)
		}

		var_a8cd03727ad8_mapped := val.(int32)

		s.Version = var_a8cd03727ad8_mapped
	}
	if properties["createdBy"] != nil {

		var_7b9d7e5ab754 := properties["createdBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_7b9d7e5ab754)

		if err != nil {
			panic(err)
		}

		var_7b9d7e5ab754_mapped := new(string)
		*var_7b9d7e5ab754_mapped = val.(string)

		s.CreatedBy = var_7b9d7e5ab754_mapped
	}
	if properties["updatedBy"] != nil {

		var_88e474d60f8d := properties["updatedBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_88e474d60f8d)

		if err != nil {
			panic(err)
		}

		var_88e474d60f8d_mapped := new(string)
		*var_88e474d60f8d_mapped = val.(string)

		s.UpdatedBy = var_88e474d60f8d_mapped
	}
	if properties["createdOn"] != nil {

		var_bc04c4a78f50 := properties["createdOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_bc04c4a78f50)

		if err != nil {
			panic(err)
		}

		var_bc04c4a78f50_mapped := new(time.Time)
		*var_bc04c4a78f50_mapped = val.(time.Time)

		s.CreatedOn = var_bc04c4a78f50_mapped
	}
	if properties["updatedOn"] != nil {

		var_5e9d3376bfa4 := properties["updatedOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_5e9d3376bfa4)

		if err != nil {
			panic(err)
		}

		var_5e9d3376bfa4_mapped := new(time.Time)
		*var_5e9d3376bfa4_mapped = val.(time.Time)

		s.UpdatedOn = var_5e9d3376bfa4_mapped
	}
	if properties["namespace"] != nil {

		var_fa50a64d9100 := properties["namespace"]
		var_fa50a64d9100_mapped := NamespaceMapperInstance.FromProperties(var_fa50a64d9100.GetStructValue().Fields)

		s.Namespace = var_fa50a64d9100_mapped
	}
	if properties["resource"] != nil {

		var_5fc5d24217ce := properties["resource"]
		var_5fc5d24217ce_mapped := ResourceMapperInstance.FromProperties(var_5fc5d24217ce.GetStructValue().Fields)

		s.Resource = var_5fc5d24217ce_mapped
	}
	if properties["property"] != nil {

		var_a5dc25ad1592 := properties["property"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_a5dc25ad1592)

		if err != nil {
			panic(err)
		}

		var_a5dc25ad1592_mapped := new(string)
		*var_a5dc25ad1592_mapped = val.(string)

		s.Property = var_a5dc25ad1592_mapped
	}
	if properties["propertyValue"] != nil {

		var_66b819cf34e9 := properties["propertyValue"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_66b819cf34e9)

		if err != nil {
			panic(err)
		}

		var_66b819cf34e9_mapped := new(string)
		*var_66b819cf34e9_mapped = val.(string)

		s.PropertyValue = var_66b819cf34e9_mapped
	}
	if properties["propertyMode"] != nil {

		var_b21673576be4 := properties["propertyMode"]
		var_b21673576be4_mapped := new(SecurityConstraintPropertyMode)
		*var_b21673576be4_mapped = (SecurityConstraintPropertyMode)(var_b21673576be4.GetStringValue())

		s.PropertyMode = var_b21673576be4_mapped
	}
	if properties["operation"] != nil {

		var_8260873d7f9b := properties["operation"]
		var_8260873d7f9b_mapped := (SecurityConstraintOperation)(var_8260873d7f9b.GetStringValue())

		s.Operation = var_8260873d7f9b_mapped
	}
	if properties["recordIds"] != nil {

		var_6d645052d2fd := properties["recordIds"]
		var_6d645052d2fd_mapped := []string{}
		for _, v := range var_6d645052d2fd.GetListValue().Values {

			var_d193e78c2004 := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_d193e78c2004)

			if err != nil {
				panic(err)
			}

			var_d193e78c2004_mapped := val.(string)

			var_6d645052d2fd_mapped = append(var_6d645052d2fd_mapped, var_d193e78c2004_mapped)
		}

		s.RecordIds = var_6d645052d2fd_mapped
	}
	if properties["before"] != nil {

		var_8ce759c87728 := properties["before"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_8ce759c87728)

		if err != nil {
			panic(err)
		}

		var_8ce759c87728_mapped := new(time.Time)
		*var_8ce759c87728_mapped = val.(time.Time)

		s.Before = var_8ce759c87728_mapped
	}
	if properties["after"] != nil {

		var_3c06dd3ca88e := properties["after"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_3c06dd3ca88e)

		if err != nil {
			panic(err)
		}

		var_3c06dd3ca88e_mapped := new(time.Time)
		*var_3c06dd3ca88e_mapped = val.(time.Time)

		s.After = var_3c06dd3ca88e_mapped
	}
	if properties["user"] != nil {

		var_e1fa7183fc72 := properties["user"]
		var_e1fa7183fc72_mapped := UserMapperInstance.FromProperties(var_e1fa7183fc72.GetStructValue().Fields)

		s.User = var_e1fa7183fc72_mapped
	}
	if properties["role"] != nil {

		var_a5f548cf2a1e := properties["role"]
		var_a5f548cf2a1e_mapped := RoleMapperInstance.FromProperties(var_a5f548cf2a1e.GetStructValue().Fields)

		s.Role = var_a5f548cf2a1e_mapped
	}
	if properties["permit"] != nil {

		var_31492e0b6164 := properties["permit"]
		var_31492e0b6164_mapped := (SecurityConstraintPermit)(var_31492e0b6164.GetStringValue())

		s.Permit = var_31492e0b6164_mapped
	}
	if properties["localFlags"] != nil {

		var_4c6c3f71d096 := properties["localFlags"]
		var_4c6c3f71d096_mapped := new(unstructured.Unstructured)
		*var_4c6c3f71d096_mapped = unstructured.FromStructValue(var_4c6c3f71d096.GetStructValue())

		s.LocalFlags = var_4c6c3f71d096_mapped
	}
	return s
}
