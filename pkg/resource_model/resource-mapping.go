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

	var_17c724985d75 := resource.Id

	if var_17c724985d75 != nil {
		var var_17c724985d75_mapped *structpb.Value

		var var_17c724985d75_err error
		var_17c724985d75_mapped, var_17c724985d75_err = types.ByResourcePropertyType(model.ResourceProperty_UUID).Pack(*var_17c724985d75)
		if var_17c724985d75_err != nil {
			panic(var_17c724985d75_err)
		}
		properties["id"] = var_17c724985d75_mapped
	}

	var_d7f8fc6eeb87 := resource.Version

	var var_d7f8fc6eeb87_mapped *structpb.Value

	var var_d7f8fc6eeb87_err error
	var_d7f8fc6eeb87_mapped, var_d7f8fc6eeb87_err = types.ByResourcePropertyType(model.ResourceProperty_INT32).Pack(var_d7f8fc6eeb87)
	if var_d7f8fc6eeb87_err != nil {
		panic(var_d7f8fc6eeb87_err)
	}
	properties["version"] = var_d7f8fc6eeb87_mapped

	var_a7e169a69e0a := resource.CreatedBy

	if var_a7e169a69e0a != nil {
		var var_a7e169a69e0a_mapped *structpb.Value

		var var_a7e169a69e0a_err error
		var_a7e169a69e0a_mapped, var_a7e169a69e0a_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_a7e169a69e0a)
		if var_a7e169a69e0a_err != nil {
			panic(var_a7e169a69e0a_err)
		}
		properties["createdBy"] = var_a7e169a69e0a_mapped
	}

	var_dc6fb5b2797e := resource.UpdatedBy

	if var_dc6fb5b2797e != nil {
		var var_dc6fb5b2797e_mapped *structpb.Value

		var var_dc6fb5b2797e_err error
		var_dc6fb5b2797e_mapped, var_dc6fb5b2797e_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_dc6fb5b2797e)
		if var_dc6fb5b2797e_err != nil {
			panic(var_dc6fb5b2797e_err)
		}
		properties["updatedBy"] = var_dc6fb5b2797e_mapped
	}

	var_e0242a4f6c40 := resource.CreatedOn

	if var_e0242a4f6c40 != nil {
		var var_e0242a4f6c40_mapped *structpb.Value

		var var_e0242a4f6c40_err error
		var_e0242a4f6c40_mapped, var_e0242a4f6c40_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_e0242a4f6c40)
		if var_e0242a4f6c40_err != nil {
			panic(var_e0242a4f6c40_err)
		}
		properties["createdOn"] = var_e0242a4f6c40_mapped
	}

	var_bd634df54057 := resource.UpdatedOn

	if var_bd634df54057 != nil {
		var var_bd634df54057_mapped *structpb.Value

		var var_bd634df54057_err error
		var_bd634df54057_mapped, var_bd634df54057_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_bd634df54057)
		if var_bd634df54057_err != nil {
			panic(var_bd634df54057_err)
		}
		properties["updatedOn"] = var_bd634df54057_mapped
	}

	var_4bf4966dc9fc := resource.Name

	var var_4bf4966dc9fc_mapped *structpb.Value

	var var_4bf4966dc9fc_err error
	var_4bf4966dc9fc_mapped, var_4bf4966dc9fc_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_4bf4966dc9fc)
	if var_4bf4966dc9fc_err != nil {
		panic(var_4bf4966dc9fc_err)
	}
	properties["name"] = var_4bf4966dc9fc_mapped

	var_8ae9e66ffc0b := resource.Namespace

	if var_8ae9e66ffc0b != nil {
		var var_8ae9e66ffc0b_mapped *structpb.Value

		var_8ae9e66ffc0b_mapped = structpb.NewStructValue(&structpb.Struct{Fields: NamespaceMapperInstance.ToProperties(var_8ae9e66ffc0b)})
		properties["namespace"] = var_8ae9e66ffc0b_mapped
	}

	var_241f37d36c72 := resource.Virtual

	var var_241f37d36c72_mapped *structpb.Value

	var var_241f37d36c72_err error
	var_241f37d36c72_mapped, var_241f37d36c72_err = types.ByResourcePropertyType(model.ResourceProperty_BOOL).Pack(var_241f37d36c72)
	if var_241f37d36c72_err != nil {
		panic(var_241f37d36c72_err)
	}
	properties["virtual"] = var_241f37d36c72_mapped

	var_326e1d643a27 := resource.Properties

	var var_326e1d643a27_mapped *structpb.Value

	var var_326e1d643a27_l []*structpb.Value
	for _, value := range var_326e1d643a27 {

		var_ba9b426a30e5 := value
		var var_ba9b426a30e5_mapped *structpb.Value

		var_ba9b426a30e5_mapped = structpb.NewStructValue(&structpb.Struct{Fields: ResourcePropertyMapperInstance.ToProperties(&var_ba9b426a30e5)})

		var_326e1d643a27_l = append(var_326e1d643a27_l, var_ba9b426a30e5_mapped)
	}
	var_326e1d643a27_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_326e1d643a27_l})
	properties["properties"] = var_326e1d643a27_mapped

	var_ab0e9b87f41b := resource.Indexes

	if var_ab0e9b87f41b != nil {
		var var_ab0e9b87f41b_mapped *structpb.Value

		var var_ab0e9b87f41b_l []*structpb.Value
		for _, value := range var_ab0e9b87f41b {

			var_8720b7177e7d := value
			var var_8720b7177e7d_mapped *structpb.Value

			var_8720b7177e7d_mapped = structpb.NewStructValue(&structpb.Struct{Fields: ResourceIndexMapperInstance.ToProperties(&var_8720b7177e7d)})

			var_ab0e9b87f41b_l = append(var_ab0e9b87f41b_l, var_8720b7177e7d_mapped)
		}
		var_ab0e9b87f41b_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_ab0e9b87f41b_l})
		properties["indexes"] = var_ab0e9b87f41b_mapped
	}

	var_bf0f3be3f648 := resource.Types

	if var_bf0f3be3f648 != nil {
		var var_bf0f3be3f648_mapped *structpb.Value

		var var_bf0f3be3f648_l []*structpb.Value
		for _, value := range var_bf0f3be3f648 {

			var_53519cc3b671 := value
			var var_53519cc3b671_mapped *structpb.Value

			var_53519cc3b671_mapped = structpb.NewStructValue(&structpb.Struct{Fields: ResourceSubTypeMapperInstance.ToProperties(&var_53519cc3b671)})

			var_bf0f3be3f648_l = append(var_bf0f3be3f648_l, var_53519cc3b671_mapped)
		}
		var_bf0f3be3f648_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_bf0f3be3f648_l})
		properties["types"] = var_bf0f3be3f648_mapped
	}

	var_a559ecad761e := resource.Immutable

	var var_a559ecad761e_mapped *structpb.Value

	var var_a559ecad761e_err error
	var_a559ecad761e_mapped, var_a559ecad761e_err = types.ByResourcePropertyType(model.ResourceProperty_BOOL).Pack(var_a559ecad761e)
	if var_a559ecad761e_err != nil {
		panic(var_a559ecad761e_err)
	}
	properties["immutable"] = var_a559ecad761e_mapped

	var_6d01e866ddeb := resource.Abstract

	var var_6d01e866ddeb_mapped *structpb.Value

	var var_6d01e866ddeb_err error
	var_6d01e866ddeb_mapped, var_6d01e866ddeb_err = types.ByResourcePropertyType(model.ResourceProperty_BOOL).Pack(var_6d01e866ddeb)
	if var_6d01e866ddeb_err != nil {
		panic(var_6d01e866ddeb_err)
	}
	properties["abstract"] = var_6d01e866ddeb_mapped

	var_e14a9c82ff3e := resource.CheckReferences

	var var_e14a9c82ff3e_mapped *structpb.Value

	var var_e14a9c82ff3e_err error
	var_e14a9c82ff3e_mapped, var_e14a9c82ff3e_err = types.ByResourcePropertyType(model.ResourceProperty_BOOL).Pack(var_e14a9c82ff3e)
	if var_e14a9c82ff3e_err != nil {
		panic(var_e14a9c82ff3e_err)
	}
	properties["checkReferences"] = var_e14a9c82ff3e_mapped

	var_46c76bd2002c := resource.DataSource

	if var_46c76bd2002c != nil {
		var var_46c76bd2002c_mapped *structpb.Value

		var_46c76bd2002c_mapped = structpb.NewStructValue(&structpb.Struct{Fields: DataSourceMapperInstance.ToProperties(var_46c76bd2002c)})
		properties["dataSource"] = var_46c76bd2002c_mapped
	}

	var_07e441228ea5 := resource.Entity

	if var_07e441228ea5 != nil {
		var var_07e441228ea5_mapped *structpb.Value

		var var_07e441228ea5_err error
		var_07e441228ea5_mapped, var_07e441228ea5_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_07e441228ea5)
		if var_07e441228ea5_err != nil {
			panic(var_07e441228ea5_err)
		}
		properties["entity"] = var_07e441228ea5_mapped
	}

	var_20d47fa3c3a7 := resource.Catalog

	if var_20d47fa3c3a7 != nil {
		var var_20d47fa3c3a7_mapped *structpb.Value

		var var_20d47fa3c3a7_err error
		var_20d47fa3c3a7_mapped, var_20d47fa3c3a7_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_20d47fa3c3a7)
		if var_20d47fa3c3a7_err != nil {
			panic(var_20d47fa3c3a7_err)
		}
		properties["catalog"] = var_20d47fa3c3a7_mapped
	}

	var_41e09f15270b := resource.Title

	if var_41e09f15270b != nil {
		var var_41e09f15270b_mapped *structpb.Value

		var var_41e09f15270b_err error
		var_41e09f15270b_mapped, var_41e09f15270b_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_41e09f15270b)
		if var_41e09f15270b_err != nil {
			panic(var_41e09f15270b_err)
		}
		properties["title"] = var_41e09f15270b_mapped
	}

	var_965f5d2a8859 := resource.Description

	if var_965f5d2a8859 != nil {
		var var_965f5d2a8859_mapped *structpb.Value

		var var_965f5d2a8859_err error
		var_965f5d2a8859_mapped, var_965f5d2a8859_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_965f5d2a8859)
		if var_965f5d2a8859_err != nil {
			panic(var_965f5d2a8859_err)
		}
		properties["description"] = var_965f5d2a8859_mapped
	}

	var_391c473fa1e2 := resource.Annotations

	if var_391c473fa1e2 != nil {
		var var_391c473fa1e2_mapped *structpb.Value

		var var_391c473fa1e2_st *structpb.Struct = new(structpb.Struct)
		var_391c473fa1e2_st.Fields = make(map[string]*structpb.Value)
		for key, value := range var_391c473fa1e2 {

			var_149310a7790f := value
			var var_149310a7790f_mapped *structpb.Value

			var var_149310a7790f_err error
			var_149310a7790f_mapped, var_149310a7790f_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_149310a7790f)
			if var_149310a7790f_err != nil {
				panic(var_149310a7790f_err)
			}

			var_391c473fa1e2_st.Fields[key] = var_149310a7790f_mapped
		}
		var_391c473fa1e2_mapped = structpb.NewStructValue(var_391c473fa1e2_st)
		properties["annotations"] = var_391c473fa1e2_mapped
	}
	return properties
}

func (m *ResourceMapper) FromProperties(properties map[string]*structpb.Value) *Resource {
	var s = m.New()
	if properties["id"] != nil {

		var_c704eebb8088 := properties["id"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_UUID).UnPack(var_c704eebb8088)

		if err != nil {
			panic(err)
		}

		var_c704eebb8088_mapped := new(uuid.UUID)
		*var_c704eebb8088_mapped = val.(uuid.UUID)

		s.Id = var_c704eebb8088_mapped
	}
	if properties["version"] != nil {

		var_1a0c262e3f62 := properties["version"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(var_1a0c262e3f62)

		if err != nil {
			panic(err)
		}

		var_1a0c262e3f62_mapped := val.(int32)

		s.Version = var_1a0c262e3f62_mapped
	}
	if properties["createdBy"] != nil {

		var_5d00989c5275 := properties["createdBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_5d00989c5275)

		if err != nil {
			panic(err)
		}

		var_5d00989c5275_mapped := new(string)
		*var_5d00989c5275_mapped = val.(string)

		s.CreatedBy = var_5d00989c5275_mapped
	}
	if properties["updatedBy"] != nil {

		var_cdf20e3db686 := properties["updatedBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_cdf20e3db686)

		if err != nil {
			panic(err)
		}

		var_cdf20e3db686_mapped := new(string)
		*var_cdf20e3db686_mapped = val.(string)

		s.UpdatedBy = var_cdf20e3db686_mapped
	}
	if properties["createdOn"] != nil {

		var_b6110e99441f := properties["createdOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_b6110e99441f)

		if err != nil {
			panic(err)
		}

		var_b6110e99441f_mapped := new(time.Time)
		*var_b6110e99441f_mapped = val.(time.Time)

		s.CreatedOn = var_b6110e99441f_mapped
	}
	if properties["updatedOn"] != nil {

		var_6fc76509e3b8 := properties["updatedOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_6fc76509e3b8)

		if err != nil {
			panic(err)
		}

		var_6fc76509e3b8_mapped := new(time.Time)
		*var_6fc76509e3b8_mapped = val.(time.Time)

		s.UpdatedOn = var_6fc76509e3b8_mapped
	}
	if properties["name"] != nil {

		var_1e661720fd35 := properties["name"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_1e661720fd35)

		if err != nil {
			panic(err)
		}

		var_1e661720fd35_mapped := val.(string)

		s.Name = var_1e661720fd35_mapped
	}
	if properties["namespace"] != nil {

		var_9c6d2d2f196d := properties["namespace"]
		var_9c6d2d2f196d_mapped := NamespaceMapperInstance.FromProperties(var_9c6d2d2f196d.GetStructValue().Fields)

		s.Namespace = var_9c6d2d2f196d_mapped
	}
	if properties["virtual"] != nil {

		var_90fdee98b3a8 := properties["virtual"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_BOOL).UnPack(var_90fdee98b3a8)

		if err != nil {
			panic(err)
		}

		var_90fdee98b3a8_mapped := val.(bool)

		s.Virtual = var_90fdee98b3a8_mapped
	}
	if properties["properties"] != nil {

		var_23f81a5fc265 := properties["properties"]
		var_23f81a5fc265_mapped := []ResourceProperty{}
		for _, v := range var_23f81a5fc265.GetListValue().Values {

			var_8e80b2cc35f8 := v
			var mappedValue = ResourcePropertyMapperInstance.FromProperties(var_8e80b2cc35f8.GetStructValue().Fields)

			var_8e80b2cc35f8_mapped := *mappedValue

			var_23f81a5fc265_mapped = append(var_23f81a5fc265_mapped, var_8e80b2cc35f8_mapped)
		}

		s.Properties = var_23f81a5fc265_mapped
	}
	if properties["indexes"] != nil {

		var_d1c4f1a19e18 := properties["indexes"]
		var_d1c4f1a19e18_mapped := []ResourceIndex{}
		for _, v := range var_d1c4f1a19e18.GetListValue().Values {

			var_4e4b3cb61726 := v
			var mappedValue = ResourceIndexMapperInstance.FromProperties(var_4e4b3cb61726.GetStructValue().Fields)

			var_4e4b3cb61726_mapped := *mappedValue

			var_d1c4f1a19e18_mapped = append(var_d1c4f1a19e18_mapped, var_4e4b3cb61726_mapped)
		}

		s.Indexes = var_d1c4f1a19e18_mapped
	}
	if properties["types"] != nil {

		var_6ad4df47dde9 := properties["types"]
		var_6ad4df47dde9_mapped := []ResourceSubType{}
		for _, v := range var_6ad4df47dde9.GetListValue().Values {

			var_900e17a962f8 := v
			var mappedValue = ResourceSubTypeMapperInstance.FromProperties(var_900e17a962f8.GetStructValue().Fields)

			var_900e17a962f8_mapped := *mappedValue

			var_6ad4df47dde9_mapped = append(var_6ad4df47dde9_mapped, var_900e17a962f8_mapped)
		}

		s.Types = var_6ad4df47dde9_mapped
	}
	if properties["immutable"] != nil {

		var_a3cec3f1ba24 := properties["immutable"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_BOOL).UnPack(var_a3cec3f1ba24)

		if err != nil {
			panic(err)
		}

		var_a3cec3f1ba24_mapped := val.(bool)

		s.Immutable = var_a3cec3f1ba24_mapped
	}
	if properties["abstract"] != nil {

		var_ea281c1f5280 := properties["abstract"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_BOOL).UnPack(var_ea281c1f5280)

		if err != nil {
			panic(err)
		}

		var_ea281c1f5280_mapped := val.(bool)

		s.Abstract = var_ea281c1f5280_mapped
	}
	if properties["checkReferences"] != nil {

		var_4544962e6661 := properties["checkReferences"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_BOOL).UnPack(var_4544962e6661)

		if err != nil {
			panic(err)
		}

		var_4544962e6661_mapped := val.(bool)

		s.CheckReferences = var_4544962e6661_mapped
	}
	if properties["dataSource"] != nil {

		var_a37cce8cf0a9 := properties["dataSource"]
		var_a37cce8cf0a9_mapped := DataSourceMapperInstance.FromProperties(var_a37cce8cf0a9.GetStructValue().Fields)

		s.DataSource = var_a37cce8cf0a9_mapped
	}
	if properties["entity"] != nil {

		var_a71e5910e5aa := properties["entity"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_a71e5910e5aa)

		if err != nil {
			panic(err)
		}

		var_a71e5910e5aa_mapped := new(string)
		*var_a71e5910e5aa_mapped = val.(string)

		s.Entity = var_a71e5910e5aa_mapped
	}
	if properties["catalog"] != nil {

		var_bbab7de54ffd := properties["catalog"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_bbab7de54ffd)

		if err != nil {
			panic(err)
		}

		var_bbab7de54ffd_mapped := new(string)
		*var_bbab7de54ffd_mapped = val.(string)

		s.Catalog = var_bbab7de54ffd_mapped
	}
	if properties["title"] != nil {

		var_ddee70c59ddc := properties["title"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_ddee70c59ddc)

		if err != nil {
			panic(err)
		}

		var_ddee70c59ddc_mapped := new(string)
		*var_ddee70c59ddc_mapped = val.(string)

		s.Title = var_ddee70c59ddc_mapped
	}
	if properties["description"] != nil {

		var_ac048c9635c3 := properties["description"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_ac048c9635c3)

		if err != nil {
			panic(err)
		}

		var_ac048c9635c3_mapped := new(string)
		*var_ac048c9635c3_mapped = val.(string)

		s.Description = var_ac048c9635c3_mapped
	}
	if properties["annotations"] != nil {

		var_ab044e2dafe5 := properties["annotations"]
		var_ab044e2dafe5_mapped := make(map[string]string)
		for k, v := range var_ab044e2dafe5.GetStructValue().Fields {

			var_99a41432ae98 := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_99a41432ae98)

			if err != nil {
				panic(err)
			}

			var_99a41432ae98_mapped := val.(string)

			var_ab044e2dafe5_mapped[k] = var_99a41432ae98_mapped
		}

		s.Annotations = var_ab044e2dafe5_mapped
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

	var_a4b63d92b3a3 := resourceProperty.Name

	var var_a4b63d92b3a3_mapped *structpb.Value

	var var_a4b63d92b3a3_err error
	var_a4b63d92b3a3_mapped, var_a4b63d92b3a3_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_a4b63d92b3a3)
	if var_a4b63d92b3a3_err != nil {
		panic(var_a4b63d92b3a3_err)
	}
	properties["name"] = var_a4b63d92b3a3_mapped

	var_95061beb65a0 := resourceProperty.Type

	var var_95061beb65a0_mapped *structpb.Value

	var var_95061beb65a0_err error
	var_95061beb65a0_mapped, var_95061beb65a0_err = types.ByResourcePropertyType(model.ResourceProperty_ENUM).Pack(string(var_95061beb65a0))
	if var_95061beb65a0_err != nil {
		panic(var_95061beb65a0_err)
	}
	properties["type"] = var_95061beb65a0_mapped

	var_fca3d4da2a94 := resourceProperty.TypeRef

	if var_fca3d4da2a94 != nil {
		var var_fca3d4da2a94_mapped *structpb.Value

		var var_fca3d4da2a94_err error
		var_fca3d4da2a94_mapped, var_fca3d4da2a94_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_fca3d4da2a94)
		if var_fca3d4da2a94_err != nil {
			panic(var_fca3d4da2a94_err)
		}
		properties["typeRef"] = var_fca3d4da2a94_mapped
	}

	var_b327022379a4 := resourceProperty.Mapping

	var var_b327022379a4_mapped *structpb.Value

	var var_b327022379a4_err error
	var_b327022379a4_mapped, var_b327022379a4_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_b327022379a4)
	if var_b327022379a4_err != nil {
		panic(var_b327022379a4_err)
	}
	properties["mapping"] = var_b327022379a4_mapped

	var_606a11898b48 := resourceProperty.Primary

	var var_606a11898b48_mapped *structpb.Value

	var var_606a11898b48_err error
	var_606a11898b48_mapped, var_606a11898b48_err = types.ByResourcePropertyType(model.ResourceProperty_BOOL).Pack(var_606a11898b48)
	if var_606a11898b48_err != nil {
		panic(var_606a11898b48_err)
	}
	properties["primary"] = var_606a11898b48_mapped

	var_666a19a693e0 := resourceProperty.Required

	var var_666a19a693e0_mapped *structpb.Value

	var var_666a19a693e0_err error
	var_666a19a693e0_mapped, var_666a19a693e0_err = types.ByResourcePropertyType(model.ResourceProperty_BOOL).Pack(var_666a19a693e0)
	if var_666a19a693e0_err != nil {
		panic(var_666a19a693e0_err)
	}
	properties["required"] = var_666a19a693e0_mapped

	var_d29b96c17737 := resourceProperty.Unique

	var var_d29b96c17737_mapped *structpb.Value

	var var_d29b96c17737_err error
	var_d29b96c17737_mapped, var_d29b96c17737_err = types.ByResourcePropertyType(model.ResourceProperty_BOOL).Pack(var_d29b96c17737)
	if var_d29b96c17737_err != nil {
		panic(var_d29b96c17737_err)
	}
	properties["unique"] = var_d29b96c17737_mapped

	var_ecfdc7637267 := resourceProperty.Immutable

	var var_ecfdc7637267_mapped *structpb.Value

	var var_ecfdc7637267_err error
	var_ecfdc7637267_mapped, var_ecfdc7637267_err = types.ByResourcePropertyType(model.ResourceProperty_BOOL).Pack(var_ecfdc7637267)
	if var_ecfdc7637267_err != nil {
		panic(var_ecfdc7637267_err)
	}
	properties["immutable"] = var_ecfdc7637267_mapped

	var_d5777fe51213 := resourceProperty.Length

	var var_d5777fe51213_mapped *structpb.Value

	var var_d5777fe51213_err error
	var_d5777fe51213_mapped, var_d5777fe51213_err = types.ByResourcePropertyType(model.ResourceProperty_INT32).Pack(var_d5777fe51213)
	if var_d5777fe51213_err != nil {
		panic(var_d5777fe51213_err)
	}
	properties["length"] = var_d5777fe51213_mapped

	var_692f56d8aa0c := resourceProperty.Item

	if var_692f56d8aa0c != nil {
		var var_692f56d8aa0c_mapped *structpb.Value

		var_692f56d8aa0c_mapped = structpb.NewStructValue(&structpb.Struct{Fields: ResourcePropertyMapperInstance.ToProperties(var_692f56d8aa0c)})
		properties["item"] = var_692f56d8aa0c_mapped
	}

	var_0053f71a5c3b := resourceProperty.Properties

	var var_0053f71a5c3b_mapped *structpb.Value

	var var_0053f71a5c3b_l []*structpb.Value
	for _, value := range var_0053f71a5c3b {

		var_9b0f9574b364 := value
		var var_9b0f9574b364_mapped *structpb.Value

		var_9b0f9574b364_mapped = structpb.NewStructValue(&structpb.Struct{Fields: ResourcePropertyMapperInstance.ToProperties(&var_9b0f9574b364)})

		var_0053f71a5c3b_l = append(var_0053f71a5c3b_l, var_9b0f9574b364_mapped)
	}
	var_0053f71a5c3b_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_0053f71a5c3b_l})
	properties["properties"] = var_0053f71a5c3b_mapped

	var_aa750bef36ce := resourceProperty.Reference

	if var_aa750bef36ce != nil {
		var var_aa750bef36ce_mapped *structpb.Value

		var_aa750bef36ce_mapped = structpb.NewStructValue(&structpb.Struct{Fields: ResourceReferenceMapperInstance.ToProperties(var_aa750bef36ce)})
		properties["reference"] = var_aa750bef36ce_mapped
	}

	var_91fb8dab6871 := resourceProperty.DefaultValue

	if var_91fb8dab6871 != nil {
		var var_91fb8dab6871_mapped *structpb.Value

		var var_91fb8dab6871_err error
		var_91fb8dab6871_mapped, var_91fb8dab6871_err = types.ByResourcePropertyType(model.ResourceProperty_OBJECT).Pack(*var_91fb8dab6871)
		if var_91fb8dab6871_err != nil {
			panic(var_91fb8dab6871_err)
		}
		properties["defaultValue"] = var_91fb8dab6871_mapped
	}

	var_33e2024f4855 := resourceProperty.EnumValues

	if var_33e2024f4855 != nil {
		var var_33e2024f4855_mapped *structpb.Value

		var var_33e2024f4855_l []*structpb.Value
		for _, value := range var_33e2024f4855 {

			var_90c7171d316d := value
			var var_90c7171d316d_mapped *structpb.Value

			var var_90c7171d316d_err error
			var_90c7171d316d_mapped, var_90c7171d316d_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_90c7171d316d)
			if var_90c7171d316d_err != nil {
				panic(var_90c7171d316d_err)
			}

			var_33e2024f4855_l = append(var_33e2024f4855_l, var_90c7171d316d_mapped)
		}
		var_33e2024f4855_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_33e2024f4855_l})
		properties["enumValues"] = var_33e2024f4855_mapped
	}

	var_96f6586cfa14 := resourceProperty.ExampleValue

	if var_96f6586cfa14 != nil {
		var var_96f6586cfa14_mapped *structpb.Value

		var var_96f6586cfa14_err error
		var_96f6586cfa14_mapped, var_96f6586cfa14_err = types.ByResourcePropertyType(model.ResourceProperty_OBJECT).Pack(*var_96f6586cfa14)
		if var_96f6586cfa14_err != nil {
			panic(var_96f6586cfa14_err)
		}
		properties["exampleValue"] = var_96f6586cfa14_mapped
	}

	var_1c69b5437446 := resourceProperty.Title

	if var_1c69b5437446 != nil {
		var var_1c69b5437446_mapped *structpb.Value

		var var_1c69b5437446_err error
		var_1c69b5437446_mapped, var_1c69b5437446_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_1c69b5437446)
		if var_1c69b5437446_err != nil {
			panic(var_1c69b5437446_err)
		}
		properties["title"] = var_1c69b5437446_mapped
	}

	var_63f8ea8c6301 := resourceProperty.Description

	if var_63f8ea8c6301 != nil {
		var var_63f8ea8c6301_mapped *structpb.Value

		var var_63f8ea8c6301_err error
		var_63f8ea8c6301_mapped, var_63f8ea8c6301_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_63f8ea8c6301)
		if var_63f8ea8c6301_err != nil {
			panic(var_63f8ea8c6301_err)
		}
		properties["description"] = var_63f8ea8c6301_mapped
	}

	var_6b829bd0985b := resourceProperty.Annotations

	if var_6b829bd0985b != nil {
		var var_6b829bd0985b_mapped *structpb.Value

		var var_6b829bd0985b_st *structpb.Struct = new(structpb.Struct)
		var_6b829bd0985b_st.Fields = make(map[string]*structpb.Value)
		for key, value := range var_6b829bd0985b {

			var_61d308815d32 := value
			var var_61d308815d32_mapped *structpb.Value

			var var_61d308815d32_err error
			var_61d308815d32_mapped, var_61d308815d32_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_61d308815d32)
			if var_61d308815d32_err != nil {
				panic(var_61d308815d32_err)
			}

			var_6b829bd0985b_st.Fields[key] = var_61d308815d32_mapped
		}
		var_6b829bd0985b_mapped = structpb.NewStructValue(var_6b829bd0985b_st)
		properties["annotations"] = var_6b829bd0985b_mapped
	}
	return properties
}

func (m *ResourcePropertyMapper) FromProperties(properties map[string]*structpb.Value) *ResourceProperty {
	var s = m.New()
	if properties["name"] != nil {

		var_9a850ab64196 := properties["name"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_9a850ab64196)

		if err != nil {
			panic(err)
		}

		var_9a850ab64196_mapped := val.(string)

		s.Name = var_9a850ab64196_mapped
	}
	if properties["type"] != nil {

		var_d18610d83dcd := properties["type"]
		var_d18610d83dcd_mapped := (ResourceType)(var_d18610d83dcd.GetStringValue())

		s.Type = var_d18610d83dcd_mapped
	}
	if properties["typeRef"] != nil {

		var_b2ed6da48049 := properties["typeRef"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_b2ed6da48049)

		if err != nil {
			panic(err)
		}

		var_b2ed6da48049_mapped := new(string)
		*var_b2ed6da48049_mapped = val.(string)

		s.TypeRef = var_b2ed6da48049_mapped
	}
	if properties["mapping"] != nil {

		var_df96a6697c46 := properties["mapping"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_df96a6697c46)

		if err != nil {
			panic(err)
		}

		var_df96a6697c46_mapped := val.(string)

		s.Mapping = var_df96a6697c46_mapped
	}
	if properties["primary"] != nil {

		var_0ae6305db9eb := properties["primary"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_BOOL).UnPack(var_0ae6305db9eb)

		if err != nil {
			panic(err)
		}

		var_0ae6305db9eb_mapped := val.(bool)

		s.Primary = var_0ae6305db9eb_mapped
	}
	if properties["required"] != nil {

		var_385018911ee5 := properties["required"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_BOOL).UnPack(var_385018911ee5)

		if err != nil {
			panic(err)
		}

		var_385018911ee5_mapped := val.(bool)

		s.Required = var_385018911ee5_mapped
	}
	if properties["unique"] != nil {

		var_4b68b6914685 := properties["unique"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_BOOL).UnPack(var_4b68b6914685)

		if err != nil {
			panic(err)
		}

		var_4b68b6914685_mapped := val.(bool)

		s.Unique = var_4b68b6914685_mapped
	}
	if properties["immutable"] != nil {

		var_4b68a2035020 := properties["immutable"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_BOOL).UnPack(var_4b68a2035020)

		if err != nil {
			panic(err)
		}

		var_4b68a2035020_mapped := val.(bool)

		s.Immutable = var_4b68a2035020_mapped
	}
	if properties["length"] != nil {

		var_b2c97c779bc1 := properties["length"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(var_b2c97c779bc1)

		if err != nil {
			panic(err)
		}

		var_b2c97c779bc1_mapped := val.(int32)

		s.Length = var_b2c97c779bc1_mapped
	}
	if properties["item"] != nil {

		var_c9a7fc6d0691 := properties["item"]
		var mappedValue = ResourcePropertyMapperInstance.FromProperties(var_c9a7fc6d0691.GetStructValue().Fields)

		var_c9a7fc6d0691_mapped := mappedValue

		s.Item = var_c9a7fc6d0691_mapped
	}
	if properties["properties"] != nil {

		var_7232cfaeb2a9 := properties["properties"]
		var_7232cfaeb2a9_mapped := []ResourceProperty{}
		for _, v := range var_7232cfaeb2a9.GetListValue().Values {

			var_2b5afdd882b0 := v
			var mappedValue = ResourcePropertyMapperInstance.FromProperties(var_2b5afdd882b0.GetStructValue().Fields)

			var_2b5afdd882b0_mapped := *mappedValue

			var_7232cfaeb2a9_mapped = append(var_7232cfaeb2a9_mapped, var_2b5afdd882b0_mapped)
		}

		s.Properties = var_7232cfaeb2a9_mapped
	}
	if properties["reference"] != nil {

		var_e98342eeb447 := properties["reference"]
		var mappedValue = ResourceReferenceMapperInstance.FromProperties(var_e98342eeb447.GetStructValue().Fields)

		var_e98342eeb447_mapped := mappedValue

		s.Reference = var_e98342eeb447_mapped
	}
	if properties["defaultValue"] != nil {

		var_d5a9b8eb57dc := properties["defaultValue"]
		var_d5a9b8eb57dc_mapped := new(unstructured.Unstructured)
		*var_d5a9b8eb57dc_mapped = unstructured.FromStructValue(var_d5a9b8eb57dc.GetStructValue())

		s.DefaultValue = var_d5a9b8eb57dc_mapped
	}
	if properties["enumValues"] != nil {

		var_f01bb6e0be7e := properties["enumValues"]
		var_f01bb6e0be7e_mapped := []string{}
		for _, v := range var_f01bb6e0be7e.GetListValue().Values {

			var_a11e5956a74c := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_a11e5956a74c)

			if err != nil {
				panic(err)
			}

			var_a11e5956a74c_mapped := val.(string)

			var_f01bb6e0be7e_mapped = append(var_f01bb6e0be7e_mapped, var_a11e5956a74c_mapped)
		}

		s.EnumValues = var_f01bb6e0be7e_mapped
	}
	if properties["exampleValue"] != nil {

		var_2684963a4bd3 := properties["exampleValue"]
		var_2684963a4bd3_mapped := new(unstructured.Unstructured)
		*var_2684963a4bd3_mapped = unstructured.FromStructValue(var_2684963a4bd3.GetStructValue())

		s.ExampleValue = var_2684963a4bd3_mapped
	}
	if properties["title"] != nil {

		var_1fee87d43a33 := properties["title"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_1fee87d43a33)

		if err != nil {
			panic(err)
		}

		var_1fee87d43a33_mapped := new(string)
		*var_1fee87d43a33_mapped = val.(string)

		s.Title = var_1fee87d43a33_mapped
	}
	if properties["description"] != nil {

		var_74f11455b5d4 := properties["description"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_74f11455b5d4)

		if err != nil {
			panic(err)
		}

		var_74f11455b5d4_mapped := new(string)
		*var_74f11455b5d4_mapped = val.(string)

		s.Description = var_74f11455b5d4_mapped
	}
	if properties["annotations"] != nil {

		var_8b5551c77d8f := properties["annotations"]
		var_8b5551c77d8f_mapped := make(map[string]string)
		for k, v := range var_8b5551c77d8f.GetStructValue().Fields {

			var_e4acf0752cc0 := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_e4acf0752cc0)

			if err != nil {
				panic(err)
			}

			var_e4acf0752cc0_mapped := val.(string)

			var_8b5551c77d8f_mapped[k] = var_e4acf0752cc0_mapped
		}

		s.Annotations = var_8b5551c77d8f_mapped
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

	var_a8c8cd5ce6ed := resourceSubType.Name

	var var_a8c8cd5ce6ed_mapped *structpb.Value

	var var_a8c8cd5ce6ed_err error
	var_a8c8cd5ce6ed_mapped, var_a8c8cd5ce6ed_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_a8c8cd5ce6ed)
	if var_a8c8cd5ce6ed_err != nil {
		panic(var_a8c8cd5ce6ed_err)
	}
	properties["name"] = var_a8c8cd5ce6ed_mapped

	var_f7a860ca0bf6 := resourceSubType.Title

	if var_f7a860ca0bf6 != nil {
		var var_f7a860ca0bf6_mapped *structpb.Value

		var var_f7a860ca0bf6_err error
		var_f7a860ca0bf6_mapped, var_f7a860ca0bf6_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_f7a860ca0bf6)
		if var_f7a860ca0bf6_err != nil {
			panic(var_f7a860ca0bf6_err)
		}
		properties["title"] = var_f7a860ca0bf6_mapped
	}

	var_b212e80efa83 := resourceSubType.Description

	if var_b212e80efa83 != nil {
		var var_b212e80efa83_mapped *structpb.Value

		var var_b212e80efa83_err error
		var_b212e80efa83_mapped, var_b212e80efa83_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_b212e80efa83)
		if var_b212e80efa83_err != nil {
			panic(var_b212e80efa83_err)
		}
		properties["description"] = var_b212e80efa83_mapped
	}

	var_98c1d3bf6dbf := resourceSubType.Properties

	var var_98c1d3bf6dbf_mapped *structpb.Value

	var var_98c1d3bf6dbf_l []*structpb.Value
	for _, value := range var_98c1d3bf6dbf {

		var_e577841a432d := value
		var var_e577841a432d_mapped *structpb.Value

		var_e577841a432d_mapped = structpb.NewStructValue(&structpb.Struct{Fields: ResourcePropertyMapperInstance.ToProperties(&var_e577841a432d)})

		var_98c1d3bf6dbf_l = append(var_98c1d3bf6dbf_l, var_e577841a432d_mapped)
	}
	var_98c1d3bf6dbf_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_98c1d3bf6dbf_l})
	properties["properties"] = var_98c1d3bf6dbf_mapped
	return properties
}

func (m *ResourceSubTypeMapper) FromProperties(properties map[string]*structpb.Value) *ResourceSubType {
	var s = m.New()
	if properties["name"] != nil {

		var_f6068896c88d := properties["name"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_f6068896c88d)

		if err != nil {
			panic(err)
		}

		var_f6068896c88d_mapped := val.(string)

		s.Name = var_f6068896c88d_mapped
	}
	if properties["title"] != nil {

		var_4290b0661e94 := properties["title"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_4290b0661e94)

		if err != nil {
			panic(err)
		}

		var_4290b0661e94_mapped := new(string)
		*var_4290b0661e94_mapped = val.(string)

		s.Title = var_4290b0661e94_mapped
	}
	if properties["description"] != nil {

		var_124bfa07d629 := properties["description"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_124bfa07d629)

		if err != nil {
			panic(err)
		}

		var_124bfa07d629_mapped := new(string)
		*var_124bfa07d629_mapped = val.(string)

		s.Description = var_124bfa07d629_mapped
	}
	if properties["properties"] != nil {

		var_c3498b891b9c := properties["properties"]
		var_c3498b891b9c_mapped := []ResourceProperty{}
		for _, v := range var_c3498b891b9c.GetListValue().Values {

			var_8465cb56812c := v
			var mappedValue = ResourcePropertyMapperInstance.FromProperties(var_8465cb56812c.GetStructValue().Fields)

			var_8465cb56812c_mapped := *mappedValue

			var_c3498b891b9c_mapped = append(var_c3498b891b9c_mapped, var_8465cb56812c_mapped)
		}

		s.Properties = var_c3498b891b9c_mapped
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

	var_5e0b9a1f0a36 := resourceIndexProperty.Name

	var var_5e0b9a1f0a36_mapped *structpb.Value

	var var_5e0b9a1f0a36_err error
	var_5e0b9a1f0a36_mapped, var_5e0b9a1f0a36_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_5e0b9a1f0a36)
	if var_5e0b9a1f0a36_err != nil {
		panic(var_5e0b9a1f0a36_err)
	}
	properties["name"] = var_5e0b9a1f0a36_mapped

	var_be7b1ea7827e := resourceIndexProperty.Order

	if var_be7b1ea7827e != nil {
		var var_be7b1ea7827e_mapped *structpb.Value

		var var_be7b1ea7827e_err error
		var_be7b1ea7827e_mapped, var_be7b1ea7827e_err = types.ByResourcePropertyType(model.ResourceProperty_ENUM).Pack(string(*var_be7b1ea7827e))
		if var_be7b1ea7827e_err != nil {
			panic(var_be7b1ea7827e_err)
		}
		properties["order"] = var_be7b1ea7827e_mapped
	}
	return properties
}

func (m *ResourceIndexPropertyMapper) FromProperties(properties map[string]*structpb.Value) *ResourceIndexProperty {
	var s = m.New()
	if properties["name"] != nil {

		var_1384c60769fc := properties["name"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_1384c60769fc)

		if err != nil {
			panic(err)
		}

		var_1384c60769fc_mapped := val.(string)

		s.Name = var_1384c60769fc_mapped
	}
	if properties["order"] != nil {

		var_ff3434784287 := properties["order"]
		var_ff3434784287_mapped := new(ResourceOrder)
		*var_ff3434784287_mapped = (ResourceOrder)(var_ff3434784287.GetStringValue())

		s.Order = var_ff3434784287_mapped
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

	var_a2e101f89f08 := resourceIndex.Properties

	if var_a2e101f89f08 != nil {
		var var_a2e101f89f08_mapped *structpb.Value

		var var_a2e101f89f08_l []*structpb.Value
		for _, value := range var_a2e101f89f08 {

			var_ae4f43fb49c1 := value
			var var_ae4f43fb49c1_mapped *structpb.Value

			var_ae4f43fb49c1_mapped = structpb.NewStructValue(&structpb.Struct{Fields: ResourceIndexPropertyMapperInstance.ToProperties(&var_ae4f43fb49c1)})

			var_a2e101f89f08_l = append(var_a2e101f89f08_l, var_ae4f43fb49c1_mapped)
		}
		var_a2e101f89f08_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_a2e101f89f08_l})
		properties["properties"] = var_a2e101f89f08_mapped
	}

	var_cd312f00678e := resourceIndex.IndexType

	if var_cd312f00678e != nil {
		var var_cd312f00678e_mapped *structpb.Value

		var var_cd312f00678e_err error
		var_cd312f00678e_mapped, var_cd312f00678e_err = types.ByResourcePropertyType(model.ResourceProperty_ENUM).Pack(string(*var_cd312f00678e))
		if var_cd312f00678e_err != nil {
			panic(var_cd312f00678e_err)
		}
		properties["indexType"] = var_cd312f00678e_mapped
	}

	var_a0df07b74835 := resourceIndex.Unique

	if var_a0df07b74835 != nil {
		var var_a0df07b74835_mapped *structpb.Value

		var var_a0df07b74835_err error
		var_a0df07b74835_mapped, var_a0df07b74835_err = types.ByResourcePropertyType(model.ResourceProperty_BOOL).Pack(*var_a0df07b74835)
		if var_a0df07b74835_err != nil {
			panic(var_a0df07b74835_err)
		}
		properties["unique"] = var_a0df07b74835_mapped
	}

	var_3b532d52f202 := resourceIndex.Annotations

	if var_3b532d52f202 != nil {
		var var_3b532d52f202_mapped *structpb.Value

		var var_3b532d52f202_st *structpb.Struct = new(structpb.Struct)
		var_3b532d52f202_st.Fields = make(map[string]*structpb.Value)
		for key, value := range var_3b532d52f202 {

			var_0fcf1632c41e := value
			var var_0fcf1632c41e_mapped *structpb.Value

			var var_0fcf1632c41e_err error
			var_0fcf1632c41e_mapped, var_0fcf1632c41e_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_0fcf1632c41e)
			if var_0fcf1632c41e_err != nil {
				panic(var_0fcf1632c41e_err)
			}

			var_3b532d52f202_st.Fields[key] = var_0fcf1632c41e_mapped
		}
		var_3b532d52f202_mapped = structpb.NewStructValue(var_3b532d52f202_st)
		properties["annotations"] = var_3b532d52f202_mapped
	}
	return properties
}

func (m *ResourceIndexMapper) FromProperties(properties map[string]*structpb.Value) *ResourceIndex {
	var s = m.New()
	if properties["properties"] != nil {

		var_11356976e7fa := properties["properties"]
		var_11356976e7fa_mapped := []ResourceIndexProperty{}
		for _, v := range var_11356976e7fa.GetListValue().Values {

			var_e036598e6ffe := v
			var mappedValue = ResourceIndexPropertyMapperInstance.FromProperties(var_e036598e6ffe.GetStructValue().Fields)

			var_e036598e6ffe_mapped := *mappedValue

			var_11356976e7fa_mapped = append(var_11356976e7fa_mapped, var_e036598e6ffe_mapped)
		}

		s.Properties = var_11356976e7fa_mapped
	}
	if properties["indexType"] != nil {

		var_5ad29bcfe650 := properties["indexType"]
		var_5ad29bcfe650_mapped := new(ResourceIndexType)
		*var_5ad29bcfe650_mapped = (ResourceIndexType)(var_5ad29bcfe650.GetStringValue())

		s.IndexType = var_5ad29bcfe650_mapped
	}
	if properties["unique"] != nil {

		var_39125b56a528 := properties["unique"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_BOOL).UnPack(var_39125b56a528)

		if err != nil {
			panic(err)
		}

		var_39125b56a528_mapped := new(bool)
		*var_39125b56a528_mapped = val.(bool)

		s.Unique = var_39125b56a528_mapped
	}
	if properties["annotations"] != nil {

		var_0cb179906b32 := properties["annotations"]
		var_0cb179906b32_mapped := make(map[string]string)
		for k, v := range var_0cb179906b32.GetStructValue().Fields {

			var_7afd3b259dc6 := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_7afd3b259dc6)

			if err != nil {
				panic(err)
			}

			var_7afd3b259dc6_mapped := val.(string)

			var_0cb179906b32_mapped[k] = var_7afd3b259dc6_mapped
		}

		s.Annotations = var_0cb179906b32_mapped
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

	var_d813454c83ff := resourceReference.Resource

	if var_d813454c83ff != nil {
		var var_d813454c83ff_mapped *structpb.Value

		var_d813454c83ff_mapped = structpb.NewStructValue(&structpb.Struct{Fields: ResourceMapperInstance.ToProperties(var_d813454c83ff)})
		properties["resource"] = var_d813454c83ff_mapped
	}

	var_21c02cf415fa := resourceReference.Cascade

	if var_21c02cf415fa != nil {
		var var_21c02cf415fa_mapped *structpb.Value

		var var_21c02cf415fa_err error
		var_21c02cf415fa_mapped, var_21c02cf415fa_err = types.ByResourcePropertyType(model.ResourceProperty_BOOL).Pack(*var_21c02cf415fa)
		if var_21c02cf415fa_err != nil {
			panic(var_21c02cf415fa_err)
		}
		properties["cascade"] = var_21c02cf415fa_mapped
	}

	var_cff0587a1f5c := resourceReference.BackReference

	if var_cff0587a1f5c != nil {
		var var_cff0587a1f5c_mapped *structpb.Value

		var var_cff0587a1f5c_err error
		var_cff0587a1f5c_mapped, var_cff0587a1f5c_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_cff0587a1f5c)
		if var_cff0587a1f5c_err != nil {
			panic(var_cff0587a1f5c_err)
		}
		properties["backReference"] = var_cff0587a1f5c_mapped
	}
	return properties
}

func (m *ResourceReferenceMapper) FromProperties(properties map[string]*structpb.Value) *ResourceReference {
	var s = m.New()
	if properties["resource"] != nil {

		var_83b1798df59e := properties["resource"]
		var_83b1798df59e_mapped := ResourceMapperInstance.FromProperties(var_83b1798df59e.GetStructValue().Fields)

		s.Resource = var_83b1798df59e_mapped
	}
	if properties["cascade"] != nil {

		var_4dd8c8133359 := properties["cascade"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_BOOL).UnPack(var_4dd8c8133359)

		if err != nil {
			panic(err)
		}

		var_4dd8c8133359_mapped := new(bool)
		*var_4dd8c8133359_mapped = val.(bool)

		s.Cascade = var_4dd8c8133359_mapped
	}
	if properties["backReference"] != nil {

		var_85a3b9924109 := properties["backReference"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_85a3b9924109)

		if err != nil {
			panic(err)
		}

		var_85a3b9924109_mapped := new(string)
		*var_85a3b9924109_mapped = val.(string)

		s.BackReference = var_85a3b9924109_mapped
	}
	return s
}
