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

	var_a6a517e81d08 := extension.Id

	if var_a6a517e81d08 != nil {
		var var_a6a517e81d08_mapped *structpb.Value

		var var_a6a517e81d08_err error
		var_a6a517e81d08_mapped, var_a6a517e81d08_err = types.ByResourcePropertyType(model.ResourceProperty_UUID).Pack(*var_a6a517e81d08)
		if var_a6a517e81d08_err != nil {
			panic(var_a6a517e81d08_err)
		}
		properties["id"] = var_a6a517e81d08_mapped
	}

	var_0c1c80ed45cd := extension.Version

	var var_0c1c80ed45cd_mapped *structpb.Value

	var var_0c1c80ed45cd_err error
	var_0c1c80ed45cd_mapped, var_0c1c80ed45cd_err = types.ByResourcePropertyType(model.ResourceProperty_INT32).Pack(var_0c1c80ed45cd)
	if var_0c1c80ed45cd_err != nil {
		panic(var_0c1c80ed45cd_err)
	}
	properties["version"] = var_0c1c80ed45cd_mapped

	var_3574182462ec := extension.CreatedBy

	if var_3574182462ec != nil {
		var var_3574182462ec_mapped *structpb.Value

		var var_3574182462ec_err error
		var_3574182462ec_mapped, var_3574182462ec_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_3574182462ec)
		if var_3574182462ec_err != nil {
			panic(var_3574182462ec_err)
		}
		properties["createdBy"] = var_3574182462ec_mapped
	}

	var_05107cede04f := extension.UpdatedBy

	if var_05107cede04f != nil {
		var var_05107cede04f_mapped *structpb.Value

		var var_05107cede04f_err error
		var_05107cede04f_mapped, var_05107cede04f_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_05107cede04f)
		if var_05107cede04f_err != nil {
			panic(var_05107cede04f_err)
		}
		properties["updatedBy"] = var_05107cede04f_mapped
	}

	var_11c04ca8b239 := extension.CreatedOn

	if var_11c04ca8b239 != nil {
		var var_11c04ca8b239_mapped *structpb.Value

		var var_11c04ca8b239_err error
		var_11c04ca8b239_mapped, var_11c04ca8b239_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_11c04ca8b239)
		if var_11c04ca8b239_err != nil {
			panic(var_11c04ca8b239_err)
		}
		properties["createdOn"] = var_11c04ca8b239_mapped
	}

	var_2727a2437e55 := extension.UpdatedOn

	if var_2727a2437e55 != nil {
		var var_2727a2437e55_mapped *structpb.Value

		var var_2727a2437e55_err error
		var_2727a2437e55_mapped, var_2727a2437e55_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_2727a2437e55)
		if var_2727a2437e55_err != nil {
			panic(var_2727a2437e55_err)
		}
		properties["updatedOn"] = var_2727a2437e55_mapped
	}

	var_c9696ef9c9cd := extension.Name

	var var_c9696ef9c9cd_mapped *structpb.Value

	var var_c9696ef9c9cd_err error
	var_c9696ef9c9cd_mapped, var_c9696ef9c9cd_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_c9696ef9c9cd)
	if var_c9696ef9c9cd_err != nil {
		panic(var_c9696ef9c9cd_err)
	}
	properties["name"] = var_c9696ef9c9cd_mapped

	var_951886abcf9d := extension.Description

	if var_951886abcf9d != nil {
		var var_951886abcf9d_mapped *structpb.Value

		var var_951886abcf9d_err error
		var_951886abcf9d_mapped, var_951886abcf9d_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_951886abcf9d)
		if var_951886abcf9d_err != nil {
			panic(var_951886abcf9d_err)
		}
		properties["description"] = var_951886abcf9d_mapped
	}

	var_0b89dbb59e32 := extension.Selector

	if var_0b89dbb59e32 != nil {
		var var_0b89dbb59e32_mapped *structpb.Value

		var_0b89dbb59e32_mapped = structpb.NewStructValue(&structpb.Struct{Fields: ExtensionEventSelectorMapperInstance.ToProperties(var_0b89dbb59e32)})
		properties["selector"] = var_0b89dbb59e32_mapped
	}

	var_0112030cd09d := extension.Order

	var var_0112030cd09d_mapped *structpb.Value

	var var_0112030cd09d_err error
	var_0112030cd09d_mapped, var_0112030cd09d_err = types.ByResourcePropertyType(model.ResourceProperty_INT32).Pack(var_0112030cd09d)
	if var_0112030cd09d_err != nil {
		panic(var_0112030cd09d_err)
	}
	properties["order"] = var_0112030cd09d_mapped

	var_c79e77c5496b := extension.Finalizes

	var var_c79e77c5496b_mapped *structpb.Value

	var var_c79e77c5496b_err error
	var_c79e77c5496b_mapped, var_c79e77c5496b_err = types.ByResourcePropertyType(model.ResourceProperty_BOOL).Pack(var_c79e77c5496b)
	if var_c79e77c5496b_err != nil {
		panic(var_c79e77c5496b_err)
	}
	properties["finalizes"] = var_c79e77c5496b_mapped

	var_fbe7de9f588c := extension.Sync

	var var_fbe7de9f588c_mapped *structpb.Value

	var var_fbe7de9f588c_err error
	var_fbe7de9f588c_mapped, var_fbe7de9f588c_err = types.ByResourcePropertyType(model.ResourceProperty_BOOL).Pack(var_fbe7de9f588c)
	if var_fbe7de9f588c_err != nil {
		panic(var_fbe7de9f588c_err)
	}
	properties["sync"] = var_fbe7de9f588c_mapped

	var_45cd579323b9 := extension.Responds

	var var_45cd579323b9_mapped *structpb.Value

	var var_45cd579323b9_err error
	var_45cd579323b9_mapped, var_45cd579323b9_err = types.ByResourcePropertyType(model.ResourceProperty_BOOL).Pack(var_45cd579323b9)
	if var_45cd579323b9_err != nil {
		panic(var_45cd579323b9_err)
	}
	properties["responds"] = var_45cd579323b9_mapped

	var_fb201025748d := extension.Call

	var var_fb201025748d_mapped *structpb.Value

	var_fb201025748d_mapped = structpb.NewStructValue(&structpb.Struct{Fields: ExtensionExternalCallMapperInstance.ToProperties(&var_fb201025748d)})
	properties["call"] = var_fb201025748d_mapped

	var_72fe2ab6c6bb := extension.Annotations

	if var_72fe2ab6c6bb != nil {
		var var_72fe2ab6c6bb_mapped *structpb.Value

		var var_72fe2ab6c6bb_st *structpb.Struct = new(structpb.Struct)
		var_72fe2ab6c6bb_st.Fields = make(map[string]*structpb.Value)
		for key, value := range var_72fe2ab6c6bb {

			var_607fc478e80a := value
			var var_607fc478e80a_mapped *structpb.Value

			var var_607fc478e80a_err error
			var_607fc478e80a_mapped, var_607fc478e80a_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_607fc478e80a)
			if var_607fc478e80a_err != nil {
				panic(var_607fc478e80a_err)
			}

			var_72fe2ab6c6bb_st.Fields[key] = var_607fc478e80a_mapped
		}
		var_72fe2ab6c6bb_mapped = structpb.NewStructValue(var_72fe2ab6c6bb_st)
		properties["annotations"] = var_72fe2ab6c6bb_mapped
	}
	return properties
}

func (m *ExtensionMapper) FromProperties(properties map[string]*structpb.Value) *Extension {
	var s = m.New()
	if properties["id"] != nil {

		var_ccb50a17de68 := properties["id"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_UUID).UnPack(var_ccb50a17de68)

		if err != nil {
			panic(err)
		}

		var_ccb50a17de68_mapped := new(uuid.UUID)
		*var_ccb50a17de68_mapped = val.(uuid.UUID)

		s.Id = var_ccb50a17de68_mapped
	}
	if properties["version"] != nil {

		var_a8a1fc6a40d2 := properties["version"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(var_a8a1fc6a40d2)

		if err != nil {
			panic(err)
		}

		var_a8a1fc6a40d2_mapped := val.(int32)

		s.Version = var_a8a1fc6a40d2_mapped
	}
	if properties["createdBy"] != nil {

		var_6981373896d9 := properties["createdBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_6981373896d9)

		if err != nil {
			panic(err)
		}

		var_6981373896d9_mapped := new(string)
		*var_6981373896d9_mapped = val.(string)

		s.CreatedBy = var_6981373896d9_mapped
	}
	if properties["updatedBy"] != nil {

		var_b5c81bc6817d := properties["updatedBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_b5c81bc6817d)

		if err != nil {
			panic(err)
		}

		var_b5c81bc6817d_mapped := new(string)
		*var_b5c81bc6817d_mapped = val.(string)

		s.UpdatedBy = var_b5c81bc6817d_mapped
	}
	if properties["createdOn"] != nil {

		var_b0c18a2778fb := properties["createdOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_b0c18a2778fb)

		if err != nil {
			panic(err)
		}

		var_b0c18a2778fb_mapped := new(time.Time)
		*var_b0c18a2778fb_mapped = val.(time.Time)

		s.CreatedOn = var_b0c18a2778fb_mapped
	}
	if properties["updatedOn"] != nil {

		var_4b03ce1902ac := properties["updatedOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_4b03ce1902ac)

		if err != nil {
			panic(err)
		}

		var_4b03ce1902ac_mapped := new(time.Time)
		*var_4b03ce1902ac_mapped = val.(time.Time)

		s.UpdatedOn = var_4b03ce1902ac_mapped
	}
	if properties["name"] != nil {

		var_2c9577b3a768 := properties["name"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_2c9577b3a768)

		if err != nil {
			panic(err)
		}

		var_2c9577b3a768_mapped := val.(string)

		s.Name = var_2c9577b3a768_mapped
	}
	if properties["description"] != nil {

		var_6f872fa066cc := properties["description"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_6f872fa066cc)

		if err != nil {
			panic(err)
		}

		var_6f872fa066cc_mapped := new(string)
		*var_6f872fa066cc_mapped = val.(string)

		s.Description = var_6f872fa066cc_mapped
	}
	if properties["selector"] != nil {

		var_0e032387146d := properties["selector"]
		var mappedValue = ExtensionEventSelectorMapperInstance.FromProperties(var_0e032387146d.GetStructValue().Fields)

		var_0e032387146d_mapped := mappedValue

		s.Selector = var_0e032387146d_mapped
	}
	if properties["order"] != nil {

		var_ad6dc532b868 := properties["order"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(var_ad6dc532b868)

		if err != nil {
			panic(err)
		}

		var_ad6dc532b868_mapped := val.(int32)

		s.Order = var_ad6dc532b868_mapped
	}
	if properties["finalizes"] != nil {

		var_9e65b8e34c06 := properties["finalizes"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_BOOL).UnPack(var_9e65b8e34c06)

		if err != nil {
			panic(err)
		}

		var_9e65b8e34c06_mapped := val.(bool)

		s.Finalizes = var_9e65b8e34c06_mapped
	}
	if properties["sync"] != nil {

		var_521a0d261f95 := properties["sync"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_BOOL).UnPack(var_521a0d261f95)

		if err != nil {
			panic(err)
		}

		var_521a0d261f95_mapped := val.(bool)

		s.Sync = var_521a0d261f95_mapped
	}
	if properties["responds"] != nil {

		var_4b47bc6d4c21 := properties["responds"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_BOOL).UnPack(var_4b47bc6d4c21)

		if err != nil {
			panic(err)
		}

		var_4b47bc6d4c21_mapped := val.(bool)

		s.Responds = var_4b47bc6d4c21_mapped
	}
	if properties["call"] != nil {

		var_a2db1b043272 := properties["call"]
		var mappedValue = ExtensionExternalCallMapperInstance.FromProperties(var_a2db1b043272.GetStructValue().Fields)

		var_a2db1b043272_mapped := *mappedValue

		s.Call = var_a2db1b043272_mapped
	}
	if properties["annotations"] != nil {

		var_784cf8eedc96 := properties["annotations"]
		var_784cf8eedc96_mapped := make(map[string]string)
		for k, v := range var_784cf8eedc96.GetStructValue().Fields {

			var_decea02b27bf := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_decea02b27bf)

			if err != nil {
				panic(err)
			}

			var_decea02b27bf_mapped := val.(string)

			var_784cf8eedc96_mapped[k] = var_decea02b27bf_mapped
		}

		s.Annotations = var_784cf8eedc96_mapped
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

	var_863d50b63f0d := extensionFunctionCall.Host

	var var_863d50b63f0d_mapped *structpb.Value

	var var_863d50b63f0d_err error
	var_863d50b63f0d_mapped, var_863d50b63f0d_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_863d50b63f0d)
	if var_863d50b63f0d_err != nil {
		panic(var_863d50b63f0d_err)
	}
	properties["host"] = var_863d50b63f0d_mapped

	var_aad528e4485a := extensionFunctionCall.FunctionName

	var var_aad528e4485a_mapped *structpb.Value

	var var_aad528e4485a_err error
	var_aad528e4485a_mapped, var_aad528e4485a_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_aad528e4485a)
	if var_aad528e4485a_err != nil {
		panic(var_aad528e4485a_err)
	}
	properties["functionName"] = var_aad528e4485a_mapped
	return properties
}

func (m *ExtensionFunctionCallMapper) FromProperties(properties map[string]*structpb.Value) *ExtensionFunctionCall {
	var s = m.New()
	if properties["host"] != nil {

		var_578256a55196 := properties["host"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_578256a55196)

		if err != nil {
			panic(err)
		}

		var_578256a55196_mapped := val.(string)

		s.Host = var_578256a55196_mapped
	}
	if properties["functionName"] != nil {

		var_96225bf8fdc9 := properties["functionName"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_96225bf8fdc9)

		if err != nil {
			panic(err)
		}

		var_96225bf8fdc9_mapped := val.(string)

		s.FunctionName = var_96225bf8fdc9_mapped
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

	var_f129fdd3f7e1 := extensionHttpCall.Uri

	var var_f129fdd3f7e1_mapped *structpb.Value

	var var_f129fdd3f7e1_err error
	var_f129fdd3f7e1_mapped, var_f129fdd3f7e1_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_f129fdd3f7e1)
	if var_f129fdd3f7e1_err != nil {
		panic(var_f129fdd3f7e1_err)
	}
	properties["uri"] = var_f129fdd3f7e1_mapped

	var_e240e725d730 := extensionHttpCall.Method

	var var_e240e725d730_mapped *structpb.Value

	var var_e240e725d730_err error
	var_e240e725d730_mapped, var_e240e725d730_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_e240e725d730)
	if var_e240e725d730_err != nil {
		panic(var_e240e725d730_err)
	}
	properties["method"] = var_e240e725d730_mapped
	return properties
}

func (m *ExtensionHttpCallMapper) FromProperties(properties map[string]*structpb.Value) *ExtensionHttpCall {
	var s = m.New()
	if properties["uri"] != nil {

		var_bc733383aa9b := properties["uri"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_bc733383aa9b)

		if err != nil {
			panic(err)
		}

		var_bc733383aa9b_mapped := val.(string)

		s.Uri = var_bc733383aa9b_mapped
	}
	if properties["method"] != nil {

		var_21d9d68599aa := properties["method"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_21d9d68599aa)

		if err != nil {
			panic(err)
		}

		var_21d9d68599aa_mapped := val.(string)

		s.Method = var_21d9d68599aa_mapped
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

	var_c65c954a514f := extensionExternalCall.FunctionCall

	if var_c65c954a514f != nil {
		var var_c65c954a514f_mapped *structpb.Value

		var_c65c954a514f_mapped = structpb.NewStructValue(&structpb.Struct{Fields: ExtensionFunctionCallMapperInstance.ToProperties(var_c65c954a514f)})
		properties["functionCall"] = var_c65c954a514f_mapped
	}

	var_4f0a7cc65794 := extensionExternalCall.HttpCall

	if var_4f0a7cc65794 != nil {
		var var_4f0a7cc65794_mapped *structpb.Value

		var_4f0a7cc65794_mapped = structpb.NewStructValue(&structpb.Struct{Fields: ExtensionHttpCallMapperInstance.ToProperties(var_4f0a7cc65794)})
		properties["httpCall"] = var_4f0a7cc65794_mapped
	}
	return properties
}

func (m *ExtensionExternalCallMapper) FromProperties(properties map[string]*structpb.Value) *ExtensionExternalCall {
	var s = m.New()
	if properties["functionCall"] != nil {

		var_e12576f205f8 := properties["functionCall"]
		var mappedValue = ExtensionFunctionCallMapperInstance.FromProperties(var_e12576f205f8.GetStructValue().Fields)

		var_e12576f205f8_mapped := mappedValue

		s.FunctionCall = var_e12576f205f8_mapped
	}
	if properties["httpCall"] != nil {

		var_683f67df088e := properties["httpCall"]
		var mappedValue = ExtensionHttpCallMapperInstance.FromProperties(var_683f67df088e.GetStructValue().Fields)

		var_683f67df088e_mapped := mappedValue

		s.HttpCall = var_683f67df088e_mapped
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

	var_562a3e8e0e73 := extensionEventSelector.Actions

	if var_562a3e8e0e73 != nil {
		var var_562a3e8e0e73_mapped *structpb.Value

		var var_562a3e8e0e73_l []*structpb.Value
		for _, value := range var_562a3e8e0e73 {

			var_64044a895634 := value
			var var_64044a895634_mapped *structpb.Value

			var var_64044a895634_err error
			var_64044a895634_mapped, var_64044a895634_err = types.ByResourcePropertyType(model.ResourceProperty_ENUM).Pack(string(var_64044a895634))
			if var_64044a895634_err != nil {
				panic(var_64044a895634_err)
			}

			var_562a3e8e0e73_l = append(var_562a3e8e0e73_l, var_64044a895634_mapped)
		}
		var_562a3e8e0e73_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_562a3e8e0e73_l})
		properties["actions"] = var_562a3e8e0e73_mapped
	}

	var_a6bfe8fe4305 := extensionEventSelector.RecordSelector

	if var_a6bfe8fe4305 != nil {
		var var_a6bfe8fe4305_mapped *structpb.Value

		var_a6bfe8fe4305_mapped = structpb.NewStructValue(&structpb.Struct{Fields: ExtensionBooleanExpressionMapperInstance.ToProperties(var_a6bfe8fe4305)})
		properties["recordSelector"] = var_a6bfe8fe4305_mapped
	}

	var_4de57fae9c91 := extensionEventSelector.Namespaces

	if var_4de57fae9c91 != nil {
		var var_4de57fae9c91_mapped *structpb.Value

		var var_4de57fae9c91_l []*structpb.Value
		for _, value := range var_4de57fae9c91 {

			var_600227c4c288 := value
			var var_600227c4c288_mapped *structpb.Value

			var var_600227c4c288_err error
			var_600227c4c288_mapped, var_600227c4c288_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_600227c4c288)
			if var_600227c4c288_err != nil {
				panic(var_600227c4c288_err)
			}

			var_4de57fae9c91_l = append(var_4de57fae9c91_l, var_600227c4c288_mapped)
		}
		var_4de57fae9c91_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_4de57fae9c91_l})
		properties["namespaces"] = var_4de57fae9c91_mapped
	}

	var_76f45425fbf2 := extensionEventSelector.Resources

	if var_76f45425fbf2 != nil {
		var var_76f45425fbf2_mapped *structpb.Value

		var var_76f45425fbf2_l []*structpb.Value
		for _, value := range var_76f45425fbf2 {

			var_c47e2b8048a2 := value
			var var_c47e2b8048a2_mapped *structpb.Value

			var var_c47e2b8048a2_err error
			var_c47e2b8048a2_mapped, var_c47e2b8048a2_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_c47e2b8048a2)
			if var_c47e2b8048a2_err != nil {
				panic(var_c47e2b8048a2_err)
			}

			var_76f45425fbf2_l = append(var_76f45425fbf2_l, var_c47e2b8048a2_mapped)
		}
		var_76f45425fbf2_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_76f45425fbf2_l})
		properties["resources"] = var_76f45425fbf2_mapped
	}

	var_86cac6f4235d := extensionEventSelector.Ids

	if var_86cac6f4235d != nil {
		var var_86cac6f4235d_mapped *structpb.Value

		var var_86cac6f4235d_l []*structpb.Value
		for _, value := range var_86cac6f4235d {

			var_9cd1a0b178f5 := value
			var var_9cd1a0b178f5_mapped *structpb.Value

			var var_9cd1a0b178f5_err error
			var_9cd1a0b178f5_mapped, var_9cd1a0b178f5_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_9cd1a0b178f5)
			if var_9cd1a0b178f5_err != nil {
				panic(var_9cd1a0b178f5_err)
			}

			var_86cac6f4235d_l = append(var_86cac6f4235d_l, var_9cd1a0b178f5_mapped)
		}
		var_86cac6f4235d_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_86cac6f4235d_l})
		properties["ids"] = var_86cac6f4235d_mapped
	}

	var_aa4e3a2af544 := extensionEventSelector.Annotations

	if var_aa4e3a2af544 != nil {
		var var_aa4e3a2af544_mapped *structpb.Value

		var var_aa4e3a2af544_st *structpb.Struct = new(structpb.Struct)
		var_aa4e3a2af544_st.Fields = make(map[string]*structpb.Value)
		for key, value := range var_aa4e3a2af544 {

			var_c8699a7d6fe6 := value
			var var_c8699a7d6fe6_mapped *structpb.Value

			var var_c8699a7d6fe6_err error
			var_c8699a7d6fe6_mapped, var_c8699a7d6fe6_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_c8699a7d6fe6)
			if var_c8699a7d6fe6_err != nil {
				panic(var_c8699a7d6fe6_err)
			}

			var_aa4e3a2af544_st.Fields[key] = var_c8699a7d6fe6_mapped
		}
		var_aa4e3a2af544_mapped = structpb.NewStructValue(var_aa4e3a2af544_st)
		properties["annotations"] = var_aa4e3a2af544_mapped
	}
	return properties
}

func (m *ExtensionEventSelectorMapper) FromProperties(properties map[string]*structpb.Value) *ExtensionEventSelector {
	var s = m.New()
	if properties["actions"] != nil {

		var_183a91da161a := properties["actions"]
		var_183a91da161a_mapped := []EventAction{}
		for _, v := range var_183a91da161a.GetListValue().Values {

			var_c5ea85f8d00e := v
			var_c5ea85f8d00e_mapped := (EventAction)(var_c5ea85f8d00e.GetStringValue())

			var_183a91da161a_mapped = append(var_183a91da161a_mapped, var_c5ea85f8d00e_mapped)
		}

		s.Actions = var_183a91da161a_mapped
	}
	if properties["recordSelector"] != nil {

		var_58900742804d := properties["recordSelector"]
		var mappedValue = ExtensionBooleanExpressionMapperInstance.FromProperties(var_58900742804d.GetStructValue().Fields)

		var_58900742804d_mapped := mappedValue

		s.RecordSelector = var_58900742804d_mapped
	}
	if properties["namespaces"] != nil {

		var_f723f5c17835 := properties["namespaces"]
		var_f723f5c17835_mapped := []string{}
		for _, v := range var_f723f5c17835.GetListValue().Values {

			var_f0f1a8abd9ab := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_f0f1a8abd9ab)

			if err != nil {
				panic(err)
			}

			var_f0f1a8abd9ab_mapped := val.(string)

			var_f723f5c17835_mapped = append(var_f723f5c17835_mapped, var_f0f1a8abd9ab_mapped)
		}

		s.Namespaces = var_f723f5c17835_mapped
	}
	if properties["resources"] != nil {

		var_44929d391cfa := properties["resources"]
		var_44929d391cfa_mapped := []string{}
		for _, v := range var_44929d391cfa.GetListValue().Values {

			var_9d3ada035a84 := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_9d3ada035a84)

			if err != nil {
				panic(err)
			}

			var_9d3ada035a84_mapped := val.(string)

			var_44929d391cfa_mapped = append(var_44929d391cfa_mapped, var_9d3ada035a84_mapped)
		}

		s.Resources = var_44929d391cfa_mapped
	}
	if properties["ids"] != nil {

		var_a26fae104dc2 := properties["ids"]
		var_a26fae104dc2_mapped := []string{}
		for _, v := range var_a26fae104dc2.GetListValue().Values {

			var_3b50cbfb86d3 := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_3b50cbfb86d3)

			if err != nil {
				panic(err)
			}

			var_3b50cbfb86d3_mapped := val.(string)

			var_a26fae104dc2_mapped = append(var_a26fae104dc2_mapped, var_3b50cbfb86d3_mapped)
		}

		s.Ids = var_a26fae104dc2_mapped
	}
	if properties["annotations"] != nil {

		var_6a367378c97a := properties["annotations"]
		var_6a367378c97a_mapped := make(map[string]string)
		for k, v := range var_6a367378c97a.GetStructValue().Fields {

			var_baeb887d6b55 := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_baeb887d6b55)

			if err != nil {
				panic(err)
			}

			var_baeb887d6b55_mapped := val.(string)

			var_6a367378c97a_mapped[k] = var_baeb887d6b55_mapped
		}

		s.Annotations = var_6a367378c97a_mapped
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

	var_0c064dab84cf := extensionRecordSearchParams.Query

	if var_0c064dab84cf != nil {
		var var_0c064dab84cf_mapped *structpb.Value

		var_0c064dab84cf_mapped = structpb.NewStructValue(&structpb.Struct{Fields: ExtensionBooleanExpressionMapperInstance.ToProperties(var_0c064dab84cf)})
		properties["query"] = var_0c064dab84cf_mapped
	}

	var_5ff316931599 := extensionRecordSearchParams.Limit

	if var_5ff316931599 != nil {
		var var_5ff316931599_mapped *structpb.Value

		var var_5ff316931599_err error
		var_5ff316931599_mapped, var_5ff316931599_err = types.ByResourcePropertyType(model.ResourceProperty_INT32).Pack(*var_5ff316931599)
		if var_5ff316931599_err != nil {
			panic(var_5ff316931599_err)
		}
		properties["limit"] = var_5ff316931599_mapped
	}

	var_2f2ec99c48d4 := extensionRecordSearchParams.Offset

	if var_2f2ec99c48d4 != nil {
		var var_2f2ec99c48d4_mapped *structpb.Value

		var var_2f2ec99c48d4_err error
		var_2f2ec99c48d4_mapped, var_2f2ec99c48d4_err = types.ByResourcePropertyType(model.ResourceProperty_INT32).Pack(*var_2f2ec99c48d4)
		if var_2f2ec99c48d4_err != nil {
			panic(var_2f2ec99c48d4_err)
		}
		properties["offset"] = var_2f2ec99c48d4_mapped
	}

	var_48fdb3e22160 := extensionRecordSearchParams.ResolveReferences

	if var_48fdb3e22160 != nil {
		var var_48fdb3e22160_mapped *structpb.Value

		var var_48fdb3e22160_l []*structpb.Value
		for _, value := range var_48fdb3e22160 {

			var_2665034ac508 := value
			var var_2665034ac508_mapped *structpb.Value

			var var_2665034ac508_err error
			var_2665034ac508_mapped, var_2665034ac508_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_2665034ac508)
			if var_2665034ac508_err != nil {
				panic(var_2665034ac508_err)
			}

			var_48fdb3e22160_l = append(var_48fdb3e22160_l, var_2665034ac508_mapped)
		}
		var_48fdb3e22160_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_48fdb3e22160_l})
		properties["resolveReferences"] = var_48fdb3e22160_mapped
	}
	return properties
}

func (m *ExtensionRecordSearchParamsMapper) FromProperties(properties map[string]*structpb.Value) *ExtensionRecordSearchParams {
	var s = m.New()
	if properties["query"] != nil {

		var_729e65b763d0 := properties["query"]
		var mappedValue = ExtensionBooleanExpressionMapperInstance.FromProperties(var_729e65b763d0.GetStructValue().Fields)

		var_729e65b763d0_mapped := mappedValue

		s.Query = var_729e65b763d0_mapped
	}
	if properties["limit"] != nil {

		var_6b2fa120499f := properties["limit"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(var_6b2fa120499f)

		if err != nil {
			panic(err)
		}

		var_6b2fa120499f_mapped := new(int32)
		*var_6b2fa120499f_mapped = val.(int32)

		s.Limit = var_6b2fa120499f_mapped
	}
	if properties["offset"] != nil {

		var_80853768bcfd := properties["offset"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(var_80853768bcfd)

		if err != nil {
			panic(err)
		}

		var_80853768bcfd_mapped := new(int32)
		*var_80853768bcfd_mapped = val.(int32)

		s.Offset = var_80853768bcfd_mapped
	}
	if properties["resolveReferences"] != nil {

		var_a8d845659e88 := properties["resolveReferences"]
		var_a8d845659e88_mapped := []string{}
		for _, v := range var_a8d845659e88.GetListValue().Values {

			var_8091740c07c8 := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_8091740c07c8)

			if err != nil {
				panic(err)
			}

			var_8091740c07c8_mapped := val.(string)

			var_a8d845659e88_mapped = append(var_a8d845659e88_mapped, var_8091740c07c8_mapped)
		}

		s.ResolveReferences = var_a8d845659e88_mapped
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

	var_4ffd255ad399 := extensionEvent.Id

	if var_4ffd255ad399 != nil {
		var var_4ffd255ad399_mapped *structpb.Value

		var var_4ffd255ad399_err error
		var_4ffd255ad399_mapped, var_4ffd255ad399_err = types.ByResourcePropertyType(model.ResourceProperty_UUID).Pack(*var_4ffd255ad399)
		if var_4ffd255ad399_err != nil {
			panic(var_4ffd255ad399_err)
		}
		properties["id"] = var_4ffd255ad399_mapped
	}

	var_bbfe98ba0afa := extensionEvent.Action

	var var_bbfe98ba0afa_mapped *structpb.Value

	var var_bbfe98ba0afa_err error
	var_bbfe98ba0afa_mapped, var_bbfe98ba0afa_err = types.ByResourcePropertyType(model.ResourceProperty_ENUM).Pack(string(var_bbfe98ba0afa))
	if var_bbfe98ba0afa_err != nil {
		panic(var_bbfe98ba0afa_err)
	}
	properties["action"] = var_bbfe98ba0afa_mapped

	var_42798d945423 := extensionEvent.RecordSearchParams

	if var_42798d945423 != nil {
		var var_42798d945423_mapped *structpb.Value

		var_42798d945423_mapped = structpb.NewStructValue(&structpb.Struct{Fields: ExtensionRecordSearchParamsMapperInstance.ToProperties(var_42798d945423)})
		properties["recordSearchParams"] = var_42798d945423_mapped
	}

	var_896198cab52c := extensionEvent.ActionSummary

	if var_896198cab52c != nil {
		var var_896198cab52c_mapped *structpb.Value

		var var_896198cab52c_err error
		var_896198cab52c_mapped, var_896198cab52c_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_896198cab52c)
		if var_896198cab52c_err != nil {
			panic(var_896198cab52c_err)
		}
		properties["actionSummary"] = var_896198cab52c_mapped
	}

	var_80e75def46ce := extensionEvent.ActionDescription

	if var_80e75def46ce != nil {
		var var_80e75def46ce_mapped *structpb.Value

		var var_80e75def46ce_err error
		var_80e75def46ce_mapped, var_80e75def46ce_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_80e75def46ce)
		if var_80e75def46ce_err != nil {
			panic(var_80e75def46ce_err)
		}
		properties["actionDescription"] = var_80e75def46ce_mapped
	}

	var_3f15f5edd426 := extensionEvent.Resource

	if var_3f15f5edd426 != nil {
		var var_3f15f5edd426_mapped *structpb.Value

		var_3f15f5edd426_mapped = structpb.NewStructValue(&structpb.Struct{Fields: ResourceMapperInstance.ToProperties(var_3f15f5edd426)})
		properties["resource"] = var_3f15f5edd426_mapped
	}

	var_c2b4310657ae := extensionEvent.Records

	if var_c2b4310657ae != nil {
		var var_c2b4310657ae_mapped *structpb.Value

		var var_c2b4310657ae_l []*structpb.Value
		for _, value := range var_c2b4310657ae {

			var_2f2f24782dc4 := value
			var var_2f2f24782dc4_mapped *structpb.Value

			var_2f2f24782dc4_mapped = structpb.NewStructValue(&structpb.Struct{Fields: RecordMapperInstance.ToProperties(var_2f2f24782dc4)})

			var_c2b4310657ae_l = append(var_c2b4310657ae_l, var_2f2f24782dc4_mapped)
		}
		var_c2b4310657ae_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_c2b4310657ae_l})
		properties["records"] = var_c2b4310657ae_mapped
	}

	var_f3f39b467937 := extensionEvent.Ids

	if var_f3f39b467937 != nil {
		var var_f3f39b467937_mapped *structpb.Value

		var var_f3f39b467937_l []*structpb.Value
		for _, value := range var_f3f39b467937 {

			var_5fcdb9c291de := value
			var var_5fcdb9c291de_mapped *structpb.Value

			var var_5fcdb9c291de_err error
			var_5fcdb9c291de_mapped, var_5fcdb9c291de_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_5fcdb9c291de)
			if var_5fcdb9c291de_err != nil {
				panic(var_5fcdb9c291de_err)
			}

			var_f3f39b467937_l = append(var_f3f39b467937_l, var_5fcdb9c291de_mapped)
		}
		var_f3f39b467937_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_f3f39b467937_l})
		properties["ids"] = var_f3f39b467937_mapped
	}

	var_3928667d34c9 := extensionEvent.Finalizes

	if var_3928667d34c9 != nil {
		var var_3928667d34c9_mapped *structpb.Value

		var var_3928667d34c9_err error
		var_3928667d34c9_mapped, var_3928667d34c9_err = types.ByResourcePropertyType(model.ResourceProperty_BOOL).Pack(*var_3928667d34c9)
		if var_3928667d34c9_err != nil {
			panic(var_3928667d34c9_err)
		}
		properties["finalizes"] = var_3928667d34c9_mapped
	}

	var_eed34c9c38d7 := extensionEvent.Sync

	if var_eed34c9c38d7 != nil {
		var var_eed34c9c38d7_mapped *structpb.Value

		var var_eed34c9c38d7_err error
		var_eed34c9c38d7_mapped, var_eed34c9c38d7_err = types.ByResourcePropertyType(model.ResourceProperty_BOOL).Pack(*var_eed34c9c38d7)
		if var_eed34c9c38d7_err != nil {
			panic(var_eed34c9c38d7_err)
		}
		properties["sync"] = var_eed34c9c38d7_mapped
	}

	var_0adf127df71e := extensionEvent.Time

	if var_0adf127df71e != nil {
		var var_0adf127df71e_mapped *structpb.Value

		var var_0adf127df71e_err error
		var_0adf127df71e_mapped, var_0adf127df71e_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_0adf127df71e)
		if var_0adf127df71e_err != nil {
			panic(var_0adf127df71e_err)
		}
		properties["time"] = var_0adf127df71e_mapped
	}

	var_c0d696c7b167 := extensionEvent.Annotations

	if var_c0d696c7b167 != nil {
		var var_c0d696c7b167_mapped *structpb.Value

		var var_c0d696c7b167_st *structpb.Struct = new(structpb.Struct)
		var_c0d696c7b167_st.Fields = make(map[string]*structpb.Value)
		for key, value := range var_c0d696c7b167 {

			var_7c4415646057 := value
			var var_7c4415646057_mapped *structpb.Value

			var var_7c4415646057_err error
			var_7c4415646057_mapped, var_7c4415646057_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_7c4415646057)
			if var_7c4415646057_err != nil {
				panic(var_7c4415646057_err)
			}

			var_c0d696c7b167_st.Fields[key] = var_7c4415646057_mapped
		}
		var_c0d696c7b167_mapped = structpb.NewStructValue(var_c0d696c7b167_st)
		properties["annotations"] = var_c0d696c7b167_mapped
	}
	return properties
}

func (m *ExtensionEventMapper) FromProperties(properties map[string]*structpb.Value) *ExtensionEvent {
	var s = m.New()
	if properties["id"] != nil {

		var_dcc30cad9507 := properties["id"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_UUID).UnPack(var_dcc30cad9507)

		if err != nil {
			panic(err)
		}

		var_dcc30cad9507_mapped := new(uuid.UUID)
		*var_dcc30cad9507_mapped = val.(uuid.UUID)

		s.Id = var_dcc30cad9507_mapped
	}
	if properties["action"] != nil {

		var_ef7cb413e1cd := properties["action"]
		var_ef7cb413e1cd_mapped := (EventAction)(var_ef7cb413e1cd.GetStringValue())

		s.Action = var_ef7cb413e1cd_mapped
	}
	if properties["recordSearchParams"] != nil {

		var_80c9cf2b4c23 := properties["recordSearchParams"]
		var mappedValue = ExtensionRecordSearchParamsMapperInstance.FromProperties(var_80c9cf2b4c23.GetStructValue().Fields)

		var_80c9cf2b4c23_mapped := mappedValue

		s.RecordSearchParams = var_80c9cf2b4c23_mapped
	}
	if properties["actionSummary"] != nil {

		var_b1ca8c9fc0b5 := properties["actionSummary"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_b1ca8c9fc0b5)

		if err != nil {
			panic(err)
		}

		var_b1ca8c9fc0b5_mapped := new(string)
		*var_b1ca8c9fc0b5_mapped = val.(string)

		s.ActionSummary = var_b1ca8c9fc0b5_mapped
	}
	if properties["actionDescription"] != nil {

		var_a90ab0f201f7 := properties["actionDescription"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_a90ab0f201f7)

		if err != nil {
			panic(err)
		}

		var_a90ab0f201f7_mapped := new(string)
		*var_a90ab0f201f7_mapped = val.(string)

		s.ActionDescription = var_a90ab0f201f7_mapped
	}
	if properties["resource"] != nil {

		var_49a3f4449f04 := properties["resource"]
		var_49a3f4449f04_mapped := ResourceMapperInstance.FromProperties(var_49a3f4449f04.GetStructValue().Fields)

		s.Resource = var_49a3f4449f04_mapped
	}
	if properties["records"] != nil {

		var_1091dd1c021c := properties["records"]
		var_1091dd1c021c_mapped := []*Record{}
		for _, v := range var_1091dd1c021c.GetListValue().Values {

			var_55e18f9e7e6e := v
			var_55e18f9e7e6e_mapped := RecordMapperInstance.FromProperties(var_55e18f9e7e6e.GetStructValue().Fields)

			var_1091dd1c021c_mapped = append(var_1091dd1c021c_mapped, var_55e18f9e7e6e_mapped)
		}

		s.Records = var_1091dd1c021c_mapped
	}
	if properties["ids"] != nil {

		var_ae4433a622c2 := properties["ids"]
		var_ae4433a622c2_mapped := []string{}
		for _, v := range var_ae4433a622c2.GetListValue().Values {

			var_7b277da4f9fb := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_7b277da4f9fb)

			if err != nil {
				panic(err)
			}

			var_7b277da4f9fb_mapped := val.(string)

			var_ae4433a622c2_mapped = append(var_ae4433a622c2_mapped, var_7b277da4f9fb_mapped)
		}

		s.Ids = var_ae4433a622c2_mapped
	}
	if properties["finalizes"] != nil {

		var_b718a2ea0273 := properties["finalizes"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_BOOL).UnPack(var_b718a2ea0273)

		if err != nil {
			panic(err)
		}

		var_b718a2ea0273_mapped := new(bool)
		*var_b718a2ea0273_mapped = val.(bool)

		s.Finalizes = var_b718a2ea0273_mapped
	}
	if properties["sync"] != nil {

		var_70cbabc98de8 := properties["sync"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_BOOL).UnPack(var_70cbabc98de8)

		if err != nil {
			panic(err)
		}

		var_70cbabc98de8_mapped := new(bool)
		*var_70cbabc98de8_mapped = val.(bool)

		s.Sync = var_70cbabc98de8_mapped
	}
	if properties["time"] != nil {

		var_1955bb7b7f6e := properties["time"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_1955bb7b7f6e)

		if err != nil {
			panic(err)
		}

		var_1955bb7b7f6e_mapped := new(time.Time)
		*var_1955bb7b7f6e_mapped = val.(time.Time)

		s.Time = var_1955bb7b7f6e_mapped
	}
	if properties["annotations"] != nil {

		var_c12c1f75a9a0 := properties["annotations"]
		var_c12c1f75a9a0_mapped := make(map[string]string)
		for k, v := range var_c12c1f75a9a0.GetStructValue().Fields {

			var_5f7b73e85021 := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_5f7b73e85021)

			if err != nil {
				panic(err)
			}

			var_5f7b73e85021_mapped := val.(string)

			var_c12c1f75a9a0_mapped[k] = var_5f7b73e85021_mapped
		}

		s.Annotations = var_c12c1f75a9a0_mapped
	}
	return s
}
