package postgres

import (
	"data-handler/service/backend"
	"data-handler/stub/model"
)

type postgresDataSourceBackend struct {
	Options      *model.PostgresqlOptions
	dataSourceId string
}

func (p postgresDataSourceBackend) GetDataSourceId() string {
	return p.dataSourceId
}

func NewPostgresDataSourceBackend(dataSourceId string, options *model.PostgresqlOptions) backend.DataSourceBackend {
	return &postgresDataSourceBackend{
		dataSourceId: dataSourceId,
		Options:      options,
	}
}
