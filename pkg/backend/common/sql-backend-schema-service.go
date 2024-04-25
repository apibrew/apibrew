package common

import (
	"context"
	"github.com/apibrew/apibrew/pkg/abs"
	"github.com/apibrew/apibrew/pkg/backend/helper"
	"github.com/apibrew/apibrew/pkg/logging"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/util"
	log "github.com/sirupsen/logrus"
)

func (p *sqlBackend) ListEntities(ctx context.Context) (result []*model.DataSourceCatalog, err error) {
	err = p.withBackend(ctx, true, func(tx helper.QueryRunner) error {
		result, err = p.resourceListEntities(ctx, tx)

		return err
	})

	return
}

func (p *sqlBackend) PrepareResourceFromEntity(ctx context.Context, catalog string, entity string) (resource *model.Resource, err error) {
	logger := log.WithFields(logging.CtxFields(ctx))

	err = p.withBackend(ctx, false, func(tx helper.QueryRunner) error {
		if resource, err = p.resourcePrepareResourceFromEntity(ctx, tx, catalog, entity); err != nil {
			logger.Errorf("[PrepareResourceFromEntity] Unable to load resource details for %s Err: %s", entity, err)
			return err
		}

		return nil
	})

	if err != nil {
		logger.Errorf("Unable to load resource for %s Err: %s", entity, err)
		return nil, err
	}

	util.RemarkResource(resource)

	return resource, nil
}

func (p *sqlBackend) UpgradeResource(ctx context.Context, params abs.UpgradeResourceParams) error {
	return p.withBackend(ctx, false, func(tx helper.QueryRunner) error {
		return p.resourceMigrateTable(ctx, tx, params)
	})
}
