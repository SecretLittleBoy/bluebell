package redis

import (
	"fmt"
	"bluebell/settings"

	"github.com/go-redis/redis"
)

var rdb *redis.Client

func Init()(err error){
	rdb = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", settings.Config.RedisConfig.Host, settings.Config.RedisConfig.Port),
		Password: settings.Config.RedisConfig.Password,
		DB: settings.Config.RedisConfig.DB,
		PoolSize: settings.Config.RedisConfig.PoolSize,
	})
	_, err = rdb.Ping().Result()
	return
}

func Close() {
	_ = rdb.Close()
}