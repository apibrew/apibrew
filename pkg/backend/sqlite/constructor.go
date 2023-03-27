package sqlite

import (
	_ "github.com/lib/pq"
	"github.com/tislib/data-handler/pkg/abs"
	"github.com/tislib/data-handler/pkg/backend/common"
	"github.com/tislib/data-handler/pkg/model"
)

func NewSqliteResourceServiceBackend(dataSource *model.DataSource) abs.Backend {
	return common.NewSqlBackend(dataSource, &sqliteBackendOptions{
		dataSource:        dataSource,
		connectionDetails: dataSource.Params.(*model.DataSource_SqliteParams),
	})
}
