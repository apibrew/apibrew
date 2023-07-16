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
	return rec
}

func (m *RoleMapper) FromRecord(record *model.Record) *Role {
	return m.FromProperties(record.Properties)
}

func (m *RoleMapper) ToProperties(role *Role) map[string]*structpb.Value {
	var properties = make(map[string]*structpb.Value)

	var_65e42637d48d := role.Id

	if var_65e42637d48d != nil {
		var var_65e42637d48d_mapped *structpb.Value

		var var_65e42637d48d_err error
		var_65e42637d48d_mapped, var_65e42637d48d_err = types.ByResourcePropertyType(model.ResourceProperty_UUID).Pack(*var_65e42637d48d)
		if var_65e42637d48d_err != nil {
			panic(var_65e42637d48d_err)
		}
		properties["id"] = var_65e42637d48d_mapped
	}

	var_1e3ccc07728e := role.Version

	var var_1e3ccc07728e_mapped *structpb.Value

	var var_1e3ccc07728e_err error
	var_1e3ccc07728e_mapped, var_1e3ccc07728e_err = types.ByResourcePropertyType(model.ResourceProperty_INT32).Pack(var_1e3ccc07728e)
	if var_1e3ccc07728e_err != nil {
		panic(var_1e3ccc07728e_err)
	}
	properties["version"] = var_1e3ccc07728e_mapped

	var_b82e0c74be2e := role.CreatedBy

	if var_b82e0c74be2e != nil {
		var var_b82e0c74be2e_mapped *structpb.Value

		var var_b82e0c74be2e_err error
		var_b82e0c74be2e_mapped, var_b82e0c74be2e_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_b82e0c74be2e)
		if var_b82e0c74be2e_err != nil {
			panic(var_b82e0c74be2e_err)
		}
		properties["createdBy"] = var_b82e0c74be2e_mapped
	}

	var_dd65a9188e3f := role.UpdatedBy

	if var_dd65a9188e3f != nil {
		var var_dd65a9188e3f_mapped *structpb.Value

		var var_dd65a9188e3f_err error
		var_dd65a9188e3f_mapped, var_dd65a9188e3f_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_dd65a9188e3f)
		if var_dd65a9188e3f_err != nil {
			panic(var_dd65a9188e3f_err)
		}
		properties["updatedBy"] = var_dd65a9188e3f_mapped
	}

	var_927aa41670de := role.CreatedOn

	if var_927aa41670de != nil {
		var var_927aa41670de_mapped *structpb.Value

		var var_927aa41670de_err error
		var_927aa41670de_mapped, var_927aa41670de_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_927aa41670de)
		if var_927aa41670de_err != nil {
			panic(var_927aa41670de_err)
		}
		properties["createdOn"] = var_927aa41670de_mapped
	}

	var_73371ba749f3 := role.UpdatedOn

	if var_73371ba749f3 != nil {
		var var_73371ba749f3_mapped *structpb.Value

		var var_73371ba749f3_err error
		var_73371ba749f3_mapped, var_73371ba749f3_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_73371ba749f3)
		if var_73371ba749f3_err != nil {
			panic(var_73371ba749f3_err)
		}
		properties["updatedOn"] = var_73371ba749f3_mapped
	}

	var_e8d25ad81293 := role.Name

	var var_e8d25ad81293_mapped *structpb.Value

	var var_e8d25ad81293_err error
	var_e8d25ad81293_mapped, var_e8d25ad81293_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_e8d25ad81293)
	if var_e8d25ad81293_err != nil {
		panic(var_e8d25ad81293_err)
	}
	properties["name"] = var_e8d25ad81293_mapped

	var_9899c6fe877d := role.SecurityConstraints

	if var_9899c6fe877d != nil {
		var var_9899c6fe877d_mapped *structpb.Value

		var var_9899c6fe877d_l []*structpb.Value
		for _, value := range var_9899c6fe877d {

			var_6ad8ee984dfc := value
			var var_6ad8ee984dfc_mapped *structpb.Value

			var_6ad8ee984dfc_mapped = structpb.NewStructValue(&structpb.Struct{Fields: SecurityConstraintMapperInstance.ToProperties(var_6ad8ee984dfc)})

			var_9899c6fe877d_l = append(var_9899c6fe877d_l, var_6ad8ee984dfc_mapped)
		}
		var_9899c6fe877d_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_9899c6fe877d_l})
		properties["securityConstraints"] = var_9899c6fe877d_mapped
	}

	var_1c2662ba8f3e := role.Details

	if var_1c2662ba8f3e != nil {
		var var_1c2662ba8f3e_mapped *structpb.Value

		var var_1c2662ba8f3e_err error
		var_1c2662ba8f3e_mapped, var_1c2662ba8f3e_err = types.ByResourcePropertyType(model.ResourceProperty_OBJECT).Pack(*var_1c2662ba8f3e)
		if var_1c2662ba8f3e_err != nil {
			panic(var_1c2662ba8f3e_err)
		}
		properties["details"] = var_1c2662ba8f3e_mapped
	}
	return properties
}

func (m *RoleMapper) FromProperties(properties map[string]*structpb.Value) *Role {
	var s = m.New()
	if properties["id"] != nil {

		var_e52ae7b6de88 := properties["id"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_UUID).UnPack(var_e52ae7b6de88)

		if err != nil {
			panic(err)
		}

		var_e52ae7b6de88_mapped := new(uuid.UUID)
		*var_e52ae7b6de88_mapped = val.(uuid.UUID)

		s.Id = var_e52ae7b6de88_mapped
	}
	if properties["version"] != nil {

		var_14ee6fcb771d := properties["version"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(var_14ee6fcb771d)

		if err != nil {
			panic(err)
		}

		var_14ee6fcb771d_mapped := val.(int32)

		s.Version = var_14ee6fcb771d_mapped
	}
	if properties["createdBy"] != nil {

		var_da1b3424337d := properties["createdBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_da1b3424337d)

		if err != nil {
			panic(err)
		}

		var_da1b3424337d_mapped := new(string)
		*var_da1b3424337d_mapped = val.(string)

		s.CreatedBy = var_da1b3424337d_mapped
	}
	if properties["updatedBy"] != nil {

		var_0767bccfa43c := properties["updatedBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_0767bccfa43c)

		if err != nil {
			panic(err)
		}

		var_0767bccfa43c_mapped := new(string)
		*var_0767bccfa43c_mapped = val.(string)

		s.UpdatedBy = var_0767bccfa43c_mapped
	}
	if properties["createdOn"] != nil {

		var_9e4959c1ad8f := properties["createdOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_9e4959c1ad8f)

		if err != nil {
			panic(err)
		}

		var_9e4959c1ad8f_mapped := new(time.Time)
		*var_9e4959c1ad8f_mapped = val.(time.Time)

		s.CreatedOn = var_9e4959c1ad8f_mapped
	}
	if properties["updatedOn"] != nil {

		var_4b57c9c64ecb := properties["updatedOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_4b57c9c64ecb)

		if err != nil {
			panic(err)
		}

		var_4b57c9c64ecb_mapped := new(time.Time)
		*var_4b57c9c64ecb_mapped = val.(time.Time)

		s.UpdatedOn = var_4b57c9c64ecb_mapped
	}
	if properties["name"] != nil {

		var_312960b6bc71 := properties["name"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_312960b6bc71)

		if err != nil {
			panic(err)
		}

		var_312960b6bc71_mapped := val.(string)

		s.Name = var_312960b6bc71_mapped
	}
	if properties["securityConstraints"] != nil {

		var_e95a7e714e86 := properties["securityConstraints"]
		var_e95a7e714e86_mapped := []*SecurityConstraint{}
		for _, v := range var_e95a7e714e86.GetListValue().Values {

			var_774e37b1de91 := v
			var_774e37b1de91_mapped := SecurityConstraintMapperInstance.FromProperties(var_774e37b1de91.GetStructValue().Fields)

			var_e95a7e714e86_mapped = append(var_e95a7e714e86_mapped, var_774e37b1de91_mapped)
		}

		s.SecurityConstraints = var_e95a7e714e86_mapped
	}
	if properties["details"] != nil {

		var_eb08fd9a263f := properties["details"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_OBJECT).UnPack(var_eb08fd9a263f)

		if err != nil {
			panic(err)
		}

		var_eb08fd9a263f_mapped := new(unstructured.Unstructured)
		*var_eb08fd9a263f_mapped = val.(unstructured.Unstructured)

		s.Details = var_eb08fd9a263f_mapped
	}
	return s
}
