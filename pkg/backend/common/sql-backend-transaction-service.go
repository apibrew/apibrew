package common

import (
	"context"
	"database/sql"
	"github.com/apibrew/apibrew/pkg/abs"
	"github.com/apibrew/apibrew/pkg/errors"
	"github.com/apibrew/apibrew/pkg/helper"
	"github.com/apibrew/apibrew/pkg/logging"
	log "github.com/sirupsen/logrus"
	"time"
)

type txData struct {
	tx     *sql.Tx
	cancel context.CancelFunc
}

func (p *sqlBackend) BeginTransaction(ctx context.Context, readOnly bool) (string, errors.ServiceError) {
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
		return "", p.handleDbError(ctx, err)
	}

	txDataInstance := &txData{
		tx:     tx,
		cancel: cancel,
	}

	transactionKey := helper.RandStringRunes(8)

	p.mu.Lock()
	p.transactionMap[transactionKey] = txDataInstance
	p.mu.Unlock()

	go func() {
		<-transactionCtx.Done()
		p.mu.Lock()
		delete(p.transactionMap, transactionKey)
		p.mu.Unlock()

		log.Println(transactionCtx.Err())

		if transactionCtx.Err() == context.DeadlineExceeded {
			logger.Warnf("Timeout Rollback transaction: " + transactionKey)

			err = tx.Rollback()

			if err != nil {
				logger.Error(err)
			}
		}
	}()

	logger.Debugf("BeginTransaction: %s / %s", p.dataSourceName, transactionKey)

	return transactionKey, nil
}

func (p *sqlBackend) CommitTransaction(ctx context.Context) (serviceError errors.ServiceError) {
	logger := log.WithFields(logging.CtxFields(ctx))

	logger.Tracef("CommitTransaction")

	transactionKey := ctx.Value(abs.TransactionContextKey)

	if transactionKey == nil {
		return errors.LogicalError.WithDetails("Transaction not found")
	}

	logger.Debugf("CommitTransaction %s", transactionKey)

	p.mu.Lock()
	txDataInstance := p.transactionMap[transactionKey.(string)]
	p.mu.Unlock()

	if txDataInstance == nil {
		return errors.LogicalError.WithDetails("Transaction not found: " + transactionKey.(string))
	}

	err := txDataInstance.tx.Commit()
	txDataInstance.cancel()

	return p.handleDbError(ctx, err)
}

func (p *sqlBackend) RollbackTransaction(ctx context.Context) (serviceError errors.ServiceError) {
	logger := log.WithFields(logging.CtxFields(ctx))

	logger.Tracef("RollbackTransaction")

	transactionKey := ctx.Value(abs.TransactionContextKey)

	if transactionKey == nil {
		return errors.LogicalError.WithDetails("Transaction not found")
	}

	logger.Debugf("RollbackTransaction %s", transactionKey)

	p.mu.Lock()
	txDataInstance := p.transactionMap[transactionKey.(string)]
	p.mu.Unlock()

	if txDataInstance == nil {
		return errors.LogicalError.WithDetails("Transaction not found: " + transactionKey.(string))
	}

	err := txDataInstance.tx.Rollback()
	txDataInstance.cancel()

	return p.handleDbError(ctx, err)
}

func (p *sqlBackend) IsTransactionAlive(_ context.Context) (isAlive bool, serviceError errors.ServiceError) {
	//TODO implement me
	panic("implement me")
}
