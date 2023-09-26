package redis

import (
	"blog/config"
	"blog/pkg/logger"
	"fmt"

	"github.com/go-redis/redis"
)

var client *redis.Client

func Init(redisConfig config.RedisConfig) (err error) {
	client = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", redisConfig.Host, redisConfig.Port),
		Password: redisConfig.Password,
		DB:       redisConfig.Db,
		PoolSize: int(redisConfig.PoolSize),
	})

	_, err = client.Ping().Result()
	if err != nil {
		logger.Error("redis client connect error:%s", err)
	}

	return
}

func Close() {
	client.Close()
}
