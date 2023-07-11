package backend

import "github.com/apibrew/apibrew/pkg/model"

var dhTestMysql = &resource_model.DataSource{
	Backend:     resource_model.DataSourceBackendType_MYSQL,
	Name:        "dh-test-mysql",
	Description: "dh-test-mysql",
	Params: &resource_model.DataSource_MysqlParams{
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

var dhTestRedis = &resource_model.DataSource{
	Backend:     resource_model.DataSourceBackendType_REDIS,
	Name:        "dh-test-redis",
	Description: "dh-test-redis",
	Params: &resource_model.DataSource_RedisParams{
		RedisParams: &model.RedisParams{
			Addr:     "localhost:6379",
			Password: "",
			Db:       0,
		},
	},
}

var dhTestMongo = &resource_model.DataSource{
	Backend:     resource_model.DataSourceBackendType_MONGODB,
	Name:        "dh-test-mongo",
	Description: "dh-test-mongo",
	Params: &resource_model.DataSource_MongoParams{
		MongoParams: &model.MongoParams{
			Uri:    "mongodb://127.0.0.1:27017",
			DbName: "dhTest",
		},
	},
}
