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

	var_703a532dca81 := securityConstraint.Id

	if var_703a532dca81 != nil {
		var var_703a532dca81_mapped *structpb.Value

		var var_703a532dca81_err error
		var_703a532dca81_mapped, var_703a532dca81_err = types.ByResourcePropertyType(model.ResourceProperty_UUID).Pack(*var_703a532dca81)
		if var_703a532dca81_err != nil {
			panic(var_703a532dca81_err)
		}
		properties["id"] = var_703a532dca81_mapped
	}

	var_21f9d9b76f9f := securityConstraint.Version

	var var_21f9d9b76f9f_mapped *structpb.Value

	var var_21f9d9b76f9f_err error
	var_21f9d9b76f9f_mapped, var_21f9d9b76f9f_err = types.ByResourcePropertyType(model.ResourceProperty_INT32).Pack(var_21f9d9b76f9f)
	if var_21f9d9b76f9f_err != nil {
		panic(var_21f9d9b76f9f_err)
	}
	properties["version"] = var_21f9d9b76f9f_mapped

	var_119190eeb697 := securityConstraint.CreatedBy

	if var_119190eeb697 != nil {
		var var_119190eeb697_mapped *structpb.Value

		var var_119190eeb697_err error
		var_119190eeb697_mapped, var_119190eeb697_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_119190eeb697)
		if var_119190eeb697_err != nil {
			panic(var_119190eeb697_err)
		}
		properties["createdBy"] = var_119190eeb697_mapped
	}

	var_76cb385b46c4 := securityConstraint.UpdatedBy

	if var_76cb385b46c4 != nil {
		var var_76cb385b46c4_mapped *structpb.Value

		var var_76cb385b46c4_err error
		var_76cb385b46c4_mapped, var_76cb385b46c4_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_76cb385b46c4)
		if var_76cb385b46c4_err != nil {
			panic(var_76cb385b46c4_err)
		}
		properties["updatedBy"] = var_76cb385b46c4_mapped
	}

	var_3f5f733c5984 := securityConstraint.CreatedOn

	if var_3f5f733c5984 != nil {
		var var_3f5f733c5984_mapped *structpb.Value

		var var_3f5f733c5984_err error
		var_3f5f733c5984_mapped, var_3f5f733c5984_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_3f5f733c5984)
		if var_3f5f733c5984_err != nil {
			panic(var_3f5f733c5984_err)
		}
		properties["createdOn"] = var_3f5f733c5984_mapped
	}

	var_b4bed3dc1203 := securityConstraint.UpdatedOn

	if var_b4bed3dc1203 != nil {
		var var_b4bed3dc1203_mapped *structpb.Value

		var var_b4bed3dc1203_err error
		var_b4bed3dc1203_mapped, var_b4bed3dc1203_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_b4bed3dc1203)
		if var_b4bed3dc1203_err != nil {
			panic(var_b4bed3dc1203_err)
		}
		properties["updatedOn"] = var_b4bed3dc1203_mapped
	}

	var_458431d42bfc := securityConstraint.Namespace

	if var_458431d42bfc != nil {
		var var_458431d42bfc_mapped *structpb.Value

		var var_458431d42bfc_err error
		var_458431d42bfc_mapped, var_458431d42bfc_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_458431d42bfc)
		if var_458431d42bfc_err != nil {
			panic(var_458431d42bfc_err)
		}
		properties["namespace"] = var_458431d42bfc_mapped
	}

	var_06386251bc34 := securityConstraint.Resource

	if var_06386251bc34 != nil {
		var var_06386251bc34_mapped *structpb.Value

		var var_06386251bc34_err error
		var_06386251bc34_mapped, var_06386251bc34_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_06386251bc34)
		if var_06386251bc34_err != nil {
			panic(var_06386251bc34_err)
		}
		properties["resource"] = var_06386251bc34_mapped
	}

	var_0f4705e99885 := securityConstraint.Property

	if var_0f4705e99885 != nil {
		var var_0f4705e99885_mapped *structpb.Value

		var var_0f4705e99885_err error
		var_0f4705e99885_mapped, var_0f4705e99885_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_0f4705e99885)
		if var_0f4705e99885_err != nil {
			panic(var_0f4705e99885_err)
		}
		properties["property"] = var_0f4705e99885_mapped
	}

	var_736aa5e2ca60 := securityConstraint.PropertyValue

	if var_736aa5e2ca60 != nil {
		var var_736aa5e2ca60_mapped *structpb.Value

		var var_736aa5e2ca60_err error
		var_736aa5e2ca60_mapped, var_736aa5e2ca60_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_736aa5e2ca60)
		if var_736aa5e2ca60_err != nil {
			panic(var_736aa5e2ca60_err)
		}
		properties["propertyValue"] = var_736aa5e2ca60_mapped
	}

	var_71212d9e73ad := securityConstraint.PropertyMode

	if var_71212d9e73ad != nil {
		var var_71212d9e73ad_mapped *structpb.Value

		var var_71212d9e73ad_err error
		var_71212d9e73ad_mapped, var_71212d9e73ad_err = types.ByResourcePropertyType(model.ResourceProperty_ENUM).Pack(string(*var_71212d9e73ad))
		if var_71212d9e73ad_err != nil {
			panic(var_71212d9e73ad_err)
		}
		properties["propertyMode"] = var_71212d9e73ad_mapped
	}

	var_9d2505b591d2 := securityConstraint.Operation

	var var_9d2505b591d2_mapped *structpb.Value

	var var_9d2505b591d2_err error
	var_9d2505b591d2_mapped, var_9d2505b591d2_err = types.ByResourcePropertyType(model.ResourceProperty_ENUM).Pack(string(var_9d2505b591d2))
	if var_9d2505b591d2_err != nil {
		panic(var_9d2505b591d2_err)
	}
	properties["operation"] = var_9d2505b591d2_mapped

	var_0d65637ffca7 := securityConstraint.RecordIds

	if var_0d65637ffca7 != nil {
		var var_0d65637ffca7_mapped *structpb.Value

		var var_0d65637ffca7_l []*structpb.Value
		for _, value := range var_0d65637ffca7 {

			var_76e0430e9c4f := value
			var var_76e0430e9c4f_mapped *structpb.Value

			var var_76e0430e9c4f_err error
			var_76e0430e9c4f_mapped, var_76e0430e9c4f_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_76e0430e9c4f)
			if var_76e0430e9c4f_err != nil {
				panic(var_76e0430e9c4f_err)
			}

			var_0d65637ffca7_l = append(var_0d65637ffca7_l, var_76e0430e9c4f_mapped)
		}
		var_0d65637ffca7_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_0d65637ffca7_l})
		properties["recordIds"] = var_0d65637ffca7_mapped
	}

	var_c4161ce5fcd8 := securityConstraint.Before

	if var_c4161ce5fcd8 != nil {
		var var_c4161ce5fcd8_mapped *structpb.Value

		var var_c4161ce5fcd8_err error
		var_c4161ce5fcd8_mapped, var_c4161ce5fcd8_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_c4161ce5fcd8)
		if var_c4161ce5fcd8_err != nil {
			panic(var_c4161ce5fcd8_err)
		}
		properties["before"] = var_c4161ce5fcd8_mapped
	}

	var_2d1b19af1e36 := securityConstraint.After

	if var_2d1b19af1e36 != nil {
		var var_2d1b19af1e36_mapped *structpb.Value

		var var_2d1b19af1e36_err error
		var_2d1b19af1e36_mapped, var_2d1b19af1e36_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_2d1b19af1e36)
		if var_2d1b19af1e36_err != nil {
			panic(var_2d1b19af1e36_err)
		}
		properties["after"] = var_2d1b19af1e36_mapped
	}

	var_70aeb72bd0a2 := securityConstraint.User

	if var_70aeb72bd0a2 != nil {
		var var_70aeb72bd0a2_mapped *structpb.Value

		var_70aeb72bd0a2_mapped = structpb.NewStructValue(&structpb.Struct{Fields: UserMapperInstance.ToProperties(var_70aeb72bd0a2)})
		properties["user"] = var_70aeb72bd0a2_mapped
	}

	var_81b4615de512 := securityConstraint.Role

	if var_81b4615de512 != nil {
		var var_81b4615de512_mapped *structpb.Value

		var_81b4615de512_mapped = structpb.NewStructValue(&structpb.Struct{Fields: RoleMapperInstance.ToProperties(var_81b4615de512)})
		properties["role"] = var_81b4615de512_mapped
	}

	var_eb1f192694bc := securityConstraint.Permit

	var var_eb1f192694bc_mapped *structpb.Value

	var var_eb1f192694bc_err error
	var_eb1f192694bc_mapped, var_eb1f192694bc_err = types.ByResourcePropertyType(model.ResourceProperty_ENUM).Pack(string(var_eb1f192694bc))
	if var_eb1f192694bc_err != nil {
		panic(var_eb1f192694bc_err)
	}
	properties["permit"] = var_eb1f192694bc_mapped

	var_11d66ab614da := securityConstraint.LocalFlags

	if var_11d66ab614da != nil {
		var var_11d66ab614da_mapped *structpb.Value

		var var_11d66ab614da_err error
		var_11d66ab614da_mapped, var_11d66ab614da_err = types.ByResourcePropertyType(model.ResourceProperty_OBJECT).Pack(*var_11d66ab614da)
		if var_11d66ab614da_err != nil {
			panic(var_11d66ab614da_err)
		}
		properties["localFlags"] = var_11d66ab614da_mapped
	}
	return properties
}

func (m *SecurityConstraintMapper) FromProperties(properties map[string]*structpb.Value) *SecurityConstraint {
	var s = m.New()
	if properties["id"] != nil {

		var_aa306614a7bc := properties["id"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_UUID).UnPack(var_aa306614a7bc)

		if err != nil {
			panic(err)
		}

		var_aa306614a7bc_mapped := new(uuid.UUID)
		*var_aa306614a7bc_mapped = val.(uuid.UUID)

		s.Id = var_aa306614a7bc_mapped
	}
	if properties["version"] != nil {

		var_cbf7dbace405 := properties["version"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(var_cbf7dbace405)

		if err != nil {
			panic(err)
		}

		var_cbf7dbace405_mapped := val.(int32)

		s.Version = var_cbf7dbace405_mapped
	}
	if properties["createdBy"] != nil {

		var_003ec08a69c7 := properties["createdBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_003ec08a69c7)

		if err != nil {
			panic(err)
		}

		var_003ec08a69c7_mapped := new(string)
		*var_003ec08a69c7_mapped = val.(string)

		s.CreatedBy = var_003ec08a69c7_mapped
	}
	if properties["updatedBy"] != nil {

		var_bd69bb6ab3bf := properties["updatedBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_bd69bb6ab3bf)

		if err != nil {
			panic(err)
		}

		var_bd69bb6ab3bf_mapped := new(string)
		*var_bd69bb6ab3bf_mapped = val.(string)

		s.UpdatedBy = var_bd69bb6ab3bf_mapped
	}
	if properties["createdOn"] != nil {

		var_f73f506883fd := properties["createdOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_f73f506883fd)

		if err != nil {
			panic(err)
		}

		var_f73f506883fd_mapped := new(time.Time)
		*var_f73f506883fd_mapped = val.(time.Time)

		s.CreatedOn = var_f73f506883fd_mapped
	}
	if properties["updatedOn"] != nil {

		var_eb11c0da4aa2 := properties["updatedOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_eb11c0da4aa2)

		if err != nil {
			panic(err)
		}

		var_eb11c0da4aa2_mapped := new(time.Time)
		*var_eb11c0da4aa2_mapped = val.(time.Time)

		s.UpdatedOn = var_eb11c0da4aa2_mapped
	}
	if properties["namespace"] != nil {

		var_b15065882623 := properties["namespace"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_b15065882623)

		if err != nil {
			panic(err)
		}

		var_b15065882623_mapped := new(string)
		*var_b15065882623_mapped = val.(string)

		s.Namespace = var_b15065882623_mapped
	}
	if properties["resource"] != nil {

		var_8a6b916c4ed5 := properties["resource"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_8a6b916c4ed5)

		if err != nil {
			panic(err)
		}

		var_8a6b916c4ed5_mapped := new(string)
		*var_8a6b916c4ed5_mapped = val.(string)

		s.Resource = var_8a6b916c4ed5_mapped
	}
	if properties["property"] != nil {

		var_1b61403d7331 := properties["property"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_1b61403d7331)

		if err != nil {
			panic(err)
		}

		var_1b61403d7331_mapped := new(string)
		*var_1b61403d7331_mapped = val.(string)

		s.Property = var_1b61403d7331_mapped
	}
	if properties["propertyValue"] != nil {

		var_4c6b05b9e4a3 := properties["propertyValue"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_4c6b05b9e4a3)

		if err != nil {
			panic(err)
		}

		var_4c6b05b9e4a3_mapped := new(string)
		*var_4c6b05b9e4a3_mapped = val.(string)

		s.PropertyValue = var_4c6b05b9e4a3_mapped
	}
	if properties["propertyMode"] != nil {

		var_d31cbc2f3ba4 := properties["propertyMode"]
		var_d31cbc2f3ba4_mapped := new(SecurityConstraintPropertyMode)
		*var_d31cbc2f3ba4_mapped = (SecurityConstraintPropertyMode)(var_d31cbc2f3ba4.GetStringValue())

		s.PropertyMode = var_d31cbc2f3ba4_mapped
	}
	if properties["operation"] != nil {

		var_a3565e21f92d := properties["operation"]
		var_a3565e21f92d_mapped := (SecurityConstraintOperation)(var_a3565e21f92d.GetStringValue())

		s.Operation = var_a3565e21f92d_mapped
	}
	if properties["recordIds"] != nil {

		var_f1c0db4b14a0 := properties["recordIds"]
		var_f1c0db4b14a0_mapped := []string{}
		for _, v := range var_f1c0db4b14a0.GetListValue().Values {

			var_4adacbf7a819 := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_4adacbf7a819)

			if err != nil {
				panic(err)
			}

			var_4adacbf7a819_mapped := val.(string)

			var_f1c0db4b14a0_mapped = append(var_f1c0db4b14a0_mapped, var_4adacbf7a819_mapped)
		}

		s.RecordIds = var_f1c0db4b14a0_mapped
	}
	if properties["before"] != nil {

		var_ccec02898583 := properties["before"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_ccec02898583)

		if err != nil {
			panic(err)
		}

		var_ccec02898583_mapped := new(time.Time)
		*var_ccec02898583_mapped = val.(time.Time)

		s.Before = var_ccec02898583_mapped
	}
	if properties["after"] != nil {

		var_dd96e57e3840 := properties["after"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_dd96e57e3840)

		if err != nil {
			panic(err)
		}

		var_dd96e57e3840_mapped := new(time.Time)
		*var_dd96e57e3840_mapped = val.(time.Time)

		s.After = var_dd96e57e3840_mapped
	}
	if properties["user"] != nil {

		var_30e07ef6ad39 := properties["user"]
		var_30e07ef6ad39_mapped := UserMapperInstance.FromProperties(var_30e07ef6ad39.GetStructValue().Fields)

		s.User = var_30e07ef6ad39_mapped
	}
	if properties["role"] != nil {

		var_bd795fe24412 := properties["role"]
		var_bd795fe24412_mapped := RoleMapperInstance.FromProperties(var_bd795fe24412.GetStructValue().Fields)

		s.Role = var_bd795fe24412_mapped
	}
	if properties["permit"] != nil {

		var_8bac683b4b74 := properties["permit"]
		var_8bac683b4b74_mapped := (SecurityConstraintPermit)(var_8bac683b4b74.GetStringValue())

		s.Permit = var_8bac683b4b74_mapped
	}
	if properties["localFlags"] != nil {

		var_7c3924e21557 := properties["localFlags"]
		var_7c3924e21557_mapped := new(unstructured.Unstructured)
		*var_7c3924e21557_mapped = unstructured.FromStructValue(var_7c3924e21557.GetStructValue())

		s.LocalFlags = var_7c3924e21557_mapped
	}
	return s
}
