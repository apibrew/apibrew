package query

import (
	"github.com/apibrew/apibrew/pkg/resource_model"
)

func WalkBooleanExpressionValues(exp *resource_model.BooleanExpression, walkFn func(val interface{}) interface{}) {
	if exp == nil {
		return
	}

	for _, subExp := range exp.And {
		WalkBooleanExpressionValues(&subExp, walkFn)
	}

	for _, subExp := range exp.Or {
		WalkBooleanExpressionValues(&subExp, walkFn)
	}

	if exp.Not != nil {
		WalkBooleanExpressionValues(exp.Not, walkFn)
	}

	if exp.Equal != nil {
		WalkExpressionValues(exp.Equal.Left, walkFn)
		WalkExpressionValues(exp.Equal.Right, walkFn)
	}

	if exp.GreaterThan != nil {
		WalkExpressionValues(exp.GreaterThan.Left, walkFn)
		WalkExpressionValues(exp.GreaterThan.Right, walkFn)
	}

	if exp.GreaterThanOrEqual != nil {
		WalkExpressionValues(exp.GreaterThanOrEqual.Left, walkFn)
		WalkExpressionValues(exp.GreaterThanOrEqual.Right, walkFn)
	}

	if exp.LessThan != nil {
		WalkExpressionValues(exp.LessThan.Left, walkFn)
		WalkExpressionValues(exp.LessThan.Right, walkFn)
	}

	if exp.LessThanOrEqual != nil {
		WalkExpressionValues(exp.LessThanOrEqual.Left, walkFn)
		WalkExpressionValues(exp.LessThanOrEqual.Right, walkFn)
	}

	if exp.In != nil {
		WalkExpressionValues(exp.In.Left, walkFn)
		WalkExpressionValues(exp.In.Right, walkFn)
	}
}

func WalkExpressionValues(exp *resource_model.Expression, fn func(val interface{}) interface{}) {
	if exp.Value != nil {
		exp.Value = fn(exp.Value)
	}
}
