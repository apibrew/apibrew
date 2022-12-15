package handlers

import (
	"data-handler/model"
	"data-handler/service"
	"data-handler/service/handler"
)

type StdHandler interface {
	Init(data *model.InitData)
}

type stdHandler struct {
	genericHandler    *handler.GenericHandler
	dataSourceHandler *dataSourceHandler
}

func (s *stdHandler) Init(data *model.InitData) {
	s.genericHandler.Register(s.dataSourceHandler.prepareHandler())
}

func NewStdHandler(genericHandler *handler.GenericHandler, dataSourceService service.DataSourceService) StdHandler {
	return &stdHandler{
		genericHandler:    genericHandler,
		dataSourceHandler: &dataSourceHandler{dataSourceService: dataSourceService},
	}
}
