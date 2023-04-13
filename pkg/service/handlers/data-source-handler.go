package handlers

import (
	"context"
	"github.com/tislib/data-handler/pkg/abs"
	"github.com/tislib/data-handler/pkg/errors"
	"github.com/tislib/data-handler/pkg/model"
	"github.com/tislib/data-handler/pkg/service/handler"
)

type dataSourceHandler struct {
	handler.BaseHandler
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

func (h *dataSourceHandler) prepareHandler() *handler.BaseHandler {
	return &handler.BaseHandler{
		BeforeCreate: h.BeforeCreate,
		BeforeList:   h.BeforeList,
	}
}
