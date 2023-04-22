package service

import (
	"context"
	"github.com/tislib/apibrew/pkg/abs"
	"github.com/tislib/apibrew/pkg/errors"
	"github.com/tislib/apibrew/pkg/model"
	handler2 "github.com/tislib/apibrew/pkg/service/handler"
	"github.com/tislib/apibrew/pkg/util"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type watchService struct {
	genericHandler *handler2.GenericHandler
}

func (w watchService) Watch(ctx context.Context, p abs.WatchParams) <-chan *model.WatchMessage {
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

	watchHandler.AfterList = func(ctx context.Context, resource *model.Resource, params abs.RecordListParams, records []*model.Record, total uint32) errors.ServiceError {
		sendEvent(records, model.EventType_LIST)
		return nil
	}

	watchHandler.AfterCreate = func(ctx context.Context, resource *model.Resource, params abs.RecordCreateParams, records []*model.Record) errors.ServiceError {
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

func NewWatchService(genericHandler *handler2.GenericHandler) abs.WatchService {
	return &watchService{genericHandler: genericHandler}
}
