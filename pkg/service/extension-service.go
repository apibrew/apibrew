package service

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/tislib/data-handler/pkg/abs"
	"github.com/tislib/data-handler/pkg/errors"
	"github.com/tislib/data-handler/pkg/ext"
	"github.com/tislib/data-handler/pkg/extension"
	"github.com/tislib/data-handler/pkg/logging"
	"github.com/tislib/data-handler/pkg/model"
	"github.com/tislib/data-handler/pkg/resources"
	mapping2 "github.com/tislib/data-handler/pkg/resources/mapping"
	"github.com/tislib/data-handler/pkg/service/handler"
	"github.com/tislib/data-handler/pkg/service/security"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"time"
)

type extensionService struct {
	recordService          abs.RecordService
	ServiceName            string
	backendProviderService abs.BackendProviderService
	extensionVersionMap    map[string]uint32
	genericHandler         *handler.GenericHandler
	extensionHandlerMap    map[abs.Extension]*handler.BaseHandler
}

func (d *extensionService) List(ctx context.Context) ([]*model.RemoteExtension, errors.ServiceError) {
	logger := log.WithFields(logging.CtxFields(ctx))
	logger.Debug("Begin data-source List")
	defer logger.Debug("End data-source List")

	result, _, err := d.recordService.List(ctx, abs.RecordListParams{
		Namespace: resources.ExtensionResource.Namespace,
		Resource:  resources.ExtensionResource.Name,
	})

	if err != nil {
		return nil, err
	}

	return mapping2.MapFromRecord(result, mapping2.ExtensionFromRecord), nil
}

func (d *extensionService) Create(ctx context.Context, extensions []*model.RemoteExtension) ([]*model.RemoteExtension, errors.ServiceError) {
	logger := log.WithFields(logging.CtxFields(ctx))
	logger.WithField("extensions", extensions).Debug("Begin data-source Create")
	defer logger.Debug("End data-source Create")

	// insert records via resource service
	records := mapping2.MapToRecord(extensions, mapping2.ExtensionToRecord)
	result, _, err := d.recordService.Create(ctx, abs.RecordCreateParams{
		Namespace: resources.ExtensionResource.Namespace,
		Records:   records,
	})

	if err != nil {
		return nil, err
	}

	return mapping2.MapFromRecord(result, mapping2.ExtensionFromRecord), nil
}

func (d *extensionService) Update(ctx context.Context, extensions []*model.RemoteExtension) ([]*model.RemoteExtension, errors.ServiceError) {
	logger := log.WithFields(logging.CtxFields(ctx))
	logger.WithField("extensions", extensions).Debug("Begin data-source Update")
	defer logger.Debug("End data-source Update")

	// insert records via resource service
	records := mapping2.MapToRecord(extensions, mapping2.ExtensionToRecord)
	result, err := d.recordService.Update(ctx, abs.RecordUpdateParams{
		Namespace: resources.ExtensionResource.Namespace,
		Records:   records,
	})

	if err != nil {
		return nil, err
	}

	return mapping2.MapFromRecord(result, mapping2.ExtensionFromRecord), nil
}

func (d *extensionService) Get(ctx context.Context, id string) (*model.RemoteExtension, errors.ServiceError) {
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

	return mapping2.ExtensionFromRecord(record), nil
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

func (d *extensionService) RegisterExtension(extension abs.Extension) {
	d.extensionHandlerMap[extension] = NewExtensionHandler(extension)
	d.genericHandler.RegisterWithSelector(d.extensionHandlerMap[extension], handler.ResourceSelector(&model.Resource{Namespace: extension.GetExtensionConfig().Namespace, Name: extension.GetExtensionConfig().Resource}))
}

func (d *extensionService) UnRegisterExtension(extension abs.Extension) {
	d.genericHandler.Unregister(d.extensionHandlerMap[extension])
	delete(d.extensionHandlerMap, extension)
}

func (d *extensionService) configureExtension(remoteExtension *model.RemoteExtension) {
	if d.extensionVersionMap[remoteExtension.Name] == remoteExtension.Version {
		return
	}

	d.extensionVersionMap[remoteExtension.Name] = remoteExtension.Version

	log.Debug("Start reconfiguring extension: " + remoteExtension.Name)
	var opts []grpc.DialOption

	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", remoteExtension.Server.Host, remoteExtension.Server.Port), opts...)
	if err != nil {
		panic(err)
	}

	client := ext.NewRecordExtensionServiceClient(conn)

	hdlr := NewExtensionHandler(extension.FromRecordExtensionServiceClient(client, remoteExtension.Config))

	d.genericHandler.RegisterWithSelector(hdlr, handler.ResourceSelector(&model.Resource{Namespace: remoteExtension.Config.Namespace, Name: remoteExtension.Config.Resource}))
}

func NewExtensionService(recordService abs.RecordService, backendProviderService abs.BackendProviderService, genericHandler *handler.GenericHandler) abs.ExtensionService {
	return &extensionService{
		ServiceName:            "ExtensionService",
		extensionVersionMap:    make(map[string]uint32),
		extensionHandlerMap:    make(map[abs.Extension]*handler.BaseHandler),
		recordService:          recordService,
		backendProviderService: backendProviderService,
		genericHandler:         genericHandler,
	}
}
