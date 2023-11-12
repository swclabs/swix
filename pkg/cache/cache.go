package cache

import (
	"context"
	"example/komposervice/internal/config"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func Connection() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", config.RedisHost, config.RedisPort),
		Password: config.RedisPassword,
		DB:       0,
	})
}

func Ping() {
	conn := Connection()
	cmd := conn.Ping(context.Background())
	if cmd.Err() != nil {
		print(cmd.Err().Error())
	}
	print("PONG")
}
