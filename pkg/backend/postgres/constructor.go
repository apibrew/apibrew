package postgres

import (
	_ "github.com/lib/pq"
	"github.com/tislib/data-handler/pkg/abs"
	"github.com/tislib/data-handler/pkg/backend/common"
	"github.com/tislib/data-handler/pkg/model"
)

func NewPostgresResourceServiceBackend(dataSource *model.DataSource) abs.Backend {
	return common.NewSqlBackend(dataSource, &postgreSqlBackendOptions{
		connectionDetails: dataSource.Params.(*model.DataSource_PostgresqlParams),
	})
}
