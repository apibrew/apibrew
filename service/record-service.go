package service

import (
	"context"
	"data-handler/service/backend"
	"data-handler/stub"
	"data-handler/stub/model"
)

type RecordService interface {
	stub.RecordServiceServer

	Init(data *model.InitData)
	InjectPostgresResourceServiceBackend(serviceBackend backend.ResourceServiceBackend)
	InjectDataSourceService(service DataSourceService)
	InjectAuthenticationService(service AuthenticationService)
}

type recordService struct {
	stub.RecordServiceServer
	postgresResourceServiceBackend backend.ResourceServiceBackend
	dataSourceService              DataSourceService
	authenticationService          AuthenticationService
	ServiceName                    string
}

func (r *recordService) InjectAuthenticationService(service AuthenticationService) {
	r.authenticationService = service
}

func (r *recordService) InjectDataSourceService(service DataSourceService) {
	r.dataSourceService = service
}

func (r *recordService) InjectPostgresResourceServiceBackend(resourceServiceBackend backend.ResourceServiceBackend) {
	r.postgresResourceServiceBackend = resourceServiceBackend
}

func (r *recordService) List(ctx context.Context, request *stub.ListRecordRequest) (*stub.ListRecordResponse, error) {
	err := r.authenticationService.Check(CheckParams{
		Ctx:     ctx,
		Token:   request.Token,
		Service: r.ServiceName,
		Method:  "List",
	})

	if err != nil {
		return &stub.ListRecordResponse{
			Error: toProtoError(err),
		}, nil
	}

	resource, err := r.postgresResourceServiceBackend.GetResourceByName(request.Resource)

	if err != nil {
		return &stub.ListRecordResponse{
			Error: toProtoError(err),
		}, nil
	}

	if err = checkSystemResourceAccess(ctx, resource); err != nil {
		return &stub.ListRecordResponse{
			Error: toProtoError(err),
		}, nil
	}

	records, total, err := r.postgresResourceServiceBackend.ListRecords(backend.ListRecordParams{
		Resource: resource,
		Query:    request.Query,
		Limit:    request.Limit,
		Offset:   request.Offset,
	})

	if err != nil {
		return &stub.ListRecordResponse{
			Error: toProtoError(err),
		}, nil
	}

	return &stub.ListRecordResponse{
		Content: records,
		Total:   total,
		Error:   nil,
	}, nil
}

func (r *recordService) Create(ctx context.Context, request *stub.CreateRecordRequest) (*stub.CreateRecordResponse, error) {
	err := r.authenticationService.Check(CheckParams{
		Ctx:     ctx,
		Token:   request.Token,
		Service: r.ServiceName,
		Method:  "Create",
	})

	if err != nil {
		return nil, err
	}

	var entityRecordMap = make(map[string][]*model.Record)

	for _, record := range request.Records {
		entityRecordMap[record.Resource] = append(entityRecordMap[record.Resource], record)
	}

	var result []*model.Record

	for _, item := range request.Records {
		item.Type = model.DataType_USER
	}

	for resourceName, list := range entityRecordMap {
		resource, err := r.postgresResourceServiceBackend.GetResourceByName(resourceName)

		if err != nil {
			return nil, err
		}

		if err = checkSystemResourceAccess(ctx, resource); err != nil {
			return nil, err
		}

		if err != nil {
			return nil, err
		}

		record, err := r.postgresResourceServiceBackend.AddRecords(backend.BulkRecordsParams{
			Resource: resource,
			Records:  list,
		})

		if err != nil {
			return nil, err
		}

		result = append(result, record...)
	}

	return &stub.CreateRecordResponse{
		Records: result,
		Error:   nil,
	}, nil
}

func (r *recordService) Update(ctx context.Context, request *stub.UpdateRecordRequest) (*stub.UpdateRecordResponse, error) {
	err := r.authenticationService.Check(CheckParams{
		Ctx:     ctx,
		Token:   request.Token,
		Service: r.ServiceName,
		Method:  "Update",
	})

	if err != nil {
		return &stub.UpdateRecordResponse{
			Error: toProtoError(err),
		}, nil
	}

	var entityRecordMap = make(map[string][]*model.Record)

	for _, record := range request.Records {
		entityRecordMap[record.Resource] = append(entityRecordMap[record.Resource], record)
	}

	var result []*model.Record

	for resourceName, list := range entityRecordMap {
		resource, err := r.postgresResourceServiceBackend.GetResourceByName(resourceName)

		if err != nil {
			return &stub.UpdateRecordResponse{
				Error: toProtoError(err),
			}, nil
		}

		if err = checkSystemResourceAccess(ctx, resource); err != nil {
			return &stub.UpdateRecordResponse{
				Error: toProtoError(err),
			}, nil
		}

		if err != nil {
			return nil, err
		}

		record, err := r.postgresResourceServiceBackend.UpdateRecords(backend.BulkRecordsParams{
			Resource:     resource,
			Records:      list,
			CheckVersion: request.CheckVersion,
		})

		if err != nil {
			return &stub.UpdateRecordResponse{
				Error: toProtoError(err),
			}, nil
		}

		result = append(result, record...)
	}

	return &stub.UpdateRecordResponse{
		Records: result,
		Error:   nil,
	}, nil
}

func (r *recordService) Get(ctx context.Context, request *stub.GetRecordRequest) (*stub.GetRecordResponse, error) {
	err := r.authenticationService.Check(CheckParams{
		Ctx:     ctx,
		Token:   request.Token,
		Service: r.ServiceName,
		Method:  "Get",
	})

	if err != nil {
		return nil, err
	}

	resource, err := r.postgresResourceServiceBackend.GetResourceByName(request.Resource)

	if err != nil {
		return &stub.GetRecordResponse{
			Error: toProtoError(err),
		}, nil
	}

	if err = checkSystemResourceAccess(ctx, resource); err != nil {
		return &stub.GetRecordResponse{
			Error: toProtoError(err),
		}, nil
	}

	record, err := r.postgresResourceServiceBackend.GetRecord(resource, request.Id)

	if err != nil {
		return &stub.GetRecordResponse{
			Error: toProtoError(err),
		}, nil
	}

	return &stub.GetRecordResponse{
		Record: record,
		Error:  nil,
	}, nil
}

func (r *recordService) Delete(ctx context.Context, request *stub.DeleteRecordRequest) (*stub.DeleteRecordResponse, error) {
	err := r.authenticationService.Check(CheckParams{
		Ctx:     ctx,
		Token:   request.Token,
		Service: r.ServiceName,
		Method:  "Delete",
	})

	if err != nil {
		return &stub.DeleteRecordResponse{
			Error: toProtoError(err),
		}, nil
	}

	resource, err := r.postgresResourceServiceBackend.GetResourceByName(request.Resource)

	if err != nil {
		return &stub.DeleteRecordResponse{
			Error: toProtoError(err),
		}, nil
	}

	if err = checkSystemResourceAccess(ctx, resource); err != nil {
		return &stub.DeleteRecordResponse{
			Error: toProtoError(err),
		}, nil
	}

	err = r.postgresResourceServiceBackend.DeleteRecords(resource, request.Ids)

	if err != nil {
		return &stub.DeleteRecordResponse{
			Error: toProtoError(err),
		}, nil
	}

	return &stub.DeleteRecordResponse{}, nil
}

func (r *recordService) Init(data *model.InitData) {

}

func NewRecordService() RecordService {
	return &recordService{ServiceName: "RecordService"}
}
