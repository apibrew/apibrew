package service

import (
	"context"
	"github.com/tislib/data-handler/pkg/errors"
	"github.com/tislib/data-handler/pkg/model"
	handler2 "github.com/tislib/data-handler/pkg/service/handler"
	params2 "github.com/tislib/data-handler/pkg/service/params"
	"github.com/tislib/data-handler/pkg/util"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type WatchService interface {
	Watch(ctx context.Context, params params2.WatchParams) <-chan *model.WatchMessage
}

type watchService struct {
	genericHandler *handler2.GenericHandler
}

func (w watchService) Watch(ctx context.Context, p params2.WatchParams) <-chan *model.WatchMessage {
	if p.BufferSize < 0 || p.BufferSize > 1000 {
		p.BufferSize = 100
	}

	out := make(chan *model.WatchMessage, p.BufferSize)
	watchHandler := &handler2.BaseHandler{}

	go func() {
		<-ctx.Done()

		w.genericHandler.Unregister(watchHandler)
		close(out)
	}()

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

	watchHandler.AfterList = func(ctx context.Context, resource *model.Resource, params params2.RecordListParams, records []*model.Record, total uint32) errors.ServiceError {
		sendEvent(records, model.EventType_LIST)
		return nil
	}

	watchHandler.AfterCreate = func(ctx context.Context, resource *model.Resource, params params2.RecordCreateParams, records []*model.Record) errors.ServiceError {
		sendEvent(records, model.EventType_CREATE)

		return nil
	}

	watchHandler.AfterGet = func(ctx context.Context, resource *model.Resource, id string, res *model.Record) errors.ServiceError {
		sendEvent([]*model.Record{res}, model.EventType_GET)

		return nil
	}

	//watchHandler.AfterDelete = func(ctx context.Context, params params.RecordDeleteParams) errors.ServiceError {
	//	sendEvent([]*model.Record{params}, model.EventType_CREATE)
	//
	//	return nil
	//}

	w.genericHandler.Register(watchHandler)

	return out
}

func NewWatchService(genericHandler *handler2.GenericHandler) WatchService {
	return &watchService{genericHandler: genericHandler}
}
