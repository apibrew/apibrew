package handlers

import (
	"github.com/tislib/data-handler/pkg/model"
	"github.com/tislib/data-handler/pkg/resources"
	"github.com/tislib/data-handler/pkg/service/handler"
)

type StdHandler interface {
	Init(data *model.InitData)
}

type stdHandler struct {
	genericHandler    *handler.GenericHandler
	dataSourceHandler *dataSourceHandler
	userHandler       *userHandler
}

func (s *stdHandler) Init(data *model.InitData) {
	s.genericHandler.RegisterWithSelector(s.dataSourceHandler.prepareHandler(), handler.ResourceSelector(resources.DataSourceResource))
	s.genericHandler.RegisterWithSelector(s.userHandler.prepareHandler(), handler.ResourceSelector(resources.UserResource))
}

func NewStdHandler(genericHandler *handler.GenericHandler) StdHandler {
	return &stdHandler{
		genericHandler:    genericHandler,
		dataSourceHandler: &dataSourceHandler{},
		userHandler:       &userHandler{},
	}
}
