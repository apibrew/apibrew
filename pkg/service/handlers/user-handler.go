package handlers

import (
	"context"
	"github.com/apibrew/apibrew/pkg/abs"
	"github.com/apibrew/apibrew/pkg/core"
	"github.com/apibrew/apibrew/pkg/resources"
	"github.com/apibrew/apibrew/pkg/service/backend-event-handler"
	"github.com/apibrew/apibrew/pkg/util"
	"google.golang.org/protobuf/types/known/structpb"
)

type userHandler struct {
}

func (h *userHandler) Register(eventHandler backend_event_handler.BackendEventHandler) {
	eventHandler.RegisterHandler(prepareStdHandler(1, core.Event_CREATE, h.BeforeCreate, resources.UserResource))
	eventHandler.RegisterHandler(prepareStdHandler(101, core.Event_LIST, h.AfterList, resources.UserResource))
	eventHandler.RegisterHandler(prepareStdHandler(101, core.Event_CREATE, h.AfterCreate, resources.UserResource))
	eventHandler.RegisterHandler(prepareStdHandler(1, core.Event_UPDATE, h.BeforeUpdate, resources.UserResource))
	eventHandler.RegisterHandler(prepareStdHandler(101, core.Event_UPDATE, h.AfterUpdate, resources.UserResource))
	eventHandler.RegisterHandler(prepareStdHandler(101, core.Event_GET, h.AfterGet, resources.UserResource))
}

func (h *userHandler) BeforeCreate(ctx context.Context, event *core.Event) (*core.Event, error) {
	for _, user := range event.Records {
		if user.GetProperties()["password"] != nil && user.GetProperties()["password"].GetStringValue() != "" {
			hashStr, err := util.EncodeKey(user.GetProperties()["password"].GetStringValue())

			if err != nil {
				panic(err)
			}

			user.GetProperties()["password"] = structpb.NewStringValue(hashStr)
		}
	}

	return event, nil
}

func (h *userHandler) AfterList(ctx context.Context, event *core.Event) (*core.Event, error) {
	if !util.IsSystemContext(ctx) {
		h.cleanPasswords(event.Records)
	}

	return event, nil
}

func (h *userHandler) AfterCreate(ctx context.Context, event *core.Event) (*core.Event, error) {
	if !util.IsSystemContext(ctx) {
		h.cleanPasswords(event.Records)
	}

	return event, nil
}

func (h *userHandler) BeforeUpdate(ctx context.Context, event *core.Event) (*core.Event, error) {
	for _, user := range event.Records {
		if user.GetProperties()["password"] != nil && user.GetProperties()["password"].GetStringValue() != "" {
			hashStr, err := util.EncodeKey(user.GetProperties()["password"].GetStringValue())

			if err != nil {
				panic(err)
			}

			user.GetProperties()["password"] = structpb.NewStringValue(hashStr)
		}
	}

	return event, nil
}

func (h *userHandler) AfterUpdate(ctx context.Context, event *core.Event) (*core.Event, error) {
	if !util.IsSystemContext(ctx) {
		h.cleanPasswords(event.Records)
	}

	return event, nil
}

func (h *userHandler) AfterGet(ctx context.Context, event *core.Event) (*core.Event, error) {
	if event.Records == nil || len(event.Records) == 0 {
		return event, nil
	}

	if !util.IsSystemContext(ctx) {
		h.cleanPasswords([]abs.RecordLike{event.Records[0]})
	}

	return event, nil
}

func (h *userHandler) cleanPasswords(users []abs.RecordLike) {
	for _, user := range users {
		delete(user.GetProperties(), "password")
	}
}
