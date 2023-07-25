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

	var_76db625398c1 := extension.Id

	if var_76db625398c1 != nil {
		var var_76db625398c1_mapped *structpb.Value

		var var_76db625398c1_err error
		var_76db625398c1_mapped, var_76db625398c1_err = types.ByResourcePropertyType(model.ResourceProperty_UUID).Pack(*var_76db625398c1)
		if var_76db625398c1_err != nil {
			panic(var_76db625398c1_err)
		}
		properties["id"] = var_76db625398c1_mapped
	}

	var_5e096de5236e := extension.Version

	var var_5e096de5236e_mapped *structpb.Value

	var var_5e096de5236e_err error
	var_5e096de5236e_mapped, var_5e096de5236e_err = types.ByResourcePropertyType(model.ResourceProperty_INT32).Pack(var_5e096de5236e)
	if var_5e096de5236e_err != nil {
		panic(var_5e096de5236e_err)
	}
	properties["version"] = var_5e096de5236e_mapped

	var_b84d4a0738b3 := extension.CreatedBy

	if var_b84d4a0738b3 != nil {
		var var_b84d4a0738b3_mapped *structpb.Value

		var var_b84d4a0738b3_err error
		var_b84d4a0738b3_mapped, var_b84d4a0738b3_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_b84d4a0738b3)
		if var_b84d4a0738b3_err != nil {
			panic(var_b84d4a0738b3_err)
		}
		properties["createdBy"] = var_b84d4a0738b3_mapped
	}

	var_7f67cf22718b := extension.UpdatedBy

	if var_7f67cf22718b != nil {
		var var_7f67cf22718b_mapped *structpb.Value

		var var_7f67cf22718b_err error
		var_7f67cf22718b_mapped, var_7f67cf22718b_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_7f67cf22718b)
		if var_7f67cf22718b_err != nil {
			panic(var_7f67cf22718b_err)
		}
		properties["updatedBy"] = var_7f67cf22718b_mapped
	}

	var_6c9fce7063f8 := extension.CreatedOn

	if var_6c9fce7063f8 != nil {
		var var_6c9fce7063f8_mapped *structpb.Value

		var var_6c9fce7063f8_err error
		var_6c9fce7063f8_mapped, var_6c9fce7063f8_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_6c9fce7063f8)
		if var_6c9fce7063f8_err != nil {
			panic(var_6c9fce7063f8_err)
		}
		properties["createdOn"] = var_6c9fce7063f8_mapped
	}

	var_bad33b8f4598 := extension.UpdatedOn

	if var_bad33b8f4598 != nil {
		var var_bad33b8f4598_mapped *structpb.Value

		var var_bad33b8f4598_err error
		var_bad33b8f4598_mapped, var_bad33b8f4598_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_bad33b8f4598)
		if var_bad33b8f4598_err != nil {
			panic(var_bad33b8f4598_err)
		}
		properties["updatedOn"] = var_bad33b8f4598_mapped
	}

	var_e08d29a882ec := extension.Name

	var var_e08d29a882ec_mapped *structpb.Value

	var var_e08d29a882ec_err error
	var_e08d29a882ec_mapped, var_e08d29a882ec_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_e08d29a882ec)
	if var_e08d29a882ec_err != nil {
		panic(var_e08d29a882ec_err)
	}
	properties["name"] = var_e08d29a882ec_mapped

	var_c1a1ef0b93a4 := extension.Description

	if var_c1a1ef0b93a4 != nil {
		var var_c1a1ef0b93a4_mapped *structpb.Value

		var var_c1a1ef0b93a4_err error
		var_c1a1ef0b93a4_mapped, var_c1a1ef0b93a4_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_c1a1ef0b93a4)
		if var_c1a1ef0b93a4_err != nil {
			panic(var_c1a1ef0b93a4_err)
		}
		properties["description"] = var_c1a1ef0b93a4_mapped
	}

	var_37e6d5d16bf6 := extension.Selector

	if var_37e6d5d16bf6 != nil {
		var var_37e6d5d16bf6_mapped *structpb.Value

		var_37e6d5d16bf6_mapped = structpb.NewStructValue(&structpb.Struct{Fields: ExtensionEventSelectorMapperInstance.ToProperties(var_37e6d5d16bf6)})
		properties["selector"] = var_37e6d5d16bf6_mapped
	}

	var_649843442e29 := extension.Order

	var var_649843442e29_mapped *structpb.Value

	var var_649843442e29_err error
	var_649843442e29_mapped, var_649843442e29_err = types.ByResourcePropertyType(model.ResourceProperty_INT32).Pack(var_649843442e29)
	if var_649843442e29_err != nil {
		panic(var_649843442e29_err)
	}
	properties["order"] = var_649843442e29_mapped

	var_e3b3f10b8853 := extension.Finalizes

	var var_e3b3f10b8853_mapped *structpb.Value

	var var_e3b3f10b8853_err error
	var_e3b3f10b8853_mapped, var_e3b3f10b8853_err = types.ByResourcePropertyType(model.ResourceProperty_BOOL).Pack(var_e3b3f10b8853)
	if var_e3b3f10b8853_err != nil {
		panic(var_e3b3f10b8853_err)
	}
	properties["finalizes"] = var_e3b3f10b8853_mapped

	var_42e84421e58a := extension.Sync

	var var_42e84421e58a_mapped *structpb.Value

	var var_42e84421e58a_err error
	var_42e84421e58a_mapped, var_42e84421e58a_err = types.ByResourcePropertyType(model.ResourceProperty_BOOL).Pack(var_42e84421e58a)
	if var_42e84421e58a_err != nil {
		panic(var_42e84421e58a_err)
	}
	properties["sync"] = var_42e84421e58a_mapped

	var_b34e4006c487 := extension.Responds

	var var_b34e4006c487_mapped *structpb.Value

	var var_b34e4006c487_err error
	var_b34e4006c487_mapped, var_b34e4006c487_err = types.ByResourcePropertyType(model.ResourceProperty_BOOL).Pack(var_b34e4006c487)
	if var_b34e4006c487_err != nil {
		panic(var_b34e4006c487_err)
	}
	properties["responds"] = var_b34e4006c487_mapped

	var_b782b800f1fd := extension.Call

	var var_b782b800f1fd_mapped *structpb.Value

	var_b782b800f1fd_mapped = structpb.NewStructValue(&structpb.Struct{Fields: ExtensionExternalCallMapperInstance.ToProperties(&var_b782b800f1fd)})
	properties["call"] = var_b782b800f1fd_mapped

	var_1c6164b998d1 := extension.Annotations

	if var_1c6164b998d1 != nil {
		var var_1c6164b998d1_mapped *structpb.Value

		var var_1c6164b998d1_st *structpb.Struct = new(structpb.Struct)
		var_1c6164b998d1_st.Fields = make(map[string]*structpb.Value)
		for key, value := range var_1c6164b998d1 {

			var_43b69469f697 := value
			var var_43b69469f697_mapped *structpb.Value

			var var_43b69469f697_err error
			var_43b69469f697_mapped, var_43b69469f697_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_43b69469f697)
			if var_43b69469f697_err != nil {
				panic(var_43b69469f697_err)
			}

			var_1c6164b998d1_st.Fields[key] = var_43b69469f697_mapped
		}
		var_1c6164b998d1_mapped = structpb.NewStructValue(var_1c6164b998d1_st)
		properties["annotations"] = var_1c6164b998d1_mapped
	}
	return properties
}

func (m *ExtensionMapper) FromProperties(properties map[string]*structpb.Value) *Extension {
	var s = m.New()
	if properties["id"] != nil {

		var_8adcb3f56c38 := properties["id"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_UUID).UnPack(var_8adcb3f56c38)

		if err != nil {
			panic(err)
		}

		var_8adcb3f56c38_mapped := new(uuid.UUID)
		*var_8adcb3f56c38_mapped = val.(uuid.UUID)

		s.Id = var_8adcb3f56c38_mapped
	}
	if properties["version"] != nil {

		var_c6ff91e32bf6 := properties["version"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(var_c6ff91e32bf6)

		if err != nil {
			panic(err)
		}

		var_c6ff91e32bf6_mapped := val.(int32)

		s.Version = var_c6ff91e32bf6_mapped
	}
	if properties["createdBy"] != nil {

		var_ac2b411141a5 := properties["createdBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_ac2b411141a5)

		if err != nil {
			panic(err)
		}

		var_ac2b411141a5_mapped := new(string)
		*var_ac2b411141a5_mapped = val.(string)

		s.CreatedBy = var_ac2b411141a5_mapped
	}
	if properties["updatedBy"] != nil {

		var_06936ed420eb := properties["updatedBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_06936ed420eb)

		if err != nil {
			panic(err)
		}

		var_06936ed420eb_mapped := new(string)
		*var_06936ed420eb_mapped = val.(string)

		s.UpdatedBy = var_06936ed420eb_mapped
	}
	if properties["createdOn"] != nil {

		var_3ab2550dabdd := properties["createdOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_3ab2550dabdd)

		if err != nil {
			panic(err)
		}

		var_3ab2550dabdd_mapped := new(time.Time)
		*var_3ab2550dabdd_mapped = val.(time.Time)

		s.CreatedOn = var_3ab2550dabdd_mapped
	}
	if properties["updatedOn"] != nil {

		var_5b9a28320856 := properties["updatedOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_5b9a28320856)

		if err != nil {
			panic(err)
		}

		var_5b9a28320856_mapped := new(time.Time)
		*var_5b9a28320856_mapped = val.(time.Time)

		s.UpdatedOn = var_5b9a28320856_mapped
	}
	if properties["name"] != nil {

		var_e0825f3e8fe0 := properties["name"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_e0825f3e8fe0)

		if err != nil {
			panic(err)
		}

		var_e0825f3e8fe0_mapped := val.(string)

		s.Name = var_e0825f3e8fe0_mapped
	}
	if properties["description"] != nil {

		var_13ebfbb07947 := properties["description"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_13ebfbb07947)

		if err != nil {
			panic(err)
		}

		var_13ebfbb07947_mapped := new(string)
		*var_13ebfbb07947_mapped = val.(string)

		s.Description = var_13ebfbb07947_mapped
	}
	if properties["selector"] != nil {

		var_d9faff7f7cfc := properties["selector"]
		var mappedValue = ExtensionEventSelectorMapperInstance.FromProperties(var_d9faff7f7cfc.GetStructValue().Fields)

		var_d9faff7f7cfc_mapped := mappedValue

		s.Selector = var_d9faff7f7cfc_mapped
	}
	if properties["order"] != nil {

		var_a7b1638980f4 := properties["order"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(var_a7b1638980f4)

		if err != nil {
			panic(err)
		}

		var_a7b1638980f4_mapped := val.(int32)

		s.Order = var_a7b1638980f4_mapped
	}
	if properties["finalizes"] != nil {

		var_69301268f691 := properties["finalizes"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_BOOL).UnPack(var_69301268f691)

		if err != nil {
			panic(err)
		}

		var_69301268f691_mapped := val.(bool)

		s.Finalizes = var_69301268f691_mapped
	}
	if properties["sync"] != nil {

		var_eb45b366886e := properties["sync"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_BOOL).UnPack(var_eb45b366886e)

		if err != nil {
			panic(err)
		}

		var_eb45b366886e_mapped := val.(bool)

		s.Sync = var_eb45b366886e_mapped
	}
	if properties["responds"] != nil {

		var_914c7d51048d := properties["responds"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_BOOL).UnPack(var_914c7d51048d)

		if err != nil {
			panic(err)
		}

		var_914c7d51048d_mapped := val.(bool)

		s.Responds = var_914c7d51048d_mapped
	}
	if properties["call"] != nil {

		var_a2265b1fc75a := properties["call"]
		var mappedValue = ExtensionExternalCallMapperInstance.FromProperties(var_a2265b1fc75a.GetStructValue().Fields)

		var_a2265b1fc75a_mapped := *mappedValue

		s.Call = var_a2265b1fc75a_mapped
	}
	if properties["annotations"] != nil {

		var_ff10bbef125c := properties["annotations"]
		var_ff10bbef125c_mapped := make(map[string]string)
		for k, v := range var_ff10bbef125c.GetStructValue().Fields {

			var_443b69855cd7 := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_443b69855cd7)

			if err != nil {
				panic(err)
			}

			var_443b69855cd7_mapped := val.(string)

			var_ff10bbef125c_mapped[k] = var_443b69855cd7_mapped
		}

		s.Annotations = var_ff10bbef125c_mapped
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

	var_11ee9e4c3a5e := extensionFunctionCall.Host

	var var_11ee9e4c3a5e_mapped *structpb.Value

	var var_11ee9e4c3a5e_err error
	var_11ee9e4c3a5e_mapped, var_11ee9e4c3a5e_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_11ee9e4c3a5e)
	if var_11ee9e4c3a5e_err != nil {
		panic(var_11ee9e4c3a5e_err)
	}
	properties["host"] = var_11ee9e4c3a5e_mapped

	var_c20f20195d4b := extensionFunctionCall.FunctionName

	var var_c20f20195d4b_mapped *structpb.Value

	var var_c20f20195d4b_err error
	var_c20f20195d4b_mapped, var_c20f20195d4b_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_c20f20195d4b)
	if var_c20f20195d4b_err != nil {
		panic(var_c20f20195d4b_err)
	}
	properties["functionName"] = var_c20f20195d4b_mapped
	return properties
}

func (m *ExtensionFunctionCallMapper) FromProperties(properties map[string]*structpb.Value) *ExtensionFunctionCall {
	var s = m.New()
	if properties["host"] != nil {

		var_f89a285933b3 := properties["host"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_f89a285933b3)

		if err != nil {
			panic(err)
		}

		var_f89a285933b3_mapped := val.(string)

		s.Host = var_f89a285933b3_mapped
	}
	if properties["functionName"] != nil {

		var_df85ac802627 := properties["functionName"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_df85ac802627)

		if err != nil {
			panic(err)
		}

		var_df85ac802627_mapped := val.(string)

		s.FunctionName = var_df85ac802627_mapped
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

	var_bde8ebeca0f3 := extensionHttpCall.Uri

	var var_bde8ebeca0f3_mapped *structpb.Value

	var var_bde8ebeca0f3_err error
	var_bde8ebeca0f3_mapped, var_bde8ebeca0f3_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_bde8ebeca0f3)
	if var_bde8ebeca0f3_err != nil {
		panic(var_bde8ebeca0f3_err)
	}
	properties["uri"] = var_bde8ebeca0f3_mapped

	var_43aef6feda70 := extensionHttpCall.Method

	var var_43aef6feda70_mapped *structpb.Value

	var var_43aef6feda70_err error
	var_43aef6feda70_mapped, var_43aef6feda70_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_43aef6feda70)
	if var_43aef6feda70_err != nil {
		panic(var_43aef6feda70_err)
	}
	properties["method"] = var_43aef6feda70_mapped
	return properties
}

func (m *ExtensionHttpCallMapper) FromProperties(properties map[string]*structpb.Value) *ExtensionHttpCall {
	var s = m.New()
	if properties["uri"] != nil {

		var_29738ef55a5a := properties["uri"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_29738ef55a5a)

		if err != nil {
			panic(err)
		}

		var_29738ef55a5a_mapped := val.(string)

		s.Uri = var_29738ef55a5a_mapped
	}
	if properties["method"] != nil {

		var_5577072b555a := properties["method"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_5577072b555a)

		if err != nil {
			panic(err)
		}

		var_5577072b555a_mapped := val.(string)

		s.Method = var_5577072b555a_mapped
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

	var_4b111ac33462 := extensionExternalCall.FunctionCall

	if var_4b111ac33462 != nil {
		var var_4b111ac33462_mapped *structpb.Value

		var_4b111ac33462_mapped = structpb.NewStructValue(&structpb.Struct{Fields: ExtensionFunctionCallMapperInstance.ToProperties(var_4b111ac33462)})
		properties["functionCall"] = var_4b111ac33462_mapped
	}

	var_07fce4ba49c6 := extensionExternalCall.HttpCall

	if var_07fce4ba49c6 != nil {
		var var_07fce4ba49c6_mapped *structpb.Value

		var_07fce4ba49c6_mapped = structpb.NewStructValue(&structpb.Struct{Fields: ExtensionHttpCallMapperInstance.ToProperties(var_07fce4ba49c6)})
		properties["httpCall"] = var_07fce4ba49c6_mapped
	}
	return properties
}

func (m *ExtensionExternalCallMapper) FromProperties(properties map[string]*structpb.Value) *ExtensionExternalCall {
	var s = m.New()
	if properties["functionCall"] != nil {

		var_c1d460576795 := properties["functionCall"]
		var mappedValue = ExtensionFunctionCallMapperInstance.FromProperties(var_c1d460576795.GetStructValue().Fields)

		var_c1d460576795_mapped := mappedValue

		s.FunctionCall = var_c1d460576795_mapped
	}
	if properties["httpCall"] != nil {

		var_c246fde346ff := properties["httpCall"]
		var mappedValue = ExtensionHttpCallMapperInstance.FromProperties(var_c246fde346ff.GetStructValue().Fields)

		var_c246fde346ff_mapped := mappedValue

		s.HttpCall = var_c246fde346ff_mapped
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

	var_f3cae89e6890 := extensionEventSelector.Actions

	if var_f3cae89e6890 != nil {
		var var_f3cae89e6890_mapped *structpb.Value

		var var_f3cae89e6890_l []*structpb.Value
		for _, value := range var_f3cae89e6890 {

			var_b3682cbe95de := value
			var var_b3682cbe95de_mapped *structpb.Value

			var var_b3682cbe95de_err error
			var_b3682cbe95de_mapped, var_b3682cbe95de_err = types.ByResourcePropertyType(model.ResourceProperty_ENUM).Pack(string(var_b3682cbe95de))
			if var_b3682cbe95de_err != nil {
				panic(var_b3682cbe95de_err)
			}

			var_f3cae89e6890_l = append(var_f3cae89e6890_l, var_b3682cbe95de_mapped)
		}
		var_f3cae89e6890_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_f3cae89e6890_l})
		properties["actions"] = var_f3cae89e6890_mapped
	}

	var_133add1f1e62 := extensionEventSelector.RecordSelector

	if var_133add1f1e62 != nil {
		var var_133add1f1e62_mapped *structpb.Value

		var_133add1f1e62_mapped = structpb.NewStructValue(&structpb.Struct{Fields: ExtensionBooleanExpressionMapperInstance.ToProperties(var_133add1f1e62)})
		properties["recordSelector"] = var_133add1f1e62_mapped
	}

	var_b4ca97aa8a05 := extensionEventSelector.Namespaces

	if var_b4ca97aa8a05 != nil {
		var var_b4ca97aa8a05_mapped *structpb.Value

		var var_b4ca97aa8a05_l []*structpb.Value
		for _, value := range var_b4ca97aa8a05 {

			var_465f0cd8867c := value
			var var_465f0cd8867c_mapped *structpb.Value

			var var_465f0cd8867c_err error
			var_465f0cd8867c_mapped, var_465f0cd8867c_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_465f0cd8867c)
			if var_465f0cd8867c_err != nil {
				panic(var_465f0cd8867c_err)
			}

			var_b4ca97aa8a05_l = append(var_b4ca97aa8a05_l, var_465f0cd8867c_mapped)
		}
		var_b4ca97aa8a05_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_b4ca97aa8a05_l})
		properties["namespaces"] = var_b4ca97aa8a05_mapped
	}

	var_7caaa2471835 := extensionEventSelector.Resources

	if var_7caaa2471835 != nil {
		var var_7caaa2471835_mapped *structpb.Value

		var var_7caaa2471835_l []*structpb.Value
		for _, value := range var_7caaa2471835 {

			var_482f7bcebad9 := value
			var var_482f7bcebad9_mapped *structpb.Value

			var var_482f7bcebad9_err error
			var_482f7bcebad9_mapped, var_482f7bcebad9_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_482f7bcebad9)
			if var_482f7bcebad9_err != nil {
				panic(var_482f7bcebad9_err)
			}

			var_7caaa2471835_l = append(var_7caaa2471835_l, var_482f7bcebad9_mapped)
		}
		var_7caaa2471835_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_7caaa2471835_l})
		properties["resources"] = var_7caaa2471835_mapped
	}

	var_4b145c1625b9 := extensionEventSelector.Ids

	if var_4b145c1625b9 != nil {
		var var_4b145c1625b9_mapped *structpb.Value

		var var_4b145c1625b9_l []*structpb.Value
		for _, value := range var_4b145c1625b9 {

			var_4e1146a27f13 := value
			var var_4e1146a27f13_mapped *structpb.Value

			var var_4e1146a27f13_err error
			var_4e1146a27f13_mapped, var_4e1146a27f13_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_4e1146a27f13)
			if var_4e1146a27f13_err != nil {
				panic(var_4e1146a27f13_err)
			}

			var_4b145c1625b9_l = append(var_4b145c1625b9_l, var_4e1146a27f13_mapped)
		}
		var_4b145c1625b9_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_4b145c1625b9_l})
		properties["ids"] = var_4b145c1625b9_mapped
	}

	var_c2c5c11af608 := extensionEventSelector.Annotations

	if var_c2c5c11af608 != nil {
		var var_c2c5c11af608_mapped *structpb.Value

		var var_c2c5c11af608_st *structpb.Struct = new(structpb.Struct)
		var_c2c5c11af608_st.Fields = make(map[string]*structpb.Value)
		for key, value := range var_c2c5c11af608 {

			var_f873d4b0fba7 := value
			var var_f873d4b0fba7_mapped *structpb.Value

			var var_f873d4b0fba7_err error
			var_f873d4b0fba7_mapped, var_f873d4b0fba7_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_f873d4b0fba7)
			if var_f873d4b0fba7_err != nil {
				panic(var_f873d4b0fba7_err)
			}

			var_c2c5c11af608_st.Fields[key] = var_f873d4b0fba7_mapped
		}
		var_c2c5c11af608_mapped = structpb.NewStructValue(var_c2c5c11af608_st)
		properties["annotations"] = var_c2c5c11af608_mapped
	}
	return properties
}

func (m *ExtensionEventSelectorMapper) FromProperties(properties map[string]*structpb.Value) *ExtensionEventSelector {
	var s = m.New()
	if properties["actions"] != nil {

		var_683efbe7b4e5 := properties["actions"]
		var_683efbe7b4e5_mapped := []EventAction{}
		for _, v := range var_683efbe7b4e5.GetListValue().Values {

			var_647d01f80f73 := v
			var_647d01f80f73_mapped := (EventAction)(var_647d01f80f73.GetStringValue())

			var_683efbe7b4e5_mapped = append(var_683efbe7b4e5_mapped, var_647d01f80f73_mapped)
		}

		s.Actions = var_683efbe7b4e5_mapped
	}
	if properties["recordSelector"] != nil {

		var_192b95ea104a := properties["recordSelector"]
		var mappedValue = ExtensionBooleanExpressionMapperInstance.FromProperties(var_192b95ea104a.GetStructValue().Fields)

		var_192b95ea104a_mapped := mappedValue

		s.RecordSelector = var_192b95ea104a_mapped
	}
	if properties["namespaces"] != nil {

		var_e59496cdedd1 := properties["namespaces"]
		var_e59496cdedd1_mapped := []string{}
		for _, v := range var_e59496cdedd1.GetListValue().Values {

			var_173999fd4a11 := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_173999fd4a11)

			if err != nil {
				panic(err)
			}

			var_173999fd4a11_mapped := val.(string)

			var_e59496cdedd1_mapped = append(var_e59496cdedd1_mapped, var_173999fd4a11_mapped)
		}

		s.Namespaces = var_e59496cdedd1_mapped
	}
	if properties["resources"] != nil {

		var_02836cb1c229 := properties["resources"]
		var_02836cb1c229_mapped := []string{}
		for _, v := range var_02836cb1c229.GetListValue().Values {

			var_39e8e3ae8c86 := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_39e8e3ae8c86)

			if err != nil {
				panic(err)
			}

			var_39e8e3ae8c86_mapped := val.(string)

			var_02836cb1c229_mapped = append(var_02836cb1c229_mapped, var_39e8e3ae8c86_mapped)
		}

		s.Resources = var_02836cb1c229_mapped
	}
	if properties["ids"] != nil {

		var_f57af122c56f := properties["ids"]
		var_f57af122c56f_mapped := []string{}
		for _, v := range var_f57af122c56f.GetListValue().Values {

			var_1c5161bae6bb := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_1c5161bae6bb)

			if err != nil {
				panic(err)
			}

			var_1c5161bae6bb_mapped := val.(string)

			var_f57af122c56f_mapped = append(var_f57af122c56f_mapped, var_1c5161bae6bb_mapped)
		}

		s.Ids = var_f57af122c56f_mapped
	}
	if properties["annotations"] != nil {

		var_c2e9380c24a3 := properties["annotations"]
		var_c2e9380c24a3_mapped := make(map[string]string)
		for k, v := range var_c2e9380c24a3.GetStructValue().Fields {

			var_c57295ad1986 := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_c57295ad1986)

			if err != nil {
				panic(err)
			}

			var_c57295ad1986_mapped := val.(string)

			var_c2e9380c24a3_mapped[k] = var_c57295ad1986_mapped
		}

		s.Annotations = var_c2e9380c24a3_mapped
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

	var_8b1d24d47b4f := extensionRecordSearchParams.Query

	if var_8b1d24d47b4f != nil {
		var var_8b1d24d47b4f_mapped *structpb.Value

		var_8b1d24d47b4f_mapped = structpb.NewStructValue(&structpb.Struct{Fields: ExtensionBooleanExpressionMapperInstance.ToProperties(var_8b1d24d47b4f)})
		properties["query"] = var_8b1d24d47b4f_mapped
	}

	var_06ba1bbe7353 := extensionRecordSearchParams.Limit

	if var_06ba1bbe7353 != nil {
		var var_06ba1bbe7353_mapped *structpb.Value

		var var_06ba1bbe7353_err error
		var_06ba1bbe7353_mapped, var_06ba1bbe7353_err = types.ByResourcePropertyType(model.ResourceProperty_INT32).Pack(*var_06ba1bbe7353)
		if var_06ba1bbe7353_err != nil {
			panic(var_06ba1bbe7353_err)
		}
		properties["limit"] = var_06ba1bbe7353_mapped
	}

	var_4650bee0a791 := extensionRecordSearchParams.Offset

	if var_4650bee0a791 != nil {
		var var_4650bee0a791_mapped *structpb.Value

		var var_4650bee0a791_err error
		var_4650bee0a791_mapped, var_4650bee0a791_err = types.ByResourcePropertyType(model.ResourceProperty_INT32).Pack(*var_4650bee0a791)
		if var_4650bee0a791_err != nil {
			panic(var_4650bee0a791_err)
		}
		properties["offset"] = var_4650bee0a791_mapped
	}

	var_ef9fe2a60c90 := extensionRecordSearchParams.ResolveReferences

	if var_ef9fe2a60c90 != nil {
		var var_ef9fe2a60c90_mapped *structpb.Value

		var var_ef9fe2a60c90_l []*structpb.Value
		for _, value := range var_ef9fe2a60c90 {

			var_15ef6cc83f5b := value
			var var_15ef6cc83f5b_mapped *structpb.Value

			var var_15ef6cc83f5b_err error
			var_15ef6cc83f5b_mapped, var_15ef6cc83f5b_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_15ef6cc83f5b)
			if var_15ef6cc83f5b_err != nil {
				panic(var_15ef6cc83f5b_err)
			}

			var_ef9fe2a60c90_l = append(var_ef9fe2a60c90_l, var_15ef6cc83f5b_mapped)
		}
		var_ef9fe2a60c90_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_ef9fe2a60c90_l})
		properties["resolveReferences"] = var_ef9fe2a60c90_mapped
	}
	return properties
}

func (m *ExtensionRecordSearchParamsMapper) FromProperties(properties map[string]*structpb.Value) *ExtensionRecordSearchParams {
	var s = m.New()
	if properties["query"] != nil {

		var_2ce1654160c0 := properties["query"]
		var mappedValue = ExtensionBooleanExpressionMapperInstance.FromProperties(var_2ce1654160c0.GetStructValue().Fields)

		var_2ce1654160c0_mapped := mappedValue

		s.Query = var_2ce1654160c0_mapped
	}
	if properties["limit"] != nil {

		var_b16e748fafdd := properties["limit"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(var_b16e748fafdd)

		if err != nil {
			panic(err)
		}

		var_b16e748fafdd_mapped := new(int32)
		*var_b16e748fafdd_mapped = val.(int32)

		s.Limit = var_b16e748fafdd_mapped
	}
	if properties["offset"] != nil {

		var_e97fb1874eb7 := properties["offset"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(var_e97fb1874eb7)

		if err != nil {
			panic(err)
		}

		var_e97fb1874eb7_mapped := new(int32)
		*var_e97fb1874eb7_mapped = val.(int32)

		s.Offset = var_e97fb1874eb7_mapped
	}
	if properties["resolveReferences"] != nil {

		var_ce5bc07b9f62 := properties["resolveReferences"]
		var_ce5bc07b9f62_mapped := []string{}
		for _, v := range var_ce5bc07b9f62.GetListValue().Values {

			var_91acda3e305e := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_91acda3e305e)

			if err != nil {
				panic(err)
			}

			var_91acda3e305e_mapped := val.(string)

			var_ce5bc07b9f62_mapped = append(var_ce5bc07b9f62_mapped, var_91acda3e305e_mapped)
		}

		s.ResolveReferences = var_ce5bc07b9f62_mapped
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

	var_c9561ddc842d := extensionEvent.Id

	if var_c9561ddc842d != nil {
		var var_c9561ddc842d_mapped *structpb.Value

		var var_c9561ddc842d_err error
		var_c9561ddc842d_mapped, var_c9561ddc842d_err = types.ByResourcePropertyType(model.ResourceProperty_UUID).Pack(*var_c9561ddc842d)
		if var_c9561ddc842d_err != nil {
			panic(var_c9561ddc842d_err)
		}
		properties["id"] = var_c9561ddc842d_mapped
	}

	var_0105aba78fe8 := extensionEvent.Action

	var var_0105aba78fe8_mapped *structpb.Value

	var var_0105aba78fe8_err error
	var_0105aba78fe8_mapped, var_0105aba78fe8_err = types.ByResourcePropertyType(model.ResourceProperty_ENUM).Pack(string(var_0105aba78fe8))
	if var_0105aba78fe8_err != nil {
		panic(var_0105aba78fe8_err)
	}
	properties["action"] = var_0105aba78fe8_mapped

	var_410219c7bea1 := extensionEvent.RecordSearchParams

	if var_410219c7bea1 != nil {
		var var_410219c7bea1_mapped *structpb.Value

		var_410219c7bea1_mapped = structpb.NewStructValue(&structpb.Struct{Fields: ExtensionRecordSearchParamsMapperInstance.ToProperties(var_410219c7bea1)})
		properties["recordSearchParams"] = var_410219c7bea1_mapped
	}

	var_27ffe20fb1ca := extensionEvent.ActionSummary

	if var_27ffe20fb1ca != nil {
		var var_27ffe20fb1ca_mapped *structpb.Value

		var var_27ffe20fb1ca_err error
		var_27ffe20fb1ca_mapped, var_27ffe20fb1ca_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_27ffe20fb1ca)
		if var_27ffe20fb1ca_err != nil {
			panic(var_27ffe20fb1ca_err)
		}
		properties["actionSummary"] = var_27ffe20fb1ca_mapped
	}

	var_34103c8fa1b7 := extensionEvent.ActionDescription

	if var_34103c8fa1b7 != nil {
		var var_34103c8fa1b7_mapped *structpb.Value

		var var_34103c8fa1b7_err error
		var_34103c8fa1b7_mapped, var_34103c8fa1b7_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_34103c8fa1b7)
		if var_34103c8fa1b7_err != nil {
			panic(var_34103c8fa1b7_err)
		}
		properties["actionDescription"] = var_34103c8fa1b7_mapped
	}

	var_fa07c047cb1b := extensionEvent.Resource

	if var_fa07c047cb1b != nil {
		var var_fa07c047cb1b_mapped *structpb.Value

		var_fa07c047cb1b_mapped = structpb.NewStructValue(&structpb.Struct{Fields: ResourceMapperInstance.ToProperties(var_fa07c047cb1b)})
		properties["resource"] = var_fa07c047cb1b_mapped
	}

	var_15da586e1bd1 := extensionEvent.Records

	if var_15da586e1bd1 != nil {
		var var_15da586e1bd1_mapped *structpb.Value

		var var_15da586e1bd1_l []*structpb.Value
		for _, value := range var_15da586e1bd1 {

			var_be7d4631c940 := value
			var var_be7d4631c940_mapped *structpb.Value

			var_be7d4631c940_mapped = structpb.NewStructValue(&structpb.Struct{Fields: RecordMapperInstance.ToProperties(var_be7d4631c940)})

			var_15da586e1bd1_l = append(var_15da586e1bd1_l, var_be7d4631c940_mapped)
		}
		var_15da586e1bd1_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_15da586e1bd1_l})
		properties["records"] = var_15da586e1bd1_mapped
	}

	var_06837468734d := extensionEvent.Ids

	if var_06837468734d != nil {
		var var_06837468734d_mapped *structpb.Value

		var var_06837468734d_l []*structpb.Value
		for _, value := range var_06837468734d {

			var_a2454e480b74 := value
			var var_a2454e480b74_mapped *structpb.Value

			var var_a2454e480b74_err error
			var_a2454e480b74_mapped, var_a2454e480b74_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_a2454e480b74)
			if var_a2454e480b74_err != nil {
				panic(var_a2454e480b74_err)
			}

			var_06837468734d_l = append(var_06837468734d_l, var_a2454e480b74_mapped)
		}
		var_06837468734d_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_06837468734d_l})
		properties["ids"] = var_06837468734d_mapped
	}

	var_6f5070991d53 := extensionEvent.Finalizes

	if var_6f5070991d53 != nil {
		var var_6f5070991d53_mapped *structpb.Value

		var var_6f5070991d53_err error
		var_6f5070991d53_mapped, var_6f5070991d53_err = types.ByResourcePropertyType(model.ResourceProperty_BOOL).Pack(*var_6f5070991d53)
		if var_6f5070991d53_err != nil {
			panic(var_6f5070991d53_err)
		}
		properties["finalizes"] = var_6f5070991d53_mapped
	}

	var_eddac9d221a0 := extensionEvent.Sync

	if var_eddac9d221a0 != nil {
		var var_eddac9d221a0_mapped *structpb.Value

		var var_eddac9d221a0_err error
		var_eddac9d221a0_mapped, var_eddac9d221a0_err = types.ByResourcePropertyType(model.ResourceProperty_BOOL).Pack(*var_eddac9d221a0)
		if var_eddac9d221a0_err != nil {
			panic(var_eddac9d221a0_err)
		}
		properties["sync"] = var_eddac9d221a0_mapped
	}

	var_9b1e95eb67f5 := extensionEvent.Time

	if var_9b1e95eb67f5 != nil {
		var var_9b1e95eb67f5_mapped *structpb.Value

		var var_9b1e95eb67f5_err error
		var_9b1e95eb67f5_mapped, var_9b1e95eb67f5_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_9b1e95eb67f5)
		if var_9b1e95eb67f5_err != nil {
			panic(var_9b1e95eb67f5_err)
		}
		properties["time"] = var_9b1e95eb67f5_mapped
	}

	var_21b3041618cf := extensionEvent.Annotations

	if var_21b3041618cf != nil {
		var var_21b3041618cf_mapped *structpb.Value

		var var_21b3041618cf_st *structpb.Struct = new(structpb.Struct)
		var_21b3041618cf_st.Fields = make(map[string]*structpb.Value)
		for key, value := range var_21b3041618cf {

			var_6a2dcd48ccfd := value
			var var_6a2dcd48ccfd_mapped *structpb.Value

			var var_6a2dcd48ccfd_err error
			var_6a2dcd48ccfd_mapped, var_6a2dcd48ccfd_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_6a2dcd48ccfd)
			if var_6a2dcd48ccfd_err != nil {
				panic(var_6a2dcd48ccfd_err)
			}

			var_21b3041618cf_st.Fields[key] = var_6a2dcd48ccfd_mapped
		}
		var_21b3041618cf_mapped = structpb.NewStructValue(var_21b3041618cf_st)
		properties["annotations"] = var_21b3041618cf_mapped
	}
	return properties
}

func (m *ExtensionEventMapper) FromProperties(properties map[string]*structpb.Value) *ExtensionEvent {
	var s = m.New()
	if properties["id"] != nil {

		var_d9f35e00d956 := properties["id"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_UUID).UnPack(var_d9f35e00d956)

		if err != nil {
			panic(err)
		}

		var_d9f35e00d956_mapped := new(uuid.UUID)
		*var_d9f35e00d956_mapped = val.(uuid.UUID)

		s.Id = var_d9f35e00d956_mapped
	}
	if properties["action"] != nil {

		var_a234fd1a40d8 := properties["action"]
		var_a234fd1a40d8_mapped := (EventAction)(var_a234fd1a40d8.GetStringValue())

		s.Action = var_a234fd1a40d8_mapped
	}
	if properties["recordSearchParams"] != nil {

		var_3fe9eae785e8 := properties["recordSearchParams"]
		var mappedValue = ExtensionRecordSearchParamsMapperInstance.FromProperties(var_3fe9eae785e8.GetStructValue().Fields)

		var_3fe9eae785e8_mapped := mappedValue

		s.RecordSearchParams = var_3fe9eae785e8_mapped
	}
	if properties["actionSummary"] != nil {

		var_c268eb74c193 := properties["actionSummary"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_c268eb74c193)

		if err != nil {
			panic(err)
		}

		var_c268eb74c193_mapped := new(string)
		*var_c268eb74c193_mapped = val.(string)

		s.ActionSummary = var_c268eb74c193_mapped
	}
	if properties["actionDescription"] != nil {

		var_175d05db9b01 := properties["actionDescription"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_175d05db9b01)

		if err != nil {
			panic(err)
		}

		var_175d05db9b01_mapped := new(string)
		*var_175d05db9b01_mapped = val.(string)

		s.ActionDescription = var_175d05db9b01_mapped
	}
	if properties["resource"] != nil {

		var_a2d05a9b3e90 := properties["resource"]
		var_a2d05a9b3e90_mapped := ResourceMapperInstance.FromProperties(var_a2d05a9b3e90.GetStructValue().Fields)

		s.Resource = var_a2d05a9b3e90_mapped
	}
	if properties["records"] != nil {

		var_20babf29b8da := properties["records"]
		var_20babf29b8da_mapped := []*Record{}
		for _, v := range var_20babf29b8da.GetListValue().Values {

			var_e9227bfd942f := v
			var_e9227bfd942f_mapped := RecordMapperInstance.FromProperties(var_e9227bfd942f.GetStructValue().Fields)

			var_20babf29b8da_mapped = append(var_20babf29b8da_mapped, var_e9227bfd942f_mapped)
		}

		s.Records = var_20babf29b8da_mapped
	}
	if properties["ids"] != nil {

		var_188a6daa1732 := properties["ids"]
		var_188a6daa1732_mapped := []string{}
		for _, v := range var_188a6daa1732.GetListValue().Values {

			var_8b5b060574b2 := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_8b5b060574b2)

			if err != nil {
				panic(err)
			}

			var_8b5b060574b2_mapped := val.(string)

			var_188a6daa1732_mapped = append(var_188a6daa1732_mapped, var_8b5b060574b2_mapped)
		}

		s.Ids = var_188a6daa1732_mapped
	}
	if properties["finalizes"] != nil {

		var_c6b922b3ceca := properties["finalizes"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_BOOL).UnPack(var_c6b922b3ceca)

		if err != nil {
			panic(err)
		}

		var_c6b922b3ceca_mapped := new(bool)
		*var_c6b922b3ceca_mapped = val.(bool)

		s.Finalizes = var_c6b922b3ceca_mapped
	}
	if properties["sync"] != nil {

		var_9beba367a4af := properties["sync"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_BOOL).UnPack(var_9beba367a4af)

		if err != nil {
			panic(err)
		}

		var_9beba367a4af_mapped := new(bool)
		*var_9beba367a4af_mapped = val.(bool)

		s.Sync = var_9beba367a4af_mapped
	}
	if properties["time"] != nil {

		var_b87e99af5135 := properties["time"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_b87e99af5135)

		if err != nil {
			panic(err)
		}

		var_b87e99af5135_mapped := new(time.Time)
		*var_b87e99af5135_mapped = val.(time.Time)

		s.Time = var_b87e99af5135_mapped
	}
	if properties["annotations"] != nil {

		var_e6ab17672f14 := properties["annotations"]
		var_e6ab17672f14_mapped := make(map[string]string)
		for k, v := range var_e6ab17672f14.GetStructValue().Fields {

			var_3c047199372b := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_3c047199372b)

			if err != nil {
				panic(err)
			}

			var_3c047199372b_mapped := val.(string)

			var_e6ab17672f14_mapped[k] = var_3c047199372b_mapped
		}

		s.Annotations = var_e6ab17672f14_mapped
	}
	return s
}
