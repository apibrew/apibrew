package test

import "data-handler/stub/model"

var dataSource1 = &model.DataSource{
	Backend: model.DataSourceBackend_POSTGRESQL,
	Options: &model.DataSource_PostgresqlParams{
		PostgresqlParams: &model.PostgresqlOptions{
			Username:      "root",
			Password:      "52fa536f0c5b85f9d806633937f06446",
			Host:          "tiswork.tisserv.net",
			Port:          5432,
			DbName:        "market",
			DefaultSchema: "public",
		},
	},
}
