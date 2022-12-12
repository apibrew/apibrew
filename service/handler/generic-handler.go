package handler

import (
	"context"
	"data-handler/model"
	"data-handler/service/errors"
	"data-handler/service/params"
)

type GenericHandler interface {
	BaseHandler
}

type genericHandler struct {
}

func (g genericHandler) BeforeList(ctx context.Context, resource *model.Resource, params params.RecordListParams) errors.ServiceError {
	return nil
}

func (g genericHandler) List() (handled bool, records []*model.Record, total uint32, err errors.ServiceError) {
	return false, nil, 0, nil
}

func (g genericHandler) AfterList(ctx context.Context, resource *model.Resource, params params.RecordListParams, records []*model.Record, total uint32) errors.ServiceError {
	return nil
}

func (g genericHandler) BeforeCreate(ctx context.Context, resource *model.Resource, params params.RecordCreateParams) errors.ServiceError {
	return nil
}

func (g genericHandler) Create(ctx context.Context, resource *model.Resource, params params.RecordCreateParams) (handled bool, records []*model.Record, inserted []bool, err errors.ServiceError) {
	return false, nil, nil, nil
}

func (g genericHandler) AfterCreate(ctx context.Context, resource *model.Resource, params params.RecordCreateParams, records []*model.Record) errors.ServiceError {
	return nil
}

func (g genericHandler) BeforeUpdate(ctx context.Context, resource *model.Resource, params params.RecordUpdateParams) errors.ServiceError {
	return nil
}

func (g genericHandler) Update(ctx context.Context, resource *model.Resource, params params.RecordUpdateParams) (handled bool, records []*model.Record, err errors.ServiceError) {
	return false, nil, nil
}

func (g genericHandler) AfterUpdate(ctx context.Context, resource *model.Resource, params params.RecordUpdateParams, records []*model.Record) errors.ServiceError {
	return nil
}

func (g genericHandler) BeforeGet(resource *model.Resource, id string) errors.ServiceError {
	return nil
}

func (g genericHandler) Get(resource *model.Resource, id string) (handled bool, record *model.Record, error errors.ServiceError) {
	return false, nil, nil
}

func (g genericHandler) AfterGet(resource *model.Resource, id string, res *model.Record) errors.ServiceError {
	return nil
}

func (g genericHandler) BeforeDelete(ctx context.Context, params params.RecordDeleteParams) errors.ServiceError {
	return nil
}

func (g genericHandler) Delete(ctx context.Context, params params.RecordDeleteParams) (handled bool, err errors.ServiceError) {
	return false, nil
}

func (g genericHandler) AfterDelete(ctx context.Context, params params.RecordDeleteParams) errors.ServiceError {
	return nil
}

func NewGenericHandler() GenericHandler {
	return &genericHandler{}
}
