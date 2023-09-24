package impl

import (
	"context"
	"github.com/apibrew/apibrew/pkg/errors"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/resource_model"
	"github.com/apibrew/apibrew/pkg/resources"
	"github.com/apibrew/apibrew/pkg/service"
	log "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/types/known/timestamppb"
	"sync"
	"time"
)

type eventSignal struct {
	ctx     context.Context
	handler chan *model.Event
}

type eventChannelService struct {
	channelChans         map[string]chan string
	eventMap             map[string]map[string]*model.Event
	eventSignalMap       map[string]map[string]eventSignal
	mu                   sync.Mutex
	authorizationService service.AuthorizationService
}

func (e *eventChannelService) WriteEvent(ctx context.Context, channelKey string, event *model.Event) errors.ServiceError {
	if err := e.authorizationService.CheckRecordAccess(ctx, service.CheckRecordAccessParams{
		Resource:  resources.ExtensionResource,
		Operation: resource_model.PermissionOperation_FULL,
	}); err != nil {
		return err
	}

	e.ensureChannel(channelKey)

	if e.eventMap[channelKey][event.Id] == nil || e.eventSignalMap[channelKey][event.Id].ctx.Err() != nil {
		log.Warn("Event is not exists or already discarded: " + event.Id)
		return errors.LogicalError.WithMessage("Event is not exists or already discarded: " + event.Id)
	}

	e.eventSignalMap[channelKey][event.Id].handler <- event

	return nil
}

func (e *eventChannelService) PollEvents(ctx context.Context, channelKey string) (chan *model.Event, errors.ServiceError) {
	log.Infof("Polling events for channel: %v", channelKey)
	if err := e.authorizationService.CheckRecordAccess(ctx, service.CheckRecordAccessParams{
		Resource:  resources.ExtensionResource,
		Operation: resource_model.PermissionOperation_FULL,
	}); err != nil {
		return nil, err
	}

	e.ensureChannel(channelKey)

	var eventChan = make(chan *model.Event)

	go func() {
		for {
			select {
			case <-ctx.Done():
				close(eventChan)
				return
			case eventId := <-e.channelChans[channelKey]:
				event := e.eventMap[channelKey][eventId]

				if event != nil {
					eventChan <- event
				} else {
					log.Warn("Event not found or already discarted: " + eventId)
					releaseEvent(e, channelKey, eventId)
				}
			case <-time.After(5 * time.Second):
				eventChan <- &model.Event{
					Id:   "heartbeat-message",
					Time: timestamppb.Now(),
				}
			}
		}
	}()

	log.Infof("Polling events for channel: %v - done", channelKey)

	return eventChan, nil
}

func (e *eventChannelService) Exec(ctx context.Context, channelKey string, event *model.Event) (*model.Event, errors.ServiceError) {
	e.ensureChannel(channelKey)
	e.eventMap[channelKey][event.Id] = event

	var handler = make(chan *model.Event)

	cctx, cancel := context.WithTimeout(ctx, 5*time.Second)

	defer cancel()
	defer releaseEvent(e, channelKey, event.Id)

	e.eventSignalMap[channelKey][event.Id] = eventSignal{
		ctx:     cctx,
		handler: handler,
	}

	select {
	case e.channelChans[channelKey] <- event.Id:
	case <-cctx.Done():
		log.Warn("Event channel timeout[send]: " + event.Id)
		cancel()
	}

	select {
	case result := <-handler:
		return result, nil
	case <-cctx.Done():
		log.Warn("Event handler timeout[receive]: " + event.Id)
	}

	return nil, errors.LogicalError.WithMessage(cctx.Err().Error())
}

func releaseEvent(e *eventChannelService, channelKey string, eventId string) {
	e.mu.Lock()
	defer e.mu.Unlock()

	delete(e.eventMap[channelKey], eventId)
	delete(e.eventSignalMap[channelKey], eventId)
}

func (e *eventChannelService) ensureChannel(key string) {
	e.mu.Lock()
	defer e.mu.Unlock()

	if e.channelChans[key] == nil {
		e.channelChans[key] = make(chan string, 100)
	}

	if e.eventMap[key] == nil {
		e.eventMap[key] = make(map[string]*model.Event)
	}

	if e.eventSignalMap[key] == nil {
		e.eventSignalMap[key] = make(map[string]eventSignal)
	}
}

func NewEventChannelService(authorizationService service.AuthorizationService) service.EventChannelService {
	return &eventChannelService{
		channelChans:         make(map[string]chan string),
		eventMap:             make(map[string]map[string]*model.Event),
		eventSignalMap:       make(map[string]map[string]eventSignal),
		authorizationService: authorizationService,
	}
}
