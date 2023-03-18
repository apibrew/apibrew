package common

import (
	"context"
	"database/sql"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/tislib/data-handler/pkg/abs"
	"github.com/tislib/data-handler/pkg/backend/helper"
	"github.com/tislib/data-handler/pkg/errors"
	"github.com/tislib/data-handler/pkg/logging"
	"github.com/tislib/data-handler/pkg/model"
)

func (p *sqlBackend) ListRecords(ctx context.Context, params abs.ListRecordParams) (result []*model.Record, total uint32, err errors.ServiceError) {
	logger := log.WithFields(logging.CtxFields(ctx))

	logger.Tracef("Begin listing: %s/%s", params.Resource.Namespace, params.Resource.Name)
	err = p.withBackend(ctx, true, func(tx helper.QueryRunner) errors.ServiceError {
		result, total, err = p.recordList(ctx, tx, params)

		return err
	})
	logger.Tracef("Finish listing: %s/%s", params.Resource.Namespace, params.Resource.Name)

	return
}

func (p *sqlBackend) AddRecords(ctx context.Context, params abs.BulkRecordsParams) ([]*model.Record, []bool, errors.ServiceError) {
	logger := log.WithFields(logging.CtxFields(ctx))

	var inserted []bool
	var err errors.ServiceError

	logger.Tracef("Begin creating: %s/%s", params.Resource.Namespace, params.Resource.Name)

	err = p.withBackend(ctx, false, func(tx helper.QueryRunner) errors.ServiceError {
		maxChunkSize := 1000
		chunkCount := len(params.Records) / maxChunkSize
		if chunkCount == 0 {
			chunkCount = 1
		}

		for i := 0; i < chunkCount; i++ {
			bi := i * maxChunkSize
			ei := (i + 1) * maxChunkSize

			if ei > len(params.Records) {
				ei = len(params.Records)
			}

			records := params.Records[bi:ei]

			id, err := p.recordInsert(ctx, tx, params.Resource, records, params.IgnoreIfExists, params.Schema)

			inserted = append(inserted, id)

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

func (p *sqlBackend) UpdateRecords(ctx context.Context, params abs.BulkRecordsParams) ([]*model.Record, errors.ServiceError) {
	err := p.withBackend(ctx, false, func(tx helper.QueryRunner) errors.ServiceError {
		for _, record := range params.Records {
			err := p.recordUpdate(ctx, tx, params.Resource, record, params.CheckVersion, params.Schema)

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

func (p *sqlBackend) GetRecord(ctx context.Context, resource *model.Resource, schema *abs.Schema, id string) (*model.Record, errors.ServiceError) {
	var record *model.Record = nil
	err := p.withBackend(ctx, true, func(tx helper.QueryRunner) errors.ServiceError {
		var err errors.ServiceError
		record, err = p.readRecord(ctx, tx, resource, schema, id)

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

func (p *sqlBackend) DeleteRecords(ctx context.Context, resource *model.Resource, ids []string) errors.ServiceError {
	logger := log.WithFields(logging.CtxFields(ctx))

	logger.Tracef("Begin deleting records: %v / %v / %v", resource.Namespace, resource.Name, ids)
	err := p.withBackend(ctx, false, func(tx helper.QueryRunner) errors.ServiceError {
		return p.deleteRecords(ctx, tx, resource, ids)
	})
	if err != nil {
		logger.Print(err)
	} else {
		logger.Tracef("records deleted: %v / %v / %v", resource.Namespace, resource.Name, ids)
	}

	return err
}
