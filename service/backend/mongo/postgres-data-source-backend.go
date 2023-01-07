package mongo

import (
	"data-handler/model"
	"data-handler/service/backend"
)

type postgresDataSourceBackend struct {
	Options      *model.PostgresqlOptions
	dataSourceId string
	backend      model.DataSourceBackend
}

func (p postgresDataSourceBackend) GetBackend() model.DataSourceBackend {
	return p.backend
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
