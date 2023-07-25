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

	var_7ac08613029d := securityConstraint.Id

	if var_7ac08613029d != nil {
		var var_7ac08613029d_mapped *structpb.Value

		var var_7ac08613029d_err error
		var_7ac08613029d_mapped, var_7ac08613029d_err = types.ByResourcePropertyType(model.ResourceProperty_UUID).Pack(*var_7ac08613029d)
		if var_7ac08613029d_err != nil {
			panic(var_7ac08613029d_err)
		}
		properties["id"] = var_7ac08613029d_mapped
	}

	var_95de0e4382a6 := securityConstraint.Version

	var var_95de0e4382a6_mapped *structpb.Value

	var var_95de0e4382a6_err error
	var_95de0e4382a6_mapped, var_95de0e4382a6_err = types.ByResourcePropertyType(model.ResourceProperty_INT32).Pack(var_95de0e4382a6)
	if var_95de0e4382a6_err != nil {
		panic(var_95de0e4382a6_err)
	}
	properties["version"] = var_95de0e4382a6_mapped

	var_2b1f3b6e7603 := securityConstraint.CreatedBy

	if var_2b1f3b6e7603 != nil {
		var var_2b1f3b6e7603_mapped *structpb.Value

		var var_2b1f3b6e7603_err error
		var_2b1f3b6e7603_mapped, var_2b1f3b6e7603_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_2b1f3b6e7603)
		if var_2b1f3b6e7603_err != nil {
			panic(var_2b1f3b6e7603_err)
		}
		properties["createdBy"] = var_2b1f3b6e7603_mapped
	}

	var_9af7abb054ea := securityConstraint.UpdatedBy

	if var_9af7abb054ea != nil {
		var var_9af7abb054ea_mapped *structpb.Value

		var var_9af7abb054ea_err error
		var_9af7abb054ea_mapped, var_9af7abb054ea_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_9af7abb054ea)
		if var_9af7abb054ea_err != nil {
			panic(var_9af7abb054ea_err)
		}
		properties["updatedBy"] = var_9af7abb054ea_mapped
	}

	var_9be18f37c7b7 := securityConstraint.CreatedOn

	if var_9be18f37c7b7 != nil {
		var var_9be18f37c7b7_mapped *structpb.Value

		var var_9be18f37c7b7_err error
		var_9be18f37c7b7_mapped, var_9be18f37c7b7_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_9be18f37c7b7)
		if var_9be18f37c7b7_err != nil {
			panic(var_9be18f37c7b7_err)
		}
		properties["createdOn"] = var_9be18f37c7b7_mapped
	}

	var_b53908acabad := securityConstraint.UpdatedOn

	if var_b53908acabad != nil {
		var var_b53908acabad_mapped *structpb.Value

		var var_b53908acabad_err error
		var_b53908acabad_mapped, var_b53908acabad_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_b53908acabad)
		if var_b53908acabad_err != nil {
			panic(var_b53908acabad_err)
		}
		properties["updatedOn"] = var_b53908acabad_mapped
	}

	var_068a9b4633df := securityConstraint.Namespace

	if var_068a9b4633df != nil {
		var var_068a9b4633df_mapped *structpb.Value

		var_068a9b4633df_mapped = structpb.NewStructValue(&structpb.Struct{Fields: NamespaceMapperInstance.ToProperties(var_068a9b4633df)})
		properties["namespace"] = var_068a9b4633df_mapped
	}

	var_1e0ad11eb37f := securityConstraint.Resource

	if var_1e0ad11eb37f != nil {
		var var_1e0ad11eb37f_mapped *structpb.Value

		var_1e0ad11eb37f_mapped = structpb.NewStructValue(&structpb.Struct{Fields: ResourceMapperInstance.ToProperties(var_1e0ad11eb37f)})
		properties["resource"] = var_1e0ad11eb37f_mapped
	}

	var_b3d8ecf8b6ab := securityConstraint.Property

	if var_b3d8ecf8b6ab != nil {
		var var_b3d8ecf8b6ab_mapped *structpb.Value

		var var_b3d8ecf8b6ab_err error
		var_b3d8ecf8b6ab_mapped, var_b3d8ecf8b6ab_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_b3d8ecf8b6ab)
		if var_b3d8ecf8b6ab_err != nil {
			panic(var_b3d8ecf8b6ab_err)
		}
		properties["property"] = var_b3d8ecf8b6ab_mapped
	}

	var_15c828885338 := securityConstraint.PropertyValue

	if var_15c828885338 != nil {
		var var_15c828885338_mapped *structpb.Value

		var var_15c828885338_err error
		var_15c828885338_mapped, var_15c828885338_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_15c828885338)
		if var_15c828885338_err != nil {
			panic(var_15c828885338_err)
		}
		properties["propertyValue"] = var_15c828885338_mapped
	}

	var_d8f819f99022 := securityConstraint.PropertyMode

	if var_d8f819f99022 != nil {
		var var_d8f819f99022_mapped *structpb.Value

		var var_d8f819f99022_err error
		var_d8f819f99022_mapped, var_d8f819f99022_err = types.ByResourcePropertyType(model.ResourceProperty_ENUM).Pack(string(*var_d8f819f99022))
		if var_d8f819f99022_err != nil {
			panic(var_d8f819f99022_err)
		}
		properties["propertyMode"] = var_d8f819f99022_mapped
	}

	var_40716bc2ea49 := securityConstraint.Operation

	var var_40716bc2ea49_mapped *structpb.Value

	var var_40716bc2ea49_err error
	var_40716bc2ea49_mapped, var_40716bc2ea49_err = types.ByResourcePropertyType(model.ResourceProperty_ENUM).Pack(string(var_40716bc2ea49))
	if var_40716bc2ea49_err != nil {
		panic(var_40716bc2ea49_err)
	}
	properties["operation"] = var_40716bc2ea49_mapped

	var_cd1b4c77cfc2 := securityConstraint.RecordIds

	if var_cd1b4c77cfc2 != nil {
		var var_cd1b4c77cfc2_mapped *structpb.Value

		var var_cd1b4c77cfc2_l []*structpb.Value
		for _, value := range var_cd1b4c77cfc2 {

			var_38d2a9943049 := value
			var var_38d2a9943049_mapped *structpb.Value

			var var_38d2a9943049_err error
			var_38d2a9943049_mapped, var_38d2a9943049_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_38d2a9943049)
			if var_38d2a9943049_err != nil {
				panic(var_38d2a9943049_err)
			}

			var_cd1b4c77cfc2_l = append(var_cd1b4c77cfc2_l, var_38d2a9943049_mapped)
		}
		var_cd1b4c77cfc2_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_cd1b4c77cfc2_l})
		properties["recordIds"] = var_cd1b4c77cfc2_mapped
	}

	var_cb336f9037f1 := securityConstraint.Before

	if var_cb336f9037f1 != nil {
		var var_cb336f9037f1_mapped *structpb.Value

		var var_cb336f9037f1_err error
		var_cb336f9037f1_mapped, var_cb336f9037f1_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_cb336f9037f1)
		if var_cb336f9037f1_err != nil {
			panic(var_cb336f9037f1_err)
		}
		properties["before"] = var_cb336f9037f1_mapped
	}

	var_87030e53c24e := securityConstraint.After

	if var_87030e53c24e != nil {
		var var_87030e53c24e_mapped *structpb.Value

		var var_87030e53c24e_err error
		var_87030e53c24e_mapped, var_87030e53c24e_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_87030e53c24e)
		if var_87030e53c24e_err != nil {
			panic(var_87030e53c24e_err)
		}
		properties["after"] = var_87030e53c24e_mapped
	}

	var_3f0c4f66a69d := securityConstraint.User

	if var_3f0c4f66a69d != nil {
		var var_3f0c4f66a69d_mapped *structpb.Value

		var_3f0c4f66a69d_mapped = structpb.NewStructValue(&structpb.Struct{Fields: UserMapperInstance.ToProperties(var_3f0c4f66a69d)})
		properties["user"] = var_3f0c4f66a69d_mapped
	}

	var_7a572e3da083 := securityConstraint.Role

	if var_7a572e3da083 != nil {
		var var_7a572e3da083_mapped *structpb.Value

		var_7a572e3da083_mapped = structpb.NewStructValue(&structpb.Struct{Fields: RoleMapperInstance.ToProperties(var_7a572e3da083)})
		properties["role"] = var_7a572e3da083_mapped
	}

	var_4d4bf17a2d82 := securityConstraint.Permit

	var var_4d4bf17a2d82_mapped *structpb.Value

	var var_4d4bf17a2d82_err error
	var_4d4bf17a2d82_mapped, var_4d4bf17a2d82_err = types.ByResourcePropertyType(model.ResourceProperty_ENUM).Pack(string(var_4d4bf17a2d82))
	if var_4d4bf17a2d82_err != nil {
		panic(var_4d4bf17a2d82_err)
	}
	properties["permit"] = var_4d4bf17a2d82_mapped

	var_c475d115af44 := securityConstraint.LocalFlags

	if var_c475d115af44 != nil {
		var var_c475d115af44_mapped *structpb.Value

		var var_c475d115af44_err error
		var_c475d115af44_mapped, var_c475d115af44_err = types.ByResourcePropertyType(model.ResourceProperty_OBJECT).Pack(*var_c475d115af44)
		if var_c475d115af44_err != nil {
			panic(var_c475d115af44_err)
		}
		properties["localFlags"] = var_c475d115af44_mapped
	}
	return properties
}

func (m *SecurityConstraintMapper) FromProperties(properties map[string]*structpb.Value) *SecurityConstraint {
	var s = m.New()
	if properties["id"] != nil {

		var_90a1ca0aca3f := properties["id"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_UUID).UnPack(var_90a1ca0aca3f)

		if err != nil {
			panic(err)
		}

		var_90a1ca0aca3f_mapped := new(uuid.UUID)
		*var_90a1ca0aca3f_mapped = val.(uuid.UUID)

		s.Id = var_90a1ca0aca3f_mapped
	}
	if properties["version"] != nil {

		var_aa4bcbe610a8 := properties["version"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(var_aa4bcbe610a8)

		if err != nil {
			panic(err)
		}

		var_aa4bcbe610a8_mapped := val.(int32)

		s.Version = var_aa4bcbe610a8_mapped
	}
	if properties["createdBy"] != nil {

		var_ec9d3ed938e3 := properties["createdBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_ec9d3ed938e3)

		if err != nil {
			panic(err)
		}

		var_ec9d3ed938e3_mapped := new(string)
		*var_ec9d3ed938e3_mapped = val.(string)

		s.CreatedBy = var_ec9d3ed938e3_mapped
	}
	if properties["updatedBy"] != nil {

		var_b4562ff4f072 := properties["updatedBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_b4562ff4f072)

		if err != nil {
			panic(err)
		}

		var_b4562ff4f072_mapped := new(string)
		*var_b4562ff4f072_mapped = val.(string)

		s.UpdatedBy = var_b4562ff4f072_mapped
	}
	if properties["createdOn"] != nil {

		var_2e9fa4a84482 := properties["createdOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_2e9fa4a84482)

		if err != nil {
			panic(err)
		}

		var_2e9fa4a84482_mapped := new(time.Time)
		*var_2e9fa4a84482_mapped = val.(time.Time)

		s.CreatedOn = var_2e9fa4a84482_mapped
	}
	if properties["updatedOn"] != nil {

		var_ce3e7de3167b := properties["updatedOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_ce3e7de3167b)

		if err != nil {
			panic(err)
		}

		var_ce3e7de3167b_mapped := new(time.Time)
		*var_ce3e7de3167b_mapped = val.(time.Time)

		s.UpdatedOn = var_ce3e7de3167b_mapped
	}
	if properties["namespace"] != nil {

		var_b474e45053aa := properties["namespace"]
		var_b474e45053aa_mapped := NamespaceMapperInstance.FromProperties(var_b474e45053aa.GetStructValue().Fields)

		s.Namespace = var_b474e45053aa_mapped
	}
	if properties["resource"] != nil {

		var_e36df70325e5 := properties["resource"]
		var_e36df70325e5_mapped := ResourceMapperInstance.FromProperties(var_e36df70325e5.GetStructValue().Fields)

		s.Resource = var_e36df70325e5_mapped
	}
	if properties["property"] != nil {

		var_aa2d13ed95a5 := properties["property"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_aa2d13ed95a5)

		if err != nil {
			panic(err)
		}

		var_aa2d13ed95a5_mapped := new(string)
		*var_aa2d13ed95a5_mapped = val.(string)

		s.Property = var_aa2d13ed95a5_mapped
	}
	if properties["propertyValue"] != nil {

		var_7800567acd31 := properties["propertyValue"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_7800567acd31)

		if err != nil {
			panic(err)
		}

		var_7800567acd31_mapped := new(string)
		*var_7800567acd31_mapped = val.(string)

		s.PropertyValue = var_7800567acd31_mapped
	}
	if properties["propertyMode"] != nil {

		var_c46d89bdc84a := properties["propertyMode"]
		var_c46d89bdc84a_mapped := new(SecurityConstraintPropertyMode)
		*var_c46d89bdc84a_mapped = (SecurityConstraintPropertyMode)(var_c46d89bdc84a.GetStringValue())

		s.PropertyMode = var_c46d89bdc84a_mapped
	}
	if properties["operation"] != nil {

		var_a05132ff9189 := properties["operation"]
		var_a05132ff9189_mapped := (SecurityConstraintOperation)(var_a05132ff9189.GetStringValue())

		s.Operation = var_a05132ff9189_mapped
	}
	if properties["recordIds"] != nil {

		var_b2a09ef71a90 := properties["recordIds"]
		var_b2a09ef71a90_mapped := []string{}
		for _, v := range var_b2a09ef71a90.GetListValue().Values {

			var_40e27e478101 := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_40e27e478101)

			if err != nil {
				panic(err)
			}

			var_40e27e478101_mapped := val.(string)

			var_b2a09ef71a90_mapped = append(var_b2a09ef71a90_mapped, var_40e27e478101_mapped)
		}

		s.RecordIds = var_b2a09ef71a90_mapped
	}
	if properties["before"] != nil {

		var_18351660578f := properties["before"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_18351660578f)

		if err != nil {
			panic(err)
		}

		var_18351660578f_mapped := new(time.Time)
		*var_18351660578f_mapped = val.(time.Time)

		s.Before = var_18351660578f_mapped
	}
	if properties["after"] != nil {

		var_45c07625fc50 := properties["after"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_45c07625fc50)

		if err != nil {
			panic(err)
		}

		var_45c07625fc50_mapped := new(time.Time)
		*var_45c07625fc50_mapped = val.(time.Time)

		s.After = var_45c07625fc50_mapped
	}
	if properties["user"] != nil {

		var_bfbc9531d3b8 := properties["user"]
		var_bfbc9531d3b8_mapped := UserMapperInstance.FromProperties(var_bfbc9531d3b8.GetStructValue().Fields)

		s.User = var_bfbc9531d3b8_mapped
	}
	if properties["role"] != nil {

		var_f3ce529fcc1c := properties["role"]
		var_f3ce529fcc1c_mapped := RoleMapperInstance.FromProperties(var_f3ce529fcc1c.GetStructValue().Fields)

		s.Role = var_f3ce529fcc1c_mapped
	}
	if properties["permit"] != nil {

		var_eca33e1e3781 := properties["permit"]
		var_eca33e1e3781_mapped := (SecurityConstraintPermit)(var_eca33e1e3781.GetStringValue())

		s.Permit = var_eca33e1e3781_mapped
	}
	if properties["localFlags"] != nil {

		var_b1740c5d1ef5 := properties["localFlags"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_OBJECT).UnPack(var_b1740c5d1ef5)

		if err != nil {
			panic(err)
		}

		var_b1740c5d1ef5_mapped := new(unstructured.Unstructured)
		*var_b1740c5d1ef5_mapped = val.(unstructured.Unstructured)

		s.LocalFlags = var_b1740c5d1ef5_mapped
	}
	return s
}
