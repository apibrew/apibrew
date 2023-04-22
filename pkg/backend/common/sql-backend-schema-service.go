package common

import (
	"context"
	log "github.com/sirupsen/logrus"
	"github.com/tislib/apibrew/pkg/abs"
	"github.com/tislib/apibrew/pkg/backend/helper"
	"github.com/tislib/apibrew/pkg/errors"
	"github.com/tislib/apibrew/pkg/logging"
	"github.com/tislib/apibrew/pkg/model"
	"github.com/tislib/apibrew/pkg/util"
)

func (p *sqlBackend) ListEntities(ctx context.Context) (result []*model.DataSourceCatalog, err errors.ServiceError) {
	err = p.withBackend(ctx, true, func(tx helper.QueryRunner) errors.ServiceError {
		result, err = p.resourceListEntities(ctx, tx)

		return err
	})

	return
}

func (p *sqlBackend) PrepareResourceFromEntity(ctx context.Context, catalog string, entity string) (resource *model.Resource, err errors.ServiceError) {
	logger := log.WithFields(logging.CtxFields(ctx))

	err = p.withBackend(ctx, false, func(tx helper.QueryRunner) errors.ServiceError {
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

func (p *sqlBackend) UpgradeResource(ctx context.Context, params abs.UpgradeResourceParams) errors.ServiceError {
	return p.withBackend(ctx, false, func(tx helper.QueryRunner) errors.ServiceError {
		return p.resourceMigrateTable(ctx, tx, params)
	})
}
