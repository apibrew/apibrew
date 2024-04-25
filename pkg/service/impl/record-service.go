package impl

import (
	"context"
	"fmt"
	"github.com/apibrew/apibrew/pkg/abs"
	"github.com/apibrew/apibrew/pkg/errors"
	"github.com/apibrew/apibrew/pkg/helper"
	"github.com/apibrew/apibrew/pkg/logging"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/resource_model"
	"github.com/apibrew/apibrew/pkg/resource_model/extramappings"
	"github.com/apibrew/apibrew/pkg/resources"
	"github.com/apibrew/apibrew/pkg/service"
	"github.com/apibrew/apibrew/pkg/service/annotations"
	backend_event_handler "github.com/apibrew/apibrew/pkg/service/backend-event-handler"
	"github.com/apibrew/apibrew/pkg/service/validate"
	"github.com/apibrew/apibrew/pkg/util"
	"github.com/hashicorp/go-metrics"
	log "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/types/known/structpb"
)

type recordService struct {
	ServiceName            string
	resourceService        service.ResourceService
	backendServiceProvider service.BackendProviderService
	authorizationService   service.AuthorizationService
	backendEventHandler    backend_event_handler.BackendEventHandler
}

func (r *recordService) PrepareQuery(resource *model.Resource, queryMap map[string]string) (*model.BooleanExpression, error) {
	return util.PrepareQuery(resource, queryMap)
}

func NewRecordService(resourceService service.ResourceService, backendProviderService service.BackendProviderService, authorizationService service.AuthorizationService, backendEventHandler backend_event_handler.BackendEventHandler) service.RecordService {
	return &recordService{
		ServiceName:            "RecordService",
		resourceService:        resourceService,
		backendServiceProvider: backendProviderService,
		authorizationService:   authorizationService,
		backendEventHandler:    backendEventHandler,
	}
}

func (r *recordService) List(ctx context.Context, params service.RecordListParams) ([]*model.Record, uint32, error) {
	resource, _ := r.resourceService.GetResourceByName(util.WithSystemContext(ctx), params.Namespace, params.Resource)

	if resource == nil {
		return nil, 0, errors.ResourceNotFoundError.WithDetails(fmt.Sprintf("%s/%s", params.Namespace, params.Resource))
	}

	exp, err := r.authorizationService.CheckRecordAccessWithRecordSelector(ctx, service.CheckRecordAccessParams{
		Resource:  resource,
		Operation: resource_model.PermissionOperation_READ,
	})

	if err != nil {
		return nil, 0, err
	}

	// begin metrics
	defer metrics.IncrCounterWithLabels([]string{"RecordService"}, 1, []metrics.Label{
		{Name: "operation", Value: "List"},
		{Name: "resource", Value: params.Resource},
		{Name: "namespace", Value: params.Namespace},
	})
	// end metrics

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
		var err error

		params.Query, err = util.PrepareQueryFromFilters(resource, params.Filters)

		if err != nil {
			return nil, 0, err
		}
	}

	if exp != nil {
		expM := extramappings.BooleanExpressionToProto(*exp)
		if params.Query == nil {
			params.Query = expM
		} else {
			params.Query = &model.BooleanExpression{
				Expression: &model.BooleanExpression_And{
					And: &model.CompoundBooleanExpression{
						Expressions: []*model.BooleanExpression{expM, params.Query},
					},
				},
			}
		}
	}

	records, total, err := r.backendServiceProvider.ListRecords(ctx, resource, abs.ListRecordParams{
		Query:       params.Query,
		Limit:       params.Limit,
		Offset:      params.Offset,
		Aggregation: params.Aggregation,
		Sorting:     params.Sorting,
	}, params.ResultChan)

	// todo implement params.PackRecords

	if err != nil {
		return nil, 0, err
	}

	// resolving references
	if err = r.ResolveReferences(ctx, resource, records, params.ResolveReferences); err != nil {
		return nil, 0, err
	}

	return records, total, nil
}

func (r *recordService) Create(ctx context.Context, params service.RecordCreateParams) ([]*model.Record, error) {
	if params.Resource == "" {
		return nil, errors.RecordValidationError.WithMessage("Resource name is empty")
	}

	resource, _ := r.resourceService.GetResourceByName(util.WithSystemContext(ctx), params.Namespace, params.Resource)

	if resource == nil {
		return nil, errors.ResourceNotFoundError.WithDetails(fmt.Sprintf("%s/%s", params.Namespace, params.Resource))
	}

	return r.CreateWithResource(ctx, resource, params)
}

func (r *recordService) CreateWithResource(ctx context.Context, resource *model.Resource, params service.RecordCreateParams) ([]*model.Record, error) {
	var result []*model.Record

	var err error

	var txCtx = ctx

	// only check before operation
	if err := r.authorizationService.CheckRecordAccess(ctx, service.CheckRecordAccessParams{
		Resource:  resource,
		Records:   &params.Records,
		Operation: resource_model.PermissionOperation_CREATE,
	}); err != nil {
		return nil, err
	}

	// begin metrics
	defer metrics.IncrCounterWithLabels([]string{"RecordService"}, 1, []metrics.Label{
		{Name: "operation", Value: "Create"},
		{Name: "resource", Value: params.Resource},
		{Name: "namespace", Value: params.Namespace},
	})
	// end metrics

	if len(params.Records) == 0 {
		return nil, nil
	}

	if isResourceRelatedResource(resource) {
		return nil, errors.LogicalError.WithDetails("resource and related resources cannot be modified from records API")
	}

	for _, record := range params.Records {
		InitRecord(ctx, resource, record)
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

	if params.Records == nil {
		return nil, nil
	}

	records, err = r.backendServiceProvider.AddRecords(txCtx, resource, params.Records)

	if annotations.IsEnabled(resource, annotations.KeepHistory) && annotations.IsEnabledOnCtx(ctx, annotations.IgnoreIfExists) {
		return nil, errors.RecordValidationError.WithMessage("IgnoreIfExists must be disabled if resource has keepHistory enabled")
	}

	// create back reference
	if err := r.applyBackReferences(txCtx, resource, records); err != nil {
		return nil, err
	}

	if annotations.IsEnabled(resource, annotations.KeepHistory) {
		historyResource := util.HistoryResource(resource)

		_, err = r.backendServiceProvider.AddRecords(txCtx, historyResource, records)

		if err != nil {
			return nil, err
		}
	}

	if err != nil {
		return nil, err
	}

	result = append(result, records...)

	return result, nil
}

func isResourceRelatedResource(resource *model.Resource) bool {
	return resource.Namespace == resources.ResourceResource.Namespace && (resource.Name == resources.ResourceResource.Name)
}

func (r *recordService) Load(ctx context.Context, namespace string, resourceName string, properties map[string]*structpb.Value, loadParams service.RecordLoadParams) (*model.Record, error) {
	resource, _ := r.resourceService.GetResourceByName(util.WithSystemContext(ctx), namespace, resourceName)

	if resource == nil {
		return nil, errors.RecordValidationError.WithMessage("Resource not found with name: " + resourceName)
	}

	identifierProps, err := util.RecordIdentifierProperties(resource, properties)

	if err != nil {
		return nil, errors.RecordValidationError.WithMessage(err.Error())
	}

	qb := helper.NewQueryBuilder()

	var listParams = service.RecordListParams{}

	listParams.Query = qb.FromProperties(resource, identifierProps)
	listParams.Namespace = namespace
	listParams.Resource = resourceName
	listParams.Limit = 1
	listParams.Offset = 0
	listParams.ResolveReferences = loadParams.ResolveReferences
	listParams.UseHistory = loadParams.UseHistory

	searchRes, total, serr := r.List(ctx, listParams)

	if serr != nil {
		return nil, serr
	}

	if total == 1 {
		return searchRes[0], nil
	} else if total > 1 {
		return nil, errors.RecordNotFoundError.WithMessage("multiple records found with properties")
	} else {
		return nil, errors.RecordNotFoundError.WithMessage("Could not locate record with properties")
	}
}

func (r *recordService) Apply(ctx context.Context, params service.RecordUpdateParams) ([]*model.Record, error) {
	if params.Resource == "" {
		return nil, errors.RecordValidationError.WithMessage("Resource name is empty")
	}

	resource, _ := r.resourceService.GetResourceByName(util.WithSystemContext(ctx), params.Namespace, params.Resource)

	if resource == nil {
		return nil, errors.RecordValidationError.WithMessage("Resource not found with name: " + params.Resource)
	}

	var result []*model.Record

	for _, record := range params.Records {

		// locate existing record
		var existingRecord *model.Record

		if !resource.Virtual {
			identifierProps, err := util.RecordIdentifierProperties(resource, record.Properties)

			if err != nil {
				return nil, errors.RecordValidationError.WithMessage(err.Error())
			}

			qb := helper.NewQueryBuilder()

			searchRes, total, serr := r.List(ctx, service.RecordListParams{
				Namespace: resource.Namespace,
				Resource:  resource.Name,
				Limit:     1,
				Query:     qb.FromProperties(resource, identifierProps),
			})

			if err != nil {
				return nil, serr
			}

			if total > 0 {
				existingRecord = searchRes[0]
			}
		}

		if existingRecord == nil {
			records, err := r.CreateWithResource(ctx, resource, service.RecordCreateParams{
				Namespace: resource.Namespace,
				Resource:  resource.Name,
				Records:   []*model.Record{record},
			})

			if err != nil {
				return nil, err
			}

			result = append(result, records...)
		} else {
			if annotations.IsEnabled(annotations.FromCtx(ctx), annotations.IgnoreIfExists) {
				result = append(result, record)
				continue
			}

			if record.Properties != nil && existingRecord.Properties != nil {
				record.Properties["id"] = existingRecord.Properties["id"]
			}

			if util.IsSameRecord(existingRecord, record) {
				return params.Records, nil
			}

			records, err := r.UpdateWithResource(ctx, resource, service.RecordUpdateParams{
				Namespace: resource.Namespace,
				Resource:  resource.Name,
				Records:   []*model.Record{record},
			})

			if err != nil {
				return nil, err
			}

			result = append(result, records...)
		}
	}

	return result, nil
}

func (r *recordService) Update(ctx context.Context, params service.RecordUpdateParams) ([]*model.Record, error) {
	if params.Resource == "" {
		return nil, errors.RecordValidationError.WithMessage("Resource name is empty")
	}

	resource, _ := r.resourceService.GetResourceByName(util.WithSystemContext(ctx), params.Namespace, params.Resource)

	if resource == nil {
		return nil, errors.RecordValidationError.WithMessage("Resource not found with name: " + params.Resource)
	}

	return r.UpdateWithResource(ctx, resource, params)
}

func (r *recordService) UpdateWithResource(ctx context.Context, resource *model.Resource, params service.RecordUpdateParams) ([]*model.Record, error) {
	var result []*model.Record
	var err error

	if isResourceRelatedResource(resource) {
		return nil, errors.LogicalError.WithDetails("resource and related resources cannot be modified from records API")
	}

	if err := r.authorizationService.CheckRecordAccess(ctx, service.CheckRecordAccessParams{
		Resource:  resource,
		Records:   &params.Records,
		Operation: resource_model.PermissionOperation_UPDATE,
	}); err != nil {
		return nil, err
	}

	// begin metrics
	defer metrics.IncrCounterWithLabels([]string{"RecordService"}, 1, []metrics.Label{
		{Name: "operation", Value: "Update"},
		{Name: "resource", Value: params.Resource},
		{Name: "namespace", Value: params.Namespace},
	})
	// end metrics

	if len(params.Records) == 0 {
		return nil, nil
	}

	if resource.Immutable {
		return nil, errors.RecordValidationError.WithMessage("Immutable resource cannot be modified or deleted: " + params.Resource)
	}

	if annotations.IsEnabled(resource, annotations.KeepHistory) && !annotations.IsEnabledOnCtx(ctx, annotations.CheckVersion) {
		return nil, errors.RecordValidationError.WithMessage("checkVersion must be enabled if resource has keepHistory enabled")
	}

	for _, record := range params.Records {
		PrepareUpdateForRecord(ctx, resource, record)
	}

	var records []*model.Record

	var txCtx = ctx

	records, err = r.backendServiceProvider.UpdateRecords(txCtx, resource, params.Records)

	if err != nil {
		return nil, err
	}

	if err := r.applyBackReferences(txCtx, resource, records); err != nil {
		return nil, err
	}

	if annotations.IsEnabled(resource, annotations.KeepHistory) {
		_, err = r.backendServiceProvider.AddRecords(txCtx, util.HistoryResource(resource), records)

		if err != nil {
			return nil, err
		}
	}

	result = append(result, records...)

	return result, nil
}

func (r *recordService) applyBackReferences(ctx context.Context, resource *model.Resource, records []*model.Record) error {
	for typ, refProps := range r.resourceService.GetSchema().ResourcePropertiesByType[resource.Namespace+"/"+resource.Name] {
		if typ == model.ResourceProperty_REFERENCE {
			for _, refProp := range refProps {
				if refProp.Property.BackReference != nil {
					backRef := refProp.Property.BackReference

					var backRefNewRecords []*model.Record

					var ids []string

					for _, record := range records {
						getter, _ := util.RecordPropertyAccessorByPath(record.Properties, refProp.Path)

						if getter == nil {
							continue
						}
						gotVal := getter()

						if gotVal == nil {
							continue
						}

						ids = append(ids, util.GetRecordId(record))

						if gotVal.GetListValue() != nil {

							for _, item := range gotVal.GetListValue().Values {
								st := item.GetStructValue()
								st.Fields[backRef.Property] = structpb.NewStructValue(&structpb.Struct{Fields: map[string]*structpb.Value{
									"id": structpb.NewStringValue(util.GetRecordId(record)),
								}})

								backRefNewRecords = append(backRefNewRecords, &model.Record{
									Properties: st.Fields,
								})
							}
						}
					}

					if len(ids) == 0 {
						continue
					}

					existingRecords, _, err := r.List(ctx, service.RecordListParams{
						Namespace: refProp.Property.Reference.Namespace,
						Resource:  refProp.Property.Reference.Resource,
						Query: util.QueryInExpression(backRef.Property, structpb.NewListValue(&structpb.ListValue{
							Values: util.ArrayMap(ids, func(t string) *structpb.Value {
								return structpb.NewStringValue(t)
							}),
						})),
						ResolveReferences: []string{},
					})

					if err != nil {
						return err
					}

					var backRefRecordsRemovalIds []string

					for _, existingRecord := range existingRecords {
						backRefRecordsRemovalIds = append(backRefRecordsRemovalIds, util.GetRecordId(existingRecord))
					}

					if len(backRefRecordsRemovalIds) > 0 {
						if err := r.Delete(ctx, service.RecordDeleteParams{
							Namespace: refProp.Property.Reference.Namespace,
							Resource:  refProp.Property.Reference.Resource,
							Ids:       backRefRecordsRemovalIds,
						}); err != nil {
							return err
						}
					}

					if len(backRefNewRecords) > 0 {
						if _, err := r.Create(ctx, service.RecordCreateParams{
							Namespace: refProp.Property.Reference.Namespace,
							Resource:  refProp.Property.Reference.Resource,
							Records:   backRefNewRecords,
						}); err != nil {
							return err
						}
					}
				}
			}
		}
	}

	return nil
}

func (r *recordService) GetRecord(ctx context.Context, namespace, resourceName, id string, resolveReferences []string) (*model.Record, error) {
	resource, _ := r.resourceService.GetResourceByName(util.WithSystemContext(ctx), namespace, resourceName)

	if resource == nil {
		return nil, errors.ResourceNotFoundError.WithDetails(fmt.Sprintf("%s/%s", namespace, resourceName))
	}

	if isResourceRelatedResource(resource) {
		return nil, errors.LogicalError.WithDetails("resource and related resources cannot be modified from records API")
	}

	exp, err := r.authorizationService.CheckRecordAccessWithRecordSelector(ctx, service.CheckRecordAccessParams{
		Resource:  resource,
		Operation: resource_model.PermissionOperation_READ,
	})

	if err != nil {
		return nil, err
	}

	// begin metrics
	defer metrics.IncrCounterWithLabels([]string{"RecordService"}, 1, []metrics.Label{
		{Name: "operation", Value: "Get_"},
		{Name: "resource", Value: resourceName},
		{Name: "namespace", Value: namespace},
	})
	// end metrics

	var query = &model.BooleanExpression{Expression: &model.BooleanExpression_Equal{Equal: &model.PairExpression{
		Left:  &model.Expression{Expression: &model.Expression_Property{Property: "id"}},
		Right: &model.Expression{Expression: &model.Expression_Value{Value: structpb.NewStringValue(id)}},
	}}}

	if exp != nil {
		expM := extramappings.BooleanExpressionToProto(*exp)
		query = &model.BooleanExpression{
			Expression: &model.BooleanExpression_And{
				And: &model.CompoundBooleanExpression{
					Expressions: []*model.BooleanExpression{expM, query},
				},
			},
		}
	}

	records, total, err := r.backendServiceProvider.ListRecords(ctx, resource, abs.ListRecordParams{
		Query: query,
		Limit: 1,
	}, nil)

	if err != nil {
		return nil, err
	}

	if total == 0 {
		return nil, errors.RecordNotFoundError.WithMessage("Record not found with id: " + id)
	}

	var res = records[0]

	// resolving references
	if err := r.ResolveReferences(ctx, resource, []*model.Record{res}, resolveReferences); err != nil {
		return nil, err
	}

	return res, nil
}

func (r *recordService) FindBy(ctx context.Context, namespace, resourceName, propertyName string, value string) (*model.Record, error) {
	logger := log.WithFields(logging.CtxFields(ctx))

	logger.Debug("Begin record-service FindBy")
	defer logger.Debug("Finish record-service FindBy")

	resource, _ := r.resourceService.GetResourceByName(util.WithSystemContext(ctx), namespace, resourceName)

	if resource == nil {
		return nil, errors.ResourceNotFoundError.WithDetails(fmt.Sprintf("%s/%s", namespace, resourceName))
	}

	queryMap := make(map[string]string)

	queryMap[propertyName] = value

	logger.Debug("Call PrepareQuery: ", queryMap)
	query, err := util.PrepareQuery(resource, queryMap)
	logger.Debug("Result record-service: ", query)

	if err != nil {
		return nil, err
	}

	res, total, err := r.List(ctx, service.RecordListParams{
		Query:      query,
		Namespace:  namespace,
		Resource:   resourceName,
		Limit:      2,
		Offset:     0,
		UseHistory: false,
	})

	if err != nil {
		return nil, err
	}

	if total == 0 {
		return nil, errors.RecordNotFoundError.WithMessage(fmt.Sprintf("Record not found with %s: %v", propertyName, value))
	}

	if total > 1 {
		return nil, errors.LogicalError.WithDetails("We have more than 1 record")
	}

	return res[0], nil
}

func (r *recordService) Get(ctx context.Context, params service.RecordGetParams) (*model.Record, error) {
	return r.GetRecord(ctx, params.Namespace, params.Resource, params.Id, params.ResolveReferences)
}

func (r *recordService) Delete(ctx context.Context, params service.RecordDeleteParams) error {
	resource, _ := r.resourceService.GetResourceByName(util.WithSystemContext(ctx), params.Namespace, params.Resource)

	if resource == nil {
		return errors.ResourceNotFoundError.WithDetails(fmt.Sprintf("%s/%s", params.Namespace, params.Resource))
	}

	exp, err := r.authorizationService.CheckRecordAccessWithRecordSelector(ctx, service.CheckRecordAccessParams{
		Resource:  resource,
		Operation: resource_model.PermissionOperation_DELETE,
	})

	if err != nil {
		return err
	}

	// begin metrics
	defer metrics.IncrCounterWithLabels([]string{"RecordService"}, 1, []metrics.Label{
		{Name: "operation", Value: "Delete_"},
		{Name: "resource", Value: params.Resource},
		{Name: "namespace", Value: params.Namespace},
	})
	// end metrics

	if isResourceRelatedResource(resource) {
		return errors.LogicalError.WithDetails("resource and related resources cannot be modified from records API")
	}

	if resource.Immutable {
		return errors.RecordValidationError.WithMessage("Immutable resource cannot be modified or deleted: " + params.Resource)
	}

	////
	var query = &model.BooleanExpression{Expression: &model.BooleanExpression_In{In: &model.PairExpression{
		Left: &model.Expression{Expression: &model.Expression_Property{Property: "id"}},
		Right: &model.Expression{Expression: &model.Expression_Value{Value: structpb.NewListValue(&structpb.ListValue{Values: util.ArrayMap(params.Ids, func(t string) *structpb.Value {
			return structpb.NewStringValue(t)
		})})}},
	}}}

	if exp != nil {
		expM := extramappings.BooleanExpressionToProto(*exp)
		query = &model.BooleanExpression{
			Expression: &model.BooleanExpression_And{
				And: &model.CompoundBooleanExpression{
					Expressions: []*model.BooleanExpression{expM, query},
				},
			},
		}
	}

	records, _, err := r.backendServiceProvider.ListRecords(ctx, resource, abs.ListRecordParams{
		Query: query,
		Limit: uint32(len(params.Ids)),
	}, nil)

	if err != nil {
		return err
	}

	for _, id := range params.Ids {
		var found = false

		for _, record := range records {
			var foundId = util.GetRecordId(record)
			if id == foundId {
				found = true
			}
		}

		if !found {
			return errors.RecordNotFoundError.WithMessage("Record not found with id: " + id)
		}
	}

	if err = r.backendServiceProvider.DeleteRecords(ctx, resource, records); err != nil {
		return err
	}

	return nil
}

func (r *recordService) Init(config *model.AppConfig) {
	r.initHandlers()

	r.initRecords(config)
}

func (r *recordService) initHandlers() {
	r.backendEventHandler.RegisterHandler(backend_event_handler.Handler{
		Id:   "record-validation-handler",
		Name: "record-validation-handler",
		Fn:   r.validateRecordHandler,
		Selector: &model.EventSelector{
			Actions: []model.Event_Action{
				model.Event_CREATE,
				model.Event_UPDATE,
			},
		},
		Order:    50,
		Responds: true,
		Sync:     true,
		Internal: true,
	})

	r.backendEventHandler.RegisterHandler(backend_event_handler.Handler{
		Id:   "record-reference-check-handler",
		Name: "record-reference-check-handler",
		Fn:   r.referenceCheckHandler,
		Selector: &model.EventSelector{
			Actions: []model.Event_Action{
				model.Event_CREATE,
				model.Event_UPDATE,
			},
		},
		Order:    51,
		Responds: true,
		Sync:     true,
		Internal: true,
	})
}

func (r *recordService) initRecords(config *model.AppConfig) {
	ctx := util.WithSystemContext(context.TODO())
	for _, initRecord := range config.InitRecords {
		subCtx := ctx

		if !initRecord.Override {
			subCtx = annotations.SetWithContext(subCtx, annotations.IgnoreIfExists, annotations.Enabled)
		}

		_, err := r.Apply(subCtx, service.RecordUpdateParams{
			Namespace: initRecord.Namespace,
			Resource:  initRecord.Resource,
			Records:   []*model.Record{initRecord.Record},
		})

		if err != nil {
			log.Fatal(err)
		}
	}
}

func (r *recordService) ResolveReferences(ctx context.Context, resource *model.Resource, records []*model.Record, referencesToResolve []string) error {
	log.Debug("Begin record-service ResolveReferences: " + resource.Namespace + "/" + resource.Name)

	defer func() {
		log.Debug("End record-service ResolveReferences: " + resource.Namespace + "/" + resource.Name)
	}()
	if len(records) == 0 {
		return nil
	}

	if len(referencesToResolve) == 0 {
		return nil
	}

	// resolving references
	references := r.resourceService.LocateReferences(resource, referencesToResolve)

	var rr = &recordResolver{
		recordService:   r,
		resourceService: r.resourceService,
		resource:        resource,
		records:         records,
		paths:           references,
	}

	return rr.resolveReferences(ctx)
}

func (r *recordService) checkReferences(ctx context.Context, resource *model.Resource, records []*model.Record) error {
	log.Debug("Begin record-service ResolveReferences: " + resource.Namespace + "/" + resource.Name)

	defer func() {
		log.Debug("End record-service ResolveReferences: " + resource.Namespace + "/" + resource.Name)
	}()
	if len(records) == 0 {
		return nil
	}

	// resolving references
	references := r.resourceService.LocateLocalReferences(resource)

	var rr = &recordResolver{
		recordService:   r,
		resourceService: r.resourceService,
		resource:        resource,
		records:         records,
		paths:           references,
	}

	return rr.checkReferences(ctx)
}

func (r *recordService) validateRecordHandler(ctx context.Context, event *model.Event) (*model.Event, error) {
	if err := validate.Records(event.Resource, event.Records, event.Action == model.Event_UPDATE); err != nil {
		return nil, err
	}

	return event, nil
}

func (r *recordService) referenceCheckHandler(ctx context.Context, event *model.Event) (*model.Event, error) {
	if event.Resource.CheckReferences {
		if err := r.checkReferences(ctx, event.Resource, event.Records); err != nil {
			return nil, err
		}
	}

	return event, nil
}
