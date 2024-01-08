package api

import (
	"context"
	"github.com/apibrew/apibrew/pkg/errors"
	"github.com/apibrew/apibrew/pkg/formats/unstructured"
	"github.com/apibrew/apibrew/pkg/resource_model"
)

// Interface
// Api interface is a facade for all api services
// /*
type Interface interface {
	Create(ctx context.Context, record unstructured.Unstructured) (unstructured.Unstructured, errors.ServiceError)
	Update(ctx context.Context, record unstructured.Unstructured) (unstructured.Unstructured, errors.ServiceError)
	Apply(ctx context.Context, record unstructured.Unstructured) (unstructured.Unstructured, errors.ServiceError)
	Save(ctx context.Context, saveMode SaveMode, record unstructured.Unstructured) (unstructured.Unstructured, errors.ServiceError)
	Load(ctx context.Context, record unstructured.Unstructured) (unstructured.Unstructured, errors.ServiceError)
	Delete(ctx context.Context, record unstructured.Unstructured) errors.ServiceError
	List(ctx context.Context, params ListParams) ([]unstructured.Unstructured, uint32, errors.ServiceError)
}

type ListParams struct {
	Query             resource_model.BooleanExpression `json:"query,omitempty"`
	Type              string                           `type:"and,omitempty"`
	Limit             uint32                           `json:"limit,omitempty"`
	Offset            uint64                           `json:"offset,omitempty"`
	UseHistory        bool                             `json:"useHistory,omitempty"`
	ResolveReferences []string                         `json:"resolveReferences,omitempty"`
	Filters           map[string]string                `json:"filters,omitempty"`
	Aggregation       *Aggregation                     `json:"aggregation,omitempty"`
}

type Aggregation struct {
	Items    []AggregationItem `json:"items,omitempty"`
	Grouping []GroupingItem    `json:"grouping,omitempty"`
}

type AggregationItem struct {
	Name      string               `json:"name,omitempty"`
	Algorithm AggregationAlgorithm `json:"algorithm,omitempty"`
	Property  string               `json:"property,omitempty"`
}

type GroupingItem struct {
	Property string `json:"property,omitempty"`
}

type SaveMode int

const (
	Create SaveMode = iota
	Update
	Apply
)

type AggregationAlgorithm string

const (
	Count AggregationAlgorithm = "COUNT"
	Sum   AggregationAlgorithm = "SUM"
	Avg   AggregationAlgorithm = "AVG"
	Max   AggregationAlgorithm = "MAX"
	Min   AggregationAlgorithm = "MIN"
)
