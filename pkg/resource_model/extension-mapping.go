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

		var_7669c851fc5c := properties["id"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_UUID).UnPack(var_7669c851fc5c)

		if err != nil {
			panic(err)
		}

		var_7669c851fc5c_mapped := new(uuid.UUID)
		*var_7669c851fc5c_mapped = val.(uuid.UUID)

		s.Id = var_7669c851fc5c_mapped
	}
	if properties["version"] != nil {

		var_37eb5e258ef7 := properties["version"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(var_37eb5e258ef7)

		if err != nil {
			panic(err)
		}

		var_37eb5e258ef7_mapped := new(int32)
		*var_37eb5e258ef7_mapped = val.(int32)

		s.Version = var_37eb5e258ef7_mapped
	}
	if properties["createdBy"] != nil {

		var_715ebc85b7da := properties["createdBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_715ebc85b7da)

		if err != nil {
			panic(err)
		}

		var_715ebc85b7da_mapped := new(string)
		*var_715ebc85b7da_mapped = val.(string)

		s.CreatedBy = var_715ebc85b7da_mapped
	}
	if properties["updatedBy"] != nil {

		var_ccaf374ed57e := properties["updatedBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_ccaf374ed57e)

		if err != nil {
			panic(err)
		}

		var_ccaf374ed57e_mapped := new(string)
		*var_ccaf374ed57e_mapped = val.(string)

		s.UpdatedBy = var_ccaf374ed57e_mapped
	}
	if properties["createdOn"] != nil {

		var_689146f1bdb1 := properties["createdOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_689146f1bdb1)

		if err != nil {
			panic(err)
		}

		var_689146f1bdb1_mapped := new(time.Time)
		*var_689146f1bdb1_mapped = val.(time.Time)

		s.CreatedOn = var_689146f1bdb1_mapped
	}
	if properties["updatedOn"] != nil {

		var_3c7d1cdcf60a := properties["updatedOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_3c7d1cdcf60a)

		if err != nil {
			panic(err)
		}

		var_3c7d1cdcf60a_mapped := new(time.Time)
		*var_3c7d1cdcf60a_mapped = val.(time.Time)

		s.UpdatedOn = var_3c7d1cdcf60a_mapped
	}
	if properties["name"] != nil {

		var_f9cbaef75e8e := properties["name"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_f9cbaef75e8e)

		if err != nil {
			panic(err)
		}

		var_f9cbaef75e8e_mapped := val.(string)

		s.Name = var_f9cbaef75e8e_mapped
	}
	if properties["description"] != nil {

		var_1bfe7373cb5d := properties["description"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_1bfe7373cb5d)

		if err != nil {
			panic(err)
		}

		var_1bfe7373cb5d_mapped := new(string)
		*var_1bfe7373cb5d_mapped = val.(string)

		s.Description = var_1bfe7373cb5d_mapped
	}
	if properties["selector"] != nil {

		var_fbcd068362b5 := properties["selector"]
		var mappedValue = ExtensionEventSelectorMapperInstance.FromProperties(var_fbcd068362b5.GetStructValue().Fields)

		var_fbcd068362b5_mapped := mappedValue

		s.Selector = var_fbcd068362b5_mapped
	}
	if properties["order"] != nil {

		var_921277b2ee93 := properties["order"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(var_921277b2ee93)

		if err != nil {
			panic(err)
		}

		var_921277b2ee93_mapped := val.(int32)

		s.Order = var_921277b2ee93_mapped
	}
	if properties["finalizes"] != nil {

		var_8660009a969e := properties["finalizes"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_BOOL).UnPack(var_8660009a969e)

		if err != nil {
			panic(err)
		}

		var_8660009a969e_mapped := val.(bool)

		s.Finalizes = var_8660009a969e_mapped
	}
	if properties["sync"] != nil {

		var_bbc6085263ea := properties["sync"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_BOOL).UnPack(var_bbc6085263ea)

		if err != nil {
			panic(err)
		}

		var_bbc6085263ea_mapped := val.(bool)

		s.Sync = var_bbc6085263ea_mapped
	}
	if properties["responds"] != nil {

		var_db858020c536 := properties["responds"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_BOOL).UnPack(var_db858020c536)

		if err != nil {
			panic(err)
		}

		var_db858020c536_mapped := val.(bool)

		s.Responds = var_db858020c536_mapped
	}
	if properties["call"] != nil {

		var_085fffd65336 := properties["call"]
		var mappedValue = ExtensionExternalCallMapperInstance.FromProperties(var_085fffd65336.GetStructValue().Fields)

		var_085fffd65336_mapped := *mappedValue

		s.Call = var_085fffd65336_mapped
	}
	if properties["annotations"] != nil {

		var_09d6fbf5eb3a := properties["annotations"]
		var_09d6fbf5eb3a_mapped := make(map[string]string)
		for k, v := range var_09d6fbf5eb3a.GetStructValue().Fields {

			var_a80106d2c112 := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_a80106d2c112)

			if err != nil {
				panic(err)
			}

			var_a80106d2c112_mapped := val.(string)

			var_09d6fbf5eb3a_mapped[k] = var_a80106d2c112_mapped
		}

		s.Annotations = var_09d6fbf5eb3a_mapped
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

		var_ea0487d9c2a2 := properties["host"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_ea0487d9c2a2)

		if err != nil {
			panic(err)
		}

		var_ea0487d9c2a2_mapped := val.(string)

		s.Host = var_ea0487d9c2a2_mapped
	}
	if properties["functionName"] != nil {

		var_7ccebec346f4 := properties["functionName"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_7ccebec346f4)

		if err != nil {
			panic(err)
		}

		var_7ccebec346f4_mapped := val.(string)

		s.FunctionName = var_7ccebec346f4_mapped
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

		var_5fb30318fae9 := properties["uri"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_5fb30318fae9)

		if err != nil {
			panic(err)
		}

		var_5fb30318fae9_mapped := val.(string)

		s.Uri = var_5fb30318fae9_mapped
	}
	if properties["method"] != nil {

		var_7c1ee26be21c := properties["method"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_7c1ee26be21c)

		if err != nil {
			panic(err)
		}

		var_7c1ee26be21c_mapped := val.(string)

		s.Method = var_7c1ee26be21c_mapped
	}
	return s
}

type ExtensionExternalCallMapper struct {
}

func NewExtensionExternalCallMapper() *ExtensionExternalCallMapper {
	return &ExtensionExternalCallMapper{}
}

var ExtensionExternalCallMapperInstance = NewExtensionExternalCallMapper()

func (m *ExtensionExternalCallMapper) New() *ExtensionExternalCall {
	return &ExtensionExternalCall{}
}

func (m *ExtensionExternalCallMapper) ToRecord(extensionExternalCall *ExtensionExternalCall) *model.Record {
	var rec = &model.Record{}
	rec.Properties = m.ToProperties(extensionExternalCall)
	return rec
}

func (m *ExtensionExternalCallMapper) FromRecord(record *model.Record) *ExtensionExternalCall {
	return m.FromProperties(record.Properties)
}

func (m *ExtensionExternalCallMapper) ToProperties(extensionExternalCall *ExtensionExternalCall) map[string]*structpb.Value {
	var properties = make(map[string]*structpb.Value)

	if extensionExternalCall.FunctionCall != nil {
	}

	if extensionExternalCall.HttpCall != nil {
	}

	return properties
}

func (m *ExtensionExternalCallMapper) FromProperties(properties map[string]*structpb.Value) *ExtensionExternalCall {
	var s = m.New()
	if properties["functionCall"] != nil {

		var_ebeb4dddbb1e := properties["functionCall"]
		var mappedValue = ExtensionFunctionCallMapperInstance.FromProperties(var_ebeb4dddbb1e.GetStructValue().Fields)

		var_ebeb4dddbb1e_mapped := mappedValue

		s.FunctionCall = var_ebeb4dddbb1e_mapped
	}
	if properties["httpCall"] != nil {

		var_43be3924ff90 := properties["httpCall"]
		var mappedValue = ExtensionHttpCallMapperInstance.FromProperties(var_43be3924ff90.GetStructValue().Fields)

		var_43be3924ff90_mapped := mappedValue

		s.HttpCall = var_43be3924ff90_mapped
	}
	return s
}

type ExtensionEventSelectorMapper struct {
}

func NewExtensionEventSelectorMapper() *ExtensionEventSelectorMapper {
	return &ExtensionEventSelectorMapper{}
}

var ExtensionEventSelectorMapperInstance = NewExtensionEventSelectorMapper()

func (m *ExtensionEventSelectorMapper) New() *ExtensionEventSelector {
	return &ExtensionEventSelector{}
}

func (m *ExtensionEventSelectorMapper) ToRecord(extensionEventSelector *ExtensionEventSelector) *model.Record {
	var rec = &model.Record{}
	rec.Properties = m.ToProperties(extensionEventSelector)
	return rec
}

func (m *ExtensionEventSelectorMapper) FromRecord(record *model.Record) *ExtensionEventSelector {
	return m.FromProperties(record.Properties)
}

func (m *ExtensionEventSelectorMapper) ToProperties(extensionEventSelector *ExtensionEventSelector) map[string]*structpb.Value {
	var properties = make(map[string]*structpb.Value)

	if extensionEventSelector.Actions != nil {
	}

	if extensionEventSelector.RecordSelector != nil {
	}

	if extensionEventSelector.Namespaces != nil {
	}

	if extensionEventSelector.Resources != nil {
	}

	if extensionEventSelector.Ids != nil {
	}

	if extensionEventSelector.Annotations != nil {
	}

	return properties
}

func (m *ExtensionEventSelectorMapper) FromProperties(properties map[string]*structpb.Value) *ExtensionEventSelector {
	var s = m.New()
	if properties["actions"] != nil {

		var_a3bde283de12 := properties["actions"]
		var_a3bde283de12_mapped := []EventAction{}
		for _, v := range var_a3bde283de12.GetListValue().Values {

			var_248b51a50f7b := v
			var_248b51a50f7b_mapped := (EventAction)(var_248b51a50f7b.GetStringValue())

			var_a3bde283de12_mapped = append(var_a3bde283de12_mapped, var_248b51a50f7b_mapped)
		}

		s.Actions = var_a3bde283de12_mapped
	}
	if properties["recordSelector"] != nil {

		var_6ffaee622b97 := properties["recordSelector"]
		var mappedValue = ExtensionBooleanExpressionMapperInstance.FromProperties(var_6ffaee622b97.GetStructValue().Fields)

		var_6ffaee622b97_mapped := mappedValue

		s.RecordSelector = var_6ffaee622b97_mapped
	}
	if properties["namespaces"] != nil {

		var_fc64e2f70b36 := properties["namespaces"]
		var_fc64e2f70b36_mapped := []string{}
		for _, v := range var_fc64e2f70b36.GetListValue().Values {

			var_ae4074a73471 := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_ae4074a73471)

			if err != nil {
				panic(err)
			}

			var_ae4074a73471_mapped := val.(string)

			var_fc64e2f70b36_mapped = append(var_fc64e2f70b36_mapped, var_ae4074a73471_mapped)
		}

		s.Namespaces = var_fc64e2f70b36_mapped
	}
	if properties["resources"] != nil {

		var_fd5004076242 := properties["resources"]
		var_fd5004076242_mapped := []string{}
		for _, v := range var_fd5004076242.GetListValue().Values {

			var_c6fb0c9891af := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_c6fb0c9891af)

			if err != nil {
				panic(err)
			}

			var_c6fb0c9891af_mapped := val.(string)

			var_fd5004076242_mapped = append(var_fd5004076242_mapped, var_c6fb0c9891af_mapped)
		}

		s.Resources = var_fd5004076242_mapped
	}
	if properties["ids"] != nil {

		var_8cc131e78ab6 := properties["ids"]
		var_8cc131e78ab6_mapped := []string{}
		for _, v := range var_8cc131e78ab6.GetListValue().Values {

			var_8ed39e7ba631 := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_8ed39e7ba631)

			if err != nil {
				panic(err)
			}

			var_8ed39e7ba631_mapped := val.(string)

			var_8cc131e78ab6_mapped = append(var_8cc131e78ab6_mapped, var_8ed39e7ba631_mapped)
		}

		s.Ids = var_8cc131e78ab6_mapped
	}
	if properties["annotations"] != nil {

		var_d053633b968b := properties["annotations"]
		var_d053633b968b_mapped := make(map[string]string)
		for k, v := range var_d053633b968b.GetStructValue().Fields {

			var_1812ad68c352 := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_1812ad68c352)

			if err != nil {
				panic(err)
			}

			var_1812ad68c352_mapped := val.(string)

			var_d053633b968b_mapped[k] = var_1812ad68c352_mapped
		}

		s.Annotations = var_d053633b968b_mapped
	}
	return s
}

type ExtensionRecordSearchParamsMapper struct {
}

func NewExtensionRecordSearchParamsMapper() *ExtensionRecordSearchParamsMapper {
	return &ExtensionRecordSearchParamsMapper{}
}

var ExtensionRecordSearchParamsMapperInstance = NewExtensionRecordSearchParamsMapper()

func (m *ExtensionRecordSearchParamsMapper) New() *ExtensionRecordSearchParams {
	return &ExtensionRecordSearchParams{}
}

func (m *ExtensionRecordSearchParamsMapper) ToRecord(extensionRecordSearchParams *ExtensionRecordSearchParams) *model.Record {
	var rec = &model.Record{}
	rec.Properties = m.ToProperties(extensionRecordSearchParams)
	return rec
}

func (m *ExtensionRecordSearchParamsMapper) FromRecord(record *model.Record) *ExtensionRecordSearchParams {
	return m.FromProperties(record.Properties)
}

func (m *ExtensionRecordSearchParamsMapper) ToProperties(extensionRecordSearchParams *ExtensionRecordSearchParams) map[string]*structpb.Value {
	var properties = make(map[string]*structpb.Value)

	if extensionRecordSearchParams.Query != nil {
	}

	if extensionRecordSearchParams.Limit != nil {
		limit, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).Pack(*extensionRecordSearchParams.Limit)
		if err != nil {
			panic(err)
		}
		properties["limit"] = limit
	}

	if extensionRecordSearchParams.Offset != nil {
		offset, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).Pack(*extensionRecordSearchParams.Offset)
		if err != nil {
			panic(err)
		}
		properties["offset"] = offset
	}

	if extensionRecordSearchParams.ResolveReferences != nil {
	}

	return properties
}

func (m *ExtensionRecordSearchParamsMapper) FromProperties(properties map[string]*structpb.Value) *ExtensionRecordSearchParams {
	var s = m.New()
	if properties["query"] != nil {

		var_ed17b8fa1408 := properties["query"]
		var mappedValue = ExtensionBooleanExpressionMapperInstance.FromProperties(var_ed17b8fa1408.GetStructValue().Fields)

		var_ed17b8fa1408_mapped := mappedValue

		s.Query = var_ed17b8fa1408_mapped
	}
	if properties["limit"] != nil {

		var_4823959b4d44 := properties["limit"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(var_4823959b4d44)

		if err != nil {
			panic(err)
		}

		var_4823959b4d44_mapped := new(int32)
		*var_4823959b4d44_mapped = val.(int32)

		s.Limit = var_4823959b4d44_mapped
	}
	if properties["offset"] != nil {

		var_a68a1460e6cf := properties["offset"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(var_a68a1460e6cf)

		if err != nil {
			panic(err)
		}

		var_a68a1460e6cf_mapped := new(int32)
		*var_a68a1460e6cf_mapped = val.(int32)

		s.Offset = var_a68a1460e6cf_mapped
	}
	if properties["resolveReferences"] != nil {

		var_35aaec2f3963 := properties["resolveReferences"]
		var_35aaec2f3963_mapped := []string{}
		for _, v := range var_35aaec2f3963.GetListValue().Values {

			var_e97b777e0cfe := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_e97b777e0cfe)

			if err != nil {
				panic(err)
			}

			var_e97b777e0cfe_mapped := val.(string)

			var_35aaec2f3963_mapped = append(var_35aaec2f3963_mapped, var_e97b777e0cfe_mapped)
		}

		s.ResolveReferences = var_35aaec2f3963_mapped
	}
	return s
}

type ExtensionEventMapper struct {
}

func NewExtensionEventMapper() *ExtensionEventMapper {
	return &ExtensionEventMapper{}
}

var ExtensionEventMapperInstance = NewExtensionEventMapper()

func (m *ExtensionEventMapper) New() *ExtensionEvent {
	return &ExtensionEvent{}
}

func (m *ExtensionEventMapper) ToRecord(extensionEvent *ExtensionEvent) *model.Record {
	var rec = &model.Record{}
	rec.Properties = m.ToProperties(extensionEvent)
	return rec
}

func (m *ExtensionEventMapper) FromRecord(record *model.Record) *ExtensionEvent {
	return m.FromProperties(record.Properties)
}

func (m *ExtensionEventMapper) ToProperties(extensionEvent *ExtensionEvent) map[string]*structpb.Value {
	var properties = make(map[string]*structpb.Value)

	if extensionEvent.Id != nil {
		id, err := types.ByResourcePropertyType(model.ResourceProperty_UUID).Pack(*extensionEvent.Id)
		if err != nil {
			panic(err)
		}
		properties["id"] = id
	}

	action, err := types.ByResourcePropertyType(model.ResourceProperty_ENUM).Pack(extensionEvent.Action)
	if err != nil {
		panic(err)
	}
	properties["action"] = action

	if extensionEvent.RecordSearchParams != nil {
	}

	if extensionEvent.ActionSummary != nil {
		actionSummary, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*extensionEvent.ActionSummary)
		if err != nil {
			panic(err)
		}
		properties["actionSummary"] = actionSummary
	}

	if extensionEvent.ActionDescription != nil {
		actionDescription, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*extensionEvent.ActionDescription)
		if err != nil {
			panic(err)
		}
		properties["actionDescription"] = actionDescription
	}

	if extensionEvent.Resource != nil {
	}

	if extensionEvent.Records != nil {
	}

	if extensionEvent.Ids != nil {
	}

	if extensionEvent.Finalizes != nil {
		finalizes, err := types.ByResourcePropertyType(model.ResourceProperty_BOOL).Pack(*extensionEvent.Finalizes)
		if err != nil {
			panic(err)
		}
		properties["finalizes"] = finalizes
	}

	if extensionEvent.Sync != nil {
		sync, err := types.ByResourcePropertyType(model.ResourceProperty_BOOL).Pack(*extensionEvent.Sync)
		if err != nil {
			panic(err)
		}
		properties["sync"] = sync
	}

	if extensionEvent.Time != nil {
		time, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*extensionEvent.Time)
		if err != nil {
			panic(err)
		}
		properties["time"] = time
	}

	if extensionEvent.Annotations != nil {
	}

	return properties
}

func (m *ExtensionEventMapper) FromProperties(properties map[string]*structpb.Value) *ExtensionEvent {
	var s = m.New()
	if properties["id"] != nil {

		var_969441cb67be := properties["id"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_UUID).UnPack(var_969441cb67be)

		if err != nil {
			panic(err)
		}

		var_969441cb67be_mapped := new(uuid.UUID)
		*var_969441cb67be_mapped = val.(uuid.UUID)

		s.Id = var_969441cb67be_mapped
	}
	if properties["action"] != nil {

		var_7e616f005de7 := properties["action"]
		var_7e616f005de7_mapped := (EventAction)(var_7e616f005de7.GetStringValue())

		s.Action = var_7e616f005de7_mapped
	}
	if properties["recordSearchParams"] != nil {

		var_750fcccf29b5 := properties["recordSearchParams"]
		var mappedValue = ExtensionRecordSearchParamsMapperInstance.FromProperties(var_750fcccf29b5.GetStructValue().Fields)

		var_750fcccf29b5_mapped := mappedValue

		s.RecordSearchParams = var_750fcccf29b5_mapped
	}
	if properties["actionSummary"] != nil {

		var_0a05d7193cab := properties["actionSummary"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_0a05d7193cab)

		if err != nil {
			panic(err)
		}

		var_0a05d7193cab_mapped := new(string)
		*var_0a05d7193cab_mapped = val.(string)

		s.ActionSummary = var_0a05d7193cab_mapped
	}
	if properties["actionDescription"] != nil {

		var_98f7fb2e9801 := properties["actionDescription"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_98f7fb2e9801)

		if err != nil {
			panic(err)
		}

		var_98f7fb2e9801_mapped := new(string)
		*var_98f7fb2e9801_mapped = val.(string)

		s.ActionDescription = var_98f7fb2e9801_mapped
	}
	if properties["resource"] != nil {

		var_c96c034c1966 := properties["resource"]
		var_c96c034c1966_mapped := ResourceMapperInstance.FromProperties(var_c96c034c1966.GetStructValue().Fields)

		s.Resource = var_c96c034c1966_mapped
	}
	if properties["records"] != nil {

		var_f87a1f21d485 := properties["records"]
		var_f87a1f21d485_mapped := []*Record{}
		for _, v := range var_f87a1f21d485.GetListValue().Values {

			var_d0557b9c9a96 := v
			var_d0557b9c9a96_mapped := RecordMapperInstance.FromProperties(var_d0557b9c9a96.GetStructValue().Fields)

			var_f87a1f21d485_mapped = append(var_f87a1f21d485_mapped, var_d0557b9c9a96_mapped)
		}

		s.Records = var_f87a1f21d485_mapped
	}
	if properties["ids"] != nil {

		var_d45079f043a2 := properties["ids"]
		var_d45079f043a2_mapped := []string{}
		for _, v := range var_d45079f043a2.GetListValue().Values {

			var_64c319e90ed2 := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_64c319e90ed2)

			if err != nil {
				panic(err)
			}

			var_64c319e90ed2_mapped := val.(string)

			var_d45079f043a2_mapped = append(var_d45079f043a2_mapped, var_64c319e90ed2_mapped)
		}

		s.Ids = var_d45079f043a2_mapped
	}
	if properties["finalizes"] != nil {

		var_0a5e3f5195cb := properties["finalizes"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_BOOL).UnPack(var_0a5e3f5195cb)

		if err != nil {
			panic(err)
		}

		var_0a5e3f5195cb_mapped := new(bool)
		*var_0a5e3f5195cb_mapped = val.(bool)

		s.Finalizes = var_0a5e3f5195cb_mapped
	}
	if properties["sync"] != nil {

		var_6d7b0cf323dd := properties["sync"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_BOOL).UnPack(var_6d7b0cf323dd)

		if err != nil {
			panic(err)
		}

		var_6d7b0cf323dd_mapped := new(bool)
		*var_6d7b0cf323dd_mapped = val.(bool)

		s.Sync = var_6d7b0cf323dd_mapped
	}
	if properties["time"] != nil {

		var_a3e9b4ebf3f6 := properties["time"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_a3e9b4ebf3f6)

		if err != nil {
			panic(err)
		}

		var_a3e9b4ebf3f6_mapped := new(time.Time)
		*var_a3e9b4ebf3f6_mapped = val.(time.Time)

		s.Time = var_a3e9b4ebf3f6_mapped
	}
	if properties["annotations"] != nil {

		var_6bb5a097b53a := properties["annotations"]
		var_6bb5a097b53a_mapped := make(map[string]string)
		for k, v := range var_6bb5a097b53a.GetStructValue().Fields {

			var_f2e3333d2547 := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_f2e3333d2547)

			if err != nil {
				panic(err)
			}

			var_f2e3333d2547_mapped := val.(string)

			var_6bb5a097b53a_mapped[k] = var_f2e3333d2547_mapped
		}

		s.Annotations = var_6bb5a097b53a_mapped
	}
	return s
}
