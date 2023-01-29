package handlers

import (
	"github.com/tislib/data-handler/model"
	"github.com/tislib/data-handler/service"
	"github.com/tislib/data-handler/service/handler"
	"github.com/tislib/data-handler/service/system"
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
	s.genericHandler.RegisterWithSelector(s.dataSourceHandler.prepareHandler(), handler.ResourceSelector(system.DataSourceResource))
	s.genericHandler.RegisterWithSelector(s.userHandler.prepareHandler(), handler.ResourceSelector(system.UserResource))
}

func NewStdHandler(genericHandler *handler.GenericHandler, dataSourceService service.DataSourceService, userService service.UserService, recordService service.RecordService) StdHandler {
	return &stdHandler{
		genericHandler:    genericHandler,
		dataSourceHandler: &dataSourceHandler{dataSourceService: dataSourceService},
		userHandler:       &userHandler{userService: userService, recordService: recordService},
	}
}
