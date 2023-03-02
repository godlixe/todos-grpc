package handler

import (
	"context"
	"todos-grpc/internal"
	"todos-grpc/pkg/proto"
)

type Todo struct {
	todoService internal.TodoSvcItf
	proto.UnimplementedTodosServer
}

func NewTodo(
	todoService internal.TodoDBItf,
) *Todo {
	res := &Todo{
		todoService: todoService,
	}

	return res
}

func (t *Todo) GetTodo(ctx context.Context, req *proto.GetTodoRequest) (*proto.GetTodoResponse, error) {
	res, err := t.todoService.GetTodoByID(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &proto.GetTodoResponse{
		Todo: &proto.Todo{
			Id:          res.ID,
			Title:       res.Title,
			Description: res.Description,
			IsDone:      res.IsDone,
			UserId:      res.UserID,
			TimeStamp:   res.TimeStamp.String(),
		},
	}, nil
}

func (t *Todo) GetAllTodos(ctx context.Context, req *proto.GetTodosRequest) (*proto.GetTodosResponse, error) {
	return nil, nil
}
func (t *Todo) CreateTodo(ctx context.Context, req *proto.CreateTodoRequest) (*proto.GetTodoResponse, error) {
	return nil, nil
}

func (t *Todo) UpdateTodo(ctx context.Context, res *proto.UpdateTodoResponse) (*proto.GetTodoResponse, error) {
	return nil, nil
}
