package grpc

import (
	"context"
	"github.com/apibrew/apibrew/pkg/service"
	"github.com/apibrew/apibrew/pkg/stub"
	log "github.com/sirupsen/logrus"
)

type eventChannelGrpcService struct {
	stub.EventChannelServer
	eventChannelService service.EventChannelService
}

func (e *eventChannelGrpcService) Poll(req *stub.EventPollRequest, srv stub.EventChannel_PollServer) error {
	localCtx, cancel := context.WithCancel(srv.Context())
	defer func() {
		cancel()
	}()

	out, err := e.eventChannelService.PollEvents(localCtx, req.ChannelKey)

	if err != nil {
		return err
	}

	for message := range out {
		err := srv.Send(message)

		if err != nil {
			cancel()
			log.Error(err)
			return err
		}
	}

	return nil
}

func (e *eventChannelGrpcService) Write(ctx context.Context, req *stub.EventWriteRequest) (*stub.EventWriteResponse, error) {
	err := e.eventChannelService.WriteEvent(ctx, req.ChannelKey, req.Event)

	if err != nil {
		log.Error(err)
		return nil, err
	}

	return &stub.EventWriteResponse{}, nil
}

func NewEventChannelGrpcService(service service.EventChannelService) stub.EventChannelServer {
	return &eventChannelGrpcService{eventChannelService: service}
}
