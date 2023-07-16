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

	var_d7d062a54d25 := dataSource.Id

	if var_d7d062a54d25 != nil {
		var var_d7d062a54d25_mapped *structpb.Value

		var var_d7d062a54d25_err error
		var_d7d062a54d25_mapped, var_d7d062a54d25_err = types.ByResourcePropertyType(model.ResourceProperty_UUID).Pack(*var_d7d062a54d25)
		if var_d7d062a54d25_err != nil {
			panic(var_d7d062a54d25_err)
		}
		properties["id"] = var_d7d062a54d25_mapped
	}

	var_ffc6992e7e07 := dataSource.Version

	var var_ffc6992e7e07_mapped *structpb.Value

	var var_ffc6992e7e07_err error
	var_ffc6992e7e07_mapped, var_ffc6992e7e07_err = types.ByResourcePropertyType(model.ResourceProperty_INT32).Pack(var_ffc6992e7e07)
	if var_ffc6992e7e07_err != nil {
		panic(var_ffc6992e7e07_err)
	}
	properties["version"] = var_ffc6992e7e07_mapped

	var_658b6636d618 := dataSource.CreatedBy

	if var_658b6636d618 != nil {
		var var_658b6636d618_mapped *structpb.Value

		var var_658b6636d618_err error
		var_658b6636d618_mapped, var_658b6636d618_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_658b6636d618)
		if var_658b6636d618_err != nil {
			panic(var_658b6636d618_err)
		}
		properties["createdBy"] = var_658b6636d618_mapped
	}

	var_8014071eac16 := dataSource.UpdatedBy

	if var_8014071eac16 != nil {
		var var_8014071eac16_mapped *structpb.Value

		var var_8014071eac16_err error
		var_8014071eac16_mapped, var_8014071eac16_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_8014071eac16)
		if var_8014071eac16_err != nil {
			panic(var_8014071eac16_err)
		}
		properties["updatedBy"] = var_8014071eac16_mapped
	}

	var_071c01a25a0d := dataSource.CreatedOn

	if var_071c01a25a0d != nil {
		var var_071c01a25a0d_mapped *structpb.Value

		var var_071c01a25a0d_err error
		var_071c01a25a0d_mapped, var_071c01a25a0d_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_071c01a25a0d)
		if var_071c01a25a0d_err != nil {
			panic(var_071c01a25a0d_err)
		}
		properties["createdOn"] = var_071c01a25a0d_mapped
	}

	var_83ef2285eb19 := dataSource.UpdatedOn

	if var_83ef2285eb19 != nil {
		var var_83ef2285eb19_mapped *structpb.Value

		var var_83ef2285eb19_err error
		var_83ef2285eb19_mapped, var_83ef2285eb19_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_83ef2285eb19)
		if var_83ef2285eb19_err != nil {
			panic(var_83ef2285eb19_err)
		}
		properties["updatedOn"] = var_83ef2285eb19_mapped
	}

	var_588ae66c2454 := dataSource.Name

	var var_588ae66c2454_mapped *structpb.Value

	var var_588ae66c2454_err error
	var_588ae66c2454_mapped, var_588ae66c2454_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_588ae66c2454)
	if var_588ae66c2454_err != nil {
		panic(var_588ae66c2454_err)
	}
	properties["name"] = var_588ae66c2454_mapped

	var_3f0fde6d7265 := dataSource.Description

	var var_3f0fde6d7265_mapped *structpb.Value

	var var_3f0fde6d7265_err error
	var_3f0fde6d7265_mapped, var_3f0fde6d7265_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_3f0fde6d7265)
	if var_3f0fde6d7265_err != nil {
		panic(var_3f0fde6d7265_err)
	}
	properties["description"] = var_3f0fde6d7265_mapped

	var_eda0da40711f := dataSource.Backend

	var var_eda0da40711f_mapped *structpb.Value

	var var_eda0da40711f_err error
	var_eda0da40711f_mapped, var_eda0da40711f_err = types.ByResourcePropertyType(model.ResourceProperty_ENUM).Pack(string(var_eda0da40711f))
	if var_eda0da40711f_err != nil {
		panic(var_eda0da40711f_err)
	}
	properties["backend"] = var_eda0da40711f_mapped

	var_66cb16f87db4 := dataSource.Options

	var var_66cb16f87db4_mapped *structpb.Value

	var var_66cb16f87db4_st *structpb.Struct = new(structpb.Struct)
	var_66cb16f87db4_st.Fields = make(map[string]*structpb.Value)
	for key, value := range var_66cb16f87db4 {

		var_3a46a4059795 := value
		var var_3a46a4059795_mapped *structpb.Value

		var var_3a46a4059795_err error
		var_3a46a4059795_mapped, var_3a46a4059795_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_3a46a4059795)
		if var_3a46a4059795_err != nil {
			panic(var_3a46a4059795_err)
		}

		var_66cb16f87db4_st.Fields[key] = var_3a46a4059795_mapped
	}
	var_66cb16f87db4_mapped = structpb.NewStructValue(var_66cb16f87db4_st)
	properties["options"] = var_66cb16f87db4_mapped
	return properties
}

func (m *DataSourceMapper) FromProperties(properties map[string]*structpb.Value) *DataSource {
	var s = m.New()
	if properties["id"] != nil {

		var_88d128c59327 := properties["id"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_UUID).UnPack(var_88d128c59327)

		if err != nil {
			panic(err)
		}

		var_88d128c59327_mapped := new(uuid.UUID)
		*var_88d128c59327_mapped = val.(uuid.UUID)

		s.Id = var_88d128c59327_mapped
	}
	if properties["version"] != nil {

		var_3b43bc37abd9 := properties["version"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(var_3b43bc37abd9)

		if err != nil {
			panic(err)
		}

		var_3b43bc37abd9_mapped := val.(int32)

		s.Version = var_3b43bc37abd9_mapped
	}
	if properties["createdBy"] != nil {

		var_700849ce511e := properties["createdBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_700849ce511e)

		if err != nil {
			panic(err)
		}

		var_700849ce511e_mapped := new(string)
		*var_700849ce511e_mapped = val.(string)

		s.CreatedBy = var_700849ce511e_mapped
	}
	if properties["updatedBy"] != nil {

		var_ce9b69c9b6e9 := properties["updatedBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_ce9b69c9b6e9)

		if err != nil {
			panic(err)
		}

		var_ce9b69c9b6e9_mapped := new(string)
		*var_ce9b69c9b6e9_mapped = val.(string)

		s.UpdatedBy = var_ce9b69c9b6e9_mapped
	}
	if properties["createdOn"] != nil {

		var_7600e9ceb5c4 := properties["createdOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_7600e9ceb5c4)

		if err != nil {
			panic(err)
		}

		var_7600e9ceb5c4_mapped := new(time.Time)
		*var_7600e9ceb5c4_mapped = val.(time.Time)

		s.CreatedOn = var_7600e9ceb5c4_mapped
	}
	if properties["updatedOn"] != nil {

		var_c4e40a064596 := properties["updatedOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_c4e40a064596)

		if err != nil {
			panic(err)
		}

		var_c4e40a064596_mapped := new(time.Time)
		*var_c4e40a064596_mapped = val.(time.Time)

		s.UpdatedOn = var_c4e40a064596_mapped
	}
	if properties["name"] != nil {

		var_2aa04e2a2ba2 := properties["name"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_2aa04e2a2ba2)

		if err != nil {
			panic(err)
		}

		var_2aa04e2a2ba2_mapped := val.(string)

		s.Name = var_2aa04e2a2ba2_mapped
	}
	if properties["description"] != nil {

		var_8eaab3e7d06a := properties["description"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_8eaab3e7d06a)

		if err != nil {
			panic(err)
		}

		var_8eaab3e7d06a_mapped := val.(string)

		s.Description = var_8eaab3e7d06a_mapped
	}
	if properties["backend"] != nil {

		var_c39551d97ea2 := properties["backend"]
		var_c39551d97ea2_mapped := (DataSourceBackend)(var_c39551d97ea2.GetStringValue())

		s.Backend = var_c39551d97ea2_mapped
	}
	if properties["options"] != nil {

		var_29ba585e9a9d := properties["options"]
		var_29ba585e9a9d_mapped := make(map[string]string)
		for k, v := range var_29ba585e9a9d.GetStructValue().Fields {

			var_5742dd9b63f7 := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_5742dd9b63f7)

			if err != nil {
				panic(err)
			}

			var_5742dd9b63f7_mapped := val.(string)

			var_29ba585e9a9d_mapped[k] = var_5742dd9b63f7_mapped
		}

		s.Options = var_29ba585e9a9d_mapped
	}
	return s
}
