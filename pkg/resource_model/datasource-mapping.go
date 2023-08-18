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

	var_5125ffb74aff := dataSource.Id

	if var_5125ffb74aff != nil {
		var var_5125ffb74aff_mapped *structpb.Value

		var var_5125ffb74aff_err error
		var_5125ffb74aff_mapped, var_5125ffb74aff_err = types.ByResourcePropertyType(model.ResourceProperty_UUID).Pack(*var_5125ffb74aff)
		if var_5125ffb74aff_err != nil {
			panic(var_5125ffb74aff_err)
		}
		properties["id"] = var_5125ffb74aff_mapped
	}

	var_445987a06405 := dataSource.Version

	var var_445987a06405_mapped *structpb.Value

	var var_445987a06405_err error
	var_445987a06405_mapped, var_445987a06405_err = types.ByResourcePropertyType(model.ResourceProperty_INT32).Pack(var_445987a06405)
	if var_445987a06405_err != nil {
		panic(var_445987a06405_err)
	}
	properties["version"] = var_445987a06405_mapped

	var_82e29fdf0fc5 := dataSource.CreatedBy

	if var_82e29fdf0fc5 != nil {
		var var_82e29fdf0fc5_mapped *structpb.Value

		var var_82e29fdf0fc5_err error
		var_82e29fdf0fc5_mapped, var_82e29fdf0fc5_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_82e29fdf0fc5)
		if var_82e29fdf0fc5_err != nil {
			panic(var_82e29fdf0fc5_err)
		}
		properties["createdBy"] = var_82e29fdf0fc5_mapped
	}

	var_decec31d903d := dataSource.UpdatedBy

	if var_decec31d903d != nil {
		var var_decec31d903d_mapped *structpb.Value

		var var_decec31d903d_err error
		var_decec31d903d_mapped, var_decec31d903d_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_decec31d903d)
		if var_decec31d903d_err != nil {
			panic(var_decec31d903d_err)
		}
		properties["updatedBy"] = var_decec31d903d_mapped
	}

	var_a707a8989b53 := dataSource.CreatedOn

	if var_a707a8989b53 != nil {
		var var_a707a8989b53_mapped *structpb.Value

		var var_a707a8989b53_err error
		var_a707a8989b53_mapped, var_a707a8989b53_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_a707a8989b53)
		if var_a707a8989b53_err != nil {
			panic(var_a707a8989b53_err)
		}
		properties["createdOn"] = var_a707a8989b53_mapped
	}

	var_20ed612f8a37 := dataSource.UpdatedOn

	if var_20ed612f8a37 != nil {
		var var_20ed612f8a37_mapped *structpb.Value

		var var_20ed612f8a37_err error
		var_20ed612f8a37_mapped, var_20ed612f8a37_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_20ed612f8a37)
		if var_20ed612f8a37_err != nil {
			panic(var_20ed612f8a37_err)
		}
		properties["updatedOn"] = var_20ed612f8a37_mapped
	}

	var_0a7c60516023 := dataSource.Name

	var var_0a7c60516023_mapped *structpb.Value

	var var_0a7c60516023_err error
	var_0a7c60516023_mapped, var_0a7c60516023_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_0a7c60516023)
	if var_0a7c60516023_err != nil {
		panic(var_0a7c60516023_err)
	}
	properties["name"] = var_0a7c60516023_mapped

	var_671acf38be41 := dataSource.Description

	var var_671acf38be41_mapped *structpb.Value

	var var_671acf38be41_err error
	var_671acf38be41_mapped, var_671acf38be41_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_671acf38be41)
	if var_671acf38be41_err != nil {
		panic(var_671acf38be41_err)
	}
	properties["description"] = var_671acf38be41_mapped

	var_7833be0dae85 := dataSource.Backend

	var var_7833be0dae85_mapped *structpb.Value

	var var_7833be0dae85_err error
	var_7833be0dae85_mapped, var_7833be0dae85_err = types.ByResourcePropertyType(model.ResourceProperty_ENUM).Pack(string(var_7833be0dae85))
	if var_7833be0dae85_err != nil {
		panic(var_7833be0dae85_err)
	}
	properties["backend"] = var_7833be0dae85_mapped

	var_5b8516550787 := dataSource.Options

	var var_5b8516550787_mapped *structpb.Value

	var var_5b8516550787_st *structpb.Struct = new(structpb.Struct)
	var_5b8516550787_st.Fields = make(map[string]*structpb.Value)
	for key, value := range var_5b8516550787 {

		var_6877338b3ec2 := value
		var var_6877338b3ec2_mapped *structpb.Value

		var var_6877338b3ec2_err error
		var_6877338b3ec2_mapped, var_6877338b3ec2_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_6877338b3ec2)
		if var_6877338b3ec2_err != nil {
			panic(var_6877338b3ec2_err)
		}

		var_5b8516550787_st.Fields[key] = var_6877338b3ec2_mapped
	}
	var_5b8516550787_mapped = structpb.NewStructValue(var_5b8516550787_st)
	properties["options"] = var_5b8516550787_mapped
	return properties
}

func (m *DataSourceMapper) FromProperties(properties map[string]*structpb.Value) *DataSource {
	var s = m.New()
	if properties["id"] != nil {

		var_73ee37eae59f := properties["id"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_UUID).UnPack(var_73ee37eae59f)

		if err != nil {
			panic(err)
		}

		var_73ee37eae59f_mapped := new(uuid.UUID)
		*var_73ee37eae59f_mapped = val.(uuid.UUID)

		s.Id = var_73ee37eae59f_mapped
	}
	if properties["version"] != nil {

		var_8a9004cd7a55 := properties["version"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(var_8a9004cd7a55)

		if err != nil {
			panic(err)
		}

		var_8a9004cd7a55_mapped := val.(int32)

		s.Version = var_8a9004cd7a55_mapped
	}
	if properties["createdBy"] != nil {

		var_3c91c093aa13 := properties["createdBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_3c91c093aa13)

		if err != nil {
			panic(err)
		}

		var_3c91c093aa13_mapped := new(string)
		*var_3c91c093aa13_mapped = val.(string)

		s.CreatedBy = var_3c91c093aa13_mapped
	}
	if properties["updatedBy"] != nil {

		var_8f4033e1ac22 := properties["updatedBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_8f4033e1ac22)

		if err != nil {
			panic(err)
		}

		var_8f4033e1ac22_mapped := new(string)
		*var_8f4033e1ac22_mapped = val.(string)

		s.UpdatedBy = var_8f4033e1ac22_mapped
	}
	if properties["createdOn"] != nil {

		var_5a10d8e30795 := properties["createdOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_5a10d8e30795)

		if err != nil {
			panic(err)
		}

		var_5a10d8e30795_mapped := new(time.Time)
		*var_5a10d8e30795_mapped = val.(time.Time)

		s.CreatedOn = var_5a10d8e30795_mapped
	}
	if properties["updatedOn"] != nil {

		var_f1102e6cc7f1 := properties["updatedOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_f1102e6cc7f1)

		if err != nil {
			panic(err)
		}

		var_f1102e6cc7f1_mapped := new(time.Time)
		*var_f1102e6cc7f1_mapped = val.(time.Time)

		s.UpdatedOn = var_f1102e6cc7f1_mapped
	}
	if properties["name"] != nil {

		var_b80865725ecb := properties["name"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_b80865725ecb)

		if err != nil {
			panic(err)
		}

		var_b80865725ecb_mapped := val.(string)

		s.Name = var_b80865725ecb_mapped
	}
	if properties["description"] != nil {

		var_b5f418d64b47 := properties["description"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_b5f418d64b47)

		if err != nil {
			panic(err)
		}

		var_b5f418d64b47_mapped := val.(string)

		s.Description = var_b5f418d64b47_mapped
	}
	if properties["backend"] != nil {

		var_2cf3958c864f := properties["backend"]
		var_2cf3958c864f_mapped := (DataSourceBackend)(var_2cf3958c864f.GetStringValue())

		s.Backend = var_2cf3958c864f_mapped
	}
	if properties["options"] != nil {

		var_606621119de3 := properties["options"]
		var_606621119de3_mapped := make(map[string]string)
		for k, v := range var_606621119de3.GetStructValue().Fields {

			var_60064b15fd24 := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_60064b15fd24)

			if err != nil {
				panic(err)
			}

			var_60064b15fd24_mapped := val.(string)

			var_606621119de3_mapped[k] = var_60064b15fd24_mapped
		}

		s.Options = var_606621119de3_mapped
	}
	return s
}
