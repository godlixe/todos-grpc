package service

import (
	"context"
	"fmt"
	"todos-grpc/internal"
	"todos-grpc/internal/entity"
)

type Todo struct {
	cache internal.TodoCacheItf
	db    internal.TodoDBItf
}

func NewTodo(
	db internal.TodoDBItf,
	cache internal.TodoCacheItf,
) *Todo {
	res := &Todo{
		db:    db,
		cache: cache,
	}

	return res
}

func (t *Todo) GetTodoByID(ctx context.Context, id int64) (*entity.Todo, error) {
	todo, err := t.cache.GetTodo(ctx, id)
	if err == nil {
		return todo, nil
	}

	todo, err = t.db.GetTodoByID(ctx, id)
	if err != nil {
		return nil, err
	}

	x := t.cache.SetTodo(ctx, todo)
	if x != nil {
		fmt.Println(x)
	}

	return todo, nil
}

func (t *Todo) GetAllTodos(ctx context.Context) ([]*entity.Todo, error) {
	todos, err := t.db.GetAllTodos(ctx)
	if err != nil {
		return nil, err
	}

	return todos, nil
}

// CreateTodo creates todo by the given details
func (t *Todo) CreateTodo(ctx context.Context, todo *entity.Todo) (*entity.Todo, error) {
	res, err := t.db.CreateTodo(ctx, todo)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// UpdateTodo updates todo by the given details
func (t *Todo) UpdateTodo(ctx context.Context, todo *entity.Todo) (*entity.Todo, error) {
	res, err := t.db.UpdateTodo(ctx, todo)
	if err != nil {
		return nil, err
	}

	return res, nil
}
