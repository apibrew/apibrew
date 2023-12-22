package nano

import (
	"context"
	"fmt"
	"github.com/apibrew/apibrew/pkg/errors"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/resources"
	"github.com/apibrew/apibrew/pkg/service"
	backend_event_handler "github.com/apibrew/apibrew/pkg/service/backend-event-handler"
	"github.com/apibrew/apibrew/pkg/util"
	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/types/known/structpb"
	"log"
)

type module struct {
	container           service.Container
	codeExecutor        *codeExecutorService
	backendEventHandler backend_event_handler.BackendEventHandler
}

func (m module) Init() {
	m.ensureNamespace()
	m.ensureResources()
	m.initCodeListeners()
	m.initExistingCodes()
}

func (m module) ensureNamespace() {
	_, err := m.container.GetRecordService().Apply(util.SystemContext, service.RecordUpdateParams{
		Namespace: resources.NamespaceResource.Namespace,
		Resource:  resources.NamespaceResource.Name,
		Records: []*model.Record{
			{
				Properties: map[string]*structpb.Value{
					"name": structpb.NewStringValue("nano"),
				},
			},
		},
	})

	if err != nil {
		log.Panic(err)
	}
}

func (m module) ensureResources() {
	var resources = []*model.Resource{CodeResource}

	for _, resource := range resources {
		existingResource, err := m.container.GetResourceService().GetResourceByName(util.SystemContext, resource.Namespace, resource.Name)

		if err == nil {
			resource.Id = existingResource.Id
			err = m.container.GetResourceService().Update(util.SystemContext, resource, true, true)

			if err != nil {
				log.Panic(err)
			}
		} else if err.Is(errors.ResourceNotFoundError) {
			_, err = m.container.GetResourceService().Create(util.SystemContext, resource, true, true)

			if err != nil {
				log.Panic(err)
			}
		} else if err != nil {
			log.Panic(err)
		}
	}
}

func (m module) initCodeListeners() {
	m.backendEventHandler.RegisterHandler(backend_event_handler.Handler{
		Id:   "nano-code-listener",
		Name: "nano-code-listener",
		Fn:   m.codeListenerHandler,
		Selector: &model.EventSelector{
			Actions: []model.Event_Action{
				model.Event_CREATE, model.Event_UPDATE, model.Event_DELETE,
			},
			Namespaces: []string{CodeResource.Namespace},
			Resources:  []string{CodeResource.Name},
		},
		Order:    90,
		Sync:     true,
		Internal: true,
	})
}

func (m module) initExistingCodes() {
	var codeRecords, _, err = m.container.GetRecordService().List(util.SystemContext, service.RecordListParams{
		Namespace: CodeResource.Namespace,
		Resource:  CodeResource.Name,
		Limit:     1000000,
	})

	if err != nil {
		log.Panic(err)
	}

	for _, codeRecord := range codeRecords {
		var code = CodeMapperInstance.FromRecord(codeRecord)

		err := m.codeExecutor.registerCode(code)

		if err != nil {
			logrus.WithField("CodeName", code.Name).Error(err)
		}
	}
}

func (m module) codeListenerHandler(ctx context.Context, event *model.Event) (*model.Event, errors.ServiceError) {
	for _, record := range event.Records {
		code := CodeMapperInstance.FromRecord(record)

		switch event.Action {
		case model.Event_CREATE:
			err := m.codeExecutor.registerCode(code)

			if err != nil {
				return nil, errors.RecordValidationError.WithMessage(fmt.Sprintf("%v", err))
			}
		case model.Event_UPDATE:
			err := m.codeExecutor.updateCode(code)

			if err != nil {
				return nil, errors.RecordValidationError.WithMessage(fmt.Sprintf("%v", err))
			}
		case model.Event_DELETE:
			err := m.codeExecutor.unRegisterCode(code)

			if err != nil {
				return nil, errors.RecordValidationError.WithMessage(fmt.Sprintf("%v", err))
			}
		}
	}

	return event, nil
}

func NewModule(container service.Container) service.Module {
	backendEventHandler := container.GetBackendEventHandler().(backend_event_handler.BackendEventHandler)
	return &module{container: container, codeExecutor: newCodeExecutorService(container, backendEventHandler), backendEventHandler: backendEventHandler}
}
