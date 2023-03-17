package common

import (
	"context"
	log "github.com/sirupsen/logrus"
	"github.com/tislib/data-handler/pkg/abs"
	"github.com/tislib/data-handler/pkg/backend/helper"
	"github.com/tislib/data-handler/pkg/errors"
	"github.com/tislib/data-handler/pkg/logging"
	"github.com/tislib/data-handler/pkg/model"
	"github.com/tislib/data-handler/pkg/service/annotations"
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

	return resource, nil
}

func (p *sqlBackend) UpgradeResource(ctx context.Context, params abs.UpgradeResourceParams) errors.ServiceError {
	return p.withBackend(ctx, false, func(tx helper.QueryRunner) errors.ServiceError {
		if err := p.resourceMigrateTable(ctx, tx, params, false); err != nil {
			return err
		}

		if annotations.IsEnabled(params.MigrationPlan.CurrentResource, annotations.KeepHistory) {
			if err := p.resourceMigrateTable(ctx, tx, params, true); err != nil {
				return err
			}
		}

		return nil
	})
}
