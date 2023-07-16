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

	var_66d125a5b576 := user.Id

	if var_66d125a5b576 != nil {
		var var_66d125a5b576_mapped *structpb.Value

		var var_66d125a5b576_err error
		var_66d125a5b576_mapped, var_66d125a5b576_err = types.ByResourcePropertyType(model.ResourceProperty_UUID).Pack(*var_66d125a5b576)
		if var_66d125a5b576_err != nil {
			panic(var_66d125a5b576_err)
		}
		properties["id"] = var_66d125a5b576_mapped
	}

	var_3f32f14e3960 := user.Version

	var var_3f32f14e3960_mapped *structpb.Value

	var var_3f32f14e3960_err error
	var_3f32f14e3960_mapped, var_3f32f14e3960_err = types.ByResourcePropertyType(model.ResourceProperty_INT32).Pack(var_3f32f14e3960)
	if var_3f32f14e3960_err != nil {
		panic(var_3f32f14e3960_err)
	}
	properties["version"] = var_3f32f14e3960_mapped

	var_9e8c047417c1 := user.CreatedBy

	if var_9e8c047417c1 != nil {
		var var_9e8c047417c1_mapped *structpb.Value

		var var_9e8c047417c1_err error
		var_9e8c047417c1_mapped, var_9e8c047417c1_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_9e8c047417c1)
		if var_9e8c047417c1_err != nil {
			panic(var_9e8c047417c1_err)
		}
		properties["createdBy"] = var_9e8c047417c1_mapped
	}

	var_823dcf757bb6 := user.UpdatedBy

	if var_823dcf757bb6 != nil {
		var var_823dcf757bb6_mapped *structpb.Value

		var var_823dcf757bb6_err error
		var_823dcf757bb6_mapped, var_823dcf757bb6_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_823dcf757bb6)
		if var_823dcf757bb6_err != nil {
			panic(var_823dcf757bb6_err)
		}
		properties["updatedBy"] = var_823dcf757bb6_mapped
	}

	var_d340c552a61a := user.CreatedOn

	if var_d340c552a61a != nil {
		var var_d340c552a61a_mapped *structpb.Value

		var var_d340c552a61a_err error
		var_d340c552a61a_mapped, var_d340c552a61a_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_d340c552a61a)
		if var_d340c552a61a_err != nil {
			panic(var_d340c552a61a_err)
		}
		properties["createdOn"] = var_d340c552a61a_mapped
	}

	var_3c89846de24f := user.UpdatedOn

	if var_3c89846de24f != nil {
		var var_3c89846de24f_mapped *structpb.Value

		var var_3c89846de24f_err error
		var_3c89846de24f_mapped, var_3c89846de24f_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_3c89846de24f)
		if var_3c89846de24f_err != nil {
			panic(var_3c89846de24f_err)
		}
		properties["updatedOn"] = var_3c89846de24f_mapped
	}

	var_1d38b6cc9102 := user.Username

	var var_1d38b6cc9102_mapped *structpb.Value

	var var_1d38b6cc9102_err error
	var_1d38b6cc9102_mapped, var_1d38b6cc9102_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_1d38b6cc9102)
	if var_1d38b6cc9102_err != nil {
		panic(var_1d38b6cc9102_err)
	}
	properties["username"] = var_1d38b6cc9102_mapped

	var_878742ec89fa := user.Password

	if var_878742ec89fa != nil {
		var var_878742ec89fa_mapped *structpb.Value

		var var_878742ec89fa_err error
		var_878742ec89fa_mapped, var_878742ec89fa_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_878742ec89fa)
		if var_878742ec89fa_err != nil {
			panic(var_878742ec89fa_err)
		}
		properties["password"] = var_878742ec89fa_mapped
	}

	var_089563314f1a := user.Roles

	if var_089563314f1a != nil {
		var var_089563314f1a_mapped *structpb.Value

		var var_089563314f1a_l []*structpb.Value
		for _, value := range var_089563314f1a {

			var_58cec76a2ad5 := value
			var var_58cec76a2ad5_mapped *structpb.Value

			var var_58cec76a2ad5_err error
			var_58cec76a2ad5_mapped, var_58cec76a2ad5_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_58cec76a2ad5)
			if var_58cec76a2ad5_err != nil {
				panic(var_58cec76a2ad5_err)
			}

			var_089563314f1a_l = append(var_089563314f1a_l, var_58cec76a2ad5_mapped)
		}
		var_089563314f1a_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_089563314f1a_l})
		properties["roles"] = var_089563314f1a_mapped
	}

	var_144a4fce97e4 := user.SecurityConstraints

	if var_144a4fce97e4 != nil {
		var var_144a4fce97e4_mapped *structpb.Value

		var var_144a4fce97e4_l []*structpb.Value
		for _, value := range var_144a4fce97e4 {

			var_5b551f334e8a := value
			var var_5b551f334e8a_mapped *structpb.Value

			var_5b551f334e8a_mapped = structpb.NewStructValue(&structpb.Struct{Fields: SecurityConstraintMapperInstance.ToProperties(var_5b551f334e8a)})

			var_144a4fce97e4_l = append(var_144a4fce97e4_l, var_5b551f334e8a_mapped)
		}
		var_144a4fce97e4_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_144a4fce97e4_l})
		properties["securityConstraints"] = var_144a4fce97e4_mapped
	}

	var_00e00a345383 := user.Details

	if var_00e00a345383 != nil {
		var var_00e00a345383_mapped *structpb.Value

		var var_00e00a345383_err error
		var_00e00a345383_mapped, var_00e00a345383_err = types.ByResourcePropertyType(model.ResourceProperty_OBJECT).Pack(*var_00e00a345383)
		if var_00e00a345383_err != nil {
			panic(var_00e00a345383_err)
		}
		properties["details"] = var_00e00a345383_mapped
	}
	return properties
}

func (m *UserMapper) FromProperties(properties map[string]*structpb.Value) *User {
	var s = m.New()
	if properties["id"] != nil {

		var_8bb2106c5515 := properties["id"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_UUID).UnPack(var_8bb2106c5515)

		if err != nil {
			panic(err)
		}

		var_8bb2106c5515_mapped := new(uuid.UUID)
		*var_8bb2106c5515_mapped = val.(uuid.UUID)

		s.Id = var_8bb2106c5515_mapped
	}
	if properties["version"] != nil {

		var_9c639af5e2b6 := properties["version"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(var_9c639af5e2b6)

		if err != nil {
			panic(err)
		}

		var_9c639af5e2b6_mapped := val.(int32)

		s.Version = var_9c639af5e2b6_mapped
	}
	if properties["createdBy"] != nil {

		var_396484b787bc := properties["createdBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_396484b787bc)

		if err != nil {
			panic(err)
		}

		var_396484b787bc_mapped := new(string)
		*var_396484b787bc_mapped = val.(string)

		s.CreatedBy = var_396484b787bc_mapped
	}
	if properties["updatedBy"] != nil {

		var_e55967325dd1 := properties["updatedBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_e55967325dd1)

		if err != nil {
			panic(err)
		}

		var_e55967325dd1_mapped := new(string)
		*var_e55967325dd1_mapped = val.(string)

		s.UpdatedBy = var_e55967325dd1_mapped
	}
	if properties["createdOn"] != nil {

		var_6d13b1e5ab32 := properties["createdOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_6d13b1e5ab32)

		if err != nil {
			panic(err)
		}

		var_6d13b1e5ab32_mapped := new(time.Time)
		*var_6d13b1e5ab32_mapped = val.(time.Time)

		s.CreatedOn = var_6d13b1e5ab32_mapped
	}
	if properties["updatedOn"] != nil {

		var_5da3599f6267 := properties["updatedOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_5da3599f6267)

		if err != nil {
			panic(err)
		}

		var_5da3599f6267_mapped := new(time.Time)
		*var_5da3599f6267_mapped = val.(time.Time)

		s.UpdatedOn = var_5da3599f6267_mapped
	}
	if properties["username"] != nil {

		var_2f2447b53438 := properties["username"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_2f2447b53438)

		if err != nil {
			panic(err)
		}

		var_2f2447b53438_mapped := val.(string)

		s.Username = var_2f2447b53438_mapped
	}
	if properties["password"] != nil {

		var_fbb5ceae0037 := properties["password"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_fbb5ceae0037)

		if err != nil {
			panic(err)
		}

		var_fbb5ceae0037_mapped := new(string)
		*var_fbb5ceae0037_mapped = val.(string)

		s.Password = var_fbb5ceae0037_mapped
	}
	if properties["roles"] != nil {

		var_6ae8115201bc := properties["roles"]
		var_6ae8115201bc_mapped := []string{}
		for _, v := range var_6ae8115201bc.GetListValue().Values {

			var_c410ee18e6b7 := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_c410ee18e6b7)

			if err != nil {
				panic(err)
			}

			var_c410ee18e6b7_mapped := val.(string)

			var_6ae8115201bc_mapped = append(var_6ae8115201bc_mapped, var_c410ee18e6b7_mapped)
		}

		s.Roles = var_6ae8115201bc_mapped
	}
	if properties["securityConstraints"] != nil {

		var_c74527f31978 := properties["securityConstraints"]
		var_c74527f31978_mapped := []*SecurityConstraint{}
		for _, v := range var_c74527f31978.GetListValue().Values {

			var_c89270e87afa := v
			var_c89270e87afa_mapped := SecurityConstraintMapperInstance.FromProperties(var_c89270e87afa.GetStructValue().Fields)

			var_c74527f31978_mapped = append(var_c74527f31978_mapped, var_c89270e87afa_mapped)
		}

		s.SecurityConstraints = var_c74527f31978_mapped
	}
	if properties["details"] != nil {

		var_487c09a2e785 := properties["details"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_OBJECT).UnPack(var_487c09a2e785)

		if err != nil {
			panic(err)
		}

		var_487c09a2e785_mapped := new(unstructured.Unstructured)
		*var_487c09a2e785_mapped = val.(unstructured.Unstructured)

		s.Details = var_487c09a2e785_mapped
	}
	return s
}
