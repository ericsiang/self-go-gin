package redis

import (
	"context"
	"fmt"
	"os"
	"self_go_gin/infra/env"
	"strconv"

	"github.com/redis/go-redis/v9"
)

var redisClient *redis.Client

func InitRedis(GetServerEnv func() *env.ServerConfig) *redis.Client {
	redisConfig := GetServerEnv().Redis
	redisAddr := redisConfig.Host + ":" + strconv.Itoa(redisConfig.Port)
	redisClient = redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: redisConfig.Password,
		DB:       0, // use default DB
	})

	ctx := context.Background()
	if err := redisClient.Ping(ctx).Err(); err != nil {
		fmt.Fprintln(os.Stderr, "redis connect failed, err:", err)
		panic(err)
	}

	fmt.Println("redis client connect ping success")

	return redisClient
}

func GetRedisClient() *redis.Client {
	return redisClient
}
