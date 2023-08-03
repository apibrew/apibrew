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

	var_7df81b55293c := dataSource.Id

	if var_7df81b55293c != nil {
		var var_7df81b55293c_mapped *structpb.Value

		var var_7df81b55293c_err error
		var_7df81b55293c_mapped, var_7df81b55293c_err = types.ByResourcePropertyType(model.ResourceProperty_UUID).Pack(*var_7df81b55293c)
		if var_7df81b55293c_err != nil {
			panic(var_7df81b55293c_err)
		}
		properties["id"] = var_7df81b55293c_mapped
	}

	var_4b835ed9dd38 := dataSource.Version

	var var_4b835ed9dd38_mapped *structpb.Value

	var var_4b835ed9dd38_err error
	var_4b835ed9dd38_mapped, var_4b835ed9dd38_err = types.ByResourcePropertyType(model.ResourceProperty_INT32).Pack(var_4b835ed9dd38)
	if var_4b835ed9dd38_err != nil {
		panic(var_4b835ed9dd38_err)
	}
	properties["version"] = var_4b835ed9dd38_mapped

	var_432323ba4489 := dataSource.CreatedBy

	if var_432323ba4489 != nil {
		var var_432323ba4489_mapped *structpb.Value

		var var_432323ba4489_err error
		var_432323ba4489_mapped, var_432323ba4489_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_432323ba4489)
		if var_432323ba4489_err != nil {
			panic(var_432323ba4489_err)
		}
		properties["createdBy"] = var_432323ba4489_mapped
	}

	var_a5eaca0c8359 := dataSource.UpdatedBy

	if var_a5eaca0c8359 != nil {
		var var_a5eaca0c8359_mapped *structpb.Value

		var var_a5eaca0c8359_err error
		var_a5eaca0c8359_mapped, var_a5eaca0c8359_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_a5eaca0c8359)
		if var_a5eaca0c8359_err != nil {
			panic(var_a5eaca0c8359_err)
		}
		properties["updatedBy"] = var_a5eaca0c8359_mapped
	}

	var_1005090d9c05 := dataSource.CreatedOn

	if var_1005090d9c05 != nil {
		var var_1005090d9c05_mapped *structpb.Value

		var var_1005090d9c05_err error
		var_1005090d9c05_mapped, var_1005090d9c05_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_1005090d9c05)
		if var_1005090d9c05_err != nil {
			panic(var_1005090d9c05_err)
		}
		properties["createdOn"] = var_1005090d9c05_mapped
	}

	var_f3c8936a7451 := dataSource.UpdatedOn

	if var_f3c8936a7451 != nil {
		var var_f3c8936a7451_mapped *structpb.Value

		var var_f3c8936a7451_err error
		var_f3c8936a7451_mapped, var_f3c8936a7451_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_f3c8936a7451)
		if var_f3c8936a7451_err != nil {
			panic(var_f3c8936a7451_err)
		}
		properties["updatedOn"] = var_f3c8936a7451_mapped
	}

	var_b7554aa76503 := dataSource.Name

	var var_b7554aa76503_mapped *structpb.Value

	var var_b7554aa76503_err error
	var_b7554aa76503_mapped, var_b7554aa76503_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_b7554aa76503)
	if var_b7554aa76503_err != nil {
		panic(var_b7554aa76503_err)
	}
	properties["name"] = var_b7554aa76503_mapped

	var_599c32e5adeb := dataSource.Description

	var var_599c32e5adeb_mapped *structpb.Value

	var var_599c32e5adeb_err error
	var_599c32e5adeb_mapped, var_599c32e5adeb_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_599c32e5adeb)
	if var_599c32e5adeb_err != nil {
		panic(var_599c32e5adeb_err)
	}
	properties["description"] = var_599c32e5adeb_mapped

	var_f8f7f449d893 := dataSource.Backend

	var var_f8f7f449d893_mapped *structpb.Value

	var var_f8f7f449d893_err error
	var_f8f7f449d893_mapped, var_f8f7f449d893_err = types.ByResourcePropertyType(model.ResourceProperty_ENUM).Pack(string(var_f8f7f449d893))
	if var_f8f7f449d893_err != nil {
		panic(var_f8f7f449d893_err)
	}
	properties["backend"] = var_f8f7f449d893_mapped

	var_f966840a4337 := dataSource.Options

	var var_f966840a4337_mapped *structpb.Value

	var var_f966840a4337_st *structpb.Struct = new(structpb.Struct)
	var_f966840a4337_st.Fields = make(map[string]*structpb.Value)
	for key, value := range var_f966840a4337 {

		var_c8fea1efcf8c := value
		var var_c8fea1efcf8c_mapped *structpb.Value

		var var_c8fea1efcf8c_err error
		var_c8fea1efcf8c_mapped, var_c8fea1efcf8c_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_c8fea1efcf8c)
		if var_c8fea1efcf8c_err != nil {
			panic(var_c8fea1efcf8c_err)
		}

		var_f966840a4337_st.Fields[key] = var_c8fea1efcf8c_mapped
	}
	var_f966840a4337_mapped = structpb.NewStructValue(var_f966840a4337_st)
	properties["options"] = var_f966840a4337_mapped
	return properties
}

func (m *DataSourceMapper) FromProperties(properties map[string]*structpb.Value) *DataSource {
	var s = m.New()
	if properties["id"] != nil {

		var_afdb5568415e := properties["id"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_UUID).UnPack(var_afdb5568415e)

		if err != nil {
			panic(err)
		}

		var_afdb5568415e_mapped := new(uuid.UUID)
		*var_afdb5568415e_mapped = val.(uuid.UUID)

		s.Id = var_afdb5568415e_mapped
	}
	if properties["version"] != nil {

		var_4d288a71001a := properties["version"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(var_4d288a71001a)

		if err != nil {
			panic(err)
		}

		var_4d288a71001a_mapped := val.(int32)

		s.Version = var_4d288a71001a_mapped
	}
	if properties["createdBy"] != nil {

		var_6881894752d4 := properties["createdBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_6881894752d4)

		if err != nil {
			panic(err)
		}

		var_6881894752d4_mapped := new(string)
		*var_6881894752d4_mapped = val.(string)

		s.CreatedBy = var_6881894752d4_mapped
	}
	if properties["updatedBy"] != nil {

		var_a0ea9298ef8d := properties["updatedBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_a0ea9298ef8d)

		if err != nil {
			panic(err)
		}

		var_a0ea9298ef8d_mapped := new(string)
		*var_a0ea9298ef8d_mapped = val.(string)

		s.UpdatedBy = var_a0ea9298ef8d_mapped
	}
	if properties["createdOn"] != nil {

		var_bb527cd0bd48 := properties["createdOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_bb527cd0bd48)

		if err != nil {
			panic(err)
		}

		var_bb527cd0bd48_mapped := new(time.Time)
		*var_bb527cd0bd48_mapped = val.(time.Time)

		s.CreatedOn = var_bb527cd0bd48_mapped
	}
	if properties["updatedOn"] != nil {

		var_25e725cea016 := properties["updatedOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_25e725cea016)

		if err != nil {
			panic(err)
		}

		var_25e725cea016_mapped := new(time.Time)
		*var_25e725cea016_mapped = val.(time.Time)

		s.UpdatedOn = var_25e725cea016_mapped
	}
	if properties["name"] != nil {

		var_af0eab27da80 := properties["name"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_af0eab27da80)

		if err != nil {
			panic(err)
		}

		var_af0eab27da80_mapped := val.(string)

		s.Name = var_af0eab27da80_mapped
	}
	if properties["description"] != nil {

		var_96b5819c7634 := properties["description"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_96b5819c7634)

		if err != nil {
			panic(err)
		}

		var_96b5819c7634_mapped := val.(string)

		s.Description = var_96b5819c7634_mapped
	}
	if properties["backend"] != nil {

		var_7d0493753f62 := properties["backend"]
		var_7d0493753f62_mapped := (DataSourceBackend)(var_7d0493753f62.GetStringValue())

		s.Backend = var_7d0493753f62_mapped
	}
	if properties["options"] != nil {

		var_731146a381b4 := properties["options"]
		var_731146a381b4_mapped := make(map[string]string)
		for k, v := range var_731146a381b4.GetStructValue().Fields {

			var_a95a9546ee36 := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_a95a9546ee36)

			if err != nil {
				panic(err)
			}

			var_a95a9546ee36_mapped := val.(string)

			var_731146a381b4_mapped[k] = var_a95a9546ee36_mapped
		}

		s.Options = var_731146a381b4_mapped
	}
	return s
}
