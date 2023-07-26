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

	var_75a76ae0d0e8 := resource.Id

	if var_75a76ae0d0e8 != nil {
		var var_75a76ae0d0e8_mapped *structpb.Value

		var var_75a76ae0d0e8_err error
		var_75a76ae0d0e8_mapped, var_75a76ae0d0e8_err = types.ByResourcePropertyType(model.ResourceProperty_UUID).Pack(*var_75a76ae0d0e8)
		if var_75a76ae0d0e8_err != nil {
			panic(var_75a76ae0d0e8_err)
		}
		properties["id"] = var_75a76ae0d0e8_mapped
	}

	var_5878b79349b3 := resource.Version

	var var_5878b79349b3_mapped *structpb.Value

	var var_5878b79349b3_err error
	var_5878b79349b3_mapped, var_5878b79349b3_err = types.ByResourcePropertyType(model.ResourceProperty_INT32).Pack(var_5878b79349b3)
	if var_5878b79349b3_err != nil {
		panic(var_5878b79349b3_err)
	}
	properties["version"] = var_5878b79349b3_mapped

	var_148441789e5a := resource.CreatedBy

	if var_148441789e5a != nil {
		var var_148441789e5a_mapped *structpb.Value

		var var_148441789e5a_err error
		var_148441789e5a_mapped, var_148441789e5a_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_148441789e5a)
		if var_148441789e5a_err != nil {
			panic(var_148441789e5a_err)
		}
		properties["createdBy"] = var_148441789e5a_mapped
	}

	var_2555a16711fa := resource.UpdatedBy

	if var_2555a16711fa != nil {
		var var_2555a16711fa_mapped *structpb.Value

		var var_2555a16711fa_err error
		var_2555a16711fa_mapped, var_2555a16711fa_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_2555a16711fa)
		if var_2555a16711fa_err != nil {
			panic(var_2555a16711fa_err)
		}
		properties["updatedBy"] = var_2555a16711fa_mapped
	}

	var_9da74b727849 := resource.CreatedOn

	if var_9da74b727849 != nil {
		var var_9da74b727849_mapped *structpb.Value

		var var_9da74b727849_err error
		var_9da74b727849_mapped, var_9da74b727849_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_9da74b727849)
		if var_9da74b727849_err != nil {
			panic(var_9da74b727849_err)
		}
		properties["createdOn"] = var_9da74b727849_mapped
	}

	var_1f5dfe6f035c := resource.UpdatedOn

	if var_1f5dfe6f035c != nil {
		var var_1f5dfe6f035c_mapped *structpb.Value

		var var_1f5dfe6f035c_err error
		var_1f5dfe6f035c_mapped, var_1f5dfe6f035c_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_1f5dfe6f035c)
		if var_1f5dfe6f035c_err != nil {
			panic(var_1f5dfe6f035c_err)
		}
		properties["updatedOn"] = var_1f5dfe6f035c_mapped
	}

	var_b5f7b1d76436 := resource.Name

	var var_b5f7b1d76436_mapped *structpb.Value

	var var_b5f7b1d76436_err error
	var_b5f7b1d76436_mapped, var_b5f7b1d76436_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_b5f7b1d76436)
	if var_b5f7b1d76436_err != nil {
		panic(var_b5f7b1d76436_err)
	}
	properties["name"] = var_b5f7b1d76436_mapped

	var_d55d588d12c1 := resource.Namespace

	if var_d55d588d12c1 != nil {
		var var_d55d588d12c1_mapped *structpb.Value

		var_d55d588d12c1_mapped = structpb.NewStructValue(&structpb.Struct{Fields: NamespaceMapperInstance.ToProperties(var_d55d588d12c1)})
		properties["namespace"] = var_d55d588d12c1_mapped
	}

	var_4892ccd3d672 := resource.Virtual

	var var_4892ccd3d672_mapped *structpb.Value

	var var_4892ccd3d672_err error
	var_4892ccd3d672_mapped, var_4892ccd3d672_err = types.ByResourcePropertyType(model.ResourceProperty_BOOL).Pack(var_4892ccd3d672)
	if var_4892ccd3d672_err != nil {
		panic(var_4892ccd3d672_err)
	}
	properties["virtual"] = var_4892ccd3d672_mapped

	var_b4e6a11ab924 := resource.Properties

	var var_b4e6a11ab924_mapped *structpb.Value

	var var_b4e6a11ab924_l []*structpb.Value
	for _, value := range var_b4e6a11ab924 {

		var_543419e5d136 := value
		var var_543419e5d136_mapped *structpb.Value

		var_543419e5d136_mapped = structpb.NewStructValue(&structpb.Struct{Fields: ResourcePropertyMapperInstance.ToProperties(&var_543419e5d136)})

		var_b4e6a11ab924_l = append(var_b4e6a11ab924_l, var_543419e5d136_mapped)
	}
	var_b4e6a11ab924_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_b4e6a11ab924_l})
	properties["properties"] = var_b4e6a11ab924_mapped

	var_f8f5e0bf9f46 := resource.Indexes

	if var_f8f5e0bf9f46 != nil {
		var var_f8f5e0bf9f46_mapped *structpb.Value

		var var_f8f5e0bf9f46_l []*structpb.Value
		for _, value := range var_f8f5e0bf9f46 {

			var_c56bdacc15d6 := value
			var var_c56bdacc15d6_mapped *structpb.Value

			var_c56bdacc15d6_mapped = structpb.NewStructValue(&structpb.Struct{Fields: ResourceIndexMapperInstance.ToProperties(&var_c56bdacc15d6)})

			var_f8f5e0bf9f46_l = append(var_f8f5e0bf9f46_l, var_c56bdacc15d6_mapped)
		}
		var_f8f5e0bf9f46_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_f8f5e0bf9f46_l})
		properties["indexes"] = var_f8f5e0bf9f46_mapped
	}

	var_2c4557e46f0e := resource.Types

	if var_2c4557e46f0e != nil {
		var var_2c4557e46f0e_mapped *structpb.Value

		var var_2c4557e46f0e_l []*structpb.Value
		for _, value := range var_2c4557e46f0e {

			var_7c7a53c706c6 := value
			var var_7c7a53c706c6_mapped *structpb.Value

			var_7c7a53c706c6_mapped = structpb.NewStructValue(&structpb.Struct{Fields: ResourceSubTypeMapperInstance.ToProperties(&var_7c7a53c706c6)})

			var_2c4557e46f0e_l = append(var_2c4557e46f0e_l, var_7c7a53c706c6_mapped)
		}
		var_2c4557e46f0e_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_2c4557e46f0e_l})
		properties["types"] = var_2c4557e46f0e_mapped
	}

	var_ed890d8b27e0 := resource.Immutable

	var var_ed890d8b27e0_mapped *structpb.Value

	var var_ed890d8b27e0_err error
	var_ed890d8b27e0_mapped, var_ed890d8b27e0_err = types.ByResourcePropertyType(model.ResourceProperty_BOOL).Pack(var_ed890d8b27e0)
	if var_ed890d8b27e0_err != nil {
		panic(var_ed890d8b27e0_err)
	}
	properties["immutable"] = var_ed890d8b27e0_mapped

	var_7a8e0bf1a96e := resource.Abstract

	var var_7a8e0bf1a96e_mapped *structpb.Value

	var var_7a8e0bf1a96e_err error
	var_7a8e0bf1a96e_mapped, var_7a8e0bf1a96e_err = types.ByResourcePropertyType(model.ResourceProperty_BOOL).Pack(var_7a8e0bf1a96e)
	if var_7a8e0bf1a96e_err != nil {
		panic(var_7a8e0bf1a96e_err)
	}
	properties["abstract"] = var_7a8e0bf1a96e_mapped

	var_0b370c4b3e53 := resource.DataSource

	if var_0b370c4b3e53 != nil {
		var var_0b370c4b3e53_mapped *structpb.Value

		var_0b370c4b3e53_mapped = structpb.NewStructValue(&structpb.Struct{Fields: DataSourceMapperInstance.ToProperties(var_0b370c4b3e53)})
		properties["dataSource"] = var_0b370c4b3e53_mapped
	}

	var_fb5042af38cb := resource.Entity

	if var_fb5042af38cb != nil {
		var var_fb5042af38cb_mapped *structpb.Value

		var var_fb5042af38cb_err error
		var_fb5042af38cb_mapped, var_fb5042af38cb_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_fb5042af38cb)
		if var_fb5042af38cb_err != nil {
			panic(var_fb5042af38cb_err)
		}
		properties["entity"] = var_fb5042af38cb_mapped
	}

	var_00bbfa929de1 := resource.Catalog

	if var_00bbfa929de1 != nil {
		var var_00bbfa929de1_mapped *structpb.Value

		var var_00bbfa929de1_err error
		var_00bbfa929de1_mapped, var_00bbfa929de1_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_00bbfa929de1)
		if var_00bbfa929de1_err != nil {
			panic(var_00bbfa929de1_err)
		}
		properties["catalog"] = var_00bbfa929de1_mapped
	}

	var_47daef6ee9ad := resource.Title

	if var_47daef6ee9ad != nil {
		var var_47daef6ee9ad_mapped *structpb.Value

		var var_47daef6ee9ad_err error
		var_47daef6ee9ad_mapped, var_47daef6ee9ad_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_47daef6ee9ad)
		if var_47daef6ee9ad_err != nil {
			panic(var_47daef6ee9ad_err)
		}
		properties["title"] = var_47daef6ee9ad_mapped
	}

	var_ef3afbe64568 := resource.Description

	if var_ef3afbe64568 != nil {
		var var_ef3afbe64568_mapped *structpb.Value

		var var_ef3afbe64568_err error
		var_ef3afbe64568_mapped, var_ef3afbe64568_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_ef3afbe64568)
		if var_ef3afbe64568_err != nil {
			panic(var_ef3afbe64568_err)
		}
		properties["description"] = var_ef3afbe64568_mapped
	}

	var_f4dbf40fe282 := resource.Annotations

	if var_f4dbf40fe282 != nil {
		var var_f4dbf40fe282_mapped *structpb.Value

		var var_f4dbf40fe282_st *structpb.Struct = new(structpb.Struct)
		var_f4dbf40fe282_st.Fields = make(map[string]*structpb.Value)
		for key, value := range var_f4dbf40fe282 {

			var_04421427ee2e := value
			var var_04421427ee2e_mapped *structpb.Value

			var var_04421427ee2e_err error
			var_04421427ee2e_mapped, var_04421427ee2e_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_04421427ee2e)
			if var_04421427ee2e_err != nil {
				panic(var_04421427ee2e_err)
			}

			var_f4dbf40fe282_st.Fields[key] = var_04421427ee2e_mapped
		}
		var_f4dbf40fe282_mapped = structpb.NewStructValue(var_f4dbf40fe282_st)
		properties["annotations"] = var_f4dbf40fe282_mapped
	}
	return properties
}

func (m *ResourceMapper) FromProperties(properties map[string]*structpb.Value) *Resource {
	var s = m.New()
	if properties["id"] != nil {

		var_82374fc19722 := properties["id"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_UUID).UnPack(var_82374fc19722)

		if err != nil {
			panic(err)
		}

		var_82374fc19722_mapped := new(uuid.UUID)
		*var_82374fc19722_mapped = val.(uuid.UUID)

		s.Id = var_82374fc19722_mapped
	}
	if properties["version"] != nil {

		var_a3ccbb0608ea := properties["version"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(var_a3ccbb0608ea)

		if err != nil {
			panic(err)
		}

		var_a3ccbb0608ea_mapped := val.(int32)

		s.Version = var_a3ccbb0608ea_mapped
	}
	if properties["createdBy"] != nil {

		var_1d1bd98e2d8f := properties["createdBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_1d1bd98e2d8f)

		if err != nil {
			panic(err)
		}

		var_1d1bd98e2d8f_mapped := new(string)
		*var_1d1bd98e2d8f_mapped = val.(string)

		s.CreatedBy = var_1d1bd98e2d8f_mapped
	}
	if properties["updatedBy"] != nil {

		var_a712006797e9 := properties["updatedBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_a712006797e9)

		if err != nil {
			panic(err)
		}

		var_a712006797e9_mapped := new(string)
		*var_a712006797e9_mapped = val.(string)

		s.UpdatedBy = var_a712006797e9_mapped
	}
	if properties["createdOn"] != nil {

		var_2568f7e337d6 := properties["createdOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_2568f7e337d6)

		if err != nil {
			panic(err)
		}

		var_2568f7e337d6_mapped := new(time.Time)
		*var_2568f7e337d6_mapped = val.(time.Time)

		s.CreatedOn = var_2568f7e337d6_mapped
	}
	if properties["updatedOn"] != nil {

		var_b23a776658fb := properties["updatedOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_b23a776658fb)

		if err != nil {
			panic(err)
		}

		var_b23a776658fb_mapped := new(time.Time)
		*var_b23a776658fb_mapped = val.(time.Time)

		s.UpdatedOn = var_b23a776658fb_mapped
	}
	if properties["name"] != nil {

		var_0cdf6aa7f270 := properties["name"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_0cdf6aa7f270)

		if err != nil {
			panic(err)
		}

		var_0cdf6aa7f270_mapped := val.(string)

		s.Name = var_0cdf6aa7f270_mapped
	}
	if properties["namespace"] != nil {

		var_deeca6982541 := properties["namespace"]
		var_deeca6982541_mapped := NamespaceMapperInstance.FromProperties(var_deeca6982541.GetStructValue().Fields)

		s.Namespace = var_deeca6982541_mapped
	}
	if properties["virtual"] != nil {

		var_4ce458ec78d5 := properties["virtual"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_BOOL).UnPack(var_4ce458ec78d5)

		if err != nil {
			panic(err)
		}

		var_4ce458ec78d5_mapped := val.(bool)

		s.Virtual = var_4ce458ec78d5_mapped
	}
	if properties["properties"] != nil {

		var_275db4e0ec70 := properties["properties"]
		var_275db4e0ec70_mapped := []ResourceProperty{}
		for _, v := range var_275db4e0ec70.GetListValue().Values {

			var_5fa60f787537 := v
			var mappedValue = ResourcePropertyMapperInstance.FromProperties(var_5fa60f787537.GetStructValue().Fields)

			var_5fa60f787537_mapped := *mappedValue

			var_275db4e0ec70_mapped = append(var_275db4e0ec70_mapped, var_5fa60f787537_mapped)
		}

		s.Properties = var_275db4e0ec70_mapped
	}
	if properties["indexes"] != nil {

		var_a5d2bf20f19c := properties["indexes"]
		var_a5d2bf20f19c_mapped := []ResourceIndex{}
		for _, v := range var_a5d2bf20f19c.GetListValue().Values {

			var_c1a517be53dd := v
			var mappedValue = ResourceIndexMapperInstance.FromProperties(var_c1a517be53dd.GetStructValue().Fields)

			var_c1a517be53dd_mapped := *mappedValue

			var_a5d2bf20f19c_mapped = append(var_a5d2bf20f19c_mapped, var_c1a517be53dd_mapped)
		}

		s.Indexes = var_a5d2bf20f19c_mapped
	}
	if properties["types"] != nil {

		var_243492252312 := properties["types"]
		var_243492252312_mapped := []ResourceSubType{}
		for _, v := range var_243492252312.GetListValue().Values {

			var_23802e40cdc8 := v
			var mappedValue = ResourceSubTypeMapperInstance.FromProperties(var_23802e40cdc8.GetStructValue().Fields)

			var_23802e40cdc8_mapped := *mappedValue

			var_243492252312_mapped = append(var_243492252312_mapped, var_23802e40cdc8_mapped)
		}

		s.Types = var_243492252312_mapped
	}
	if properties["immutable"] != nil {

		var_e7857e3fca71 := properties["immutable"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_BOOL).UnPack(var_e7857e3fca71)

		if err != nil {
			panic(err)
		}

		var_e7857e3fca71_mapped := val.(bool)

		s.Immutable = var_e7857e3fca71_mapped
	}
	if properties["abstract"] != nil {

		var_a780f0b3db40 := properties["abstract"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_BOOL).UnPack(var_a780f0b3db40)

		if err != nil {
			panic(err)
		}

		var_a780f0b3db40_mapped := val.(bool)

		s.Abstract = var_a780f0b3db40_mapped
	}
	if properties["dataSource"] != nil {

		var_b49ee21d90e6 := properties["dataSource"]
		var_b49ee21d90e6_mapped := DataSourceMapperInstance.FromProperties(var_b49ee21d90e6.GetStructValue().Fields)

		s.DataSource = var_b49ee21d90e6_mapped
	}
	if properties["entity"] != nil {

		var_f10446a845c9 := properties["entity"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_f10446a845c9)

		if err != nil {
			panic(err)
		}

		var_f10446a845c9_mapped := new(string)
		*var_f10446a845c9_mapped = val.(string)

		s.Entity = var_f10446a845c9_mapped
	}
	if properties["catalog"] != nil {

		var_cb318eeca76e := properties["catalog"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_cb318eeca76e)

		if err != nil {
			panic(err)
		}

		var_cb318eeca76e_mapped := new(string)
		*var_cb318eeca76e_mapped = val.(string)

		s.Catalog = var_cb318eeca76e_mapped
	}
	if properties["title"] != nil {

		var_84bf1db7108a := properties["title"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_84bf1db7108a)

		if err != nil {
			panic(err)
		}

		var_84bf1db7108a_mapped := new(string)
		*var_84bf1db7108a_mapped = val.(string)

		s.Title = var_84bf1db7108a_mapped
	}
	if properties["description"] != nil {

		var_d237c66cc643 := properties["description"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_d237c66cc643)

		if err != nil {
			panic(err)
		}

		var_d237c66cc643_mapped := new(string)
		*var_d237c66cc643_mapped = val.(string)

		s.Description = var_d237c66cc643_mapped
	}
	if properties["annotations"] != nil {

		var_4c366a00a8cf := properties["annotations"]
		var_4c366a00a8cf_mapped := make(map[string]string)
		for k, v := range var_4c366a00a8cf.GetStructValue().Fields {

			var_28a4b9280e90 := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_28a4b9280e90)

			if err != nil {
				panic(err)
			}

			var_28a4b9280e90_mapped := val.(string)

			var_4c366a00a8cf_mapped[k] = var_28a4b9280e90_mapped
		}

		s.Annotations = var_4c366a00a8cf_mapped
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

	var_93bbc8373f4c := resourceProperty.Name

	var var_93bbc8373f4c_mapped *structpb.Value

	var var_93bbc8373f4c_err error
	var_93bbc8373f4c_mapped, var_93bbc8373f4c_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_93bbc8373f4c)
	if var_93bbc8373f4c_err != nil {
		panic(var_93bbc8373f4c_err)
	}
	properties["name"] = var_93bbc8373f4c_mapped

	var_bca349c3b03b := resourceProperty.Type

	var var_bca349c3b03b_mapped *structpb.Value

	var var_bca349c3b03b_err error
	var_bca349c3b03b_mapped, var_bca349c3b03b_err = types.ByResourcePropertyType(model.ResourceProperty_INT32).Pack(var_bca349c3b03b)
	if var_bca349c3b03b_err != nil {
		panic(var_bca349c3b03b_err)
	}
	properties["type"] = var_bca349c3b03b_mapped

	var_f239bd99ec28 := resourceProperty.TypeRef

	if var_f239bd99ec28 != nil {
		var var_f239bd99ec28_mapped *structpb.Value

		var var_f239bd99ec28_err error
		var_f239bd99ec28_mapped, var_f239bd99ec28_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_f239bd99ec28)
		if var_f239bd99ec28_err != nil {
			panic(var_f239bd99ec28_err)
		}
		properties["typeRef"] = var_f239bd99ec28_mapped
	}

	var_c95b8bbbc015 := resourceProperty.Mapping

	var var_c95b8bbbc015_mapped *structpb.Value

	var var_c95b8bbbc015_err error
	var_c95b8bbbc015_mapped, var_c95b8bbbc015_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_c95b8bbbc015)
	if var_c95b8bbbc015_err != nil {
		panic(var_c95b8bbbc015_err)
	}
	properties["mapping"] = var_c95b8bbbc015_mapped

	var_ee70e75ac993 := resourceProperty.Primary

	var var_ee70e75ac993_mapped *structpb.Value

	var var_ee70e75ac993_err error
	var_ee70e75ac993_mapped, var_ee70e75ac993_err = types.ByResourcePropertyType(model.ResourceProperty_BOOL).Pack(var_ee70e75ac993)
	if var_ee70e75ac993_err != nil {
		panic(var_ee70e75ac993_err)
	}
	properties["primary"] = var_ee70e75ac993_mapped

	var_87ac1ed8719d := resourceProperty.Required

	var var_87ac1ed8719d_mapped *structpb.Value

	var var_87ac1ed8719d_err error
	var_87ac1ed8719d_mapped, var_87ac1ed8719d_err = types.ByResourcePropertyType(model.ResourceProperty_BOOL).Pack(var_87ac1ed8719d)
	if var_87ac1ed8719d_err != nil {
		panic(var_87ac1ed8719d_err)
	}
	properties["required"] = var_87ac1ed8719d_mapped

	var_43e1c9787cc8 := resourceProperty.Unique

	var var_43e1c9787cc8_mapped *structpb.Value

	var var_43e1c9787cc8_err error
	var_43e1c9787cc8_mapped, var_43e1c9787cc8_err = types.ByResourcePropertyType(model.ResourceProperty_BOOL).Pack(var_43e1c9787cc8)
	if var_43e1c9787cc8_err != nil {
		panic(var_43e1c9787cc8_err)
	}
	properties["unique"] = var_43e1c9787cc8_mapped

	var_238ccf6b59fa := resourceProperty.Immutable

	var var_238ccf6b59fa_mapped *structpb.Value

	var var_238ccf6b59fa_err error
	var_238ccf6b59fa_mapped, var_238ccf6b59fa_err = types.ByResourcePropertyType(model.ResourceProperty_BOOL).Pack(var_238ccf6b59fa)
	if var_238ccf6b59fa_err != nil {
		panic(var_238ccf6b59fa_err)
	}
	properties["immutable"] = var_238ccf6b59fa_mapped

	var_23d11ea5ba16 := resourceProperty.Length

	var var_23d11ea5ba16_mapped *structpb.Value

	var var_23d11ea5ba16_err error
	var_23d11ea5ba16_mapped, var_23d11ea5ba16_err = types.ByResourcePropertyType(model.ResourceProperty_INT32).Pack(var_23d11ea5ba16)
	if var_23d11ea5ba16_err != nil {
		panic(var_23d11ea5ba16_err)
	}
	properties["length"] = var_23d11ea5ba16_mapped

	var_1b848b8fd610 := resourceProperty.Item

	if var_1b848b8fd610 != nil {
		var var_1b848b8fd610_mapped *structpb.Value

		var_1b848b8fd610_mapped = structpb.NewStructValue(&structpb.Struct{Fields: ResourcePropertyMapperInstance.ToProperties(var_1b848b8fd610)})
		properties["item"] = var_1b848b8fd610_mapped
	}

	var_1a0f1a0ea039 := resourceProperty.Properties

	var var_1a0f1a0ea039_mapped *structpb.Value

	var var_1a0f1a0ea039_l []*structpb.Value
	for _, value := range var_1a0f1a0ea039 {

		var_a2f5a9c53a2f := value
		var var_a2f5a9c53a2f_mapped *structpb.Value

		var_a2f5a9c53a2f_mapped = structpb.NewStructValue(&structpb.Struct{Fields: ResourcePropertyMapperInstance.ToProperties(&var_a2f5a9c53a2f)})

		var_1a0f1a0ea039_l = append(var_1a0f1a0ea039_l, var_a2f5a9c53a2f_mapped)
	}
	var_1a0f1a0ea039_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_1a0f1a0ea039_l})
	properties["properties"] = var_1a0f1a0ea039_mapped

	var_d082d6211c3e := resourceProperty.Reference

	if var_d082d6211c3e != nil {
		var var_d082d6211c3e_mapped *structpb.Value

		var_d082d6211c3e_mapped = structpb.NewStructValue(&structpb.Struct{Fields: ResourceReferenceMapperInstance.ToProperties(var_d082d6211c3e)})
		properties["reference"] = var_d082d6211c3e_mapped
	}

	var_1c816970fa99 := resourceProperty.DefaultValue

	if var_1c816970fa99 != nil {
		var var_1c816970fa99_mapped *structpb.Value

		var var_1c816970fa99_err error
		var_1c816970fa99_mapped, var_1c816970fa99_err = types.ByResourcePropertyType(model.ResourceProperty_OBJECT).Pack(*var_1c816970fa99)
		if var_1c816970fa99_err != nil {
			panic(var_1c816970fa99_err)
		}
		properties["defaultValue"] = var_1c816970fa99_mapped
	}

	var_36fd7a8fadff := resourceProperty.EnumValues

	if var_36fd7a8fadff != nil {
		var var_36fd7a8fadff_mapped *structpb.Value

		var var_36fd7a8fadff_l []*structpb.Value
		for _, value := range var_36fd7a8fadff {

			var_33f9e1b9b865 := value
			var var_33f9e1b9b865_mapped *structpb.Value

			var var_33f9e1b9b865_err error
			var_33f9e1b9b865_mapped, var_33f9e1b9b865_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_33f9e1b9b865)
			if var_33f9e1b9b865_err != nil {
				panic(var_33f9e1b9b865_err)
			}

			var_36fd7a8fadff_l = append(var_36fd7a8fadff_l, var_33f9e1b9b865_mapped)
		}
		var_36fd7a8fadff_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_36fd7a8fadff_l})
		properties["enumValues"] = var_36fd7a8fadff_mapped
	}

	var_2b28b7e730f7 := resourceProperty.ExampleValue

	if var_2b28b7e730f7 != nil {
		var var_2b28b7e730f7_mapped *structpb.Value

		var var_2b28b7e730f7_err error
		var_2b28b7e730f7_mapped, var_2b28b7e730f7_err = types.ByResourcePropertyType(model.ResourceProperty_OBJECT).Pack(*var_2b28b7e730f7)
		if var_2b28b7e730f7_err != nil {
			panic(var_2b28b7e730f7_err)
		}
		properties["exampleValue"] = var_2b28b7e730f7_mapped
	}

	var_061e2fb16ed3 := resourceProperty.Title

	if var_061e2fb16ed3 != nil {
		var var_061e2fb16ed3_mapped *structpb.Value

		var var_061e2fb16ed3_err error
		var_061e2fb16ed3_mapped, var_061e2fb16ed3_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_061e2fb16ed3)
		if var_061e2fb16ed3_err != nil {
			panic(var_061e2fb16ed3_err)
		}
		properties["title"] = var_061e2fb16ed3_mapped
	}

	var_d4e5e87961c7 := resourceProperty.Description

	if var_d4e5e87961c7 != nil {
		var var_d4e5e87961c7_mapped *structpb.Value

		var var_d4e5e87961c7_err error
		var_d4e5e87961c7_mapped, var_d4e5e87961c7_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_d4e5e87961c7)
		if var_d4e5e87961c7_err != nil {
			panic(var_d4e5e87961c7_err)
		}
		properties["description"] = var_d4e5e87961c7_mapped
	}

	var_428d29b9ce87 := resourceProperty.Annotations

	if var_428d29b9ce87 != nil {
		var var_428d29b9ce87_mapped *structpb.Value

		var var_428d29b9ce87_st *structpb.Struct = new(structpb.Struct)
		var_428d29b9ce87_st.Fields = make(map[string]*structpb.Value)
		for key, value := range var_428d29b9ce87 {

			var_25b57704ca67 := value
			var var_25b57704ca67_mapped *structpb.Value

			var var_25b57704ca67_err error
			var_25b57704ca67_mapped, var_25b57704ca67_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_25b57704ca67)
			if var_25b57704ca67_err != nil {
				panic(var_25b57704ca67_err)
			}

			var_428d29b9ce87_st.Fields[key] = var_25b57704ca67_mapped
		}
		var_428d29b9ce87_mapped = structpb.NewStructValue(var_428d29b9ce87_st)
		properties["annotations"] = var_428d29b9ce87_mapped
	}
	return properties
}

func (m *ResourcePropertyMapper) FromProperties(properties map[string]*structpb.Value) *ResourceProperty {
	var s = m.New()
	if properties["name"] != nil {

		var_7ee9a83b0775 := properties["name"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_7ee9a83b0775)

		if err != nil {
			panic(err)
		}

		var_7ee9a83b0775_mapped := val.(string)

		s.Name = var_7ee9a83b0775_mapped
	}
	if properties["type"] != nil {

		var_626882d8794a := properties["type"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(var_626882d8794a)

		if err != nil {
			panic(err)
		}

		var_626882d8794a_mapped := val.(int32)

		s.Type = var_626882d8794a_mapped
	}
	if properties["typeRef"] != nil {

		var_c52e1b0abc64 := properties["typeRef"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_c52e1b0abc64)

		if err != nil {
			panic(err)
		}

		var_c52e1b0abc64_mapped := new(string)
		*var_c52e1b0abc64_mapped = val.(string)

		s.TypeRef = var_c52e1b0abc64_mapped
	}
	if properties["mapping"] != nil {

		var_bdaab50367ea := properties["mapping"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_bdaab50367ea)

		if err != nil {
			panic(err)
		}

		var_bdaab50367ea_mapped := val.(string)

		s.Mapping = var_bdaab50367ea_mapped
	}
	if properties["primary"] != nil {

		var_024781914dac := properties["primary"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_BOOL).UnPack(var_024781914dac)

		if err != nil {
			panic(err)
		}

		var_024781914dac_mapped := val.(bool)

		s.Primary = var_024781914dac_mapped
	}
	if properties["required"] != nil {

		var_820dd12c5487 := properties["required"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_BOOL).UnPack(var_820dd12c5487)

		if err != nil {
			panic(err)
		}

		var_820dd12c5487_mapped := val.(bool)

		s.Required = var_820dd12c5487_mapped
	}
	if properties["unique"] != nil {

		var_6413baef78d1 := properties["unique"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_BOOL).UnPack(var_6413baef78d1)

		if err != nil {
			panic(err)
		}

		var_6413baef78d1_mapped := val.(bool)

		s.Unique = var_6413baef78d1_mapped
	}
	if properties["immutable"] != nil {

		var_6bc49fbd8697 := properties["immutable"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_BOOL).UnPack(var_6bc49fbd8697)

		if err != nil {
			panic(err)
		}

		var_6bc49fbd8697_mapped := val.(bool)

		s.Immutable = var_6bc49fbd8697_mapped
	}
	if properties["length"] != nil {

		var_77bdb0cdfffb := properties["length"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(var_77bdb0cdfffb)

		if err != nil {
			panic(err)
		}

		var_77bdb0cdfffb_mapped := val.(int32)

		s.Length = var_77bdb0cdfffb_mapped
	}
	if properties["item"] != nil {

		var_9e14d1057d03 := properties["item"]
		var mappedValue = ResourcePropertyMapperInstance.FromProperties(var_9e14d1057d03.GetStructValue().Fields)

		var_9e14d1057d03_mapped := mappedValue

		s.Item = var_9e14d1057d03_mapped
	}
	if properties["properties"] != nil {

		var_cbab87d043fc := properties["properties"]
		var_cbab87d043fc_mapped := []ResourceProperty{}
		for _, v := range var_cbab87d043fc.GetListValue().Values {

			var_a5bd90fee5a9 := v
			var mappedValue = ResourcePropertyMapperInstance.FromProperties(var_a5bd90fee5a9.GetStructValue().Fields)

			var_a5bd90fee5a9_mapped := *mappedValue

			var_cbab87d043fc_mapped = append(var_cbab87d043fc_mapped, var_a5bd90fee5a9_mapped)
		}

		s.Properties = var_cbab87d043fc_mapped
	}
	if properties["reference"] != nil {

		var_7ab0c2bf85dc := properties["reference"]
		var mappedValue = ResourceReferenceMapperInstance.FromProperties(var_7ab0c2bf85dc.GetStructValue().Fields)

		var_7ab0c2bf85dc_mapped := mappedValue

		s.Reference = var_7ab0c2bf85dc_mapped
	}
	if properties["defaultValue"] != nil {

		var_fb71845b6430 := properties["defaultValue"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_OBJECT).UnPack(var_fb71845b6430)

		if err != nil {
			panic(err)
		}

		var_fb71845b6430_mapped := new(unstructured.Unstructured)
		*var_fb71845b6430_mapped = val.(unstructured.Unstructured)

		s.DefaultValue = var_fb71845b6430_mapped
	}
	if properties["enumValues"] != nil {

		var_47bf674de0a7 := properties["enumValues"]
		var_47bf674de0a7_mapped := []string{}
		for _, v := range var_47bf674de0a7.GetListValue().Values {

			var_2856eefee2ed := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_2856eefee2ed)

			if err != nil {
				panic(err)
			}

			var_2856eefee2ed_mapped := val.(string)

			var_47bf674de0a7_mapped = append(var_47bf674de0a7_mapped, var_2856eefee2ed_mapped)
		}

		s.EnumValues = var_47bf674de0a7_mapped
	}
	if properties["exampleValue"] != nil {

		var_1b5a79da5328 := properties["exampleValue"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_OBJECT).UnPack(var_1b5a79da5328)

		if err != nil {
			panic(err)
		}

		var_1b5a79da5328_mapped := new(unstructured.Unstructured)
		*var_1b5a79da5328_mapped = val.(unstructured.Unstructured)

		s.ExampleValue = var_1b5a79da5328_mapped
	}
	if properties["title"] != nil {

		var_bbc67f788540 := properties["title"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_bbc67f788540)

		if err != nil {
			panic(err)
		}

		var_bbc67f788540_mapped := new(string)
		*var_bbc67f788540_mapped = val.(string)

		s.Title = var_bbc67f788540_mapped
	}
	if properties["description"] != nil {

		var_f0a69e843598 := properties["description"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_f0a69e843598)

		if err != nil {
			panic(err)
		}

		var_f0a69e843598_mapped := new(string)
		*var_f0a69e843598_mapped = val.(string)

		s.Description = var_f0a69e843598_mapped
	}
	if properties["annotations"] != nil {

		var_17758e36f1c5 := properties["annotations"]
		var_17758e36f1c5_mapped := make(map[string]string)
		for k, v := range var_17758e36f1c5.GetStructValue().Fields {

			var_a83469b59c12 := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_a83469b59c12)

			if err != nil {
				panic(err)
			}

			var_a83469b59c12_mapped := val.(string)

			var_17758e36f1c5_mapped[k] = var_a83469b59c12_mapped
		}

		s.Annotations = var_17758e36f1c5_mapped
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

	var_10d94768a74f := resourceSubType.Name

	var var_10d94768a74f_mapped *structpb.Value

	var var_10d94768a74f_err error
	var_10d94768a74f_mapped, var_10d94768a74f_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_10d94768a74f)
	if var_10d94768a74f_err != nil {
		panic(var_10d94768a74f_err)
	}
	properties["name"] = var_10d94768a74f_mapped

	var_bd2d6229dbac := resourceSubType.Properties

	var var_bd2d6229dbac_mapped *structpb.Value

	var var_bd2d6229dbac_l []*structpb.Value
	for _, value := range var_bd2d6229dbac {

		var_3405d5006df4 := value
		var var_3405d5006df4_mapped *structpb.Value

		var_3405d5006df4_mapped = structpb.NewStructValue(&structpb.Struct{Fields: ResourcePropertyMapperInstance.ToProperties(&var_3405d5006df4)})

		var_bd2d6229dbac_l = append(var_bd2d6229dbac_l, var_3405d5006df4_mapped)
	}
	var_bd2d6229dbac_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_bd2d6229dbac_l})
	properties["properties"] = var_bd2d6229dbac_mapped
	return properties
}

func (m *ResourceSubTypeMapper) FromProperties(properties map[string]*structpb.Value) *ResourceSubType {
	var s = m.New()
	if properties["name"] != nil {

		var_f9ff9c5c8980 := properties["name"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_f9ff9c5c8980)

		if err != nil {
			panic(err)
		}

		var_f9ff9c5c8980_mapped := val.(string)

		s.Name = var_f9ff9c5c8980_mapped
	}
	if properties["properties"] != nil {

		var_a97d3516d4c1 := properties["properties"]
		var_a97d3516d4c1_mapped := []ResourceProperty{}
		for _, v := range var_a97d3516d4c1.GetListValue().Values {

			var_45b449cc7e27 := v
			var mappedValue = ResourcePropertyMapperInstance.FromProperties(var_45b449cc7e27.GetStructValue().Fields)

			var_45b449cc7e27_mapped := *mappedValue

			var_a97d3516d4c1_mapped = append(var_a97d3516d4c1_mapped, var_45b449cc7e27_mapped)
		}

		s.Properties = var_a97d3516d4c1_mapped
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

	var_ae26db42bff2 := resourceIndexProperty.Name

	var var_ae26db42bff2_mapped *structpb.Value

	var var_ae26db42bff2_err error
	var_ae26db42bff2_mapped, var_ae26db42bff2_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_ae26db42bff2)
	if var_ae26db42bff2_err != nil {
		panic(var_ae26db42bff2_err)
	}
	properties["name"] = var_ae26db42bff2_mapped

	var_e8c47758b070 := resourceIndexProperty.Order

	if var_e8c47758b070 != nil {
		var var_e8c47758b070_mapped *structpb.Value

		var var_e8c47758b070_err error
		var_e8c47758b070_mapped, var_e8c47758b070_err = types.ByResourcePropertyType(model.ResourceProperty_ENUM).Pack(string(*var_e8c47758b070))
		if var_e8c47758b070_err != nil {
			panic(var_e8c47758b070_err)
		}
		properties["order"] = var_e8c47758b070_mapped
	}
	return properties
}

func (m *ResourceIndexPropertyMapper) FromProperties(properties map[string]*structpb.Value) *ResourceIndexProperty {
	var s = m.New()
	if properties["name"] != nil {

		var_3714b33a86a5 := properties["name"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_3714b33a86a5)

		if err != nil {
			panic(err)
		}

		var_3714b33a86a5_mapped := val.(string)

		s.Name = var_3714b33a86a5_mapped
	}
	if properties["order"] != nil {

		var_e30c1fa89196 := properties["order"]
		var_e30c1fa89196_mapped := new(ResourceOrder)
		*var_e30c1fa89196_mapped = (ResourceOrder)(var_e30c1fa89196.GetStringValue())

		s.Order = var_e30c1fa89196_mapped
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

	var_54f60f1e293a := resourceIndex.Properties

	if var_54f60f1e293a != nil {
		var var_54f60f1e293a_mapped *structpb.Value

		var var_54f60f1e293a_l []*structpb.Value
		for _, value := range var_54f60f1e293a {

			var_1c9c4ad33a69 := value
			var var_1c9c4ad33a69_mapped *structpb.Value

			var_1c9c4ad33a69_mapped = structpb.NewStructValue(&structpb.Struct{Fields: ResourceIndexPropertyMapperInstance.ToProperties(&var_1c9c4ad33a69)})

			var_54f60f1e293a_l = append(var_54f60f1e293a_l, var_1c9c4ad33a69_mapped)
		}
		var_54f60f1e293a_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_54f60f1e293a_l})
		properties["properties"] = var_54f60f1e293a_mapped
	}

	var_0177fd5a5b80 := resourceIndex.IndexType

	if var_0177fd5a5b80 != nil {
		var var_0177fd5a5b80_mapped *structpb.Value

		var var_0177fd5a5b80_err error
		var_0177fd5a5b80_mapped, var_0177fd5a5b80_err = types.ByResourcePropertyType(model.ResourceProperty_ENUM).Pack(string(*var_0177fd5a5b80))
		if var_0177fd5a5b80_err != nil {
			panic(var_0177fd5a5b80_err)
		}
		properties["indexType"] = var_0177fd5a5b80_mapped
	}

	var_838105afda54 := resourceIndex.Unique

	if var_838105afda54 != nil {
		var var_838105afda54_mapped *structpb.Value

		var var_838105afda54_err error
		var_838105afda54_mapped, var_838105afda54_err = types.ByResourcePropertyType(model.ResourceProperty_BOOL).Pack(*var_838105afda54)
		if var_838105afda54_err != nil {
			panic(var_838105afda54_err)
		}
		properties["unique"] = var_838105afda54_mapped
	}

	var_158e459a3c11 := resourceIndex.Annotations

	if var_158e459a3c11 != nil {
		var var_158e459a3c11_mapped *structpb.Value

		var var_158e459a3c11_st *structpb.Struct = new(structpb.Struct)
		var_158e459a3c11_st.Fields = make(map[string]*structpb.Value)
		for key, value := range var_158e459a3c11 {

			var_ec6f39e5484b := value
			var var_ec6f39e5484b_mapped *structpb.Value

			var var_ec6f39e5484b_err error
			var_ec6f39e5484b_mapped, var_ec6f39e5484b_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_ec6f39e5484b)
			if var_ec6f39e5484b_err != nil {
				panic(var_ec6f39e5484b_err)
			}

			var_158e459a3c11_st.Fields[key] = var_ec6f39e5484b_mapped
		}
		var_158e459a3c11_mapped = structpb.NewStructValue(var_158e459a3c11_st)
		properties["annotations"] = var_158e459a3c11_mapped
	}
	return properties
}

func (m *ResourceIndexMapper) FromProperties(properties map[string]*structpb.Value) *ResourceIndex {
	var s = m.New()
	if properties["properties"] != nil {

		var_ff65e25ac94d := properties["properties"]
		var_ff65e25ac94d_mapped := []ResourceIndexProperty{}
		for _, v := range var_ff65e25ac94d.GetListValue().Values {

			var_d201e5d6a492 := v
			var mappedValue = ResourceIndexPropertyMapperInstance.FromProperties(var_d201e5d6a492.GetStructValue().Fields)

			var_d201e5d6a492_mapped := *mappedValue

			var_ff65e25ac94d_mapped = append(var_ff65e25ac94d_mapped, var_d201e5d6a492_mapped)
		}

		s.Properties = var_ff65e25ac94d_mapped
	}
	if properties["indexType"] != nil {

		var_64134384b05e := properties["indexType"]
		var_64134384b05e_mapped := new(ResourceIndexType)
		*var_64134384b05e_mapped = (ResourceIndexType)(var_64134384b05e.GetStringValue())

		s.IndexType = var_64134384b05e_mapped
	}
	if properties["unique"] != nil {

		var_28454dfbfed4 := properties["unique"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_BOOL).UnPack(var_28454dfbfed4)

		if err != nil {
			panic(err)
		}

		var_28454dfbfed4_mapped := new(bool)
		*var_28454dfbfed4_mapped = val.(bool)

		s.Unique = var_28454dfbfed4_mapped
	}
	if properties["annotations"] != nil {

		var_705e2d3f6618 := properties["annotations"]
		var_705e2d3f6618_mapped := make(map[string]string)
		for k, v := range var_705e2d3f6618.GetStructValue().Fields {

			var_36c644b5ab18 := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_36c644b5ab18)

			if err != nil {
				panic(err)
			}

			var_36c644b5ab18_mapped := val.(string)

			var_705e2d3f6618_mapped[k] = var_36c644b5ab18_mapped
		}

		s.Annotations = var_705e2d3f6618_mapped
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

	var_84a8a9b9e39f := resourceReference.Resource

	if var_84a8a9b9e39f != nil {
		var var_84a8a9b9e39f_mapped *structpb.Value

		var_84a8a9b9e39f_mapped = structpb.NewStructValue(&structpb.Struct{Fields: ResourceMapperInstance.ToProperties(var_84a8a9b9e39f)})
		properties["resource"] = var_84a8a9b9e39f_mapped
	}

	var_64b9386e91e4 := resourceReference.Cascade

	if var_64b9386e91e4 != nil {
		var var_64b9386e91e4_mapped *structpb.Value

		var var_64b9386e91e4_err error
		var_64b9386e91e4_mapped, var_64b9386e91e4_err = types.ByResourcePropertyType(model.ResourceProperty_BOOL).Pack(*var_64b9386e91e4)
		if var_64b9386e91e4_err != nil {
			panic(var_64b9386e91e4_err)
		}
		properties["cascade"] = var_64b9386e91e4_mapped
	}

	var_a4bfc134763f := resourceReference.BackReference

	if var_a4bfc134763f != nil {
		var var_a4bfc134763f_mapped *structpb.Value

		var var_a4bfc134763f_err error
		var_a4bfc134763f_mapped, var_a4bfc134763f_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_a4bfc134763f)
		if var_a4bfc134763f_err != nil {
			panic(var_a4bfc134763f_err)
		}
		properties["backReference"] = var_a4bfc134763f_mapped
	}
	return properties
}

func (m *ResourceReferenceMapper) FromProperties(properties map[string]*structpb.Value) *ResourceReference {
	var s = m.New()
	if properties["resource"] != nil {

		var_1fd3382cef98 := properties["resource"]
		var_1fd3382cef98_mapped := ResourceMapperInstance.FromProperties(var_1fd3382cef98.GetStructValue().Fields)

		s.Resource = var_1fd3382cef98_mapped
	}
	if properties["cascade"] != nil {

		var_217fbbb5ddbf := properties["cascade"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_BOOL).UnPack(var_217fbbb5ddbf)

		if err != nil {
			panic(err)
		}

		var_217fbbb5ddbf_mapped := new(bool)
		*var_217fbbb5ddbf_mapped = val.(bool)

		s.Cascade = var_217fbbb5ddbf_mapped
	}
	if properties["backReference"] != nil {

		var_9231bc899fb5 := properties["backReference"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_9231bc899fb5)

		if err != nil {
			panic(err)
		}

		var_9231bc899fb5_mapped := new(string)
		*var_9231bc899fb5_mapped = val.(string)

		s.BackReference = var_9231bc899fb5_mapped
	}
	return s
}
