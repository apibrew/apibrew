package service

import (
	"context"
	log "github.com/sirupsen/logrus"
	"github.com/tislib/apibrew/pkg/abs"
	"github.com/tislib/apibrew/pkg/errors"
	"github.com/tislib/apibrew/pkg/helper"
	"github.com/tislib/apibrew/pkg/logging"
	"github.com/tislib/apibrew/pkg/model"
	"github.com/tislib/apibrew/pkg/resources"
	"github.com/tislib/apibrew/pkg/service/annotations"
	"github.com/tislib/apibrew/pkg/util"
	"google.golang.org/protobuf/types/known/structpb"
)

type recordService struct {
	ServiceName            string
	resourceService        abs.ResourceService
	backendServiceProvider abs.BackendProviderService
}

func (r *recordService) PrepareQuery(resource *model.Resource, queryMap map[string]interface{}) (*model.BooleanExpression, errors.ServiceError) {
	return util.PrepareQuery(resource, queryMap)
}

func NewRecordService(resourceService abs.ResourceService, backendProviderService abs.BackendProviderService) abs.RecordService {
	return &recordService{
		ServiceName:            "RecordService",
		resourceService:        resourceService,
		backendServiceProvider: backendProviderService,
	}
}

var virtualResourceBackendAccessError = errors.LogicalError.WithMessage("Virtual resource is trying to access real backend")

func (r *recordService) List(ctx context.Context, params abs.RecordListParams) ([]*model.Record, uint32, errors.ServiceError) {
	resource := r.resourceService.GetResourceByName(ctx, params.Namespace, params.Resource)

	if resource == nil {
		return nil, 0, errors.ResourceNotFoundError
	}

	if err := checkAccess(ctx, checkAccessParams{
		Resource:  resource,
		Operation: model.OperationType_OPERATION_TYPE_READ,
	}); err != nil {
		return nil, 0, err
	}

	if resource.Virtual {
		return nil, 0, virtualResourceBackendAccessError
	}

	bck, err := r.backendServiceProvider.GetBackendByDataSourceName(ctx, resource.GetSourceConfig().DataSource)

	if err != nil {
		return nil, 0, err
	}

	if params.UseHistory {
		if !annotations.IsEnabled(resource, annotations.KeepHistory) {
			return nil, 0, errors.LogicalError.WithDetails("History is not enabled on resource")
		}
		resource = util.HistoryResource(resource)
	}

	if params.Query != nil && params.Filters != nil {
		return nil, 0, errors.LogicalError.WithDetails("Both query and filters cannot be set at the same time")
	}

	if params.Query == nil && params.Filters != nil {
		var err errors.ServiceError

		params.Query, err = util.PrepareQueryFromFilters(resource, params.Filters)

		if err != nil {
			return nil, 0, err
		}
	}

	records, total, err := bck.ListRecords(ctx, resource, abs.ListRecordParams{
		Query:             params.Query,
		Limit:             params.Limit,
		Offset:            params.Offset,
		ResolveReferences: params.ResolveReferences,
	}, params.ResultChan)

	// todo implement params.PackRecords

	if err != nil {
		return nil, 0, err
	}

	for _, record := range records {
		util.DeNormalizeRecord(resource, record)
	}

	if err = checkAccess(ctx, checkAccessParams{
		Resource:  resource,
		Records:   &records,
		Operation: model.OperationType_OPERATION_TYPE_READ,
	}); err != nil {
		return nil, 0, err
	}

	return records, total, nil
}

func (r *recordService) Create(ctx context.Context, params abs.RecordCreateParams) ([]*model.Record, []bool, errors.ServiceError) {
	if params.Resource == "" {
		return nil, nil, errors.RecordValidationError.WithMessage("Resource name is empty")
	}

	resource := r.resourceService.GetResourceByName(ctx, params.Namespace, params.Resource)

	if resource == nil {
		return nil, nil, errors.ResourceNotFoundError
	}

	return r.CreateWithResource(ctx, resource, params)
}

func (r *recordService) CreateWithResource(ctx context.Context, resource *model.Resource, params abs.RecordCreateParams) ([]*model.Record, []bool, errors.ServiceError) {
	var result []*model.Record

	var err errors.ServiceError

	var success = true
	var txCtx context.Context

	if err = checkAccess(ctx, checkAccessParams{
		Resource:  resource,
		Records:   &params.Records,
		Operation: model.OperationType_OPERATION_TYPE_CREATE,
	}); err != nil {
		return nil, nil, err
	}

	if len(params.Records) == 0 {
		return nil, nil, nil
	}

	if isResourceRelatedResource(resource) {
		return nil, nil, errors.LogicalError.WithDetails("resource and related resources cannot be modified from records API")
	}

	for _, record := range params.Records {
		util.InitRecord(ctx, resource, record)
		util.NormalizeRecord(resource, record)
		log.Print("Normalized record: " + record.Id)
	}

	if err = validateRecords(resource, params.Records, false); err != nil {
		return nil, nil, err
	}

	// prepare default values
	var defaultValueMap = make(map[string]*structpb.Value)
	for _, prop := range resource.Properties {
		if prop.DefaultValue != nil && prop.DefaultValue.AsInterface() != nil {
			defaultValueMap[prop.Name] = prop.DefaultValue
		}
	}
	// set default values
	if len(defaultValueMap) > 0 {
		for _, record := range params.Records {
			for key, value := range defaultValueMap {
				_, exists := record.Properties[key]

				if !exists {
					record.Properties[key] = value
				}
			}
		}
	}

	var records []*model.Record
	var inserted []bool

	if resource.Virtual {
		return nil, nil, virtualResourceBackendAccessError
	}

	if params.Records == nil {
		return nil, nil, nil
	}

	bck, err := r.backendServiceProvider.GetBackendByDataSourceName(ctx, resource.GetSourceConfig().DataSource)

	if err != nil {
		success = false
		return nil, []bool{}, err
	}

	tx, err := bck.BeginTransaction(ctx, false)

	if err != nil {
		success = false
		return nil, []bool{}, err
	}

	txCtx = context.WithValue(ctx, abs.TransactionContextKey, tx)

	defer func() {
		if success {
			err = bck.CommitTransaction(txCtx)

			if err != nil {
				log.Print(err)
				success = false
			}
		} else {
			err = bck.RollbackTransaction(txCtx)

			if err != nil {
				log.Print(err)
			}
		}
	}()

	records, inserted, err = bck.AddRecords(txCtx, resource, params.Records)

	if annotations.IsEnabled(resource, annotations.KeepHistory) {
		var historyRecords []*model.Record
		historyResource := util.HistoryResource(resource)

		for index, rec := range inserted {
			if rec {
				historyRecords = append(historyRecords, records[index])
			}
		}

		_, _, err = bck.AddRecords(txCtx, historyResource, historyRecords)

		if err != nil {
			success = false
			return nil, nil, err
		}
	}

	if err != nil {
		success = false
		return nil, nil, err
	}

	result = append(result, records...)

	return result, inserted, nil
}

func isResourceRelatedResource(resource *model.Resource) bool {
	return resource.Namespace == resources.ResourceResource.Namespace && (resource.Name == resources.ResourceResource.Name || resource.Name == resources.ResourcePropertyResource.Name)
}

func (r *recordService) Apply(ctx context.Context, params abs.RecordUpdateParams) ([]*model.Record, errors.ServiceError) {
	if params.Resource == "" {
		return nil, errors.RecordValidationError.WithMessage("Resource name is empty")
	}

	resource := r.resourceService.GetResourceByName(ctx, params.Namespace, params.Resource)

	if resource == nil {
		return nil, errors.RecordValidationError.WithMessage("Resource not found with name: " + params.Resource)
	}

	var result []*model.Record

	for _, record := range params.Records {

		// locate existing record
		var existingRecord *model.Record

		identifierProps, err := util.RecordIdentifierProperties(resource, record.Properties)

		if err != nil {
			return nil, errors.RecordValidationError.WithMessage(err.Error())
		}

		qb := helper.NewQueryBuilder()

		searchRes, total, err := r.List(ctx, abs.RecordListParams{
			Namespace: resource.Namespace,
			Resource:  resource.Name,
			Limit:     1,
			Query:     qb.FromProperties(identifierProps),
		})

		if err != nil {
			return nil, errors.RecordValidationError.WithMessage(err.Error())
		}

		if total > 0 {
			existingRecord = searchRes[0]
		}

		if existingRecord == nil {
			records, _, err := r.CreateWithResource(ctx, resource, abs.RecordCreateParams{
				Namespace: resource.Namespace,
				Resource:  resource.Name,
				Records:   []*model.Record{record},
			})

			if err != nil {
				return nil, errors.RecordValidationError.WithMessage(err.Error())
			}

			result = append(result, records...)
		} else {
			record.Id = existingRecord.Id

			if util.IsSameRecord(existingRecord, record) {
				return params.Records, nil
			}

			records, err := r.UpdateWithResource(ctx, resource, abs.RecordUpdateParams{
				Namespace: resource.Namespace,
				Resource:  resource.Name,
				Records:   []*model.Record{record},
			})

			if err != nil {
				return nil, errors.RecordValidationError.WithMessage(err.Error())
			}

			result = append(result, records...)
		}
	}

	return result, nil
}

func (r *recordService) Update(ctx context.Context, params abs.RecordUpdateParams) ([]*model.Record, errors.ServiceError) {
	if params.Resource == "" {
		return nil, errors.RecordValidationError.WithMessage("Resource name is empty")
	}

	resource := r.resourceService.GetResourceByName(ctx, params.Namespace, params.Resource)

	if resource == nil {
		return nil, errors.RecordValidationError.WithMessage("Resource not found with name: " + params.Resource)
	}

	return r.UpdateWithResource(ctx, resource, params)
}

func (r *recordService) UpdateWithResource(ctx context.Context, resource *model.Resource, params abs.RecordUpdateParams) ([]*model.Record, errors.ServiceError) {
	var result []*model.Record
	var err errors.ServiceError

	var success = true

	if isResourceRelatedResource(resource) {
		return nil, errors.LogicalError.WithDetails("resource and related resources cannot be modified from records API")
	}

	if err = checkAccess(ctx, checkAccessParams{
		Resource:  resource,
		Records:   &params.Records,
		Operation: model.OperationType_OPERATION_TYPE_UPDATE,
	}); err != nil {
		return nil, err
	}

	if len(params.Records) == 0 {
		return nil, nil
	}

	if resource.Immutable {
		return nil, errors.RecordValidationError.WithMessage("Immutable resource cannot be modified or deleted: " + params.Resource)
	}

	if annotations.IsEnabled(resource, annotations.KeepHistory) && !annotations.IsEnabledOnCtx(ctx, annotations.CheckVersion) {
		success = false
		return nil, errors.RecordValidationError.WithMessage("checkVersion must be enabled if resource has keepHistory enabled")
	}

	for _, record := range params.Records {
		util.PrepareUpdateForRecord(ctx, resource, record)
		util.NormalizeRecord(resource, record)
	}

	err = validateRecords(resource, params.Records, true)

	if err != nil {
		success = false
		return nil, err
	}

	if resource.Virtual {
		return nil, virtualResourceBackendAccessError
	}

	var records []*model.Record

	bck, err := r.backendServiceProvider.GetBackendByDataSourceName(ctx, resource.GetSourceConfig().DataSource)

	if err != nil {
		success = false
		return nil, err
	}

	tx, err := bck.BeginTransaction(ctx, false)

	if err != nil {
		success = false
		return nil, err
	}

	txCtx := context.WithValue(ctx, abs.TransactionContextKey, tx)

	defer func() {
		if success {
			err = bck.CommitTransaction(txCtx)

			if err != nil {
				log.Print(err)
				success = false
			}
		} else {
			err = bck.RollbackTransaction(txCtx)

			if err != nil {
				log.Print(err)
			}
		}
	}()

	records, err = bck.UpdateRecords(txCtx, resource, params.Records)

	if err != nil {
		success = false
		return nil, err
	}

	if annotations.IsEnabled(resource, annotations.KeepHistory) {
		_, _, err = bck.AddRecords(txCtx, util.HistoryResource(resource), records)

		if err != nil {
			success = false
			return nil, err
		}
	}

	result = append(result, records...)

	return result, nil
}

func (r *recordService) GetRecord(ctx context.Context, namespace, resourceName, id string) (*model.Record, errors.ServiceError) {
	resource := r.resourceService.GetResourceByName(ctx, namespace, resourceName)

	if resource == nil {
		return nil, errors.ResourceNotFoundError
	}

	if isResourceRelatedResource(resource) {
		return nil, errors.LogicalError.WithDetails("resource and related resources cannot be modified from records API")
	}

	if err := checkAccess(ctx, checkAccessParams{
		Resource: resource,
		Records: &[]*model.Record{
			{
				Id: id,
			},
		},
		Operation: model.OperationType_OPERATION_TYPE_READ,
	}); err != nil {
		return nil, err
	}

	if resource.Virtual {
		return nil, virtualResourceBackendAccessError
	}

	bck, err := r.backendServiceProvider.GetBackendByDataSourceName(ctx, resource.GetSourceConfig().DataSource)

	if err != nil {
		return nil, err
	}

	res, err := bck.GetRecord(ctx, resource, id)

	if err != nil {
		return nil, err
	}

	util.DeNormalizeRecord(resource, res)

	return res, nil
}

func (r *recordService) FindBy(ctx context.Context, namespace, resourceName, propertyName string, value interface{}) (*model.Record, errors.ServiceError) {
	logger := log.WithFields(logging.CtxFields(ctx))

	logger.Debug("Begin record-service FindBy")
	defer logger.Debug("Finish record-service FindBy")

	resource := r.resourceService.GetResourceByName(ctx, namespace, resourceName)

	if resource == nil {
		return nil, errors.ResourceNotFoundError
	}

	queryMap := make(map[string]interface{})

	queryMap[propertyName] = value

	logger.Debug("Call PrepareQuery: ", queryMap)
	query, err := util.PrepareQuery(resource, queryMap)
	logger.Debug("Result record-service: ", query)

	if err != nil {
		return nil, err
	}

	res, total, err := r.List(ctx, abs.RecordListParams{
		Query:             query,
		Namespace:         namespace,
		Resource:          resourceName,
		Limit:             2,
		Offset:            0,
		UseHistory:        false,
		ResolveReferences: []string{},
	})

	if err != nil {
		return nil, err
	}

	if total == 0 {
		return nil, errors.RecordNotFoundError
	}

	if total > 1 {
		return nil, errors.LogicalError.WithDetails("We have more than 1 record")
	}

	return res[0], nil
}

func (r *recordService) Get(ctx context.Context, params abs.RecordGetParams) (*model.Record, errors.ServiceError) {
	return r.GetRecord(ctx, params.Namespace, params.Resource, params.Id)
}

func (r *recordService) Delete(ctx context.Context, params abs.RecordDeleteParams) errors.ServiceError {
	resource := r.resourceService.GetResourceByName(ctx, params.Namespace, params.Resource)

	if resource == nil {
		return errors.ResourceNotFoundError
	}

	var recordForCheck = util.ArrayMap(params.Ids, func(t string) *model.Record {
		return &model.Record{
			Id: t,
		}
	})

	if err := checkAccess(ctx, checkAccessParams{
		Resource:  resource,
		Records:   &recordForCheck,
		Operation: model.OperationType_OPERATION_TYPE_DELETE,
	}); err != nil {
		return err
	}

	if isResourceRelatedResource(resource) {
		return errors.LogicalError.WithDetails("resource and related resources cannot be modified from records API")
	}

	if resource.Immutable {
		return errors.RecordValidationError.WithMessage("Immutable resource cannot be modified or deleted: " + params.Resource)
	}

	if resource.Virtual {
		return virtualResourceBackendAccessError
	}

	bck, err := r.backendServiceProvider.GetBackendByDataSourceName(ctx, resource.GetSourceConfig().DataSource)

	if err != nil {
		return err
	}

	if err = bck.DeleteRecords(ctx, resource, params.Ids); err != nil {
		return err
	}

	return nil
}
