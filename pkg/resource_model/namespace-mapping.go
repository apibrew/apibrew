package resource_model

import (
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/types"
	"google.golang.org/protobuf/types/known/structpb"
)

import "github.com/google/uuid"
import "time"
import "github.com/apibrew/apibrew/pkg/formats/unstructured"

type NamespaceMapper struct {
}

func NewNamespaceMapper() *NamespaceMapper {
	return &NamespaceMapper{}
}

var NamespaceMapperInstance = NewNamespaceMapper()

func (m *NamespaceMapper) New() *Namespace {
	return &Namespace{}
}

func (m *NamespaceMapper) ToRecord(namespace *Namespace) *model.Record {
	var rec = &model.Record{}
	rec.Properties = m.ToProperties(namespace)

	if namespace.Id != nil {
		rec.Id = namespace.Id.String()
	}

	return rec
}

func (m *NamespaceMapper) FromRecord(record *model.Record) *Namespace {
	return m.FromProperties(record.Properties)
}

func (m *NamespaceMapper) ToProperties(namespace *Namespace) map[string]*structpb.Value {
	var properties = make(map[string]*structpb.Value)

	var_cf783e1cc028 := namespace.Id

	if var_cf783e1cc028 != nil {
		var var_cf783e1cc028_mapped *structpb.Value

		var var_cf783e1cc028_err error
		var_cf783e1cc028_mapped, var_cf783e1cc028_err = types.ByResourcePropertyType(model.ResourceProperty_UUID).Pack(*var_cf783e1cc028)
		if var_cf783e1cc028_err != nil {
			panic(var_cf783e1cc028_err)
		}
		properties["id"] = var_cf783e1cc028_mapped
	}

	var_70b3f2f0f240 := namespace.Version

	var var_70b3f2f0f240_mapped *structpb.Value

	var var_70b3f2f0f240_err error
	var_70b3f2f0f240_mapped, var_70b3f2f0f240_err = types.ByResourcePropertyType(model.ResourceProperty_INT32).Pack(var_70b3f2f0f240)
	if var_70b3f2f0f240_err != nil {
		panic(var_70b3f2f0f240_err)
	}
	properties["version"] = var_70b3f2f0f240_mapped

	var_c27dda1ed4d6 := namespace.CreatedBy

	if var_c27dda1ed4d6 != nil {
		var var_c27dda1ed4d6_mapped *structpb.Value

		var var_c27dda1ed4d6_err error
		var_c27dda1ed4d6_mapped, var_c27dda1ed4d6_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_c27dda1ed4d6)
		if var_c27dda1ed4d6_err != nil {
			panic(var_c27dda1ed4d6_err)
		}
		properties["createdBy"] = var_c27dda1ed4d6_mapped
	}

	var_8dda6ca383d9 := namespace.UpdatedBy

	if var_8dda6ca383d9 != nil {
		var var_8dda6ca383d9_mapped *structpb.Value

		var var_8dda6ca383d9_err error
		var_8dda6ca383d9_mapped, var_8dda6ca383d9_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_8dda6ca383d9)
		if var_8dda6ca383d9_err != nil {
			panic(var_8dda6ca383d9_err)
		}
		properties["updatedBy"] = var_8dda6ca383d9_mapped
	}

	var_c9fa88641e60 := namespace.CreatedOn

	if var_c9fa88641e60 != nil {
		var var_c9fa88641e60_mapped *structpb.Value

		var var_c9fa88641e60_err error
		var_c9fa88641e60_mapped, var_c9fa88641e60_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_c9fa88641e60)
		if var_c9fa88641e60_err != nil {
			panic(var_c9fa88641e60_err)
		}
		properties["createdOn"] = var_c9fa88641e60_mapped
	}

	var_e1f5255434f9 := namespace.UpdatedOn

	if var_e1f5255434f9 != nil {
		var var_e1f5255434f9_mapped *structpb.Value

		var var_e1f5255434f9_err error
		var_e1f5255434f9_mapped, var_e1f5255434f9_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_e1f5255434f9)
		if var_e1f5255434f9_err != nil {
			panic(var_e1f5255434f9_err)
		}
		properties["updatedOn"] = var_e1f5255434f9_mapped
	}

	var_419535f791e6 := namespace.Name

	var var_419535f791e6_mapped *structpb.Value

	var var_419535f791e6_err error
	var_419535f791e6_mapped, var_419535f791e6_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_419535f791e6)
	if var_419535f791e6_err != nil {
		panic(var_419535f791e6_err)
	}
	properties["name"] = var_419535f791e6_mapped

	var_f99ddf4f39c2 := namespace.Description

	if var_f99ddf4f39c2 != nil {
		var var_f99ddf4f39c2_mapped *structpb.Value

		var var_f99ddf4f39c2_err error
		var_f99ddf4f39c2_mapped, var_f99ddf4f39c2_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_f99ddf4f39c2)
		if var_f99ddf4f39c2_err != nil {
			panic(var_f99ddf4f39c2_err)
		}
		properties["description"] = var_f99ddf4f39c2_mapped
	}

	var_e0689830037f := namespace.Details

	if var_e0689830037f != nil {
		var var_e0689830037f_mapped *structpb.Value

		var var_e0689830037f_err error
		var_e0689830037f_mapped, var_e0689830037f_err = types.ByResourcePropertyType(model.ResourceProperty_OBJECT).Pack(*var_e0689830037f)
		if var_e0689830037f_err != nil {
			panic(var_e0689830037f_err)
		}
		properties["details"] = var_e0689830037f_mapped
	}

	var_1203325f1493 := namespace.SecurityConstraints

	if var_1203325f1493 != nil {
		var var_1203325f1493_mapped *structpb.Value

		var var_1203325f1493_l []*structpb.Value
		for _, value := range var_1203325f1493 {

			var_ad067c8dfef8 := value
			var var_ad067c8dfef8_mapped *structpb.Value

			var_ad067c8dfef8_mapped = structpb.NewStructValue(&structpb.Struct{Fields: SecurityConstraintMapperInstance.ToProperties(var_ad067c8dfef8)})

			var_1203325f1493_l = append(var_1203325f1493_l, var_ad067c8dfef8_mapped)
		}
		var_1203325f1493_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_1203325f1493_l})
		properties["securityConstraints"] = var_1203325f1493_mapped
	}
	return properties
}

func (m *NamespaceMapper) FromProperties(properties map[string]*structpb.Value) *Namespace {
	var s = m.New()
	if properties["id"] != nil {

		var_ff1ac95fd1eb := properties["id"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_UUID).UnPack(var_ff1ac95fd1eb)

		if err != nil {
			panic(err)
		}

		var_ff1ac95fd1eb_mapped := new(uuid.UUID)
		*var_ff1ac95fd1eb_mapped = val.(uuid.UUID)

		s.Id = var_ff1ac95fd1eb_mapped
	}
	if properties["version"] != nil {

		var_a42a5ee8dc3c := properties["version"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(var_a42a5ee8dc3c)

		if err != nil {
			panic(err)
		}

		var_a42a5ee8dc3c_mapped := val.(int32)

		s.Version = var_a42a5ee8dc3c_mapped
	}
	if properties["createdBy"] != nil {

		var_d7428c54f11b := properties["createdBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_d7428c54f11b)

		if err != nil {
			panic(err)
		}

		var_d7428c54f11b_mapped := new(string)
		*var_d7428c54f11b_mapped = val.(string)

		s.CreatedBy = var_d7428c54f11b_mapped
	}
	if properties["updatedBy"] != nil {

		var_65ce37f276a8 := properties["updatedBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_65ce37f276a8)

		if err != nil {
			panic(err)
		}

		var_65ce37f276a8_mapped := new(string)
		*var_65ce37f276a8_mapped = val.(string)

		s.UpdatedBy = var_65ce37f276a8_mapped
	}
	if properties["createdOn"] != nil {

		var_6d6a9cbea349 := properties["createdOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_6d6a9cbea349)

		if err != nil {
			panic(err)
		}

		var_6d6a9cbea349_mapped := new(time.Time)
		*var_6d6a9cbea349_mapped = val.(time.Time)

		s.CreatedOn = var_6d6a9cbea349_mapped
	}
	if properties["updatedOn"] != nil {

		var_788ac52d55ca := properties["updatedOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_788ac52d55ca)

		if err != nil {
			panic(err)
		}

		var_788ac52d55ca_mapped := new(time.Time)
		*var_788ac52d55ca_mapped = val.(time.Time)

		s.UpdatedOn = var_788ac52d55ca_mapped
	}
	if properties["name"] != nil {

		var_219f107d0019 := properties["name"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_219f107d0019)

		if err != nil {
			panic(err)
		}

		var_219f107d0019_mapped := val.(string)

		s.Name = var_219f107d0019_mapped
	}
	if properties["description"] != nil {

		var_68291097f6f1 := properties["description"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_68291097f6f1)

		if err != nil {
			panic(err)
		}

		var_68291097f6f1_mapped := new(string)
		*var_68291097f6f1_mapped = val.(string)

		s.Description = var_68291097f6f1_mapped
	}
	if properties["details"] != nil {

		var_7c8d125cd662 := properties["details"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_OBJECT).UnPack(var_7c8d125cd662)

		if err != nil {
			panic(err)
		}

		var_7c8d125cd662_mapped := new(unstructured.Unstructured)
		*var_7c8d125cd662_mapped = val.(unstructured.Unstructured)

		s.Details = var_7c8d125cd662_mapped
	}
	if properties["securityConstraints"] != nil {

		var_2a8c8cae3443 := properties["securityConstraints"]
		var_2a8c8cae3443_mapped := []*SecurityConstraint{}
		for _, v := range var_2a8c8cae3443.GetListValue().Values {

			var_208db97479ff := v
			var_208db97479ff_mapped := SecurityConstraintMapperInstance.FromProperties(var_208db97479ff.GetStructValue().Fields)

			var_2a8c8cae3443_mapped = append(var_2a8c8cae3443_mapped, var_208db97479ff_mapped)
		}

		s.SecurityConstraints = var_2a8c8cae3443_mapped
	}
	return s
}
