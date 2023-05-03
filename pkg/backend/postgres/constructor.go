package postgres

import (
	"github.com/apibrew/apibrew/pkg/abs"
	"github.com/apibrew/apibrew/pkg/backend/common"
	"github.com/apibrew/apibrew/pkg/model"
	_ "github.com/lib/pq"
)

func NewPostgresResourceServiceBackend(dataSource *model.DataSource) abs.Backend {
	return common.NewSqlBackend(dataSource, &postgreSqlBackendOptions{
		connectionDetails: dataSource.Params.(*model.DataSource_PostgresqlParams),
	})
}
