package test

import "data-handler/stub/model"

var systemDataSource = &model.DataSource{
	Id:      "system",
	Backend: model.DataSourceBackend_POSTGRESQL,
	Type:    model.DataType_SYSTEM,
	Options: &model.DataSource_PostgresqlParams{
		PostgresqlParams: &model.PostgresqlOptions{
			Username:      "root",
			Password:      "root",
			Host:          "127.0.0.1",
			Port:          5432,
			DbName:        "postgres",
			DefaultSchema: "public",
		},
	},
}

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

var dataSource1WrongPassword = &model.DataSource{
	Backend: model.DataSourceBackend_POSTGRESQL,
	Options: &model.DataSource_PostgresqlParams{
		PostgresqlParams: &model.PostgresqlOptions{
			Username:      "root",
			Password:      "root",
			Host:          "tiswork.tisserv.net",
			Port:          5432,
			DbName:        "market",
			DefaultSchema: "public",
		},
	},
}

var dataSourceDhTest = &model.DataSource{
	Backend: model.DataSourceBackend_POSTGRESQL,
	Options: &model.DataSource_PostgresqlParams{
		PostgresqlParams: &model.PostgresqlOptions{
			Username:      "dh_test",
			Password:      "dh_test",
			Host:          "127.0.0.1",
			Port:          5432,
			DbName:        "dh_test",
			DefaultSchema: "public",
		},
	},
}
