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

		var_384dd1ea993f := properties["id"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_UUID).UnPack(var_384dd1ea993f)

		if err != nil {
			panic(err)
		}

		var_384dd1ea993f_mapped := new(uuid.UUID)
		*var_384dd1ea993f_mapped = val.(uuid.UUID)

		s.Id = var_384dd1ea993f_mapped
	}
	if properties["version"] != nil {

		var_06de0b0f5e17 := properties["version"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(var_06de0b0f5e17)

		if err != nil {
			panic(err)
		}

		var_06de0b0f5e17_mapped := new(int32)
		*var_06de0b0f5e17_mapped = val.(int32)

		s.Version = var_06de0b0f5e17_mapped
	}
	if properties["createdBy"] != nil {

		var_1b0a8f73574c := properties["createdBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_1b0a8f73574c)

		if err != nil {
			panic(err)
		}

		var_1b0a8f73574c_mapped := new(string)
		*var_1b0a8f73574c_mapped = val.(string)

		s.CreatedBy = var_1b0a8f73574c_mapped
	}
	if properties["updatedBy"] != nil {

		var_7ee9d069d2b7 := properties["updatedBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_7ee9d069d2b7)

		if err != nil {
			panic(err)
		}

		var_7ee9d069d2b7_mapped := new(string)
		*var_7ee9d069d2b7_mapped = val.(string)

		s.UpdatedBy = var_7ee9d069d2b7_mapped
	}
	if properties["createdOn"] != nil {

		var_ca174ad345c6 := properties["createdOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_ca174ad345c6)

		if err != nil {
			panic(err)
		}

		var_ca174ad345c6_mapped := new(time.Time)
		*var_ca174ad345c6_mapped = val.(time.Time)

		s.CreatedOn = var_ca174ad345c6_mapped
	}
	if properties["updatedOn"] != nil {

		var_5740ff6bbc52 := properties["updatedOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_5740ff6bbc52)

		if err != nil {
			panic(err)
		}

		var_5740ff6bbc52_mapped := new(time.Time)
		*var_5740ff6bbc52_mapped = val.(time.Time)

		s.UpdatedOn = var_5740ff6bbc52_mapped
	}
	if properties["name"] != nil {

		var_186ef8776f92 := properties["name"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_186ef8776f92)

		if err != nil {
			panic(err)
		}

		var_186ef8776f92_mapped := val.(string)

		s.Name = var_186ef8776f92_mapped
	}
	if properties["description"] != nil {

		var_c66527a9100d := properties["description"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_c66527a9100d)

		if err != nil {
			panic(err)
		}

		var_c66527a9100d_mapped := new(string)
		*var_c66527a9100d_mapped = val.(string)

		s.Description = var_c66527a9100d_mapped
	}
	if properties["backend"] != nil {

		var_db0505ead5be := properties["backend"]
		var_db0505ead5be_mapped := (DataSourceBackend)(var_db0505ead5be.GetStringValue())

		s.Backend = var_db0505ead5be_mapped
	}
	if properties["options"] != nil {

		var_045305bea9f5 := properties["options"]
		var_045305bea9f5_mapped := make(map[string]string)
		for k, v := range var_045305bea9f5.GetStructValue().Fields {

			var_8c087c8b25a1 := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_8c087c8b25a1)

			if err != nil {
				panic(err)
			}

			var_8c087c8b25a1_mapped := val.(string)

			var_045305bea9f5_mapped[k] = var_8c087c8b25a1_mapped
		}

		s.Options = var_045305bea9f5_mapped
	}
	return s
}
