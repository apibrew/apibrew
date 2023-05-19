package common

import (
	"context"
	"database/sql"
	"github.com/apibrew/apibrew/pkg/errors"
	"github.com/apibrew/apibrew/pkg/logging"
	log "github.com/sirupsen/logrus"
)

func (p *sqlBackend) acquireConnection(ctx context.Context) (*sql.DB, errors.ServiceError) {
	logger := log.WithFields(logging.CtxFields(ctx))

	if p.connection == nil {

		connStr := p.options.GetConnectionString()

		// Connect to database
		conn, sqlErr := sql.Open(p.options.GetDriverName(), connStr)
		err := p.handleDbError(ctx, sqlErr)

		if err != nil {
			return nil, err
		}

		p.connection = conn

		logger.Infof("Connecting to Datasource: %s", p.dataSourceName)

		_, _, err = p.GetStatus(ctx)

		if err != nil {
			return nil, err
		}

		logger.Infof("Connected to Datasource: %s", p.dataSourceName)
	}

	return p.connection, nil
}
