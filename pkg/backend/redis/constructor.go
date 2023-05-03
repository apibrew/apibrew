package redis

import (
	"github.com/apibrew/apibrew/pkg/abs"
	"github.com/apibrew/apibrew/pkg/model"
	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"
)

func NewRedisResourceServiceBackend(dataSource *model.DataSource) abs.Backend {
	redisOptions := dataSource.Params.(*model.DataSource_RedisParams)

	bck := &redisBackend{
		dataSource: dataSource,
		rdb: redis.NewClient(&redis.Options{
			Addr:     redisOptions.RedisParams.Addr,
			Password: redisOptions.RedisParams.Password,
			DB:       int(redisOptions.RedisParams.Db),
		}),
	}

	return bck
}
