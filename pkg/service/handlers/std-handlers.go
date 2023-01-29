package handlers

import (
	"github.com/tislib/data-handler/pkg/model"
	service2 "github.com/tislib/data-handler/pkg/service"
	"github.com/tislib/data-handler/pkg/service/handler"
	"github.com/tislib/data-handler/pkg/system"
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

func NewStdHandler(genericHandler *handler.GenericHandler, dataSourceService service2.DataSourceService, userService service2.UserService, recordService service2.RecordService) StdHandler {
	return &stdHandler{
		genericHandler:    genericHandler,
		dataSourceHandler: &dataSourceHandler{dataSourceService: dataSourceService},
		userHandler:       &userHandler{userService: userService, recordService: recordService},
	}
}
