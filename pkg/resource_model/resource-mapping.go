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

	var_25adc3864143 := resource.Id

	if var_25adc3864143 != nil {
		var var_25adc3864143_mapped *structpb.Value

		var var_25adc3864143_err error
		var_25adc3864143_mapped, var_25adc3864143_err = types.ByResourcePropertyType(model.ResourceProperty_UUID).Pack(*var_25adc3864143)
		if var_25adc3864143_err != nil {
			panic(var_25adc3864143_err)
		}
		properties["id"] = var_25adc3864143_mapped
	}

	var_43ab622aaf9d := resource.Version

	var var_43ab622aaf9d_mapped *structpb.Value

	var var_43ab622aaf9d_err error
	var_43ab622aaf9d_mapped, var_43ab622aaf9d_err = types.ByResourcePropertyType(model.ResourceProperty_INT32).Pack(var_43ab622aaf9d)
	if var_43ab622aaf9d_err != nil {
		panic(var_43ab622aaf9d_err)
	}
	properties["version"] = var_43ab622aaf9d_mapped

	var_767b366b768d := resource.CreatedBy

	if var_767b366b768d != nil {
		var var_767b366b768d_mapped *structpb.Value

		var var_767b366b768d_err error
		var_767b366b768d_mapped, var_767b366b768d_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_767b366b768d)
		if var_767b366b768d_err != nil {
			panic(var_767b366b768d_err)
		}
		properties["createdBy"] = var_767b366b768d_mapped
	}

	var_fa33827f40b8 := resource.UpdatedBy

	if var_fa33827f40b8 != nil {
		var var_fa33827f40b8_mapped *structpb.Value

		var var_fa33827f40b8_err error
		var_fa33827f40b8_mapped, var_fa33827f40b8_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_fa33827f40b8)
		if var_fa33827f40b8_err != nil {
			panic(var_fa33827f40b8_err)
		}
		properties["updatedBy"] = var_fa33827f40b8_mapped
	}

	var_9ad12a483fdc := resource.CreatedOn

	if var_9ad12a483fdc != nil {
		var var_9ad12a483fdc_mapped *structpb.Value

		var var_9ad12a483fdc_err error
		var_9ad12a483fdc_mapped, var_9ad12a483fdc_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_9ad12a483fdc)
		if var_9ad12a483fdc_err != nil {
			panic(var_9ad12a483fdc_err)
		}
		properties["createdOn"] = var_9ad12a483fdc_mapped
	}

	var_65f7f9782495 := resource.UpdatedOn

	if var_65f7f9782495 != nil {
		var var_65f7f9782495_mapped *structpb.Value

		var var_65f7f9782495_err error
		var_65f7f9782495_mapped, var_65f7f9782495_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_65f7f9782495)
		if var_65f7f9782495_err != nil {
			panic(var_65f7f9782495_err)
		}
		properties["updatedOn"] = var_65f7f9782495_mapped
	}

	var_3a372071a8a3 := resource.Name

	var var_3a372071a8a3_mapped *structpb.Value

	var var_3a372071a8a3_err error
	var_3a372071a8a3_mapped, var_3a372071a8a3_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_3a372071a8a3)
	if var_3a372071a8a3_err != nil {
		panic(var_3a372071a8a3_err)
	}
	properties["name"] = var_3a372071a8a3_mapped

	var_816d2695481c := resource.Namespace

	if var_816d2695481c != nil {
		var var_816d2695481c_mapped *structpb.Value

		var_816d2695481c_mapped = structpb.NewStructValue(&structpb.Struct{Fields: NamespaceMapperInstance.ToProperties(var_816d2695481c)})
		properties["namespace"] = var_816d2695481c_mapped
	}

	var_1d0134d89e7c := resource.Virtual

	var var_1d0134d89e7c_mapped *structpb.Value

	var var_1d0134d89e7c_err error
	var_1d0134d89e7c_mapped, var_1d0134d89e7c_err = types.ByResourcePropertyType(model.ResourceProperty_BOOL).Pack(var_1d0134d89e7c)
	if var_1d0134d89e7c_err != nil {
		panic(var_1d0134d89e7c_err)
	}
	properties["virtual"] = var_1d0134d89e7c_mapped

	var_b57fe8049118 := resource.Properties

	var var_b57fe8049118_mapped *structpb.Value

	var var_b57fe8049118_l []*structpb.Value
	for _, value := range var_b57fe8049118 {

		var_c7ddfaa7f32a := value
		var var_c7ddfaa7f32a_mapped *structpb.Value

		var_c7ddfaa7f32a_mapped = structpb.NewStructValue(&structpb.Struct{Fields: ResourcePropertyMapperInstance.ToProperties(&var_c7ddfaa7f32a)})

		var_b57fe8049118_l = append(var_b57fe8049118_l, var_c7ddfaa7f32a_mapped)
	}
	var_b57fe8049118_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_b57fe8049118_l})
	properties["properties"] = var_b57fe8049118_mapped

	var_5e66779d5b51 := resource.Types

	if var_5e66779d5b51 != nil {
		var var_5e66779d5b51_mapped *structpb.Value

		var var_5e66779d5b51_l []*structpb.Value
		for _, value := range var_5e66779d5b51 {

			var_9e2172fbf41c := value
			var var_9e2172fbf41c_mapped *structpb.Value

			var_9e2172fbf41c_mapped = structpb.NewStructValue(&structpb.Struct{Fields: ResourceSubTypeMapperInstance.ToProperties(&var_9e2172fbf41c)})

			var_5e66779d5b51_l = append(var_5e66779d5b51_l, var_9e2172fbf41c_mapped)
		}
		var_5e66779d5b51_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_5e66779d5b51_l})
		properties["types"] = var_5e66779d5b51_mapped
	}

	var_fc97210d1244 := resource.Immutable

	var var_fc97210d1244_mapped *structpb.Value

	var var_fc97210d1244_err error
	var_fc97210d1244_mapped, var_fc97210d1244_err = types.ByResourcePropertyType(model.ResourceProperty_BOOL).Pack(var_fc97210d1244)
	if var_fc97210d1244_err != nil {
		panic(var_fc97210d1244_err)
	}
	properties["immutable"] = var_fc97210d1244_mapped

	var_ba3cae10e1b1 := resource.Abstract

	var var_ba3cae10e1b1_mapped *structpb.Value

	var var_ba3cae10e1b1_err error
	var_ba3cae10e1b1_mapped, var_ba3cae10e1b1_err = types.ByResourcePropertyType(model.ResourceProperty_BOOL).Pack(var_ba3cae10e1b1)
	if var_ba3cae10e1b1_err != nil {
		panic(var_ba3cae10e1b1_err)
	}
	properties["abstract"] = var_ba3cae10e1b1_mapped

	var_47ef3c16ddb2 := resource.DataSource

	if var_47ef3c16ddb2 != nil {
		var var_47ef3c16ddb2_mapped *structpb.Value

		var_47ef3c16ddb2_mapped = structpb.NewStructValue(&structpb.Struct{Fields: DataSourceMapperInstance.ToProperties(var_47ef3c16ddb2)})
		properties["dataSource"] = var_47ef3c16ddb2_mapped
	}

	var_7f4a66e6f59b := resource.Entity

	if var_7f4a66e6f59b != nil {
		var var_7f4a66e6f59b_mapped *structpb.Value

		var var_7f4a66e6f59b_err error
		var_7f4a66e6f59b_mapped, var_7f4a66e6f59b_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_7f4a66e6f59b)
		if var_7f4a66e6f59b_err != nil {
			panic(var_7f4a66e6f59b_err)
		}
		properties["entity"] = var_7f4a66e6f59b_mapped
	}

	var_d0ae5df7c5c6 := resource.Catalog

	if var_d0ae5df7c5c6 != nil {
		var var_d0ae5df7c5c6_mapped *structpb.Value

		var var_d0ae5df7c5c6_err error
		var_d0ae5df7c5c6_mapped, var_d0ae5df7c5c6_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_d0ae5df7c5c6)
		if var_d0ae5df7c5c6_err != nil {
			panic(var_d0ae5df7c5c6_err)
		}
		properties["catalog"] = var_d0ae5df7c5c6_mapped
	}

	var_eb1b1979790b := resource.Annotations

	if var_eb1b1979790b != nil {
		var var_eb1b1979790b_mapped *structpb.Value

		var var_eb1b1979790b_st *structpb.Struct = new(structpb.Struct)
		var_eb1b1979790b_st.Fields = make(map[string]*structpb.Value)
		for key, value := range var_eb1b1979790b {

			var_187bd2e508c4 := value
			var var_187bd2e508c4_mapped *structpb.Value

			var var_187bd2e508c4_err error
			var_187bd2e508c4_mapped, var_187bd2e508c4_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_187bd2e508c4)
			if var_187bd2e508c4_err != nil {
				panic(var_187bd2e508c4_err)
			}

			var_eb1b1979790b_st.Fields[key] = var_187bd2e508c4_mapped
		}
		var_eb1b1979790b_mapped = structpb.NewStructValue(var_eb1b1979790b_st)
		properties["annotations"] = var_eb1b1979790b_mapped
	}

	var_63e6589b1555 := resource.Indexes

	if var_63e6589b1555 != nil {
		var var_63e6589b1555_mapped *structpb.Value

		var var_63e6589b1555_err error
		var_63e6589b1555_mapped, var_63e6589b1555_err = types.ByResourcePropertyType(model.ResourceProperty_OBJECT).Pack(*var_63e6589b1555)
		if var_63e6589b1555_err != nil {
			panic(var_63e6589b1555_err)
		}
		properties["indexes"] = var_63e6589b1555_mapped
	}

	var_6f8e4ac9f904 := resource.Title

	if var_6f8e4ac9f904 != nil {
		var var_6f8e4ac9f904_mapped *structpb.Value

		var var_6f8e4ac9f904_err error
		var_6f8e4ac9f904_mapped, var_6f8e4ac9f904_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_6f8e4ac9f904)
		if var_6f8e4ac9f904_err != nil {
			panic(var_6f8e4ac9f904_err)
		}
		properties["title"] = var_6f8e4ac9f904_mapped
	}

	var_215fb95f8b30 := resource.Description

	if var_215fb95f8b30 != nil {
		var var_215fb95f8b30_mapped *structpb.Value

		var var_215fb95f8b30_err error
		var_215fb95f8b30_mapped, var_215fb95f8b30_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_215fb95f8b30)
		if var_215fb95f8b30_err != nil {
			panic(var_215fb95f8b30_err)
		}
		properties["description"] = var_215fb95f8b30_mapped
	}
	return properties
}

func (m *ResourceMapper) FromProperties(properties map[string]*structpb.Value) *Resource {
	var s = m.New()
	if properties["id"] != nil {

		var_2cc931e0ab43 := properties["id"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_UUID).UnPack(var_2cc931e0ab43)

		if err != nil {
			panic(err)
		}

		var_2cc931e0ab43_mapped := new(uuid.UUID)
		*var_2cc931e0ab43_mapped = val.(uuid.UUID)

		s.Id = var_2cc931e0ab43_mapped
	}
	if properties["version"] != nil {

		var_31d1af9bfa83 := properties["version"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(var_31d1af9bfa83)

		if err != nil {
			panic(err)
		}

		var_31d1af9bfa83_mapped := val.(int32)

		s.Version = var_31d1af9bfa83_mapped
	}
	if properties["createdBy"] != nil {

		var_5dd98d7f854d := properties["createdBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_5dd98d7f854d)

		if err != nil {
			panic(err)
		}

		var_5dd98d7f854d_mapped := new(string)
		*var_5dd98d7f854d_mapped = val.(string)

		s.CreatedBy = var_5dd98d7f854d_mapped
	}
	if properties["updatedBy"] != nil {

		var_6fc9679b5bb1 := properties["updatedBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_6fc9679b5bb1)

		if err != nil {
			panic(err)
		}

		var_6fc9679b5bb1_mapped := new(string)
		*var_6fc9679b5bb1_mapped = val.(string)

		s.UpdatedBy = var_6fc9679b5bb1_mapped
	}
	if properties["createdOn"] != nil {

		var_027394785ab5 := properties["createdOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_027394785ab5)

		if err != nil {
			panic(err)
		}

		var_027394785ab5_mapped := new(time.Time)
		*var_027394785ab5_mapped = val.(time.Time)

		s.CreatedOn = var_027394785ab5_mapped
	}
	if properties["updatedOn"] != nil {

		var_ddd879c97428 := properties["updatedOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_ddd879c97428)

		if err != nil {
			panic(err)
		}

		var_ddd879c97428_mapped := new(time.Time)
		*var_ddd879c97428_mapped = val.(time.Time)

		s.UpdatedOn = var_ddd879c97428_mapped
	}
	if properties["name"] != nil {

		var_9e9a97c6c565 := properties["name"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_9e9a97c6c565)

		if err != nil {
			panic(err)
		}

		var_9e9a97c6c565_mapped := val.(string)

		s.Name = var_9e9a97c6c565_mapped
	}
	if properties["namespace"] != nil {

		var_f75a0ac48b66 := properties["namespace"]
		var_f75a0ac48b66_mapped := NamespaceMapperInstance.FromProperties(var_f75a0ac48b66.GetStructValue().Fields)

		s.Namespace = var_f75a0ac48b66_mapped
	}
	if properties["virtual"] != nil {

		var_97d3ab9ad938 := properties["virtual"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_BOOL).UnPack(var_97d3ab9ad938)

		if err != nil {
			panic(err)
		}

		var_97d3ab9ad938_mapped := val.(bool)

		s.Virtual = var_97d3ab9ad938_mapped
	}
	if properties["properties"] != nil {

		var_e4fb7e683cb6 := properties["properties"]
		var_e4fb7e683cb6_mapped := []ResourceProperty{}
		for _, v := range var_e4fb7e683cb6.GetListValue().Values {

			var_a5df972b93bd := v
			var mappedValue = ResourcePropertyMapperInstance.FromProperties(var_a5df972b93bd.GetStructValue().Fields)

			var_a5df972b93bd_mapped := *mappedValue

			var_e4fb7e683cb6_mapped = append(var_e4fb7e683cb6_mapped, var_a5df972b93bd_mapped)
		}

		s.Properties = var_e4fb7e683cb6_mapped
	}
	if properties["types"] != nil {

		var_389144310e3a := properties["types"]
		var_389144310e3a_mapped := []ResourceSubType{}
		for _, v := range var_389144310e3a.GetListValue().Values {

			var_c2467f5610ad := v
			var mappedValue = ResourceSubTypeMapperInstance.FromProperties(var_c2467f5610ad.GetStructValue().Fields)

			var_c2467f5610ad_mapped := *mappedValue

			var_389144310e3a_mapped = append(var_389144310e3a_mapped, var_c2467f5610ad_mapped)
		}

		s.Types = var_389144310e3a_mapped
	}
	if properties["immutable"] != nil {

		var_2e3ada96da16 := properties["immutable"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_BOOL).UnPack(var_2e3ada96da16)

		if err != nil {
			panic(err)
		}

		var_2e3ada96da16_mapped := val.(bool)

		s.Immutable = var_2e3ada96da16_mapped
	}
	if properties["abstract"] != nil {

		var_3ef604ecdd4d := properties["abstract"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_BOOL).UnPack(var_3ef604ecdd4d)

		if err != nil {
			panic(err)
		}

		var_3ef604ecdd4d_mapped := val.(bool)

		s.Abstract = var_3ef604ecdd4d_mapped
	}
	if properties["dataSource"] != nil {

		var_c8edcf6eb288 := properties["dataSource"]
		var_c8edcf6eb288_mapped := DataSourceMapperInstance.FromProperties(var_c8edcf6eb288.GetStructValue().Fields)

		s.DataSource = var_c8edcf6eb288_mapped
	}
	if properties["entity"] != nil {

		var_1ecb60e36d1e := properties["entity"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_1ecb60e36d1e)

		if err != nil {
			panic(err)
		}

		var_1ecb60e36d1e_mapped := new(string)
		*var_1ecb60e36d1e_mapped = val.(string)

		s.Entity = var_1ecb60e36d1e_mapped
	}
	if properties["catalog"] != nil {

		var_ffe72c0fd3a2 := properties["catalog"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_ffe72c0fd3a2)

		if err != nil {
			panic(err)
		}

		var_ffe72c0fd3a2_mapped := new(string)
		*var_ffe72c0fd3a2_mapped = val.(string)

		s.Catalog = var_ffe72c0fd3a2_mapped
	}
	if properties["annotations"] != nil {

		var_080254044d7d := properties["annotations"]
		var_080254044d7d_mapped := make(map[string]string)
		for k, v := range var_080254044d7d.GetStructValue().Fields {

			var_c6938db65840 := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_c6938db65840)

			if err != nil {
				panic(err)
			}

			var_c6938db65840_mapped := val.(string)

			var_080254044d7d_mapped[k] = var_c6938db65840_mapped
		}

		s.Annotations = var_080254044d7d_mapped
	}
	if properties["indexes"] != nil {

		var_f370a3673607 := properties["indexes"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_OBJECT).UnPack(var_f370a3673607)

		if err != nil {
			panic(err)
		}

		var_f370a3673607_mapped := new(unstructured.Unstructured)
		*var_f370a3673607_mapped = val.(unstructured.Unstructured)

		s.Indexes = var_f370a3673607_mapped
	}
	if properties["title"] != nil {

		var_158ce3448bc8 := properties["title"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_158ce3448bc8)

		if err != nil {
			panic(err)
		}

		var_158ce3448bc8_mapped := new(string)
		*var_158ce3448bc8_mapped = val.(string)

		s.Title = var_158ce3448bc8_mapped
	}
	if properties["description"] != nil {

		var_490e6e061c52 := properties["description"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_490e6e061c52)

		if err != nil {
			panic(err)
		}

		var_490e6e061c52_mapped := new(string)
		*var_490e6e061c52_mapped = val.(string)

		s.Description = var_490e6e061c52_mapped
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

	var_6ac446c5e3c1 := resourceProperty.Name

	var var_6ac446c5e3c1_mapped *structpb.Value

	var var_6ac446c5e3c1_err error
	var_6ac446c5e3c1_mapped, var_6ac446c5e3c1_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_6ac446c5e3c1)
	if var_6ac446c5e3c1_err != nil {
		panic(var_6ac446c5e3c1_err)
	}
	properties["name"] = var_6ac446c5e3c1_mapped

	var_789d4295bdf4 := resourceProperty.Type

	var var_789d4295bdf4_mapped *structpb.Value

	var var_789d4295bdf4_err error
	var_789d4295bdf4_mapped, var_789d4295bdf4_err = types.ByResourcePropertyType(model.ResourceProperty_INT32).Pack(var_789d4295bdf4)
	if var_789d4295bdf4_err != nil {
		panic(var_789d4295bdf4_err)
	}
	properties["type"] = var_789d4295bdf4_mapped

	var_11d3196c2db1 := resourceProperty.TypeRef

	if var_11d3196c2db1 != nil {
		var var_11d3196c2db1_mapped *structpb.Value

		var var_11d3196c2db1_err error
		var_11d3196c2db1_mapped, var_11d3196c2db1_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_11d3196c2db1)
		if var_11d3196c2db1_err != nil {
			panic(var_11d3196c2db1_err)
		}
		properties["typeRef"] = var_11d3196c2db1_mapped
	}

	var_9a3ce5fc47b2 := resourceProperty.Mapping

	var var_9a3ce5fc47b2_mapped *structpb.Value

	var var_9a3ce5fc47b2_err error
	var_9a3ce5fc47b2_mapped, var_9a3ce5fc47b2_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_9a3ce5fc47b2)
	if var_9a3ce5fc47b2_err != nil {
		panic(var_9a3ce5fc47b2_err)
	}
	properties["mapping"] = var_9a3ce5fc47b2_mapped

	var_caa1de17a70b := resourceProperty.Primary

	var var_caa1de17a70b_mapped *structpb.Value

	var var_caa1de17a70b_err error
	var_caa1de17a70b_mapped, var_caa1de17a70b_err = types.ByResourcePropertyType(model.ResourceProperty_BOOL).Pack(var_caa1de17a70b)
	if var_caa1de17a70b_err != nil {
		panic(var_caa1de17a70b_err)
	}
	properties["primary"] = var_caa1de17a70b_mapped

	var_eb4aa7378bea := resourceProperty.Required

	var var_eb4aa7378bea_mapped *structpb.Value

	var var_eb4aa7378bea_err error
	var_eb4aa7378bea_mapped, var_eb4aa7378bea_err = types.ByResourcePropertyType(model.ResourceProperty_BOOL).Pack(var_eb4aa7378bea)
	if var_eb4aa7378bea_err != nil {
		panic(var_eb4aa7378bea_err)
	}
	properties["required"] = var_eb4aa7378bea_mapped

	var_90b0e728a318 := resourceProperty.Unique

	var var_90b0e728a318_mapped *structpb.Value

	var var_90b0e728a318_err error
	var_90b0e728a318_mapped, var_90b0e728a318_err = types.ByResourcePropertyType(model.ResourceProperty_BOOL).Pack(var_90b0e728a318)
	if var_90b0e728a318_err != nil {
		panic(var_90b0e728a318_err)
	}
	properties["unique"] = var_90b0e728a318_mapped

	var_ead234b961ca := resourceProperty.Immutable

	var var_ead234b961ca_mapped *structpb.Value

	var var_ead234b961ca_err error
	var_ead234b961ca_mapped, var_ead234b961ca_err = types.ByResourcePropertyType(model.ResourceProperty_BOOL).Pack(var_ead234b961ca)
	if var_ead234b961ca_err != nil {
		panic(var_ead234b961ca_err)
	}
	properties["immutable"] = var_ead234b961ca_mapped

	var_11ea8914d1ab := resourceProperty.Length

	var var_11ea8914d1ab_mapped *structpb.Value

	var var_11ea8914d1ab_err error
	var_11ea8914d1ab_mapped, var_11ea8914d1ab_err = types.ByResourcePropertyType(model.ResourceProperty_INT32).Pack(var_11ea8914d1ab)
	if var_11ea8914d1ab_err != nil {
		panic(var_11ea8914d1ab_err)
	}
	properties["length"] = var_11ea8914d1ab_mapped

	var_b0e2fb35d2fd := resourceProperty.Resource

	if var_b0e2fb35d2fd != nil {
		var var_b0e2fb35d2fd_mapped *structpb.Value

		var_b0e2fb35d2fd_mapped = structpb.NewStructValue(&structpb.Struct{Fields: ResourceMapperInstance.ToProperties(var_b0e2fb35d2fd)})
		properties["resource"] = var_b0e2fb35d2fd_mapped
	}

	var_6de6b9f5017f := resourceProperty.Item

	if var_6de6b9f5017f != nil {
		var var_6de6b9f5017f_mapped *structpb.Value

		var_6de6b9f5017f_mapped = structpb.NewStructValue(&structpb.Struct{Fields: ResourcePropertyMapperInstance.ToProperties(var_6de6b9f5017f)})
		properties["item"] = var_6de6b9f5017f_mapped
	}

	var_6d8feb082bfc := resourceProperty.Properties

	var var_6d8feb082bfc_mapped *structpb.Value

	var var_6d8feb082bfc_l []*structpb.Value
	for _, value := range var_6d8feb082bfc {

		var_d583ca259405 := value
		var var_d583ca259405_mapped *structpb.Value

		var_d583ca259405_mapped = structpb.NewStructValue(&structpb.Struct{Fields: ResourcePropertyMapperInstance.ToProperties(&var_d583ca259405)})

		var_6d8feb082bfc_l = append(var_6d8feb082bfc_l, var_d583ca259405_mapped)
	}
	var_6d8feb082bfc_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_6d8feb082bfc_l})
	properties["properties"] = var_6d8feb082bfc_mapped

	var_9e27b7688603 := resourceProperty.ReferenceResource

	if var_9e27b7688603 != nil {
		var var_9e27b7688603_mapped *structpb.Value

		var_9e27b7688603_mapped = structpb.NewStructValue(&structpb.Struct{Fields: ResourceMapperInstance.ToProperties(var_9e27b7688603)})
		properties["reference_resource"] = var_9e27b7688603_mapped
	}

	var_830c9edec4af := resourceProperty.ReferenceCascade

	if var_830c9edec4af != nil {
		var var_830c9edec4af_mapped *structpb.Value

		var var_830c9edec4af_err error
		var_830c9edec4af_mapped, var_830c9edec4af_err = types.ByResourcePropertyType(model.ResourceProperty_BOOL).Pack(*var_830c9edec4af)
		if var_830c9edec4af_err != nil {
			panic(var_830c9edec4af_err)
		}
		properties["reference_cascade"] = var_830c9edec4af_mapped
	}

	var_22289356bdee := resourceProperty.BackReferenceProperty

	if var_22289356bdee != nil {
		var var_22289356bdee_mapped *structpb.Value

		var var_22289356bdee_err error
		var_22289356bdee_mapped, var_22289356bdee_err = types.ByResourcePropertyType(model.ResourceProperty_BOOL).Pack(*var_22289356bdee)
		if var_22289356bdee_err != nil {
			panic(var_22289356bdee_err)
		}
		properties["back_reference_property"] = var_22289356bdee_mapped
	}

	var_a27550306802 := resourceProperty.DefaultValue

	if var_a27550306802 != nil {
		var var_a27550306802_mapped *structpb.Value

		var var_a27550306802_err error
		var_a27550306802_mapped, var_a27550306802_err = types.ByResourcePropertyType(model.ResourceProperty_OBJECT).Pack(*var_a27550306802)
		if var_a27550306802_err != nil {
			panic(var_a27550306802_err)
		}
		properties["defaultValue"] = var_a27550306802_mapped
	}

	var_0cb223e366dd := resourceProperty.EnumValues

	if var_0cb223e366dd != nil {
		var var_0cb223e366dd_mapped *structpb.Value

		var var_0cb223e366dd_l []*structpb.Value
		for _, value := range var_0cb223e366dd {

			var_27cfb7e69ad2 := value
			var var_27cfb7e69ad2_mapped *structpb.Value

			var var_27cfb7e69ad2_err error
			var_27cfb7e69ad2_mapped, var_27cfb7e69ad2_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_27cfb7e69ad2)
			if var_27cfb7e69ad2_err != nil {
				panic(var_27cfb7e69ad2_err)
			}

			var_0cb223e366dd_l = append(var_0cb223e366dd_l, var_27cfb7e69ad2_mapped)
		}
		var_0cb223e366dd_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_0cb223e366dd_l})
		properties["enumValues"] = var_0cb223e366dd_mapped
	}

	var_66ceaf514621 := resourceProperty.ExampleValue

	if var_66ceaf514621 != nil {
		var var_66ceaf514621_mapped *structpb.Value

		var var_66ceaf514621_err error
		var_66ceaf514621_mapped, var_66ceaf514621_err = types.ByResourcePropertyType(model.ResourceProperty_OBJECT).Pack(*var_66ceaf514621)
		if var_66ceaf514621_err != nil {
			panic(var_66ceaf514621_err)
		}
		properties["exampleValue"] = var_66ceaf514621_mapped
	}

	var_6435c96611d5 := resourceProperty.Title

	if var_6435c96611d5 != nil {
		var var_6435c96611d5_mapped *structpb.Value

		var var_6435c96611d5_err error
		var_6435c96611d5_mapped, var_6435c96611d5_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_6435c96611d5)
		if var_6435c96611d5_err != nil {
			panic(var_6435c96611d5_err)
		}
		properties["title"] = var_6435c96611d5_mapped
	}

	var_3a4d53f895c8 := resourceProperty.Description

	if var_3a4d53f895c8 != nil {
		var var_3a4d53f895c8_mapped *structpb.Value

		var var_3a4d53f895c8_err error
		var_3a4d53f895c8_mapped, var_3a4d53f895c8_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_3a4d53f895c8)
		if var_3a4d53f895c8_err != nil {
			panic(var_3a4d53f895c8_err)
		}
		properties["description"] = var_3a4d53f895c8_mapped
	}

	var_027e32f77952 := resourceProperty.Annotations

	if var_027e32f77952 != nil {
		var var_027e32f77952_mapped *structpb.Value

		var var_027e32f77952_st *structpb.Struct = new(structpb.Struct)
		var_027e32f77952_st.Fields = make(map[string]*structpb.Value)
		for key, value := range var_027e32f77952 {

			var_6a67bfe4ff39 := value
			var var_6a67bfe4ff39_mapped *structpb.Value

			var var_6a67bfe4ff39_err error
			var_6a67bfe4ff39_mapped, var_6a67bfe4ff39_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_6a67bfe4ff39)
			if var_6a67bfe4ff39_err != nil {
				panic(var_6a67bfe4ff39_err)
			}

			var_027e32f77952_st.Fields[key] = var_6a67bfe4ff39_mapped
		}
		var_027e32f77952_mapped = structpb.NewStructValue(var_027e32f77952_st)
		properties["annotations"] = var_027e32f77952_mapped
	}
	return properties
}

func (m *ResourcePropertyMapper) FromProperties(properties map[string]*structpb.Value) *ResourceProperty {
	var s = m.New()
	if properties["name"] != nil {

		var_81752753f075 := properties["name"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_81752753f075)

		if err != nil {
			panic(err)
		}

		var_81752753f075_mapped := val.(string)

		s.Name = var_81752753f075_mapped
	}
	if properties["type"] != nil {

		var_2d5535945f13 := properties["type"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(var_2d5535945f13)

		if err != nil {
			panic(err)
		}

		var_2d5535945f13_mapped := val.(int32)

		s.Type = var_2d5535945f13_mapped
	}
	if properties["typeRef"] != nil {

		var_9bc447b6671e := properties["typeRef"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_9bc447b6671e)

		if err != nil {
			panic(err)
		}

		var_9bc447b6671e_mapped := new(string)
		*var_9bc447b6671e_mapped = val.(string)

		s.TypeRef = var_9bc447b6671e_mapped
	}
	if properties["mapping"] != nil {

		var_a37574a45f25 := properties["mapping"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_a37574a45f25)

		if err != nil {
			panic(err)
		}

		var_a37574a45f25_mapped := val.(string)

		s.Mapping = var_a37574a45f25_mapped
	}
	if properties["primary"] != nil {

		var_2ed291222412 := properties["primary"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_BOOL).UnPack(var_2ed291222412)

		if err != nil {
			panic(err)
		}

		var_2ed291222412_mapped := val.(bool)

		s.Primary = var_2ed291222412_mapped
	}
	if properties["required"] != nil {

		var_1af14f09ef41 := properties["required"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_BOOL).UnPack(var_1af14f09ef41)

		if err != nil {
			panic(err)
		}

		var_1af14f09ef41_mapped := val.(bool)

		s.Required = var_1af14f09ef41_mapped
	}
	if properties["unique"] != nil {

		var_cdbbf2cc8082 := properties["unique"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_BOOL).UnPack(var_cdbbf2cc8082)

		if err != nil {
			panic(err)
		}

		var_cdbbf2cc8082_mapped := val.(bool)

		s.Unique = var_cdbbf2cc8082_mapped
	}
	if properties["immutable"] != nil {

		var_4444a1430a45 := properties["immutable"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_BOOL).UnPack(var_4444a1430a45)

		if err != nil {
			panic(err)
		}

		var_4444a1430a45_mapped := val.(bool)

		s.Immutable = var_4444a1430a45_mapped
	}
	if properties["length"] != nil {

		var_a261aa4bacc6 := properties["length"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(var_a261aa4bacc6)

		if err != nil {
			panic(err)
		}

		var_a261aa4bacc6_mapped := val.(int32)

		s.Length = var_a261aa4bacc6_mapped
	}
	if properties["resource"] != nil {

		var_13b0dff640f3 := properties["resource"]
		var_13b0dff640f3_mapped := ResourceMapperInstance.FromProperties(var_13b0dff640f3.GetStructValue().Fields)

		s.Resource = var_13b0dff640f3_mapped
	}
	if properties["item"] != nil {

		var_986327cdaf96 := properties["item"]
		var mappedValue = ResourcePropertyMapperInstance.FromProperties(var_986327cdaf96.GetStructValue().Fields)

		var_986327cdaf96_mapped := mappedValue

		s.Item = var_986327cdaf96_mapped
	}
	if properties["properties"] != nil {

		var_34ac7d032cf4 := properties["properties"]
		var_34ac7d032cf4_mapped := []ResourceProperty{}
		for _, v := range var_34ac7d032cf4.GetListValue().Values {

			var_bbbe96a20f97 := v
			var mappedValue = ResourcePropertyMapperInstance.FromProperties(var_bbbe96a20f97.GetStructValue().Fields)

			var_bbbe96a20f97_mapped := *mappedValue

			var_34ac7d032cf4_mapped = append(var_34ac7d032cf4_mapped, var_bbbe96a20f97_mapped)
		}

		s.Properties = var_34ac7d032cf4_mapped
	}
	if properties["reference_resource"] != nil {

		var_153fa2d3c696 := properties["reference_resource"]
		var_153fa2d3c696_mapped := ResourceMapperInstance.FromProperties(var_153fa2d3c696.GetStructValue().Fields)

		s.ReferenceResource = var_153fa2d3c696_mapped
	}
	if properties["reference_cascade"] != nil {

		var_e929872c2f98 := properties["reference_cascade"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_BOOL).UnPack(var_e929872c2f98)

		if err != nil {
			panic(err)
		}

		var_e929872c2f98_mapped := new(bool)
		*var_e929872c2f98_mapped = val.(bool)

		s.ReferenceCascade = var_e929872c2f98_mapped
	}
	if properties["back_reference_property"] != nil {

		var_0cf1af71ec6f := properties["back_reference_property"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_BOOL).UnPack(var_0cf1af71ec6f)

		if err != nil {
			panic(err)
		}

		var_0cf1af71ec6f_mapped := new(bool)
		*var_0cf1af71ec6f_mapped = val.(bool)

		s.BackReferenceProperty = var_0cf1af71ec6f_mapped
	}
	if properties["defaultValue"] != nil {

		var_c4a51171f813 := properties["defaultValue"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_OBJECT).UnPack(var_c4a51171f813)

		if err != nil {
			panic(err)
		}

		var_c4a51171f813_mapped := new(unstructured.Unstructured)
		*var_c4a51171f813_mapped = val.(unstructured.Unstructured)

		s.DefaultValue = var_c4a51171f813_mapped
	}
	if properties["enumValues"] != nil {

		var_ac1bd6f4de0b := properties["enumValues"]
		var_ac1bd6f4de0b_mapped := []string{}
		for _, v := range var_ac1bd6f4de0b.GetListValue().Values {

			var_8407273904ed := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_8407273904ed)

			if err != nil {
				panic(err)
			}

			var_8407273904ed_mapped := val.(string)

			var_ac1bd6f4de0b_mapped = append(var_ac1bd6f4de0b_mapped, var_8407273904ed_mapped)
		}

		s.EnumValues = var_ac1bd6f4de0b_mapped
	}
	if properties["exampleValue"] != nil {

		var_fed346bd49e3 := properties["exampleValue"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_OBJECT).UnPack(var_fed346bd49e3)

		if err != nil {
			panic(err)
		}

		var_fed346bd49e3_mapped := new(unstructured.Unstructured)
		*var_fed346bd49e3_mapped = val.(unstructured.Unstructured)

		s.ExampleValue = var_fed346bd49e3_mapped
	}
	if properties["title"] != nil {

		var_dca18a2aff74 := properties["title"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_dca18a2aff74)

		if err != nil {
			panic(err)
		}

		var_dca18a2aff74_mapped := new(string)
		*var_dca18a2aff74_mapped = val.(string)

		s.Title = var_dca18a2aff74_mapped
	}
	if properties["description"] != nil {

		var_bbe5c86b9291 := properties["description"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_bbe5c86b9291)

		if err != nil {
			panic(err)
		}

		var_bbe5c86b9291_mapped := new(string)
		*var_bbe5c86b9291_mapped = val.(string)

		s.Description = var_bbe5c86b9291_mapped
	}
	if properties["annotations"] != nil {

		var_c61edf87c56a := properties["annotations"]
		var_c61edf87c56a_mapped := make(map[string]string)
		for k, v := range var_c61edf87c56a.GetStructValue().Fields {

			var_25df02f4f902 := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_25df02f4f902)

			if err != nil {
				panic(err)
			}

			var_25df02f4f902_mapped := val.(string)

			var_c61edf87c56a_mapped[k] = var_25df02f4f902_mapped
		}

		s.Annotations = var_c61edf87c56a_mapped
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

	var_4dff19fd1083 := resourceSubType.Name

	var var_4dff19fd1083_mapped *structpb.Value

	var var_4dff19fd1083_err error
	var_4dff19fd1083_mapped, var_4dff19fd1083_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_4dff19fd1083)
	if var_4dff19fd1083_err != nil {
		panic(var_4dff19fd1083_err)
	}
	properties["name"] = var_4dff19fd1083_mapped

	var_106c8150cae0 := resourceSubType.Properties

	var var_106c8150cae0_mapped *structpb.Value

	var var_106c8150cae0_l []*structpb.Value
	for _, value := range var_106c8150cae0 {

		var_4135130fc495 := value
		var var_4135130fc495_mapped *structpb.Value

		var_4135130fc495_mapped = structpb.NewStructValue(&structpb.Struct{Fields: ResourcePropertyMapperInstance.ToProperties(&var_4135130fc495)})

		var_106c8150cae0_l = append(var_106c8150cae0_l, var_4135130fc495_mapped)
	}
	var_106c8150cae0_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_106c8150cae0_l})
	properties["properties"] = var_106c8150cae0_mapped
	return properties
}

func (m *ResourceSubTypeMapper) FromProperties(properties map[string]*structpb.Value) *ResourceSubType {
	var s = m.New()
	if properties["name"] != nil {

		var_e2b4d8f9b4b2 := properties["name"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_e2b4d8f9b4b2)

		if err != nil {
			panic(err)
		}

		var_e2b4d8f9b4b2_mapped := val.(string)

		s.Name = var_e2b4d8f9b4b2_mapped
	}
	if properties["properties"] != nil {

		var_ac0c296a230c := properties["properties"]
		var_ac0c296a230c_mapped := []ResourceProperty{}
		for _, v := range var_ac0c296a230c.GetListValue().Values {

			var_9bda6a9b35aa := v
			var mappedValue = ResourcePropertyMapperInstance.FromProperties(var_9bda6a9b35aa.GetStructValue().Fields)

			var_9bda6a9b35aa_mapped := *mappedValue

			var_ac0c296a230c_mapped = append(var_ac0c296a230c_mapped, var_9bda6a9b35aa_mapped)
		}

		s.Properties = var_ac0c296a230c_mapped
	}
	return s
}
