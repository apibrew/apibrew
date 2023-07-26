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

	var_7e900e6e2a1c := dataSource.Id

	if var_7e900e6e2a1c != nil {
		var var_7e900e6e2a1c_mapped *structpb.Value

		var var_7e900e6e2a1c_err error
		var_7e900e6e2a1c_mapped, var_7e900e6e2a1c_err = types.ByResourcePropertyType(model.ResourceProperty_UUID).Pack(*var_7e900e6e2a1c)
		if var_7e900e6e2a1c_err != nil {
			panic(var_7e900e6e2a1c_err)
		}
		properties["id"] = var_7e900e6e2a1c_mapped
	}

	var_1fb2bf09b6a3 := dataSource.Version

	var var_1fb2bf09b6a3_mapped *structpb.Value

	var var_1fb2bf09b6a3_err error
	var_1fb2bf09b6a3_mapped, var_1fb2bf09b6a3_err = types.ByResourcePropertyType(model.ResourceProperty_INT32).Pack(var_1fb2bf09b6a3)
	if var_1fb2bf09b6a3_err != nil {
		panic(var_1fb2bf09b6a3_err)
	}
	properties["version"] = var_1fb2bf09b6a3_mapped

	var_b6174e029bf3 := dataSource.CreatedBy

	if var_b6174e029bf3 != nil {
		var var_b6174e029bf3_mapped *structpb.Value

		var var_b6174e029bf3_err error
		var_b6174e029bf3_mapped, var_b6174e029bf3_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_b6174e029bf3)
		if var_b6174e029bf3_err != nil {
			panic(var_b6174e029bf3_err)
		}
		properties["createdBy"] = var_b6174e029bf3_mapped
	}

	var_ddf5e668faa9 := dataSource.UpdatedBy

	if var_ddf5e668faa9 != nil {
		var var_ddf5e668faa9_mapped *structpb.Value

		var var_ddf5e668faa9_err error
		var_ddf5e668faa9_mapped, var_ddf5e668faa9_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_ddf5e668faa9)
		if var_ddf5e668faa9_err != nil {
			panic(var_ddf5e668faa9_err)
		}
		properties["updatedBy"] = var_ddf5e668faa9_mapped
	}

	var_e216ef17c7db := dataSource.CreatedOn

	if var_e216ef17c7db != nil {
		var var_e216ef17c7db_mapped *structpb.Value

		var var_e216ef17c7db_err error
		var_e216ef17c7db_mapped, var_e216ef17c7db_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_e216ef17c7db)
		if var_e216ef17c7db_err != nil {
			panic(var_e216ef17c7db_err)
		}
		properties["createdOn"] = var_e216ef17c7db_mapped
	}

	var_1507599efbe9 := dataSource.UpdatedOn

	if var_1507599efbe9 != nil {
		var var_1507599efbe9_mapped *structpb.Value

		var var_1507599efbe9_err error
		var_1507599efbe9_mapped, var_1507599efbe9_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_1507599efbe9)
		if var_1507599efbe9_err != nil {
			panic(var_1507599efbe9_err)
		}
		properties["updatedOn"] = var_1507599efbe9_mapped
	}

	var_e593bc77aed7 := dataSource.Name

	var var_e593bc77aed7_mapped *structpb.Value

	var var_e593bc77aed7_err error
	var_e593bc77aed7_mapped, var_e593bc77aed7_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_e593bc77aed7)
	if var_e593bc77aed7_err != nil {
		panic(var_e593bc77aed7_err)
	}
	properties["name"] = var_e593bc77aed7_mapped

	var_d1c5092bb483 := dataSource.Description

	var var_d1c5092bb483_mapped *structpb.Value

	var var_d1c5092bb483_err error
	var_d1c5092bb483_mapped, var_d1c5092bb483_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_d1c5092bb483)
	if var_d1c5092bb483_err != nil {
		panic(var_d1c5092bb483_err)
	}
	properties["description"] = var_d1c5092bb483_mapped

	var_daba7feace45 := dataSource.Backend

	var var_daba7feace45_mapped *structpb.Value

	var var_daba7feace45_err error
	var_daba7feace45_mapped, var_daba7feace45_err = types.ByResourcePropertyType(model.ResourceProperty_ENUM).Pack(string(var_daba7feace45))
	if var_daba7feace45_err != nil {
		panic(var_daba7feace45_err)
	}
	properties["backend"] = var_daba7feace45_mapped

	var_595a9ab9e0d4 := dataSource.Options

	var var_595a9ab9e0d4_mapped *structpb.Value

	var var_595a9ab9e0d4_st *structpb.Struct = new(structpb.Struct)
	var_595a9ab9e0d4_st.Fields = make(map[string]*structpb.Value)
	for key, value := range var_595a9ab9e0d4 {

		var_b6170e369720 := value
		var var_b6170e369720_mapped *structpb.Value

		var var_b6170e369720_err error
		var_b6170e369720_mapped, var_b6170e369720_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_b6170e369720)
		if var_b6170e369720_err != nil {
			panic(var_b6170e369720_err)
		}

		var_595a9ab9e0d4_st.Fields[key] = var_b6170e369720_mapped
	}
	var_595a9ab9e0d4_mapped = structpb.NewStructValue(var_595a9ab9e0d4_st)
	properties["options"] = var_595a9ab9e0d4_mapped
	return properties
}

func (m *DataSourceMapper) FromProperties(properties map[string]*structpb.Value) *DataSource {
	var s = m.New()
	if properties["id"] != nil {

		var_123701649c17 := properties["id"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_UUID).UnPack(var_123701649c17)

		if err != nil {
			panic(err)
		}

		var_123701649c17_mapped := new(uuid.UUID)
		*var_123701649c17_mapped = val.(uuid.UUID)

		s.Id = var_123701649c17_mapped
	}
	if properties["version"] != nil {

		var_438ee2bb85f6 := properties["version"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(var_438ee2bb85f6)

		if err != nil {
			panic(err)
		}

		var_438ee2bb85f6_mapped := val.(int32)

		s.Version = var_438ee2bb85f6_mapped
	}
	if properties["createdBy"] != nil {

		var_05eacd7a9032 := properties["createdBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_05eacd7a9032)

		if err != nil {
			panic(err)
		}

		var_05eacd7a9032_mapped := new(string)
		*var_05eacd7a9032_mapped = val.(string)

		s.CreatedBy = var_05eacd7a9032_mapped
	}
	if properties["updatedBy"] != nil {

		var_4f9c1a0db2ee := properties["updatedBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_4f9c1a0db2ee)

		if err != nil {
			panic(err)
		}

		var_4f9c1a0db2ee_mapped := new(string)
		*var_4f9c1a0db2ee_mapped = val.(string)

		s.UpdatedBy = var_4f9c1a0db2ee_mapped
	}
	if properties["createdOn"] != nil {

		var_c2240ceb7ee0 := properties["createdOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_c2240ceb7ee0)

		if err != nil {
			panic(err)
		}

		var_c2240ceb7ee0_mapped := new(time.Time)
		*var_c2240ceb7ee0_mapped = val.(time.Time)

		s.CreatedOn = var_c2240ceb7ee0_mapped
	}
	if properties["updatedOn"] != nil {

		var_44d24dd180c6 := properties["updatedOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_44d24dd180c6)

		if err != nil {
			panic(err)
		}

		var_44d24dd180c6_mapped := new(time.Time)
		*var_44d24dd180c6_mapped = val.(time.Time)

		s.UpdatedOn = var_44d24dd180c6_mapped
	}
	if properties["name"] != nil {

		var_fa2c3acb49f4 := properties["name"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_fa2c3acb49f4)

		if err != nil {
			panic(err)
		}

		var_fa2c3acb49f4_mapped := val.(string)

		s.Name = var_fa2c3acb49f4_mapped
	}
	if properties["description"] != nil {

		var_d58dcebf9cc6 := properties["description"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_d58dcebf9cc6)

		if err != nil {
			panic(err)
		}

		var_d58dcebf9cc6_mapped := val.(string)

		s.Description = var_d58dcebf9cc6_mapped
	}
	if properties["backend"] != nil {

		var_6070b4f4b756 := properties["backend"]
		var_6070b4f4b756_mapped := (DataSourceBackend)(var_6070b4f4b756.GetStringValue())

		s.Backend = var_6070b4f4b756_mapped
	}
	if properties["options"] != nil {

		var_bd9b822439a1 := properties["options"]
		var_bd9b822439a1_mapped := make(map[string]string)
		for k, v := range var_bd9b822439a1.GetStructValue().Fields {

			var_9400798c1d6c := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_9400798c1d6c)

			if err != nil {
				panic(err)
			}

			var_9400798c1d6c_mapped := val.(string)

			var_bd9b822439a1_mapped[k] = var_9400798c1d6c_mapped
		}

		s.Options = var_bd9b822439a1_mapped
	}
	return s
}
