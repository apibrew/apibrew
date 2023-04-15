package service

import (
	"github.com/tislib/data-handler/pkg/abs"
	"github.com/tislib/data-handler/pkg/errors"
	"github.com/tislib/data-handler/pkg/model"
	"github.com/tislib/data-handler/pkg/service/handler"
	"github.com/tislib/data-handler/pkg/util"
)

type recordService struct {
	ServiceName            string
	resourceService        abs.ResourceService
	genericHandler         *handler.GenericHandler
	backendServiceProvider abs.BackendProviderService
}

func (r *recordService) PrepareQuery(resource *model.Resource, queryMap map[string]interface{}) (*model.BooleanExpression, errors.ServiceError) {
	return util.PrepareQuery(resource, queryMap)
}

func (r *recordService) Init(data *model.InitData) {

}

func NewRecordService(resourceService abs.ResourceService, backendProviderService abs.BackendProviderService, genericHandler *handler.GenericHandler) abs.RecordService {
	return &recordService{
		ServiceName:            "RecordService",
		resourceService:        resourceService,
		backendServiceProvider: backendProviderService,
		genericHandler:         genericHandler,
	}
}
