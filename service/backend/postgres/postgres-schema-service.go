package postgres

import (
	"context"
	"data-handler/model"
	"data-handler/service/annotations"
	"data-handler/service/errors"
	"database/sql"
	log "github.com/sirupsen/logrus"
)

func (p *postgresResourceServiceBackend) ListEntities(ctx context.Context) (result []string, err errors.ServiceError) {
	err = p.withBackend(ctx, true, func(tx *sql.Tx) errors.ServiceError {
		result, err = resourceListEntities(ctx, tx)

		return err
	})

	return
}

func (p *postgresResourceServiceBackend) PrepareResourceFromEntity(ctx context.Context, catalog string, entity string) (resource *model.Resource, err errors.ServiceError) {
	err = p.withBackend(ctx, false, func(tx *sql.Tx) errors.ServiceError {
		if resource, err = resourcePrepareResourceFromEntity(ctx, tx, catalog, entity); err != nil {
			log.Errorf("[PrepareResourceFromEntity] Unable to load resource details for %s Err: %s", entity, err)
			return err
		}

		return nil
	})

	if err != nil {
		log.Errorf("Unable to load resource for %s Err: %s", entity, err)
		return nil, err
	}

	return resource, nil
}

func (p *postgresResourceServiceBackend) UpgradeResource(ctx context.Context, currentResource *model.Resource, resource *model.Resource, forceMigration bool) errors.ServiceError {
	return p.withBackend(ctx, false, func(tx *sql.Tx) errors.ServiceError {
		if err := resourceCreateTable(tx, resource); err != nil {
			return err
		}

		if err := resourceMigrateTable(ctx, tx, resource, forceMigration, false); err != nil {
			return err
		}

		if annotations.IsEnabled(resource, annotations.KeepHistory) {
			if err := resourceCreateHistoryTable(tx, resource); err != nil {
				return err
			}

			if err := resourceMigrateTable(ctx, tx, resource, forceMigration, true); err != nil {
				return err
			}
		}

		return nil
	})
}

func (p *postgresResourceServiceBackend) DowngradeResource(ctx context.Context, resource *model.Resource, forceMigration bool) errors.ServiceError {
	return p.withBackend(ctx, false, func(tx *sql.Tx) errors.ServiceError {
		err := resourceDropTable(tx, getTableName(resource.SourceConfig, false))

		if err != nil {
			return err
		}

		if annotations.IsEnabled(resource, annotations.KeepHistory) {
			err = resourceDropTable(tx, getTableName(resource.SourceConfig, true))

			if err != nil {
				return err
			}
		}

		return nil
	})
}
