package impl

import (
	"context"
	"github.com/apibrew/apibrew/pkg/errors"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/resource_model"
	"github.com/apibrew/apibrew/pkg/resources"
	"github.com/apibrew/apibrew/pkg/service"
	"github.com/apibrew/apibrew/pkg/service/annotations"
	backend_event_handler "github.com/apibrew/apibrew/pkg/service/backend-event-handler"
	jwt_model "github.com/apibrew/apibrew/pkg/util/jwt-model"
	"time"
)

type auditService struct {
	backendEventHandler backend_event_handler.BackendEventHandler
	recordService       service.RecordService
}

func (a *auditService) Init(config *model.AppConfig) {
	a.backendEventHandler.RegisterHandler(a.prepareHandler())
}

func (a *auditService) prepareHandler() backend_event_handler.Handler {
	return backend_event_handler.Handler{
		Id:   "audit-handler",
		Name: "audit-handler",
		Fn:   a.handle,
		Selector: &model.EventSelector{
			Annotations: map[string]string{
				annotations.EnableAudit: annotations.Enabled,
			},
		},
		Order:    200,
		Sync:     false,
		Internal: true,
	}
}

func (a *auditService) handle(ctx context.Context, event *model.Event) (*model.Event, errors.ServiceError) {

	if event.Action == model.Event_GET || event.Action == model.Event_LIST {
		return event, nil
	}

	// prepare audit
	auditLog := &resource_model.AuditLog{
		Namespace:   event.Resource.Namespace,
		Resource:    event.Resource.Name,
		Time:        time.Now(),
		Operation:   resource_model.AuditLogOperation(event.Action.String()),
		Annotations: event.Annotations,
	}

	userDetails := jwt_model.GetUserDetailsFromContext(ctx)

	if userDetails != nil {
		auditLog.Username = userDetails.Username
	} else {
		auditLog.Username = "internal"
	}

	if event.Records != nil && len(event.Records) > 0 {
		for _, record := range event.Records {
			auditLog.RecordId = record.Id

			_, err := a.recordService.Create(ctx, service.RecordCreateParams{
				Namespace: resources.AuditLogResource.Namespace,
				Resource:  resources.AuditLogResource.Name,
				Records:   []*model.Record{resource_model.AuditLogMapperInstance.ToRecord(auditLog)},
			})

			if err != nil {
				return nil, err
			}
		}
	} else if event.Ids != nil && len(event.Ids) > 0 {
		for _, recordId := range event.Ids {
			auditLog.RecordId = recordId

			_, err := a.recordService.Create(ctx, service.RecordCreateParams{
				Namespace: resources.AuditLogResource.Namespace,
				Resource:  resources.AuditLogResource.Name,
				Records:   []*model.Record{resource_model.AuditLogMapperInstance.ToRecord(auditLog)},
			})

			if err != nil {
				return nil, err
			}
		}
	}

	return event, nil
}

func NewAuditService(backendEventHandler backend_event_handler.BackendEventHandler, recordService service.RecordService) service.AuditService {
	return &auditService{
		backendEventHandler,
		recordService,
	}
}
