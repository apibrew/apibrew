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

	version, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).Pack(extension.Version)
	if err != nil {
		panic(err)
	}
	properties["version"] = version

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

		var_61a5e069d580 := properties["id"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_UUID).UnPack(var_61a5e069d580)

		if err != nil {
			panic(err)
		}

		var_61a5e069d580_mapped := new(uuid.UUID)
		*var_61a5e069d580_mapped = val.(uuid.UUID)

		s.Id = var_61a5e069d580_mapped
	}
	if properties["version"] != nil {

		var_4d3fe211c570 := properties["version"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(var_4d3fe211c570)

		if err != nil {
			panic(err)
		}

		var_4d3fe211c570_mapped := val.(int32)

		s.Version = var_4d3fe211c570_mapped
	}
	if properties["createdBy"] != nil {

		var_18cff8983e0f := properties["createdBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_18cff8983e0f)

		if err != nil {
			panic(err)
		}

		var_18cff8983e0f_mapped := new(string)
		*var_18cff8983e0f_mapped = val.(string)

		s.CreatedBy = var_18cff8983e0f_mapped
	}
	if properties["updatedBy"] != nil {

		var_8e2c949019f7 := properties["updatedBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_8e2c949019f7)

		if err != nil {
			panic(err)
		}

		var_8e2c949019f7_mapped := new(string)
		*var_8e2c949019f7_mapped = val.(string)

		s.UpdatedBy = var_8e2c949019f7_mapped
	}
	if properties["createdOn"] != nil {

		var_57709d12b973 := properties["createdOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_57709d12b973)

		if err != nil {
			panic(err)
		}

		var_57709d12b973_mapped := new(time.Time)
		*var_57709d12b973_mapped = val.(time.Time)

		s.CreatedOn = var_57709d12b973_mapped
	}
	if properties["updatedOn"] != nil {

		var_9f645893d049 := properties["updatedOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_9f645893d049)

		if err != nil {
			panic(err)
		}

		var_9f645893d049_mapped := new(time.Time)
		*var_9f645893d049_mapped = val.(time.Time)

		s.UpdatedOn = var_9f645893d049_mapped
	}
	if properties["name"] != nil {

		var_b7bdf9e20b90 := properties["name"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_b7bdf9e20b90)

		if err != nil {
			panic(err)
		}

		var_b7bdf9e20b90_mapped := val.(string)

		s.Name = var_b7bdf9e20b90_mapped
	}
	if properties["description"] != nil {

		var_40fb99c78fe1 := properties["description"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_40fb99c78fe1)

		if err != nil {
			panic(err)
		}

		var_40fb99c78fe1_mapped := new(string)
		*var_40fb99c78fe1_mapped = val.(string)

		s.Description = var_40fb99c78fe1_mapped
	}
	if properties["selector"] != nil {

		var_f4cf922d018c := properties["selector"]
		var mappedValue = ExtensionEventSelectorMapperInstance.FromProperties(var_f4cf922d018c.GetStructValue().Fields)

		var_f4cf922d018c_mapped := mappedValue

		s.Selector = var_f4cf922d018c_mapped
	}
	if properties["order"] != nil {

		var_4fe469fb475d := properties["order"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(var_4fe469fb475d)

		if err != nil {
			panic(err)
		}

		var_4fe469fb475d_mapped := val.(int32)

		s.Order = var_4fe469fb475d_mapped
	}
	if properties["finalizes"] != nil {

		var_cc22556be709 := properties["finalizes"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_BOOL).UnPack(var_cc22556be709)

		if err != nil {
			panic(err)
		}

		var_cc22556be709_mapped := val.(bool)

		s.Finalizes = var_cc22556be709_mapped
	}
	if properties["sync"] != nil {

		var_84fa77833ef8 := properties["sync"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_BOOL).UnPack(var_84fa77833ef8)

		if err != nil {
			panic(err)
		}

		var_84fa77833ef8_mapped := val.(bool)

		s.Sync = var_84fa77833ef8_mapped
	}
	if properties["responds"] != nil {

		var_fb4dc0e96faf := properties["responds"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_BOOL).UnPack(var_fb4dc0e96faf)

		if err != nil {
			panic(err)
		}

		var_fb4dc0e96faf_mapped := val.(bool)

		s.Responds = var_fb4dc0e96faf_mapped
	}
	if properties["call"] != nil {

		var_dd8ec95b80ae := properties["call"]
		var mappedValue = ExtensionExternalCallMapperInstance.FromProperties(var_dd8ec95b80ae.GetStructValue().Fields)

		var_dd8ec95b80ae_mapped := *mappedValue

		s.Call = var_dd8ec95b80ae_mapped
	}
	if properties["annotations"] != nil {

		var_4bc2bd52a0ad := properties["annotations"]
		var_4bc2bd52a0ad_mapped := make(map[string]string)
		for k, v := range var_4bc2bd52a0ad.GetStructValue().Fields {

			var_2f8fd2056d6b := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_2f8fd2056d6b)

			if err != nil {
				panic(err)
			}

			var_2f8fd2056d6b_mapped := val.(string)

			var_4bc2bd52a0ad_mapped[k] = var_2f8fd2056d6b_mapped
		}

		s.Annotations = var_4bc2bd52a0ad_mapped
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

		var_26578160968c := properties["host"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_26578160968c)

		if err != nil {
			panic(err)
		}

		var_26578160968c_mapped := val.(string)

		s.Host = var_26578160968c_mapped
	}
	if properties["functionName"] != nil {

		var_24c10302c718 := properties["functionName"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_24c10302c718)

		if err != nil {
			panic(err)
		}

		var_24c10302c718_mapped := val.(string)

		s.FunctionName = var_24c10302c718_mapped
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

		var_a24f6d265763 := properties["uri"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_a24f6d265763)

		if err != nil {
			panic(err)
		}

		var_a24f6d265763_mapped := val.(string)

		s.Uri = var_a24f6d265763_mapped
	}
	if properties["method"] != nil {

		var_d4a095089f09 := properties["method"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_d4a095089f09)

		if err != nil {
			panic(err)
		}

		var_d4a095089f09_mapped := val.(string)

		s.Method = var_d4a095089f09_mapped
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

		var_89427ae23300 := properties["functionCall"]
		var mappedValue = ExtensionFunctionCallMapperInstance.FromProperties(var_89427ae23300.GetStructValue().Fields)

		var_89427ae23300_mapped := mappedValue

		s.FunctionCall = var_89427ae23300_mapped
	}
	if properties["httpCall"] != nil {

		var_dc77a0b063cd := properties["httpCall"]
		var mappedValue = ExtensionHttpCallMapperInstance.FromProperties(var_dc77a0b063cd.GetStructValue().Fields)

		var_dc77a0b063cd_mapped := mappedValue

		s.HttpCall = var_dc77a0b063cd_mapped
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

		var_4406f2ddfe04 := properties["actions"]
		var_4406f2ddfe04_mapped := []EventAction{}
		for _, v := range var_4406f2ddfe04.GetListValue().Values {

			var_e30626b13da6 := v
			var_e30626b13da6_mapped := (EventAction)(var_e30626b13da6.GetStringValue())

			var_4406f2ddfe04_mapped = append(var_4406f2ddfe04_mapped, var_e30626b13da6_mapped)
		}

		s.Actions = var_4406f2ddfe04_mapped
	}
	if properties["recordSelector"] != nil {

		var_9b7c0cc04dac := properties["recordSelector"]
		var mappedValue = ExtensionBooleanExpressionMapperInstance.FromProperties(var_9b7c0cc04dac.GetStructValue().Fields)

		var_9b7c0cc04dac_mapped := mappedValue

		s.RecordSelector = var_9b7c0cc04dac_mapped
	}
	if properties["namespaces"] != nil {

		var_9eebfb34b3c2 := properties["namespaces"]
		var_9eebfb34b3c2_mapped := []string{}
		for _, v := range var_9eebfb34b3c2.GetListValue().Values {

			var_363d998cce1d := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_363d998cce1d)

			if err != nil {
				panic(err)
			}

			var_363d998cce1d_mapped := val.(string)

			var_9eebfb34b3c2_mapped = append(var_9eebfb34b3c2_mapped, var_363d998cce1d_mapped)
		}

		s.Namespaces = var_9eebfb34b3c2_mapped
	}
	if properties["resources"] != nil {

		var_9c3b39ea8bf8 := properties["resources"]
		var_9c3b39ea8bf8_mapped := []string{}
		for _, v := range var_9c3b39ea8bf8.GetListValue().Values {

			var_447c7fa44940 := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_447c7fa44940)

			if err != nil {
				panic(err)
			}

			var_447c7fa44940_mapped := val.(string)

			var_9c3b39ea8bf8_mapped = append(var_9c3b39ea8bf8_mapped, var_447c7fa44940_mapped)
		}

		s.Resources = var_9c3b39ea8bf8_mapped
	}
	if properties["ids"] != nil {

		var_5f87af086673 := properties["ids"]
		var_5f87af086673_mapped := []string{}
		for _, v := range var_5f87af086673.GetListValue().Values {

			var_50ffe4b07682 := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_50ffe4b07682)

			if err != nil {
				panic(err)
			}

			var_50ffe4b07682_mapped := val.(string)

			var_5f87af086673_mapped = append(var_5f87af086673_mapped, var_50ffe4b07682_mapped)
		}

		s.Ids = var_5f87af086673_mapped
	}
	if properties["annotations"] != nil {

		var_547b5aea7b88 := properties["annotations"]
		var_547b5aea7b88_mapped := make(map[string]string)
		for k, v := range var_547b5aea7b88.GetStructValue().Fields {

			var_ecf831dd60a2 := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_ecf831dd60a2)

			if err != nil {
				panic(err)
			}

			var_ecf831dd60a2_mapped := val.(string)

			var_547b5aea7b88_mapped[k] = var_ecf831dd60a2_mapped
		}

		s.Annotations = var_547b5aea7b88_mapped
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

		var_a996c4514c15 := properties["query"]
		var mappedValue = ExtensionBooleanExpressionMapperInstance.FromProperties(var_a996c4514c15.GetStructValue().Fields)

		var_a996c4514c15_mapped := mappedValue

		s.Query = var_a996c4514c15_mapped
	}
	if properties["limit"] != nil {

		var_b1d736da3109 := properties["limit"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(var_b1d736da3109)

		if err != nil {
			panic(err)
		}

		var_b1d736da3109_mapped := new(int32)
		*var_b1d736da3109_mapped = val.(int32)

		s.Limit = var_b1d736da3109_mapped
	}
	if properties["offset"] != nil {

		var_80ed7cef3d7a := properties["offset"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(var_80ed7cef3d7a)

		if err != nil {
			panic(err)
		}

		var_80ed7cef3d7a_mapped := new(int32)
		*var_80ed7cef3d7a_mapped = val.(int32)

		s.Offset = var_80ed7cef3d7a_mapped
	}
	if properties["resolveReferences"] != nil {

		var_cceb2af6316d := properties["resolveReferences"]
		var_cceb2af6316d_mapped := []string{}
		for _, v := range var_cceb2af6316d.GetListValue().Values {

			var_f20020f4267b := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_f20020f4267b)

			if err != nil {
				panic(err)
			}

			var_f20020f4267b_mapped := val.(string)

			var_cceb2af6316d_mapped = append(var_cceb2af6316d_mapped, var_f20020f4267b_mapped)
		}

		s.ResolveReferences = var_cceb2af6316d_mapped
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

		var_e252531f854e := properties["id"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_UUID).UnPack(var_e252531f854e)

		if err != nil {
			panic(err)
		}

		var_e252531f854e_mapped := new(uuid.UUID)
		*var_e252531f854e_mapped = val.(uuid.UUID)

		s.Id = var_e252531f854e_mapped
	}
	if properties["action"] != nil {

		var_6d2a9505cd05 := properties["action"]
		var_6d2a9505cd05_mapped := (EventAction)(var_6d2a9505cd05.GetStringValue())

		s.Action = var_6d2a9505cd05_mapped
	}
	if properties["recordSearchParams"] != nil {

		var_800fc73f21ea := properties["recordSearchParams"]
		var mappedValue = ExtensionRecordSearchParamsMapperInstance.FromProperties(var_800fc73f21ea.GetStructValue().Fields)

		var_800fc73f21ea_mapped := mappedValue

		s.RecordSearchParams = var_800fc73f21ea_mapped
	}
	if properties["actionSummary"] != nil {

		var_a0b681285ce4 := properties["actionSummary"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_a0b681285ce4)

		if err != nil {
			panic(err)
		}

		var_a0b681285ce4_mapped := new(string)
		*var_a0b681285ce4_mapped = val.(string)

		s.ActionSummary = var_a0b681285ce4_mapped
	}
	if properties["actionDescription"] != nil {

		var_ad2e663b9f23 := properties["actionDescription"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_ad2e663b9f23)

		if err != nil {
			panic(err)
		}

		var_ad2e663b9f23_mapped := new(string)
		*var_ad2e663b9f23_mapped = val.(string)

		s.ActionDescription = var_ad2e663b9f23_mapped
	}
	if properties["resource"] != nil {

		var_56a602100b8b := properties["resource"]
		var_56a602100b8b_mapped := ResourceMapperInstance.FromProperties(var_56a602100b8b.GetStructValue().Fields)

		s.Resource = var_56a602100b8b_mapped
	}
	if properties["records"] != nil {

		var_0788202aa797 := properties["records"]
		var_0788202aa797_mapped := []*Record{}
		for _, v := range var_0788202aa797.GetListValue().Values {

			var_65342a495585 := v
			var_65342a495585_mapped := RecordMapperInstance.FromProperties(var_65342a495585.GetStructValue().Fields)

			var_0788202aa797_mapped = append(var_0788202aa797_mapped, var_65342a495585_mapped)
		}

		s.Records = var_0788202aa797_mapped
	}
	if properties["ids"] != nil {

		var_c7f7b76dffc8 := properties["ids"]
		var_c7f7b76dffc8_mapped := []string{}
		for _, v := range var_c7f7b76dffc8.GetListValue().Values {

			var_37d7a95481a6 := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_37d7a95481a6)

			if err != nil {
				panic(err)
			}

			var_37d7a95481a6_mapped := val.(string)

			var_c7f7b76dffc8_mapped = append(var_c7f7b76dffc8_mapped, var_37d7a95481a6_mapped)
		}

		s.Ids = var_c7f7b76dffc8_mapped
	}
	if properties["finalizes"] != nil {

		var_c3c1d1bad24e := properties["finalizes"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_BOOL).UnPack(var_c3c1d1bad24e)

		if err != nil {
			panic(err)
		}

		var_c3c1d1bad24e_mapped := new(bool)
		*var_c3c1d1bad24e_mapped = val.(bool)

		s.Finalizes = var_c3c1d1bad24e_mapped
	}
	if properties["sync"] != nil {

		var_c9a7569bb13b := properties["sync"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_BOOL).UnPack(var_c9a7569bb13b)

		if err != nil {
			panic(err)
		}

		var_c9a7569bb13b_mapped := new(bool)
		*var_c9a7569bb13b_mapped = val.(bool)

		s.Sync = var_c9a7569bb13b_mapped
	}
	if properties["time"] != nil {

		var_b0ab953d4b06 := properties["time"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_b0ab953d4b06)

		if err != nil {
			panic(err)
		}

		var_b0ab953d4b06_mapped := new(time.Time)
		*var_b0ab953d4b06_mapped = val.(time.Time)

		s.Time = var_b0ab953d4b06_mapped
	}
	if properties["annotations"] != nil {

		var_c8bdb3092cde := properties["annotations"]
		var_c8bdb3092cde_mapped := make(map[string]string)
		for k, v := range var_c8bdb3092cde.GetStructValue().Fields {

			var_eb05b978ab03 := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_eb05b978ab03)

			if err != nil {
				panic(err)
			}

			var_eb05b978ab03_mapped := val.(string)

			var_c8bdb3092cde_mapped[k] = var_eb05b978ab03_mapped
		}

		s.Annotations = var_c8bdb3092cde_mapped
	}
	return s
}
