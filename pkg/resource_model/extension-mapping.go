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

	var_86aae22bfb63 := extension.Id

	if var_86aae22bfb63 != nil {
		var var_86aae22bfb63_mapped *structpb.Value

		var var_86aae22bfb63_err error
		var_86aae22bfb63_mapped, var_86aae22bfb63_err = types.ByResourcePropertyType(model.ResourceProperty_UUID).Pack(*var_86aae22bfb63)
		if var_86aae22bfb63_err != nil {
			panic(var_86aae22bfb63_err)
		}
		properties["id"] = var_86aae22bfb63_mapped
	}

	var_117095481a0f := extension.Version

	var var_117095481a0f_mapped *structpb.Value

	var var_117095481a0f_err error
	var_117095481a0f_mapped, var_117095481a0f_err = types.ByResourcePropertyType(model.ResourceProperty_INT32).Pack(var_117095481a0f)
	if var_117095481a0f_err != nil {
		panic(var_117095481a0f_err)
	}
	properties["version"] = var_117095481a0f_mapped

	var_14fb5d4fa7e1 := extension.CreatedBy

	if var_14fb5d4fa7e1 != nil {
		var var_14fb5d4fa7e1_mapped *structpb.Value

		var var_14fb5d4fa7e1_err error
		var_14fb5d4fa7e1_mapped, var_14fb5d4fa7e1_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_14fb5d4fa7e1)
		if var_14fb5d4fa7e1_err != nil {
			panic(var_14fb5d4fa7e1_err)
		}
		properties["createdBy"] = var_14fb5d4fa7e1_mapped
	}

	var_223fde541f3b := extension.UpdatedBy

	if var_223fde541f3b != nil {
		var var_223fde541f3b_mapped *structpb.Value

		var var_223fde541f3b_err error
		var_223fde541f3b_mapped, var_223fde541f3b_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_223fde541f3b)
		if var_223fde541f3b_err != nil {
			panic(var_223fde541f3b_err)
		}
		properties["updatedBy"] = var_223fde541f3b_mapped
	}

	var_7db0d10da6b6 := extension.CreatedOn

	if var_7db0d10da6b6 != nil {
		var var_7db0d10da6b6_mapped *structpb.Value

		var var_7db0d10da6b6_err error
		var_7db0d10da6b6_mapped, var_7db0d10da6b6_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_7db0d10da6b6)
		if var_7db0d10da6b6_err != nil {
			panic(var_7db0d10da6b6_err)
		}
		properties["createdOn"] = var_7db0d10da6b6_mapped
	}

	var_602e6092bc9f := extension.UpdatedOn

	if var_602e6092bc9f != nil {
		var var_602e6092bc9f_mapped *structpb.Value

		var var_602e6092bc9f_err error
		var_602e6092bc9f_mapped, var_602e6092bc9f_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_602e6092bc9f)
		if var_602e6092bc9f_err != nil {
			panic(var_602e6092bc9f_err)
		}
		properties["updatedOn"] = var_602e6092bc9f_mapped
	}

	var_1c3b05a22f5d := extension.Name

	var var_1c3b05a22f5d_mapped *structpb.Value

	var var_1c3b05a22f5d_err error
	var_1c3b05a22f5d_mapped, var_1c3b05a22f5d_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_1c3b05a22f5d)
	if var_1c3b05a22f5d_err != nil {
		panic(var_1c3b05a22f5d_err)
	}
	properties["name"] = var_1c3b05a22f5d_mapped

	var_c7fb7e0d688f := extension.Description

	if var_c7fb7e0d688f != nil {
		var var_c7fb7e0d688f_mapped *structpb.Value

		var var_c7fb7e0d688f_err error
		var_c7fb7e0d688f_mapped, var_c7fb7e0d688f_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_c7fb7e0d688f)
		if var_c7fb7e0d688f_err != nil {
			panic(var_c7fb7e0d688f_err)
		}
		properties["description"] = var_c7fb7e0d688f_mapped
	}

	var_f0b876226688 := extension.Selector

	if var_f0b876226688 != nil {
		var var_f0b876226688_mapped *structpb.Value

		var_f0b876226688_mapped = structpb.NewStructValue(&structpb.Struct{Fields: ExtensionEventSelectorMapperInstance.ToProperties(var_f0b876226688)})
		properties["selector"] = var_f0b876226688_mapped
	}

	var_9d5a372789e0 := extension.Order

	var var_9d5a372789e0_mapped *structpb.Value

	var var_9d5a372789e0_err error
	var_9d5a372789e0_mapped, var_9d5a372789e0_err = types.ByResourcePropertyType(model.ResourceProperty_INT32).Pack(var_9d5a372789e0)
	if var_9d5a372789e0_err != nil {
		panic(var_9d5a372789e0_err)
	}
	properties["order"] = var_9d5a372789e0_mapped

	var_0b9c98890dcd := extension.Finalizes

	var var_0b9c98890dcd_mapped *structpb.Value

	var var_0b9c98890dcd_err error
	var_0b9c98890dcd_mapped, var_0b9c98890dcd_err = types.ByResourcePropertyType(model.ResourceProperty_BOOL).Pack(var_0b9c98890dcd)
	if var_0b9c98890dcd_err != nil {
		panic(var_0b9c98890dcd_err)
	}
	properties["finalizes"] = var_0b9c98890dcd_mapped

	var_cacb3bc31959 := extension.Sync

	var var_cacb3bc31959_mapped *structpb.Value

	var var_cacb3bc31959_err error
	var_cacb3bc31959_mapped, var_cacb3bc31959_err = types.ByResourcePropertyType(model.ResourceProperty_BOOL).Pack(var_cacb3bc31959)
	if var_cacb3bc31959_err != nil {
		panic(var_cacb3bc31959_err)
	}
	properties["sync"] = var_cacb3bc31959_mapped

	var_a36b6cd13480 := extension.Responds

	var var_a36b6cd13480_mapped *structpb.Value

	var var_a36b6cd13480_err error
	var_a36b6cd13480_mapped, var_a36b6cd13480_err = types.ByResourcePropertyType(model.ResourceProperty_BOOL).Pack(var_a36b6cd13480)
	if var_a36b6cd13480_err != nil {
		panic(var_a36b6cd13480_err)
	}
	properties["responds"] = var_a36b6cd13480_mapped

	var_cbcf8a933ad9 := extension.Call

	var var_cbcf8a933ad9_mapped *structpb.Value

	var_cbcf8a933ad9_mapped = structpb.NewStructValue(&structpb.Struct{Fields: ExtensionExternalCallMapperInstance.ToProperties(&var_cbcf8a933ad9)})
	properties["call"] = var_cbcf8a933ad9_mapped

	var_e1f264d17ad6 := extension.Annotations

	if var_e1f264d17ad6 != nil {
		var var_e1f264d17ad6_mapped *structpb.Value

		var var_e1f264d17ad6_st *structpb.Struct = new(structpb.Struct)
		var_e1f264d17ad6_st.Fields = make(map[string]*structpb.Value)
		for key, value := range var_e1f264d17ad6 {

			var_cc34d627a627 := value
			var var_cc34d627a627_mapped *structpb.Value

			var var_cc34d627a627_err error
			var_cc34d627a627_mapped, var_cc34d627a627_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_cc34d627a627)
			if var_cc34d627a627_err != nil {
				panic(var_cc34d627a627_err)
			}

			var_e1f264d17ad6_st.Fields[key] = var_cc34d627a627_mapped
		}
		var_e1f264d17ad6_mapped = structpb.NewStructValue(var_e1f264d17ad6_st)
		properties["annotations"] = var_e1f264d17ad6_mapped
	}
	return properties
}

func (m *ExtensionMapper) FromProperties(properties map[string]*structpb.Value) *Extension {
	var s = m.New()
	if properties["id"] != nil {

		var_f991192ce994 := properties["id"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_UUID).UnPack(var_f991192ce994)

		if err != nil {
			panic(err)
		}

		var_f991192ce994_mapped := new(uuid.UUID)
		*var_f991192ce994_mapped = val.(uuid.UUID)

		s.Id = var_f991192ce994_mapped
	}
	if properties["version"] != nil {

		var_b53da1955348 := properties["version"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(var_b53da1955348)

		if err != nil {
			panic(err)
		}

		var_b53da1955348_mapped := val.(int32)

		s.Version = var_b53da1955348_mapped
	}
	if properties["createdBy"] != nil {

		var_63c7dbcdd04d := properties["createdBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_63c7dbcdd04d)

		if err != nil {
			panic(err)
		}

		var_63c7dbcdd04d_mapped := new(string)
		*var_63c7dbcdd04d_mapped = val.(string)

		s.CreatedBy = var_63c7dbcdd04d_mapped
	}
	if properties["updatedBy"] != nil {

		var_8b5cc60dbbd2 := properties["updatedBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_8b5cc60dbbd2)

		if err != nil {
			panic(err)
		}

		var_8b5cc60dbbd2_mapped := new(string)
		*var_8b5cc60dbbd2_mapped = val.(string)

		s.UpdatedBy = var_8b5cc60dbbd2_mapped
	}
	if properties["createdOn"] != nil {

		var_e5340a1ddcff := properties["createdOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_e5340a1ddcff)

		if err != nil {
			panic(err)
		}

		var_e5340a1ddcff_mapped := new(time.Time)
		*var_e5340a1ddcff_mapped = val.(time.Time)

		s.CreatedOn = var_e5340a1ddcff_mapped
	}
	if properties["updatedOn"] != nil {

		var_1aebae2e28f0 := properties["updatedOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_1aebae2e28f0)

		if err != nil {
			panic(err)
		}

		var_1aebae2e28f0_mapped := new(time.Time)
		*var_1aebae2e28f0_mapped = val.(time.Time)

		s.UpdatedOn = var_1aebae2e28f0_mapped
	}
	if properties["name"] != nil {

		var_d1242234d954 := properties["name"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_d1242234d954)

		if err != nil {
			panic(err)
		}

		var_d1242234d954_mapped := val.(string)

		s.Name = var_d1242234d954_mapped
	}
	if properties["description"] != nil {

		var_e29bac1ea257 := properties["description"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_e29bac1ea257)

		if err != nil {
			panic(err)
		}

		var_e29bac1ea257_mapped := new(string)
		*var_e29bac1ea257_mapped = val.(string)

		s.Description = var_e29bac1ea257_mapped
	}
	if properties["selector"] != nil {

		var_61b97fdde1f7 := properties["selector"]
		var mappedValue = ExtensionEventSelectorMapperInstance.FromProperties(var_61b97fdde1f7.GetStructValue().Fields)

		var_61b97fdde1f7_mapped := mappedValue

		s.Selector = var_61b97fdde1f7_mapped
	}
	if properties["order"] != nil {

		var_a739d91f6d06 := properties["order"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(var_a739d91f6d06)

		if err != nil {
			panic(err)
		}

		var_a739d91f6d06_mapped := val.(int32)

		s.Order = var_a739d91f6d06_mapped
	}
	if properties["finalizes"] != nil {

		var_cae4617efd35 := properties["finalizes"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_BOOL).UnPack(var_cae4617efd35)

		if err != nil {
			panic(err)
		}

		var_cae4617efd35_mapped := val.(bool)

		s.Finalizes = var_cae4617efd35_mapped
	}
	if properties["sync"] != nil {

		var_325acd058305 := properties["sync"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_BOOL).UnPack(var_325acd058305)

		if err != nil {
			panic(err)
		}

		var_325acd058305_mapped := val.(bool)

		s.Sync = var_325acd058305_mapped
	}
	if properties["responds"] != nil {

		var_f7399b236e2d := properties["responds"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_BOOL).UnPack(var_f7399b236e2d)

		if err != nil {
			panic(err)
		}

		var_f7399b236e2d_mapped := val.(bool)

		s.Responds = var_f7399b236e2d_mapped
	}
	if properties["call"] != nil {

		var_24841c97e150 := properties["call"]
		var mappedValue = ExtensionExternalCallMapperInstance.FromProperties(var_24841c97e150.GetStructValue().Fields)

		var_24841c97e150_mapped := *mappedValue

		s.Call = var_24841c97e150_mapped
	}
	if properties["annotations"] != nil {

		var_f2e2d9439365 := properties["annotations"]
		var_f2e2d9439365_mapped := make(map[string]string)
		for k, v := range var_f2e2d9439365.GetStructValue().Fields {

			var_8319130ebff1 := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_8319130ebff1)

			if err != nil {
				panic(err)
			}

			var_8319130ebff1_mapped := val.(string)

			var_f2e2d9439365_mapped[k] = var_8319130ebff1_mapped
		}

		s.Annotations = var_f2e2d9439365_mapped
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

	var_16f269d08751 := extensionFunctionCall.Host

	var var_16f269d08751_mapped *structpb.Value

	var var_16f269d08751_err error
	var_16f269d08751_mapped, var_16f269d08751_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_16f269d08751)
	if var_16f269d08751_err != nil {
		panic(var_16f269d08751_err)
	}
	properties["host"] = var_16f269d08751_mapped

	var_5ed32725050f := extensionFunctionCall.FunctionName

	var var_5ed32725050f_mapped *structpb.Value

	var var_5ed32725050f_err error
	var_5ed32725050f_mapped, var_5ed32725050f_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_5ed32725050f)
	if var_5ed32725050f_err != nil {
		panic(var_5ed32725050f_err)
	}
	properties["functionName"] = var_5ed32725050f_mapped
	return properties
}

func (m *ExtensionFunctionCallMapper) FromProperties(properties map[string]*structpb.Value) *ExtensionFunctionCall {
	var s = m.New()
	if properties["host"] != nil {

		var_f5037588ce68 := properties["host"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_f5037588ce68)

		if err != nil {
			panic(err)
		}

		var_f5037588ce68_mapped := val.(string)

		s.Host = var_f5037588ce68_mapped
	}
	if properties["functionName"] != nil {

		var_5d8f71497295 := properties["functionName"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_5d8f71497295)

		if err != nil {
			panic(err)
		}

		var_5d8f71497295_mapped := val.(string)

		s.FunctionName = var_5d8f71497295_mapped
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

	var_c863d3901224 := extensionHttpCall.Uri

	var var_c863d3901224_mapped *structpb.Value

	var var_c863d3901224_err error
	var_c863d3901224_mapped, var_c863d3901224_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_c863d3901224)
	if var_c863d3901224_err != nil {
		panic(var_c863d3901224_err)
	}
	properties["uri"] = var_c863d3901224_mapped

	var_c7373c2fb2df := extensionHttpCall.Method

	var var_c7373c2fb2df_mapped *structpb.Value

	var var_c7373c2fb2df_err error
	var_c7373c2fb2df_mapped, var_c7373c2fb2df_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_c7373c2fb2df)
	if var_c7373c2fb2df_err != nil {
		panic(var_c7373c2fb2df_err)
	}
	properties["method"] = var_c7373c2fb2df_mapped
	return properties
}

func (m *ExtensionHttpCallMapper) FromProperties(properties map[string]*structpb.Value) *ExtensionHttpCall {
	var s = m.New()
	if properties["uri"] != nil {

		var_ee89b9f36e1a := properties["uri"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_ee89b9f36e1a)

		if err != nil {
			panic(err)
		}

		var_ee89b9f36e1a_mapped := val.(string)

		s.Uri = var_ee89b9f36e1a_mapped
	}
	if properties["method"] != nil {

		var_b21eb6abf872 := properties["method"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_b21eb6abf872)

		if err != nil {
			panic(err)
		}

		var_b21eb6abf872_mapped := val.(string)

		s.Method = var_b21eb6abf872_mapped
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

	var_880f7cfce688 := extensionExternalCall.FunctionCall

	if var_880f7cfce688 != nil {
		var var_880f7cfce688_mapped *structpb.Value

		var_880f7cfce688_mapped = structpb.NewStructValue(&structpb.Struct{Fields: ExtensionFunctionCallMapperInstance.ToProperties(var_880f7cfce688)})
		properties["functionCall"] = var_880f7cfce688_mapped
	}

	var_611cb4aa89b8 := extensionExternalCall.HttpCall

	if var_611cb4aa89b8 != nil {
		var var_611cb4aa89b8_mapped *structpb.Value

		var_611cb4aa89b8_mapped = structpb.NewStructValue(&structpb.Struct{Fields: ExtensionHttpCallMapperInstance.ToProperties(var_611cb4aa89b8)})
		properties["httpCall"] = var_611cb4aa89b8_mapped
	}
	return properties
}

func (m *ExtensionExternalCallMapper) FromProperties(properties map[string]*structpb.Value) *ExtensionExternalCall {
	var s = m.New()
	if properties["functionCall"] != nil {

		var_3d06b57eaafd := properties["functionCall"]
		var mappedValue = ExtensionFunctionCallMapperInstance.FromProperties(var_3d06b57eaafd.GetStructValue().Fields)

		var_3d06b57eaafd_mapped := mappedValue

		s.FunctionCall = var_3d06b57eaafd_mapped
	}
	if properties["httpCall"] != nil {

		var_f2db26e0a7af := properties["httpCall"]
		var mappedValue = ExtensionHttpCallMapperInstance.FromProperties(var_f2db26e0a7af.GetStructValue().Fields)

		var_f2db26e0a7af_mapped := mappedValue

		s.HttpCall = var_f2db26e0a7af_mapped
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

	var_9846721977be := extensionEventSelector.Actions

	if var_9846721977be != nil {
		var var_9846721977be_mapped *structpb.Value

		var var_9846721977be_l []*structpb.Value
		for _, value := range var_9846721977be {

			var_2ddc49284074 := value
			var var_2ddc49284074_mapped *structpb.Value

			var var_2ddc49284074_err error
			var_2ddc49284074_mapped, var_2ddc49284074_err = types.ByResourcePropertyType(model.ResourceProperty_ENUM).Pack(string(var_2ddc49284074))
			if var_2ddc49284074_err != nil {
				panic(var_2ddc49284074_err)
			}

			var_9846721977be_l = append(var_9846721977be_l, var_2ddc49284074_mapped)
		}
		var_9846721977be_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_9846721977be_l})
		properties["actions"] = var_9846721977be_mapped
	}

	var_52d624f0023f := extensionEventSelector.RecordSelector

	if var_52d624f0023f != nil {
		var var_52d624f0023f_mapped *structpb.Value

		var_52d624f0023f_mapped = structpb.NewStructValue(&structpb.Struct{Fields: ExtensionBooleanExpressionMapperInstance.ToProperties(var_52d624f0023f)})
		properties["recordSelector"] = var_52d624f0023f_mapped
	}

	var_acd3f55ad3c1 := extensionEventSelector.Namespaces

	if var_acd3f55ad3c1 != nil {
		var var_acd3f55ad3c1_mapped *structpb.Value

		var var_acd3f55ad3c1_l []*structpb.Value
		for _, value := range var_acd3f55ad3c1 {

			var_5a12c781abeb := value
			var var_5a12c781abeb_mapped *structpb.Value

			var var_5a12c781abeb_err error
			var_5a12c781abeb_mapped, var_5a12c781abeb_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_5a12c781abeb)
			if var_5a12c781abeb_err != nil {
				panic(var_5a12c781abeb_err)
			}

			var_acd3f55ad3c1_l = append(var_acd3f55ad3c1_l, var_5a12c781abeb_mapped)
		}
		var_acd3f55ad3c1_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_acd3f55ad3c1_l})
		properties["namespaces"] = var_acd3f55ad3c1_mapped
	}

	var_dcdc212f4a35 := extensionEventSelector.Resources

	if var_dcdc212f4a35 != nil {
		var var_dcdc212f4a35_mapped *structpb.Value

		var var_dcdc212f4a35_l []*structpb.Value
		for _, value := range var_dcdc212f4a35 {

			var_4d2d3be417de := value
			var var_4d2d3be417de_mapped *structpb.Value

			var var_4d2d3be417de_err error
			var_4d2d3be417de_mapped, var_4d2d3be417de_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_4d2d3be417de)
			if var_4d2d3be417de_err != nil {
				panic(var_4d2d3be417de_err)
			}

			var_dcdc212f4a35_l = append(var_dcdc212f4a35_l, var_4d2d3be417de_mapped)
		}
		var_dcdc212f4a35_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_dcdc212f4a35_l})
		properties["resources"] = var_dcdc212f4a35_mapped
	}

	var_cd674a2ebd93 := extensionEventSelector.Ids

	if var_cd674a2ebd93 != nil {
		var var_cd674a2ebd93_mapped *structpb.Value

		var var_cd674a2ebd93_l []*structpb.Value
		for _, value := range var_cd674a2ebd93 {

			var_f6511f3444e7 := value
			var var_f6511f3444e7_mapped *structpb.Value

			var var_f6511f3444e7_err error
			var_f6511f3444e7_mapped, var_f6511f3444e7_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_f6511f3444e7)
			if var_f6511f3444e7_err != nil {
				panic(var_f6511f3444e7_err)
			}

			var_cd674a2ebd93_l = append(var_cd674a2ebd93_l, var_f6511f3444e7_mapped)
		}
		var_cd674a2ebd93_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_cd674a2ebd93_l})
		properties["ids"] = var_cd674a2ebd93_mapped
	}

	var_47ab1aaa5da7 := extensionEventSelector.Annotations

	if var_47ab1aaa5da7 != nil {
		var var_47ab1aaa5da7_mapped *structpb.Value

		var var_47ab1aaa5da7_st *structpb.Struct = new(structpb.Struct)
		var_47ab1aaa5da7_st.Fields = make(map[string]*structpb.Value)
		for key, value := range var_47ab1aaa5da7 {

			var_e2ef6273f43c := value
			var var_e2ef6273f43c_mapped *structpb.Value

			var var_e2ef6273f43c_err error
			var_e2ef6273f43c_mapped, var_e2ef6273f43c_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_e2ef6273f43c)
			if var_e2ef6273f43c_err != nil {
				panic(var_e2ef6273f43c_err)
			}

			var_47ab1aaa5da7_st.Fields[key] = var_e2ef6273f43c_mapped
		}
		var_47ab1aaa5da7_mapped = structpb.NewStructValue(var_47ab1aaa5da7_st)
		properties["annotations"] = var_47ab1aaa5da7_mapped
	}
	return properties
}

func (m *ExtensionEventSelectorMapper) FromProperties(properties map[string]*structpb.Value) *ExtensionEventSelector {
	var s = m.New()
	if properties["actions"] != nil {

		var_cacdd098c3d7 := properties["actions"]
		var_cacdd098c3d7_mapped := []EventAction{}
		for _, v := range var_cacdd098c3d7.GetListValue().Values {

			var_b458723f549a := v
			var_b458723f549a_mapped := (EventAction)(var_b458723f549a.GetStringValue())

			var_cacdd098c3d7_mapped = append(var_cacdd098c3d7_mapped, var_b458723f549a_mapped)
		}

		s.Actions = var_cacdd098c3d7_mapped
	}
	if properties["recordSelector"] != nil {

		var_d01bcef0ddf7 := properties["recordSelector"]
		var mappedValue = ExtensionBooleanExpressionMapperInstance.FromProperties(var_d01bcef0ddf7.GetStructValue().Fields)

		var_d01bcef0ddf7_mapped := mappedValue

		s.RecordSelector = var_d01bcef0ddf7_mapped
	}
	if properties["namespaces"] != nil {

		var_1f31792c5e26 := properties["namespaces"]
		var_1f31792c5e26_mapped := []string{}
		for _, v := range var_1f31792c5e26.GetListValue().Values {

			var_ac25788767da := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_ac25788767da)

			if err != nil {
				panic(err)
			}

			var_ac25788767da_mapped := val.(string)

			var_1f31792c5e26_mapped = append(var_1f31792c5e26_mapped, var_ac25788767da_mapped)
		}

		s.Namespaces = var_1f31792c5e26_mapped
	}
	if properties["resources"] != nil {

		var_20de38532f5d := properties["resources"]
		var_20de38532f5d_mapped := []string{}
		for _, v := range var_20de38532f5d.GetListValue().Values {

			var_b24e8362c63a := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_b24e8362c63a)

			if err != nil {
				panic(err)
			}

			var_b24e8362c63a_mapped := val.(string)

			var_20de38532f5d_mapped = append(var_20de38532f5d_mapped, var_b24e8362c63a_mapped)
		}

		s.Resources = var_20de38532f5d_mapped
	}
	if properties["ids"] != nil {

		var_a7c91cf1f013 := properties["ids"]
		var_a7c91cf1f013_mapped := []string{}
		for _, v := range var_a7c91cf1f013.GetListValue().Values {

			var_3ea0c68979e5 := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_3ea0c68979e5)

			if err != nil {
				panic(err)
			}

			var_3ea0c68979e5_mapped := val.(string)

			var_a7c91cf1f013_mapped = append(var_a7c91cf1f013_mapped, var_3ea0c68979e5_mapped)
		}

		s.Ids = var_a7c91cf1f013_mapped
	}
	if properties["annotations"] != nil {

		var_14880db19184 := properties["annotations"]
		var_14880db19184_mapped := make(map[string]string)
		for k, v := range var_14880db19184.GetStructValue().Fields {

			var_e5271506f89c := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_e5271506f89c)

			if err != nil {
				panic(err)
			}

			var_e5271506f89c_mapped := val.(string)

			var_14880db19184_mapped[k] = var_e5271506f89c_mapped
		}

		s.Annotations = var_14880db19184_mapped
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

	var_d9ddc5405521 := extensionRecordSearchParams.Query

	if var_d9ddc5405521 != nil {
		var var_d9ddc5405521_mapped *structpb.Value

		var_d9ddc5405521_mapped = structpb.NewStructValue(&structpb.Struct{Fields: ExtensionBooleanExpressionMapperInstance.ToProperties(var_d9ddc5405521)})
		properties["query"] = var_d9ddc5405521_mapped
	}

	var_5a6137c547aa := extensionRecordSearchParams.Limit

	if var_5a6137c547aa != nil {
		var var_5a6137c547aa_mapped *structpb.Value

		var var_5a6137c547aa_err error
		var_5a6137c547aa_mapped, var_5a6137c547aa_err = types.ByResourcePropertyType(model.ResourceProperty_INT32).Pack(*var_5a6137c547aa)
		if var_5a6137c547aa_err != nil {
			panic(var_5a6137c547aa_err)
		}
		properties["limit"] = var_5a6137c547aa_mapped
	}

	var_e42f933cbafa := extensionRecordSearchParams.Offset

	if var_e42f933cbafa != nil {
		var var_e42f933cbafa_mapped *structpb.Value

		var var_e42f933cbafa_err error
		var_e42f933cbafa_mapped, var_e42f933cbafa_err = types.ByResourcePropertyType(model.ResourceProperty_INT32).Pack(*var_e42f933cbafa)
		if var_e42f933cbafa_err != nil {
			panic(var_e42f933cbafa_err)
		}
		properties["offset"] = var_e42f933cbafa_mapped
	}

	var_eb467b0741d4 := extensionRecordSearchParams.ResolveReferences

	if var_eb467b0741d4 != nil {
		var var_eb467b0741d4_mapped *structpb.Value

		var var_eb467b0741d4_l []*structpb.Value
		for _, value := range var_eb467b0741d4 {

			var_126cbb214304 := value
			var var_126cbb214304_mapped *structpb.Value

			var var_126cbb214304_err error
			var_126cbb214304_mapped, var_126cbb214304_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_126cbb214304)
			if var_126cbb214304_err != nil {
				panic(var_126cbb214304_err)
			}

			var_eb467b0741d4_l = append(var_eb467b0741d4_l, var_126cbb214304_mapped)
		}
		var_eb467b0741d4_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_eb467b0741d4_l})
		properties["resolveReferences"] = var_eb467b0741d4_mapped
	}
	return properties
}

func (m *ExtensionRecordSearchParamsMapper) FromProperties(properties map[string]*structpb.Value) *ExtensionRecordSearchParams {
	var s = m.New()
	if properties["query"] != nil {

		var_0a9dd1b0ef29 := properties["query"]
		var mappedValue = ExtensionBooleanExpressionMapperInstance.FromProperties(var_0a9dd1b0ef29.GetStructValue().Fields)

		var_0a9dd1b0ef29_mapped := mappedValue

		s.Query = var_0a9dd1b0ef29_mapped
	}
	if properties["limit"] != nil {

		var_8e884024ea0c := properties["limit"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(var_8e884024ea0c)

		if err != nil {
			panic(err)
		}

		var_8e884024ea0c_mapped := new(int32)
		*var_8e884024ea0c_mapped = val.(int32)

		s.Limit = var_8e884024ea0c_mapped
	}
	if properties["offset"] != nil {

		var_39b5284149bc := properties["offset"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(var_39b5284149bc)

		if err != nil {
			panic(err)
		}

		var_39b5284149bc_mapped := new(int32)
		*var_39b5284149bc_mapped = val.(int32)

		s.Offset = var_39b5284149bc_mapped
	}
	if properties["resolveReferences"] != nil {

		var_dc1d8c35156e := properties["resolveReferences"]
		var_dc1d8c35156e_mapped := []string{}
		for _, v := range var_dc1d8c35156e.GetListValue().Values {

			var_0e0fad4b1196 := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_0e0fad4b1196)

			if err != nil {
				panic(err)
			}

			var_0e0fad4b1196_mapped := val.(string)

			var_dc1d8c35156e_mapped = append(var_dc1d8c35156e_mapped, var_0e0fad4b1196_mapped)
		}

		s.ResolveReferences = var_dc1d8c35156e_mapped
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

	var_5efcb3b50feb := extensionEvent.Id

	if var_5efcb3b50feb != nil {
		var var_5efcb3b50feb_mapped *structpb.Value

		var var_5efcb3b50feb_err error
		var_5efcb3b50feb_mapped, var_5efcb3b50feb_err = types.ByResourcePropertyType(model.ResourceProperty_UUID).Pack(*var_5efcb3b50feb)
		if var_5efcb3b50feb_err != nil {
			panic(var_5efcb3b50feb_err)
		}
		properties["id"] = var_5efcb3b50feb_mapped
	}

	var_7a62f03bdb64 := extensionEvent.Action

	var var_7a62f03bdb64_mapped *structpb.Value

	var var_7a62f03bdb64_err error
	var_7a62f03bdb64_mapped, var_7a62f03bdb64_err = types.ByResourcePropertyType(model.ResourceProperty_ENUM).Pack(string(var_7a62f03bdb64))
	if var_7a62f03bdb64_err != nil {
		panic(var_7a62f03bdb64_err)
	}
	properties["action"] = var_7a62f03bdb64_mapped

	var_1900e8de8a50 := extensionEvent.RecordSearchParams

	if var_1900e8de8a50 != nil {
		var var_1900e8de8a50_mapped *structpb.Value

		var_1900e8de8a50_mapped = structpb.NewStructValue(&structpb.Struct{Fields: ExtensionRecordSearchParamsMapperInstance.ToProperties(var_1900e8de8a50)})
		properties["recordSearchParams"] = var_1900e8de8a50_mapped
	}

	var_c3a4dd163638 := extensionEvent.ActionSummary

	if var_c3a4dd163638 != nil {
		var var_c3a4dd163638_mapped *structpb.Value

		var var_c3a4dd163638_err error
		var_c3a4dd163638_mapped, var_c3a4dd163638_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_c3a4dd163638)
		if var_c3a4dd163638_err != nil {
			panic(var_c3a4dd163638_err)
		}
		properties["actionSummary"] = var_c3a4dd163638_mapped
	}

	var_8d773da4973a := extensionEvent.ActionDescription

	if var_8d773da4973a != nil {
		var var_8d773da4973a_mapped *structpb.Value

		var var_8d773da4973a_err error
		var_8d773da4973a_mapped, var_8d773da4973a_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_8d773da4973a)
		if var_8d773da4973a_err != nil {
			panic(var_8d773da4973a_err)
		}
		properties["actionDescription"] = var_8d773da4973a_mapped
	}

	var_e345bb98dc07 := extensionEvent.Resource

	if var_e345bb98dc07 != nil {
		var var_e345bb98dc07_mapped *structpb.Value

		var_e345bb98dc07_mapped = structpb.NewStructValue(&structpb.Struct{Fields: ResourceMapperInstance.ToProperties(var_e345bb98dc07)})
		properties["resource"] = var_e345bb98dc07_mapped
	}

	var_42c85bcbee30 := extensionEvent.Records

	if var_42c85bcbee30 != nil {
		var var_42c85bcbee30_mapped *structpb.Value

		var var_42c85bcbee30_l []*structpb.Value
		for _, value := range var_42c85bcbee30 {

			var_d34ae6e1d18b := value
			var var_d34ae6e1d18b_mapped *structpb.Value

			var_d34ae6e1d18b_mapped = structpb.NewStructValue(&structpb.Struct{Fields: RecordMapperInstance.ToProperties(var_d34ae6e1d18b)})

			var_42c85bcbee30_l = append(var_42c85bcbee30_l, var_d34ae6e1d18b_mapped)
		}
		var_42c85bcbee30_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_42c85bcbee30_l})
		properties["records"] = var_42c85bcbee30_mapped
	}

	var_eae2c632bdf7 := extensionEvent.Ids

	if var_eae2c632bdf7 != nil {
		var var_eae2c632bdf7_mapped *structpb.Value

		var var_eae2c632bdf7_l []*structpb.Value
		for _, value := range var_eae2c632bdf7 {

			var_bfc1db44281e := value
			var var_bfc1db44281e_mapped *structpb.Value

			var var_bfc1db44281e_err error
			var_bfc1db44281e_mapped, var_bfc1db44281e_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_bfc1db44281e)
			if var_bfc1db44281e_err != nil {
				panic(var_bfc1db44281e_err)
			}

			var_eae2c632bdf7_l = append(var_eae2c632bdf7_l, var_bfc1db44281e_mapped)
		}
		var_eae2c632bdf7_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_eae2c632bdf7_l})
		properties["ids"] = var_eae2c632bdf7_mapped
	}

	var_6111bae67081 := extensionEvent.Finalizes

	if var_6111bae67081 != nil {
		var var_6111bae67081_mapped *structpb.Value

		var var_6111bae67081_err error
		var_6111bae67081_mapped, var_6111bae67081_err = types.ByResourcePropertyType(model.ResourceProperty_BOOL).Pack(*var_6111bae67081)
		if var_6111bae67081_err != nil {
			panic(var_6111bae67081_err)
		}
		properties["finalizes"] = var_6111bae67081_mapped
	}

	var_be2832093c08 := extensionEvent.Sync

	if var_be2832093c08 != nil {
		var var_be2832093c08_mapped *structpb.Value

		var var_be2832093c08_err error
		var_be2832093c08_mapped, var_be2832093c08_err = types.ByResourcePropertyType(model.ResourceProperty_BOOL).Pack(*var_be2832093c08)
		if var_be2832093c08_err != nil {
			panic(var_be2832093c08_err)
		}
		properties["sync"] = var_be2832093c08_mapped
	}

	var_79fb0eaf6192 := extensionEvent.Time

	if var_79fb0eaf6192 != nil {
		var var_79fb0eaf6192_mapped *structpb.Value

		var var_79fb0eaf6192_err error
		var_79fb0eaf6192_mapped, var_79fb0eaf6192_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_79fb0eaf6192)
		if var_79fb0eaf6192_err != nil {
			panic(var_79fb0eaf6192_err)
		}
		properties["time"] = var_79fb0eaf6192_mapped
	}

	var_1211dbb2a259 := extensionEvent.Annotations

	if var_1211dbb2a259 != nil {
		var var_1211dbb2a259_mapped *structpb.Value

		var var_1211dbb2a259_st *structpb.Struct = new(structpb.Struct)
		var_1211dbb2a259_st.Fields = make(map[string]*structpb.Value)
		for key, value := range var_1211dbb2a259 {

			var_a4b186b67671 := value
			var var_a4b186b67671_mapped *structpb.Value

			var var_a4b186b67671_err error
			var_a4b186b67671_mapped, var_a4b186b67671_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_a4b186b67671)
			if var_a4b186b67671_err != nil {
				panic(var_a4b186b67671_err)
			}

			var_1211dbb2a259_st.Fields[key] = var_a4b186b67671_mapped
		}
		var_1211dbb2a259_mapped = structpb.NewStructValue(var_1211dbb2a259_st)
		properties["annotations"] = var_1211dbb2a259_mapped
	}
	return properties
}

func (m *ExtensionEventMapper) FromProperties(properties map[string]*structpb.Value) *ExtensionEvent {
	var s = m.New()
	if properties["id"] != nil {

		var_79e72e9e1d48 := properties["id"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_UUID).UnPack(var_79e72e9e1d48)

		if err != nil {
			panic(err)
		}

		var_79e72e9e1d48_mapped := new(uuid.UUID)
		*var_79e72e9e1d48_mapped = val.(uuid.UUID)

		s.Id = var_79e72e9e1d48_mapped
	}
	if properties["action"] != nil {

		var_1243f0919acd := properties["action"]
		var_1243f0919acd_mapped := (EventAction)(var_1243f0919acd.GetStringValue())

		s.Action = var_1243f0919acd_mapped
	}
	if properties["recordSearchParams"] != nil {

		var_9c708bc54cef := properties["recordSearchParams"]
		var mappedValue = ExtensionRecordSearchParamsMapperInstance.FromProperties(var_9c708bc54cef.GetStructValue().Fields)

		var_9c708bc54cef_mapped := mappedValue

		s.RecordSearchParams = var_9c708bc54cef_mapped
	}
	if properties["actionSummary"] != nil {

		var_5c019b9b019e := properties["actionSummary"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_5c019b9b019e)

		if err != nil {
			panic(err)
		}

		var_5c019b9b019e_mapped := new(string)
		*var_5c019b9b019e_mapped = val.(string)

		s.ActionSummary = var_5c019b9b019e_mapped
	}
	if properties["actionDescription"] != nil {

		var_0f98ec6f3115 := properties["actionDescription"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_0f98ec6f3115)

		if err != nil {
			panic(err)
		}

		var_0f98ec6f3115_mapped := new(string)
		*var_0f98ec6f3115_mapped = val.(string)

		s.ActionDescription = var_0f98ec6f3115_mapped
	}
	if properties["resource"] != nil {

		var_2010afcbf93b := properties["resource"]
		var_2010afcbf93b_mapped := ResourceMapperInstance.FromProperties(var_2010afcbf93b.GetStructValue().Fields)

		s.Resource = var_2010afcbf93b_mapped
	}
	if properties["records"] != nil {

		var_2f7998e51c1b := properties["records"]
		var_2f7998e51c1b_mapped := []*Record{}
		for _, v := range var_2f7998e51c1b.GetListValue().Values {

			var_4059e043511b := v
			var_4059e043511b_mapped := RecordMapperInstance.FromProperties(var_4059e043511b.GetStructValue().Fields)

			var_2f7998e51c1b_mapped = append(var_2f7998e51c1b_mapped, var_4059e043511b_mapped)
		}

		s.Records = var_2f7998e51c1b_mapped
	}
	if properties["ids"] != nil {

		var_f64eff4d2736 := properties["ids"]
		var_f64eff4d2736_mapped := []string{}
		for _, v := range var_f64eff4d2736.GetListValue().Values {

			var_df82b24558f4 := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_df82b24558f4)

			if err != nil {
				panic(err)
			}

			var_df82b24558f4_mapped := val.(string)

			var_f64eff4d2736_mapped = append(var_f64eff4d2736_mapped, var_df82b24558f4_mapped)
		}

		s.Ids = var_f64eff4d2736_mapped
	}
	if properties["finalizes"] != nil {

		var_b077f5519ce4 := properties["finalizes"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_BOOL).UnPack(var_b077f5519ce4)

		if err != nil {
			panic(err)
		}

		var_b077f5519ce4_mapped := new(bool)
		*var_b077f5519ce4_mapped = val.(bool)

		s.Finalizes = var_b077f5519ce4_mapped
	}
	if properties["sync"] != nil {

		var_41c665c25d73 := properties["sync"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_BOOL).UnPack(var_41c665c25d73)

		if err != nil {
			panic(err)
		}

		var_41c665c25d73_mapped := new(bool)
		*var_41c665c25d73_mapped = val.(bool)

		s.Sync = var_41c665c25d73_mapped
	}
	if properties["time"] != nil {

		var_6957eb10c664 := properties["time"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_6957eb10c664)

		if err != nil {
			panic(err)
		}

		var_6957eb10c664_mapped := new(time.Time)
		*var_6957eb10c664_mapped = val.(time.Time)

		s.Time = var_6957eb10c664_mapped
	}
	if properties["annotations"] != nil {

		var_7e1d97454000 := properties["annotations"]
		var_7e1d97454000_mapped := make(map[string]string)
		for k, v := range var_7e1d97454000.GetStructValue().Fields {

			var_b5889cfffbad := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_b5889cfffbad)

			if err != nil {
				panic(err)
			}

			var_b5889cfffbad_mapped := val.(string)

			var_7e1d97454000_mapped[k] = var_b5889cfffbad_mapped
		}

		s.Annotations = var_7e1d97454000_mapped
	}
	return s
}
