package grpc_service

import (
	"data-handler/grpc/stub"
	"data-handler/service"
	"data-handler/service/params"
)

type watchGrpcService struct {
	stub.WatchServiceServer
	watchService service.WatchService
}

func (w *watchGrpcService) Watch(req *stub.WatchRequest, res stub.WatchService_WatchServer) error {
	out := w.watchService.Watch(res.Context(), params.WatchParams{
		Workspace:  req.Workspace,
		Resource:   req.Resource,
		Query:      req.Query,
		BufferSize: 500,
	})

	for item := range out {
		res.Send(&stub.WatchResponse{Message: item})
	}
}

func NewWatchServiceServer(service service.WatchService) stub.WatchServiceServer {
	return &watchGrpcService{watchService: service}
}
