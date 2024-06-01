package handlers

import (
	"context"
	"github.com/apibrew/apibrew/pkg/core"
	"github.com/apibrew/apibrew/pkg/resources"
	"github.com/apibrew/apibrew/pkg/service"
	"github.com/apibrew/apibrew/pkg/service/backend-event-handler"
	"github.com/apibrew/apibrew/pkg/util"
)

type dataSourceHandler struct {
	backendProviderService service.BackendProviderService
}

func (h *dataSourceHandler) Register(eventHandler backend_event_handler.BackendEventHandler) {
	eventHandler.RegisterHandler(prepareStdHandler(101, core.Event_UPDATE, h.AfterUpdate, resources.DataSourceResource))
	eventHandler.RegisterHandler(prepareStdHandler(101, core.Event_DELETE, h.AfterDelete, resources.DataSourceResource))
}

func (h *dataSourceHandler) AfterUpdate(ctx context.Context, event *core.Event) (*core.Event, error) {
	for _, dataSource := range event.Records {
		err := h.backendProviderService.DestroyDataSource(ctx, util.GetRecordId(dataSource))

		if err != nil {
			return nil, err
		}
	}

	return event, nil
}

func (h *dataSourceHandler) AfterDelete(ctx context.Context, event *core.Event) (*core.Event, error) {
	for _, record := range event.Records {
		err := h.backendProviderService.DestroyDataSource(ctx, util.GetRecordId(record))

		if err != nil {
			return nil, err
		}
	}

	return event, nil
}
