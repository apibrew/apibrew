package mongo

import (
	"data-handler/model"
	"data-handler/service/backend"
)

type postgresDataSourceBackend struct {
	Options      *model.PostgresqlOptions
	dataSourceId string
	backend      model.DataSourceBackendType
}

func (p postgresDataSourceBackend) GetBackendType() model.DataSourceBackendType {
	return p.backend
}

func (p postgresDataSourceBackend) GetDataSourceId() string {
	return p.dataSourceId
}

func NewPostgresDataSourceBackend(dataSourceId string, options *model.PostgresqlOptions) backend.DataSourceConnectionDetails {
	return &postgresDataSourceBackend{
		dataSourceId: dataSourceId,
		Options:      options,
	}
}
