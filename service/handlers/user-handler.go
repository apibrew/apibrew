package handlers

import (
	"context"
	"github.com/tislib/data-handler/model"
	"github.com/tislib/data-handler/service"
	"github.com/tislib/data-handler/service/errors"
	"github.com/tislib/data-handler/service/handler"
	"github.com/tislib/data-handler/service/params"
	"github.com/tislib/data-handler/service/security"
	"github.com/tislib/data-handler/service/system"
	"google.golang.org/protobuf/types/known/structpb"
)

type userHandler struct {
	handler.BaseHandler
	userService   service.UserService
	recordService service.RecordService
}

func (h *userHandler) BeforeCreate(ctx context.Context, resource *model.Resource, params params.RecordCreateParams) errors.ServiceError {
	for _, user := range params.Records {
		if user.Properties["password"] != nil && user.Properties["password"].GetStringValue() != "" {
			hashStr, err := security.EncodeKey(user.Properties["password"].GetStringValue())

			if err != nil {
				panic(err)
			}

			user.Properties["password"] = structpb.NewStringValue(hashStr)
		}
	}

	return nil
}

func (h *userHandler) AfterList(ctx context.Context, resource *model.Resource, params params.RecordListParams, records []*model.Record, total uint32) errors.ServiceError {
	if !security.IsSystemContext(ctx) {
		h.cleanPasswords(records)
	}

	return nil
}

func (h *userHandler) AfterCreate(ctx context.Context, resource *model.Resource, params params.RecordCreateParams, records []*model.Record) errors.ServiceError {
	if !security.IsSystemContext(ctx) {
		h.cleanPasswords(records)
	}

	return nil
}

func (h *userHandler) BeforeUpdate(ctx context.Context, resource *model.Resource, params2 params.RecordUpdateParams) errors.ServiceError {
	for _, user := range params2.Records {
		if user.Properties["password"] != nil && user.Properties["password"].GetStringValue() != "" {
			hashStr, err := security.EncodeKey(user.Properties["password"].GetStringValue())

			if err != nil {
				panic(err)
			}

			user.Properties["password"] = structpb.NewStringValue(hashStr)
		} else {
			record, err := h.recordService.Get(ctx, params.RecordGetParams{
				Namespace: system.UserResource.Namespace,
				Resource:  system.UserResource.Name,
				Id:        user.Id,
			})

			if err != nil {
				return err
			}

			user.Properties["password"] = record.Properties["password"]
		}
	}

	return nil
}

func (h *userHandler) AfterUpdate(ctx context.Context, resource *model.Resource, params params.RecordUpdateParams, records []*model.Record) errors.ServiceError {
	if !security.IsSystemContext(ctx) {
		h.cleanPasswords(records)
	}

	return nil
}

func (h *userHandler) AfterGet(ctx context.Context, resource *model.Resource, id string, record *model.Record) errors.ServiceError {
	if !security.IsSystemContext(ctx) {
		h.cleanPasswords([]*model.Record{record})
	}

	return nil
}

func (h *userHandler) prepareHandler() *handler.BaseHandler {
	return &handler.BaseHandler{
		BeforeCreate: h.BeforeCreate,
		AfterCreate:  h.AfterCreate,

		BeforeUpdate: h.BeforeUpdate,
		AfterUpdate:  h.AfterUpdate,

		AfterList: h.AfterList,

		AfterGet: h.AfterGet,
	}
}

func (h *userHandler) cleanPasswords(users []*model.Record) {
	for _, user := range users {
		user.Properties["password"] = nil
	}
}
