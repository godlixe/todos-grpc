package internal

import (
	"context"
	"todos-grpc/internal/entity"
)

type TodoDBItf interface {

	// GetTodoByID gets a single todo by id
	GetTodoByID(ctx context.Context, id int64) (*entity.Todo, error)

	// // GetAllTodos gets multiple todos
	GetAllTodos(ctx context.Context) ([]entity.Todo, error)

	// // CreateTodo creates todo by the given details
	CreateTodo(ctx context.Context, todo *entity.Todo) error

	// // UpdateTodo updates todo by the given details
	UpdateTodo(ctx context.Context, todo *entity.Todo) error
}
