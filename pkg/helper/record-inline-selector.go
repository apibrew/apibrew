package helper

import (
	"context"
	"github.com/apibrew/apibrew/pkg/errors"
	"github.com/apibrew/apibrew/pkg/formats/unstructured"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/resource_model"
	"github.com/apibrew/apibrew/pkg/types"
	"github.com/apibrew/apibrew/pkg/util"
)

type RecordInlineSelector struct {
}

func (s RecordInlineSelector) SelectRecords(ctx context.Context, resource *model.Resource, records *[]*model.Record, selector *resource_model.BooleanExpression) ([]*model.Record, error) {
	var result []*model.Record

	for _, record := range *records {
		ok, err := s.EvaluateRecord(ctx, resource, record, selector)
		if err != nil {
			return nil, err
		}
		if ok {
			result = append(result, record)
		}
	}

	return result, nil
}

func (s RecordInlineSelector) EvaluateRecord(ctx context.Context, resource *model.Resource, record *model.Record, selector *resource_model.BooleanExpression) (bool, error) {
	if selector.And != nil {
		for _, expression := range selector.And {
			if ok, err := s.EvaluateRecord(ctx, resource, record, &expression); !ok || err != nil {
				return ok, err
			}
		}
	}

	if selector.Or != nil {
		for _, expression := range selector.Or {
			if ok, err := s.EvaluateRecord(ctx, resource, record, &expression); ok || err != nil {
				return ok, err
			}
		}
	}

	if selector.Not != nil {
		ok, err := s.EvaluateRecord(ctx, resource, record, selector.Not)

		if err != nil {
			return false, err
		}

		return !ok, nil
	}

	if selector.Equal != nil {
		left, right, prop, err := s.resolve(resource, record, selector.Equal)

		if err != nil {
			return false, err
		}

		typ := types.ByResourcePropertyType(prop.Type)

		if typ.Equals(left, right) {
			return true, nil
		} else {
			return false, nil
		}
	}

	return false, errors.UnsupportedOperation.WithDetails("Unknown boolean expression")
}

func (s RecordInlineSelector) resolve(resource *model.Resource, record *model.Record, than *resource_model.PairExpression) (unstructured.Any, unstructured.Any, *model.ResourceProperty, error) {
	namedProps := util.GetNamedMap(resource.Properties)

	var left unstructured.Any
	var right unstructured.Any
	var prop *model.ResourceProperty

	if than.Left.Property != nil {
		prop = namedProps[*than.Left.Property]

		if prop == nil {
			return nil, nil, nil, errors.PropertyNotFoundError.WithDetails("Property not found: " + *than.Left.Property)
		}

		left = unstructured.FromValue(record.Properties[*than.Left.Property])
	} else {
		left = than.Left.Value
	}

	if than.Right.Property != nil {
		prop = namedProps[*than.Right.Property]

		if prop == nil {
			return nil, nil, nil, errors.PropertyNotFoundError.WithDetails("Property not found: " + *than.Left.Property)
		}

		right = unstructured.FromValue(record.Properties[*than.Right.Property])
	} else {
		right = than.Right.Value
	}

	return left, right, prop, nil
}
