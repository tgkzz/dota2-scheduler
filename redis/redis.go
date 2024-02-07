package redis

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

func NewRedisClient() (*redis.Client, error) {
	var ctx = context.Background()

	client := redis.NewClient(
		&redis.Options{
			Username:     "kamal",
			Password:     "Adizero50!",
			Addr:         "redis-15640.c55.eu-central-1-1.ec2.cloud.redislabs.com:15640",
			ReadTimeout:  30 * time.Second,
			WriteTimeout: 30 * time.Second,
		})

	res, err := client.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}

	fmt.Println(res)

	return client, err
}
