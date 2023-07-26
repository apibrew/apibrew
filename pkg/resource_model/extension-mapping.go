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

	var_92d2b2ce0410 := extension.Id

	if var_92d2b2ce0410 != nil {
		var var_92d2b2ce0410_mapped *structpb.Value

		var var_92d2b2ce0410_err error
		var_92d2b2ce0410_mapped, var_92d2b2ce0410_err = types.ByResourcePropertyType(model.ResourceProperty_UUID).Pack(*var_92d2b2ce0410)
		if var_92d2b2ce0410_err != nil {
			panic(var_92d2b2ce0410_err)
		}
		properties["id"] = var_92d2b2ce0410_mapped
	}

	var_52e38d7b5373 := extension.Version

	var var_52e38d7b5373_mapped *structpb.Value

	var var_52e38d7b5373_err error
	var_52e38d7b5373_mapped, var_52e38d7b5373_err = types.ByResourcePropertyType(model.ResourceProperty_INT32).Pack(var_52e38d7b5373)
	if var_52e38d7b5373_err != nil {
		panic(var_52e38d7b5373_err)
	}
	properties["version"] = var_52e38d7b5373_mapped

	var_dcdf3f99d3a9 := extension.CreatedBy

	if var_dcdf3f99d3a9 != nil {
		var var_dcdf3f99d3a9_mapped *structpb.Value

		var var_dcdf3f99d3a9_err error
		var_dcdf3f99d3a9_mapped, var_dcdf3f99d3a9_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_dcdf3f99d3a9)
		if var_dcdf3f99d3a9_err != nil {
			panic(var_dcdf3f99d3a9_err)
		}
		properties["createdBy"] = var_dcdf3f99d3a9_mapped
	}

	var_6e9d3d95c86a := extension.UpdatedBy

	if var_6e9d3d95c86a != nil {
		var var_6e9d3d95c86a_mapped *structpb.Value

		var var_6e9d3d95c86a_err error
		var_6e9d3d95c86a_mapped, var_6e9d3d95c86a_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_6e9d3d95c86a)
		if var_6e9d3d95c86a_err != nil {
			panic(var_6e9d3d95c86a_err)
		}
		properties["updatedBy"] = var_6e9d3d95c86a_mapped
	}

	var_a2924697fbe9 := extension.CreatedOn

	if var_a2924697fbe9 != nil {
		var var_a2924697fbe9_mapped *structpb.Value

		var var_a2924697fbe9_err error
		var_a2924697fbe9_mapped, var_a2924697fbe9_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_a2924697fbe9)
		if var_a2924697fbe9_err != nil {
			panic(var_a2924697fbe9_err)
		}
		properties["createdOn"] = var_a2924697fbe9_mapped
	}

	var_39442422ba31 := extension.UpdatedOn

	if var_39442422ba31 != nil {
		var var_39442422ba31_mapped *structpb.Value

		var var_39442422ba31_err error
		var_39442422ba31_mapped, var_39442422ba31_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_39442422ba31)
		if var_39442422ba31_err != nil {
			panic(var_39442422ba31_err)
		}
		properties["updatedOn"] = var_39442422ba31_mapped
	}

	var_680e1cdcc414 := extension.Name

	var var_680e1cdcc414_mapped *structpb.Value

	var var_680e1cdcc414_err error
	var_680e1cdcc414_mapped, var_680e1cdcc414_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_680e1cdcc414)
	if var_680e1cdcc414_err != nil {
		panic(var_680e1cdcc414_err)
	}
	properties["name"] = var_680e1cdcc414_mapped

	var_389332804df2 := extension.Description

	if var_389332804df2 != nil {
		var var_389332804df2_mapped *structpb.Value

		var var_389332804df2_err error
		var_389332804df2_mapped, var_389332804df2_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_389332804df2)
		if var_389332804df2_err != nil {
			panic(var_389332804df2_err)
		}
		properties["description"] = var_389332804df2_mapped
	}

	var_c045957e3e80 := extension.Selector

	if var_c045957e3e80 != nil {
		var var_c045957e3e80_mapped *structpb.Value

		var_c045957e3e80_mapped = structpb.NewStructValue(&structpb.Struct{Fields: ExtensionEventSelectorMapperInstance.ToProperties(var_c045957e3e80)})
		properties["selector"] = var_c045957e3e80_mapped
	}

	var_3b03043024ac := extension.Order

	var var_3b03043024ac_mapped *structpb.Value

	var var_3b03043024ac_err error
	var_3b03043024ac_mapped, var_3b03043024ac_err = types.ByResourcePropertyType(model.ResourceProperty_INT32).Pack(var_3b03043024ac)
	if var_3b03043024ac_err != nil {
		panic(var_3b03043024ac_err)
	}
	properties["order"] = var_3b03043024ac_mapped

	var_4358c61b2286 := extension.Finalizes

	var var_4358c61b2286_mapped *structpb.Value

	var var_4358c61b2286_err error
	var_4358c61b2286_mapped, var_4358c61b2286_err = types.ByResourcePropertyType(model.ResourceProperty_BOOL).Pack(var_4358c61b2286)
	if var_4358c61b2286_err != nil {
		panic(var_4358c61b2286_err)
	}
	properties["finalizes"] = var_4358c61b2286_mapped

	var_df8f020604af := extension.Sync

	var var_df8f020604af_mapped *structpb.Value

	var var_df8f020604af_err error
	var_df8f020604af_mapped, var_df8f020604af_err = types.ByResourcePropertyType(model.ResourceProperty_BOOL).Pack(var_df8f020604af)
	if var_df8f020604af_err != nil {
		panic(var_df8f020604af_err)
	}
	properties["sync"] = var_df8f020604af_mapped

	var_92790fa93c51 := extension.Responds

	var var_92790fa93c51_mapped *structpb.Value

	var var_92790fa93c51_err error
	var_92790fa93c51_mapped, var_92790fa93c51_err = types.ByResourcePropertyType(model.ResourceProperty_BOOL).Pack(var_92790fa93c51)
	if var_92790fa93c51_err != nil {
		panic(var_92790fa93c51_err)
	}
	properties["responds"] = var_92790fa93c51_mapped

	var_3ed8d8a8d534 := extension.Call

	var var_3ed8d8a8d534_mapped *structpb.Value

	var_3ed8d8a8d534_mapped = structpb.NewStructValue(&structpb.Struct{Fields: ExtensionExternalCallMapperInstance.ToProperties(&var_3ed8d8a8d534)})
	properties["call"] = var_3ed8d8a8d534_mapped

	var_30288944b6f5 := extension.Annotations

	if var_30288944b6f5 != nil {
		var var_30288944b6f5_mapped *structpb.Value

		var var_30288944b6f5_st *structpb.Struct = new(structpb.Struct)
		var_30288944b6f5_st.Fields = make(map[string]*structpb.Value)
		for key, value := range var_30288944b6f5 {

			var_46591127974f := value
			var var_46591127974f_mapped *structpb.Value

			var var_46591127974f_err error
			var_46591127974f_mapped, var_46591127974f_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_46591127974f)
			if var_46591127974f_err != nil {
				panic(var_46591127974f_err)
			}

			var_30288944b6f5_st.Fields[key] = var_46591127974f_mapped
		}
		var_30288944b6f5_mapped = structpb.NewStructValue(var_30288944b6f5_st)
		properties["annotations"] = var_30288944b6f5_mapped
	}
	return properties
}

func (m *ExtensionMapper) FromProperties(properties map[string]*structpb.Value) *Extension {
	var s = m.New()
	if properties["id"] != nil {

		var_af4f043feb8a := properties["id"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_UUID).UnPack(var_af4f043feb8a)

		if err != nil {
			panic(err)
		}

		var_af4f043feb8a_mapped := new(uuid.UUID)
		*var_af4f043feb8a_mapped = val.(uuid.UUID)

		s.Id = var_af4f043feb8a_mapped
	}
	if properties["version"] != nil {

		var_402efc95c592 := properties["version"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(var_402efc95c592)

		if err != nil {
			panic(err)
		}

		var_402efc95c592_mapped := val.(int32)

		s.Version = var_402efc95c592_mapped
	}
	if properties["createdBy"] != nil {

		var_52d3af2f2537 := properties["createdBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_52d3af2f2537)

		if err != nil {
			panic(err)
		}

		var_52d3af2f2537_mapped := new(string)
		*var_52d3af2f2537_mapped = val.(string)

		s.CreatedBy = var_52d3af2f2537_mapped
	}
	if properties["updatedBy"] != nil {

		var_6ae6c3f8d003 := properties["updatedBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_6ae6c3f8d003)

		if err != nil {
			panic(err)
		}

		var_6ae6c3f8d003_mapped := new(string)
		*var_6ae6c3f8d003_mapped = val.(string)

		s.UpdatedBy = var_6ae6c3f8d003_mapped
	}
	if properties["createdOn"] != nil {

		var_549da032d86a := properties["createdOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_549da032d86a)

		if err != nil {
			panic(err)
		}

		var_549da032d86a_mapped := new(time.Time)
		*var_549da032d86a_mapped = val.(time.Time)

		s.CreatedOn = var_549da032d86a_mapped
	}
	if properties["updatedOn"] != nil {

		var_d29215274e13 := properties["updatedOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_d29215274e13)

		if err != nil {
			panic(err)
		}

		var_d29215274e13_mapped := new(time.Time)
		*var_d29215274e13_mapped = val.(time.Time)

		s.UpdatedOn = var_d29215274e13_mapped
	}
	if properties["name"] != nil {

		var_47310a874d6e := properties["name"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_47310a874d6e)

		if err != nil {
			panic(err)
		}

		var_47310a874d6e_mapped := val.(string)

		s.Name = var_47310a874d6e_mapped
	}
	if properties["description"] != nil {

		var_451bff5d1b12 := properties["description"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_451bff5d1b12)

		if err != nil {
			panic(err)
		}

		var_451bff5d1b12_mapped := new(string)
		*var_451bff5d1b12_mapped = val.(string)

		s.Description = var_451bff5d1b12_mapped
	}
	if properties["selector"] != nil {

		var_b75f25e2c90b := properties["selector"]
		var mappedValue = ExtensionEventSelectorMapperInstance.FromProperties(var_b75f25e2c90b.GetStructValue().Fields)

		var_b75f25e2c90b_mapped := mappedValue

		s.Selector = var_b75f25e2c90b_mapped
	}
	if properties["order"] != nil {

		var_5bc9e0a8c2d7 := properties["order"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(var_5bc9e0a8c2d7)

		if err != nil {
			panic(err)
		}

		var_5bc9e0a8c2d7_mapped := val.(int32)

		s.Order = var_5bc9e0a8c2d7_mapped
	}
	if properties["finalizes"] != nil {

		var_49d987bcbb07 := properties["finalizes"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_BOOL).UnPack(var_49d987bcbb07)

		if err != nil {
			panic(err)
		}

		var_49d987bcbb07_mapped := val.(bool)

		s.Finalizes = var_49d987bcbb07_mapped
	}
	if properties["sync"] != nil {

		var_cd88ce426ec6 := properties["sync"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_BOOL).UnPack(var_cd88ce426ec6)

		if err != nil {
			panic(err)
		}

		var_cd88ce426ec6_mapped := val.(bool)

		s.Sync = var_cd88ce426ec6_mapped
	}
	if properties["responds"] != nil {

		var_297668d13735 := properties["responds"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_BOOL).UnPack(var_297668d13735)

		if err != nil {
			panic(err)
		}

		var_297668d13735_mapped := val.(bool)

		s.Responds = var_297668d13735_mapped
	}
	if properties["call"] != nil {

		var_879c42668b89 := properties["call"]
		var mappedValue = ExtensionExternalCallMapperInstance.FromProperties(var_879c42668b89.GetStructValue().Fields)

		var_879c42668b89_mapped := *mappedValue

		s.Call = var_879c42668b89_mapped
	}
	if properties["annotations"] != nil {

		var_7087213a30bd := properties["annotations"]
		var_7087213a30bd_mapped := make(map[string]string)
		for k, v := range var_7087213a30bd.GetStructValue().Fields {

			var_8ed13b315fd1 := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_8ed13b315fd1)

			if err != nil {
				panic(err)
			}

			var_8ed13b315fd1_mapped := val.(string)

			var_7087213a30bd_mapped[k] = var_8ed13b315fd1_mapped
		}

		s.Annotations = var_7087213a30bd_mapped
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

	var_7769357621c3 := extensionFunctionCall.Host

	var var_7769357621c3_mapped *structpb.Value

	var var_7769357621c3_err error
	var_7769357621c3_mapped, var_7769357621c3_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_7769357621c3)
	if var_7769357621c3_err != nil {
		panic(var_7769357621c3_err)
	}
	properties["host"] = var_7769357621c3_mapped

	var_b61bcb2afde7 := extensionFunctionCall.FunctionName

	var var_b61bcb2afde7_mapped *structpb.Value

	var var_b61bcb2afde7_err error
	var_b61bcb2afde7_mapped, var_b61bcb2afde7_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_b61bcb2afde7)
	if var_b61bcb2afde7_err != nil {
		panic(var_b61bcb2afde7_err)
	}
	properties["functionName"] = var_b61bcb2afde7_mapped
	return properties
}

func (m *ExtensionFunctionCallMapper) FromProperties(properties map[string]*structpb.Value) *ExtensionFunctionCall {
	var s = m.New()
	if properties["host"] != nil {

		var_a2af95b2ea32 := properties["host"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_a2af95b2ea32)

		if err != nil {
			panic(err)
		}

		var_a2af95b2ea32_mapped := val.(string)

		s.Host = var_a2af95b2ea32_mapped
	}
	if properties["functionName"] != nil {

		var_207c926a712c := properties["functionName"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_207c926a712c)

		if err != nil {
			panic(err)
		}

		var_207c926a712c_mapped := val.(string)

		s.FunctionName = var_207c926a712c_mapped
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

	var_f4dbea83342d := extensionHttpCall.Uri

	var var_f4dbea83342d_mapped *structpb.Value

	var var_f4dbea83342d_err error
	var_f4dbea83342d_mapped, var_f4dbea83342d_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_f4dbea83342d)
	if var_f4dbea83342d_err != nil {
		panic(var_f4dbea83342d_err)
	}
	properties["uri"] = var_f4dbea83342d_mapped

	var_20077cbe1e18 := extensionHttpCall.Method

	var var_20077cbe1e18_mapped *structpb.Value

	var var_20077cbe1e18_err error
	var_20077cbe1e18_mapped, var_20077cbe1e18_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_20077cbe1e18)
	if var_20077cbe1e18_err != nil {
		panic(var_20077cbe1e18_err)
	}
	properties["method"] = var_20077cbe1e18_mapped
	return properties
}

func (m *ExtensionHttpCallMapper) FromProperties(properties map[string]*structpb.Value) *ExtensionHttpCall {
	var s = m.New()
	if properties["uri"] != nil {

		var_6b85782eda06 := properties["uri"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_6b85782eda06)

		if err != nil {
			panic(err)
		}

		var_6b85782eda06_mapped := val.(string)

		s.Uri = var_6b85782eda06_mapped
	}
	if properties["method"] != nil {

		var_a7265a395b69 := properties["method"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_a7265a395b69)

		if err != nil {
			panic(err)
		}

		var_a7265a395b69_mapped := val.(string)

		s.Method = var_a7265a395b69_mapped
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

	var_d80ebe8cb9c4 := extensionExternalCall.FunctionCall

	if var_d80ebe8cb9c4 != nil {
		var var_d80ebe8cb9c4_mapped *structpb.Value

		var_d80ebe8cb9c4_mapped = structpb.NewStructValue(&structpb.Struct{Fields: ExtensionFunctionCallMapperInstance.ToProperties(var_d80ebe8cb9c4)})
		properties["functionCall"] = var_d80ebe8cb9c4_mapped
	}

	var_399e8be1cea8 := extensionExternalCall.HttpCall

	if var_399e8be1cea8 != nil {
		var var_399e8be1cea8_mapped *structpb.Value

		var_399e8be1cea8_mapped = structpb.NewStructValue(&structpb.Struct{Fields: ExtensionHttpCallMapperInstance.ToProperties(var_399e8be1cea8)})
		properties["httpCall"] = var_399e8be1cea8_mapped
	}
	return properties
}

func (m *ExtensionExternalCallMapper) FromProperties(properties map[string]*structpb.Value) *ExtensionExternalCall {
	var s = m.New()
	if properties["functionCall"] != nil {

		var_19c617c37c75 := properties["functionCall"]
		var mappedValue = ExtensionFunctionCallMapperInstance.FromProperties(var_19c617c37c75.GetStructValue().Fields)

		var_19c617c37c75_mapped := mappedValue

		s.FunctionCall = var_19c617c37c75_mapped
	}
	if properties["httpCall"] != nil {

		var_2b56e5b0a77e := properties["httpCall"]
		var mappedValue = ExtensionHttpCallMapperInstance.FromProperties(var_2b56e5b0a77e.GetStructValue().Fields)

		var_2b56e5b0a77e_mapped := mappedValue

		s.HttpCall = var_2b56e5b0a77e_mapped
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

	var_6c9309df6132 := extensionEventSelector.Actions

	if var_6c9309df6132 != nil {
		var var_6c9309df6132_mapped *structpb.Value

		var var_6c9309df6132_l []*structpb.Value
		for _, value := range var_6c9309df6132 {

			var_21d3eca9d23a := value
			var var_21d3eca9d23a_mapped *structpb.Value

			var var_21d3eca9d23a_err error
			var_21d3eca9d23a_mapped, var_21d3eca9d23a_err = types.ByResourcePropertyType(model.ResourceProperty_ENUM).Pack(string(var_21d3eca9d23a))
			if var_21d3eca9d23a_err != nil {
				panic(var_21d3eca9d23a_err)
			}

			var_6c9309df6132_l = append(var_6c9309df6132_l, var_21d3eca9d23a_mapped)
		}
		var_6c9309df6132_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_6c9309df6132_l})
		properties["actions"] = var_6c9309df6132_mapped
	}

	var_25ebb886ce20 := extensionEventSelector.RecordSelector

	if var_25ebb886ce20 != nil {
		var var_25ebb886ce20_mapped *structpb.Value

		var_25ebb886ce20_mapped = structpb.NewStructValue(&structpb.Struct{Fields: ExtensionBooleanExpressionMapperInstance.ToProperties(var_25ebb886ce20)})
		properties["recordSelector"] = var_25ebb886ce20_mapped
	}

	var_ee3e10eb2471 := extensionEventSelector.Namespaces

	if var_ee3e10eb2471 != nil {
		var var_ee3e10eb2471_mapped *structpb.Value

		var var_ee3e10eb2471_l []*structpb.Value
		for _, value := range var_ee3e10eb2471 {

			var_7a6b94cb3eec := value
			var var_7a6b94cb3eec_mapped *structpb.Value

			var var_7a6b94cb3eec_err error
			var_7a6b94cb3eec_mapped, var_7a6b94cb3eec_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_7a6b94cb3eec)
			if var_7a6b94cb3eec_err != nil {
				panic(var_7a6b94cb3eec_err)
			}

			var_ee3e10eb2471_l = append(var_ee3e10eb2471_l, var_7a6b94cb3eec_mapped)
		}
		var_ee3e10eb2471_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_ee3e10eb2471_l})
		properties["namespaces"] = var_ee3e10eb2471_mapped
	}

	var_3f1155e39780 := extensionEventSelector.Resources

	if var_3f1155e39780 != nil {
		var var_3f1155e39780_mapped *structpb.Value

		var var_3f1155e39780_l []*structpb.Value
		for _, value := range var_3f1155e39780 {

			var_4d08726e0608 := value
			var var_4d08726e0608_mapped *structpb.Value

			var var_4d08726e0608_err error
			var_4d08726e0608_mapped, var_4d08726e0608_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_4d08726e0608)
			if var_4d08726e0608_err != nil {
				panic(var_4d08726e0608_err)
			}

			var_3f1155e39780_l = append(var_3f1155e39780_l, var_4d08726e0608_mapped)
		}
		var_3f1155e39780_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_3f1155e39780_l})
		properties["resources"] = var_3f1155e39780_mapped
	}

	var_a50c481ecc14 := extensionEventSelector.Ids

	if var_a50c481ecc14 != nil {
		var var_a50c481ecc14_mapped *structpb.Value

		var var_a50c481ecc14_l []*structpb.Value
		for _, value := range var_a50c481ecc14 {

			var_0187d994b3a0 := value
			var var_0187d994b3a0_mapped *structpb.Value

			var var_0187d994b3a0_err error
			var_0187d994b3a0_mapped, var_0187d994b3a0_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_0187d994b3a0)
			if var_0187d994b3a0_err != nil {
				panic(var_0187d994b3a0_err)
			}

			var_a50c481ecc14_l = append(var_a50c481ecc14_l, var_0187d994b3a0_mapped)
		}
		var_a50c481ecc14_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_a50c481ecc14_l})
		properties["ids"] = var_a50c481ecc14_mapped
	}

	var_e61e342b62b6 := extensionEventSelector.Annotations

	if var_e61e342b62b6 != nil {
		var var_e61e342b62b6_mapped *structpb.Value

		var var_e61e342b62b6_st *structpb.Struct = new(structpb.Struct)
		var_e61e342b62b6_st.Fields = make(map[string]*structpb.Value)
		for key, value := range var_e61e342b62b6 {

			var_0fe451d4cc95 := value
			var var_0fe451d4cc95_mapped *structpb.Value

			var var_0fe451d4cc95_err error
			var_0fe451d4cc95_mapped, var_0fe451d4cc95_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_0fe451d4cc95)
			if var_0fe451d4cc95_err != nil {
				panic(var_0fe451d4cc95_err)
			}

			var_e61e342b62b6_st.Fields[key] = var_0fe451d4cc95_mapped
		}
		var_e61e342b62b6_mapped = structpb.NewStructValue(var_e61e342b62b6_st)
		properties["annotations"] = var_e61e342b62b6_mapped
	}
	return properties
}

func (m *ExtensionEventSelectorMapper) FromProperties(properties map[string]*structpb.Value) *ExtensionEventSelector {
	var s = m.New()
	if properties["actions"] != nil {

		var_8bb04bbed56e := properties["actions"]
		var_8bb04bbed56e_mapped := []EventAction{}
		for _, v := range var_8bb04bbed56e.GetListValue().Values {

			var_f6a1660005c5 := v
			var_f6a1660005c5_mapped := (EventAction)(var_f6a1660005c5.GetStringValue())

			var_8bb04bbed56e_mapped = append(var_8bb04bbed56e_mapped, var_f6a1660005c5_mapped)
		}

		s.Actions = var_8bb04bbed56e_mapped
	}
	if properties["recordSelector"] != nil {

		var_75f7edc263b7 := properties["recordSelector"]
		var mappedValue = ExtensionBooleanExpressionMapperInstance.FromProperties(var_75f7edc263b7.GetStructValue().Fields)

		var_75f7edc263b7_mapped := mappedValue

		s.RecordSelector = var_75f7edc263b7_mapped
	}
	if properties["namespaces"] != nil {

		var_4a7927d62677 := properties["namespaces"]
		var_4a7927d62677_mapped := []string{}
		for _, v := range var_4a7927d62677.GetListValue().Values {

			var_9fada13cfdf1 := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_9fada13cfdf1)

			if err != nil {
				panic(err)
			}

			var_9fada13cfdf1_mapped := val.(string)

			var_4a7927d62677_mapped = append(var_4a7927d62677_mapped, var_9fada13cfdf1_mapped)
		}

		s.Namespaces = var_4a7927d62677_mapped
	}
	if properties["resources"] != nil {

		var_83e376045d43 := properties["resources"]
		var_83e376045d43_mapped := []string{}
		for _, v := range var_83e376045d43.GetListValue().Values {

			var_652c9966e20f := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_652c9966e20f)

			if err != nil {
				panic(err)
			}

			var_652c9966e20f_mapped := val.(string)

			var_83e376045d43_mapped = append(var_83e376045d43_mapped, var_652c9966e20f_mapped)
		}

		s.Resources = var_83e376045d43_mapped
	}
	if properties["ids"] != nil {

		var_a6ff26e2bbdd := properties["ids"]
		var_a6ff26e2bbdd_mapped := []string{}
		for _, v := range var_a6ff26e2bbdd.GetListValue().Values {

			var_2048d79b6d5a := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_2048d79b6d5a)

			if err != nil {
				panic(err)
			}

			var_2048d79b6d5a_mapped := val.(string)

			var_a6ff26e2bbdd_mapped = append(var_a6ff26e2bbdd_mapped, var_2048d79b6d5a_mapped)
		}

		s.Ids = var_a6ff26e2bbdd_mapped
	}
	if properties["annotations"] != nil {

		var_1da91857f280 := properties["annotations"]
		var_1da91857f280_mapped := make(map[string]string)
		for k, v := range var_1da91857f280.GetStructValue().Fields {

			var_398a68a05173 := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_398a68a05173)

			if err != nil {
				panic(err)
			}

			var_398a68a05173_mapped := val.(string)

			var_1da91857f280_mapped[k] = var_398a68a05173_mapped
		}

		s.Annotations = var_1da91857f280_mapped
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

	var_c606ac1ff014 := extensionRecordSearchParams.Query

	if var_c606ac1ff014 != nil {
		var var_c606ac1ff014_mapped *structpb.Value

		var_c606ac1ff014_mapped = structpb.NewStructValue(&structpb.Struct{Fields: ExtensionBooleanExpressionMapperInstance.ToProperties(var_c606ac1ff014)})
		properties["query"] = var_c606ac1ff014_mapped
	}

	var_005125edcb6d := extensionRecordSearchParams.Limit

	if var_005125edcb6d != nil {
		var var_005125edcb6d_mapped *structpb.Value

		var var_005125edcb6d_err error
		var_005125edcb6d_mapped, var_005125edcb6d_err = types.ByResourcePropertyType(model.ResourceProperty_INT32).Pack(*var_005125edcb6d)
		if var_005125edcb6d_err != nil {
			panic(var_005125edcb6d_err)
		}
		properties["limit"] = var_005125edcb6d_mapped
	}

	var_9f66f9f0974e := extensionRecordSearchParams.Offset

	if var_9f66f9f0974e != nil {
		var var_9f66f9f0974e_mapped *structpb.Value

		var var_9f66f9f0974e_err error
		var_9f66f9f0974e_mapped, var_9f66f9f0974e_err = types.ByResourcePropertyType(model.ResourceProperty_INT32).Pack(*var_9f66f9f0974e)
		if var_9f66f9f0974e_err != nil {
			panic(var_9f66f9f0974e_err)
		}
		properties["offset"] = var_9f66f9f0974e_mapped
	}

	var_3fca8a40cda0 := extensionRecordSearchParams.ResolveReferences

	if var_3fca8a40cda0 != nil {
		var var_3fca8a40cda0_mapped *structpb.Value

		var var_3fca8a40cda0_l []*structpb.Value
		for _, value := range var_3fca8a40cda0 {

			var_0c8a2d56f1b2 := value
			var var_0c8a2d56f1b2_mapped *structpb.Value

			var var_0c8a2d56f1b2_err error
			var_0c8a2d56f1b2_mapped, var_0c8a2d56f1b2_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_0c8a2d56f1b2)
			if var_0c8a2d56f1b2_err != nil {
				panic(var_0c8a2d56f1b2_err)
			}

			var_3fca8a40cda0_l = append(var_3fca8a40cda0_l, var_0c8a2d56f1b2_mapped)
		}
		var_3fca8a40cda0_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_3fca8a40cda0_l})
		properties["resolveReferences"] = var_3fca8a40cda0_mapped
	}
	return properties
}

func (m *ExtensionRecordSearchParamsMapper) FromProperties(properties map[string]*structpb.Value) *ExtensionRecordSearchParams {
	var s = m.New()
	if properties["query"] != nil {

		var_e5eecbbb915a := properties["query"]
		var mappedValue = ExtensionBooleanExpressionMapperInstance.FromProperties(var_e5eecbbb915a.GetStructValue().Fields)

		var_e5eecbbb915a_mapped := mappedValue

		s.Query = var_e5eecbbb915a_mapped
	}
	if properties["limit"] != nil {

		var_ecc019853890 := properties["limit"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(var_ecc019853890)

		if err != nil {
			panic(err)
		}

		var_ecc019853890_mapped := new(int32)
		*var_ecc019853890_mapped = val.(int32)

		s.Limit = var_ecc019853890_mapped
	}
	if properties["offset"] != nil {

		var_214219c828e0 := properties["offset"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(var_214219c828e0)

		if err != nil {
			panic(err)
		}

		var_214219c828e0_mapped := new(int32)
		*var_214219c828e0_mapped = val.(int32)

		s.Offset = var_214219c828e0_mapped
	}
	if properties["resolveReferences"] != nil {

		var_e76dd64245f6 := properties["resolveReferences"]
		var_e76dd64245f6_mapped := []string{}
		for _, v := range var_e76dd64245f6.GetListValue().Values {

			var_fb2ed6adf6cb := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_fb2ed6adf6cb)

			if err != nil {
				panic(err)
			}

			var_fb2ed6adf6cb_mapped := val.(string)

			var_e76dd64245f6_mapped = append(var_e76dd64245f6_mapped, var_fb2ed6adf6cb_mapped)
		}

		s.ResolveReferences = var_e76dd64245f6_mapped
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

	var_f4a60b8b45ff := extensionEvent.Id

	if var_f4a60b8b45ff != nil {
		var var_f4a60b8b45ff_mapped *structpb.Value

		var var_f4a60b8b45ff_err error
		var_f4a60b8b45ff_mapped, var_f4a60b8b45ff_err = types.ByResourcePropertyType(model.ResourceProperty_UUID).Pack(*var_f4a60b8b45ff)
		if var_f4a60b8b45ff_err != nil {
			panic(var_f4a60b8b45ff_err)
		}
		properties["id"] = var_f4a60b8b45ff_mapped
	}

	var_48a3c69f2ed0 := extensionEvent.Action

	var var_48a3c69f2ed0_mapped *structpb.Value

	var var_48a3c69f2ed0_err error
	var_48a3c69f2ed0_mapped, var_48a3c69f2ed0_err = types.ByResourcePropertyType(model.ResourceProperty_ENUM).Pack(string(var_48a3c69f2ed0))
	if var_48a3c69f2ed0_err != nil {
		panic(var_48a3c69f2ed0_err)
	}
	properties["action"] = var_48a3c69f2ed0_mapped

	var_bfd8d7f37431 := extensionEvent.RecordSearchParams

	if var_bfd8d7f37431 != nil {
		var var_bfd8d7f37431_mapped *structpb.Value

		var_bfd8d7f37431_mapped = structpb.NewStructValue(&structpb.Struct{Fields: ExtensionRecordSearchParamsMapperInstance.ToProperties(var_bfd8d7f37431)})
		properties["recordSearchParams"] = var_bfd8d7f37431_mapped
	}

	var_67ff23c78f0f := extensionEvent.ActionSummary

	if var_67ff23c78f0f != nil {
		var var_67ff23c78f0f_mapped *structpb.Value

		var var_67ff23c78f0f_err error
		var_67ff23c78f0f_mapped, var_67ff23c78f0f_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_67ff23c78f0f)
		if var_67ff23c78f0f_err != nil {
			panic(var_67ff23c78f0f_err)
		}
		properties["actionSummary"] = var_67ff23c78f0f_mapped
	}

	var_436028126a7c := extensionEvent.ActionDescription

	if var_436028126a7c != nil {
		var var_436028126a7c_mapped *structpb.Value

		var var_436028126a7c_err error
		var_436028126a7c_mapped, var_436028126a7c_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_436028126a7c)
		if var_436028126a7c_err != nil {
			panic(var_436028126a7c_err)
		}
		properties["actionDescription"] = var_436028126a7c_mapped
	}

	var_d0223c1d551c := extensionEvent.Resource

	if var_d0223c1d551c != nil {
		var var_d0223c1d551c_mapped *structpb.Value

		var_d0223c1d551c_mapped = structpb.NewStructValue(&structpb.Struct{Fields: ResourceMapperInstance.ToProperties(var_d0223c1d551c)})
		properties["resource"] = var_d0223c1d551c_mapped
	}

	var_f033415d357c := extensionEvent.Records

	if var_f033415d357c != nil {
		var var_f033415d357c_mapped *structpb.Value

		var var_f033415d357c_l []*structpb.Value
		for _, value := range var_f033415d357c {

			var_3a5a818adf0d := value
			var var_3a5a818adf0d_mapped *structpb.Value

			var_3a5a818adf0d_mapped = structpb.NewStructValue(&structpb.Struct{Fields: RecordMapperInstance.ToProperties(var_3a5a818adf0d)})

			var_f033415d357c_l = append(var_f033415d357c_l, var_3a5a818adf0d_mapped)
		}
		var_f033415d357c_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_f033415d357c_l})
		properties["records"] = var_f033415d357c_mapped
	}

	var_18c1e3cc0e7c := extensionEvent.Ids

	if var_18c1e3cc0e7c != nil {
		var var_18c1e3cc0e7c_mapped *structpb.Value

		var var_18c1e3cc0e7c_l []*structpb.Value
		for _, value := range var_18c1e3cc0e7c {

			var_8458a70aa2d4 := value
			var var_8458a70aa2d4_mapped *structpb.Value

			var var_8458a70aa2d4_err error
			var_8458a70aa2d4_mapped, var_8458a70aa2d4_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_8458a70aa2d4)
			if var_8458a70aa2d4_err != nil {
				panic(var_8458a70aa2d4_err)
			}

			var_18c1e3cc0e7c_l = append(var_18c1e3cc0e7c_l, var_8458a70aa2d4_mapped)
		}
		var_18c1e3cc0e7c_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_18c1e3cc0e7c_l})
		properties["ids"] = var_18c1e3cc0e7c_mapped
	}

	var_9fff7e099410 := extensionEvent.Finalizes

	if var_9fff7e099410 != nil {
		var var_9fff7e099410_mapped *structpb.Value

		var var_9fff7e099410_err error
		var_9fff7e099410_mapped, var_9fff7e099410_err = types.ByResourcePropertyType(model.ResourceProperty_BOOL).Pack(*var_9fff7e099410)
		if var_9fff7e099410_err != nil {
			panic(var_9fff7e099410_err)
		}
		properties["finalizes"] = var_9fff7e099410_mapped
	}

	var_f0664c86d687 := extensionEvent.Sync

	if var_f0664c86d687 != nil {
		var var_f0664c86d687_mapped *structpb.Value

		var var_f0664c86d687_err error
		var_f0664c86d687_mapped, var_f0664c86d687_err = types.ByResourcePropertyType(model.ResourceProperty_BOOL).Pack(*var_f0664c86d687)
		if var_f0664c86d687_err != nil {
			panic(var_f0664c86d687_err)
		}
		properties["sync"] = var_f0664c86d687_mapped
	}

	var_4fbd673e53ca := extensionEvent.Time

	if var_4fbd673e53ca != nil {
		var var_4fbd673e53ca_mapped *structpb.Value

		var var_4fbd673e53ca_err error
		var_4fbd673e53ca_mapped, var_4fbd673e53ca_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_4fbd673e53ca)
		if var_4fbd673e53ca_err != nil {
			panic(var_4fbd673e53ca_err)
		}
		properties["time"] = var_4fbd673e53ca_mapped
	}

	var_c1cb1ca5d4a5 := extensionEvent.Annotations

	if var_c1cb1ca5d4a5 != nil {
		var var_c1cb1ca5d4a5_mapped *structpb.Value

		var var_c1cb1ca5d4a5_st *structpb.Struct = new(structpb.Struct)
		var_c1cb1ca5d4a5_st.Fields = make(map[string]*structpb.Value)
		for key, value := range var_c1cb1ca5d4a5 {

			var_83fab06a3e45 := value
			var var_83fab06a3e45_mapped *structpb.Value

			var var_83fab06a3e45_err error
			var_83fab06a3e45_mapped, var_83fab06a3e45_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_83fab06a3e45)
			if var_83fab06a3e45_err != nil {
				panic(var_83fab06a3e45_err)
			}

			var_c1cb1ca5d4a5_st.Fields[key] = var_83fab06a3e45_mapped
		}
		var_c1cb1ca5d4a5_mapped = structpb.NewStructValue(var_c1cb1ca5d4a5_st)
		properties["annotations"] = var_c1cb1ca5d4a5_mapped
	}
	return properties
}

func (m *ExtensionEventMapper) FromProperties(properties map[string]*structpb.Value) *ExtensionEvent {
	var s = m.New()
	if properties["id"] != nil {

		var_86a792db573e := properties["id"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_UUID).UnPack(var_86a792db573e)

		if err != nil {
			panic(err)
		}

		var_86a792db573e_mapped := new(uuid.UUID)
		*var_86a792db573e_mapped = val.(uuid.UUID)

		s.Id = var_86a792db573e_mapped
	}
	if properties["action"] != nil {

		var_ddbc9ecea3b1 := properties["action"]
		var_ddbc9ecea3b1_mapped := (EventAction)(var_ddbc9ecea3b1.GetStringValue())

		s.Action = var_ddbc9ecea3b1_mapped
	}
	if properties["recordSearchParams"] != nil {

		var_1730c970eb18 := properties["recordSearchParams"]
		var mappedValue = ExtensionRecordSearchParamsMapperInstance.FromProperties(var_1730c970eb18.GetStructValue().Fields)

		var_1730c970eb18_mapped := mappedValue

		s.RecordSearchParams = var_1730c970eb18_mapped
	}
	if properties["actionSummary"] != nil {

		var_9715c5473800 := properties["actionSummary"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_9715c5473800)

		if err != nil {
			panic(err)
		}

		var_9715c5473800_mapped := new(string)
		*var_9715c5473800_mapped = val.(string)

		s.ActionSummary = var_9715c5473800_mapped
	}
	if properties["actionDescription"] != nil {

		var_21eb61ba1c77 := properties["actionDescription"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_21eb61ba1c77)

		if err != nil {
			panic(err)
		}

		var_21eb61ba1c77_mapped := new(string)
		*var_21eb61ba1c77_mapped = val.(string)

		s.ActionDescription = var_21eb61ba1c77_mapped
	}
	if properties["resource"] != nil {

		var_e3cda09aabb4 := properties["resource"]
		var_e3cda09aabb4_mapped := ResourceMapperInstance.FromProperties(var_e3cda09aabb4.GetStructValue().Fields)

		s.Resource = var_e3cda09aabb4_mapped
	}
	if properties["records"] != nil {

		var_c2bf1dee969d := properties["records"]
		var_c2bf1dee969d_mapped := []*Record{}
		for _, v := range var_c2bf1dee969d.GetListValue().Values {

			var_d051ce2b0c14 := v
			var_d051ce2b0c14_mapped := RecordMapperInstance.FromProperties(var_d051ce2b0c14.GetStructValue().Fields)

			var_c2bf1dee969d_mapped = append(var_c2bf1dee969d_mapped, var_d051ce2b0c14_mapped)
		}

		s.Records = var_c2bf1dee969d_mapped
	}
	if properties["ids"] != nil {

		var_08b56c71c418 := properties["ids"]
		var_08b56c71c418_mapped := []string{}
		for _, v := range var_08b56c71c418.GetListValue().Values {

			var_94a144fe6956 := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_94a144fe6956)

			if err != nil {
				panic(err)
			}

			var_94a144fe6956_mapped := val.(string)

			var_08b56c71c418_mapped = append(var_08b56c71c418_mapped, var_94a144fe6956_mapped)
		}

		s.Ids = var_08b56c71c418_mapped
	}
	if properties["finalizes"] != nil {

		var_a718b198ad16 := properties["finalizes"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_BOOL).UnPack(var_a718b198ad16)

		if err != nil {
			panic(err)
		}

		var_a718b198ad16_mapped := new(bool)
		*var_a718b198ad16_mapped = val.(bool)

		s.Finalizes = var_a718b198ad16_mapped
	}
	if properties["sync"] != nil {

		var_807044505ca4 := properties["sync"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_BOOL).UnPack(var_807044505ca4)

		if err != nil {
			panic(err)
		}

		var_807044505ca4_mapped := new(bool)
		*var_807044505ca4_mapped = val.(bool)

		s.Sync = var_807044505ca4_mapped
	}
	if properties["time"] != nil {

		var_b88e49b87264 := properties["time"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_b88e49b87264)

		if err != nil {
			panic(err)
		}

		var_b88e49b87264_mapped := new(time.Time)
		*var_b88e49b87264_mapped = val.(time.Time)

		s.Time = var_b88e49b87264_mapped
	}
	if properties["annotations"] != nil {

		var_5d1a4fee6abd := properties["annotations"]
		var_5d1a4fee6abd_mapped := make(map[string]string)
		for k, v := range var_5d1a4fee6abd.GetStructValue().Fields {

			var_b1c3e0f9e1ad := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_b1c3e0f9e1ad)

			if err != nil {
				panic(err)
			}

			var_b1c3e0f9e1ad_mapped := val.(string)

			var_5d1a4fee6abd_mapped[k] = var_b1c3e0f9e1ad_mapped
		}

		s.Annotations = var_5d1a4fee6abd_mapped
	}
	return s
}
