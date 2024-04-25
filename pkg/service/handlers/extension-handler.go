package handlers

import (
	"context"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/resources"
	"github.com/apibrew/apibrew/pkg/service"
	"github.com/apibrew/apibrew/pkg/service/backend-event-handler"
)

type extensionHandler struct {
	extensionService service.ExtensionService
}

func (h *extensionHandler) Register(eventHandler backend_event_handler.BackendEventHandler) {
	eventHandler.RegisterHandler(prepareStdHandler(101, model.Event_CREATE, h.AfterChange, resources.ExtensionResource))
	eventHandler.RegisterHandler(prepareStdHandler(101, model.Event_UPDATE, h.AfterChange, resources.ExtensionResource))
	eventHandler.RegisterHandler(prepareStdHandler(101, model.Event_DELETE, h.AfterChange, resources.ExtensionResource))
}

func (h *extensionHandler) AfterChange(ctx context.Context, event *model.Event) (*model.Event, error) {

	go h.extensionService.Reload()

	return event, nil
}
