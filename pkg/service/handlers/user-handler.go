package handlers

import (
	"context"
	"github.com/tislib/apibrew/pkg/errors"
	"github.com/tislib/apibrew/pkg/model"
	"github.com/tislib/apibrew/pkg/resources"
	backend_event_handler "github.com/tislib/apibrew/pkg/service/backend-event-handler"
	"github.com/tislib/apibrew/pkg/service/security"
	"google.golang.org/protobuf/types/known/structpb"
)

type userHandler struct {
}

func (h *userHandler) Register(eventHandler backend_event_handler.BackendEventHandler) {
	eventHandler.RegisterHandler(prepareStdHandler(1, model.Event_CREATE, h.BeforeCreate, resources.UserResource))
	eventHandler.RegisterHandler(prepareStdHandler(101, model.Event_LIST, h.AfterList, resources.UserResource))
	eventHandler.RegisterHandler(prepareStdHandler(101, model.Event_CREATE, h.AfterCreate, resources.UserResource))
	eventHandler.RegisterHandler(prepareStdHandler(1, model.Event_UPDATE, h.BeforeUpdate, resources.UserResource))
	eventHandler.RegisterHandler(prepareStdHandler(101, model.Event_UPDATE, h.AfterUpdate, resources.UserResource))
	eventHandler.RegisterHandler(prepareStdHandler(101, model.Event_GET, h.AfterGet, resources.UserResource))
}

func (h *userHandler) BeforeCreate(ctx context.Context, event *model.Event) (*model.Event, errors.ServiceError) {
	for _, user := range event.Records {
		if user.Properties["password"] != nil && user.Properties["password"].GetStringValue() != "" {
			hashStr, err := security.EncodeKey(user.Properties["password"].GetStringValue())

			if err != nil {
				panic(err)
			}

			user.Properties["password"] = structpb.NewStringValue(hashStr)
		}
	}

	return event, nil
}

func (h *userHandler) AfterList(ctx context.Context, event *model.Event) (*model.Event, errors.ServiceError) {
	if !security.IsSystemContext(ctx) {
		h.cleanPasswords(event.Records)
	}

	return event, nil
}

func (h *userHandler) AfterCreate(ctx context.Context, event *model.Event) (*model.Event, errors.ServiceError) {
	if !security.IsSystemContext(ctx) {
		h.cleanPasswords(event.Records)
	}

	return event, nil
}

func (h *userHandler) BeforeUpdate(ctx context.Context, event *model.Event) (*model.Event, errors.ServiceError) {
	for _, user := range event.Records {
		if user.Properties["password"] != nil && user.Properties["password"].GetStringValue() != "" {
			hashStr, err := security.EncodeKey(user.Properties["password"].GetStringValue())

			if err != nil {
				panic(err)
			}

			user.Properties["password"] = structpb.NewStringValue(hashStr)
		}
	}

	return event, nil
}

func (h *userHandler) AfterUpdate(ctx context.Context, event *model.Event) (*model.Event, errors.ServiceError) {
	if !security.IsSystemContext(ctx) {
		h.cleanPasswords(event.Records)
	}

	return event, nil
}

func (h *userHandler) AfterGet(ctx context.Context, event *model.Event) (*model.Event, errors.ServiceError) {
	if event.Records == nil || len(event.Records) == 0 {
		return event, nil
	}

	if !security.IsSystemContext(ctx) {
		h.cleanPasswords([]*model.Record{event.Records[0]})
	}

	return event, nil
}

func (h *userHandler) cleanPasswords(users []*model.Record) {
	for _, user := range users {
		delete(user.Properties, "password")
	}
}
