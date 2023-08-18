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

	var_9724f62a3fe6 := securityConstraint.Id

	if var_9724f62a3fe6 != nil {
		var var_9724f62a3fe6_mapped *structpb.Value

		var var_9724f62a3fe6_err error
		var_9724f62a3fe6_mapped, var_9724f62a3fe6_err = types.ByResourcePropertyType(model.ResourceProperty_UUID).Pack(*var_9724f62a3fe6)
		if var_9724f62a3fe6_err != nil {
			panic(var_9724f62a3fe6_err)
		}
		properties["id"] = var_9724f62a3fe6_mapped
	}

	var_b46eba942d75 := securityConstraint.Version

	var var_b46eba942d75_mapped *structpb.Value

	var var_b46eba942d75_err error
	var_b46eba942d75_mapped, var_b46eba942d75_err = types.ByResourcePropertyType(model.ResourceProperty_INT32).Pack(var_b46eba942d75)
	if var_b46eba942d75_err != nil {
		panic(var_b46eba942d75_err)
	}
	properties["version"] = var_b46eba942d75_mapped

	var_81e2c69ca863 := securityConstraint.CreatedBy

	if var_81e2c69ca863 != nil {
		var var_81e2c69ca863_mapped *structpb.Value

		var var_81e2c69ca863_err error
		var_81e2c69ca863_mapped, var_81e2c69ca863_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_81e2c69ca863)
		if var_81e2c69ca863_err != nil {
			panic(var_81e2c69ca863_err)
		}
		properties["createdBy"] = var_81e2c69ca863_mapped
	}

	var_32c0fa93f24a := securityConstraint.UpdatedBy

	if var_32c0fa93f24a != nil {
		var var_32c0fa93f24a_mapped *structpb.Value

		var var_32c0fa93f24a_err error
		var_32c0fa93f24a_mapped, var_32c0fa93f24a_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_32c0fa93f24a)
		if var_32c0fa93f24a_err != nil {
			panic(var_32c0fa93f24a_err)
		}
		properties["updatedBy"] = var_32c0fa93f24a_mapped
	}

	var_d27fbbc50c66 := securityConstraint.CreatedOn

	if var_d27fbbc50c66 != nil {
		var var_d27fbbc50c66_mapped *structpb.Value

		var var_d27fbbc50c66_err error
		var_d27fbbc50c66_mapped, var_d27fbbc50c66_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_d27fbbc50c66)
		if var_d27fbbc50c66_err != nil {
			panic(var_d27fbbc50c66_err)
		}
		properties["createdOn"] = var_d27fbbc50c66_mapped
	}

	var_795e75b3b2d1 := securityConstraint.UpdatedOn

	if var_795e75b3b2d1 != nil {
		var var_795e75b3b2d1_mapped *structpb.Value

		var var_795e75b3b2d1_err error
		var_795e75b3b2d1_mapped, var_795e75b3b2d1_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_795e75b3b2d1)
		if var_795e75b3b2d1_err != nil {
			panic(var_795e75b3b2d1_err)
		}
		properties["updatedOn"] = var_795e75b3b2d1_mapped
	}

	var_a8ba40bbd601 := securityConstraint.Namespace

	if var_a8ba40bbd601 != nil {
		var var_a8ba40bbd601_mapped *structpb.Value

		var var_a8ba40bbd601_err error
		var_a8ba40bbd601_mapped, var_a8ba40bbd601_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_a8ba40bbd601)
		if var_a8ba40bbd601_err != nil {
			panic(var_a8ba40bbd601_err)
		}
		properties["namespace"] = var_a8ba40bbd601_mapped
	}

	var_0b6d249abcdc := securityConstraint.Resource

	if var_0b6d249abcdc != nil {
		var var_0b6d249abcdc_mapped *structpb.Value

		var var_0b6d249abcdc_err error
		var_0b6d249abcdc_mapped, var_0b6d249abcdc_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_0b6d249abcdc)
		if var_0b6d249abcdc_err != nil {
			panic(var_0b6d249abcdc_err)
		}
		properties["resource"] = var_0b6d249abcdc_mapped
	}

	var_900c419ab097 := securityConstraint.Property

	if var_900c419ab097 != nil {
		var var_900c419ab097_mapped *structpb.Value

		var var_900c419ab097_err error
		var_900c419ab097_mapped, var_900c419ab097_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_900c419ab097)
		if var_900c419ab097_err != nil {
			panic(var_900c419ab097_err)
		}
		properties["property"] = var_900c419ab097_mapped
	}

	var_f61417810a88 := securityConstraint.PropertyValue

	if var_f61417810a88 != nil {
		var var_f61417810a88_mapped *structpb.Value

		var var_f61417810a88_err error
		var_f61417810a88_mapped, var_f61417810a88_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_f61417810a88)
		if var_f61417810a88_err != nil {
			panic(var_f61417810a88_err)
		}
		properties["propertyValue"] = var_f61417810a88_mapped
	}

	var_cb3978af1618 := securityConstraint.PropertyMode

	if var_cb3978af1618 != nil {
		var var_cb3978af1618_mapped *structpb.Value

		var var_cb3978af1618_err error
		var_cb3978af1618_mapped, var_cb3978af1618_err = types.ByResourcePropertyType(model.ResourceProperty_ENUM).Pack(string(*var_cb3978af1618))
		if var_cb3978af1618_err != nil {
			panic(var_cb3978af1618_err)
		}
		properties["propertyMode"] = var_cb3978af1618_mapped
	}

	var_8c05e31ffa8b := securityConstraint.Operation

	var var_8c05e31ffa8b_mapped *structpb.Value

	var var_8c05e31ffa8b_err error
	var_8c05e31ffa8b_mapped, var_8c05e31ffa8b_err = types.ByResourcePropertyType(model.ResourceProperty_ENUM).Pack(string(var_8c05e31ffa8b))
	if var_8c05e31ffa8b_err != nil {
		panic(var_8c05e31ffa8b_err)
	}
	properties["operation"] = var_8c05e31ffa8b_mapped

	var_34c20d5d2f8c := securityConstraint.RecordIds

	if var_34c20d5d2f8c != nil {
		var var_34c20d5d2f8c_mapped *structpb.Value

		var var_34c20d5d2f8c_l []*structpb.Value
		for _, value := range var_34c20d5d2f8c {

			var_ff3d9188901a := value
			var var_ff3d9188901a_mapped *structpb.Value

			var var_ff3d9188901a_err error
			var_ff3d9188901a_mapped, var_ff3d9188901a_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_ff3d9188901a)
			if var_ff3d9188901a_err != nil {
				panic(var_ff3d9188901a_err)
			}

			var_34c20d5d2f8c_l = append(var_34c20d5d2f8c_l, var_ff3d9188901a_mapped)
		}
		var_34c20d5d2f8c_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_34c20d5d2f8c_l})
		properties["recordIds"] = var_34c20d5d2f8c_mapped
	}

	var_044076246342 := securityConstraint.Before

	if var_044076246342 != nil {
		var var_044076246342_mapped *structpb.Value

		var var_044076246342_err error
		var_044076246342_mapped, var_044076246342_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_044076246342)
		if var_044076246342_err != nil {
			panic(var_044076246342_err)
		}
		properties["before"] = var_044076246342_mapped
	}

	var_be13a83d6816 := securityConstraint.After

	if var_be13a83d6816 != nil {
		var var_be13a83d6816_mapped *structpb.Value

		var var_be13a83d6816_err error
		var_be13a83d6816_mapped, var_be13a83d6816_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_be13a83d6816)
		if var_be13a83d6816_err != nil {
			panic(var_be13a83d6816_err)
		}
		properties["after"] = var_be13a83d6816_mapped
	}

	var_82deabedf186 := securityConstraint.User

	if var_82deabedf186 != nil {
		var var_82deabedf186_mapped *structpb.Value

		var_82deabedf186_mapped = structpb.NewStructValue(&structpb.Struct{Fields: UserMapperInstance.ToProperties(var_82deabedf186)})
		properties["user"] = var_82deabedf186_mapped
	}

	var_cfc82588aa5a := securityConstraint.Role

	if var_cfc82588aa5a != nil {
		var var_cfc82588aa5a_mapped *structpb.Value

		var_cfc82588aa5a_mapped = structpb.NewStructValue(&structpb.Struct{Fields: RoleMapperInstance.ToProperties(var_cfc82588aa5a)})
		properties["role"] = var_cfc82588aa5a_mapped
	}

	var_ba613a23a729 := securityConstraint.Permit

	var var_ba613a23a729_mapped *structpb.Value

	var var_ba613a23a729_err error
	var_ba613a23a729_mapped, var_ba613a23a729_err = types.ByResourcePropertyType(model.ResourceProperty_ENUM).Pack(string(var_ba613a23a729))
	if var_ba613a23a729_err != nil {
		panic(var_ba613a23a729_err)
	}
	properties["permit"] = var_ba613a23a729_mapped

	var_ed2a4959ba87 := securityConstraint.LocalFlags

	if var_ed2a4959ba87 != nil {
		var var_ed2a4959ba87_mapped *structpb.Value

		var var_ed2a4959ba87_err error
		var_ed2a4959ba87_mapped, var_ed2a4959ba87_err = types.ByResourcePropertyType(model.ResourceProperty_OBJECT).Pack(*var_ed2a4959ba87)
		if var_ed2a4959ba87_err != nil {
			panic(var_ed2a4959ba87_err)
		}
		properties["localFlags"] = var_ed2a4959ba87_mapped
	}
	return properties
}

func (m *SecurityConstraintMapper) FromProperties(properties map[string]*structpb.Value) *SecurityConstraint {
	var s = m.New()
	if properties["id"] != nil {

		var_be0777f73576 := properties["id"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_UUID).UnPack(var_be0777f73576)

		if err != nil {
			panic(err)
		}

		var_be0777f73576_mapped := new(uuid.UUID)
		*var_be0777f73576_mapped = val.(uuid.UUID)

		s.Id = var_be0777f73576_mapped
	}
	if properties["version"] != nil {

		var_88859cef4bfd := properties["version"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(var_88859cef4bfd)

		if err != nil {
			panic(err)
		}

		var_88859cef4bfd_mapped := val.(int32)

		s.Version = var_88859cef4bfd_mapped
	}
	if properties["createdBy"] != nil {

		var_ad3a93db53e1 := properties["createdBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_ad3a93db53e1)

		if err != nil {
			panic(err)
		}

		var_ad3a93db53e1_mapped := new(string)
		*var_ad3a93db53e1_mapped = val.(string)

		s.CreatedBy = var_ad3a93db53e1_mapped
	}
	if properties["updatedBy"] != nil {

		var_1f0404ebb765 := properties["updatedBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_1f0404ebb765)

		if err != nil {
			panic(err)
		}

		var_1f0404ebb765_mapped := new(string)
		*var_1f0404ebb765_mapped = val.(string)

		s.UpdatedBy = var_1f0404ebb765_mapped
	}
	if properties["createdOn"] != nil {

		var_c9324d746b14 := properties["createdOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_c9324d746b14)

		if err != nil {
			panic(err)
		}

		var_c9324d746b14_mapped := new(time.Time)
		*var_c9324d746b14_mapped = val.(time.Time)

		s.CreatedOn = var_c9324d746b14_mapped
	}
	if properties["updatedOn"] != nil {

		var_f13eab58a9b2 := properties["updatedOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_f13eab58a9b2)

		if err != nil {
			panic(err)
		}

		var_f13eab58a9b2_mapped := new(time.Time)
		*var_f13eab58a9b2_mapped = val.(time.Time)

		s.UpdatedOn = var_f13eab58a9b2_mapped
	}
	if properties["namespace"] != nil {

		var_fcd884d2a502 := properties["namespace"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_fcd884d2a502)

		if err != nil {
			panic(err)
		}

		var_fcd884d2a502_mapped := new(string)
		*var_fcd884d2a502_mapped = val.(string)

		s.Namespace = var_fcd884d2a502_mapped
	}
	if properties["resource"] != nil {

		var_28bf9712d40b := properties["resource"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_28bf9712d40b)

		if err != nil {
			panic(err)
		}

		var_28bf9712d40b_mapped := new(string)
		*var_28bf9712d40b_mapped = val.(string)

		s.Resource = var_28bf9712d40b_mapped
	}
	if properties["property"] != nil {

		var_aa838cd13835 := properties["property"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_aa838cd13835)

		if err != nil {
			panic(err)
		}

		var_aa838cd13835_mapped := new(string)
		*var_aa838cd13835_mapped = val.(string)

		s.Property = var_aa838cd13835_mapped
	}
	if properties["propertyValue"] != nil {

		var_f2259f065a2c := properties["propertyValue"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_f2259f065a2c)

		if err != nil {
			panic(err)
		}

		var_f2259f065a2c_mapped := new(string)
		*var_f2259f065a2c_mapped = val.(string)

		s.PropertyValue = var_f2259f065a2c_mapped
	}
	if properties["propertyMode"] != nil {

		var_b821ce8999ef := properties["propertyMode"]
		var_b821ce8999ef_mapped := new(SecurityConstraintPropertyMode)
		*var_b821ce8999ef_mapped = (SecurityConstraintPropertyMode)(var_b821ce8999ef.GetStringValue())

		s.PropertyMode = var_b821ce8999ef_mapped
	}
	if properties["operation"] != nil {

		var_f76220f3d45b := properties["operation"]
		var_f76220f3d45b_mapped := (SecurityConstraintOperation)(var_f76220f3d45b.GetStringValue())

		s.Operation = var_f76220f3d45b_mapped
	}
	if properties["recordIds"] != nil {

		var_f9914c89ce9f := properties["recordIds"]
		var_f9914c89ce9f_mapped := []string{}
		for _, v := range var_f9914c89ce9f.GetListValue().Values {

			var_f775e4060670 := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_f775e4060670)

			if err != nil {
				panic(err)
			}

			var_f775e4060670_mapped := val.(string)

			var_f9914c89ce9f_mapped = append(var_f9914c89ce9f_mapped, var_f775e4060670_mapped)
		}

		s.RecordIds = var_f9914c89ce9f_mapped
	}
	if properties["before"] != nil {

		var_43195def4e01 := properties["before"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_43195def4e01)

		if err != nil {
			panic(err)
		}

		var_43195def4e01_mapped := new(time.Time)
		*var_43195def4e01_mapped = val.(time.Time)

		s.Before = var_43195def4e01_mapped
	}
	if properties["after"] != nil {

		var_5a981ac7346d := properties["after"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_5a981ac7346d)

		if err != nil {
			panic(err)
		}

		var_5a981ac7346d_mapped := new(time.Time)
		*var_5a981ac7346d_mapped = val.(time.Time)

		s.After = var_5a981ac7346d_mapped
	}
	if properties["user"] != nil {

		var_0ddf94191a61 := properties["user"]
		var_0ddf94191a61_mapped := UserMapperInstance.FromProperties(var_0ddf94191a61.GetStructValue().Fields)

		s.User = var_0ddf94191a61_mapped
	}
	if properties["role"] != nil {

		var_311fb1355bea := properties["role"]
		var_311fb1355bea_mapped := RoleMapperInstance.FromProperties(var_311fb1355bea.GetStructValue().Fields)

		s.Role = var_311fb1355bea_mapped
	}
	if properties["permit"] != nil {

		var_c6b2946568d0 := properties["permit"]
		var_c6b2946568d0_mapped := (SecurityConstraintPermit)(var_c6b2946568d0.GetStringValue())

		s.Permit = var_c6b2946568d0_mapped
	}
	if properties["localFlags"] != nil {

		var_85f56fde1e49 := properties["localFlags"]
		var_85f56fde1e49_mapped := new(unstructured.Unstructured)
		*var_85f56fde1e49_mapped = unstructured.FromStructValue(var_85f56fde1e49.GetStructValue())

		s.LocalFlags = var_85f56fde1e49_mapped
	}
	return s
}
