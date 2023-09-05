package client

import (
	"context"
	model2 "github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/resource_model"
	"strconv"
)

type RecordProcessFunc[Entity interface{}] func(ctx context.Context, instance Entity) (Entity, error)

type Handler[Entity interface{}] interface {
	Name(string) Handler[Entity]
	Before() Handler[Entity]
	After() Handler[Entity]
	Instead() Handler[Entity]
	Create(RecordProcessFunc[Entity]) Handler[Entity]
	Update(RecordProcessFunc[Entity]) Handler[Entity]
	Delete(RecordProcessFunc[Entity]) Handler[Entity]
}

type handler[Entity interface{}] struct {
	ext           Extension
	repository    Repository[Entity]
	dhClient      DhClient
	extensionRepo Repository[*resource_model.Extension]
	action        resource_model.EventAction
	order         int32
	finalizes     bool
	sync          bool
	name          string
	responds      bool
}

func (h handler[Entity]) Name(name string) Handler[Entity] {
	h.name = name
	return h
}

func (h handler[Entity]) Before() Handler[Entity] {
	h.order = 90
	return h
}

func (h handler[Entity]) After() Handler[Entity] {
	h.order = 110
	return h
}

func (h handler[Entity]) Instead() Handler[Entity] {
	h.order = 100
	h.finalizes = true
	return h
}

func (h handler[Entity]) handle(processFunc RecordProcessFunc[Entity]) {
	h.prepareExtension()

	h.ext.RegisterFunction(h.prepareName(), h.prepareProcessFunc(processFunc))
}

func (h handler[Entity]) prepareProcessFunc(processFunc RecordProcessFunc[Entity]) func(ctx context.Context, req *model2.Event) (*model2.Event, error) {
	return func(ctx context.Context, req *model2.Event) (*model2.Event, error) {
		processedRecords := make([]*model2.Record, len(req.Records))

		for i, record := range req.Records {
			processedRecord, err := processFunc(ctx, h.repository.Mapper().FromRecord(record))

			if err != nil {
				return nil, err
			}

			processedRecords[i] = h.repository.Mapper().ToRecord(processedRecord)
		}

		req.Records = processedRecords

		return req, nil
	}
}

func (h handler[Entity]) Create(processFunc RecordProcessFunc[Entity]) Handler[Entity] {
	h.action = resource_model.EventAction_CREATE

	h.handle(processFunc)

	return h
}

func (h handler[Entity]) Update(processFunc RecordProcessFunc[Entity]) Handler[Entity] {
	h.action = resource_model.EventAction_UPDATE

	h.handle(processFunc)

	return h
}

func (h handler[Entity]) Delete(processFunc RecordProcessFunc[Entity]) Handler[Entity] {
	h.action = resource_model.EventAction_DELETE

	h.handle(processFunc)

	return h
}

func (h handler[Entity]) prepareExtension() {
	ri := h.repository.Mapper().ResourceIdentity()

	h.name = h.prepareName()

	extension := &resource_model.Extension{
		Name:        h.name,
		Description: nil,
		Selector: &resource_model.ExtensionEventSelector{
			Actions:    []resource_model.EventAction{h.action},
			Namespaces: []string{ri.Namespace},
			Resources:  []string{ri.Name},
		},
		Order:     h.order,
		Finalizes: h.finalizes,
		Sync:      h.sync,
		Responds:  h.responds,
		Call: resource_model.ExtensionExternalCall{
			FunctionCall: &resource_model.ExtensionFunctionCall{
				Host:         h.ext.GetRemoteHost(),
				FunctionName: h.name,
			},
		},
	}

	_, err := h.extensionRepo.Apply(context.TODO(), extension)

	if err != nil {
		panic(err)
	}
}

func (h handler[Entity]) prepareName() string {
	ri := h.repository.Mapper().ResourceIdentity()

	if h.name == "" {
		h.name = "def_" + ri.Namespace + "_" + ri.Name + "_" + string(h.action) + "_" + strconv.Itoa(int(h.order))
	}

	return h.name
}

func NewHandler[Entity interface{}](dhClient DhClient, ext Extension, repository Repository[Entity]) Handler[Entity] {
	extensionRepo := R[*resource_model.Extension](dhClient, resource_model.ExtensionMapperInstance)

	return handler[Entity]{
		dhClient:      dhClient,
		ext:           ext,
		sync:          true,
		repository:    repository,
		extensionRepo: extensionRepo,
		responds:      true,
		order:         100,
	}
}
