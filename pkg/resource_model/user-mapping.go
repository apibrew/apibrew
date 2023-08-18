package resource_model

import (
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/types"
	"google.golang.org/protobuf/types/known/structpb"
)

import "github.com/google/uuid"
import "time"
import "github.com/apibrew/apibrew/pkg/formats/unstructured"

type UserMapper struct {
}

func NewUserMapper() *UserMapper {
	return &UserMapper{}
}

var UserMapperInstance = NewUserMapper()

func (m *UserMapper) New() *User {
	return &User{}
}

func (m *UserMapper) ToRecord(user *User) *model.Record {
	var rec = &model.Record{}
	rec.Properties = m.ToProperties(user)

	if user.Id != nil {
		rec.Id = user.Id.String()
	}

	return rec
}

func (m *UserMapper) FromRecord(record *model.Record) *User {
	return m.FromProperties(record.Properties)
}

func (m *UserMapper) ToProperties(user *User) map[string]*structpb.Value {
	var properties = make(map[string]*structpb.Value)

	var_7226007566d4 := user.Id

	if var_7226007566d4 != nil {
		var var_7226007566d4_mapped *structpb.Value

		var var_7226007566d4_err error
		var_7226007566d4_mapped, var_7226007566d4_err = types.ByResourcePropertyType(model.ResourceProperty_UUID).Pack(*var_7226007566d4)
		if var_7226007566d4_err != nil {
			panic(var_7226007566d4_err)
		}
		properties["id"] = var_7226007566d4_mapped
	}

	var_03050259da65 := user.Version

	var var_03050259da65_mapped *structpb.Value

	var var_03050259da65_err error
	var_03050259da65_mapped, var_03050259da65_err = types.ByResourcePropertyType(model.ResourceProperty_INT32).Pack(var_03050259da65)
	if var_03050259da65_err != nil {
		panic(var_03050259da65_err)
	}
	properties["version"] = var_03050259da65_mapped

	var_037fd0ce4123 := user.CreatedBy

	if var_037fd0ce4123 != nil {
		var var_037fd0ce4123_mapped *structpb.Value

		var var_037fd0ce4123_err error
		var_037fd0ce4123_mapped, var_037fd0ce4123_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_037fd0ce4123)
		if var_037fd0ce4123_err != nil {
			panic(var_037fd0ce4123_err)
		}
		properties["createdBy"] = var_037fd0ce4123_mapped
	}

	var_358cae88a1b6 := user.UpdatedBy

	if var_358cae88a1b6 != nil {
		var var_358cae88a1b6_mapped *structpb.Value

		var var_358cae88a1b6_err error
		var_358cae88a1b6_mapped, var_358cae88a1b6_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_358cae88a1b6)
		if var_358cae88a1b6_err != nil {
			panic(var_358cae88a1b6_err)
		}
		properties["updatedBy"] = var_358cae88a1b6_mapped
	}

	var_0143d3a57a63 := user.CreatedOn

	if var_0143d3a57a63 != nil {
		var var_0143d3a57a63_mapped *structpb.Value

		var var_0143d3a57a63_err error
		var_0143d3a57a63_mapped, var_0143d3a57a63_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_0143d3a57a63)
		if var_0143d3a57a63_err != nil {
			panic(var_0143d3a57a63_err)
		}
		properties["createdOn"] = var_0143d3a57a63_mapped
	}

	var_6b898d721d6f := user.UpdatedOn

	if var_6b898d721d6f != nil {
		var var_6b898d721d6f_mapped *structpb.Value

		var var_6b898d721d6f_err error
		var_6b898d721d6f_mapped, var_6b898d721d6f_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_6b898d721d6f)
		if var_6b898d721d6f_err != nil {
			panic(var_6b898d721d6f_err)
		}
		properties["updatedOn"] = var_6b898d721d6f_mapped
	}

	var_e5ea7a372bf2 := user.Username

	var var_e5ea7a372bf2_mapped *structpb.Value

	var var_e5ea7a372bf2_err error
	var_e5ea7a372bf2_mapped, var_e5ea7a372bf2_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_e5ea7a372bf2)
	if var_e5ea7a372bf2_err != nil {
		panic(var_e5ea7a372bf2_err)
	}
	properties["username"] = var_e5ea7a372bf2_mapped

	var_8a7725301ed2 := user.Password

	if var_8a7725301ed2 != nil {
		var var_8a7725301ed2_mapped *structpb.Value

		var var_8a7725301ed2_err error
		var_8a7725301ed2_mapped, var_8a7725301ed2_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_8a7725301ed2)
		if var_8a7725301ed2_err != nil {
			panic(var_8a7725301ed2_err)
		}
		properties["password"] = var_8a7725301ed2_mapped
	}

	var_05355f4a9265 := user.Roles

	if var_05355f4a9265 != nil {
		var var_05355f4a9265_mapped *structpb.Value

		var var_05355f4a9265_l []*structpb.Value
		for _, value := range var_05355f4a9265 {

			var_121644eca09a := value
			var var_121644eca09a_mapped *structpb.Value

			var_121644eca09a_mapped = structpb.NewStructValue(&structpb.Struct{Fields: RoleMapperInstance.ToProperties(var_121644eca09a)})

			var_05355f4a9265_l = append(var_05355f4a9265_l, var_121644eca09a_mapped)
		}
		var_05355f4a9265_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_05355f4a9265_l})
		properties["roles"] = var_05355f4a9265_mapped
	}

	var_1a0012f2cf07 := user.SecurityConstraints

	if var_1a0012f2cf07 != nil {
		var var_1a0012f2cf07_mapped *structpb.Value

		var var_1a0012f2cf07_l []*structpb.Value
		for _, value := range var_1a0012f2cf07 {

			var_5ecce5a79500 := value
			var var_5ecce5a79500_mapped *structpb.Value

			var_5ecce5a79500_mapped = structpb.NewStructValue(&structpb.Struct{Fields: SecurityConstraintMapperInstance.ToProperties(var_5ecce5a79500)})

			var_1a0012f2cf07_l = append(var_1a0012f2cf07_l, var_5ecce5a79500_mapped)
		}
		var_1a0012f2cf07_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_1a0012f2cf07_l})
		properties["securityConstraints"] = var_1a0012f2cf07_mapped
	}

	var_cbc50f6b0a8a := user.Details

	if var_cbc50f6b0a8a != nil {
		var var_cbc50f6b0a8a_mapped *structpb.Value

		var var_cbc50f6b0a8a_err error
		var_cbc50f6b0a8a_mapped, var_cbc50f6b0a8a_err = types.ByResourcePropertyType(model.ResourceProperty_OBJECT).Pack(*var_cbc50f6b0a8a)
		if var_cbc50f6b0a8a_err != nil {
			panic(var_cbc50f6b0a8a_err)
		}
		properties["details"] = var_cbc50f6b0a8a_mapped
	}
	return properties
}

func (m *UserMapper) FromProperties(properties map[string]*structpb.Value) *User {
	var s = m.New()
	if properties["id"] != nil {

		var_ca60e3ef8d26 := properties["id"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_UUID).UnPack(var_ca60e3ef8d26)

		if err != nil {
			panic(err)
		}

		var_ca60e3ef8d26_mapped := new(uuid.UUID)
		*var_ca60e3ef8d26_mapped = val.(uuid.UUID)

		s.Id = var_ca60e3ef8d26_mapped
	}
	if properties["version"] != nil {

		var_c63831c5a4b6 := properties["version"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(var_c63831c5a4b6)

		if err != nil {
			panic(err)
		}

		var_c63831c5a4b6_mapped := val.(int32)

		s.Version = var_c63831c5a4b6_mapped
	}
	if properties["createdBy"] != nil {

		var_611452260235 := properties["createdBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_611452260235)

		if err != nil {
			panic(err)
		}

		var_611452260235_mapped := new(string)
		*var_611452260235_mapped = val.(string)

		s.CreatedBy = var_611452260235_mapped
	}
	if properties["updatedBy"] != nil {

		var_79236cbe3719 := properties["updatedBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_79236cbe3719)

		if err != nil {
			panic(err)
		}

		var_79236cbe3719_mapped := new(string)
		*var_79236cbe3719_mapped = val.(string)

		s.UpdatedBy = var_79236cbe3719_mapped
	}
	if properties["createdOn"] != nil {

		var_27e2bda839c7 := properties["createdOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_27e2bda839c7)

		if err != nil {
			panic(err)
		}

		var_27e2bda839c7_mapped := new(time.Time)
		*var_27e2bda839c7_mapped = val.(time.Time)

		s.CreatedOn = var_27e2bda839c7_mapped
	}
	if properties["updatedOn"] != nil {

		var_0da8dabf19ba := properties["updatedOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_0da8dabf19ba)

		if err != nil {
			panic(err)
		}

		var_0da8dabf19ba_mapped := new(time.Time)
		*var_0da8dabf19ba_mapped = val.(time.Time)

		s.UpdatedOn = var_0da8dabf19ba_mapped
	}
	if properties["username"] != nil {

		var_a25c2bc5a4bf := properties["username"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_a25c2bc5a4bf)

		if err != nil {
			panic(err)
		}

		var_a25c2bc5a4bf_mapped := val.(string)

		s.Username = var_a25c2bc5a4bf_mapped
	}
	if properties["password"] != nil {

		var_0aee1ca5b65f := properties["password"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_0aee1ca5b65f)

		if err != nil {
			panic(err)
		}

		var_0aee1ca5b65f_mapped := new(string)
		*var_0aee1ca5b65f_mapped = val.(string)

		s.Password = var_0aee1ca5b65f_mapped
	}
	if properties["roles"] != nil {

		var_2e1f58c1264b := properties["roles"]
		var_2e1f58c1264b_mapped := []*Role{}
		for _, v := range var_2e1f58c1264b.GetListValue().Values {

			var_db343fc1f006 := v
			var_db343fc1f006_mapped := RoleMapperInstance.FromProperties(var_db343fc1f006.GetStructValue().Fields)

			var_2e1f58c1264b_mapped = append(var_2e1f58c1264b_mapped, var_db343fc1f006_mapped)
		}

		s.Roles = var_2e1f58c1264b_mapped
	}
	if properties["securityConstraints"] != nil {

		var_a2e4fd9a722a := properties["securityConstraints"]
		var_a2e4fd9a722a_mapped := []*SecurityConstraint{}
		for _, v := range var_a2e4fd9a722a.GetListValue().Values {

			var_647d91ad9151 := v
			var_647d91ad9151_mapped := SecurityConstraintMapperInstance.FromProperties(var_647d91ad9151.GetStructValue().Fields)

			var_a2e4fd9a722a_mapped = append(var_a2e4fd9a722a_mapped, var_647d91ad9151_mapped)
		}

		s.SecurityConstraints = var_a2e4fd9a722a_mapped
	}
	if properties["details"] != nil {

		var_d2fa02bf78a5 := properties["details"]
		var_d2fa02bf78a5_mapped := new(unstructured.Unstructured)
		*var_d2fa02bf78a5_mapped = unstructured.FromStructValue(var_d2fa02bf78a5.GetStructValue())

		s.Details = var_d2fa02bf78a5_mapped
	}
	return s
}
