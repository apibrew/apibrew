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

	var_9963c5d55fa5 := user.Id

	if var_9963c5d55fa5 != nil {
		var var_9963c5d55fa5_mapped *structpb.Value

		var var_9963c5d55fa5_err error
		var_9963c5d55fa5_mapped, var_9963c5d55fa5_err = types.ByResourcePropertyType(model.ResourceProperty_UUID).Pack(*var_9963c5d55fa5)
		if var_9963c5d55fa5_err != nil {
			panic(var_9963c5d55fa5_err)
		}
		properties["id"] = var_9963c5d55fa5_mapped
	}

	var_03a4fe307714 := user.Version

	var var_03a4fe307714_mapped *structpb.Value

	var var_03a4fe307714_err error
	var_03a4fe307714_mapped, var_03a4fe307714_err = types.ByResourcePropertyType(model.ResourceProperty_INT32).Pack(var_03a4fe307714)
	if var_03a4fe307714_err != nil {
		panic(var_03a4fe307714_err)
	}
	properties["version"] = var_03a4fe307714_mapped

	var_367acd76fe2b := user.CreatedBy

	if var_367acd76fe2b != nil {
		var var_367acd76fe2b_mapped *structpb.Value

		var var_367acd76fe2b_err error
		var_367acd76fe2b_mapped, var_367acd76fe2b_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_367acd76fe2b)
		if var_367acd76fe2b_err != nil {
			panic(var_367acd76fe2b_err)
		}
		properties["createdBy"] = var_367acd76fe2b_mapped
	}

	var_9accad59fc4e := user.UpdatedBy

	if var_9accad59fc4e != nil {
		var var_9accad59fc4e_mapped *structpb.Value

		var var_9accad59fc4e_err error
		var_9accad59fc4e_mapped, var_9accad59fc4e_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_9accad59fc4e)
		if var_9accad59fc4e_err != nil {
			panic(var_9accad59fc4e_err)
		}
		properties["updatedBy"] = var_9accad59fc4e_mapped
	}

	var_ff5832c3b511 := user.CreatedOn

	if var_ff5832c3b511 != nil {
		var var_ff5832c3b511_mapped *structpb.Value

		var var_ff5832c3b511_err error
		var_ff5832c3b511_mapped, var_ff5832c3b511_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_ff5832c3b511)
		if var_ff5832c3b511_err != nil {
			panic(var_ff5832c3b511_err)
		}
		properties["createdOn"] = var_ff5832c3b511_mapped
	}

	var_2738ba5ac490 := user.UpdatedOn

	if var_2738ba5ac490 != nil {
		var var_2738ba5ac490_mapped *structpb.Value

		var var_2738ba5ac490_err error
		var_2738ba5ac490_mapped, var_2738ba5ac490_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_2738ba5ac490)
		if var_2738ba5ac490_err != nil {
			panic(var_2738ba5ac490_err)
		}
		properties["updatedOn"] = var_2738ba5ac490_mapped
	}

	var_818db23ac88b := user.Username

	var var_818db23ac88b_mapped *structpb.Value

	var var_818db23ac88b_err error
	var_818db23ac88b_mapped, var_818db23ac88b_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_818db23ac88b)
	if var_818db23ac88b_err != nil {
		panic(var_818db23ac88b_err)
	}
	properties["username"] = var_818db23ac88b_mapped

	var_28cc82a241c7 := user.Password

	if var_28cc82a241c7 != nil {
		var var_28cc82a241c7_mapped *structpb.Value

		var var_28cc82a241c7_err error
		var_28cc82a241c7_mapped, var_28cc82a241c7_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_28cc82a241c7)
		if var_28cc82a241c7_err != nil {
			panic(var_28cc82a241c7_err)
		}
		properties["password"] = var_28cc82a241c7_mapped
	}

	var_0cbdeddc29a3 := user.Roles

	if var_0cbdeddc29a3 != nil {
		var var_0cbdeddc29a3_mapped *structpb.Value

		var var_0cbdeddc29a3_l []*structpb.Value
		for _, value := range var_0cbdeddc29a3 {

			var_6dc6dfaf6b8f := value
			var var_6dc6dfaf6b8f_mapped *structpb.Value

			var_6dc6dfaf6b8f_mapped = structpb.NewStructValue(&structpb.Struct{Fields: RoleMapperInstance.ToProperties(var_6dc6dfaf6b8f)})

			var_0cbdeddc29a3_l = append(var_0cbdeddc29a3_l, var_6dc6dfaf6b8f_mapped)
		}
		var_0cbdeddc29a3_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_0cbdeddc29a3_l})
		properties["roles"] = var_0cbdeddc29a3_mapped
	}

	var_551af8519dc9 := user.SecurityConstraints

	if var_551af8519dc9 != nil {
		var var_551af8519dc9_mapped *structpb.Value

		var var_551af8519dc9_l []*structpb.Value
		for _, value := range var_551af8519dc9 {

			var_9b5ce399b935 := value
			var var_9b5ce399b935_mapped *structpb.Value

			var_9b5ce399b935_mapped = structpb.NewStructValue(&structpb.Struct{Fields: SecurityConstraintMapperInstance.ToProperties(var_9b5ce399b935)})

			var_551af8519dc9_l = append(var_551af8519dc9_l, var_9b5ce399b935_mapped)
		}
		var_551af8519dc9_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_551af8519dc9_l})
		properties["securityConstraints"] = var_551af8519dc9_mapped
	}

	var_8955c97e0e8f := user.Details

	if var_8955c97e0e8f != nil {
		var var_8955c97e0e8f_mapped *structpb.Value

		var var_8955c97e0e8f_err error
		var_8955c97e0e8f_mapped, var_8955c97e0e8f_err = types.ByResourcePropertyType(model.ResourceProperty_OBJECT).Pack(*var_8955c97e0e8f)
		if var_8955c97e0e8f_err != nil {
			panic(var_8955c97e0e8f_err)
		}
		properties["details"] = var_8955c97e0e8f_mapped
	}
	return properties
}

func (m *UserMapper) FromProperties(properties map[string]*structpb.Value) *User {
	var s = m.New()
	if properties["id"] != nil {

		var_a19824080054 := properties["id"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_UUID).UnPack(var_a19824080054)

		if err != nil {
			panic(err)
		}

		var_a19824080054_mapped := new(uuid.UUID)
		*var_a19824080054_mapped = val.(uuid.UUID)

		s.Id = var_a19824080054_mapped
	}
	if properties["version"] != nil {

		var_6ec0711c86c1 := properties["version"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(var_6ec0711c86c1)

		if err != nil {
			panic(err)
		}

		var_6ec0711c86c1_mapped := val.(int32)

		s.Version = var_6ec0711c86c1_mapped
	}
	if properties["createdBy"] != nil {

		var_7a53db3a6e7c := properties["createdBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_7a53db3a6e7c)

		if err != nil {
			panic(err)
		}

		var_7a53db3a6e7c_mapped := new(string)
		*var_7a53db3a6e7c_mapped = val.(string)

		s.CreatedBy = var_7a53db3a6e7c_mapped
	}
	if properties["updatedBy"] != nil {

		var_716c56772b0c := properties["updatedBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_716c56772b0c)

		if err != nil {
			panic(err)
		}

		var_716c56772b0c_mapped := new(string)
		*var_716c56772b0c_mapped = val.(string)

		s.UpdatedBy = var_716c56772b0c_mapped
	}
	if properties["createdOn"] != nil {

		var_b7da16314e45 := properties["createdOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_b7da16314e45)

		if err != nil {
			panic(err)
		}

		var_b7da16314e45_mapped := new(time.Time)
		*var_b7da16314e45_mapped = val.(time.Time)

		s.CreatedOn = var_b7da16314e45_mapped
	}
	if properties["updatedOn"] != nil {

		var_b6a7847c20a9 := properties["updatedOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_b6a7847c20a9)

		if err != nil {
			panic(err)
		}

		var_b6a7847c20a9_mapped := new(time.Time)
		*var_b6a7847c20a9_mapped = val.(time.Time)

		s.UpdatedOn = var_b6a7847c20a9_mapped
	}
	if properties["username"] != nil {

		var_427f2cfb79e6 := properties["username"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_427f2cfb79e6)

		if err != nil {
			panic(err)
		}

		var_427f2cfb79e6_mapped := val.(string)

		s.Username = var_427f2cfb79e6_mapped
	}
	if properties["password"] != nil {

		var_af197ad99634 := properties["password"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_af197ad99634)

		if err != nil {
			panic(err)
		}

		var_af197ad99634_mapped := new(string)
		*var_af197ad99634_mapped = val.(string)

		s.Password = var_af197ad99634_mapped
	}
	if properties["roles"] != nil {

		var_3f24402c24d5 := properties["roles"]
		var_3f24402c24d5_mapped := []*Role{}
		for _, v := range var_3f24402c24d5.GetListValue().Values {

			var_a26689a6bea7 := v
			var_a26689a6bea7_mapped := RoleMapperInstance.FromProperties(var_a26689a6bea7.GetStructValue().Fields)

			var_3f24402c24d5_mapped = append(var_3f24402c24d5_mapped, var_a26689a6bea7_mapped)
		}

		s.Roles = var_3f24402c24d5_mapped
	}
	if properties["securityConstraints"] != nil {

		var_6875dc7c282d := properties["securityConstraints"]
		var_6875dc7c282d_mapped := []*SecurityConstraint{}
		for _, v := range var_6875dc7c282d.GetListValue().Values {

			var_ff4f723590d8 := v
			var_ff4f723590d8_mapped := SecurityConstraintMapperInstance.FromProperties(var_ff4f723590d8.GetStructValue().Fields)

			var_6875dc7c282d_mapped = append(var_6875dc7c282d_mapped, var_ff4f723590d8_mapped)
		}

		s.SecurityConstraints = var_6875dc7c282d_mapped
	}
	if properties["details"] != nil {

		var_c4a794218f15 := properties["details"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_OBJECT).UnPack(var_c4a794218f15)

		if err != nil {
			panic(err)
		}

		var_c4a794218f15_mapped := new(unstructured.Unstructured)
		*var_c4a794218f15_mapped = val.(unstructured.Unstructured)

		s.Details = var_c4a794218f15_mapped
	}
	return s
}
