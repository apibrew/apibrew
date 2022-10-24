package service

import (
	"context"
	"data-handler/service/backend"
	"data-handler/service/errors"
	"data-handler/service/security"
	"data-handler/service/system"
	"data-handler/service/types"
	"data-handler/stub"
	"data-handler/stub/model"
	"google.golang.org/protobuf/types/known/structpb"
)

type RecordServiceInternal interface {
	stub.RecordServiceServer
	PrepareQuery(resource *model.Resource, queryMap map[string]interface{}) (*model.BooleanExpression, errors.ServiceError)
	GetRecord(ctx context.Context, workspace, resourceName, id string) (*model.Record, errors.ServiceError)
	FindBy(ctx context.Context, workspace, resourceName, propertyName string, value interface{}) (*model.Record, errors.ServiceError)
}

type RecordService interface {
	stub.RecordServiceServer
	RecordServiceInternal
	Init(data *model.InitData)
	InjectPostgresResourceServiceBackend(serviceBackend backend.ResourceServiceBackend)
	InjectDataSourceService(service DataSourceService)
	InjectAuthenticationService(service AuthenticationService)
	InjectResourceService(service ResourceService)
}

type recordService struct {
	stub.RecordServiceServer
	postgresResourceServiceBackend backend.ResourceServiceBackend
	dataSourceService              DataSourceServiceInternal
	authenticationService          AuthenticationServiceInternal
	ServiceName                    string
	resourceService                ResourceServiceInternal
}

func (r *recordService) InjectAuthenticationService(service AuthenticationService) {
	r.authenticationService = service
}

func (r *recordService) InjectResourceService(service ResourceService) {
	r.resourceService = service
}

func (r *recordService) InjectDataSourceService(service DataSourceService) {
	r.dataSourceService = service
}

func (r *recordService) InjectPostgresResourceServiceBackend(resourceServiceBackend backend.ResourceServiceBackend) {
	r.postgresResourceServiceBackend = resourceServiceBackend
}

func (r *recordService) List(ctx context.Context, request *stub.ListRecordRequest) (*stub.ListRecordResponse, error) {
	resource, err := r.resourceService.GetResourceByName(ctx, request.Workspace, request.Resource)

	if err != nil {
		return &stub.ListRecordResponse{
			Error: toProtoError(err),
		}, nil
	}

	if resource == nil {
		return &stub.ListRecordResponse{
			Error: toProtoError(errors.RecordValidationError.WithDetails("resource not found: " + request.Resource)),
		}, nil
	}

	if err = security.CheckSystemResourceAccess(ctx, resource); err != nil {
		return &stub.ListRecordResponse{
			Error: toProtoError(err),
		}, nil
	}

	records, total, err := r.postgresResourceServiceBackend.ListRecords(backend.ListRecordParams{
		Resource:   resource,
		Query:      request.Query,
		Limit:      request.Limit,
		Offset:     request.Offset,
		UseHistory: request.UseHistory,
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

func (r *recordService) PrepareQuery(resource *model.Resource, queryMap map[string]interface{}) (*model.BooleanExpression, errors.ServiceError) {
	var criteria []*model.BooleanExpression
	for _, property := range resource.Properties {
		if queryMap[property.Name] != nil {
			var val *structpb.Value
			val, err := structpb.NewValue(queryMap[property.Name])
			if err != nil {
				return nil, errors.RecordValidationError
			}
			criteria = append(criteria, r.newEqualExpression(property.Name, val))
		}
	}

	var additionalProperties = []string{
		"id", "version",
	}

	for _, property := range additionalProperties {
		if queryMap[property] != nil {
			var val *structpb.Value
			val, err := structpb.NewValue(queryMap[property])
			if err != nil {
				return nil, errors.RecordValidationError
			}
			criteria = append(criteria, r.newEqualExpression(property, val))
		}
	}

	var query *model.BooleanExpression

	if len(criteria) > 0 {
		query = &model.BooleanExpression{Expression: &model.BooleanExpression_And{And: &model.CompoundBooleanExpression{Expressions: criteria}}}
	}
	return query, nil
}

func (r *recordService) newEqualExpression(propertyName string, val *structpb.Value) *model.BooleanExpression {
	return &model.BooleanExpression{
		Expression: &model.BooleanExpression_Equal{
			Equal: &model.PairExpression{
				Left: &model.Expression{
					Expression: &model.Expression_Property{
						Property: propertyName,
					},
				},
				Right: &model.Expression{
					Expression: &model.Expression_Value{
						Value: val,
					},
				},
			},
		},
	}
}

func (r *recordService) Create(ctx context.Context, request *stub.CreateRecordRequest) (*stub.CreateRecordResponse, error) {
	var entityRecordMap = make(map[string][]*model.Record)

	for _, record := range request.Records {
		entityRecordMap[record.Resource] = append(entityRecordMap[record.Resource], record)
	}

	var result []*model.Record

	for _, item := range request.Records {
		item.Type = model.DataType_USER
	}

	var insertedArray []bool
	var err error

	for resourceName, list := range entityRecordMap {
		var resource *model.Resource
		resource, err = r.resourceService.GetResourceByName(ctx, request.Workspace, resourceName)

		if err != nil {
			return &stub.CreateRecordResponse{
				Error: toProtoError(err),
			}, nil
		}

		if resource == nil {
			return &stub.CreateRecordResponse{
				Error: toProtoError(errors.RecordValidationError.WithDetails("resource not found: " + resourceName)),
			}, nil
		}

		if err = security.CheckSystemResourceAccess(ctx, resource); err != nil {
			return nil, err
		}

		if err != nil {
			return &stub.CreateRecordResponse{
				Error: toProtoError(err),
			}, nil
		}

		err = r.validateRecords(resource, list)

		if err != nil {
			return &stub.CreateRecordResponse{
				Error: toProtoError(err),
			}, nil
		}

		var records []*model.Record
		var inserted bool
		records, inserted, err = r.postgresResourceServiceBackend.AddRecords(backend.BulkRecordsParams{
			Resource:       resource,
			Records:        list,
			IgnoreIfExists: request.IgnoreIfExists,
		})

		insertedArray = append(insertedArray, inserted)

		if err != nil {
			return &stub.CreateRecordResponse{
				Error: toProtoError(err),
			}, nil
		}

		result = append(result, records...)
	}

	return &stub.CreateRecordResponse{
		Records:  result,
		Error:    nil,
		Inserted: insertedArray,
	}, nil
}

func (r *recordService) Update(ctx context.Context, request *stub.UpdateRecordRequest) (*stub.UpdateRecordResponse, error) {
	var entityRecordMap = make(map[string][]*model.Record)

	for _, record := range request.Records {
		entityRecordMap[record.Resource] = append(entityRecordMap[record.Resource], record)
	}

	var result []*model.Record
	var err error

	for resourceName, list := range entityRecordMap {
		var resource *model.Resource
		resource, err = r.resourceService.GetResourceByName(ctx, request.Workspace, resourceName)

		if err != nil {
			return &stub.UpdateRecordResponse{
				Error: toProtoError(err),
			}, nil
		}

		if resource == nil {
			return &stub.UpdateRecordResponse{
				Error: toProtoError(errors.RecordValidationError.WithDetails("resource not found: " + resourceName)),
			}, nil
		}

		if err = security.CheckSystemResourceAccess(ctx, resource); err != nil {
			return &stub.UpdateRecordResponse{
				Error: toProtoError(err),
			}, nil
		}

		if err != nil {
			return &stub.UpdateRecordResponse{
				Error: toProtoError(err),
			}, nil
		}

		if resource.Flags.KeepHistory && !request.CheckVersion {
			return &stub.UpdateRecordResponse{
				Error: toProtoError(errors.RecordValidationError.WithMessage("checkVersion must be enabled if resource has keepHistory enabled")),
			}, nil
		}

		err = r.validateRecords(resource, list)

		if err != nil {
			return &stub.UpdateRecordResponse{
				Error: toProtoError(err),
			}, nil
		}

		var records []*model.Record
		records, err = r.postgresResourceServiceBackend.UpdateRecords(backend.BulkRecordsParams{
			Resource:     resource,
			Records:      list,
			CheckVersion: request.CheckVersion,
		})

		if err != nil {
			return &stub.UpdateRecordResponse{
				Error: toProtoError(err),
			}, nil
		}

		result = append(result, records...)
	}

	return &stub.UpdateRecordResponse{
		Records: result,
		Error:   nil,
	}, nil
}

func (r *recordService) GetRecord(ctx context.Context, workspace, resourceName, id string) (*model.Record, errors.ServiceError) {
	resource, err := r.resourceService.GetResourceByName(ctx, workspace, resourceName)

	if err != nil {
		return nil, err
	}

	if err = security.CheckSystemResourceAccess(ctx, resource); err != nil {
		return nil, err
	}

	return r.postgresResourceServiceBackend.GetRecord(resource, id)
}

func (r *recordService) FindBy(ctx context.Context, workspace, resourceName, propertyName string, value interface{}) (*model.Record, errors.ServiceError) {
	resource, err := r.resourceService.GetResourceByName(ctx, workspace, resourceName)

	if err != nil {
		return nil, err
	}

	queryMap := make(map[string]interface{})

	queryMap[propertyName] = value

	query, err := r.PrepareQuery(system.UserResource, queryMap)

	if err != nil {
		return nil, err
	}

	res, total, err := r.postgresResourceServiceBackend.ListRecords(backend.ListRecordParams{
		Resource:   resource,
		Query:      query,
		Limit:      1,
		Offset:     0,
		UseHistory: false,
	})

	if total == 0 {
		return nil, errors.NotFoundError
	}

	if total > 1 {
		return nil, errors.LogicalError.WithDetails("We have more than 1 record")
	}

	return res[0], nil
}

func (r *recordService) Get(ctx context.Context, request *stub.GetRecordRequest) (*stub.GetRecordResponse, error) {
	record, err := r.GetRecord(ctx, request.Workspace, request.Resource, request.Id)

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
	resource, err := r.resourceService.GetResourceByName(ctx, request.Workspace, request.Resource)

	if err != nil {
		return &stub.DeleteRecordResponse{
			Error: toProtoError(err),
		}, nil
	}

	if resource == nil {
		return &stub.DeleteRecordResponse{
			Error: toProtoError(errors.RecordValidationError.WithDetails("resource not found: " + request.Resource)),
		}, nil
	}

	if err = security.CheckSystemResourceAccess(ctx, resource); err != nil {
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

func (r *recordService) validateRecords(resource *model.Resource, list []*model.Record) error {
	var fieldErrors []*model.ErrorField

	var resourcePropertyExists = make(map[string]bool)

	for _, property := range resource.Properties {
		resourcePropertyExists[property.Name] = true
	}

	for _, record := range list {
		propertyMap := record.Properties.AsMap()
		for _, property := range resource.Properties {
			propertyType := types.ByResourcePropertyType(property.Type)
			val := propertyMap[property.Name]
			isEmpty := propertyType.IsEmpty(val)

			if property.Required && isEmpty {
				fieldErrors = append(fieldErrors, &model.ErrorField{
					RecordId: record.Id,
					Property: property.Name,
					Message:  "required",
				})
			}

			if !isEmpty {
				err := propertyType.ValidateValue(val)

				if err != nil {
					fieldErrors = append(fieldErrors, &model.ErrorField{
						RecordId: record.Id,
						Property: property.Name,
						Message:  err.Error(),
					})
				}
			}
		}

		for key := range propertyMap {
			if !resourcePropertyExists[key] {
				fieldErrors = append(fieldErrors, &model.ErrorField{
					RecordId: record.Id,
					Property: key,
					Message:  "there are no such property",
				})
			}
		}
	}

	if len(fieldErrors) == 0 {
		return nil
	}

	return errors.RecordValidationError.WithErrorFields(fieldErrors)
}

func NewRecordService() RecordService {
	return &recordService{ServiceName: "RecordService"}
}
