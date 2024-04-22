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

type AuditLogMapper struct {
}

func NewAuditLogMapper() *AuditLogMapper {
	return &AuditLogMapper{}
}

var AuditLogMapperInstance = NewAuditLogMapper()

func (m *AuditLogMapper) New() *AuditLog {
	return &AuditLog{}
}

func (m *AuditLogMapper) ResourceIdentity() abs.ResourceIdentity {
	return abs.ResourceIdentity{
		Namespace: "system",
		Name:      "AuditLog",
	}
}

func (m *AuditLogMapper) ToRecord(auditLog *AuditLog) *model.Record {
	var rec = &model.Record{}
	rec.Properties = m.ToProperties(auditLog)
	return rec
}

func (m *AuditLogMapper) FromRecord(record *model.Record) *AuditLog {
	return m.FromProperties(record.Properties)
}

func (m *AuditLogMapper) ToProperties(auditLog *AuditLog) map[string]*structpb.Value {
	var properties = make(map[string]*structpb.Value)

	var_Id := auditLog.Id

	if var_Id != nil {
		var var_Id_mapped *structpb.Value

		var var_Id_err error
		var_Id_mapped, var_Id_err = types.ByResourcePropertyType(model.ResourceProperty_UUID).Pack(*var_Id)
		if var_Id_err != nil {
			panic(var_Id_err)
		}
		properties["id"] = var_Id_mapped
	}

	var_Version := auditLog.Version

	var var_Version_mapped *structpb.Value

	var var_Version_err error
	var_Version_mapped, var_Version_err = types.ByResourcePropertyType(model.ResourceProperty_INT32).Pack(var_Version)
	if var_Version_err != nil {
		panic(var_Version_err)
	}
	properties["version"] = var_Version_mapped

	var_Namespace := auditLog.Namespace

	var var_Namespace_mapped *structpb.Value

	var var_Namespace_err error
	var_Namespace_mapped, var_Namespace_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_Namespace)
	if var_Namespace_err != nil {
		panic(var_Namespace_err)
	}
	properties["namespace"] = var_Namespace_mapped

	var_Resource := auditLog.Resource

	var var_Resource_mapped *structpb.Value

	var var_Resource_err error
	var_Resource_mapped, var_Resource_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_Resource)
	if var_Resource_err != nil {
		panic(var_Resource_err)
	}
	properties["resource"] = var_Resource_mapped

	var_RecordId := auditLog.RecordId

	var var_RecordId_mapped *structpb.Value

	var var_RecordId_err error
	var_RecordId_mapped, var_RecordId_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_RecordId)
	if var_RecordId_err != nil {
		panic(var_RecordId_err)
	}
	properties["recordId"] = var_RecordId_mapped

	var_Time := auditLog.Time

	var var_Time_mapped *structpb.Value

	var var_Time_err error
	var_Time_mapped, var_Time_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(var_Time)
	if var_Time_err != nil {
		panic(var_Time_err)
	}
	properties["time"] = var_Time_mapped

	var_Username := auditLog.Username

	var var_Username_mapped *structpb.Value

	var var_Username_err error
	var_Username_mapped, var_Username_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_Username)
	if var_Username_err != nil {
		panic(var_Username_err)
	}
	properties["username"] = var_Username_mapped

	var_Operation := auditLog.Operation

	var var_Operation_mapped *structpb.Value

	var var_Operation_err error
	var_Operation_mapped, var_Operation_err = types.ByResourcePropertyType(model.ResourceProperty_ENUM).Pack(string(var_Operation))
	if var_Operation_err != nil {
		panic(var_Operation_err)
	}
	properties["operation"] = var_Operation_mapped

	var_Properties := auditLog.Properties

	if var_Properties != nil {
		var var_Properties_mapped *structpb.Value

		var var_Properties_err error
		var_Properties_mapped, var_Properties_err = types.ByResourcePropertyType(model.ResourceProperty_OBJECT).Pack(var_Properties)
		if var_Properties_err != nil {
			panic(var_Properties_err)
		}
		properties["properties"] = var_Properties_mapped
	}

	var_Annotations := auditLog.Annotations

	if var_Annotations != nil {
		var var_Annotations_mapped *structpb.Value

		var var_Annotations_st *structpb.Struct = new(structpb.Struct)
		var_Annotations_st.Fields = make(map[string]*structpb.Value)
		for key, value := range var_Annotations {

			var_1x := value
			var var_1x_mapped *structpb.Value

			var var_1x_err error
			var_1x_mapped, var_1x_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_1x)
			if var_1x_err != nil {
				panic(var_1x_err)
			}

			var_Annotations_st.Fields[key] = var_1x_mapped
		}
		var_Annotations_mapped = structpb.NewStructValue(var_Annotations_st)
		properties["annotations"] = var_Annotations_mapped
	}
	return properties
}

func (m *AuditLogMapper) FromProperties(properties map[string]*structpb.Value) *AuditLog {
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
	if properties["namespace"] != nil && properties["namespace"].AsInterface() != nil {

		var_Namespace := properties["namespace"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_Namespace)

		if err != nil {
			panic(err)
		}

		var_Namespace_mapped := val.(string)

		s.Namespace = var_Namespace_mapped
	}
	if properties["resource"] != nil && properties["resource"].AsInterface() != nil {

		var_Resource := properties["resource"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_Resource)

		if err != nil {
			panic(err)
		}

		var_Resource_mapped := val.(string)

		s.Resource = var_Resource_mapped
	}
	if properties["recordId"] != nil && properties["recordId"].AsInterface() != nil {

		var_RecordId := properties["recordId"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_RecordId)

		if err != nil {
			panic(err)
		}

		var_RecordId_mapped := val.(string)

		s.RecordId = var_RecordId_mapped
	}
	if properties["time"] != nil && properties["time"].AsInterface() != nil {

		var_Time := properties["time"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_Time)

		if err != nil {
			panic(err)
		}

		var_Time_mapped := val.(time.Time)

		s.Time = var_Time_mapped
	}
	if properties["username"] != nil && properties["username"].AsInterface() != nil {

		var_Username := properties["username"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_Username)

		if err != nil {
			panic(err)
		}

		var_Username_mapped := val.(string)

		s.Username = var_Username_mapped
	}
	if properties["operation"] != nil && properties["operation"].AsInterface() != nil {

		var_Operation := properties["operation"]
		var_Operation_mapped := (AuditLogOperation)(var_Operation.GetStringValue())

		s.Operation = var_Operation_mapped
	}
	if properties["properties"] != nil && properties["properties"].AsInterface() != nil {

		var_Properties := properties["properties"]
		var_Properties_mapped := new(interface{})
		*var_Properties_mapped = unstructured.FromValue(var_Properties)

		s.Properties = var_Properties_mapped
	}
	if properties["annotations"] != nil && properties["annotations"].AsInterface() != nil {

		var_Annotations := properties["annotations"]
		var_Annotations_mapped := make(map[string]string)
		for k, v := range var_Annotations.GetStructValue().Fields {

			var_3x := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_3x)

			if err != nil {
				panic(err)
			}

			var_3x_mapped := val.(string)

			var_Annotations_mapped[k] = var_3x_mapped
		}

		s.Annotations = var_Annotations_mapped
	}
	return s
}

func (m *AuditLogMapper) ToUnstructured(auditLog *AuditLog) unstructured.Unstructured {
	var properties unstructured.Unstructured = make(unstructured.Unstructured)
	properties["type"] = "system/AuditLog"

	var_Id := auditLog.Id

	if var_Id != nil {
		var var_Id_mapped interface{}

		var_Id_mapped = var_Id.String()
		properties["id"] = var_Id_mapped
	}

	var_Version := auditLog.Version

	var var_Version_mapped interface{}

	var_Version_mapped = var_Version
	properties["version"] = var_Version_mapped

	var_Namespace := auditLog.Namespace

	var var_Namespace_mapped interface{}

	var_Namespace_mapped = var_Namespace
	properties["namespace"] = var_Namespace_mapped

	var_Resource := auditLog.Resource

	var var_Resource_mapped interface{}

	var_Resource_mapped = var_Resource
	properties["resource"] = var_Resource_mapped

	var_RecordId := auditLog.RecordId

	var var_RecordId_mapped interface{}

	var_RecordId_mapped = var_RecordId
	properties["recordId"] = var_RecordId_mapped

	var_Time := auditLog.Time

	var var_Time_mapped interface{}

	var_Time_mapped = var_Time
	properties["time"] = var_Time_mapped

	var_Username := auditLog.Username

	var var_Username_mapped interface{}

	var_Username_mapped = var_Username
	properties["username"] = var_Username_mapped

	var_Operation := auditLog.Operation

	var var_Operation_mapped interface{}

	var_Operation_mapped = string(var_Operation)
	properties["operation"] = var_Operation_mapped

	var_Properties := auditLog.Properties

	if var_Properties != nil {
		var var_Properties_mapped interface{}

		var_Properties_mapped = var_Properties
		properties["properties"] = var_Properties_mapped
	}

	var_Annotations := auditLog.Annotations

	if var_Annotations != nil {
		var var_Annotations_mapped interface{}

		var var_Annotations_st map[string]interface{} = make(map[string]interface{})
		for key, value := range var_Annotations {

			var_1x := value
			var var_1x_mapped interface{}

			var_1x_mapped = var_1x

			var_Annotations_st[key] = var_1x_mapped
		}
		var_Annotations_mapped = var_Annotations_st
		properties["annotations"] = var_Annotations_mapped
	}

	return properties
}
