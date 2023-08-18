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

	var_ef7541f7b942 := extension.Id

	if var_ef7541f7b942 != nil {
		var var_ef7541f7b942_mapped *structpb.Value

		var var_ef7541f7b942_err error
		var_ef7541f7b942_mapped, var_ef7541f7b942_err = types.ByResourcePropertyType(model.ResourceProperty_UUID).Pack(*var_ef7541f7b942)
		if var_ef7541f7b942_err != nil {
			panic(var_ef7541f7b942_err)
		}
		properties["id"] = var_ef7541f7b942_mapped
	}

	var_e23e75d22f11 := extension.Version

	var var_e23e75d22f11_mapped *structpb.Value

	var var_e23e75d22f11_err error
	var_e23e75d22f11_mapped, var_e23e75d22f11_err = types.ByResourcePropertyType(model.ResourceProperty_INT32).Pack(var_e23e75d22f11)
	if var_e23e75d22f11_err != nil {
		panic(var_e23e75d22f11_err)
	}
	properties["version"] = var_e23e75d22f11_mapped

	var_51cc859068fa := extension.CreatedBy

	if var_51cc859068fa != nil {
		var var_51cc859068fa_mapped *structpb.Value

		var var_51cc859068fa_err error
		var_51cc859068fa_mapped, var_51cc859068fa_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_51cc859068fa)
		if var_51cc859068fa_err != nil {
			panic(var_51cc859068fa_err)
		}
		properties["createdBy"] = var_51cc859068fa_mapped
	}

	var_12c1a789c3e3 := extension.UpdatedBy

	if var_12c1a789c3e3 != nil {
		var var_12c1a789c3e3_mapped *structpb.Value

		var var_12c1a789c3e3_err error
		var_12c1a789c3e3_mapped, var_12c1a789c3e3_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_12c1a789c3e3)
		if var_12c1a789c3e3_err != nil {
			panic(var_12c1a789c3e3_err)
		}
		properties["updatedBy"] = var_12c1a789c3e3_mapped
	}

	var_e5f08bcf20b3 := extension.CreatedOn

	if var_e5f08bcf20b3 != nil {
		var var_e5f08bcf20b3_mapped *structpb.Value

		var var_e5f08bcf20b3_err error
		var_e5f08bcf20b3_mapped, var_e5f08bcf20b3_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_e5f08bcf20b3)
		if var_e5f08bcf20b3_err != nil {
			panic(var_e5f08bcf20b3_err)
		}
		properties["createdOn"] = var_e5f08bcf20b3_mapped
	}

	var_7cf547c89cd5 := extension.UpdatedOn

	if var_7cf547c89cd5 != nil {
		var var_7cf547c89cd5_mapped *structpb.Value

		var var_7cf547c89cd5_err error
		var_7cf547c89cd5_mapped, var_7cf547c89cd5_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_7cf547c89cd5)
		if var_7cf547c89cd5_err != nil {
			panic(var_7cf547c89cd5_err)
		}
		properties["updatedOn"] = var_7cf547c89cd5_mapped
	}

	var_9579a3506ce0 := extension.Name

	var var_9579a3506ce0_mapped *structpb.Value

	var var_9579a3506ce0_err error
	var_9579a3506ce0_mapped, var_9579a3506ce0_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_9579a3506ce0)
	if var_9579a3506ce0_err != nil {
		panic(var_9579a3506ce0_err)
	}
	properties["name"] = var_9579a3506ce0_mapped

	var_c85aacd7b10e := extension.Description

	if var_c85aacd7b10e != nil {
		var var_c85aacd7b10e_mapped *structpb.Value

		var var_c85aacd7b10e_err error
		var_c85aacd7b10e_mapped, var_c85aacd7b10e_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_c85aacd7b10e)
		if var_c85aacd7b10e_err != nil {
			panic(var_c85aacd7b10e_err)
		}
		properties["description"] = var_c85aacd7b10e_mapped
	}

	var_59b88e2a355c := extension.Selector

	if var_59b88e2a355c != nil {
		var var_59b88e2a355c_mapped *structpb.Value

		var_59b88e2a355c_mapped = structpb.NewStructValue(&structpb.Struct{Fields: ExtensionEventSelectorMapperInstance.ToProperties(var_59b88e2a355c)})
		properties["selector"] = var_59b88e2a355c_mapped
	}

	var_6722e3875beb := extension.Order

	var var_6722e3875beb_mapped *structpb.Value

	var var_6722e3875beb_err error
	var_6722e3875beb_mapped, var_6722e3875beb_err = types.ByResourcePropertyType(model.ResourceProperty_INT32).Pack(var_6722e3875beb)
	if var_6722e3875beb_err != nil {
		panic(var_6722e3875beb_err)
	}
	properties["order"] = var_6722e3875beb_mapped

	var_7835c9215eb9 := extension.Finalizes

	var var_7835c9215eb9_mapped *structpb.Value

	var var_7835c9215eb9_err error
	var_7835c9215eb9_mapped, var_7835c9215eb9_err = types.ByResourcePropertyType(model.ResourceProperty_BOOL).Pack(var_7835c9215eb9)
	if var_7835c9215eb9_err != nil {
		panic(var_7835c9215eb9_err)
	}
	properties["finalizes"] = var_7835c9215eb9_mapped

	var_d3fdba511488 := extension.Sync

	var var_d3fdba511488_mapped *structpb.Value

	var var_d3fdba511488_err error
	var_d3fdba511488_mapped, var_d3fdba511488_err = types.ByResourcePropertyType(model.ResourceProperty_BOOL).Pack(var_d3fdba511488)
	if var_d3fdba511488_err != nil {
		panic(var_d3fdba511488_err)
	}
	properties["sync"] = var_d3fdba511488_mapped

	var_31352606c2d5 := extension.Responds

	var var_31352606c2d5_mapped *structpb.Value

	var var_31352606c2d5_err error
	var_31352606c2d5_mapped, var_31352606c2d5_err = types.ByResourcePropertyType(model.ResourceProperty_BOOL).Pack(var_31352606c2d5)
	if var_31352606c2d5_err != nil {
		panic(var_31352606c2d5_err)
	}
	properties["responds"] = var_31352606c2d5_mapped

	var_0995a19d3576 := extension.Call

	var var_0995a19d3576_mapped *structpb.Value

	var_0995a19d3576_mapped = structpb.NewStructValue(&structpb.Struct{Fields: ExtensionExternalCallMapperInstance.ToProperties(&var_0995a19d3576)})
	properties["call"] = var_0995a19d3576_mapped

	var_2604ff076066 := extension.Annotations

	if var_2604ff076066 != nil {
		var var_2604ff076066_mapped *structpb.Value

		var var_2604ff076066_st *structpb.Struct = new(structpb.Struct)
		var_2604ff076066_st.Fields = make(map[string]*structpb.Value)
		for key, value := range var_2604ff076066 {

			var_bdc72263dbc8 := value
			var var_bdc72263dbc8_mapped *structpb.Value

			var var_bdc72263dbc8_err error
			var_bdc72263dbc8_mapped, var_bdc72263dbc8_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_bdc72263dbc8)
			if var_bdc72263dbc8_err != nil {
				panic(var_bdc72263dbc8_err)
			}

			var_2604ff076066_st.Fields[key] = var_bdc72263dbc8_mapped
		}
		var_2604ff076066_mapped = structpb.NewStructValue(var_2604ff076066_st)
		properties["annotations"] = var_2604ff076066_mapped
	}
	return properties
}

func (m *ExtensionMapper) FromProperties(properties map[string]*structpb.Value) *Extension {
	var s = m.New()
	if properties["id"] != nil {

		var_e122f4d8d8ab := properties["id"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_UUID).UnPack(var_e122f4d8d8ab)

		if err != nil {
			panic(err)
		}

		var_e122f4d8d8ab_mapped := new(uuid.UUID)
		*var_e122f4d8d8ab_mapped = val.(uuid.UUID)

		s.Id = var_e122f4d8d8ab_mapped
	}
	if properties["version"] != nil {

		var_a2420c656d1e := properties["version"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(var_a2420c656d1e)

		if err != nil {
			panic(err)
		}

		var_a2420c656d1e_mapped := val.(int32)

		s.Version = var_a2420c656d1e_mapped
	}
	if properties["createdBy"] != nil {

		var_4607315666fc := properties["createdBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_4607315666fc)

		if err != nil {
			panic(err)
		}

		var_4607315666fc_mapped := new(string)
		*var_4607315666fc_mapped = val.(string)

		s.CreatedBy = var_4607315666fc_mapped
	}
	if properties["updatedBy"] != nil {

		var_8943abf862b7 := properties["updatedBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_8943abf862b7)

		if err != nil {
			panic(err)
		}

		var_8943abf862b7_mapped := new(string)
		*var_8943abf862b7_mapped = val.(string)

		s.UpdatedBy = var_8943abf862b7_mapped
	}
	if properties["createdOn"] != nil {

		var_12f0a4652ed2 := properties["createdOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_12f0a4652ed2)

		if err != nil {
			panic(err)
		}

		var_12f0a4652ed2_mapped := new(time.Time)
		*var_12f0a4652ed2_mapped = val.(time.Time)

		s.CreatedOn = var_12f0a4652ed2_mapped
	}
	if properties["updatedOn"] != nil {

		var_a20e9ab44772 := properties["updatedOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_a20e9ab44772)

		if err != nil {
			panic(err)
		}

		var_a20e9ab44772_mapped := new(time.Time)
		*var_a20e9ab44772_mapped = val.(time.Time)

		s.UpdatedOn = var_a20e9ab44772_mapped
	}
	if properties["name"] != nil {

		var_7ca85575419a := properties["name"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_7ca85575419a)

		if err != nil {
			panic(err)
		}

		var_7ca85575419a_mapped := val.(string)

		s.Name = var_7ca85575419a_mapped
	}
	if properties["description"] != nil {

		var_8e76ef753a29 := properties["description"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_8e76ef753a29)

		if err != nil {
			panic(err)
		}

		var_8e76ef753a29_mapped := new(string)
		*var_8e76ef753a29_mapped = val.(string)

		s.Description = var_8e76ef753a29_mapped
	}
	if properties["selector"] != nil {

		var_1b00417318f6 := properties["selector"]
		var mappedValue = ExtensionEventSelectorMapperInstance.FromProperties(var_1b00417318f6.GetStructValue().Fields)

		var_1b00417318f6_mapped := mappedValue

		s.Selector = var_1b00417318f6_mapped
	}
	if properties["order"] != nil {

		var_36b4ba1e2da4 := properties["order"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(var_36b4ba1e2da4)

		if err != nil {
			panic(err)
		}

		var_36b4ba1e2da4_mapped := val.(int32)

		s.Order = var_36b4ba1e2da4_mapped
	}
	if properties["finalizes"] != nil {

		var_a8ffc3d38de2 := properties["finalizes"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_BOOL).UnPack(var_a8ffc3d38de2)

		if err != nil {
			panic(err)
		}

		var_a8ffc3d38de2_mapped := val.(bool)

		s.Finalizes = var_a8ffc3d38de2_mapped
	}
	if properties["sync"] != nil {

		var_81323e24accb := properties["sync"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_BOOL).UnPack(var_81323e24accb)

		if err != nil {
			panic(err)
		}

		var_81323e24accb_mapped := val.(bool)

		s.Sync = var_81323e24accb_mapped
	}
	if properties["responds"] != nil {

		var_786f7cd0c81f := properties["responds"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_BOOL).UnPack(var_786f7cd0c81f)

		if err != nil {
			panic(err)
		}

		var_786f7cd0c81f_mapped := val.(bool)

		s.Responds = var_786f7cd0c81f_mapped
	}
	if properties["call"] != nil {

		var_f134c47acfdc := properties["call"]
		var mappedValue = ExtensionExternalCallMapperInstance.FromProperties(var_f134c47acfdc.GetStructValue().Fields)

		var_f134c47acfdc_mapped := *mappedValue

		s.Call = var_f134c47acfdc_mapped
	}
	if properties["annotations"] != nil {

		var_462327809906 := properties["annotations"]
		var_462327809906_mapped := make(map[string]string)
		for k, v := range var_462327809906.GetStructValue().Fields {

			var_8813462e2973 := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_8813462e2973)

			if err != nil {
				panic(err)
			}

			var_8813462e2973_mapped := val.(string)

			var_462327809906_mapped[k] = var_8813462e2973_mapped
		}

		s.Annotations = var_462327809906_mapped
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

	var_0ec2c3e18990 := extensionFunctionCall.Host

	var var_0ec2c3e18990_mapped *structpb.Value

	var var_0ec2c3e18990_err error
	var_0ec2c3e18990_mapped, var_0ec2c3e18990_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_0ec2c3e18990)
	if var_0ec2c3e18990_err != nil {
		panic(var_0ec2c3e18990_err)
	}
	properties["host"] = var_0ec2c3e18990_mapped

	var_4be449b00e6b := extensionFunctionCall.FunctionName

	var var_4be449b00e6b_mapped *structpb.Value

	var var_4be449b00e6b_err error
	var_4be449b00e6b_mapped, var_4be449b00e6b_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_4be449b00e6b)
	if var_4be449b00e6b_err != nil {
		panic(var_4be449b00e6b_err)
	}
	properties["functionName"] = var_4be449b00e6b_mapped
	return properties
}

func (m *ExtensionFunctionCallMapper) FromProperties(properties map[string]*structpb.Value) *ExtensionFunctionCall {
	var s = m.New()
	if properties["host"] != nil {

		var_429a983e5d26 := properties["host"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_429a983e5d26)

		if err != nil {
			panic(err)
		}

		var_429a983e5d26_mapped := val.(string)

		s.Host = var_429a983e5d26_mapped
	}
	if properties["functionName"] != nil {

		var_607d5baa5529 := properties["functionName"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_607d5baa5529)

		if err != nil {
			panic(err)
		}

		var_607d5baa5529_mapped := val.(string)

		s.FunctionName = var_607d5baa5529_mapped
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

	var_31d8d425d222 := extensionHttpCall.Uri

	var var_31d8d425d222_mapped *structpb.Value

	var var_31d8d425d222_err error
	var_31d8d425d222_mapped, var_31d8d425d222_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_31d8d425d222)
	if var_31d8d425d222_err != nil {
		panic(var_31d8d425d222_err)
	}
	properties["uri"] = var_31d8d425d222_mapped

	var_77ca391879a3 := extensionHttpCall.Method

	var var_77ca391879a3_mapped *structpb.Value

	var var_77ca391879a3_err error
	var_77ca391879a3_mapped, var_77ca391879a3_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_77ca391879a3)
	if var_77ca391879a3_err != nil {
		panic(var_77ca391879a3_err)
	}
	properties["method"] = var_77ca391879a3_mapped
	return properties
}

func (m *ExtensionHttpCallMapper) FromProperties(properties map[string]*structpb.Value) *ExtensionHttpCall {
	var s = m.New()
	if properties["uri"] != nil {

		var_daab90fce09f := properties["uri"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_daab90fce09f)

		if err != nil {
			panic(err)
		}

		var_daab90fce09f_mapped := val.(string)

		s.Uri = var_daab90fce09f_mapped
	}
	if properties["method"] != nil {

		var_3dc135fb92e1 := properties["method"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_3dc135fb92e1)

		if err != nil {
			panic(err)
		}

		var_3dc135fb92e1_mapped := val.(string)

		s.Method = var_3dc135fb92e1_mapped
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

	var_cd2e09102bd6 := extensionExternalCall.FunctionCall

	if var_cd2e09102bd6 != nil {
		var var_cd2e09102bd6_mapped *structpb.Value

		var_cd2e09102bd6_mapped = structpb.NewStructValue(&structpb.Struct{Fields: ExtensionFunctionCallMapperInstance.ToProperties(var_cd2e09102bd6)})
		properties["functionCall"] = var_cd2e09102bd6_mapped
	}

	var_a5eb7b38edc7 := extensionExternalCall.HttpCall

	if var_a5eb7b38edc7 != nil {
		var var_a5eb7b38edc7_mapped *structpb.Value

		var_a5eb7b38edc7_mapped = structpb.NewStructValue(&structpb.Struct{Fields: ExtensionHttpCallMapperInstance.ToProperties(var_a5eb7b38edc7)})
		properties["httpCall"] = var_a5eb7b38edc7_mapped
	}
	return properties
}

func (m *ExtensionExternalCallMapper) FromProperties(properties map[string]*structpb.Value) *ExtensionExternalCall {
	var s = m.New()
	if properties["functionCall"] != nil {

		var_4559c5dbd312 := properties["functionCall"]
		var mappedValue = ExtensionFunctionCallMapperInstance.FromProperties(var_4559c5dbd312.GetStructValue().Fields)

		var_4559c5dbd312_mapped := mappedValue

		s.FunctionCall = var_4559c5dbd312_mapped
	}
	if properties["httpCall"] != nil {

		var_13b91f7eeafe := properties["httpCall"]
		var mappedValue = ExtensionHttpCallMapperInstance.FromProperties(var_13b91f7eeafe.GetStructValue().Fields)

		var_13b91f7eeafe_mapped := mappedValue

		s.HttpCall = var_13b91f7eeafe_mapped
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

	var_4fc2314e38cc := extensionEventSelector.Actions

	if var_4fc2314e38cc != nil {
		var var_4fc2314e38cc_mapped *structpb.Value

		var var_4fc2314e38cc_l []*structpb.Value
		for _, value := range var_4fc2314e38cc {

			var_18a54c6a28c4 := value
			var var_18a54c6a28c4_mapped *structpb.Value

			var var_18a54c6a28c4_err error
			var_18a54c6a28c4_mapped, var_18a54c6a28c4_err = types.ByResourcePropertyType(model.ResourceProperty_ENUM).Pack(string(var_18a54c6a28c4))
			if var_18a54c6a28c4_err != nil {
				panic(var_18a54c6a28c4_err)
			}

			var_4fc2314e38cc_l = append(var_4fc2314e38cc_l, var_18a54c6a28c4_mapped)
		}
		var_4fc2314e38cc_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_4fc2314e38cc_l})
		properties["actions"] = var_4fc2314e38cc_mapped
	}

	var_edc2732f650c := extensionEventSelector.RecordSelector

	if var_edc2732f650c != nil {
		var var_edc2732f650c_mapped *structpb.Value

		var_edc2732f650c_mapped = structpb.NewStructValue(&structpb.Struct{Fields: ExtensionBooleanExpressionMapperInstance.ToProperties(var_edc2732f650c)})
		properties["recordSelector"] = var_edc2732f650c_mapped
	}

	var_80be2d70207b := extensionEventSelector.Namespaces

	if var_80be2d70207b != nil {
		var var_80be2d70207b_mapped *structpb.Value

		var var_80be2d70207b_l []*structpb.Value
		for _, value := range var_80be2d70207b {

			var_94a3d6db2856 := value
			var var_94a3d6db2856_mapped *structpb.Value

			var var_94a3d6db2856_err error
			var_94a3d6db2856_mapped, var_94a3d6db2856_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_94a3d6db2856)
			if var_94a3d6db2856_err != nil {
				panic(var_94a3d6db2856_err)
			}

			var_80be2d70207b_l = append(var_80be2d70207b_l, var_94a3d6db2856_mapped)
		}
		var_80be2d70207b_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_80be2d70207b_l})
		properties["namespaces"] = var_80be2d70207b_mapped
	}

	var_9c01a041d86f := extensionEventSelector.Resources

	if var_9c01a041d86f != nil {
		var var_9c01a041d86f_mapped *structpb.Value

		var var_9c01a041d86f_l []*structpb.Value
		for _, value := range var_9c01a041d86f {

			var_4ea25fb7f5d7 := value
			var var_4ea25fb7f5d7_mapped *structpb.Value

			var var_4ea25fb7f5d7_err error
			var_4ea25fb7f5d7_mapped, var_4ea25fb7f5d7_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_4ea25fb7f5d7)
			if var_4ea25fb7f5d7_err != nil {
				panic(var_4ea25fb7f5d7_err)
			}

			var_9c01a041d86f_l = append(var_9c01a041d86f_l, var_4ea25fb7f5d7_mapped)
		}
		var_9c01a041d86f_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_9c01a041d86f_l})
		properties["resources"] = var_9c01a041d86f_mapped
	}

	var_d516645e9de8 := extensionEventSelector.Ids

	if var_d516645e9de8 != nil {
		var var_d516645e9de8_mapped *structpb.Value

		var var_d516645e9de8_l []*structpb.Value
		for _, value := range var_d516645e9de8 {

			var_fcf165bfbb2d := value
			var var_fcf165bfbb2d_mapped *structpb.Value

			var var_fcf165bfbb2d_err error
			var_fcf165bfbb2d_mapped, var_fcf165bfbb2d_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_fcf165bfbb2d)
			if var_fcf165bfbb2d_err != nil {
				panic(var_fcf165bfbb2d_err)
			}

			var_d516645e9de8_l = append(var_d516645e9de8_l, var_fcf165bfbb2d_mapped)
		}
		var_d516645e9de8_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_d516645e9de8_l})
		properties["ids"] = var_d516645e9de8_mapped
	}

	var_f0cbeba881f9 := extensionEventSelector.Annotations

	if var_f0cbeba881f9 != nil {
		var var_f0cbeba881f9_mapped *structpb.Value

		var var_f0cbeba881f9_st *structpb.Struct = new(structpb.Struct)
		var_f0cbeba881f9_st.Fields = make(map[string]*structpb.Value)
		for key, value := range var_f0cbeba881f9 {

			var_1903b30ffa91 := value
			var var_1903b30ffa91_mapped *structpb.Value

			var var_1903b30ffa91_err error
			var_1903b30ffa91_mapped, var_1903b30ffa91_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_1903b30ffa91)
			if var_1903b30ffa91_err != nil {
				panic(var_1903b30ffa91_err)
			}

			var_f0cbeba881f9_st.Fields[key] = var_1903b30ffa91_mapped
		}
		var_f0cbeba881f9_mapped = structpb.NewStructValue(var_f0cbeba881f9_st)
		properties["annotations"] = var_f0cbeba881f9_mapped
	}
	return properties
}

func (m *ExtensionEventSelectorMapper) FromProperties(properties map[string]*structpb.Value) *ExtensionEventSelector {
	var s = m.New()
	if properties["actions"] != nil {

		var_a5527711ec26 := properties["actions"]
		var_a5527711ec26_mapped := []EventAction{}
		for _, v := range var_a5527711ec26.GetListValue().Values {

			var_03954b6885ac := v
			var_03954b6885ac_mapped := (EventAction)(var_03954b6885ac.GetStringValue())

			var_a5527711ec26_mapped = append(var_a5527711ec26_mapped, var_03954b6885ac_mapped)
		}

		s.Actions = var_a5527711ec26_mapped
	}
	if properties["recordSelector"] != nil {

		var_8e73b259d453 := properties["recordSelector"]
		var mappedValue = ExtensionBooleanExpressionMapperInstance.FromProperties(var_8e73b259d453.GetStructValue().Fields)

		var_8e73b259d453_mapped := mappedValue

		s.RecordSelector = var_8e73b259d453_mapped
	}
	if properties["namespaces"] != nil {

		var_39512cecaa6b := properties["namespaces"]
		var_39512cecaa6b_mapped := []string{}
		for _, v := range var_39512cecaa6b.GetListValue().Values {

			var_19f4812009ad := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_19f4812009ad)

			if err != nil {
				panic(err)
			}

			var_19f4812009ad_mapped := val.(string)

			var_39512cecaa6b_mapped = append(var_39512cecaa6b_mapped, var_19f4812009ad_mapped)
		}

		s.Namespaces = var_39512cecaa6b_mapped
	}
	if properties["resources"] != nil {

		var_ed5adfa4ce7d := properties["resources"]
		var_ed5adfa4ce7d_mapped := []string{}
		for _, v := range var_ed5adfa4ce7d.GetListValue().Values {

			var_9b24588a0005 := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_9b24588a0005)

			if err != nil {
				panic(err)
			}

			var_9b24588a0005_mapped := val.(string)

			var_ed5adfa4ce7d_mapped = append(var_ed5adfa4ce7d_mapped, var_9b24588a0005_mapped)
		}

		s.Resources = var_ed5adfa4ce7d_mapped
	}
	if properties["ids"] != nil {

		var_dcd1f4827d67 := properties["ids"]
		var_dcd1f4827d67_mapped := []string{}
		for _, v := range var_dcd1f4827d67.GetListValue().Values {

			var_f9bbbb3f36ad := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_f9bbbb3f36ad)

			if err != nil {
				panic(err)
			}

			var_f9bbbb3f36ad_mapped := val.(string)

			var_dcd1f4827d67_mapped = append(var_dcd1f4827d67_mapped, var_f9bbbb3f36ad_mapped)
		}

		s.Ids = var_dcd1f4827d67_mapped
	}
	if properties["annotations"] != nil {

		var_6ef3a07f7bb3 := properties["annotations"]
		var_6ef3a07f7bb3_mapped := make(map[string]string)
		for k, v := range var_6ef3a07f7bb3.GetStructValue().Fields {

			var_404679fb539f := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_404679fb539f)

			if err != nil {
				panic(err)
			}

			var_404679fb539f_mapped := val.(string)

			var_6ef3a07f7bb3_mapped[k] = var_404679fb539f_mapped
		}

		s.Annotations = var_6ef3a07f7bb3_mapped
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

	var_21e4fcec6ae9 := extensionRecordSearchParams.Query

	if var_21e4fcec6ae9 != nil {
		var var_21e4fcec6ae9_mapped *structpb.Value

		var_21e4fcec6ae9_mapped = structpb.NewStructValue(&structpb.Struct{Fields: ExtensionBooleanExpressionMapperInstance.ToProperties(var_21e4fcec6ae9)})
		properties["query"] = var_21e4fcec6ae9_mapped
	}

	var_45ef13d363bb := extensionRecordSearchParams.Limit

	if var_45ef13d363bb != nil {
		var var_45ef13d363bb_mapped *structpb.Value

		var var_45ef13d363bb_err error
		var_45ef13d363bb_mapped, var_45ef13d363bb_err = types.ByResourcePropertyType(model.ResourceProperty_INT32).Pack(*var_45ef13d363bb)
		if var_45ef13d363bb_err != nil {
			panic(var_45ef13d363bb_err)
		}
		properties["limit"] = var_45ef13d363bb_mapped
	}

	var_92025d4cb123 := extensionRecordSearchParams.Offset

	if var_92025d4cb123 != nil {
		var var_92025d4cb123_mapped *structpb.Value

		var var_92025d4cb123_err error
		var_92025d4cb123_mapped, var_92025d4cb123_err = types.ByResourcePropertyType(model.ResourceProperty_INT32).Pack(*var_92025d4cb123)
		if var_92025d4cb123_err != nil {
			panic(var_92025d4cb123_err)
		}
		properties["offset"] = var_92025d4cb123_mapped
	}

	var_836ccf4bac68 := extensionRecordSearchParams.ResolveReferences

	if var_836ccf4bac68 != nil {
		var var_836ccf4bac68_mapped *structpb.Value

		var var_836ccf4bac68_l []*structpb.Value
		for _, value := range var_836ccf4bac68 {

			var_16fe7b9df01d := value
			var var_16fe7b9df01d_mapped *structpb.Value

			var var_16fe7b9df01d_err error
			var_16fe7b9df01d_mapped, var_16fe7b9df01d_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_16fe7b9df01d)
			if var_16fe7b9df01d_err != nil {
				panic(var_16fe7b9df01d_err)
			}

			var_836ccf4bac68_l = append(var_836ccf4bac68_l, var_16fe7b9df01d_mapped)
		}
		var_836ccf4bac68_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_836ccf4bac68_l})
		properties["resolveReferences"] = var_836ccf4bac68_mapped
	}
	return properties
}

func (m *ExtensionRecordSearchParamsMapper) FromProperties(properties map[string]*structpb.Value) *ExtensionRecordSearchParams {
	var s = m.New()
	if properties["query"] != nil {

		var_7038a88ea537 := properties["query"]
		var mappedValue = ExtensionBooleanExpressionMapperInstance.FromProperties(var_7038a88ea537.GetStructValue().Fields)

		var_7038a88ea537_mapped := mappedValue

		s.Query = var_7038a88ea537_mapped
	}
	if properties["limit"] != nil {

		var_c744d200a77a := properties["limit"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(var_c744d200a77a)

		if err != nil {
			panic(err)
		}

		var_c744d200a77a_mapped := new(int32)
		*var_c744d200a77a_mapped = val.(int32)

		s.Limit = var_c744d200a77a_mapped
	}
	if properties["offset"] != nil {

		var_435487351448 := properties["offset"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(var_435487351448)

		if err != nil {
			panic(err)
		}

		var_435487351448_mapped := new(int32)
		*var_435487351448_mapped = val.(int32)

		s.Offset = var_435487351448_mapped
	}
	if properties["resolveReferences"] != nil {

		var_80163a74f7f9 := properties["resolveReferences"]
		var_80163a74f7f9_mapped := []string{}
		for _, v := range var_80163a74f7f9.GetListValue().Values {

			var_7ce41f0de896 := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_7ce41f0de896)

			if err != nil {
				panic(err)
			}

			var_7ce41f0de896_mapped := val.(string)

			var_80163a74f7f9_mapped = append(var_80163a74f7f9_mapped, var_7ce41f0de896_mapped)
		}

		s.ResolveReferences = var_80163a74f7f9_mapped
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

	var_f27622656d93 := extensionEvent.Id

	if var_f27622656d93 != nil {
		var var_f27622656d93_mapped *structpb.Value

		var var_f27622656d93_err error
		var_f27622656d93_mapped, var_f27622656d93_err = types.ByResourcePropertyType(model.ResourceProperty_UUID).Pack(*var_f27622656d93)
		if var_f27622656d93_err != nil {
			panic(var_f27622656d93_err)
		}
		properties["id"] = var_f27622656d93_mapped
	}

	var_4aff342de207 := extensionEvent.Action

	var var_4aff342de207_mapped *structpb.Value

	var var_4aff342de207_err error
	var_4aff342de207_mapped, var_4aff342de207_err = types.ByResourcePropertyType(model.ResourceProperty_ENUM).Pack(string(var_4aff342de207))
	if var_4aff342de207_err != nil {
		panic(var_4aff342de207_err)
	}
	properties["action"] = var_4aff342de207_mapped

	var_cee99f5a59ef := extensionEvent.RecordSearchParams

	if var_cee99f5a59ef != nil {
		var var_cee99f5a59ef_mapped *structpb.Value

		var_cee99f5a59ef_mapped = structpb.NewStructValue(&structpb.Struct{Fields: ExtensionRecordSearchParamsMapperInstance.ToProperties(var_cee99f5a59ef)})
		properties["recordSearchParams"] = var_cee99f5a59ef_mapped
	}

	var_53b09a2c7c6c := extensionEvent.ActionSummary

	if var_53b09a2c7c6c != nil {
		var var_53b09a2c7c6c_mapped *structpb.Value

		var var_53b09a2c7c6c_err error
		var_53b09a2c7c6c_mapped, var_53b09a2c7c6c_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_53b09a2c7c6c)
		if var_53b09a2c7c6c_err != nil {
			panic(var_53b09a2c7c6c_err)
		}
		properties["actionSummary"] = var_53b09a2c7c6c_mapped
	}

	var_2bd2c347a0b4 := extensionEvent.ActionDescription

	if var_2bd2c347a0b4 != nil {
		var var_2bd2c347a0b4_mapped *structpb.Value

		var var_2bd2c347a0b4_err error
		var_2bd2c347a0b4_mapped, var_2bd2c347a0b4_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_2bd2c347a0b4)
		if var_2bd2c347a0b4_err != nil {
			panic(var_2bd2c347a0b4_err)
		}
		properties["actionDescription"] = var_2bd2c347a0b4_mapped
	}

	var_c13a025bb186 := extensionEvent.Resource

	if var_c13a025bb186 != nil {
		var var_c13a025bb186_mapped *structpb.Value

		var_c13a025bb186_mapped = structpb.NewStructValue(&structpb.Struct{Fields: ResourceMapperInstance.ToProperties(var_c13a025bb186)})
		properties["resource"] = var_c13a025bb186_mapped
	}

	var_60459a82fa49 := extensionEvent.Records

	if var_60459a82fa49 != nil {
		var var_60459a82fa49_mapped *structpb.Value

		var var_60459a82fa49_l []*structpb.Value
		for _, value := range var_60459a82fa49 {

			var_0e21a1791f71 := value
			var var_0e21a1791f71_mapped *structpb.Value

			var_0e21a1791f71_mapped = structpb.NewStructValue(&structpb.Struct{Fields: RecordMapperInstance.ToProperties(var_0e21a1791f71)})

			var_60459a82fa49_l = append(var_60459a82fa49_l, var_0e21a1791f71_mapped)
		}
		var_60459a82fa49_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_60459a82fa49_l})
		properties["records"] = var_60459a82fa49_mapped
	}

	var_6ed71debf8d2 := extensionEvent.Ids

	if var_6ed71debf8d2 != nil {
		var var_6ed71debf8d2_mapped *structpb.Value

		var var_6ed71debf8d2_l []*structpb.Value
		for _, value := range var_6ed71debf8d2 {

			var_eb282792e927 := value
			var var_eb282792e927_mapped *structpb.Value

			var var_eb282792e927_err error
			var_eb282792e927_mapped, var_eb282792e927_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_eb282792e927)
			if var_eb282792e927_err != nil {
				panic(var_eb282792e927_err)
			}

			var_6ed71debf8d2_l = append(var_6ed71debf8d2_l, var_eb282792e927_mapped)
		}
		var_6ed71debf8d2_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_6ed71debf8d2_l})
		properties["ids"] = var_6ed71debf8d2_mapped
	}

	var_2cda38ba5203 := extensionEvent.Finalizes

	if var_2cda38ba5203 != nil {
		var var_2cda38ba5203_mapped *structpb.Value

		var var_2cda38ba5203_err error
		var_2cda38ba5203_mapped, var_2cda38ba5203_err = types.ByResourcePropertyType(model.ResourceProperty_BOOL).Pack(*var_2cda38ba5203)
		if var_2cda38ba5203_err != nil {
			panic(var_2cda38ba5203_err)
		}
		properties["finalizes"] = var_2cda38ba5203_mapped
	}

	var_bf75dc8fc8ff := extensionEvent.Sync

	if var_bf75dc8fc8ff != nil {
		var var_bf75dc8fc8ff_mapped *structpb.Value

		var var_bf75dc8fc8ff_err error
		var_bf75dc8fc8ff_mapped, var_bf75dc8fc8ff_err = types.ByResourcePropertyType(model.ResourceProperty_BOOL).Pack(*var_bf75dc8fc8ff)
		if var_bf75dc8fc8ff_err != nil {
			panic(var_bf75dc8fc8ff_err)
		}
		properties["sync"] = var_bf75dc8fc8ff_mapped
	}

	var_2723210b11b8 := extensionEvent.Time

	if var_2723210b11b8 != nil {
		var var_2723210b11b8_mapped *structpb.Value

		var var_2723210b11b8_err error
		var_2723210b11b8_mapped, var_2723210b11b8_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_2723210b11b8)
		if var_2723210b11b8_err != nil {
			panic(var_2723210b11b8_err)
		}
		properties["time"] = var_2723210b11b8_mapped
	}

	var_4522aea72368 := extensionEvent.Annotations

	if var_4522aea72368 != nil {
		var var_4522aea72368_mapped *structpb.Value

		var var_4522aea72368_st *structpb.Struct = new(structpb.Struct)
		var_4522aea72368_st.Fields = make(map[string]*structpb.Value)
		for key, value := range var_4522aea72368 {

			var_966036bb29db := value
			var var_966036bb29db_mapped *structpb.Value

			var var_966036bb29db_err error
			var_966036bb29db_mapped, var_966036bb29db_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_966036bb29db)
			if var_966036bb29db_err != nil {
				panic(var_966036bb29db_err)
			}

			var_4522aea72368_st.Fields[key] = var_966036bb29db_mapped
		}
		var_4522aea72368_mapped = structpb.NewStructValue(var_4522aea72368_st)
		properties["annotations"] = var_4522aea72368_mapped
	}
	return properties
}

func (m *ExtensionEventMapper) FromProperties(properties map[string]*structpb.Value) *ExtensionEvent {
	var s = m.New()
	if properties["id"] != nil {

		var_3a0c8ccb1eb2 := properties["id"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_UUID).UnPack(var_3a0c8ccb1eb2)

		if err != nil {
			panic(err)
		}

		var_3a0c8ccb1eb2_mapped := new(uuid.UUID)
		*var_3a0c8ccb1eb2_mapped = val.(uuid.UUID)

		s.Id = var_3a0c8ccb1eb2_mapped
	}
	if properties["action"] != nil {

		var_174a2bf9b4c7 := properties["action"]
		var_174a2bf9b4c7_mapped := (EventAction)(var_174a2bf9b4c7.GetStringValue())

		s.Action = var_174a2bf9b4c7_mapped
	}
	if properties["recordSearchParams"] != nil {

		var_404979be7990 := properties["recordSearchParams"]
		var mappedValue = ExtensionRecordSearchParamsMapperInstance.FromProperties(var_404979be7990.GetStructValue().Fields)

		var_404979be7990_mapped := mappedValue

		s.RecordSearchParams = var_404979be7990_mapped
	}
	if properties["actionSummary"] != nil {

		var_56c2ec41581d := properties["actionSummary"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_56c2ec41581d)

		if err != nil {
			panic(err)
		}

		var_56c2ec41581d_mapped := new(string)
		*var_56c2ec41581d_mapped = val.(string)

		s.ActionSummary = var_56c2ec41581d_mapped
	}
	if properties["actionDescription"] != nil {

		var_84693878955b := properties["actionDescription"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_84693878955b)

		if err != nil {
			panic(err)
		}

		var_84693878955b_mapped := new(string)
		*var_84693878955b_mapped = val.(string)

		s.ActionDescription = var_84693878955b_mapped
	}
	if properties["resource"] != nil {

		var_2d5258887910 := properties["resource"]
		var_2d5258887910_mapped := ResourceMapperInstance.FromProperties(var_2d5258887910.GetStructValue().Fields)

		s.Resource = var_2d5258887910_mapped
	}
	if properties["records"] != nil {

		var_08e5d73b89f3 := properties["records"]
		var_08e5d73b89f3_mapped := []*Record{}
		for _, v := range var_08e5d73b89f3.GetListValue().Values {

			var_66fa73e2dceb := v
			var_66fa73e2dceb_mapped := RecordMapperInstance.FromProperties(var_66fa73e2dceb.GetStructValue().Fields)

			var_08e5d73b89f3_mapped = append(var_08e5d73b89f3_mapped, var_66fa73e2dceb_mapped)
		}

		s.Records = var_08e5d73b89f3_mapped
	}
	if properties["ids"] != nil {

		var_48b1737d3d48 := properties["ids"]
		var_48b1737d3d48_mapped := []string{}
		for _, v := range var_48b1737d3d48.GetListValue().Values {

			var_d422c7c9b343 := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_d422c7c9b343)

			if err != nil {
				panic(err)
			}

			var_d422c7c9b343_mapped := val.(string)

			var_48b1737d3d48_mapped = append(var_48b1737d3d48_mapped, var_d422c7c9b343_mapped)
		}

		s.Ids = var_48b1737d3d48_mapped
	}
	if properties["finalizes"] != nil {

		var_2d43513de68a := properties["finalizes"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_BOOL).UnPack(var_2d43513de68a)

		if err != nil {
			panic(err)
		}

		var_2d43513de68a_mapped := new(bool)
		*var_2d43513de68a_mapped = val.(bool)

		s.Finalizes = var_2d43513de68a_mapped
	}
	if properties["sync"] != nil {

		var_c26e73262e7a := properties["sync"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_BOOL).UnPack(var_c26e73262e7a)

		if err != nil {
			panic(err)
		}

		var_c26e73262e7a_mapped := new(bool)
		*var_c26e73262e7a_mapped = val.(bool)

		s.Sync = var_c26e73262e7a_mapped
	}
	if properties["time"] != nil {

		var_96a5b611e4fc := properties["time"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_96a5b611e4fc)

		if err != nil {
			panic(err)
		}

		var_96a5b611e4fc_mapped := new(time.Time)
		*var_96a5b611e4fc_mapped = val.(time.Time)

		s.Time = var_96a5b611e4fc_mapped
	}
	if properties["annotations"] != nil {

		var_4008a239205a := properties["annotations"]
		var_4008a239205a_mapped := make(map[string]string)
		for k, v := range var_4008a239205a.GetStructValue().Fields {

			var_ae6a9b889651 := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_ae6a9b889651)

			if err != nil {
				panic(err)
			}

			var_ae6a9b889651_mapped := val.(string)

			var_4008a239205a_mapped[k] = var_ae6a9b889651_mapped
		}

		s.Annotations = var_4008a239205a_mapped
	}
	return s
}
