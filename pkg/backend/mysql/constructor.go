package mysql

import (
	"github.com/apibrew/apibrew/pkg/abs"
	"github.com/apibrew/apibrew/pkg/backend/common"
	"github.com/apibrew/apibrew/pkg/modelnew"
	_ "github.com/lib/pq"
)

func NewMysqlResourceServiceBackend(dataSource *modelnew.DataSource) abs.Backend {
	return common.NewSqlBackend(dataSource, &mysqlBackendOptions{
		dataSource: dataSource,
	})
}
