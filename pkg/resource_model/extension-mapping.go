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

	var_59296887f91d := extension.Id

	if var_59296887f91d != nil {
		var var_59296887f91d_mapped *structpb.Value

		var var_59296887f91d_err error
		var_59296887f91d_mapped, var_59296887f91d_err = types.ByResourcePropertyType(model.ResourceProperty_UUID).Pack(*var_59296887f91d)
		if var_59296887f91d_err != nil {
			panic(var_59296887f91d_err)
		}
		properties["id"] = var_59296887f91d_mapped
	}

	var_2081bf3802f5 := extension.Version

	var var_2081bf3802f5_mapped *structpb.Value

	var var_2081bf3802f5_err error
	var_2081bf3802f5_mapped, var_2081bf3802f5_err = types.ByResourcePropertyType(model.ResourceProperty_INT32).Pack(var_2081bf3802f5)
	if var_2081bf3802f5_err != nil {
		panic(var_2081bf3802f5_err)
	}
	properties["version"] = var_2081bf3802f5_mapped

	var_7cfe16a56caf := extension.CreatedBy

	if var_7cfe16a56caf != nil {
		var var_7cfe16a56caf_mapped *structpb.Value

		var var_7cfe16a56caf_err error
		var_7cfe16a56caf_mapped, var_7cfe16a56caf_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_7cfe16a56caf)
		if var_7cfe16a56caf_err != nil {
			panic(var_7cfe16a56caf_err)
		}
		properties["createdBy"] = var_7cfe16a56caf_mapped
	}

	var_4e75a0e21d4b := extension.UpdatedBy

	if var_4e75a0e21d4b != nil {
		var var_4e75a0e21d4b_mapped *structpb.Value

		var var_4e75a0e21d4b_err error
		var_4e75a0e21d4b_mapped, var_4e75a0e21d4b_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_4e75a0e21d4b)
		if var_4e75a0e21d4b_err != nil {
			panic(var_4e75a0e21d4b_err)
		}
		properties["updatedBy"] = var_4e75a0e21d4b_mapped
	}

	var_6cffc7e87db2 := extension.CreatedOn

	if var_6cffc7e87db2 != nil {
		var var_6cffc7e87db2_mapped *structpb.Value

		var var_6cffc7e87db2_err error
		var_6cffc7e87db2_mapped, var_6cffc7e87db2_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_6cffc7e87db2)
		if var_6cffc7e87db2_err != nil {
			panic(var_6cffc7e87db2_err)
		}
		properties["createdOn"] = var_6cffc7e87db2_mapped
	}

	var_f0662419144b := extension.UpdatedOn

	if var_f0662419144b != nil {
		var var_f0662419144b_mapped *structpb.Value

		var var_f0662419144b_err error
		var_f0662419144b_mapped, var_f0662419144b_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_f0662419144b)
		if var_f0662419144b_err != nil {
			panic(var_f0662419144b_err)
		}
		properties["updatedOn"] = var_f0662419144b_mapped
	}

	var_5793618f792a := extension.Name

	var var_5793618f792a_mapped *structpb.Value

	var var_5793618f792a_err error
	var_5793618f792a_mapped, var_5793618f792a_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_5793618f792a)
	if var_5793618f792a_err != nil {
		panic(var_5793618f792a_err)
	}
	properties["name"] = var_5793618f792a_mapped

	var_7200da44666c := extension.Description

	if var_7200da44666c != nil {
		var var_7200da44666c_mapped *structpb.Value

		var var_7200da44666c_err error
		var_7200da44666c_mapped, var_7200da44666c_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_7200da44666c)
		if var_7200da44666c_err != nil {
			panic(var_7200da44666c_err)
		}
		properties["description"] = var_7200da44666c_mapped
	}

	var_29d651f47eeb := extension.Selector

	if var_29d651f47eeb != nil {
		var var_29d651f47eeb_mapped *structpb.Value

		var_29d651f47eeb_mapped = structpb.NewStructValue(&structpb.Struct{Fields: ExtensionEventSelectorMapperInstance.ToProperties(var_29d651f47eeb)})
		properties["selector"] = var_29d651f47eeb_mapped
	}

	var_f7d16765d158 := extension.Order

	var var_f7d16765d158_mapped *structpb.Value

	var var_f7d16765d158_err error
	var_f7d16765d158_mapped, var_f7d16765d158_err = types.ByResourcePropertyType(model.ResourceProperty_INT32).Pack(var_f7d16765d158)
	if var_f7d16765d158_err != nil {
		panic(var_f7d16765d158_err)
	}
	properties["order"] = var_f7d16765d158_mapped

	var_faadc90bc723 := extension.Finalizes

	var var_faadc90bc723_mapped *structpb.Value

	var var_faadc90bc723_err error
	var_faadc90bc723_mapped, var_faadc90bc723_err = types.ByResourcePropertyType(model.ResourceProperty_BOOL).Pack(var_faadc90bc723)
	if var_faadc90bc723_err != nil {
		panic(var_faadc90bc723_err)
	}
	properties["finalizes"] = var_faadc90bc723_mapped

	var_dc6bfaa53a5b := extension.Sync

	var var_dc6bfaa53a5b_mapped *structpb.Value

	var var_dc6bfaa53a5b_err error
	var_dc6bfaa53a5b_mapped, var_dc6bfaa53a5b_err = types.ByResourcePropertyType(model.ResourceProperty_BOOL).Pack(var_dc6bfaa53a5b)
	if var_dc6bfaa53a5b_err != nil {
		panic(var_dc6bfaa53a5b_err)
	}
	properties["sync"] = var_dc6bfaa53a5b_mapped

	var_c6f49bb6d909 := extension.Responds

	var var_c6f49bb6d909_mapped *structpb.Value

	var var_c6f49bb6d909_err error
	var_c6f49bb6d909_mapped, var_c6f49bb6d909_err = types.ByResourcePropertyType(model.ResourceProperty_BOOL).Pack(var_c6f49bb6d909)
	if var_c6f49bb6d909_err != nil {
		panic(var_c6f49bb6d909_err)
	}
	properties["responds"] = var_c6f49bb6d909_mapped

	var_b72d11a5dfe4 := extension.Call

	var var_b72d11a5dfe4_mapped *structpb.Value

	var_b72d11a5dfe4_mapped = structpb.NewStructValue(&structpb.Struct{Fields: ExtensionExternalCallMapperInstance.ToProperties(&var_b72d11a5dfe4)})
	properties["call"] = var_b72d11a5dfe4_mapped

	var_ce26a51fb8c4 := extension.Annotations

	if var_ce26a51fb8c4 != nil {
		var var_ce26a51fb8c4_mapped *structpb.Value

		var var_ce26a51fb8c4_st *structpb.Struct = new(structpb.Struct)
		var_ce26a51fb8c4_st.Fields = make(map[string]*structpb.Value)
		for key, value := range var_ce26a51fb8c4 {

			var_1dc2fe8684e3 := value
			var var_1dc2fe8684e3_mapped *structpb.Value

			var var_1dc2fe8684e3_err error
			var_1dc2fe8684e3_mapped, var_1dc2fe8684e3_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_1dc2fe8684e3)
			if var_1dc2fe8684e3_err != nil {
				panic(var_1dc2fe8684e3_err)
			}

			var_ce26a51fb8c4_st.Fields[key] = var_1dc2fe8684e3_mapped
		}
		var_ce26a51fb8c4_mapped = structpb.NewStructValue(var_ce26a51fb8c4_st)
		properties["annotations"] = var_ce26a51fb8c4_mapped
	}
	return properties
}

func (m *ExtensionMapper) FromProperties(properties map[string]*structpb.Value) *Extension {
	var s = m.New()
	if properties["id"] != nil {

		var_18d8c79dd79b := properties["id"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_UUID).UnPack(var_18d8c79dd79b)

		if err != nil {
			panic(err)
		}

		var_18d8c79dd79b_mapped := new(uuid.UUID)
		*var_18d8c79dd79b_mapped = val.(uuid.UUID)

		s.Id = var_18d8c79dd79b_mapped
	}
	if properties["version"] != nil {

		var_0e4d57c721a8 := properties["version"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(var_0e4d57c721a8)

		if err != nil {
			panic(err)
		}

		var_0e4d57c721a8_mapped := val.(int32)

		s.Version = var_0e4d57c721a8_mapped
	}
	if properties["createdBy"] != nil {

		var_747f037e8d7f := properties["createdBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_747f037e8d7f)

		if err != nil {
			panic(err)
		}

		var_747f037e8d7f_mapped := new(string)
		*var_747f037e8d7f_mapped = val.(string)

		s.CreatedBy = var_747f037e8d7f_mapped
	}
	if properties["updatedBy"] != nil {

		var_8a43e37d24cb := properties["updatedBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_8a43e37d24cb)

		if err != nil {
			panic(err)
		}

		var_8a43e37d24cb_mapped := new(string)
		*var_8a43e37d24cb_mapped = val.(string)

		s.UpdatedBy = var_8a43e37d24cb_mapped
	}
	if properties["createdOn"] != nil {

		var_fe2882e9792f := properties["createdOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_fe2882e9792f)

		if err != nil {
			panic(err)
		}

		var_fe2882e9792f_mapped := new(time.Time)
		*var_fe2882e9792f_mapped = val.(time.Time)

		s.CreatedOn = var_fe2882e9792f_mapped
	}
	if properties["updatedOn"] != nil {

		var_85f5749a40e7 := properties["updatedOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_85f5749a40e7)

		if err != nil {
			panic(err)
		}

		var_85f5749a40e7_mapped := new(time.Time)
		*var_85f5749a40e7_mapped = val.(time.Time)

		s.UpdatedOn = var_85f5749a40e7_mapped
	}
	if properties["name"] != nil {

		var_d777e5fcdfab := properties["name"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_d777e5fcdfab)

		if err != nil {
			panic(err)
		}

		var_d777e5fcdfab_mapped := val.(string)

		s.Name = var_d777e5fcdfab_mapped
	}
	if properties["description"] != nil {

		var_1129b55fab3f := properties["description"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_1129b55fab3f)

		if err != nil {
			panic(err)
		}

		var_1129b55fab3f_mapped := new(string)
		*var_1129b55fab3f_mapped = val.(string)

		s.Description = var_1129b55fab3f_mapped
	}
	if properties["selector"] != nil {

		var_7528f1c84a70 := properties["selector"]
		var mappedValue = ExtensionEventSelectorMapperInstance.FromProperties(var_7528f1c84a70.GetStructValue().Fields)

		var_7528f1c84a70_mapped := mappedValue

		s.Selector = var_7528f1c84a70_mapped
	}
	if properties["order"] != nil {

		var_36c86ecf01a8 := properties["order"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(var_36c86ecf01a8)

		if err != nil {
			panic(err)
		}

		var_36c86ecf01a8_mapped := val.(int32)

		s.Order = var_36c86ecf01a8_mapped
	}
	if properties["finalizes"] != nil {

		var_b3b4511de7e8 := properties["finalizes"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_BOOL).UnPack(var_b3b4511de7e8)

		if err != nil {
			panic(err)
		}

		var_b3b4511de7e8_mapped := val.(bool)

		s.Finalizes = var_b3b4511de7e8_mapped
	}
	if properties["sync"] != nil {

		var_369862496ffc := properties["sync"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_BOOL).UnPack(var_369862496ffc)

		if err != nil {
			panic(err)
		}

		var_369862496ffc_mapped := val.(bool)

		s.Sync = var_369862496ffc_mapped
	}
	if properties["responds"] != nil {

		var_a99a4f0c8b9c := properties["responds"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_BOOL).UnPack(var_a99a4f0c8b9c)

		if err != nil {
			panic(err)
		}

		var_a99a4f0c8b9c_mapped := val.(bool)

		s.Responds = var_a99a4f0c8b9c_mapped
	}
	if properties["call"] != nil {

		var_9421df8b9cb6 := properties["call"]
		var mappedValue = ExtensionExternalCallMapperInstance.FromProperties(var_9421df8b9cb6.GetStructValue().Fields)

		var_9421df8b9cb6_mapped := *mappedValue

		s.Call = var_9421df8b9cb6_mapped
	}
	if properties["annotations"] != nil {

		var_6c67fa76a324 := properties["annotations"]
		var_6c67fa76a324_mapped := make(map[string]string)
		for k, v := range var_6c67fa76a324.GetStructValue().Fields {

			var_63bf2649c5b4 := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_63bf2649c5b4)

			if err != nil {
				panic(err)
			}

			var_63bf2649c5b4_mapped := val.(string)

			var_6c67fa76a324_mapped[k] = var_63bf2649c5b4_mapped
		}

		s.Annotations = var_6c67fa76a324_mapped
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

	var_b4085ef8da3e := extensionFunctionCall.Host

	var var_b4085ef8da3e_mapped *structpb.Value

	var var_b4085ef8da3e_err error
	var_b4085ef8da3e_mapped, var_b4085ef8da3e_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_b4085ef8da3e)
	if var_b4085ef8da3e_err != nil {
		panic(var_b4085ef8da3e_err)
	}
	properties["host"] = var_b4085ef8da3e_mapped

	var_ee9a0a4185b9 := extensionFunctionCall.FunctionName

	var var_ee9a0a4185b9_mapped *structpb.Value

	var var_ee9a0a4185b9_err error
	var_ee9a0a4185b9_mapped, var_ee9a0a4185b9_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_ee9a0a4185b9)
	if var_ee9a0a4185b9_err != nil {
		panic(var_ee9a0a4185b9_err)
	}
	properties["functionName"] = var_ee9a0a4185b9_mapped
	return properties
}

func (m *ExtensionFunctionCallMapper) FromProperties(properties map[string]*structpb.Value) *ExtensionFunctionCall {
	var s = m.New()
	if properties["host"] != nil {

		var_3474f5bf4035 := properties["host"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_3474f5bf4035)

		if err != nil {
			panic(err)
		}

		var_3474f5bf4035_mapped := val.(string)

		s.Host = var_3474f5bf4035_mapped
	}
	if properties["functionName"] != nil {

		var_7065e9f64772 := properties["functionName"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_7065e9f64772)

		if err != nil {
			panic(err)
		}

		var_7065e9f64772_mapped := val.(string)

		s.FunctionName = var_7065e9f64772_mapped
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

	var_9b94ff429c9e := extensionHttpCall.Uri

	var var_9b94ff429c9e_mapped *structpb.Value

	var var_9b94ff429c9e_err error
	var_9b94ff429c9e_mapped, var_9b94ff429c9e_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_9b94ff429c9e)
	if var_9b94ff429c9e_err != nil {
		panic(var_9b94ff429c9e_err)
	}
	properties["uri"] = var_9b94ff429c9e_mapped

	var_51c1e73d29b5 := extensionHttpCall.Method

	var var_51c1e73d29b5_mapped *structpb.Value

	var var_51c1e73d29b5_err error
	var_51c1e73d29b5_mapped, var_51c1e73d29b5_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_51c1e73d29b5)
	if var_51c1e73d29b5_err != nil {
		panic(var_51c1e73d29b5_err)
	}
	properties["method"] = var_51c1e73d29b5_mapped
	return properties
}

func (m *ExtensionHttpCallMapper) FromProperties(properties map[string]*structpb.Value) *ExtensionHttpCall {
	var s = m.New()
	if properties["uri"] != nil {

		var_e56c9ac78613 := properties["uri"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_e56c9ac78613)

		if err != nil {
			panic(err)
		}

		var_e56c9ac78613_mapped := val.(string)

		s.Uri = var_e56c9ac78613_mapped
	}
	if properties["method"] != nil {

		var_fdc1d94ae9ed := properties["method"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_fdc1d94ae9ed)

		if err != nil {
			panic(err)
		}

		var_fdc1d94ae9ed_mapped := val.(string)

		s.Method = var_fdc1d94ae9ed_mapped
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

	var_06e19bae9e14 := extensionExternalCall.FunctionCall

	if var_06e19bae9e14 != nil {
		var var_06e19bae9e14_mapped *structpb.Value

		var_06e19bae9e14_mapped = structpb.NewStructValue(&structpb.Struct{Fields: ExtensionFunctionCallMapperInstance.ToProperties(var_06e19bae9e14)})
		properties["functionCall"] = var_06e19bae9e14_mapped
	}

	var_3b1c2111fe19 := extensionExternalCall.HttpCall

	if var_3b1c2111fe19 != nil {
		var var_3b1c2111fe19_mapped *structpb.Value

		var_3b1c2111fe19_mapped = structpb.NewStructValue(&structpb.Struct{Fields: ExtensionHttpCallMapperInstance.ToProperties(var_3b1c2111fe19)})
		properties["httpCall"] = var_3b1c2111fe19_mapped
	}
	return properties
}

func (m *ExtensionExternalCallMapper) FromProperties(properties map[string]*structpb.Value) *ExtensionExternalCall {
	var s = m.New()
	if properties["functionCall"] != nil {

		var_a033d5f41329 := properties["functionCall"]
		var mappedValue = ExtensionFunctionCallMapperInstance.FromProperties(var_a033d5f41329.GetStructValue().Fields)

		var_a033d5f41329_mapped := mappedValue

		s.FunctionCall = var_a033d5f41329_mapped
	}
	if properties["httpCall"] != nil {

		var_82e230c72f12 := properties["httpCall"]
		var mappedValue = ExtensionHttpCallMapperInstance.FromProperties(var_82e230c72f12.GetStructValue().Fields)

		var_82e230c72f12_mapped := mappedValue

		s.HttpCall = var_82e230c72f12_mapped
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

	var_f380c9b3cf28 := extensionEventSelector.Actions

	if var_f380c9b3cf28 != nil {
		var var_f380c9b3cf28_mapped *structpb.Value

		var var_f380c9b3cf28_l []*structpb.Value
		for _, value := range var_f380c9b3cf28 {

			var_a0e012159c25 := value
			var var_a0e012159c25_mapped *structpb.Value

			var var_a0e012159c25_err error
			var_a0e012159c25_mapped, var_a0e012159c25_err = types.ByResourcePropertyType(model.ResourceProperty_ENUM).Pack(string(var_a0e012159c25))
			if var_a0e012159c25_err != nil {
				panic(var_a0e012159c25_err)
			}

			var_f380c9b3cf28_l = append(var_f380c9b3cf28_l, var_a0e012159c25_mapped)
		}
		var_f380c9b3cf28_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_f380c9b3cf28_l})
		properties["actions"] = var_f380c9b3cf28_mapped
	}

	var_2cfcd88b3492 := extensionEventSelector.RecordSelector

	if var_2cfcd88b3492 != nil {
		var var_2cfcd88b3492_mapped *structpb.Value

		var_2cfcd88b3492_mapped = structpb.NewStructValue(&structpb.Struct{Fields: ExtensionBooleanExpressionMapperInstance.ToProperties(var_2cfcd88b3492)})
		properties["recordSelector"] = var_2cfcd88b3492_mapped
	}

	var_ef134c28822e := extensionEventSelector.Namespaces

	if var_ef134c28822e != nil {
		var var_ef134c28822e_mapped *structpb.Value

		var var_ef134c28822e_l []*structpb.Value
		for _, value := range var_ef134c28822e {

			var_c5ef5cf0fd4d := value
			var var_c5ef5cf0fd4d_mapped *structpb.Value

			var var_c5ef5cf0fd4d_err error
			var_c5ef5cf0fd4d_mapped, var_c5ef5cf0fd4d_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_c5ef5cf0fd4d)
			if var_c5ef5cf0fd4d_err != nil {
				panic(var_c5ef5cf0fd4d_err)
			}

			var_ef134c28822e_l = append(var_ef134c28822e_l, var_c5ef5cf0fd4d_mapped)
		}
		var_ef134c28822e_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_ef134c28822e_l})
		properties["namespaces"] = var_ef134c28822e_mapped
	}

	var_168cf49c3589 := extensionEventSelector.Resources

	if var_168cf49c3589 != nil {
		var var_168cf49c3589_mapped *structpb.Value

		var var_168cf49c3589_l []*structpb.Value
		for _, value := range var_168cf49c3589 {

			var_6a115b5f24be := value
			var var_6a115b5f24be_mapped *structpb.Value

			var var_6a115b5f24be_err error
			var_6a115b5f24be_mapped, var_6a115b5f24be_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_6a115b5f24be)
			if var_6a115b5f24be_err != nil {
				panic(var_6a115b5f24be_err)
			}

			var_168cf49c3589_l = append(var_168cf49c3589_l, var_6a115b5f24be_mapped)
		}
		var_168cf49c3589_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_168cf49c3589_l})
		properties["resources"] = var_168cf49c3589_mapped
	}

	var_0d08cd9cb6b6 := extensionEventSelector.Ids

	if var_0d08cd9cb6b6 != nil {
		var var_0d08cd9cb6b6_mapped *structpb.Value

		var var_0d08cd9cb6b6_l []*structpb.Value
		for _, value := range var_0d08cd9cb6b6 {

			var_7f7acaf9b339 := value
			var var_7f7acaf9b339_mapped *structpb.Value

			var var_7f7acaf9b339_err error
			var_7f7acaf9b339_mapped, var_7f7acaf9b339_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_7f7acaf9b339)
			if var_7f7acaf9b339_err != nil {
				panic(var_7f7acaf9b339_err)
			}

			var_0d08cd9cb6b6_l = append(var_0d08cd9cb6b6_l, var_7f7acaf9b339_mapped)
		}
		var_0d08cd9cb6b6_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_0d08cd9cb6b6_l})
		properties["ids"] = var_0d08cd9cb6b6_mapped
	}

	var_9b64d78ad471 := extensionEventSelector.Annotations

	if var_9b64d78ad471 != nil {
		var var_9b64d78ad471_mapped *structpb.Value

		var var_9b64d78ad471_st *structpb.Struct = new(structpb.Struct)
		var_9b64d78ad471_st.Fields = make(map[string]*structpb.Value)
		for key, value := range var_9b64d78ad471 {

			var_1e3a4832d4ff := value
			var var_1e3a4832d4ff_mapped *structpb.Value

			var var_1e3a4832d4ff_err error
			var_1e3a4832d4ff_mapped, var_1e3a4832d4ff_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_1e3a4832d4ff)
			if var_1e3a4832d4ff_err != nil {
				panic(var_1e3a4832d4ff_err)
			}

			var_9b64d78ad471_st.Fields[key] = var_1e3a4832d4ff_mapped
		}
		var_9b64d78ad471_mapped = structpb.NewStructValue(var_9b64d78ad471_st)
		properties["annotations"] = var_9b64d78ad471_mapped
	}
	return properties
}

func (m *ExtensionEventSelectorMapper) FromProperties(properties map[string]*structpb.Value) *ExtensionEventSelector {
	var s = m.New()
	if properties["actions"] != nil {

		var_3fae033829a0 := properties["actions"]
		var_3fae033829a0_mapped := []EventAction{}
		for _, v := range var_3fae033829a0.GetListValue().Values {

			var_e76a3118c631 := v
			var_e76a3118c631_mapped := (EventAction)(var_e76a3118c631.GetStringValue())

			var_3fae033829a0_mapped = append(var_3fae033829a0_mapped, var_e76a3118c631_mapped)
		}

		s.Actions = var_3fae033829a0_mapped
	}
	if properties["recordSelector"] != nil {

		var_3e7d61b92521 := properties["recordSelector"]
		var mappedValue = ExtensionBooleanExpressionMapperInstance.FromProperties(var_3e7d61b92521.GetStructValue().Fields)

		var_3e7d61b92521_mapped := mappedValue

		s.RecordSelector = var_3e7d61b92521_mapped
	}
	if properties["namespaces"] != nil {

		var_e4b9962b49f8 := properties["namespaces"]
		var_e4b9962b49f8_mapped := []string{}
		for _, v := range var_e4b9962b49f8.GetListValue().Values {

			var_3b19fe7b825e := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_3b19fe7b825e)

			if err != nil {
				panic(err)
			}

			var_3b19fe7b825e_mapped := val.(string)

			var_e4b9962b49f8_mapped = append(var_e4b9962b49f8_mapped, var_3b19fe7b825e_mapped)
		}

		s.Namespaces = var_e4b9962b49f8_mapped
	}
	if properties["resources"] != nil {

		var_8e6852b83fc2 := properties["resources"]
		var_8e6852b83fc2_mapped := []string{}
		for _, v := range var_8e6852b83fc2.GetListValue().Values {

			var_370a8cd8efd4 := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_370a8cd8efd4)

			if err != nil {
				panic(err)
			}

			var_370a8cd8efd4_mapped := val.(string)

			var_8e6852b83fc2_mapped = append(var_8e6852b83fc2_mapped, var_370a8cd8efd4_mapped)
		}

		s.Resources = var_8e6852b83fc2_mapped
	}
	if properties["ids"] != nil {

		var_6e4e5b044fbc := properties["ids"]
		var_6e4e5b044fbc_mapped := []string{}
		for _, v := range var_6e4e5b044fbc.GetListValue().Values {

			var_6af54972aebe := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_6af54972aebe)

			if err != nil {
				panic(err)
			}

			var_6af54972aebe_mapped := val.(string)

			var_6e4e5b044fbc_mapped = append(var_6e4e5b044fbc_mapped, var_6af54972aebe_mapped)
		}

		s.Ids = var_6e4e5b044fbc_mapped
	}
	if properties["annotations"] != nil {

		var_88734d6a9521 := properties["annotations"]
		var_88734d6a9521_mapped := make(map[string]string)
		for k, v := range var_88734d6a9521.GetStructValue().Fields {

			var_ac2694aa9060 := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_ac2694aa9060)

			if err != nil {
				panic(err)
			}

			var_ac2694aa9060_mapped := val.(string)

			var_88734d6a9521_mapped[k] = var_ac2694aa9060_mapped
		}

		s.Annotations = var_88734d6a9521_mapped
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

	var_b0fe6512e95e := extensionRecordSearchParams.Query

	if var_b0fe6512e95e != nil {
		var var_b0fe6512e95e_mapped *structpb.Value

		var_b0fe6512e95e_mapped = structpb.NewStructValue(&structpb.Struct{Fields: ExtensionBooleanExpressionMapperInstance.ToProperties(var_b0fe6512e95e)})
		properties["query"] = var_b0fe6512e95e_mapped
	}

	var_592b2c4c137b := extensionRecordSearchParams.Limit

	if var_592b2c4c137b != nil {
		var var_592b2c4c137b_mapped *structpb.Value

		var var_592b2c4c137b_err error
		var_592b2c4c137b_mapped, var_592b2c4c137b_err = types.ByResourcePropertyType(model.ResourceProperty_INT32).Pack(*var_592b2c4c137b)
		if var_592b2c4c137b_err != nil {
			panic(var_592b2c4c137b_err)
		}
		properties["limit"] = var_592b2c4c137b_mapped
	}

	var_94320b711667 := extensionRecordSearchParams.Offset

	if var_94320b711667 != nil {
		var var_94320b711667_mapped *structpb.Value

		var var_94320b711667_err error
		var_94320b711667_mapped, var_94320b711667_err = types.ByResourcePropertyType(model.ResourceProperty_INT32).Pack(*var_94320b711667)
		if var_94320b711667_err != nil {
			panic(var_94320b711667_err)
		}
		properties["offset"] = var_94320b711667_mapped
	}

	var_bd7e5df82cf6 := extensionRecordSearchParams.ResolveReferences

	if var_bd7e5df82cf6 != nil {
		var var_bd7e5df82cf6_mapped *structpb.Value

		var var_bd7e5df82cf6_l []*structpb.Value
		for _, value := range var_bd7e5df82cf6 {

			var_c9afe4e12462 := value
			var var_c9afe4e12462_mapped *structpb.Value

			var var_c9afe4e12462_err error
			var_c9afe4e12462_mapped, var_c9afe4e12462_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_c9afe4e12462)
			if var_c9afe4e12462_err != nil {
				panic(var_c9afe4e12462_err)
			}

			var_bd7e5df82cf6_l = append(var_bd7e5df82cf6_l, var_c9afe4e12462_mapped)
		}
		var_bd7e5df82cf6_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_bd7e5df82cf6_l})
		properties["resolveReferences"] = var_bd7e5df82cf6_mapped
	}
	return properties
}

func (m *ExtensionRecordSearchParamsMapper) FromProperties(properties map[string]*structpb.Value) *ExtensionRecordSearchParams {
	var s = m.New()
	if properties["query"] != nil {

		var_8b1ad7cd34f7 := properties["query"]
		var mappedValue = ExtensionBooleanExpressionMapperInstance.FromProperties(var_8b1ad7cd34f7.GetStructValue().Fields)

		var_8b1ad7cd34f7_mapped := mappedValue

		s.Query = var_8b1ad7cd34f7_mapped
	}
	if properties["limit"] != nil {

		var_e510974811ef := properties["limit"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(var_e510974811ef)

		if err != nil {
			panic(err)
		}

		var_e510974811ef_mapped := new(int32)
		*var_e510974811ef_mapped = val.(int32)

		s.Limit = var_e510974811ef_mapped
	}
	if properties["offset"] != nil {

		var_cee6d53bc023 := properties["offset"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(var_cee6d53bc023)

		if err != nil {
			panic(err)
		}

		var_cee6d53bc023_mapped := new(int32)
		*var_cee6d53bc023_mapped = val.(int32)

		s.Offset = var_cee6d53bc023_mapped
	}
	if properties["resolveReferences"] != nil {

		var_6f217a9d01a6 := properties["resolveReferences"]
		var_6f217a9d01a6_mapped := []string{}
		for _, v := range var_6f217a9d01a6.GetListValue().Values {

			var_08d39c34121b := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_08d39c34121b)

			if err != nil {
				panic(err)
			}

			var_08d39c34121b_mapped := val.(string)

			var_6f217a9d01a6_mapped = append(var_6f217a9d01a6_mapped, var_08d39c34121b_mapped)
		}

		s.ResolveReferences = var_6f217a9d01a6_mapped
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

	var_f8142d0c2c8d := extensionEvent.Id

	if var_f8142d0c2c8d != nil {
		var var_f8142d0c2c8d_mapped *structpb.Value

		var var_f8142d0c2c8d_err error
		var_f8142d0c2c8d_mapped, var_f8142d0c2c8d_err = types.ByResourcePropertyType(model.ResourceProperty_UUID).Pack(*var_f8142d0c2c8d)
		if var_f8142d0c2c8d_err != nil {
			panic(var_f8142d0c2c8d_err)
		}
		properties["id"] = var_f8142d0c2c8d_mapped
	}

	var_112b934cb002 := extensionEvent.Action

	var var_112b934cb002_mapped *structpb.Value

	var var_112b934cb002_err error
	var_112b934cb002_mapped, var_112b934cb002_err = types.ByResourcePropertyType(model.ResourceProperty_ENUM).Pack(string(var_112b934cb002))
	if var_112b934cb002_err != nil {
		panic(var_112b934cb002_err)
	}
	properties["action"] = var_112b934cb002_mapped

	var_b0f8518322d4 := extensionEvent.RecordSearchParams

	if var_b0f8518322d4 != nil {
		var var_b0f8518322d4_mapped *structpb.Value

		var_b0f8518322d4_mapped = structpb.NewStructValue(&structpb.Struct{Fields: ExtensionRecordSearchParamsMapperInstance.ToProperties(var_b0f8518322d4)})
		properties["recordSearchParams"] = var_b0f8518322d4_mapped
	}

	var_ecac1f879747 := extensionEvent.ActionSummary

	if var_ecac1f879747 != nil {
		var var_ecac1f879747_mapped *structpb.Value

		var var_ecac1f879747_err error
		var_ecac1f879747_mapped, var_ecac1f879747_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_ecac1f879747)
		if var_ecac1f879747_err != nil {
			panic(var_ecac1f879747_err)
		}
		properties["actionSummary"] = var_ecac1f879747_mapped
	}

	var_a1cec7bb92f0 := extensionEvent.ActionDescription

	if var_a1cec7bb92f0 != nil {
		var var_a1cec7bb92f0_mapped *structpb.Value

		var var_a1cec7bb92f0_err error
		var_a1cec7bb92f0_mapped, var_a1cec7bb92f0_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_a1cec7bb92f0)
		if var_a1cec7bb92f0_err != nil {
			panic(var_a1cec7bb92f0_err)
		}
		properties["actionDescription"] = var_a1cec7bb92f0_mapped
	}

	var_e84c99a2ded4 := extensionEvent.Resource

	if var_e84c99a2ded4 != nil {
		var var_e84c99a2ded4_mapped *structpb.Value

		var_e84c99a2ded4_mapped = structpb.NewStructValue(&structpb.Struct{Fields: ResourceMapperInstance.ToProperties(var_e84c99a2ded4)})
		properties["resource"] = var_e84c99a2ded4_mapped
	}

	var_b2eb1c2d73e1 := extensionEvent.Records

	if var_b2eb1c2d73e1 != nil {
		var var_b2eb1c2d73e1_mapped *structpb.Value

		var var_b2eb1c2d73e1_l []*structpb.Value
		for _, value := range var_b2eb1c2d73e1 {

			var_83f4e0a15ac1 := value
			var var_83f4e0a15ac1_mapped *structpb.Value

			var_83f4e0a15ac1_mapped = structpb.NewStructValue(&structpb.Struct{Fields: RecordMapperInstance.ToProperties(var_83f4e0a15ac1)})

			var_b2eb1c2d73e1_l = append(var_b2eb1c2d73e1_l, var_83f4e0a15ac1_mapped)
		}
		var_b2eb1c2d73e1_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_b2eb1c2d73e1_l})
		properties["records"] = var_b2eb1c2d73e1_mapped
	}

	var_7d6955964297 := extensionEvent.Ids

	if var_7d6955964297 != nil {
		var var_7d6955964297_mapped *structpb.Value

		var var_7d6955964297_l []*structpb.Value
		for _, value := range var_7d6955964297 {

			var_ad6336af6b30 := value
			var var_ad6336af6b30_mapped *structpb.Value

			var var_ad6336af6b30_err error
			var_ad6336af6b30_mapped, var_ad6336af6b30_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_ad6336af6b30)
			if var_ad6336af6b30_err != nil {
				panic(var_ad6336af6b30_err)
			}

			var_7d6955964297_l = append(var_7d6955964297_l, var_ad6336af6b30_mapped)
		}
		var_7d6955964297_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_7d6955964297_l})
		properties["ids"] = var_7d6955964297_mapped
	}

	var_bd48b41bd9ed := extensionEvent.Finalizes

	if var_bd48b41bd9ed != nil {
		var var_bd48b41bd9ed_mapped *structpb.Value

		var var_bd48b41bd9ed_err error
		var_bd48b41bd9ed_mapped, var_bd48b41bd9ed_err = types.ByResourcePropertyType(model.ResourceProperty_BOOL).Pack(*var_bd48b41bd9ed)
		if var_bd48b41bd9ed_err != nil {
			panic(var_bd48b41bd9ed_err)
		}
		properties["finalizes"] = var_bd48b41bd9ed_mapped
	}

	var_97e89f5db56a := extensionEvent.Sync

	if var_97e89f5db56a != nil {
		var var_97e89f5db56a_mapped *structpb.Value

		var var_97e89f5db56a_err error
		var_97e89f5db56a_mapped, var_97e89f5db56a_err = types.ByResourcePropertyType(model.ResourceProperty_BOOL).Pack(*var_97e89f5db56a)
		if var_97e89f5db56a_err != nil {
			panic(var_97e89f5db56a_err)
		}
		properties["sync"] = var_97e89f5db56a_mapped
	}

	var_43bf8c86b4ee := extensionEvent.Time

	if var_43bf8c86b4ee != nil {
		var var_43bf8c86b4ee_mapped *structpb.Value

		var var_43bf8c86b4ee_err error
		var_43bf8c86b4ee_mapped, var_43bf8c86b4ee_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_43bf8c86b4ee)
		if var_43bf8c86b4ee_err != nil {
			panic(var_43bf8c86b4ee_err)
		}
		properties["time"] = var_43bf8c86b4ee_mapped
	}

	var_326a177c61e4 := extensionEvent.Annotations

	if var_326a177c61e4 != nil {
		var var_326a177c61e4_mapped *structpb.Value

		var var_326a177c61e4_st *structpb.Struct = new(structpb.Struct)
		var_326a177c61e4_st.Fields = make(map[string]*structpb.Value)
		for key, value := range var_326a177c61e4 {

			var_7ad69d465bb2 := value
			var var_7ad69d465bb2_mapped *structpb.Value

			var var_7ad69d465bb2_err error
			var_7ad69d465bb2_mapped, var_7ad69d465bb2_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_7ad69d465bb2)
			if var_7ad69d465bb2_err != nil {
				panic(var_7ad69d465bb2_err)
			}

			var_326a177c61e4_st.Fields[key] = var_7ad69d465bb2_mapped
		}
		var_326a177c61e4_mapped = structpb.NewStructValue(var_326a177c61e4_st)
		properties["annotations"] = var_326a177c61e4_mapped
	}
	return properties
}

func (m *ExtensionEventMapper) FromProperties(properties map[string]*structpb.Value) *ExtensionEvent {
	var s = m.New()
	if properties["id"] != nil {

		var_ab9a3b83469b := properties["id"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_UUID).UnPack(var_ab9a3b83469b)

		if err != nil {
			panic(err)
		}

		var_ab9a3b83469b_mapped := new(uuid.UUID)
		*var_ab9a3b83469b_mapped = val.(uuid.UUID)

		s.Id = var_ab9a3b83469b_mapped
	}
	if properties["action"] != nil {

		var_68161a35deff := properties["action"]
		var_68161a35deff_mapped := (EventAction)(var_68161a35deff.GetStringValue())

		s.Action = var_68161a35deff_mapped
	}
	if properties["recordSearchParams"] != nil {

		var_65151f3e180e := properties["recordSearchParams"]
		var mappedValue = ExtensionRecordSearchParamsMapperInstance.FromProperties(var_65151f3e180e.GetStructValue().Fields)

		var_65151f3e180e_mapped := mappedValue

		s.RecordSearchParams = var_65151f3e180e_mapped
	}
	if properties["actionSummary"] != nil {

		var_9bc4cb9c336e := properties["actionSummary"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_9bc4cb9c336e)

		if err != nil {
			panic(err)
		}

		var_9bc4cb9c336e_mapped := new(string)
		*var_9bc4cb9c336e_mapped = val.(string)

		s.ActionSummary = var_9bc4cb9c336e_mapped
	}
	if properties["actionDescription"] != nil {

		var_51c17572091e := properties["actionDescription"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_51c17572091e)

		if err != nil {
			panic(err)
		}

		var_51c17572091e_mapped := new(string)
		*var_51c17572091e_mapped = val.(string)

		s.ActionDescription = var_51c17572091e_mapped
	}
	if properties["resource"] != nil {

		var_37397265fe63 := properties["resource"]
		var_37397265fe63_mapped := ResourceMapperInstance.FromProperties(var_37397265fe63.GetStructValue().Fields)

		s.Resource = var_37397265fe63_mapped
	}
	if properties["records"] != nil {

		var_836bcae458a9 := properties["records"]
		var_836bcae458a9_mapped := []*Record{}
		for _, v := range var_836bcae458a9.GetListValue().Values {

			var_85e487748681 := v
			var_85e487748681_mapped := RecordMapperInstance.FromProperties(var_85e487748681.GetStructValue().Fields)

			var_836bcae458a9_mapped = append(var_836bcae458a9_mapped, var_85e487748681_mapped)
		}

		s.Records = var_836bcae458a9_mapped
	}
	if properties["ids"] != nil {

		var_e944779e014d := properties["ids"]
		var_e944779e014d_mapped := []string{}
		for _, v := range var_e944779e014d.GetListValue().Values {

			var_ac55fbddf13f := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_ac55fbddf13f)

			if err != nil {
				panic(err)
			}

			var_ac55fbddf13f_mapped := val.(string)

			var_e944779e014d_mapped = append(var_e944779e014d_mapped, var_ac55fbddf13f_mapped)
		}

		s.Ids = var_e944779e014d_mapped
	}
	if properties["finalizes"] != nil {

		var_6f2fdf2cac6c := properties["finalizes"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_BOOL).UnPack(var_6f2fdf2cac6c)

		if err != nil {
			panic(err)
		}

		var_6f2fdf2cac6c_mapped := new(bool)
		*var_6f2fdf2cac6c_mapped = val.(bool)

		s.Finalizes = var_6f2fdf2cac6c_mapped
	}
	if properties["sync"] != nil {

		var_3c459910bf7a := properties["sync"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_BOOL).UnPack(var_3c459910bf7a)

		if err != nil {
			panic(err)
		}

		var_3c459910bf7a_mapped := new(bool)
		*var_3c459910bf7a_mapped = val.(bool)

		s.Sync = var_3c459910bf7a_mapped
	}
	if properties["time"] != nil {

		var_4bd72e5b788f := properties["time"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_4bd72e5b788f)

		if err != nil {
			panic(err)
		}

		var_4bd72e5b788f_mapped := new(time.Time)
		*var_4bd72e5b788f_mapped = val.(time.Time)

		s.Time = var_4bd72e5b788f_mapped
	}
	if properties["annotations"] != nil {

		var_3b219454b15b := properties["annotations"]
		var_3b219454b15b_mapped := make(map[string]string)
		for k, v := range var_3b219454b15b.GetStructValue().Fields {

			var_d46024ad6e29 := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_d46024ad6e29)

			if err != nil {
				panic(err)
			}

			var_d46024ad6e29_mapped := val.(string)

			var_3b219454b15b_mapped[k] = var_d46024ad6e29_mapped
		}

		s.Annotations = var_3b219454b15b_mapped
	}
	return s
}
