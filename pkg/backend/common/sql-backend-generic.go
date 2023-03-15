package common

import (
	"context"
	log "github.com/sirupsen/logrus"
	"github.com/tislib/data-handler/pkg/errors"
	"github.com/tislib/data-handler/pkg/logging"
)

func (p *sqlBackend) DestroyDataSource(ctx context.Context) {
	logger := log.WithFields(logging.CtxFields(ctx))

	if p.connection != nil {
		err := p.connection.Close()

		if err != nil {
			logger.Error(err)
		}

		p.connection = nil
	}
}

func (p *sqlBackend) GetStatus(ctx context.Context) (connectionAlreadyInitiated bool, testConnection bool, err errors.ServiceError) {
	conn, err := p.acquireConnection(ctx)

	if err != nil {
		return
	}

	err = p.handleDbError(ctx, conn.Ping())

	testConnection = err == nil

	return
}
