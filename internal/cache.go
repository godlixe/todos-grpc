package internal

import (
	"context"
	"todos-grpc/internal/entity"
)

type TodoCacheItf interface {
	SetTodo(ctx context.Context, todo *entity.Todo) error
	GetTodo(ctx context.Context, id int64) (*entity.Todo, error)
}
