package cache

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

type RedisClient struct {
	client redis.Client
}

func NewRedisClient() (*RedisClient, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "root",
		DB:       0,
	})

	redisClient := RedisClient{
		client: *client,
	}

	return &redisClient, nil
}

func (r *RedisClient) Get(ctx context.Context, key string) ([]byte, error) {
	res, err := r.client.Get(key).Result()
	if err != nil {
		return nil, nil
	}
	fmt.Println("getting from cache")
	return []byte(res), nil
}

func (r *RedisClient) Set(ctx context.Context, key string, value string) error {
	_, err := r.client.Set(key, value, 2*time.Hour).Result()
	if err != nil {
		return err
	}
	fmt.Println("setting cache", key, value)
	return nil
}
