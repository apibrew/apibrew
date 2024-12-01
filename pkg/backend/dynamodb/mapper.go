package dynamodb

import (
	"fmt"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/util"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"google.golang.org/protobuf/types/known/structpb"
	"strconv"
)

func (d *dynamoDbBackend) recordToAttributeMap(resource *model.Resource, record *model.Record) map[string]types.AttributeValue {
	attributeMap := make(map[string]types.AttributeValue)

	for _, field := range resource.Properties {
		if record.Properties[field.Name] == nil {
			continue
		}

		attributeMap[field.Name] = d.recordPropertyValueToAttributeValue(field, record.Properties[field.Name])
	}

	attributeMap["PK"] = d.getPKForResource(resource)
	attributeMap["SK"] = &types.AttributeValueMemberS{Value: util.GetRecordId(record)}

	return attributeMap
}

func (d *dynamoDbBackend) recordPropertyValueToAttributeValue(field *model.ResourceProperty, value *structpb.Value) types.AttributeValue {
	switch field.Type {
	case model.ResourceProperty_STRING, model.ResourceProperty_ENUM, model.ResourceProperty_UUID, model.ResourceProperty_DATE, model.ResourceProperty_TIME, model.ResourceProperty_TIMESTAMP:
		return &types.AttributeValueMemberS{Value: value.GetStringValue()}
	case model.ResourceProperty_FLOAT32, model.ResourceProperty_FLOAT64, model.ResourceProperty_INT32, model.ResourceProperty_INT64:
		return &types.AttributeValueMemberN{Value: fmt.Sprintf("%f", value.GetNumberValue())}
	case model.ResourceProperty_BOOL:
		return &types.AttributeValueMemberBOOL{Value: value.GetBoolValue()}
	case model.ResourceProperty_BYTES:
		return &types.AttributeValueMemberS{Value: value.GetStringValue()}
	case model.ResourceProperty_LIST, model.ResourceProperty_OBJECT, model.ResourceProperty_MAP, model.ResourceProperty_REFERENCE, model.ResourceProperty_STRUCT:
		return d.convertStructToAttributeValue(value)
	default:
		panic("unknown type: " + field.Type.String())
	}
}

func (d *dynamoDbBackend) convertStructToAttributeValue(value *structpb.Value) types.AttributeValue {
	if value == nil {
		return &types.AttributeValueMemberNULL{Value: true}
	}
	switch value.Kind.(type) {
	case *structpb.Value_NullValue:
		return &types.AttributeValueMemberNULL{Value: true}
	case *structpb.Value_StringValue:
		return &types.AttributeValueMemberS{Value: value.GetStringValue()}
	case *structpb.Value_NumberValue:
		return &types.AttributeValueMemberN{Value: fmt.Sprintf("%f", value.GetNumberValue())}
	case *structpb.Value_BoolValue:
		return &types.AttributeValueMemberBOOL{Value: value.GetBoolValue()}
	case *structpb.Value_StructValue:
		result := make(map[string]types.AttributeValue)

		for key, val := range value.GetStructValue().Fields {
			result[key] = d.convertStructToAttributeValue(val)
		}

		return &types.AttributeValueMemberM{Value: result}
	case *structpb.Value_ListValue:
		result := make([]types.AttributeValue, 0, len(value.GetListValue().Values))

		for _, val := range value.GetListValue().Values {
			result = append(result, d.convertStructToAttributeValue(val))
		}

		return &types.AttributeValueMemberL{Value: result}
	default:
		panic("unknown type")
	}
}

func (d *dynamoDbBackend) convertAttributeMapToRecord(resource *model.Resource, item map[string]types.AttributeValue) (*model.Record, error) {
	var record = &model.Record{}

	delete(item, "PK")
	delete(item, "SK")

	record.Properties = d.convertAttributeMapToPropertyMap(item)

	return record, nil
}

func (d *dynamoDbBackend) convertAttributeMapToPropertyMap(item map[string]types.AttributeValue) map[string]*structpb.Value {
	result := make(map[string]*structpb.Value)

	for key, val := range item {
		result[key] = d.convertAttributeValueToValue(val)
	}

	return result
}

func (d *dynamoDbBackend) convertAttributeValueToValue(val types.AttributeValue) *structpb.Value {
	switch val.(type) {
	case *types.AttributeValueMemberS:
		return &structpb.Value{Kind: &structpb.Value_StringValue{StringValue: val.(*types.AttributeValueMemberS).Value}}
	case *types.AttributeValueMemberN:
		var floatValue, err = strconv.ParseFloat(val.(*types.AttributeValueMemberN).Value, 64)

		if err != nil {
			panic(err)
		}
		return &structpb.Value{Kind: &structpb.Value_NumberValue{NumberValue: floatValue}}
	case *types.AttributeValueMemberBOOL:
		return &structpb.Value{Kind: &structpb.Value_BoolValue{BoolValue: val.(*types.AttributeValueMemberBOOL).Value}}
	case *types.AttributeValueMemberM:
		result := make(map[string]*structpb.Value)

		for key, val := range val.(*types.AttributeValueMemberM).Value {
			result[key] = d.convertAttributeValueToValue(val)
		}

		return &structpb.Value{Kind: &structpb.Value_StructValue{StructValue: &structpb.Struct{Fields: result}}}
	case *types.AttributeValueMemberL:
		result := make([]*structpb.Value, 0, len(val.(*types.AttributeValueMemberL).Value))

		for _, val := range val.(*types.AttributeValueMemberL).Value {
			result = append(result, d.convertAttributeValueToValue(val))
		}

		return &structpb.Value{Kind: &structpb.Value_ListValue{ListValue: &structpb.ListValue{Values: result}}}
	case *types.AttributeValueMemberNULL:
		return &structpb.Value{Kind: &structpb.Value_NullValue{NullValue: structpb.NullValue_NULL_VALUE}}
	default:
		panic("unknown type")
	}
}
