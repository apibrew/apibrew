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

		var_249e63739473 := properties["id"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_UUID).UnPack(var_249e63739473)

		if err != nil {
			panic(err)
		}

		var_249e63739473_mapped := new(uuid.UUID)
		*var_249e63739473_mapped = val.(uuid.UUID)

		s.Id = var_249e63739473_mapped
	}
	if properties["version"] != nil {

		var_d3c97abc47e9 := properties["version"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(var_d3c97abc47e9)

		if err != nil {
			panic(err)
		}

		var_d3c97abc47e9_mapped := new(int32)
		*var_d3c97abc47e9_mapped = val.(int32)

		s.Version = var_d3c97abc47e9_mapped
	}
	if properties["createdBy"] != nil {

		var_9bd7270c09cd := properties["createdBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_9bd7270c09cd)

		if err != nil {
			panic(err)
		}

		var_9bd7270c09cd_mapped := new(string)
		*var_9bd7270c09cd_mapped = val.(string)

		s.CreatedBy = var_9bd7270c09cd_mapped
	}
	if properties["updatedBy"] != nil {

		var_3aed35cca6c3 := properties["updatedBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_3aed35cca6c3)

		if err != nil {
			panic(err)
		}

		var_3aed35cca6c3_mapped := new(string)
		*var_3aed35cca6c3_mapped = val.(string)

		s.UpdatedBy = var_3aed35cca6c3_mapped
	}
	if properties["createdOn"] != nil {

		var_33c60cc798dd := properties["createdOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_33c60cc798dd)

		if err != nil {
			panic(err)
		}

		var_33c60cc798dd_mapped := new(time.Time)
		*var_33c60cc798dd_mapped = val.(time.Time)

		s.CreatedOn = var_33c60cc798dd_mapped
	}
	if properties["updatedOn"] != nil {

		var_9cd6585172e2 := properties["updatedOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_9cd6585172e2)

		if err != nil {
			panic(err)
		}

		var_9cd6585172e2_mapped := new(time.Time)
		*var_9cd6585172e2_mapped = val.(time.Time)

		s.UpdatedOn = var_9cd6585172e2_mapped
	}
	if properties["name"] != nil {

		var_a82773b02576 := properties["name"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_a82773b02576)

		if err != nil {
			panic(err)
		}

		var_a82773b02576_mapped := val.(string)

		s.Name = var_a82773b02576_mapped
	}
	if properties["description"] != nil {

		var_fbdb5ad10932 := properties["description"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_fbdb5ad10932)

		if err != nil {
			panic(err)
		}

		var_fbdb5ad10932_mapped := new(string)
		*var_fbdb5ad10932_mapped = val.(string)

		s.Description = var_fbdb5ad10932_mapped
	}
	if properties["selector"] != nil {

		var_34922d1ca81d := properties["selector"]
		var mappedValue = ExtensionEventSelectorMapperInstance.FromProperties(var_34922d1ca81d.GetStructValue().Fields)

		var_34922d1ca81d_mapped := mappedValue

		s.Selector = var_34922d1ca81d_mapped
	}
	if properties["order"] != nil {

		var_4e1dd55536ab := properties["order"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(var_4e1dd55536ab)

		if err != nil {
			panic(err)
		}

		var_4e1dd55536ab_mapped := val.(int32)

		s.Order = var_4e1dd55536ab_mapped
	}
	if properties["finalizes"] != nil {

		var_bf25c581dc43 := properties["finalizes"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_BOOL).UnPack(var_bf25c581dc43)

		if err != nil {
			panic(err)
		}

		var_bf25c581dc43_mapped := val.(bool)

		s.Finalizes = var_bf25c581dc43_mapped
	}
	if properties["sync"] != nil {

		var_26159b241ba8 := properties["sync"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_BOOL).UnPack(var_26159b241ba8)

		if err != nil {
			panic(err)
		}

		var_26159b241ba8_mapped := val.(bool)

		s.Sync = var_26159b241ba8_mapped
	}
	if properties["responds"] != nil {

		var_7e3d842a801e := properties["responds"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_BOOL).UnPack(var_7e3d842a801e)

		if err != nil {
			panic(err)
		}

		var_7e3d842a801e_mapped := val.(bool)

		s.Responds = var_7e3d842a801e_mapped
	}
	if properties["call"] != nil {

		var_ad81fb930b31 := properties["call"]
		var mappedValue = ExtensionExternalCallMapperInstance.FromProperties(var_ad81fb930b31.GetStructValue().Fields)

		var_ad81fb930b31_mapped := *mappedValue

		s.Call = var_ad81fb930b31_mapped
	}
	if properties["annotations"] != nil {

		var_3c39c7ef530d := properties["annotations"]
		var_3c39c7ef530d_mapped := make(map[string]string)
		for k, v := range var_3c39c7ef530d.GetStructValue().Fields {

			var_e8a13aac490b := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_e8a13aac490b)

			if err != nil {
				panic(err)
			}

			var_e8a13aac490b_mapped := val.(string)

			var_3c39c7ef530d_mapped[k] = var_e8a13aac490b_mapped
		}

		s.Annotations = var_3c39c7ef530d_mapped
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

		var_6aa383783f49 := properties["host"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_6aa383783f49)

		if err != nil {
			panic(err)
		}

		var_6aa383783f49_mapped := val.(string)

		s.Host = var_6aa383783f49_mapped
	}
	if properties["functionName"] != nil {

		var_21f9af1777af := properties["functionName"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_21f9af1777af)

		if err != nil {
			panic(err)
		}

		var_21f9af1777af_mapped := val.(string)

		s.FunctionName = var_21f9af1777af_mapped
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

		var_c961b7ec7a13 := properties["uri"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_c961b7ec7a13)

		if err != nil {
			panic(err)
		}

		var_c961b7ec7a13_mapped := val.(string)

		s.Uri = var_c961b7ec7a13_mapped
	}
	if properties["method"] != nil {

		var_5c31914799da := properties["method"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_5c31914799da)

		if err != nil {
			panic(err)
		}

		var_5c31914799da_mapped := val.(string)

		s.Method = var_5c31914799da_mapped
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

		var_41efaafc1256 := properties["functionCall"]
		var mappedValue = ExtensionFunctionCallMapperInstance.FromProperties(var_41efaafc1256.GetStructValue().Fields)

		var_41efaafc1256_mapped := mappedValue

		s.FunctionCall = var_41efaafc1256_mapped
	}
	if properties["httpCall"] != nil {

		var_3555b465098a := properties["httpCall"]
		var mappedValue = ExtensionHttpCallMapperInstance.FromProperties(var_3555b465098a.GetStructValue().Fields)

		var_3555b465098a_mapped := mappedValue

		s.HttpCall = var_3555b465098a_mapped
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

		var_5ee706909cf8 := properties["actions"]
		var_5ee706909cf8_mapped := []EventAction{}
		for _, v := range var_5ee706909cf8.GetListValue().Values {

			var_eb377ba57c35 := v
			var_eb377ba57c35_mapped := (EventAction)(var_eb377ba57c35.GetStringValue())

			var_5ee706909cf8_mapped = append(var_5ee706909cf8_mapped, var_eb377ba57c35_mapped)
		}

		s.Actions = var_5ee706909cf8_mapped
	}
	if properties["recordSelector"] != nil {

		var_722982175248 := properties["recordSelector"]
		var mappedValue = ExtensionBooleanExpressionMapperInstance.FromProperties(var_722982175248.GetStructValue().Fields)

		var_722982175248_mapped := mappedValue

		s.RecordSelector = var_722982175248_mapped
	}
	if properties["namespaces"] != nil {

		var_ae597a17a453 := properties["namespaces"]
		var_ae597a17a453_mapped := []string{}
		for _, v := range var_ae597a17a453.GetListValue().Values {

			var_c9086b65f719 := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_c9086b65f719)

			if err != nil {
				panic(err)
			}

			var_c9086b65f719_mapped := val.(string)

			var_ae597a17a453_mapped = append(var_ae597a17a453_mapped, var_c9086b65f719_mapped)
		}

		s.Namespaces = var_ae597a17a453_mapped
	}
	if properties["resources"] != nil {

		var_b1ebb7455aad := properties["resources"]
		var_b1ebb7455aad_mapped := []string{}
		for _, v := range var_b1ebb7455aad.GetListValue().Values {

			var_cef401f1c815 := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_cef401f1c815)

			if err != nil {
				panic(err)
			}

			var_cef401f1c815_mapped := val.(string)

			var_b1ebb7455aad_mapped = append(var_b1ebb7455aad_mapped, var_cef401f1c815_mapped)
		}

		s.Resources = var_b1ebb7455aad_mapped
	}
	if properties["ids"] != nil {

		var_79aaeb44f486 := properties["ids"]
		var_79aaeb44f486_mapped := []string{}
		for _, v := range var_79aaeb44f486.GetListValue().Values {

			var_6002a5fa9e3b := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_6002a5fa9e3b)

			if err != nil {
				panic(err)
			}

			var_6002a5fa9e3b_mapped := val.(string)

			var_79aaeb44f486_mapped = append(var_79aaeb44f486_mapped, var_6002a5fa9e3b_mapped)
		}

		s.Ids = var_79aaeb44f486_mapped
	}
	if properties["annotations"] != nil {

		var_9769731d08c1 := properties["annotations"]
		var_9769731d08c1_mapped := make(map[string]string)
		for k, v := range var_9769731d08c1.GetStructValue().Fields {

			var_c612ccd606f8 := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_c612ccd606f8)

			if err != nil {
				panic(err)
			}

			var_c612ccd606f8_mapped := val.(string)

			var_9769731d08c1_mapped[k] = var_c612ccd606f8_mapped
		}

		s.Annotations = var_9769731d08c1_mapped
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

		var_0c6ec4fe50a7 := properties["query"]
		var mappedValue = ExtensionBooleanExpressionMapperInstance.FromProperties(var_0c6ec4fe50a7.GetStructValue().Fields)

		var_0c6ec4fe50a7_mapped := mappedValue

		s.Query = var_0c6ec4fe50a7_mapped
	}
	if properties["limit"] != nil {

		var_212723d26349 := properties["limit"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(var_212723d26349)

		if err != nil {
			panic(err)
		}

		var_212723d26349_mapped := new(int32)
		*var_212723d26349_mapped = val.(int32)

		s.Limit = var_212723d26349_mapped
	}
	if properties["offset"] != nil {

		var_19adfaf7fe30 := properties["offset"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(var_19adfaf7fe30)

		if err != nil {
			panic(err)
		}

		var_19adfaf7fe30_mapped := new(int32)
		*var_19adfaf7fe30_mapped = val.(int32)

		s.Offset = var_19adfaf7fe30_mapped
	}
	if properties["resolveReferences"] != nil {

		var_483a1a2ed24e := properties["resolveReferences"]
		var_483a1a2ed24e_mapped := []string{}
		for _, v := range var_483a1a2ed24e.GetListValue().Values {

			var_191269c16ec7 := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_191269c16ec7)

			if err != nil {
				panic(err)
			}

			var_191269c16ec7_mapped := val.(string)

			var_483a1a2ed24e_mapped = append(var_483a1a2ed24e_mapped, var_191269c16ec7_mapped)
		}

		s.ResolveReferences = var_483a1a2ed24e_mapped
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

		var_bee5fe2a2e26 := properties["id"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_UUID).UnPack(var_bee5fe2a2e26)

		if err != nil {
			panic(err)
		}

		var_bee5fe2a2e26_mapped := new(uuid.UUID)
		*var_bee5fe2a2e26_mapped = val.(uuid.UUID)

		s.Id = var_bee5fe2a2e26_mapped
	}
	if properties["action"] != nil {

		var_f7d94bc9dd14 := properties["action"]
		var_f7d94bc9dd14_mapped := (EventAction)(var_f7d94bc9dd14.GetStringValue())

		s.Action = var_f7d94bc9dd14_mapped
	}
	if properties["recordSearchParams"] != nil {

		var_01dfb7818682 := properties["recordSearchParams"]
		var mappedValue = ExtensionRecordSearchParamsMapperInstance.FromProperties(var_01dfb7818682.GetStructValue().Fields)

		var_01dfb7818682_mapped := mappedValue

		s.RecordSearchParams = var_01dfb7818682_mapped
	}
	if properties["actionSummary"] != nil {

		var_a706bd836346 := properties["actionSummary"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_a706bd836346)

		if err != nil {
			panic(err)
		}

		var_a706bd836346_mapped := new(string)
		*var_a706bd836346_mapped = val.(string)

		s.ActionSummary = var_a706bd836346_mapped
	}
	if properties["actionDescription"] != nil {

		var_7c5356de04fb := properties["actionDescription"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_7c5356de04fb)

		if err != nil {
			panic(err)
		}

		var_7c5356de04fb_mapped := new(string)
		*var_7c5356de04fb_mapped = val.(string)

		s.ActionDescription = var_7c5356de04fb_mapped
	}
	if properties["resource"] != nil {

		var_903d63525798 := properties["resource"]
		var_903d63525798_mapped := ResourceMapperInstance.FromProperties(var_903d63525798.GetStructValue().Fields)

		s.Resource = var_903d63525798_mapped
	}
	if properties["records"] != nil {

		var_afd2465d70fe := properties["records"]
		var_afd2465d70fe_mapped := []*Record{}
		for _, v := range var_afd2465d70fe.GetListValue().Values {

			var_9a7bae280c04 := v
			var_9a7bae280c04_mapped := RecordMapperInstance.FromProperties(var_9a7bae280c04.GetStructValue().Fields)

			var_afd2465d70fe_mapped = append(var_afd2465d70fe_mapped, var_9a7bae280c04_mapped)
		}

		s.Records = var_afd2465d70fe_mapped
	}
	if properties["ids"] != nil {

		var_8b4082d7a29d := properties["ids"]
		var_8b4082d7a29d_mapped := []string{}
		for _, v := range var_8b4082d7a29d.GetListValue().Values {

			var_86b7d27b314b := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_86b7d27b314b)

			if err != nil {
				panic(err)
			}

			var_86b7d27b314b_mapped := val.(string)

			var_8b4082d7a29d_mapped = append(var_8b4082d7a29d_mapped, var_86b7d27b314b_mapped)
		}

		s.Ids = var_8b4082d7a29d_mapped
	}
	if properties["finalizes"] != nil {

		var_7345ac0229dc := properties["finalizes"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_BOOL).UnPack(var_7345ac0229dc)

		if err != nil {
			panic(err)
		}

		var_7345ac0229dc_mapped := new(bool)
		*var_7345ac0229dc_mapped = val.(bool)

		s.Finalizes = var_7345ac0229dc_mapped
	}
	if properties["sync"] != nil {

		var_8b23bcdc7e1b := properties["sync"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_BOOL).UnPack(var_8b23bcdc7e1b)

		if err != nil {
			panic(err)
		}

		var_8b23bcdc7e1b_mapped := new(bool)
		*var_8b23bcdc7e1b_mapped = val.(bool)

		s.Sync = var_8b23bcdc7e1b_mapped
	}
	if properties["time"] != nil {

		var_e434be6db32a := properties["time"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_e434be6db32a)

		if err != nil {
			panic(err)
		}

		var_e434be6db32a_mapped := new(time.Time)
		*var_e434be6db32a_mapped = val.(time.Time)

		s.Time = var_e434be6db32a_mapped
	}
	if properties["annotations"] != nil {

		var_e60df16d7159 := properties["annotations"]
		var_e60df16d7159_mapped := make(map[string]string)
		for k, v := range var_e60df16d7159.GetStructValue().Fields {

			var_50ee4091b76b := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_50ee4091b76b)

			if err != nil {
				panic(err)
			}

			var_50ee4091b76b_mapped := val.(string)

			var_e60df16d7159_mapped[k] = var_50ee4091b76b_mapped
		}

		s.Annotations = var_e60df16d7159_mapped
	}
	return s
}
