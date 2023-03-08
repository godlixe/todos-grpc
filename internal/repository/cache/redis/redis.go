package redis

import (
	"todos-grpc/internal"
	"todos-grpc/pkg/cache"
)

type Cache struct {
	Todo   internal.TodoCacheItf
	client cache.RedisClientItf
}

func New() (*Cache, error) {

	redisClient, err := cache.NewRedisClient()
	if err != nil {
		return nil, err
	}

	todoRepository := NewTodo(*redisClient)

	res := &Cache{
		client: redisClient,
		Todo:   todoRepository,
	}

	return res, nil
}
