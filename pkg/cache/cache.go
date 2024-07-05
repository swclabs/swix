package cache

import (
	"context"
	"fmt"

	"go.uber.org/fx"

	"swclabs/swipecore/internal/config"

	"github.com/redis/go-redis/v9"
)

func Connection() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", config.RedisHost, config.RedisPort),
		Password: config.RedisPassword,
		DB:       0,
	})
}

func CreateRedisConnection(lc fx.Lifecycle) *redis.Client {
	conn := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", config.RedisHost, config.RedisPort),
		Password: config.RedisPassword,
		DB:       0,
	})
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return conn.Close()
		},
	})
	return conn
}
