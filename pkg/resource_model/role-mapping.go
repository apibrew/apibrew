package resource_model

import (
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/types"
	"google.golang.org/protobuf/types/known/structpb"
)

import "github.com/google/uuid"
import "time"
import "github.com/apibrew/apibrew/pkg/formats/unstructured"

type RoleMapper struct {
}

func NewRoleMapper() *RoleMapper {
	return &RoleMapper{}
}

var RoleMapperInstance = NewRoleMapper()

func (m *RoleMapper) New() *Role {
	return &Role{}
}

func (m *RoleMapper) ToRecord(role *Role) *model.Record {
	var rec = &model.Record{}
	rec.Properties = m.ToProperties(role)

	if role.Id != nil {
		rec.Id = role.Id.String()
	}

	return rec
}

func (m *RoleMapper) FromRecord(record *model.Record) *Role {
	return m.FromProperties(record.Properties)
}

func (m *RoleMapper) ToProperties(role *Role) map[string]*structpb.Value {
	var properties = make(map[string]*structpb.Value)

	var_d55862ffa5a2 := role.Id

	if var_d55862ffa5a2 != nil {
		var var_d55862ffa5a2_mapped *structpb.Value

		var var_d55862ffa5a2_err error
		var_d55862ffa5a2_mapped, var_d55862ffa5a2_err = types.ByResourcePropertyType(model.ResourceProperty_UUID).Pack(*var_d55862ffa5a2)
		if var_d55862ffa5a2_err != nil {
			panic(var_d55862ffa5a2_err)
		}
		properties["id"] = var_d55862ffa5a2_mapped
	}

	var_41ad2c4ecd91 := role.Version

	var var_41ad2c4ecd91_mapped *structpb.Value

	var var_41ad2c4ecd91_err error
	var_41ad2c4ecd91_mapped, var_41ad2c4ecd91_err = types.ByResourcePropertyType(model.ResourceProperty_INT32).Pack(var_41ad2c4ecd91)
	if var_41ad2c4ecd91_err != nil {
		panic(var_41ad2c4ecd91_err)
	}
	properties["version"] = var_41ad2c4ecd91_mapped

	var_a0859a677297 := role.CreatedBy

	if var_a0859a677297 != nil {
		var var_a0859a677297_mapped *structpb.Value

		var var_a0859a677297_err error
		var_a0859a677297_mapped, var_a0859a677297_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_a0859a677297)
		if var_a0859a677297_err != nil {
			panic(var_a0859a677297_err)
		}
		properties["createdBy"] = var_a0859a677297_mapped
	}

	var_87b34cb86508 := role.UpdatedBy

	if var_87b34cb86508 != nil {
		var var_87b34cb86508_mapped *structpb.Value

		var var_87b34cb86508_err error
		var_87b34cb86508_mapped, var_87b34cb86508_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_87b34cb86508)
		if var_87b34cb86508_err != nil {
			panic(var_87b34cb86508_err)
		}
		properties["updatedBy"] = var_87b34cb86508_mapped
	}

	var_2a2ed3269c47 := role.CreatedOn

	if var_2a2ed3269c47 != nil {
		var var_2a2ed3269c47_mapped *structpb.Value

		var var_2a2ed3269c47_err error
		var_2a2ed3269c47_mapped, var_2a2ed3269c47_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_2a2ed3269c47)
		if var_2a2ed3269c47_err != nil {
			panic(var_2a2ed3269c47_err)
		}
		properties["createdOn"] = var_2a2ed3269c47_mapped
	}

	var_e83287a3253c := role.UpdatedOn

	if var_e83287a3253c != nil {
		var var_e83287a3253c_mapped *structpb.Value

		var var_e83287a3253c_err error
		var_e83287a3253c_mapped, var_e83287a3253c_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_e83287a3253c)
		if var_e83287a3253c_err != nil {
			panic(var_e83287a3253c_err)
		}
		properties["updatedOn"] = var_e83287a3253c_mapped
	}

	var_bf8f5b17b592 := role.Name

	var var_bf8f5b17b592_mapped *structpb.Value

	var var_bf8f5b17b592_err error
	var_bf8f5b17b592_mapped, var_bf8f5b17b592_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_bf8f5b17b592)
	if var_bf8f5b17b592_err != nil {
		panic(var_bf8f5b17b592_err)
	}
	properties["name"] = var_bf8f5b17b592_mapped

	var_22b5fa782a6d := role.SecurityConstraints

	if var_22b5fa782a6d != nil {
		var var_22b5fa782a6d_mapped *structpb.Value

		var var_22b5fa782a6d_l []*structpb.Value
		for _, value := range var_22b5fa782a6d {

			var_cb272d4f0933 := value
			var var_cb272d4f0933_mapped *structpb.Value

			var_cb272d4f0933_mapped = structpb.NewStructValue(&structpb.Struct{Fields: SecurityConstraintMapperInstance.ToProperties(var_cb272d4f0933)})

			var_22b5fa782a6d_l = append(var_22b5fa782a6d_l, var_cb272d4f0933_mapped)
		}
		var_22b5fa782a6d_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_22b5fa782a6d_l})
		properties["securityConstraints"] = var_22b5fa782a6d_mapped
	}

	var_abf774b14247 := role.Details

	if var_abf774b14247 != nil {
		var var_abf774b14247_mapped *structpb.Value

		var var_abf774b14247_err error
		var_abf774b14247_mapped, var_abf774b14247_err = types.ByResourcePropertyType(model.ResourceProperty_OBJECT).Pack(*var_abf774b14247)
		if var_abf774b14247_err != nil {
			panic(var_abf774b14247_err)
		}
		properties["details"] = var_abf774b14247_mapped
	}
	return properties
}

func (m *RoleMapper) FromProperties(properties map[string]*structpb.Value) *Role {
	var s = m.New()
	if properties["id"] != nil {

		var_f765d054f4ed := properties["id"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_UUID).UnPack(var_f765d054f4ed)

		if err != nil {
			panic(err)
		}

		var_f765d054f4ed_mapped := new(uuid.UUID)
		*var_f765d054f4ed_mapped = val.(uuid.UUID)

		s.Id = var_f765d054f4ed_mapped
	}
	if properties["version"] != nil {

		var_9e7ae6b7dada := properties["version"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(var_9e7ae6b7dada)

		if err != nil {
			panic(err)
		}

		var_9e7ae6b7dada_mapped := val.(int32)

		s.Version = var_9e7ae6b7dada_mapped
	}
	if properties["createdBy"] != nil {

		var_f7962a821310 := properties["createdBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_f7962a821310)

		if err != nil {
			panic(err)
		}

		var_f7962a821310_mapped := new(string)
		*var_f7962a821310_mapped = val.(string)

		s.CreatedBy = var_f7962a821310_mapped
	}
	if properties["updatedBy"] != nil {

		var_6de2d2e08859 := properties["updatedBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_6de2d2e08859)

		if err != nil {
			panic(err)
		}

		var_6de2d2e08859_mapped := new(string)
		*var_6de2d2e08859_mapped = val.(string)

		s.UpdatedBy = var_6de2d2e08859_mapped
	}
	if properties["createdOn"] != nil {

		var_f700b9fcbe0a := properties["createdOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_f700b9fcbe0a)

		if err != nil {
			panic(err)
		}

		var_f700b9fcbe0a_mapped := new(time.Time)
		*var_f700b9fcbe0a_mapped = val.(time.Time)

		s.CreatedOn = var_f700b9fcbe0a_mapped
	}
	if properties["updatedOn"] != nil {

		var_a480282e6a5b := properties["updatedOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_a480282e6a5b)

		if err != nil {
			panic(err)
		}

		var_a480282e6a5b_mapped := new(time.Time)
		*var_a480282e6a5b_mapped = val.(time.Time)

		s.UpdatedOn = var_a480282e6a5b_mapped
	}
	if properties["name"] != nil {

		var_b1178cdfaf01 := properties["name"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_b1178cdfaf01)

		if err != nil {
			panic(err)
		}

		var_b1178cdfaf01_mapped := val.(string)

		s.Name = var_b1178cdfaf01_mapped
	}
	if properties["securityConstraints"] != nil {

		var_e2ae70d88de2 := properties["securityConstraints"]
		var_e2ae70d88de2_mapped := []*SecurityConstraint{}
		for _, v := range var_e2ae70d88de2.GetListValue().Values {

			var_8ae999589b16 := v
			var_8ae999589b16_mapped := SecurityConstraintMapperInstance.FromProperties(var_8ae999589b16.GetStructValue().Fields)

			var_e2ae70d88de2_mapped = append(var_e2ae70d88de2_mapped, var_8ae999589b16_mapped)
		}

		s.SecurityConstraints = var_e2ae70d88de2_mapped
	}
	if properties["details"] != nil {

		var_65da11994217 := properties["details"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_OBJECT).UnPack(var_65da11994217)

		if err != nil {
			panic(err)
		}

		var_65da11994217_mapped := new(unstructured.Unstructured)
		*var_65da11994217_mapped = val.(unstructured.Unstructured)

		s.Details = var_65da11994217_mapped
	}
	return s
}
