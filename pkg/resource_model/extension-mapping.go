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

	var_6cb32fab32dc := extension.Id

	if var_6cb32fab32dc != nil {
		var var_6cb32fab32dc_mapped *structpb.Value

		var var_6cb32fab32dc_err error
		var_6cb32fab32dc_mapped, var_6cb32fab32dc_err = types.ByResourcePropertyType(model.ResourceProperty_UUID).Pack(*var_6cb32fab32dc)
		if var_6cb32fab32dc_err != nil {
			panic(var_6cb32fab32dc_err)
		}
		properties["id"] = var_6cb32fab32dc_mapped
	}

	var_c6f8505cc66f := extension.Version

	var var_c6f8505cc66f_mapped *structpb.Value

	var var_c6f8505cc66f_err error
	var_c6f8505cc66f_mapped, var_c6f8505cc66f_err = types.ByResourcePropertyType(model.ResourceProperty_INT32).Pack(var_c6f8505cc66f)
	if var_c6f8505cc66f_err != nil {
		panic(var_c6f8505cc66f_err)
	}
	properties["version"] = var_c6f8505cc66f_mapped

	var_2f0c568d43fe := extension.CreatedBy

	if var_2f0c568d43fe != nil {
		var var_2f0c568d43fe_mapped *structpb.Value

		var var_2f0c568d43fe_err error
		var_2f0c568d43fe_mapped, var_2f0c568d43fe_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_2f0c568d43fe)
		if var_2f0c568d43fe_err != nil {
			panic(var_2f0c568d43fe_err)
		}
		properties["createdBy"] = var_2f0c568d43fe_mapped
	}

	var_adc15d3bd523 := extension.UpdatedBy

	if var_adc15d3bd523 != nil {
		var var_adc15d3bd523_mapped *structpb.Value

		var var_adc15d3bd523_err error
		var_adc15d3bd523_mapped, var_adc15d3bd523_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_adc15d3bd523)
		if var_adc15d3bd523_err != nil {
			panic(var_adc15d3bd523_err)
		}
		properties["updatedBy"] = var_adc15d3bd523_mapped
	}

	var_b4d491193df3 := extension.CreatedOn

	if var_b4d491193df3 != nil {
		var var_b4d491193df3_mapped *structpb.Value

		var var_b4d491193df3_err error
		var_b4d491193df3_mapped, var_b4d491193df3_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_b4d491193df3)
		if var_b4d491193df3_err != nil {
			panic(var_b4d491193df3_err)
		}
		properties["createdOn"] = var_b4d491193df3_mapped
	}

	var_0bd31dee7fbf := extension.UpdatedOn

	if var_0bd31dee7fbf != nil {
		var var_0bd31dee7fbf_mapped *structpb.Value

		var var_0bd31dee7fbf_err error
		var_0bd31dee7fbf_mapped, var_0bd31dee7fbf_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_0bd31dee7fbf)
		if var_0bd31dee7fbf_err != nil {
			panic(var_0bd31dee7fbf_err)
		}
		properties["updatedOn"] = var_0bd31dee7fbf_mapped
	}

	var_f83650227ef9 := extension.Name

	var var_f83650227ef9_mapped *structpb.Value

	var var_f83650227ef9_err error
	var_f83650227ef9_mapped, var_f83650227ef9_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_f83650227ef9)
	if var_f83650227ef9_err != nil {
		panic(var_f83650227ef9_err)
	}
	properties["name"] = var_f83650227ef9_mapped

	var_6338e94481e0 := extension.Description

	if var_6338e94481e0 != nil {
		var var_6338e94481e0_mapped *structpb.Value

		var var_6338e94481e0_err error
		var_6338e94481e0_mapped, var_6338e94481e0_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_6338e94481e0)
		if var_6338e94481e0_err != nil {
			panic(var_6338e94481e0_err)
		}
		properties["description"] = var_6338e94481e0_mapped
	}

	var_936faa767c28 := extension.Selector

	if var_936faa767c28 != nil {
		var var_936faa767c28_mapped *structpb.Value

		var_936faa767c28_mapped = structpb.NewStructValue(&structpb.Struct{Fields: ExtensionEventSelectorMapperInstance.ToProperties(var_936faa767c28)})
		properties["selector"] = var_936faa767c28_mapped
	}

	var_f1b0f4dc432b := extension.Order

	var var_f1b0f4dc432b_mapped *structpb.Value

	var var_f1b0f4dc432b_err error
	var_f1b0f4dc432b_mapped, var_f1b0f4dc432b_err = types.ByResourcePropertyType(model.ResourceProperty_INT32).Pack(var_f1b0f4dc432b)
	if var_f1b0f4dc432b_err != nil {
		panic(var_f1b0f4dc432b_err)
	}
	properties["order"] = var_f1b0f4dc432b_mapped

	var_7edff584a90e := extension.Finalizes

	var var_7edff584a90e_mapped *structpb.Value

	var var_7edff584a90e_err error
	var_7edff584a90e_mapped, var_7edff584a90e_err = types.ByResourcePropertyType(model.ResourceProperty_BOOL).Pack(var_7edff584a90e)
	if var_7edff584a90e_err != nil {
		panic(var_7edff584a90e_err)
	}
	properties["finalizes"] = var_7edff584a90e_mapped

	var_785bbc25b57b := extension.Sync

	var var_785bbc25b57b_mapped *structpb.Value

	var var_785bbc25b57b_err error
	var_785bbc25b57b_mapped, var_785bbc25b57b_err = types.ByResourcePropertyType(model.ResourceProperty_BOOL).Pack(var_785bbc25b57b)
	if var_785bbc25b57b_err != nil {
		panic(var_785bbc25b57b_err)
	}
	properties["sync"] = var_785bbc25b57b_mapped

	var_aa5c468996c5 := extension.Responds

	var var_aa5c468996c5_mapped *structpb.Value

	var var_aa5c468996c5_err error
	var_aa5c468996c5_mapped, var_aa5c468996c5_err = types.ByResourcePropertyType(model.ResourceProperty_BOOL).Pack(var_aa5c468996c5)
	if var_aa5c468996c5_err != nil {
		panic(var_aa5c468996c5_err)
	}
	properties["responds"] = var_aa5c468996c5_mapped

	var_2bf14cb3bd26 := extension.Call

	var var_2bf14cb3bd26_mapped *structpb.Value

	var_2bf14cb3bd26_mapped = structpb.NewStructValue(&structpb.Struct{Fields: ExtensionExternalCallMapperInstance.ToProperties(&var_2bf14cb3bd26)})
	properties["call"] = var_2bf14cb3bd26_mapped

	var_00c2ecdcdfde := extension.Annotations

	if var_00c2ecdcdfde != nil {
		var var_00c2ecdcdfde_mapped *structpb.Value

		var var_00c2ecdcdfde_st *structpb.Struct = new(structpb.Struct)
		var_00c2ecdcdfde_st.Fields = make(map[string]*structpb.Value)
		for key, value := range var_00c2ecdcdfde {

			var_aed70692dca9 := value
			var var_aed70692dca9_mapped *structpb.Value

			var var_aed70692dca9_err error
			var_aed70692dca9_mapped, var_aed70692dca9_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_aed70692dca9)
			if var_aed70692dca9_err != nil {
				panic(var_aed70692dca9_err)
			}

			var_00c2ecdcdfde_st.Fields[key] = var_aed70692dca9_mapped
		}
		var_00c2ecdcdfde_mapped = structpb.NewStructValue(var_00c2ecdcdfde_st)
		properties["annotations"] = var_00c2ecdcdfde_mapped
	}
	return properties
}

func (m *ExtensionMapper) FromProperties(properties map[string]*structpb.Value) *Extension {
	var s = m.New()
	if properties["id"] != nil {

		var_0aedcc8466c3 := properties["id"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_UUID).UnPack(var_0aedcc8466c3)

		if err != nil {
			panic(err)
		}

		var_0aedcc8466c3_mapped := new(uuid.UUID)
		*var_0aedcc8466c3_mapped = val.(uuid.UUID)

		s.Id = var_0aedcc8466c3_mapped
	}
	if properties["version"] != nil {

		var_51bf6de2f207 := properties["version"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(var_51bf6de2f207)

		if err != nil {
			panic(err)
		}

		var_51bf6de2f207_mapped := val.(int32)

		s.Version = var_51bf6de2f207_mapped
	}
	if properties["createdBy"] != nil {

		var_0133f082ced3 := properties["createdBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_0133f082ced3)

		if err != nil {
			panic(err)
		}

		var_0133f082ced3_mapped := new(string)
		*var_0133f082ced3_mapped = val.(string)

		s.CreatedBy = var_0133f082ced3_mapped
	}
	if properties["updatedBy"] != nil {

		var_6d368fc68c08 := properties["updatedBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_6d368fc68c08)

		if err != nil {
			panic(err)
		}

		var_6d368fc68c08_mapped := new(string)
		*var_6d368fc68c08_mapped = val.(string)

		s.UpdatedBy = var_6d368fc68c08_mapped
	}
	if properties["createdOn"] != nil {

		var_8286d6627683 := properties["createdOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_8286d6627683)

		if err != nil {
			panic(err)
		}

		var_8286d6627683_mapped := new(time.Time)
		*var_8286d6627683_mapped = val.(time.Time)

		s.CreatedOn = var_8286d6627683_mapped
	}
	if properties["updatedOn"] != nil {

		var_b37f643dd3c6 := properties["updatedOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_b37f643dd3c6)

		if err != nil {
			panic(err)
		}

		var_b37f643dd3c6_mapped := new(time.Time)
		*var_b37f643dd3c6_mapped = val.(time.Time)

		s.UpdatedOn = var_b37f643dd3c6_mapped
	}
	if properties["name"] != nil {

		var_a59a770efbdb := properties["name"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_a59a770efbdb)

		if err != nil {
			panic(err)
		}

		var_a59a770efbdb_mapped := val.(string)

		s.Name = var_a59a770efbdb_mapped
	}
	if properties["description"] != nil {

		var_66403ada7ce1 := properties["description"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_66403ada7ce1)

		if err != nil {
			panic(err)
		}

		var_66403ada7ce1_mapped := new(string)
		*var_66403ada7ce1_mapped = val.(string)

		s.Description = var_66403ada7ce1_mapped
	}
	if properties["selector"] != nil {

		var_eee8b06c8937 := properties["selector"]
		var mappedValue = ExtensionEventSelectorMapperInstance.FromProperties(var_eee8b06c8937.GetStructValue().Fields)

		var_eee8b06c8937_mapped := mappedValue

		s.Selector = var_eee8b06c8937_mapped
	}
	if properties["order"] != nil {

		var_3917bda22520 := properties["order"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(var_3917bda22520)

		if err != nil {
			panic(err)
		}

		var_3917bda22520_mapped := val.(int32)

		s.Order = var_3917bda22520_mapped
	}
	if properties["finalizes"] != nil {

		var_e94721e2c4b9 := properties["finalizes"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_BOOL).UnPack(var_e94721e2c4b9)

		if err != nil {
			panic(err)
		}

		var_e94721e2c4b9_mapped := val.(bool)

		s.Finalizes = var_e94721e2c4b9_mapped
	}
	if properties["sync"] != nil {

		var_57044fdc7631 := properties["sync"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_BOOL).UnPack(var_57044fdc7631)

		if err != nil {
			panic(err)
		}

		var_57044fdc7631_mapped := val.(bool)

		s.Sync = var_57044fdc7631_mapped
	}
	if properties["responds"] != nil {

		var_bc34202b799c := properties["responds"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_BOOL).UnPack(var_bc34202b799c)

		if err != nil {
			panic(err)
		}

		var_bc34202b799c_mapped := val.(bool)

		s.Responds = var_bc34202b799c_mapped
	}
	if properties["call"] != nil {

		var_bc1255f27ec3 := properties["call"]
		var mappedValue = ExtensionExternalCallMapperInstance.FromProperties(var_bc1255f27ec3.GetStructValue().Fields)

		var_bc1255f27ec3_mapped := *mappedValue

		s.Call = var_bc1255f27ec3_mapped
	}
	if properties["annotations"] != nil {

		var_2576afc5b0c7 := properties["annotations"]
		var_2576afc5b0c7_mapped := make(map[string]string)
		for k, v := range var_2576afc5b0c7.GetStructValue().Fields {

			var_da357c857460 := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_da357c857460)

			if err != nil {
				panic(err)
			}

			var_da357c857460_mapped := val.(string)

			var_2576afc5b0c7_mapped[k] = var_da357c857460_mapped
		}

		s.Annotations = var_2576afc5b0c7_mapped
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

	var_0ffd0dc0ae7b := extensionFunctionCall.Host

	var var_0ffd0dc0ae7b_mapped *structpb.Value

	var var_0ffd0dc0ae7b_err error
	var_0ffd0dc0ae7b_mapped, var_0ffd0dc0ae7b_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_0ffd0dc0ae7b)
	if var_0ffd0dc0ae7b_err != nil {
		panic(var_0ffd0dc0ae7b_err)
	}
	properties["host"] = var_0ffd0dc0ae7b_mapped

	var_a99620146cd9 := extensionFunctionCall.FunctionName

	var var_a99620146cd9_mapped *structpb.Value

	var var_a99620146cd9_err error
	var_a99620146cd9_mapped, var_a99620146cd9_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_a99620146cd9)
	if var_a99620146cd9_err != nil {
		panic(var_a99620146cd9_err)
	}
	properties["functionName"] = var_a99620146cd9_mapped
	return properties
}

func (m *ExtensionFunctionCallMapper) FromProperties(properties map[string]*structpb.Value) *ExtensionFunctionCall {
	var s = m.New()
	if properties["host"] != nil {

		var_b260ffa947df := properties["host"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_b260ffa947df)

		if err != nil {
			panic(err)
		}

		var_b260ffa947df_mapped := val.(string)

		s.Host = var_b260ffa947df_mapped
	}
	if properties["functionName"] != nil {

		var_65d53aa6086b := properties["functionName"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_65d53aa6086b)

		if err != nil {
			panic(err)
		}

		var_65d53aa6086b_mapped := val.(string)

		s.FunctionName = var_65d53aa6086b_mapped
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

	var_0a99927a96ef := extensionHttpCall.Uri

	var var_0a99927a96ef_mapped *structpb.Value

	var var_0a99927a96ef_err error
	var_0a99927a96ef_mapped, var_0a99927a96ef_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_0a99927a96ef)
	if var_0a99927a96ef_err != nil {
		panic(var_0a99927a96ef_err)
	}
	properties["uri"] = var_0a99927a96ef_mapped

	var_ce95bd41c5c5 := extensionHttpCall.Method

	var var_ce95bd41c5c5_mapped *structpb.Value

	var var_ce95bd41c5c5_err error
	var_ce95bd41c5c5_mapped, var_ce95bd41c5c5_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_ce95bd41c5c5)
	if var_ce95bd41c5c5_err != nil {
		panic(var_ce95bd41c5c5_err)
	}
	properties["method"] = var_ce95bd41c5c5_mapped
	return properties
}

func (m *ExtensionHttpCallMapper) FromProperties(properties map[string]*structpb.Value) *ExtensionHttpCall {
	var s = m.New()
	if properties["uri"] != nil {

		var_bdd3cb6985f8 := properties["uri"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_bdd3cb6985f8)

		if err != nil {
			panic(err)
		}

		var_bdd3cb6985f8_mapped := val.(string)

		s.Uri = var_bdd3cb6985f8_mapped
	}
	if properties["method"] != nil {

		var_24b0e00e67de := properties["method"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_24b0e00e67de)

		if err != nil {
			panic(err)
		}

		var_24b0e00e67de_mapped := val.(string)

		s.Method = var_24b0e00e67de_mapped
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

	var_6f2f3ba6450f := extensionExternalCall.FunctionCall

	if var_6f2f3ba6450f != nil {
		var var_6f2f3ba6450f_mapped *structpb.Value

		var_6f2f3ba6450f_mapped = structpb.NewStructValue(&structpb.Struct{Fields: ExtensionFunctionCallMapperInstance.ToProperties(var_6f2f3ba6450f)})
		properties["functionCall"] = var_6f2f3ba6450f_mapped
	}

	var_ec467565b56f := extensionExternalCall.HttpCall

	if var_ec467565b56f != nil {
		var var_ec467565b56f_mapped *structpb.Value

		var_ec467565b56f_mapped = structpb.NewStructValue(&structpb.Struct{Fields: ExtensionHttpCallMapperInstance.ToProperties(var_ec467565b56f)})
		properties["httpCall"] = var_ec467565b56f_mapped
	}
	return properties
}

func (m *ExtensionExternalCallMapper) FromProperties(properties map[string]*structpb.Value) *ExtensionExternalCall {
	var s = m.New()
	if properties["functionCall"] != nil {

		var_c1e018f0d5cf := properties["functionCall"]
		var mappedValue = ExtensionFunctionCallMapperInstance.FromProperties(var_c1e018f0d5cf.GetStructValue().Fields)

		var_c1e018f0d5cf_mapped := mappedValue

		s.FunctionCall = var_c1e018f0d5cf_mapped
	}
	if properties["httpCall"] != nil {

		var_925251ba2b34 := properties["httpCall"]
		var mappedValue = ExtensionHttpCallMapperInstance.FromProperties(var_925251ba2b34.GetStructValue().Fields)

		var_925251ba2b34_mapped := mappedValue

		s.HttpCall = var_925251ba2b34_mapped
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

	var_a29752f2c254 := extensionEventSelector.Actions

	if var_a29752f2c254 != nil {
		var var_a29752f2c254_mapped *structpb.Value

		var var_a29752f2c254_l []*structpb.Value
		for _, value := range var_a29752f2c254 {

			var_d83841ed6a92 := value
			var var_d83841ed6a92_mapped *structpb.Value

			var var_d83841ed6a92_err error
			var_d83841ed6a92_mapped, var_d83841ed6a92_err = types.ByResourcePropertyType(model.ResourceProperty_ENUM).Pack(string(var_d83841ed6a92))
			if var_d83841ed6a92_err != nil {
				panic(var_d83841ed6a92_err)
			}

			var_a29752f2c254_l = append(var_a29752f2c254_l, var_d83841ed6a92_mapped)
		}
		var_a29752f2c254_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_a29752f2c254_l})
		properties["actions"] = var_a29752f2c254_mapped
	}

	var_b9d2ddf4e5b7 := extensionEventSelector.RecordSelector

	if var_b9d2ddf4e5b7 != nil {
		var var_b9d2ddf4e5b7_mapped *structpb.Value

		var_b9d2ddf4e5b7_mapped = structpb.NewStructValue(&structpb.Struct{Fields: ExtensionBooleanExpressionMapperInstance.ToProperties(var_b9d2ddf4e5b7)})
		properties["recordSelector"] = var_b9d2ddf4e5b7_mapped
	}

	var_710db6d3bb92 := extensionEventSelector.Namespaces

	if var_710db6d3bb92 != nil {
		var var_710db6d3bb92_mapped *structpb.Value

		var var_710db6d3bb92_l []*structpb.Value
		for _, value := range var_710db6d3bb92 {

			var_7f00d60cac4b := value
			var var_7f00d60cac4b_mapped *structpb.Value

			var var_7f00d60cac4b_err error
			var_7f00d60cac4b_mapped, var_7f00d60cac4b_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_7f00d60cac4b)
			if var_7f00d60cac4b_err != nil {
				panic(var_7f00d60cac4b_err)
			}

			var_710db6d3bb92_l = append(var_710db6d3bb92_l, var_7f00d60cac4b_mapped)
		}
		var_710db6d3bb92_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_710db6d3bb92_l})
		properties["namespaces"] = var_710db6d3bb92_mapped
	}

	var_0ba32822c870 := extensionEventSelector.Resources

	if var_0ba32822c870 != nil {
		var var_0ba32822c870_mapped *structpb.Value

		var var_0ba32822c870_l []*structpb.Value
		for _, value := range var_0ba32822c870 {

			var_28d6aceeb3e5 := value
			var var_28d6aceeb3e5_mapped *structpb.Value

			var var_28d6aceeb3e5_err error
			var_28d6aceeb3e5_mapped, var_28d6aceeb3e5_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_28d6aceeb3e5)
			if var_28d6aceeb3e5_err != nil {
				panic(var_28d6aceeb3e5_err)
			}

			var_0ba32822c870_l = append(var_0ba32822c870_l, var_28d6aceeb3e5_mapped)
		}
		var_0ba32822c870_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_0ba32822c870_l})
		properties["resources"] = var_0ba32822c870_mapped
	}

	var_68c5faae9b58 := extensionEventSelector.Ids

	if var_68c5faae9b58 != nil {
		var var_68c5faae9b58_mapped *structpb.Value

		var var_68c5faae9b58_l []*structpb.Value
		for _, value := range var_68c5faae9b58 {

			var_468f62bbbc06 := value
			var var_468f62bbbc06_mapped *structpb.Value

			var var_468f62bbbc06_err error
			var_468f62bbbc06_mapped, var_468f62bbbc06_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_468f62bbbc06)
			if var_468f62bbbc06_err != nil {
				panic(var_468f62bbbc06_err)
			}

			var_68c5faae9b58_l = append(var_68c5faae9b58_l, var_468f62bbbc06_mapped)
		}
		var_68c5faae9b58_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_68c5faae9b58_l})
		properties["ids"] = var_68c5faae9b58_mapped
	}

	var_4eadf0a67550 := extensionEventSelector.Annotations

	if var_4eadf0a67550 != nil {
		var var_4eadf0a67550_mapped *structpb.Value

		var var_4eadf0a67550_st *structpb.Struct = new(structpb.Struct)
		var_4eadf0a67550_st.Fields = make(map[string]*structpb.Value)
		for key, value := range var_4eadf0a67550 {

			var_de87d7297ffb := value
			var var_de87d7297ffb_mapped *structpb.Value

			var var_de87d7297ffb_err error
			var_de87d7297ffb_mapped, var_de87d7297ffb_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_de87d7297ffb)
			if var_de87d7297ffb_err != nil {
				panic(var_de87d7297ffb_err)
			}

			var_4eadf0a67550_st.Fields[key] = var_de87d7297ffb_mapped
		}
		var_4eadf0a67550_mapped = structpb.NewStructValue(var_4eadf0a67550_st)
		properties["annotations"] = var_4eadf0a67550_mapped
	}
	return properties
}

func (m *ExtensionEventSelectorMapper) FromProperties(properties map[string]*structpb.Value) *ExtensionEventSelector {
	var s = m.New()
	if properties["actions"] != nil {

		var_4d74a6b4084b := properties["actions"]
		var_4d74a6b4084b_mapped := []EventAction{}
		for _, v := range var_4d74a6b4084b.GetListValue().Values {

			var_d18901bc680a := v
			var_d18901bc680a_mapped := (EventAction)(var_d18901bc680a.GetStringValue())

			var_4d74a6b4084b_mapped = append(var_4d74a6b4084b_mapped, var_d18901bc680a_mapped)
		}

		s.Actions = var_4d74a6b4084b_mapped
	}
	if properties["recordSelector"] != nil {

		var_ae21181d2ed8 := properties["recordSelector"]
		var mappedValue = ExtensionBooleanExpressionMapperInstance.FromProperties(var_ae21181d2ed8.GetStructValue().Fields)

		var_ae21181d2ed8_mapped := mappedValue

		s.RecordSelector = var_ae21181d2ed8_mapped
	}
	if properties["namespaces"] != nil {

		var_d58345607035 := properties["namespaces"]
		var_d58345607035_mapped := []string{}
		for _, v := range var_d58345607035.GetListValue().Values {

			var_fdf4a73ef8a2 := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_fdf4a73ef8a2)

			if err != nil {
				panic(err)
			}

			var_fdf4a73ef8a2_mapped := val.(string)

			var_d58345607035_mapped = append(var_d58345607035_mapped, var_fdf4a73ef8a2_mapped)
		}

		s.Namespaces = var_d58345607035_mapped
	}
	if properties["resources"] != nil {

		var_c1d81536cd22 := properties["resources"]
		var_c1d81536cd22_mapped := []string{}
		for _, v := range var_c1d81536cd22.GetListValue().Values {

			var_4d197e4b8a91 := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_4d197e4b8a91)

			if err != nil {
				panic(err)
			}

			var_4d197e4b8a91_mapped := val.(string)

			var_c1d81536cd22_mapped = append(var_c1d81536cd22_mapped, var_4d197e4b8a91_mapped)
		}

		s.Resources = var_c1d81536cd22_mapped
	}
	if properties["ids"] != nil {

		var_bd583635dce6 := properties["ids"]
		var_bd583635dce6_mapped := []string{}
		for _, v := range var_bd583635dce6.GetListValue().Values {

			var_11d6ebe4f345 := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_11d6ebe4f345)

			if err != nil {
				panic(err)
			}

			var_11d6ebe4f345_mapped := val.(string)

			var_bd583635dce6_mapped = append(var_bd583635dce6_mapped, var_11d6ebe4f345_mapped)
		}

		s.Ids = var_bd583635dce6_mapped
	}
	if properties["annotations"] != nil {

		var_948bc809472a := properties["annotations"]
		var_948bc809472a_mapped := make(map[string]string)
		for k, v := range var_948bc809472a.GetStructValue().Fields {

			var_1a2805fcd6de := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_1a2805fcd6de)

			if err != nil {
				panic(err)
			}

			var_1a2805fcd6de_mapped := val.(string)

			var_948bc809472a_mapped[k] = var_1a2805fcd6de_mapped
		}

		s.Annotations = var_948bc809472a_mapped
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

	var_c19f8a874996 := extensionRecordSearchParams.Query

	if var_c19f8a874996 != nil {
		var var_c19f8a874996_mapped *structpb.Value

		var_c19f8a874996_mapped = structpb.NewStructValue(&structpb.Struct{Fields: ExtensionBooleanExpressionMapperInstance.ToProperties(var_c19f8a874996)})
		properties["query"] = var_c19f8a874996_mapped
	}

	var_8d8a85db4a60 := extensionRecordSearchParams.Limit

	if var_8d8a85db4a60 != nil {
		var var_8d8a85db4a60_mapped *structpb.Value

		var var_8d8a85db4a60_err error
		var_8d8a85db4a60_mapped, var_8d8a85db4a60_err = types.ByResourcePropertyType(model.ResourceProperty_INT32).Pack(*var_8d8a85db4a60)
		if var_8d8a85db4a60_err != nil {
			panic(var_8d8a85db4a60_err)
		}
		properties["limit"] = var_8d8a85db4a60_mapped
	}

	var_ffe18f1fbbb7 := extensionRecordSearchParams.Offset

	if var_ffe18f1fbbb7 != nil {
		var var_ffe18f1fbbb7_mapped *structpb.Value

		var var_ffe18f1fbbb7_err error
		var_ffe18f1fbbb7_mapped, var_ffe18f1fbbb7_err = types.ByResourcePropertyType(model.ResourceProperty_INT32).Pack(*var_ffe18f1fbbb7)
		if var_ffe18f1fbbb7_err != nil {
			panic(var_ffe18f1fbbb7_err)
		}
		properties["offset"] = var_ffe18f1fbbb7_mapped
	}

	var_f044778b29f2 := extensionRecordSearchParams.ResolveReferences

	if var_f044778b29f2 != nil {
		var var_f044778b29f2_mapped *structpb.Value

		var var_f044778b29f2_l []*structpb.Value
		for _, value := range var_f044778b29f2 {

			var_de8585971864 := value
			var var_de8585971864_mapped *structpb.Value

			var var_de8585971864_err error
			var_de8585971864_mapped, var_de8585971864_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_de8585971864)
			if var_de8585971864_err != nil {
				panic(var_de8585971864_err)
			}

			var_f044778b29f2_l = append(var_f044778b29f2_l, var_de8585971864_mapped)
		}
		var_f044778b29f2_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_f044778b29f2_l})
		properties["resolveReferences"] = var_f044778b29f2_mapped
	}
	return properties
}

func (m *ExtensionRecordSearchParamsMapper) FromProperties(properties map[string]*structpb.Value) *ExtensionRecordSearchParams {
	var s = m.New()
	if properties["query"] != nil {

		var_79df2d834b47 := properties["query"]
		var mappedValue = ExtensionBooleanExpressionMapperInstance.FromProperties(var_79df2d834b47.GetStructValue().Fields)

		var_79df2d834b47_mapped := mappedValue

		s.Query = var_79df2d834b47_mapped
	}
	if properties["limit"] != nil {

		var_0d281f07ee83 := properties["limit"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(var_0d281f07ee83)

		if err != nil {
			panic(err)
		}

		var_0d281f07ee83_mapped := new(int32)
		*var_0d281f07ee83_mapped = val.(int32)

		s.Limit = var_0d281f07ee83_mapped
	}
	if properties["offset"] != nil {

		var_d24f06989423 := properties["offset"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(var_d24f06989423)

		if err != nil {
			panic(err)
		}

		var_d24f06989423_mapped := new(int32)
		*var_d24f06989423_mapped = val.(int32)

		s.Offset = var_d24f06989423_mapped
	}
	if properties["resolveReferences"] != nil {

		var_aa04edf55921 := properties["resolveReferences"]
		var_aa04edf55921_mapped := []string{}
		for _, v := range var_aa04edf55921.GetListValue().Values {

			var_b0559b67213a := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_b0559b67213a)

			if err != nil {
				panic(err)
			}

			var_b0559b67213a_mapped := val.(string)

			var_aa04edf55921_mapped = append(var_aa04edf55921_mapped, var_b0559b67213a_mapped)
		}

		s.ResolveReferences = var_aa04edf55921_mapped
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

	var_373caa2fe7c2 := extensionEvent.Id

	if var_373caa2fe7c2 != nil {
		var var_373caa2fe7c2_mapped *structpb.Value

		var var_373caa2fe7c2_err error
		var_373caa2fe7c2_mapped, var_373caa2fe7c2_err = types.ByResourcePropertyType(model.ResourceProperty_UUID).Pack(*var_373caa2fe7c2)
		if var_373caa2fe7c2_err != nil {
			panic(var_373caa2fe7c2_err)
		}
		properties["id"] = var_373caa2fe7c2_mapped
	}

	var_ca0268211a49 := extensionEvent.Action

	var var_ca0268211a49_mapped *structpb.Value

	var var_ca0268211a49_err error
	var_ca0268211a49_mapped, var_ca0268211a49_err = types.ByResourcePropertyType(model.ResourceProperty_ENUM).Pack(string(var_ca0268211a49))
	if var_ca0268211a49_err != nil {
		panic(var_ca0268211a49_err)
	}
	properties["action"] = var_ca0268211a49_mapped

	var_ecb68aa7f12b := extensionEvent.RecordSearchParams

	if var_ecb68aa7f12b != nil {
		var var_ecb68aa7f12b_mapped *structpb.Value

		var_ecb68aa7f12b_mapped = structpb.NewStructValue(&structpb.Struct{Fields: ExtensionRecordSearchParamsMapperInstance.ToProperties(var_ecb68aa7f12b)})
		properties["recordSearchParams"] = var_ecb68aa7f12b_mapped
	}

	var_92d2ab235849 := extensionEvent.ActionSummary

	if var_92d2ab235849 != nil {
		var var_92d2ab235849_mapped *structpb.Value

		var var_92d2ab235849_err error
		var_92d2ab235849_mapped, var_92d2ab235849_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_92d2ab235849)
		if var_92d2ab235849_err != nil {
			panic(var_92d2ab235849_err)
		}
		properties["actionSummary"] = var_92d2ab235849_mapped
	}

	var_7ce2edc2db4e := extensionEvent.ActionDescription

	if var_7ce2edc2db4e != nil {
		var var_7ce2edc2db4e_mapped *structpb.Value

		var var_7ce2edc2db4e_err error
		var_7ce2edc2db4e_mapped, var_7ce2edc2db4e_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_7ce2edc2db4e)
		if var_7ce2edc2db4e_err != nil {
			panic(var_7ce2edc2db4e_err)
		}
		properties["actionDescription"] = var_7ce2edc2db4e_mapped
	}

	var_780897ea710e := extensionEvent.Resource

	if var_780897ea710e != nil {
		var var_780897ea710e_mapped *structpb.Value

		var_780897ea710e_mapped = structpb.NewStructValue(&structpb.Struct{Fields: ResourceMapperInstance.ToProperties(var_780897ea710e)})
		properties["resource"] = var_780897ea710e_mapped
	}

	var_72a17b398f3b := extensionEvent.Records

	if var_72a17b398f3b != nil {
		var var_72a17b398f3b_mapped *structpb.Value

		var var_72a17b398f3b_l []*structpb.Value
		for _, value := range var_72a17b398f3b {

			var_d7893cadd7b4 := value
			var var_d7893cadd7b4_mapped *structpb.Value

			var_d7893cadd7b4_mapped = structpb.NewStructValue(&structpb.Struct{Fields: RecordMapperInstance.ToProperties(var_d7893cadd7b4)})

			var_72a17b398f3b_l = append(var_72a17b398f3b_l, var_d7893cadd7b4_mapped)
		}
		var_72a17b398f3b_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_72a17b398f3b_l})
		properties["records"] = var_72a17b398f3b_mapped
	}

	var_783192e98596 := extensionEvent.Ids

	if var_783192e98596 != nil {
		var var_783192e98596_mapped *structpb.Value

		var var_783192e98596_l []*structpb.Value
		for _, value := range var_783192e98596 {

			var_f3527a14e92e := value
			var var_f3527a14e92e_mapped *structpb.Value

			var var_f3527a14e92e_err error
			var_f3527a14e92e_mapped, var_f3527a14e92e_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_f3527a14e92e)
			if var_f3527a14e92e_err != nil {
				panic(var_f3527a14e92e_err)
			}

			var_783192e98596_l = append(var_783192e98596_l, var_f3527a14e92e_mapped)
		}
		var_783192e98596_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_783192e98596_l})
		properties["ids"] = var_783192e98596_mapped
	}

	var_86929c47a122 := extensionEvent.Finalizes

	if var_86929c47a122 != nil {
		var var_86929c47a122_mapped *structpb.Value

		var var_86929c47a122_err error
		var_86929c47a122_mapped, var_86929c47a122_err = types.ByResourcePropertyType(model.ResourceProperty_BOOL).Pack(*var_86929c47a122)
		if var_86929c47a122_err != nil {
			panic(var_86929c47a122_err)
		}
		properties["finalizes"] = var_86929c47a122_mapped
	}

	var_53f89fe1993a := extensionEvent.Sync

	if var_53f89fe1993a != nil {
		var var_53f89fe1993a_mapped *structpb.Value

		var var_53f89fe1993a_err error
		var_53f89fe1993a_mapped, var_53f89fe1993a_err = types.ByResourcePropertyType(model.ResourceProperty_BOOL).Pack(*var_53f89fe1993a)
		if var_53f89fe1993a_err != nil {
			panic(var_53f89fe1993a_err)
		}
		properties["sync"] = var_53f89fe1993a_mapped
	}

	var_c907e8a82c39 := extensionEvent.Time

	if var_c907e8a82c39 != nil {
		var var_c907e8a82c39_mapped *structpb.Value

		var var_c907e8a82c39_err error
		var_c907e8a82c39_mapped, var_c907e8a82c39_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_c907e8a82c39)
		if var_c907e8a82c39_err != nil {
			panic(var_c907e8a82c39_err)
		}
		properties["time"] = var_c907e8a82c39_mapped
	}

	var_2d9f285e7836 := extensionEvent.Annotations

	if var_2d9f285e7836 != nil {
		var var_2d9f285e7836_mapped *structpb.Value

		var var_2d9f285e7836_st *structpb.Struct = new(structpb.Struct)
		var_2d9f285e7836_st.Fields = make(map[string]*structpb.Value)
		for key, value := range var_2d9f285e7836 {

			var_db88d8c84eb4 := value
			var var_db88d8c84eb4_mapped *structpb.Value

			var var_db88d8c84eb4_err error
			var_db88d8c84eb4_mapped, var_db88d8c84eb4_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_db88d8c84eb4)
			if var_db88d8c84eb4_err != nil {
				panic(var_db88d8c84eb4_err)
			}

			var_2d9f285e7836_st.Fields[key] = var_db88d8c84eb4_mapped
		}
		var_2d9f285e7836_mapped = structpb.NewStructValue(var_2d9f285e7836_st)
		properties["annotations"] = var_2d9f285e7836_mapped
	}
	return properties
}

func (m *ExtensionEventMapper) FromProperties(properties map[string]*structpb.Value) *ExtensionEvent {
	var s = m.New()
	if properties["id"] != nil {

		var_b3ca817c3b80 := properties["id"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_UUID).UnPack(var_b3ca817c3b80)

		if err != nil {
			panic(err)
		}

		var_b3ca817c3b80_mapped := new(uuid.UUID)
		*var_b3ca817c3b80_mapped = val.(uuid.UUID)

		s.Id = var_b3ca817c3b80_mapped
	}
	if properties["action"] != nil {

		var_3b291df29851 := properties["action"]
		var_3b291df29851_mapped := (EventAction)(var_3b291df29851.GetStringValue())

		s.Action = var_3b291df29851_mapped
	}
	if properties["recordSearchParams"] != nil {

		var_6e27d5e00b2e := properties["recordSearchParams"]
		var mappedValue = ExtensionRecordSearchParamsMapperInstance.FromProperties(var_6e27d5e00b2e.GetStructValue().Fields)

		var_6e27d5e00b2e_mapped := mappedValue

		s.RecordSearchParams = var_6e27d5e00b2e_mapped
	}
	if properties["actionSummary"] != nil {

		var_401162d66513 := properties["actionSummary"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_401162d66513)

		if err != nil {
			panic(err)
		}

		var_401162d66513_mapped := new(string)
		*var_401162d66513_mapped = val.(string)

		s.ActionSummary = var_401162d66513_mapped
	}
	if properties["actionDescription"] != nil {

		var_68ca4cac15d4 := properties["actionDescription"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_68ca4cac15d4)

		if err != nil {
			panic(err)
		}

		var_68ca4cac15d4_mapped := new(string)
		*var_68ca4cac15d4_mapped = val.(string)

		s.ActionDescription = var_68ca4cac15d4_mapped
	}
	if properties["resource"] != nil {

		var_fc648c9588eb := properties["resource"]
		var_fc648c9588eb_mapped := ResourceMapperInstance.FromProperties(var_fc648c9588eb.GetStructValue().Fields)

		s.Resource = var_fc648c9588eb_mapped
	}
	if properties["records"] != nil {

		var_2e4491b7408a := properties["records"]
		var_2e4491b7408a_mapped := []*Record{}
		for _, v := range var_2e4491b7408a.GetListValue().Values {

			var_2a1885817c27 := v
			var_2a1885817c27_mapped := RecordMapperInstance.FromProperties(var_2a1885817c27.GetStructValue().Fields)

			var_2e4491b7408a_mapped = append(var_2e4491b7408a_mapped, var_2a1885817c27_mapped)
		}

		s.Records = var_2e4491b7408a_mapped
	}
	if properties["ids"] != nil {

		var_4f140e61c316 := properties["ids"]
		var_4f140e61c316_mapped := []string{}
		for _, v := range var_4f140e61c316.GetListValue().Values {

			var_1df0a56a8849 := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_1df0a56a8849)

			if err != nil {
				panic(err)
			}

			var_1df0a56a8849_mapped := val.(string)

			var_4f140e61c316_mapped = append(var_4f140e61c316_mapped, var_1df0a56a8849_mapped)
		}

		s.Ids = var_4f140e61c316_mapped
	}
	if properties["finalizes"] != nil {

		var_56b29e04a8d2 := properties["finalizes"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_BOOL).UnPack(var_56b29e04a8d2)

		if err != nil {
			panic(err)
		}

		var_56b29e04a8d2_mapped := new(bool)
		*var_56b29e04a8d2_mapped = val.(bool)

		s.Finalizes = var_56b29e04a8d2_mapped
	}
	if properties["sync"] != nil {

		var_6a2e457af2e0 := properties["sync"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_BOOL).UnPack(var_6a2e457af2e0)

		if err != nil {
			panic(err)
		}

		var_6a2e457af2e0_mapped := new(bool)
		*var_6a2e457af2e0_mapped = val.(bool)

		s.Sync = var_6a2e457af2e0_mapped
	}
	if properties["time"] != nil {

		var_2ce7cecda0bb := properties["time"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_2ce7cecda0bb)

		if err != nil {
			panic(err)
		}

		var_2ce7cecda0bb_mapped := new(time.Time)
		*var_2ce7cecda0bb_mapped = val.(time.Time)

		s.Time = var_2ce7cecda0bb_mapped
	}
	if properties["annotations"] != nil {

		var_caee2e363c18 := properties["annotations"]
		var_caee2e363c18_mapped := make(map[string]string)
		for k, v := range var_caee2e363c18.GetStructValue().Fields {

			var_348b5269a9cb := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_348b5269a9cb)

			if err != nil {
				panic(err)
			}

			var_348b5269a9cb_mapped := val.(string)

			var_caee2e363c18_mapped[k] = var_348b5269a9cb_mapped
		}

		s.Annotations = var_caee2e363c18_mapped
	}
	return s
}
