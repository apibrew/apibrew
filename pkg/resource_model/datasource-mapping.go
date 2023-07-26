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

	var_201cbd3103bf := dataSource.Id

	if var_201cbd3103bf != nil {
		var var_201cbd3103bf_mapped *structpb.Value

		var var_201cbd3103bf_err error
		var_201cbd3103bf_mapped, var_201cbd3103bf_err = types.ByResourcePropertyType(model.ResourceProperty_UUID).Pack(*var_201cbd3103bf)
		if var_201cbd3103bf_err != nil {
			panic(var_201cbd3103bf_err)
		}
		properties["id"] = var_201cbd3103bf_mapped
	}

	var_0723ac459b28 := dataSource.Version

	var var_0723ac459b28_mapped *structpb.Value

	var var_0723ac459b28_err error
	var_0723ac459b28_mapped, var_0723ac459b28_err = types.ByResourcePropertyType(model.ResourceProperty_INT32).Pack(var_0723ac459b28)
	if var_0723ac459b28_err != nil {
		panic(var_0723ac459b28_err)
	}
	properties["version"] = var_0723ac459b28_mapped

	var_b9556d043f60 := dataSource.CreatedBy

	if var_b9556d043f60 != nil {
		var var_b9556d043f60_mapped *structpb.Value

		var var_b9556d043f60_err error
		var_b9556d043f60_mapped, var_b9556d043f60_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_b9556d043f60)
		if var_b9556d043f60_err != nil {
			panic(var_b9556d043f60_err)
		}
		properties["createdBy"] = var_b9556d043f60_mapped
	}

	var_83da3e7d8491 := dataSource.UpdatedBy

	if var_83da3e7d8491 != nil {
		var var_83da3e7d8491_mapped *structpb.Value

		var var_83da3e7d8491_err error
		var_83da3e7d8491_mapped, var_83da3e7d8491_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_83da3e7d8491)
		if var_83da3e7d8491_err != nil {
			panic(var_83da3e7d8491_err)
		}
		properties["updatedBy"] = var_83da3e7d8491_mapped
	}

	var_c49057a46aba := dataSource.CreatedOn

	if var_c49057a46aba != nil {
		var var_c49057a46aba_mapped *structpb.Value

		var var_c49057a46aba_err error
		var_c49057a46aba_mapped, var_c49057a46aba_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_c49057a46aba)
		if var_c49057a46aba_err != nil {
			panic(var_c49057a46aba_err)
		}
		properties["createdOn"] = var_c49057a46aba_mapped
	}

	var_794e1e3f7e9e := dataSource.UpdatedOn

	if var_794e1e3f7e9e != nil {
		var var_794e1e3f7e9e_mapped *structpb.Value

		var var_794e1e3f7e9e_err error
		var_794e1e3f7e9e_mapped, var_794e1e3f7e9e_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_794e1e3f7e9e)
		if var_794e1e3f7e9e_err != nil {
			panic(var_794e1e3f7e9e_err)
		}
		properties["updatedOn"] = var_794e1e3f7e9e_mapped
	}

	var_84c4393ddab9 := dataSource.Name

	var var_84c4393ddab9_mapped *structpb.Value

	var var_84c4393ddab9_err error
	var_84c4393ddab9_mapped, var_84c4393ddab9_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_84c4393ddab9)
	if var_84c4393ddab9_err != nil {
		panic(var_84c4393ddab9_err)
	}
	properties["name"] = var_84c4393ddab9_mapped

	var_d7e19bda1b62 := dataSource.Description

	var var_d7e19bda1b62_mapped *structpb.Value

	var var_d7e19bda1b62_err error
	var_d7e19bda1b62_mapped, var_d7e19bda1b62_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_d7e19bda1b62)
	if var_d7e19bda1b62_err != nil {
		panic(var_d7e19bda1b62_err)
	}
	properties["description"] = var_d7e19bda1b62_mapped

	var_6fda77f9ae08 := dataSource.Backend

	var var_6fda77f9ae08_mapped *structpb.Value

	var var_6fda77f9ae08_err error
	var_6fda77f9ae08_mapped, var_6fda77f9ae08_err = types.ByResourcePropertyType(model.ResourceProperty_ENUM).Pack(string(var_6fda77f9ae08))
	if var_6fda77f9ae08_err != nil {
		panic(var_6fda77f9ae08_err)
	}
	properties["backend"] = var_6fda77f9ae08_mapped

	var_4c09e848e3d0 := dataSource.Options

	var var_4c09e848e3d0_mapped *structpb.Value

	var var_4c09e848e3d0_st *structpb.Struct = new(structpb.Struct)
	var_4c09e848e3d0_st.Fields = make(map[string]*structpb.Value)
	for key, value := range var_4c09e848e3d0 {

		var_dd57bc98a514 := value
		var var_dd57bc98a514_mapped *structpb.Value

		var var_dd57bc98a514_err error
		var_dd57bc98a514_mapped, var_dd57bc98a514_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_dd57bc98a514)
		if var_dd57bc98a514_err != nil {
			panic(var_dd57bc98a514_err)
		}

		var_4c09e848e3d0_st.Fields[key] = var_dd57bc98a514_mapped
	}
	var_4c09e848e3d0_mapped = structpb.NewStructValue(var_4c09e848e3d0_st)
	properties["options"] = var_4c09e848e3d0_mapped
	return properties
}

func (m *DataSourceMapper) FromProperties(properties map[string]*structpb.Value) *DataSource {
	var s = m.New()
	if properties["id"] != nil {

		var_c8c0bb8c17b7 := properties["id"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_UUID).UnPack(var_c8c0bb8c17b7)

		if err != nil {
			panic(err)
		}

		var_c8c0bb8c17b7_mapped := new(uuid.UUID)
		*var_c8c0bb8c17b7_mapped = val.(uuid.UUID)

		s.Id = var_c8c0bb8c17b7_mapped
	}
	if properties["version"] != nil {

		var_8bcd3c3a0ed9 := properties["version"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(var_8bcd3c3a0ed9)

		if err != nil {
			panic(err)
		}

		var_8bcd3c3a0ed9_mapped := val.(int32)

		s.Version = var_8bcd3c3a0ed9_mapped
	}
	if properties["createdBy"] != nil {

		var_33740e3849b2 := properties["createdBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_33740e3849b2)

		if err != nil {
			panic(err)
		}

		var_33740e3849b2_mapped := new(string)
		*var_33740e3849b2_mapped = val.(string)

		s.CreatedBy = var_33740e3849b2_mapped
	}
	if properties["updatedBy"] != nil {

		var_9ae3c68e91b5 := properties["updatedBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_9ae3c68e91b5)

		if err != nil {
			panic(err)
		}

		var_9ae3c68e91b5_mapped := new(string)
		*var_9ae3c68e91b5_mapped = val.(string)

		s.UpdatedBy = var_9ae3c68e91b5_mapped
	}
	if properties["createdOn"] != nil {

		var_cc02d83e3260 := properties["createdOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_cc02d83e3260)

		if err != nil {
			panic(err)
		}

		var_cc02d83e3260_mapped := new(time.Time)
		*var_cc02d83e3260_mapped = val.(time.Time)

		s.CreatedOn = var_cc02d83e3260_mapped
	}
	if properties["updatedOn"] != nil {

		var_1e633e5e2446 := properties["updatedOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_1e633e5e2446)

		if err != nil {
			panic(err)
		}

		var_1e633e5e2446_mapped := new(time.Time)
		*var_1e633e5e2446_mapped = val.(time.Time)

		s.UpdatedOn = var_1e633e5e2446_mapped
	}
	if properties["name"] != nil {

		var_d954149152b3 := properties["name"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_d954149152b3)

		if err != nil {
			panic(err)
		}

		var_d954149152b3_mapped := val.(string)

		s.Name = var_d954149152b3_mapped
	}
	if properties["description"] != nil {

		var_41e4035ffa88 := properties["description"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_41e4035ffa88)

		if err != nil {
			panic(err)
		}

		var_41e4035ffa88_mapped := val.(string)

		s.Description = var_41e4035ffa88_mapped
	}
	if properties["backend"] != nil {

		var_117a619cf0f9 := properties["backend"]
		var_117a619cf0f9_mapped := (DataSourceBackend)(var_117a619cf0f9.GetStringValue())

		s.Backend = var_117a619cf0f9_mapped
	}
	if properties["options"] != nil {

		var_468e9317a23c := properties["options"]
		var_468e9317a23c_mapped := make(map[string]string)
		for k, v := range var_468e9317a23c.GetStructValue().Fields {

			var_12398c4a37b0 := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_12398c4a37b0)

			if err != nil {
				panic(err)
			}

			var_12398c4a37b0_mapped := val.(string)

			var_468e9317a23c_mapped[k] = var_12398c4a37b0_mapped
		}

		s.Options = var_468e9317a23c_mapped
	}
	return s
}
