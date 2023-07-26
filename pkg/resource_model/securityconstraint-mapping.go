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

	var_62d2b64df1c9 := securityConstraint.Id

	if var_62d2b64df1c9 != nil {
		var var_62d2b64df1c9_mapped *structpb.Value

		var var_62d2b64df1c9_err error
		var_62d2b64df1c9_mapped, var_62d2b64df1c9_err = types.ByResourcePropertyType(model.ResourceProperty_UUID).Pack(*var_62d2b64df1c9)
		if var_62d2b64df1c9_err != nil {
			panic(var_62d2b64df1c9_err)
		}
		properties["id"] = var_62d2b64df1c9_mapped
	}

	var_0ec648c454e6 := securityConstraint.Version

	var var_0ec648c454e6_mapped *structpb.Value

	var var_0ec648c454e6_err error
	var_0ec648c454e6_mapped, var_0ec648c454e6_err = types.ByResourcePropertyType(model.ResourceProperty_INT32).Pack(var_0ec648c454e6)
	if var_0ec648c454e6_err != nil {
		panic(var_0ec648c454e6_err)
	}
	properties["version"] = var_0ec648c454e6_mapped

	var_19d0e145a549 := securityConstraint.CreatedBy

	if var_19d0e145a549 != nil {
		var var_19d0e145a549_mapped *structpb.Value

		var var_19d0e145a549_err error
		var_19d0e145a549_mapped, var_19d0e145a549_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_19d0e145a549)
		if var_19d0e145a549_err != nil {
			panic(var_19d0e145a549_err)
		}
		properties["createdBy"] = var_19d0e145a549_mapped
	}

	var_72e6b9194e38 := securityConstraint.UpdatedBy

	if var_72e6b9194e38 != nil {
		var var_72e6b9194e38_mapped *structpb.Value

		var var_72e6b9194e38_err error
		var_72e6b9194e38_mapped, var_72e6b9194e38_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_72e6b9194e38)
		if var_72e6b9194e38_err != nil {
			panic(var_72e6b9194e38_err)
		}
		properties["updatedBy"] = var_72e6b9194e38_mapped
	}

	var_5a6c3f49b242 := securityConstraint.CreatedOn

	if var_5a6c3f49b242 != nil {
		var var_5a6c3f49b242_mapped *structpb.Value

		var var_5a6c3f49b242_err error
		var_5a6c3f49b242_mapped, var_5a6c3f49b242_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_5a6c3f49b242)
		if var_5a6c3f49b242_err != nil {
			panic(var_5a6c3f49b242_err)
		}
		properties["createdOn"] = var_5a6c3f49b242_mapped
	}

	var_d833a9b6f02e := securityConstraint.UpdatedOn

	if var_d833a9b6f02e != nil {
		var var_d833a9b6f02e_mapped *structpb.Value

		var var_d833a9b6f02e_err error
		var_d833a9b6f02e_mapped, var_d833a9b6f02e_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_d833a9b6f02e)
		if var_d833a9b6f02e_err != nil {
			panic(var_d833a9b6f02e_err)
		}
		properties["updatedOn"] = var_d833a9b6f02e_mapped
	}

	var_684d4ddb882c := securityConstraint.Namespace

	if var_684d4ddb882c != nil {
		var var_684d4ddb882c_mapped *structpb.Value

		var_684d4ddb882c_mapped = structpb.NewStructValue(&structpb.Struct{Fields: NamespaceMapperInstance.ToProperties(var_684d4ddb882c)})
		properties["namespace"] = var_684d4ddb882c_mapped
	}

	var_b271f0b3b799 := securityConstraint.Resource

	if var_b271f0b3b799 != nil {
		var var_b271f0b3b799_mapped *structpb.Value

		var_b271f0b3b799_mapped = structpb.NewStructValue(&structpb.Struct{Fields: ResourceMapperInstance.ToProperties(var_b271f0b3b799)})
		properties["resource"] = var_b271f0b3b799_mapped
	}

	var_572dab6a1568 := securityConstraint.Property

	if var_572dab6a1568 != nil {
		var var_572dab6a1568_mapped *structpb.Value

		var var_572dab6a1568_err error
		var_572dab6a1568_mapped, var_572dab6a1568_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_572dab6a1568)
		if var_572dab6a1568_err != nil {
			panic(var_572dab6a1568_err)
		}
		properties["property"] = var_572dab6a1568_mapped
	}

	var_5865b313c395 := securityConstraint.PropertyValue

	if var_5865b313c395 != nil {
		var var_5865b313c395_mapped *structpb.Value

		var var_5865b313c395_err error
		var_5865b313c395_mapped, var_5865b313c395_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_5865b313c395)
		if var_5865b313c395_err != nil {
			panic(var_5865b313c395_err)
		}
		properties["propertyValue"] = var_5865b313c395_mapped
	}

	var_4f121e77af3d := securityConstraint.PropertyMode

	if var_4f121e77af3d != nil {
		var var_4f121e77af3d_mapped *structpb.Value

		var var_4f121e77af3d_err error
		var_4f121e77af3d_mapped, var_4f121e77af3d_err = types.ByResourcePropertyType(model.ResourceProperty_ENUM).Pack(string(*var_4f121e77af3d))
		if var_4f121e77af3d_err != nil {
			panic(var_4f121e77af3d_err)
		}
		properties["propertyMode"] = var_4f121e77af3d_mapped
	}

	var_dac54bab8c14 := securityConstraint.Operation

	var var_dac54bab8c14_mapped *structpb.Value

	var var_dac54bab8c14_err error
	var_dac54bab8c14_mapped, var_dac54bab8c14_err = types.ByResourcePropertyType(model.ResourceProperty_ENUM).Pack(string(var_dac54bab8c14))
	if var_dac54bab8c14_err != nil {
		panic(var_dac54bab8c14_err)
	}
	properties["operation"] = var_dac54bab8c14_mapped

	var_3159dc6d171c := securityConstraint.RecordIds

	if var_3159dc6d171c != nil {
		var var_3159dc6d171c_mapped *structpb.Value

		var var_3159dc6d171c_l []*structpb.Value
		for _, value := range var_3159dc6d171c {

			var_3869e2674fcb := value
			var var_3869e2674fcb_mapped *structpb.Value

			var var_3869e2674fcb_err error
			var_3869e2674fcb_mapped, var_3869e2674fcb_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_3869e2674fcb)
			if var_3869e2674fcb_err != nil {
				panic(var_3869e2674fcb_err)
			}

			var_3159dc6d171c_l = append(var_3159dc6d171c_l, var_3869e2674fcb_mapped)
		}
		var_3159dc6d171c_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_3159dc6d171c_l})
		properties["recordIds"] = var_3159dc6d171c_mapped
	}

	var_0f6b1e0dd23b := securityConstraint.Before

	if var_0f6b1e0dd23b != nil {
		var var_0f6b1e0dd23b_mapped *structpb.Value

		var var_0f6b1e0dd23b_err error
		var_0f6b1e0dd23b_mapped, var_0f6b1e0dd23b_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_0f6b1e0dd23b)
		if var_0f6b1e0dd23b_err != nil {
			panic(var_0f6b1e0dd23b_err)
		}
		properties["before"] = var_0f6b1e0dd23b_mapped
	}

	var_5acf0248c920 := securityConstraint.After

	if var_5acf0248c920 != nil {
		var var_5acf0248c920_mapped *structpb.Value

		var var_5acf0248c920_err error
		var_5acf0248c920_mapped, var_5acf0248c920_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_5acf0248c920)
		if var_5acf0248c920_err != nil {
			panic(var_5acf0248c920_err)
		}
		properties["after"] = var_5acf0248c920_mapped
	}

	var_9490a9b71120 := securityConstraint.User

	if var_9490a9b71120 != nil {
		var var_9490a9b71120_mapped *structpb.Value

		var_9490a9b71120_mapped = structpb.NewStructValue(&structpb.Struct{Fields: UserMapperInstance.ToProperties(var_9490a9b71120)})
		properties["user"] = var_9490a9b71120_mapped
	}

	var_7045e5f81170 := securityConstraint.Role

	if var_7045e5f81170 != nil {
		var var_7045e5f81170_mapped *structpb.Value

		var_7045e5f81170_mapped = structpb.NewStructValue(&structpb.Struct{Fields: RoleMapperInstance.ToProperties(var_7045e5f81170)})
		properties["role"] = var_7045e5f81170_mapped
	}

	var_0c89095b3c66 := securityConstraint.Permit

	var var_0c89095b3c66_mapped *structpb.Value

	var var_0c89095b3c66_err error
	var_0c89095b3c66_mapped, var_0c89095b3c66_err = types.ByResourcePropertyType(model.ResourceProperty_ENUM).Pack(string(var_0c89095b3c66))
	if var_0c89095b3c66_err != nil {
		panic(var_0c89095b3c66_err)
	}
	properties["permit"] = var_0c89095b3c66_mapped

	var_2bcc26b5dd5a := securityConstraint.LocalFlags

	if var_2bcc26b5dd5a != nil {
		var var_2bcc26b5dd5a_mapped *structpb.Value

		var var_2bcc26b5dd5a_err error
		var_2bcc26b5dd5a_mapped, var_2bcc26b5dd5a_err = types.ByResourcePropertyType(model.ResourceProperty_OBJECT).Pack(*var_2bcc26b5dd5a)
		if var_2bcc26b5dd5a_err != nil {
			panic(var_2bcc26b5dd5a_err)
		}
		properties["localFlags"] = var_2bcc26b5dd5a_mapped
	}
	return properties
}

func (m *SecurityConstraintMapper) FromProperties(properties map[string]*structpb.Value) *SecurityConstraint {
	var s = m.New()
	if properties["id"] != nil {

		var_073198168f6e := properties["id"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_UUID).UnPack(var_073198168f6e)

		if err != nil {
			panic(err)
		}

		var_073198168f6e_mapped := new(uuid.UUID)
		*var_073198168f6e_mapped = val.(uuid.UUID)

		s.Id = var_073198168f6e_mapped
	}
	if properties["version"] != nil {

		var_50441db2bc7c := properties["version"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(var_50441db2bc7c)

		if err != nil {
			panic(err)
		}

		var_50441db2bc7c_mapped := val.(int32)

		s.Version = var_50441db2bc7c_mapped
	}
	if properties["createdBy"] != nil {

		var_1610917ec28d := properties["createdBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_1610917ec28d)

		if err != nil {
			panic(err)
		}

		var_1610917ec28d_mapped := new(string)
		*var_1610917ec28d_mapped = val.(string)

		s.CreatedBy = var_1610917ec28d_mapped
	}
	if properties["updatedBy"] != nil {

		var_e3fe5d62fd82 := properties["updatedBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_e3fe5d62fd82)

		if err != nil {
			panic(err)
		}

		var_e3fe5d62fd82_mapped := new(string)
		*var_e3fe5d62fd82_mapped = val.(string)

		s.UpdatedBy = var_e3fe5d62fd82_mapped
	}
	if properties["createdOn"] != nil {

		var_a70da324225a := properties["createdOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_a70da324225a)

		if err != nil {
			panic(err)
		}

		var_a70da324225a_mapped := new(time.Time)
		*var_a70da324225a_mapped = val.(time.Time)

		s.CreatedOn = var_a70da324225a_mapped
	}
	if properties["updatedOn"] != nil {

		var_2ddd6704ac93 := properties["updatedOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_2ddd6704ac93)

		if err != nil {
			panic(err)
		}

		var_2ddd6704ac93_mapped := new(time.Time)
		*var_2ddd6704ac93_mapped = val.(time.Time)

		s.UpdatedOn = var_2ddd6704ac93_mapped
	}
	if properties["namespace"] != nil {

		var_bc9ae97ee776 := properties["namespace"]
		var_bc9ae97ee776_mapped := NamespaceMapperInstance.FromProperties(var_bc9ae97ee776.GetStructValue().Fields)

		s.Namespace = var_bc9ae97ee776_mapped
	}
	if properties["resource"] != nil {

		var_bcf024d57392 := properties["resource"]
		var_bcf024d57392_mapped := ResourceMapperInstance.FromProperties(var_bcf024d57392.GetStructValue().Fields)

		s.Resource = var_bcf024d57392_mapped
	}
	if properties["property"] != nil {

		var_c92058dde9d1 := properties["property"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_c92058dde9d1)

		if err != nil {
			panic(err)
		}

		var_c92058dde9d1_mapped := new(string)
		*var_c92058dde9d1_mapped = val.(string)

		s.Property = var_c92058dde9d1_mapped
	}
	if properties["propertyValue"] != nil {

		var_32f5a0658d23 := properties["propertyValue"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_32f5a0658d23)

		if err != nil {
			panic(err)
		}

		var_32f5a0658d23_mapped := new(string)
		*var_32f5a0658d23_mapped = val.(string)

		s.PropertyValue = var_32f5a0658d23_mapped
	}
	if properties["propertyMode"] != nil {

		var_9a44894b74b9 := properties["propertyMode"]
		var_9a44894b74b9_mapped := new(SecurityConstraintPropertyMode)
		*var_9a44894b74b9_mapped = (SecurityConstraintPropertyMode)(var_9a44894b74b9.GetStringValue())

		s.PropertyMode = var_9a44894b74b9_mapped
	}
	if properties["operation"] != nil {

		var_c3c2328df122 := properties["operation"]
		var_c3c2328df122_mapped := (SecurityConstraintOperation)(var_c3c2328df122.GetStringValue())

		s.Operation = var_c3c2328df122_mapped
	}
	if properties["recordIds"] != nil {

		var_2fa0c9f0faed := properties["recordIds"]
		var_2fa0c9f0faed_mapped := []string{}
		for _, v := range var_2fa0c9f0faed.GetListValue().Values {

			var_04b73cb1277c := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_04b73cb1277c)

			if err != nil {
				panic(err)
			}

			var_04b73cb1277c_mapped := val.(string)

			var_2fa0c9f0faed_mapped = append(var_2fa0c9f0faed_mapped, var_04b73cb1277c_mapped)
		}

		s.RecordIds = var_2fa0c9f0faed_mapped
	}
	if properties["before"] != nil {

		var_fabf68bd1af7 := properties["before"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_fabf68bd1af7)

		if err != nil {
			panic(err)
		}

		var_fabf68bd1af7_mapped := new(time.Time)
		*var_fabf68bd1af7_mapped = val.(time.Time)

		s.Before = var_fabf68bd1af7_mapped
	}
	if properties["after"] != nil {

		var_954dfd67ab6e := properties["after"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_954dfd67ab6e)

		if err != nil {
			panic(err)
		}

		var_954dfd67ab6e_mapped := new(time.Time)
		*var_954dfd67ab6e_mapped = val.(time.Time)

		s.After = var_954dfd67ab6e_mapped
	}
	if properties["user"] != nil {

		var_5b76893c7d50 := properties["user"]
		var_5b76893c7d50_mapped := UserMapperInstance.FromProperties(var_5b76893c7d50.GetStructValue().Fields)

		s.User = var_5b76893c7d50_mapped
	}
	if properties["role"] != nil {

		var_40ea28b5e142 := properties["role"]
		var_40ea28b5e142_mapped := RoleMapperInstance.FromProperties(var_40ea28b5e142.GetStructValue().Fields)

		s.Role = var_40ea28b5e142_mapped
	}
	if properties["permit"] != nil {

		var_4ace06b332ff := properties["permit"]
		var_4ace06b332ff_mapped := (SecurityConstraintPermit)(var_4ace06b332ff.GetStringValue())

		s.Permit = var_4ace06b332ff_mapped
	}
	if properties["localFlags"] != nil {

		var_917cb5d68882 := properties["localFlags"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_OBJECT).UnPack(var_917cb5d68882)

		if err != nil {
			panic(err)
		}

		var_917cb5d68882_mapped := new(unstructured.Unstructured)
		*var_917cb5d68882_mapped = val.(unstructured.Unstructured)

		s.LocalFlags = var_917cb5d68882_mapped
	}
	return s
}
