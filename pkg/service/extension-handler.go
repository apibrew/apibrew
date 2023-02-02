package service

import (
	"context"
	log "github.com/sirupsen/logrus"
	"github.com/tislib/data-handler/pkg/abs"
	"github.com/tislib/data-handler/pkg/errors"
	"github.com/tislib/data-handler/pkg/ext"
	"github.com/tislib/data-handler/pkg/model"
	"github.com/tislib/data-handler/pkg/service/handler"
)

type extensionHandler struct {
	handler.BaseHandler
	service abs.Extension
}

func (h *extensionHandler) BeforeList(ctx context.Context, resource *model.Resource, params abs.RecordListParams) errors.ServiceError {
	_, err := h.service.BeforeList(ctx, &ext.BeforeListRecordRequest{
		Resource: resource,
		Params: &ext.RecordListParams{
			Namespace:         resource.Namespace,
			Resource:          resource.Name,
			Query:             params.Query,
			Limit:             params.Limit,
			Offset:            params.Offset,
			UseHistory:        params.UseHistory,
			ResolveReferences: params.ResolveReferences,
		},
	})

	if err != nil {
		log.Error(err)
		return errors.InternalError.WithDetails("External communication error")
	}

	return nil
}

func (h *extensionHandler) List(ctx context.Context, resource *model.Resource, params abs.RecordListParams) (bool, []*model.Record, uint32, errors.ServiceError) {
	resp, err := h.service.List(ctx, &ext.ListRecordRequest{
		Resource: resource,
		Params: &ext.RecordListParams{
			Namespace:         resource.Namespace,
			Resource:          resource.Name,
			Query:             params.Query,
			Limit:             params.Limit,
			Offset:            params.Offset,
			UseHistory:        params.UseHistory,
			ResolveReferences: params.ResolveReferences,
		},
	})

	if err != nil {
		log.Error(err)
		return true, nil, 0, errors.InternalError.WithDetails("External communication error")
	}

	return true, resp.Records, resp.Total, nil
}

func (h *extensionHandler) AfterList(ctx context.Context, resource *model.Resource, params abs.RecordListParams, records []*model.Record, total uint32) errors.ServiceError {
	_, err := h.service.AfterList(ctx, &ext.AfterListRecordRequest{
		Resource: resource,
		Params: &ext.RecordListParams{
			Namespace:         resource.Namespace,
			Resource:          resource.Name,
			Query:             params.Query,
			Limit:             params.Limit,
			Offset:            params.Offset,
			UseHistory:        params.UseHistory,
			ResolveReferences: params.ResolveReferences,
		},
		Records: records,
		Total:   total,
	})

	if err != nil {
		log.Error(err)
		return errors.InternalError.WithDetails("External communication error")
	}

	return nil
}

func (h *extensionHandler) Create(ctx context.Context, resource *model.Resource, params abs.RecordCreateParams) (bool, []*model.Record, []bool, errors.ServiceError) {
	resp, err := h.service.Create(ctx, &ext.CreateRecordRequest{
		Resource: resource,
		Params: &ext.RecordCreateParams{
			Namespace:      resource.Namespace,
			Records:        params.Records,
			IgnoreIfExists: params.IgnoreIfExists,
		},
	})

	if err != nil {
		log.Error(err)
		return true, nil, nil, errors.InternalError.WithDetails("External communication error")
	}

	return true, resp.Records, resp.Inserted, nil
}

func (h *extensionHandler) Update(ctx context.Context, resource *model.Resource, params abs.RecordUpdateParams) (bool, []*model.Record, errors.ServiceError) {
	resp, err := h.service.Update(ctx, &ext.UpdateRecordRequest{
		Resource: resource,
		Params: &ext.RecordUpdateParams{
			Namespace:    resource.Namespace,
			Records:      params.Records,
			CheckVersion: params.CheckVersion,
		},
	})

	if err != nil {
		log.Error(err)
		return true, nil, errors.InternalError.WithDetails("External communication error")
	}

	return true, resp.Records, nil
}

func (h *extensionHandler) Get(ctx context.Context, resource *model.Resource, id string) (handled bool, record *model.Record, error errors.ServiceError) {
	resp, err := h.service.Get(ctx, &ext.GetRecordRequest{
		Resource: resource,
		Id:       id,
	})

	if err != nil {
		log.Error(err)
		return true, nil, errors.InternalError.WithDetails("External communication error")
	}

	return true, resp.Record, nil
}

func (h *extensionHandler) Delete(ctx context.Context, resource *model.Resource, params abs.RecordDeleteParams) (bool, errors.ServiceError) {
	_, err := h.service.Delete(ctx, &ext.DeleteRecordRequest{
		Resource: resource,
		Params: &ext.RecordDeleteParams{
			Namespace: params.Namespace,
			Resource:  params.Resource,
			Ids:       params.Ids,
		},
	})

	if err != nil {
		log.Error(err)
		return true, errors.InternalError.WithDetails("External communication error")
	}

	return true, nil
}

func NewExtensionHandler(service abs.Extension) *handler.BaseHandler {
	h := &handler.BaseHandler{}

	extensionConfig := service.GetExtensionConfig()
	exth := &extensionHandler{service: service}

	for _, op := range extensionConfig.Operations {
		if op.OperationType == model.ExtensionOperationType_ExtensionOperationTypeList {
			if op.Step == model.ExtensionOperationStep_ExtensionOperationStepBefore {
				if op.Sync {
					h.BeforeList = exth.BeforeList
				} else {
					h.BeforeList = func(ctx context.Context, resource *model.Resource, params abs.RecordListParams) errors.ServiceError {
						go func() {
							err := exth.BeforeList(context.TODO(), resource, params)

							if err != nil {
								log.Error(err)
							}
						}()
						return nil
					}
				}
			} else if op.Step == model.ExtensionOperationStep_ExtensionOperationStepInstead {
				h.List = exth.List
			} else if op.Step == model.ExtensionOperationStep_ExtensionOperationStepAfter {
				if op.Sync {
					h.AfterList = exth.AfterList
				} else {
					h.AfterList = func(ctx context.Context, resource *model.Resource, params abs.RecordListParams, records []*model.Record, total uint32) errors.ServiceError {
						go func() {
							err := exth.AfterList(context.TODO(), resource, params, records, total)

							if err != nil {
								log.Error(err)
							}
						}()
						return nil
					}
				}
			} else {
				panic("Unknown step:" + op.Step.String())
			}
		} else if op.OperationType == model.ExtensionOperationType_ExtensionOperationTypeCreate {
			if op.Step == model.ExtensionOperationStep_ExtensionOperationStepBefore {
				if op.Sync {
					h.BeforeCreate = exth.BeforeCreate
				} else {
					h.BeforeCreate = func(ctx context.Context, resource *model.Resource, params abs.RecordCreateParams) errors.ServiceError {
						go func() {
							err := exth.BeforeCreate(context.TODO(), resource, params)

							if err != nil {
								log.Error(err)
							}
						}()
						return nil
					}
				}
			} else if op.Step == model.ExtensionOperationStep_ExtensionOperationStepInstead {
				h.Create = exth.Create
			} else if op.Step == model.ExtensionOperationStep_ExtensionOperationStepAfter {
				if op.Sync {
					h.AfterCreate = exth.AfterCreate
				} else {
					h.AfterCreate = func(ctx context.Context, resource *model.Resource, params abs.RecordCreateParams, records []*model.Record) errors.ServiceError {
						go func() {
							err := exth.AfterCreate(context.TODO(), resource, params, records)

							if err != nil {
								log.Error(err)
							}
						}()
						return nil
					}
				}
			} else {
				panic("Unknown step:" + op.Step.String())
			}
		} else if op.OperationType == model.ExtensionOperationType_ExtensionOperationTypeUpdate {
			if op.Step == model.ExtensionOperationStep_ExtensionOperationStepBefore {
				h.BeforeUpdate = exth.BeforeUpdate
			} else if op.Step == model.ExtensionOperationStep_ExtensionOperationStepInstead {
				h.Update = exth.Update
			} else if op.Step == model.ExtensionOperationStep_ExtensionOperationStepAfter {
				h.AfterUpdate = exth.AfterUpdate
			} else {
				panic("Unknown step:" + op.Step.String())
			}
		} else if op.OperationType == model.ExtensionOperationType_ExtensionOperationTypeGet {
			if op.Step == model.ExtensionOperationStep_ExtensionOperationStepBefore {
				h.BeforeGet = exth.BeforeGet
			} else if op.Step == model.ExtensionOperationStep_ExtensionOperationStepInstead {
				h.Get = exth.Get
			} else if op.Step == model.ExtensionOperationStep_ExtensionOperationStepAfter {
				h.AfterGet = exth.AfterGet
			} else {
				panic("Unknown step:" + op.Step.String())
			}
		} else if op.OperationType == model.ExtensionOperationType_ExtensionOperationTypeDelete {
			if op.Step == model.ExtensionOperationStep_ExtensionOperationStepBefore {
				h.BeforeDelete = exth.BeforeDelete
			} else if op.Step == model.ExtensionOperationStep_ExtensionOperationStepInstead {
				h.Delete = exth.Delete
			} else if op.Step == model.ExtensionOperationStep_ExtensionOperationStepAfter {
				h.AfterDelete = exth.AfterDelete
			} else {
				panic("Unknown step:" + op.Step.String())
			}
		}
	}

	return h
}
