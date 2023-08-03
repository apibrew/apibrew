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

	var_6ccafa09f059 := resource.Id

	if var_6ccafa09f059 != nil {
		var var_6ccafa09f059_mapped *structpb.Value

		var var_6ccafa09f059_err error
		var_6ccafa09f059_mapped, var_6ccafa09f059_err = types.ByResourcePropertyType(model.ResourceProperty_UUID).Pack(*var_6ccafa09f059)
		if var_6ccafa09f059_err != nil {
			panic(var_6ccafa09f059_err)
		}
		properties["id"] = var_6ccafa09f059_mapped
	}

	var_fe215b82b2b8 := resource.Version

	var var_fe215b82b2b8_mapped *structpb.Value

	var var_fe215b82b2b8_err error
	var_fe215b82b2b8_mapped, var_fe215b82b2b8_err = types.ByResourcePropertyType(model.ResourceProperty_INT32).Pack(var_fe215b82b2b8)
	if var_fe215b82b2b8_err != nil {
		panic(var_fe215b82b2b8_err)
	}
	properties["version"] = var_fe215b82b2b8_mapped

	var_aa7e24a89043 := resource.CreatedBy

	if var_aa7e24a89043 != nil {
		var var_aa7e24a89043_mapped *structpb.Value

		var var_aa7e24a89043_err error
		var_aa7e24a89043_mapped, var_aa7e24a89043_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_aa7e24a89043)
		if var_aa7e24a89043_err != nil {
			panic(var_aa7e24a89043_err)
		}
		properties["createdBy"] = var_aa7e24a89043_mapped
	}

	var_9b9c864e64bd := resource.UpdatedBy

	if var_9b9c864e64bd != nil {
		var var_9b9c864e64bd_mapped *structpb.Value

		var var_9b9c864e64bd_err error
		var_9b9c864e64bd_mapped, var_9b9c864e64bd_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_9b9c864e64bd)
		if var_9b9c864e64bd_err != nil {
			panic(var_9b9c864e64bd_err)
		}
		properties["updatedBy"] = var_9b9c864e64bd_mapped
	}

	var_4935b89b4f0e := resource.CreatedOn

	if var_4935b89b4f0e != nil {
		var var_4935b89b4f0e_mapped *structpb.Value

		var var_4935b89b4f0e_err error
		var_4935b89b4f0e_mapped, var_4935b89b4f0e_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_4935b89b4f0e)
		if var_4935b89b4f0e_err != nil {
			panic(var_4935b89b4f0e_err)
		}
		properties["createdOn"] = var_4935b89b4f0e_mapped
	}

	var_c8f84fbefd60 := resource.UpdatedOn

	if var_c8f84fbefd60 != nil {
		var var_c8f84fbefd60_mapped *structpb.Value

		var var_c8f84fbefd60_err error
		var_c8f84fbefd60_mapped, var_c8f84fbefd60_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_c8f84fbefd60)
		if var_c8f84fbefd60_err != nil {
			panic(var_c8f84fbefd60_err)
		}
		properties["updatedOn"] = var_c8f84fbefd60_mapped
	}

	var_825d7c849933 := resource.Name

	var var_825d7c849933_mapped *structpb.Value

	var var_825d7c849933_err error
	var_825d7c849933_mapped, var_825d7c849933_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_825d7c849933)
	if var_825d7c849933_err != nil {
		panic(var_825d7c849933_err)
	}
	properties["name"] = var_825d7c849933_mapped

	var_883984570611 := resource.Namespace

	if var_883984570611 != nil {
		var var_883984570611_mapped *structpb.Value

		var_883984570611_mapped = structpb.NewStructValue(&structpb.Struct{Fields: NamespaceMapperInstance.ToProperties(var_883984570611)})
		properties["namespace"] = var_883984570611_mapped
	}

	var_062ac35e9957 := resource.Virtual

	var var_062ac35e9957_mapped *structpb.Value

	var var_062ac35e9957_err error
	var_062ac35e9957_mapped, var_062ac35e9957_err = types.ByResourcePropertyType(model.ResourceProperty_BOOL).Pack(var_062ac35e9957)
	if var_062ac35e9957_err != nil {
		panic(var_062ac35e9957_err)
	}
	properties["virtual"] = var_062ac35e9957_mapped

	var_888a853e6d2a := resource.Properties

	var var_888a853e6d2a_mapped *structpb.Value

	var var_888a853e6d2a_l []*structpb.Value
	for _, value := range var_888a853e6d2a {

		var_d7b1d9ceab8d := value
		var var_d7b1d9ceab8d_mapped *structpb.Value

		var_d7b1d9ceab8d_mapped = structpb.NewStructValue(&structpb.Struct{Fields: ResourcePropertyMapperInstance.ToProperties(&var_d7b1d9ceab8d)})

		var_888a853e6d2a_l = append(var_888a853e6d2a_l, var_d7b1d9ceab8d_mapped)
	}
	var_888a853e6d2a_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_888a853e6d2a_l})
	properties["properties"] = var_888a853e6d2a_mapped

	var_d9cda76a0a8b := resource.Indexes

	if var_d9cda76a0a8b != nil {
		var var_d9cda76a0a8b_mapped *structpb.Value

		var var_d9cda76a0a8b_l []*structpb.Value
		for _, value := range var_d9cda76a0a8b {

			var_61a20964365b := value
			var var_61a20964365b_mapped *structpb.Value

			var_61a20964365b_mapped = structpb.NewStructValue(&structpb.Struct{Fields: ResourceIndexMapperInstance.ToProperties(&var_61a20964365b)})

			var_d9cda76a0a8b_l = append(var_d9cda76a0a8b_l, var_61a20964365b_mapped)
		}
		var_d9cda76a0a8b_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_d9cda76a0a8b_l})
		properties["indexes"] = var_d9cda76a0a8b_mapped
	}

	var_d4482508e7d7 := resource.Types

	if var_d4482508e7d7 != nil {
		var var_d4482508e7d7_mapped *structpb.Value

		var var_d4482508e7d7_l []*structpb.Value
		for _, value := range var_d4482508e7d7 {

			var_29fab64eecd0 := value
			var var_29fab64eecd0_mapped *structpb.Value

			var_29fab64eecd0_mapped = structpb.NewStructValue(&structpb.Struct{Fields: ResourceSubTypeMapperInstance.ToProperties(&var_29fab64eecd0)})

			var_d4482508e7d7_l = append(var_d4482508e7d7_l, var_29fab64eecd0_mapped)
		}
		var_d4482508e7d7_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_d4482508e7d7_l})
		properties["types"] = var_d4482508e7d7_mapped
	}

	var_7574f581238d := resource.Immutable

	var var_7574f581238d_mapped *structpb.Value

	var var_7574f581238d_err error
	var_7574f581238d_mapped, var_7574f581238d_err = types.ByResourcePropertyType(model.ResourceProperty_BOOL).Pack(var_7574f581238d)
	if var_7574f581238d_err != nil {
		panic(var_7574f581238d_err)
	}
	properties["immutable"] = var_7574f581238d_mapped

	var_193d1dc12fa0 := resource.Abstract

	var var_193d1dc12fa0_mapped *structpb.Value

	var var_193d1dc12fa0_err error
	var_193d1dc12fa0_mapped, var_193d1dc12fa0_err = types.ByResourcePropertyType(model.ResourceProperty_BOOL).Pack(var_193d1dc12fa0)
	if var_193d1dc12fa0_err != nil {
		panic(var_193d1dc12fa0_err)
	}
	properties["abstract"] = var_193d1dc12fa0_mapped

	var_9e19fd48fef9 := resource.DataSource

	if var_9e19fd48fef9 != nil {
		var var_9e19fd48fef9_mapped *structpb.Value

		var_9e19fd48fef9_mapped = structpb.NewStructValue(&structpb.Struct{Fields: DataSourceMapperInstance.ToProperties(var_9e19fd48fef9)})
		properties["dataSource"] = var_9e19fd48fef9_mapped
	}

	var_bfc0d7e1eadf := resource.Entity

	if var_bfc0d7e1eadf != nil {
		var var_bfc0d7e1eadf_mapped *structpb.Value

		var var_bfc0d7e1eadf_err error
		var_bfc0d7e1eadf_mapped, var_bfc0d7e1eadf_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_bfc0d7e1eadf)
		if var_bfc0d7e1eadf_err != nil {
			panic(var_bfc0d7e1eadf_err)
		}
		properties["entity"] = var_bfc0d7e1eadf_mapped
	}

	var_2abc091d8e68 := resource.Catalog

	if var_2abc091d8e68 != nil {
		var var_2abc091d8e68_mapped *structpb.Value

		var var_2abc091d8e68_err error
		var_2abc091d8e68_mapped, var_2abc091d8e68_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_2abc091d8e68)
		if var_2abc091d8e68_err != nil {
			panic(var_2abc091d8e68_err)
		}
		properties["catalog"] = var_2abc091d8e68_mapped
	}

	var_ed64aac48707 := resource.Title

	if var_ed64aac48707 != nil {
		var var_ed64aac48707_mapped *structpb.Value

		var var_ed64aac48707_err error
		var_ed64aac48707_mapped, var_ed64aac48707_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_ed64aac48707)
		if var_ed64aac48707_err != nil {
			panic(var_ed64aac48707_err)
		}
		properties["title"] = var_ed64aac48707_mapped
	}

	var_7459f0417124 := resource.Description

	if var_7459f0417124 != nil {
		var var_7459f0417124_mapped *structpb.Value

		var var_7459f0417124_err error
		var_7459f0417124_mapped, var_7459f0417124_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_7459f0417124)
		if var_7459f0417124_err != nil {
			panic(var_7459f0417124_err)
		}
		properties["description"] = var_7459f0417124_mapped
	}

	var_c76a5f4f605f := resource.Annotations

	if var_c76a5f4f605f != nil {
		var var_c76a5f4f605f_mapped *structpb.Value

		var var_c76a5f4f605f_st *structpb.Struct = new(structpb.Struct)
		var_c76a5f4f605f_st.Fields = make(map[string]*structpb.Value)
		for key, value := range var_c76a5f4f605f {

			var_2879dc1a7de3 := value
			var var_2879dc1a7de3_mapped *structpb.Value

			var var_2879dc1a7de3_err error
			var_2879dc1a7de3_mapped, var_2879dc1a7de3_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_2879dc1a7de3)
			if var_2879dc1a7de3_err != nil {
				panic(var_2879dc1a7de3_err)
			}

			var_c76a5f4f605f_st.Fields[key] = var_2879dc1a7de3_mapped
		}
		var_c76a5f4f605f_mapped = structpb.NewStructValue(var_c76a5f4f605f_st)
		properties["annotations"] = var_c76a5f4f605f_mapped
	}
	return properties
}

func (m *ResourceMapper) FromProperties(properties map[string]*structpb.Value) *Resource {
	var s = m.New()
	if properties["id"] != nil {

		var_824adffaca96 := properties["id"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_UUID).UnPack(var_824adffaca96)

		if err != nil {
			panic(err)
		}

		var_824adffaca96_mapped := new(uuid.UUID)
		*var_824adffaca96_mapped = val.(uuid.UUID)

		s.Id = var_824adffaca96_mapped
	}
	if properties["version"] != nil {

		var_f6c39e1fd273 := properties["version"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(var_f6c39e1fd273)

		if err != nil {
			panic(err)
		}

		var_f6c39e1fd273_mapped := val.(int32)

		s.Version = var_f6c39e1fd273_mapped
	}
	if properties["createdBy"] != nil {

		var_fbf51df80585 := properties["createdBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_fbf51df80585)

		if err != nil {
			panic(err)
		}

		var_fbf51df80585_mapped := new(string)
		*var_fbf51df80585_mapped = val.(string)

		s.CreatedBy = var_fbf51df80585_mapped
	}
	if properties["updatedBy"] != nil {

		var_58f3b81e3086 := properties["updatedBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_58f3b81e3086)

		if err != nil {
			panic(err)
		}

		var_58f3b81e3086_mapped := new(string)
		*var_58f3b81e3086_mapped = val.(string)

		s.UpdatedBy = var_58f3b81e3086_mapped
	}
	if properties["createdOn"] != nil {

		var_12fb0fd0d385 := properties["createdOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_12fb0fd0d385)

		if err != nil {
			panic(err)
		}

		var_12fb0fd0d385_mapped := new(time.Time)
		*var_12fb0fd0d385_mapped = val.(time.Time)

		s.CreatedOn = var_12fb0fd0d385_mapped
	}
	if properties["updatedOn"] != nil {

		var_1aa9fdaafdaa := properties["updatedOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_1aa9fdaafdaa)

		if err != nil {
			panic(err)
		}

		var_1aa9fdaafdaa_mapped := new(time.Time)
		*var_1aa9fdaafdaa_mapped = val.(time.Time)

		s.UpdatedOn = var_1aa9fdaafdaa_mapped
	}
	if properties["name"] != nil {

		var_d746c5b8a6ba := properties["name"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_d746c5b8a6ba)

		if err != nil {
			panic(err)
		}

		var_d746c5b8a6ba_mapped := val.(string)

		s.Name = var_d746c5b8a6ba_mapped
	}
	if properties["namespace"] != nil {

		var_10155f7bd115 := properties["namespace"]
		var_10155f7bd115_mapped := NamespaceMapperInstance.FromProperties(var_10155f7bd115.GetStructValue().Fields)

		s.Namespace = var_10155f7bd115_mapped
	}
	if properties["virtual"] != nil {

		var_4618773d3ada := properties["virtual"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_BOOL).UnPack(var_4618773d3ada)

		if err != nil {
			panic(err)
		}

		var_4618773d3ada_mapped := val.(bool)

		s.Virtual = var_4618773d3ada_mapped
	}
	if properties["properties"] != nil {

		var_0c13835b4f89 := properties["properties"]
		var_0c13835b4f89_mapped := []ResourceProperty{}
		for _, v := range var_0c13835b4f89.GetListValue().Values {

			var_8ad37dacc997 := v
			var mappedValue = ResourcePropertyMapperInstance.FromProperties(var_8ad37dacc997.GetStructValue().Fields)

			var_8ad37dacc997_mapped := *mappedValue

			var_0c13835b4f89_mapped = append(var_0c13835b4f89_mapped, var_8ad37dacc997_mapped)
		}

		s.Properties = var_0c13835b4f89_mapped
	}
	if properties["indexes"] != nil {

		var_97590bef0bd7 := properties["indexes"]
		var_97590bef0bd7_mapped := []ResourceIndex{}
		for _, v := range var_97590bef0bd7.GetListValue().Values {

			var_18684647324f := v
			var mappedValue = ResourceIndexMapperInstance.FromProperties(var_18684647324f.GetStructValue().Fields)

			var_18684647324f_mapped := *mappedValue

			var_97590bef0bd7_mapped = append(var_97590bef0bd7_mapped, var_18684647324f_mapped)
		}

		s.Indexes = var_97590bef0bd7_mapped
	}
	if properties["types"] != nil {

		var_be8a80dd36a6 := properties["types"]
		var_be8a80dd36a6_mapped := []ResourceSubType{}
		for _, v := range var_be8a80dd36a6.GetListValue().Values {

			var_c99509d95928 := v
			var mappedValue = ResourceSubTypeMapperInstance.FromProperties(var_c99509d95928.GetStructValue().Fields)

			var_c99509d95928_mapped := *mappedValue

			var_be8a80dd36a6_mapped = append(var_be8a80dd36a6_mapped, var_c99509d95928_mapped)
		}

		s.Types = var_be8a80dd36a6_mapped
	}
	if properties["immutable"] != nil {

		var_e47bbca523c9 := properties["immutable"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_BOOL).UnPack(var_e47bbca523c9)

		if err != nil {
			panic(err)
		}

		var_e47bbca523c9_mapped := val.(bool)

		s.Immutable = var_e47bbca523c9_mapped
	}
	if properties["abstract"] != nil {

		var_e2df5a8a6d71 := properties["abstract"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_BOOL).UnPack(var_e2df5a8a6d71)

		if err != nil {
			panic(err)
		}

		var_e2df5a8a6d71_mapped := val.(bool)

		s.Abstract = var_e2df5a8a6d71_mapped
	}
	if properties["dataSource"] != nil {

		var_d8312d9d0b14 := properties["dataSource"]
		var_d8312d9d0b14_mapped := DataSourceMapperInstance.FromProperties(var_d8312d9d0b14.GetStructValue().Fields)

		s.DataSource = var_d8312d9d0b14_mapped
	}
	if properties["entity"] != nil {

		var_8a748eefe367 := properties["entity"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_8a748eefe367)

		if err != nil {
			panic(err)
		}

		var_8a748eefe367_mapped := new(string)
		*var_8a748eefe367_mapped = val.(string)

		s.Entity = var_8a748eefe367_mapped
	}
	if properties["catalog"] != nil {

		var_1d7aa68134a4 := properties["catalog"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_1d7aa68134a4)

		if err != nil {
			panic(err)
		}

		var_1d7aa68134a4_mapped := new(string)
		*var_1d7aa68134a4_mapped = val.(string)

		s.Catalog = var_1d7aa68134a4_mapped
	}
	if properties["title"] != nil {

		var_4a95d78a63d8 := properties["title"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_4a95d78a63d8)

		if err != nil {
			panic(err)
		}

		var_4a95d78a63d8_mapped := new(string)
		*var_4a95d78a63d8_mapped = val.(string)

		s.Title = var_4a95d78a63d8_mapped
	}
	if properties["description"] != nil {

		var_bdfe0eed6e75 := properties["description"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_bdfe0eed6e75)

		if err != nil {
			panic(err)
		}

		var_bdfe0eed6e75_mapped := new(string)
		*var_bdfe0eed6e75_mapped = val.(string)

		s.Description = var_bdfe0eed6e75_mapped
	}
	if properties["annotations"] != nil {

		var_9a96bbb4150b := properties["annotations"]
		var_9a96bbb4150b_mapped := make(map[string]string)
		for k, v := range var_9a96bbb4150b.GetStructValue().Fields {

			var_97a320e3fd65 := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_97a320e3fd65)

			if err != nil {
				panic(err)
			}

			var_97a320e3fd65_mapped := val.(string)

			var_9a96bbb4150b_mapped[k] = var_97a320e3fd65_mapped
		}

		s.Annotations = var_9a96bbb4150b_mapped
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

	var_af5daaec2ef3 := resourceProperty.Name

	var var_af5daaec2ef3_mapped *structpb.Value

	var var_af5daaec2ef3_err error
	var_af5daaec2ef3_mapped, var_af5daaec2ef3_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_af5daaec2ef3)
	if var_af5daaec2ef3_err != nil {
		panic(var_af5daaec2ef3_err)
	}
	properties["name"] = var_af5daaec2ef3_mapped

	var_29ae28e9f648 := resourceProperty.Type

	var var_29ae28e9f648_mapped *structpb.Value

	var var_29ae28e9f648_err error
	var_29ae28e9f648_mapped, var_29ae28e9f648_err = types.ByResourcePropertyType(model.ResourceProperty_INT32).Pack(var_29ae28e9f648)
	if var_29ae28e9f648_err != nil {
		panic(var_29ae28e9f648_err)
	}
	properties["type"] = var_29ae28e9f648_mapped

	var_9933229e0fab := resourceProperty.TypeRef

	if var_9933229e0fab != nil {
		var var_9933229e0fab_mapped *structpb.Value

		var var_9933229e0fab_err error
		var_9933229e0fab_mapped, var_9933229e0fab_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_9933229e0fab)
		if var_9933229e0fab_err != nil {
			panic(var_9933229e0fab_err)
		}
		properties["typeRef"] = var_9933229e0fab_mapped
	}

	var_763e88c4b8be := resourceProperty.Mapping

	var var_763e88c4b8be_mapped *structpb.Value

	var var_763e88c4b8be_err error
	var_763e88c4b8be_mapped, var_763e88c4b8be_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_763e88c4b8be)
	if var_763e88c4b8be_err != nil {
		panic(var_763e88c4b8be_err)
	}
	properties["mapping"] = var_763e88c4b8be_mapped

	var_8aafd49ca703 := resourceProperty.Primary

	var var_8aafd49ca703_mapped *structpb.Value

	var var_8aafd49ca703_err error
	var_8aafd49ca703_mapped, var_8aafd49ca703_err = types.ByResourcePropertyType(model.ResourceProperty_BOOL).Pack(var_8aafd49ca703)
	if var_8aafd49ca703_err != nil {
		panic(var_8aafd49ca703_err)
	}
	properties["primary"] = var_8aafd49ca703_mapped

	var_b4ddfaadec59 := resourceProperty.Required

	var var_b4ddfaadec59_mapped *structpb.Value

	var var_b4ddfaadec59_err error
	var_b4ddfaadec59_mapped, var_b4ddfaadec59_err = types.ByResourcePropertyType(model.ResourceProperty_BOOL).Pack(var_b4ddfaadec59)
	if var_b4ddfaadec59_err != nil {
		panic(var_b4ddfaadec59_err)
	}
	properties["required"] = var_b4ddfaadec59_mapped

	var_ccf00394d890 := resourceProperty.Unique

	var var_ccf00394d890_mapped *structpb.Value

	var var_ccf00394d890_err error
	var_ccf00394d890_mapped, var_ccf00394d890_err = types.ByResourcePropertyType(model.ResourceProperty_BOOL).Pack(var_ccf00394d890)
	if var_ccf00394d890_err != nil {
		panic(var_ccf00394d890_err)
	}
	properties["unique"] = var_ccf00394d890_mapped

	var_4d46bc17d01e := resourceProperty.Immutable

	var var_4d46bc17d01e_mapped *structpb.Value

	var var_4d46bc17d01e_err error
	var_4d46bc17d01e_mapped, var_4d46bc17d01e_err = types.ByResourcePropertyType(model.ResourceProperty_BOOL).Pack(var_4d46bc17d01e)
	if var_4d46bc17d01e_err != nil {
		panic(var_4d46bc17d01e_err)
	}
	properties["immutable"] = var_4d46bc17d01e_mapped

	var_c0116208be93 := resourceProperty.Length

	var var_c0116208be93_mapped *structpb.Value

	var var_c0116208be93_err error
	var_c0116208be93_mapped, var_c0116208be93_err = types.ByResourcePropertyType(model.ResourceProperty_INT32).Pack(var_c0116208be93)
	if var_c0116208be93_err != nil {
		panic(var_c0116208be93_err)
	}
	properties["length"] = var_c0116208be93_mapped

	var_c83f2980ab91 := resourceProperty.Item

	if var_c83f2980ab91 != nil {
		var var_c83f2980ab91_mapped *structpb.Value

		var_c83f2980ab91_mapped = structpb.NewStructValue(&structpb.Struct{Fields: ResourcePropertyMapperInstance.ToProperties(var_c83f2980ab91)})
		properties["item"] = var_c83f2980ab91_mapped
	}

	var_0e172fe87de1 := resourceProperty.Properties

	var var_0e172fe87de1_mapped *structpb.Value

	var var_0e172fe87de1_l []*structpb.Value
	for _, value := range var_0e172fe87de1 {

		var_46d046fa188e := value
		var var_46d046fa188e_mapped *structpb.Value

		var_46d046fa188e_mapped = structpb.NewStructValue(&structpb.Struct{Fields: ResourcePropertyMapperInstance.ToProperties(&var_46d046fa188e)})

		var_0e172fe87de1_l = append(var_0e172fe87de1_l, var_46d046fa188e_mapped)
	}
	var_0e172fe87de1_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_0e172fe87de1_l})
	properties["properties"] = var_0e172fe87de1_mapped

	var_62c133854198 := resourceProperty.Reference

	if var_62c133854198 != nil {
		var var_62c133854198_mapped *structpb.Value

		var_62c133854198_mapped = structpb.NewStructValue(&structpb.Struct{Fields: ResourceReferenceMapperInstance.ToProperties(var_62c133854198)})
		properties["reference"] = var_62c133854198_mapped
	}

	var_73dc2d7cf476 := resourceProperty.DefaultValue

	if var_73dc2d7cf476 != nil {
		var var_73dc2d7cf476_mapped *structpb.Value

		var var_73dc2d7cf476_err error
		var_73dc2d7cf476_mapped, var_73dc2d7cf476_err = types.ByResourcePropertyType(model.ResourceProperty_OBJECT).Pack(*var_73dc2d7cf476)
		if var_73dc2d7cf476_err != nil {
			panic(var_73dc2d7cf476_err)
		}
		properties["defaultValue"] = var_73dc2d7cf476_mapped
	}

	var_e645da0b24ad := resourceProperty.EnumValues

	if var_e645da0b24ad != nil {
		var var_e645da0b24ad_mapped *structpb.Value

		var var_e645da0b24ad_l []*structpb.Value
		for _, value := range var_e645da0b24ad {

			var_9f1dd3d4cf7f := value
			var var_9f1dd3d4cf7f_mapped *structpb.Value

			var var_9f1dd3d4cf7f_err error
			var_9f1dd3d4cf7f_mapped, var_9f1dd3d4cf7f_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_9f1dd3d4cf7f)
			if var_9f1dd3d4cf7f_err != nil {
				panic(var_9f1dd3d4cf7f_err)
			}

			var_e645da0b24ad_l = append(var_e645da0b24ad_l, var_9f1dd3d4cf7f_mapped)
		}
		var_e645da0b24ad_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_e645da0b24ad_l})
		properties["enumValues"] = var_e645da0b24ad_mapped
	}

	var_8f54b83fac17 := resourceProperty.ExampleValue

	if var_8f54b83fac17 != nil {
		var var_8f54b83fac17_mapped *structpb.Value

		var var_8f54b83fac17_err error
		var_8f54b83fac17_mapped, var_8f54b83fac17_err = types.ByResourcePropertyType(model.ResourceProperty_OBJECT).Pack(*var_8f54b83fac17)
		if var_8f54b83fac17_err != nil {
			panic(var_8f54b83fac17_err)
		}
		properties["exampleValue"] = var_8f54b83fac17_mapped
	}

	var_5dcd3c518ab7 := resourceProperty.Title

	if var_5dcd3c518ab7 != nil {
		var var_5dcd3c518ab7_mapped *structpb.Value

		var var_5dcd3c518ab7_err error
		var_5dcd3c518ab7_mapped, var_5dcd3c518ab7_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_5dcd3c518ab7)
		if var_5dcd3c518ab7_err != nil {
			panic(var_5dcd3c518ab7_err)
		}
		properties["title"] = var_5dcd3c518ab7_mapped
	}

	var_8239352f5f83 := resourceProperty.Description

	if var_8239352f5f83 != nil {
		var var_8239352f5f83_mapped *structpb.Value

		var var_8239352f5f83_err error
		var_8239352f5f83_mapped, var_8239352f5f83_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_8239352f5f83)
		if var_8239352f5f83_err != nil {
			panic(var_8239352f5f83_err)
		}
		properties["description"] = var_8239352f5f83_mapped
	}

	var_bb237df176f1 := resourceProperty.Annotations

	if var_bb237df176f1 != nil {
		var var_bb237df176f1_mapped *structpb.Value

		var var_bb237df176f1_st *structpb.Struct = new(structpb.Struct)
		var_bb237df176f1_st.Fields = make(map[string]*structpb.Value)
		for key, value := range var_bb237df176f1 {

			var_edf19d099459 := value
			var var_edf19d099459_mapped *structpb.Value

			var var_edf19d099459_err error
			var_edf19d099459_mapped, var_edf19d099459_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_edf19d099459)
			if var_edf19d099459_err != nil {
				panic(var_edf19d099459_err)
			}

			var_bb237df176f1_st.Fields[key] = var_edf19d099459_mapped
		}
		var_bb237df176f1_mapped = structpb.NewStructValue(var_bb237df176f1_st)
		properties["annotations"] = var_bb237df176f1_mapped
	}
	return properties
}

func (m *ResourcePropertyMapper) FromProperties(properties map[string]*structpb.Value) *ResourceProperty {
	var s = m.New()
	if properties["name"] != nil {

		var_b252afb359ef := properties["name"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_b252afb359ef)

		if err != nil {
			panic(err)
		}

		var_b252afb359ef_mapped := val.(string)

		s.Name = var_b252afb359ef_mapped
	}
	if properties["type"] != nil {

		var_2b918b6e964f := properties["type"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(var_2b918b6e964f)

		if err != nil {
			panic(err)
		}

		var_2b918b6e964f_mapped := val.(int32)

		s.Type = var_2b918b6e964f_mapped
	}
	if properties["typeRef"] != nil {

		var_0ccb4395e619 := properties["typeRef"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_0ccb4395e619)

		if err != nil {
			panic(err)
		}

		var_0ccb4395e619_mapped := new(string)
		*var_0ccb4395e619_mapped = val.(string)

		s.TypeRef = var_0ccb4395e619_mapped
	}
	if properties["mapping"] != nil {

		var_2b1f605851c7 := properties["mapping"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_2b1f605851c7)

		if err != nil {
			panic(err)
		}

		var_2b1f605851c7_mapped := val.(string)

		s.Mapping = var_2b1f605851c7_mapped
	}
	if properties["primary"] != nil {

		var_9055f393182c := properties["primary"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_BOOL).UnPack(var_9055f393182c)

		if err != nil {
			panic(err)
		}

		var_9055f393182c_mapped := val.(bool)

		s.Primary = var_9055f393182c_mapped
	}
	if properties["required"] != nil {

		var_969519d41ad1 := properties["required"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_BOOL).UnPack(var_969519d41ad1)

		if err != nil {
			panic(err)
		}

		var_969519d41ad1_mapped := val.(bool)

		s.Required = var_969519d41ad1_mapped
	}
	if properties["unique"] != nil {

		var_66d44823b618 := properties["unique"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_BOOL).UnPack(var_66d44823b618)

		if err != nil {
			panic(err)
		}

		var_66d44823b618_mapped := val.(bool)

		s.Unique = var_66d44823b618_mapped
	}
	if properties["immutable"] != nil {

		var_b25493df7614 := properties["immutable"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_BOOL).UnPack(var_b25493df7614)

		if err != nil {
			panic(err)
		}

		var_b25493df7614_mapped := val.(bool)

		s.Immutable = var_b25493df7614_mapped
	}
	if properties["length"] != nil {

		var_ed5f0a063814 := properties["length"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(var_ed5f0a063814)

		if err != nil {
			panic(err)
		}

		var_ed5f0a063814_mapped := val.(int32)

		s.Length = var_ed5f0a063814_mapped
	}
	if properties["item"] != nil {

		var_7f2135dc905f := properties["item"]
		var mappedValue = ResourcePropertyMapperInstance.FromProperties(var_7f2135dc905f.GetStructValue().Fields)

		var_7f2135dc905f_mapped := mappedValue

		s.Item = var_7f2135dc905f_mapped
	}
	if properties["properties"] != nil {

		var_534d8463c7e9 := properties["properties"]
		var_534d8463c7e9_mapped := []ResourceProperty{}
		for _, v := range var_534d8463c7e9.GetListValue().Values {

			var_3234e56b1fd5 := v
			var mappedValue = ResourcePropertyMapperInstance.FromProperties(var_3234e56b1fd5.GetStructValue().Fields)

			var_3234e56b1fd5_mapped := *mappedValue

			var_534d8463c7e9_mapped = append(var_534d8463c7e9_mapped, var_3234e56b1fd5_mapped)
		}

		s.Properties = var_534d8463c7e9_mapped
	}
	if properties["reference"] != nil {

		var_7cb87c70c600 := properties["reference"]
		var mappedValue = ResourceReferenceMapperInstance.FromProperties(var_7cb87c70c600.GetStructValue().Fields)

		var_7cb87c70c600_mapped := mappedValue

		s.Reference = var_7cb87c70c600_mapped
	}
	if properties["defaultValue"] != nil {

		var_fdc0c8cba409 := properties["defaultValue"]
		var_fdc0c8cba409_mapped := new(unstructured.Unstructured)
		*var_fdc0c8cba409_mapped = unstructured.FromStructValue(var_fdc0c8cba409.GetStructValue())

		s.DefaultValue = var_fdc0c8cba409_mapped
	}
	if properties["enumValues"] != nil {

		var_c2f644edb452 := properties["enumValues"]
		var_c2f644edb452_mapped := []string{}
		for _, v := range var_c2f644edb452.GetListValue().Values {

			var_1474d455ab83 := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_1474d455ab83)

			if err != nil {
				panic(err)
			}

			var_1474d455ab83_mapped := val.(string)

			var_c2f644edb452_mapped = append(var_c2f644edb452_mapped, var_1474d455ab83_mapped)
		}

		s.EnumValues = var_c2f644edb452_mapped
	}
	if properties["exampleValue"] != nil {

		var_8af6fb7c1eab := properties["exampleValue"]
		var_8af6fb7c1eab_mapped := new(unstructured.Unstructured)
		*var_8af6fb7c1eab_mapped = unstructured.FromStructValue(var_8af6fb7c1eab.GetStructValue())

		s.ExampleValue = var_8af6fb7c1eab_mapped
	}
	if properties["title"] != nil {

		var_4f731e16af61 := properties["title"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_4f731e16af61)

		if err != nil {
			panic(err)
		}

		var_4f731e16af61_mapped := new(string)
		*var_4f731e16af61_mapped = val.(string)

		s.Title = var_4f731e16af61_mapped
	}
	if properties["description"] != nil {

		var_484645e3fb4a := properties["description"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_484645e3fb4a)

		if err != nil {
			panic(err)
		}

		var_484645e3fb4a_mapped := new(string)
		*var_484645e3fb4a_mapped = val.(string)

		s.Description = var_484645e3fb4a_mapped
	}
	if properties["annotations"] != nil {

		var_c80df6425f70 := properties["annotations"]
		var_c80df6425f70_mapped := make(map[string]string)
		for k, v := range var_c80df6425f70.GetStructValue().Fields {

			var_3ab11fdc4781 := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_3ab11fdc4781)

			if err != nil {
				panic(err)
			}

			var_3ab11fdc4781_mapped := val.(string)

			var_c80df6425f70_mapped[k] = var_3ab11fdc4781_mapped
		}

		s.Annotations = var_c80df6425f70_mapped
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

	var_7c10a9dcac25 := resourceSubType.Name

	var var_7c10a9dcac25_mapped *structpb.Value

	var var_7c10a9dcac25_err error
	var_7c10a9dcac25_mapped, var_7c10a9dcac25_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_7c10a9dcac25)
	if var_7c10a9dcac25_err != nil {
		panic(var_7c10a9dcac25_err)
	}
	properties["name"] = var_7c10a9dcac25_mapped

	var_2f303aa8d0b2 := resourceSubType.Properties

	var var_2f303aa8d0b2_mapped *structpb.Value

	var var_2f303aa8d0b2_l []*structpb.Value
	for _, value := range var_2f303aa8d0b2 {

		var_7558f96280cc := value
		var var_7558f96280cc_mapped *structpb.Value

		var_7558f96280cc_mapped = structpb.NewStructValue(&structpb.Struct{Fields: ResourcePropertyMapperInstance.ToProperties(&var_7558f96280cc)})

		var_2f303aa8d0b2_l = append(var_2f303aa8d0b2_l, var_7558f96280cc_mapped)
	}
	var_2f303aa8d0b2_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_2f303aa8d0b2_l})
	properties["properties"] = var_2f303aa8d0b2_mapped
	return properties
}

func (m *ResourceSubTypeMapper) FromProperties(properties map[string]*structpb.Value) *ResourceSubType {
	var s = m.New()
	if properties["name"] != nil {

		var_0ed15357b143 := properties["name"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_0ed15357b143)

		if err != nil {
			panic(err)
		}

		var_0ed15357b143_mapped := val.(string)

		s.Name = var_0ed15357b143_mapped
	}
	if properties["properties"] != nil {

		var_4320ff4b416d := properties["properties"]
		var_4320ff4b416d_mapped := []ResourceProperty{}
		for _, v := range var_4320ff4b416d.GetListValue().Values {

			var_f72ac19ab986 := v
			var mappedValue = ResourcePropertyMapperInstance.FromProperties(var_f72ac19ab986.GetStructValue().Fields)

			var_f72ac19ab986_mapped := *mappedValue

			var_4320ff4b416d_mapped = append(var_4320ff4b416d_mapped, var_f72ac19ab986_mapped)
		}

		s.Properties = var_4320ff4b416d_mapped
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

	var_50b5731d66a6 := resourceIndexProperty.Name

	var var_50b5731d66a6_mapped *structpb.Value

	var var_50b5731d66a6_err error
	var_50b5731d66a6_mapped, var_50b5731d66a6_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_50b5731d66a6)
	if var_50b5731d66a6_err != nil {
		panic(var_50b5731d66a6_err)
	}
	properties["name"] = var_50b5731d66a6_mapped

	var_35e78cce7ebe := resourceIndexProperty.Order

	if var_35e78cce7ebe != nil {
		var var_35e78cce7ebe_mapped *structpb.Value

		var var_35e78cce7ebe_err error
		var_35e78cce7ebe_mapped, var_35e78cce7ebe_err = types.ByResourcePropertyType(model.ResourceProperty_ENUM).Pack(string(*var_35e78cce7ebe))
		if var_35e78cce7ebe_err != nil {
			panic(var_35e78cce7ebe_err)
		}
		properties["order"] = var_35e78cce7ebe_mapped
	}
	return properties
}

func (m *ResourceIndexPropertyMapper) FromProperties(properties map[string]*structpb.Value) *ResourceIndexProperty {
	var s = m.New()
	if properties["name"] != nil {

		var_47d8e74f4df1 := properties["name"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_47d8e74f4df1)

		if err != nil {
			panic(err)
		}

		var_47d8e74f4df1_mapped := val.(string)

		s.Name = var_47d8e74f4df1_mapped
	}
	if properties["order"] != nil {

		var_eb97acf80a12 := properties["order"]
		var_eb97acf80a12_mapped := new(ResourceOrder)
		*var_eb97acf80a12_mapped = (ResourceOrder)(var_eb97acf80a12.GetStringValue())

		s.Order = var_eb97acf80a12_mapped
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

	var_fa6257245abd := resourceIndex.Properties

	if var_fa6257245abd != nil {
		var var_fa6257245abd_mapped *structpb.Value

		var var_fa6257245abd_l []*structpb.Value
		for _, value := range var_fa6257245abd {

			var_704cfa1f4009 := value
			var var_704cfa1f4009_mapped *structpb.Value

			var_704cfa1f4009_mapped = structpb.NewStructValue(&structpb.Struct{Fields: ResourceIndexPropertyMapperInstance.ToProperties(&var_704cfa1f4009)})

			var_fa6257245abd_l = append(var_fa6257245abd_l, var_704cfa1f4009_mapped)
		}
		var_fa6257245abd_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_fa6257245abd_l})
		properties["properties"] = var_fa6257245abd_mapped
	}

	var_52fa0748cacf := resourceIndex.IndexType

	if var_52fa0748cacf != nil {
		var var_52fa0748cacf_mapped *structpb.Value

		var var_52fa0748cacf_err error
		var_52fa0748cacf_mapped, var_52fa0748cacf_err = types.ByResourcePropertyType(model.ResourceProperty_ENUM).Pack(string(*var_52fa0748cacf))
		if var_52fa0748cacf_err != nil {
			panic(var_52fa0748cacf_err)
		}
		properties["indexType"] = var_52fa0748cacf_mapped
	}

	var_7021544b464f := resourceIndex.Unique

	if var_7021544b464f != nil {
		var var_7021544b464f_mapped *structpb.Value

		var var_7021544b464f_err error
		var_7021544b464f_mapped, var_7021544b464f_err = types.ByResourcePropertyType(model.ResourceProperty_BOOL).Pack(*var_7021544b464f)
		if var_7021544b464f_err != nil {
			panic(var_7021544b464f_err)
		}
		properties["unique"] = var_7021544b464f_mapped
	}

	var_d8ca21f30d41 := resourceIndex.Annotations

	if var_d8ca21f30d41 != nil {
		var var_d8ca21f30d41_mapped *structpb.Value

		var var_d8ca21f30d41_st *structpb.Struct = new(structpb.Struct)
		var_d8ca21f30d41_st.Fields = make(map[string]*structpb.Value)
		for key, value := range var_d8ca21f30d41 {

			var_9075a6f05b66 := value
			var var_9075a6f05b66_mapped *structpb.Value

			var var_9075a6f05b66_err error
			var_9075a6f05b66_mapped, var_9075a6f05b66_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_9075a6f05b66)
			if var_9075a6f05b66_err != nil {
				panic(var_9075a6f05b66_err)
			}

			var_d8ca21f30d41_st.Fields[key] = var_9075a6f05b66_mapped
		}
		var_d8ca21f30d41_mapped = structpb.NewStructValue(var_d8ca21f30d41_st)
		properties["annotations"] = var_d8ca21f30d41_mapped
	}
	return properties
}

func (m *ResourceIndexMapper) FromProperties(properties map[string]*structpb.Value) *ResourceIndex {
	var s = m.New()
	if properties["properties"] != nil {

		var_a2d0f65ef5e3 := properties["properties"]
		var_a2d0f65ef5e3_mapped := []ResourceIndexProperty{}
		for _, v := range var_a2d0f65ef5e3.GetListValue().Values {

			var_d6c7c55fd594 := v
			var mappedValue = ResourceIndexPropertyMapperInstance.FromProperties(var_d6c7c55fd594.GetStructValue().Fields)

			var_d6c7c55fd594_mapped := *mappedValue

			var_a2d0f65ef5e3_mapped = append(var_a2d0f65ef5e3_mapped, var_d6c7c55fd594_mapped)
		}

		s.Properties = var_a2d0f65ef5e3_mapped
	}
	if properties["indexType"] != nil {

		var_4200635f7334 := properties["indexType"]
		var_4200635f7334_mapped := new(ResourceIndexType)
		*var_4200635f7334_mapped = (ResourceIndexType)(var_4200635f7334.GetStringValue())

		s.IndexType = var_4200635f7334_mapped
	}
	if properties["unique"] != nil {

		var_28c29c72f959 := properties["unique"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_BOOL).UnPack(var_28c29c72f959)

		if err != nil {
			panic(err)
		}

		var_28c29c72f959_mapped := new(bool)
		*var_28c29c72f959_mapped = val.(bool)

		s.Unique = var_28c29c72f959_mapped
	}
	if properties["annotations"] != nil {

		var_30ffd1b33135 := properties["annotations"]
		var_30ffd1b33135_mapped := make(map[string]string)
		for k, v := range var_30ffd1b33135.GetStructValue().Fields {

			var_acf9e87ff8d6 := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_acf9e87ff8d6)

			if err != nil {
				panic(err)
			}

			var_acf9e87ff8d6_mapped := val.(string)

			var_30ffd1b33135_mapped[k] = var_acf9e87ff8d6_mapped
		}

		s.Annotations = var_30ffd1b33135_mapped
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

	var_10710784a099 := resourceReference.Resource

	if var_10710784a099 != nil {
		var var_10710784a099_mapped *structpb.Value

		var_10710784a099_mapped = structpb.NewStructValue(&structpb.Struct{Fields: ResourceMapperInstance.ToProperties(var_10710784a099)})
		properties["resource"] = var_10710784a099_mapped
	}

	var_2097d9d2d9a8 := resourceReference.Cascade

	if var_2097d9d2d9a8 != nil {
		var var_2097d9d2d9a8_mapped *structpb.Value

		var var_2097d9d2d9a8_err error
		var_2097d9d2d9a8_mapped, var_2097d9d2d9a8_err = types.ByResourcePropertyType(model.ResourceProperty_BOOL).Pack(*var_2097d9d2d9a8)
		if var_2097d9d2d9a8_err != nil {
			panic(var_2097d9d2d9a8_err)
		}
		properties["cascade"] = var_2097d9d2d9a8_mapped
	}

	var_dab32c0ec9c8 := resourceReference.BackReference

	if var_dab32c0ec9c8 != nil {
		var var_dab32c0ec9c8_mapped *structpb.Value

		var var_dab32c0ec9c8_err error
		var_dab32c0ec9c8_mapped, var_dab32c0ec9c8_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_dab32c0ec9c8)
		if var_dab32c0ec9c8_err != nil {
			panic(var_dab32c0ec9c8_err)
		}
		properties["backReference"] = var_dab32c0ec9c8_mapped
	}
	return properties
}

func (m *ResourceReferenceMapper) FromProperties(properties map[string]*structpb.Value) *ResourceReference {
	var s = m.New()
	if properties["resource"] != nil {

		var_98e9ee283718 := properties["resource"]
		var_98e9ee283718_mapped := ResourceMapperInstance.FromProperties(var_98e9ee283718.GetStructValue().Fields)

		s.Resource = var_98e9ee283718_mapped
	}
	if properties["cascade"] != nil {

		var_a2f6c630e374 := properties["cascade"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_BOOL).UnPack(var_a2f6c630e374)

		if err != nil {
			panic(err)
		}

		var_a2f6c630e374_mapped := new(bool)
		*var_a2f6c630e374_mapped = val.(bool)

		s.Cascade = var_a2f6c630e374_mapped
	}
	if properties["backReference"] != nil {

		var_774894acbda3 := properties["backReference"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_774894acbda3)

		if err != nil {
			panic(err)
		}

		var_774894acbda3_mapped := new(string)
		*var_774894acbda3_mapped = val.(string)

		s.BackReference = var_774894acbda3_mapped
	}
	return s
}
