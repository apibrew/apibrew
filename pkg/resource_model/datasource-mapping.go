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

	var_c3fb8a26a7fa := dataSource.Id

	if var_c3fb8a26a7fa != nil {
		var var_c3fb8a26a7fa_mapped *structpb.Value

		var var_c3fb8a26a7fa_err error
		var_c3fb8a26a7fa_mapped, var_c3fb8a26a7fa_err = types.ByResourcePropertyType(model.ResourceProperty_UUID).Pack(*var_c3fb8a26a7fa)
		if var_c3fb8a26a7fa_err != nil {
			panic(var_c3fb8a26a7fa_err)
		}
		properties["id"] = var_c3fb8a26a7fa_mapped
	}

	var_d81e04080f1c := dataSource.Version

	var var_d81e04080f1c_mapped *structpb.Value

	var var_d81e04080f1c_err error
	var_d81e04080f1c_mapped, var_d81e04080f1c_err = types.ByResourcePropertyType(model.ResourceProperty_INT32).Pack(var_d81e04080f1c)
	if var_d81e04080f1c_err != nil {
		panic(var_d81e04080f1c_err)
	}
	properties["version"] = var_d81e04080f1c_mapped

	var_2d3bdbe9ea58 := dataSource.CreatedBy

	if var_2d3bdbe9ea58 != nil {
		var var_2d3bdbe9ea58_mapped *structpb.Value

		var var_2d3bdbe9ea58_err error
		var_2d3bdbe9ea58_mapped, var_2d3bdbe9ea58_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_2d3bdbe9ea58)
		if var_2d3bdbe9ea58_err != nil {
			panic(var_2d3bdbe9ea58_err)
		}
		properties["createdBy"] = var_2d3bdbe9ea58_mapped
	}

	var_27683c6755db := dataSource.UpdatedBy

	if var_27683c6755db != nil {
		var var_27683c6755db_mapped *structpb.Value

		var var_27683c6755db_err error
		var_27683c6755db_mapped, var_27683c6755db_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_27683c6755db)
		if var_27683c6755db_err != nil {
			panic(var_27683c6755db_err)
		}
		properties["updatedBy"] = var_27683c6755db_mapped
	}

	var_93e4b258e3d0 := dataSource.CreatedOn

	if var_93e4b258e3d0 != nil {
		var var_93e4b258e3d0_mapped *structpb.Value

		var var_93e4b258e3d0_err error
		var_93e4b258e3d0_mapped, var_93e4b258e3d0_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_93e4b258e3d0)
		if var_93e4b258e3d0_err != nil {
			panic(var_93e4b258e3d0_err)
		}
		properties["createdOn"] = var_93e4b258e3d0_mapped
	}

	var_00d30f0f6a21 := dataSource.UpdatedOn

	if var_00d30f0f6a21 != nil {
		var var_00d30f0f6a21_mapped *structpb.Value

		var var_00d30f0f6a21_err error
		var_00d30f0f6a21_mapped, var_00d30f0f6a21_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_00d30f0f6a21)
		if var_00d30f0f6a21_err != nil {
			panic(var_00d30f0f6a21_err)
		}
		properties["updatedOn"] = var_00d30f0f6a21_mapped
	}

	var_f41453c0890b := dataSource.Name

	var var_f41453c0890b_mapped *structpb.Value

	var var_f41453c0890b_err error
	var_f41453c0890b_mapped, var_f41453c0890b_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_f41453c0890b)
	if var_f41453c0890b_err != nil {
		panic(var_f41453c0890b_err)
	}
	properties["name"] = var_f41453c0890b_mapped

	var_8924b5e64929 := dataSource.Description

	var var_8924b5e64929_mapped *structpb.Value

	var var_8924b5e64929_err error
	var_8924b5e64929_mapped, var_8924b5e64929_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_8924b5e64929)
	if var_8924b5e64929_err != nil {
		panic(var_8924b5e64929_err)
	}
	properties["description"] = var_8924b5e64929_mapped

	var_a048f3842088 := dataSource.Backend

	var var_a048f3842088_mapped *structpb.Value

	var var_a048f3842088_err error
	var_a048f3842088_mapped, var_a048f3842088_err = types.ByResourcePropertyType(model.ResourceProperty_ENUM).Pack(string(var_a048f3842088))
	if var_a048f3842088_err != nil {
		panic(var_a048f3842088_err)
	}
	properties["backend"] = var_a048f3842088_mapped

	var_bb37ebbf0f2c := dataSource.Options

	var var_bb37ebbf0f2c_mapped *structpb.Value

	var var_bb37ebbf0f2c_st *structpb.Struct = new(structpb.Struct)
	var_bb37ebbf0f2c_st.Fields = make(map[string]*structpb.Value)
	for key, value := range var_bb37ebbf0f2c {

		var_5ec69640350a := value
		var var_5ec69640350a_mapped *structpb.Value

		var var_5ec69640350a_err error
		var_5ec69640350a_mapped, var_5ec69640350a_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_5ec69640350a)
		if var_5ec69640350a_err != nil {
			panic(var_5ec69640350a_err)
		}

		var_bb37ebbf0f2c_st.Fields[key] = var_5ec69640350a_mapped
	}
	var_bb37ebbf0f2c_mapped = structpb.NewStructValue(var_bb37ebbf0f2c_st)
	properties["options"] = var_bb37ebbf0f2c_mapped
	return properties
}

func (m *DataSourceMapper) FromProperties(properties map[string]*structpb.Value) *DataSource {
	var s = m.New()
	if properties["id"] != nil {

		var_d382c2966c48 := properties["id"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_UUID).UnPack(var_d382c2966c48)

		if err != nil {
			panic(err)
		}

		var_d382c2966c48_mapped := new(uuid.UUID)
		*var_d382c2966c48_mapped = val.(uuid.UUID)

		s.Id = var_d382c2966c48_mapped
	}
	if properties["version"] != nil {

		var_b73e9bfccceb := properties["version"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(var_b73e9bfccceb)

		if err != nil {
			panic(err)
		}

		var_b73e9bfccceb_mapped := val.(int32)

		s.Version = var_b73e9bfccceb_mapped
	}
	if properties["createdBy"] != nil {

		var_2e124eb3da16 := properties["createdBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_2e124eb3da16)

		if err != nil {
			panic(err)
		}

		var_2e124eb3da16_mapped := new(string)
		*var_2e124eb3da16_mapped = val.(string)

		s.CreatedBy = var_2e124eb3da16_mapped
	}
	if properties["updatedBy"] != nil {

		var_f28df641d000 := properties["updatedBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_f28df641d000)

		if err != nil {
			panic(err)
		}

		var_f28df641d000_mapped := new(string)
		*var_f28df641d000_mapped = val.(string)

		s.UpdatedBy = var_f28df641d000_mapped
	}
	if properties["createdOn"] != nil {

		var_5a76fedf7dad := properties["createdOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_5a76fedf7dad)

		if err != nil {
			panic(err)
		}

		var_5a76fedf7dad_mapped := new(time.Time)
		*var_5a76fedf7dad_mapped = val.(time.Time)

		s.CreatedOn = var_5a76fedf7dad_mapped
	}
	if properties["updatedOn"] != nil {

		var_800f0baf5bc0 := properties["updatedOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_800f0baf5bc0)

		if err != nil {
			panic(err)
		}

		var_800f0baf5bc0_mapped := new(time.Time)
		*var_800f0baf5bc0_mapped = val.(time.Time)

		s.UpdatedOn = var_800f0baf5bc0_mapped
	}
	if properties["name"] != nil {

		var_dde53a5a9e18 := properties["name"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_dde53a5a9e18)

		if err != nil {
			panic(err)
		}

		var_dde53a5a9e18_mapped := val.(string)

		s.Name = var_dde53a5a9e18_mapped
	}
	if properties["description"] != nil {

		var_3058dbc24dfe := properties["description"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_3058dbc24dfe)

		if err != nil {
			panic(err)
		}

		var_3058dbc24dfe_mapped := val.(string)

		s.Description = var_3058dbc24dfe_mapped
	}
	if properties["backend"] != nil {

		var_571e8b04e1dd := properties["backend"]
		var_571e8b04e1dd_mapped := (DataSourceBackend)(var_571e8b04e1dd.GetStringValue())

		s.Backend = var_571e8b04e1dd_mapped
	}
	if properties["options"] != nil {

		var_f6b2e2cd7878 := properties["options"]
		var_f6b2e2cd7878_mapped := make(map[string]string)
		for k, v := range var_f6b2e2cd7878.GetStructValue().Fields {

			var_d5e3ef219ee2 := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_d5e3ef219ee2)

			if err != nil {
				panic(err)
			}

			var_d5e3ef219ee2_mapped := val.(string)

			var_f6b2e2cd7878_mapped[k] = var_d5e3ef219ee2_mapped
		}

		s.Options = var_f6b2e2cd7878_mapped
	}
	return s
}
