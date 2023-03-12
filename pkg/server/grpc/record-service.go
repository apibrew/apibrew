package grpc

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/tislib/data-handler/pkg/abs"
	"github.com/tislib/data-handler/pkg/model"
	"github.com/tislib/data-handler/pkg/service/annotations"
	"github.com/tislib/data-handler/pkg/stub"
	"github.com/tislib/data-handler/pkg/util"
)

type recordServer struct {
	stub.RecordServer
	service               abs.RecordService
	authenticationService abs.AuthenticationService
}

func (r *recordServer) List(ctx context.Context, request *stub.ListRecordRequest) (*stub.ListRecordResponse, error) {
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
	}, util.ToStatusError(err)
}

func (r *recordServer) Search(ctx context.Context, request *stub.SearchRecordRequest) (*stub.SearchRecordResponse, error) {
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
	}, util.ToStatusError(err)
}

func (r *recordServer) ReadStream(request *stub.ReadStreamRequest, resp stub.Record_ReadStreamServer) error {
	ictx, err := interceptRequest(r.authenticationService, resp.Context(), request)
	ctx, cancel := context.WithCancel(ictx)

	defer func() {
		log.Println("cancelled")
		cancel()
	}()

	if err != nil {
		return err
	}

	resultChan := make(chan *model.Record, 100)

	defer func() {
		log.Print("Closing chan")
		close(resultChan)
	}()

	go func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in f", r)
				cancel()
				close(resultChan)
			}
		}()

		for record := range resultChan {
			err2 := resp.Send(record)

			if err2 != nil {
				cancel()
				break
			}
		}
	}()

	_, _, err = r.service.List(annotations.WithContext(ctx, request), abs.RecordListParams{
		Namespace:         request.Namespace,
		Resource:          request.Resource,
		Limit:             request.Limit,
		Query:             request.Query,
		Offset:            request.Offset,
		UseHistory:        request.UseHistory,
		ResolveReferences: request.ResolveReferences,
		PackRecords:       request.PackRecords,
		ResultChan:        resultChan,
	})

	if err != nil {
		return err
	}

	return nil
}

func (r *recordServer) Create(ctx context.Context, request *stub.CreateRecordRequest) (*stub.CreateRecordResponse, error) {
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
	}, util.ToStatusError(err)
}

func (r *recordServer) Update(ctx context.Context, request *stub.UpdateRecordRequest) (*stub.UpdateRecordResponse, error) {
	records, err := r.service.Update(annotations.WithContext(ctx, request), abs.RecordUpdateParams{
		Namespace:    request.Namespace,
		Resource:     request.Resource,
		Records:      util.ArrayPrepend(request.Records, request.Record),
		CheckVersion: request.CheckVersion,
	})

	return &stub.UpdateRecordResponse{
		Record:  util.ArrayFirst(records),
		Records: records,
	}, util.ToStatusError(err)
}

func (r *recordServer) Get(ctx context.Context, request *stub.GetRecordRequest) (*stub.GetRecordResponse, error) {
	record, err := r.service.Get(annotations.WithContext(ctx, request), abs.RecordGetParams{
		Namespace: request.Namespace,
		Resource:  request.Resource,
		Id:        request.Id,
	})

	return &stub.GetRecordResponse{
		Record: record,
	}, util.ToStatusError(err)
}

func (r *recordServer) Delete(ctx context.Context, request *stub.DeleteRecordRequest) (*stub.DeleteRecordResponse, error) {
	err := r.service.Delete(annotations.WithContext(ctx, request), abs.RecordDeleteParams{
		Namespace: request.Namespace,
		Resource:  request.Resource,
		Ids:       request.Ids,
	})

	return &stub.DeleteRecordResponse{}, util.ToStatusError(err)
}

func NewRecordServer(service abs.RecordService, authenticationService abs.AuthenticationService) stub.RecordServer {
	return &recordServer{service: service, authenticationService: authenticationService}
}
