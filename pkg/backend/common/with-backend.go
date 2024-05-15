package common

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/apibrew/apibrew/pkg/backend/helper"
	"github.com/apibrew/apibrew/pkg/logging"
	log "github.com/sirupsen/logrus"
)

type queryLoggerStruct struct {
	delegate       helper.QueryRunner
	dataSourceName string
}

func (q queryLoggerStruct) logQuery(ctx context.Context, query string, args ...any) {
	logger := log.WithFields(logging.CtxFields(ctx))

	var argsToPrint []string

	for _, arg := range args {
		str := fmt.Sprintf("%v", arg)

		if len(str) > 10 {
			str = fmt.Sprintf("%s...", str[:10])
		}

		argsToPrint = append(argsToPrint, str)
	}

	logger.Debugf("Log SQL[%s]: %s ; Bind Params: %v", q.dataSourceName, query, argsToPrint)
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

func (p *sqlBackend) queryLogger(dataSourceName string, runner helper.QueryRunner) helper.QueryRunner {
	return queryLoggerStruct{dataSourceName: dataSourceName, delegate: runner}
}

func (p *sqlBackend) withBackend(ctx context.Context, readOnly bool, fn func(tx helper.QueryRunner) error) error {
	logger := log.WithFields(logging.CtxFields(ctx))

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

	serviceErr = fn(p.queryLogger(p.dataSourceName, tx))

	if serviceErr != nil {
		logger.Errorf("Rollback: %s", serviceErr)
		return serviceErr
	}

	serviceErr = p.handleDbError(ctx, tx.Commit())
	logger.Tracef("end transaction readonly=%v", readOnly)

	return serviceErr
}
