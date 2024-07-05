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

func CreateRedisConnection(lc fx.Lifecycle, env config.Env) *redis.Client {
	conn := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", env.RedisHost, env.RedisPort),
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

func Ping() {
	conn := Connection()
	cmd := conn.Ping(context.Background())
	if cmd.Err() != nil {
		print(cmd.Err().Error())
	}
	print("PONG")
}
