package util

import (
	"github.com/apibrew/apibrew/pkg/errors"
	"github.com/apibrew/apibrew/pkg/formats/unstructured"
	"github.com/apibrew/apibrew/pkg/model"
	log "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/types/known/structpb"
)

func PrepareQuery(resource *model.Resource, queryMap map[string]interface{}) (*model.BooleanExpression, error) {
	var filters = make(map[string]*structpb.Value)

	for _, property := range resource.Properties {
		if queryMap[property.Name] != "" && queryMap[property.Name] != nil {
			if !IsFilterableProperty(property.Type) {
				return nil, errors.RecordValidationError.WithDetails("property is not filterable: " + property.Name)
			}

			val, err := unstructured.ToValue(queryMap[property.Name])

			if err != nil {
				log.Error("error converting value: "+property.Name, queryMap[property.Name])
				return nil, errors.RecordValidationError.WithDetails("error converting value: " + property.Name)
			}

			filters[property.Name] = val
		}
	}

	if len(filters) == 0 {
		return nil, nil
	}

	return &model.BooleanExpression{Filters: filters}, nil
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
