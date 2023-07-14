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

		var_4fcc52ece8cf := properties["id"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_UUID).UnPack(var_4fcc52ece8cf)

		if err != nil {
			panic(err)
		}

		var_4fcc52ece8cf_mapped := new(uuid.UUID)
		*var_4fcc52ece8cf_mapped = val.(uuid.UUID)

		s.Id = var_4fcc52ece8cf_mapped
	}
	if properties["version"] != nil {

		var_6aabfce62e5d := properties["version"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(var_6aabfce62e5d)

		if err != nil {
			panic(err)
		}

		var_6aabfce62e5d_mapped := new(int32)
		*var_6aabfce62e5d_mapped = val.(int32)

		s.Version = var_6aabfce62e5d_mapped
	}
	if properties["createdBy"] != nil {

		var_ae778e40e1f9 := properties["createdBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_ae778e40e1f9)

		if err != nil {
			panic(err)
		}

		var_ae778e40e1f9_mapped := new(string)
		*var_ae778e40e1f9_mapped = val.(string)

		s.CreatedBy = var_ae778e40e1f9_mapped
	}
	if properties["updatedBy"] != nil {

		var_684d7d8f8328 := properties["updatedBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_684d7d8f8328)

		if err != nil {
			panic(err)
		}

		var_684d7d8f8328_mapped := new(string)
		*var_684d7d8f8328_mapped = val.(string)

		s.UpdatedBy = var_684d7d8f8328_mapped
	}
	if properties["createdOn"] != nil {

		var_5238be6f8e82 := properties["createdOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_5238be6f8e82)

		if err != nil {
			panic(err)
		}

		var_5238be6f8e82_mapped := new(time.Time)
		*var_5238be6f8e82_mapped = val.(time.Time)

		s.CreatedOn = var_5238be6f8e82_mapped
	}
	if properties["updatedOn"] != nil {

		var_79d03c0b10aa := properties["updatedOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_79d03c0b10aa)

		if err != nil {
			panic(err)
		}

		var_79d03c0b10aa_mapped := new(time.Time)
		*var_79d03c0b10aa_mapped = val.(time.Time)

		s.UpdatedOn = var_79d03c0b10aa_mapped
	}
	if properties["name"] != nil {

		var_895b91b78b48 := properties["name"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_895b91b78b48)

		if err != nil {
			panic(err)
		}

		var_895b91b78b48_mapped := val.(string)

		s.Name = var_895b91b78b48_mapped
	}
	if properties["description"] != nil {

		var_8b1a51f9bb5c := properties["description"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_8b1a51f9bb5c)

		if err != nil {
			panic(err)
		}

		var_8b1a51f9bb5c_mapped := new(string)
		*var_8b1a51f9bb5c_mapped = val.(string)

		s.Description = var_8b1a51f9bb5c_mapped
	}
	if properties["selector"] != nil {

		var_e707c63a5187 := properties["selector"]
		var mappedValue = ExtensionEventSelectorMapperInstance.FromProperties(var_e707c63a5187.GetStructValue().Fields)

		var_e707c63a5187_mapped := mappedValue

		s.Selector = var_e707c63a5187_mapped
	}
	if properties["order"] != nil {

		var_cb5bf2040c69 := properties["order"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(var_cb5bf2040c69)

		if err != nil {
			panic(err)
		}

		var_cb5bf2040c69_mapped := val.(int32)

		s.Order = var_cb5bf2040c69_mapped
	}
	if properties["finalizes"] != nil {

		var_96e9b86f3719 := properties["finalizes"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_BOOL).UnPack(var_96e9b86f3719)

		if err != nil {
			panic(err)
		}

		var_96e9b86f3719_mapped := val.(bool)

		s.Finalizes = var_96e9b86f3719_mapped
	}
	if properties["sync"] != nil {

		var_9430bf8d319a := properties["sync"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_BOOL).UnPack(var_9430bf8d319a)

		if err != nil {
			panic(err)
		}

		var_9430bf8d319a_mapped := val.(bool)

		s.Sync = var_9430bf8d319a_mapped
	}
	if properties["responds"] != nil {

		var_bfe78ee47793 := properties["responds"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_BOOL).UnPack(var_bfe78ee47793)

		if err != nil {
			panic(err)
		}

		var_bfe78ee47793_mapped := val.(bool)

		s.Responds = var_bfe78ee47793_mapped
	}
	if properties["call"] != nil {

		var_fe2101498ab1 := properties["call"]
		var mappedValue = ExtensionExternalCallMapperInstance.FromProperties(var_fe2101498ab1.GetStructValue().Fields)

		var_fe2101498ab1_mapped := *mappedValue

		s.Call = var_fe2101498ab1_mapped
	}
	if properties["annotations"] != nil {

		var_1f3de182dc6b := properties["annotations"]
		var_1f3de182dc6b_mapped := make(map[string]string)
		for k, v := range var_1f3de182dc6b.GetStructValue().Fields {

			var_eba6615fc671 := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_eba6615fc671)

			if err != nil {
				panic(err)
			}

			var_eba6615fc671_mapped := val.(string)

			var_1f3de182dc6b_mapped[k] = var_eba6615fc671_mapped
		}

		s.Annotations = var_1f3de182dc6b_mapped
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

		var_6e61d53a7692 := properties["host"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_6e61d53a7692)

		if err != nil {
			panic(err)
		}

		var_6e61d53a7692_mapped := val.(string)

		s.Host = var_6e61d53a7692_mapped
	}
	if properties["functionName"] != nil {

		var_59270339b1a2 := properties["functionName"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_59270339b1a2)

		if err != nil {
			panic(err)
		}

		var_59270339b1a2_mapped := val.(string)

		s.FunctionName = var_59270339b1a2_mapped
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

		var_bb278a6e909c := properties["uri"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_bb278a6e909c)

		if err != nil {
			panic(err)
		}

		var_bb278a6e909c_mapped := val.(string)

		s.Uri = var_bb278a6e909c_mapped
	}
	if properties["method"] != nil {

		var_31d5ef8b4e33 := properties["method"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_31d5ef8b4e33)

		if err != nil {
			panic(err)
		}

		var_31d5ef8b4e33_mapped := val.(string)

		s.Method = var_31d5ef8b4e33_mapped
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

		var_0019557e25d0 := properties["functionCall"]
		var mappedValue = ExtensionFunctionCallMapperInstance.FromProperties(var_0019557e25d0.GetStructValue().Fields)

		var_0019557e25d0_mapped := mappedValue

		s.FunctionCall = var_0019557e25d0_mapped
	}
	if properties["httpCall"] != nil {

		var_61ce036359c8 := properties["httpCall"]
		var mappedValue = ExtensionHttpCallMapperInstance.FromProperties(var_61ce036359c8.GetStructValue().Fields)

		var_61ce036359c8_mapped := mappedValue

		s.HttpCall = var_61ce036359c8_mapped
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

		var_584ca5e992d0 := properties["actions"]
		var_584ca5e992d0_mapped := []EventAction{}
		for _, v := range var_584ca5e992d0.GetListValue().Values {

			var_b4495161ecbb := v
			var_b4495161ecbb_mapped := (EventAction)(var_b4495161ecbb.GetStringValue())

			var_584ca5e992d0_mapped = append(var_584ca5e992d0_mapped, var_b4495161ecbb_mapped)
		}

		s.Actions = var_584ca5e992d0_mapped
	}
	if properties["recordSelector"] != nil {

		var_ef0db9f8b6d8 := properties["recordSelector"]
		var mappedValue = ExtensionBooleanExpressionMapperInstance.FromProperties(var_ef0db9f8b6d8.GetStructValue().Fields)

		var_ef0db9f8b6d8_mapped := mappedValue

		s.RecordSelector = var_ef0db9f8b6d8_mapped
	}
	if properties["namespaces"] != nil {

		var_a4a5c8f1751f := properties["namespaces"]
		var_a4a5c8f1751f_mapped := []string{}
		for _, v := range var_a4a5c8f1751f.GetListValue().Values {

			var_903a145795f0 := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_903a145795f0)

			if err != nil {
				panic(err)
			}

			var_903a145795f0_mapped := val.(string)

			var_a4a5c8f1751f_mapped = append(var_a4a5c8f1751f_mapped, var_903a145795f0_mapped)
		}

		s.Namespaces = var_a4a5c8f1751f_mapped
	}
	if properties["resources"] != nil {

		var_ca49d0915b3b := properties["resources"]
		var_ca49d0915b3b_mapped := []string{}
		for _, v := range var_ca49d0915b3b.GetListValue().Values {

			var_fbaeefa2b09f := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_fbaeefa2b09f)

			if err != nil {
				panic(err)
			}

			var_fbaeefa2b09f_mapped := val.(string)

			var_ca49d0915b3b_mapped = append(var_ca49d0915b3b_mapped, var_fbaeefa2b09f_mapped)
		}

		s.Resources = var_ca49d0915b3b_mapped
	}
	if properties["ids"] != nil {

		var_0df630657ccc := properties["ids"]
		var_0df630657ccc_mapped := []string{}
		for _, v := range var_0df630657ccc.GetListValue().Values {

			var_bb84f1885e33 := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_bb84f1885e33)

			if err != nil {
				panic(err)
			}

			var_bb84f1885e33_mapped := val.(string)

			var_0df630657ccc_mapped = append(var_0df630657ccc_mapped, var_bb84f1885e33_mapped)
		}

		s.Ids = var_0df630657ccc_mapped
	}
	if properties["annotations"] != nil {

		var_2665ece16d58 := properties["annotations"]
		var_2665ece16d58_mapped := make(map[string]string)
		for k, v := range var_2665ece16d58.GetStructValue().Fields {

			var_f7aec8d08f15 := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_f7aec8d08f15)

			if err != nil {
				panic(err)
			}

			var_f7aec8d08f15_mapped := val.(string)

			var_2665ece16d58_mapped[k] = var_f7aec8d08f15_mapped
		}

		s.Annotations = var_2665ece16d58_mapped
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

		var_d77a3a92bd5b := properties["query"]
		var mappedValue = ExtensionBooleanExpressionMapperInstance.FromProperties(var_d77a3a92bd5b.GetStructValue().Fields)

		var_d77a3a92bd5b_mapped := mappedValue

		s.Query = var_d77a3a92bd5b_mapped
	}
	if properties["limit"] != nil {

		var_17e8a9227e9a := properties["limit"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(var_17e8a9227e9a)

		if err != nil {
			panic(err)
		}

		var_17e8a9227e9a_mapped := new(int32)
		*var_17e8a9227e9a_mapped = val.(int32)

		s.Limit = var_17e8a9227e9a_mapped
	}
	if properties["offset"] != nil {

		var_0a01cdef92e6 := properties["offset"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(var_0a01cdef92e6)

		if err != nil {
			panic(err)
		}

		var_0a01cdef92e6_mapped := new(int32)
		*var_0a01cdef92e6_mapped = val.(int32)

		s.Offset = var_0a01cdef92e6_mapped
	}
	if properties["resolveReferences"] != nil {

		var_25dc70ccb6e0 := properties["resolveReferences"]
		var_25dc70ccb6e0_mapped := []string{}
		for _, v := range var_25dc70ccb6e0.GetListValue().Values {

			var_3b2d761f22e6 := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_3b2d761f22e6)

			if err != nil {
				panic(err)
			}

			var_3b2d761f22e6_mapped := val.(string)

			var_25dc70ccb6e0_mapped = append(var_25dc70ccb6e0_mapped, var_3b2d761f22e6_mapped)
		}

		s.ResolveReferences = var_25dc70ccb6e0_mapped
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

		var_bf4147b3f116 := properties["id"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_UUID).UnPack(var_bf4147b3f116)

		if err != nil {
			panic(err)
		}

		var_bf4147b3f116_mapped := new(uuid.UUID)
		*var_bf4147b3f116_mapped = val.(uuid.UUID)

		s.Id = var_bf4147b3f116_mapped
	}
	if properties["action"] != nil {

		var_0aebe272cb6f := properties["action"]
		var_0aebe272cb6f_mapped := (EventAction)(var_0aebe272cb6f.GetStringValue())

		s.Action = var_0aebe272cb6f_mapped
	}
	if properties["recordSearchParams"] != nil {

		var_9758780d7c6b := properties["recordSearchParams"]
		var mappedValue = ExtensionRecordSearchParamsMapperInstance.FromProperties(var_9758780d7c6b.GetStructValue().Fields)

		var_9758780d7c6b_mapped := mappedValue

		s.RecordSearchParams = var_9758780d7c6b_mapped
	}
	if properties["actionSummary"] != nil {

		var_9aea458f2955 := properties["actionSummary"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_9aea458f2955)

		if err != nil {
			panic(err)
		}

		var_9aea458f2955_mapped := new(string)
		*var_9aea458f2955_mapped = val.(string)

		s.ActionSummary = var_9aea458f2955_mapped
	}
	if properties["actionDescription"] != nil {

		var_3637df191d5e := properties["actionDescription"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_3637df191d5e)

		if err != nil {
			panic(err)
		}

		var_3637df191d5e_mapped := new(string)
		*var_3637df191d5e_mapped = val.(string)

		s.ActionDescription = var_3637df191d5e_mapped
	}
	if properties["resource"] != nil {

		var_2d1260ea7528 := properties["resource"]
		var_2d1260ea7528_mapped := ResourceMapperInstance.FromProperties(var_2d1260ea7528.GetStructValue().Fields)

		s.Resource = var_2d1260ea7528_mapped
	}
	if properties["records"] != nil {

		var_c4ea3b238739 := properties["records"]
		var_c4ea3b238739_mapped := []*Record{}
		for _, v := range var_c4ea3b238739.GetListValue().Values {

			var_7e0508cf9d8b := v
			var_7e0508cf9d8b_mapped := RecordMapperInstance.FromProperties(var_7e0508cf9d8b.GetStructValue().Fields)

			var_c4ea3b238739_mapped = append(var_c4ea3b238739_mapped, var_7e0508cf9d8b_mapped)
		}

		s.Records = var_c4ea3b238739_mapped
	}
	if properties["ids"] != nil {

		var_05dd6b42172f := properties["ids"]
		var_05dd6b42172f_mapped := []string{}
		for _, v := range var_05dd6b42172f.GetListValue().Values {

			var_a372ac775801 := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_a372ac775801)

			if err != nil {
				panic(err)
			}

			var_a372ac775801_mapped := val.(string)

			var_05dd6b42172f_mapped = append(var_05dd6b42172f_mapped, var_a372ac775801_mapped)
		}

		s.Ids = var_05dd6b42172f_mapped
	}
	if properties["finalizes"] != nil {

		var_37e4f9ec6f08 := properties["finalizes"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_BOOL).UnPack(var_37e4f9ec6f08)

		if err != nil {
			panic(err)
		}

		var_37e4f9ec6f08_mapped := new(bool)
		*var_37e4f9ec6f08_mapped = val.(bool)

		s.Finalizes = var_37e4f9ec6f08_mapped
	}
	if properties["sync"] != nil {

		var_906076ad836a := properties["sync"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_BOOL).UnPack(var_906076ad836a)

		if err != nil {
			panic(err)
		}

		var_906076ad836a_mapped := new(bool)
		*var_906076ad836a_mapped = val.(bool)

		s.Sync = var_906076ad836a_mapped
	}
	if properties["time"] != nil {

		var_dbbdf3cc1930 := properties["time"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_dbbdf3cc1930)

		if err != nil {
			panic(err)
		}

		var_dbbdf3cc1930_mapped := new(time.Time)
		*var_dbbdf3cc1930_mapped = val.(time.Time)

		s.Time = var_dbbdf3cc1930_mapped
	}
	if properties["annotations"] != nil {

		var_4708683c2c0a := properties["annotations"]
		var_4708683c2c0a_mapped := make(map[string]string)
		for k, v := range var_4708683c2c0a.GetStructValue().Fields {

			var_931732234653 := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_931732234653)

			if err != nil {
				panic(err)
			}

			var_931732234653_mapped := val.(string)

			var_4708683c2c0a_mapped[k] = var_931732234653_mapped
		}

		s.Annotations = var_4708683c2c0a_mapped
	}
	return s
}
