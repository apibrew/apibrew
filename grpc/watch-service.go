package grpc_service

import (
	"context"
	"data-handler/grpc/stub"
	"data-handler/service"
	"data-handler/service/params"
)

type watchGrpcService struct {
	stub.WatchServiceServer
	watchService service.WatchService
}

func (w *watchGrpcService) Watch(req *stub.WatchRequest, res stub.WatchService_WatchServer) error {
	localCtx, cancel := context.WithCancel(res.Context())
	defer cancel()

	out := w.watchService.Watch(localCtx, params.WatchParams{
		Workspace:  req.Workspace,
		Resource:   req.Resource,
		Query:      req.Query,
		BufferSize: 500,
	})

	for item := range out {
		err := res.Send(&stub.WatchResponse{Message: item})

		if err != nil {
			cancel()
			return err
		}
	}

	return nil
}

func NewWatchServiceServer(service service.WatchService) stub.WatchServiceServer {
	return &watchGrpcService{watchService: service}
}
