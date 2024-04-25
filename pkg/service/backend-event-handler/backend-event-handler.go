package backend_event_handler

import (
	"context"
	"github.com/apibrew/apibrew/pkg/errors"
	"github.com/apibrew/apibrew/pkg/logging"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/util"
	log "github.com/sirupsen/logrus"
	"sort"
)

type backendEventHandler struct {
	handlers                      []Handler
	extensionEventSelectorMatcher *ExtensionEventSelectorMatcher
}

func (b *backendEventHandler) Handle(ctx context.Context, originalEvent *model.Event) (*model.Event, error) {
	nextEvent := originalEvent

	var handlers = b.filterHandlersForEvent(nextEvent)

	if len(handlers) == 0 || (len(handlers) == 1 && handlers[0].Id == "audit-handler") {
		log.Tracef("No handlers found for event: %s", ShortEventInfo(nextEvent))
		return nextEvent, errors.LogicalError.WithDetails("No handlers found for event")
	}

	sort.Sort(ByOrder(handlers))

	logger := log.WithFields(logging.CtxFields(ctx))

	logger.Debugf("Starting handler chain: %d", len(handlers))
	for _, handler := range handlers {
		nextEvent.Resource = originalEvent.Resource
		if !handler.Sync {
			nextEvent.Sync = false
			go func(localHandler Handler) {
				// wait for until current request to be done
				<-ctx.Done()

				logger.Debugf("Calling handler[%d sync: %v]: %s - %s", localHandler.Order, localHandler.Sync, localHandler.Name, ShortEventInfo(nextEvent))
				logger.Tracef("Processing event[body]: %s", nextEvent)

				_, err := localHandler.Fn(util.NewContextWithValues(context.TODO(), ctx), nextEvent)

				if err != nil {
					logger.Error("Error from async handler", err)
				}
			}(handler)
		} else {
			logger.Debugf("Calling handler[%d sync: %v]: %s - %s", handler.Order, handler.Sync, handler.Name, ShortEventInfo(nextEvent))
			logger.Tracef("Processing event[body]: %s", nextEvent)

			nextEvent.Sync = true
			result, err := handler.Fn(ctx, nextEvent)

			if err != nil {
				logger.Warnf("Handler [%s] responded with error: %v", handler.Name, err)
				return nil, err
			}

			logger.Debugf("Handler responded: %s - %s", handler.Name, ShortEventInfo(result))
			logger.Tracef("Handler responded: %s", result)

			if result != nil && result.Error != nil {
				logger.Warnf("Handler [%s] responded with error: %v", handler.Name, result.Error)
				return nil, errors.FromProtoError(result.Error)
			}

			if handler.Responds {
				logger.Debugf("Handler [%s] responded with result", handler.Name)
				nextEvent = result
			}

			if nextEvent == nil {
				logger.Debugf("Handler [%s] breaks", handler.Name)
				break
			}

			if nextEvent.Error != nil {
				logger.Warnf("Handler [%s] responded with error: %v", handler.Name, nextEvent.Error)
				return nil, errors.FromProtoError(nextEvent.Error)
			}

			if handler.Finalizes {
				logger.Debugf("Handler [%s] finalizes", handler.Name)
				break
			}

		}
	}
	logger.Debugf("Finished handler chain - %s", ShortEventInfo(nextEvent))
	logger.Tracef("Processed event: %s", nextEvent)

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
		if b.extensionEventSelectorMatcher.SelectorMatches(incoming, handler.Selector) {
			log.Tracef("Handler matches: %s [%v]", handler.Id, handler)
			result = append(result, handler)
		}
	}

	return result
}

func NewBackendEventHandler() BackendEventHandler {
	return &backendEventHandler{extensionEventSelectorMatcher: &ExtensionEventSelectorMatcher{}}
}

type BackendEventHandler interface {
	RegisterHandler(handler Handler)
	UnRegisterHandler(handler Handler)
	Handle(ctx context.Context, incoming *model.Event) (*model.Event, error)
}
