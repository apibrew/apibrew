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

	var_f1dca72bdc72 := securityConstraint.Id

	if var_f1dca72bdc72 != nil {
		var var_f1dca72bdc72_mapped *structpb.Value

		var var_f1dca72bdc72_err error
		var_f1dca72bdc72_mapped, var_f1dca72bdc72_err = types.ByResourcePropertyType(model.ResourceProperty_UUID).Pack(*var_f1dca72bdc72)
		if var_f1dca72bdc72_err != nil {
			panic(var_f1dca72bdc72_err)
		}
		properties["id"] = var_f1dca72bdc72_mapped
	}

	var_df0269ef5079 := securityConstraint.Version

	var var_df0269ef5079_mapped *structpb.Value

	var var_df0269ef5079_err error
	var_df0269ef5079_mapped, var_df0269ef5079_err = types.ByResourcePropertyType(model.ResourceProperty_INT32).Pack(var_df0269ef5079)
	if var_df0269ef5079_err != nil {
		panic(var_df0269ef5079_err)
	}
	properties["version"] = var_df0269ef5079_mapped

	var_2d156e65bffe := securityConstraint.CreatedBy

	if var_2d156e65bffe != nil {
		var var_2d156e65bffe_mapped *structpb.Value

		var var_2d156e65bffe_err error
		var_2d156e65bffe_mapped, var_2d156e65bffe_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_2d156e65bffe)
		if var_2d156e65bffe_err != nil {
			panic(var_2d156e65bffe_err)
		}
		properties["createdBy"] = var_2d156e65bffe_mapped
	}

	var_490ffd278d40 := securityConstraint.UpdatedBy

	if var_490ffd278d40 != nil {
		var var_490ffd278d40_mapped *structpb.Value

		var var_490ffd278d40_err error
		var_490ffd278d40_mapped, var_490ffd278d40_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_490ffd278d40)
		if var_490ffd278d40_err != nil {
			panic(var_490ffd278d40_err)
		}
		properties["updatedBy"] = var_490ffd278d40_mapped
	}

	var_b68e445bf1c2 := securityConstraint.CreatedOn

	if var_b68e445bf1c2 != nil {
		var var_b68e445bf1c2_mapped *structpb.Value

		var var_b68e445bf1c2_err error
		var_b68e445bf1c2_mapped, var_b68e445bf1c2_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_b68e445bf1c2)
		if var_b68e445bf1c2_err != nil {
			panic(var_b68e445bf1c2_err)
		}
		properties["createdOn"] = var_b68e445bf1c2_mapped
	}

	var_23529feae86c := securityConstraint.UpdatedOn

	if var_23529feae86c != nil {
		var var_23529feae86c_mapped *structpb.Value

		var var_23529feae86c_err error
		var_23529feae86c_mapped, var_23529feae86c_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_23529feae86c)
		if var_23529feae86c_err != nil {
			panic(var_23529feae86c_err)
		}
		properties["updatedOn"] = var_23529feae86c_mapped
	}

	var_19ae3d297cba := securityConstraint.Namespace

	if var_19ae3d297cba != nil {
		var var_19ae3d297cba_mapped *structpb.Value

		var_19ae3d297cba_mapped = structpb.NewStructValue(&structpb.Struct{Fields: NamespaceMapperInstance.ToProperties(var_19ae3d297cba)})
		properties["namespace"] = var_19ae3d297cba_mapped
	}

	var_6cb3b351a655 := securityConstraint.Resource

	if var_6cb3b351a655 != nil {
		var var_6cb3b351a655_mapped *structpb.Value

		var_6cb3b351a655_mapped = structpb.NewStructValue(&structpb.Struct{Fields: ResourceMapperInstance.ToProperties(var_6cb3b351a655)})
		properties["resource"] = var_6cb3b351a655_mapped
	}

	var_c14b34246518 := securityConstraint.Property

	if var_c14b34246518 != nil {
		var var_c14b34246518_mapped *structpb.Value

		var var_c14b34246518_err error
		var_c14b34246518_mapped, var_c14b34246518_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_c14b34246518)
		if var_c14b34246518_err != nil {
			panic(var_c14b34246518_err)
		}
		properties["property"] = var_c14b34246518_mapped
	}

	var_f3932ad7f7f8 := securityConstraint.PropertyValue

	if var_f3932ad7f7f8 != nil {
		var var_f3932ad7f7f8_mapped *structpb.Value

		var var_f3932ad7f7f8_err error
		var_f3932ad7f7f8_mapped, var_f3932ad7f7f8_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_f3932ad7f7f8)
		if var_f3932ad7f7f8_err != nil {
			panic(var_f3932ad7f7f8_err)
		}
		properties["propertyValue"] = var_f3932ad7f7f8_mapped
	}

	var_7db4622ca31a := securityConstraint.PropertyMode

	if var_7db4622ca31a != nil {
		var var_7db4622ca31a_mapped *structpb.Value

		var var_7db4622ca31a_err error
		var_7db4622ca31a_mapped, var_7db4622ca31a_err = types.ByResourcePropertyType(model.ResourceProperty_ENUM).Pack(string(*var_7db4622ca31a))
		if var_7db4622ca31a_err != nil {
			panic(var_7db4622ca31a_err)
		}
		properties["propertyMode"] = var_7db4622ca31a_mapped
	}

	var_f4459690e4f3 := securityConstraint.Operation

	var var_f4459690e4f3_mapped *structpb.Value

	var var_f4459690e4f3_err error
	var_f4459690e4f3_mapped, var_f4459690e4f3_err = types.ByResourcePropertyType(model.ResourceProperty_ENUM).Pack(string(var_f4459690e4f3))
	if var_f4459690e4f3_err != nil {
		panic(var_f4459690e4f3_err)
	}
	properties["operation"] = var_f4459690e4f3_mapped

	var_2ddff8f6f166 := securityConstraint.RecordIds

	if var_2ddff8f6f166 != nil {
		var var_2ddff8f6f166_mapped *structpb.Value

		var var_2ddff8f6f166_l []*structpb.Value
		for _, value := range var_2ddff8f6f166 {

			var_65f4bbe97486 := value
			var var_65f4bbe97486_mapped *structpb.Value

			var var_65f4bbe97486_err error
			var_65f4bbe97486_mapped, var_65f4bbe97486_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_65f4bbe97486)
			if var_65f4bbe97486_err != nil {
				panic(var_65f4bbe97486_err)
			}

			var_2ddff8f6f166_l = append(var_2ddff8f6f166_l, var_65f4bbe97486_mapped)
		}
		var_2ddff8f6f166_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_2ddff8f6f166_l})
		properties["recordIds"] = var_2ddff8f6f166_mapped
	}

	var_a6f5e03abbe1 := securityConstraint.Before

	if var_a6f5e03abbe1 != nil {
		var var_a6f5e03abbe1_mapped *structpb.Value

		var var_a6f5e03abbe1_err error
		var_a6f5e03abbe1_mapped, var_a6f5e03abbe1_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_a6f5e03abbe1)
		if var_a6f5e03abbe1_err != nil {
			panic(var_a6f5e03abbe1_err)
		}
		properties["before"] = var_a6f5e03abbe1_mapped
	}

	var_dc639f317400 := securityConstraint.After

	if var_dc639f317400 != nil {
		var var_dc639f317400_mapped *structpb.Value

		var var_dc639f317400_err error
		var_dc639f317400_mapped, var_dc639f317400_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_dc639f317400)
		if var_dc639f317400_err != nil {
			panic(var_dc639f317400_err)
		}
		properties["after"] = var_dc639f317400_mapped
	}

	var_5b44b811f669 := securityConstraint.User

	if var_5b44b811f669 != nil {
		var var_5b44b811f669_mapped *structpb.Value

		var_5b44b811f669_mapped = structpb.NewStructValue(&structpb.Struct{Fields: UserMapperInstance.ToProperties(var_5b44b811f669)})
		properties["user"] = var_5b44b811f669_mapped
	}

	var_f62162315b28 := securityConstraint.Role

	if var_f62162315b28 != nil {
		var var_f62162315b28_mapped *structpb.Value

		var_f62162315b28_mapped = structpb.NewStructValue(&structpb.Struct{Fields: RoleMapperInstance.ToProperties(var_f62162315b28)})
		properties["role"] = var_f62162315b28_mapped
	}

	var_4be7bcb6c2f3 := securityConstraint.Permit

	var var_4be7bcb6c2f3_mapped *structpb.Value

	var var_4be7bcb6c2f3_err error
	var_4be7bcb6c2f3_mapped, var_4be7bcb6c2f3_err = types.ByResourcePropertyType(model.ResourceProperty_ENUM).Pack(string(var_4be7bcb6c2f3))
	if var_4be7bcb6c2f3_err != nil {
		panic(var_4be7bcb6c2f3_err)
	}
	properties["permit"] = var_4be7bcb6c2f3_mapped

	var_14fc7b27c9a3 := securityConstraint.LocalFlags

	if var_14fc7b27c9a3 != nil {
		var var_14fc7b27c9a3_mapped *structpb.Value

		var var_14fc7b27c9a3_err error
		var_14fc7b27c9a3_mapped, var_14fc7b27c9a3_err = types.ByResourcePropertyType(model.ResourceProperty_OBJECT).Pack(*var_14fc7b27c9a3)
		if var_14fc7b27c9a3_err != nil {
			panic(var_14fc7b27c9a3_err)
		}
		properties["localFlags"] = var_14fc7b27c9a3_mapped
	}
	return properties
}

func (m *SecurityConstraintMapper) FromProperties(properties map[string]*structpb.Value) *SecurityConstraint {
	var s = m.New()
	if properties["id"] != nil {

		var_a2e1bc5339c3 := properties["id"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_UUID).UnPack(var_a2e1bc5339c3)

		if err != nil {
			panic(err)
		}

		var_a2e1bc5339c3_mapped := new(uuid.UUID)
		*var_a2e1bc5339c3_mapped = val.(uuid.UUID)

		s.Id = var_a2e1bc5339c3_mapped
	}
	if properties["version"] != nil {

		var_b0a5cb20c9c5 := properties["version"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(var_b0a5cb20c9c5)

		if err != nil {
			panic(err)
		}

		var_b0a5cb20c9c5_mapped := val.(int32)

		s.Version = var_b0a5cb20c9c5_mapped
	}
	if properties["createdBy"] != nil {

		var_48f06e6f80e5 := properties["createdBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_48f06e6f80e5)

		if err != nil {
			panic(err)
		}

		var_48f06e6f80e5_mapped := new(string)
		*var_48f06e6f80e5_mapped = val.(string)

		s.CreatedBy = var_48f06e6f80e5_mapped
	}
	if properties["updatedBy"] != nil {

		var_84fb06e079ea := properties["updatedBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_84fb06e079ea)

		if err != nil {
			panic(err)
		}

		var_84fb06e079ea_mapped := new(string)
		*var_84fb06e079ea_mapped = val.(string)

		s.UpdatedBy = var_84fb06e079ea_mapped
	}
	if properties["createdOn"] != nil {

		var_0d80c51ef80a := properties["createdOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_0d80c51ef80a)

		if err != nil {
			panic(err)
		}

		var_0d80c51ef80a_mapped := new(time.Time)
		*var_0d80c51ef80a_mapped = val.(time.Time)

		s.CreatedOn = var_0d80c51ef80a_mapped
	}
	if properties["updatedOn"] != nil {

		var_9dc1da1465ab := properties["updatedOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_9dc1da1465ab)

		if err != nil {
			panic(err)
		}

		var_9dc1da1465ab_mapped := new(time.Time)
		*var_9dc1da1465ab_mapped = val.(time.Time)

		s.UpdatedOn = var_9dc1da1465ab_mapped
	}
	if properties["namespace"] != nil {

		var_4f6bb6e48351 := properties["namespace"]
		var_4f6bb6e48351_mapped := NamespaceMapperInstance.FromProperties(var_4f6bb6e48351.GetStructValue().Fields)

		s.Namespace = var_4f6bb6e48351_mapped
	}
	if properties["resource"] != nil {

		var_a2f1506bcc23 := properties["resource"]
		var_a2f1506bcc23_mapped := ResourceMapperInstance.FromProperties(var_a2f1506bcc23.GetStructValue().Fields)

		s.Resource = var_a2f1506bcc23_mapped
	}
	if properties["property"] != nil {

		var_5bb1ab784606 := properties["property"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_5bb1ab784606)

		if err != nil {
			panic(err)
		}

		var_5bb1ab784606_mapped := new(string)
		*var_5bb1ab784606_mapped = val.(string)

		s.Property = var_5bb1ab784606_mapped
	}
	if properties["propertyValue"] != nil {

		var_c3447c60dd24 := properties["propertyValue"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_c3447c60dd24)

		if err != nil {
			panic(err)
		}

		var_c3447c60dd24_mapped := new(string)
		*var_c3447c60dd24_mapped = val.(string)

		s.PropertyValue = var_c3447c60dd24_mapped
	}
	if properties["propertyMode"] != nil {

		var_3952f0a559f0 := properties["propertyMode"]
		var_3952f0a559f0_mapped := new(SecurityConstraintPropertyMode)
		*var_3952f0a559f0_mapped = (SecurityConstraintPropertyMode)(var_3952f0a559f0.GetStringValue())

		s.PropertyMode = var_3952f0a559f0_mapped
	}
	if properties["operation"] != nil {

		var_fb91be04d4ac := properties["operation"]
		var_fb91be04d4ac_mapped := (SecurityConstraintOperation)(var_fb91be04d4ac.GetStringValue())

		s.Operation = var_fb91be04d4ac_mapped
	}
	if properties["recordIds"] != nil {

		var_ed8c8a52bebf := properties["recordIds"]
		var_ed8c8a52bebf_mapped := []string{}
		for _, v := range var_ed8c8a52bebf.GetListValue().Values {

			var_981c1de4db48 := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_981c1de4db48)

			if err != nil {
				panic(err)
			}

			var_981c1de4db48_mapped := val.(string)

			var_ed8c8a52bebf_mapped = append(var_ed8c8a52bebf_mapped, var_981c1de4db48_mapped)
		}

		s.RecordIds = var_ed8c8a52bebf_mapped
	}
	if properties["before"] != nil {

		var_f3f010ffba86 := properties["before"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_f3f010ffba86)

		if err != nil {
			panic(err)
		}

		var_f3f010ffba86_mapped := new(time.Time)
		*var_f3f010ffba86_mapped = val.(time.Time)

		s.Before = var_f3f010ffba86_mapped
	}
	if properties["after"] != nil {

		var_96a02983b597 := properties["after"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_96a02983b597)

		if err != nil {
			panic(err)
		}

		var_96a02983b597_mapped := new(time.Time)
		*var_96a02983b597_mapped = val.(time.Time)

		s.After = var_96a02983b597_mapped
	}
	if properties["user"] != nil {

		var_b66560a5f174 := properties["user"]
		var_b66560a5f174_mapped := UserMapperInstance.FromProperties(var_b66560a5f174.GetStructValue().Fields)

		s.User = var_b66560a5f174_mapped
	}
	if properties["role"] != nil {

		var_ddb17da833f0 := properties["role"]
		var_ddb17da833f0_mapped := RoleMapperInstance.FromProperties(var_ddb17da833f0.GetStructValue().Fields)

		s.Role = var_ddb17da833f0_mapped
	}
	if properties["permit"] != nil {

		var_3a7053b04209 := properties["permit"]
		var_3a7053b04209_mapped := (SecurityConstraintPermit)(var_3a7053b04209.GetStringValue())

		s.Permit = var_3a7053b04209_mapped
	}
	if properties["localFlags"] != nil {

		var_2d99848ea484 := properties["localFlags"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_OBJECT).UnPack(var_2d99848ea484)

		if err != nil {
			panic(err)
		}

		var_2d99848ea484_mapped := new(unstructured.Unstructured)
		*var_2d99848ea484_mapped = val.(unstructured.Unstructured)

		s.LocalFlags = var_2d99848ea484_mapped
	}
	return s
}
