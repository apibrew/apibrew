package helper

import (
	"github.com/apibrew/apibrew/pkg/model"
)

type property[GoType any, QueryBuilderType PropertyTypeSpecificQueryBuilder[GoType]] struct {
	name         string
	typ          model.ResourceProperty_Type
	queryBuilder QueryBuilderType
}

func (p property[GoType, QueryBuilderType]) GetName() string {
	return p.name
}

func (p property[GoType, QueryBuilderType]) GetType() model.ResourceProperty_Type {
	return p.typ
}

func (p property[GoType, QueryBuilderType]) Query() QueryBuilderType {
	return p.queryBuilder
}

type Property[GoType any, QueryBuilderType PropertyTypeSpecificQueryBuilder[GoType]] interface {
	GetName() string
	GetType() model.ResourceProperty_Type
	Query() QueryBuilderType
}

func DefineProperty[GoType any, QueryBuilderType PropertyTypeSpecificQueryBuilder[GoType]](name string, typ model.ResourceProperty_Type, queryBuilderType QueryBuilderType) Property[GoType, QueryBuilderType] {
	return &property[GoType, QueryBuilderType]{
		name:         name,
		typ:          typ,
		queryBuilder: queryBuilderType,
	}
}
