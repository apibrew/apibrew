package service

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/tislib/data-handler/pkg/errors"
	"github.com/tislib/data-handler/pkg/ext"
	"github.com/tislib/data-handler/pkg/model"
	"github.com/tislib/data-handler/pkg/service/handler"
	"github.com/tislib/data-handler/pkg/service/params"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type extensionHandler struct {
	handler.BaseHandler
	recordClient ext.RecordExtensionServiceClient
}

func (h *extensionHandler) List(ctx context.Context, params params.RecordListParams) (bool, []*model.Record, uint32, errors.ServiceError) {
	resp, err := h.recordClient.List(ctx, &ext.ListRecordRequest{
		Namespace:         params.Namespace,
		Resource:          params.Resource,
		Query:             params.Query,
		Limit:             params.Limit,
		Offset:            params.Offset,
		UseHistory:        params.UseHistory,
		ResolveReferences: params.ResolveReferences,
	})

	if err != nil {
		log.Error(err)
		return true, nil, 0, errors.InternalError.WithDetails("External communication error")
	}

	return true, resp.Content, resp.Total, nil
}

func NewExtensionHandler(extension *model.Extension) *handler.BaseHandler {
	h := new(extensionHandler)

	var opts []grpc.DialOption

	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", extension.Server.Host, extension.Server.Port), opts...)
	if err != nil {
		panic(err)
	}

	h.recordClient = ext.NewRecordExtensionServiceClient(conn)

	return &handler.BaseHandler{
		List: h.List,
	}
}
