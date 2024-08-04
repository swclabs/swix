// Package cache connect to redis
package cache

import (
	"context"
	"fmt"
	"time"

	"go.uber.org/fx"

	"swclabs/swipecore/internal/config"
	"swclabs/swipecore/pkg/lib/logger"

	"github.com/redis/go-redis/v9"
)

// New creates a new redis connection.
func New(lc fx.Lifecycle) ICache {
	conn := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", config.RedisHost, config.RedisPort),
		Password: config.RedisPassword,
		DB:       0,
	})
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			_, err := conn.Ping(ctx).Result()
			if err != nil {
				return err
			}
			logger.Info(fmt.Sprintf("%s to %s", logger.Green.Add("Connect"), logger.Red.Add("Redis")))
			return nil
		},
		OnStop: func(_ context.Context) error {
			logger.Info(fmt.Sprintf("%s %s connection", logger.Green.Add("Closed"), logger.Red.Add("Redis")))
			return conn.Close()
		},
	})
	return &Cache{conn: conn}
}

// ICache interface for cache infrastructure
type ICache interface {
	Set(ctx context.Context, key, val string) error
	Get(ctx context.Context, key string) (string, error)
	Del(ctx context.Context, key string) error
}

var _ ICache = (*Cache)(nil)

// Cache struct for cache
type Cache struct {
	conn *redis.Client
}

// Del implements ICache.
func (c *Cache) Del(ctx context.Context, key string) error {
	return c.conn.Del(ctx, key).Err()
}

// Get implements ICache.
func (c *Cache) Get(ctx context.Context, key string) (string, error) {
	return c.conn.Get(ctx, key).Result()
}

// Set implements ICache.
func (c *Cache) Set(ctx context.Context, key string, val string) error {
	return c.conn.Set(ctx, key, val, time.Duration(time.Minute*10)).Err()
}
