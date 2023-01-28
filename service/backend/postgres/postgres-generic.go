package postgres

import (
	"context"
	"data-handler/logging"
	"data-handler/service/errors"
	log "github.com/sirupsen/logrus"
)

func (p *postgresResourceServiceBackend) DestroyDataSource(ctx context.Context) {
	logger := log.WithFields(logging.CtxFields(ctx))

	if p.connection != nil {
		err := p.connection.Close()

		if err != nil {
			logger.Error(err)
		}

		p.connection = nil
	}
}

func (p *postgresResourceServiceBackend) GetStatus(ctx context.Context) (connectionAlreadyInitiated bool, testConnection bool, err errors.ServiceError) {
	conn, err := p.acquireConnection(ctx)

	if err != nil {
		return
	}

	err = handleDbError(ctx, conn.Ping())

	testConnection = err == nil

	return
}
