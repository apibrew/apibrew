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

	var_7af215b55c2b := dataSource.Id

	if var_7af215b55c2b != nil {
		var var_7af215b55c2b_mapped *structpb.Value

		var var_7af215b55c2b_err error
		var_7af215b55c2b_mapped, var_7af215b55c2b_err = types.ByResourcePropertyType(model.ResourceProperty_UUID).Pack(*var_7af215b55c2b)
		if var_7af215b55c2b_err != nil {
			panic(var_7af215b55c2b_err)
		}
		properties["id"] = var_7af215b55c2b_mapped
	}

	var_acc3f8969e30 := dataSource.Version

	var var_acc3f8969e30_mapped *structpb.Value

	var var_acc3f8969e30_err error
	var_acc3f8969e30_mapped, var_acc3f8969e30_err = types.ByResourcePropertyType(model.ResourceProperty_INT32).Pack(var_acc3f8969e30)
	if var_acc3f8969e30_err != nil {
		panic(var_acc3f8969e30_err)
	}
	properties["version"] = var_acc3f8969e30_mapped

	var_b33d654b9909 := dataSource.CreatedBy

	if var_b33d654b9909 != nil {
		var var_b33d654b9909_mapped *structpb.Value

		var var_b33d654b9909_err error
		var_b33d654b9909_mapped, var_b33d654b9909_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_b33d654b9909)
		if var_b33d654b9909_err != nil {
			panic(var_b33d654b9909_err)
		}
		properties["createdBy"] = var_b33d654b9909_mapped
	}

	var_1976eea1d611 := dataSource.UpdatedBy

	if var_1976eea1d611 != nil {
		var var_1976eea1d611_mapped *structpb.Value

		var var_1976eea1d611_err error
		var_1976eea1d611_mapped, var_1976eea1d611_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_1976eea1d611)
		if var_1976eea1d611_err != nil {
			panic(var_1976eea1d611_err)
		}
		properties["updatedBy"] = var_1976eea1d611_mapped
	}

	var_810a87cf611c := dataSource.CreatedOn

	if var_810a87cf611c != nil {
		var var_810a87cf611c_mapped *structpb.Value

		var var_810a87cf611c_err error
		var_810a87cf611c_mapped, var_810a87cf611c_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_810a87cf611c)
		if var_810a87cf611c_err != nil {
			panic(var_810a87cf611c_err)
		}
		properties["createdOn"] = var_810a87cf611c_mapped
	}

	var_86cef9b3bdfb := dataSource.UpdatedOn

	if var_86cef9b3bdfb != nil {
		var var_86cef9b3bdfb_mapped *structpb.Value

		var var_86cef9b3bdfb_err error
		var_86cef9b3bdfb_mapped, var_86cef9b3bdfb_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_86cef9b3bdfb)
		if var_86cef9b3bdfb_err != nil {
			panic(var_86cef9b3bdfb_err)
		}
		properties["updatedOn"] = var_86cef9b3bdfb_mapped
	}

	var_bd734e41b40c := dataSource.Name

	var var_bd734e41b40c_mapped *structpb.Value

	var var_bd734e41b40c_err error
	var_bd734e41b40c_mapped, var_bd734e41b40c_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_bd734e41b40c)
	if var_bd734e41b40c_err != nil {
		panic(var_bd734e41b40c_err)
	}
	properties["name"] = var_bd734e41b40c_mapped

	var_44daf018afac := dataSource.Description

	var var_44daf018afac_mapped *structpb.Value

	var var_44daf018afac_err error
	var_44daf018afac_mapped, var_44daf018afac_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_44daf018afac)
	if var_44daf018afac_err != nil {
		panic(var_44daf018afac_err)
	}
	properties["description"] = var_44daf018afac_mapped

	var_666ccaaa2486 := dataSource.Backend

	var var_666ccaaa2486_mapped *structpb.Value

	var var_666ccaaa2486_err error
	var_666ccaaa2486_mapped, var_666ccaaa2486_err = types.ByResourcePropertyType(model.ResourceProperty_ENUM).Pack(string(var_666ccaaa2486))
	if var_666ccaaa2486_err != nil {
		panic(var_666ccaaa2486_err)
	}
	properties["backend"] = var_666ccaaa2486_mapped

	var_6e4ffb3c261a := dataSource.Options

	var var_6e4ffb3c261a_mapped *structpb.Value

	var var_6e4ffb3c261a_st *structpb.Struct = new(structpb.Struct)
	var_6e4ffb3c261a_st.Fields = make(map[string]*structpb.Value)
	for key, value := range var_6e4ffb3c261a {

		var_31a2c3876093 := value
		var var_31a2c3876093_mapped *structpb.Value

		var var_31a2c3876093_err error
		var_31a2c3876093_mapped, var_31a2c3876093_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_31a2c3876093)
		if var_31a2c3876093_err != nil {
			panic(var_31a2c3876093_err)
		}

		var_6e4ffb3c261a_st.Fields[key] = var_31a2c3876093_mapped
	}
	var_6e4ffb3c261a_mapped = structpb.NewStructValue(var_6e4ffb3c261a_st)
	properties["options"] = var_6e4ffb3c261a_mapped
	return properties
}

func (m *DataSourceMapper) FromProperties(properties map[string]*structpb.Value) *DataSource {
	var s = m.New()
	if properties["id"] != nil {

		var_fb15e03e7fd8 := properties["id"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_UUID).UnPack(var_fb15e03e7fd8)

		if err != nil {
			panic(err)
		}

		var_fb15e03e7fd8_mapped := new(uuid.UUID)
		*var_fb15e03e7fd8_mapped = val.(uuid.UUID)

		s.Id = var_fb15e03e7fd8_mapped
	}
	if properties["version"] != nil {

		var_123a3af8d9b3 := properties["version"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(var_123a3af8d9b3)

		if err != nil {
			panic(err)
		}

		var_123a3af8d9b3_mapped := val.(int32)

		s.Version = var_123a3af8d9b3_mapped
	}
	if properties["createdBy"] != nil {

		var_9555f125346a := properties["createdBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_9555f125346a)

		if err != nil {
			panic(err)
		}

		var_9555f125346a_mapped := new(string)
		*var_9555f125346a_mapped = val.(string)

		s.CreatedBy = var_9555f125346a_mapped
	}
	if properties["updatedBy"] != nil {

		var_87f31b527422 := properties["updatedBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_87f31b527422)

		if err != nil {
			panic(err)
		}

		var_87f31b527422_mapped := new(string)
		*var_87f31b527422_mapped = val.(string)

		s.UpdatedBy = var_87f31b527422_mapped
	}
	if properties["createdOn"] != nil {

		var_27da025fd5b0 := properties["createdOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_27da025fd5b0)

		if err != nil {
			panic(err)
		}

		var_27da025fd5b0_mapped := new(time.Time)
		*var_27da025fd5b0_mapped = val.(time.Time)

		s.CreatedOn = var_27da025fd5b0_mapped
	}
	if properties["updatedOn"] != nil {

		var_fb2540f4dbb8 := properties["updatedOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_fb2540f4dbb8)

		if err != nil {
			panic(err)
		}

		var_fb2540f4dbb8_mapped := new(time.Time)
		*var_fb2540f4dbb8_mapped = val.(time.Time)

		s.UpdatedOn = var_fb2540f4dbb8_mapped
	}
	if properties["name"] != nil {

		var_74190e6d757f := properties["name"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_74190e6d757f)

		if err != nil {
			panic(err)
		}

		var_74190e6d757f_mapped := val.(string)

		s.Name = var_74190e6d757f_mapped
	}
	if properties["description"] != nil {

		var_68a6b848795c := properties["description"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_68a6b848795c)

		if err != nil {
			panic(err)
		}

		var_68a6b848795c_mapped := val.(string)

		s.Description = var_68a6b848795c_mapped
	}
	if properties["backend"] != nil {

		var_e449f7928010 := properties["backend"]
		var_e449f7928010_mapped := (DataSourceBackend)(var_e449f7928010.GetStringValue())

		s.Backend = var_e449f7928010_mapped
	}
	if properties["options"] != nil {

		var_5139a7750311 := properties["options"]
		var_5139a7750311_mapped := make(map[string]string)
		for k, v := range var_5139a7750311.GetStructValue().Fields {

			var_fd4a1ffcfd13 := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_fd4a1ffcfd13)

			if err != nil {
				panic(err)
			}

			var_fd4a1ffcfd13_mapped := val.(string)

			var_5139a7750311_mapped[k] = var_fd4a1ffcfd13_mapped
		}

		s.Options = var_5139a7750311_mapped
	}
	return s
}
