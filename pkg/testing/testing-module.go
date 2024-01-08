package testing

import (
	"context"
	"fmt"
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
	testExecution.Logs = util.Pointer("")

	// executing test
	return m.executeTestCase(ctx, testExecution), testExecution.Stored
}

func (m module) executeTestCase(ctx context.Context, execution *TestExecution) errors.ServiceError {
	m.log(execution, "Executing test case %s", execution.TestCase.Name)
	// executing steps
	for _, step := range execution.TestCase.Steps {
		err := m.executeStep(ctx, execution, step)

		if err != nil {
			return err
		}
	}
	m.log(execution, "Test case %s executed", execution.TestCase.Name)
	return nil
}

func (m module) executeStep(ctx context.Context, execution *TestExecution, step TestCaseTestCaseStep) errors.ServiceError {
	switch step.Operation {
	case TestCaseOperation_CREATE:
		return m.executeCreate(ctx, execution, step)
	case TestCaseOperation_UPDATE:
		return m.executeUpdate(ctx, execution, step)
	case TestCaseOperation_DELETE:
		return m.executeDelete(ctx, execution, step)
	case TestCaseOperation_GET:
		return m.executeGet(ctx, execution, step)
	case TestCaseOperation_LIST:
		return m.executeList(ctx, execution, step)
	case TestCaseOperation_APPLY:
		return m.executeApply(ctx, execution, step)
	case TestCaseOperation_NANO:
		return m.executeNano(ctx, execution, step)
	}

	return nil
}

func (m module) executeCreate(ctx context.Context, execution *TestExecution, step TestCaseTestCaseStep) errors.ServiceError {
	res, err := m.container.GetRecordService().Apply(ctx, service.RecordUpdateParams{
		Namespace: "",
		Resource:  "",
		Records:   []*model.Record{},
	})

	return err
}

func (m module) executeUpdate(ctx context.Context, execution *TestExecution, step TestCaseTestCaseStep) errors.ServiceError {

}

func (m module) executeDelete(ctx context.Context, execution *TestExecution, step TestCaseTestCaseStep) errors.ServiceError {

}

func (m module) executeGet(ctx context.Context, execution *TestExecution, step TestCaseTestCaseStep) errors.ServiceError {

}

func (m module) executeList(ctx context.Context, execution *TestExecution, step TestCaseTestCaseStep) errors.ServiceError {

}

func (m module) executeApply(ctx context.Context, execution *TestExecution, step TestCaseTestCaseStep) errors.ServiceError {

}

func (m module) executeNano(ctx context.Context, execution *TestExecution, step TestCaseTestCaseStep) errors.ServiceError {

}

func (m module) log(execution *TestExecution, args ...interface{}) {
	logStr := fmt.Sprint(args...)

	log.Printf("[TESTING] %s: %s", execution.TestCase.Name, logStr)

	execution.Logs = util.Pointer(fmt.Sprintf("%s\n%s", *execution.Logs, logStr))
}

func NewModule(container service.Container) service.Module {
	backendEventHandler := container.GetBackendEventHandler().(backend_event_handler.BackendEventHandler)
	return &module{container: container, backendEventHandler: backendEventHandler}
}
