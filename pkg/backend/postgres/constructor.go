package postgres

import (
	"github.com/apibrew/apibrew/pkg/abs"
	"github.com/apibrew/apibrew/pkg/backend/common"
	"github.com/apibrew/apibrew/pkg/resource_model"
	_ "github.com/lib/pq"
)

func NewPostgresResourceServiceBackend(dataSource abs.DataSource) abs.Backend {
	return common.NewSqlBackend(dataSource, &postgreSqlBackendOptions{
		dataSource: dataSource.(*resource_model.DataSource),
	})
}
