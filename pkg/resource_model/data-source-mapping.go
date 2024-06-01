// Code generated by apbr generate. DO NOT EDIT.
// versions:
// 	apbr generate v1.2

//go:build !codeanalysis

package resource_model

import (
	"github.com/apibrew/apibrew/pkg/abs"
	"github.com/apibrew/apibrew/pkg/formats/unstructured"
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

func (m *DataSourceMapper) ResourceIdentity() abs.ResourceIdentity {
	return abs.ResourceIdentity{
		Namespace: "system",
		Name:      "DataSource",
	}
}

func (m *DataSourceMapper) ToRecord(dataSource *DataSource) abs.RecordLike {
	return abs.NewRecordLikeWithProperties(m.ToProperties(dataSource))
}

func (m *DataSourceMapper) FromRecord(record abs.RecordLike) *DataSource {
	return m.FromProperties(record.ToStruct().GetFields())
}

func (m *DataSourceMapper) ToProperties(dataSource *DataSource) map[string]*structpb.Value {
	var properties = make(map[string]*structpb.Value)

	var_Id := dataSource.Id

	if var_Id != nil {
		var var_Id_mapped *structpb.Value

		var var_Id_err error
		var_Id_mapped, var_Id_err = types.ByResourcePropertyType(model.ResourceProperty_UUID).Pack(*var_Id)
		if var_Id_err != nil {
			panic(var_Id_err)
		}
		properties["id"] = var_Id_mapped
	}

	var_Version := dataSource.Version

	var var_Version_mapped *structpb.Value

	var var_Version_err error
	var_Version_mapped, var_Version_err = types.ByResourcePropertyType(model.ResourceProperty_INT32).Pack(var_Version)
	if var_Version_err != nil {
		panic(var_Version_err)
	}
	properties["version"] = var_Version_mapped

	var_AuditData := dataSource.AuditData

	if var_AuditData != nil {
		var var_AuditData_mapped *structpb.Value

		var_AuditData_mapped = structpb.NewStructValue(&structpb.Struct{Fields: DataSourceAuditDataMapperInstance.ToProperties(var_AuditData)})
		properties["auditData"] = var_AuditData_mapped
	}

	var_Name := dataSource.Name

	var var_Name_mapped *structpb.Value

	var var_Name_err error
	var_Name_mapped, var_Name_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_Name)
	if var_Name_err != nil {
		panic(var_Name_err)
	}
	properties["name"] = var_Name_mapped

	var_Description := dataSource.Description

	var var_Description_mapped *structpb.Value

	var var_Description_err error
	var_Description_mapped, var_Description_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_Description)
	if var_Description_err != nil {
		panic(var_Description_err)
	}
	properties["description"] = var_Description_mapped

	var_Backend := dataSource.Backend

	var var_Backend_mapped *structpb.Value

	var var_Backend_err error
	var_Backend_mapped, var_Backend_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_Backend)
	if var_Backend_err != nil {
		panic(var_Backend_err)
	}
	properties["backend"] = var_Backend_mapped

	var_Options := dataSource.Options

	var var_Options_mapped *structpb.Value

	var var_Options_st *structpb.Struct = new(structpb.Struct)
	var_Options_st.Fields = make(map[string]*structpb.Value)
	for key, value := range var_Options {

		var_1x := value
		var var_1x_mapped *structpb.Value

		var var_1x_err error
		var_1x_mapped, var_1x_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_1x)
		if var_1x_err != nil {
			panic(var_1x_err)
		}

		var_Options_st.Fields[key] = var_1x_mapped
	}
	var_Options_mapped = structpb.NewStructValue(var_Options_st)
	properties["options"] = var_Options_mapped
	return properties
}

func (m *DataSourceMapper) FromProperties(properties map[string]*structpb.Value) *DataSource {
	var s = m.New()
	if properties["id"] != nil && properties["id"].AsInterface() != nil {

		var_Id := properties["id"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_UUID).UnPack(var_Id)

		if err != nil {
			panic(err)
		}

		var_Id_mapped := new(uuid.UUID)
		*var_Id_mapped = val.(uuid.UUID)

		s.Id = var_Id_mapped
	}
	if properties["version"] != nil && properties["version"].AsInterface() != nil {

		var_Version := properties["version"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(var_Version)

		if err != nil {
			panic(err)
		}

		var_Version_mapped := val.(int32)

		s.Version = var_Version_mapped
	}
	if properties["auditData"] != nil && properties["auditData"].AsInterface() != nil {

		var_AuditData := properties["auditData"]
		var mappedValue = DataSourceAuditDataMapperInstance.FromProperties(var_AuditData.GetStructValue().Fields)

		var_AuditData_mapped := mappedValue

		s.AuditData = var_AuditData_mapped
	}
	if properties["name"] != nil && properties["name"].AsInterface() != nil {

		var_Name := properties["name"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_Name)

		if err != nil {
			panic(err)
		}

		var_Name_mapped := val.(string)

		s.Name = var_Name_mapped
	}
	if properties["description"] != nil && properties["description"].AsInterface() != nil {

		var_Description := properties["description"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_Description)

		if err != nil {
			panic(err)
		}

		var_Description_mapped := val.(string)

		s.Description = var_Description_mapped
	}
	if properties["backend"] != nil && properties["backend"].AsInterface() != nil {

		var_Backend := properties["backend"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_Backend)

		if err != nil {
			panic(err)
		}

		var_Backend_mapped := val.(string)

		s.Backend = var_Backend_mapped
	}
	if properties["options"] != nil && properties["options"].AsInterface() != nil {

		var_Options := properties["options"]
		var_Options_mapped := make(map[string]string)
		for k, v := range var_Options.GetStructValue().Fields {

			var_3x := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_3x)

			if err != nil {
				panic(err)
			}

			var_3x_mapped := val.(string)

			var_Options_mapped[k] = var_3x_mapped
		}

		s.Options = var_Options_mapped
	}
	return s
}

func (m *DataSourceMapper) ToUnstructured(dataSource *DataSource) unstructured.Unstructured {
	var properties unstructured.Unstructured = make(unstructured.Unstructured)
	properties["type"] = "system/DataSource"

	var_Id := dataSource.Id

	if var_Id != nil {
		var var_Id_mapped interface{}

		var_Id_mapped = var_Id.String()
		properties["id"] = var_Id_mapped
	}

	var_Version := dataSource.Version

	var var_Version_mapped interface{}

	var_Version_mapped = var_Version
	properties["version"] = var_Version_mapped

	var_AuditData := dataSource.AuditData

	if var_AuditData != nil {
		var var_AuditData_mapped interface{}

		var_AuditData_mapped = DataSourceAuditDataMapperInstance.ToUnstructured(var_AuditData)
		properties["auditData"] = var_AuditData_mapped
	}

	var_Name := dataSource.Name

	var var_Name_mapped interface{}

	var_Name_mapped = var_Name
	properties["name"] = var_Name_mapped

	var_Description := dataSource.Description

	var var_Description_mapped interface{}

	var_Description_mapped = var_Description
	properties["description"] = var_Description_mapped

	var_Backend := dataSource.Backend

	var var_Backend_mapped interface{}

	var_Backend_mapped = var_Backend
	properties["backend"] = var_Backend_mapped

	var_Options := dataSource.Options

	var var_Options_mapped interface{}

	var var_Options_st map[string]interface{} = make(map[string]interface{})
	for key, value := range var_Options {

		var_1x := value
		var var_1x_mapped interface{}

		var_1x_mapped = var_1x

		var_Options_st[key] = var_1x_mapped
	}
	var_Options_mapped = var_Options_st
	properties["options"] = var_Options_mapped

	return properties
}

type DataSourceAuditDataMapper struct {
}

func NewDataSourceAuditDataMapper() *DataSourceAuditDataMapper {
	return &DataSourceAuditDataMapper{}
}

var DataSourceAuditDataMapperInstance = NewDataSourceAuditDataMapper()

func (m *DataSourceAuditDataMapper) New() *DataSourceAuditData {
	return &DataSourceAuditData{}
}

func (m *DataSourceAuditDataMapper) ResourceIdentity() abs.ResourceIdentity {
	return abs.ResourceIdentity{
		Namespace: "system",
		Name:      "DataSource",
	}
}

func (m *DataSourceAuditDataMapper) ToProperties(dataSourceAuditData *DataSourceAuditData) map[string]*structpb.Value {
	var properties = make(map[string]*structpb.Value)

	var_CreatedBy := dataSourceAuditData.CreatedBy

	if var_CreatedBy != nil {
		var var_CreatedBy_mapped *structpb.Value

		var var_CreatedBy_err error
		var_CreatedBy_mapped, var_CreatedBy_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_CreatedBy)
		if var_CreatedBy_err != nil {
			panic(var_CreatedBy_err)
		}
		properties["createdBy"] = var_CreatedBy_mapped
	}

	var_UpdatedBy := dataSourceAuditData.UpdatedBy

	if var_UpdatedBy != nil {
		var var_UpdatedBy_mapped *structpb.Value

		var var_UpdatedBy_err error
		var_UpdatedBy_mapped, var_UpdatedBy_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_UpdatedBy)
		if var_UpdatedBy_err != nil {
			panic(var_UpdatedBy_err)
		}
		properties["updatedBy"] = var_UpdatedBy_mapped
	}

	var_CreatedOn := dataSourceAuditData.CreatedOn

	if var_CreatedOn != nil {
		var var_CreatedOn_mapped *structpb.Value

		var var_CreatedOn_err error
		var_CreatedOn_mapped, var_CreatedOn_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_CreatedOn)
		if var_CreatedOn_err != nil {
			panic(var_CreatedOn_err)
		}
		properties["createdOn"] = var_CreatedOn_mapped
	}

	var_UpdatedOn := dataSourceAuditData.UpdatedOn

	if var_UpdatedOn != nil {
		var var_UpdatedOn_mapped *structpb.Value

		var var_UpdatedOn_err error
		var_UpdatedOn_mapped, var_UpdatedOn_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_UpdatedOn)
		if var_UpdatedOn_err != nil {
			panic(var_UpdatedOn_err)
		}
		properties["updatedOn"] = var_UpdatedOn_mapped
	}
	return properties
}

func (m *DataSourceAuditDataMapper) FromProperties(properties map[string]*structpb.Value) *DataSourceAuditData {
	var s = m.New()
	if properties["createdBy"] != nil && properties["createdBy"].AsInterface() != nil {

		var_CreatedBy := properties["createdBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_CreatedBy)

		if err != nil {
			panic(err)
		}

		var_CreatedBy_mapped := new(string)
		*var_CreatedBy_mapped = val.(string)

		s.CreatedBy = var_CreatedBy_mapped
	}
	if properties["updatedBy"] != nil && properties["updatedBy"].AsInterface() != nil {

		var_UpdatedBy := properties["updatedBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_UpdatedBy)

		if err != nil {
			panic(err)
		}

		var_UpdatedBy_mapped := new(string)
		*var_UpdatedBy_mapped = val.(string)

		s.UpdatedBy = var_UpdatedBy_mapped
	}
	if properties["createdOn"] != nil && properties["createdOn"].AsInterface() != nil {

		var_CreatedOn := properties["createdOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_CreatedOn)

		if err != nil {
			panic(err)
		}

		var_CreatedOn_mapped := new(time.Time)
		*var_CreatedOn_mapped = val.(time.Time)

		s.CreatedOn = var_CreatedOn_mapped
	}
	if properties["updatedOn"] != nil && properties["updatedOn"].AsInterface() != nil {

		var_UpdatedOn := properties["updatedOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_UpdatedOn)

		if err != nil {
			panic(err)
		}

		var_UpdatedOn_mapped := new(time.Time)
		*var_UpdatedOn_mapped = val.(time.Time)

		s.UpdatedOn = var_UpdatedOn_mapped
	}
	return s
}

func (m *DataSourceAuditDataMapper) ToUnstructured(dataSourceAuditData *DataSourceAuditData) unstructured.Unstructured {
	var properties unstructured.Unstructured = make(unstructured.Unstructured)

	var_CreatedBy := dataSourceAuditData.CreatedBy

	if var_CreatedBy != nil {
		var var_CreatedBy_mapped interface{}

		var_CreatedBy_mapped = *var_CreatedBy
		properties["createdBy"] = var_CreatedBy_mapped
	}

	var_UpdatedBy := dataSourceAuditData.UpdatedBy

	if var_UpdatedBy != nil {
		var var_UpdatedBy_mapped interface{}

		var_UpdatedBy_mapped = *var_UpdatedBy
		properties["updatedBy"] = var_UpdatedBy_mapped
	}

	var_CreatedOn := dataSourceAuditData.CreatedOn

	if var_CreatedOn != nil {
		var var_CreatedOn_mapped interface{}

		var_CreatedOn_mapped = *var_CreatedOn
		properties["createdOn"] = var_CreatedOn_mapped
	}

	var_UpdatedOn := dataSourceAuditData.UpdatedOn

	if var_UpdatedOn != nil {
		var var_UpdatedOn_mapped interface{}

		var_UpdatedOn_mapped = *var_UpdatedOn
		properties["updatedOn"] = var_UpdatedOn_mapped
	}

	return properties
}
