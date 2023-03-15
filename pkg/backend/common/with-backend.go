package common

import (
	"context"
	"database/sql"
	log "github.com/sirupsen/logrus"
	"github.com/tislib/data-handler/pkg/abs"
	"github.com/tislib/data-handler/pkg/errors"
	"github.com/tislib/data-handler/pkg/logging"
)

type QueryRunner interface {
	QueryRow(query string, args ...any) *sql.Row
	QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row
	Exec(query string, args ...any) (sql.Result, error)
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
	Query(query string, args ...any) (*sql.Rows, error)
	QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error)
}

type queryLoggerStruct struct {
	delegate       QueryRunner
	dataSourceName string
	transactionKey string
}

func (q queryLoggerStruct) logQuery(ctx context.Context, query string, args ...any) {
	logger := log.WithFields(logging.CtxFields(ctx))

	logger.Debugf("Log SQL[%s/%s]: %s ; Bind Params: %v", q.dataSourceName, q.transactionKey, query, args)
}

func (q queryLoggerStruct) QueryRow(query string, args ...any) *sql.Row {
	q.logQuery(context.TODO(), query, args...)
	return q.delegate.QueryRow(query, args...)
}

func (q queryLoggerStruct) QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row {
	q.logQuery(ctx, query, args...)
	return q.delegate.QueryRowContext(ctx, query, args...)
}

func (q queryLoggerStruct) Exec(query string, args ...any) (sql.Result, error) {
	q.logQuery(context.TODO(), query, args...)
	return q.delegate.Exec(query, args...)
}

func (q queryLoggerStruct) ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error) {
	q.logQuery(ctx, query, args...)
	return q.delegate.ExecContext(ctx, query, args...)
}

func (q queryLoggerStruct) Query(query string, args ...any) (*sql.Rows, error) {
	q.logQuery(context.TODO(), query, args...)
	return q.delegate.Query(query, args...)
}

func (q queryLoggerStruct) QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error) {
	q.logQuery(ctx, query, args...)
	return q.delegate.QueryContext(ctx, query, args...)
}

func (p *sqlBackend) queryLogger(transactionKey, dataSourceName string, runner QueryRunner) QueryRunner {
	if transactionKey == "" {
		transactionKey = "default"
	}
	return queryLoggerStruct{transactionKey: transactionKey, dataSourceName: dataSourceName, delegate: runner}
}

func (p *sqlBackend) withBackend(ctx context.Context, readOnly bool, fn func(tx QueryRunner) errors.ServiceError) errors.ServiceError {
	logger := log.WithFields(logging.CtxFields(ctx))

	transactionKey := ctx.Value(abs.TransactionContextKey)

	if transactionKey != nil {
		txDataInstance := p.transactionMap[transactionKey.(string)]

		if txDataInstance == nil {
			return errors.LogicalError.WithDetails("Transaction not found: " + transactionKey.(string))
		}

		return fn(p.queryLogger(transactionKey.(string), p.dataSourceName, txDataInstance.tx))
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
		return p.handleDbError(ctx, err)
	}

	defer func(tx *sql.Tx) {
		_ = tx.Rollback()
	}(tx)

	serviceErr = fn(p.queryLogger("", p.dataSourceName, tx))

	if serviceErr != nil {
		logger.Errorf("Rollback: %s", serviceErr)
		return serviceErr
	}

	serviceErr = p.handleDbError(ctx, tx.Commit())
	logger.Tracef("end transaction readonly=%v", readOnly)

	return serviceErr
}
