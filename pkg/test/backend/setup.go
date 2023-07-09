package backend

import "github.com/apibrew/apibrew/pkg/model"

var dhTestMysql = &modelnew.DataSource{
	Backend:     modelnew.DataSourceBackendType_MYSQL,
	Name:        "dh-test-mysql",
	Description: "dh-test-mysql",
	Params: &modelnew.DataSource_MysqlParams{
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

var dhTestRedis = &modelnew.DataSource{
	Backend:     modelnew.DataSourceBackendType_REDIS,
	Name:        "dh-test-redis",
	Description: "dh-test-redis",
	Params: &modelnew.DataSource_RedisParams{
		RedisParams: &model.RedisParams{
			Addr:     "localhost:6379",
			Password: "",
			Db:       0,
		},
	},
}

var dhTestMongo = &modelnew.DataSource{
	Backend:     modelnew.DataSourceBackendType_MONGODB,
	Name:        "dh-test-mongo",
	Description: "dh-test-mongo",
	Params: &modelnew.DataSource_MongoParams{
		MongoParams: &model.MongoParams{
			Uri:    "mongodb://127.0.0.1:27017",
			DbName: "dhTest",
		},
	},
}
