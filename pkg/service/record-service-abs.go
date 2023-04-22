package service

import (
	"github.com/tislib/apibrew/pkg/abs"
	"github.com/tislib/apibrew/pkg/errors"
	"github.com/tislib/apibrew/pkg/model"
	"github.com/tislib/apibrew/pkg/service/handler"
	"github.com/tislib/apibrew/pkg/util"
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
