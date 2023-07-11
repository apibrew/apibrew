package resource_model

import (
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/types"
	"google.golang.org/protobuf/types/known/structpb"
)

import "github.com/google/uuid"
import "time"

type ExtensionMapper struct {
}

func NewExtensionMapper() *ExtensionMapper {
	return &ExtensionMapper{}
}

var ExtensionMapperInstance = NewExtensionMapper()

func (m *ExtensionMapper) New() *Extension {
	return &Extension{}
}

func (m *ExtensionMapper) ToRecord(extension *Extension) *model.Record {
	var rec = &model.Record{}
	rec.Properties = m.ToProperties(extension)
	return rec
}

func (m *ExtensionMapper) FromRecord(record *model.Record) *Extension {
	return m.FromProperties(record.Properties)
}

func (m *ExtensionMapper) ToProperties(extension *Extension) map[string]*structpb.Value {
	var properties = make(map[string]*structpb.Value)

	if extension.Id != nil {
		id, err := types.ByResourcePropertyType(model.ResourceProperty_UUID).Pack(*extension.Id)
		if err != nil {
			panic(err)
		}
		properties["id"] = id
	}

	if extension.Version != nil {
		version, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).Pack(*extension.Version)
		if err != nil {
			panic(err)
		}
		properties["version"] = version
	}

	if extension.CreatedBy != nil {
		createdBy, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*extension.CreatedBy)
		if err != nil {
			panic(err)
		}
		properties["createdBy"] = createdBy
	}

	if extension.UpdatedBy != nil {
		updatedBy, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*extension.UpdatedBy)
		if err != nil {
			panic(err)
		}
		properties["updatedBy"] = updatedBy
	}

	if extension.CreatedOn != nil {
		createdOn, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*extension.CreatedOn)
		if err != nil {
			panic(err)
		}
		properties["createdOn"] = createdOn
	}

	if extension.UpdatedOn != nil {
		updatedOn, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*extension.UpdatedOn)
		if err != nil {
			panic(err)
		}
		properties["updatedOn"] = updatedOn
	}

	name, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(extension.Name)
	if err != nil {
		panic(err)
	}
	properties["name"] = name

	if extension.Description != nil {
		description, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*extension.Description)
		if err != nil {
			panic(err)
		}
		properties["description"] = description
	}

	if extension.Selector != nil {
	}

	order, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).Pack(extension.Order)
	if err != nil {
		panic(err)
	}
	properties["order"] = order

	finalizes, err := types.ByResourcePropertyType(model.ResourceProperty_BOOL).Pack(extension.Finalizes)
	if err != nil {
		panic(err)
	}
	properties["finalizes"] = finalizes

	sync, err := types.ByResourcePropertyType(model.ResourceProperty_BOOL).Pack(extension.Sync)
	if err != nil {
		panic(err)
	}
	properties["sync"] = sync

	responds, err := types.ByResourcePropertyType(model.ResourceProperty_BOOL).Pack(extension.Responds)
	if err != nil {
		panic(err)
	}
	properties["responds"] = responds

	if extension.Annotations != nil {
	}

	return properties
}

func (m *ExtensionMapper) FromProperties(properties map[string]*structpb.Value) *Extension {
	var s = m.New()
	if properties["id"] != nil {
		val, err := types.ByResourcePropertyType(model.ResourceProperty_UUID).UnPack(properties["id"])

		if err != nil {
			panic(err)
		}

		s.Id = new(uuid.UUID)
		*s.Id = val.(uuid.UUID)
	}
	if properties["version"] != nil {
		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(properties["version"])

		if err != nil {
			panic(err)
		}

		s.Version = new(int32)
		*s.Version = val.(int32)
	}
	if properties["createdBy"] != nil {
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(properties["createdBy"])

		if err != nil {
			panic(err)
		}

		s.CreatedBy = new(string)
		*s.CreatedBy = val.(string)
	}
	if properties["updatedBy"] != nil {
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(properties["updatedBy"])

		if err != nil {
			panic(err)
		}

		s.UpdatedBy = new(string)
		*s.UpdatedBy = val.(string)
	}
	if properties["createdOn"] != nil {
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(properties["createdOn"])

		if err != nil {
			panic(err)
		}

		s.CreatedOn = new(time.Time)
		*s.CreatedOn = val.(time.Time)
	}
	if properties["updatedOn"] != nil {
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(properties["updatedOn"])

		if err != nil {
			panic(err)
		}

		s.UpdatedOn = new(time.Time)
		*s.UpdatedOn = val.(time.Time)
	}
	if properties["name"] != nil {
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(properties["name"])

		if err != nil {
			panic(err)
		}

		s.Name = val.(string)
	}
	if properties["description"] != nil {
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(properties["description"])

		if err != nil {
			panic(err)
		}

		s.Description = new(string)
		*s.Description = val.(string)
	}
	if properties["selector"] != nil {
		var mappedValue = ExtensionSelectorMapperInstance.FromProperties(properties["selector"].GetStructValue().Fields)

		s.Selector = mappedValue

	}
	if properties["order"] != nil {
		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(properties["order"])

		if err != nil {
			panic(err)
		}

		s.Order = val.(int32)
	}
	if properties["finalizes"] != nil {
		val, err := types.ByResourcePropertyType(model.ResourceProperty_BOOL).UnPack(properties["finalizes"])

		if err != nil {
			panic(err)
		}

		s.Finalizes = val.(bool)
	}
	if properties["sync"] != nil {
		val, err := types.ByResourcePropertyType(model.ResourceProperty_BOOL).UnPack(properties["sync"])

		if err != nil {
			panic(err)
		}

		s.Sync = val.(bool)
	}
	if properties["responds"] != nil {
		val, err := types.ByResourcePropertyType(model.ResourceProperty_BOOL).UnPack(properties["responds"])

		if err != nil {
			panic(err)
		}

		s.Responds = val.(bool)
	}
	if properties["call"] != nil {
		var mappedValue = ExtensionCallMapperInstance.FromProperties(properties["call"].GetStructValue().Fields)

		s.Call = *mappedValue

	}
	if properties["annotations"] != nil {
		s.Annotations = make(map[string]string)
		for k, v := range properties["annotations"].GetStructValue().Fields {
			s.Annotations[k] = v.AsInterface().(string)
		}
	}
	return s
}

type ExtensionFunctionCallMapper struct {
}

func NewExtensionFunctionCallMapper() *ExtensionFunctionCallMapper {
	return &ExtensionFunctionCallMapper{}
}

var ExtensionFunctionCallMapperInstance = NewExtensionFunctionCallMapper()

func (m *ExtensionFunctionCallMapper) New() *ExtensionFunctionCall {
	return &ExtensionFunctionCall{}
}

func (m *ExtensionFunctionCallMapper) ToRecord(extensionFunctionCall *ExtensionFunctionCall) *model.Record {
	var rec = &model.Record{}
	rec.Properties = m.ToProperties(extensionFunctionCall)
	return rec
}

func (m *ExtensionFunctionCallMapper) FromRecord(record *model.Record) *ExtensionFunctionCall {
	return m.FromProperties(record.Properties)
}

func (m *ExtensionFunctionCallMapper) ToProperties(extensionFunctionCall *ExtensionFunctionCall) map[string]*structpb.Value {
	var properties = make(map[string]*structpb.Value)

	host, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(extensionFunctionCall.Host)
	if err != nil {
		panic(err)
	}
	properties["host"] = host

	functionName, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(extensionFunctionCall.FunctionName)
	if err != nil {
		panic(err)
	}
	properties["functionName"] = functionName

	return properties
}

func (m *ExtensionFunctionCallMapper) FromProperties(properties map[string]*structpb.Value) *ExtensionFunctionCall {
	var s = m.New()
	if properties["host"] != nil {
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(properties["host"])

		if err != nil {
			panic(err)
		}

		s.Host = val.(string)
	}
	if properties["functionName"] != nil {
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(properties["functionName"])

		if err != nil {
			panic(err)
		}

		s.FunctionName = val.(string)
	}
	return s
}

type ExtensionHttpCallMapper struct {
}

func NewExtensionHttpCallMapper() *ExtensionHttpCallMapper {
	return &ExtensionHttpCallMapper{}
}

var ExtensionHttpCallMapperInstance = NewExtensionHttpCallMapper()

func (m *ExtensionHttpCallMapper) New() *ExtensionHttpCall {
	return &ExtensionHttpCall{}
}

func (m *ExtensionHttpCallMapper) ToRecord(extensionHttpCall *ExtensionHttpCall) *model.Record {
	var rec = &model.Record{}
	rec.Properties = m.ToProperties(extensionHttpCall)
	return rec
}

func (m *ExtensionHttpCallMapper) FromRecord(record *model.Record) *ExtensionHttpCall {
	return m.FromProperties(record.Properties)
}

func (m *ExtensionHttpCallMapper) ToProperties(extensionHttpCall *ExtensionHttpCall) map[string]*structpb.Value {
	var properties = make(map[string]*structpb.Value)

	uri, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(extensionHttpCall.Uri)
	if err != nil {
		panic(err)
	}
	properties["uri"] = uri

	method, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(extensionHttpCall.Method)
	if err != nil {
		panic(err)
	}
	properties["method"] = method

	return properties
}

func (m *ExtensionHttpCallMapper) FromProperties(properties map[string]*structpb.Value) *ExtensionHttpCall {
	var s = m.New()
	if properties["uri"] != nil {
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(properties["uri"])

		if err != nil {
			panic(err)
		}

		s.Uri = val.(string)
	}
	if properties["method"] != nil {
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(properties["method"])

		if err != nil {
			panic(err)
		}

		s.Method = val.(string)
	}
	return s
}

type ExtensionCallMapper struct {
}

func NewExtensionCallMapper() *ExtensionCallMapper {
	return &ExtensionCallMapper{}
}

var ExtensionCallMapperInstance = NewExtensionCallMapper()

func (m *ExtensionCallMapper) New() *ExtensionCall {
	return &ExtensionCall{}
}

func (m *ExtensionCallMapper) ToRecord(extensionCall *ExtensionCall) *model.Record {
	var rec = &model.Record{}
	rec.Properties = m.ToProperties(extensionCall)
	return rec
}

func (m *ExtensionCallMapper) FromRecord(record *model.Record) *ExtensionCall {
	return m.FromProperties(record.Properties)
}

func (m *ExtensionCallMapper) ToProperties(extensionCall *ExtensionCall) map[string]*structpb.Value {
	var properties = make(map[string]*structpb.Value)

	if extensionCall.FunctionCall != nil {
	}

	if extensionCall.HttpCall != nil {
	}

	return properties
}

func (m *ExtensionCallMapper) FromProperties(properties map[string]*structpb.Value) *ExtensionCall {
	var s = m.New()
	if properties["functionCall"] != nil {
		var mappedValue = ExtensionFunctionCallMapperInstance.FromProperties(properties["functionCall"].GetStructValue().Fields)

		s.FunctionCall = mappedValue

	}
	if properties["httpCall"] != nil {
		var mappedValue = ExtensionHttpCallMapperInstance.FromProperties(properties["httpCall"].GetStructValue().Fields)

		s.HttpCall = mappedValue

	}
	return s
}

type ExtensionBooleanExpressionMapper struct {
}

func NewExtensionBooleanExpressionMapper() *ExtensionBooleanExpressionMapper {
	return &ExtensionBooleanExpressionMapper{}
}

var ExtensionBooleanExpressionMapperInstance = NewExtensionBooleanExpressionMapper()

func (m *ExtensionBooleanExpressionMapper) New() *ExtensionBooleanExpression {
	return &ExtensionBooleanExpression{}
}

func (m *ExtensionBooleanExpressionMapper) ToRecord(extensionBooleanExpression *ExtensionBooleanExpression) *model.Record {
	var rec = &model.Record{}
	rec.Properties = m.ToProperties(extensionBooleanExpression)
	return rec
}

func (m *ExtensionBooleanExpressionMapper) FromRecord(record *model.Record) *ExtensionBooleanExpression {
	return m.FromProperties(record.Properties)
}

func (m *ExtensionBooleanExpressionMapper) ToProperties(extensionBooleanExpression *ExtensionBooleanExpression) map[string]*structpb.Value {
	var properties = make(map[string]*structpb.Value)

	return properties
}

func (m *ExtensionBooleanExpressionMapper) FromProperties(properties map[string]*structpb.Value) *ExtensionBooleanExpression {
	var s = m.New()
	return s
}

type ExtensionSelectorMapper struct {
}

func NewExtensionSelectorMapper() *ExtensionSelectorMapper {
	return &ExtensionSelectorMapper{}
}

var ExtensionSelectorMapperInstance = NewExtensionSelectorMapper()

func (m *ExtensionSelectorMapper) New() *ExtensionSelector {
	return &ExtensionSelector{}
}

func (m *ExtensionSelectorMapper) ToRecord(extensionSelector *ExtensionSelector) *model.Record {
	var rec = &model.Record{}
	rec.Properties = m.ToProperties(extensionSelector)
	return rec
}

func (m *ExtensionSelectorMapper) FromRecord(record *model.Record) *ExtensionSelector {
	return m.FromProperties(record.Properties)
}

func (m *ExtensionSelectorMapper) ToProperties(extensionSelector *ExtensionSelector) map[string]*structpb.Value {
	var properties = make(map[string]*structpb.Value)

	if extensionSelector.Actions != nil {
	}

	if extensionSelector.RecordSelector != nil {
	}

	if extensionSelector.Namespaces != nil {
	}

	if extensionSelector.Resources != nil {
	}

	if extensionSelector.Ids != nil {
	}

	if extensionSelector.Annotations != nil {
	}

	return properties
}

func (m *ExtensionSelectorMapper) FromProperties(properties map[string]*structpb.Value) *ExtensionSelector {
	var s = m.New()
	if properties["actions"] != nil {
		s.Actions = []ExtensionActions{}
		for _, v := range properties["actions"].AsInterface().([]interface{}) {
			s.Actions = append(s.Actions, v.(ExtensionActions))
		}
	}
	if properties["recordSelector"] != nil {
		var mappedValue = ExtensionBooleanExpressionMapperInstance.FromProperties(properties["recordSelector"].GetStructValue().Fields)

		s.RecordSelector = mappedValue

	}
	if properties["namespaces"] != nil {
		s.Namespaces = []string{}
		for _, v := range properties["namespaces"].AsInterface().([]interface{}) {
			s.Namespaces = append(s.Namespaces, v.(string))
		}
	}
	if properties["resources"] != nil {
		s.Resources = []string{}
		for _, v := range properties["resources"].AsInterface().([]interface{}) {
			s.Resources = append(s.Resources, v.(string))
		}
	}
	if properties["ids"] != nil {
		s.Ids = []string{}
		for _, v := range properties["ids"].AsInterface().([]interface{}) {
			s.Ids = append(s.Ids, v.(string))
		}
	}
	if properties["annotations"] != nil {
		s.Annotations = make(map[string]string)
		for k, v := range properties["annotations"].GetStructValue().Fields {
			s.Annotations[k] = v.AsInterface().(string)
		}
	}
	return s
}
