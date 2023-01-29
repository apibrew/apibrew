package service

import (
	"context"
	log "github.com/sirupsen/logrus"
	"github.com/tislib/data-handler/pkg/errors"
	"github.com/tislib/data-handler/pkg/model"
	"github.com/tislib/data-handler/pkg/service/mapping"
	"github.com/tislib/data-handler/pkg/service/params"
	"github.com/tislib/data-handler/pkg/service/security"
	"github.com/tislib/data-handler/pkg/system"
)

type NamespaceService interface {
	InjectRecordService(service RecordService)
	InjectResourceService(service ResourceService)
	Init(data *model.InitData)
	Create(ctx context.Context, namespaces []*model.Namespace) ([]*model.Namespace, errors.ServiceError)
	Update(ctx context.Context, namespaces []*model.Namespace) ([]*model.Namespace, errors.ServiceError)
	Delete(ctx context.Context, ids []string) errors.ServiceError
	Get(ctx context.Context, id string) (*model.Namespace, errors.ServiceError)
	List(ctx context.Context) ([]*model.Namespace, errors.ServiceError)
	InjectBackendProviderService(service BackendProviderService)
}

type namespaceService struct {
	recordService          RecordService
	serviceName            string
	resourceService        ResourceService
	backendProviderService BackendProviderService
}

func (u *namespaceService) InjectBackendProviderService(backendProviderService BackendProviderService) {
	u.backendProviderService = backendProviderService
}

func (u *namespaceService) InjectResourceService(service ResourceService) {
	u.resourceService = service
}

func (u *namespaceService) InjectRecordService(service RecordService) {
	u.recordService = service
}

func (u *namespaceService) Create(ctx context.Context, namespaces []*model.Namespace) ([]*model.Namespace, errors.ServiceError) {
	// insert records via resource service
	records := mapping.MapToRecord(namespaces, mapping.NamespaceToRecord)

	result, _, err := u.recordService.Create(ctx, params.RecordCreateParams{
		Namespace: system.NamespaceResource.Namespace,
		Records:   records,
	})

	if err != nil {
		return nil, err
	}

	return mapping.MapFromRecord(result, mapping.NamespaceFromRecord), nil
}

func (u *namespaceService) Update(ctx context.Context, namespaces []*model.Namespace) ([]*model.Namespace, errors.ServiceError) {
	// insert records via resource service
	records := mapping.MapToRecord(namespaces, mapping.NamespaceToRecord)

	result, err := u.recordService.Update(ctx, params.RecordUpdateParams{
		Namespace: system.NamespaceResource.Namespace,
		Records:   records,
	})

	if err != nil {
		return nil, err
	}

	return mapping.MapFromRecord(result, mapping.NamespaceFromRecord), nil
}

func (u *namespaceService) Delete(ctx context.Context, ids []string) errors.ServiceError {

	return u.recordService.Delete(ctx, params.RecordDeleteParams{
		Namespace: system.NamespaceResource.Namespace,
		Resource:  system.NamespaceResource.Name,
		Ids:       ids,
	})
}

func (u *namespaceService) Get(ctx context.Context, id string) (*model.Namespace, errors.ServiceError) {

	record, err := u.recordService.Get(ctx, params.RecordGetParams{
		Namespace: system.NamespaceResource.Namespace,
		Resource:  system.NamespaceResource.Name,
		Id:        id,
	})

	if err != nil {
		return nil, err
	}

	return mapping.NamespaceFromRecord(record), nil
}

func (u *namespaceService) List(ctx context.Context) ([]*model.Namespace, errors.ServiceError) {

	result, _, err := u.recordService.List(ctx, params.RecordListParams{
		Namespace: system.NamespaceResource.Namespace,
		Resource:  system.NamespaceResource.Name,
	})

	if err != nil {
		return nil, err
	}

	return mapping.MapFromRecord(result, mapping.NamespaceFromRecord), err
}

func (d *namespaceService) Init(data *model.InitData) {
	d.backendProviderService.MigrateResource(system.NamespaceResource, nil)

	if len(data.InitNamespaces) > 0 {
		_, _, err := d.recordService.Create(security.SystemContext, params.RecordCreateParams{
			Namespace:      system.NamespaceResource.Namespace,
			Records:        mapping.MapToRecord(data.InitNamespaces, mapping.NamespaceToRecord),
			IgnoreIfExists: true,
		})

		if err != nil {
			log.Error(err)
		}
	}
}

func NewNamespaceService() NamespaceService {
	return &namespaceService{
		serviceName: "NamespaceService",
	}
}
