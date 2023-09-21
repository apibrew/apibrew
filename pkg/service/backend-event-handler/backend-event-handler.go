package backend_event_handler

import (
	"context"
	"github.com/apibrew/apibrew/pkg/errors"
	"github.com/apibrew/apibrew/pkg/logging"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/util"
	log "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/proto"
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

func (b *backendEventHandler) HandleInternalOperation(ctx context.Context, originalEvent *model.Event, actualHandler HandlerFunc) (*model.Event, errors.ServiceError) {
	nextEvent := originalEvent

	handlers := b.filterHandlersForEvent(nextEvent)

	if !nextEvent.Resource.Virtual {
		handlers = append(handlers, Handler{
			Id:        "actualHandler",
			Name:      "actualHandler",
			Fn:        actualHandler,
			Order:     NaturalOrder,
			Finalizes: false,
			Sync:      true,
			Responds:  true,
		})
	}

	sort.Sort(ByOrder(handlers))

	logger := log.WithFields(logging.CtxFields(ctx))

	logger.Debugf("Starting handler chain: %d", len(handlers))
	for _, handler := range handlers {
		nextEvent.Resource = originalEvent.Resource

		logger.Debugf("Calling handler: %s", handler.Name)
		if !handler.Sync {
			nextEvent.Sync = false
			go func(localHandler Handler) {
				_, err := localHandler.Fn(context.TODO(), nextEvent)

				if err != nil {
					logger.Error("Error from async handler", err)
				}
			}(handler)
		} else {
			nextEvent.Sync = true
			result, err := handler.Fn(ctx, nextEvent)
			logger.Debugf("Handler responded: %s", handler.Name)

			if err != nil {
				logger.Warnf("Handler [%s] responded with error: %v", handler.Name, err)
				return nil, err
			}

			if handler.Responds {
				logger.Debugf("Handler [%s] responded with result", handler.Name)
				nextEvent = result
			}

			if nextEvent == nil {
				break
			}

			if handler.Finalizes {
				logger.Debugf("Handler [%s] finalizes", handler.Name)
				break
			}

		}
	}
	logger.Debugf("Finished handler chain")

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
			log.Tracef("Handler matches: %s [%v]", handler.Id, handler)
			result = append(result, handler)
		}
	}

	return result
}

func (b *backendEventHandler) SelectorMatches(incoming *model.Event, selector *model.EventSelector) bool {
	if selector == nil {
		return true
	}

	if len(selector.Resources) > 0 {
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

	if len(selector.Actions) > 0 {
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

	if len(selector.Ids) > 0 {
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

	if len(selector.Namespaces) > 0 {
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

	if len(selector.Annotations) > 0 {
		for key, value := range selector.Annotations {
			if incoming.Resource.Annotations[key] == value {
				break
			}

			if value == "*" {
				if _, ok := incoming.Resource.Annotations[key]; ok {
					break
				}
			}

			return false
		}
	}

	if selector.RecordSelector != nil {
		return recordSelectorMatches(incoming, selector.RecordSelector)
	}

	return true
}

func recordSelectorMatches(incoming *model.Event, selector *model.BooleanExpression) bool {
	if selector == nil {
		return true
	}

	if selector.GetAnd() != nil {
		for _, child := range selector.GetAnd().Expressions {
			if !recordSelectorMatches(incoming, child) {
				return false
			}
		}

		return true
	}

	if selector.GetOr() != nil {
		for _, child := range selector.GetOr().Expressions {
			if recordSelectorMatches(incoming, child) {
				return true
			}
		}

		return false
	}

	if selector.GetNot() != nil {
		return !recordSelectorMatches(incoming, selector.GetNot())
	}

	if selector.GetEqual() != nil {
		left := resolve(incoming, selector.GetEqual().Left)
		right := resolve(incoming, selector.GetEqual().Right)

		return proto.Equal(left, right)
	}

	return true
}

func resolve(incoming *model.Event, left *model.Expression) proto.Message {
	if left.GetProperty() != "" {
		if len(incoming.Records) == 0 {
			return nil
		}
		return incoming.Records[0].Properties[left.GetProperty()]
	}

	if left.GetRefValue() != nil {
		return left.GetRefValue()
	}

	if left.GetValue() != nil {
		return left.GetValue()
	}

	return nil
}

func NewBackendEventHandler() BackendEventHandler {
	return &backendEventHandler{}
}
