package postgres

import (
	"context"
	log "github.com/sirupsen/logrus"
	"github.com/tislib/data-handler/pkg/abs"
	"github.com/tislib/data-handler/pkg/errors"
	"github.com/tislib/data-handler/pkg/logging"
	"github.com/tislib/data-handler/pkg/model"
	annotations "github.com/tislib/data-handler/pkg/service/annotations"
)

func (p *postgresResourceServiceBackend) ListEntities(ctx context.Context) (result []*model.DataSourceCatalog, err errors.ServiceError) {
	err = p.withBackend(ctx, true, func(tx QueryRunner) errors.ServiceError {
		result, err = resourceListEntities(ctx, tx)

		return err
	})

	return
}

func (p *postgresResourceServiceBackend) PrepareResourceFromEntity(ctx context.Context, catalog string, entity string) (resource *model.Resource, err errors.ServiceError) {
	logger := log.WithFields(logging.CtxFields(ctx))

	err = p.withBackend(ctx, false, func(tx QueryRunner) errors.ServiceError {
		if resource, err = resourcePrepareResourceFromEntity(ctx, tx, catalog, entity); err != nil {
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

func (p *postgresResourceServiceBackend) UpgradeResource(ctx context.Context, params abs.UpgradeResourceParams) errors.ServiceError {
	return p.withBackend(ctx, false, func(tx QueryRunner) errors.ServiceError {
		if err := resourceCreateTable(ctx, tx, params.Resource); err != nil {
			return err
		}

		if err := resourceMigrateTable(ctx, tx, params, false); err != nil {
			return err
		}

		if annotations.IsEnabled(params.Resource, annotations.KeepHistory) {
			if err := resourceCreateHistoryTable(ctx, tx, params.Resource); err != nil {
				return err
			}

			if err := resourceMigrateTable(ctx, tx, params, true); err != nil {
				return err
			}
		}

		return nil
	})
}

func (p *postgresResourceServiceBackend) DowngradeResource(ctx context.Context, resource *model.Resource, forceMigration bool) errors.ServiceError {
	return p.withBackend(ctx, false, func(tx QueryRunner) errors.ServiceError {
		err := resourceDropTable(ctx, tx, getFullTableName(resource.SourceConfig, false), forceMigration)

		if err != nil {
			return err
		}

		if annotations.IsEnabled(resource, annotations.KeepHistory) {
			err = resourceDropTable(ctx, tx, getFullTableName(resource.SourceConfig, true), forceMigration)

			if err != nil {
				return err
			}
		}

		return nil
	})
}
