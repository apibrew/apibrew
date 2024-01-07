package testing

import (
	"context"
	"github.com/apibrew/apibrew/pkg/errors"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/resources"
	"github.com/apibrew/apibrew/pkg/service"
	backend_event_handler "github.com/apibrew/apibrew/pkg/service/backend-event-handler"
	"github.com/apibrew/apibrew/pkg/util"
	"google.golang.org/protobuf/types/known/structpb"
	"log"
)

type module struct {
	container           service.Container
	backendEventHandler backend_event_handler.BackendEventHandler
}

func (m module) Init() {
	m.ensureNamespace()
	m.ensureResources()
	m.initTestExecutionListeners()
}

func (m module) ensureNamespace() {
	_, err := m.container.GetRecordService().Apply(util.SystemContext, service.RecordUpdateParams{
		Namespace: resources.NamespaceResource.Namespace,
		Resource:  resources.NamespaceResource.Name,
		Records: []*model.Record{
			{
				Properties: map[string]*structpb.Value{
					"name": structpb.NewStringValue("testing"),
				},
			},
		},
	})

	if err != nil {
		log.Panic(err)
	}
}

func (m module) ensureResources() {
	var resources = []*model.Resource{TestSuiteResource, TestCaseResource, TestExecutionResource}

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

func (m module) initTestExecutionListeners() {
	m.backendEventHandler.RegisterHandler(backend_event_handler.Handler{
		Id:   "test-execution-listener",
		Name: "test-execution-listener",
		Fn:   m.handleTestExecution,
		Selector: &model.EventSelector{
			Actions: []model.Event_Action{
				model.Event_CREATE,
			},
			Namespaces: []string{TestExecutionResource.Namespace},
			Resources:  []string{TestExecutionResource.Name},
		},
		Order:    90,
		Sync:     true,
		Internal: true,
	})
}

func (m module) handleTestExecution(ctx context.Context, event *model.Event) (*model.Event, errors.ServiceError) {
	var records []*model.Record
	for _, record := range event.Records {
		err, store := m.executeTest(ctx, record)

		if err != nil {
			return nil, err
		}

		if store {
			records = append(records, record)
		}
	}

	if len(records) > 0 {
		event.Records = records
		return event, nil
	} else {
		return nil, nil
	}

}

func (m module) executeTest(ctx context.Context, record *model.Record) (errors.ServiceError, bool) {
	// locating records
	// locating test execution

	err := m.container.GetRecordService().ResolveReferences(ctx, TestExecutionResource, []*model.Record{record}, []string{"$.testCase"})

	if err != nil {
		return err, false
	}

	testExecution := TestExecutionMapperInstance.FromRecord(record)

	// executing test
	return m.executeTestCase(ctx, testExecution), testExecution.Stored
}

func (m module) executeTestCase(ctx context.Context, execution *TestExecution) errors.ServiceError {
	return nil
}

func NewModule(container service.Container) service.Module {
	backendEventHandler := container.GetBackendEventHandler().(backend_event_handler.BackendEventHandler)
	return &module{container: container, backendEventHandler: backendEventHandler}
}
