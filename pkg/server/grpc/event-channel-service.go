package grpc

import (
	"context"
	"github.com/apibrew/apibrew/pkg/service"
	"github.com/apibrew/apibrew/pkg/stub"
	log "github.com/sirupsen/logrus"
)

type eventChannelGrpcService struct {
	stub.EventChannelServer
	eventChannelService   service.EventChannelService
	authenticationService service.AuthenticationService
}

func (e *eventChannelGrpcService) Poll(req *stub.EventPollRequest, srv stub.EventChannel_PollServer) error {
	ictx, err := interceptRequest(e.authenticationService, srv.Context(), req)

	localCtx, cancel := context.WithCancel(ictx)
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

func NewEventChannelGrpcService(service service.EventChannelService, authenticationService service.AuthenticationService) stub.EventChannelServer {
	return &eventChannelGrpcService{eventChannelService: service, authenticationService: authenticationService}
}
