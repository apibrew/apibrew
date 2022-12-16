package service

import (
	"context"
	"data-handler/service/handler"
	"data-handler/service/params"
)

type WatchService interface {
	Watch(ctx context.Context, params params.WatchParams)
}

type watchService struct {
	genericHandler *handler.GenericHandler
}

func (w watchService) Watch(ctx context.Context, params params.WatchParams) <- model. {
}

func NewWatchService(genericHandler *handler.GenericHandler) WatchService {
	return &watchService{genericHandler: genericHandler}
}
