package grpc_service

import (
	"context"
	"data-handler/grpc/stub"
	"data-handler/service"
)

type recordServiceServer struct {
	stub.RecordServiceServer
	service service.RecordService
}

func (r *recordServiceServer) List(ctx context.Context, request *stub.ListRecordRequest) (*stub.ListRecordResponse, error) {
	records, total, err := r.service.List(ctx, service.RecordListParams{
		Workspace:  request.Workspace,
		Resource:   request.Resource,
		Query:      request.Query,
		Limit:      request.Limit,
		Offset:     request.Offset,
		UseHistory: request.UseHistory,
	})

	return &stub.ListRecordResponse{
		Content: records,
		Total:   total,
		Error:   toProtoError(err),
	}, nil
}

func (r *recordServiceServer) Create(ctx context.Context, request *stub.CreateRecordRequest) (*stub.CreateRecordResponse, error) {
	records, inserted, err := r.service.Create(ctx, service.RecordCreateParams{
		Workspace:      request.Workspace,
		Records:        request.Records,
		IgnoreIfExists: request.IgnoreIfExists,
	})

	return &stub.CreateRecordResponse{
		Records:  records,
		Inserted: inserted,
		Error:    toProtoError(err),
	}, nil
}

func (r *recordServiceServer) Update(ctx context.Context, request *stub.UpdateRecordRequest) (*stub.UpdateRecordResponse, error) {
	records, err := r.service.Update(ctx, service.RecordUpdateParams{
		Workspace:    request.Workspace,
		Records:      request.Records,
		CheckVersion: request.CheckVersion,
	})

	return &stub.UpdateRecordResponse{
		Records: records,
		Error:   toProtoError(err),
	}, nil
}

func (r *recordServiceServer) Get(ctx context.Context, request *stub.GetRecordRequest) (*stub.GetRecordResponse, error) {
	record, err := r.service.Get(ctx, service.RecordGetParams{
		Workspace: request.Workspace,
		Resource:  request.Resource,
		Id:        request.Id,
	})

	return &stub.GetRecordResponse{
		Record: record,
		Error:  toProtoError(err),
	}, nil
}

func (r *recordServiceServer) Delete(ctx context.Context, request *stub.DeleteRecordRequest) (*stub.DeleteRecordResponse, error) {
	err := r.service.Delete(ctx, service.RecordDeleteParams{
		Workspace: request.Workspace,
		Resource:  request.Resource,
		Ids:       request.Ids,
	})

	return &stub.DeleteRecordResponse{
		Error: toProtoError(err),
	}, nil
}

func NewRecordServiceServer(service service.RecordService) stub.RecordServiceServer {
	return &recordServiceServer{service: service}
}
