package service

import (
	"context"
	log "github.com/sirupsen/logrus"
	"github.com/tislib/data-handler/pkg/abs"
	"github.com/tislib/data-handler/pkg/errors"
	"github.com/tislib/data-handler/pkg/logging"
	"github.com/tislib/data-handler/pkg/model"
	"github.com/tislib/data-handler/pkg/resources"
	"github.com/tislib/data-handler/pkg/service/annotations"
	"github.com/tislib/data-handler/pkg/util"
	"google.golang.org/protobuf/types/known/structpb"
)

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

	if err := r.genericHandler.BeforeList(ctx, resource, params); err != nil {
		return nil, 0, err
	}

	if handled, records, total, err := r.genericHandler.List(ctx, resource, params); handled {
		return records, total, err
	}

	if resource.Virtual {
		return nil, 0, virtualResourceBackendAccessError
	}

	bck, err := r.backendServiceProvider.GetBackendByDataSourceName(ctx, resource.GetSourceConfig().DataSource)

	if err != nil {
		return nil, 0, err
	}

	records, total, err := bck.ListRecords(ctx, abs.ListRecordParams{
		Resource:          resource,
		Query:             params.Query,
		Limit:             params.Limit,
		Offset:            params.Offset,
		UseHistory:        params.UseHistory,
		ResolveReferences: params.ResolveReferences,
		Schema:            r.resourceService.GetSchema(),
		PackRecords:       params.PackRecords,
		ResultChan:        params.ResultChan,
	})

	if err != nil {
		return nil, 0, err
	}

	if err := checkAccess(ctx, checkAccessParams{
		Resource:  resource,
		Records:   &records,
		Operation: model.OperationType_OPERATION_TYPE_READ,
	}); err != nil {
		return nil, 0, err
	}

	if err = r.genericHandler.AfterList(ctx, resource, params, records, total); err != nil {
		return nil, 0, err
	}

	return records, total, err
}

func (r *recordService) Create(ctx context.Context, params abs.RecordCreateParams) ([]*model.Record, []bool, errors.ServiceError) {
	var result []*model.Record

	var insertedArray []bool
	var err errors.ServiceError

	var success = true
	var txCtx context.Context

	if params.Resource == "" {
		return nil, nil, errors.RecordValidationError.WithMessage("Resource name is empty")
	}

	resource := r.resourceService.GetResourceByName(ctx, params.Namespace, params.Resource)

	if resource == nil {
		success = false
		return nil, nil, errors.ResourceNotFoundError
	}

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

	if err = r.validateRecords(resource, params.Records, false); err != nil {
		return nil, nil, err
	}

	if err = r.genericHandler.BeforeCreate(ctx, resource, params); err != nil {
		return nil, nil, err
	}

	// prepare default values
	var defaultValueMap = make(map[string]*structpb.Value)
	for _, prop := range resource.Properties {
		if prop.DefaultValue != nil {
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
	var inserted bool

	if handled, records, inserted, err := r.genericHandler.Create(ctx, resource, params); handled {
		return records, inserted, err
	}

	if resource.Virtual {
		return nil, nil, virtualResourceBackendAccessError
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

	records, inserted, err = bck.AddRecords(txCtx, abs.BulkRecordsParams{
		Resource:       resource,
		Records:        params.Records,
		IgnoreIfExists: params.IgnoreIfExists,
		Schema:         r.resourceService.GetSchema(),
	})

	insertedArray = append(insertedArray, inserted)

	if err != nil {
		success = false
		return nil, nil, err
	}

	if err = r.genericHandler.AfterCreate(ctx, resource, params, records); err != nil {
		success = false
		return nil, nil, err
	}

	result = append(result, records...)

	return result, insertedArray, nil
}

func isResourceRelatedResource(resource *model.Resource) bool {
	return resource.Namespace == resources.ResourceResource.Namespace && (resource.Name == resources.ResourceResource.Name || resource.Name == resources.ResourcePropertyResource.Name)
}

func (r *recordService) Update(ctx context.Context, params abs.RecordUpdateParams) ([]*model.Record, errors.ServiceError) {
	var result []*model.Record
	var err errors.ServiceError

	var success = true

	if params.Resource == "" {
		return nil, errors.RecordValidationError.WithMessage("Resource name is empty")
	}

	resource := r.resourceService.GetResourceByName(ctx, params.Namespace, params.Resource)

	if resource == nil {
		return nil, errors.RecordValidationError.WithMessage("Resource not found with name: " + params.Resource)
	}

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

	if annotations.IsEnabled(resource, annotations.KeepHistory) && !params.CheckVersion {
		success = false
		return nil, errors.RecordValidationError.WithMessage("checkVersion must be enabled if resource has keepHistory enabled")
	}

	err = r.validateRecords(resource, params.Records, true)

	if err != nil {
		success = false
		return nil, err
	}

	// prevent immutable properties to be updated

	var immutableColsMap = make(map[string]bool)
	for _, prop := range resource.Properties {
		if prop.Immutable {
			immutableColsMap[prop.Name] = true
		}
	}

	for _, record := range params.Records {
		for key := range immutableColsMap {
			delete(record.Properties, key)
		}
	}

	if err = r.genericHandler.BeforeUpdate(ctx, resource, params); err != nil {
		success = false
		return nil, err
	}

	if handled, records, err := r.genericHandler.Update(ctx, resource, params); handled {
		success = false
		return records, err
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

	records, err = bck.UpdateRecords(txCtx, abs.BulkRecordsParams{
		Resource:     resource,
		Records:      params.Records,
		CheckVersion: params.CheckVersion,
		Schema:       r.resourceService.GetSchema(),
	})

	if err != nil {
		success = false
		return nil, err
	}

	if err = r.genericHandler.AfterUpdate(ctx, resource, params, records); err != nil {
		return nil, err
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

	if err := r.genericHandler.BeforeGet(ctx, resource, id); err != nil {
		return nil, err
	}

	if handled, res, err := r.genericHandler.Get(ctx, resource, id); handled {
		return res, err
	}

	if resource.Virtual {
		return nil, virtualResourceBackendAccessError
	}

	bck, err := r.backendServiceProvider.GetBackendByDataSourceName(ctx, resource.GetSourceConfig().DataSource)

	if err != nil {
		return nil, err
	}

	res, err := bck.GetRecord(ctx, resource, r.resourceService.GetSchema(), id)

	if err != nil {
		return nil, err
	}

	if err = r.genericHandler.AfterGet(ctx, resource, id, res); err != nil {
		return nil, err
	}

	return res, err
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
	query, err := PrepareQuery(resource, queryMap)
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

	if err := r.genericHandler.BeforeDelete(ctx, resource, params); err != nil {
		return err
	}

	if handled, err := r.genericHandler.Delete(ctx, resource, params); handled {
		return err
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

	if err = r.genericHandler.AfterDelete(ctx, resource, params); err != nil {
		return err
	}

	return nil
}
