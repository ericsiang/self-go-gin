package initialize

import (
	"context"
	"strconv"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

var redisClient *redis.Client

func initRedis() *redis.Client {
	redisConfig := GetServerEnv().Redis
	redisAddr := redisConfig.Host + ":" + strconv.Itoa(redisConfig.Port)
	redisClient = redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: redisConfig.Password,
		DB:       0, // use default DB
	})

	ctx := context.Background()
	if err := redisClient.Ping(ctx).Err(); err != nil {
		zap.S().Error("redis connect failed, err:%v\n", err.Error())
		panic(err)
	}

	zap.S().Info("redis client connect ping success")

	return redisClient
}

func GetRedisClient() *redis.Client {
	return redisClient
}
