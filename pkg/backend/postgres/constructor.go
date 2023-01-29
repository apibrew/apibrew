package postgres

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/tislib/data-handler/pkg/backend"
	"github.com/tislib/data-handler/pkg/model"
)

const DbNameType = "VARCHAR(64)"

type postgresResourceServiceBackend struct {
	connectionDetails *model.DataSource_PostgresqlParams
	connection        *sql.DB
	transactionMap    map[string]*txData
}

func NewPostgresResourceServiceBackend(connectionDetails backend.DataSourceConnectionDetails) backend.Backend {
	return &postgresResourceServiceBackend{
		connectionDetails: connectionDetails.(*model.DataSource_PostgresqlParams),
		transactionMap:    make(map[string]*txData),
	}
}
