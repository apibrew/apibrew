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
	return rec
}

func (m *DataSourceMapper) FromRecord(record *model.Record) *DataSource {
	return m.FromProperties(record.Properties)
}

func (m *DataSourceMapper) ToProperties(dataSource *DataSource) map[string]*structpb.Value {
	var properties = make(map[string]*structpb.Value)

	if dataSource.Id != nil {
		id, err := types.ByResourcePropertyType(model.ResourceProperty_UUID).Pack(*dataSource.Id)
		if err != nil {
			panic(err)
		}
		properties["id"] = id
	}

	if dataSource.Version != nil {
		version, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).Pack(*dataSource.Version)
		if err != nil {
			panic(err)
		}
		properties["version"] = version
	}

	if dataSource.CreatedBy != nil {
		createdBy, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*dataSource.CreatedBy)
		if err != nil {
			panic(err)
		}
		properties["createdBy"] = createdBy
	}

	if dataSource.UpdatedBy != nil {
		updatedBy, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*dataSource.UpdatedBy)
		if err != nil {
			panic(err)
		}
		properties["updatedBy"] = updatedBy
	}

	if dataSource.CreatedOn != nil {
		createdOn, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*dataSource.CreatedOn)
		if err != nil {
			panic(err)
		}
		properties["createdOn"] = createdOn
	}

	if dataSource.UpdatedOn != nil {
		updatedOn, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*dataSource.UpdatedOn)
		if err != nil {
			panic(err)
		}
		properties["updatedOn"] = updatedOn
	}

	name, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(dataSource.Name)
	if err != nil {
		panic(err)
	}
	properties["name"] = name

	if dataSource.Description != nil {
		description, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*dataSource.Description)
		if err != nil {
			panic(err)
		}
		properties["description"] = description
	}

	backend, err := types.ByResourcePropertyType(model.ResourceProperty_ENUM).Pack(dataSource.Backend)
	if err != nil {
		panic(err)
	}
	properties["backend"] = backend

	return properties
}

func (m *DataSourceMapper) FromProperties(properties map[string]*structpb.Value) *DataSource {
	var s = m.New()
	if properties["id"] != nil {

		var_1dad339af38f := properties["id"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_UUID).UnPack(var_1dad339af38f)

		if err != nil {
			panic(err)
		}

		var_1dad339af38f_mapped := new(uuid.UUID)
		*var_1dad339af38f_mapped = val.(uuid.UUID)

		s.Id = var_1dad339af38f_mapped
	}
	if properties["version"] != nil {

		var_83d83569ab6a := properties["version"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(var_83d83569ab6a)

		if err != nil {
			panic(err)
		}

		var_83d83569ab6a_mapped := new(int32)
		*var_83d83569ab6a_mapped = val.(int32)

		s.Version = var_83d83569ab6a_mapped
	}
	if properties["createdBy"] != nil {

		var_444ce9b5d2f9 := properties["createdBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_444ce9b5d2f9)

		if err != nil {
			panic(err)
		}

		var_444ce9b5d2f9_mapped := new(string)
		*var_444ce9b5d2f9_mapped = val.(string)

		s.CreatedBy = var_444ce9b5d2f9_mapped
	}
	if properties["updatedBy"] != nil {

		var_033098f1232d := properties["updatedBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_033098f1232d)

		if err != nil {
			panic(err)
		}

		var_033098f1232d_mapped := new(string)
		*var_033098f1232d_mapped = val.(string)

		s.UpdatedBy = var_033098f1232d_mapped
	}
	if properties["createdOn"] != nil {

		var_290c5ac3ac30 := properties["createdOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_290c5ac3ac30)

		if err != nil {
			panic(err)
		}

		var_290c5ac3ac30_mapped := new(time.Time)
		*var_290c5ac3ac30_mapped = val.(time.Time)

		s.CreatedOn = var_290c5ac3ac30_mapped
	}
	if properties["updatedOn"] != nil {

		var_3f2700eba157 := properties["updatedOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_3f2700eba157)

		if err != nil {
			panic(err)
		}

		var_3f2700eba157_mapped := new(time.Time)
		*var_3f2700eba157_mapped = val.(time.Time)

		s.UpdatedOn = var_3f2700eba157_mapped
	}
	if properties["name"] != nil {

		var_e64ff55d22f3 := properties["name"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_e64ff55d22f3)

		if err != nil {
			panic(err)
		}

		var_e64ff55d22f3_mapped := val.(string)

		s.Name = var_e64ff55d22f3_mapped
	}
	if properties["description"] != nil {

		var_d92062ecbff0 := properties["description"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_d92062ecbff0)

		if err != nil {
			panic(err)
		}

		var_d92062ecbff0_mapped := new(string)
		*var_d92062ecbff0_mapped = val.(string)

		s.Description = var_d92062ecbff0_mapped
	}
	if properties["backend"] != nil {

		var_e82ce81284f9 := properties["backend"]
		var_e82ce81284f9_mapped := (DataSourceBackend)(var_e82ce81284f9.GetStringValue())

		s.Backend = var_e82ce81284f9_mapped
	}
	if properties["options"] != nil {

		var_1496e0bbd218 := properties["options"]
		var_1496e0bbd218_mapped := make(map[string]string)
		for k, v := range var_1496e0bbd218.GetStructValue().Fields {

			var_f99a51981473 := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_f99a51981473)

			if err != nil {
				panic(err)
			}

			var_f99a51981473_mapped := val.(string)

			var_1496e0bbd218_mapped[k] = var_f99a51981473_mapped
		}

		s.Options = var_1496e0bbd218_mapped
	}
	return s
}
