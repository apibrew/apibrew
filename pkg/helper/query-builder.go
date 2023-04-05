package helper

import (
	"github.com/google/uuid"
	"github.com/tislib/data-handler/pkg/abs"
	"github.com/tislib/data-handler/pkg/model"
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

func (q QueryBuilder) Equal(property string, value *structpb.Value) *model.BooleanExpression {
	return &model.BooleanExpression{
		Expression: &model.BooleanExpression_Equal{
			Equal: &model.PairExpression{
				Left: &model.Expression{
					Expression: &model.Expression_Property{Property: property},
				},
				Right: &model.Expression{
					Expression: &model.Expression_Value{Value: value},
				},
			},
		},
	}
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

type ReferenceQueryBuilder[RefType abs.Entity[RefType]] struct {
	PropName string
}

func (r ReferenceQueryBuilder[RefType]) Equals(val RefType) *model.BooleanExpression {
	qb := NewQueryBuilder()
	return qb.Equal(r.PropName, structpb.NewStructValue(&structpb.Struct{Fields: val.ToProperties()}))
}

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
