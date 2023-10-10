package common

import (
	"context"
	"github.com/apibrew/apibrew/pkg/errors"
	"github.com/apibrew/apibrew/pkg/logging"
	log "github.com/sirupsen/logrus"
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

	_, intErr := conn.ExecContext(ctx, "SELECT 1")

	err = p.handleDbError(ctx, intErr)

	if err != nil {
		return
	}

	err = p.handleDbError(ctx, conn.Ping())

	testConnection = err == nil

	log.Debugf("Connection status: %s => %v", p.dataSourceName, testConnection)

	return
}
