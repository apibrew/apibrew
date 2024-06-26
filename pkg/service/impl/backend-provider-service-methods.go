package impl

import (
	"context"
	"fmt"
	"github.com/apibrew/apibrew/pkg/abs"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/service/annotations"
	"github.com/apibrew/apibrew/pkg/util"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (b *backendProviderService) GetStatus(ctx context.Context, dataSourceId string) (connectionAlreadyInitiated bool, testConnection bool, err error) {
	bck, err := b.getBackendByDataSourceId(ctx, dataSourceId)

	if err != nil {
		return false, false, err
	}

	return bck.GetStatus(ctx)
}

func (b backendProviderService) DestroyDataSource(ctx context.Context, dataSourceId string) error {
	return b.DestroyBackend(ctx, dataSourceId)
}

func (b backendProviderService) AddRecords(ctx context.Context, resource *model.Resource, records []*model.Record) ([]*model.Record, error) {
	endEvent, err := b.eventHandler.Handle(ctx, b.PrepareInternalEvent(ctx, &model.Event{
		Action:   model.Event_CREATE,
		Resource: resource,
		Records:  records,
	}))

	if err != nil {
		return nil, err
	}

	if endEvent == nil {
		return nil, nil
	}

	return endEvent.Records, nil
}

func (b backendProviderService) UpdateRecords(ctx context.Context, resource *model.Resource, records []*model.Record) ([]*model.Record, error) {
	endEvent, err := b.eventHandler.Handle(ctx, b.PrepareInternalEvent(ctx, &model.Event{
		Action:   model.Event_UPDATE,
		Resource: resource,
		Records:  records,
	}))

	if err != nil {
		return nil, err
	}

	if endEvent == nil {
		return nil, nil
	}

	return endEvent.Records, nil
}

func (b backendProviderService) GetRecord(ctx context.Context, resource *model.Resource, id string, resolveReferences []string) (*model.Record, error) {
	endEvent, err := b.eventHandler.Handle(ctx, b.PrepareInternalEvent(ctx, &model.Event{
		Action:   model.Event_GET,
		Resource: resource,
		Records:  []*model.Record{util.IdRecord(id)},
		RecordSearchParams: &model.Event_RecordSearchParams{
			ResolveReferences: resolveReferences,
		},
	}))

	if err != nil {
		return nil, err
	}

	if endEvent == nil {
		return nil, nil
	}

	if len(endEvent.Records) == 0 {
		return nil, nil
	}

	return endEvent.Records[0], nil
}

func (b backendProviderService) DeleteRecords(ctx context.Context, resource *model.Resource, list []*model.Record) error {
	_, err := b.eventHandler.Handle(ctx, b.PrepareInternalEvent(ctx, &model.Event{
		Action:   model.Event_DELETE,
		Resource: resource,
		Records:  list,
	}))

	return err
}

func (b backendProviderService) ListRecords(ctx context.Context, resource *model.Resource, params abs.ListRecordParams, resultChan chan<- *model.Record) ([]*model.Record, uint32, error) {
	endEvent, err := b.eventHandler.Handle(ctx, b.PrepareInternalEvent(ctx, &model.Event{
		Action:   model.Event_LIST,
		Resource: resource,
		RecordSearchParams: &model.Event_RecordSearchParams{
			Query:             params.Query,
			Limit:             params.Limit,
			Offset:            params.Offset,
			ResolveReferences: params.ResolveReferences,
			Aggregation:       params.Aggregation,
			Sorting:           params.Sorting,
		},
	}))

	if endEvent == nil || err != nil {
		return nil, 0, err
	}

	return endEvent.Records, uint32(endEvent.Total), nil
}

func (b backendProviderService) PrepareInternalEvent(ctx context.Context, event *model.Event) *model.Event {
	event.Id = fmt.Sprintf("internal-event-%s-%s-%s-%s", event.Resource.Namespace, event.Resource.Name, event.Action, util.RandomHex(6))
	event.Time = timestamppb.Now()
	event.Annotations = annotations.FromCtx(annotations.WithContext(ctx, event.Resource)).GetAnnotations()

	return event
}

func (b backendProviderService) ListEntities(ctx context.Context, dataSourceId string) ([]*model.DataSourceCatalog, error) {
	bck, err := b.getBackendByDataSourceId(ctx, dataSourceId)

	if err != nil {
		return nil, err
	}

	return bck.ListEntities(ctx)
}

func (b backendProviderService) PrepareResourceFromEntity(ctx context.Context, dataSourceName string, catalog, entity string) (*model.Resource, error) {
	bck, err := b.getBackendByDataSourceName(ctx, dataSourceName)

	if err != nil {
		return nil, err
	}

	return bck.PrepareResourceFromEntity(ctx, catalog, entity)
}

func (b backendProviderService) UpgradeResource(ctx context.Context, dataSourceName string, params abs.UpgradeResourceParams) error {
	bck, err := b.getBackendByDataSourceName(ctx, dataSourceName)

	if err != nil {
		return err
	}

	return bck.UpgradeResource(ctx, params)
}
