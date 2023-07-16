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
	return rec
}

func (m *ResourceMapper) FromRecord(record *model.Record) *Resource {
	return m.FromProperties(record.Properties)
}

func (m *ResourceMapper) ToProperties(resource *Resource) map[string]*structpb.Value {
	var properties = make(map[string]*structpb.Value)

	var_8a74059ee453 := resource.Id

	if var_8a74059ee453 != nil {
		var var_8a74059ee453_mapped *structpb.Value

		var var_8a74059ee453_err error
		var_8a74059ee453_mapped, var_8a74059ee453_err = types.ByResourcePropertyType(model.ResourceProperty_UUID).Pack(*var_8a74059ee453)
		if var_8a74059ee453_err != nil {
			panic(var_8a74059ee453_err)
		}
		properties["id"] = var_8a74059ee453_mapped
	}

	var_700f9843d0e0 := resource.Version

	var var_700f9843d0e0_mapped *structpb.Value

	var var_700f9843d0e0_err error
	var_700f9843d0e0_mapped, var_700f9843d0e0_err = types.ByResourcePropertyType(model.ResourceProperty_INT32).Pack(var_700f9843d0e0)
	if var_700f9843d0e0_err != nil {
		panic(var_700f9843d0e0_err)
	}
	properties["version"] = var_700f9843d0e0_mapped

	var_e945d9ca338b := resource.CreatedBy

	if var_e945d9ca338b != nil {
		var var_e945d9ca338b_mapped *structpb.Value

		var var_e945d9ca338b_err error
		var_e945d9ca338b_mapped, var_e945d9ca338b_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_e945d9ca338b)
		if var_e945d9ca338b_err != nil {
			panic(var_e945d9ca338b_err)
		}
		properties["createdBy"] = var_e945d9ca338b_mapped
	}

	var_0d08eaf52a1f := resource.UpdatedBy

	if var_0d08eaf52a1f != nil {
		var var_0d08eaf52a1f_mapped *structpb.Value

		var var_0d08eaf52a1f_err error
		var_0d08eaf52a1f_mapped, var_0d08eaf52a1f_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_0d08eaf52a1f)
		if var_0d08eaf52a1f_err != nil {
			panic(var_0d08eaf52a1f_err)
		}
		properties["updatedBy"] = var_0d08eaf52a1f_mapped
	}

	var_fd857d4d7b4a := resource.CreatedOn

	if var_fd857d4d7b4a != nil {
		var var_fd857d4d7b4a_mapped *structpb.Value

		var var_fd857d4d7b4a_err error
		var_fd857d4d7b4a_mapped, var_fd857d4d7b4a_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_fd857d4d7b4a)
		if var_fd857d4d7b4a_err != nil {
			panic(var_fd857d4d7b4a_err)
		}
		properties["createdOn"] = var_fd857d4d7b4a_mapped
	}

	var_1eaf4059dc57 := resource.UpdatedOn

	if var_1eaf4059dc57 != nil {
		var var_1eaf4059dc57_mapped *structpb.Value

		var var_1eaf4059dc57_err error
		var_1eaf4059dc57_mapped, var_1eaf4059dc57_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_1eaf4059dc57)
		if var_1eaf4059dc57_err != nil {
			panic(var_1eaf4059dc57_err)
		}
		properties["updatedOn"] = var_1eaf4059dc57_mapped
	}

	var_dc1f52402baf := resource.Name

	var var_dc1f52402baf_mapped *structpb.Value

	var var_dc1f52402baf_err error
	var_dc1f52402baf_mapped, var_dc1f52402baf_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_dc1f52402baf)
	if var_dc1f52402baf_err != nil {
		panic(var_dc1f52402baf_err)
	}
	properties["name"] = var_dc1f52402baf_mapped

	var_117120c7c020 := resource.Namespace

	if var_117120c7c020 != nil {
		var var_117120c7c020_mapped *structpb.Value

		var_117120c7c020_mapped = structpb.NewStructValue(&structpb.Struct{Fields: NamespaceMapperInstance.ToProperties(var_117120c7c020)})
		properties["namespace"] = var_117120c7c020_mapped
	}

	var_53764256e105 := resource.Virtual

	var var_53764256e105_mapped *structpb.Value

	var var_53764256e105_err error
	var_53764256e105_mapped, var_53764256e105_err = types.ByResourcePropertyType(model.ResourceProperty_BOOL).Pack(var_53764256e105)
	if var_53764256e105_err != nil {
		panic(var_53764256e105_err)
	}
	properties["virtual"] = var_53764256e105_mapped

	var_c4f1cb6a3533 := resource.Types

	if var_c4f1cb6a3533 != nil {
		var var_c4f1cb6a3533_mapped *structpb.Value

		var var_c4f1cb6a3533_err error
		var_c4f1cb6a3533_mapped, var_c4f1cb6a3533_err = types.ByResourcePropertyType(model.ResourceProperty_OBJECT).Pack(*var_c4f1cb6a3533)
		if var_c4f1cb6a3533_err != nil {
			panic(var_c4f1cb6a3533_err)
		}
		properties["types"] = var_c4f1cb6a3533_mapped
	}

	var_9c2aaec28a57 := resource.Immutable

	var var_9c2aaec28a57_mapped *structpb.Value

	var var_9c2aaec28a57_err error
	var_9c2aaec28a57_mapped, var_9c2aaec28a57_err = types.ByResourcePropertyType(model.ResourceProperty_BOOL).Pack(var_9c2aaec28a57)
	if var_9c2aaec28a57_err != nil {
		panic(var_9c2aaec28a57_err)
	}
	properties["immutable"] = var_9c2aaec28a57_mapped

	var_44f784ee02b8 := resource.Abstract

	var var_44f784ee02b8_mapped *structpb.Value

	var var_44f784ee02b8_err error
	var_44f784ee02b8_mapped, var_44f784ee02b8_err = types.ByResourcePropertyType(model.ResourceProperty_BOOL).Pack(var_44f784ee02b8)
	if var_44f784ee02b8_err != nil {
		panic(var_44f784ee02b8_err)
	}
	properties["abstract"] = var_44f784ee02b8_mapped

	var_3d21d189aacf := resource.DataSource

	if var_3d21d189aacf != nil {
		var var_3d21d189aacf_mapped *structpb.Value

		var_3d21d189aacf_mapped = structpb.NewStructValue(&structpb.Struct{Fields: DataSourceMapperInstance.ToProperties(var_3d21d189aacf)})
		properties["dataSource"] = var_3d21d189aacf_mapped
	}

	var_c15252bb4c06 := resource.Entity

	if var_c15252bb4c06 != nil {
		var var_c15252bb4c06_mapped *structpb.Value

		var var_c15252bb4c06_err error
		var_c15252bb4c06_mapped, var_c15252bb4c06_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_c15252bb4c06)
		if var_c15252bb4c06_err != nil {
			panic(var_c15252bb4c06_err)
		}
		properties["entity"] = var_c15252bb4c06_mapped
	}

	var_aa474762e334 := resource.Catalog

	if var_aa474762e334 != nil {
		var var_aa474762e334_mapped *structpb.Value

		var var_aa474762e334_err error
		var_aa474762e334_mapped, var_aa474762e334_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_aa474762e334)
		if var_aa474762e334_err != nil {
			panic(var_aa474762e334_err)
		}
		properties["catalog"] = var_aa474762e334_mapped
	}

	var_01417f5d8b92 := resource.Annotations

	if var_01417f5d8b92 != nil {
		var var_01417f5d8b92_mapped *structpb.Value

		var var_01417f5d8b92_st *structpb.Struct = new(structpb.Struct)
		var_01417f5d8b92_st.Fields = make(map[string]*structpb.Value)
		for key, value := range var_01417f5d8b92 {

			var_b32f16fe3e69 := value
			var var_b32f16fe3e69_mapped *structpb.Value

			var var_b32f16fe3e69_err error
			var_b32f16fe3e69_mapped, var_b32f16fe3e69_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_b32f16fe3e69)
			if var_b32f16fe3e69_err != nil {
				panic(var_b32f16fe3e69_err)
			}

			var_01417f5d8b92_st.Fields[key] = var_b32f16fe3e69_mapped
		}
		var_01417f5d8b92_mapped = structpb.NewStructValue(var_01417f5d8b92_st)
		properties["annotations"] = var_01417f5d8b92_mapped
	}

	var_4846d0918982 := resource.Indexes

	if var_4846d0918982 != nil {
		var var_4846d0918982_mapped *structpb.Value

		var var_4846d0918982_err error
		var_4846d0918982_mapped, var_4846d0918982_err = types.ByResourcePropertyType(model.ResourceProperty_OBJECT).Pack(*var_4846d0918982)
		if var_4846d0918982_err != nil {
			panic(var_4846d0918982_err)
		}
		properties["indexes"] = var_4846d0918982_mapped
	}

	var_dc60ccc46890 := resource.SecurityConstraints

	if var_dc60ccc46890 != nil {
		var var_dc60ccc46890_mapped *structpb.Value

		var var_dc60ccc46890_l []*structpb.Value
		for _, value := range var_dc60ccc46890 {

			var_18f93e78f86c := value
			var var_18f93e78f86c_mapped *structpb.Value

			var_18f93e78f86c_mapped = structpb.NewStructValue(&structpb.Struct{Fields: SecurityConstraintMapperInstance.ToProperties(var_18f93e78f86c)})

			var_dc60ccc46890_l = append(var_dc60ccc46890_l, var_18f93e78f86c_mapped)
		}
		var_dc60ccc46890_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_dc60ccc46890_l})
		properties["securityConstraints"] = var_dc60ccc46890_mapped
	}

	var_c937a710cfc5 := resource.Title

	if var_c937a710cfc5 != nil {
		var var_c937a710cfc5_mapped *structpb.Value

		var var_c937a710cfc5_err error
		var_c937a710cfc5_mapped, var_c937a710cfc5_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_c937a710cfc5)
		if var_c937a710cfc5_err != nil {
			panic(var_c937a710cfc5_err)
		}
		properties["title"] = var_c937a710cfc5_mapped
	}

	var_28c41a6b5fb9 := resource.Description

	if var_28c41a6b5fb9 != nil {
		var var_28c41a6b5fb9_mapped *structpb.Value

		var var_28c41a6b5fb9_err error
		var_28c41a6b5fb9_mapped, var_28c41a6b5fb9_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_28c41a6b5fb9)
		if var_28c41a6b5fb9_err != nil {
			panic(var_28c41a6b5fb9_err)
		}
		properties["description"] = var_28c41a6b5fb9_mapped
	}
	return properties
}

func (m *ResourceMapper) FromProperties(properties map[string]*structpb.Value) *Resource {
	var s = m.New()
	if properties["id"] != nil {

		var_c3e79cb70125 := properties["id"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_UUID).UnPack(var_c3e79cb70125)

		if err != nil {
			panic(err)
		}

		var_c3e79cb70125_mapped := new(uuid.UUID)
		*var_c3e79cb70125_mapped = val.(uuid.UUID)

		s.Id = var_c3e79cb70125_mapped
	}
	if properties["version"] != nil {

		var_df0c3d745df1 := properties["version"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(var_df0c3d745df1)

		if err != nil {
			panic(err)
		}

		var_df0c3d745df1_mapped := val.(int32)

		s.Version = var_df0c3d745df1_mapped
	}
	if properties["createdBy"] != nil {

		var_504659afbe73 := properties["createdBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_504659afbe73)

		if err != nil {
			panic(err)
		}

		var_504659afbe73_mapped := new(string)
		*var_504659afbe73_mapped = val.(string)

		s.CreatedBy = var_504659afbe73_mapped
	}
	if properties["updatedBy"] != nil {

		var_a05c5260c852 := properties["updatedBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_a05c5260c852)

		if err != nil {
			panic(err)
		}

		var_a05c5260c852_mapped := new(string)
		*var_a05c5260c852_mapped = val.(string)

		s.UpdatedBy = var_a05c5260c852_mapped
	}
	if properties["createdOn"] != nil {

		var_a8fb5fda9d89 := properties["createdOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_a8fb5fda9d89)

		if err != nil {
			panic(err)
		}

		var_a8fb5fda9d89_mapped := new(time.Time)
		*var_a8fb5fda9d89_mapped = val.(time.Time)

		s.CreatedOn = var_a8fb5fda9d89_mapped
	}
	if properties["updatedOn"] != nil {

		var_8d97b7caaf91 := properties["updatedOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_8d97b7caaf91)

		if err != nil {
			panic(err)
		}

		var_8d97b7caaf91_mapped := new(time.Time)
		*var_8d97b7caaf91_mapped = val.(time.Time)

		s.UpdatedOn = var_8d97b7caaf91_mapped
	}
	if properties["name"] != nil {

		var_2842a10fe410 := properties["name"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_2842a10fe410)

		if err != nil {
			panic(err)
		}

		var_2842a10fe410_mapped := val.(string)

		s.Name = var_2842a10fe410_mapped
	}
	if properties["namespace"] != nil {

		var_dd137f22ba10 := properties["namespace"]
		var_dd137f22ba10_mapped := NamespaceMapperInstance.FromProperties(var_dd137f22ba10.GetStructValue().Fields)

		s.Namespace = var_dd137f22ba10_mapped
	}
	if properties["virtual"] != nil {

		var_0ca56e1da045 := properties["virtual"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_BOOL).UnPack(var_0ca56e1da045)

		if err != nil {
			panic(err)
		}

		var_0ca56e1da045_mapped := val.(bool)

		s.Virtual = var_0ca56e1da045_mapped
	}
	if properties["types"] != nil {

		var_4fd9a4c853c6 := properties["types"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_OBJECT).UnPack(var_4fd9a4c853c6)

		if err != nil {
			panic(err)
		}

		var_4fd9a4c853c6_mapped := new(unstructured.Unstructured)
		*var_4fd9a4c853c6_mapped = val.(unstructured.Unstructured)

		s.Types = var_4fd9a4c853c6_mapped
	}
	if properties["immutable"] != nil {

		var_f54f2f2b6a9f := properties["immutable"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_BOOL).UnPack(var_f54f2f2b6a9f)

		if err != nil {
			panic(err)
		}

		var_f54f2f2b6a9f_mapped := val.(bool)

		s.Immutable = var_f54f2f2b6a9f_mapped
	}
	if properties["abstract"] != nil {

		var_4653fcb07aa2 := properties["abstract"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_BOOL).UnPack(var_4653fcb07aa2)

		if err != nil {
			panic(err)
		}

		var_4653fcb07aa2_mapped := val.(bool)

		s.Abstract = var_4653fcb07aa2_mapped
	}
	if properties["dataSource"] != nil {

		var_4c15c16ac5bb := properties["dataSource"]
		var_4c15c16ac5bb_mapped := DataSourceMapperInstance.FromProperties(var_4c15c16ac5bb.GetStructValue().Fields)

		s.DataSource = var_4c15c16ac5bb_mapped
	}
	if properties["entity"] != nil {

		var_a533e2685f08 := properties["entity"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_a533e2685f08)

		if err != nil {
			panic(err)
		}

		var_a533e2685f08_mapped := new(string)
		*var_a533e2685f08_mapped = val.(string)

		s.Entity = var_a533e2685f08_mapped
	}
	if properties["catalog"] != nil {

		var_6385fd949475 := properties["catalog"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_6385fd949475)

		if err != nil {
			panic(err)
		}

		var_6385fd949475_mapped := new(string)
		*var_6385fd949475_mapped = val.(string)

		s.Catalog = var_6385fd949475_mapped
	}
	if properties["annotations"] != nil {

		var_739ef10f05a6 := properties["annotations"]
		var_739ef10f05a6_mapped := make(map[string]string)
		for k, v := range var_739ef10f05a6.GetStructValue().Fields {

			var_2892ab8b6f73 := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_2892ab8b6f73)

			if err != nil {
				panic(err)
			}

			var_2892ab8b6f73_mapped := val.(string)

			var_739ef10f05a6_mapped[k] = var_2892ab8b6f73_mapped
		}

		s.Annotations = var_739ef10f05a6_mapped
	}
	if properties["indexes"] != nil {

		var_b670af9c874f := properties["indexes"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_OBJECT).UnPack(var_b670af9c874f)

		if err != nil {
			panic(err)
		}

		var_b670af9c874f_mapped := new(unstructured.Unstructured)
		*var_b670af9c874f_mapped = val.(unstructured.Unstructured)

		s.Indexes = var_b670af9c874f_mapped
	}
	if properties["securityConstraints"] != nil {

		var_f204c78a81f1 := properties["securityConstraints"]
		var_f204c78a81f1_mapped := []*SecurityConstraint{}
		for _, v := range var_f204c78a81f1.GetListValue().Values {

			var_048f1bb0d3fc := v
			var_048f1bb0d3fc_mapped := SecurityConstraintMapperInstance.FromProperties(var_048f1bb0d3fc.GetStructValue().Fields)

			var_f204c78a81f1_mapped = append(var_f204c78a81f1_mapped, var_048f1bb0d3fc_mapped)
		}

		s.SecurityConstraints = var_f204c78a81f1_mapped
	}
	if properties["title"] != nil {

		var_03a367fa7509 := properties["title"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_03a367fa7509)

		if err != nil {
			panic(err)
		}

		var_03a367fa7509_mapped := new(string)
		*var_03a367fa7509_mapped = val.(string)

		s.Title = var_03a367fa7509_mapped
	}
	if properties["description"] != nil {

		var_e39cc918c9d2 := properties["description"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_e39cc918c9d2)

		if err != nil {
			panic(err)
		}

		var_e39cc918c9d2_mapped := new(string)
		*var_e39cc918c9d2_mapped = val.(string)

		s.Description = var_e39cc918c9d2_mapped
	}
	return s
}
