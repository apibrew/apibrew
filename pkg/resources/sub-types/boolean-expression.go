package sub_types

import (
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/service/annotations"
	"github.com/apibrew/apibrew/pkg/util"
)

var PairExpression = &model.ResourceSubType{
	Name: "PairExpression",
	Annotations: map[string]string{
		annotations.CommonType: annotations.Enabled,
	},
	Properties: map[string]*model.ResourceProperty{
		"left": {
			Type:    model.ResourceProperty_STRUCT,
			TypeRef: util.Pointer(Expression.Name),
		},
		"right": {
			Type:    model.ResourceProperty_STRUCT,
			TypeRef: util.Pointer(Expression.Name),
		},
	},
}

var RegexMatchExpression = &model.ResourceSubType{
	Name: "RegexMatchExpression",
	Annotations: map[string]string{
		annotations.CommonType: annotations.Enabled,
	},
	Properties: map[string]*model.ResourceProperty{
		"pattern": {
			Type: model.ResourceProperty_STRING,
		},
		"expression": {
			Type:    model.ResourceProperty_STRUCT,
			TypeRef: util.Pointer(Expression.Name),
		},
	},
}

var Expression = &model.ResourceSubType{
	Name: "Expression",
	Annotations: map[string]string{
		annotations.CommonType: annotations.Enabled,
	},
	Properties: map[string]*model.ResourceProperty{
		"property": {
			Type: model.ResourceProperty_STRING,
		},
		"value": {
			Type: model.ResourceProperty_OBJECT,
		},
	},
}

var BooleanExpression = &model.ResourceSubType{
	Name: "BooleanExpression",
	Annotations: map[string]string{
		annotations.CommonType: annotations.Enabled,
	},
	Properties: map[string]*model.ResourceProperty{
		"and": {
			Type: model.ResourceProperty_LIST,
			Item: &model.ResourceProperty{
				Type:    model.ResourceProperty_STRUCT,
				TypeRef: util.Pointer("BooleanExpression"),
			},
		},
		"or": {
			Type: model.ResourceProperty_LIST,
			Item: &model.ResourceProperty{
				Type:    model.ResourceProperty_STRUCT,
				TypeRef: util.Pointer("BooleanExpression"),
			},
		},
		"not": {
			Type:    model.ResourceProperty_STRUCT,
			TypeRef: util.Pointer("BooleanExpression"),
		},
		"equal": {
			Type:    model.ResourceProperty_STRUCT,
			TypeRef: util.Pointer(PairExpression.Name),
		},
		"lessThan": {
			Type:    model.ResourceProperty_STRUCT,
			TypeRef: util.Pointer(PairExpression.Name),
		},
		"greaterThan": {
			Type:    model.ResourceProperty_STRUCT,
			TypeRef: util.Pointer(PairExpression.Name),
		},
		"lessThanOrEqual": {
			Type:    model.ResourceProperty_STRUCT,
			TypeRef: util.Pointer(PairExpression.Name),
		},
		"greaterThanOrEqual": {
			Type:    model.ResourceProperty_STRUCT,
			TypeRef: util.Pointer(PairExpression.Name),
		},
		"in": {
			Type:    model.ResourceProperty_STRUCT,
			TypeRef: util.Pointer(PairExpression.Name),
		},
		"isNull": {
			Type:    model.ResourceProperty_STRUCT,
			TypeRef: util.Pointer(Expression.Name),
		},
		"regexMatch": {
			Type:    model.ResourceProperty_STRUCT,
			TypeRef: util.Pointer(RegexMatchExpression.Name),
		},
	},
}
