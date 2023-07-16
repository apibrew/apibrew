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

	var_aaaf2881d40d := resource.Id

	if var_aaaf2881d40d != nil {
		var var_aaaf2881d40d_mapped *structpb.Value

		var var_aaaf2881d40d_err error
		var_aaaf2881d40d_mapped, var_aaaf2881d40d_err = types.ByResourcePropertyType(model.ResourceProperty_UUID).Pack(*var_aaaf2881d40d)
		if var_aaaf2881d40d_err != nil {
			panic(var_aaaf2881d40d_err)
		}
		properties["id"] = var_aaaf2881d40d_mapped
	}

	var_2246b7ff20f5 := resource.Version

	var var_2246b7ff20f5_mapped *structpb.Value

	var var_2246b7ff20f5_err error
	var_2246b7ff20f5_mapped, var_2246b7ff20f5_err = types.ByResourcePropertyType(model.ResourceProperty_INT32).Pack(var_2246b7ff20f5)
	if var_2246b7ff20f5_err != nil {
		panic(var_2246b7ff20f5_err)
	}
	properties["version"] = var_2246b7ff20f5_mapped

	var_8c785d1d0fee := resource.CreatedBy

	if var_8c785d1d0fee != nil {
		var var_8c785d1d0fee_mapped *structpb.Value

		var var_8c785d1d0fee_err error
		var_8c785d1d0fee_mapped, var_8c785d1d0fee_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_8c785d1d0fee)
		if var_8c785d1d0fee_err != nil {
			panic(var_8c785d1d0fee_err)
		}
		properties["createdBy"] = var_8c785d1d0fee_mapped
	}

	var_014c5e640dcc := resource.UpdatedBy

	if var_014c5e640dcc != nil {
		var var_014c5e640dcc_mapped *structpb.Value

		var var_014c5e640dcc_err error
		var_014c5e640dcc_mapped, var_014c5e640dcc_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_014c5e640dcc)
		if var_014c5e640dcc_err != nil {
			panic(var_014c5e640dcc_err)
		}
		properties["updatedBy"] = var_014c5e640dcc_mapped
	}

	var_d297677eabaa := resource.CreatedOn

	if var_d297677eabaa != nil {
		var var_d297677eabaa_mapped *structpb.Value

		var var_d297677eabaa_err error
		var_d297677eabaa_mapped, var_d297677eabaa_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_d297677eabaa)
		if var_d297677eabaa_err != nil {
			panic(var_d297677eabaa_err)
		}
		properties["createdOn"] = var_d297677eabaa_mapped
	}

	var_9e39b645020a := resource.UpdatedOn

	if var_9e39b645020a != nil {
		var var_9e39b645020a_mapped *structpb.Value

		var var_9e39b645020a_err error
		var_9e39b645020a_mapped, var_9e39b645020a_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_9e39b645020a)
		if var_9e39b645020a_err != nil {
			panic(var_9e39b645020a_err)
		}
		properties["updatedOn"] = var_9e39b645020a_mapped
	}

	var_158ebfa8df44 := resource.Name

	var var_158ebfa8df44_mapped *structpb.Value

	var var_158ebfa8df44_err error
	var_158ebfa8df44_mapped, var_158ebfa8df44_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_158ebfa8df44)
	if var_158ebfa8df44_err != nil {
		panic(var_158ebfa8df44_err)
	}
	properties["name"] = var_158ebfa8df44_mapped

	var_6d290d35777f := resource.Namespace

	if var_6d290d35777f != nil {
		var var_6d290d35777f_mapped *structpb.Value

		var_6d290d35777f_mapped = structpb.NewStructValue(&structpb.Struct{Fields: NamespaceMapperInstance.ToProperties(var_6d290d35777f)})
		properties["namespace"] = var_6d290d35777f_mapped
	}

	var_2524ce2f66af := resource.Virtual

	var var_2524ce2f66af_mapped *structpb.Value

	var var_2524ce2f66af_err error
	var_2524ce2f66af_mapped, var_2524ce2f66af_err = types.ByResourcePropertyType(model.ResourceProperty_BOOL).Pack(var_2524ce2f66af)
	if var_2524ce2f66af_err != nil {
		panic(var_2524ce2f66af_err)
	}
	properties["virtual"] = var_2524ce2f66af_mapped

	var_b08164ff55c9 := resource.Types

	if var_b08164ff55c9 != nil {
		var var_b08164ff55c9_mapped *structpb.Value

		var var_b08164ff55c9_err error
		var_b08164ff55c9_mapped, var_b08164ff55c9_err = types.ByResourcePropertyType(model.ResourceProperty_OBJECT).Pack(*var_b08164ff55c9)
		if var_b08164ff55c9_err != nil {
			panic(var_b08164ff55c9_err)
		}
		properties["types"] = var_b08164ff55c9_mapped
	}

	var_842768ccaf18 := resource.Immutable

	var var_842768ccaf18_mapped *structpb.Value

	var var_842768ccaf18_err error
	var_842768ccaf18_mapped, var_842768ccaf18_err = types.ByResourcePropertyType(model.ResourceProperty_BOOL).Pack(var_842768ccaf18)
	if var_842768ccaf18_err != nil {
		panic(var_842768ccaf18_err)
	}
	properties["immutable"] = var_842768ccaf18_mapped

	var_bf4a69f3997a := resource.Abstract

	var var_bf4a69f3997a_mapped *structpb.Value

	var var_bf4a69f3997a_err error
	var_bf4a69f3997a_mapped, var_bf4a69f3997a_err = types.ByResourcePropertyType(model.ResourceProperty_BOOL).Pack(var_bf4a69f3997a)
	if var_bf4a69f3997a_err != nil {
		panic(var_bf4a69f3997a_err)
	}
	properties["abstract"] = var_bf4a69f3997a_mapped

	var_f9a1f630f132 := resource.DataSource

	if var_f9a1f630f132 != nil {
		var var_f9a1f630f132_mapped *structpb.Value

		var_f9a1f630f132_mapped = structpb.NewStructValue(&structpb.Struct{Fields: DataSourceMapperInstance.ToProperties(var_f9a1f630f132)})
		properties["dataSource"] = var_f9a1f630f132_mapped
	}

	var_23bc76a641ba := resource.Entity

	if var_23bc76a641ba != nil {
		var var_23bc76a641ba_mapped *structpb.Value

		var var_23bc76a641ba_err error
		var_23bc76a641ba_mapped, var_23bc76a641ba_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_23bc76a641ba)
		if var_23bc76a641ba_err != nil {
			panic(var_23bc76a641ba_err)
		}
		properties["entity"] = var_23bc76a641ba_mapped
	}

	var_1662c279f8e4 := resource.Catalog

	if var_1662c279f8e4 != nil {
		var var_1662c279f8e4_mapped *structpb.Value

		var var_1662c279f8e4_err error
		var_1662c279f8e4_mapped, var_1662c279f8e4_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_1662c279f8e4)
		if var_1662c279f8e4_err != nil {
			panic(var_1662c279f8e4_err)
		}
		properties["catalog"] = var_1662c279f8e4_mapped
	}

	var_b2fcc1dcfe4f := resource.Annotations

	if var_b2fcc1dcfe4f != nil {
		var var_b2fcc1dcfe4f_mapped *structpb.Value

		var var_b2fcc1dcfe4f_st *structpb.Struct = new(structpb.Struct)
		var_b2fcc1dcfe4f_st.Fields = make(map[string]*structpb.Value)
		for key, value := range var_b2fcc1dcfe4f {

			var_f36ac47c3a06 := value
			var var_f36ac47c3a06_mapped *structpb.Value

			var var_f36ac47c3a06_err error
			var_f36ac47c3a06_mapped, var_f36ac47c3a06_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_f36ac47c3a06)
			if var_f36ac47c3a06_err != nil {
				panic(var_f36ac47c3a06_err)
			}

			var_b2fcc1dcfe4f_st.Fields[key] = var_f36ac47c3a06_mapped
		}
		var_b2fcc1dcfe4f_mapped = structpb.NewStructValue(var_b2fcc1dcfe4f_st)
		properties["annotations"] = var_b2fcc1dcfe4f_mapped
	}

	var_2aded745e145 := resource.Indexes

	if var_2aded745e145 != nil {
		var var_2aded745e145_mapped *structpb.Value

		var var_2aded745e145_err error
		var_2aded745e145_mapped, var_2aded745e145_err = types.ByResourcePropertyType(model.ResourceProperty_OBJECT).Pack(*var_2aded745e145)
		if var_2aded745e145_err != nil {
			panic(var_2aded745e145_err)
		}
		properties["indexes"] = var_2aded745e145_mapped
	}

	var_3cb2ee602651 := resource.SecurityConstraints

	if var_3cb2ee602651 != nil {
		var var_3cb2ee602651_mapped *structpb.Value

		var var_3cb2ee602651_l []*structpb.Value
		for _, value := range var_3cb2ee602651 {

			var_a48a1757f0c3 := value
			var var_a48a1757f0c3_mapped *structpb.Value

			var_a48a1757f0c3_mapped = structpb.NewStructValue(&structpb.Struct{Fields: SecurityConstraintMapperInstance.ToProperties(var_a48a1757f0c3)})

			var_3cb2ee602651_l = append(var_3cb2ee602651_l, var_a48a1757f0c3_mapped)
		}
		var_3cb2ee602651_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_3cb2ee602651_l})
		properties["securityConstraints"] = var_3cb2ee602651_mapped
	}

	var_69d3f1b876e4 := resource.Title

	if var_69d3f1b876e4 != nil {
		var var_69d3f1b876e4_mapped *structpb.Value

		var var_69d3f1b876e4_err error
		var_69d3f1b876e4_mapped, var_69d3f1b876e4_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_69d3f1b876e4)
		if var_69d3f1b876e4_err != nil {
			panic(var_69d3f1b876e4_err)
		}
		properties["title"] = var_69d3f1b876e4_mapped
	}

	var_d26a27b53eee := resource.Description

	if var_d26a27b53eee != nil {
		var var_d26a27b53eee_mapped *structpb.Value

		var var_d26a27b53eee_err error
		var_d26a27b53eee_mapped, var_d26a27b53eee_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_d26a27b53eee)
		if var_d26a27b53eee_err != nil {
			panic(var_d26a27b53eee_err)
		}
		properties["description"] = var_d26a27b53eee_mapped
	}
	return properties
}

func (m *ResourceMapper) FromProperties(properties map[string]*structpb.Value) *Resource {
	var s = m.New()
	if properties["id"] != nil {

		var_87e79013f81f := properties["id"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_UUID).UnPack(var_87e79013f81f)

		if err != nil {
			panic(err)
		}

		var_87e79013f81f_mapped := new(uuid.UUID)
		*var_87e79013f81f_mapped = val.(uuid.UUID)

		s.Id = var_87e79013f81f_mapped
	}
	if properties["version"] != nil {

		var_f0755525c82a := properties["version"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(var_f0755525c82a)

		if err != nil {
			panic(err)
		}

		var_f0755525c82a_mapped := val.(int32)

		s.Version = var_f0755525c82a_mapped
	}
	if properties["createdBy"] != nil {

		var_8f6749e9a482 := properties["createdBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_8f6749e9a482)

		if err != nil {
			panic(err)
		}

		var_8f6749e9a482_mapped := new(string)
		*var_8f6749e9a482_mapped = val.(string)

		s.CreatedBy = var_8f6749e9a482_mapped
	}
	if properties["updatedBy"] != nil {

		var_ac1a49c927c5 := properties["updatedBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_ac1a49c927c5)

		if err != nil {
			panic(err)
		}

		var_ac1a49c927c5_mapped := new(string)
		*var_ac1a49c927c5_mapped = val.(string)

		s.UpdatedBy = var_ac1a49c927c5_mapped
	}
	if properties["createdOn"] != nil {

		var_18ceece5a3da := properties["createdOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_18ceece5a3da)

		if err != nil {
			panic(err)
		}

		var_18ceece5a3da_mapped := new(time.Time)
		*var_18ceece5a3da_mapped = val.(time.Time)

		s.CreatedOn = var_18ceece5a3da_mapped
	}
	if properties["updatedOn"] != nil {

		var_ae2ef250ac2b := properties["updatedOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_ae2ef250ac2b)

		if err != nil {
			panic(err)
		}

		var_ae2ef250ac2b_mapped := new(time.Time)
		*var_ae2ef250ac2b_mapped = val.(time.Time)

		s.UpdatedOn = var_ae2ef250ac2b_mapped
	}
	if properties["name"] != nil {

		var_b7e1937db401 := properties["name"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_b7e1937db401)

		if err != nil {
			panic(err)
		}

		var_b7e1937db401_mapped := val.(string)

		s.Name = var_b7e1937db401_mapped
	}
	if properties["namespace"] != nil {

		var_69184e071ddc := properties["namespace"]
		var_69184e071ddc_mapped := NamespaceMapperInstance.FromProperties(var_69184e071ddc.GetStructValue().Fields)

		s.Namespace = var_69184e071ddc_mapped
	}
	if properties["virtual"] != nil {

		var_2e6718d4e55f := properties["virtual"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_BOOL).UnPack(var_2e6718d4e55f)

		if err != nil {
			panic(err)
		}

		var_2e6718d4e55f_mapped := val.(bool)

		s.Virtual = var_2e6718d4e55f_mapped
	}
	if properties["types"] != nil {

		var_833e92e74170 := properties["types"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_OBJECT).UnPack(var_833e92e74170)

		if err != nil {
			panic(err)
		}

		var_833e92e74170_mapped := new(unstructured.Unstructured)
		*var_833e92e74170_mapped = val.(unstructured.Unstructured)

		s.Types = var_833e92e74170_mapped
	}
	if properties["immutable"] != nil {

		var_7252a70d5a27 := properties["immutable"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_BOOL).UnPack(var_7252a70d5a27)

		if err != nil {
			panic(err)
		}

		var_7252a70d5a27_mapped := val.(bool)

		s.Immutable = var_7252a70d5a27_mapped
	}
	if properties["abstract"] != nil {

		var_badc4a189c2f := properties["abstract"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_BOOL).UnPack(var_badc4a189c2f)

		if err != nil {
			panic(err)
		}

		var_badc4a189c2f_mapped := val.(bool)

		s.Abstract = var_badc4a189c2f_mapped
	}
	if properties["dataSource"] != nil {

		var_bc5bbc97662d := properties["dataSource"]
		var_bc5bbc97662d_mapped := DataSourceMapperInstance.FromProperties(var_bc5bbc97662d.GetStructValue().Fields)

		s.DataSource = var_bc5bbc97662d_mapped
	}
	if properties["entity"] != nil {

		var_535e6bf71774 := properties["entity"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_535e6bf71774)

		if err != nil {
			panic(err)
		}

		var_535e6bf71774_mapped := new(string)
		*var_535e6bf71774_mapped = val.(string)

		s.Entity = var_535e6bf71774_mapped
	}
	if properties["catalog"] != nil {

		var_6519bcdf180f := properties["catalog"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_6519bcdf180f)

		if err != nil {
			panic(err)
		}

		var_6519bcdf180f_mapped := new(string)
		*var_6519bcdf180f_mapped = val.(string)

		s.Catalog = var_6519bcdf180f_mapped
	}
	if properties["annotations"] != nil {

		var_387311d19ce8 := properties["annotations"]
		var_387311d19ce8_mapped := make(map[string]string)
		for k, v := range var_387311d19ce8.GetStructValue().Fields {

			var_7dc447936d49 := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_7dc447936d49)

			if err != nil {
				panic(err)
			}

			var_7dc447936d49_mapped := val.(string)

			var_387311d19ce8_mapped[k] = var_7dc447936d49_mapped
		}

		s.Annotations = var_387311d19ce8_mapped
	}
	if properties["indexes"] != nil {

		var_9f85d85400b1 := properties["indexes"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_OBJECT).UnPack(var_9f85d85400b1)

		if err != nil {
			panic(err)
		}

		var_9f85d85400b1_mapped := new(unstructured.Unstructured)
		*var_9f85d85400b1_mapped = val.(unstructured.Unstructured)

		s.Indexes = var_9f85d85400b1_mapped
	}
	if properties["securityConstraints"] != nil {

		var_f9056c8bde20 := properties["securityConstraints"]
		var_f9056c8bde20_mapped := []*SecurityConstraint{}
		for _, v := range var_f9056c8bde20.GetListValue().Values {

			var_64d686f83394 := v
			var_64d686f83394_mapped := SecurityConstraintMapperInstance.FromProperties(var_64d686f83394.GetStructValue().Fields)

			var_f9056c8bde20_mapped = append(var_f9056c8bde20_mapped, var_64d686f83394_mapped)
		}

		s.SecurityConstraints = var_f9056c8bde20_mapped
	}
	if properties["title"] != nil {

		var_5b4a41367922 := properties["title"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_5b4a41367922)

		if err != nil {
			panic(err)
		}

		var_5b4a41367922_mapped := new(string)
		*var_5b4a41367922_mapped = val.(string)

		s.Title = var_5b4a41367922_mapped
	}
	if properties["description"] != nil {

		var_e4b681d0a4f1 := properties["description"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_e4b681d0a4f1)

		if err != nil {
			panic(err)
		}

		var_e4b681d0a4f1_mapped := new(string)
		*var_e4b681d0a4f1_mapped = val.(string)

		s.Description = var_e4b681d0a4f1_mapped
	}
	return s
}
