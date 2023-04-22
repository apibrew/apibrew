package postgres

import (
	_ "github.com/lib/pq"
	"github.com/tislib/apibrew/pkg/abs"
	"github.com/tislib/apibrew/pkg/backend/common"
	"github.com/tislib/apibrew/pkg/model"
)

func NewPostgresResourceServiceBackend(dataSource *model.DataSource) abs.Backend {
	return common.NewSqlBackend(dataSource, &postgreSqlBackendOptions{
		connectionDetails: dataSource.Params.(*model.DataSource_PostgresqlParams),
	})
}
