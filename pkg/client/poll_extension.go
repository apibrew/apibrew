package client

import (
	"context"
	"github.com/apibrew/apibrew/pkg/helper"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/resource_model"
	"github.com/apibrew/apibrew/pkg/util"
	log "github.com/sirupsen/logrus"
)

type pollExtension struct {
	serviceKey                    string
	client                        *dhClient
	functions                     map[string]ExternalFunction
	registeredExtensions          []*resource_model.Extension
	extensionEventSelectorMatcher *helper.ExtensionEventSelectorMatcher
}

func (e *pollExtension) PrepareCall(extension *resource_model.Extension) resource_model.ExtensionExternalCall {
	return resource_model.ExtensionExternalCall{
		ChannelCall: &resource_model.ExtensionChannelCall{
			ChannelKey: e.serviceKey,
		},
	}
}

func (e *pollExtension) getServiceKey() string {
	return e.serviceKey
}

func (e *pollExtension) RegisterExtension(newExtension *resource_model.Extension) {
	e.registeredExtensions = append(e.registeredExtensions, newExtension)
}

func (e *pollExtension) RegisterFunction(name string, handler ExternalFunction) {
	e.functions[name] = handler
}

// WithServiceKey
func (e *pollExtension) WithServiceKey(serviceKey string) Extension {
	e.serviceKey = serviceKey
	return e
}

func (e *pollExtension) Run(ctx context.Context) error {
	eventsChan, err := e.client.PollEvents(ctx, e.serviceKey)

	if err != nil {
		return err
	}

	for event := range eventsChan {
		e.processEvent(event)
	}

	return nil
}

func (e *pollExtension) prepareExtSelector(extension *resource_model.Extension) *model.EventSelector {
	return &model.EventSelector{
		Actions: util.ArrayMap(extension.Selector.Actions, func(s resource_model.EventAction) model.Event_Action {
			return model.Event_Action(model.Event_Action_value[string(s)])
		}),
		RecordSelector: nil,
		Namespaces:     extension.Selector.Namespaces,
		Resources:      extension.Selector.Resources,
		Ids:            extension.Selector.Ids,
		Annotations:    extension.Selector.Annotations,
	}
}

func (e *pollExtension) processEvent(originalEvent *model.Event) {
	var processedEvent = originalEvent

	for _, ext := range e.registeredExtensions {
		if e.extensionEventSelectorMatcher.SelectorMatches(originalEvent, e.prepareExtSelector(ext)) {
			funcName := ext.Name

			if e.functions[funcName] == nil {
				log.Warnf("External function not found: " + funcName)
			}

			var err error
			processedEvent, err = e.functions[funcName](context.TODO(), processedEvent)

			if processedEvent == nil {
				processedEvent = originalEvent
			}

			if err != nil {
				processedEvent.Error = &model.Error{
					Code:    model.ErrorCode_INTERNAL_ERROR,
					Message: err.Error(),
				}
			}

			err = e.client.WriteEvent(context.TODO(), e.serviceKey, processedEvent)

			if err != nil {
				log.Error("Error while writing event: ", err)
			}
		}
	}
}

func (d *dhClient) NewPollExtension() Extension {
	return &pollExtension{
		client:                        d,
		serviceKey:                    "golang-ext",
		functions:                     make(map[string]ExternalFunction),
		extensionEventSelectorMatcher: &helper.ExtensionEventSelectorMatcher{},
	}
}
