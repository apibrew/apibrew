package dynamodb

import (
	"errors"
	"fmt"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/util"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"strings"
)

func (d *dynamoDbBackend) prepareQuery(query *model.BooleanExpression) (string, map[string]types.AttributeValue, error) {
	if query.GetAnd() != nil {
		var parts []string
		var allValues = make(map[string]types.AttributeValue)

		for _, expr := range query.GetAnd().GetExpressions() {
			filter, values, err := d.prepareQuery(expr)
			if err != nil {
				return "", nil, err
			}
			parts = append(parts, filter)
			allValues = util.MapMerge(allValues, values)
		}

		return fmt.Sprintf("(%s)", strings.Join(parts, " AND ")), allValues, nil
	} else if query.GetOr() != nil {
		var parts []string
		var allValues = make(map[string]types.AttributeValue)

		for _, expr := range query.GetOr().GetExpressions() {
			filter, values, err := d.prepareQuery(expr)
			if err != nil {
				return "", nil, err
			}
			parts = append(parts, filter)
			allValues = util.MapMerge(allValues, values)
		}

		return fmt.Sprintf("(%s)", strings.Join(parts, " OR ")), allValues, nil
	} else if query.GetNot() != nil {
		filter, values, err := d.prepareQuery(query.GetNot())
		if err != nil {
			return "", nil, err
		}

		return fmt.Sprintf("NOT (%s)", filter), values, nil
	} else if query.GetEqual() != nil {
		left, right, values := d.prepareQueryPair(query.GetEqual())

		return fmt.Sprintf("%s = %s", left, right), values, nil
	} else if query.GetGreaterThanOrEqual() != nil {
		left, right, values := d.prepareQueryPair(query.GetGreaterThanOrEqual())

		return fmt.Sprintf("%s >= %s", left, right), values, nil
	} else if query.GetGreaterThan() != nil {
		left, right, values := d.prepareQueryPair(query.GetGreaterThan())

		return fmt.Sprintf("%s > %s", left, right), values, nil
	} else if query.GetLessThanOrEqual() != nil {
		left, right, values := d.prepareQueryPair(query.GetLessThanOrEqual())

		return fmt.Sprintf("%s <= %s", left, right), values, nil
	} else if query.GetLessThan() != nil {
		left, right, values := d.prepareQueryPair(query.GetLessThan())

		return fmt.Sprintf("%s < %s", left, right), values, nil
	} else if query.GetIn() != nil {
		left, right, values := d.prepareQueryPair(query.GetLessThan())

		return fmt.Sprintf("%s IN (%s)", left, right), values, nil
	} else if query.GetLike() != nil {
		left, right, values := d.prepareQueryPair(query.GetLike())

		return fmt.Sprintf("contains(%s, %s)", left, right), values, nil
	} else {
		return "", make(map[string]types.AttributeValue), errors.New("unsupported query type")
	}

	return "", make(map[string]types.AttributeValue), nil
}

func (d *dynamoDbBackend) prepareQueryPair(pair *model.PairExpression) (string, string, map[string]types.AttributeValue) {
	left, values := d.prepareQueryExpression(pair.GetLeft())
	right, rightValues := d.prepareQueryExpression(pair.GetRight())

	return left, right, util.MapMerge(values, rightValues)
}

func (d *dynamoDbBackend) prepareQueryExpression(exp *model.Expression) (string, map[string]types.AttributeValue) {
	var key = util.RandomHex(6)

	switch exp.GetExpression().(type) {
	case *model.Expression_Property:
		return exp.GetProperty(), make(map[string]types.AttributeValue)
	case *model.Expression_Value:
		return fmt.Sprintf(":%s", key), map[string]types.AttributeValue{
			fmt.Sprintf(":%s", key): d.convertStructToAttributeValue(exp.GetValue()),
		}
	default:
		panic("unsupported expression type")
	}
}
