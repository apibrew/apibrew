package handler

import (
	"context"
	"data-handler/model"
	"data-handler/service/errors"
	"data-handler/service/params"
)

type Event struct {
}

type BeforeList func(ctx context.Context, resource *model.Resource, params params.RecordListParams) errors.ServiceError
type List func(ctx context.Context, params params.RecordListParams) (handled bool, records []*model.Record, total uint32, err errors.ServiceError)
type AfterList func(ctx context.Context, resource *model.Resource, params params.RecordListParams, records []*model.Record, total uint32) errors.ServiceError

type BeforeCreate func(ctx context.Context, resource *model.Resource, params params.RecordCreateParams) errors.ServiceError
type Create func(ctx context.Context, resource *model.Resource, params params.RecordCreateParams) (handled bool, records []*model.Record, inserted []bool, err errors.ServiceError)
type AfterCreate func(ctx context.Context, resource *model.Resource, params params.RecordCreateParams, records []*model.Record) errors.ServiceError

type BeforeUpdate func(ctx context.Context, resource *model.Resource, params params.RecordUpdateParams) errors.ServiceError
type Update func(ctx context.Context, resource *model.Resource, params params.RecordUpdateParams) (handled bool, records []*model.Record, err errors.ServiceError)
type AfterUpdate func(ctx context.Context, resource *model.Resource, params params.RecordUpdateParams, records []*model.Record) errors.ServiceError

type BeforeGet func(ctx context.Context, resource *model.Resource, id string) errors.ServiceError
type Get func(ctx context.Context, resource *model.Resource, id string) (handled bool, record *model.Record, error errors.ServiceError)
type AfterGet func(ctx context.Context, resource *model.Resource, id string, res *model.Record) errors.ServiceError

type BeforeDelete func(ctx context.Context, params params.RecordDeleteParams) errors.ServiceError
type Delete func(ctx context.Context, params params.RecordDeleteParams) (handled bool, err errors.ServiceError)
type AfterDelete func(ctx context.Context, params params.RecordDeleteParams) errors.ServiceError

type BaseHandler struct {
	Id string
	// list
	BeforeList BeforeList
	List       List
	AfterList  AfterList

	// create
	BeforeCreate BeforeCreate
	Create       Create
	AfterCreate  AfterCreate

	// update
	BeforeUpdate BeforeUpdate
	Update       Update
	AfterUpdate  AfterUpdate

	// get
	BeforeGet BeforeGet
	Get       Get
	AfterGet  AfterGet

	// delete
	BeforeDelete BeforeDelete
	Delete       Delete
	AfterDelete  AfterDelete
}
