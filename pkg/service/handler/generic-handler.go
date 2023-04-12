package handler

import (
	"context"
	"fmt"
	"github.com/tislib/data-handler/pkg/abs"
	"github.com/tislib/data-handler/pkg/errors"
	"github.com/tislib/data-handler/pkg/model"
	"math/rand"
)

type EventSelector func(ctx context.Context, resource *model.Resource) bool

func ResourceSelector(expectedResource *model.Resource) EventSelector {
	return func(ctx context.Context, resource *model.Resource) bool {
		return expectedResource.Namespace == resource.Namespace && expectedResource.Name == resource.Name
	}
}

type GenericHandler struct {
	BaseHandler
	handlers    []*BaseHandler
	selectorMap map[*BaseHandler]EventSelector
}

func (g *GenericHandler) Register(handler *BaseHandler) {
	handler.Id = fmt.Sprintf("%v", rand.Float64())
	g.handlers = append(g.handlers, handler)
}

func (g *GenericHandler) RegisterWithSelector(handler *BaseHandler, selector EventSelector) {
	handler.Id = fmt.Sprintf("%v", rand.Float64())
	g.handlers = append(g.handlers, handler)
	g.selectorMap[handler] = selector
}

func (g *GenericHandler) BeforeList(ctx context.Context, resource *model.Resource, params abs.RecordListParams) errors.ServiceError {
	for _, item := range g.handlers {
		if g.selectorMap[item] != nil && !g.selectorMap[item](ctx, resource) {
			continue
		}
		if item.BeforeList != nil {
			if err := item.BeforeList(ctx, resource, params); err != nil {
				return err
			}
		}
	}

	return nil
}

func (g *GenericHandler) List(ctx context.Context, resource *model.Resource, params abs.RecordListParams) (handled bool, records []*model.Record, total uint32, err errors.ServiceError) {
	for _, item := range g.handlers {
		if g.selectorMap[item] != nil && !g.selectorMap[item](ctx, resource) {
			continue
		}
		if item.List != nil {
			if handled, records, total, err = item.List(ctx, resource, params); handled {
				return
			}
		}
	}

	return
}

func (g *GenericHandler) AfterList(ctx context.Context, resource *model.Resource, params abs.RecordListParams, records []*model.Record, total uint32) errors.ServiceError {
	for _, item := range g.handlers {
		if g.selectorMap[item] != nil && !g.selectorMap[item](ctx, resource) {
			continue
		}
		if item.AfterList != nil {
			if err := item.AfterList(ctx, resource, params, records, total); err != nil {
				return err
			}
		}
	}

	return nil
}

func (g *GenericHandler) BeforeCreate(ctx context.Context, resource *model.Resource, params abs.RecordCreateParams) errors.ServiceError {
	for _, item := range g.handlers {
		if g.selectorMap[item] != nil && !g.selectorMap[item](ctx, resource) {
			continue
		}
		if item.BeforeCreate == nil {
			continue
		}
		if err := item.BeforeCreate(ctx, resource, params); err != nil {
			return err
		}
	}

	return nil
}

func (g *GenericHandler) Create(ctx context.Context, resource *model.Resource, params abs.RecordCreateParams) (handled bool, records []*model.Record, inserted []bool, err errors.ServiceError) {
	records = params.Records

	for _, item := range g.handlers {
		if g.selectorMap[item] != nil && !g.selectorMap[item](ctx, resource) {
			continue
		}
		if item.List != nil {
			if handled, records, inserted, err = item.Create(ctx, resource, params); handled {
				return
			}
		}
	}

	return
}

func (g *GenericHandler) AfterCreate(ctx context.Context, resource *model.Resource, params abs.RecordCreateParams, records []*model.Record) errors.ServiceError {
	for _, item := range g.handlers {
		if g.selectorMap[item] != nil && !g.selectorMap[item](ctx, resource) {
			continue
		}
		if item.AfterCreate == nil {
			continue
		}

		if err := item.AfterCreate(ctx, resource, params, records); err != nil {
			return err
		}
	}

	return nil
}

func (g *GenericHandler) BeforeUpdate(ctx context.Context, resource *model.Resource, params abs.RecordUpdateParams) errors.ServiceError {
	for _, item := range g.handlers {
		if g.selectorMap[item] != nil && !g.selectorMap[item](ctx, resource) {
			continue
		}
		if item.BeforeUpdate == nil {
			continue
		}

		if err := item.BeforeUpdate(ctx, resource, params); err != nil {
			return err
		}
	}

	return nil
}

func (g *GenericHandler) Update(ctx context.Context, resource *model.Resource, params abs.RecordUpdateParams) (handled bool, records []*model.Record, err errors.ServiceError) {
	records = params.Records

	for _, item := range g.handlers {
		if g.selectorMap[item] != nil && !g.selectorMap[item](ctx, resource) {
			continue
		}
		if item.List != nil {
			if handled, records, err = item.Update(ctx, resource, params); handled {
				return
			}
		}
	}

	return
}

func (g *GenericHandler) AfterUpdate(ctx context.Context, resource *model.Resource, params abs.RecordUpdateParams, records []*model.Record) errors.ServiceError {
	for _, item := range g.handlers {
		if g.selectorMap[item] != nil && !g.selectorMap[item](ctx, resource) {
			continue
		}
		if item.AfterUpdate == nil {
			continue
		}

		if err := item.AfterUpdate(ctx, resource, params, records); err != nil {
			return err
		}
	}

	return nil
}

func (g *GenericHandler) BeforeGet(ctx context.Context, resource *model.Resource, id string) errors.ServiceError {
	for _, item := range g.handlers {
		if g.selectorMap[item] != nil && !g.selectorMap[item](ctx, resource) {
			continue
		}
		if item.BeforeGet == nil {
			continue
		}

		if err := item.BeforeGet(ctx, resource, id); err != nil {
			return err
		}
	}

	return nil
}

func (g *GenericHandler) Get(ctx context.Context, resource *model.Resource, id string) (handled bool, record *model.Record, error errors.ServiceError) {
	for _, item := range g.handlers {
		if g.selectorMap[item] != nil && !g.selectorMap[item](ctx, resource) {
			continue
		}
		if item.List != nil {
			if handled, record, error = item.Get(ctx, resource, id); handled {
				return
			}
		}
	}

	return
}

func (g *GenericHandler) AfterGet(ctx context.Context, resource *model.Resource, id string, res *model.Record) errors.ServiceError {
	for _, item := range g.handlers {
		if g.selectorMap[item] != nil && !g.selectorMap[item](ctx, resource) {
			continue
		}
		if item.AfterGet == nil {
			continue
		}

		if err := item.AfterGet(ctx, resource, id, res); err != nil {
			return err
		}
	}

	return nil
}

func (g *GenericHandler) BeforeDelete(ctx context.Context, resource *model.Resource, params abs.RecordDeleteParams) errors.ServiceError {
	for _, item := range g.handlers {
		if g.selectorMap[item] != nil && !g.selectorMap[item](ctx, &model.Resource{Namespace: params.Namespace, Name: params.Resource}) {
			continue
		}
		if item.BeforeDelete == nil {
			continue
		}

		if err := item.BeforeDelete(ctx, resource, params); err != nil {
			return err
		}
	}

	return nil
}

func (g *GenericHandler) Delete(ctx context.Context, resource *model.Resource, params abs.RecordDeleteParams) (handled bool, err errors.ServiceError) {
	for _, item := range g.handlers {
		if g.selectorMap[item] != nil && !g.selectorMap[item](ctx, resource) {
			continue
		}
		if item.List != nil {
			if handled, err = item.Delete(ctx, resource, params); handled {
				return
			}
		}
	}

	return
}

func (g *GenericHandler) AfterDelete(ctx context.Context, resource *model.Resource, params abs.RecordDeleteParams) errors.ServiceError {
	for _, item := range g.handlers {
		if g.selectorMap[item] != nil && !g.selectorMap[item](ctx, &model.Resource{Namespace: params.Namespace, Name: params.Resource}) {
			continue
		}
		if item.AfterDelete == nil {
			continue
		}

		if err := item.AfterDelete(ctx, resource, params); err != nil {
			return err
		}
	}

	return nil
}

func (g *GenericHandler) Unregister(handler *BaseHandler) {
	var newHandlers []*BaseHandler

	for _, item := range g.handlers {
		if handler == item {
			continue
		}

		newHandlers = append(newHandlers, item)
	}

	g.handlers = newHandlers
}

func NewGenericHandler() *GenericHandler {
	return &GenericHandler{
		selectorMap: make(map[*BaseHandler]EventSelector),
	}
}
