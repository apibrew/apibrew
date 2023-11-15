package common

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/apibrew/apibrew/pkg/abs"
	"github.com/apibrew/apibrew/pkg/backend/helper"
	"github.com/apibrew/apibrew/pkg/errors"
	"github.com/apibrew/apibrew/pkg/logging"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/service/annotations"
	"github.com/apibrew/apibrew/pkg/util"
	log "github.com/sirupsen/logrus"
)

func (p *sqlBackend) ListRecords(ctx context.Context, resource *model.Resource, params abs.ListRecordParams, resultChan chan<- *model.Record) (result []*model.Record, total uint32, err errors.ServiceError) {
	logger := log.WithFields(logging.CtxFields(ctx))

	logger.Tracef("Begin listing: %s/%s", resource.Namespace, resource.Name)
	err = p.withBackend(ctx, true, func(tx helper.QueryRunner) errors.ServiceError {
		result, total, err = p.recordList(ctx, tx, resource, params, resultChan)

		return err
	})
	logger.Tracef("Finish listing: %s/%s", resource.Namespace, resource.Name)

	return
}

func (p *sqlBackend) AddRecords(ctx context.Context, resource *model.Resource, records []*model.Record) ([]*model.Record, errors.ServiceError) {
	logger := log.WithFields(logging.CtxFields(ctx))

	var err errors.ServiceError

	logger.Tracef("Begin creating: %s/%s", resource.Namespace, resource.Name)

	err = p.withBackend(ctx, false, func(tx helper.QueryRunner) errors.ServiceError {
		maxChunkSize := 1000
		chunkCount := len(records) / maxChunkSize
		if chunkCount == 0 {
			chunkCount = 1
		}

		for i := 0; i < chunkCount; i++ {
			bi := i * maxChunkSize
			ei := (i + 1) * maxChunkSize

			if ei > len(records) {
				ei = len(records)
			}

			batch := records[bi:ei]

			err = p.recordInsert(ctx, tx, resource, batch, annotations.IsEnabled(annotations.FromCtx(ctx), annotations.IgnoreIfExists), p.schema)

			if err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	logger.Tracef("Records created: %s/%s", resource.Namespace, resource.Name)

	return records, nil
}

func (p *sqlBackend) UpdateRecords(ctx context.Context, resource *model.Resource, records []*model.Record) ([]*model.Record, errors.ServiceError) {
	err := p.withBackend(ctx, false, func(tx helper.QueryRunner) errors.ServiceError {
		for _, record := range records {
			err := p.recordUpdate(ctx, tx, resource, record, annotations.IsEnabledOnCtx(ctx, annotations.CheckVersion), p.schema)

			if err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return records, nil
}

func (p *sqlBackend) GetRecord(ctx context.Context, resource *model.Resource, id string, resolveReferences []string) (*model.Record, errors.ServiceError) {
	var record *model.Record = nil
	err := p.withBackend(ctx, true, func(tx helper.QueryRunner) errors.ServiceError {
		var err errors.ServiceError
		record, err = p.readRecord(ctx, tx, resource, id, resolveReferences)

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

func (p *sqlBackend) DeleteRecords(ctx context.Context, resource *model.Resource, records []*model.Record) errors.ServiceError {
	var ids = util.ArrayMap(records, func(record *model.Record) string {
		return util.GetRecordId(resource, record)
	})
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
