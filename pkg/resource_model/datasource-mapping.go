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

	version, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).Pack(dataSource.Version)
	if err != nil {
		panic(err)
	}
	properties["version"] = version

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

	description, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(dataSource.Description)
	if err != nil {
		panic(err)
	}
	properties["description"] = description

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

		var_85ca8bd7077c := properties["id"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_UUID).UnPack(var_85ca8bd7077c)

		if err != nil {
			panic(err)
		}

		var_85ca8bd7077c_mapped := new(uuid.UUID)
		*var_85ca8bd7077c_mapped = val.(uuid.UUID)

		s.Id = var_85ca8bd7077c_mapped
	}
	if properties["version"] != nil {

		var_eafda6ba4c29 := properties["version"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(var_eafda6ba4c29)

		if err != nil {
			panic(err)
		}

		var_eafda6ba4c29_mapped := val.(int32)

		s.Version = var_eafda6ba4c29_mapped
	}
	if properties["createdBy"] != nil {

		var_0a55cb5b8443 := properties["createdBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_0a55cb5b8443)

		if err != nil {
			panic(err)
		}

		var_0a55cb5b8443_mapped := new(string)
		*var_0a55cb5b8443_mapped = val.(string)

		s.CreatedBy = var_0a55cb5b8443_mapped
	}
	if properties["updatedBy"] != nil {

		var_bc016ce00751 := properties["updatedBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_bc016ce00751)

		if err != nil {
			panic(err)
		}

		var_bc016ce00751_mapped := new(string)
		*var_bc016ce00751_mapped = val.(string)

		s.UpdatedBy = var_bc016ce00751_mapped
	}
	if properties["createdOn"] != nil {

		var_78d8a65b147e := properties["createdOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_78d8a65b147e)

		if err != nil {
			panic(err)
		}

		var_78d8a65b147e_mapped := new(time.Time)
		*var_78d8a65b147e_mapped = val.(time.Time)

		s.CreatedOn = var_78d8a65b147e_mapped
	}
	if properties["updatedOn"] != nil {

		var_85533f92b4f5 := properties["updatedOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_85533f92b4f5)

		if err != nil {
			panic(err)
		}

		var_85533f92b4f5_mapped := new(time.Time)
		*var_85533f92b4f5_mapped = val.(time.Time)

		s.UpdatedOn = var_85533f92b4f5_mapped
	}
	if properties["name"] != nil {

		var_5d472c126034 := properties["name"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_5d472c126034)

		if err != nil {
			panic(err)
		}

		var_5d472c126034_mapped := val.(string)

		s.Name = var_5d472c126034_mapped
	}
	if properties["description"] != nil {

		var_b5bfdb156b7b := properties["description"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_b5bfdb156b7b)

		if err != nil {
			panic(err)
		}

		var_b5bfdb156b7b_mapped := val.(string)

		s.Description = var_b5bfdb156b7b_mapped
	}
	if properties["backend"] != nil {

		var_03add09dcb6f := properties["backend"]
		var_03add09dcb6f_mapped := (DataSourceBackend)(var_03add09dcb6f.GetStringValue())

		s.Backend = var_03add09dcb6f_mapped
	}
	if properties["options"] != nil {

		var_509472cbe647 := properties["options"]
		var_509472cbe647_mapped := make(map[string]string)
		for k, v := range var_509472cbe647.GetStructValue().Fields {

			var_956d41187d6f := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_956d41187d6f)

			if err != nil {
				panic(err)
			}

			var_956d41187d6f_mapped := val.(string)

			var_509472cbe647_mapped[k] = var_956d41187d6f_mapped
		}

		s.Options = var_509472cbe647_mapped
	}
	return s
}
