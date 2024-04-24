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
	log "github.com/sirupsen/logrus"
)

type extensionService struct {
	recordService       service.RecordService
	ServiceName         string
	extensionVersionMap map[string]uint32
	extensionHandlerMap map[string]*backend_event_handler.Handler
	externalService     service.ExternalService
	backendEventHandler backend_event_handler.BackendEventHandler
}

func (d *extensionService) Reload() {
	d.runConfigureExtensions()
}

func (d *extensionService) Init(config *model.AppConfig) {
	d.runConfigureExtensions()
}

func (d *extensionService) runConfigureExtensions() {
	log.Trace("Start reconfiguring extension services")

	records, _, err := d.recordService.List(util.WithSystemContext(context.TODO()), service.RecordListParams{
		Namespace: resources.ExtensionResource.Namespace,
		Resource:  resources.ExtensionResource.Name,
		Limit:     10000,
	})

	if err != nil {
		panic(err)
	}

	var extensions = util.ArrayMap(records, resource_model.ExtensionMapperInstance.FromRecord)

	for _, ext := range extensions {
		log.Tracef("Configure extension: %v", ext)
		d.configureExtension(ext)
	}

	log.Trace("Finish reconfiguring extension services")
}

func (d *extensionService) RegisterExtension(extension *resource_model.Extension) {
	d.configureExtension(extension)
}

func (d *extensionService) UnRegisterExtension(extension *resource_model.Extension) {
	if d.extensionHandlerMap[extension.Id.String()] == nil {
		log.Warn("Trying to unregister extension that is not registered")
		return
	}

	d.backendEventHandler.UnRegisterHandler(*d.extensionHandlerMap[extension.Id.String()])

	delete(d.extensionHandlerMap, extension.Id.String())
	delete(d.extensionVersionMap, extension.Id.String())
}

func (d *extensionService) configureExtension(extension *resource_model.Extension) {
	if d.extensionVersionMap[extension.Id.String()] == uint32(extension.Version) {
		return
	}

	d.extensionVersionMap[extension.Id.String()] = uint32(extension.Version)

	extensionHandler := d.prepareExtensionHandler(extension)
	if d.extensionHandlerMap[extension.Id.String()] != nil {
		d.backendEventHandler.UnRegisterHandler(*d.extensionHandlerMap[extension.Id.String()])
	}
	d.extensionHandlerMap[extension.Id.String()] = &extensionHandler

	d.backendEventHandler.RegisterHandler(*d.extensionHandlerMap[extension.Id.String()])
}

func (d *extensionService) prepareExtensionHandler(extension *resource_model.Extension) backend_event_handler.Handler {
	return backend_event_handler.Handler{
		Id:   "extension-" + extension.Id.String(),
		Name: "extension-" + extension.Name,
		Selector: &model.EventSelector{
			Actions: util.ArrayMap(extension.Selector.Actions, func(s resource_model.EventAction) model.Event_Action {
				return model.Event_Action(model.Event_Action_value[string(s)])
			}),
			RecordSelector: nil,
			Namespaces:     extension.Selector.Namespaces,
			Resources:      extension.Selector.Resources,
			Ids:            extension.Selector.Ids,
			Annotations:    extension.Selector.Annotations,
		},
		Order:     int(extension.Order),
		Finalizes: extension.Finalizes,
		Sync:      extension.Sync,
		Responds:  extension.Responds,
		Fn: func(ctx context.Context, event *model.Event) (*model.Event, errors.ServiceError) {
			if event.Annotations == nil {
				event.Annotations = make(map[string]string)
			}

			event.Annotations[annotations.ExtensionId] = extension.Id.String()

			return d.externalService.Call(ctx, extension.Call, event)
		},
	}
}

func NewExtensionService(recordService service.RecordService, backendEventHandler backend_event_handler.BackendEventHandler, externalService service.ExternalService) service.ExtensionService {
	return &extensionService{
		ServiceName:         "ExtensionService",
		extensionVersionMap: make(map[string]uint32),
		extensionHandlerMap: make(map[string]*backend_event_handler.Handler),
		recordService:       recordService,
		backendEventHandler: backendEventHandler,
		externalService:     externalService,
	}
}
