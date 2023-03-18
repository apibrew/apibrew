package backend

import "github.com/tislib/data-handler/pkg/model"

var dhTestMysql = &model.DataSource{
	Backend:     model.DataSourceBackendType_MYSQL,
	Name:        "dh-test-mysql",
	Description: "dh-test-mysql",
	Options: &model.DataSource_MysqlParams{
		MysqlParams: &model.MysqlOptions{
			Username:      "root",
			Password:      "",
			Host:          "127.0.0.1",
			Port:          3306,
			DbName:        "dh_test",
			DefaultSchema: "public",
		},
	},
}
