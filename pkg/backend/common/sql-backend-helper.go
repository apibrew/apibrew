package common

import (
	"context"
	"database/sql"
	log "github.com/sirupsen/logrus"
	"github.com/tislib/apibrew/pkg/errors"
	"github.com/tislib/apibrew/pkg/logging"
	"github.com/tislib/apibrew/pkg/model"
	"net"
	"runtime/debug"
)

func (p *sqlBackend) handleDbError(ctx context.Context, err error) errors.ServiceError {
	logger := log.WithFields(logging.CtxFields(ctx))

	if err == nil {
		return nil
	}

	logger.Errorf("Db Error: %s", err)

	if err == sql.ErrNoRows {
		return errors.RecordNotFoundError
	}

	logger.Debug("Stack: " + string(debug.Stack()))

	if err == sql.ErrTxDone {
		logger.Panic("Illegal situation")
	}

	if _, ok := err.(errors.ServiceError); ok {
		logger.Panic("database error is expected: ", err)
	}

	if handledErr, handled := p.options.HandleError(err); handled {
		return handledErr
	}

	if netErr, ok := err.(*net.OpError); ok {
		return errors.InternalError.WithDetails(netErr.Error())
	}

	if err.Error() == "context cancelled" {
		return errors.InternalError.WithDetails(err.Error())
	}

	logger.Print("Unhandled Error: ", err)
	return errors.InternalError.WithDetails(err.Error())
}

func (p *sqlBackend) getFullTableName(sourceConfig *model.ResourceSourceConfig) string {
	return p.options.GetFullTableName(sourceConfig)
}

func (p *sqlBackend) prepareResourceRecordCols(resource *model.Resource) []string {
	var cols []string

	for _, property := range resource.Properties {
		col := p.options.Quote(property.Mapping)
		cols = append(cols, col)
	}

	return cols
}
