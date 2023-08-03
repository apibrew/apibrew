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

	var_e427cb8093b8 := extension.Id

	if var_e427cb8093b8 != nil {
		var var_e427cb8093b8_mapped *structpb.Value

		var var_e427cb8093b8_err error
		var_e427cb8093b8_mapped, var_e427cb8093b8_err = types.ByResourcePropertyType(model.ResourceProperty_UUID).Pack(*var_e427cb8093b8)
		if var_e427cb8093b8_err != nil {
			panic(var_e427cb8093b8_err)
		}
		properties["id"] = var_e427cb8093b8_mapped
	}

	var_fd9f6c6e5e51 := extension.Version

	var var_fd9f6c6e5e51_mapped *structpb.Value

	var var_fd9f6c6e5e51_err error
	var_fd9f6c6e5e51_mapped, var_fd9f6c6e5e51_err = types.ByResourcePropertyType(model.ResourceProperty_INT32).Pack(var_fd9f6c6e5e51)
	if var_fd9f6c6e5e51_err != nil {
		panic(var_fd9f6c6e5e51_err)
	}
	properties["version"] = var_fd9f6c6e5e51_mapped

	var_2ea68554a469 := extension.CreatedBy

	if var_2ea68554a469 != nil {
		var var_2ea68554a469_mapped *structpb.Value

		var var_2ea68554a469_err error
		var_2ea68554a469_mapped, var_2ea68554a469_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_2ea68554a469)
		if var_2ea68554a469_err != nil {
			panic(var_2ea68554a469_err)
		}
		properties["createdBy"] = var_2ea68554a469_mapped
	}

	var_216a1fadb2e7 := extension.UpdatedBy

	if var_216a1fadb2e7 != nil {
		var var_216a1fadb2e7_mapped *structpb.Value

		var var_216a1fadb2e7_err error
		var_216a1fadb2e7_mapped, var_216a1fadb2e7_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_216a1fadb2e7)
		if var_216a1fadb2e7_err != nil {
			panic(var_216a1fadb2e7_err)
		}
		properties["updatedBy"] = var_216a1fadb2e7_mapped
	}

	var_6a9fb15adfe1 := extension.CreatedOn

	if var_6a9fb15adfe1 != nil {
		var var_6a9fb15adfe1_mapped *structpb.Value

		var var_6a9fb15adfe1_err error
		var_6a9fb15adfe1_mapped, var_6a9fb15adfe1_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_6a9fb15adfe1)
		if var_6a9fb15adfe1_err != nil {
			panic(var_6a9fb15adfe1_err)
		}
		properties["createdOn"] = var_6a9fb15adfe1_mapped
	}

	var_8b1814cf5f35 := extension.UpdatedOn

	if var_8b1814cf5f35 != nil {
		var var_8b1814cf5f35_mapped *structpb.Value

		var var_8b1814cf5f35_err error
		var_8b1814cf5f35_mapped, var_8b1814cf5f35_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_8b1814cf5f35)
		if var_8b1814cf5f35_err != nil {
			panic(var_8b1814cf5f35_err)
		}
		properties["updatedOn"] = var_8b1814cf5f35_mapped
	}

	var_249c41645085 := extension.Name

	var var_249c41645085_mapped *structpb.Value

	var var_249c41645085_err error
	var_249c41645085_mapped, var_249c41645085_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_249c41645085)
	if var_249c41645085_err != nil {
		panic(var_249c41645085_err)
	}
	properties["name"] = var_249c41645085_mapped

	var_2716041a081d := extension.Description

	if var_2716041a081d != nil {
		var var_2716041a081d_mapped *structpb.Value

		var var_2716041a081d_err error
		var_2716041a081d_mapped, var_2716041a081d_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_2716041a081d)
		if var_2716041a081d_err != nil {
			panic(var_2716041a081d_err)
		}
		properties["description"] = var_2716041a081d_mapped
	}

	var_141b56f87d2d := extension.Selector

	if var_141b56f87d2d != nil {
		var var_141b56f87d2d_mapped *structpb.Value

		var_141b56f87d2d_mapped = structpb.NewStructValue(&structpb.Struct{Fields: ExtensionEventSelectorMapperInstance.ToProperties(var_141b56f87d2d)})
		properties["selector"] = var_141b56f87d2d_mapped
	}

	var_2e9e38f9d387 := extension.Order

	var var_2e9e38f9d387_mapped *structpb.Value

	var var_2e9e38f9d387_err error
	var_2e9e38f9d387_mapped, var_2e9e38f9d387_err = types.ByResourcePropertyType(model.ResourceProperty_INT32).Pack(var_2e9e38f9d387)
	if var_2e9e38f9d387_err != nil {
		panic(var_2e9e38f9d387_err)
	}
	properties["order"] = var_2e9e38f9d387_mapped

	var_5dbfa00e9129 := extension.Finalizes

	var var_5dbfa00e9129_mapped *structpb.Value

	var var_5dbfa00e9129_err error
	var_5dbfa00e9129_mapped, var_5dbfa00e9129_err = types.ByResourcePropertyType(model.ResourceProperty_BOOL).Pack(var_5dbfa00e9129)
	if var_5dbfa00e9129_err != nil {
		panic(var_5dbfa00e9129_err)
	}
	properties["finalizes"] = var_5dbfa00e9129_mapped

	var_6c38e251daaf := extension.Sync

	var var_6c38e251daaf_mapped *structpb.Value

	var var_6c38e251daaf_err error
	var_6c38e251daaf_mapped, var_6c38e251daaf_err = types.ByResourcePropertyType(model.ResourceProperty_BOOL).Pack(var_6c38e251daaf)
	if var_6c38e251daaf_err != nil {
		panic(var_6c38e251daaf_err)
	}
	properties["sync"] = var_6c38e251daaf_mapped

	var_08e39f582ccf := extension.Responds

	var var_08e39f582ccf_mapped *structpb.Value

	var var_08e39f582ccf_err error
	var_08e39f582ccf_mapped, var_08e39f582ccf_err = types.ByResourcePropertyType(model.ResourceProperty_BOOL).Pack(var_08e39f582ccf)
	if var_08e39f582ccf_err != nil {
		panic(var_08e39f582ccf_err)
	}
	properties["responds"] = var_08e39f582ccf_mapped

	var_b6831c7c505d := extension.Call

	var var_b6831c7c505d_mapped *structpb.Value

	var_b6831c7c505d_mapped = structpb.NewStructValue(&structpb.Struct{Fields: ExtensionExternalCallMapperInstance.ToProperties(&var_b6831c7c505d)})
	properties["call"] = var_b6831c7c505d_mapped

	var_3bef49e17425 := extension.Annotations

	if var_3bef49e17425 != nil {
		var var_3bef49e17425_mapped *structpb.Value

		var var_3bef49e17425_st *structpb.Struct = new(structpb.Struct)
		var_3bef49e17425_st.Fields = make(map[string]*structpb.Value)
		for key, value := range var_3bef49e17425 {

			var_c09db3916e53 := value
			var var_c09db3916e53_mapped *structpb.Value

			var var_c09db3916e53_err error
			var_c09db3916e53_mapped, var_c09db3916e53_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_c09db3916e53)
			if var_c09db3916e53_err != nil {
				panic(var_c09db3916e53_err)
			}

			var_3bef49e17425_st.Fields[key] = var_c09db3916e53_mapped
		}
		var_3bef49e17425_mapped = structpb.NewStructValue(var_3bef49e17425_st)
		properties["annotations"] = var_3bef49e17425_mapped
	}
	return properties
}

func (m *ExtensionMapper) FromProperties(properties map[string]*structpb.Value) *Extension {
	var s = m.New()
	if properties["id"] != nil {

		var_d66a204c5294 := properties["id"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_UUID).UnPack(var_d66a204c5294)

		if err != nil {
			panic(err)
		}

		var_d66a204c5294_mapped := new(uuid.UUID)
		*var_d66a204c5294_mapped = val.(uuid.UUID)

		s.Id = var_d66a204c5294_mapped
	}
	if properties["version"] != nil {

		var_7a907a7172a8 := properties["version"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(var_7a907a7172a8)

		if err != nil {
			panic(err)
		}

		var_7a907a7172a8_mapped := val.(int32)

		s.Version = var_7a907a7172a8_mapped
	}
	if properties["createdBy"] != nil {

		var_021e2ec8bf29 := properties["createdBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_021e2ec8bf29)

		if err != nil {
			panic(err)
		}

		var_021e2ec8bf29_mapped := new(string)
		*var_021e2ec8bf29_mapped = val.(string)

		s.CreatedBy = var_021e2ec8bf29_mapped
	}
	if properties["updatedBy"] != nil {

		var_2a6cd85eb13e := properties["updatedBy"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_2a6cd85eb13e)

		if err != nil {
			panic(err)
		}

		var_2a6cd85eb13e_mapped := new(string)
		*var_2a6cd85eb13e_mapped = val.(string)

		s.UpdatedBy = var_2a6cd85eb13e_mapped
	}
	if properties["createdOn"] != nil {

		var_2f90baf5f288 := properties["createdOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_2f90baf5f288)

		if err != nil {
			panic(err)
		}

		var_2f90baf5f288_mapped := new(time.Time)
		*var_2f90baf5f288_mapped = val.(time.Time)

		s.CreatedOn = var_2f90baf5f288_mapped
	}
	if properties["updatedOn"] != nil {

		var_b6b8628e2706 := properties["updatedOn"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_b6b8628e2706)

		if err != nil {
			panic(err)
		}

		var_b6b8628e2706_mapped := new(time.Time)
		*var_b6b8628e2706_mapped = val.(time.Time)

		s.UpdatedOn = var_b6b8628e2706_mapped
	}
	if properties["name"] != nil {

		var_e038e44cc939 := properties["name"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_e038e44cc939)

		if err != nil {
			panic(err)
		}

		var_e038e44cc939_mapped := val.(string)

		s.Name = var_e038e44cc939_mapped
	}
	if properties["description"] != nil {

		var_3fa301323472 := properties["description"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_3fa301323472)

		if err != nil {
			panic(err)
		}

		var_3fa301323472_mapped := new(string)
		*var_3fa301323472_mapped = val.(string)

		s.Description = var_3fa301323472_mapped
	}
	if properties["selector"] != nil {

		var_a3755afb024f := properties["selector"]
		var mappedValue = ExtensionEventSelectorMapperInstance.FromProperties(var_a3755afb024f.GetStructValue().Fields)

		var_a3755afb024f_mapped := mappedValue

		s.Selector = var_a3755afb024f_mapped
	}
	if properties["order"] != nil {

		var_4b1afa92fe58 := properties["order"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(var_4b1afa92fe58)

		if err != nil {
			panic(err)
		}

		var_4b1afa92fe58_mapped := val.(int32)

		s.Order = var_4b1afa92fe58_mapped
	}
	if properties["finalizes"] != nil {

		var_ede9a6d71d32 := properties["finalizes"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_BOOL).UnPack(var_ede9a6d71d32)

		if err != nil {
			panic(err)
		}

		var_ede9a6d71d32_mapped := val.(bool)

		s.Finalizes = var_ede9a6d71d32_mapped
	}
	if properties["sync"] != nil {

		var_b2b66fce0d66 := properties["sync"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_BOOL).UnPack(var_b2b66fce0d66)

		if err != nil {
			panic(err)
		}

		var_b2b66fce0d66_mapped := val.(bool)

		s.Sync = var_b2b66fce0d66_mapped
	}
	if properties["responds"] != nil {

		var_3157762aa4d4 := properties["responds"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_BOOL).UnPack(var_3157762aa4d4)

		if err != nil {
			panic(err)
		}

		var_3157762aa4d4_mapped := val.(bool)

		s.Responds = var_3157762aa4d4_mapped
	}
	if properties["call"] != nil {

		var_51949384c561 := properties["call"]
		var mappedValue = ExtensionExternalCallMapperInstance.FromProperties(var_51949384c561.GetStructValue().Fields)

		var_51949384c561_mapped := *mappedValue

		s.Call = var_51949384c561_mapped
	}
	if properties["annotations"] != nil {

		var_dab9256fefa0 := properties["annotations"]
		var_dab9256fefa0_mapped := make(map[string]string)
		for k, v := range var_dab9256fefa0.GetStructValue().Fields {

			var_311c227ae957 := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_311c227ae957)

			if err != nil {
				panic(err)
			}

			var_311c227ae957_mapped := val.(string)

			var_dab9256fefa0_mapped[k] = var_311c227ae957_mapped
		}

		s.Annotations = var_dab9256fefa0_mapped
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

	var_1a9f208541e1 := extensionFunctionCall.Host

	var var_1a9f208541e1_mapped *structpb.Value

	var var_1a9f208541e1_err error
	var_1a9f208541e1_mapped, var_1a9f208541e1_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_1a9f208541e1)
	if var_1a9f208541e1_err != nil {
		panic(var_1a9f208541e1_err)
	}
	properties["host"] = var_1a9f208541e1_mapped

	var_8fba9418111d := extensionFunctionCall.FunctionName

	var var_8fba9418111d_mapped *structpb.Value

	var var_8fba9418111d_err error
	var_8fba9418111d_mapped, var_8fba9418111d_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_8fba9418111d)
	if var_8fba9418111d_err != nil {
		panic(var_8fba9418111d_err)
	}
	properties["functionName"] = var_8fba9418111d_mapped
	return properties
}

func (m *ExtensionFunctionCallMapper) FromProperties(properties map[string]*structpb.Value) *ExtensionFunctionCall {
	var s = m.New()
	if properties["host"] != nil {

		var_9b96511e2840 := properties["host"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_9b96511e2840)

		if err != nil {
			panic(err)
		}

		var_9b96511e2840_mapped := val.(string)

		s.Host = var_9b96511e2840_mapped
	}
	if properties["functionName"] != nil {

		var_779aef461bc9 := properties["functionName"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_779aef461bc9)

		if err != nil {
			panic(err)
		}

		var_779aef461bc9_mapped := val.(string)

		s.FunctionName = var_779aef461bc9_mapped
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

	var_cfeb8b72aeb6 := extensionHttpCall.Uri

	var var_cfeb8b72aeb6_mapped *structpb.Value

	var var_cfeb8b72aeb6_err error
	var_cfeb8b72aeb6_mapped, var_cfeb8b72aeb6_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_cfeb8b72aeb6)
	if var_cfeb8b72aeb6_err != nil {
		panic(var_cfeb8b72aeb6_err)
	}
	properties["uri"] = var_cfeb8b72aeb6_mapped

	var_148a83c691f6 := extensionHttpCall.Method

	var var_148a83c691f6_mapped *structpb.Value

	var var_148a83c691f6_err error
	var_148a83c691f6_mapped, var_148a83c691f6_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_148a83c691f6)
	if var_148a83c691f6_err != nil {
		panic(var_148a83c691f6_err)
	}
	properties["method"] = var_148a83c691f6_mapped
	return properties
}

func (m *ExtensionHttpCallMapper) FromProperties(properties map[string]*structpb.Value) *ExtensionHttpCall {
	var s = m.New()
	if properties["uri"] != nil {

		var_93dc0c8981ad := properties["uri"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_93dc0c8981ad)

		if err != nil {
			panic(err)
		}

		var_93dc0c8981ad_mapped := val.(string)

		s.Uri = var_93dc0c8981ad_mapped
	}
	if properties["method"] != nil {

		var_89d5e7bf5a78 := properties["method"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_89d5e7bf5a78)

		if err != nil {
			panic(err)
		}

		var_89d5e7bf5a78_mapped := val.(string)

		s.Method = var_89d5e7bf5a78_mapped
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

	var_abaa748e6951 := extensionExternalCall.FunctionCall

	if var_abaa748e6951 != nil {
		var var_abaa748e6951_mapped *structpb.Value

		var_abaa748e6951_mapped = structpb.NewStructValue(&structpb.Struct{Fields: ExtensionFunctionCallMapperInstance.ToProperties(var_abaa748e6951)})
		properties["functionCall"] = var_abaa748e6951_mapped
	}

	var_b5dcb64b5585 := extensionExternalCall.HttpCall

	if var_b5dcb64b5585 != nil {
		var var_b5dcb64b5585_mapped *structpb.Value

		var_b5dcb64b5585_mapped = structpb.NewStructValue(&structpb.Struct{Fields: ExtensionHttpCallMapperInstance.ToProperties(var_b5dcb64b5585)})
		properties["httpCall"] = var_b5dcb64b5585_mapped
	}
	return properties
}

func (m *ExtensionExternalCallMapper) FromProperties(properties map[string]*structpb.Value) *ExtensionExternalCall {
	var s = m.New()
	if properties["functionCall"] != nil {

		var_7e3bcb7e4340 := properties["functionCall"]
		var mappedValue = ExtensionFunctionCallMapperInstance.FromProperties(var_7e3bcb7e4340.GetStructValue().Fields)

		var_7e3bcb7e4340_mapped := mappedValue

		s.FunctionCall = var_7e3bcb7e4340_mapped
	}
	if properties["httpCall"] != nil {

		var_59bb336c4078 := properties["httpCall"]
		var mappedValue = ExtensionHttpCallMapperInstance.FromProperties(var_59bb336c4078.GetStructValue().Fields)

		var_59bb336c4078_mapped := mappedValue

		s.HttpCall = var_59bb336c4078_mapped
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

	var_c93c350157c4 := extensionEventSelector.Actions

	if var_c93c350157c4 != nil {
		var var_c93c350157c4_mapped *structpb.Value

		var var_c93c350157c4_l []*structpb.Value
		for _, value := range var_c93c350157c4 {

			var_a4b280dbe163 := value
			var var_a4b280dbe163_mapped *structpb.Value

			var var_a4b280dbe163_err error
			var_a4b280dbe163_mapped, var_a4b280dbe163_err = types.ByResourcePropertyType(model.ResourceProperty_ENUM).Pack(string(var_a4b280dbe163))
			if var_a4b280dbe163_err != nil {
				panic(var_a4b280dbe163_err)
			}

			var_c93c350157c4_l = append(var_c93c350157c4_l, var_a4b280dbe163_mapped)
		}
		var_c93c350157c4_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_c93c350157c4_l})
		properties["actions"] = var_c93c350157c4_mapped
	}

	var_33bf9b26e13d := extensionEventSelector.RecordSelector

	if var_33bf9b26e13d != nil {
		var var_33bf9b26e13d_mapped *structpb.Value

		var_33bf9b26e13d_mapped = structpb.NewStructValue(&structpb.Struct{Fields: ExtensionBooleanExpressionMapperInstance.ToProperties(var_33bf9b26e13d)})
		properties["recordSelector"] = var_33bf9b26e13d_mapped
	}

	var_0dd468ca2b4c := extensionEventSelector.Namespaces

	if var_0dd468ca2b4c != nil {
		var var_0dd468ca2b4c_mapped *structpb.Value

		var var_0dd468ca2b4c_l []*structpb.Value
		for _, value := range var_0dd468ca2b4c {

			var_eac9e65278dd := value
			var var_eac9e65278dd_mapped *structpb.Value

			var var_eac9e65278dd_err error
			var_eac9e65278dd_mapped, var_eac9e65278dd_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_eac9e65278dd)
			if var_eac9e65278dd_err != nil {
				panic(var_eac9e65278dd_err)
			}

			var_0dd468ca2b4c_l = append(var_0dd468ca2b4c_l, var_eac9e65278dd_mapped)
		}
		var_0dd468ca2b4c_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_0dd468ca2b4c_l})
		properties["namespaces"] = var_0dd468ca2b4c_mapped
	}

	var_97d3a65583bb := extensionEventSelector.Resources

	if var_97d3a65583bb != nil {
		var var_97d3a65583bb_mapped *structpb.Value

		var var_97d3a65583bb_l []*structpb.Value
		for _, value := range var_97d3a65583bb {

			var_3ce17ddd6089 := value
			var var_3ce17ddd6089_mapped *structpb.Value

			var var_3ce17ddd6089_err error
			var_3ce17ddd6089_mapped, var_3ce17ddd6089_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_3ce17ddd6089)
			if var_3ce17ddd6089_err != nil {
				panic(var_3ce17ddd6089_err)
			}

			var_97d3a65583bb_l = append(var_97d3a65583bb_l, var_3ce17ddd6089_mapped)
		}
		var_97d3a65583bb_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_97d3a65583bb_l})
		properties["resources"] = var_97d3a65583bb_mapped
	}

	var_373b3d0d2cd2 := extensionEventSelector.Ids

	if var_373b3d0d2cd2 != nil {
		var var_373b3d0d2cd2_mapped *structpb.Value

		var var_373b3d0d2cd2_l []*structpb.Value
		for _, value := range var_373b3d0d2cd2 {

			var_31648865f7e5 := value
			var var_31648865f7e5_mapped *structpb.Value

			var var_31648865f7e5_err error
			var_31648865f7e5_mapped, var_31648865f7e5_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_31648865f7e5)
			if var_31648865f7e5_err != nil {
				panic(var_31648865f7e5_err)
			}

			var_373b3d0d2cd2_l = append(var_373b3d0d2cd2_l, var_31648865f7e5_mapped)
		}
		var_373b3d0d2cd2_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_373b3d0d2cd2_l})
		properties["ids"] = var_373b3d0d2cd2_mapped
	}

	var_94ac719927fb := extensionEventSelector.Annotations

	if var_94ac719927fb != nil {
		var var_94ac719927fb_mapped *structpb.Value

		var var_94ac719927fb_st *structpb.Struct = new(structpb.Struct)
		var_94ac719927fb_st.Fields = make(map[string]*structpb.Value)
		for key, value := range var_94ac719927fb {

			var_0c8dd6f95cb1 := value
			var var_0c8dd6f95cb1_mapped *structpb.Value

			var var_0c8dd6f95cb1_err error
			var_0c8dd6f95cb1_mapped, var_0c8dd6f95cb1_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_0c8dd6f95cb1)
			if var_0c8dd6f95cb1_err != nil {
				panic(var_0c8dd6f95cb1_err)
			}

			var_94ac719927fb_st.Fields[key] = var_0c8dd6f95cb1_mapped
		}
		var_94ac719927fb_mapped = structpb.NewStructValue(var_94ac719927fb_st)
		properties["annotations"] = var_94ac719927fb_mapped
	}
	return properties
}

func (m *ExtensionEventSelectorMapper) FromProperties(properties map[string]*structpb.Value) *ExtensionEventSelector {
	var s = m.New()
	if properties["actions"] != nil {

		var_6acd3ca90246 := properties["actions"]
		var_6acd3ca90246_mapped := []EventAction{}
		for _, v := range var_6acd3ca90246.GetListValue().Values {

			var_9d00573044ba := v
			var_9d00573044ba_mapped := (EventAction)(var_9d00573044ba.GetStringValue())

			var_6acd3ca90246_mapped = append(var_6acd3ca90246_mapped, var_9d00573044ba_mapped)
		}

		s.Actions = var_6acd3ca90246_mapped
	}
	if properties["recordSelector"] != nil {

		var_51a3e3e9f09f := properties["recordSelector"]
		var mappedValue = ExtensionBooleanExpressionMapperInstance.FromProperties(var_51a3e3e9f09f.GetStructValue().Fields)

		var_51a3e3e9f09f_mapped := mappedValue

		s.RecordSelector = var_51a3e3e9f09f_mapped
	}
	if properties["namespaces"] != nil {

		var_f3c471449795 := properties["namespaces"]
		var_f3c471449795_mapped := []string{}
		for _, v := range var_f3c471449795.GetListValue().Values {

			var_015f55efee75 := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_015f55efee75)

			if err != nil {
				panic(err)
			}

			var_015f55efee75_mapped := val.(string)

			var_f3c471449795_mapped = append(var_f3c471449795_mapped, var_015f55efee75_mapped)
		}

		s.Namespaces = var_f3c471449795_mapped
	}
	if properties["resources"] != nil {

		var_65a28aff0b38 := properties["resources"]
		var_65a28aff0b38_mapped := []string{}
		for _, v := range var_65a28aff0b38.GetListValue().Values {

			var_957f920f5d14 := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_957f920f5d14)

			if err != nil {
				panic(err)
			}

			var_957f920f5d14_mapped := val.(string)

			var_65a28aff0b38_mapped = append(var_65a28aff0b38_mapped, var_957f920f5d14_mapped)
		}

		s.Resources = var_65a28aff0b38_mapped
	}
	if properties["ids"] != nil {

		var_827f89c110df := properties["ids"]
		var_827f89c110df_mapped := []string{}
		for _, v := range var_827f89c110df.GetListValue().Values {

			var_1eab0840ba92 := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_1eab0840ba92)

			if err != nil {
				panic(err)
			}

			var_1eab0840ba92_mapped := val.(string)

			var_827f89c110df_mapped = append(var_827f89c110df_mapped, var_1eab0840ba92_mapped)
		}

		s.Ids = var_827f89c110df_mapped
	}
	if properties["annotations"] != nil {

		var_9c65bc056d32 := properties["annotations"]
		var_9c65bc056d32_mapped := make(map[string]string)
		for k, v := range var_9c65bc056d32.GetStructValue().Fields {

			var_0dc9e5687424 := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_0dc9e5687424)

			if err != nil {
				panic(err)
			}

			var_0dc9e5687424_mapped := val.(string)

			var_9c65bc056d32_mapped[k] = var_0dc9e5687424_mapped
		}

		s.Annotations = var_9c65bc056d32_mapped
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

	var_dc453c8512ea := extensionRecordSearchParams.Query

	if var_dc453c8512ea != nil {
		var var_dc453c8512ea_mapped *structpb.Value

		var_dc453c8512ea_mapped = structpb.NewStructValue(&structpb.Struct{Fields: ExtensionBooleanExpressionMapperInstance.ToProperties(var_dc453c8512ea)})
		properties["query"] = var_dc453c8512ea_mapped
	}

	var_ab98a4d693b2 := extensionRecordSearchParams.Limit

	if var_ab98a4d693b2 != nil {
		var var_ab98a4d693b2_mapped *structpb.Value

		var var_ab98a4d693b2_err error
		var_ab98a4d693b2_mapped, var_ab98a4d693b2_err = types.ByResourcePropertyType(model.ResourceProperty_INT32).Pack(*var_ab98a4d693b2)
		if var_ab98a4d693b2_err != nil {
			panic(var_ab98a4d693b2_err)
		}
		properties["limit"] = var_ab98a4d693b2_mapped
	}

	var_3b217862a90c := extensionRecordSearchParams.Offset

	if var_3b217862a90c != nil {
		var var_3b217862a90c_mapped *structpb.Value

		var var_3b217862a90c_err error
		var_3b217862a90c_mapped, var_3b217862a90c_err = types.ByResourcePropertyType(model.ResourceProperty_INT32).Pack(*var_3b217862a90c)
		if var_3b217862a90c_err != nil {
			panic(var_3b217862a90c_err)
		}
		properties["offset"] = var_3b217862a90c_mapped
	}

	var_6165b1b2edd4 := extensionRecordSearchParams.ResolveReferences

	if var_6165b1b2edd4 != nil {
		var var_6165b1b2edd4_mapped *structpb.Value

		var var_6165b1b2edd4_l []*structpb.Value
		for _, value := range var_6165b1b2edd4 {

			var_a3d2304bb588 := value
			var var_a3d2304bb588_mapped *structpb.Value

			var var_a3d2304bb588_err error
			var_a3d2304bb588_mapped, var_a3d2304bb588_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_a3d2304bb588)
			if var_a3d2304bb588_err != nil {
				panic(var_a3d2304bb588_err)
			}

			var_6165b1b2edd4_l = append(var_6165b1b2edd4_l, var_a3d2304bb588_mapped)
		}
		var_6165b1b2edd4_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_6165b1b2edd4_l})
		properties["resolveReferences"] = var_6165b1b2edd4_mapped
	}
	return properties
}

func (m *ExtensionRecordSearchParamsMapper) FromProperties(properties map[string]*structpb.Value) *ExtensionRecordSearchParams {
	var s = m.New()
	if properties["query"] != nil {

		var_877bae58b3dc := properties["query"]
		var mappedValue = ExtensionBooleanExpressionMapperInstance.FromProperties(var_877bae58b3dc.GetStructValue().Fields)

		var_877bae58b3dc_mapped := mappedValue

		s.Query = var_877bae58b3dc_mapped
	}
	if properties["limit"] != nil {

		var_e978c85239a2 := properties["limit"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(var_e978c85239a2)

		if err != nil {
			panic(err)
		}

		var_e978c85239a2_mapped := new(int32)
		*var_e978c85239a2_mapped = val.(int32)

		s.Limit = var_e978c85239a2_mapped
	}
	if properties["offset"] != nil {

		var_d769ea7075fe := properties["offset"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_INT32).UnPack(var_d769ea7075fe)

		if err != nil {
			panic(err)
		}

		var_d769ea7075fe_mapped := new(int32)
		*var_d769ea7075fe_mapped = val.(int32)

		s.Offset = var_d769ea7075fe_mapped
	}
	if properties["resolveReferences"] != nil {

		var_263a7f3103ed := properties["resolveReferences"]
		var_263a7f3103ed_mapped := []string{}
		for _, v := range var_263a7f3103ed.GetListValue().Values {

			var_432c95f82281 := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_432c95f82281)

			if err != nil {
				panic(err)
			}

			var_432c95f82281_mapped := val.(string)

			var_263a7f3103ed_mapped = append(var_263a7f3103ed_mapped, var_432c95f82281_mapped)
		}

		s.ResolveReferences = var_263a7f3103ed_mapped
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

	var_035fd6b8f644 := extensionEvent.Id

	if var_035fd6b8f644 != nil {
		var var_035fd6b8f644_mapped *structpb.Value

		var var_035fd6b8f644_err error
		var_035fd6b8f644_mapped, var_035fd6b8f644_err = types.ByResourcePropertyType(model.ResourceProperty_UUID).Pack(*var_035fd6b8f644)
		if var_035fd6b8f644_err != nil {
			panic(var_035fd6b8f644_err)
		}
		properties["id"] = var_035fd6b8f644_mapped
	}

	var_8f5b377463b8 := extensionEvent.Action

	var var_8f5b377463b8_mapped *structpb.Value

	var var_8f5b377463b8_err error
	var_8f5b377463b8_mapped, var_8f5b377463b8_err = types.ByResourcePropertyType(model.ResourceProperty_ENUM).Pack(string(var_8f5b377463b8))
	if var_8f5b377463b8_err != nil {
		panic(var_8f5b377463b8_err)
	}
	properties["action"] = var_8f5b377463b8_mapped

	var_1c0f5fcace06 := extensionEvent.RecordSearchParams

	if var_1c0f5fcace06 != nil {
		var var_1c0f5fcace06_mapped *structpb.Value

		var_1c0f5fcace06_mapped = structpb.NewStructValue(&structpb.Struct{Fields: ExtensionRecordSearchParamsMapperInstance.ToProperties(var_1c0f5fcace06)})
		properties["recordSearchParams"] = var_1c0f5fcace06_mapped
	}

	var_768ddece59de := extensionEvent.ActionSummary

	if var_768ddece59de != nil {
		var var_768ddece59de_mapped *structpb.Value

		var var_768ddece59de_err error
		var_768ddece59de_mapped, var_768ddece59de_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_768ddece59de)
		if var_768ddece59de_err != nil {
			panic(var_768ddece59de_err)
		}
		properties["actionSummary"] = var_768ddece59de_mapped
	}

	var_b0b28113b5d1 := extensionEvent.ActionDescription

	if var_b0b28113b5d1 != nil {
		var var_b0b28113b5d1_mapped *structpb.Value

		var var_b0b28113b5d1_err error
		var_b0b28113b5d1_mapped, var_b0b28113b5d1_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(*var_b0b28113b5d1)
		if var_b0b28113b5d1_err != nil {
			panic(var_b0b28113b5d1_err)
		}
		properties["actionDescription"] = var_b0b28113b5d1_mapped
	}

	var_0e8b17a30c9f := extensionEvent.Resource

	if var_0e8b17a30c9f != nil {
		var var_0e8b17a30c9f_mapped *structpb.Value

		var_0e8b17a30c9f_mapped = structpb.NewStructValue(&structpb.Struct{Fields: ResourceMapperInstance.ToProperties(var_0e8b17a30c9f)})
		properties["resource"] = var_0e8b17a30c9f_mapped
	}

	var_8039bf614513 := extensionEvent.Records

	if var_8039bf614513 != nil {
		var var_8039bf614513_mapped *structpb.Value

		var var_8039bf614513_l []*structpb.Value
		for _, value := range var_8039bf614513 {

			var_ba05dfb22728 := value
			var var_ba05dfb22728_mapped *structpb.Value

			var_ba05dfb22728_mapped = structpb.NewStructValue(&structpb.Struct{Fields: RecordMapperInstance.ToProperties(var_ba05dfb22728)})

			var_8039bf614513_l = append(var_8039bf614513_l, var_ba05dfb22728_mapped)
		}
		var_8039bf614513_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_8039bf614513_l})
		properties["records"] = var_8039bf614513_mapped
	}

	var_520c80ae60b0 := extensionEvent.Ids

	if var_520c80ae60b0 != nil {
		var var_520c80ae60b0_mapped *structpb.Value

		var var_520c80ae60b0_l []*structpb.Value
		for _, value := range var_520c80ae60b0 {

			var_d17152c0c04f := value
			var var_d17152c0c04f_mapped *structpb.Value

			var var_d17152c0c04f_err error
			var_d17152c0c04f_mapped, var_d17152c0c04f_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_d17152c0c04f)
			if var_d17152c0c04f_err != nil {
				panic(var_d17152c0c04f_err)
			}

			var_520c80ae60b0_l = append(var_520c80ae60b0_l, var_d17152c0c04f_mapped)
		}
		var_520c80ae60b0_mapped = structpb.NewListValue(&structpb.ListValue{Values: var_520c80ae60b0_l})
		properties["ids"] = var_520c80ae60b0_mapped
	}

	var_99a98f4fc0ce := extensionEvent.Finalizes

	if var_99a98f4fc0ce != nil {
		var var_99a98f4fc0ce_mapped *structpb.Value

		var var_99a98f4fc0ce_err error
		var_99a98f4fc0ce_mapped, var_99a98f4fc0ce_err = types.ByResourcePropertyType(model.ResourceProperty_BOOL).Pack(*var_99a98f4fc0ce)
		if var_99a98f4fc0ce_err != nil {
			panic(var_99a98f4fc0ce_err)
		}
		properties["finalizes"] = var_99a98f4fc0ce_mapped
	}

	var_29f24dbc9e8e := extensionEvent.Sync

	if var_29f24dbc9e8e != nil {
		var var_29f24dbc9e8e_mapped *structpb.Value

		var var_29f24dbc9e8e_err error
		var_29f24dbc9e8e_mapped, var_29f24dbc9e8e_err = types.ByResourcePropertyType(model.ResourceProperty_BOOL).Pack(*var_29f24dbc9e8e)
		if var_29f24dbc9e8e_err != nil {
			panic(var_29f24dbc9e8e_err)
		}
		properties["sync"] = var_29f24dbc9e8e_mapped
	}

	var_becd4a336a8a := extensionEvent.Time

	if var_becd4a336a8a != nil {
		var var_becd4a336a8a_mapped *structpb.Value

		var var_becd4a336a8a_err error
		var_becd4a336a8a_mapped, var_becd4a336a8a_err = types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).Pack(*var_becd4a336a8a)
		if var_becd4a336a8a_err != nil {
			panic(var_becd4a336a8a_err)
		}
		properties["time"] = var_becd4a336a8a_mapped
	}

	var_127a57a5c836 := extensionEvent.Annotations

	if var_127a57a5c836 != nil {
		var var_127a57a5c836_mapped *structpb.Value

		var var_127a57a5c836_st *structpb.Struct = new(structpb.Struct)
		var_127a57a5c836_st.Fields = make(map[string]*structpb.Value)
		for key, value := range var_127a57a5c836 {

			var_51f2e96b65ef := value
			var var_51f2e96b65ef_mapped *structpb.Value

			var var_51f2e96b65ef_err error
			var_51f2e96b65ef_mapped, var_51f2e96b65ef_err = types.ByResourcePropertyType(model.ResourceProperty_STRING).Pack(var_51f2e96b65ef)
			if var_51f2e96b65ef_err != nil {
				panic(var_51f2e96b65ef_err)
			}

			var_127a57a5c836_st.Fields[key] = var_51f2e96b65ef_mapped
		}
		var_127a57a5c836_mapped = structpb.NewStructValue(var_127a57a5c836_st)
		properties["annotations"] = var_127a57a5c836_mapped
	}
	return properties
}

func (m *ExtensionEventMapper) FromProperties(properties map[string]*structpb.Value) *ExtensionEvent {
	var s = m.New()
	if properties["id"] != nil {

		var_5027ac2daff4 := properties["id"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_UUID).UnPack(var_5027ac2daff4)

		if err != nil {
			panic(err)
		}

		var_5027ac2daff4_mapped := new(uuid.UUID)
		*var_5027ac2daff4_mapped = val.(uuid.UUID)

		s.Id = var_5027ac2daff4_mapped
	}
	if properties["action"] != nil {

		var_36888223137f := properties["action"]
		var_36888223137f_mapped := (EventAction)(var_36888223137f.GetStringValue())

		s.Action = var_36888223137f_mapped
	}
	if properties["recordSearchParams"] != nil {

		var_0bec0138781c := properties["recordSearchParams"]
		var mappedValue = ExtensionRecordSearchParamsMapperInstance.FromProperties(var_0bec0138781c.GetStructValue().Fields)

		var_0bec0138781c_mapped := mappedValue

		s.RecordSearchParams = var_0bec0138781c_mapped
	}
	if properties["actionSummary"] != nil {

		var_bb98409df415 := properties["actionSummary"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_bb98409df415)

		if err != nil {
			panic(err)
		}

		var_bb98409df415_mapped := new(string)
		*var_bb98409df415_mapped = val.(string)

		s.ActionSummary = var_bb98409df415_mapped
	}
	if properties["actionDescription"] != nil {

		var_8b3da29b2b72 := properties["actionDescription"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_8b3da29b2b72)

		if err != nil {
			panic(err)
		}

		var_8b3da29b2b72_mapped := new(string)
		*var_8b3da29b2b72_mapped = val.(string)

		s.ActionDescription = var_8b3da29b2b72_mapped
	}
	if properties["resource"] != nil {

		var_6892f84aa461 := properties["resource"]
		var_6892f84aa461_mapped := ResourceMapperInstance.FromProperties(var_6892f84aa461.GetStructValue().Fields)

		s.Resource = var_6892f84aa461_mapped
	}
	if properties["records"] != nil {

		var_a200ccbc3b24 := properties["records"]
		var_a200ccbc3b24_mapped := []*Record{}
		for _, v := range var_a200ccbc3b24.GetListValue().Values {

			var_d9c00bab67b4 := v
			var_d9c00bab67b4_mapped := RecordMapperInstance.FromProperties(var_d9c00bab67b4.GetStructValue().Fields)

			var_a200ccbc3b24_mapped = append(var_a200ccbc3b24_mapped, var_d9c00bab67b4_mapped)
		}

		s.Records = var_a200ccbc3b24_mapped
	}
	if properties["ids"] != nil {

		var_a447578bc822 := properties["ids"]
		var_a447578bc822_mapped := []string{}
		for _, v := range var_a447578bc822.GetListValue().Values {

			var_9f4d0c510cee := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_9f4d0c510cee)

			if err != nil {
				panic(err)
			}

			var_9f4d0c510cee_mapped := val.(string)

			var_a447578bc822_mapped = append(var_a447578bc822_mapped, var_9f4d0c510cee_mapped)
		}

		s.Ids = var_a447578bc822_mapped
	}
	if properties["finalizes"] != nil {

		var_0f92891722e5 := properties["finalizes"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_BOOL).UnPack(var_0f92891722e5)

		if err != nil {
			panic(err)
		}

		var_0f92891722e5_mapped := new(bool)
		*var_0f92891722e5_mapped = val.(bool)

		s.Finalizes = var_0f92891722e5_mapped
	}
	if properties["sync"] != nil {

		var_b926ac54bc9d := properties["sync"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_BOOL).UnPack(var_b926ac54bc9d)

		if err != nil {
			panic(err)
		}

		var_b926ac54bc9d_mapped := new(bool)
		*var_b926ac54bc9d_mapped = val.(bool)

		s.Sync = var_b926ac54bc9d_mapped
	}
	if properties["time"] != nil {

		var_59f530155552 := properties["time"]
		val, err := types.ByResourcePropertyType(model.ResourceProperty_TIMESTAMP).UnPack(var_59f530155552)

		if err != nil {
			panic(err)
		}

		var_59f530155552_mapped := new(time.Time)
		*var_59f530155552_mapped = val.(time.Time)

		s.Time = var_59f530155552_mapped
	}
	if properties["annotations"] != nil {

		var_5b693009a037 := properties["annotations"]
		var_5b693009a037_mapped := make(map[string]string)
		for k, v := range var_5b693009a037.GetStructValue().Fields {

			var_5bccf3f66e2a := v
			val, err := types.ByResourcePropertyType(model.ResourceProperty_STRING).UnPack(var_5bccf3f66e2a)

			if err != nil {
				panic(err)
			}

			var_5bccf3f66e2a_mapped := val.(string)

			var_5b693009a037_mapped[k] = var_5bccf3f66e2a_mapped
		}

		s.Annotations = var_5b693009a037_mapped
	}
	return s
}
