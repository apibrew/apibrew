package helper

import (
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/structpb"
	"time"
)

type QueryBuilder struct {
}

func (q QueryBuilder) And(list ...*model.BooleanExpression) *model.BooleanExpression {
	return &model.BooleanExpression{Expression: &model.BooleanExpression_And{
		And: &model.CompoundBooleanExpression{
			Expressions: list,
		},
	}}
}

func (q QueryBuilder) Or(list ...*model.BooleanExpression) *model.BooleanExpression {
	return &model.BooleanExpression{Expression: &model.BooleanExpression_Or{
		Or: &model.CompoundBooleanExpression{
			Expressions: list,
		},
	}}
}

func (q QueryBuilder) Not(condition *model.BooleanExpression) *model.BooleanExpression {
	return &model.BooleanExpression{Expression: &model.BooleanExpression_Not{
		Not: condition,
	}}
}

func (q QueryBuilder) In(property string, values []interface{}) *model.BooleanExpression {
	lval, err := structpb.NewList(values)

	if err != nil {
		panic(err)
	}

	return &model.BooleanExpression{Expression: &model.BooleanExpression_In{
		In: &model.PairExpression{
			Left: &model.Expression{
				Expression: &model.Expression_Property{Property: property},
			},
			Right: &model.Expression{
				Expression: &model.Expression_Value{Value: structpb.NewListValue(lval)},
			},
		},
	}}
}

func (q QueryBuilder) Equal(propertyName string, value *structpb.Value) *model.BooleanExpression {
	return &model.BooleanExpression{
		Expression: &model.BooleanExpression_Equal{
			Equal: &model.PairExpression{
				Left: &model.Expression{
					Expression: &model.Expression_Property{Property: propertyName},
				},
				Right: &model.Expression{
					Expression: &model.Expression_Value{Value: value},
				},
			},
		},
	}
}

func (q QueryBuilder) FromProperties(props map[string]*structpb.Value) *model.BooleanExpression {
	var list []*model.BooleanExpression

	for propName, value := range props {
		list = append(list, q.Equal(propName, value))
	}

	return q.And(list...)
}

func NewQueryBuilder() QueryBuilder {
	return QueryBuilder{}
}

type PropertyTypeSpecificQueryBuilder[GoType any] interface {
	Equals(val GoType) *model.BooleanExpression
}

type StringQueryBuilder struct {
	PropName string
}

func (sq StringQueryBuilder) Equals(val string) *model.BooleanExpression {
	return nil
}

func (sq StringQueryBuilder) Contains(val string) *model.BooleanExpression {
	return nil
}

type ListQueryBuilder struct {
	PropName string
}

func (sq ListQueryBuilder) Equals(val []interface{}) *model.BooleanExpression {
	return nil
}

func (sq ListQueryBuilder) Contains(val []interface{}) *model.BooleanExpression {
	return nil
}

type MapQueryBuilder struct {
	PropName string
}

func (sq MapQueryBuilder) Equals(val map[string]interface{}) *model.BooleanExpression {
	return nil
}

type StructQueryBuilder struct {
	PropName string
}

func (sq MapQueryBuilder) Contains(val map[string]interface{}) *model.BooleanExpression {
	return nil
}

func (sq StructQueryBuilder) Equals(val map[string]interface{}) *model.BooleanExpression {
	return nil
}

func (sq StructQueryBuilder) Contains(val map[string]interface{}) *model.BooleanExpression {
	return nil
}

type EnumQueryBuilder struct {
	PropName string
}

func (sq EnumQueryBuilder) Equals(val string) *model.BooleanExpression {
	return nil
}

//type ReferenceQueryBuilder[RefType abs.Entity[RefType]] struct {
//	PropName string
//}
//
//func (r ReferenceQueryBuilder[RefType]) Equals(val RefType) *model.BooleanExpression {
//	properties := val.ToProperties()
//
//	return &model.BooleanExpression{
//		Expression: &model.BooleanExpression_Equal{
//			Equal: &model.PairExpression{
//				Left: &model.Expression{
//					Expression: &model.Expression_Property{Property: r.PropName},
//				},
//				Right: &model.Expression{
//					Expression: &model.Expression_RefValue{RefValue: &model.RefValue{
//						Namespace:  val.GetNamespace(),
//						Resource:   val.GetResourceName(),
//						Properties: properties,
//					}},
//				},
//			},
//		},
//	}
//}

type UuidQueryBuilder struct {
	PropName string
}

func (u UuidQueryBuilder) Equals(val uuid.UUID) *model.BooleanExpression {
	//TODO implement me
	panic("implement me")
}

type Int32QueryBuilder struct {
	PropName string
}

func (i Int32QueryBuilder) Equals(val int32) *model.BooleanExpression {
	//TODO implement me
	panic("implement me")
}

type Int64QueryBuilder struct {
	PropName string
}

func (i Int64QueryBuilder) Equals(val int64) *model.BooleanExpression {
	//TODO implement me
	panic("implement me")
}

type DateQueryBuilder struct {
	PropName string
}

func (i DateQueryBuilder) Equals(val time.Time) *model.BooleanExpression {
	//TODO implement me
	panic("implement me")
}

type TimeQueryBuilder struct {
	PropName string
}

func (i TimeQueryBuilder) Equals(val time.Time) *model.BooleanExpression {
	//TODO implement me
	panic("implement me")
}

type TimestampQueryBuilder struct {
	PropName string
}

func (i TimestampQueryBuilder) Equals(val time.Time) *model.BooleanExpression {
	//TODO implement me
	panic("implement me")
}

type BoolQueryBuilder struct {
	PropName string
}

func (i BoolQueryBuilder) Equals(val bool) *model.BooleanExpression {
	//TODO implement me
	panic("implement me")
}

type ObjectQueryBuilder struct {
	PropName string
}

func (i ObjectQueryBuilder) Equals(val map[string]interface{}) *model.BooleanExpression {
	//TODO implement me
	panic("implement me")
}

type Float32QueryBuilder struct {
	PropName string
}

func (i Float32QueryBuilder) Equals(val float32) *model.BooleanExpression {
	//TODO implement me
	panic("implement me")
}

type Float64QueryBuilder struct {
	PropName string
}

func (i Float64QueryBuilder) Equals(val float64) *model.BooleanExpression {
	//TODO implement me
	panic("implement me")
}

type BytesQueryBuilder struct {
	PropName string
}

func (i BytesQueryBuilder) Equals(val []byte) *model.BooleanExpression {
	//TODO implement me
	panic("implement me")
}
