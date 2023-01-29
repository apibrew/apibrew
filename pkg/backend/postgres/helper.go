package postgres

import (
	"context"
	"database/sql"
	"github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"github.com/tislib/data-handler/pkg/errors"
	"github.com/tislib/data-handler/pkg/logging"
	"github.com/tislib/data-handler/pkg/model"
	"net"
	"runtime/debug"
)

func locatePropertyByName(resource *model.Resource, propertyName string) *model.ResourceProperty {
	for _, property := range resource.Properties {
		if property.Name == propertyName {
			return property
		}
	}

	return nil
}

func handleDbError(ctx context.Context, err error) errors.ServiceError {
	logger := log.WithFields(logging.CtxFields(ctx))

	if err == nil {
		return nil
	}

	if err == sql.ErrNoRows {
		return errors.RecordNotFoundError
	}

	logger.Errorf("Db Error: %s", err)
	logger.Tracef("Stack: " + string(debug.Stack()))

	if err == sql.ErrTxDone {
		logger.Panic("Illegal situation")
	}

	if _, ok := err.(errors.ServiceError); ok {
		logger.Panic("database error is expected: ", err)
	}

	if pqErr, ok := err.(*pq.Error); ok {
		return handlePqErr(pqErr)
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

func handlePqErr(err *pq.Error) errors.ServiceError {
	switch err.Code {
	case "28000":
		return errors.BackendConnectionAuthenticationError.WithMessage(err.Message)
	case "28P01":
		return errors.BackendConnectionAuthenticationError.WithMessage(err.Message)
	case "23505":
		return errors.UniqueViolation.WithDetails(err.Message)
	case "23503":
		return errors.ForeignKeyViolation.WithDetails(err.Message)
	default:
		return errors.InternalError.WithMessage(err.Message)
	}
}
