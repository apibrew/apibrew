package handlers

import (
	"github.com/tislib/apibrew/pkg/model"
	backend_event_handler "github.com/tislib/apibrew/pkg/service/backend-event-handler"
)

type StdHandlers interface {
	Init(data *model.InitData)
}

type stdHandlers struct {
	backendEventHandler backend_event_handler.BackendEventHandler
	dataSourceHandler   *dataSourceHandler
	userHandler         *userHandler
}

func (s *stdHandlers) Init(data *model.InitData) {
	s.dataSourceHandler.Register(s.backendEventHandler)
	s.userHandler.Register(s.backendEventHandler)
}

func NewStdHandler(backendEventHandler backend_event_handler.BackendEventHandler) StdHandlers {
	return &stdHandlers{
		backendEventHandler: backendEventHandler,
		dataSourceHandler:   &dataSourceHandler{},
		userHandler:         &userHandler{},
	}
}
