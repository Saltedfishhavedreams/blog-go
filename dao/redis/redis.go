package redis

import (
	"blog/config"
	"blog/pkg/logger"
	"fmt"
	"time"

	"github.com/go-redis/redis"
	"go.uber.org/zap"
)

var RDB *redis.Client

const DataDefaultExpire = time.Second * 3600 * 24 * 14

func Init(redisConfig *config.RedisConfig) (err error) {
	RDB = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", redisConfig.Host, redisConfig.Port),
		Password: redisConfig.Password,
		DB:       redisConfig.Db,
		PoolSize: int(redisConfig.PoolSize),
	})

	_, err = RDB.Ping().Result()
	if err != nil {
		logger.Error("redis client connect error", zap.Error(err))
	}

	return
}

func Close() {
	RDB.Close()
}
