package backend

import "github.com/tislib/data-handler/pkg/model"

var dhTestMysql = &model.DataSource{
	Backend:     model.DataSourceBackendType_MYSQL,
	Name:        "dh-test-mysql",
	Description: "dh-test-mysql",
	Options: &model.DataSource_MysqlParams{
		MysqlParams: &model.MysqlOptions{
			Username:      "dh_test",
			Password:      "dh_test",
			Host:          "127.0.0.1",
			Port:          3306,
			DbName:        "dh_test",
			DefaultSchema: "public",
		},
	},
}

var dhTestRedis = &model.DataSource{
	Backend:     model.DataSourceBackendType_REDIS,
	Name:        "dh-test-redis",
	Description: "dh-test-redis",
	Options: &model.DataSource_RedisOptions{
		RedisOptions: &model.RedisOptions{
			Addr:     "localhost:6379",
			Password: "",
			Db:       0,
		},
	},
}
