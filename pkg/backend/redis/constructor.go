package redis

import (
	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"
	"github.com/tislib/data-handler/pkg/abs"
	"github.com/tislib/data-handler/pkg/model"
)

func NewRedisResourceServiceBackend(dataSource *model.DataSource) abs.Backend {
	redisOptions := dataSource.Options.(*model.DataSource_RedisOptions)

	bck := redisBackend{
		dataSource: dataSource,
		rdb: redis.NewClient(&redis.Options{
			Addr:     redisOptions.RedisOptions.Addr,
			Password: redisOptions.RedisOptions.Password,
			DB:       int(redisOptions.RedisOptions.Db),
		}),
	}

	return bck
}
