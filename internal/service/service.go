package service

import (
	"todos-grpc/internal"
	"todos-grpc/internal/repository/cache/redis"
	"todos-grpc/internal/repository/db/postgre"
)

type Domain struct {
	Todo  internal.TodoSvcItf
	Cache internal.TodoCacheItf
}

func New(
	postgreDB *postgre.DB,
	redisCache *redis.Cache,
) *Domain {
	todoSvc := NewTodo(
		postgreDB.Todo,
		redisCache.Todo,
	)

	res := &Domain{
		Todo: todoSvc,
	}

	return res
}
