package postgres

import (
	"context"
	"database/sql"
	log "github.com/sirupsen/logrus"
	"github.com/tislib/data-handler/pkg/abs"
	"github.com/tislib/data-handler/pkg/errors"
	"github.com/tislib/data-handler/pkg/helper"
	"github.com/tislib/data-handler/pkg/logging"
	"time"
)

type txData struct {
	tx     *sql.Tx
	cancel context.CancelFunc
}

func (p *postgresResourceServiceBackend) BeginTransaction(ctx context.Context, readOnly bool) (string, errors.ServiceError) {
	logger := log.WithFields(logging.CtxFields(ctx))

	logger.Tracef("begin transaction readonly=%v", readOnly)
	conn, serviceErr := p.acquireConnection(ctx)

	if serviceErr != nil {
		return "", serviceErr
	}

	transactionCtx, cancel := context.WithTimeout(context.TODO(), time.Minute*10)

	tx, err := conn.BeginTx(transactionCtx, &sql.TxOptions{
		ReadOnly: readOnly,
	})

	if err != nil {
		cancel()
		return "", handleDbError(ctx, err)
	}

	txDataInstance := &txData{
		tx:     tx,
		cancel: cancel,
	}

	transactionKey := helper.RandStringRunes(8)

	p.transactionMap[transactionKey] = txDataInstance

	go func() {
		<-transactionCtx.Done()
		delete(p.transactionMap, transactionKey)

		log.Println(transactionCtx.Err())

		if transactionCtx.Err() == context.DeadlineExceeded {
			logger.Warnf("Timeout Rollback transaction: " + transactionKey)

			err = tx.Rollback()

			if err != nil {
				logger.Error(err)
			}
		}
	}()

	logger.Debugf("BeginTransaction: " + transactionKey)

	return transactionKey, nil
}

func (p *postgresResourceServiceBackend) CommitTransaction(ctx context.Context) (serviceError errors.ServiceError) {
	logger := log.WithFields(logging.CtxFields(ctx))

	logger.Tracef("CommitTransaction")

	transactionKey := ctx.Value(abs.TransactionContextKey)

	if transactionKey == nil {
		return errors.LogicalError.WithDetails("Transaction not found")
	}

	logger.Debugf("CommitTransaction %s", transactionKey)

	txDataInstance := p.transactionMap[transactionKey.(string)]

	if txDataInstance == nil {
		return errors.LogicalError.WithDetails("Transaction not found: " + transactionKey.(string))
	}

	err := txDataInstance.tx.Commit()
	txDataInstance.cancel()

	return handleDbError(ctx, err)
}

func (p *postgresResourceServiceBackend) RollbackTransaction(ctx context.Context) (serviceError errors.ServiceError) {
	logger := log.WithFields(logging.CtxFields(ctx))

	logger.Tracef("RollbackTransaction")

	transactionKey := ctx.Value(abs.TransactionContextKey)

	if transactionKey == nil {
		return errors.LogicalError.WithDetails("Transaction not found")
	}

	logger.Debugf("RollbackTransaction %s", transactionKey)

	txDataInstance := p.transactionMap[transactionKey.(string)]

	if txDataInstance == nil {
		return errors.LogicalError.WithDetails("Transaction not found: " + transactionKey.(string))
	}

	err := txDataInstance.tx.Rollback()
	txDataInstance.cancel()

	return handleDbError(ctx, err)
}

func (p *postgresResourceServiceBackend) IsTransactionAlive(_ context.Context) (isAlive bool, serviceError errors.ServiceError) {
	//TODO implement me
	panic("implement me")
}
