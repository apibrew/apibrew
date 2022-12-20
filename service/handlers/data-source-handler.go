package handlers

import (
	"context"
	"data-handler/model"
	"data-handler/service"
	"data-handler/service/errors"
	"data-handler/service/handler"
	"data-handler/service/params"
)

type dataSourceHandler struct {
	handler.BaseHandler
	dataSourceService service.DataSourceService
}

func (h *dataSourceHandler) BeforeCreate(ctx context.Context, resource *model.Resource, params params.RecordCreateParams) errors.ServiceError {
	if resource.Workspace != "system" || resource.Name != "data-source" {
		return nil
	}
	return nil
}

func (h *dataSourceHandler) BeforeList(ctx context.Context, resource *model.Resource, params params.RecordListParams) errors.ServiceError {
	if resource.Workspace != "system" || resource.Name != "data-source" {
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
