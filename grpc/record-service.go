package grpc_service

import (
	"context"
	"data-handler/grpc/stub"
	"data-handler/service"
	"data-handler/service/params"
	"data-handler/util"
)

type recordServiceServer struct {
	stub.RecordServiceServer
	service service.RecordService
}

func (r *recordServiceServer) List(ctx context.Context, request *stub.ListRecordRequest) (*stub.ListRecordResponse, error) {
	records, total, err := r.service.List(ctx, params.RecordListParams{
		Namespace:         request.Namespace,
		Resource:          request.Resource,
		Limit:             request.Limit,
		Offset:            request.Offset,
		UseHistory:        request.UseHistory,
		ResolveReferences: request.ResolveReferences,
	})

	return &stub.ListRecordResponse{
		Content: records,
		Total:   total,
		Error:   toProtoError(err),
	}, err
}

func (r *recordServiceServer) Search(ctx context.Context, request *stub.SearchRecordRequest) (*stub.SearchRecordResponse, error) {
	records, total, err := r.service.List(ctx, params.RecordListParams{
		Namespace:         request.Namespace,
		Resource:          request.Resource,
		Limit:             request.Limit,
		Query:             request.Query,
		Offset:            request.Offset,
		UseHistory:        request.UseHistory,
		ResolveReferences: request.ResolveReferences,
	})

	return &stub.SearchRecordResponse{
		Content: records,
		Total:   total,
		Error:   toProtoError(err),
	}, err
}

func (r *recordServiceServer) Create(ctx context.Context, request *stub.CreateRecordRequest) (*stub.CreateRecordResponse, error) {
	records, inserted, err := r.service.Create(ctx, params.RecordCreateParams{
		Namespace:      request.Namespace,
		Records:        util.ArrayPrepend(request.Records, request.Record),
		IgnoreIfExists: request.IgnoreIfExists,
	})

	return &stub.CreateRecordResponse{
		Record:   util.ArrayFirst(records),
		Records:  util.ArrayCut(records, 0),
		Inserted: inserted,
		Error:    toProtoError(err),
	}, err
}

func (r *recordServiceServer) Update(ctx context.Context, request *stub.UpdateRecordRequest) (*stub.UpdateRecordResponse, error) {
	records, err := r.service.Update(ctx, params.RecordUpdateParams{
		Namespace:    request.Namespace,
		Records:      util.ArrayPrepend(request.Records, request.Record),
		CheckVersion: request.CheckVersion,
	})

	return &stub.UpdateRecordResponse{
		Record:  util.ArrayFirst(records),
		Records: util.ArrayCut(records, 0),
		Error:   toProtoError(err),
	}, err
}

func (r *recordServiceServer) Get(ctx context.Context, request *stub.GetRecordRequest) (*stub.GetRecordResponse, error) {
	record, err := r.service.Get(ctx, params.RecordGetParams{
		Namespace: request.Namespace,
		Resource:  request.Resource,
		Id:        request.Id,
	})

	return &stub.GetRecordResponse{
		Record: record,
		Error:  toProtoError(err),
	}, err
}

func (r *recordServiceServer) Delete(ctx context.Context, request *stub.DeleteRecordRequest) (*stub.DeleteRecordResponse, error) {
	err := r.service.Delete(ctx, params.RecordDeleteParams{
		Namespace: request.Namespace,
		Resource:  request.Resource,
		Ids:       request.Ids,
	})

	return &stub.DeleteRecordResponse{
		Error: toProtoError(err),
	}, err
}

func NewRecordServiceServer(service service.RecordService) stub.RecordServiceServer {
	return &recordServiceServer{service: service}
}
