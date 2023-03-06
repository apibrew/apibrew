package postgres

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/tislib/data-handler/pkg/abs"
	"github.com/tislib/data-handler/pkg/model"
)

const DbNameType = "VARCHAR(64)"

type postgresResourceServiceBackend struct {
	connectionDetails *model.DataSource_PostgresqlParams
	connection        *sql.DB
	transactionMap    map[string]*txData
	dataSourceName    string
}

func NewPostgresResourceServiceBackend(dataSource *model.DataSource) abs.Backend {
	connectionDetails := dataSource.GetOptions()
	return &postgresResourceServiceBackend{
		dataSourceName:    dataSource.Name,
		connectionDetails: connectionDetails.(*model.DataSource_PostgresqlParams),
		transactionMap:    make(map[string]*txData),
	}
}
