package common

import (
	"context"
	"database/sql"
	"github.com/apibrew/apibrew/pkg/backend/helper"
	"github.com/apibrew/apibrew/pkg/errors"
	"github.com/apibrew/apibrew/pkg/logging"
	"github.com/apibrew/apibrew/pkg/model"
	log "github.com/sirupsen/logrus"
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

	for propertyName, property := range resource.Properties {
		if helper.IsPropertyOmitted(property) {
			continue
		}

		col := p.options.Quote(propertyName)
		cols = append(cols, col)
	}

	return cols
}
