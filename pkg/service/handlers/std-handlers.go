package handlers

import (
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/service"
	"github.com/apibrew/apibrew/pkg/service/backend-event-handler"
)

type StdHandlers interface {
	Init(config *model.AppConfig)
}

type stdHandlers struct {
	backendEventHandler backend_event_handler.BackendEventHandler
	userHandler         *userHandler
	extensionHandler    *extensionHandler
}

func (s *stdHandlers) Init(config *model.AppConfig) {
	s.userHandler.Register(s.backendEventHandler)
	s.extensionHandler.Register(s.backendEventHandler)
}

func NewStdHandler(backendEventHandler backend_event_handler.BackendEventHandler, extensionService service.ExtensionService) StdHandlers {
	return &stdHandlers{
		backendEventHandler: backendEventHandler,
		userHandler:         &userHandler{},
		extensionHandler:    &extensionHandler{extensionService: extensionService},
	}
}
