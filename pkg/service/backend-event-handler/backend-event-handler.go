package backend_event_handler

import (
	"context"
	log "github.com/sirupsen/logrus"
	"github.com/tislib/apibrew/pkg/errors"
	"github.com/tislib/apibrew/pkg/model"
	"github.com/tislib/apibrew/pkg/util"
	"google.golang.org/protobuf/types/known/timestamppb"
	"sort"
)

type BackendEventHandler interface {
	RegisterHandler(handler Handler)
	UnRegisterHandler(handler Handler)
	HandleInternalOperation(ctx context.Context, incoming *model.Event, actualHandler HandlerFunc) (*model.Event, errors.ServiceError)
	PrepareInternalEvent(ctx context.Context, event *model.Event) *model.Event
}

type backendEventHandler struct {
	handlers []Handler
}

func (b *backendEventHandler) PrepareInternalEvent(ctx context.Context, event *model.Event) *model.Event {
	event.Id = "internal-event-" + util.RandomHex(6)
	event.Time = timestamppb.Now()

	return event
}

func (b *backendEventHandler) HandleInternalOperation(ctx context.Context, nextEvent *model.Event, actualHandler HandlerFunc) (*model.Event, errors.ServiceError) {
	handlers := b.filterHandlersForEvent(nextEvent)

	handlers = append(handlers, Handler{
		Id:        "actualHandler",
		Fn:        actualHandler,
		Order:     NaturalOrder,
		Finalizes: false,
		Sync:      true,
		Responds:  true,
	})

	sort.Sort(ByOrder(handlers))

	for _, handler := range handlers {
		if !handler.Sync {
			go func() {
				_, err := handler.Fn(ctx, nextEvent)

				if err != nil {
					log.Error("Error from async handler", err)
				}
			}()
		} else {
			result, err := handler.Fn(ctx, nextEvent)

			if err != nil {
				return nil, err
			}

			if handler.Responds {
				nextEvent = result
			}

			if handler.Finalizes {
				break
			}

		}
	}

	return nextEvent, nil
}

func (b *backendEventHandler) RegisterHandler(handler Handler) {
	log.Debugf("Registering handler: %s [%v]", handler.Id, handler)

	b.handlers = append(b.handlers, handler)
}

func (b *backendEventHandler) UnRegisterHandler(handler Handler) {
	log.Debugf("Unregister handler: %s [%v]", handler.Id, handler)

	for i, h := range b.handlers {
		if h.Id == handler.Id {
			b.handlers = append(b.handlers[:i], b.handlers[i+1:]...)
			return
		}
	}

	log.Debugf("Unregister handler[not found]: %s [%v]", handler.Id, handler)
}

func (b *backendEventHandler) filterHandlersForEvent(incoming *model.Event) []Handler {
	var result []Handler

	for _, handler := range b.handlers {
		if b.SelectorMatches(incoming, handler.Selector) {
			result = append(result, handler)
		}
	}

	return result
}

func (b *backendEventHandler) SelectorMatches(incoming *model.Event, selector *model.EventSelector) bool {
	if selector == nil {
		return true
	}

	if selector.Resources != nil {
		var found = false
		for _, resource := range selector.Resources {
			if resource == incoming.Resource.Name {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}

	if selector.Actions != nil {
		var found = false
		for _, action := range selector.Actions {
			if action == incoming.Action {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}

	if selector.Ids != nil {
		var found = false
		for _, id := range selector.Ids {
			if id == incoming.Id {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}

	if selector.Namespaces != nil {
		var found = false
		for _, namespace := range selector.Namespaces {
			if namespace == incoming.Resource.Namespace {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}

	if selector.Annotations != nil {
		for key, value := range selector.Annotations {
			var found = false
			if incoming.Resource.Annotations[key] == value {
				found = true
				break
			}

			if value == "*" {
				if _, ok := incoming.Resource.Annotations[key]; ok {
					break
				}
			}

			if !found {
				return false
			}
		}

	}

	if selector.RecordSelector != nil {
		//TODO implement me
	}

	return true
}

func NewBackendEventHandler() BackendEventHandler {
	return &backendEventHandler{}
}
