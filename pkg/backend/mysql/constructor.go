package mysql

import (
	"github.com/apibrew/apibrew/pkg/abs"
	"github.com/apibrew/apibrew/pkg/backend/common"
	"github.com/apibrew/apibrew/pkg/model"
	_ "github.com/lib/pq"
)

func NewMysqlResourceServiceBackend(dataSource *model.DataSource) abs.Backend {
	return common.NewSqlBackend(dataSource, &mysqlBackendOptions{
		dataSource:        dataSource,
		connectionDetails: dataSource.Params.(*model.DataSource_MysqlParams),
	})
}
