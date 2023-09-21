package extramappings

import (
	"github.com/apibrew/apibrew/pkg/formats/unstructured"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/resource_model"
	"github.com/apibrew/apibrew/pkg/util"
	"google.golang.org/protobuf/types/known/structpb"
)

func BooleanExpressionFromProto(exp *model.BooleanExpression) resource_model.ExtensionBooleanExpression {
	var result = resource_model.ExtensionBooleanExpression{}

	if exp.Expression != nil {
		if exp.GetAnd() != nil {
			result.And = util.ArrayMap(exp.GetAnd().Expressions, BooleanExpressionFromProto)
		}

		if exp.GetOr() != nil {
			result.Or = util.ArrayMap(exp.GetOr().Expressions, BooleanExpressionFromProto)
		}

		if exp.GetNot() != nil {
			result.Not = util.Pointer(BooleanExpressionFromProto(exp.GetNot()))
		}

		if exp.GetEqual() != nil {
			result.Equal = util.Pointer(PairExpressionFromProto(exp.GetEqual()))
		}

		if exp.GetGreaterThan() != nil {
			result.GreaterThan = util.Pointer(PairExpressionFromProto(exp.GetGreaterThan()))
		}

		if exp.GetLessThan() != nil {
			result.LessThan = util.Pointer(PairExpressionFromProto(exp.GetLessThan()))
		}

		if exp.GetGreaterThanOrEqual() != nil {
			result.GreaterThanOrEqual = util.Pointer(PairExpressionFromProto(exp.GetGreaterThanOrEqual()))
		}

		if exp.GetLessThanOrEqual() != nil {
			result.LessThanOrEqual = util.Pointer(PairExpressionFromProto(exp.GetLessThanOrEqual()))
		}

		if exp.GetRegexMatch() != nil {
			result.RegexMatch = util.Pointer(RegexMatchExpressionFromProto(exp.GetRegexMatch()))
		}
	}

	return result
}

func PairExpressionFromProto(equal *model.PairExpression) resource_model.ExtensionPairExpression {
	var result = resource_model.ExtensionPairExpression{}

	if equal.Left != nil {
		result.Left = util.Pointer(ExpressionFromProto(equal.Left))
	}

	if equal.Right != nil {
		result.Right = util.Pointer(ExpressionFromProto(equal.Right))
	}

	return result
}

func ExpressionFromProto(exp *model.Expression) resource_model.ExtensionExpression {
	var result = resource_model.ExtensionExpression{}

	if exp.GetProperty() != "" {
		result.Property = util.Pointer(exp.GetProperty())
	}

	if exp.GetValue() != nil {
		result.Value = unstructured.FromValue(exp.GetValue())
	}

	if exp.GetRefValue() != nil {
		result.RefValue = util.Pointer(RefValueFromProto(exp.GetRefValue()))
	}

	return result
}

func RefValueFromProto(value *model.RefValue) resource_model.ExtensionRefValue {
	var result = resource_model.ExtensionRefValue{}

	result.Namespace = util.Pointer(value.GetNamespace())
	result.Resource = util.Pointer(value.GetResource())
	if value.GetProperties() != nil {
		result.Properties = make(map[string]interface{})

		for k, v := range value.GetProperties() {
			result.Properties[k] = unstructured.FromValue(v)
		}
	}

	return result
}

func RegexMatchExpressionFromProto(match *model.RegexMatchExpression) resource_model.ExtensionRegexMatchExpression {
	var result = resource_model.ExtensionRegexMatchExpression{}

	result.Pattern = util.Pointer(match.GetPattern())
	result.Expression = util.Pointer(ExpressionFromProto(match.GetExpression()))

	return result
}

func BooleanExpressionToProto(exp resource_model.ExtensionBooleanExpression) *model.BooleanExpression {
	var result = new(model.BooleanExpression)

	if exp.And != nil {
		result.Expression = &model.BooleanExpression_And{
			And: &model.CompoundBooleanExpression{
				Expressions: util.ArrayMap(exp.And, BooleanExpressionToProto),
			},
		}
	}

	if exp.Or != nil {
		result.Expression = &model.BooleanExpression_Or{
			Or: &model.CompoundBooleanExpression{
				Expressions: util.ArrayMap(exp.Or, BooleanExpressionToProto),
			},
		}
	}

	if exp.Not != nil {
		result.Expression = &model.BooleanExpression_Not{
			Not: BooleanExpressionToProto(*exp.Not),
		}
	}

	if exp.Equal != nil {
		result.Expression = &model.BooleanExpression_Equal{
			Equal: PairExpressionToProto(*exp.Equal),
		}
	}

	if exp.GreaterThan != nil {
		result.Expression = &model.BooleanExpression_GreaterThan{
			GreaterThan: PairExpressionToProto(*exp.GreaterThan),
		}
	}

	if exp.LessThan != nil {
		result.Expression = &model.BooleanExpression_LessThan{
			LessThan: PairExpressionToProto(*exp.LessThan),
		}
	}

	if exp.GreaterThanOrEqual != nil {
		result.Expression = &model.BooleanExpression_GreaterThanOrEqual{
			GreaterThanOrEqual: PairExpressionToProto(*exp.GreaterThanOrEqual),
		}
	}

	if exp.LessThanOrEqual != nil {
		result.Expression = &model.BooleanExpression_LessThanOrEqual{
			LessThanOrEqual: PairExpressionToProto(*exp.LessThanOrEqual),
		}
	}

	if exp.RegexMatch != nil {
		result.Expression = &model.BooleanExpression_RegexMatch{
			RegexMatch: RegexMatchExpressionToProto(*exp.RegexMatch),
		}
	}

	return result
}

func RegexMatchExpressionToProto(expression resource_model.ExtensionRegexMatchExpression) *model.RegexMatchExpression {
	var result = new(model.RegexMatchExpression)

	result.Pattern = util.DePointer(expression.Pattern, "")

	if expression.Expression != nil {
		result.Expression = ExpressionToProto(*expression.Expression)
	}

	return result
}

func ExpressionToProto(expression resource_model.ExtensionExpression) *model.Expression {
	var result = new(model.Expression)

	if expression.Property != nil {
		result.Expression = &model.Expression_Property{
			Property: util.DePointer(expression.Property, ""),
		}
	}

	if expression.Value != nil {
		val, err := structpb.NewValue(expression.Value)

		if err != nil {
			panic(err)
		}

		result.Expression = &model.Expression_Value{
			Value: val,
		}
	}

	if expression.RefValue != nil {
		result.Expression = &model.Expression_RefValue{
			RefValue: RefValueToProto(*expression.RefValue),
		}
	}

	return result
}

func RefValueToProto(value resource_model.ExtensionRefValue) *model.RefValue {
	var result = new(model.RefValue)

	result.Namespace = util.DePointer(value.Namespace, "")
	result.Resource = util.DePointer(value.Resource, "")

	if value.Properties != nil {
		result.Properties = make(map[string]*structpb.Value)

		for k, v := range value.Properties {
			val, err := structpb.NewValue(v)

			if err != nil {
				panic(err)
			}

			result.Properties[k] = val
		}
	}

	return result
}

func PairExpressionToProto(expression resource_model.ExtensionPairExpression) *model.PairExpression {
	var result = new(model.PairExpression)

	if expression.Left != nil {
		result.Left = ExpressionToProto(*expression.Left)
	}

	if expression.Right != nil {
		result.Right = ExpressionToProto(*expression.Right)
	}

	return result
}
