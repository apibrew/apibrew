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
	"github.com/apibrew/apibrew/pkg/util"
	jwt_model "github.com/apibrew/apibrew/pkg/util/jwt-model"
	log "github.com/sirupsen/logrus"
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
	log.Debug("Handled by audit-handler")

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

	ctx = annotations.SetWithContext(ctx, annotations.BypassExtensions, annotations.Enabled)

	if event.Records != nil && len(event.Records) > 0 {
		for _, record := range event.Records {
			if record.Properties["id"] == nil {
				log.Warnf("Audit log cannot be created for record %s as it does not have an id", record.Id)
				continue
			}
			auditLog.RecordId = record.Properties["id"].GetStringValue()

			_, err := a.recordService.Create(util.WithSystemContext(ctx), service.RecordCreateParams{
				Namespace: resources.AuditLogResource.Namespace,
				Resource:  resources.AuditLogResource.Name,
				Records:   []*model.Record{resource_model.AuditLogMapperInstance.ToRecord(auditLog)},
			})

			if err != nil {
				log.Error(err)
				return nil, err
			}
			log.Debugf("Audit log created for record %s", record.Id)
		}
	} else if event.Ids != nil && len(event.Ids) > 0 {
		for _, recordId := range event.Ids {
			auditLog.RecordId = recordId

			_, err := a.recordService.Create(util.WithSystemContext(ctx), service.RecordCreateParams{
				Namespace: resources.AuditLogResource.Namespace,
				Resource:  resources.AuditLogResource.Name,
				Records:   []*model.Record{resource_model.AuditLogMapperInstance.ToRecord(auditLog)},
			})

			if err != nil {
				log.Error(err)
				return nil, err
			}
			log.Debugf("Audit log created for record %s", recordId)
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
