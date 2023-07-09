package redis

import (
	"github.com/apibrew/apibrew/pkg/abs"
	"github.com/apibrew/apibrew/pkg/modelnew"
	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"
	"strconv"
)

func NewRedisResourceServiceBackend(dataSource *modelnew.DataSource) abs.Backend {

	db, _ := strconv.Atoi(dataSource.Options["db"])

	bck := &redisBackend{
		dataSource: dataSource,
		rdb: redis.NewClient(&redis.Options{
			Addr:     dataSource.Options["addr"],
			Password: dataSource.Options["password"],
			DB:       db,
		}),
	}

	return bck
}
