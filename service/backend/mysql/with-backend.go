package mysql

import (
	"context"
	"data-handler/service/errors"
	"database/sql"
	log "github.com/sirupsen/logrus"
)

func (p *postgresResourceServiceBackend) withSystemBackend(ctx context.Context, readOnly bool, fn func(tx *sql.Tx) errors.ServiceError) errors.ServiceError {
	return p.withBackend(ctx, p.systemBackend.GetDataSourceId(), readOnly, fn)
}

func (p *postgresResourceServiceBackend) withBackend(ctx context.Context, dataSourceId string, readOnly bool, fn func(tx *sql.Tx) errors.ServiceError) errors.ServiceError {
	log.Tracef("begin transaction: %s, readonly=%v", dataSourceId, readOnly)
	conn, serviceErr := p.acquireConnection(ctx, dataSourceId)

	if serviceErr != nil {
		return serviceErr
	}

	tx, err := conn.BeginTx(context.TODO(), &sql.TxOptions{
		ReadOnly: readOnly,
	})

	if err != nil {
		log.Errorf("Unable to begin transaction: %s %s", err, dataSourceId)
		return handleDbError(err)
	}

	defer func(tx *sql.Tx) {
		_ = tx.Rollback()
	}(tx)

	serviceErr = fn(tx)

	if serviceErr != nil {
		log.Errorf("Rollback: %s", serviceErr)
		return serviceErr
	}

	serviceErr = handleDbError(tx.Commit())
	log.Tracef("end transaction: %s, readonly=%v", dataSourceId, readOnly)

	return serviceErr
}
