package service

import (
	"context"
	"data-handler/model"
	"data-handler/service/errors"
	"data-handler/service/handler"
	"data-handler/service/params"
	"data-handler/util"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type WatchService interface {
	Watch(ctx context.Context, params params.WatchParams) <-chan *model.WatchMessage
}

type watchService struct {
	genericHandler *handler.GenericHandler
}

func (w watchService) Watch(ctx context.Context, p params.WatchParams) <-chan *model.WatchMessage {
	if p.BufferSize < 0 || p.BufferSize > 1000 {
		p.BufferSize = 100
	}

	out := make(chan *model.WatchMessage, p.BufferSize)
	watchHandler := &handler.BaseHandler{}

	//go func() {
	//	<-ctx.Done()
	//
	//	w.genericHandler.Unregister(watchHandler)
	//	close(out)
	//}()

	sendEvent := func(records []*model.Record, event model.EventType) {
		select {
		case out <- &model.WatchMessage{
			Changes: nil,
			RecordIds: util.ArrayMap[*model.Record, string](records, func(r *model.Record) string {
				return r.Id
			}),
			Event:   event,
			EventOn: timestamppb.New(time.Now()),
		}:
		default:
		}
	}

	watchHandler.AfterList = func(ctx context.Context, resource *model.Resource, params params.RecordListParams, records []*model.Record, total uint32) errors.ServiceError {
		sendEvent(records, model.EventType_LIST)
		return nil
	}

	watchHandler.AfterCreate = func(ctx context.Context, resource *model.Resource, params params.RecordCreateParams, records []*model.Record) errors.ServiceError {
		sendEvent(records, model.EventType_CREATE)

		return nil
	}

	watchHandler.AfterGet = func(ctx context.Context, resource *model.Resource, id string, res *model.Record) errors.ServiceError {
		sendEvent([]*model.Record{res}, model.EventType_GET)

		return nil
	}

	w.genericHandler.Register(watchHandler)

	//watchHandler.AfterDelete = func(ctx context.Context, params params.RecordDeleteParams) errors.ServiceError {
	//	sendEvent([]*model.Record{params.}, model.EventType_CREATE)
	//
	//	return nil
	//}

	return out
}

func NewWatchService(genericHandler *handler.GenericHandler) WatchService {
	return &watchService{genericHandler: genericHandler}
}
