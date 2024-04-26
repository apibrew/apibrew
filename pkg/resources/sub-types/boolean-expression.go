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
	Properties: []*model.ResourceProperty{
		{
			Name:    "left",
			Type:    model.ResourceProperty_STRUCT,
			TypeRef: util.Pointer(Expression.Name),
		},
		{
			Name:    "right",
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
	Properties: []*model.ResourceProperty{
		{
			Name: "pattern",
			Type: model.ResourceProperty_STRING,
		},
		{
			Name:    "expression",
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
	Properties: []*model.ResourceProperty{
		{
			Name: "property",
			Type: model.ResourceProperty_STRING,
		},
		{
			Name: "value",
			Type: model.ResourceProperty_OBJECT,
		},
	},
}

var BooleanExpression = &model.ResourceSubType{
	Name: "BooleanExpression",
	Annotations: map[string]string{
		annotations.CommonType: annotations.Enabled,
	},
	Properties: []*model.ResourceProperty{
		{
			Name: "and",
			Type: model.ResourceProperty_LIST,
			Item: &model.ResourceProperty{
				Type:    model.ResourceProperty_STRUCT,
				TypeRef: util.Pointer("BooleanExpression"),
			},
		},
		{
			Name: "or",
			Type: model.ResourceProperty_LIST,
			Item: &model.ResourceProperty{
				Type:    model.ResourceProperty_STRUCT,
				TypeRef: util.Pointer("BooleanExpression"),
			},
		},
		{
			Name:    "not",
			Type:    model.ResourceProperty_STRUCT,
			TypeRef: util.Pointer("BooleanExpression"),
		},
		{
			Name:    "equal",
			Type:    model.ResourceProperty_STRUCT,
			TypeRef: util.Pointer(PairExpression.Name),
		},
		{
			Name:    "lessThan",
			Type:    model.ResourceProperty_STRUCT,
			TypeRef: util.Pointer(PairExpression.Name),
		},
		{
			Name:    "greaterThan",
			Type:    model.ResourceProperty_STRUCT,
			TypeRef: util.Pointer(PairExpression.Name),
		},
		{
			Name:    "lessThanOrEqual",
			Type:    model.ResourceProperty_STRUCT,
			TypeRef: util.Pointer(PairExpression.Name),
		},
		{
			Name:    "greaterThanOrEqual",
			Type:    model.ResourceProperty_STRUCT,
			TypeRef: util.Pointer(PairExpression.Name),
		},
		{
			Name:    "in",
			Type:    model.ResourceProperty_STRUCT,
			TypeRef: util.Pointer(PairExpression.Name),
		},
		{
			Name:    "isNull",
			Type:    model.ResourceProperty_STRUCT,
			TypeRef: util.Pointer(Expression.Name),
		},
		{
			Name: "filters",
			Type: model.ResourceProperty_MAP,
			Item: &model.ResourceProperty{Type: model.ResourceProperty_OBJECT},
		},
		{
			Name:    "regexMatch",
			Type:    model.ResourceProperty_STRUCT,
			TypeRef: util.Pointer(RegexMatchExpression.Name),
		},
	},
}
