package resource_model

import (
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/types"
	"google.golang.org/protobuf/types/known/structpb"
)

import "github.com/google/uuid"
import "time"

type DataSourceMapper struct {
}

func NewDataSourceMapper() *DataSourceMapper {
	return &DataSourceMapper{}
}

var DataSourceMapperInstance = NewDataSourceMapper()

func (m *DataSourceMapper) New() *DataSource {
	return &DataSource{}
}

func (m *DataSourceMapper) ToRecord(dataSource *DataSource) *model.Record {
	var rec = &model.Record{}
	rec.Properties = m.ToProperties(dataSource)

	if dataSource.Id != nil {
		rec.Id = dataSource.Id.String()
	}

	return rec
}

func (m *DataSourceMapper) FromRecord(record *model.Record) *DataSource {
	return m.FromProperties(record.Properties)
}

func (m *DataSourceMapper) ToProperties(dataSource *DataSource) map[string]*structpb.Value {
	var properties = make(map[string]*structpb.Value)

	var_fdce3e3527e8 := dataSource.Id

	if var_fdce3e3527e8 != nil {
		var var_fdce3e3527e8_mapped *structpb.Value

		var var_fdce3e3527e8_err error
		var_fdce3e3527e8_mapped, var_fdce3e3527e8_err = types.ByResourcePropertyType(model.ResourceProperty_UUID).Pack(*var_fdce3e3527e8)
		if var_fdce3e3527e8_err != nil {
			panic(var_fdce3e3527e8_err)
		}
		properties["id"] = var_fdce3e3527e8_mapped
	}

	var_4dc82b2e1703 := dataSource.Version

	var var_4dc82b2e1703_mapped *structpb.Value

	var var_4dc82b2e1703_err error
	var_4dc82b2e1703_mapped, var_4dc82b2e1703_err = types.ByResourcePropertyType(model.ResourceProperty_INT32).Pack(var_4dc82b2e1703)
	if var_4dc82b2e1703_err != nil {
		panic(var_4dc82b2e1703_err)
	}
	properties["version"] = var_4dc82b2e1703_mapped

	var_ec5e06d5e8db := dataSource.CreatedBy

	if var_ec5e06d5e8db != nil {
		var var_ec5e06d5e8db_mapped *structpb.Value

		var var_ec5e06d5e8db_err error
		var_ec5e06d5e8db_mapped, var_ec5e06d5e8db_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_ec5e06d5e8db)
		if var_ec5e06d5e8db_err != nil {
			panic(var_ec5e06d5e8db_err)
		}
		properties["createdBy"] = var_ec5e06d5e8db_mapped
	}

	var_76a05eecb608 := dataSource.UpdatedBy

	if var_76a05eecb608 != nil {
		var var_76a05eecb608_mapped *structpb.Value

		var var_76a05eecb608_err error
		var_76a05eecb608_mapped, var_76a05eecb608_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_76a05eecb608)
		if var_76a05eecb608_err != nil {
			panic(var_76a05eecb608_err)
		}
		properties["updatedBy"] = var_76a05eecb608_mapped
	}

	var_693673c76eaa := dataSource.CreatedOn

	if var_693673c76eaa != nil {
		var var_693673c76eaa_mapped *structpb.Value

		var var_693673c76eaa_err error
		var_693673c76eaa_mapped, var_693673c76eaa_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_693673c76eaa)
		if var_693673c76eaa_err != nil {
			panic(var_693673c76eaa_err)
		}
		properties["createdOn"] = var_693673c76eaa_mapped
	}

	var_6847bc0ba87d := dataSource.UpdatedOn

	if var_6847bc0ba87d != nil {
		var var_6847bc0ba87d_mapped *structpb.Value

		var var_6847bc0ba87d_err error
		var_6847bc0ba87d_mapped, var_6847bc0ba87d_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_6847bc0ba87d)
		if var_6847bc0ba87d_err != nil {
			panic(var_6847bc0ba87d_err)
		}
		properties["updatedOn"] = var_6847bc0ba87d_mapped
	}

	var_cba486c89b94 := dataSource.Name

	var var_cba486c89b94_mapped *structpb.Value

	var var_cba486c89b94_err error
	var_cba486c89b94_mapped, var_cba486c89b94_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_cba486c89b94)
	if var_cba486c89b94_err != nil {
		panic(var_cba486c89b94_err)
	}
	properties["name"] = var_cba486c89b94_mapped

	var_55fed68a1290 := dataSource.Description

	var var_55fed68a1290_mapped *structpb.Value

	var var_55fed68a1290_err error
	var_55fed68a1290_mapped, var_55fed68a1290_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_55fed68a1290)
	if var_55fed68a1290_err != nil {
		panic(var_55fed68a1290_err)
	}
	properties["description"] = var_55fed68a1290_mapped

	var_0f8e043d947f := dataSource.Backend

	var var_0f8e043d947f_mapped *structpb.Value

	var var_0f8e043d947f_err error
	var_0f8e043d947f_mapped, var_0f8e043d947f_err = types.ByResourcePropertyType(model.ResourceProperty_ENUM).Pack(string(var_0f8e043d947f))
	if var_0f8e043d947f_err != nil {
		panic(var_0f8e043d947f_err)
	}
	properties["backend"] = var_0f8e043d947f_mapped

	var_38172bb2ded7 := dataSource.Options

	var var_38172bb2ded7_mapped *structpb.Value

	var var_38172bb2ded7_st *structpb.Struct = new(structpb.Struct)
	var_38172bb2ded7_st.Fields = make(map[string]*structpb.Value)
	for key, value := range var_38172bb2ded7 {

		var_49ea637f00ed := value
		var var_49ea637f00ed_mapped *structpb.Value

		var var_49ea637f00ed_err error
		var_49ea637f00ed_mapped, var_49ea637f00ed_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_49ea637f00ed)
		if var_49ea637f00ed_err != nil {
			panic(var_49ea637f00ed_err)
		}

		var_38172bb2ded7_st.Fields[key] = var_49ea637f00ed_mapped
	}
	var_38172bb2ded7_mapped = structpb.NewStructValue(var_38172bb2ded7_st)
	properties["options"] = var_38172bb2ded7_mapped
	return properties
}

func (m *DataSourceMapper) FromProperties(properties map[string]*structpb.Value) *DataSource {
	var s = m.New()
	if properties["id"] != nil {

		var_05e750aeaaa7 := properties["id"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_UUID).UnPack(var_05e750aeaaa7)

		if err != nil {
			panic(err)
		}

		var_05e750aeaaa7_mapped := new(uuid.UUID)
		*var_05e750aeaaa7_mapped = val.(uuid.UUID)

		s.Id = var_05e750aeaaa7_mapped
	}
	if properties["version"] != nil {

		var_5818b901d96a := properties["version"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(var_5818b901d96a)

		if err != nil {
			panic(err)
		}

		var_5818b901d96a_mapped := val.(int32)

		s.Version = var_5818b901d96a_mapped
	}
	if properties["createdBy"] != nil {

		var_872752174628 := properties["createdBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_872752174628)

		if err != nil {
			panic(err)
		}

		var_872752174628_mapped := new(string)
		*var_872752174628_mapped = val.(string)

		s.CreatedBy = var_872752174628_mapped
	}
	if properties["updatedBy"] != nil {

		var_d821a8cb5ce5 := properties["updatedBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_d821a8cb5ce5)

		if err != nil {
			panic(err)
		}

		var_d821a8cb5ce5_mapped := new(string)
		*var_d821a8cb5ce5_mapped = val.(string)

		s.UpdatedBy = var_d821a8cb5ce5_mapped
	}
	if properties["createdOn"] != nil {

		var_8fd3124644c5 := properties["createdOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_8fd3124644c5)

		if err != nil {
			panic(err)
		}

		var_8fd3124644c5_mapped := new(time.Time)
		*var_8fd3124644c5_mapped = val.(time.Time)

		s.CreatedOn = var_8fd3124644c5_mapped
	}
	if properties["updatedOn"] != nil {

		var_8b59bc63aedd := properties["updatedOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_8b59bc63aedd)

		if err != nil {
			panic(err)
		}

		var_8b59bc63aedd_mapped := new(time.Time)
		*var_8b59bc63aedd_mapped = val.(time.Time)

		s.UpdatedOn = var_8b59bc63aedd_mapped
	}
	if properties["name"] != nil {

		var_e7981335abf3 := properties["name"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_e7981335abf3)

		if err != nil {
			panic(err)
		}

		var_e7981335abf3_mapped := val.(string)

		s.Name = var_e7981335abf3_mapped
	}
	if properties["description"] != nil {

		var_36eb31ed733a := properties["description"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_36eb31ed733a)

		if err != nil {
			panic(err)
		}

		var_36eb31ed733a_mapped := val.(string)

		s.Description = var_36eb31ed733a_mapped
	}
	if properties["backend"] != nil {

		var_438827e6440f := properties["backend"]
		var_438827e6440f_mapped := (DataSourceBackend)(var_438827e6440f.GetStringValue())

		s.Backend = var_438827e6440f_mapped
	}
	if properties["options"] != nil {

		var_43bb219afe09 := properties["options"]
		var_43bb219afe09_mapped := make(map[string]string)
		for k, v := range var_43bb219afe09.GetStructValue().Fields {

			var_6cf6976aa764 := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_6cf6976aa764)

			if err != nil {
				panic(err)
			}

			var_6cf6976aa764_mapped := val.(string)

			var_43bb219afe09_mapped[k] = var_6cf6976aa764_mapped
		}

		s.Options = var_43bb219afe09_mapped
	}
	return s
}
