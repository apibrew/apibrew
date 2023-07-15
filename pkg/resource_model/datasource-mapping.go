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

		var_b52b35ec4cf3 := properties["id"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_UUID).UnPack(var_b52b35ec4cf3)

		if err != nil {
			panic(err)
		}

		var_b52b35ec4cf3_mapped := new(uuid.UUID)
		*var_b52b35ec4cf3_mapped = val.(uuid.UUID)

		s.Id = var_b52b35ec4cf3_mapped
	}
	if properties["version"] != nil {

		var_1e3f4a6097b0 := properties["version"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(var_1e3f4a6097b0)

		if err != nil {
			panic(err)
		}

		var_1e3f4a6097b0_mapped := new(int32)
		*var_1e3f4a6097b0_mapped = val.(int32)

		s.Version = var_1e3f4a6097b0_mapped
	}
	if properties["createdBy"] != nil {

		var_74e7573102c1 := properties["createdBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_74e7573102c1)

		if err != nil {
			panic(err)
		}

		var_74e7573102c1_mapped := new(string)
		*var_74e7573102c1_mapped = val.(string)

		s.CreatedBy = var_74e7573102c1_mapped
	}
	if properties["updatedBy"] != nil {

		var_e01be0e7dd4d := properties["updatedBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_e01be0e7dd4d)

		if err != nil {
			panic(err)
		}

		var_e01be0e7dd4d_mapped := new(string)
		*var_e01be0e7dd4d_mapped = val.(string)

		s.UpdatedBy = var_e01be0e7dd4d_mapped
	}
	if properties["createdOn"] != nil {

		var_9e292a9af649 := properties["createdOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_9e292a9af649)

		if err != nil {
			panic(err)
		}

		var_9e292a9af649_mapped := new(time.Time)
		*var_9e292a9af649_mapped = val.(time.Time)

		s.CreatedOn = var_9e292a9af649_mapped
	}
	if properties["updatedOn"] != nil {

		var_b96ed9732dfb := properties["updatedOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_b96ed9732dfb)

		if err != nil {
			panic(err)
		}

		var_b96ed9732dfb_mapped := new(time.Time)
		*var_b96ed9732dfb_mapped = val.(time.Time)

		s.UpdatedOn = var_b96ed9732dfb_mapped
	}
	if properties["name"] != nil {

		var_166ee8f407d2 := properties["name"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_166ee8f407d2)

		if err != nil {
			panic(err)
		}

		var_166ee8f407d2_mapped := val.(string)

		s.Name = var_166ee8f407d2_mapped
	}
	if properties["description"] != nil {

		var_18425b59a7f0 := properties["description"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_18425b59a7f0)

		if err != nil {
			panic(err)
		}

		var_18425b59a7f0_mapped := new(string)
		*var_18425b59a7f0_mapped = val.(string)

		s.Description = var_18425b59a7f0_mapped
	}
	if properties["backend"] != nil {

		var_99b3c14853d5 := properties["backend"]
		var_99b3c14853d5_mapped := (DataSourceBackend)(var_99b3c14853d5.GetStringValue())

		s.Backend = var_99b3c14853d5_mapped
	}
	if properties["options"] != nil {

		var_d079905f0480 := properties["options"]
		var_d079905f0480_mapped := make(map[string]string)
		for k, v := range var_d079905f0480.GetStructValue().Fields {

			var_6f5d55545548 := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_6f5d55545548)

			if err != nil {
				panic(err)
			}

			var_6f5d55545548_mapped := val.(string)

			var_d079905f0480_mapped[k] = var_6f5d55545548_mapped
		}

		s.Options = var_d079905f0480_mapped
	}
	return s
}
