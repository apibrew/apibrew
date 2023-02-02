package handler

import (
	"context"
	"github.com/tislib/data-handler/pkg/abs"
	"github.com/tislib/data-handler/pkg/errors"
	"github.com/tislib/data-handler/pkg/model"
)

type Event struct {
}

type BeforeList func(ctx context.Context, resource *model.Resource, params abs.RecordListParams) errors.ServiceError
type List func(ctx context.Context, resource *model.Resource, params abs.RecordListParams) (handled bool, records []*model.Record, total uint32, err errors.ServiceError)
type AfterList func(ctx context.Context, resource *model.Resource, params abs.RecordListParams, records []*model.Record, total uint32) errors.ServiceError

type BeforeCreate func(ctx context.Context, resource *model.Resource, params abs.RecordCreateParams) errors.ServiceError
type Create func(ctx context.Context, resource *model.Resource, params abs.RecordCreateParams) (handled bool, records []*model.Record, inserted []bool, err errors.ServiceError)
type AfterCreate func(ctx context.Context, resource *model.Resource, params abs.RecordCreateParams, records []*model.Record) errors.ServiceError

type BeforeUpdate func(ctx context.Context, resource *model.Resource, params abs.RecordUpdateParams) errors.ServiceError
type Update func(ctx context.Context, resource *model.Resource, params abs.RecordUpdateParams) (handled bool, records []*model.Record, err errors.ServiceError)
type AfterUpdate func(ctx context.Context, resource *model.Resource, params abs.RecordUpdateParams, records []*model.Record) errors.ServiceError

type BeforeGet func(ctx context.Context, resource *model.Resource, id string) errors.ServiceError
type Get func(ctx context.Context, resource *model.Resource, id string) (handled bool, record *model.Record, error errors.ServiceError)
type AfterGet func(ctx context.Context, resource *model.Resource, id string, res *model.Record) errors.ServiceError

type BeforeDelete func(ctx context.Context, resource *model.Resource, params abs.RecordDeleteParams) errors.ServiceError
type Delete func(ctx context.Context, resource *model.Resource, params abs.RecordDeleteParams) (handled bool, err errors.ServiceError)
type AfterDelete func(ctx context.Context, resource *model.Resource, params abs.RecordDeleteParams) errors.ServiceError

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
