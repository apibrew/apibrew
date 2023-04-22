package client

import (
	"context"
	"errors"
	log "github.com/sirupsen/logrus"
	errors2 "github.com/tislib/apibrew/pkg/errors"
	"github.com/tislib/apibrew/pkg/model"
	"github.com/tislib/apibrew/pkg/stub"
	"google.golang.org/protobuf/proto"
)

func (d *dhClient) Apply(ctx context.Context, msg proto.Message) error {
	switch msgTyped := msg.(type) {
	case *model.Resource:
		return d.ApplyResource(ctx, msgTyped, true, false)
	case *model.DataSource:
		return d.ApplyDataSource(ctx, msgTyped)
	case *model.User:
		return d.ApplyUser(ctx, msgTyped)
	case *model.Extension:
		return d.ApplyExtension(ctx, msgTyped)
	case *model.Namespace:
		return d.ApplyNamespace(ctx, msgTyped)
	default:
		return errors.New(string("Unknown message type: " + msg.ProtoReflect().Descriptor().FullName()))
	}
}

func (d *dhClient) ApplyResource(ctx context.Context, resource *model.Resource, doMigration, forceMigration bool) error {
	resp, err := d.resourceClient.GetByName(ctx, &stub.GetResourceByNameRequest{
		Token:     d.GetToken(),
		Namespace: resource.Namespace,
		Name:      resource.Name,
	})

	if !errors2.ResourceNotFoundError.Is(err) && err != nil {
		return err
	}

	if errors2.ResourceNotFoundError.Is(err) || resp.Resource == nil { // create
		_, err = d.resourceClient.Create(ctx, &stub.CreateResourceRequest{
			Token:          d.GetToken(),
			Resources:      []*model.Resource{resource},
			DoMigration:    doMigration,
			ForceMigration: forceMigration,
		})

		if err != nil {
			return err
		}

		log.Info("Resource created: " + resource.Name)

		return nil
	} else {
		resource.Id = resp.Resource.Id
		_, err = d.resourceClient.Update(ctx, &stub.UpdateResourceRequest{
			Token:          d.GetToken(),
			Resources:      []*model.Resource{resource},
			DoMigration:    doMigration,
			ForceMigration: forceMigration,
		})

		if err != nil {
			return err
		}

		log.Info("Resource updated: " + resource.Name)

		return nil
	}
}

func (d *dhClient) ApplyDataSource(ctx context.Context, dataSource *model.DataSource) error {
	resp, err := d.dataSourceClient.List(ctx, &stub.ListDataSourceRequest{
		Token: d.GetToken(),
	})

	if err != nil {
		return err
	}

	var existingDataSource *model.DataSource

	for _, item := range resp.Content {
		if item.Name == dataSource.Name {
			existingDataSource = item
		}
	}

	if existingDataSource == nil { // create
		_, err = d.dataSourceClient.Create(ctx, &stub.CreateDataSourceRequest{
			Token:       d.GetToken(),
			DataSources: []*model.DataSource{dataSource},
		})

		if err != nil {
			return err
		}

		log.Info("DataSource created: " + dataSource.Name)

		return nil
	} else {
		dataSource.Id = existingDataSource.Id
		_, err = d.dataSourceClient.Update(ctx, &stub.UpdateDataSourceRequest{
			Token:       d.GetToken(),
			DataSources: []*model.DataSource{dataSource},
		})

		if err != nil {
			return err
		}

		log.Info("DataSource updated: " + dataSource.Name)

		return nil
	}
}

func (d *dhClient) ApplyUser(ctx context.Context, user *model.User) error {
	resp, err := d.userClient.List(ctx, &stub.ListUserRequest{
		Token: d.GetToken(),
	})

	if err != nil {
		return err
	}

	var existingUser *model.User

	for _, item := range resp.Content {
		if item.Username == user.Username {
			existingUser = item
		}
	}

	if existingUser == nil { // create
		_, err = d.userClient.Create(ctx, &stub.CreateUserRequest{
			Token: d.GetToken(),
			Users: []*model.User{user},
		})

		if err != nil {
			return err
		}

		log.Info("User created: " + user.Username)

		return nil
	} else {
		user.Id = existingUser.Id
		_, err = d.userClient.Update(ctx, &stub.UpdateUserRequest{
			Token: d.GetToken(),
			Users: []*model.User{user},
		})

		if err != nil {
			return err
		}

		log.Info("User updated: " + user.Username)

		return nil
	}
}

func (d *dhClient) ApplyExtension(ctx context.Context, extension *model.Extension) error {
	resp, err := d.extensionClient.List(ctx, &stub.ListExtensionRequest{
		Token: d.GetToken(),
	})

	if err != nil {
		return err
	}

	var existingExtension *model.Extension

	for _, item := range resp.Content {
		if item.Name == extension.Name {
			existingExtension = item
		}
	}

	if existingExtension == nil { // create
		_, err = d.extensionClient.Create(ctx, &stub.CreateExtensionRequest{
			Token:      d.GetToken(),
			Extensions: []*model.Extension{extension},
		})

		if err != nil {
			return err
		}

		log.Info("Extension created: " + extension.Name)

		return nil
	} else {
		extension.Id = existingExtension.Id
		_, err = d.extensionClient.Update(ctx, &stub.UpdateExtensionRequest{
			Token:      d.GetToken(),
			Extensions: []*model.Extension{extension},
		})

		if err != nil {
			return err
		}

		log.Info("Extension updated: " + extension.Name)

		return nil
	}
}

func (d *dhClient) ApplyNamespace(ctx context.Context, namespace *model.Namespace) error {
	resp, err := d.namespaceClient.List(ctx, &stub.ListNamespaceRequest{
		Token: d.GetToken(),
	})

	if err != nil {
		return err
	}

	var existingNamespace *model.Namespace

	for _, item := range resp.Content {
		if item.Name == namespace.Name {
			existingNamespace = item
		}
	}

	if existingNamespace == nil { // create
		_, err = d.namespaceClient.Create(ctx, &stub.CreateNamespaceRequest{
			Token:      d.GetToken(),
			Namespaces: []*model.Namespace{namespace},
		})

		if err != nil {
			return err
		}

		log.Info("Namespace created: " + namespace.Name)

		return nil
	} else {
		namespace.Id = existingNamespace.Id
		_, err = d.namespaceClient.Update(ctx, &stub.UpdateNamespaceRequest{
			Token:      d.GetToken(),
			Namespaces: []*model.Namespace{namespace},
		})

		if err != nil {
			return err
		}

		log.Info("Namespace updated: " + namespace.Name)

		return nil
	}
}

func (d *dhClient) ApplyRecord(ctx context.Context, resource *model.Resource, record *model.Record) error {
	_, err := d.recordClient.Apply(ctx, &stub.ApplyRecordRequest{
		Token:     d.GetToken(),
		Namespace: resource.Namespace,
		Resource:  resource.Name,
		Record:    record,
	})

	return err
}
