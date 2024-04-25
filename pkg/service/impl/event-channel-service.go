package impl

import (
	"context"
	"github.com/apibrew/apibrew/pkg/errors"
	"github.com/apibrew/apibrew/pkg/model"
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
	channelChans         sync.Map
	eventSignalMap       sync.Map
	authorizationService service.AuthorizationService
	config               *model.EventChannelConfig
}

func (e *eventChannelService) Init(config *model.AppConfig) {
	e.config = config.EventChannelConfig

	if e.config == nil {
		e.config = &model.EventChannelConfig{}
	}

	if config.EventChannelConfig == nil || config.EventChannelConfig.MaxWaitTimeMs == 0 {
		e.config.MaxWaitTimeMs = 5000
	}

	if config.EventChannelConfig == nil || config.EventChannelConfig.MaxChannelSize == 0 {
		e.config.MaxChannelSize = 100
	}
}

func (e *eventChannelService) WriteEvent(ctx context.Context, event *model.Event) error {
	if err := e.authorizationService.CheckIsExtensionController(ctx); err != nil {
		return err
	}

	value, ok := e.eventSignalMap.Load(event.Id)

	if !ok {
		log.Warn("Event is not exists or already discarded: " + event.Id)
		return errors.LogicalError.WithMessage("Event is not exists or already discarded: " + event.Id)
	}

	signalMap := value.(eventSignal)

	signalMap.handler <- event

	return nil
}

func (e *eventChannelService) PollEvents(ctx context.Context, channelKey string) (chan *model.Event, error) {
	log.Infof("Polling events for channel: %v", channelKey)
	if err := e.authorizationService.CheckIsExtensionController(ctx); err != nil {
		return nil, err
	}

	eventChan := e.ensureChannel(channelKey)
	out := make(chan *model.Event, 100)

	go func() {
		for {
			select {
			case <-ctx.Done():
				close(out)
				return
			case event := <-eventChan:
				if event != nil {
					out <- event
					log.Tracef("Event sent to channel: %v", channelKey)
				}
			case <-time.After(3 * time.Second):
				log.Tracef("Heartbeat message sent to channel: %v", channelKey)
				out <- &model.Event{
					Id:   "heartbeat-message",
					Time: timestamppb.Now(),
				}
			}
		}
	}()

	log.Infof("Polling events for channel: %v - done", channelKey)

	return out, nil
}

func (e *eventChannelService) Exec(ctx context.Context, channelKey string, event *model.Event) (*model.Event, error) {
	eventChan := e.ensureChannel(channelKey)

	var handler chan *model.Event

	cctx, cancel := context.WithTimeout(ctx, time.Duration(e.config.MaxWaitTimeMs)*time.Millisecond)

	defer cancel()
	defer releaseEvent(e, event.Id)

	if event.Sync {
		handler = make(chan *model.Event)
		e.eventSignalMap.Store(event.Id, eventSignal{ctx: cctx, handler: handler})
	}

	select {
	case eventChan <- event:
	case <-cctx.Done():
		log.Warn("Event channel timeout[send]: " + event.Id + "/" + channelKey)
		cancel()
	}

	if !event.Sync {
		return nil, nil
	}

	select {
	case result := <-handler:
		return result, nil
	case <-cctx.Done():
		log.Warn("Event handler timeout[receive]: " + event.Id + "/" + channelKey)
	}

	return nil, errors.LogicalError.WithMessage(cctx.Err().Error())
}

func releaseEvent(e *eventChannelService, eventId string) {
	e.eventSignalMap.Delete(eventId)
}

func (e *eventChannelService) ensureChannel(key string) chan *model.Event {
	value, ok := e.channelChans.Load(key)

	var ch chan *model.Event
	if !ok || value == nil {
		ch = make(chan *model.Event, 100)
		e.channelChans.Store(key, ch)
	} else {
		ch = value.(chan *model.Event)
	}

	return ch
}

func NewEventChannelService(authorizationService service.AuthorizationService) service.EventChannelService {
	return &eventChannelService{
		channelChans:         sync.Map{},
		eventSignalMap:       sync.Map{},
		authorizationService: authorizationService,
	}
}
