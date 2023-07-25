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

	var_2efa3b014a72 := resource.Id

	if var_2efa3b014a72 != nil {
		var var_2efa3b014a72_mapped *structpb.Value

		var var_2efa3b014a72_err error
		var_2efa3b014a72_mapped, var_2efa3b014a72_err = types.ByResourcePropertyType(model.ResourceProperty_UUID).Pack(*var_2efa3b014a72)
		if var_2efa3b014a72_err != nil {
			panic(var_2efa3b014a72_err)
		}
		properties["id"] = var_2efa3b014a72_mapped
	}

	var_0c63a20236a1 := resource.Version

	var var_0c63a20236a1_mapped *structpb.Value

	var var_0c63a20236a1_err error
	var_0c63a20236a1_mapped, var_0c63a20236a1_err = types.ByResourcePropertyType(model.ResourceProperty_INT32).Pack(var_0c63a20236a1)
	if var_0c63a20236a1_err != nil {
		panic(var_0c63a20236a1_err)
	}
	properties["version"] = var_0c63a20236a1_mapped

	var_c0cc10f7d606 := resource.CreatedBy

	if var_c0cc10f7d606 != nil {
		var var_c0cc10f7d606_mapped *structpb.Value

		var var_c0cc10f7d606_err error
		var_c0cc10f7d606_mapped, var_c0cc10f7d606_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_c0cc10f7d606)
		if var_c0cc10f7d606_err != nil {
			panic(var_c0cc10f7d606_err)
		}
		properties["createdBy"] = var_c0cc10f7d606_mapped
	}

	var_4acb365bac66 := resource.UpdatedBy

	if var_4acb365bac66 != nil {
		var var_4acb365bac66_mapped *structpb.Value

		var var_4acb365bac66_err error
		var_4acb365bac66_mapped, var_4acb365bac66_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_4acb365bac66)
		if var_4acb365bac66_err != nil {
			panic(var_4acb365bac66_err)
		}
		properties["updatedBy"] = var_4acb365bac66_mapped
	}

	var_73309288ac37 := resource.CreatedOn

	if var_73309288ac37 != nil {
		var var_73309288ac37_mapped *structpb.Value

		var var_73309288ac37_err error
		var_73309288ac37_mapped, var_73309288ac37_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_73309288ac37)
		if var_73309288ac37_err != nil {
			panic(var_73309288ac37_err)
		}
		properties["createdOn"] = var_73309288ac37_mapped
	}

	var_acc3900db56e := resource.UpdatedOn

	if var_acc3900db56e != nil {
		var var_acc3900db56e_mapped *structpb.Value

		var var_acc3900db56e_err error
		var_acc3900db56e_mapped, var_acc3900db56e_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_acc3900db56e)
		if var_acc3900db56e_err != nil {
			panic(var_acc3900db56e_err)
		}
		properties["updatedOn"] = var_acc3900db56e_mapped
	}

	var_6d69f98df14b := resource.Name

	var var_6d69f98df14b_mapped *structpb.Value

	var var_6d69f98df14b_err error
	var_6d69f98df14b_mapped, var_6d69f98df14b_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_6d69f98df14b)
	if var_6d69f98df14b_err != nil {
		panic(var_6d69f98df14b_err)
	}
	properties["name"] = var_6d69f98df14b_mapped

	var_51f4b8055d57 := resource.Namespace

	if var_51f4b8055d57 != nil {
		var var_51f4b8055d57_mapped *structpb.Value

		var_51f4b8055d57_mapped = structpb.NewStructValue(&structpb.Struct{Fields: NamespaceMapperInstance.ToProperties(var_51f4b8055d57)})
		properties["namespace"] = var_51f4b8055d57_mapped
	}

	var_bee160ad21d9 := resource.Virtual

	var var_bee160ad21d9_mapped *structpb.Value

	var var_bee160ad21d9_err error
	var_bee160ad21d9_mapped, var_bee160ad21d9_err = types.ByResourcePropertyType(model.ResourceProperty_BOOL).Pack(var_bee160ad21d9)
	if var_bee160ad21d9_err != nil {
		panic(var_bee160ad21d9_err)
	}
	properties["virtual"] = var_bee160ad21d9_mapped

	var_0bb92e4ed93d := resource.Types

	if var_0bb92e4ed93d != nil {
		var var_0bb92e4ed93d_mapped *structpb.Value

		var var_0bb92e4ed93d_err error
		var_0bb92e4ed93d_mapped, var_0bb92e4ed93d_err = types.ByResourcePropertyType(model.ResourceProperty_OBJECT).Pack(*var_0bb92e4ed93d)
		if var_0bb92e4ed93d_err != nil {
			panic(var_0bb92e4ed93d_err)
		}
		properties["types"] = var_0bb92e4ed93d_mapped
	}

	var_6beffb095b7a := resource.Immutable

	var var_6beffb095b7a_mapped *structpb.Value

	var var_6beffb095b7a_err error
	var_6beffb095b7a_mapped, var_6beffb095b7a_err = types.ByResourcePropertyType(model.ResourceProperty_BOOL).Pack(var_6beffb095b7a)
	if var_6beffb095b7a_err != nil {
		panic(var_6beffb095b7a_err)
	}
	properties["immutable"] = var_6beffb095b7a_mapped

	var_5da0fd6ec839 := resource.Abstract

	var var_5da0fd6ec839_mapped *structpb.Value

	var var_5da0fd6ec839_err error
	var_5da0fd6ec839_mapped, var_5da0fd6ec839_err = types.ByResourcePropertyType(model.ResourceProperty_BOOL).Pack(var_5da0fd6ec839)
	if var_5da0fd6ec839_err != nil {
		panic(var_5da0fd6ec839_err)
	}
	properties["abstract"] = var_5da0fd6ec839_mapped

	var_8d06b7dc2b79 := resource.DataSource

	if var_8d06b7dc2b79 != nil {
		var var_8d06b7dc2b79_mapped *structpb.Value

		var_8d06b7dc2b79_mapped = structpb.NewStructValue(&structpb.Struct{Fields: DataSourceMapperInstance.ToProperties(var_8d06b7dc2b79)})
		properties["dataSource"] = var_8d06b7dc2b79_mapped
	}

	var_e70f9dd7f058 := resource.Entity

	if var_e70f9dd7f058 != nil {
		var var_e70f9dd7f058_mapped *structpb.Value

		var var_e70f9dd7f058_err error
		var_e70f9dd7f058_mapped, var_e70f9dd7f058_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_e70f9dd7f058)
		if var_e70f9dd7f058_err != nil {
			panic(var_e70f9dd7f058_err)
		}
		properties["entity"] = var_e70f9dd7f058_mapped
	}

	var_597fbdbed2c4 := resource.Catalog

	if var_597fbdbed2c4 != nil {
		var var_597fbdbed2c4_mapped *structpb.Value

		var var_597fbdbed2c4_err error
		var_597fbdbed2c4_mapped, var_597fbdbed2c4_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_597fbdbed2c4)
		if var_597fbdbed2c4_err != nil {
			panic(var_597fbdbed2c4_err)
		}
		properties["catalog"] = var_597fbdbed2c4_mapped
	}

	var_5b89f14fd392 := resource.Annotations

	if var_5b89f14fd392 != nil {
		var var_5b89f14fd392_mapped *structpb.Value

		var var_5b89f14fd392_st *structpb.Struct = new(structpb.Struct)
		var_5b89f14fd392_st.Fields = make(map[string]*structpb.Value)
		for key, value := range var_5b89f14fd392 {

			var_429af064ad34 := value
			var var_429af064ad34_mapped *structpb.Value

			var var_429af064ad34_err error
			var_429af064ad34_mapped, var_429af064ad34_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_429af064ad34)
			if var_429af064ad34_err != nil {
				panic(var_429af064ad34_err)
			}

			var_5b89f14fd392_st.Fields[key] = var_429af064ad34_mapped
		}
		var_5b89f14fd392_mapped = structpb.NewStructValue(var_5b89f14fd392_st)
		properties["annotations"] = var_5b89f14fd392_mapped
	}

	var_487c14d791f7 := resource.Indexes

	if var_487c14d791f7 != nil {
		var var_487c14d791f7_mapped *structpb.Value

		var var_487c14d791f7_err error
		var_487c14d791f7_mapped, var_487c14d791f7_err = types.ByResourcePropertyType(model.ResourceProperty_OBJECT).Pack(*var_487c14d791f7)
		if var_487c14d791f7_err != nil {
			panic(var_487c14d791f7_err)
		}
		properties["indexes"] = var_487c14d791f7_mapped
	}

	var_e55c91e5f994 := resource.Title

	if var_e55c91e5f994 != nil {
		var var_e55c91e5f994_mapped *structpb.Value

		var var_e55c91e5f994_err error
		var_e55c91e5f994_mapped, var_e55c91e5f994_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_e55c91e5f994)
		if var_e55c91e5f994_err != nil {
			panic(var_e55c91e5f994_err)
		}
		properties["title"] = var_e55c91e5f994_mapped
	}

	var_6e2e499978b5 := resource.Description

	if var_6e2e499978b5 != nil {
		var var_6e2e499978b5_mapped *structpb.Value

		var var_6e2e499978b5_err error
		var_6e2e499978b5_mapped, var_6e2e499978b5_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_6e2e499978b5)
		if var_6e2e499978b5_err != nil {
			panic(var_6e2e499978b5_err)
		}
		properties["description"] = var_6e2e499978b5_mapped
	}
	return properties
}

func (m *ResourceMapper) FromProperties(properties map[string]*structpb.Value) *Resource {
	var s = m.New()
	if properties["id"] != nil {

		var_aa1c4d9c2afb := properties["id"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_UUID).UnPack(var_aa1c4d9c2afb)

		if err != nil {
			panic(err)
		}

		var_aa1c4d9c2afb_mapped := new(uuid.UUID)
		*var_aa1c4d9c2afb_mapped = val.(uuid.UUID)

		s.Id = var_aa1c4d9c2afb_mapped
	}
	if properties["version"] != nil {

		var_4ec2d0176c06 := properties["version"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(var_4ec2d0176c06)

		if err != nil {
			panic(err)
		}

		var_4ec2d0176c06_mapped := val.(int32)

		s.Version = var_4ec2d0176c06_mapped
	}
	if properties["createdBy"] != nil {

		var_efe6901246bd := properties["createdBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_efe6901246bd)

		if err != nil {
			panic(err)
		}

		var_efe6901246bd_mapped := new(string)
		*var_efe6901246bd_mapped = val.(string)

		s.CreatedBy = var_efe6901246bd_mapped
	}
	if properties["updatedBy"] != nil {

		var_3f7a26a69762 := properties["updatedBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_3f7a26a69762)

		if err != nil {
			panic(err)
		}

		var_3f7a26a69762_mapped := new(string)
		*var_3f7a26a69762_mapped = val.(string)

		s.UpdatedBy = var_3f7a26a69762_mapped
	}
	if properties["createdOn"] != nil {

		var_069e85bce294 := properties["createdOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_069e85bce294)

		if err != nil {
			panic(err)
		}

		var_069e85bce294_mapped := new(time.Time)
		*var_069e85bce294_mapped = val.(time.Time)

		s.CreatedOn = var_069e85bce294_mapped
	}
	if properties["updatedOn"] != nil {

		var_52ba20c4f92a := properties["updatedOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_52ba20c4f92a)

		if err != nil {
			panic(err)
		}

		var_52ba20c4f92a_mapped := new(time.Time)
		*var_52ba20c4f92a_mapped = val.(time.Time)

		s.UpdatedOn = var_52ba20c4f92a_mapped
	}
	if properties["name"] != nil {

		var_9e3eb1ae1d1f := properties["name"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_9e3eb1ae1d1f)

		if err != nil {
			panic(err)
		}

		var_9e3eb1ae1d1f_mapped := val.(string)

		s.Name = var_9e3eb1ae1d1f_mapped
	}
	if properties["namespace"] != nil {

		var_ce0e946b2aff := properties["namespace"]
		var_ce0e946b2aff_mapped := NamespaceMapperInstance.FromProperties(var_ce0e946b2aff.GetStructValue().Fields)

		s.Namespace = var_ce0e946b2aff_mapped
	}
	if properties["virtual"] != nil {

		var_ba9ba783d109 := properties["virtual"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_BOOL).UnPack(var_ba9ba783d109)

		if err != nil {
			panic(err)
		}

		var_ba9ba783d109_mapped := val.(bool)

		s.Virtual = var_ba9ba783d109_mapped
	}
	if properties["types"] != nil {

		var_0f9646e97b67 := properties["types"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_OBJECT).UnPack(var_0f9646e97b67)

		if err != nil {
			panic(err)
		}

		var_0f9646e97b67_mapped := new(unstructured.Unstructured)
		*var_0f9646e97b67_mapped = val.(unstructured.Unstructured)

		s.Types = var_0f9646e97b67_mapped
	}
	if properties["immutable"] != nil {

		var_8dbebf1fe01d := properties["immutable"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_BOOL).UnPack(var_8dbebf1fe01d)

		if err != nil {
			panic(err)
		}

		var_8dbebf1fe01d_mapped := val.(bool)

		s.Immutable = var_8dbebf1fe01d_mapped
	}
	if properties["abstract"] != nil {

		var_ee73ca41d99b := properties["abstract"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_BOOL).UnPack(var_ee73ca41d99b)

		if err != nil {
			panic(err)
		}

		var_ee73ca41d99b_mapped := val.(bool)

		s.Abstract = var_ee73ca41d99b_mapped
	}
	if properties["dataSource"] != nil {

		var_cbeed03c10d1 := properties["dataSource"]
		var_cbeed03c10d1_mapped := DataSourceMapperInstance.FromProperties(var_cbeed03c10d1.GetStructValue().Fields)

		s.DataSource = var_cbeed03c10d1_mapped
	}
	if properties["entity"] != nil {

		var_a8c3ea0f2c25 := properties["entity"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_a8c3ea0f2c25)

		if err != nil {
			panic(err)
		}

		var_a8c3ea0f2c25_mapped := new(string)
		*var_a8c3ea0f2c25_mapped = val.(string)

		s.Entity = var_a8c3ea0f2c25_mapped
	}
	if properties["catalog"] != nil {

		var_773e6db3904b := properties["catalog"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_773e6db3904b)

		if err != nil {
			panic(err)
		}

		var_773e6db3904b_mapped := new(string)
		*var_773e6db3904b_mapped = val.(string)

		s.Catalog = var_773e6db3904b_mapped
	}
	if properties["annotations"] != nil {

		var_1debaea24188 := properties["annotations"]
		var_1debaea24188_mapped := make(map[string]string)
		for k, v := range var_1debaea24188.GetStructValue().Fields {

			var_d9ff18780c1f := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_d9ff18780c1f)

			if err != nil {
				panic(err)
			}

			var_d9ff18780c1f_mapped := val.(string)

			var_1debaea24188_mapped[k] = var_d9ff18780c1f_mapped
		}

		s.Annotations = var_1debaea24188_mapped
	}
	if properties["indexes"] != nil {

		var_0e36ccd929a2 := properties["indexes"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_OBJECT).UnPack(var_0e36ccd929a2)

		if err != nil {
			panic(err)
		}

		var_0e36ccd929a2_mapped := new(unstructured.Unstructured)
		*var_0e36ccd929a2_mapped = val.(unstructured.Unstructured)

		s.Indexes = var_0e36ccd929a2_mapped
	}
	if properties["title"] != nil {

		var_af2c19216384 := properties["title"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_af2c19216384)

		if err != nil {
			panic(err)
		}

		var_af2c19216384_mapped := new(string)
		*var_af2c19216384_mapped = val.(string)

		s.Title = var_af2c19216384_mapped
	}
	if properties["description"] != nil {

		var_7d6b59a8d225 := properties["description"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_7d6b59a8d225)

		if err != nil {
			panic(err)
		}

		var_7d6b59a8d225_mapped := new(string)
		*var_7d6b59a8d225_mapped = val.(string)

		s.Description = var_7d6b59a8d225_mapped
	}
	return s
}
