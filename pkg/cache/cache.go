package cache

import "context"

type RedisClientItf interface {
	Get(ctx context.Context, key string) ([]byte, error)
	Set(ctx context.Context, key string, value string) error
}
