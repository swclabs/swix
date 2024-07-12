// Package cache connect to redis
package cache

import (
	"context"
	"fmt"

	"go.uber.org/fx"

	"swclabs/swipecore/internal/config"

	"github.com/redis/go-redis/v9"
)

func New(lc fx.Lifecycle) *redis.Client {
	conn := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", config.RedisHost, config.RedisPort),
		Password: config.RedisPassword,
		DB:       0,
	})
	lc.Append(fx.Hook{
		OnStart: func(_ context.Context) error {
			return nil
		},
		OnStop: func(_ context.Context) error {
			return conn.Close()
		},
	})
	return conn
}
