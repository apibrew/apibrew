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

	var_330741bc581b := namespace.Id

	if var_330741bc581b != nil {
		var var_330741bc581b_mapped *structpb.Value

		var var_330741bc581b_err error
		var_330741bc581b_mapped, var_330741bc581b_err = types.ByResourcePropertyType(model.ResourceProperty_UUID).Pack(*var_330741bc581b)
		if var_330741bc581b_err != nil {
			panic(var_330741bc581b_err)
		}
		properties["id"] = var_330741bc581b_mapped
	}

	var_b9bbe5775a9e := namespace.Version

	var var_b9bbe5775a9e_mapped *structpb.Value

	var var_b9bbe5775a9e_err error
	var_b9bbe5775a9e_mapped, var_b9bbe5775a9e_err = types.ByResourcePropertyType(model.ResourceProperty_INT32).Pack(var_b9bbe5775a9e)
	if var_b9bbe5775a9e_err != nil {
		panic(var_b9bbe5775a9e_err)
	}
	properties["version"] = var_b9bbe5775a9e_mapped

	var_a86dda635d7b := namespace.CreatedBy

	if var_a86dda635d7b != nil {
		var var_a86dda635d7b_mapped *structpb.Value

		var var_a86dda635d7b_err error
		var_a86dda635d7b_mapped, var_a86dda635d7b_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_a86dda635d7b)
		if var_a86dda635d7b_err != nil {
			panic(var_a86dda635d7b_err)
		}
		properties["createdBy"] = var_a86dda635d7b_mapped
	}

	var_ca647086a539 := namespace.UpdatedBy

	if var_ca647086a539 != nil {
		var var_ca647086a539_mapped *structpb.Value

		var var_ca647086a539_err error
		var_ca647086a539_mapped, var_ca647086a539_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_ca647086a539)
		if var_ca647086a539_err != nil {
			panic(var_ca647086a539_err)
		}
		properties["updatedBy"] = var_ca647086a539_mapped
	}

	var_117be9ec24a9 := namespace.CreatedOn

	if var_117be9ec24a9 != nil {
		var var_117be9ec24a9_mapped *structpb.Value

		var var_117be9ec24a9_err error
		var_117be9ec24a9_mapped, var_117be9ec24a9_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_117be9ec24a9)
		if var_117be9ec24a9_err != nil {
			panic(var_117be9ec24a9_err)
		}
		properties["createdOn"] = var_117be9ec24a9_mapped
	}

	var_3ec8a25a3176 := namespace.UpdatedOn

	if var_3ec8a25a3176 != nil {
		var var_3ec8a25a3176_mapped *structpb.Value

		var var_3ec8a25a3176_err error
		var_3ec8a25a3176_mapped, var_3ec8a25a3176_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_3ec8a25a3176)
		if var_3ec8a25a3176_err != nil {
			panic(var_3ec8a25a3176_err)
		}
		properties["updatedOn"] = var_3ec8a25a3176_mapped
	}

	var_1b7abdd8868a := namespace.Name

	var var_1b7abdd8868a_mapped *structpb.Value

	var var_1b7abdd8868a_err error
	var_1b7abdd8868a_mapped, var_1b7abdd8868a_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_1b7abdd8868a)
	if var_1b7abdd8868a_err != nil {
		panic(var_1b7abdd8868a_err)
	}
	properties["name"] = var_1b7abdd8868a_mapped

	var_dca9c6b620d2 := namespace.Description

	if var_dca9c6b620d2 != nil {
		var var_dca9c6b620d2_mapped *structpb.Value

		var var_dca9c6b620d2_err error
		var_dca9c6b620d2_mapped, var_dca9c6b620d2_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_dca9c6b620d2)
		if var_dca9c6b620d2_err != nil {
			panic(var_dca9c6b620d2_err)
		}
		properties["description"] = var_dca9c6b620d2_mapped
	}

	var_b33abdaf8d2a := namespace.Details

	if var_b33abdaf8d2a != nil {
		var var_b33abdaf8d2a_mapped *structpb.Value

		var var_b33abdaf8d2a_err error
		var_b33abdaf8d2a_mapped, var_b33abdaf8d2a_err = types.ByResourcePropertyType(model.ResourceProperty_OBJECT).Pack(*var_b33abdaf8d2a)
		if var_b33abdaf8d2a_err != nil {
			panic(var_b33abdaf8d2a_err)
		}
		properties["details"] = var_b33abdaf8d2a_mapped
	}

	var_eaf5932c8f2c := namespace.SecurityConstraints

	if var_eaf5932c8f2c != nil {
		var var_eaf5932c8f2c_mapped *structpb.Value

		var var_eaf5932c8f2c_l []*structpb.Value
		for _, value := range var_eaf5932c8f2c {

			var_bcbe5c285db3 := value
			var var_bcbe5c285db3_mapped *structpb.Value

			var_bcbe5c285db3_mapped = structpb.NewStructValue(&structpb.Struct{Fields: SecurityConstraintMapperInstance.ToProperties(var_bcbe5c285db3)})

			var_eaf5932c8f2c_l = append(var_eaf5932c8f2c_l, var_bcbe5c285db3_mapped)
		}
		var_eaf5932c8f2c_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_eaf5932c8f2c_l})
		properties["securityConstraints"] = var_eaf5932c8f2c_mapped
	}
	return properties
}

func (m *NamespaceMapper) FromProperties(properties map[string]*structpb.Value) *Namespace {
	var s = m.New()
	if properties["id"] != nil {

		var_36e884fe2614 := properties["id"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_UUID).UnPack(var_36e884fe2614)

		if err != nil {
			panic(err)
		}

		var_36e884fe2614_mapped := new(uuid.UUID)
		*var_36e884fe2614_mapped = val.(uuid.UUID)

		s.Id = var_36e884fe2614_mapped
	}
	if properties["version"] != nil {

		var_fab3835deccb := properties["version"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(var_fab3835deccb)

		if err != nil {
			panic(err)
		}

		var_fab3835deccb_mapped := val.(int32)

		s.Version = var_fab3835deccb_mapped
	}
	if properties["createdBy"] != nil {

		var_1486841d99ff := properties["createdBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_1486841d99ff)

		if err != nil {
			panic(err)
		}

		var_1486841d99ff_mapped := new(string)
		*var_1486841d99ff_mapped = val.(string)

		s.CreatedBy = var_1486841d99ff_mapped
	}
	if properties["updatedBy"] != nil {

		var_8578ba264279 := properties["updatedBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_8578ba264279)

		if err != nil {
			panic(err)
		}

		var_8578ba264279_mapped := new(string)
		*var_8578ba264279_mapped = val.(string)

		s.UpdatedBy = var_8578ba264279_mapped
	}
	if properties["createdOn"] != nil {

		var_214a67a88a4f := properties["createdOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_214a67a88a4f)

		if err != nil {
			panic(err)
		}

		var_214a67a88a4f_mapped := new(time.Time)
		*var_214a67a88a4f_mapped = val.(time.Time)

		s.CreatedOn = var_214a67a88a4f_mapped
	}
	if properties["updatedOn"] != nil {

		var_6614c1c75f49 := properties["updatedOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_6614c1c75f49)

		if err != nil {
			panic(err)
		}

		var_6614c1c75f49_mapped := new(time.Time)
		*var_6614c1c75f49_mapped = val.(time.Time)

		s.UpdatedOn = var_6614c1c75f49_mapped
	}
	if properties["name"] != nil {

		var_b8703c8728b4 := properties["name"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_b8703c8728b4)

		if err != nil {
			panic(err)
		}

		var_b8703c8728b4_mapped := val.(string)

		s.Name = var_b8703c8728b4_mapped
	}
	if properties["description"] != nil {

		var_3eabb5992e3d := properties["description"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_3eabb5992e3d)

		if err != nil {
			panic(err)
		}

		var_3eabb5992e3d_mapped := new(string)
		*var_3eabb5992e3d_mapped = val.(string)

		s.Description = var_3eabb5992e3d_mapped
	}
	if properties["details"] != nil {

		var_f23f8157f030 := properties["details"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_OBJECT).UnPack(var_f23f8157f030)

		if err != nil {
			panic(err)
		}

		var_f23f8157f030_mapped := new(unstructured.Unstructured)
		*var_f23f8157f030_mapped = val.(unstructured.Unstructured)

		s.Details = var_f23f8157f030_mapped
	}
	if properties["securityConstraints"] != nil {

		var_dbed584ceab0 := properties["securityConstraints"]
		var_dbed584ceab0_mapped := []*SecurityConstraint{}
		for _, v := range var_dbed584ceab0.GetListValue().Values {

			var_ff5294a03a34 := v
			var_ff5294a03a34_mapped := SecurityConstraintMapperInstance.FromProperties(var_ff5294a03a34.GetStructValue().Fields)

			var_dbed584ceab0_mapped = append(var_dbed584ceab0_mapped, var_ff5294a03a34_mapped)
		}

		s.SecurityConstraints = var_dbed584ceab0_mapped
	}
	return s
}
