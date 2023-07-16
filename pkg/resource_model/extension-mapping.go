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

	if extension.Id != nil {
		rec.Id = extension.Id.String()
	}

	return rec
}

func (m *ExtensionMapper) FromRecord(record *model.Record) *Extension {
	return m.FromProperties(record.Properties)
}

func (m *ExtensionMapper) ToProperties(extension *Extension) map[string]*structpb.Value {
	var properties = make(map[string]*structpb.Value)

	var_e97382d75578 := extension.Id

	if var_e97382d75578 != nil {
		var var_e97382d75578_mapped *structpb.Value

		var var_e97382d75578_err error
		var_e97382d75578_mapped, var_e97382d75578_err = types.ByResourcePropertyType(model.ResourceProperty_UUID).Pack(*var_e97382d75578)
		if var_e97382d75578_err != nil {
			panic(var_e97382d75578_err)
		}
		properties["id"] = var_e97382d75578_mapped
	}

	var_615b4434c691 := extension.Version

	var var_615b4434c691_mapped *structpb.Value

	var var_615b4434c691_err error
	var_615b4434c691_mapped, var_615b4434c691_err = types.ByResourcePropertyType(model.ResourceProperty_INT32).Pack(var_615b4434c691)
	if var_615b4434c691_err != nil {
		panic(var_615b4434c691_err)
	}
	properties["version"] = var_615b4434c691_mapped

	var_ebd41c673dd1 := extension.CreatedBy

	if var_ebd41c673dd1 != nil {
		var var_ebd41c673dd1_mapped *structpb.Value

		var var_ebd41c673dd1_err error
		var_ebd41c673dd1_mapped, var_ebd41c673dd1_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_ebd41c673dd1)
		if var_ebd41c673dd1_err != nil {
			panic(var_ebd41c673dd1_err)
		}
		properties["createdBy"] = var_ebd41c673dd1_mapped
	}

	var_a999aac4294a := extension.UpdatedBy

	if var_a999aac4294a != nil {
		var var_a999aac4294a_mapped *structpb.Value

		var var_a999aac4294a_err error
		var_a999aac4294a_mapped, var_a999aac4294a_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_a999aac4294a)
		if var_a999aac4294a_err != nil {
			panic(var_a999aac4294a_err)
		}
		properties["updatedBy"] = var_a999aac4294a_mapped
	}

	var_e8e85c12312a := extension.CreatedOn

	if var_e8e85c12312a != nil {
		var var_e8e85c12312a_mapped *structpb.Value

		var var_e8e85c12312a_err error
		var_e8e85c12312a_mapped, var_e8e85c12312a_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_e8e85c12312a)
		if var_e8e85c12312a_err != nil {
			panic(var_e8e85c12312a_err)
		}
		properties["createdOn"] = var_e8e85c12312a_mapped
	}

	var_e759679436a5 := extension.UpdatedOn

	if var_e759679436a5 != nil {
		var var_e759679436a5_mapped *structpb.Value

		var var_e759679436a5_err error
		var_e759679436a5_mapped, var_e759679436a5_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_e759679436a5)
		if var_e759679436a5_err != nil {
			panic(var_e759679436a5_err)
		}
		properties["updatedOn"] = var_e759679436a5_mapped
	}

	var_007cc3fdece5 := extension.Name

	var var_007cc3fdece5_mapped *structpb.Value

	var var_007cc3fdece5_err error
	var_007cc3fdece5_mapped, var_007cc3fdece5_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_007cc3fdece5)
	if var_007cc3fdece5_err != nil {
		panic(var_007cc3fdece5_err)
	}
	properties["name"] = var_007cc3fdece5_mapped

	var_f78b4434ffb1 := extension.Description

	if var_f78b4434ffb1 != nil {
		var var_f78b4434ffb1_mapped *structpb.Value

		var var_f78b4434ffb1_err error
		var_f78b4434ffb1_mapped, var_f78b4434ffb1_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_f78b4434ffb1)
		if var_f78b4434ffb1_err != nil {
			panic(var_f78b4434ffb1_err)
		}
		properties["description"] = var_f78b4434ffb1_mapped
	}

	var_10914bc6dab7 := extension.Selector

	if var_10914bc6dab7 != nil {
		var var_10914bc6dab7_mapped *structpb.Value

		var_10914bc6dab7_mapped = structpb.NewStructValue(&structpb.Struct{Fields: ExtensionEventSelectorMapperInstance.ToProperties(var_10914bc6dab7)})
		properties["selector"] = var_10914bc6dab7_mapped
	}

	var_2776ffe0a1db := extension.Order

	var var_2776ffe0a1db_mapped *structpb.Value

	var var_2776ffe0a1db_err error
	var_2776ffe0a1db_mapped, var_2776ffe0a1db_err = types.ByResourcePropertyType(model.ResourceProperty_INT32).Pack(var_2776ffe0a1db)
	if var_2776ffe0a1db_err != nil {
		panic(var_2776ffe0a1db_err)
	}
	properties["order"] = var_2776ffe0a1db_mapped

	var_c39ceaebe900 := extension.Finalizes

	var var_c39ceaebe900_mapped *structpb.Value

	var var_c39ceaebe900_err error
	var_c39ceaebe900_mapped, var_c39ceaebe900_err = types.ByResourcePropertyType(model.ResourceProperty_BOOL).Pack(var_c39ceaebe900)
	if var_c39ceaebe900_err != nil {
		panic(var_c39ceaebe900_err)
	}
	properties["finalizes"] = var_c39ceaebe900_mapped

	var_6007c7a1185c := extension.Sync

	var var_6007c7a1185c_mapped *structpb.Value

	var var_6007c7a1185c_err error
	var_6007c7a1185c_mapped, var_6007c7a1185c_err = types.ByResourcePropertyType(model.ResourceProperty_BOOL).Pack(var_6007c7a1185c)
	if var_6007c7a1185c_err != nil {
		panic(var_6007c7a1185c_err)
	}
	properties["sync"] = var_6007c7a1185c_mapped

	var_0fbaf13789bd := extension.Responds

	var var_0fbaf13789bd_mapped *structpb.Value

	var var_0fbaf13789bd_err error
	var_0fbaf13789bd_mapped, var_0fbaf13789bd_err = types.ByResourcePropertyType(model.ResourceProperty_BOOL).Pack(var_0fbaf13789bd)
	if var_0fbaf13789bd_err != nil {
		panic(var_0fbaf13789bd_err)
	}
	properties["responds"] = var_0fbaf13789bd_mapped

	var_41fdc441003b := extension.Call

	var var_41fdc441003b_mapped *structpb.Value

	var_41fdc441003b_mapped = structpb.NewStructValue(&structpb.Struct{Fields: ExtensionExternalCallMapperInstance.ToProperties(&var_41fdc441003b)})
	properties["call"] = var_41fdc441003b_mapped

	var_e6515fa6582f := extension.Annotations

	if var_e6515fa6582f != nil {
		var var_e6515fa6582f_mapped *structpb.Value

		var var_e6515fa6582f_st *structpb.Struct = new(structpb.Struct)
		var_e6515fa6582f_st.Fields = make(map[string]*structpb.Value)
		for key, value := range var_e6515fa6582f {

			var_35c4bde0c24f := value
			var var_35c4bde0c24f_mapped *structpb.Value

			var var_35c4bde0c24f_err error
			var_35c4bde0c24f_mapped, var_35c4bde0c24f_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_35c4bde0c24f)
			if var_35c4bde0c24f_err != nil {
				panic(var_35c4bde0c24f_err)
			}

			var_e6515fa6582f_st.Fields[key] = var_35c4bde0c24f_mapped
		}
		var_e6515fa6582f_mapped = structpb.NewStructValue(var_e6515fa6582f_st)
		properties["annotations"] = var_e6515fa6582f_mapped
	}
	return properties
}

func (m *ExtensionMapper) FromProperties(properties map[string]*structpb.Value) *Extension {
	var s = m.New()
	if properties["id"] != nil {

		var_42b56d28451c := properties["id"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_UUID).UnPack(var_42b56d28451c)

		if err != nil {
			panic(err)
		}

		var_42b56d28451c_mapped := new(uuid.UUID)
		*var_42b56d28451c_mapped = val.(uuid.UUID)

		s.Id = var_42b56d28451c_mapped
	}
	if properties["version"] != nil {

		var_b776e229e5c4 := properties["version"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(var_b776e229e5c4)

		if err != nil {
			panic(err)
		}

		var_b776e229e5c4_mapped := val.(int32)

		s.Version = var_b776e229e5c4_mapped
	}
	if properties["createdBy"] != nil {

		var_9291cd5ed8e7 := properties["createdBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_9291cd5ed8e7)

		if err != nil {
			panic(err)
		}

		var_9291cd5ed8e7_mapped := new(string)
		*var_9291cd5ed8e7_mapped = val.(string)

		s.CreatedBy = var_9291cd5ed8e7_mapped
	}
	if properties["updatedBy"] != nil {

		var_8657351347b6 := properties["updatedBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_8657351347b6)

		if err != nil {
			panic(err)
		}

		var_8657351347b6_mapped := new(string)
		*var_8657351347b6_mapped = val.(string)

		s.UpdatedBy = var_8657351347b6_mapped
	}
	if properties["createdOn"] != nil {

		var_bf55544a1eec := properties["createdOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_bf55544a1eec)

		if err != nil {
			panic(err)
		}

		var_bf55544a1eec_mapped := new(time.Time)
		*var_bf55544a1eec_mapped = val.(time.Time)

		s.CreatedOn = var_bf55544a1eec_mapped
	}
	if properties["updatedOn"] != nil {

		var_767238d52383 := properties["updatedOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_767238d52383)

		if err != nil {
			panic(err)
		}

		var_767238d52383_mapped := new(time.Time)
		*var_767238d52383_mapped = val.(time.Time)

		s.UpdatedOn = var_767238d52383_mapped
	}
	if properties["name"] != nil {

		var_95949f857822 := properties["name"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_95949f857822)

		if err != nil {
			panic(err)
		}

		var_95949f857822_mapped := val.(string)

		s.Name = var_95949f857822_mapped
	}
	if properties["description"] != nil {

		var_1ca5a7a3d30e := properties["description"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_1ca5a7a3d30e)

		if err != nil {
			panic(err)
		}

		var_1ca5a7a3d30e_mapped := new(string)
		*var_1ca5a7a3d30e_mapped = val.(string)

		s.Description = var_1ca5a7a3d30e_mapped
	}
	if properties["selector"] != nil {

		var_1b52173ef4d1 := properties["selector"]
		var mappedValue = ExtensionEventSelectorMapperInstance.FromProperties(var_1b52173ef4d1.GetStructValue().Fields)

		var_1b52173ef4d1_mapped := mappedValue

		s.Selector = var_1b52173ef4d1_mapped
	}
	if properties["order"] != nil {

		var_dfa30297d601 := properties["order"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(var_dfa30297d601)

		if err != nil {
			panic(err)
		}

		var_dfa30297d601_mapped := val.(int32)

		s.Order = var_dfa30297d601_mapped
	}
	if properties["finalizes"] != nil {

		var_0e432484243e := properties["finalizes"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_BOOL).UnPack(var_0e432484243e)

		if err != nil {
			panic(err)
		}

		var_0e432484243e_mapped := val.(bool)

		s.Finalizes = var_0e432484243e_mapped
	}
	if properties["sync"] != nil {

		var_16386c8e212f := properties["sync"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_BOOL).UnPack(var_16386c8e212f)

		if err != nil {
			panic(err)
		}

		var_16386c8e212f_mapped := val.(bool)

		s.Sync = var_16386c8e212f_mapped
	}
	if properties["responds"] != nil {

		var_e59ee1471d96 := properties["responds"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_BOOL).UnPack(var_e59ee1471d96)

		if err != nil {
			panic(err)
		}

		var_e59ee1471d96_mapped := val.(bool)

		s.Responds = var_e59ee1471d96_mapped
	}
	if properties["call"] != nil {

		var_6f66c45cf6dc := properties["call"]
		var mappedValue = ExtensionExternalCallMapperInstance.FromProperties(var_6f66c45cf6dc.GetStructValue().Fields)

		var_6f66c45cf6dc_mapped := *mappedValue

		s.Call = var_6f66c45cf6dc_mapped
	}
	if properties["annotations"] != nil {

		var_0d7ab14e9b2c := properties["annotations"]
		var_0d7ab14e9b2c_mapped := make(map[string]string)
		for k, v := range var_0d7ab14e9b2c.GetStructValue().Fields {

			var_2a3cdd5c9d5b := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_2a3cdd5c9d5b)

			if err != nil {
				panic(err)
			}

			var_2a3cdd5c9d5b_mapped := val.(string)

			var_0d7ab14e9b2c_mapped[k] = var_2a3cdd5c9d5b_mapped
		}

		s.Annotations = var_0d7ab14e9b2c_mapped
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

func (m *ExtensionFunctionCallMapper) ToProperties(extensionFunctionCall *ExtensionFunctionCall) map[string]*structpb.Value {
	var properties = make(map[string]*structpb.Value)

	var_27422feaf13e := extensionFunctionCall.Host

	var var_27422feaf13e_mapped *structpb.Value

	var var_27422feaf13e_err error
	var_27422feaf13e_mapped, var_27422feaf13e_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_27422feaf13e)
	if var_27422feaf13e_err != nil {
		panic(var_27422feaf13e_err)
	}
	properties["host"] = var_27422feaf13e_mapped

	var_20c57e796e21 := extensionFunctionCall.FunctionName

	var var_20c57e796e21_mapped *structpb.Value

	var var_20c57e796e21_err error
	var_20c57e796e21_mapped, var_20c57e796e21_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_20c57e796e21)
	if var_20c57e796e21_err != nil {
		panic(var_20c57e796e21_err)
	}
	properties["functionName"] = var_20c57e796e21_mapped
	return properties
}

func (m *ExtensionFunctionCallMapper) FromProperties(properties map[string]*structpb.Value) *ExtensionFunctionCall {
	var s = m.New()
	if properties["host"] != nil {

		var_6927b1aa1335 := properties["host"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_6927b1aa1335)

		if err != nil {
			panic(err)
		}

		var_6927b1aa1335_mapped := val.(string)

		s.Host = var_6927b1aa1335_mapped
	}
	if properties["functionName"] != nil {

		var_5a8fc36d8f8f := properties["functionName"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_5a8fc36d8f8f)

		if err != nil {
			panic(err)
		}

		var_5a8fc36d8f8f_mapped := val.(string)

		s.FunctionName = var_5a8fc36d8f8f_mapped
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

func (m *ExtensionHttpCallMapper) ToProperties(extensionHttpCall *ExtensionHttpCall) map[string]*structpb.Value {
	var properties = make(map[string]*structpb.Value)

	var_a2b742645f7b := extensionHttpCall.Uri

	var var_a2b742645f7b_mapped *structpb.Value

	var var_a2b742645f7b_err error
	var_a2b742645f7b_mapped, var_a2b742645f7b_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_a2b742645f7b)
	if var_a2b742645f7b_err != nil {
		panic(var_a2b742645f7b_err)
	}
	properties["uri"] = var_a2b742645f7b_mapped

	var_2cfa3249c453 := extensionHttpCall.Method

	var var_2cfa3249c453_mapped *structpb.Value

	var var_2cfa3249c453_err error
	var_2cfa3249c453_mapped, var_2cfa3249c453_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_2cfa3249c453)
	if var_2cfa3249c453_err != nil {
		panic(var_2cfa3249c453_err)
	}
	properties["method"] = var_2cfa3249c453_mapped
	return properties
}

func (m *ExtensionHttpCallMapper) FromProperties(properties map[string]*structpb.Value) *ExtensionHttpCall {
	var s = m.New()
	if properties["uri"] != nil {

		var_91b2e7ee74a9 := properties["uri"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_91b2e7ee74a9)

		if err != nil {
			panic(err)
		}

		var_91b2e7ee74a9_mapped := val.(string)

		s.Uri = var_91b2e7ee74a9_mapped
	}
	if properties["method"] != nil {

		var_4e270e796c7e := properties["method"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_4e270e796c7e)

		if err != nil {
			panic(err)
		}

		var_4e270e796c7e_mapped := val.(string)

		s.Method = var_4e270e796c7e_mapped
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

func (m *ExtensionExternalCallMapper) ToProperties(extensionExternalCall *ExtensionExternalCall) map[string]*structpb.Value {
	var properties = make(map[string]*structpb.Value)

	var_9483ae07f95d := extensionExternalCall.FunctionCall

	if var_9483ae07f95d != nil {
		var var_9483ae07f95d_mapped *structpb.Value

		var_9483ae07f95d_mapped = structpb.NewStructValue(&structpb.Struct{Fields: ExtensionFunctionCallMapperInstance.ToProperties(var_9483ae07f95d)})
		properties["functionCall"] = var_9483ae07f95d_mapped
	}

	var_3c8f6327ddb8 := extensionExternalCall.HttpCall

	if var_3c8f6327ddb8 != nil {
		var var_3c8f6327ddb8_mapped *structpb.Value

		var_3c8f6327ddb8_mapped = structpb.NewStructValue(&structpb.Struct{Fields: ExtensionHttpCallMapperInstance.ToProperties(var_3c8f6327ddb8)})
		properties["httpCall"] = var_3c8f6327ddb8_mapped
	}
	return properties
}

func (m *ExtensionExternalCallMapper) FromProperties(properties map[string]*structpb.Value) *ExtensionExternalCall {
	var s = m.New()
	if properties["functionCall"] != nil {

		var_ab4f665bfd53 := properties["functionCall"]
		var mappedValue = ExtensionFunctionCallMapperInstance.FromProperties(var_ab4f665bfd53.GetStructValue().Fields)

		var_ab4f665bfd53_mapped := mappedValue

		s.FunctionCall = var_ab4f665bfd53_mapped
	}
	if properties["httpCall"] != nil {

		var_a849687584f5 := properties["httpCall"]
		var mappedValue = ExtensionHttpCallMapperInstance.FromProperties(var_a849687584f5.GetStructValue().Fields)

		var_a849687584f5_mapped := mappedValue

		s.HttpCall = var_a849687584f5_mapped
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

func (m *ExtensionEventSelectorMapper) ToProperties(extensionEventSelector *ExtensionEventSelector) map[string]*structpb.Value {
	var properties = make(map[string]*structpb.Value)

	var_732535e478eb := extensionEventSelector.Actions

	if var_732535e478eb != nil {
		var var_732535e478eb_mapped *structpb.Value

		var var_732535e478eb_l []*structpb.Value
		for _, value := range var_732535e478eb {

			var_73a885e18191 := value
			var var_73a885e18191_mapped *structpb.Value

			var var_73a885e18191_err error
			var_73a885e18191_mapped, var_73a885e18191_err = types.ByResourcePropertyType(model.ResourceProperty_ENUM).Pack(string(var_73a885e18191))
			if var_73a885e18191_err != nil {
				panic(var_73a885e18191_err)
			}

			var_732535e478eb_l = append(var_732535e478eb_l, var_73a885e18191_mapped)
		}
		var_732535e478eb_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_732535e478eb_l})
		properties["actions"] = var_732535e478eb_mapped
	}

	var_f5ec8a4e2d49 := extensionEventSelector.RecordSelector

	if var_f5ec8a4e2d49 != nil {
		var var_f5ec8a4e2d49_mapped *structpb.Value

		var_f5ec8a4e2d49_mapped = structpb.NewStructValue(&structpb.Struct{Fields: ExtensionBooleanExpressionMapperInstance.ToProperties(var_f5ec8a4e2d49)})
		properties["recordSelector"] = var_f5ec8a4e2d49_mapped
	}

	var_2acef4b4668b := extensionEventSelector.Namespaces

	if var_2acef4b4668b != nil {
		var var_2acef4b4668b_mapped *structpb.Value

		var var_2acef4b4668b_l []*structpb.Value
		for _, value := range var_2acef4b4668b {

			var_652cf78b2fb6 := value
			var var_652cf78b2fb6_mapped *structpb.Value

			var var_652cf78b2fb6_err error
			var_652cf78b2fb6_mapped, var_652cf78b2fb6_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_652cf78b2fb6)
			if var_652cf78b2fb6_err != nil {
				panic(var_652cf78b2fb6_err)
			}

			var_2acef4b4668b_l = append(var_2acef4b4668b_l, var_652cf78b2fb6_mapped)
		}
		var_2acef4b4668b_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_2acef4b4668b_l})
		properties["namespaces"] = var_2acef4b4668b_mapped
	}

	var_931b98cca0a4 := extensionEventSelector.Resources

	if var_931b98cca0a4 != nil {
		var var_931b98cca0a4_mapped *structpb.Value

		var var_931b98cca0a4_l []*structpb.Value
		for _, value := range var_931b98cca0a4 {

			var_769e57e24708 := value
			var var_769e57e24708_mapped *structpb.Value

			var var_769e57e24708_err error
			var_769e57e24708_mapped, var_769e57e24708_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_769e57e24708)
			if var_769e57e24708_err != nil {
				panic(var_769e57e24708_err)
			}

			var_931b98cca0a4_l = append(var_931b98cca0a4_l, var_769e57e24708_mapped)
		}
		var_931b98cca0a4_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_931b98cca0a4_l})
		properties["resources"] = var_931b98cca0a4_mapped
	}

	var_b50aa3ce32cf := extensionEventSelector.Ids

	if var_b50aa3ce32cf != nil {
		var var_b50aa3ce32cf_mapped *structpb.Value

		var var_b50aa3ce32cf_l []*structpb.Value
		for _, value := range var_b50aa3ce32cf {

			var_9bc75b43cd59 := value
			var var_9bc75b43cd59_mapped *structpb.Value

			var var_9bc75b43cd59_err error
			var_9bc75b43cd59_mapped, var_9bc75b43cd59_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_9bc75b43cd59)
			if var_9bc75b43cd59_err != nil {
				panic(var_9bc75b43cd59_err)
			}

			var_b50aa3ce32cf_l = append(var_b50aa3ce32cf_l, var_9bc75b43cd59_mapped)
		}
		var_b50aa3ce32cf_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_b50aa3ce32cf_l})
		properties["ids"] = var_b50aa3ce32cf_mapped
	}

	var_b25e08249240 := extensionEventSelector.Annotations

	if var_b25e08249240 != nil {
		var var_b25e08249240_mapped *structpb.Value

		var var_b25e08249240_st *structpb.Struct = new(structpb.Struct)
		var_b25e08249240_st.Fields = make(map[string]*structpb.Value)
		for key, value := range var_b25e08249240 {

			var_c44ccf96677f := value
			var var_c44ccf96677f_mapped *structpb.Value

			var var_c44ccf96677f_err error
			var_c44ccf96677f_mapped, var_c44ccf96677f_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_c44ccf96677f)
			if var_c44ccf96677f_err != nil {
				panic(var_c44ccf96677f_err)
			}

			var_b25e08249240_st.Fields[key] = var_c44ccf96677f_mapped
		}
		var_b25e08249240_mapped = structpb.NewStructValue(var_b25e08249240_st)
		properties["annotations"] = var_b25e08249240_mapped
	}
	return properties
}

func (m *ExtensionEventSelectorMapper) FromProperties(properties map[string]*structpb.Value) *ExtensionEventSelector {
	var s = m.New()
	if properties["actions"] != nil {

		var_d4c8233b3b82 := properties["actions"]
		var_d4c8233b3b82_mapped := []EventAction{}
		for _, v := range var_d4c8233b3b82.GetListValue().Values {

			var_5e491397fe5f := v
			var_5e491397fe5f_mapped := (EventAction)(var_5e491397fe5f.GetStringValue())

			var_d4c8233b3b82_mapped = append(var_d4c8233b3b82_mapped, var_5e491397fe5f_mapped)
		}

		s.Actions = var_d4c8233b3b82_mapped
	}
	if properties["recordSelector"] != nil {

		var_09d6f1efcbdf := properties["recordSelector"]
		var mappedValue = ExtensionBooleanExpressionMapperInstance.FromProperties(var_09d6f1efcbdf.GetStructValue().Fields)

		var_09d6f1efcbdf_mapped := mappedValue

		s.RecordSelector = var_09d6f1efcbdf_mapped
	}
	if properties["namespaces"] != nil {

		var_95320a01a642 := properties["namespaces"]
		var_95320a01a642_mapped := []string{}
		for _, v := range var_95320a01a642.GetListValue().Values {

			var_8f36822761e8 := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_8f36822761e8)

			if err != nil {
				panic(err)
			}

			var_8f36822761e8_mapped := val.(string)

			var_95320a01a642_mapped = append(var_95320a01a642_mapped, var_8f36822761e8_mapped)
		}

		s.Namespaces = var_95320a01a642_mapped
	}
	if properties["resources"] != nil {

		var_719277d237ae := properties["resources"]
		var_719277d237ae_mapped := []string{}
		for _, v := range var_719277d237ae.GetListValue().Values {

			var_3bc3631f127e := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_3bc3631f127e)

			if err != nil {
				panic(err)
			}

			var_3bc3631f127e_mapped := val.(string)

			var_719277d237ae_mapped = append(var_719277d237ae_mapped, var_3bc3631f127e_mapped)
		}

		s.Resources = var_719277d237ae_mapped
	}
	if properties["ids"] != nil {

		var_6138a7c01e37 := properties["ids"]
		var_6138a7c01e37_mapped := []string{}
		for _, v := range var_6138a7c01e37.GetListValue().Values {

			var_33591d978e9d := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_33591d978e9d)

			if err != nil {
				panic(err)
			}

			var_33591d978e9d_mapped := val.(string)

			var_6138a7c01e37_mapped = append(var_6138a7c01e37_mapped, var_33591d978e9d_mapped)
		}

		s.Ids = var_6138a7c01e37_mapped
	}
	if properties["annotations"] != nil {

		var_ba0dc567bed7 := properties["annotations"]
		var_ba0dc567bed7_mapped := make(map[string]string)
		for k, v := range var_ba0dc567bed7.GetStructValue().Fields {

			var_d37e59311da9 := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_d37e59311da9)

			if err != nil {
				panic(err)
			}

			var_d37e59311da9_mapped := val.(string)

			var_ba0dc567bed7_mapped[k] = var_d37e59311da9_mapped
		}

		s.Annotations = var_ba0dc567bed7_mapped
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

func (m *ExtensionRecordSearchParamsMapper) ToProperties(extensionRecordSearchParams *ExtensionRecordSearchParams) map[string]*structpb.Value {
	var properties = make(map[string]*structpb.Value)

	var_aa397ed98732 := extensionRecordSearchParams.Query

	if var_aa397ed98732 != nil {
		var var_aa397ed98732_mapped *structpb.Value

		var_aa397ed98732_mapped = structpb.NewStructValue(&structpb.Struct{Fields: ExtensionBooleanExpressionMapperInstance.ToProperties(var_aa397ed98732)})
		properties["query"] = var_aa397ed98732_mapped
	}

	var_e4be7de92e70 := extensionRecordSearchParams.Limit

	if var_e4be7de92e70 != nil {
		var var_e4be7de92e70_mapped *structpb.Value

		var var_e4be7de92e70_err error
		var_e4be7de92e70_mapped, var_e4be7de92e70_err = types.ByResourcePropertyType(model.ResourceProperty_INT32).Pack(*var_e4be7de92e70)
		if var_e4be7de92e70_err != nil {
			panic(var_e4be7de92e70_err)
		}
		properties["limit"] = var_e4be7de92e70_mapped
	}

	var_b07ec079dab8 := extensionRecordSearchParams.Offset

	if var_b07ec079dab8 != nil {
		var var_b07ec079dab8_mapped *structpb.Value

		var var_b07ec079dab8_err error
		var_b07ec079dab8_mapped, var_b07ec079dab8_err = types.ByResourcePropertyType(model.ResourceProperty_INT32).Pack(*var_b07ec079dab8)
		if var_b07ec079dab8_err != nil {
			panic(var_b07ec079dab8_err)
		}
		properties["offset"] = var_b07ec079dab8_mapped
	}

	var_95f40fbb9396 := extensionRecordSearchParams.ResolveReferences

	if var_95f40fbb9396 != nil {
		var var_95f40fbb9396_mapped *structpb.Value

		var var_95f40fbb9396_l []*structpb.Value
		for _, value := range var_95f40fbb9396 {

			var_f010bea6f03d := value
			var var_f010bea6f03d_mapped *structpb.Value

			var var_f010bea6f03d_err error
			var_f010bea6f03d_mapped, var_f010bea6f03d_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_f010bea6f03d)
			if var_f010bea6f03d_err != nil {
				panic(var_f010bea6f03d_err)
			}

			var_95f40fbb9396_l = append(var_95f40fbb9396_l, var_f010bea6f03d_mapped)
		}
		var_95f40fbb9396_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_95f40fbb9396_l})
		properties["resolveReferences"] = var_95f40fbb9396_mapped
	}
	return properties
}

func (m *ExtensionRecordSearchParamsMapper) FromProperties(properties map[string]*structpb.Value) *ExtensionRecordSearchParams {
	var s = m.New()
	if properties["query"] != nil {

		var_b96964b449ef := properties["query"]
		var mappedValue = ExtensionBooleanExpressionMapperInstance.FromProperties(var_b96964b449ef.GetStructValue().Fields)

		var_b96964b449ef_mapped := mappedValue

		s.Query = var_b96964b449ef_mapped
	}
	if properties["limit"] != nil {

		var_da2dbed10466 := properties["limit"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(var_da2dbed10466)

		if err != nil {
			panic(err)
		}

		var_da2dbed10466_mapped := new(int32)
		*var_da2dbed10466_mapped = val.(int32)

		s.Limit = var_da2dbed10466_mapped
	}
	if properties["offset"] != nil {

		var_2f1a9fc5f067 := properties["offset"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(var_2f1a9fc5f067)

		if err != nil {
			panic(err)
		}

		var_2f1a9fc5f067_mapped := new(int32)
		*var_2f1a9fc5f067_mapped = val.(int32)

		s.Offset = var_2f1a9fc5f067_mapped
	}
	if properties["resolveReferences"] != nil {

		var_92b4e3ac233c := properties["resolveReferences"]
		var_92b4e3ac233c_mapped := []string{}
		for _, v := range var_92b4e3ac233c.GetListValue().Values {

			var_281ee5f9b53d := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_281ee5f9b53d)

			if err != nil {
				panic(err)
			}

			var_281ee5f9b53d_mapped := val.(string)

			var_92b4e3ac233c_mapped = append(var_92b4e3ac233c_mapped, var_281ee5f9b53d_mapped)
		}

		s.ResolveReferences = var_92b4e3ac233c_mapped
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

func (m *ExtensionEventMapper) ToProperties(extensionEvent *ExtensionEvent) map[string]*structpb.Value {
	var properties = make(map[string]*structpb.Value)

	var_36fefa3bca31 := extensionEvent.Id

	if var_36fefa3bca31 != nil {
		var var_36fefa3bca31_mapped *structpb.Value

		var var_36fefa3bca31_err error
		var_36fefa3bca31_mapped, var_36fefa3bca31_err = types.ByResourcePropertyType(model.ResourceProperty_UUID).Pack(*var_36fefa3bca31)
		if var_36fefa3bca31_err != nil {
			panic(var_36fefa3bca31_err)
		}
		properties["id"] = var_36fefa3bca31_mapped
	}

	var_db3e15233265 := extensionEvent.Action

	var var_db3e15233265_mapped *structpb.Value

	var var_db3e15233265_err error
	var_db3e15233265_mapped, var_db3e15233265_err = types.ByResourcePropertyType(model.ResourceProperty_ENUM).Pack(string(var_db3e15233265))
	if var_db3e15233265_err != nil {
		panic(var_db3e15233265_err)
	}
	properties["action"] = var_db3e15233265_mapped

	var_f5047a0b9397 := extensionEvent.RecordSearchParams

	if var_f5047a0b9397 != nil {
		var var_f5047a0b9397_mapped *structpb.Value

		var_f5047a0b9397_mapped = structpb.NewStructValue(&structpb.Struct{Fields: ExtensionRecordSearchParamsMapperInstance.ToProperties(var_f5047a0b9397)})
		properties["recordSearchParams"] = var_f5047a0b9397_mapped
	}

	var_138d7bd08f53 := extensionEvent.ActionSummary

	if var_138d7bd08f53 != nil {
		var var_138d7bd08f53_mapped *structpb.Value

		var var_138d7bd08f53_err error
		var_138d7bd08f53_mapped, var_138d7bd08f53_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_138d7bd08f53)
		if var_138d7bd08f53_err != nil {
			panic(var_138d7bd08f53_err)
		}
		properties["actionSummary"] = var_138d7bd08f53_mapped
	}

	var_261b6a38ea5d := extensionEvent.ActionDescription

	if var_261b6a38ea5d != nil {
		var var_261b6a38ea5d_mapped *structpb.Value

		var var_261b6a38ea5d_err error
		var_261b6a38ea5d_mapped, var_261b6a38ea5d_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_261b6a38ea5d)
		if var_261b6a38ea5d_err != nil {
			panic(var_261b6a38ea5d_err)
		}
		properties["actionDescription"] = var_261b6a38ea5d_mapped
	}

	var_0d0a519b71d4 := extensionEvent.Resource

	if var_0d0a519b71d4 != nil {
		var var_0d0a519b71d4_mapped *structpb.Value

		var_0d0a519b71d4_mapped = structpb.NewStructValue(&structpb.Struct{Fields: ResourceMapperInstance.ToProperties(var_0d0a519b71d4)})
		properties["resource"] = var_0d0a519b71d4_mapped
	}

	var_0b7333064f3a := extensionEvent.Records

	if var_0b7333064f3a != nil {
		var var_0b7333064f3a_mapped *structpb.Value

		var var_0b7333064f3a_l []*structpb.Value
		for _, value := range var_0b7333064f3a {

			var_25a6d17a12b2 := value
			var var_25a6d17a12b2_mapped *structpb.Value

			var_25a6d17a12b2_mapped = structpb.NewStructValue(&structpb.Struct{Fields: RecordMapperInstance.ToProperties(var_25a6d17a12b2)})

			var_0b7333064f3a_l = append(var_0b7333064f3a_l, var_25a6d17a12b2_mapped)
		}
		var_0b7333064f3a_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_0b7333064f3a_l})
		properties["records"] = var_0b7333064f3a_mapped
	}

	var_54e23ce0c3a7 := extensionEvent.Ids

	if var_54e23ce0c3a7 != nil {
		var var_54e23ce0c3a7_mapped *structpb.Value

		var var_54e23ce0c3a7_l []*structpb.Value
		for _, value := range var_54e23ce0c3a7 {

			var_3b8ea261aa66 := value
			var var_3b8ea261aa66_mapped *structpb.Value

			var var_3b8ea261aa66_err error
			var_3b8ea261aa66_mapped, var_3b8ea261aa66_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_3b8ea261aa66)
			if var_3b8ea261aa66_err != nil {
				panic(var_3b8ea261aa66_err)
			}

			var_54e23ce0c3a7_l = append(var_54e23ce0c3a7_l, var_3b8ea261aa66_mapped)
		}
		var_54e23ce0c3a7_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_54e23ce0c3a7_l})
		properties["ids"] = var_54e23ce0c3a7_mapped
	}

	var_77f65c36984d := extensionEvent.Finalizes

	if var_77f65c36984d != nil {
		var var_77f65c36984d_mapped *structpb.Value

		var var_77f65c36984d_err error
		var_77f65c36984d_mapped, var_77f65c36984d_err = types.ByResourcePropertyType(model.ResourceProperty_BOOL).Pack(*var_77f65c36984d)
		if var_77f65c36984d_err != nil {
			panic(var_77f65c36984d_err)
		}
		properties["finalizes"] = var_77f65c36984d_mapped
	}

	var_d17bca79caa7 := extensionEvent.Sync

	if var_d17bca79caa7 != nil {
		var var_d17bca79caa7_mapped *structpb.Value

		var var_d17bca79caa7_err error
		var_d17bca79caa7_mapped, var_d17bca79caa7_err = types.ByResourcePropertyType(model.ResourceProperty_BOOL).Pack(*var_d17bca79caa7)
		if var_d17bca79caa7_err != nil {
			panic(var_d17bca79caa7_err)
		}
		properties["sync"] = var_d17bca79caa7_mapped
	}

	var_c30364ed8775 := extensionEvent.Time

	if var_c30364ed8775 != nil {
		var var_c30364ed8775_mapped *structpb.Value

		var var_c30364ed8775_err error
		var_c30364ed8775_mapped, var_c30364ed8775_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_c30364ed8775)
		if var_c30364ed8775_err != nil {
			panic(var_c30364ed8775_err)
		}
		properties["time"] = var_c30364ed8775_mapped
	}

	var_88365b84ecd4 := extensionEvent.Annotations

	if var_88365b84ecd4 != nil {
		var var_88365b84ecd4_mapped *structpb.Value

		var var_88365b84ecd4_st *structpb.Struct = new(structpb.Struct)
		var_88365b84ecd4_st.Fields = make(map[string]*structpb.Value)
		for key, value := range var_88365b84ecd4 {

			var_2bfbdf10823e := value
			var var_2bfbdf10823e_mapped *structpb.Value

			var var_2bfbdf10823e_err error
			var_2bfbdf10823e_mapped, var_2bfbdf10823e_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_2bfbdf10823e)
			if var_2bfbdf10823e_err != nil {
				panic(var_2bfbdf10823e_err)
			}

			var_88365b84ecd4_st.Fields[key] = var_2bfbdf10823e_mapped
		}
		var_88365b84ecd4_mapped = structpb.NewStructValue(var_88365b84ecd4_st)
		properties["annotations"] = var_88365b84ecd4_mapped
	}
	return properties
}

func (m *ExtensionEventMapper) FromProperties(properties map[string]*structpb.Value) *ExtensionEvent {
	var s = m.New()
	if properties["id"] != nil {

		var_1b5b97b52c86 := properties["id"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_UUID).UnPack(var_1b5b97b52c86)

		if err != nil {
			panic(err)
		}

		var_1b5b97b52c86_mapped := new(uuid.UUID)
		*var_1b5b97b52c86_mapped = val.(uuid.UUID)

		s.Id = var_1b5b97b52c86_mapped
	}
	if properties["action"] != nil {

		var_259e1f8e6263 := properties["action"]
		var_259e1f8e6263_mapped := (EventAction)(var_259e1f8e6263.GetStringValue())

		s.Action = var_259e1f8e6263_mapped
	}
	if properties["recordSearchParams"] != nil {

		var_fbf8a730d93b := properties["recordSearchParams"]
		var mappedValue = ExtensionRecordSearchParamsMapperInstance.FromProperties(var_fbf8a730d93b.GetStructValue().Fields)

		var_fbf8a730d93b_mapped := mappedValue

		s.RecordSearchParams = var_fbf8a730d93b_mapped
	}
	if properties["actionSummary"] != nil {

		var_cb2b6fdbb12c := properties["actionSummary"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_cb2b6fdbb12c)

		if err != nil {
			panic(err)
		}

		var_cb2b6fdbb12c_mapped := new(string)
		*var_cb2b6fdbb12c_mapped = val.(string)

		s.ActionSummary = var_cb2b6fdbb12c_mapped
	}
	if properties["actionDescription"] != nil {

		var_bdd1c0d29d70 := properties["actionDescription"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_bdd1c0d29d70)

		if err != nil {
			panic(err)
		}

		var_bdd1c0d29d70_mapped := new(string)
		*var_bdd1c0d29d70_mapped = val.(string)

		s.ActionDescription = var_bdd1c0d29d70_mapped
	}
	if properties["resource"] != nil {

		var_26684d504fe7 := properties["resource"]
		var_26684d504fe7_mapped := ResourceMapperInstance.FromProperties(var_26684d504fe7.GetStructValue().Fields)

		s.Resource = var_26684d504fe7_mapped
	}
	if properties["records"] != nil {

		var_9b39daf0424e := properties["records"]
		var_9b39daf0424e_mapped := []*Record{}
		for _, v := range var_9b39daf0424e.GetListValue().Values {

			var_f219c295b162 := v
			var_f219c295b162_mapped := RecordMapperInstance.FromProperties(var_f219c295b162.GetStructValue().Fields)

			var_9b39daf0424e_mapped = append(var_9b39daf0424e_mapped, var_f219c295b162_mapped)
		}

		s.Records = var_9b39daf0424e_mapped
	}
	if properties["ids"] != nil {

		var_cdfa70a85647 := properties["ids"]
		var_cdfa70a85647_mapped := []string{}
		for _, v := range var_cdfa70a85647.GetListValue().Values {

			var_aee3be584c06 := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_aee3be584c06)

			if err != nil {
				panic(err)
			}

			var_aee3be584c06_mapped := val.(string)

			var_cdfa70a85647_mapped = append(var_cdfa70a85647_mapped, var_aee3be584c06_mapped)
		}

		s.Ids = var_cdfa70a85647_mapped
	}
	if properties["finalizes"] != nil {

		var_9ce60c9eca7c := properties["finalizes"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_BOOL).UnPack(var_9ce60c9eca7c)

		if err != nil {
			panic(err)
		}

		var_9ce60c9eca7c_mapped := new(bool)
		*var_9ce60c9eca7c_mapped = val.(bool)

		s.Finalizes = var_9ce60c9eca7c_mapped
	}
	if properties["sync"] != nil {

		var_e6b1b53fe3b0 := properties["sync"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_BOOL).UnPack(var_e6b1b53fe3b0)

		if err != nil {
			panic(err)
		}

		var_e6b1b53fe3b0_mapped := new(bool)
		*var_e6b1b53fe3b0_mapped = val.(bool)

		s.Sync = var_e6b1b53fe3b0_mapped
	}
	if properties["time"] != nil {

		var_0151138a2a3f := properties["time"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_0151138a2a3f)

		if err != nil {
			panic(err)
		}

		var_0151138a2a3f_mapped := new(time.Time)
		*var_0151138a2a3f_mapped = val.(time.Time)

		s.Time = var_0151138a2a3f_mapped
	}
	if properties["annotations"] != nil {

		var_42326fc9a20f := properties["annotations"]
		var_42326fc9a20f_mapped := make(map[string]string)
		for k, v := range var_42326fc9a20f.GetStructValue().Fields {

			var_fbf88641dedd := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_fbf88641dedd)

			if err != nil {
				panic(err)
			}

			var_fbf88641dedd_mapped := val.(string)

			var_42326fc9a20f_mapped[k] = var_fbf88641dedd_mapped
		}

		s.Annotations = var_42326fc9a20f_mapped
	}
	return s
}
