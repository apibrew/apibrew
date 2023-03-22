package service

import (
	"context"
	"github.com/tislib/data-handler/pkg/abs"
	"github.com/tislib/data-handler/pkg/errors"
	"github.com/tislib/data-handler/pkg/model"
	"github.com/tislib/data-handler/pkg/service/handler"
	"github.com/tislib/data-handler/pkg/stub"
	"github.com/tislib/data-handler/pkg/util"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func (d *extensionService) prepareExtensionHandler(extension *model.Extension) *handler.BaseHandler {
	return &handler.BaseHandler{
		Id: extension.Id,
		BeforeList: func(ctx context.Context, resource *model.Resource, params abs.RecordListParams) errors.ServiceError {
			var request = map[string]proto.Message{
				"resource": resource,
				"request":  params.ToRequest(),
				"action":   wrapperspb.String("List"),
			}

			var response = map[string]proto.Message{}

			if extension.Before != nil {
				return util.CoalesceThen[model.ExternalCall](func(externalCall *model.ExternalCall) errors.ServiceError {
					return d.externalService.Call(ctx, externalCall, request, response)
				}, extension.Before.List, extension.Before.All)
			}

			return nil
		},
		List: func(ctx context.Context, resource *model.Resource, params abs.RecordListParams) (handled bool, records []*model.Record, total uint32, err errors.ServiceError) {
			var listRecordResponse = &stub.ListRecordResponse{}

			var request = map[string]proto.Message{
				"resource": resource,
				"request":  params.ToRequest(),
				"action":   wrapperspb.String("List"),
			}

			var response = map[string]proto.Message{
				"response": listRecordResponse,
			}

			if extension.Instead != nil {
				err = util.CoalesceThen(func(externalCall *model.ExternalCall) errors.ServiceError {
					handled = true
					return d.externalService.Call(ctx, externalCall, request, response)
				}, extension.Instead.List, extension.Instead.All)
			}

			return handled, listRecordResponse.Content, listRecordResponse.Total, err
		},
		AfterList: func(ctx context.Context, resource *model.Resource, params abs.RecordListParams, records []*model.Record, total uint32) errors.ServiceError {
			var request = map[string]proto.Message{
				"resource": resource,
				"request":  params.ToRequest(),
				"action":   wrapperspb.String("AfterList"),
				"response": &stub.ListRecordResponse{
					Total:   total,
					Content: records,
				},
			}

			var response = map[string]proto.Message{}

			if extension.After != nil {
				return util.CoalesceThen[model.ExternalCall](func(externalCall *model.ExternalCall) errors.ServiceError {
					return d.externalService.Call(ctx, externalCall, request, response)
				}, extension.After.List, extension.After.All)
			}

			return nil
		},
		BeforeCreate: func(ctx context.Context, resource *model.Resource, params abs.RecordCreateParams) errors.ServiceError {
			var request = map[string]proto.Message{
				"resource": resource,
				"request":  params.ToRequest(),
				"action":   wrapperspb.String("BeforeCreate"),
			}

			var response = map[string]proto.Message{}

			if extension.Before != nil {
				return util.CoalesceThen[model.ExternalCall](func(externalCall *model.ExternalCall) errors.ServiceError {
					return d.externalService.Call(ctx, externalCall, request, response)
				}, extension.Before.Create, extension.Before.All)
			}

			return nil
		},
		Create: func(ctx context.Context, resource *model.Resource, params abs.RecordCreateParams) (handled bool, records []*model.Record, inserted []bool, err errors.ServiceError) {
			var listRecordResponse = &stub.CreateRecordResponse{}

			var request = map[string]proto.Message{
				"resource": resource,
				"request":  params.ToRequest(),
				"action":   wrapperspb.String("Create"),
			}

			var response = map[string]proto.Message{
				"response": listRecordResponse,
			}

			if extension.Instead != nil {
				err = util.CoalesceThen(func(externalCall *model.ExternalCall) errors.ServiceError {
					handled = true
					return d.externalService.Call(ctx, externalCall, request, response)
				}, extension.Instead.Create, extension.Instead.All)
			}

			return handled, listRecordResponse.Records, listRecordResponse.Inserted, err
		},
		AfterCreate: func(ctx context.Context, resource *model.Resource, params abs.RecordCreateParams, records []*model.Record) errors.ServiceError {
			var request = map[string]proto.Message{
				"resource": resource,
				"request":  params.ToRequest(),
				"action":   wrapperspb.String("AfterCreate"),
				"response": &stub.CreateRecordResponse{
					Records: records,
				},
			}

			var response = map[string]proto.Message{}

			if extension.After != nil {
				return util.CoalesceThen[model.ExternalCall](func(externalCall *model.ExternalCall) errors.ServiceError {
					return d.externalService.Call(ctx, externalCall, request, response)
				}, extension.After.Create, extension.After.All)
			}

			return nil
		},
		BeforeUpdate: func(ctx context.Context, resource *model.Resource, params abs.RecordUpdateParams) errors.ServiceError {
			var request = map[string]proto.Message{
				"resource": resource,
				"request":  params.ToRequest(),
				"action":   wrapperspb.String("BeforeUpdate"),
			}

			var response = map[string]proto.Message{}

			if extension.Before != nil {
				return util.CoalesceThen[model.ExternalCall](func(externalCall *model.ExternalCall) errors.ServiceError {
					return d.externalService.Call(ctx, externalCall, request, response)
				}, extension.Before.Update, extension.Before.All)
			}

			return nil
		},
		Update: func(ctx context.Context, resource *model.Resource, params abs.RecordUpdateParams) (handled bool, records []*model.Record, err errors.ServiceError) {
			var listRecordResponse = &stub.UpdateRecordResponse{}

			var request = map[string]proto.Message{
				"resource": resource,
				"request":  params.ToRequest(),
				"action":   wrapperspb.String("Update"),
			}

			var response = map[string]proto.Message{
				"response": listRecordResponse,
			}

			if extension.Instead != nil {
				err = util.CoalesceThen(func(externalCall *model.ExternalCall) errors.ServiceError {
					handled = true
					return d.externalService.Call(ctx, externalCall, request, response)
				}, extension.Instead.Update, extension.Instead.All)
			}

			return handled, listRecordResponse.Records, err
		},
		AfterUpdate: func(ctx context.Context, resource *model.Resource, params abs.RecordUpdateParams, records []*model.Record) errors.ServiceError {
			var request = map[string]proto.Message{
				"resource": resource,
				"request":  params.ToRequest(),
				"action":   wrapperspb.String("AfterUpdate"),
				"response": &stub.UpdateRecordResponse{
					Records: records,
				},
			}

			var response = map[string]proto.Message{}

			if extension.After != nil {
				return util.CoalesceThen[model.ExternalCall](func(externalCall *model.ExternalCall) errors.ServiceError {
					return d.externalService.Call(ctx, externalCall, request, response)
				}, extension.After.Update, extension.After.All)
			}

			return nil
		},
		BeforeGet: func(ctx context.Context, resource *model.Resource, id string) errors.ServiceError {
			var request = map[string]proto.Message{
				"resource": resource,
				"request": &stub.GetRecordRequest{
					Namespace: resource.Namespace,
					Resource:  resource.Name,
					Id:        id,
				},
				"action": wrapperspb.String("BeforeGet"),
			}

			var response = map[string]proto.Message{}

			if extension.Before != nil {
				return util.CoalesceThen[model.ExternalCall](func(externalCall *model.ExternalCall) errors.ServiceError {
					return d.externalService.Call(ctx, externalCall, request, response)
				}, extension.Before.Get, extension.Before.All)
			}

			return nil
		},
		Get: func(ctx context.Context, resource *model.Resource, id string) (handled bool, record *model.Record, err errors.ServiceError) {
			var listRecordResponse = &stub.GetRecordResponse{}

			var request = map[string]proto.Message{
				"resource": resource,
				"request": &stub.GetRecordRequest{
					Namespace: resource.Namespace,
					Resource:  resource.Name,
					Id:        id,
				},
				"action": wrapperspb.String("Get"),
			}

			var response = map[string]proto.Message{
				"response": listRecordResponse,
			}

			if extension.Instead != nil {
				err = util.CoalesceThen(func(externalCall *model.ExternalCall) errors.ServiceError {
					handled = true
					return d.externalService.Call(ctx, externalCall, request, response)
				}, extension.Instead.Get, extension.Instead.All)
			}

			return handled, listRecordResponse.Record, err
		},
		AfterGet: func(ctx context.Context, resource *model.Resource, id string, record *model.Record) errors.ServiceError {
			var request = map[string]proto.Message{
				"resource": resource,
				"request": &stub.GetRecordRequest{
					Namespace: resource.Namespace,
					Resource:  resource.Name,
					Id:        id,
				},
				"action": wrapperspb.String("AfterGet"),
				"response": &stub.GetRecordResponse{
					Record: record,
				},
			}

			var response = map[string]proto.Message{}

			if extension.After != nil {
				return util.CoalesceThen[model.ExternalCall](func(externalCall *model.ExternalCall) errors.ServiceError {
					return d.externalService.Call(ctx, externalCall, request, response)
				}, extension.After.Get, extension.After.All)
			}

			return nil
		},
		BeforeDelete: func(ctx context.Context, resource *model.Resource, params abs.RecordDeleteParams) errors.ServiceError {
			var request = map[string]proto.Message{
				"resource": resource,
				"request":  params.ToRequest(),
				"action":   wrapperspb.String("BeforeDelete"),
			}

			var response = map[string]proto.Message{}

			if extension.Before != nil {
				return util.CoalesceThen[model.ExternalCall](func(externalCall *model.ExternalCall) errors.ServiceError {
					return d.externalService.Call(ctx, externalCall, request, response)
				}, extension.Before.Delete, extension.Before.All)
			}

			return nil
		},
		Delete: func(ctx context.Context, resource *model.Resource, params abs.RecordDeleteParams) (handled bool, err errors.ServiceError) {
			var listRecordResponse = &stub.DeleteRecordResponse{}

			var request = map[string]proto.Message{
				"resource": resource,
				"request":  params.ToRequest(),
				"action":   wrapperspb.String("Delete"),
			}

			var response = map[string]proto.Message{
				"response": listRecordResponse,
			}

			if extension.Instead != nil {
				err = util.CoalesceThen(func(externalCall *model.ExternalCall) errors.ServiceError {
					handled = true
					return d.externalService.Call(ctx, externalCall, request, response)
				}, extension.Instead.Delete, extension.Instead.All)
			}

			return handled, err
		},
		AfterDelete: func(ctx context.Context, resource *model.Resource, params abs.RecordDeleteParams) errors.ServiceError {
			var request = map[string]proto.Message{
				"resource": resource,
				"request":  params.ToRequest(),
				"action":   wrapperspb.String("AfterDelete"),
				"response": &stub.DeleteRecordResponse{},
			}

			var response = map[string]proto.Message{}

			if extension.After != nil {
				return util.CoalesceThen[model.ExternalCall](func(externalCall *model.ExternalCall) errors.ServiceError {
					return d.externalService.Call(ctx, externalCall, request, response)
				}, extension.After.Delete, extension.After.All)
			}

			return nil
		},
	}
}
