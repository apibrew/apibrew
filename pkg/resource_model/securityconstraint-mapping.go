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

	var_380ac3a33e05 := securityConstraint.Id

	if var_380ac3a33e05 != nil {
		var var_380ac3a33e05_mapped *structpb.Value

		var var_380ac3a33e05_err error
		var_380ac3a33e05_mapped, var_380ac3a33e05_err = types.ByResourcePropertyType(model.ResourceProperty_UUID).Pack(*var_380ac3a33e05)
		if var_380ac3a33e05_err != nil {
			panic(var_380ac3a33e05_err)
		}
		properties["id"] = var_380ac3a33e05_mapped
	}

	var_135f02c54faf := securityConstraint.Version

	var var_135f02c54faf_mapped *structpb.Value

	var var_135f02c54faf_err error
	var_135f02c54faf_mapped, var_135f02c54faf_err = types.ByResourcePropertyType(model.ResourceProperty_INT32).Pack(var_135f02c54faf)
	if var_135f02c54faf_err != nil {
		panic(var_135f02c54faf_err)
	}
	properties["version"] = var_135f02c54faf_mapped

	var_10162ef74bae := securityConstraint.CreatedBy

	if var_10162ef74bae != nil {
		var var_10162ef74bae_mapped *structpb.Value

		var var_10162ef74bae_err error
		var_10162ef74bae_mapped, var_10162ef74bae_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_10162ef74bae)
		if var_10162ef74bae_err != nil {
			panic(var_10162ef74bae_err)
		}
		properties["createdBy"] = var_10162ef74bae_mapped
	}

	var_8a4c717e6a98 := securityConstraint.UpdatedBy

	if var_8a4c717e6a98 != nil {
		var var_8a4c717e6a98_mapped *structpb.Value

		var var_8a4c717e6a98_err error
		var_8a4c717e6a98_mapped, var_8a4c717e6a98_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_8a4c717e6a98)
		if var_8a4c717e6a98_err != nil {
			panic(var_8a4c717e6a98_err)
		}
		properties["updatedBy"] = var_8a4c717e6a98_mapped
	}

	var_9f8befb67fa6 := securityConstraint.CreatedOn

	if var_9f8befb67fa6 != nil {
		var var_9f8befb67fa6_mapped *structpb.Value

		var var_9f8befb67fa6_err error
		var_9f8befb67fa6_mapped, var_9f8befb67fa6_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_9f8befb67fa6)
		if var_9f8befb67fa6_err != nil {
			panic(var_9f8befb67fa6_err)
		}
		properties["createdOn"] = var_9f8befb67fa6_mapped
	}

	var_5d46b3a70077 := securityConstraint.UpdatedOn

	if var_5d46b3a70077 != nil {
		var var_5d46b3a70077_mapped *structpb.Value

		var var_5d46b3a70077_err error
		var_5d46b3a70077_mapped, var_5d46b3a70077_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_5d46b3a70077)
		if var_5d46b3a70077_err != nil {
			panic(var_5d46b3a70077_err)
		}
		properties["updatedOn"] = var_5d46b3a70077_mapped
	}

	var_4f6240772e32 := securityConstraint.Namespace

	if var_4f6240772e32 != nil {
		var var_4f6240772e32_mapped *structpb.Value

		var_4f6240772e32_mapped = structpb.NewStructValue(&structpb.Struct{Fields: NamespaceMapperInstance.ToProperties(var_4f6240772e32)})
		properties["namespace"] = var_4f6240772e32_mapped
	}

	var_cb96caa12c34 := securityConstraint.Resource

	if var_cb96caa12c34 != nil {
		var var_cb96caa12c34_mapped *structpb.Value

		var_cb96caa12c34_mapped = structpb.NewStructValue(&structpb.Struct{Fields: ResourceMapperInstance.ToProperties(var_cb96caa12c34)})
		properties["resource"] = var_cb96caa12c34_mapped
	}

	var_2e264a7ee30e := securityConstraint.Property

	if var_2e264a7ee30e != nil {
		var var_2e264a7ee30e_mapped *structpb.Value

		var var_2e264a7ee30e_err error
		var_2e264a7ee30e_mapped, var_2e264a7ee30e_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_2e264a7ee30e)
		if var_2e264a7ee30e_err != nil {
			panic(var_2e264a7ee30e_err)
		}
		properties["property"] = var_2e264a7ee30e_mapped
	}

	var_f7ffa8363605 := securityConstraint.PropertyValue

	if var_f7ffa8363605 != nil {
		var var_f7ffa8363605_mapped *structpb.Value

		var var_f7ffa8363605_err error
		var_f7ffa8363605_mapped, var_f7ffa8363605_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_f7ffa8363605)
		if var_f7ffa8363605_err != nil {
			panic(var_f7ffa8363605_err)
		}
		properties["propertyValue"] = var_f7ffa8363605_mapped
	}

	var_f7f32051fd8d := securityConstraint.PropertyMode

	if var_f7f32051fd8d != nil {
		var var_f7f32051fd8d_mapped *structpb.Value

		var var_f7f32051fd8d_err error
		var_f7f32051fd8d_mapped, var_f7f32051fd8d_err = types.ByResourcePropertyType(model.ResourceProperty_ENUM).Pack(string(*var_f7f32051fd8d))
		if var_f7f32051fd8d_err != nil {
			panic(var_f7f32051fd8d_err)
		}
		properties["propertyMode"] = var_f7f32051fd8d_mapped
	}

	var_ce300970c22e := securityConstraint.Operation

	var var_ce300970c22e_mapped *structpb.Value

	var var_ce300970c22e_err error
	var_ce300970c22e_mapped, var_ce300970c22e_err = types.ByResourcePropertyType(model.ResourceProperty_ENUM).Pack(string(var_ce300970c22e))
	if var_ce300970c22e_err != nil {
		panic(var_ce300970c22e_err)
	}
	properties["operation"] = var_ce300970c22e_mapped

	var_0a1f5e0b181b := securityConstraint.RecordIds

	if var_0a1f5e0b181b != nil {
		var var_0a1f5e0b181b_mapped *structpb.Value

		var var_0a1f5e0b181b_l []*structpb.Value
		for _, value := range var_0a1f5e0b181b {

			var_f1150264a9ea := value
			var var_f1150264a9ea_mapped *structpb.Value

			var var_f1150264a9ea_err error
			var_f1150264a9ea_mapped, var_f1150264a9ea_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_f1150264a9ea)
			if var_f1150264a9ea_err != nil {
				panic(var_f1150264a9ea_err)
			}

			var_0a1f5e0b181b_l = append(var_0a1f5e0b181b_l, var_f1150264a9ea_mapped)
		}
		var_0a1f5e0b181b_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_0a1f5e0b181b_l})
		properties["recordIds"] = var_0a1f5e0b181b_mapped
	}

	var_a6d77caa6dd6 := securityConstraint.Before

	if var_a6d77caa6dd6 != nil {
		var var_a6d77caa6dd6_mapped *structpb.Value

		var var_a6d77caa6dd6_err error
		var_a6d77caa6dd6_mapped, var_a6d77caa6dd6_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_a6d77caa6dd6)
		if var_a6d77caa6dd6_err != nil {
			panic(var_a6d77caa6dd6_err)
		}
		properties["before"] = var_a6d77caa6dd6_mapped
	}

	var_bccab682bccc := securityConstraint.After

	if var_bccab682bccc != nil {
		var var_bccab682bccc_mapped *structpb.Value

		var var_bccab682bccc_err error
		var_bccab682bccc_mapped, var_bccab682bccc_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_bccab682bccc)
		if var_bccab682bccc_err != nil {
			panic(var_bccab682bccc_err)
		}
		properties["after"] = var_bccab682bccc_mapped
	}

	var_6165d6d477a9 := securityConstraint.User

	if var_6165d6d477a9 != nil {
		var var_6165d6d477a9_mapped *structpb.Value

		var_6165d6d477a9_mapped = structpb.NewStructValue(&structpb.Struct{Fields: UserMapperInstance.ToProperties(var_6165d6d477a9)})
		properties["user"] = var_6165d6d477a9_mapped
	}

	var_f87c75ee5e71 := securityConstraint.Role

	if var_f87c75ee5e71 != nil {
		var var_f87c75ee5e71_mapped *structpb.Value

		var_f87c75ee5e71_mapped = structpb.NewStructValue(&structpb.Struct{Fields: RoleMapperInstance.ToProperties(var_f87c75ee5e71)})
		properties["role"] = var_f87c75ee5e71_mapped
	}

	var_2904eb656e18 := securityConstraint.Permit

	var var_2904eb656e18_mapped *structpb.Value

	var var_2904eb656e18_err error
	var_2904eb656e18_mapped, var_2904eb656e18_err = types.ByResourcePropertyType(model.ResourceProperty_ENUM).Pack(string(var_2904eb656e18))
	if var_2904eb656e18_err != nil {
		panic(var_2904eb656e18_err)
	}
	properties["permit"] = var_2904eb656e18_mapped

	var_86f785a8a5bd := securityConstraint.LocalFlags

	if var_86f785a8a5bd != nil {
		var var_86f785a8a5bd_mapped *structpb.Value

		var var_86f785a8a5bd_err error
		var_86f785a8a5bd_mapped, var_86f785a8a5bd_err = types.ByResourcePropertyType(model.ResourceProperty_OBJECT).Pack(*var_86f785a8a5bd)
		if var_86f785a8a5bd_err != nil {
			panic(var_86f785a8a5bd_err)
		}
		properties["localFlags"] = var_86f785a8a5bd_mapped
	}
	return properties
}

func (m *SecurityConstraintMapper) FromProperties(properties map[string]*structpb.Value) *SecurityConstraint {
	var s = m.New()
	if properties["id"] != nil {

		var_3dfcb1ccde5f := properties["id"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_UUID).UnPack(var_3dfcb1ccde5f)

		if err != nil {
			panic(err)
		}

		var_3dfcb1ccde5f_mapped := new(uuid.UUID)
		*var_3dfcb1ccde5f_mapped = val.(uuid.UUID)

		s.Id = var_3dfcb1ccde5f_mapped
	}
	if properties["version"] != nil {

		var_eaaf75c51259 := properties["version"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(var_eaaf75c51259)

		if err != nil {
			panic(err)
		}

		var_eaaf75c51259_mapped := val.(int32)

		s.Version = var_eaaf75c51259_mapped
	}
	if properties["createdBy"] != nil {

		var_6f490e7c348b := properties["createdBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_6f490e7c348b)

		if err != nil {
			panic(err)
		}

		var_6f490e7c348b_mapped := new(string)
		*var_6f490e7c348b_mapped = val.(string)

		s.CreatedBy = var_6f490e7c348b_mapped
	}
	if properties["updatedBy"] != nil {

		var_eefce7f11cfb := properties["updatedBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_eefce7f11cfb)

		if err != nil {
			panic(err)
		}

		var_eefce7f11cfb_mapped := new(string)
		*var_eefce7f11cfb_mapped = val.(string)

		s.UpdatedBy = var_eefce7f11cfb_mapped
	}
	if properties["createdOn"] != nil {

		var_5464a464d0d8 := properties["createdOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_5464a464d0d8)

		if err != nil {
			panic(err)
		}

		var_5464a464d0d8_mapped := new(time.Time)
		*var_5464a464d0d8_mapped = val.(time.Time)

		s.CreatedOn = var_5464a464d0d8_mapped
	}
	if properties["updatedOn"] != nil {

		var_6a1558481273 := properties["updatedOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_6a1558481273)

		if err != nil {
			panic(err)
		}

		var_6a1558481273_mapped := new(time.Time)
		*var_6a1558481273_mapped = val.(time.Time)

		s.UpdatedOn = var_6a1558481273_mapped
	}
	if properties["namespace"] != nil {

		var_ba99d0bd6252 := properties["namespace"]
		var_ba99d0bd6252_mapped := NamespaceMapperInstance.FromProperties(var_ba99d0bd6252.GetStructValue().Fields)

		s.Namespace = var_ba99d0bd6252_mapped
	}
	if properties["resource"] != nil {

		var_fa4f59d3be82 := properties["resource"]
		var_fa4f59d3be82_mapped := ResourceMapperInstance.FromProperties(var_fa4f59d3be82.GetStructValue().Fields)

		s.Resource = var_fa4f59d3be82_mapped
	}
	if properties["property"] != nil {

		var_b72db4ebdb2e := properties["property"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_b72db4ebdb2e)

		if err != nil {
			panic(err)
		}

		var_b72db4ebdb2e_mapped := new(string)
		*var_b72db4ebdb2e_mapped = val.(string)

		s.Property = var_b72db4ebdb2e_mapped
	}
	if properties["propertyValue"] != nil {

		var_26cebc007ff7 := properties["propertyValue"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_26cebc007ff7)

		if err != nil {
			panic(err)
		}

		var_26cebc007ff7_mapped := new(string)
		*var_26cebc007ff7_mapped = val.(string)

		s.PropertyValue = var_26cebc007ff7_mapped
	}
	if properties["propertyMode"] != nil {

		var_2795bbcab552 := properties["propertyMode"]
		var_2795bbcab552_mapped := new(SecurityConstraintPropertyMode)
		*var_2795bbcab552_mapped = (SecurityConstraintPropertyMode)(var_2795bbcab552.GetStringValue())

		s.PropertyMode = var_2795bbcab552_mapped
	}
	if properties["operation"] != nil {

		var_601c661b56df := properties["operation"]
		var_601c661b56df_mapped := (SecurityConstraintOperation)(var_601c661b56df.GetStringValue())

		s.Operation = var_601c661b56df_mapped
	}
	if properties["recordIds"] != nil {

		var_193bcd0ced24 := properties["recordIds"]
		var_193bcd0ced24_mapped := []string{}
		for _, v := range var_193bcd0ced24.GetListValue().Values {

			var_28168c34ce7e := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_28168c34ce7e)

			if err != nil {
				panic(err)
			}

			var_28168c34ce7e_mapped := val.(string)

			var_193bcd0ced24_mapped = append(var_193bcd0ced24_mapped, var_28168c34ce7e_mapped)
		}

		s.RecordIds = var_193bcd0ced24_mapped
	}
	if properties["before"] != nil {

		var_31f0c2ad09e1 := properties["before"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_31f0c2ad09e1)

		if err != nil {
			panic(err)
		}

		var_31f0c2ad09e1_mapped := new(time.Time)
		*var_31f0c2ad09e1_mapped = val.(time.Time)

		s.Before = var_31f0c2ad09e1_mapped
	}
	if properties["after"] != nil {

		var_2c754dfd3804 := properties["after"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_2c754dfd3804)

		if err != nil {
			panic(err)
		}

		var_2c754dfd3804_mapped := new(time.Time)
		*var_2c754dfd3804_mapped = val.(time.Time)

		s.After = var_2c754dfd3804_mapped
	}
	if properties["user"] != nil {

		var_db7753f26ffa := properties["user"]
		var_db7753f26ffa_mapped := UserMapperInstance.FromProperties(var_db7753f26ffa.GetStructValue().Fields)

		s.User = var_db7753f26ffa_mapped
	}
	if properties["role"] != nil {

		var_14a438d5d63b := properties["role"]
		var_14a438d5d63b_mapped := RoleMapperInstance.FromProperties(var_14a438d5d63b.GetStructValue().Fields)

		s.Role = var_14a438d5d63b_mapped
	}
	if properties["permit"] != nil {

		var_39c34751f350 := properties["permit"]
		var_39c34751f350_mapped := (SecurityConstraintPermit)(var_39c34751f350.GetStringValue())

		s.Permit = var_39c34751f350_mapped
	}
	if properties["localFlags"] != nil {

		var_8d6def7a491d := properties["localFlags"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_OBJECT).UnPack(var_8d6def7a491d)

		if err != nil {
			panic(err)
		}

		var_8d6def7a491d_mapped := new(unstructured.Unstructured)
		*var_8d6def7a491d_mapped = val.(unstructured.Unstructured)

		s.LocalFlags = var_8d6def7a491d_mapped
	}
	return s
}
