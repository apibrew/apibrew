package pkg

import (
	"context"
	model2 "github.com/apibrew/apibrew/modules/tenant/model"
	"github.com/apibrew/apibrew/pkg/api"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/resources"
	"github.com/apibrew/apibrew/pkg/service"
	backend_event_handler "github.com/apibrew/apibrew/pkg/service/backend-event-handler"
	"google.golang.org/protobuf/proto"
	"sync"
)

type tenantProcessor struct {
	m                   sync.Mutex
	api                 api.Interface
	backendEventHandler backend_event_handler.BackendEventHandler
	resourceService     service.ResourceService
}

func (f *tenantProcessor) MapperTo(record *model.Record) *model2.Tenant {
	return model2.TenantMapperInstance.FromRecord(record)
}

func (f *tenantProcessor) Register(ctx context.Context, entity *model2.Tenant) error {
	f.m.Lock()
	defer f.m.Unlock()

	// setup root resources
	tenantCtx := WithTenant(ctx, *entity.Name)

	f.MigrateResource(tenantCtx, resources.NamespaceResource)
	f.MigrateResource(tenantCtx, resources.ResourceResource)
	f.MigrateResource(tenantCtx, resources.UserResource)
	f.MigrateResource(tenantCtx, resources.RoleResource)
	f.MigrateResource(tenantCtx, resources.PermissionResource)
	f.MigrateResource(tenantCtx, resources.ExtensionResource)
	f.MigrateResource(tenantCtx, resources.AuditLogResource)

	return nil
}

func (f *tenantProcessor) MigrateResource(ctx context.Context, resource *model.Resource) {
	resourceCopy := proto.Clone(resource).(*model.Resource)
	resourceCopy.SourceConfig.Catalog = GetTenant(ctx)

	f.resourceService.MigrateResource(ctx, resourceCopy)
}

func (f *tenantProcessor) Update(ctx context.Context, entity *model2.Tenant) error {
	if err := f.UnRegister(ctx, entity); err != nil {
		return err
	}

	return f.Register(ctx, entity)
}

func (f *tenantProcessor) UnRegister(ctx context.Context, entity *model2.Tenant) error {
	f.m.Lock()
	defer f.m.Unlock()

	return nil
}
