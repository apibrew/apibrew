package service

import (
	"context"
	"github.com/apibrew/apibrew/pkg/abs"
	"github.com/apibrew/apibrew/pkg/errors"
	"github.com/apibrew/apibrew/pkg/logging"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/resources"
	"github.com/apibrew/apibrew/pkg/resources/mapping"
	backend_event_handler "github.com/apibrew/apibrew/pkg/service/backend-event-handler"
	"github.com/apibrew/apibrew/pkg/service/security"
	log "github.com/sirupsen/logrus"
	"time"
)

type extensionService struct {
	recordService          abs.RecordService
	ServiceName            string
	backendProviderService abs.BackendProviderService
	extensionVersionMap    map[string]uint32
	extensionHandlerMap    map[string]*backend_event_handler.Handler
	externalService        abs.ExternalService
	backendEventHandler    backend_event_handler.BackendEventHandler
}

func (d *extensionService) List(ctx context.Context) ([]*model.Extension, errors.ServiceError) {
	logger := log.WithFields(logging.CtxFields(ctx))
	logger.Trace("Begin extension List")
	defer logger.Trace("End extension List")

	result, _, err := d.recordService.List(ctx, abs.RecordListParams{
		Namespace: resources.ExtensionResource.Namespace,
		Resource:  resources.ExtensionResource.Name,
		Limit:     1000,
	})

	if err != nil {
		return nil, err
	}

	return mapping.MapFromRecord(result, mapping.ExtensionFromRecord), nil
}

func (d *extensionService) Create(ctx context.Context, extensions []*model.Extension) ([]*model.Extension, errors.ServiceError) {
	logger := log.WithFields(logging.CtxFields(ctx))
	logger.WithField("extensions", extensions).Debug("Begin data-source Create")

	// insert records via resource service
	records := mapping.MapToRecord(extensions, mapping.ExtensionToRecord)
	result, err := d.recordService.Create(ctx, abs.RecordCreateParams{
		Namespace: resources.ExtensionResource.Namespace,
		Resource:  resources.ExtensionResource.Name,
		Records:   records,
	})

	if err != nil {
		return nil, err
	}

	return mapping.MapFromRecord(result, mapping.ExtensionFromRecord), nil
}

func (d *extensionService) Update(ctx context.Context, extensions []*model.Extension) ([]*model.Extension, errors.ServiceError) {
	logger := log.WithFields(logging.CtxFields(ctx))
	logger.WithField("extensions", extensions).Debug("Begin data-source Update")
	defer logger.Debug("End data-source Update")

	// insert records via resource service
	records := mapping.MapToRecord(extensions, mapping.ExtensionToRecord)
	result, err := d.recordService.Update(ctx, abs.RecordUpdateParams{
		Namespace: resources.ExtensionResource.Namespace,
		Resource:  resources.ExtensionResource.Name,
		Records:   records,
	})

	if err != nil {
		return nil, err
	}

	return mapping.MapFromRecord(result, mapping.ExtensionFromRecord), nil
}

func (d *extensionService) Apply(ctx context.Context, extensions []*model.Extension) ([]*model.Extension, errors.ServiceError) {
	logger := log.WithFields(logging.CtxFields(ctx))
	logger.WithField("extensions", extensions).Debug("Begin data-source Update")
	defer logger.Debug("End data-source Update")

	// insert records via resource service
	records := mapping.MapToRecord(extensions, mapping.ExtensionToRecord)
	result, err := d.recordService.Apply(ctx, abs.RecordUpdateParams{
		Namespace: resources.ExtensionResource.Namespace,
		Resource:  resources.ExtensionResource.Name,
		Records:   records,
	})

	if err != nil {
		return nil, err
	}

	return mapping.MapFromRecord(result, mapping.ExtensionFromRecord), nil
}

func (d *extensionService) Get(ctx context.Context, id string) (*model.Extension, errors.ServiceError) {
	logger := log.WithFields(logging.CtxFields(ctx))
	logger.WithField("id", id).Debug("Begin data-source Get")
	defer logger.Debug("End data-source Get")

	record, err := d.recordService.Get(ctx, abs.RecordGetParams{
		Namespace: resources.ExtensionResource.Namespace,
		Resource:  resources.ExtensionResource.Name,
		Id:        id,
	})

	if err != nil {
		return nil, err
	}

	return mapping.ExtensionFromRecord(record), nil
}

func (d *extensionService) Delete(ctx context.Context, ids []string) errors.ServiceError {
	logger := log.WithFields(logging.CtxFields(ctx))
	logger.WithField("ids", ids).Debug("Begin data-source Delete")
	defer logger.Debug("End data-source Delete")

	return d.recordService.Delete(ctx, abs.RecordDeleteParams{
		Namespace: resources.ExtensionResource.Namespace,
		Resource:  resources.ExtensionResource.Name,
		Ids:       ids,
	})
}

func (d *extensionService) Init(data *model.InitData) {
	d.runConfigureExtensions()

	go d.keepExtensionsRunning()
}

func (d *extensionService) keepExtensionsRunning() {
	for {
		time.Sleep(10 * time.Second)

		d.runConfigureExtensions()
	}
}

func (d *extensionService) runConfigureExtensions() {
	log.Trace("Start reconfiguring extension services")

	extensions, err := d.List(security.WithSystemContext(context.TODO()))

	if err != nil {
		panic(err)
	}

	for _, ext := range extensions {
		log.Tracef("Configure extension: %v", ext)
		d.configureExtension(ext)
	}

	log.Trace("Finish reconfiguring extension services")
}

func (d *extensionService) RegisterExtension(extension *model.Extension) {
	d.configureExtension(extension)
}

func (d *extensionService) UnRegisterExtension(extension *model.Extension) {
	if d.extensionHandlerMap[extension.Id] == nil {
		log.Warn("Trying to unregister extension that is not registered")
		return
	}

	d.backendEventHandler.UnRegisterHandler(*d.extensionHandlerMap[extension.Id])

	delete(d.extensionHandlerMap, extension.Id)
	delete(d.extensionVersionMap, extension.Id)
}

func (d *extensionService) configureExtension(extension *model.Extension) {
	if d.extensionVersionMap[extension.Id] == extension.Version {
		return
	}

	d.extensionVersionMap[extension.Id] = extension.Version

	extensionHandler := d.prepareExtensionHandler(extension)
	if d.extensionHandlerMap[extension.Id] != nil {
		d.backendEventHandler.UnRegisterHandler(*d.extensionHandlerMap[extension.Id])
	}
	d.extensionHandlerMap[extension.Id] = &extensionHandler

	d.backendEventHandler.RegisterHandler(*d.extensionHandlerMap[extension.Id])
}

func (d *extensionService) prepareExtensionHandler(extension *model.Extension) backend_event_handler.Handler {
	return backend_event_handler.Handler{
		Id:        "extension-" + extension.Id,
		Name:      "extension-" + extension.Name,
		Selector:  extension.Selector,
		Order:     int(extension.Order),
		Finalizes: extension.Finalizes,
		Sync:      extension.Sync,
		Responds:  extension.Responds,
		Fn: func(ctx context.Context, event *model.Event) (*model.Event, errors.ServiceError) {
			return d.externalService.Call(ctx, extension.Call, event)
		},
	}
}

func NewExtensionService(recordService abs.RecordService, backendProviderService abs.BackendProviderService, backendEventHandler backend_event_handler.BackendEventHandler, externalService abs.ExternalService) abs.ExtensionService {
	return &extensionService{
		ServiceName:            "ExtensionService",
		extensionVersionMap:    make(map[string]uint32),
		extensionHandlerMap:    make(map[string]*backend_event_handler.Handler),
		recordService:          recordService,
		backendProviderService: backendProviderService,
		backendEventHandler:    backendEventHandler,
		externalService:        externalService,
	}
}
