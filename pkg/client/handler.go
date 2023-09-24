package client

import (
	"context"
	"github.com/apibrew/apibrew/pkg/abs"
	model2 "github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/resource_model"
	"github.com/apibrew/apibrew/pkg/service/annotations"
	log "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/types/known/structpb"
	"reflect"
	"strconv"
)

type RecordProcessFunc[Entity interface{}] func(ctx context.Context, instance Entity) (Entity, error)
type LambdaProcessFunc[Entity interface{}] func(ctx context.Context, instance Entity) error

type Handler[Entity interface{}] interface {
	Name(string) Handler[Entity]
	Before() Handler[Entity]
	PreProcess(RecordProcessFunc[Entity]) Handler[Entity]
	PostProcess(RecordProcessFunc[Entity]) Handler[Entity]
	Lambda(action string, processor LambdaProcessFunc[Entity]) Handler[Entity]
	Fire(ctx context.Context, action string, payload Entity) error
	After() Handler[Entity]
	Instead() Handler[Entity]
	Create(RecordProcessFunc[Entity]) Handler[Entity]
	Update(RecordProcessFunc[Entity]) Handler[Entity]
	Delete(RecordProcessFunc[Entity]) Handler[Entity]
}

type handler[Entity interface{}] struct {
	ext           Extension
	mapper        abs.EntityMapper[Entity]
	dhClient      Client
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

func (h handler[Entity]) Lambda(action string, processor LambdaProcessFunc[Entity]) Handler[Entity] {
	h.sync = false
	h.responds = false
	h.finalizes = true
	h.order = 1
	h.action = resource_model.EventAction_CREATE
	h.name = h.prepareName() + "_" + action
	h.prepareExtension()

	h.ext.RegisterFunction(h.prepareName(), h.prepareLambdaProcessFunc(action, processor))

	return h
}

func (h handler[Entity]) handle(processFunc RecordProcessFunc[Entity]) {
	h.name = h.prepareName()
	h.prepareExtension()

	h.ext.RegisterFunction(h.prepareName(), h.prepareProcessFunc(processFunc))
}

func (h handler[Entity]) prepareProcessFunc(processFunc RecordProcessFunc[Entity]) func(ctx context.Context, req *model2.Event) (*model2.Event, error) {
	return func(ctx context.Context, req *model2.Event) (*model2.Event, error) {
		processedRecords := make([]*model2.Record, len(req.Records))

		for i, record := range req.Records {
			processedRecord, err := processFunc(ctx, h.mapper.FromRecord(record))

			if err != nil {
				return nil, err
			}

			processedRecords[i] = h.mapper.ToRecord(processedRecord)
		}

		req.Records = processedRecords

		return req, nil
	}
}

func (h handler[Entity]) Fire(ctx context.Context, action string, payload Entity) error {
	if reflect.ValueOf(payload).IsZero() {
		payload = h.mapper.New()
	}

	rec := h.mapper.ToRecord(payload)
	ri := h.mapper.ResourceIdentity()

	rec.Properties["action"] = structpb.NewStringValue(action)

	_, err := h.dhClient.CreateRecord(ctx, ri.Namespace, ri.Name, rec)

	if err != nil {
		log.Error("Error while firing event: ", err)
	}

	return err
}

func (h handler[Entity]) prepareLambdaProcessFunc(action string, processFunc LambdaProcessFunc[Entity]) func(ctx context.Context, req *model2.Event) (*model2.Event, error) {
	return func(ctx context.Context, req *model2.Event) (*model2.Event, error) {

		for _, record := range req.Records {
			if record.Properties["action"] == nil || record.Properties["action"].GetStringValue() != action { // @todo this logic should be in server side
				continue
			}
			err := processFunc(ctx, h.mapper.FromRecord(record))

			if err != nil {
				return nil, err
			}
		}

		return req, nil
	}
}

func (h handler[Entity]) Create(processFunc RecordProcessFunc[Entity]) Handler[Entity] {
	h.action = resource_model.EventAction_CREATE

	h.handle(processFunc)

	return h
}

func (h handler[Entity]) PreProcess(processFunc RecordProcessFunc[Entity]) Handler[Entity] {
	return h.Before().Create(processFunc).Update(processFunc).Delete(processFunc)
}

func (h handler[Entity]) PostProcess(processFunc RecordProcessFunc[Entity]) Handler[Entity] {
	return h.After().Create(processFunc).Update(processFunc).Delete(processFunc)
}

func (h handler[Entity]) Process(processFunc RecordProcessFunc[Entity]) Handler[Entity] {
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
	ri := h.mapper.ResourceIdentity()

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
		Annotations: map[string]string{
			annotations.ServiceKey: h.ext.getServiceKey(),
		},
	}

	extension.Call = h.ext.PrepareCall(extension)

	newExtension, err := h.extensionRepo.Apply(context.TODO(), extension)

	if err != nil {
		panic(err)
	}

	h.ext.RegisterExtension(newExtension)
}

func (h handler[Entity]) prepareName() string {
	ri := h.mapper.ResourceIdentity()

	if h.name == "" {
		h.name = "def_" + ri.Namespace + "_" + ri.Name + "_" + string(h.action) + "_" + strconv.Itoa(int(h.order))
	}

	return h.name
}

func NewHandler[Entity interface{}](dhClient Client, ext Extension, mapper abs.EntityMapper[Entity]) Handler[Entity] {
	extensionRepo := R[*resource_model.Extension](dhClient, resource_model.ExtensionMapperInstance)

	return handler[Entity]{
		dhClient:      dhClient,
		ext:           ext,
		sync:          true,
		mapper:        mapper,
		extensionRepo: extensionRepo,
		responds:      true,
		order:         100,
	}
}
