package postgres

import (
	"context"
	log "github.com/sirupsen/logrus"
	"github.com/tislib/data-handler/logging"
	"github.com/tislib/data-handler/service/errors"
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
