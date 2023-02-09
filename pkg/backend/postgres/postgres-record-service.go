package postgres

import (
	"context"
	"database/sql"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/tislib/data-handler/pkg/abs"
	"github.com/tislib/data-handler/pkg/errors"
	"github.com/tislib/data-handler/pkg/logging"
	"github.com/tislib/data-handler/pkg/model"
	annotations "github.com/tislib/data-handler/pkg/service/annotations"
)

func (p *postgresResourceServiceBackend) ListRecords(ctx context.Context, params abs.ListRecordParams) (result []*model.Record, total uint32, err errors.ServiceError) {
	logger := log.WithFields(logging.CtxFields(ctx))

	logger.Tracef("Begin listing: %s/%s", params.Resource.Namespace, params.Resource.Name)
	err = p.withBackend(ctx, true, func(tx *sql.Tx) errors.ServiceError {
		result, total, err = recordList(ctx, tx, params)

		return err
	})
	logger.Tracef("Finish listing: %s/%s", params.Resource.Namespace, params.Resource.Name)

	return
}

func (p *postgresResourceServiceBackend) AddRecords(ctx context.Context, params abs.BulkRecordsParams) ([]*model.Record, bool, errors.ServiceError) {
	logger := log.WithFields(logging.CtxFields(ctx))

	var inserted bool
	var err errors.ServiceError

	logger.Tracef("Begin creating: %s/%s", params.Resource.Namespace, params.Resource.Name)

	err = p.withBackend(ctx, false, func(tx *sql.Tx) errors.ServiceError {
		inserted, err = recordInsert(ctx, tx, params.Resource, params.Records, params.IgnoreIfExists, params.Schema, false)

		if err != nil {
			return err
		}

		if inserted && annotations.IsEnabled(params.Resource, annotations.KeepHistory) {
			_, err = recordInsert(ctx, tx, params.Resource, params.Records, false, params.Schema, true)

			if err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return nil, inserted, err
	}

	logger.Tracef("Records created: %s/%s", params.Resource.Namespace, params.Resource.Name)

	return params.Records, inserted, nil
}

func (p *postgresResourceServiceBackend) UpdateRecords(ctx context.Context, params abs.BulkRecordsParams) ([]*model.Record, errors.ServiceError) {
	err := p.withBackend(ctx, false, func(tx *sql.Tx) errors.ServiceError {
		for _, record := range params.Records {
			err := recordUpdate(ctx, tx, params.Resource, record, params.CheckVersion)

			if err != nil {
				return err
			}
		}

		if annotations.IsEnabled(params.Resource, annotations.KeepHistory) {
			_, err := recordInsert(ctx, tx, params.Resource, params.Records, false, params.Schema, false)

			if err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return params.Records, nil
}

func (p *postgresResourceServiceBackend) GetRecord(ctx context.Context, resource *model.Resource, schema *abs.Schema, id string) (*model.Record, errors.ServiceError) {
	var record *model.Record = nil
	err := p.withBackend(ctx, true, func(tx *sql.Tx) errors.ServiceError {
		var err errors.ServiceError
		record, err = readRecord(ctx, tx, resource, schema, id)

		if err == sql.ErrNoRows {
			return errors.RecordNotFoundError.WithDetails(fmt.Sprintf("namespace %s; resource %s; id %v", resource.Namespace, resource.Name, id))
		}

		if err != nil {
			return err
		}

		return err
	})

	return record, err
}

func (p *postgresResourceServiceBackend) DeleteRecords(ctx context.Context, resource *model.Resource, ids []string) errors.ServiceError {
	logger := log.WithFields(logging.CtxFields(ctx))

	logger.Tracef("Begin deleting records: %v / %v / %v", resource.Namespace, resource.Name, ids)
	err := p.withBackend(ctx, false, func(tx *sql.Tx) errors.ServiceError {
		return deleteRecords(ctx, tx, resource, ids)
	})
	if err != nil {
		logger.Print(err)
	} else {
		logger.Tracef("records deleted: %v / %v / %v", resource.Namespace, resource.Name, ids)
	}

	return err
}
