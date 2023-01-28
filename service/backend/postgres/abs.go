package postgres

import (
	"context"
	"data-handler/logging"
	"data-handler/service/errors"
	"database/sql"
	"fmt"
	log "github.com/sirupsen/logrus"
)

func (p *postgresResourceServiceBackend) acquireConnection(ctx context.Context) (*sql.DB, errors.ServiceError) {
	logger := log.WithFields(logging.CtxFields(ctx))

	if p.connection == nil {

		params := p.connectionDetails.PostgresqlParams

		connStr := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?sslmode=disable", params.Username, params.Password, params.Host, params.Port, params.DbName)
		// Connect to database
		conn, sqlErr := sql.Open("postgres", connStr)
		err := handleDbError(ctx, sqlErr)

		if err != nil {
			return nil, err
		}

		p.connection = conn

		logger.Infof("Connected to Datasource: %s@%s:%d/%s", params.Username, params.Host, params.Port, params.DefaultSchema)
	}

	return p.connection, nil
}
