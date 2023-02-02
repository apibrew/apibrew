package handlers

import (
	"context"
	"github.com/tislib/data-handler/pkg/abs"
	"github.com/tislib/data-handler/pkg/errors"
	"github.com/tislib/data-handler/pkg/model"
	"github.com/tislib/data-handler/pkg/service/handler"
	"github.com/tislib/data-handler/pkg/service/security"
	"google.golang.org/protobuf/types/known/structpb"
)

type userHandler struct {
	handler.BaseHandler
}

func (h *userHandler) BeforeCreate(ctx context.Context, resource *model.Resource, params abs.RecordCreateParams) errors.ServiceError {
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

func (h *userHandler) AfterList(ctx context.Context, resource *model.Resource, params abs.RecordListParams, records []*model.Record, total uint32) errors.ServiceError {
	if !security.IsSystemContext(ctx) {
		h.cleanPasswords(records)
	}

	return nil
}

func (h *userHandler) AfterCreate(ctx context.Context, resource *model.Resource, params abs.RecordCreateParams, records []*model.Record) errors.ServiceError {
	if !security.IsSystemContext(ctx) {
		h.cleanPasswords(records)
	}

	return nil
}

func (h *userHandler) BeforeUpdate(ctx context.Context, resource *model.Resource, params2 abs.RecordUpdateParams) errors.ServiceError {
	for _, user := range params2.Records {
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

func (h *userHandler) AfterUpdate(ctx context.Context, resource *model.Resource, params abs.RecordUpdateParams, records []*model.Record) errors.ServiceError {
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
		delete(user.Properties, "password")
	}
}
