package service

import (
	"context"
	"data-handler/model"
	"data-handler/service/backend"
	"data-handler/service/errors"
	"data-handler/service/params"
	"data-handler/service/security"
)

func (r *recordService) List(ctx context.Context, params params.RecordListParams) ([]*model.Record, uint32, errors.ServiceError) {
	resource, err := r.resourceService.GetResourceByName(ctx, params.Workspace, params.Resource)

	if err != nil {
		return nil, 0, err
	}

	if err = security.CheckSystemResourceAccess(ctx, resource); err != nil {
		return nil, 0, err
	}

	if err = r.genericHandler.BeforeList(ctx, resource, params); err != nil {
		return nil, 0, err
	}

	if handled, records, total, err := r.genericHandler.List(); handled {
		return records, total, err
	}

	records, total, err := r.postgresResourceServiceBackend.ListRecords(backend.ListRecordParams{
		Resource:          resource,
		Query:             params.Query,
		Limit:             params.Limit,
		Offset:            params.Offset,
		UseHistory:        params.UseHistory,
		ResolveReferences: params.ResolveReferences,
	})

	if err != nil {
		return nil, 0, err
	}

	if err = r.genericHandler.AfterList(ctx, resource, params, records, total); err != nil {
		return nil, 0, err
	}

	return records, total, err
}

func (r *recordService) Create(ctx context.Context, params params.RecordCreateParams) ([]*model.Record, []bool, errors.ServiceError) {
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

		if err = r.validateRecords(resource, list); err != nil {
			return nil, nil, err
		}

		if err = r.genericHandler.BeforeCreate(ctx, resource, params); err != nil {
			return nil, nil, err
		}

		var records []*model.Record
		var inserted bool

		if handled, records, inserted, err := r.genericHandler.Create(ctx, resource, params); handled {
			return records, inserted, err
		}

		records, inserted, err = r.postgresResourceServiceBackend.AddRecords(backend.BulkRecordsParams{
			Resource:       resource,
			Records:        list,
			IgnoreIfExists: params.IgnoreIfExists,
		})

		insertedArray = append(insertedArray, inserted)

		if err != nil {
			return nil, nil, err
		}

		if err = r.genericHandler.AfterCreate(ctx, resource, params, records); err != nil {
			return nil, nil, err
		}

		result = append(result, records...)
	}

	return result, insertedArray, nil
}

func (r *recordService) Update(ctx context.Context, params params.RecordUpdateParams) ([]*model.Record, errors.ServiceError) {
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

		if err = r.genericHandler.BeforeUpdate(ctx, resource, params); err != nil {
			return nil, err
		}

		if handled, records, err := r.genericHandler.Update(ctx, resource, params); handled {
			return records, err
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

		if err = r.genericHandler.AfterUpdate(ctx, resource, params, records); err != nil {
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

	if err = r.genericHandler.BeforeGet(resource, id); err != nil {
		return nil, err
	}

	if handled, res, err := r.genericHandler.Get(resource, id); handled {
		return res, err
	}

	res, err := r.postgresResourceServiceBackend.GetRecord(resource, id)

	if err != nil {
		return nil, err
	}

	if err = r.genericHandler.AfterGet(resource, id, res); err != nil {
		return nil, err
	}

	return res, err
}

func (r *recordService) FindBy(ctx context.Context, workspace, resourceName, propertyName string, value interface{}) (*model.Record, errors.ServiceError) {
	resource, err := r.resourceService.GetResourceByName(ctx, workspace, resourceName)

	if err != nil {
		return nil, err
	}

	queryMap := make(map[string]interface{})

	queryMap[propertyName] = value

	query, err := r.PrepareQuery(resource, queryMap)

	if err != nil {
		return nil, err
	}

	res, total, err := r.List(ctx, params.RecordListParams{
		Query:             query,
		Workspace:         workspace,
		Resource:          resourceName,
		Limit:             2,
		Offset:            0,
		UseHistory:        false,
		ResolveReferences: false,
	})

	if total == 0 {
		return nil, errors.NotFoundError
	}

	if total > 1 {
		return nil, errors.LogicalError.WithDetails("We have more than 1 record")
	}

	return res[0], nil
}

func (r *recordService) Get(ctx context.Context, params params.RecordGetParams) (*model.Record, errors.ServiceError) {
	return r.GetRecord(ctx, params.Workspace, params.Resource, params.Id)
}

func (r *recordService) Delete(ctx context.Context, params params.RecordDeleteParams) errors.ServiceError {
	resource, err := r.resourceService.GetResourceByName(ctx, params.Workspace, params.Resource)

	if err != nil {
		return err
	}

	if err = security.CheckSystemResourceAccess(ctx, resource); err != nil {
		return err
	}

	if err = r.genericHandler.BeforeDelete(ctx, params); err != nil {
		return err
	}

	if handled, err := r.genericHandler.Delete(ctx, params); handled {
		return err
	}

	if err = r.postgresResourceServiceBackend.DeleteRecords(resource, params.Ids); err != nil {
		return err
	}

	if err = r.genericHandler.AfterDelete(ctx, params); err != nil {
		return err
	}

	return nil
}
