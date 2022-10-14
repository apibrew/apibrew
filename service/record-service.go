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
}

type recordService struct {
	stub.RecordServiceServer
	postgresResourceServiceBackend backend.ResourceServiceBackend
	dataSourceService              DataSourceService
}

func (r *recordService) InjectDataSourceService(service DataSourceService) {
	r.dataSourceService = service
}

func (r *recordService) InjectPostgresResourceServiceBackend(resourceServiceBackend backend.ResourceServiceBackend) {
	r.postgresResourceServiceBackend = resourceServiceBackend
}

func (r *recordService) Create(ctx context.Context, request *stub.CreateRecordRequest) (*stub.CreateRecordResponse, error) {
	var entityRecordMap = make(map[string][]*model.Record)

	for _, record := range request.Records {
		entityRecordMap[record.Resource] = append(entityRecordMap[record.Resource], record)
	}

	var result []*model.Record

	for resourceName, list := range entityRecordMap {
		resource, err := r.postgresResourceServiceBackend.GetResourceByName(resourceName)

		if err != nil {
			return nil, err
		}

		if err != nil {
			return nil, err
		}

		record, err := r.postgresResourceServiceBackend.AddRecords(backend.AddRecordsParams{
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

func (r *recordService) Get(ctx context.Context, request *stub.GetRecordRequest) (*stub.GetRecordResponse, error) {
	resource, err := r.postgresResourceServiceBackend.GetResourceByName(request.Resource)

	if err != nil {
		return nil, err
	}

	record, err := r.postgresResourceServiceBackend.GetRecord(resource, request.Id)

	if err != nil {
		return nil, err
	}

	return &stub.GetRecordResponse{
		Record: record,
		Error:  nil,
	}, nil
}

func (r *recordService) Delete(ctx context.Context, request *stub.DeleteRecordRequest) (*stub.DeleteRecordResponse, error) {
	resource, err := r.postgresResourceServiceBackend.GetResourceByName(request.Resource)

	if err != nil {
		return nil, err
	}

	err = r.postgresResourceServiceBackend.DeleteRecords(resource, request.Ids)

	if err != nil {
		return nil, err
	}

	return &stub.DeleteRecordResponse{}, nil
}

func (r *recordService) Init(data *model.InitData) {
}

func NewRecordService() RecordService {
	return &recordService{}
}
