package service

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/tislib/data-handler/pkg/errors"
	"github.com/tislib/data-handler/pkg/ext"
	"github.com/tislib/data-handler/pkg/model"
	"github.com/tislib/data-handler/pkg/service/handler"
	"github.com/tislib/data-handler/pkg/service/params"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type extensionHandler struct {
	handler.BaseHandler
	recordClient ext.RecordExtensionServiceClient
}

func (h *extensionHandler) List(ctx context.Context, resource *model.Resource, params params.RecordListParams) (bool, []*model.Record, uint32, errors.ServiceError) {
	resp, err := h.recordClient.List(ctx, &ext.ListRecordRequest{
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

func (h *extensionHandler) Create(ctx context.Context, resource *model.Resource, params params.RecordCreateParams) (bool, []*model.Record, []bool, errors.ServiceError) {
	resp, err := h.recordClient.Create(ctx, &ext.CreateRecordRequest{
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

func (h *extensionHandler) Update(ctx context.Context, resource *model.Resource, params params.RecordUpdateParams) (bool, []*model.Record, errors.ServiceError) {
	resp, err := h.recordClient.Update(ctx, &ext.UpdateRecordRequest{
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
	resp, err := h.recordClient.Get(ctx, &ext.GetRecordRequest{
		Resource: resource,
		Id:       id,
	})

	if err != nil {
		log.Error(err)
		return true, nil, errors.InternalError.WithDetails("External communication error")
	}

	return true, resp.Record, nil
}

func (h *extensionHandler) Delete(ctx context.Context, resource *model.Resource, params params.RecordDeleteParams) (bool, errors.ServiceError) {
	_, err := h.recordClient.Delete(ctx, &ext.DeleteRecordRequest{
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

func NewExtensionHandler(extension *model.Extension) *handler.BaseHandler {
	exth := new(extensionHandler)

	var opts []grpc.DialOption

	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", extension.Server.Host, extension.Server.Port), opts...)
	if err != nil {
		panic(err)
	}

	exth.recordClient = ext.NewRecordExtensionServiceClient(conn)

	h := &handler.BaseHandler{}

	for _, op := range extension.Operations {
		if op.OperationType == model.ExtensionOperationType_ExtensionOperationTypeList {
			if op.Step == model.ExtensionOperationStep_ExtensionOperationStepBefore {
				if op.Sync {
					h.BeforeList = exth.BeforeList
				} else {
					h.BeforeList = func(ctx context.Context, resource *model.Resource, params params.RecordListParams) errors.ServiceError {
						go func() {
							err := exth.BeforeList(ctx, resource, params)

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
					h.AfterList = func(ctx context.Context, resource *model.Resource, params params.RecordListParams, records []*model.Record, total uint32) errors.ServiceError {
						go func() {
							err := exth.AfterList(ctx, resource, params, records, total)

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
					h.BeforeCreate = func(ctx context.Context, resource *model.Resource, params params.RecordCreateParams) errors.ServiceError {
						go func() {
							err := exth.BeforeCreate(ctx, resource, params)

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
					h.AfterCreate = func(ctx context.Context, resource *model.Resource, params params.RecordCreateParams, records []*model.Record) errors.ServiceError {
						go func() {
							err := exth.AfterCreate(ctx, resource, params, records)

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
