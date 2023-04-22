package rest

import (
	"context"
	"github.com/tislib/apibrew/pkg/abs"
	"github.com/tislib/apibrew/pkg/model"
	"github.com/tislib/apibrew/pkg/service/annotations"
	"github.com/tislib/apibrew/pkg/stub/rest"
	"github.com/tislib/apibrew/pkg/util"
)

type recordService struct {
	rest.RecordServer
	service abs.RecordService
}

func (r *recordService) Create(ctx context.Context, request *rest.CreateRecordRequest) (*rest.CreateRecordResponse, error) {
	records, _, err := r.service.Create(annotations.WithContext(ctx, request), abs.RecordCreateParams{
		Namespace: request.Namespace,
		Resource:  request.Resource,
		Records: []*model.Record{{
			Properties: request.Properties,
		}},
	})

	if err != nil {
		return nil, util.ToStatusError(err)
	}

	return &rest.CreateRecordResponse{
		Id:         records[0].Id,
		Properties: records[0].Properties,
	}, util.ToStatusError(err)
}

func (r *recordService) Apply(ctx context.Context, request *rest.ApplyRecordRequest) (*rest.ApplyRecordResponse, error) {
	records, err := r.service.Apply(annotations.WithContext(ctx, request), abs.RecordUpdateParams{
		Namespace: request.Namespace,
		Resource:  request.Resource,
		Records: []*model.Record{{
			Properties: request.Properties,
		}},
	})

	if err != nil {
		return nil, util.ToStatusError(err)
	}

	return &rest.ApplyRecordResponse{
		Properties: records[0].Properties,
	}, util.ToStatusError(err)
}

func (r *recordService) Update(ctx context.Context, request *rest.UpdateRecordRequest) (*rest.UpdateRecordResponse, error) {
	records, err := r.service.Update(annotations.WithContext(ctx, request), abs.RecordUpdateParams{
		Namespace: request.Namespace,
		Resource:  request.Resource,
		Records: []*model.Record{{
			Id:         request.Id,
			Properties: request.Properties,
		}},
	})

	if err != nil {
		return nil, util.ToStatusError(err)
	}

	return &rest.UpdateRecordResponse{
		Properties: records[0].Properties,
	}, util.ToStatusError(err)
}

func (r *recordService) Delete(ctx context.Context, request *rest.DeleteRecordRequest) (*rest.DeleteRecordResponse, error) {
	err := r.service.Delete(annotations.WithContext(ctx, request), abs.RecordDeleteParams{
		Namespace: request.Namespace,
		Resource:  request.Resource,
		Ids:       []string{request.Id},
	})

	if err != nil {
		return nil, util.ToStatusError(err)
	}

	return &rest.DeleteRecordResponse{}, util.ToStatusError(err)
}

func newRecordService(service abs.RecordService) rest.RecordServer {
	return &recordService{service: service}
}
