package grpc

import (
	"context"
	log "github.com/sirupsen/logrus"
	"github.com/tislib/apibrew/pkg/abs"
	"github.com/tislib/apibrew/pkg/stub"
)

type watchGrpcService struct {
	stub.WatchServer
	watchService abs.WatchService
}

func (w *watchGrpcService) Watch(req *stub.WatchRequest, res stub.Watch_WatchServer) error {
	localCtx, cancel := context.WithCancel(res.Context())
	defer func() {
		cancel()
	}()

	out := w.watchService.Watch(localCtx, abs.WatchParams{
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

func NewWatchServer(service abs.WatchService) stub.WatchServer {
	return &watchGrpcService{watchService: service}
}
