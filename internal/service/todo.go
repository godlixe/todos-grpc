package service

import (
	"context"
	"todos-grpc/internal"
	"todos-grpc/internal/entity"
)

type Todo struct {
	db internal.TodoDBItf
}

func NewTodo(
	db internal.TodoDBItf,
) *Todo {
	res := &Todo{
		db: db,
	}

	return res
}

func (t *Todo) GetTodoByID(ctx context.Context, id int64) (*entity.Todo, error) {
	todo, err := t.db.GetTodoByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return todo, nil
}

func (t *Todo) GetAllTodos(ctx context.Context) ([]entity.Todo, error) {
	return nil, nil
}

// // CreateTodo creates todo by the given details
func (t *Todo) CreateTodo(ctx context.Context, todo *entity.Todo) error {
	return nil
}

// // UpdateTodo updates todo by the given details
func (t *Todo) UpdateTodo(ctx context.Context, todo *entity.Todo) error {
	return nil
}
