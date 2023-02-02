package service

import (
	"context"
	log "github.com/sirupsen/logrus"
	"github.com/tislib/data-handler/pkg/errors"
	"github.com/tislib/data-handler/pkg/logging"
	"github.com/tislib/data-handler/pkg/model"
	"github.com/tislib/data-handler/pkg/service/handler"
	"github.com/tislib/data-handler/pkg/service/mapping"
	"github.com/tislib/data-handler/pkg/service/params"
	"github.com/tislib/data-handler/pkg/service/security"
	"github.com/tislib/data-handler/pkg/system"
	"time"
)

type ExtensionService interface {
	Init(*model.InitData)
	InjectRecordService(service RecordService)
	List(ctx context.Context) ([]*model.Extension, errors.ServiceError)
	Create(ctx context.Context, sources []*model.Extension) ([]*model.Extension, errors.ServiceError)
	Update(ctx context.Context, sources []*model.Extension) ([]*model.Extension, errors.ServiceError)
	Get(ctx context.Context, id string) (*model.Extension, errors.ServiceError)
	Delete(ctx context.Context, ids []string) errors.ServiceError
	InjectBackendProviderService(service BackendProviderService)
	InjectGenericHandler(handler *handler.GenericHandler)
}

type extensionService struct {
	recordService          RecordService
	ServiceName            string
	backendProviderService BackendProviderService
	extensionVersionMap    map[string]uint32
	genericHandler         *handler.GenericHandler
}

func (d *extensionService) InjectGenericHandler(genericHandler *handler.GenericHandler) {
	d.genericHandler = genericHandler
}

func (d *extensionService) InjectBackendProviderService(backendProviderService BackendProviderService) {
	d.backendProviderService = backendProviderService
}

func (d *extensionService) List(ctx context.Context) ([]*model.Extension, errors.ServiceError) {
	logger := log.WithFields(logging.CtxFields(ctx))
	logger.Debug("Begin data-source List")
	defer logger.Debug("End data-source List")

	result, _, err := d.recordService.List(ctx, params.RecordListParams{
		Namespace: system.ExtensionResource.Namespace,
		Resource:  system.ExtensionResource.Name,
	})

	if err != nil {
		return nil, err
	}

	return mapping.MapFromRecord(result, mapping.ExtensionFromRecord), nil
}

func (d *extensionService) Create(ctx context.Context, extensions []*model.Extension) ([]*model.Extension, errors.ServiceError) {
	logger := log.WithFields(logging.CtxFields(ctx))
	logger.WithField("extensions", extensions).Debug("Begin data-source Create")
	defer logger.Debug("End data-source Create")

	// insert records via resource service
	records := mapping.MapToRecord(extensions, mapping.ExtensionToRecord)
	result, _, err := d.recordService.Create(ctx, params.RecordCreateParams{
		Namespace: system.ExtensionResource.Namespace,
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
	result, err := d.recordService.Update(ctx, params.RecordUpdateParams{
		Namespace: system.ExtensionResource.Namespace,
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

	record, err := d.recordService.Get(ctx, params.RecordGetParams{
		Namespace: system.ExtensionResource.Namespace,
		Resource:  system.ExtensionResource.Name,
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

	return d.recordService.Delete(ctx, params.RecordDeleteParams{
		Namespace: system.ExtensionResource.Namespace,
		Resource:  system.ExtensionResource.Name,
		Ids:       ids,
	})
}

func (d *extensionService) InjectRecordService(service RecordService) {
	d.recordService = service
}

func (d *extensionService) Init(data *model.InitData) {
	d.backendProviderService.MigrateResource(system.ExtensionResource, nil)

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
	log.Debug("Start reconfiguring extension services")

	extensions, err := d.List(security.WithSystemContext(context.TODO()))

	if err != nil {
		panic(err)
	}

	for _, extension := range extensions {
		d.configureExtension(extension)
	}

	log.Debug("Finish reconfiguring extension services")
}

func (d *extensionService) configureExtension(extension *model.Extension) {
	if d.extensionVersionMap[extension.Name] == extension.Version {
		return
	}

	d.extensionVersionMap[extension.Name] = extension.Version

	log.Debug("Start reconfiguring extension: " + extension.Name)

	hdlr := NewExtensionHandler(extension)

	d.genericHandler.RegisterWithSelector(hdlr, handler.ResourceSelector(&model.Resource{Namespace: extension.Namespace, Name: extension.Resource}))
}

func NewExtensionService(recordService RecordService, backendProviderService BackendProviderService, genericHandler *handler.GenericHandler) ExtensionService {
	return &extensionService{
		ServiceName:            "ExtensionService",
		extensionVersionMap:    make(map[string]uint32),
		recordService:          recordService,
		backendProviderService: backendProviderService,
		genericHandler:         genericHandler,
	}
}
