package grpc

import (
	"context"
	"github.com/apibrew/apibrew/pkg/service"
	"github.com/apibrew/apibrew/pkg/stub"
	log "github.com/sirupsen/logrus"
)

type watchGrpcService struct {
	stub.WatchServer
	watchService service.WatchService
}

func (w *watchGrpcService) Watch(req *stub.WatchRequest, res stub.Watch_WatchServer) error {
	localCtx, cancel := context.WithCancel(res.Context())
	defer func() {
		cancel()
	}()

	out := w.watchService.Watch(localCtx, service.WatchParams{
		Selector:   req.Selector,
		BufferSize: 500,
	})

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

func NewWatchServer(service service.WatchService) stub.WatchServer {
	return &watchGrpcService{watchService: service}
}
