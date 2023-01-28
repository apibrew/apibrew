package postgres

import (
	"context"
	"database/sql"
	log "github.com/sirupsen/logrus"
	"github.com/tislib/data-handler/logging"
	"github.com/tislib/data-handler/service/errors"
	"time"
)

const ctxTransactionKey = "transactionKey"

type txData struct {
	tx     *sql.Tx
	cancel context.CancelFunc
}

func (p *postgresResourceServiceBackend) BeginTransaction(ctx context.Context, readOnly bool) (transactionKey string, serviceError errors.ServiceError) {
	logger := log.WithFields(logging.CtxFields(ctx))

	logger.Tracef("begin transaction readonly=%v", readOnly)
	conn, serviceErr := p.acquireConnection(ctx)

	if serviceErr != nil {
		return "", serviceErr
	}

	transactionCtx, cancel := context.WithTimeout(context.TODO(), time.Second*30)

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

	p.transactionMap[transactionKey] = txDataInstance

	go func() {
		<-transactionCtx.Done()
		delete(p.transactionMap, transactionKey)
	}()

	return transactionKey, nil
}

func (p *postgresResourceServiceBackend) CommitTransaction(ctx context.Context) (serviceError errors.ServiceError) {
	transactionKey := ctx.Value(ctxTransactionKey)

	if transactionKey == nil {
		return errors.LogicalError.WithDetails("Transaction not found")
	}

	txDataInstance := p.transactionMap[transactionKey.(string)]

	if txDataInstance == nil {
		return errors.LogicalError.WithDetails("Transaction not found: " + transactionKey.(string))
	}

	err := txDataInstance.tx.Commit()
	txDataInstance.cancel()

	return handleDbError(ctx, err)
}

func (p *postgresResourceServiceBackend) RollbackTransaction(ctx context.Context) (serviceError errors.ServiceError) {
	transactionKey := ctx.Value(ctxTransactionKey)

	if transactionKey == nil {
		return errors.LogicalError.WithDetails("Transaction not found")
	}

	txDataInstance := p.transactionMap[transactionKey.(string)]

	if txDataInstance == nil {
		return errors.LogicalError.WithDetails("Transaction not found: " + transactionKey.(string))
	}

	err := txDataInstance.tx.Rollback()
	txDataInstance.cancel()

	return handleDbError(ctx, err)
}

func (p *postgresResourceServiceBackend) IsTransactionAlive(ctx context.Context) (isAlive bool, serviceError errors.ServiceError) {
	//TODO implement me
	panic("implement me")
}
