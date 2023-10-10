package client

import (
	"context"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/resource_model"
)

type ExternalFunction func(ctx context.Context, req *model.Event) (*model.Event, error)

type Extension interface {
	Run(ctx context.Context) error
	RegisterFunction(s string, f ExternalFunction)
	RegisterExtension(newExtension *resource_model.Extension)
	getServiceKey() string
	WithServiceKey(serviceKey string) Extension
	PrepareCall(*resource_model.Extension) resource_model.ExternalCall
}
