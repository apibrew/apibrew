package nano

import (
	"github.com/apibrew/apibrew/pkg/errors"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/resources"
	"github.com/apibrew/apibrew/pkg/service"
	"github.com/apibrew/apibrew/pkg/util"
	"google.golang.org/protobuf/types/known/structpb"
	"log"
)

type module struct {
	container    service.Container
	codeExecutor *codeExecutorService
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

		m.codeExecutor.registerCode(code)
	}
}

func NewModule(container service.Container) service.Module {
	return &module{container: container, codeExecutor: newCodeExecutorService()}
}
