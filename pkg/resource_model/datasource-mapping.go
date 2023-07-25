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

	var_8e2b4448ea5b := dataSource.Id

	if var_8e2b4448ea5b != nil {
		var var_8e2b4448ea5b_mapped *structpb.Value

		var var_8e2b4448ea5b_err error
		var_8e2b4448ea5b_mapped, var_8e2b4448ea5b_err = types.ByResourcePropertyType(model.ResourceProperty_UUID).Pack(*var_8e2b4448ea5b)
		if var_8e2b4448ea5b_err != nil {
			panic(var_8e2b4448ea5b_err)
		}
		properties["id"] = var_8e2b4448ea5b_mapped
	}

	var_f5ea23d59681 := dataSource.Version

	var var_f5ea23d59681_mapped *structpb.Value

	var var_f5ea23d59681_err error
	var_f5ea23d59681_mapped, var_f5ea23d59681_err = types.ByResourcePropertyType(model.ResourceProperty_INT32).Pack(var_f5ea23d59681)
	if var_f5ea23d59681_err != nil {
		panic(var_f5ea23d59681_err)
	}
	properties["version"] = var_f5ea23d59681_mapped

	var_a7b1458d03c1 := dataSource.CreatedBy

	if var_a7b1458d03c1 != nil {
		var var_a7b1458d03c1_mapped *structpb.Value

		var var_a7b1458d03c1_err error
		var_a7b1458d03c1_mapped, var_a7b1458d03c1_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_a7b1458d03c1)
		if var_a7b1458d03c1_err != nil {
			panic(var_a7b1458d03c1_err)
		}
		properties["createdBy"] = var_a7b1458d03c1_mapped
	}

	var_a5369b5165d3 := dataSource.UpdatedBy

	if var_a5369b5165d3 != nil {
		var var_a5369b5165d3_mapped *structpb.Value

		var var_a5369b5165d3_err error
		var_a5369b5165d3_mapped, var_a5369b5165d3_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_a5369b5165d3)
		if var_a5369b5165d3_err != nil {
			panic(var_a5369b5165d3_err)
		}
		properties["updatedBy"] = var_a5369b5165d3_mapped
	}

	var_6e60f895d9fd := dataSource.CreatedOn

	if var_6e60f895d9fd != nil {
		var var_6e60f895d9fd_mapped *structpb.Value

		var var_6e60f895d9fd_err error
		var_6e60f895d9fd_mapped, var_6e60f895d9fd_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_6e60f895d9fd)
		if var_6e60f895d9fd_err != nil {
			panic(var_6e60f895d9fd_err)
		}
		properties["createdOn"] = var_6e60f895d9fd_mapped
	}

	var_48de5cd1d3f2 := dataSource.UpdatedOn

	if var_48de5cd1d3f2 != nil {
		var var_48de5cd1d3f2_mapped *structpb.Value

		var var_48de5cd1d3f2_err error
		var_48de5cd1d3f2_mapped, var_48de5cd1d3f2_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_48de5cd1d3f2)
		if var_48de5cd1d3f2_err != nil {
			panic(var_48de5cd1d3f2_err)
		}
		properties["updatedOn"] = var_48de5cd1d3f2_mapped
	}

	var_4e81a6e39220 := dataSource.Name

	var var_4e81a6e39220_mapped *structpb.Value

	var var_4e81a6e39220_err error
	var_4e81a6e39220_mapped, var_4e81a6e39220_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_4e81a6e39220)
	if var_4e81a6e39220_err != nil {
		panic(var_4e81a6e39220_err)
	}
	properties["name"] = var_4e81a6e39220_mapped

	var_7c981df5904a := dataSource.Description

	var var_7c981df5904a_mapped *structpb.Value

	var var_7c981df5904a_err error
	var_7c981df5904a_mapped, var_7c981df5904a_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_7c981df5904a)
	if var_7c981df5904a_err != nil {
		panic(var_7c981df5904a_err)
	}
	properties["description"] = var_7c981df5904a_mapped

	var_6a27704a1e7a := dataSource.Backend

	var var_6a27704a1e7a_mapped *structpb.Value

	var var_6a27704a1e7a_err error
	var_6a27704a1e7a_mapped, var_6a27704a1e7a_err = types.ByResourcePropertyType(model.ResourceProperty_ENUM).Pack(string(var_6a27704a1e7a))
	if var_6a27704a1e7a_err != nil {
		panic(var_6a27704a1e7a_err)
	}
	properties["backend"] = var_6a27704a1e7a_mapped

	var_88542b56cfc2 := dataSource.Options

	var var_88542b56cfc2_mapped *structpb.Value

	var var_88542b56cfc2_st *structpb.Struct = new(structpb.Struct)
	var_88542b56cfc2_st.Fields = make(map[string]*structpb.Value)
	for key, value := range var_88542b56cfc2 {

		var_92dc59f7d755 := value
		var var_92dc59f7d755_mapped *structpb.Value

		var var_92dc59f7d755_err error
		var_92dc59f7d755_mapped, var_92dc59f7d755_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_92dc59f7d755)
		if var_92dc59f7d755_err != nil {
			panic(var_92dc59f7d755_err)
		}

		var_88542b56cfc2_st.Fields[key] = var_92dc59f7d755_mapped
	}
	var_88542b56cfc2_mapped = structpb.NewStructValue(var_88542b56cfc2_st)
	properties["options"] = var_88542b56cfc2_mapped
	return properties
}

func (m *DataSourceMapper) FromProperties(properties map[string]*structpb.Value) *DataSource {
	var s = m.New()
	if properties["id"] != nil {

		var_4dd6117ad77e := properties["id"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_UUID).UnPack(var_4dd6117ad77e)

		if err != nil {
			panic(err)
		}

		var_4dd6117ad77e_mapped := new(uuid.UUID)
		*var_4dd6117ad77e_mapped = val.(uuid.UUID)

		s.Id = var_4dd6117ad77e_mapped
	}
	if properties["version"] != nil {

		var_e20843638544 := properties["version"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(var_e20843638544)

		if err != nil {
			panic(err)
		}

		var_e20843638544_mapped := val.(int32)

		s.Version = var_e20843638544_mapped
	}
	if properties["createdBy"] != nil {

		var_1c36440ca659 := properties["createdBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_1c36440ca659)

		if err != nil {
			panic(err)
		}

		var_1c36440ca659_mapped := new(string)
		*var_1c36440ca659_mapped = val.(string)

		s.CreatedBy = var_1c36440ca659_mapped
	}
	if properties["updatedBy"] != nil {

		var_b5a47a56e434 := properties["updatedBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_b5a47a56e434)

		if err != nil {
			panic(err)
		}

		var_b5a47a56e434_mapped := new(string)
		*var_b5a47a56e434_mapped = val.(string)

		s.UpdatedBy = var_b5a47a56e434_mapped
	}
	if properties["createdOn"] != nil {

		var_2dd464637ee9 := properties["createdOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_2dd464637ee9)

		if err != nil {
			panic(err)
		}

		var_2dd464637ee9_mapped := new(time.Time)
		*var_2dd464637ee9_mapped = val.(time.Time)

		s.CreatedOn = var_2dd464637ee9_mapped
	}
	if properties["updatedOn"] != nil {

		var_9b68cc9957bd := properties["updatedOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_9b68cc9957bd)

		if err != nil {
			panic(err)
		}

		var_9b68cc9957bd_mapped := new(time.Time)
		*var_9b68cc9957bd_mapped = val.(time.Time)

		s.UpdatedOn = var_9b68cc9957bd_mapped
	}
	if properties["name"] != nil {

		var_7d55a56670bd := properties["name"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_7d55a56670bd)

		if err != nil {
			panic(err)
		}

		var_7d55a56670bd_mapped := val.(string)

		s.Name = var_7d55a56670bd_mapped
	}
	if properties["description"] != nil {

		var_0d5cf1cbccdb := properties["description"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_0d5cf1cbccdb)

		if err != nil {
			panic(err)
		}

		var_0d5cf1cbccdb_mapped := val.(string)

		s.Description = var_0d5cf1cbccdb_mapped
	}
	if properties["backend"] != nil {

		var_8beba85fa85e := properties["backend"]
		var_8beba85fa85e_mapped := (DataSourceBackend)(var_8beba85fa85e.GetStringValue())

		s.Backend = var_8beba85fa85e_mapped
	}
	if properties["options"] != nil {

		var_477015769139 := properties["options"]
		var_477015769139_mapped := make(map[string]string)
		for k, v := range var_477015769139.GetStructValue().Fields {

			var_499c6baa4ad7 := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_499c6baa4ad7)

			if err != nil {
				panic(err)
			}

			var_499c6baa4ad7_mapped := val.(string)

			var_477015769139_mapped[k] = var_499c6baa4ad7_mapped
		}

		s.Options = var_477015769139_mapped
	}
	return s
}
