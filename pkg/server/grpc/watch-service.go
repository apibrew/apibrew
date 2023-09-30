package grpc

import (
	"context"
	"github.com/apibrew/apibrew/pkg/service"
	"github.com/apibrew/apibrew/pkg/stub"
	log "github.com/sirupsen/logrus"
)

type watchGrpcService struct {
	stub.WatchServer
	watchService          service.WatchService
	authenticationService service.AuthenticationService
}

func (w *watchGrpcService) Watch(req *stub.WatchRequest, res stub.Watch_WatchServer) error {
	ictx, err := interceptRequest(w.authenticationService, res.Context(), req)

	if err != nil {
		return err
	}

	localCtx, cancel := context.WithCancel(ictx)
	defer func() {
		cancel()
	}()

	out, err := w.watchService.Watch(localCtx, service.WatchParams{
		Selector:   req.Selector,
		BufferSize: 500,
	})

	if err != nil {
		return err
	}

	for message := range out {
		err := res.Send(message)

		if err != nil {
			cancel()
			log.Error(err)
			return err
		}
	}

	return nil
}

func NewWatchServer(service service.WatchService, authenticationService service.AuthenticationService) stub.WatchServer {
	return &watchGrpcService{watchService: service, authenticationService: authenticationService}
}
