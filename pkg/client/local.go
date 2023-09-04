package client

import (
	"context"
	"github.com/apibrew/apibrew/pkg/errors"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/service"
	"github.com/apibrew/apibrew/pkg/stub"
	log "github.com/sirupsen/logrus"
)

type localClient struct {
	container service.Container
}

func (l localClient) ListResources(ctx context.Context) ([]*model.Resource, error) {
	return l.container.GetResourceService().List(ctx)
}

func (l localClient) ApplyRecord(ctx context.Context, resource *model.Resource, record *model.Record) error {
	_, err := l.container.GetRecordService().Apply(ctx, service.RecordUpdateParams{
		Namespace: resource.Namespace,
		Resource:  resource.Name,
		Records:   []*model.Record{record},
	})

	return err
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

func (l localClient) GetAuthenticationClient() stub.AuthenticationClient {
	panic("not supported")
}

func (l localClient) GetDataSourceClient() stub.DataSourceClient {
	panic("not supported")
}

func (l localClient) GetResourceClient() stub.ResourceClient {
	panic("not supported")
}

func (l localClient) GetRecordClient() stub.RecordClient {
	panic("not supported")
}

func (l localClient) GetGenericClient() stub.GenericClient {
	panic("not supported")
}

func (l localClient) GetToken() string {
	panic("not supported")
}

func (l localClient) AuthenticateWithToken(token string) {
	panic("not supported")
}

func (l localClient) AuthenticateWithUsernameAndPassword(username string, password string) error {
	panic("not supported")
}

func (l localClient) NewExtension(host string) Extension {
	panic("not supported")
}

func (l localClient) UpdateTokenFromContext(ctx context.Context) {
	panic("not supported")
}

func NewLocalClient(container service.Container) DhClient {
	return &localClient{container: container}
}
