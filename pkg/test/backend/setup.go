package backend

import "github.com/tislib/data-handler/pkg/model"

var dhTestMysql = &model.DataSource{
	Backend:     model.DataSourceBackendType_MYSQL,
	Name:        "dh-test-mysql",
	Description: "dh-test-mysql",
	Params: &model.DataSource_MysqlParams{
		MysqlParams: &model.MysqlParams{
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
	Params: &model.DataSource_RedisParams{
		RedisParams: &model.RedisParams{
			Addr:     "localhost:6379",
			Password: "",
			Db:       0,
		},
	},
}

var dhTestMongo = &model.DataSource{
	Backend:     model.DataSourceBackendType_MONGODB,
	Name:        "dh-test-mongo",
	Description: "dh-test-mongo",
	Params: &model.DataSource_MongoParams{
		MongoParams: &model.MongoParams{
			Uri:    "mongodb://127.0.0.1:27017",
			DbName: "dhTest",
		},
	},
}

var dhTestSqlite = &model.DataSource{
	Backend:     model.DataSourceBackendType_SQLITE,
	Name:        "dh-test-sqlite",
	Description: "dh-test-sqlite",
	Params: &model.DataSource_SqliteParams{
		SqliteParams: &model.SqliteParams{
			Path: "./test.db",
		},
	},
}
