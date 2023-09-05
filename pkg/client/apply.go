package client

import (
	"context"
	"errors"
	errors2 "github.com/apibrew/apibrew/pkg/errors"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/stub"
	log "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/proto"
)

func (d *dhClient) Apply(ctx context.Context, msg proto.Message) error {
	switch msgTyped := msg.(type) {
	case *model.Resource:
		return d.ApplyResource(ctx, msgTyped, true, false)
	default:
		return errors.New(string("Unknown message type: " + msg.ProtoReflect().Descriptor().FullName()))
	}
}

func (d *dhClient) ApplyResource(ctx context.Context, resource *model.Resource, doMigration, forceMigration bool) error {
	resp, err := d.resourceClient.GetByName(ctx, &stub.GetResourceByNameRequest{
		Token:     d.token,
		Namespace: resource.Namespace,
		Name:      resource.Name,
	})

	if !errors2.ResourceNotFoundError.Is(err) && err != nil {
		return err
	}

	if errors2.ResourceNotFoundError.Is(err) || resp.Resource == nil { // create
		_, err = d.resourceClient.Create(ctx, &stub.CreateResourceRequest{
			Token:          d.token,
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
			Token:          d.token,
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
