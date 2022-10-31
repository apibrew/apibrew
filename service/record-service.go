package service

import (
	"context"
	"data-handler/model"
	"data-handler/service/backend"
	"data-handler/service/errors"
	"data-handler/service/security"
	"data-handler/service/system"
	"data-handler/service/types"
	"data-handler/util"
	"google.golang.org/protobuf/types/known/structpb"
	"strings"
)

type RecordListParams struct {
	Query      *model.BooleanExpression
	Workspace  string
	Resource   string
	Limit      uint32
	Offset     uint64
	UseHistory bool
}

type RecordCreateParams struct {
	Workspace      string
	Resource       string
	Records        []*model.Record
	IgnoreIfExists bool
}

type RecordUpdateParams struct {
	Workspace    string
	Records      []*model.Record
	CheckVersion bool
}

type RecordGetParams struct {
	Workspace string
	Resource  string
	Id        string
}

type RecordDeleteParams struct {
	Workspace string
	Resource  string
	Ids       []string
}

type RecordService interface {
	PrepareQuery(resource *model.Resource, queryMap map[string]interface{}) (*model.BooleanExpression, errors.ServiceError)
	GetRecord(ctx context.Context, workspace, resourceName, id string) (*model.Record, errors.ServiceError)
	FindBy(ctx context.Context, workspace, resourceName, propertyName string, value interface{}) (*model.Record, errors.ServiceError)

	Init(data *model.InitData)
	InjectPostgresResourceServiceBackend(serviceBackend backend.ResourceServiceBackend)
	InjectDataSourceService(service DataSourceService)
	InjectAuthenticationService(service AuthenticationService)
	InjectResourceService(service ResourceService)

	List(ctx context.Context, params RecordListParams) ([]*model.Record, uint32, errors.ServiceError)
	Create(ctx context.Context, params RecordCreateParams) ([]*model.Record, []bool, errors.ServiceError)
	Update(ctx context.Context, params RecordUpdateParams) ([]*model.Record, errors.ServiceError)
	Get(ctx context.Context, params RecordGetParams) (*model.Record, errors.ServiceError)
	Delete(ctx context.Context, params RecordDeleteParams) errors.ServiceError
}

type recordService struct {
	postgresResourceServiceBackend backend.ResourceServiceBackend
	dataSourceService              DataSourceService
	authenticationService          AuthenticationService
	ServiceName                    string
	resourceService                ResourceService
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

func (r *recordService) List(ctx context.Context, params RecordListParams) ([]*model.Record, uint32, errors.ServiceError) {
	resource, err := r.resourceService.GetResourceByName(ctx, params.Workspace, params.Resource)

	if err != nil {
		return nil, 0, err
	}

	if err = security.CheckSystemResourceAccess(ctx, resource); err != nil {
		return nil, 0, err
	}

	records, total, err := r.postgresResourceServiceBackend.ListRecords(backend.ListRecordParams{
		Resource:   resource,
		Query:      params.Query,
		Limit:      params.Limit,
		Offset:     params.Offset,
		UseHistory: params.UseHistory,
	})

	if err != nil {
		return nil, 0, err
	}

	return records, total, err
}

func (r *recordService) PrepareQuery(resource *model.Resource, queryMap map[string]interface{}) (*model.BooleanExpression, errors.ServiceError) {
	var criteria []*model.BooleanExpression
	for _, property := range resource.Properties {
		if queryMap[property.Name] != nil {
			var val *structpb.Value
			val, err := structpb.NewValue(queryMap[property.Name])
			if err != nil {
				return nil, errors.RecordValidationError.WithDetails(err.Error())
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
				return nil, errors.RecordValidationError.WithDetails(err.Error())
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

func (r *recordService) Create(ctx context.Context, params RecordCreateParams) ([]*model.Record, []bool, errors.ServiceError) {
	var entityRecordMap = make(map[string][]*model.Record)

	for _, record := range params.Records {
		entityRecordMap[record.Resource] = append(entityRecordMap[record.Resource], record)
	}

	var result []*model.Record

	for _, item := range params.Records {
		item.Type = model.DataType_USER
	}

	var insertedArray []bool
	var err errors.ServiceError

	for resourceName, list := range entityRecordMap {
		var resource *model.Resource
		resource, err = r.resourceService.GetResourceByName(ctx, params.Workspace, resourceName)

		if err != nil {
			return nil, nil, err
		}

		if err = security.CheckSystemResourceAccess(ctx, resource); err != nil {
			return nil, nil, err
		}

		if err != nil {
			return nil, nil, err
		}

		err = r.validateRecords(resource, list)

		if err != nil {
			return nil, nil, err
		}

		var records []*model.Record
		var inserted bool
		records, inserted, err = r.postgresResourceServiceBackend.AddRecords(backend.BulkRecordsParams{
			Resource:       resource,
			Records:        list,
			IgnoreIfExists: params.IgnoreIfExists,
		})

		insertedArray = append(insertedArray, inserted)

		if err != nil {
			return nil, nil, err
		}

		result = append(result, records...)
	}

	return result, insertedArray, nil
}

func (r *recordService) Update(ctx context.Context, params RecordUpdateParams) ([]*model.Record, errors.ServiceError) {
	var entityRecordMap = make(map[string][]*model.Record)

	for _, record := range params.Records {
		entityRecordMap[record.Resource] = append(entityRecordMap[record.Resource], record)
	}

	var result []*model.Record
	var err errors.ServiceError

	for resourceName, list := range entityRecordMap {
		var resource *model.Resource
		if resource, err = r.resourceService.GetResourceByName(ctx, params.Workspace, resourceName); err != nil {
			return nil, err
		}

		if err = security.CheckSystemResourceAccess(ctx, resource); err != nil {
			return nil, err
		}

		if resource.Flags.KeepHistory && !params.CheckVersion {
			return nil, errors.RecordValidationError.WithMessage("checkVersion must be enabled if resource has keepHistory enabled")
		}

		err = r.validateRecords(resource, list)

		if err != nil {
			return nil, err
		}

		var records []*model.Record
		records, err = r.postgresResourceServiceBackend.UpdateRecords(backend.BulkRecordsParams{
			Resource:     resource,
			Records:      list,
			CheckVersion: params.CheckVersion,
		})

		if err != nil {
			return nil, err
		}

		result = append(result, records...)
	}

	return result, nil
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

func (r *recordService) Get(ctx context.Context, params RecordGetParams) (*model.Record, errors.ServiceError) {
	return r.GetRecord(ctx, params.Workspace, params.Resource, params.Id)
}

func (r *recordService) Delete(ctx context.Context, params RecordDeleteParams) errors.ServiceError {
	resource, err := r.resourceService.GetResourceByName(ctx, params.Workspace, params.Resource)

	if err != nil {
		return err
	}

	if err = security.CheckSystemResourceAccess(ctx, resource); err != nil {
		return err
	}

	return r.postgresResourceServiceBackend.DeleteRecords(resource, params.Ids)
}

func (r *recordService) Init(data *model.InitData) {

}

func (r *recordService) validateRecords(resource *model.Resource, list []*model.Record) errors.ServiceError {
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

	return errors.RecordValidationError.WithDetails("Validation failed on some fields:" + strings.Join(util.ArrayMap[*model.ErrorField, string](fieldErrors, func(fieldError *model.ErrorField) string {
		return fieldError.Property + ":" + fieldError.Message
	}), ";")).WithErrorFields(fieldErrors)
}

func NewRecordService() RecordService {
	return &recordService{ServiceName: "RecordService"}
}
