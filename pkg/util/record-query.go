package util

import (
	"github.com/apibrew/apibrew/pkg/errors"
	"github.com/apibrew/apibrew/pkg/model"
	"google.golang.org/protobuf/types/known/structpb"
	"strconv"
)

func PrepareQueryFromFilters(resource *model.Resource, filters map[string]string) (*model.BooleanExpression, errors.ServiceError) {
	var criteria []*model.BooleanExpression
	for _, property := range resource.Properties {
		if filters[property.Name] != "" {
			val, err := deStringifyPropertyValue(filters[property.Name], property)
			if err != nil {
				return nil, errors.RecordValidationError.WithDetails(err.Error())
			}

			if val.GetListValue() != nil {
				criteria = append(criteria, QueryInExpression(property.Name, val))
			} else {
				criteria = append(criteria, QueryEqualExpression(property.Name, val))
			}

		}
	}

	var query *model.BooleanExpression

	if len(criteria) > 0 {
		query = &model.BooleanExpression{Expression: &model.BooleanExpression_And{And: &model.CompoundBooleanExpression{Expressions: criteria}}}
	}
	return query, nil
}

func deStringifyPropertyValue(actualValue string, property *model.ResourceProperty) (*structpb.Value, error) {
	parseFloatVal := func() (*structpb.Value, error) {
		floatVal, err := strconv.ParseFloat(actualValue, 64)

		return structpb.NewNumberValue(floatVal), err
	}

	switch property.Type {
	case model.ResourceProperty_INT32:
		return parseFloatVal()
	case model.ResourceProperty_INT64:
		return parseFloatVal()
	case model.ResourceProperty_FLOAT32:
		return parseFloatVal()
	case model.ResourceProperty_FLOAT64:
		return parseFloatVal()
	case model.ResourceProperty_STRING:
		return structpb.NewStringValue(actualValue), nil
	case model.ResourceProperty_TIME:
		return structpb.NewStringValue(actualValue), nil
	case model.ResourceProperty_TIMESTAMP:
		return structpb.NewStringValue(actualValue), nil
	case model.ResourceProperty_DATE:
		return structpb.NewStringValue(actualValue), nil
	case model.ResourceProperty_BYTES:
		return structpb.NewStringValue(actualValue), nil
	case model.ResourceProperty_UUID:
		return structpb.NewStringValue(actualValue), nil
	case model.ResourceProperty_OBJECT:
		return nil, errors.UnsupportedOperation.WithDetails("Object de-serialization from string is not supported")
	case model.ResourceProperty_STRUCT:
		return nil, errors.UnsupportedOperation.WithDetails("Struct de-serialization from string is not supported")
	case model.ResourceProperty_MAP:
		return nil, errors.UnsupportedOperation.WithDetails("Map de-serialization from string is not supported")
	case model.ResourceProperty_LIST:
		return nil, errors.UnsupportedOperation.WithDetails("List de-serialization from string is not supported")
	case model.ResourceProperty_ENUM:
		return structpb.NewStringValue(actualValue), nil
	case model.ResourceProperty_REFERENCE:
		return structpb.NewStringValue(actualValue), nil
	case model.ResourceProperty_BOOL:
		boolVal, err := strconv.ParseBool(actualValue)

		return structpb.NewBoolValue(boolVal), err
	default:
		return structpb.NewStringValue(actualValue), nil
	}
}

func PrepareQuery(resource *model.Resource, queryMap map[string]string) (*model.BooleanExpression, errors.ServiceError) {
	var criteria []*model.BooleanExpression
	for _, property := range resource.Properties {
		if queryMap[property.Name] != "" {
			var val *structpb.Value
			val, err := structpb.NewValue(queryMap[property.Name])
			if err != nil {
				return nil, errors.RecordValidationError.WithDetails(err.Error())
			}

			if val.GetListValue() != nil {
				criteria = append(criteria, QueryInExpression(property.Name, val))
			} else {
				criteria = append(criteria, QueryEqualExpression(property.Name, val))
			}
		}
	}

	var query *model.BooleanExpression

	if len(criteria) > 0 {
		query = &model.BooleanExpression{Expression: &model.BooleanExpression_And{And: &model.CompoundBooleanExpression{Expressions: criteria}}}
	}
	return query, nil
}

func QueryEqualExpression(propertyName string, val *structpb.Value) *model.BooleanExpression {
	return &model.BooleanExpression{
		Expression: &model.BooleanExpression_Equal{
			Equal: &model.PairExpression{
				Left: &model.Expression{
					Expression: &model.Expression_Property{
						Property: propertyName,
					},
				},
				Right: &model.Expression{
					Expression: &model.Expression_Value{
						Value: val,
					},
				},
			},
		},
	}
}

func QueryAndExpression(left *model.BooleanExpression, right *model.BooleanExpression) *model.BooleanExpression {
	return &model.BooleanExpression{
		Expression: &model.BooleanExpression_And{
			And: &model.CompoundBooleanExpression{
				Expressions: []*model.BooleanExpression{left, right},
			},
		},
	}
}

func QueryOrExpression(left *model.BooleanExpression, right *model.BooleanExpression) *model.BooleanExpression {
	return &model.BooleanExpression{
		Expression: &model.BooleanExpression_Or{
			Or: &model.CompoundBooleanExpression{
				Expressions: []*model.BooleanExpression{left, right},
			},
		},
	}
}

func QueryInExpression(propertyName string, val *structpb.Value) *model.BooleanExpression {
	return &model.BooleanExpression{
		Expression: &model.BooleanExpression_In{
			In: &model.PairExpression{
				Left: &model.Expression{
					Expression: &model.Expression_Property{
						Property: propertyName,
					},
				},
				Right: &model.Expression{
					Expression: &model.Expression_Value{
						Value: val,
					},
				},
			},
		},
	}
}

func RecordIdentifierQuery(resource *model.Resource, properties map[string]*structpb.Value) (*model.BooleanExpression, error) {
	identifiableProperties, err := RecordIdentifierProperties(resource, properties)

	if err != nil {
		return nil, err
	}

	var criteria []*model.BooleanExpression
	for key, value := range identifiableProperties {
		criteria = append(criteria, QueryEqualExpression(key, value))
	}

	var query *model.BooleanExpression

	if len(criteria) > 0 {
		query = &model.BooleanExpression{Expression: &model.BooleanExpression_And{And: &model.CompoundBooleanExpression{Expressions: criteria}}}
	}
	return query, nil

}
