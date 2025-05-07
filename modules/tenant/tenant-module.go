package pkg

import (
	"context"
	errors2 "errors"
	"github.com/apibrew/apibrew/modules/common"
	model2 "github.com/apibrew/apibrew/modules/tenant/model"
	"github.com/apibrew/apibrew/pkg/api"
	"github.com/apibrew/apibrew/pkg/errors"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/resources"
	"github.com/apibrew/apibrew/pkg/service"
	"github.com/apibrew/apibrew/pkg/service/backend-event-handler"
	"github.com/apibrew/apibrew/pkg/util"
	"google.golang.org/protobuf/types/known/structpb"
	"log"
)

type module struct {
	container           service.Container
	backendEventHandler backend_event_handler.BackendEventHandler
	api                 api.Interface
	disabled            bool
	options             map[string]string
	serviceId           string
}

func (m module) Init() {
	if m.disabled {
		return
	}

	m.ensureResources()

	if err := common.RegisterResourceProcessor[*model2.Tenant](
		"tenant-listener",
		&tenantProcessor{
			backendEventHandler: m.backendEventHandler,
			resourceService:     m.container.GetResourceService(),
			api:                 api.NewInterface(m.container),
		},
		m.backendEventHandler,
		m.container,
		model2.TenantResource,
	); err != nil {
		log.Fatal(err)
	}

	m.registerTenantListener()

	log.Println("Initialized module tenant")
}

func (m module) registerTenantListener() {
	m.backendEventHandler.RegisterHandler(backend_event_handler.Handler{
		Id:   "metrics-listener",
		Name: "metrics-listener",
		Fn: func(ctx context.Context, event *model.Event) (*model.Event, error) {
			tenant := GetTenant(ctx)

			if tenant != "" {
				event.Resource.SourceConfig.Catalog = tenant

				if event.Resource.Namespace == resources.ResourceResource.Namespace && event.Resource.Name == resources.ResourceResource.Name {
					for _, record := range event.Records {
						record.Properties["catalog"] = structpb.NewStringValue(tenant)
					}
				}

				return event, nil
			}
			// end metrics
			return event, nil
		},
		Order:    1,
		Sync:     true,
		Internal: true,
	})
}

func (m module) ensureResources() {
	var list = []*model.Resource{
		model2.TenantResource,
	}

	for _, resource := range list {
		existingResource, err := m.container.GetResourceService().GetResourceByName(util.SystemContext, resource.Namespace, resource.Name)

		if err == nil {
			resource.Id = existingResource.Id
			err = m.container.GetResourceService().Update(util.SystemContext, resource, true, true)

			if err != nil {
				log.Fatal(err)
			}
		} else if errors2.Is(err, errors.ResourceNotFoundError) {
			_, err = m.container.GetResourceService().Create(util.SystemContext, resource, true, true)

			if err != nil {
				log.Fatal(err)
			}
		} else if err != nil {
			log.Fatal(err)
		}
	}
}

func NewModule(container service.Container) service.Module {
	a := api.NewInterface(container)

	var config = container.GetAppConfig().Modules["tenant"]

	if config == nil {
		config = &model.ModuleConfig{
			Disabled: true,
		}
	}

	backendEventHandler := container.GetBackendEventHandler().(backend_event_handler.BackendEventHandler)
	return &module{container: container,
		api:                 a,
		disabled:            config.Disabled,
		options:             config.Options,
		serviceId:           container.GetAppConfig().ServiceId,
		backendEventHandler: backendEventHandler}
}
