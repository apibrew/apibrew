package handler

import (
	"context"
	"data-handler/model"
	"data-handler/service/errors"
	"data-handler/service/params"
)

type Event struct {
}

type BaseHandler interface {
	// list
	BeforeList(ctx context.Context, resource *model.Resource, params params.RecordListParams) errors.ServiceError
	List() (handled bool, records []*model.Record, total uint32, err errors.ServiceError)
	AfterList(ctx context.Context, resource *model.Resource, params params.RecordListParams, records []*model.Record, total uint32) errors.ServiceError

	// create
	BeforeCreate(ctx context.Context, resource *model.Resource, params params.RecordCreateParams) errors.ServiceError
	Create(ctx context.Context, resource *model.Resource, params params.RecordCreateParams) (handled bool, records []*model.Record, inserted []bool, err errors.ServiceError)
	AfterCreate(ctx context.Context, resource *model.Resource, params params.RecordCreateParams, records []*model.Record) errors.ServiceError

	// update
	BeforeUpdate(ctx context.Context, resource *model.Resource, params params.RecordUpdateParams) errors.ServiceError

	Update(ctx context.Context, resource *model.Resource, params params.RecordUpdateParams) (handled bool, records []*model.Record, err errors.ServiceError)
	AfterUpdate(ctx context.Context, resource *model.Resource, params params.RecordUpdateParams, records []*model.Record) errors.ServiceError

	// get
	BeforeGet(resource *model.Resource, id string) errors.ServiceError
	Get(resource *model.Resource, id string) (handled bool, record *model.Record, error errors.ServiceError)
	AfterGet(resource *model.Resource, id string, res *model.Record) errors.ServiceError

	// delete
	BeforeDelete(ctx context.Context, params params.RecordDeleteParams) errors.ServiceError
	Delete(ctx context.Context, params params.RecordDeleteParams) (handled bool, err errors.ServiceError)
	AfterDelete(ctx context.Context, params params.RecordDeleteParams) errors.ServiceError
}
