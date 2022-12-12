package service

import (
	"data-handler/model"
	"data-handler/service/errors"
	"google.golang.org/protobuf/types/known/structpb"
)

func (r *recordService) PrepareQuery(resource *model.Resource, queryMap map[string]interface{}) (*model.BooleanExpression, errors.ServiceError) {
	var criteria []*model.BooleanExpression
	for _, property := range resource.Properties {
		if queryMap[property.Name] != nil {
			var val *structpb.Value
			val, err := structpb.NewValue(queryMap[property.Name])
			if err != nil {
				return nil, errors.RecordValidationError.WithDetails(err.Error())
			}
			criteria = append(criteria, r.newEqualExpression(property.Name, val))
		}
	}

	var additionalProperties = []string{
		"id", "version",
	}

	for _, property := range additionalProperties {
		if queryMap[property] != nil {
			var val *structpb.Value
			val, err := structpb.NewValue(queryMap[property])
			if err != nil {
				return nil, errors.RecordValidationError.WithDetails(err.Error())
			}
			criteria = append(criteria, r.newEqualExpression(property, val))
		}
	}

	var query *model.BooleanExpression

	if len(criteria) > 0 {
		query = &model.BooleanExpression{Expression: &model.BooleanExpression_And{And: &model.CompoundBooleanExpression{Expressions: criteria}}}
	}
	return query, nil
}

func (r *recordService) newEqualExpression(propertyName string, val *structpb.Value) *model.BooleanExpression {
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
