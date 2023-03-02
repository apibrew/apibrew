package grpc

import (
	"context"
	log "github.com/sirupsen/logrus"
	"github.com/tislib/data-handler/pkg/abs"
	"github.com/tislib/data-handler/pkg/model"
	util2 "github.com/tislib/data-handler/pkg/server/util"
	"github.com/tislib/data-handler/pkg/service/annotations"
	"github.com/tislib/data-handler/pkg/stub"
	"github.com/tislib/data-handler/pkg/util"
)

type recordServiceServer struct {
	stub.RecordServiceServer
	service               abs.RecordService
	authenticationService abs.AuthenticationService
}

func (r *recordServiceServer) List(ctx context.Context, request *stub.ListRecordRequest) (*stub.ListRecordResponse, error) {
	records, total, err := r.service.List(annotations.WithContext(ctx, request), abs.RecordListParams{
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
	}, util2.ToStatusError(err)
}

func (r *recordServiceServer) Search(ctx context.Context, request *stub.SearchRecordRequest) (*stub.SearchRecordResponse, error) {
	records, total, err := r.service.List(annotations.WithContext(ctx, request), abs.RecordListParams{
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
	}, util2.ToStatusError(err)
}

func (r *recordServiceServer) ReadStream(request *stub.ReadStreamRequest, resp stub.RecordService_ReadStreamServer) error {
	ctx, err := interceptRequest(r.authenticationService, resp.Context(), request)

	if err != nil {
		return err
	}

	resultChan := make(chan *model.Record, 100)

	go func() {
		_, total, err := r.service.List(annotations.WithContext(ctx, request), abs.RecordListParams{
			Namespace:         request.Namespace,
			Resource:          request.Resource,
			Limit:             request.Limit,
			Query:             request.Query,
			Offset:            request.Offset,
			UseHistory:        request.UseHistory,
			ResolveReferences: request.ResolveReferences,
			ResultChan:        resultChan,
		})

		close(resultChan)
		resp.Context().Done()

		if err != nil || total == 0 {
			log.Print(err)
		}
	}()

	for record := range resultChan {
		err2 := resp.Send(record)

		if err2 != nil {
			return err2
		}
	}

	return nil
}

func (r *recordServiceServer) Create(ctx context.Context, request *stub.CreateRecordRequest) (*stub.CreateRecordResponse, error) {
	records, inserted, err := r.service.Create(annotations.WithContext(ctx, request), abs.RecordCreateParams{
		Namespace:      request.Namespace,
		Resource:       request.Resource,
		Records:        util.ArrayPrepend(request.Records, request.Record),
		IgnoreIfExists: request.IgnoreIfExists,
	})

	return &stub.CreateRecordResponse{
		Record:   util.ArrayFirst(records),
		Records:  records,
		Inserted: inserted,
	}, util2.ToStatusError(err)
}

func (r *recordServiceServer) Update(ctx context.Context, request *stub.UpdateRecordRequest) (*stub.UpdateRecordResponse, error) {
	records, err := r.service.Update(annotations.WithContext(ctx, request), abs.RecordUpdateParams{
		Namespace:    request.Namespace,
		Resource:     request.Resource,
		Records:      util.ArrayPrepend(request.Records, request.Record),
		CheckVersion: request.CheckVersion,
	})

	return &stub.UpdateRecordResponse{
		Record:  util.ArrayFirst(records),
		Records: records,
	}, util2.ToStatusError(err)
}

func (r *recordServiceServer) Get(ctx context.Context, request *stub.GetRecordRequest) (*stub.GetRecordResponse, error) {
	record, err := r.service.Get(annotations.WithContext(ctx, request), abs.RecordGetParams{
		Namespace: request.Namespace,
		Resource:  request.Resource,
		Id:        request.Id,
	})

	return &stub.GetRecordResponse{
		Record: record,
	}, util2.ToStatusError(err)
}

func (r *recordServiceServer) Delete(ctx context.Context, request *stub.DeleteRecordRequest) (*stub.DeleteRecordResponse, error) {
	err := r.service.Delete(annotations.WithContext(ctx, request), abs.RecordDeleteParams{
		Namespace: request.Namespace,
		Resource:  request.Resource,
		Ids:       request.Ids,
	})

	return &stub.DeleteRecordResponse{}, util2.ToStatusError(err)
}

func NewRecordServiceServer(service abs.RecordService, authenticationService abs.AuthenticationService) stub.RecordServiceServer {
	return &recordServiceServer{service: service, authenticationService: authenticationService}
}
