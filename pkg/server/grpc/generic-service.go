package grpc

import (
	"context"
	"github.com/tislib/apibrew/pkg/abs"
	"github.com/tislib/apibrew/pkg/model"
	"github.com/tislib/apibrew/pkg/resources/mapping"
	"github.com/tislib/apibrew/pkg/service/annotations"
	"github.com/tislib/apibrew/pkg/stub"
	"github.com/tislib/apibrew/pkg/util"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

type GenericGrpcService interface {
	stub.GenericServer
}

type genericServer struct {
	stub.GenericServer
	service abs.RecordService
}

func (g *genericServer) Create(ctx context.Context, request *stub.CreateRequest) (*stub.CreateResponse, error) {
	records, err := g.itemsToRecords(request.Items)

	if err != nil {
		return nil, err
	}

	records, inserted, serviceErr := g.service.Create(annotations.WithContext(ctx, request), abs.RecordCreateParams{
		Namespace:      request.Namespace,
		Resource:       request.Resource,
		Records:        records,
		IgnoreIfExists: request.IgnoreIfExists,
	})

	items, err := g.recordsToItems(request.Resource, request.Namespace, records)

	if err != nil {
		return nil, err
	}

	return &stub.CreateResponse{
		Items:    items,
		Inserted: inserted,
	}, util.ToStatusError(serviceErr)
}

func (g *genericServer) Update(ctx context.Context, request *stub.UpdateRequest) (*stub.UpdateResponse, error) {
	records, err := g.itemsToRecords(request.Items)

	if err != nil {
		return nil, err
	}

	records, serviceErr := g.service.Update(annotations.WithContext(ctx, request), abs.RecordUpdateParams{
		Namespace:    request.Namespace,
		Resource:     request.Resource,
		Records:      records,
		CheckVersion: request.CheckVersion,
	})

	items, err := g.recordsToItems(request.Resource, request.Namespace, records)

	if err != nil {
		return nil, err
	}

	return &stub.UpdateResponse{
		Items: items,
	}, util.ToStatusError(serviceErr)
}
func (g *genericServer) UpdateMulti(ctx context.Context, request *stub.UpdateMultiRequest) (*stub.UpdateMultiResponse, error) {
	return nil, nil
}

func (g *genericServer) Delete(ctx context.Context, request *stub.DeleteRequest) (*stub.DeleteResponse, error) {
	err := g.service.Delete(annotations.WithContext(ctx, request), abs.RecordDeleteParams{
		Namespace: request.Namespace,
		Resource:  request.Resource,
		Ids:       request.Ids,
	})

	return &stub.DeleteResponse{}, util.ToStatusError(err)
}

func (g *genericServer) List(ctx context.Context, request *stub.ListRequest) (*stub.ListResponse, error) {
	records, total, serviceErr := g.service.List(annotations.WithContext(ctx, request), abs.RecordListParams{
		Namespace:         request.Namespace,
		Resource:          request.Resource,
		Limit:             request.Limit,
		Offset:            request.Offset,
		UseHistory:        request.UseHistory,
		ResolveReferences: request.ResolveReferences,
	})

	items, err := g.recordsToItems(request.Resource, request.Namespace, records)

	if err != nil {
		return nil, err
	}

	return &stub.ListResponse{
		Content: items,
		Total:   total,
	}, util.ToStatusError(serviceErr)
}

func (g *genericServer) Search(ctx context.Context, request *stub.SearchRequest) (*stub.SearchResponse, error) {
	records, total, serviceErr := g.service.List(annotations.WithContext(ctx, request), abs.RecordListParams{
		Namespace:         request.Namespace,
		Resource:          request.Resource,
		Limit:             request.Limit,
		Query:             request.Query,
		Offset:            request.Offset,
		UseHistory:        request.UseHistory,
		ResolveReferences: request.ResolveReferences,
	})

	items, err := g.recordsToItems(request.Resource, request.Namespace, records)

	if err != nil {
		return nil, err
	}

	return &stub.SearchResponse{
		Content: items,
		Total:   total,
	}, util.ToStatusError(serviceErr)
}

func (g *genericServer) Get(ctx context.Context, request *stub.GetRequest) (*stub.GetResponse, error) {
	record, serviceErr := g.service.Get(annotations.WithContext(ctx, request), abs.RecordGetParams{
		Namespace: request.Namespace,
		Resource:  request.Resource,
		Id:        request.Id,
	})

	item := new(anypb.Any)

	message := mapping.MessageFromRecord(request.Resource, request.Namespace, record)

	err := anypb.MarshalFrom(item, message, proto.MarshalOptions{})

	if err != nil {
		return nil, err
	}

	return &stub.GetResponse{
		Item: item,
	}, util.ToStatusError(serviceErr)
}

func (g *genericServer) recordsToItems(resource, namespace string, records []*model.Record) ([]*anypb.Any, error) {
	var items []*anypb.Any

	for _, record := range records {
		item := new(anypb.Any)

		message := mapping.MessageFromRecord(resource, namespace, record)

		err := anypb.MarshalFrom(item, message, proto.MarshalOptions{})

		if err != nil {
			return nil, err
		}

		items = append(items, item)
	}

	return items, nil
}

func (g *genericServer) itemsToRecords(items []*anypb.Any) ([]*model.Record, error) {
	var records []*model.Record
	for _, item := range items {
		message, err := item.UnmarshalNew()

		if err != nil {
			return nil, err
		}

		records = append(records, mapping.MessageToRecord(message))
	}

	return records, nil
}

func NewGenericService(service abs.RecordService) stub.GenericServer {
	return &genericServer{service: service}
}
