package mysql

import (
	_ "github.com/lib/pq"
	"github.com/tislib/apibrew/pkg/abs"
	"github.com/tislib/apibrew/pkg/backend/common"
	"github.com/tislib/apibrew/pkg/model"
)

func NewMysqlResourceServiceBackend(dataSource *model.DataSource) abs.Backend {
	return common.NewSqlBackend(dataSource, &mysqlBackendOptions{
		dataSource:        dataSource,
		connectionDetails: dataSource.Params.(*model.DataSource_MysqlParams),
	})
}
