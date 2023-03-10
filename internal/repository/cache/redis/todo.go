package redis

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"todos-grpc/internal/entity"
	"todos-grpc/pkg/cache"
)

// Todo cache repository
type Todo struct {
	client cache.RedisClientItf
}

func NewTodo(
	client cache.RedisClient,
) *Todo {
	todoRepository := &Todo{
		client: &client,
	}
	return todoRepository
}

func (r *Todo) SetTodo(ctx context.Context, todo *entity.Todo) error {

	res, err := json.Marshal(todo)
	if err != nil {
		return err
	}

	key := fmt.Sprintf("todo:%v", todo.ID)
	err = r.client.Set(ctx, key, string(res))
	if err != nil {
		return err
	}

	return nil
}

func (r *Todo) GetTodo(ctx context.Context, id int64) (*entity.Todo, error) {

	key := fmt.Sprintf("todo:%v", id)
	res, err := r.client.Get(ctx, key)
	if err != nil {
		return nil, err
	}

	if res == nil {
		return nil, errors.New("cache miss")
	}

	var todo entity.Todo
	err = json.Unmarshal(res, &todo)
	if err != nil {
		return nil, nil
	}
	return &todo, nil
}
