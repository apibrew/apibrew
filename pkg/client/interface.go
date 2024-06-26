package client

import (
	"context"
	"github.com/apibrew/apibrew/pkg/api"
	"github.com/apibrew/apibrew/pkg/errors"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/service"
	"github.com/apibrew/apibrew/pkg/util"
	"google.golang.org/protobuf/types/known/structpb"
)

type clientResourceService struct {
	client Client
}

func (c clientResourceService) GetResourceByName(ctx context.Context, namespace, resource string) (*model.Resource, error) {
	res, err := c.client.GetResourceByName(ctx, namespace, resource)

	if err != nil {
		return nil, errors.FromGrpcError(err)
	}

	return res, nil
}

func (c clientResourceService) Create(ctx context.Context, resource *model.Resource, doMigration bool, forceMigration bool) (*model.Resource, error) {
	if err := c.client.CreateResource(ctx, resource, doMigration, forceMigration); err != nil {
		return nil, errors.FromGrpcError(err)
	}

	return c.GetResourceByName(ctx, resource.Namespace, resource.Name)
}

func (c clientResourceService) Update(ctx context.Context, resource *model.Resource, doMigration bool, forceMigration bool) error {
	if err := c.client.UpdateResource(ctx, resource, doMigration, forceMigration); err != nil {
		return errors.FromGrpcError(err)
	}

	return nil
}

func (c clientResourceService) Delete(ctx context.Context, ids []string, doMigration bool, forceMigration bool) error {
	for _, id := range ids {
		if err := c.client.DeleteResource(ctx, id, doMigration, forceMigration); err != nil {
			return errors.FromGrpcError(err)
		}
	}

	return nil
}

func (c clientResourceService) List(ctx context.Context) ([]*model.Resource, error) {
	res, err := c.client.ListResources(ctx)

	if err != nil {
		return nil, errors.FromGrpcError(err)
	}

	return res, nil
}

type clientRecordService struct {
	client Client
}

func (c clientRecordService) Create(ctx context.Context, params service.RecordCreateParams) ([]*model.Record, error) {
	var result []*model.Record

	for _, record := range params.Records {
		created, err := c.client.CreateRecord(ctx, params.Namespace, params.Resource, record)

		if err != nil {
			return nil, errors.FromGrpcError(err)
		}

		result = append(result, created)
	}

	return result, nil
}

func (c clientRecordService) Update(ctx context.Context, params service.RecordUpdateParams) ([]*model.Record, error) {
	var result []*model.Record

	for _, record := range params.Records {
		updated, err := c.client.UpdateRecord(ctx, params.Namespace, params.Resource, record)

		if err != nil {
			return nil, errors.FromGrpcError(err)
		}

		result = append(result, updated)
	}

	return result, nil
}

func (c clientRecordService) Apply(ctx context.Context, params service.RecordUpdateParams) ([]*model.Record, error) {
	var result []*model.Record

	for _, record := range params.Records {
		applied, err := c.client.ApplyRecord(ctx, params.Namespace, params.Resource, record)

		if err != nil {
			return nil, errors.FromGrpcError(err)
		}

		result = append(result, applied)
	}

	return result, nil
}

func (c clientRecordService) Delete(ctx context.Context, params service.RecordDeleteParams) error {
	for _, id := range params.Ids {
		if err := c.client.DeleteRecord(ctx, params.Namespace, params.Resource, util.IdRecord(id)); err != nil {
			return errors.FromGrpcError(err)
		}
	}

	return nil
}

func (c clientRecordService) Load(ctx context.Context, namespace string, name string, properties map[string]*structpb.Value, loadParams service.RecordLoadParams) (*model.Record, error) {
	res, err := c.client.LoadRecord(ctx, namespace, name, properties, loadParams)

	if err != nil {
		return nil, errors.FromGrpcError(err)
	}

	return res, nil
}

func (c clientRecordService) List(ctx context.Context, params service.RecordListParams) ([]*model.Record, uint32, error) {
	res, total, err := c.client.ListRecords(ctx, params)

	if err != nil {
		return nil, 0, errors.FromGrpcError(err)
	}

	return res, total, nil
}

func NewInterface(client Client) api.Interface {
	return api.NewInterface2(&clientResourceService{client: client}, &clientRecordService{client: client})
}
