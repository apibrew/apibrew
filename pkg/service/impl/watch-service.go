package impl

import (
	"context"
	"github.com/apibrew/apibrew/pkg/errors"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/resource_model"
	"github.com/apibrew/apibrew/pkg/resources"
	"github.com/apibrew/apibrew/pkg/service"
	backend_event_handler "github.com/apibrew/apibrew/pkg/service/backend-event-handler"
	"github.com/apibrew/apibrew/pkg/util"
)

type watchService struct {
	backendEventHandler  backend_event_handler.BackendEventHandler
	authorizationService service.AuthorizationService
}

func (w watchService) Watch(ctx context.Context, p service.WatchParams) (<-chan *model.Event, errors.ServiceError) {
	if err := w.authorizationService.CheckRecordAccess(ctx, service.CheckRecordAccessParams{
		Resource:  resources.ExtensionResource,
		Operation: resource_model.PermissionOperation_FULL,
	}); err != nil {
		return nil, err
	}

	if p.BufferSize < 0 || p.BufferSize > 1000 {
		p.BufferSize = 100
	}

	out := make(chan *model.Event, p.BufferSize)
	watchHandlerId := util.RandomHex(6)
	watchHandler := backend_event_handler.Handler{
		Id:   "watch-handler-" + watchHandlerId,
		Name: "watch-handler-" + watchHandlerId,
		Fn: func(ctx context.Context, event *model.Event) (*model.Event, errors.ServiceError) {
			out <- event

			return event, nil
		},
		Selector: p.Selector,
		Order:    101,
	}

	go func() {
		<-ctx.Done()

		w.backendEventHandler.UnRegisterHandler(watchHandler)
		close(out)
	}()

	w.backendEventHandler.RegisterHandler(watchHandler)

	return out, nil
}

func NewWatchService(backendEventHandler backend_event_handler.BackendEventHandler, authorizationService service.AuthorizationService) service.WatchService {
	return &watchService{backendEventHandler: backendEventHandler, authorizationService: authorizationService}
}
