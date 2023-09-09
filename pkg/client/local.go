package client

import (
	"context"
	"github.com/apibrew/apibrew/pkg/errors"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/service"
	log "github.com/sirupsen/logrus"
)

type localClient struct {
	container service.Container
}

func (l localClient) ListenRecords(ctx context.Context, namespace string, resource string, consumer func(records []*model.Record)) error {
	//TODO implement me
	panic("implement me")
}

func (l localClient) NewExtension(host string, remoteHost string) Extension {
	//TODO implement me
	panic("implement me")
}

func (l localClient) DeleteResource(ctx context.Context, id string, doMigration bool, forceMigration bool) error {
	//TODO implement me
	panic("implement me")
}

func (l localClient) ListRecords(ctx context.Context, params service.RecordListParams) ([]*model.Record, uint32, error) {
	return l.container.GetRecordService().List(ctx, params)
}

func (l localClient) GetResourceByName(ctx context.Context, namespace string, getType string) (*model.Resource, error) {
	return l.container.GetResourceService().GetResourceByName(ctx, namespace, getType)
}

func (l localClient) ReadRecordStream(ctx context.Context, params service.RecordListParams, recordsChan chan *model.Record) error {
	panic("Unsupported")
}

func (l localClient) AuthenticateWithToken(token string) {
	panic("Unsupported")
}

func (l localClient) AuthenticateWithUsernameAndPassword(username string, password string) error {
	panic("Unsupported")
}

func (l localClient) UpdateTokenFromContext(ctx context.Context) {
	panic("Unsupported")
}

func (l localClient) ApplyRecord(ctx context.Context, namespace string, resource string, record *model.Record) (*model.Record, error) {
	resp, err := l.container.GetRecordService().Apply(ctx, service.RecordUpdateParams{
		Namespace: namespace,
		Resource:  resource,
		Records:   []*model.Record{record},
	})

	if err != nil {
		return nil, err
	}

	return resp[0], err
}

func (l localClient) CreateRecord(ctx context.Context, namespace string, resource string, record *model.Record) (*model.Record, error) {
	resp, err := l.container.GetRecordService().Create(ctx, service.RecordCreateParams{
		Namespace: namespace,
		Resource:  resource,
		Records:   []*model.Record{record},
	})

	if err != nil {
		return nil, err
	}

	return resp[0], err
}

func (l localClient) UpdateRecord(ctx context.Context, namespace string, resource string, record *model.Record) (*model.Record, error) {
	resp, err := l.container.GetRecordService().Update(ctx, service.RecordUpdateParams{
		Namespace: namespace,
		Resource:  resource,
		Records:   []*model.Record{record},
	})

	if err != nil {
		return nil, err
	}

	return resp[0], err
}

func (l localClient) GetRecord(ctx context.Context, namespace string, resource string, id string) (*model.Record, error) {
	return l.container.GetRecordService().Get(ctx, service.RecordGetParams{
		Namespace: namespace,
		Resource:  resource,
		Id:        id,
	})
}

func (l localClient) FindRecords(ctx context.Context, params service.RecordListParams) ([]*model.Record, uint32, error) {
	return l.container.GetRecordService().List(ctx, params)
}

func (l localClient) ListResources(ctx context.Context) ([]*model.Resource, error) {
	return l.container.GetResourceService().List(ctx)
}

func (l localClient) ApplyResource(ctx context.Context, resource *model.Resource, doMigration, forceMigration bool) error {
	resp, err := l.container.GetResourceService().GetResourceByName(ctx, resource.Namespace, resource.Name)

	if !errors.ResourceNotFoundError.Is(err) && err != nil {
		return err
	}

	if errors.ResourceNotFoundError.Is(err) || resp == nil { // create
		_, err = l.container.GetResourceService().Create(ctx, resource, doMigration, forceMigration)

		if err != nil {
			return err
		}

		log.Info("Resource created: " + resource.Name)

		return nil
	} else {
		resource.Id = resp.Id
		err = l.container.GetResourceService().Update(ctx, resource, doMigration, forceMigration)

		if err != nil {
			return err
		}

		log.Info("Resource updated: " + resource.Name)

		return nil
	}
}

func NewLocalClient(container service.Container) DhClient {
	return &localClient{container: container}
}
