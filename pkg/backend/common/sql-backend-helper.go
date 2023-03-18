package common

import (
	"context"
	"database/sql"
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"github.com/tislib/data-handler/pkg/errors"
	"github.com/tislib/data-handler/pkg/logging"
	"github.com/tislib/data-handler/pkg/model"
	"github.com/tislib/data-handler/pkg/types"
	"google.golang.org/protobuf/types/known/structpb"
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

func (p *sqlBackend) DbEncode(property *model.ResourceProperty, packedVal *structpb.Value) (interface{}, errors.ServiceError) {
	propertyType := types.ByResourcePropertyType(property.Type)
	var val interface{}

	if property.Type == model.ResourceProperty_OBJECT || property.Type == model.ResourceProperty_ENUM || property.Type == model.ResourceProperty_MAP || property.Type == model.ResourceProperty_LIST {
		var err error
		val, err = json.Marshal(packedVal.AsInterface())

		if err != nil {
			return nil, errors.InternalError.WithDetails(err.Error())
		}
		val = string(val.([]byte))
	} else {
		var err error
		val, err = propertyType.UnPack(packedVal)

		if err != nil {
			return nil, errors.InternalError.WithDetails(err.Error())
		}
	}
	return val, nil
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
