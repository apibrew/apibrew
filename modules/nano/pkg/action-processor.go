package nano

import (
	"context"
	model2 "github.com/apibrew/apibrew/modules/nano/pkg/model"
	"github.com/apibrew/apibrew/pkg/api"
	"github.com/apibrew/apibrew/pkg/formats/unstructured"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/resource_model"
	"github.com/apibrew/apibrew/pkg/service/annotations"
	backend_event_handler "github.com/apibrew/apibrew/pkg/service/backend-event-handler"
	"github.com/apibrew/apibrew/pkg/util"
	"sync"
)

type actionProcessor struct {
	m                   sync.Mutex
	api                 api.Interface
	codeExecutor        *codeExecutorService
	backendEventHandler backend_event_handler.BackendEventHandler
}

func (f *actionProcessor) MapperTo(record *model.Record) *model2.Action {
	return model2.ActionMapperInstance.FromRecord(record)
}

func (f *actionProcessor) Register(ctx context.Context, entity *model2.Action) error {
	f.m.Lock()
	defer f.m.Unlock()

	resource := f.prepareResource(entity)

	entity.Resource = resource

	resourceUn := resource_model.ResourceMapperInstance.ToUnstructured(resource)

	_, err := f.api.Apply(ctx, resourceUn)

	if err != nil {
		return err
	}

	var inlineCode = f.prepareInlineCode(entity)

	err = f.codeExecutor.registerCode(ctx, inlineCode)

	if err != nil {
		return err
	}

	return nil
}

func (f *actionProcessor) prepareResource(entity *model2.Action) *resource_model.Resource {
	var resource = &resource_model.Resource{
		Name: entity.Name,
		Namespace: &resource_model.Namespace{
			Name: "actions",
		},
		Annotations: map[string]string{
			annotations.EnableAudit: annotations.Enabled,
			annotations.ActionApi:   annotations.Enabled,
		},
		Properties: map[string]resource_model.Property{
			"input": {
				Type: resource_model.ResourceType_OBJECT,
			},
			"output": {
				Type: resource_model.ResourceType_OBJECT,
			},
		},
	}

	if entity.RestPath != "" {
		resource.Annotations[annotations.OpenApiRestPath] = entity.RestPath
	}

	if entity.InputSchema != nil {
		resource.Types = append(resource.Types, resource_model.SubType{
			Name:       "Input",
			Properties: entity.InputSchema,
		})
		resource.Properties["input"] = resource_model.Property{
			Type:    resource_model.ResourceType_STRUCT,
			TypeRef: util.Pointer("Input"),
		}
	}
	if entity.OutputSchema != nil {
		resource.Types = append(resource.Types, resource_model.SubType{
			Name:       "Output",
			Properties: entity.InputSchema,
		})

		resource.Properties["output"] = resource_model.Property{
			Type:    resource_model.ResourceType_STRUCT,
			TypeRef: util.Pointer("Output"),
		}
	}

	return resource
}

func (f *actionProcessor) Update(ctx context.Context, entity *model2.Action) error {
	if err := f.UnRegister(ctx, entity); err != nil {
		return err
	}

	return f.Register(ctx, entity)
}

func (f *actionProcessor) UnRegister(ctx context.Context, entity *model2.Action) error {
	f.m.Lock()
	defer f.m.Unlock()

	var existing, err = f.api.Load(ctx, map[string]unstructured.Any{
		"type": "nano/Action",
		"id":   entity.Id.String(),
	}, api.LoadParams{})

	if err != nil {
		return err
	}

	// locate the action resource
	return f.codeExecutor.unRegisterCode(ctx, &model2.Code{
		Id:   entity.Id,
		Name: "Actions/" + existing["name"].(string),
	})
}

func (f *actionProcessor) prepareInlineCode(entity *model2.Action) *model2.Code {
	return &model2.Code{
		Id:   entity.Id,
		Name: "Actions/" + entity.Name,
		Content: `
const fn = ` + entity.Source + `

resource('` + entity.Resource.Namespace.Name + "/" + entity.Resource.Name + `').beforeCreate(record => {
	record.output = fn(record.input)
	
	return record
})
`,
		ContentFormat: model2.CodeContentFormat(entity.ContentFormat),
		Annotations:   entity.Annotations,
		Language:      model2.CodeLanguage(entity.Language),
	}
}
