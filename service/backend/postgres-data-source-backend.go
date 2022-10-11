package backend

import "data-handler/stub/model"

type postgresDataSourceBackend struct {
	Options      *model.PostgresqlOptions
	dataSourceId string
}

func (p postgresDataSourceBackend) getDataSourceId() string {
	return p.dataSourceId
}

func NewPostgresDataSourceBackend(dataSourceId string, options *model.PostgresqlOptions) DataSourceBackend {
	return &postgresDataSourceBackend{
		dataSourceId: dataSourceId,
		Options:      options,
	}
}
