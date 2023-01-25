package postgres

import (
	"context"
	"data-handler/service/errors"
	"database/sql"
	log "github.com/sirupsen/logrus"
)

func (p *postgresResourceServiceBackend) withBackend(ctx context.Context, readOnly bool, fn func(tx *sql.Tx) errors.ServiceError) errors.ServiceError {
	transactionKey := ctx.Value(ctxTransactionKey)

	if transactionKey != nil {
		txDataInstance := p.transactionMap[transactionKey.(string)]

		return fn(txDataInstance.tx)
	}

	log.Tracef("begin transaction readonly=%v", readOnly)
	conn, serviceErr := p.acquireConnection(ctx)

	if serviceErr != nil {
		return serviceErr
	}

	tx, err := conn.BeginTx(context.TODO(), &sql.TxOptions{
		ReadOnly: readOnly,
	})

	if err != nil {
		log.Errorf("Unable to begin transaction: %s", err)
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
	log.Tracef("end transaction readonly=%v", readOnly)

	return serviceErr
}
