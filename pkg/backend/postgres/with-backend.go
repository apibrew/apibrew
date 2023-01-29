package postgres

import (
	"context"
	"database/sql"
	log "github.com/sirupsen/logrus"
	"github.com/tislib/data-handler/pkg/errors"
	"github.com/tislib/data-handler/pkg/logging"
)

func (p *postgresResourceServiceBackend) withBackend(ctx context.Context, readOnly bool, fn func(tx *sql.Tx) errors.ServiceError) errors.ServiceError {
	logger := log.WithFields(logging.CtxFields(ctx))

	transactionKey := ctx.Value(ctxTransactionKey)

	if transactionKey != nil {
		txDataInstance := p.transactionMap[transactionKey.(string)]

		if txDataInstance == nil {
			return errors.LogicalError.WithDetails("Transaction not found: " + transactionKey.(string))
		}

		return fn(txDataInstance.tx)
	}

	logger.Tracef("begin transaction readonly=%v", readOnly)
	conn, serviceErr := p.acquireConnection(ctx)

	if serviceErr != nil {
		return serviceErr
	}

	tx, err := conn.BeginTx(context.TODO(), &sql.TxOptions{
		ReadOnly: readOnly,
	})

	if err != nil {
		logger.Errorf("Unable to begin transaction: %s", err)
		return handleDbError(ctx, err)
	}

	defer func(tx *sql.Tx) {
		_ = tx.Rollback()
	}(tx)

	serviceErr = fn(tx)

	if serviceErr != nil {
		logger.Errorf("Rollback: %s", serviceErr)
		return serviceErr
	}

	serviceErr = handleDbError(ctx, tx.Commit())
	logger.Tracef("end transaction readonly=%v", readOnly)

	return serviceErr
}
