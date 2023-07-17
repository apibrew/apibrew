package handlers

import (
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/service"
	backend_event_handler "github.com/apibrew/apibrew/pkg/service/backend-event-handler"
)

type StdHandlers interface {
	Init(config *model.AppConfig)
}

type stdHandlers struct {
	backendEventHandler backend_event_handler.BackendEventHandler
	dataSourceHandler   *dataSourceHandler
	userHandler         *userHandler
}

func (s *stdHandlers) Init(config *model.AppConfig) {
	s.dataSourceHandler.Register(s.backendEventHandler)
	s.userHandler.Register(s.backendEventHandler)
}

func NewStdHandler(backendEventHandler backend_event_handler.BackendEventHandler, backendProviderService service.BackendProviderService) StdHandlers {
	return &stdHandlers{
		backendEventHandler: backendEventHandler,
		dataSourceHandler:   &dataSourceHandler{backendProviderService: backendProviderService},
		userHandler:         &userHandler{},
	}
}
