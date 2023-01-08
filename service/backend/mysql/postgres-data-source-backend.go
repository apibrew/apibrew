package mysql

import (
	"data-handler/model"
	"data-handler/service/backend"
)

type postgresDataSourceBackend struct {
	Options      *model.PostgresqlOptions
	dataSourceId string
}

func (p postgresDataSourceBackend) GetBackendType() model.DataSourceBackendType {
	//TODO implement me
	panic("implement me")
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
