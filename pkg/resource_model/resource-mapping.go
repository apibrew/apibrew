package resource_model

import (
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/types"
	"google.golang.org/protobuf/types/known/structpb"
)

import "github.com/google/uuid"
import "time"
import "github.com/apibrew/apibrew/pkg/formats/unstructured"

type ResourceMapper struct {
}

func NewResourceMapper() *ResourceMapper {
	return &ResourceMapper{}
}

var ResourceMapperInstance = NewResourceMapper()

func (m *ResourceMapper) New() *Resource {
	return &Resource{}
}

func (m *ResourceMapper) ToRecord(resource *Resource) *model.Record {
	var rec = &model.Record{}
	rec.Properties = m.ToProperties(resource)

	if resource.Id != nil {
		rec.Id = resource.Id.String()
	}

	return rec
}

func (m *ResourceMapper) FromRecord(record *model.Record) *Resource {
	return m.FromProperties(record.Properties)
}

func (m *ResourceMapper) ToProperties(resource *Resource) map[string]*structpb.Value {
	var properties = make(map[string]*structpb.Value)

	var_93a28a126e9e := resource.Id

	if var_93a28a126e9e != nil {
		var var_93a28a126e9e_mapped *structpb.Value

		var var_93a28a126e9e_err error
		var_93a28a126e9e_mapped, var_93a28a126e9e_err = types.ByResourcePropertyType(model.ResourceProperty_UUID).Pack(*var_93a28a126e9e)
		if var_93a28a126e9e_err != nil {
			panic(var_93a28a126e9e_err)
		}
		properties["id"] = var_93a28a126e9e_mapped
	}

	var_979f685e7d5e := resource.Version

	var var_979f685e7d5e_mapped *structpb.Value

	var var_979f685e7d5e_err error
	var_979f685e7d5e_mapped, var_979f685e7d5e_err = types.ByResourcePropertyType(model.ResourceProperty_INT32).Pack(var_979f685e7d5e)
	if var_979f685e7d5e_err != nil {
		panic(var_979f685e7d5e_err)
	}
	properties["version"] = var_979f685e7d5e_mapped

	var_49f5a517b6c5 := resource.CreatedBy

	if var_49f5a517b6c5 != nil {
		var var_49f5a517b6c5_mapped *structpb.Value

		var var_49f5a517b6c5_err error
		var_49f5a517b6c5_mapped, var_49f5a517b6c5_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_49f5a517b6c5)
		if var_49f5a517b6c5_err != nil {
			panic(var_49f5a517b6c5_err)
		}
		properties["createdBy"] = var_49f5a517b6c5_mapped
	}

	var_0411f3e0a5f1 := resource.UpdatedBy

	if var_0411f3e0a5f1 != nil {
		var var_0411f3e0a5f1_mapped *structpb.Value

		var var_0411f3e0a5f1_err error
		var_0411f3e0a5f1_mapped, var_0411f3e0a5f1_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_0411f3e0a5f1)
		if var_0411f3e0a5f1_err != nil {
			panic(var_0411f3e0a5f1_err)
		}
		properties["updatedBy"] = var_0411f3e0a5f1_mapped
	}

	var_b3a668822b0d := resource.CreatedOn

	if var_b3a668822b0d != nil {
		var var_b3a668822b0d_mapped *structpb.Value

		var var_b3a668822b0d_err error
		var_b3a668822b0d_mapped, var_b3a668822b0d_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_b3a668822b0d)
		if var_b3a668822b0d_err != nil {
			panic(var_b3a668822b0d_err)
		}
		properties["createdOn"] = var_b3a668822b0d_mapped
	}

	var_87603a34b352 := resource.UpdatedOn

	if var_87603a34b352 != nil {
		var var_87603a34b352_mapped *structpb.Value

		var var_87603a34b352_err error
		var_87603a34b352_mapped, var_87603a34b352_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_87603a34b352)
		if var_87603a34b352_err != nil {
			panic(var_87603a34b352_err)
		}
		properties["updatedOn"] = var_87603a34b352_mapped
	}

	var_5f46d3ac3adf := resource.Name

	var var_5f46d3ac3adf_mapped *structpb.Value

	var var_5f46d3ac3adf_err error
	var_5f46d3ac3adf_mapped, var_5f46d3ac3adf_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_5f46d3ac3adf)
	if var_5f46d3ac3adf_err != nil {
		panic(var_5f46d3ac3adf_err)
	}
	properties["name"] = var_5f46d3ac3adf_mapped

	var_1bd6fae748fb := resource.Namespace

	if var_1bd6fae748fb != nil {
		var var_1bd6fae748fb_mapped *structpb.Value

		var_1bd6fae748fb_mapped = structpb.NewStructValue(&structpb.Struct{Fields: NamespaceMapperInstance.ToProperties(var_1bd6fae748fb)})
		properties["namespace"] = var_1bd6fae748fb_mapped
	}

	var_f39fc7e01233 := resource.Virtual

	var var_f39fc7e01233_mapped *structpb.Value

	var var_f39fc7e01233_err error
	var_f39fc7e01233_mapped, var_f39fc7e01233_err = types.ByResourcePropertyType(model.ResourceProperty_BOOL).Pack(var_f39fc7e01233)
	if var_f39fc7e01233_err != nil {
		panic(var_f39fc7e01233_err)
	}
	properties["virtual"] = var_f39fc7e01233_mapped

	var_37f3e92216c8 := resource.Properties

	var var_37f3e92216c8_mapped *structpb.Value

	var var_37f3e92216c8_l []*structpb.Value
	for _, value := range var_37f3e92216c8 {

		var_46f1df943cc6 := value
		var var_46f1df943cc6_mapped *structpb.Value

		var_46f1df943cc6_mapped = structpb.NewStructValue(&structpb.Struct{Fields: ResourcePropertyMapperInstance.ToProperties(&var_46f1df943cc6)})

		var_37f3e92216c8_l = append(var_37f3e92216c8_l, var_46f1df943cc6_mapped)
	}
	var_37f3e92216c8_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_37f3e92216c8_l})
	properties["properties"] = var_37f3e92216c8_mapped

	var_787b2dfe6b4c := resource.Indexes

	if var_787b2dfe6b4c != nil {
		var var_787b2dfe6b4c_mapped *structpb.Value

		var var_787b2dfe6b4c_l []*structpb.Value
		for _, value := range var_787b2dfe6b4c {

			var_8641a8d15efe := value
			var var_8641a8d15efe_mapped *structpb.Value

			var_8641a8d15efe_mapped = structpb.NewStructValue(&structpb.Struct{Fields: ResourceIndexMapperInstance.ToProperties(&var_8641a8d15efe)})

			var_787b2dfe6b4c_l = append(var_787b2dfe6b4c_l, var_8641a8d15efe_mapped)
		}
		var_787b2dfe6b4c_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_787b2dfe6b4c_l})
		properties["indexes"] = var_787b2dfe6b4c_mapped
	}

	var_e72a1323eda5 := resource.Types

	if var_e72a1323eda5 != nil {
		var var_e72a1323eda5_mapped *structpb.Value

		var var_e72a1323eda5_l []*structpb.Value
		for _, value := range var_e72a1323eda5 {

			var_eb1173220923 := value
			var var_eb1173220923_mapped *structpb.Value

			var_eb1173220923_mapped = structpb.NewStructValue(&structpb.Struct{Fields: ResourceSubTypeMapperInstance.ToProperties(&var_eb1173220923)})

			var_e72a1323eda5_l = append(var_e72a1323eda5_l, var_eb1173220923_mapped)
		}
		var_e72a1323eda5_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_e72a1323eda5_l})
		properties["types"] = var_e72a1323eda5_mapped
	}

	var_fc74e266401d := resource.Immutable

	var var_fc74e266401d_mapped *structpb.Value

	var var_fc74e266401d_err error
	var_fc74e266401d_mapped, var_fc74e266401d_err = types.ByResourcePropertyType(model.ResourceProperty_BOOL).Pack(var_fc74e266401d)
	if var_fc74e266401d_err != nil {
		panic(var_fc74e266401d_err)
	}
	properties["immutable"] = var_fc74e266401d_mapped

	var_8b335eb397e1 := resource.Abstract

	var var_8b335eb397e1_mapped *structpb.Value

	var var_8b335eb397e1_err error
	var_8b335eb397e1_mapped, var_8b335eb397e1_err = types.ByResourcePropertyType(model.ResourceProperty_BOOL).Pack(var_8b335eb397e1)
	if var_8b335eb397e1_err != nil {
		panic(var_8b335eb397e1_err)
	}
	properties["abstract"] = var_8b335eb397e1_mapped

	var_a374bc77fe54 := resource.DataSource

	if var_a374bc77fe54 != nil {
		var var_a374bc77fe54_mapped *structpb.Value

		var_a374bc77fe54_mapped = structpb.NewStructValue(&structpb.Struct{Fields: DataSourceMapperInstance.ToProperties(var_a374bc77fe54)})
		properties["dataSource"] = var_a374bc77fe54_mapped
	}

	var_51f7ea53307d := resource.Entity

	if var_51f7ea53307d != nil {
		var var_51f7ea53307d_mapped *structpb.Value

		var var_51f7ea53307d_err error
		var_51f7ea53307d_mapped, var_51f7ea53307d_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_51f7ea53307d)
		if var_51f7ea53307d_err != nil {
			panic(var_51f7ea53307d_err)
		}
		properties["entity"] = var_51f7ea53307d_mapped
	}

	var_510ad408c48a := resource.Catalog

	if var_510ad408c48a != nil {
		var var_510ad408c48a_mapped *structpb.Value

		var var_510ad408c48a_err error
		var_510ad408c48a_mapped, var_510ad408c48a_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_510ad408c48a)
		if var_510ad408c48a_err != nil {
			panic(var_510ad408c48a_err)
		}
		properties["catalog"] = var_510ad408c48a_mapped
	}

	var_864dce4e694a := resource.Title

	if var_864dce4e694a != nil {
		var var_864dce4e694a_mapped *structpb.Value

		var var_864dce4e694a_err error
		var_864dce4e694a_mapped, var_864dce4e694a_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_864dce4e694a)
		if var_864dce4e694a_err != nil {
			panic(var_864dce4e694a_err)
		}
		properties["title"] = var_864dce4e694a_mapped
	}

	var_c0e04695f934 := resource.Description

	if var_c0e04695f934 != nil {
		var var_c0e04695f934_mapped *structpb.Value

		var var_c0e04695f934_err error
		var_c0e04695f934_mapped, var_c0e04695f934_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_c0e04695f934)
		if var_c0e04695f934_err != nil {
			panic(var_c0e04695f934_err)
		}
		properties["description"] = var_c0e04695f934_mapped
	}

	var_6ad53eb7c1f1 := resource.Annotations

	if var_6ad53eb7c1f1 != nil {
		var var_6ad53eb7c1f1_mapped *structpb.Value

		var var_6ad53eb7c1f1_st *structpb.Struct = new(structpb.Struct)
		var_6ad53eb7c1f1_st.Fields = make(map[string]*structpb.Value)
		for key, value := range var_6ad53eb7c1f1 {

			var_d4ebc926bd75 := value
			var var_d4ebc926bd75_mapped *structpb.Value

			var var_d4ebc926bd75_err error
			var_d4ebc926bd75_mapped, var_d4ebc926bd75_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_d4ebc926bd75)
			if var_d4ebc926bd75_err != nil {
				panic(var_d4ebc926bd75_err)
			}

			var_6ad53eb7c1f1_st.Fields[key] = var_d4ebc926bd75_mapped
		}
		var_6ad53eb7c1f1_mapped = structpb.NewStructValue(var_6ad53eb7c1f1_st)
		properties["annotations"] = var_6ad53eb7c1f1_mapped
	}
	return properties
}

func (m *ResourceMapper) FromProperties(properties map[string]*structpb.Value) *Resource {
	var s = m.New()
	if properties["id"] != nil {

		var_b667c43d9b77 := properties["id"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_UUID).UnPack(var_b667c43d9b77)

		if err != nil {
			panic(err)
		}

		var_b667c43d9b77_mapped := new(uuid.UUID)
		*var_b667c43d9b77_mapped = val.(uuid.UUID)

		s.Id = var_b667c43d9b77_mapped
	}
	if properties["version"] != nil {

		var_5a7a2f798209 := properties["version"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(var_5a7a2f798209)

		if err != nil {
			panic(err)
		}

		var_5a7a2f798209_mapped := val.(int32)

		s.Version = var_5a7a2f798209_mapped
	}
	if properties["createdBy"] != nil {

		var_c20ef2038efd := properties["createdBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_c20ef2038efd)

		if err != nil {
			panic(err)
		}

		var_c20ef2038efd_mapped := new(string)
		*var_c20ef2038efd_mapped = val.(string)

		s.CreatedBy = var_c20ef2038efd_mapped
	}
	if properties["updatedBy"] != nil {

		var_8ca115c81d33 := properties["updatedBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_8ca115c81d33)

		if err != nil {
			panic(err)
		}

		var_8ca115c81d33_mapped := new(string)
		*var_8ca115c81d33_mapped = val.(string)

		s.UpdatedBy = var_8ca115c81d33_mapped
	}
	if properties["createdOn"] != nil {

		var_7ca89ed9a2a4 := properties["createdOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_7ca89ed9a2a4)

		if err != nil {
			panic(err)
		}

		var_7ca89ed9a2a4_mapped := new(time.Time)
		*var_7ca89ed9a2a4_mapped = val.(time.Time)

		s.CreatedOn = var_7ca89ed9a2a4_mapped
	}
	if properties["updatedOn"] != nil {

		var_f73b6dbd6bce := properties["updatedOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_f73b6dbd6bce)

		if err != nil {
			panic(err)
		}

		var_f73b6dbd6bce_mapped := new(time.Time)
		*var_f73b6dbd6bce_mapped = val.(time.Time)

		s.UpdatedOn = var_f73b6dbd6bce_mapped
	}
	if properties["name"] != nil {

		var_67ddd3628a56 := properties["name"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_67ddd3628a56)

		if err != nil {
			panic(err)
		}

		var_67ddd3628a56_mapped := val.(string)

		s.Name = var_67ddd3628a56_mapped
	}
	if properties["namespace"] != nil {

		var_495fe8fcd076 := properties["namespace"]
		var_495fe8fcd076_mapped := NamespaceMapperInstance.FromProperties(var_495fe8fcd076.GetStructValue().Fields)

		s.Namespace = var_495fe8fcd076_mapped
	}
	if properties["virtual"] != nil {

		var_d19e4c3eaa8c := properties["virtual"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_BOOL).UnPack(var_d19e4c3eaa8c)

		if err != nil {
			panic(err)
		}

		var_d19e4c3eaa8c_mapped := val.(bool)

		s.Virtual = var_d19e4c3eaa8c_mapped
	}
	if properties["properties"] != nil {

		var_1bfc22048d82 := properties["properties"]
		var_1bfc22048d82_mapped := []ResourceProperty{}
		for _, v := range var_1bfc22048d82.GetListValue().Values {

			var_19a857887bcd := v
			var mappedValue = ResourcePropertyMapperInstance.FromProperties(var_19a857887bcd.GetStructValue().Fields)

			var_19a857887bcd_mapped := *mappedValue

			var_1bfc22048d82_mapped = append(var_1bfc22048d82_mapped, var_19a857887bcd_mapped)
		}

		s.Properties = var_1bfc22048d82_mapped
	}
	if properties["indexes"] != nil {

		var_9400d47e31d9 := properties["indexes"]
		var_9400d47e31d9_mapped := []ResourceIndex{}
		for _, v := range var_9400d47e31d9.GetListValue().Values {

			var_027be610edc3 := v
			var mappedValue = ResourceIndexMapperInstance.FromProperties(var_027be610edc3.GetStructValue().Fields)

			var_027be610edc3_mapped := *mappedValue

			var_9400d47e31d9_mapped = append(var_9400d47e31d9_mapped, var_027be610edc3_mapped)
		}

		s.Indexes = var_9400d47e31d9_mapped
	}
	if properties["types"] != nil {

		var_492f82672557 := properties["types"]
		var_492f82672557_mapped := []ResourceSubType{}
		for _, v := range var_492f82672557.GetListValue().Values {

			var_85868c043fe4 := v
			var mappedValue = ResourceSubTypeMapperInstance.FromProperties(var_85868c043fe4.GetStructValue().Fields)

			var_85868c043fe4_mapped := *mappedValue

			var_492f82672557_mapped = append(var_492f82672557_mapped, var_85868c043fe4_mapped)
		}

		s.Types = var_492f82672557_mapped
	}
	if properties["immutable"] != nil {

		var_67b3c5bf1c8a := properties["immutable"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_BOOL).UnPack(var_67b3c5bf1c8a)

		if err != nil {
			panic(err)
		}

		var_67b3c5bf1c8a_mapped := val.(bool)

		s.Immutable = var_67b3c5bf1c8a_mapped
	}
	if properties["abstract"] != nil {

		var_aea82d7eaa3f := properties["abstract"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_BOOL).UnPack(var_aea82d7eaa3f)

		if err != nil {
			panic(err)
		}

		var_aea82d7eaa3f_mapped := val.(bool)

		s.Abstract = var_aea82d7eaa3f_mapped
	}
	if properties["dataSource"] != nil {

		var_ef128cafba20 := properties["dataSource"]
		var_ef128cafba20_mapped := DataSourceMapperInstance.FromProperties(var_ef128cafba20.GetStructValue().Fields)

		s.DataSource = var_ef128cafba20_mapped
	}
	if properties["entity"] != nil {

		var_cf8d0771a9f7 := properties["entity"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_cf8d0771a9f7)

		if err != nil {
			panic(err)
		}

		var_cf8d0771a9f7_mapped := new(string)
		*var_cf8d0771a9f7_mapped = val.(string)

		s.Entity = var_cf8d0771a9f7_mapped
	}
	if properties["catalog"] != nil {

		var_c9b4bd8e5af2 := properties["catalog"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_c9b4bd8e5af2)

		if err != nil {
			panic(err)
		}

		var_c9b4bd8e5af2_mapped := new(string)
		*var_c9b4bd8e5af2_mapped = val.(string)

		s.Catalog = var_c9b4bd8e5af2_mapped
	}
	if properties["title"] != nil {

		var_c40347da7bb0 := properties["title"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_c40347da7bb0)

		if err != nil {
			panic(err)
		}

		var_c40347da7bb0_mapped := new(string)
		*var_c40347da7bb0_mapped = val.(string)

		s.Title = var_c40347da7bb0_mapped
	}
	if properties["description"] != nil {

		var_c2d55b11c9ec := properties["description"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_c2d55b11c9ec)

		if err != nil {
			panic(err)
		}

		var_c2d55b11c9ec_mapped := new(string)
		*var_c2d55b11c9ec_mapped = val.(string)

		s.Description = var_c2d55b11c9ec_mapped
	}
	if properties["annotations"] != nil {

		var_a1691190c164 := properties["annotations"]
		var_a1691190c164_mapped := make(map[string]string)
		for k, v := range var_a1691190c164.GetStructValue().Fields {

			var_5ef7789c553c := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_5ef7789c553c)

			if err != nil {
				panic(err)
			}

			var_5ef7789c553c_mapped := val.(string)

			var_a1691190c164_mapped[k] = var_5ef7789c553c_mapped
		}

		s.Annotations = var_a1691190c164_mapped
	}
	return s
}

type ResourcePropertyMapper struct {
}

func NewResourcePropertyMapper() *ResourcePropertyMapper {
	return &ResourcePropertyMapper{}
}

var ResourcePropertyMapperInstance = NewResourcePropertyMapper()

func (m *ResourcePropertyMapper) New() *ResourceProperty {
	return &ResourceProperty{}
}

func (m *ResourcePropertyMapper) ToProperties(resourceProperty *ResourceProperty) map[string]*structpb.Value {
	var properties = make(map[string]*structpb.Value)

	var_ee3759653b97 := resourceProperty.Name

	var var_ee3759653b97_mapped *structpb.Value

	var var_ee3759653b97_err error
	var_ee3759653b97_mapped, var_ee3759653b97_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_ee3759653b97)
	if var_ee3759653b97_err != nil {
		panic(var_ee3759653b97_err)
	}
	properties["name"] = var_ee3759653b97_mapped

	var_e85962a68446 := resourceProperty.Type

	var var_e85962a68446_mapped *structpb.Value

	var var_e85962a68446_err error
	var_e85962a68446_mapped, var_e85962a68446_err = types.ByResourcePropertyType(model.ResourceProperty_INT32).Pack(var_e85962a68446)
	if var_e85962a68446_err != nil {
		panic(var_e85962a68446_err)
	}
	properties["type"] = var_e85962a68446_mapped

	var_51fe325e2531 := resourceProperty.TypeRef

	if var_51fe325e2531 != nil {
		var var_51fe325e2531_mapped *structpb.Value

		var var_51fe325e2531_err error
		var_51fe325e2531_mapped, var_51fe325e2531_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_51fe325e2531)
		if var_51fe325e2531_err != nil {
			panic(var_51fe325e2531_err)
		}
		properties["typeRef"] = var_51fe325e2531_mapped
	}

	var_9214e7a1c7c4 := resourceProperty.Mapping

	var var_9214e7a1c7c4_mapped *structpb.Value

	var var_9214e7a1c7c4_err error
	var_9214e7a1c7c4_mapped, var_9214e7a1c7c4_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_9214e7a1c7c4)
	if var_9214e7a1c7c4_err != nil {
		panic(var_9214e7a1c7c4_err)
	}
	properties["mapping"] = var_9214e7a1c7c4_mapped

	var_98a7187e9afd := resourceProperty.Primary

	var var_98a7187e9afd_mapped *structpb.Value

	var var_98a7187e9afd_err error
	var_98a7187e9afd_mapped, var_98a7187e9afd_err = types.ByResourcePropertyType(model.ResourceProperty_BOOL).Pack(var_98a7187e9afd)
	if var_98a7187e9afd_err != nil {
		panic(var_98a7187e9afd_err)
	}
	properties["primary"] = var_98a7187e9afd_mapped

	var_57723eea2aab := resourceProperty.Required

	var var_57723eea2aab_mapped *structpb.Value

	var var_57723eea2aab_err error
	var_57723eea2aab_mapped, var_57723eea2aab_err = types.ByResourcePropertyType(model.ResourceProperty_BOOL).Pack(var_57723eea2aab)
	if var_57723eea2aab_err != nil {
		panic(var_57723eea2aab_err)
	}
	properties["required"] = var_57723eea2aab_mapped

	var_f920d3aff278 := resourceProperty.Unique

	var var_f920d3aff278_mapped *structpb.Value

	var var_f920d3aff278_err error
	var_f920d3aff278_mapped, var_f920d3aff278_err = types.ByResourcePropertyType(model.ResourceProperty_BOOL).Pack(var_f920d3aff278)
	if var_f920d3aff278_err != nil {
		panic(var_f920d3aff278_err)
	}
	properties["unique"] = var_f920d3aff278_mapped

	var_7b914b71ca0c := resourceProperty.Immutable

	var var_7b914b71ca0c_mapped *structpb.Value

	var var_7b914b71ca0c_err error
	var_7b914b71ca0c_mapped, var_7b914b71ca0c_err = types.ByResourcePropertyType(model.ResourceProperty_BOOL).Pack(var_7b914b71ca0c)
	if var_7b914b71ca0c_err != nil {
		panic(var_7b914b71ca0c_err)
	}
	properties["immutable"] = var_7b914b71ca0c_mapped

	var_49f5089df92b := resourceProperty.Length

	var var_49f5089df92b_mapped *structpb.Value

	var var_49f5089df92b_err error
	var_49f5089df92b_mapped, var_49f5089df92b_err = types.ByResourcePropertyType(model.ResourceProperty_INT32).Pack(var_49f5089df92b)
	if var_49f5089df92b_err != nil {
		panic(var_49f5089df92b_err)
	}
	properties["length"] = var_49f5089df92b_mapped

	var_a69ae7825e01 := resourceProperty.Item

	if var_a69ae7825e01 != nil {
		var var_a69ae7825e01_mapped *structpb.Value

		var_a69ae7825e01_mapped = structpb.NewStructValue(&structpb.Struct{Fields: ResourcePropertyMapperInstance.ToProperties(var_a69ae7825e01)})
		properties["item"] = var_a69ae7825e01_mapped
	}

	var_90abd77516ae := resourceProperty.Properties

	var var_90abd77516ae_mapped *structpb.Value

	var var_90abd77516ae_l []*structpb.Value
	for _, value := range var_90abd77516ae {

		var_70330b5900bc := value
		var var_70330b5900bc_mapped *structpb.Value

		var_70330b5900bc_mapped = structpb.NewStructValue(&structpb.Struct{Fields: ResourcePropertyMapperInstance.ToProperties(&var_70330b5900bc)})

		var_90abd77516ae_l = append(var_90abd77516ae_l, var_70330b5900bc_mapped)
	}
	var_90abd77516ae_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_90abd77516ae_l})
	properties["properties"] = var_90abd77516ae_mapped

	var_167385fcdf19 := resourceProperty.Reference

	if var_167385fcdf19 != nil {
		var var_167385fcdf19_mapped *structpb.Value

		var_167385fcdf19_mapped = structpb.NewStructValue(&structpb.Struct{Fields: ResourceReferenceMapperInstance.ToProperties(var_167385fcdf19)})
		properties["reference"] = var_167385fcdf19_mapped
	}

	var_2ffaa6cf9e62 := resourceProperty.DefaultValue

	if var_2ffaa6cf9e62 != nil {
		var var_2ffaa6cf9e62_mapped *structpb.Value

		var var_2ffaa6cf9e62_err error
		var_2ffaa6cf9e62_mapped, var_2ffaa6cf9e62_err = types.ByResourcePropertyType(model.ResourceProperty_OBJECT).Pack(*var_2ffaa6cf9e62)
		if var_2ffaa6cf9e62_err != nil {
			panic(var_2ffaa6cf9e62_err)
		}
		properties["defaultValue"] = var_2ffaa6cf9e62_mapped
	}

	var_eb111d81ec8b := resourceProperty.EnumValues

	if var_eb111d81ec8b != nil {
		var var_eb111d81ec8b_mapped *structpb.Value

		var var_eb111d81ec8b_l []*structpb.Value
		for _, value := range var_eb111d81ec8b {

			var_3cab3208f63a := value
			var var_3cab3208f63a_mapped *structpb.Value

			var var_3cab3208f63a_err error
			var_3cab3208f63a_mapped, var_3cab3208f63a_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_3cab3208f63a)
			if var_3cab3208f63a_err != nil {
				panic(var_3cab3208f63a_err)
			}

			var_eb111d81ec8b_l = append(var_eb111d81ec8b_l, var_3cab3208f63a_mapped)
		}
		var_eb111d81ec8b_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_eb111d81ec8b_l})
		properties["enumValues"] = var_eb111d81ec8b_mapped
	}

	var_f812dd78034b := resourceProperty.ExampleValue

	if var_f812dd78034b != nil {
		var var_f812dd78034b_mapped *structpb.Value

		var var_f812dd78034b_err error
		var_f812dd78034b_mapped, var_f812dd78034b_err = types.ByResourcePropertyType(model.ResourceProperty_OBJECT).Pack(*var_f812dd78034b)
		if var_f812dd78034b_err != nil {
			panic(var_f812dd78034b_err)
		}
		properties["exampleValue"] = var_f812dd78034b_mapped
	}

	var_0578bf6ab02b := resourceProperty.Title

	if var_0578bf6ab02b != nil {
		var var_0578bf6ab02b_mapped *structpb.Value

		var var_0578bf6ab02b_err error
		var_0578bf6ab02b_mapped, var_0578bf6ab02b_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_0578bf6ab02b)
		if var_0578bf6ab02b_err != nil {
			panic(var_0578bf6ab02b_err)
		}
		properties["title"] = var_0578bf6ab02b_mapped
	}

	var_b01dcd686759 := resourceProperty.Description

	if var_b01dcd686759 != nil {
		var var_b01dcd686759_mapped *structpb.Value

		var var_b01dcd686759_err error
		var_b01dcd686759_mapped, var_b01dcd686759_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_b01dcd686759)
		if var_b01dcd686759_err != nil {
			panic(var_b01dcd686759_err)
		}
		properties["description"] = var_b01dcd686759_mapped
	}

	var_5691aba48dc1 := resourceProperty.Annotations

	if var_5691aba48dc1 != nil {
		var var_5691aba48dc1_mapped *structpb.Value

		var var_5691aba48dc1_st *structpb.Struct = new(structpb.Struct)
		var_5691aba48dc1_st.Fields = make(map[string]*structpb.Value)
		for key, value := range var_5691aba48dc1 {

			var_4679aafc25ae := value
			var var_4679aafc25ae_mapped *structpb.Value

			var var_4679aafc25ae_err error
			var_4679aafc25ae_mapped, var_4679aafc25ae_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_4679aafc25ae)
			if var_4679aafc25ae_err != nil {
				panic(var_4679aafc25ae_err)
			}

			var_5691aba48dc1_st.Fields[key] = var_4679aafc25ae_mapped
		}
		var_5691aba48dc1_mapped = structpb.NewStructValue(var_5691aba48dc1_st)
		properties["annotations"] = var_5691aba48dc1_mapped
	}
	return properties
}

func (m *ResourcePropertyMapper) FromProperties(properties map[string]*structpb.Value) *ResourceProperty {
	var s = m.New()
	if properties["name"] != nil {

		var_4772f1be3813 := properties["name"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_4772f1be3813)

		if err != nil {
			panic(err)
		}

		var_4772f1be3813_mapped := val.(string)

		s.Name = var_4772f1be3813_mapped
	}
	if properties["type"] != nil {

		var_593710c350bf := properties["type"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(var_593710c350bf)

		if err != nil {
			panic(err)
		}

		var_593710c350bf_mapped := val.(int32)

		s.Type = var_593710c350bf_mapped
	}
	if properties["typeRef"] != nil {

		var_80cdeb472e9a := properties["typeRef"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_80cdeb472e9a)

		if err != nil {
			panic(err)
		}

		var_80cdeb472e9a_mapped := new(string)
		*var_80cdeb472e9a_mapped = val.(string)

		s.TypeRef = var_80cdeb472e9a_mapped
	}
	if properties["mapping"] != nil {

		var_26894ee461b7 := properties["mapping"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_26894ee461b7)

		if err != nil {
			panic(err)
		}

		var_26894ee461b7_mapped := val.(string)

		s.Mapping = var_26894ee461b7_mapped
	}
	if properties["primary"] != nil {

		var_0a10d6329d8d := properties["primary"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_BOOL).UnPack(var_0a10d6329d8d)

		if err != nil {
			panic(err)
		}

		var_0a10d6329d8d_mapped := val.(bool)

		s.Primary = var_0a10d6329d8d_mapped
	}
	if properties["required"] != nil {

		var_f9c3db1f68cf := properties["required"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_BOOL).UnPack(var_f9c3db1f68cf)

		if err != nil {
			panic(err)
		}

		var_f9c3db1f68cf_mapped := val.(bool)

		s.Required = var_f9c3db1f68cf_mapped
	}
	if properties["unique"] != nil {

		var_e0e73400835d := properties["unique"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_BOOL).UnPack(var_e0e73400835d)

		if err != nil {
			panic(err)
		}

		var_e0e73400835d_mapped := val.(bool)

		s.Unique = var_e0e73400835d_mapped
	}
	if properties["immutable"] != nil {

		var_75dc837ff165 := properties["immutable"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_BOOL).UnPack(var_75dc837ff165)

		if err != nil {
			panic(err)
		}

		var_75dc837ff165_mapped := val.(bool)

		s.Immutable = var_75dc837ff165_mapped
	}
	if properties["length"] != nil {

		var_2aa18d378a4e := properties["length"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(var_2aa18d378a4e)

		if err != nil {
			panic(err)
		}

		var_2aa18d378a4e_mapped := val.(int32)

		s.Length = var_2aa18d378a4e_mapped
	}
	if properties["item"] != nil {

		var_af6cafeba172 := properties["item"]
		var mappedValue = ResourcePropertyMapperInstance.FromProperties(var_af6cafeba172.GetStructValue().Fields)

		var_af6cafeba172_mapped := mappedValue

		s.Item = var_af6cafeba172_mapped
	}
	if properties["properties"] != nil {

		var_70735acd3eac := properties["properties"]
		var_70735acd3eac_mapped := []ResourceProperty{}
		for _, v := range var_70735acd3eac.GetListValue().Values {

			var_b3ac29e8beb0 := v
			var mappedValue = ResourcePropertyMapperInstance.FromProperties(var_b3ac29e8beb0.GetStructValue().Fields)

			var_b3ac29e8beb0_mapped := *mappedValue

			var_70735acd3eac_mapped = append(var_70735acd3eac_mapped, var_b3ac29e8beb0_mapped)
		}

		s.Properties = var_70735acd3eac_mapped
	}
	if properties["reference"] != nil {

		var_9429a5a62c03 := properties["reference"]
		var mappedValue = ResourceReferenceMapperInstance.FromProperties(var_9429a5a62c03.GetStructValue().Fields)

		var_9429a5a62c03_mapped := mappedValue

		s.Reference = var_9429a5a62c03_mapped
	}
	if properties["defaultValue"] != nil {

		var_2ac39f8b9dc2 := properties["defaultValue"]
		var_2ac39f8b9dc2_mapped := new(unstructured.Unstructured)
		*var_2ac39f8b9dc2_mapped = unstructured.FromStructValue(var_2ac39f8b9dc2.GetStructValue())

		s.DefaultValue = var_2ac39f8b9dc2_mapped
	}
	if properties["enumValues"] != nil {

		var_35691aebf0e5 := properties["enumValues"]
		var_35691aebf0e5_mapped := []string{}
		for _, v := range var_35691aebf0e5.GetListValue().Values {

			var_a9442b8bf771 := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_a9442b8bf771)

			if err != nil {
				panic(err)
			}

			var_a9442b8bf771_mapped := val.(string)

			var_35691aebf0e5_mapped = append(var_35691aebf0e5_mapped, var_a9442b8bf771_mapped)
		}

		s.EnumValues = var_35691aebf0e5_mapped
	}
	if properties["exampleValue"] != nil {

		var_b7f5ea28ae80 := properties["exampleValue"]
		var_b7f5ea28ae80_mapped := new(unstructured.Unstructured)
		*var_b7f5ea28ae80_mapped = unstructured.FromStructValue(var_b7f5ea28ae80.GetStructValue())

		s.ExampleValue = var_b7f5ea28ae80_mapped
	}
	if properties["title"] != nil {

		var_c4d475198975 := properties["title"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_c4d475198975)

		if err != nil {
			panic(err)
		}

		var_c4d475198975_mapped := new(string)
		*var_c4d475198975_mapped = val.(string)

		s.Title = var_c4d475198975_mapped
	}
	if properties["description"] != nil {

		var_6215fea612b1 := properties["description"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_6215fea612b1)

		if err != nil {
			panic(err)
		}

		var_6215fea612b1_mapped := new(string)
		*var_6215fea612b1_mapped = val.(string)

		s.Description = var_6215fea612b1_mapped
	}
	if properties["annotations"] != nil {

		var_6971c59c29ea := properties["annotations"]
		var_6971c59c29ea_mapped := make(map[string]string)
		for k, v := range var_6971c59c29ea.GetStructValue().Fields {

			var_4ef700577eea := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_4ef700577eea)

			if err != nil {
				panic(err)
			}

			var_4ef700577eea_mapped := val.(string)

			var_6971c59c29ea_mapped[k] = var_4ef700577eea_mapped
		}

		s.Annotations = var_6971c59c29ea_mapped
	}
	return s
}

type ResourceSubTypeMapper struct {
}

func NewResourceSubTypeMapper() *ResourceSubTypeMapper {
	return &ResourceSubTypeMapper{}
}

var ResourceSubTypeMapperInstance = NewResourceSubTypeMapper()

func (m *ResourceSubTypeMapper) New() *ResourceSubType {
	return &ResourceSubType{}
}

func (m *ResourceSubTypeMapper) ToProperties(resourceSubType *ResourceSubType) map[string]*structpb.Value {
	var properties = make(map[string]*structpb.Value)

	var_f8b0297a6dc8 := resourceSubType.Name

	var var_f8b0297a6dc8_mapped *structpb.Value

	var var_f8b0297a6dc8_err error
	var_f8b0297a6dc8_mapped, var_f8b0297a6dc8_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_f8b0297a6dc8)
	if var_f8b0297a6dc8_err != nil {
		panic(var_f8b0297a6dc8_err)
	}
	properties["name"] = var_f8b0297a6dc8_mapped

	var_a7b852a05c70 := resourceSubType.Properties

	var var_a7b852a05c70_mapped *structpb.Value

	var var_a7b852a05c70_l []*structpb.Value
	for _, value := range var_a7b852a05c70 {

		var_6e1fdf222379 := value
		var var_6e1fdf222379_mapped *structpb.Value

		var_6e1fdf222379_mapped = structpb.NewStructValue(&structpb.Struct{Fields: ResourcePropertyMapperInstance.ToProperties(&var_6e1fdf222379)})

		var_a7b852a05c70_l = append(var_a7b852a05c70_l, var_6e1fdf222379_mapped)
	}
	var_a7b852a05c70_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_a7b852a05c70_l})
	properties["properties"] = var_a7b852a05c70_mapped
	return properties
}

func (m *ResourceSubTypeMapper) FromProperties(properties map[string]*structpb.Value) *ResourceSubType {
	var s = m.New()
	if properties["name"] != nil {

		var_14fdb03c7078 := properties["name"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_14fdb03c7078)

		if err != nil {
			panic(err)
		}

		var_14fdb03c7078_mapped := val.(string)

		s.Name = var_14fdb03c7078_mapped
	}
	if properties["properties"] != nil {

		var_6b866ecf73c6 := properties["properties"]
		var_6b866ecf73c6_mapped := []ResourceProperty{}
		for _, v := range var_6b866ecf73c6.GetListValue().Values {

			var_fcb1a92db687 := v
			var mappedValue = ResourcePropertyMapperInstance.FromProperties(var_fcb1a92db687.GetStructValue().Fields)

			var_fcb1a92db687_mapped := *mappedValue

			var_6b866ecf73c6_mapped = append(var_6b866ecf73c6_mapped, var_fcb1a92db687_mapped)
		}

		s.Properties = var_6b866ecf73c6_mapped
	}
	return s
}

type ResourceIndexPropertyMapper struct {
}

func NewResourceIndexPropertyMapper() *ResourceIndexPropertyMapper {
	return &ResourceIndexPropertyMapper{}
}

var ResourceIndexPropertyMapperInstance = NewResourceIndexPropertyMapper()

func (m *ResourceIndexPropertyMapper) New() *ResourceIndexProperty {
	return &ResourceIndexProperty{}
}

func (m *ResourceIndexPropertyMapper) ToProperties(resourceIndexProperty *ResourceIndexProperty) map[string]*structpb.Value {
	var properties = make(map[string]*structpb.Value)

	var_728b6f8451c4 := resourceIndexProperty.Name

	var var_728b6f8451c4_mapped *structpb.Value

	var var_728b6f8451c4_err error
	var_728b6f8451c4_mapped, var_728b6f8451c4_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_728b6f8451c4)
	if var_728b6f8451c4_err != nil {
		panic(var_728b6f8451c4_err)
	}
	properties["name"] = var_728b6f8451c4_mapped

	var_824b6c783e69 := resourceIndexProperty.Order

	if var_824b6c783e69 != nil {
		var var_824b6c783e69_mapped *structpb.Value

		var var_824b6c783e69_err error
		var_824b6c783e69_mapped, var_824b6c783e69_err = types.ByResourcePropertyType(model.ResourceProperty_ENUM).Pack(string(*var_824b6c783e69))
		if var_824b6c783e69_err != nil {
			panic(var_824b6c783e69_err)
		}
		properties["order"] = var_824b6c783e69_mapped
	}
	return properties
}

func (m *ResourceIndexPropertyMapper) FromProperties(properties map[string]*structpb.Value) *ResourceIndexProperty {
	var s = m.New()
	if properties["name"] != nil {

		var_1d1f5712f96c := properties["name"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_1d1f5712f96c)

		if err != nil {
			panic(err)
		}

		var_1d1f5712f96c_mapped := val.(string)

		s.Name = var_1d1f5712f96c_mapped
	}
	if properties["order"] != nil {

		var_3b03099c663c := properties["order"]
		var_3b03099c663c_mapped := new(ResourceOrder)
		*var_3b03099c663c_mapped = (ResourceOrder)(var_3b03099c663c.GetStringValue())

		s.Order = var_3b03099c663c_mapped
	}
	return s
}

type ResourceIndexMapper struct {
}

func NewResourceIndexMapper() *ResourceIndexMapper {
	return &ResourceIndexMapper{}
}

var ResourceIndexMapperInstance = NewResourceIndexMapper()

func (m *ResourceIndexMapper) New() *ResourceIndex {
	return &ResourceIndex{}
}

func (m *ResourceIndexMapper) ToProperties(resourceIndex *ResourceIndex) map[string]*structpb.Value {
	var properties = make(map[string]*structpb.Value)

	var_4e3ec0074331 := resourceIndex.Properties

	if var_4e3ec0074331 != nil {
		var var_4e3ec0074331_mapped *structpb.Value

		var var_4e3ec0074331_l []*structpb.Value
		for _, value := range var_4e3ec0074331 {

			var_3d9ec3f6165d := value
			var var_3d9ec3f6165d_mapped *structpb.Value

			var_3d9ec3f6165d_mapped = structpb.NewStructValue(&structpb.Struct{Fields: ResourceIndexPropertyMapperInstance.ToProperties(&var_3d9ec3f6165d)})

			var_4e3ec0074331_l = append(var_4e3ec0074331_l, var_3d9ec3f6165d_mapped)
		}
		var_4e3ec0074331_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_4e3ec0074331_l})
		properties["properties"] = var_4e3ec0074331_mapped
	}

	var_d4ca855c5fca := resourceIndex.IndexType

	if var_d4ca855c5fca != nil {
		var var_d4ca855c5fca_mapped *structpb.Value

		var var_d4ca855c5fca_err error
		var_d4ca855c5fca_mapped, var_d4ca855c5fca_err = types.ByResourcePropertyType(model.ResourceProperty_ENUM).Pack(string(*var_d4ca855c5fca))
		if var_d4ca855c5fca_err != nil {
			panic(var_d4ca855c5fca_err)
		}
		properties["indexType"] = var_d4ca855c5fca_mapped
	}

	var_a6fb51e6089b := resourceIndex.Unique

	if var_a6fb51e6089b != nil {
		var var_a6fb51e6089b_mapped *structpb.Value

		var var_a6fb51e6089b_err error
		var_a6fb51e6089b_mapped, var_a6fb51e6089b_err = types.ByResourcePropertyType(model.ResourceProperty_BOOL).Pack(*var_a6fb51e6089b)
		if var_a6fb51e6089b_err != nil {
			panic(var_a6fb51e6089b_err)
		}
		properties["unique"] = var_a6fb51e6089b_mapped
	}

	var_e256ecab822d := resourceIndex.Annotations

	if var_e256ecab822d != nil {
		var var_e256ecab822d_mapped *structpb.Value

		var var_e256ecab822d_st *structpb.Struct = new(structpb.Struct)
		var_e256ecab822d_st.Fields = make(map[string]*structpb.Value)
		for key, value := range var_e256ecab822d {

			var_d1a4b0aaea02 := value
			var var_d1a4b0aaea02_mapped *structpb.Value

			var var_d1a4b0aaea02_err error
			var_d1a4b0aaea02_mapped, var_d1a4b0aaea02_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_d1a4b0aaea02)
			if var_d1a4b0aaea02_err != nil {
				panic(var_d1a4b0aaea02_err)
			}

			var_e256ecab822d_st.Fields[key] = var_d1a4b0aaea02_mapped
		}
		var_e256ecab822d_mapped = structpb.NewStructValue(var_e256ecab822d_st)
		properties["annotations"] = var_e256ecab822d_mapped
	}
	return properties
}

func (m *ResourceIndexMapper) FromProperties(properties map[string]*structpb.Value) *ResourceIndex {
	var s = m.New()
	if properties["properties"] != nil {

		var_b882ce6abb36 := properties["properties"]
		var_b882ce6abb36_mapped := []ResourceIndexProperty{}
		for _, v := range var_b882ce6abb36.GetListValue().Values {

			var_26364efd3c2e := v
			var mappedValue = ResourceIndexPropertyMapperInstance.FromProperties(var_26364efd3c2e.GetStructValue().Fields)

			var_26364efd3c2e_mapped := *mappedValue

			var_b882ce6abb36_mapped = append(var_b882ce6abb36_mapped, var_26364efd3c2e_mapped)
		}

		s.Properties = var_b882ce6abb36_mapped
	}
	if properties["indexType"] != nil {

		var_e2e2eaba7cb1 := properties["indexType"]
		var_e2e2eaba7cb1_mapped := new(ResourceIndexType)
		*var_e2e2eaba7cb1_mapped = (ResourceIndexType)(var_e2e2eaba7cb1.GetStringValue())

		s.IndexType = var_e2e2eaba7cb1_mapped
	}
	if properties["unique"] != nil {

		var_5640ce92bf09 := properties["unique"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_BOOL).UnPack(var_5640ce92bf09)

		if err != nil {
			panic(err)
		}

		var_5640ce92bf09_mapped := new(bool)
		*var_5640ce92bf09_mapped = val.(bool)

		s.Unique = var_5640ce92bf09_mapped
	}
	if properties["annotations"] != nil {

		var_38e1b2a15bb4 := properties["annotations"]
		var_38e1b2a15bb4_mapped := make(map[string]string)
		for k, v := range var_38e1b2a15bb4.GetStructValue().Fields {

			var_0f36a924e491 := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_0f36a924e491)

			if err != nil {
				panic(err)
			}

			var_0f36a924e491_mapped := val.(string)

			var_38e1b2a15bb4_mapped[k] = var_0f36a924e491_mapped
		}

		s.Annotations = var_38e1b2a15bb4_mapped
	}
	return s
}

type ResourceReferenceMapper struct {
}

func NewResourceReferenceMapper() *ResourceReferenceMapper {
	return &ResourceReferenceMapper{}
}

var ResourceReferenceMapperInstance = NewResourceReferenceMapper()

func (m *ResourceReferenceMapper) New() *ResourceReference {
	return &ResourceReference{}
}

func (m *ResourceReferenceMapper) ToProperties(resourceReference *ResourceReference) map[string]*structpb.Value {
	var properties = make(map[string]*structpb.Value)

	var_3314055a9e09 := resourceReference.Resource

	if var_3314055a9e09 != nil {
		var var_3314055a9e09_mapped *structpb.Value

		var_3314055a9e09_mapped = structpb.NewStructValue(&structpb.Struct{Fields: ResourceMapperInstance.ToProperties(var_3314055a9e09)})
		properties["resource"] = var_3314055a9e09_mapped
	}

	var_0f2bc4c42ad0 := resourceReference.Cascade

	if var_0f2bc4c42ad0 != nil {
		var var_0f2bc4c42ad0_mapped *structpb.Value

		var var_0f2bc4c42ad0_err error
		var_0f2bc4c42ad0_mapped, var_0f2bc4c42ad0_err = types.ByResourcePropertyType(model.ResourceProperty_BOOL).Pack(*var_0f2bc4c42ad0)
		if var_0f2bc4c42ad0_err != nil {
			panic(var_0f2bc4c42ad0_err)
		}
		properties["cascade"] = var_0f2bc4c42ad0_mapped
	}

	var_3858bd83d27c := resourceReference.BackReference

	if var_3858bd83d27c != nil {
		var var_3858bd83d27c_mapped *structpb.Value

		var var_3858bd83d27c_err error
		var_3858bd83d27c_mapped, var_3858bd83d27c_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_3858bd83d27c)
		if var_3858bd83d27c_err != nil {
			panic(var_3858bd83d27c_err)
		}
		properties["backReference"] = var_3858bd83d27c_mapped
	}
	return properties
}

func (m *ResourceReferenceMapper) FromProperties(properties map[string]*structpb.Value) *ResourceReference {
	var s = m.New()
	if properties["resource"] != nil {

		var_bd2c8397039b := properties["resource"]
		var_bd2c8397039b_mapped := ResourceMapperInstance.FromProperties(var_bd2c8397039b.GetStructValue().Fields)

		s.Resource = var_bd2c8397039b_mapped
	}
	if properties["cascade"] != nil {

		var_7434a0271f7e := properties["cascade"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_BOOL).UnPack(var_7434a0271f7e)

		if err != nil {
			panic(err)
		}

		var_7434a0271f7e_mapped := new(bool)
		*var_7434a0271f7e_mapped = val.(bool)

		s.Cascade = var_7434a0271f7e_mapped
	}
	if properties["backReference"] != nil {

		var_67205786af7d := properties["backReference"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_67205786af7d)

		if err != nil {
			panic(err)
		}

		var_67205786af7d_mapped := new(string)
		*var_67205786af7d_mapped = val.(string)

		s.BackReference = var_67205786af7d_mapped
	}
	return s
}
