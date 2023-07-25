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

	var_4f878db25045 := role.Id

	if var_4f878db25045 != nil {
		var var_4f878db25045_mapped *structpb.Value

		var var_4f878db25045_err error
		var_4f878db25045_mapped, var_4f878db25045_err = types.ByResourcePropertyType(model.ResourceProperty_UUID).Pack(*var_4f878db25045)
		if var_4f878db25045_err != nil {
			panic(var_4f878db25045_err)
		}
		properties["id"] = var_4f878db25045_mapped
	}

	var_17b37e17c80f := role.Version

	var var_17b37e17c80f_mapped *structpb.Value

	var var_17b37e17c80f_err error
	var_17b37e17c80f_mapped, var_17b37e17c80f_err = types.ByResourcePropertyType(model.ResourceProperty_INT32).Pack(var_17b37e17c80f)
	if var_17b37e17c80f_err != nil {
		panic(var_17b37e17c80f_err)
	}
	properties["version"] = var_17b37e17c80f_mapped

	var_ce5cd08fdb01 := role.CreatedBy

	if var_ce5cd08fdb01 != nil {
		var var_ce5cd08fdb01_mapped *structpb.Value

		var var_ce5cd08fdb01_err error
		var_ce5cd08fdb01_mapped, var_ce5cd08fdb01_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_ce5cd08fdb01)
		if var_ce5cd08fdb01_err != nil {
			panic(var_ce5cd08fdb01_err)
		}
		properties["createdBy"] = var_ce5cd08fdb01_mapped
	}

	var_f530285c86dc := role.UpdatedBy

	if var_f530285c86dc != nil {
		var var_f530285c86dc_mapped *structpb.Value

		var var_f530285c86dc_err error
		var_f530285c86dc_mapped, var_f530285c86dc_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_f530285c86dc)
		if var_f530285c86dc_err != nil {
			panic(var_f530285c86dc_err)
		}
		properties["updatedBy"] = var_f530285c86dc_mapped
	}

	var_b2b32d03a134 := role.CreatedOn

	if var_b2b32d03a134 != nil {
		var var_b2b32d03a134_mapped *structpb.Value

		var var_b2b32d03a134_err error
		var_b2b32d03a134_mapped, var_b2b32d03a134_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_b2b32d03a134)
		if var_b2b32d03a134_err != nil {
			panic(var_b2b32d03a134_err)
		}
		properties["createdOn"] = var_b2b32d03a134_mapped
	}

	var_b1ef290b8d72 := role.UpdatedOn

	if var_b1ef290b8d72 != nil {
		var var_b1ef290b8d72_mapped *structpb.Value

		var var_b1ef290b8d72_err error
		var_b1ef290b8d72_mapped, var_b1ef290b8d72_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_b1ef290b8d72)
		if var_b1ef290b8d72_err != nil {
			panic(var_b1ef290b8d72_err)
		}
		properties["updatedOn"] = var_b1ef290b8d72_mapped
	}

	var_d4d25df0fe41 := role.Name

	var var_d4d25df0fe41_mapped *structpb.Value

	var var_d4d25df0fe41_err error
	var_d4d25df0fe41_mapped, var_d4d25df0fe41_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_d4d25df0fe41)
	if var_d4d25df0fe41_err != nil {
		panic(var_d4d25df0fe41_err)
	}
	properties["name"] = var_d4d25df0fe41_mapped

	var_96c3510ea93f := role.SecurityConstraints

	if var_96c3510ea93f != nil {
		var var_96c3510ea93f_mapped *structpb.Value

		var var_96c3510ea93f_l []*structpb.Value
		for _, value := range var_96c3510ea93f {

			var_a58d8c02360c := value
			var var_a58d8c02360c_mapped *structpb.Value

			var_a58d8c02360c_mapped = structpb.NewStructValue(&structpb.Struct{Fields: SecurityConstraintMapperInstance.ToProperties(var_a58d8c02360c)})

			var_96c3510ea93f_l = append(var_96c3510ea93f_l, var_a58d8c02360c_mapped)
		}
		var_96c3510ea93f_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_96c3510ea93f_l})
		properties["securityConstraints"] = var_96c3510ea93f_mapped
	}

	var_210ff9665771 := role.Details

	if var_210ff9665771 != nil {
		var var_210ff9665771_mapped *structpb.Value

		var var_210ff9665771_err error
		var_210ff9665771_mapped, var_210ff9665771_err = types.ByResourcePropertyType(model.ResourceProperty_OBJECT).Pack(*var_210ff9665771)
		if var_210ff9665771_err != nil {
			panic(var_210ff9665771_err)
		}
		properties["details"] = var_210ff9665771_mapped
	}
	return properties
}

func (m *RoleMapper) FromProperties(properties map[string]*structpb.Value) *Role {
	var s = m.New()
	if properties["id"] != nil {

		var_1ab1160cb35d := properties["id"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_UUID).UnPack(var_1ab1160cb35d)

		if err != nil {
			panic(err)
		}

		var_1ab1160cb35d_mapped := new(uuid.UUID)
		*var_1ab1160cb35d_mapped = val.(uuid.UUID)

		s.Id = var_1ab1160cb35d_mapped
	}
	if properties["version"] != nil {

		var_6241a92ed74a := properties["version"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(var_6241a92ed74a)

		if err != nil {
			panic(err)
		}

		var_6241a92ed74a_mapped := val.(int32)

		s.Version = var_6241a92ed74a_mapped
	}
	if properties["createdBy"] != nil {

		var_21a90d24624a := properties["createdBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_21a90d24624a)

		if err != nil {
			panic(err)
		}

		var_21a90d24624a_mapped := new(string)
		*var_21a90d24624a_mapped = val.(string)

		s.CreatedBy = var_21a90d24624a_mapped
	}
	if properties["updatedBy"] != nil {

		var_6023922267d3 := properties["updatedBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_6023922267d3)

		if err != nil {
			panic(err)
		}

		var_6023922267d3_mapped := new(string)
		*var_6023922267d3_mapped = val.(string)

		s.UpdatedBy = var_6023922267d3_mapped
	}
	if properties["createdOn"] != nil {

		var_9f81c19387a8 := properties["createdOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_9f81c19387a8)

		if err != nil {
			panic(err)
		}

		var_9f81c19387a8_mapped := new(time.Time)
		*var_9f81c19387a8_mapped = val.(time.Time)

		s.CreatedOn = var_9f81c19387a8_mapped
	}
	if properties["updatedOn"] != nil {

		var_be8f12b8d819 := properties["updatedOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_be8f12b8d819)

		if err != nil {
			panic(err)
		}

		var_be8f12b8d819_mapped := new(time.Time)
		*var_be8f12b8d819_mapped = val.(time.Time)

		s.UpdatedOn = var_be8f12b8d819_mapped
	}
	if properties["name"] != nil {

		var_22c414c77fa7 := properties["name"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_22c414c77fa7)

		if err != nil {
			panic(err)
		}

		var_22c414c77fa7_mapped := val.(string)

		s.Name = var_22c414c77fa7_mapped
	}
	if properties["securityConstraints"] != nil {

		var_9f1b1bfc774e := properties["securityConstraints"]
		var_9f1b1bfc774e_mapped := []*SecurityConstraint{}
		for _, v := range var_9f1b1bfc774e.GetListValue().Values {

			var_1914b7c722ac := v
			var_1914b7c722ac_mapped := SecurityConstraintMapperInstance.FromProperties(var_1914b7c722ac.GetStructValue().Fields)

			var_9f1b1bfc774e_mapped = append(var_9f1b1bfc774e_mapped, var_1914b7c722ac_mapped)
		}

		s.SecurityConstraints = var_9f1b1bfc774e_mapped
	}
	if properties["details"] != nil {

		var_9a2fa78d9819 := properties["details"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_OBJECT).UnPack(var_9a2fa78d9819)

		if err != nil {
			panic(err)
		}

		var_9a2fa78d9819_mapped := new(unstructured.Unstructured)
		*var_9a2fa78d9819_mapped = val.(unstructured.Unstructured)

		s.Details = var_9a2fa78d9819_mapped
	}
	return s
}
