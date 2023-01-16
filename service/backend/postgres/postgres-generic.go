package postgres

import (
	"context"
	"data-handler/service/errors"
	log "github.com/sirupsen/logrus"
)

func (p *postgresResourceServiceBackend) DestroyDataSource(ctx context.Context) {
	if p.connection != nil {
		err := p.connection.Close()

		if err != nil {
			log.Error(err)
		}

		p.connection = nil
	}
}

func (p *postgresResourceServiceBackend) GetStatus(ctx context.Context) (connectionAlreadyInitiated bool, testConnection bool, err errors.ServiceError) {
	conn, err := p.acquireConnection(ctx)

	if err != nil {
		return
	}

	err = handleDbError(conn.Ping())

	testConnection = err == nil

	return
}
