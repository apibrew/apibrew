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

	var_db8650a065ca := resource.Id

	if var_db8650a065ca != nil {
		var var_db8650a065ca_mapped *structpb.Value

		var var_db8650a065ca_err error
		var_db8650a065ca_mapped, var_db8650a065ca_err = types.ByResourcePropertyType(model.ResourceProperty_UUID).Pack(*var_db8650a065ca)
		if var_db8650a065ca_err != nil {
			panic(var_db8650a065ca_err)
		}
		properties["id"] = var_db8650a065ca_mapped
	}

	var_f885afbe4d13 := resource.Version

	var var_f885afbe4d13_mapped *structpb.Value

	var var_f885afbe4d13_err error
	var_f885afbe4d13_mapped, var_f885afbe4d13_err = types.ByResourcePropertyType(model.ResourceProperty_INT32).Pack(var_f885afbe4d13)
	if var_f885afbe4d13_err != nil {
		panic(var_f885afbe4d13_err)
	}
	properties["version"] = var_f885afbe4d13_mapped

	var_24e1da8ccecd := resource.CreatedBy

	if var_24e1da8ccecd != nil {
		var var_24e1da8ccecd_mapped *structpb.Value

		var var_24e1da8ccecd_err error
		var_24e1da8ccecd_mapped, var_24e1da8ccecd_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_24e1da8ccecd)
		if var_24e1da8ccecd_err != nil {
			panic(var_24e1da8ccecd_err)
		}
		properties["createdBy"] = var_24e1da8ccecd_mapped
	}

	var_77c6138869d1 := resource.UpdatedBy

	if var_77c6138869d1 != nil {
		var var_77c6138869d1_mapped *structpb.Value

		var var_77c6138869d1_err error
		var_77c6138869d1_mapped, var_77c6138869d1_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_77c6138869d1)
		if var_77c6138869d1_err != nil {
			panic(var_77c6138869d1_err)
		}
		properties["updatedBy"] = var_77c6138869d1_mapped
	}

	var_db88e6b1240e := resource.CreatedOn

	if var_db88e6b1240e != nil {
		var var_db88e6b1240e_mapped *structpb.Value

		var var_db88e6b1240e_err error
		var_db88e6b1240e_mapped, var_db88e6b1240e_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_db88e6b1240e)
		if var_db88e6b1240e_err != nil {
			panic(var_db88e6b1240e_err)
		}
		properties["createdOn"] = var_db88e6b1240e_mapped
	}

	var_72f86d83c078 := resource.UpdatedOn

	if var_72f86d83c078 != nil {
		var var_72f86d83c078_mapped *structpb.Value

		var var_72f86d83c078_err error
		var_72f86d83c078_mapped, var_72f86d83c078_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_72f86d83c078)
		if var_72f86d83c078_err != nil {
			panic(var_72f86d83c078_err)
		}
		properties["updatedOn"] = var_72f86d83c078_mapped
	}

	var_8004488c2af1 := resource.Name

	var var_8004488c2af1_mapped *structpb.Value

	var var_8004488c2af1_err error
	var_8004488c2af1_mapped, var_8004488c2af1_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_8004488c2af1)
	if var_8004488c2af1_err != nil {
		panic(var_8004488c2af1_err)
	}
	properties["name"] = var_8004488c2af1_mapped

	var_aa84eea500e8 := resource.Namespace

	if var_aa84eea500e8 != nil {
		var var_aa84eea500e8_mapped *structpb.Value

		var_aa84eea500e8_mapped = structpb.NewStructValue(&structpb.Struct{Fields: NamespaceMapperInstance.ToProperties(var_aa84eea500e8)})
		properties["namespace"] = var_aa84eea500e8_mapped
	}

	var_acda2de54918 := resource.Virtual

	var var_acda2de54918_mapped *structpb.Value

	var var_acda2de54918_err error
	var_acda2de54918_mapped, var_acda2de54918_err = types.ByResourcePropertyType(model.ResourceProperty_BOOL).Pack(var_acda2de54918)
	if var_acda2de54918_err != nil {
		panic(var_acda2de54918_err)
	}
	properties["virtual"] = var_acda2de54918_mapped

	var_5c3e22295e54 := resource.Types

	if var_5c3e22295e54 != nil {
		var var_5c3e22295e54_mapped *structpb.Value

		var var_5c3e22295e54_err error
		var_5c3e22295e54_mapped, var_5c3e22295e54_err = types.ByResourcePropertyType(model.ResourceProperty_OBJECT).Pack(*var_5c3e22295e54)
		if var_5c3e22295e54_err != nil {
			panic(var_5c3e22295e54_err)
		}
		properties["types"] = var_5c3e22295e54_mapped
	}

	var_26650ddd487e := resource.Immutable

	var var_26650ddd487e_mapped *structpb.Value

	var var_26650ddd487e_err error
	var_26650ddd487e_mapped, var_26650ddd487e_err = types.ByResourcePropertyType(model.ResourceProperty_BOOL).Pack(var_26650ddd487e)
	if var_26650ddd487e_err != nil {
		panic(var_26650ddd487e_err)
	}
	properties["immutable"] = var_26650ddd487e_mapped

	var_2a4c04e72bc5 := resource.Abstract

	var var_2a4c04e72bc5_mapped *structpb.Value

	var var_2a4c04e72bc5_err error
	var_2a4c04e72bc5_mapped, var_2a4c04e72bc5_err = types.ByResourcePropertyType(model.ResourceProperty_BOOL).Pack(var_2a4c04e72bc5)
	if var_2a4c04e72bc5_err != nil {
		panic(var_2a4c04e72bc5_err)
	}
	properties["abstract"] = var_2a4c04e72bc5_mapped

	var_0306debfd64c := resource.DataSource

	if var_0306debfd64c != nil {
		var var_0306debfd64c_mapped *structpb.Value

		var_0306debfd64c_mapped = structpb.NewStructValue(&structpb.Struct{Fields: DataSourceMapperInstance.ToProperties(var_0306debfd64c)})
		properties["dataSource"] = var_0306debfd64c_mapped
	}

	var_74f4ce7fb42c := resource.Entity

	if var_74f4ce7fb42c != nil {
		var var_74f4ce7fb42c_mapped *structpb.Value

		var var_74f4ce7fb42c_err error
		var_74f4ce7fb42c_mapped, var_74f4ce7fb42c_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_74f4ce7fb42c)
		if var_74f4ce7fb42c_err != nil {
			panic(var_74f4ce7fb42c_err)
		}
		properties["entity"] = var_74f4ce7fb42c_mapped
	}

	var_f2ce9e092355 := resource.Catalog

	if var_f2ce9e092355 != nil {
		var var_f2ce9e092355_mapped *structpb.Value

		var var_f2ce9e092355_err error
		var_f2ce9e092355_mapped, var_f2ce9e092355_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_f2ce9e092355)
		if var_f2ce9e092355_err != nil {
			panic(var_f2ce9e092355_err)
		}
		properties["catalog"] = var_f2ce9e092355_mapped
	}

	var_4b7cb81638fb := resource.Annotations

	if var_4b7cb81638fb != nil {
		var var_4b7cb81638fb_mapped *structpb.Value

		var var_4b7cb81638fb_st *structpb.Struct = new(structpb.Struct)
		var_4b7cb81638fb_st.Fields = make(map[string]*structpb.Value)
		for key, value := range var_4b7cb81638fb {

			var_9e7744412b98 := value
			var var_9e7744412b98_mapped *structpb.Value

			var var_9e7744412b98_err error
			var_9e7744412b98_mapped, var_9e7744412b98_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_9e7744412b98)
			if var_9e7744412b98_err != nil {
				panic(var_9e7744412b98_err)
			}

			var_4b7cb81638fb_st.Fields[key] = var_9e7744412b98_mapped
		}
		var_4b7cb81638fb_mapped = structpb.NewStructValue(var_4b7cb81638fb_st)
		properties["annotations"] = var_4b7cb81638fb_mapped
	}

	var_e5b3db0c9ebf := resource.Indexes

	if var_e5b3db0c9ebf != nil {
		var var_e5b3db0c9ebf_mapped *structpb.Value

		var var_e5b3db0c9ebf_err error
		var_e5b3db0c9ebf_mapped, var_e5b3db0c9ebf_err = types.ByResourcePropertyType(model.ResourceProperty_OBJECT).Pack(*var_e5b3db0c9ebf)
		if var_e5b3db0c9ebf_err != nil {
			panic(var_e5b3db0c9ebf_err)
		}
		properties["indexes"] = var_e5b3db0c9ebf_mapped
	}

	var_e3e2c393ac79 := resource.Title

	if var_e3e2c393ac79 != nil {
		var var_e3e2c393ac79_mapped *structpb.Value

		var var_e3e2c393ac79_err error
		var_e3e2c393ac79_mapped, var_e3e2c393ac79_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_e3e2c393ac79)
		if var_e3e2c393ac79_err != nil {
			panic(var_e3e2c393ac79_err)
		}
		properties["title"] = var_e3e2c393ac79_mapped
	}

	var_118e87bb82d9 := resource.Description

	if var_118e87bb82d9 != nil {
		var var_118e87bb82d9_mapped *structpb.Value

		var var_118e87bb82d9_err error
		var_118e87bb82d9_mapped, var_118e87bb82d9_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_118e87bb82d9)
		if var_118e87bb82d9_err != nil {
			panic(var_118e87bb82d9_err)
		}
		properties["description"] = var_118e87bb82d9_mapped
	}
	return properties
}

func (m *ResourceMapper) FromProperties(properties map[string]*structpb.Value) *Resource {
	var s = m.New()
	if properties["id"] != nil {

		var_d1a8ba3f3c97 := properties["id"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_UUID).UnPack(var_d1a8ba3f3c97)

		if err != nil {
			panic(err)
		}

		var_d1a8ba3f3c97_mapped := new(uuid.UUID)
		*var_d1a8ba3f3c97_mapped = val.(uuid.UUID)

		s.Id = var_d1a8ba3f3c97_mapped
	}
	if properties["version"] != nil {

		var_e24909e87c8e := properties["version"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(var_e24909e87c8e)

		if err != nil {
			panic(err)
		}

		var_e24909e87c8e_mapped := val.(int32)

		s.Version = var_e24909e87c8e_mapped
	}
	if properties["createdBy"] != nil {

		var_b9dbda4a4c9b := properties["createdBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_b9dbda4a4c9b)

		if err != nil {
			panic(err)
		}

		var_b9dbda4a4c9b_mapped := new(string)
		*var_b9dbda4a4c9b_mapped = val.(string)

		s.CreatedBy = var_b9dbda4a4c9b_mapped
	}
	if properties["updatedBy"] != nil {

		var_1e5d6fe044c9 := properties["updatedBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_1e5d6fe044c9)

		if err != nil {
			panic(err)
		}

		var_1e5d6fe044c9_mapped := new(string)
		*var_1e5d6fe044c9_mapped = val.(string)

		s.UpdatedBy = var_1e5d6fe044c9_mapped
	}
	if properties["createdOn"] != nil {

		var_cba3c083f806 := properties["createdOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_cba3c083f806)

		if err != nil {
			panic(err)
		}

		var_cba3c083f806_mapped := new(time.Time)
		*var_cba3c083f806_mapped = val.(time.Time)

		s.CreatedOn = var_cba3c083f806_mapped
	}
	if properties["updatedOn"] != nil {

		var_b1ff09e07561 := properties["updatedOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_b1ff09e07561)

		if err != nil {
			panic(err)
		}

		var_b1ff09e07561_mapped := new(time.Time)
		*var_b1ff09e07561_mapped = val.(time.Time)

		s.UpdatedOn = var_b1ff09e07561_mapped
	}
	if properties["name"] != nil {

		var_ab7621458b30 := properties["name"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_ab7621458b30)

		if err != nil {
			panic(err)
		}

		var_ab7621458b30_mapped := val.(string)

		s.Name = var_ab7621458b30_mapped
	}
	if properties["namespace"] != nil {

		var_c8c6b97e4ec1 := properties["namespace"]
		var_c8c6b97e4ec1_mapped := NamespaceMapperInstance.FromProperties(var_c8c6b97e4ec1.GetStructValue().Fields)

		s.Namespace = var_c8c6b97e4ec1_mapped
	}
	if properties["virtual"] != nil {

		var_ff3dee4278fe := properties["virtual"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_BOOL).UnPack(var_ff3dee4278fe)

		if err != nil {
			panic(err)
		}

		var_ff3dee4278fe_mapped := val.(bool)

		s.Virtual = var_ff3dee4278fe_mapped
	}
	if properties["types"] != nil {

		var_83085100a93c := properties["types"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_OBJECT).UnPack(var_83085100a93c)

		if err != nil {
			panic(err)
		}

		var_83085100a93c_mapped := new(unstructured.Unstructured)
		*var_83085100a93c_mapped = val.(unstructured.Unstructured)

		s.Types = var_83085100a93c_mapped
	}
	if properties["immutable"] != nil {

		var_0861a235d7ce := properties["immutable"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_BOOL).UnPack(var_0861a235d7ce)

		if err != nil {
			panic(err)
		}

		var_0861a235d7ce_mapped := val.(bool)

		s.Immutable = var_0861a235d7ce_mapped
	}
	if properties["abstract"] != nil {

		var_c9f3fef08b63 := properties["abstract"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_BOOL).UnPack(var_c9f3fef08b63)

		if err != nil {
			panic(err)
		}

		var_c9f3fef08b63_mapped := val.(bool)

		s.Abstract = var_c9f3fef08b63_mapped
	}
	if properties["dataSource"] != nil {

		var_a36443b55dda := properties["dataSource"]
		var_a36443b55dda_mapped := DataSourceMapperInstance.FromProperties(var_a36443b55dda.GetStructValue().Fields)

		s.DataSource = var_a36443b55dda_mapped
	}
	if properties["entity"] != nil {

		var_61252b577774 := properties["entity"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_61252b577774)

		if err != nil {
			panic(err)
		}

		var_61252b577774_mapped := new(string)
		*var_61252b577774_mapped = val.(string)

		s.Entity = var_61252b577774_mapped
	}
	if properties["catalog"] != nil {

		var_fa9dfc75a09b := properties["catalog"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_fa9dfc75a09b)

		if err != nil {
			panic(err)
		}

		var_fa9dfc75a09b_mapped := new(string)
		*var_fa9dfc75a09b_mapped = val.(string)

		s.Catalog = var_fa9dfc75a09b_mapped
	}
	if properties["annotations"] != nil {

		var_02c25417342d := properties["annotations"]
		var_02c25417342d_mapped := make(map[string]string)
		for k, v := range var_02c25417342d.GetStructValue().Fields {

			var_b9c55edcc23b := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_b9c55edcc23b)

			if err != nil {
				panic(err)
			}

			var_b9c55edcc23b_mapped := val.(string)

			var_02c25417342d_mapped[k] = var_b9c55edcc23b_mapped
		}

		s.Annotations = var_02c25417342d_mapped
	}
	if properties["indexes"] != nil {

		var_166796135b1d := properties["indexes"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_OBJECT).UnPack(var_166796135b1d)

		if err != nil {
			panic(err)
		}

		var_166796135b1d_mapped := new(unstructured.Unstructured)
		*var_166796135b1d_mapped = val.(unstructured.Unstructured)

		s.Indexes = var_166796135b1d_mapped
	}
	if properties["title"] != nil {

		var_3110ab5a88ce := properties["title"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_3110ab5a88ce)

		if err != nil {
			panic(err)
		}

		var_3110ab5a88ce_mapped := new(string)
		*var_3110ab5a88ce_mapped = val.(string)

		s.Title = var_3110ab5a88ce_mapped
	}
	if properties["description"] != nil {

		var_f64682246686 := properties["description"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_f64682246686)

		if err != nil {
			panic(err)
		}

		var_f64682246686_mapped := new(string)
		*var_f64682246686_mapped = val.(string)

		s.Description = var_f64682246686_mapped
	}
	return s
}
