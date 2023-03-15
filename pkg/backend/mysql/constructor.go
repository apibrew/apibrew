package mysql

import (
	_ "github.com/lib/pq"
	"github.com/tislib/data-handler/pkg/abs"
	"github.com/tislib/data-handler/pkg/backend/common"
	_ "github.com/tislib/data-handler/pkg/backend/postgres/sql/statik"
	"github.com/tislib/data-handler/pkg/model"
)

func NewMysqlResourceServiceBackend(dataSource *model.DataSource) abs.Backend {
	return common.NewSqlBackend(dataSource, &mysqlBackendOptions{
		connectionDetails: dataSource.Options.(*model.DataSource_MysqlParams),
	})
	//connectionDetails := dataSource.GetOptions()
	//return &postgresResourceServiceBackend{
	//	dataSourceName:    dataSource.Name,
	//	connectionDetails: connectionDetails.(*model.DataSource_PostgresqlParams),
	//	transactionMap:    make(map[string]*txData),
	//}
}
