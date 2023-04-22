package common

import (
	"context"
	"database/sql"
	log "github.com/sirupsen/logrus"
	"github.com/tislib/apibrew/pkg/errors"
	"github.com/tislib/apibrew/pkg/logging"
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

		logger.Infof("Connected to Datasource: %s", p.dataSourceName)
	}

	return p.connection, nil
}
