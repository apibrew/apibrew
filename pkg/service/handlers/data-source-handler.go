package handlers

import (
	"context"
	"github.com/apibrew/apibrew/pkg/abs"
	"github.com/apibrew/apibrew/pkg/errors"
	"github.com/apibrew/apibrew/pkg/model"
	backend_event_handler "github.com/apibrew/apibrew/pkg/service/backend-event-handler"
)

type dataSourceHandler struct {
}

func (h *dataSourceHandler) BeforeCreate(ctx context.Context, resource *model.Resource, params abs.RecordCreateParams) errors.ServiceError {
	if resource.Namespace != "system" || resource.Name != "data-source" {
		return nil
	}

	return nil
}

func (h *dataSourceHandler) BeforeList(ctx context.Context, resource *model.Resource, params abs.RecordListParams) errors.ServiceError {
	if resource.Namespace != "system" || resource.Name != "data-source" {
		return nil
	}

	return nil
}

func (h *dataSourceHandler) Register(eventHandler backend_event_handler.BackendEventHandler) {

}
